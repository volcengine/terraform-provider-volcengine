package nlb_server_group

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineNlbServerGroupService struct {
	Client *ve.SdkClient
}

func NewNlbServerGroupService(c *ve.SdkClient) *VolcengineNlbServerGroupService {
	return &VolcengineNlbServerGroupService{
		Client: c,
	}
}

func (s *VolcengineNlbServerGroupService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineNlbServerGroupService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	return ve.WithNextTokenQuery(m, "MaxResults", "NextToken", 100, nil, func(condition map[string]interface{}) ([]interface{}, string, error) {
		action := "DescribeNLBServerGroups"
		bytes, _ := json.Marshal(condition)
		logger.Debug(logger.ReqFormat, action, string(bytes))
		resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(action), &condition)
		if err != nil {
			return nil, "", err
		}
		respBytes, _ := json.Marshal(resp)
		logger.Debug(logger.RespFormat, action, condition, string(respBytes))
		results, err := ve.ObtainSdkValue("Result.ServerGroups", *resp)
		if err != nil {
			return nil, "", err
		}
		if results == nil {
			results = []interface{}{}
		}

		var nextTokenStr string
		nextToken, _ := ve.ObtainSdkValue("Result.NextToken", *resp)
		if nextToken != nil {
			if s, ok := nextToken.(string); ok {
				nextTokenStr = s
			}
		}

		if data, ok := results.([]interface{}); ok {
			return data, nextTokenStr, nil
		}
		return nil, "", errors.New("Result.ServerGroups is not Slice")
	})
}

func (s *VolcengineNlbServerGroupService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	if id == "" {
		id = resourceData.Id()
	}
	req := map[string]interface{}{
		"ServerGroupIds.1": id,
	}
	results, err := s.ReadResources(req)
	if err != nil {
		return nil, err
	}
	for _, v := range results {
		if tempData, ok := v.(map[string]interface{}); ok {
			if sgId, ok := tempData["ServerGroupId"].(string); ok && sgId == id {
				return tempData, nil
			}
		}
	}
	return nil, fmt.Errorf("resource not found")
}

func (s *VolcengineNlbServerGroupService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{"Creating", "Configuring"},
		Target:     target,
		Refresh:    ve.ResourceStateRefreshFunc(resourceData, s.ReadResource, id),
		Timeout:    timeout,
		Delay:      5 * time.Second,
		MinTimeout: 2 * time.Second,
	}
}

