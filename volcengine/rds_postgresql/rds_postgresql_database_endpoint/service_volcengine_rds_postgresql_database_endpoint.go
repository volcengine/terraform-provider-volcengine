package rds_postgresql_database_endpoint

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
	rdsPgInstance "github.com/volcengine/terraform-provider-volcengine/volcengine/rds_postgresql/rds_postgresql_instance"
)

type VolcengineRdsPostgresqlDatabaseEndpointService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewRdsPostgresqlDatabaseEndpointService(c *ve.SdkClient) *VolcengineRdsPostgresqlDatabaseEndpointService {
	return &VolcengineRdsPostgresqlDatabaseEndpointService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineRdsPostgresqlDatabaseEndpointService) GetClient() *ve.SdkClient {
	return s.Client
}

// 考虑只保留 resource 相关的方法，ReadResources 和 ReadResource 直接返回 nil, nil
func (s *VolcengineRdsPostgresqlDatabaseEndpointService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithSimpleQuery(m, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeDBInstanceDetail"
		bytes, _ := json.Marshal(condition)
		logger.Debug(logger.ReqFormat, action, string(bytes))
		if condition == nil {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), nil)
			if err != nil {
				return data, err
			}
		} else {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &condition)
			if err != nil {
				return data, err
			}
		}
		respBytes, _ := json.Marshal(resp)
		logger.Debug(logger.RespFormat, action, condition, string(respBytes))
		results, err = ve.ObtainSdkValue("Result.Endpoints", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.Endpoints is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineRdsPostgresqlDatabaseEndpointService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results      []interface{}
		ok           bool
		endpointData map[string]interface{}
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	parts := strings.Split(id, ":")
	if len(parts) < 2 {
		return data, fmt.Errorf("invalid id %s", id)
	}
	instanceId := parts[0]
	key := parts[1]
	req := map[string]interface{}{
		"InstanceId": instanceId,
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		var item map[string]interface{}
		if item, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("Value is not map ")
		}
		if eid, ok1 := item["EndpointId"]; ok1 && fmt.Sprint(eid) == key {
			endpointData = item
			break
		}
		if ename, ok2 := item["EndpointName"]; ok2 && fmt.Sprint(ename) == key {
			endpointData = item
			break
		}
	}
	if len(endpointData) == 0 {
		return data, fmt.Errorf("rds_postgresql_database_endpoint %s not exist ", id)
	}
	addresses := endpointData["Addresses"]
	if addresses != nil {
		for _, addr := range addresses.([]interface{}) {
			am := addr.(map[string]interface{})
			if fmt.Sprint(am["NetworkType"]) == "Private" {
				endpointData["Domain"] = am["Domain"]
				endpointData["DNSVisibility"] = am["DNSVisibility"]
				endpointData["Port"] = am["Port"]
				break
			}
		}
	}
	data = endpointData
	return data, err
}

func (s *VolcengineRdsPostgresqlDatabaseEndpointService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      5 * time.Second,
		MinTimeout: 5 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				d          map[string]interface{}
				status     interface{}
				failStates []string
			)
			failStates = append(failStates, "Failed")
			d, err = s.ReadResource(resourceData, id)
			if err != nil {
				return nil, "", err
			}
			status, err = ve.ObtainSdkValue("Status", d)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("rds_postgresql_database_endpoint status error, status: %s", status.(string))
				}
			}
			return d, status.(string), err
		},
	}
}