func (s *VolcengineNlbServerGroupService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action: "CreateNLBServerGroup",
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				param, err := ve.ResourceDateToRequest(d, resource, false, s.createRequestConvert(), ve.RequestConvertAll, ve.ContentTypeDefault)
				if err != nil {
					return nil, err
				}

				// 手动添加 health_check 参数
				if hcList, ok := d.GetOk("health_check"); ok {
					if hcList, ok := hcList.([]interface{}); ok && len(hcList) > 0 {
						if hc, ok := hcList[0].(map[string]interface{}); ok {
							if v, ok := hc["enabled"].(bool); ok {
								param["HealthCheck.Enabled"] = v
							}
							if v, ok := hc["type"].(string); ok && v != "" {
								param["HealthCheck.Type"] = v
							}
							if v, ok := hc["port"].(int); ok {
								param["HealthCheck.Port"] = v
							}
							if v, ok := hc["method"].(string); ok && v != "" {
								param["HealthCheck.Method"] = v
							}
							if v, ok := hc["uri"].(string); ok && v != "" {
								param["HealthCheck.URI"] = v
							}
							if v, ok := hc["domain"].(string); ok && v != "" {
								param["HealthCheck.Domain"] = v
							}
							if v, ok := hc["http_code"].(string); ok && v != "" {
								param["HealthCheck.HttpCode"] = v
							}
							if v, ok := hc["interval"].(int); ok {
								param["HealthCheck.Interval"] = v
							}
							if v, ok := hc["timeout"].(int); ok {
								param["HealthCheck.Timeout"] = v
							}
							if v, ok := hc["healthy_threshold"].(int); ok {
								param["HealthCheck.HealthyThreshold"] = v
							}
							if v, ok := hc["unhealthy_threshold"].(int); ok {
								param["HealthCheck.UnhealthyThreshold"] = v
							}
							if v, ok := hc["udp_request"].(string); ok && v != "" {
								param["HealthCheck.UdpRequest"] = v
							}
							if v, ok := hc["udp_expect"].(string); ok && v != "" {
								param["HealthCheck.UdpExpect"] = v
							}
							if v, ok := hc["udp_connect_timeout"].(int); ok {
								param["HealthCheck.UdpConnectTimeout"] = v
							}
						}
					}
				}

				// 手动添加 servers 参数
				if serversList, ok := d.GetOk("servers"); ok {
					if serversList, ok := serversList.([]interface{}); ok {
						for i, server := range serversList {
							serverIndex := i + 1
							if m, ok := server.(map[string]interface{}); ok {
								if v, ok := m["instance_id"].(string); ok && v != "" {
									param[fmt.Sprintf("Servers.%d.InstanceId", serverIndex)] = v
								}
								if v, ok := m["type"].(string); ok && v != "" {
									param[fmt.Sprintf("Servers.%d.Type", serverIndex)] = v
								}
								if v, ok := m["ip"].(string); ok && v != "" {
									param[fmt.Sprintf("Servers.%d.Ip", serverIndex)] = v
								}
								if v, ok := m["port"].(int); ok {
									param[fmt.Sprintf("Servers.%d.Port", serverIndex)] = v
								}
								if v, ok := m["weight"].(int); ok {
									param[fmt.Sprintf("Servers.%d.Weight", serverIndex)] = v
								}
								if v, ok := m["description"].(string); ok && v != "" {
									param[fmt.Sprintf("Servers.%d.Description", serverIndex)] = v
								}
								if v, ok := m["zone_id"].(string); ok && v != "" {
									param[fmt.Sprintf("Servers.%d.ZoneId", serverIndex)] = v
								}
							}
						}
					}
				}

				// 手动添加布尔参数
				if v, ok := d.Get("preserve_client_ip_enabled").(bool); ok {
					param["PreserveClientIpEnabled"] = v
				}
				if v, ok := d.Get("any_port_enabled").(bool); ok {
					param["AnyPortEnabled"] = v
				}
				if v, ok := d.Get("connection_drain_enabled").(bool); ok {
					param["ConnectionDrainEnabled"] = v
				}
				if v, ok := d.Get("session_persistence_enabled").(bool); ok {
					param["SessionPersistenceEnabled"] = v
				}
				if v, ok := d.Get("bypass_security_group_enabled").(bool); ok {
					param["BypassSecurityGroupEnabled"] = v
				}
				if v, ok := d.Get("timestamp_remove_enabled").(bool); ok {
					param["TimestampRemoveEnabled"] = v
				}

				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), &param)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, err := ve.ObtainSdkValue("Result.ServerGroupId", *resp)
				if err != nil {
					return err
				}
				if s, ok := id.(string); ok && s != "" {
					d.SetId(s)
				} else {
					return errors.New("Result.ServerGroupId is not string")
				}
				_, err = s.RefreshResourceState(d, []string{"Active"}, 10*time.Minute, d.Id()).WaitForState()
				return err
			},
		},
	}
	callbacks := []ve.Callback{callback}
	setResourceTagsCallbacks := ve.SetResourceTags(s.Client, "TagNLBResources", "UntagNLBResources", "nlb_servergroup", resourceData, getUniversalInfo)
	callbacks = append(callbacks, setResourceTagsCallbacks...)
	return callbacks
}

func (s *VolcengineNlbServerGroupService) createRequestConvert() map[string]ve.RequestConvert {
	return map[string]ve.RequestConvert{
		"server_group_name": {
			TargetField: "ServerGroupName",
		},
		"description": {
			TargetField: "Description",
		},
		"vpc_id": {
			TargetField: "VpcId",
		},
		"type": {
			TargetField: "Type",
		},
		"protocol": {
			TargetField: "Protocol",
		},
		"scheduler": {
			TargetField: "Scheduler",
		},
		"ip_address_version": {
			TargetField: "IpAddressVersion",
		},
		"any_port_enabled": {
			Ignore: true,
		},
		"connection_drain_enabled": {
			Ignore: true,
		},
		"connection_drain_timeout": {
			TargetField: "ConnectionDrainTimeout",
		},
		"preserve_client_ip_enabled": {
			Ignore: true,
		},
		"session_persistence_enabled": {
			Ignore: true,
		},
		"session_persistence_timeout": {
			TargetField: "SessionPersistenceTimeout",
		},
		"proxy_protocol_type": {
			TargetField: "ProxyProtocolType",
		},
		"bypass_security_group_enabled": {
			Ignore: true,
		},
		"timestamp_remove_enabled": {
			Ignore: true,
		},
		"servers": {
			Ignore: true,
		},
		"health_check": {
			Ignore: true,
		},
		"project_name": {
			TargetField: "ProjectName",
		},
		"tags": {
			Ignore: true,
		},
	}
}

func (s *VolcengineNlbServerGroupService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action: "ModifyNLBServerGroupAttributes",
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				param, err := ve.ResourceDateToRequest(d, resource, true, s.createRequestConvert(), ve.RequestConvertInConvert, ve.ContentTypeDefault)
				if err != nil {
					return nil, err
				}

				// 手动添加 health_check 参数
				if hcList, ok := d.GetOk("health_check"); ok {
					if hcList, ok := hcList.([]interface{}); ok && len(hcList) > 0 {
						if hc, ok := hcList[0].(map[string]interface{}); ok {
							if v, ok := hc["enabled"].(bool); ok {
								param["HealthCheck.Enabled"] = v
							}
							if v, ok := hc["type"].(string); ok && v != "" {
								param["HealthCheck.Type"] = v
							}
							if v, ok := hc["port"].(int); ok {
								param["HealthCheck.Port"] = v
							}
							if v, ok := hc["method"].(string); ok && v != "" {
								param["HealthCheck.Method"] = v
							}
							if v, ok := hc["uri"].(string); ok && v != "" {
								param["HealthCheck.URI"] = v
							}
							if v, ok := hc["domain"].(string); ok && v != "" {
								param["HealthCheck.Domain"] = v
							}
							if v, ok := hc["http_code"].(string); ok && v != "" {
								param["HealthCheck.HttpCode"] = v
							}
							if v, ok := hc["interval"].(int); ok {
								param["HealthCheck.Interval"] = v
							}
							if v, ok := hc["timeout"].(int); ok {
								param["HealthCheck.Timeout"] = v
							}
							if v, ok := hc["healthy_threshold"].(int); ok {
								param["HealthCheck.HealthyThreshold"] = v
							}
							if v, ok := hc["unhealthy_threshold"].(int); ok {
								param["HealthCheck.UnhealthyThreshold"] = v
							}
							if v, ok := hc["udp_request"].(string); ok && v != "" {
								param["HealthCheck.UdpRequest"] = v
							}
							if v, ok := hc["udp_expect"].(string); ok && v != "" {
								param["HealthCheck.UdpExpect"] = v
							}
							if v, ok := hc["udp_connect_timeout"].(int); ok {
								param["HealthCheck.UdpConnectTimeout"] = v
							}
						}
					}
				}

				// 手动添加 servers 参数
				if serversList, ok := d.GetOk("servers"); ok {
					if serversList, ok := serversList.([]interface{}); ok {
						for i, server := range serversList {
							serverIndex := i + 1
							if m, ok := server.(map[string]interface{}); ok {
								if v, ok := m["instance_id"].(string); ok && v != "" {
									param[fmt.Sprintf("Servers.%d.InstanceId", serverIndex)] = v
								}
								if v, ok := m["type"].(string); ok && v != "" {
									param[fmt.Sprintf("Servers.%d.Type", serverIndex)] = v
								}
								if v, ok := m["ip"].(string); ok && v != "" {
									param[fmt.Sprintf("Servers.%d.Ip", serverIndex)] = v
								}
								if v, ok := m["port"].(int); ok {
									param[fmt.Sprintf("Servers.%d.Port", serverIndex)] = v
								}
								if v, ok := m["weight"].(int); ok {
									param[fmt.Sprintf("Servers.%d.Weight", serverIndex)] = v
								}
								if v, ok := m["description"].(string); ok && v != "" {
									param[fmt.Sprintf("Servers.%d.Description", serverIndex)] = v
								}
								if v, ok := m["zone_id"].(string); ok && v != "" {
									param[fmt.Sprintf("Servers.%d.ZoneId", serverIndex)] = v
								}
							}
						}
					}
				}

				// 手动添加布尔参数
				if v, ok := d.Get("preserve_client_ip_enabled").(bool); ok {
					param["PreserveClientIpEnabled"] = v
				}
				if v, ok := d.Get("any_port_enabled").(bool); ok {
					param["AnyPortEnabled"] = v
				}
				if v, ok := d.Get("connection_drain_enabled").(bool); ok {
					param["ConnectionDrainEnabled"] = v
				}
				if v, ok := d.Get("session_persistence_enabled").(bool); ok {
					param["SessionPersistenceEnabled"] = v
				}
				if v, ok := d.Get("bypass_security_group_enabled").(bool); ok {
					param["BypassSecurityGroupEnabled"] = v
				}
				if v, ok := d.Get("timestamp_remove_enabled").(bool); ok {
					param["TimestampRemoveEnabled"] = v
				}

				param["ServerGroupId"] = d.Id()
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), &param)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				_, err := s.RefreshResourceState(d, []string{"Active"}, 10*time.Minute, d.Id()).WaitForState()
				return err
			},
		},
	}
	callbacks := []ve.Callback{callback}
	setResourceTagsCallbacks := ve.SetResourceTags(s.Client, "TagNLBResources", "UntagNLBResources", "nlb_servergroup", resourceData, getUniversalInfo)
	callbacks = append(callbacks, setResourceTagsCallbacks...)
	return callbacks
}