func (s *VolcengineRdsPostgresqlDatabaseEndpointService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	instanceId := resourceData.Get("instance_id").(string)
	endpointName := resourceData.Get("endpoint_name").(string)
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateDBEndpoint",
			ConvertMode: ve.RequestConvertInConvert,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"instance_id":   {TargetField: "InstanceId"},
				"endpoint_name": {TargetField: "EndpointName"},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				// endpoint_type 目前只支持取值 Custom
				if v, ok := d.GetOk("endpoint_type"); ok {
					if v.(string) != "Custom" {
						return false, fmt.Errorf("endpoint_type only support Custom")
					}
				}
				(*call.SdkParam)["EndpointType"] = "Custom"
				if nodes, ok := d.GetOk("nodes"); ok {
					(*call.SdkParam)["Nodes"] = nodes.(string)
				} else {
					return false, fmt.Errorf("nodes is required when end_point_type is Custom")
				}
				if v, ok := d.GetOk("read_write_mode"); ok {
					(*call.SdkParam)["ReadWriteMode"] = v.(string)
				}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				rdsPgInstance.NewRdsPostgresqlInstanceService(s.Client): {
					Target:     []string{"Running"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					ResourceId: instanceId,
				},
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				endpointId := ""
				current, _ := s.ReadResource(resourceData, fmt.Sprintf("%s:%s", instanceId, endpointName))
				if v, ok := current["EndpointId"].(string); ok {
					endpointId = v
				}
				d.SetId(fmt.Sprintf("%s:%s", instanceId, endpointId))
				return nil
			},
			LockId: func(d *schema.ResourceData) string { return instanceId },
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineRdsPostgresqlDatabaseEndpointService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"EndpointId":                   {TargetField: "endpoint_id"},
			"EndpointName":                 {TargetField: "endpoint_name"},
			"EndpointType":                 {TargetField: "endpoint_type"},
			"ReadWriteMode":                {TargetField: "read_write_mode"},
			"ReadWriteSpliting":            {TargetField: "read_write_splitting"},
			"ReadOnlyNodeMaxDelayTime":     {TargetField: "read_only_node_max_delay_time"},
			"ReadOnlyNodeDistributionType": {TargetField: "read_only_node_distribution_type"},
			"ReadOnlyNodeWeight":           {TargetField: "read_only_node_weight"},
			"WriteNodeHaltWriting":         {TargetField: "write_node_halt_writing"},
			"ReadWriteProxyConnection":     {TargetField: "read_write_proxy_connection"},
			"Domain":                       {TargetField: "domain"},
			"Port":                         {TargetField: "port"},
			"DNSVisibility":                {TargetField: "dns_visibility"},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineRdsPostgresqlDatabaseEndpointService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback
	parts := strings.Split(resourceData.Id(), ":")
	instanceId := parts[0]
	endpointId := ""
	current, _ := s.ReadResource(resourceData, resourceData.Id())
	if v, ok := current["EndpointId"].(string); ok {
		fmt.Printf("endpointId: %s\n", v)
		endpointId = v
	}
	var defaultEndpoint bool
	if endpointId == fmt.Sprintf("%s-cluster", instanceId) {
		defaultEndpoint = true
	} else {
		defaultEndpoint = false
	}

	var desiredSplit bool
	if v, ok := resourceData.GetOk("read_write_splitting"); ok {
		desiredSplit = v.(bool)
	} else {
		if cv, ok := current["EnableReadWriteSplitting"].(string); ok {
			if cv == "Enable" {
				desiredSplit = true
			} else {
				desiredSplit = false
			}
		}
	}
	hasSplitChange := resourceData.HasChange("read_write_splitting")
	hasWeightChange := resourceData.HasChanges("read_only_node_distribution_type", "read_only_node_weight", "write_node_halt_writing")
	hasDelayChange := resourceData.HasChange("read_only_node_max_delay_time")
	hasProxyChange := resourceData.HasChange("read_write_proxy_connection")
	if !defaultEndpoint && (hasSplitChange || hasWeightChange || hasDelayChange || hasProxyChange) {
		cb := ve.Callback{
			Call: ve.SdkCall{
				ConvertMode: ve.RequestConvertIgnore,
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					return false, fmt.Errorf("only default endpoint supports read-write splitting features, custom endpoint not support")
				},
			},
		}
		return append(callbacks, cb)
	}

	// 默认终端和普通终端均支持
	if resourceData.HasChange("endpoint_name") {
		modifyDBEndpointNameCallBack := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyDBEndpointName",
				ConvertMode: ve.RequestConvertIgnore,
				ContentType: ve.ContentTypeJson,
				SdkParam:    &map[string]interface{}{},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["InstanceId"] = instanceId
					(*call.SdkParam)["EndpointId"] = endpointId
					(*call.SdkParam)["EndpointName"] = d.Get("endpoint_name")
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
				ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
					rdsPgInstance.NewRdsPostgresqlInstanceService(s.Client): {
						Target:     []string{"Running"},
						Timeout:    resourceData.Timeout(schema.TimeoutUpdate),
						ResourceId: instanceId},
				},
				LockId: func(d *schema.ResourceData) string { return instanceId },
			},
		}
		callbacks = append(callbacks, modifyDBEndpointNameCallBack)
	}

	if resourceData.HasChange("global_read_only") {
		modifyDBInstanceConfigCallBack := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyDBInstanceConfig",
				ConvertMode: ve.RequestConvertIgnore,
				ContentType: ve.ContentTypeJson,
				SdkParam: &map[string]interface{}{
					"InstanceId":     instanceId,
					"GlobalReadOnly": resourceData.Get("global_read_only"),
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
				ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
					rdsPgInstance.NewRdsPostgresqlInstanceService(s.Client): {
						Target:     []string{"Running"},
						Timeout:    resourceData.Timeout(schema.TimeoutUpdate),
						ResourceId: instanceId,
					},
				},
				LockId: func(d *schema.ResourceData) string { return instanceId },
			},
		}
		callbacks = append(callbacks, modifyDBInstanceConfigCallBack)
	}

	if resourceData.HasChange("dns_visibility") {
		modifyDBEndpointDNSCallBack := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyDBEndpointDNS",
				ConvertMode: ve.RequestConvertIgnore,
				ContentType: ve.ContentTypeJson,
				SdkParam:    &map[string]interface{}{},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["InstanceId"] = instanceId
					(*call.SdkParam)["EndpointId"] = endpointId
					// NetworkType 要求必填，目前仅支持 Private。
					(*call.SdkParam)["NetworkType"] = "Private"
					if v, ok := d.GetOk("dns_visibility"); ok {
						(*call.SdkParam)["DNSVisibility"] = v.(bool)
					}
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
				ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
					rdsPgInstance.NewRdsPostgresqlInstanceService(s.Client): {
						Target:     []string{"Running"},
						Timeout:    resourceData.Timeout(schema.TimeoutUpdate),
						ResourceId: instanceId,
					},
				},
				LockId: func(d *schema.ResourceData) string { return instanceId },
			},
		}
		callbacks = append(callbacks, modifyDBEndpointDNSCallBack)
	}

	// 在一次调用中不能同时修改私网连接地址的前缀和端口，只能对其中一项进行修改。
	// 不能同时不为请求参数 DomainPrefix 和 Port 传值，也不可同时为两者传值。
	// domain_prefix 和 port 一次只能传一个值，如果用户都修改了，分两次传递
	if resourceData.HasChange("domain_prefix") {
		modifyDBEndpointAddressCallback1 := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyDBEndpointAddress",
				ConvertMode: ve.RequestConvertIgnore,
				ContentType: ve.ContentTypeJson,
				SdkParam:    &map[string]interface{}{},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["InstanceId"] = instanceId
					(*call.SdkParam)["EndpointId"] = endpointId
					// NetworkType 要求必填，目前仅支持 Private。
					(*call.SdkParam)["NetworkType"] = "Private"
					if v, ok := d.GetOk("domain_prefix"); ok {
						(*call.SdkParam)["DomainPrefix"] = v.(string)
					}
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
				ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
					rdsPgInstance.NewRdsPostgresqlInstanceService(s.Client): {
						Target:     []string{"Running"},
						Timeout:    resourceData.Timeout(schema.TimeoutUpdate),
						ResourceId: instanceId,
					},
				},
				LockId: func(d *schema.ResourceData) string { return instanceId },
			},
		}
		callbacks = append(callbacks, modifyDBEndpointAddressCallback1)
	}
	if resourceData.HasChange("port") {
		modifyDBEndpointAddressCallback2 := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyDBEndpointAddress",
				ConvertMode: ve.RequestConvertIgnore,
				ContentType: ve.ContentTypeJson,
				SdkParam:    &map[string]interface{}{},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["InstanceId"] = instanceId
					(*call.SdkParam)["EndpointId"] = endpointId
					(*call.SdkParam)["NetworkType"] = "Private"
					if v, ok := d.GetOk("port"); ok {
						(*call.SdkParam)["Port"] = v.(string)
					}
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
				ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
					rdsPgInstance.NewRdsPostgresqlInstanceService(s.Client): {
						Target:     []string{"Running"},
						Timeout:    resourceData.Timeout(schema.TimeoutUpdate),
						ResourceId: instanceId,
					},
				},
				LockId: func(d *schema.ResourceData) string { return instanceId },
			},
		}
		callbacks = append(callbacks, modifyDBEndpointAddressCallback2)
	}

	// 以下操作仅 默认终端 支持，先开启读写分离，再修改参数
	if hasSplitChange {
		modifyDBEndpointReadWriteFlagCallBack := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyDBEndpointReadWriteFlag",
				ConvertMode: ve.RequestConvertIgnore,
				ContentType: ve.ContentTypeJson,
				SdkParam:    &map[string]interface{}{},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["InstanceId"] = instanceId
					(*call.SdkParam)["EndpointId"] = endpointId
					(*call.SdkParam)["ReadWriteSpliting"] = desiredSplit
					if desiredSplit {
						if v, ok := d.GetOk("read_only_node_max_delay_time"); ok {
							(*call.SdkParam)["ReadOnlyNodeMaxDelayTime"] = v
						}
						if v, ok := d.GetOk("read_only_node_distribution_type"); ok {
							(*call.SdkParam)["ReadOnlyNodeDistributionType"] = v
							if v == "Custom" {
								if w, ok := d.GetOk("read_only_node_weight"); ok {
									weights := make([]map[string]interface{}, 0)
									for _, it := range w.([]interface{}) {
										m := it.(map[string]interface{})
										nt, _ := m["node_type"].(string)
										if nt == "Primary" {
											item := map[string]interface{}{
												"NodeType": "Primary",
											}
											if wt, ok2 := m["weight"]; ok2 {
												item["Weight"] = wt
											}
											weights = append(weights, item)
											continue
										}
										if nt == "ReadOnly" {
											nidStr, _ := m["node_id"].(string)
											if nidStr == "" {
												continue
											}
											item := map[string]interface{}{
												"NodeId":   nidStr,
												"NodeType": "ReadOnly",
											}
											if wt, ok2 := m["weight"]; ok2 {
												item["Weight"] = wt
											}
											weights = append(weights, item)
										}
									}
									(*call.SdkParam)["ReadOnlyNodeWeight"] = weights
								} else {
									return false, fmt.Errorf("read_only_node_weight is required when read_only_node_distribution_type is Custom")
								}
							}
						}
						if v, ok := d.GetOk("read_write_proxy_connection"); ok {
							(*call.SdkParam)["ReadWriteProxyConnection"] = v
						}
						if v, ok := d.GetOk("write_node_halt_writing"); ok {
							(*call.SdkParam)["WriteNodeHaltWriting"] = v
						}
					}
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
				ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
					rdsPgInstance.NewRdsPostgresqlInstanceService(s.Client): {
						Target:     []string{"Running"},
						Timeout:    resourceData.Timeout(schema.TimeoutUpdate),
						ResourceId: instanceId},
				},
				LockId: func(d *schema.ResourceData) string { return instanceId },
			},
		}
		callbacks = append(callbacks, modifyDBEndpointReadWriteFlagCallBack)
	}

	if !hasSplitChange && desiredSplit && hasWeightChange {
		// 仅修改了已开启读写分离的默认终端的只读节点权重
		modifyDBEndpointReadWeightCallBack := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyDBEndpointReadWeight",
				ConvertMode: ve.RequestConvertIgnore,
				ContentType: ve.ContentTypeJson,
				SdkParam:    &map[string]interface{}{},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["InstanceId"] = instanceId
					(*call.SdkParam)["EndpointId"] = endpointId
					if v, ok := d.GetOk("write_node_halt_writing"); ok {
						(*call.SdkParam)["WriteNodeHaltWriting"] = v
					}
					if v, ok := d.GetOk("read_only_node_distribution_type"); ok {
						(*call.SdkParam)["ReadOnlyNodeDistributionType"] = v
						if v == "Custom" {
							if w, ok := d.GetOk("read_only_node_weight"); ok {
								weights := make([]map[string]interface{}, 0)
								for _, it := range w.([]interface{}) {
									m := it.(map[string]interface{})
									nt, _ := m["node_type"].(string)
									if nt == "Primary" {
										item := map[string]interface{}{
											"NodeType": "Primary",
										}
										if wt, ok2 := m["weight"]; ok2 {
											item["Weight"] = wt
										}
										weights = append(weights, item)
										continue
									}
									if nt == "ReadOnly" {
										nidStr, _ := m["node_id"].(string)
										if nidStr == "" {
											continue
										}
										item := map[string]interface{}{
											"NodeId":   nidStr,
											"NodeType": "ReadOnly",
										}
										if wt, ok2 := m["weight"]; ok2 {
											item["Weight"] = wt
										}
										weights = append(weights, item)
									}
								}
								(*call.SdkParam)["ReadOnlyNodeWeight"] = weights
							} else {
								return false, fmt.Errorf("read_only_node_weight is required when read_only_node_distribution_type is Custom")
							}
						}
					}
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
				ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
					rdsPgInstance.NewRdsPostgresqlInstanceService(s.Client): {
						Target:     []string{"Running"},
						Timeout:    resourceData.Timeout(schema.TimeoutUpdate),
						ResourceId: instanceId,
					},
				},
				LockId: func(d *schema.ResourceData) string { return instanceId },
			},
		}
		callbacks = append(callbacks, modifyDBEndpointReadWeightCallBack)
	}

	if !hasSplitChange && desiredSplit && hasProxyChange {
		// 仅修改了已开启读写分离的默认终端的读写代理连接配置
		modifyDBEndpointProxyConfigCallBack := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyDBEndpointProxyConfig",
				ConvertMode: ve.RequestConvertIgnore,
				ContentType: ve.ContentTypeJson,
				SdkParam:    &map[string]interface{}{},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["InstanceId"] = instanceId
					(*call.SdkParam)["EndpointId"] = endpointId
					(*call.SdkParam)["ReadWriteProxyConnection"] = d.Get("read_write_proxy_connection")
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
				ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
					rdsPgInstance.NewRdsPostgresqlInstanceService(s.Client): {
						Target:     []string{"Running"},
						Timeout:    resourceData.Timeout(schema.TimeoutUpdate),
						ResourceId: instanceId,
					},
				},
				LockId: func(d *schema.ResourceData) string { return instanceId },
			},
		}
		callbacks = append(callbacks, modifyDBEndpointProxyConfigCallBack)
	}

	if !hasSplitChange && desiredSplit && hasDelayChange {
		// 仅修改了已开启读写分离的默认终端的只读节点延迟阈值
		modifyDBEndpointReadWriteDelayThresholdCallBack := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyDBEndpointReadWriteDelayThreshold",
				ConvertMode: ve.RequestConvertIgnore,
				ContentType: ve.ContentTypeJson,
				SdkParam:    &map[string]interface{}{},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["InstanceId"] = instanceId
					(*call.SdkParam)["EndpointId"] = endpointId
					(*call.SdkParam)["ReadOnlyNodeMaxDelayTime"] = d.Get("read_only_node_max_delay_time")
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
				ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
					rdsPgInstance.NewRdsPostgresqlInstanceService(s.Client): {
						Target:     []string{"Running"},
						Timeout:    resourceData.Timeout(schema.TimeoutUpdate),
						ResourceId: instanceId,
					},
				},
				LockId: func(d *schema.ResourceData) string { return instanceId },
			},
		}
		callbacks = append(callbacks, modifyDBEndpointReadWriteDelayThresholdCallBack)
	}

	return callbacks
}