func (s *VolcengineNlbServerGroupService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteNLBServerGroup",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"ServerGroupId": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 10*time.Minute)
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineNlbServerGroupService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"ServerGroupId": {
				TargetField: "server_group_id",
			},
			"ServerGroupName": {
				TargetField: "server_group_name",
			},
			"Description": {
				TargetField: "description",
			},
			"VpcId": {
				TargetField: "vpc_id",
			},
			"Type": {
				TargetField: "type",
			},
			"Protocol": {
				TargetField: "protocol",
			},
			"Scheduler": {
				TargetField: "scheduler",
			},
			"IpAddressVersion": {
				TargetField: "ip_address_version",
			},
			"AnyPortEnabled": {
				TargetField: "any_port_enabled",
			},
			"ConnectionDrainEnabled": {
				TargetField: "connection_drain_enabled",
			},
			"ConnectionDrainTimeout": {
				TargetField: "connection_drain_timeout",
			},
			"PreserveClientIpEnabled": {
				TargetField: "preserve_client_ip_enabled",
			},
			"SessionPersistenceEnabled": {
				TargetField: "session_persistence_enabled",
			},
			"SessionPersistenceTimeout": {
				TargetField: "session_persistence_timeout",
			},
			"ProxyProtocolType": {
				TargetField: "proxy_protocol_type",
			},
			"BypassSecurityGroupEnabled": {
				TargetField: "bypass_security_group_enabled",
			},
			"TimestampRemoveEnabled": {
				TargetField: "timestamp_remove_enabled",
			},
			"ServerCount": {
				TargetField: "server_count",
			},
			"Servers": {
				TargetField: "servers",
				Convert:     transServersToResponse,
			},
			"HealthCheck": {
				TargetField: "health_check",
				Convert:     transHealthCheckToResponse,
			},
			"ProjectName": {
				TargetField: "project_name",
			},
			"Tags": {
				TargetField: "tags",
				Convert:     ve.TransTagsToResponse,
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineNlbServerGroupService) DatasourceResources(d *schema.ResourceData, r *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"vpc_id": {
				TargetField: "VpcId",
			},
			"server_group_name": {
				TargetField: "ServerGroupName",
			},
			"instance_ids": {
				TargetField: "InstanceIds",
				ConvertType: ve.ConvertWithN,
			},
			"server_ips": {
				TargetField: "ServerIps",
				ConvertType: ve.ConvertWithN,
			},
			"type": {
				TargetField: "Type",
			},
			"server_group_ids": {
				TargetField: "ServerGroupIds",
				ConvertType: ve.ConvertWithN,
			},
			"project_name": {
				TargetField: "ProjectName",
			},
			"tags": {
				TargetField: "TagFilters",
				Convert:     ve.TransTagFiltersToRequest,
			},
		},
		CollectField: "server_groups",
		IdField:      "ServerGroupId",
		ResponseConverts: map[string]ve.ResponseConvert{
			"ServerGroupId": {
				TargetField: "server_group_id",
			},
			"ServerGroupName": {
				TargetField: "server_group_name",
			},
			"Description": {
				TargetField: "description",
			},
			"AccountId": {
				TargetField: "account_id",
			},
			"CreateTime": {
				TargetField: "create_time",
			},
			"UpdateTime": {
				TargetField: "update_time",
			},
			"Status": {
				TargetField: "status",
			},
			"VpcId": {
				TargetField: "vpc_id",
			},
			"Type": {
				TargetField: "type",
			},
			"Protocol": {
				TargetField: "protocol",
			},
			"Scheduler": {
				TargetField: "scheduler",
			},
			"IpAddressVersion": {
				TargetField: "ip_address_version",
			},
			"ProjectName": {
				TargetField: "project_name",
			},
			"AnyPortEnabled": {
				TargetField: "any_port_enabled",
			},
			"ConnectionDrainEnabled": {
				TargetField: "connection_drain_enabled",
			},
			"ConnectionDrainTimeout": {
				TargetField: "connection_drain_timeout",
			},
			"PreserveClientIpEnabled": {
				TargetField: "preserve_client_ip_enabled",
			},
			"SessionPersistenceEnabled": {
				TargetField: "session_persistence_enabled",
			},
			"SessionPersistenceTimeout": {
				TargetField: "session_persistence_timeout",
			},
			"ProxyProtocolType": {
				TargetField: "proxy_protocol_type",
			},
			"BypassSecurityGroupEnabled": {
				TargetField: "bypass_security_group_enabled",
			},
			"TimestampRemoveEnabled": {
				TargetField: "timestamp_remove_enabled",
			},
			"ServerCount": {
				TargetField: "server_count",
			},
			"Servers": {
				TargetField: "servers",
				Convert:     transServersToResponse,
			},
			"RelatedLoadBalancerIds": {
				TargetField: "related_load_balancer_ids",
			},
			"HealthCheck": {
				TargetField: "health_check",
				Convert:     transHealthCheckToResponse,
			},
			"Tags": {
				TargetField: "tags",
				Convert:     ve.TransTagsToResponse,
			},
		},
	}
}

func (s *VolcengineNlbServerGroupService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "clb",
		Action:      actionName,
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		RegionType:  ve.Regional,
	}
}

func transTagsToRequest(d *schema.ResourceData, i interface{}) interface{} {
	if i == nil {
		return nil
	}
	var tags []map[string]interface{}
	if set, ok := i.(*schema.Set); ok {
		for _, v := range set.List() {
			if m, ok := v.(map[string]interface{}); ok {
				tag := make(map[string]interface{})
				if key, ok := m["key"].(string); ok {
					tag["Key"] = key
				}
				if value, ok := m["value"].(string); ok {
					tag["Value"] = value
				}
				tags = append(tags, tag)
			}
		}
	}
	return tags
}

func transServersToRequest(d *schema.ResourceData, i interface{}) interface{} {
	if i == nil {
		return nil
	}
	res := make(map[string]interface{})
	if list, ok := i.([]interface{}); ok {
		for idx, item := range list {
			serverIndex := idx + 1
			if m, ok := item.(map[string]interface{}); ok {
				if v, ok := m["instance_id"].(string); ok && v != "" {
					res[fmt.Sprintf("Servers.%d.InstanceId", serverIndex)] = v
				}
				if v, ok := m["type"].(string); ok && v != "" {
					res[fmt.Sprintf("Servers.%d.Type", serverIndex)] = v
				}
				if v, ok := m["ip"].(string); ok && v != "" {
					res[fmt.Sprintf("Servers.%d.Ip", serverIndex)] = v
				}
				if v, ok := m["port"].(int); ok {
					res[fmt.Sprintf("Servers.%d.Port", serverIndex)] = v
				}
				if v, ok := m["weight"].(int); ok {
					res[fmt.Sprintf("Servers.%d.Weight", serverIndex)] = v
				}
				if v, ok := m["description"].(string); ok && v != "" {
					res[fmt.Sprintf("Servers.%d.Description", serverIndex)] = v
				}
				if v, ok := m["zone_id"].(string); ok && v != "" {
					res[fmt.Sprintf("Servers.%d.ZoneId", serverIndex)] = v
				}
			}
		}
	}
	return res
}