func (s *VolcengineRdsPostgresqlDatabaseEndpointService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	parts := strings.Split(resourceData.Id(), ":")
	instanceId := parts[0]
	endpointId := ""
	current, _ := s.ReadResource(resourceData, resourceData.Id())
	if v, ok := current["EndpointId"].(string); ok {
		endpointId = v
	}
	if endpointId == fmt.Sprintf("%s-cluster", instanceId) {
		// 默认终端无法删除，仅在Terraform中移除管理
		logger.Info(logger.RespFormat, "rds_postgresql_database_endpoint %s is default endpoint, can not be deleted, only remove from terraform state", endpointId)
		return []ve.Callback{}
	}
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteDBEndpoint",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"InstanceId": instanceId,
				"EndpointId": endpointId,
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 5*time.Minute)
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				rdsPgInstance.NewRdsPostgresqlInstanceService(s.Client): {
					Target:     []string{"Running"},
					Timeout:    resourceData.Timeout(schema.TimeoutDelete),
					ResourceId: instanceId,
				},
			},
			LockId: func(d *schema.ResourceData) string { return instanceId },
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineRdsPostgresqlDatabaseEndpointService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"instance_id": {TargetField: "InstanceId"},
		},
		NameField:    "EndpointName",
		IdField:      "EndpointId",
		CollectField: "endpoints",
		ContentType:  ve.ContentTypeJson,
		ResponseConverts: map[string]ve.ResponseConvert{
			"EndpointId":    {TargetField: "id", KeepDefault: true},
			"EndpointName":  {TargetField: "endpoint_name"},
			"ReadWriteMode": {TargetField: "read_write_mode"},
			"DNSVisibility": {TargetField: "dns_visibility"},
		},
	}
}

func (s *VolcengineRdsPostgresqlDatabaseEndpointService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "rds_postgresql",
		Version:     "2022-01-01",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}