func transHealthCheckToRequest(d *schema.ResourceData, i interface{}) interface{} {
	if i == nil {
		return nil
	}
	res := make(map[string]interface{})
	if list, ok := i.([]interface{}); ok && len(list) > 0 {
		if m, ok := list[0].(map[string]interface{}); ok {
			if v, ok := m["enabled"].(bool); ok {
				res["HealthCheck.Enabled"] = v
			}
			if v, ok := m["type"].(string); ok && v != "" {
				res["HealthCheck.Type"] = v
			}
			if v, ok := m["port"].(int); ok {
				res["HealthCheck.Port"] = v
			}
			if v, ok := m["method"].(string); ok && v != "" {
				res["HealthCheck.Method"] = v
			}
			if v, ok := m["uri"].(string); ok && v != "" {
				res["HealthCheck.URI"] = v
			}
			if v, ok := m["domain"].(string); ok && v != "" {
				res["HealthCheck.Domain"] = v
			}
			if v, ok := m["http_code"].(string); ok && v != "" {
				res["HealthCheck.HttpCode"] = v
			}
			if v, ok := m["interval"].(int); ok {
				res["HealthCheck.Interval"] = v
			}
			if v, ok := m["timeout"].(int); ok {
				res["HealthCheck.Timeout"] = v
			}
			if v, ok := m["healthy_threshold"].(int); ok {
				res["HealthCheck.HealthyThreshold"] = v
			}
			if v, ok := m["unhealthy_threshold"].(int); ok {
				res["HealthCheck.UnhealthyThreshold"] = v
			}
			if v, ok := m["udp_request"].(string); ok && v != "" {
				res["HealthCheck.UdpRequest"] = v
			}
			if v, ok := m["udp_expect"].(string); ok && v != "" {
				res["HealthCheck.UdpExpect"] = v
			}
			if v, ok := m["udp_connect_timeout"].(int); ok {
				res["HealthCheck.UdpConnectTimeout"] = v
			}
		}
	}
	return res
}

func transHealthCheckToResponse(i interface{}) interface{} {
	if i == nil {
		return nil
	}
	m, ok := i.(map[string]interface{})
	if !ok {
		return i
	}
	res := make(map[string]interface{})
	for k, v := range m {
		res[ve.HumpToDownLine(k)] = v
	}
	return res
}

func transServersToResponse(i interface{}) interface{} {
	if i == nil {
		return nil
	}
	list, ok := i.([]interface{})
	if !ok {
		return i
	}
	var res []interface{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		if !ok {
			res = append(res, item)
			continue
		}
		newMap := make(map[string]interface{})
		for k, v := range m {
			newMap[ve.HumpToDownLine(k)] = v
		}
		res = append(res, newMap)
	}
	return res
}
