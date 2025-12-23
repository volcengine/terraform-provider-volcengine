package rds_postgresql_instance

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

type VolcengineRdsPostgresqlInstanceService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewRdsPostgresqlInstanceService(c *ve.SdkClient) *VolcengineRdsPostgresqlInstanceService {
	return &VolcengineRdsPostgresqlInstanceService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineRdsPostgresqlInstanceService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineRdsPostgresqlInstanceService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeDBInstances"
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
		results, err = ve.ObtainSdkValue("Result.Instances", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.Instances is not Slice")
		}
		// append details
		for _, v := range data {
			var (
				basicInfo    interface{}
				endpointInfo interface{}
				nodeInfo     interface{}
			)
			action = "DescribeDBInstanceDetail"
			instance := v.(map[string]interface{})

			// DescribeDBInstanceDetail
			req := map[string]interface{}{
				"InstanceId": instance["InstanceId"],
			}
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &req)
			if err != nil {
				logger.Info("DescribeDBInstanceDetail error:", err)
				continue
			}
			respBytes, _ = json.Marshal(resp)
			logger.Debug(logger.RespFormat, action, req, string(respBytes))

			// append basic info
			basicInfo, err = ve.ObtainSdkValue("Result.BasicInfo", *resp)
			if err != nil {
				logger.Info("ObtainSdkValue Result.BasicInfo error:", err)
				continue
			}
			if basicInfoMap, ok := basicInfo.(map[string]interface{}); ok {
				instance["VCPU"] = basicInfoMap["VCPU"]
				instance["Memory"] = basicInfoMap["Memory"]
				instance["UpdateTime"] = basicInfoMap["UpdateTime"]
				// DescribeDBInstanceDetail API 的返回字段中已没有 BackupUse 字段，赋默认值 0
				instance["BackupUse"] = 0
				instance["DataSyncMode"] = basicInfoMap["DataSyncMode"]
				instance["StorageDataUse"] = basicInfoMap["StorageDataUse"]
				instance["StorageLogUse"] = basicInfoMap["StorageLogUse"]
				instance["StorageTempUse"] = basicInfoMap["StorageTempUse"]
				instance["StorageUse"] = basicInfoMap["StorageUse"]
				instance["StorageWALUse"] = basicInfoMap["StorageWALUse"]
			}

			// append endpoint info
			endpointInfo, err = ve.ObtainSdkValue("Result.Endpoints", *resp)
			if err != nil {
				logger.Info("ObtainSdkValue Result.Endpoints error:", err)
				continue
			}
			if infos, ok := endpointInfo.([]interface{}); ok {
				instance["Endpoints"] = infos
			} else {
				// 接口返回nil
				instance["Endpoints"] = []interface{}{}
			}

			// append node info
			nodeInfo, err = ve.ObtainSdkValue("Result.Nodes", *resp)
			if err != nil {
				logger.Info("ObtainSdkValue Result.Nodes error:", err)
				continue
			}
			if infos, ok := nodeInfo.([]interface{}); ok {
				instance["Nodes"] = infos
			} else {
				// 接口返回nil
				instance["Nodes"] = []interface{}{}
			}

			// 默认化估算结果，避免 plan 阶段出现 known after apply 的差异
			if _, ok := instance["EstimationResult"]; !ok {
				instance["EstimationResult"] = []interface{}{}
			}
		}
		return data, err
	})
}

func (s *VolcengineRdsPostgresqlInstanceService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"InstanceId": id,
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		if data, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("value is not map")
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("the Rds PostgreSQL instance %s not exist", id)
	}

	if nodeArr, ok := data["Nodes"].([]interface{}); ok {
		for _, node := range nodeArr {
			if nodeMap, ok1 := node.(map[string]interface{}); ok1 {
				if nodeMap["NodeType"] == "Primary" {
					data["PrimaryZoneId"] = nodeMap["ZoneId"]
				} else if nodeMap["NodeType"] == "Secondary" {
					data["SecondaryZoneId"] = nodeMap["ZoneId"]
				}
			}
		}
	}

	// Set特殊处理
	if parameterSet, ok := resourceData.GetOk("parameters"); ok {
		data["Parameters"] = parameterSet.(*schema.Set).List()
	}

	data["ChargeInfo"] = data["ChargeDetail"]

	if v, ok := resourceData.GetOk("estimation_result"); ok {
		data["EstimationResult"] = v
	}

	return data, err
}

func (s *VolcengineRdsPostgresqlInstanceService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				demo       map[string]interface{}
				status     interface{}
				failStates []string
			)
			failStates = append(failStates, "Error")

			if err = resource.Retry(20*time.Minute, func() *resource.RetryError {
				demo, err = s.ReadResource(resourceData, id)
				if err != nil {
					if ve.ResourceNotFoundError(err) {
						return resource.RetryableError(err)
					} else {
						return resource.NonRetryableError(err)
					}
				}
				return nil
			}); err != nil {
				return nil, "", err
			}

			status, err = ve.ObtainSdkValue("InstanceStatus", demo)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("the Rds PostgreSQL instance status error, status:%s ", status.(string))
				}
			}
			//注意 返回的第一个参数不能为空 否则会一直等下去
			return demo, status.(string), err
		},
	}
}

func (s *VolcengineRdsPostgresqlInstanceService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback
	// 如果指定了 src_instance_id，则走 RestoreToNewInstance
	if _, hasRestore := resourceData.GetOk("src_instance_id"); hasRestore {
		restore := ve.Callback{
			Call: ve.SdkCall{
				Action:      "RestoreToNewInstance",
				ConvertMode: ve.RequestConvertAll,
				ContentType: ve.ContentTypeJson,
				Convert: map[string]ve.RequestConvert{
					"src_instance_id":   {TargetField: "SrcInstanceId"},
					"backup_id":         {TargetField: "BackupId"},
					"restore_time":      {TargetField: "RestoreTime"},
					"db_engine_version": {Ignore: true},
					"tags":              {TargetField: "Tags", ConvertType: ve.ConvertJsonObjectArray},
					"charge_info":       {ConvertType: ve.ConvertJsonObject},
					"allow_list_ids":    {TargetField: "AllowListIds", ConvertType: ve.ConvertJsonArray},
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					// RestoreToNewInstance API 没有 DBEngineVersion 参数，需要移除可能残留的值
					delete(*call.SdkParam, "db_engine_version")
					delete(*call.SdkParam, "DBEngineVersion")
					// 直接复用正常创建的beforecall
					var (
						nodeInfos []interface{}
						subnets   []interface{}
						results   interface{}
						ok        bool
					)
					// add vpc id
					subnetId := d.Get("subnet_id")
					req := map[string]interface{}{
						"SubnetIds.1": subnetId,
					}
					action := "DescribeSubnets"
					resp, err := s.Client.UniversalClient.DoCall(getVPCUniversalInfo(action), &req)
					if err != nil {
						return false, err
					}
					results, err = ve.ObtainSdkValue("Result.Subnets", *resp)
					if err != nil {
						return false, err
					}
					if results == nil {
						results = []interface{}{}
					}
					if subnets, ok = results.([]interface{}); !ok {
						return false, errors.New("Result.Subnets is not Slice")
					}
					if len(subnets) == 0 {
						return false, fmt.Errorf("subnet %s not exist", subnetId.(string))
					}
					vpcId := subnets[0].(map[string]interface{})["VpcId"]

					(*call.SdkParam)["VpcId"] = vpcId

					// add NodeInfo，默认一主一备的高可用架构
					primaryNodeInfo := make(map[string]interface{})
					primaryNodeInfo["NodeType"] = "Primary"
					primaryNodeInfo["ZoneId"] = d.Get("primary_zone_id")
					primaryNodeInfo["NodeSpec"] = d.Get("node_spec")
					primaryNodeInfo["NodeOperateType"] = "Create"
					nodeInfos = append(nodeInfos, primaryNodeInfo)

					secondaryNodeInfo := make(map[string]interface{})
					secondaryNodeInfo["NodeType"] = "Secondary"
					secondaryNodeInfo["ZoneId"] = d.Get("secondary_zone_id")
					secondaryNodeInfo["NodeSpec"] = d.Get("node_spec")
					secondaryNodeInfo["NodeOperateType"] = "Create"
					nodeInfos = append(nodeInfos, secondaryNodeInfo)

					(*call.SdkParam)["NodeInfo"] = nodeInfos
					(*call.SdkParam)["StorageType"] = "LocalSSD"

					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
				AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
					id, _ := ve.ObtainSdkValue("Result.InstanceId", *resp)
					d.SetId(id.(string))
					return nil
				},
				Refresh: &ve.StateRefresh{Target: []string{"Running"}, Timeout: resourceData.Timeout(schema.TimeoutCreate)},
			},
		}
		callbacks = append(callbacks, restore)
		return callbacks
	}

	// 否则走常规创建实例
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateDBInstance",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"db_engine_version": {
					TargetField: "DBEngineVersion",
				},
				"tags": {
					TargetField: "Tags",
					ConvertType: ve.ConvertJsonObjectArray,
				},
				"charge_info": {
					ConvertType: ve.ConvertJsonObject,
				},
				"allow_list_ids": {
					TargetField: "AllowListIds",
					ConvertType: ve.ConvertJsonArray,
				},
				// node ignore
				"node_spec": {
					Ignore: true,
				},
				"primary_zone_id": {
					Ignore: true,
				},
				"secondary_zone_id": {
					Ignore: true,
				},
				"parameters": {
					Ignore: true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				var (
					nodeInfos []interface{}
					subnets   []interface{}
					results   interface{}
					ok        bool
				)
				// add vpc id
				subnetId := d.Get("subnet_id")
				req := map[string]interface{}{
					"SubnetIds.1": subnetId,
				}
				action := "DescribeSubnets"
				resp, err := s.Client.UniversalClient.DoCall(getVPCUniversalInfo(action), &req)
				if err != nil {
					return false, err
				}
				results, err = ve.ObtainSdkValue("Result.Subnets", *resp)
				if err != nil {
					return false, err
				}
				if results == nil {
					results = []interface{}{}
				}
				if subnets, ok = results.([]interface{}); !ok {
					return false, errors.New("Result.Subnets is not Slice")
				}
				if len(subnets) == 0 {
					return false, fmt.Errorf("subnet %s not exist", subnetId.(string))
				}
				vpcId := subnets[0].(map[string]interface{})["VpcId"]

				(*call.SdkParam)["VpcId"] = vpcId

				// add NodeInfo，默认一主一备的高可用架构
				primaryNodeInfo := make(map[string]interface{})
				primaryNodeInfo["NodeType"] = "Primary"
				primaryNodeInfo["ZoneId"] = d.Get("primary_zone_id")
				primaryNodeInfo["NodeSpec"] = d.Get("node_spec")
				primaryNodeInfo["NodeOperateType"] = "Create"
				nodeInfos = append(nodeInfos, primaryNodeInfo)

				secondaryNodeInfo := make(map[string]interface{})
				secondaryNodeInfo["NodeType"] = "Secondary"
				secondaryNodeInfo["ZoneId"] = d.Get("secondary_zone_id")
				secondaryNodeInfo["NodeSpec"] = d.Get("node_spec")
				secondaryNodeInfo["NodeOperateType"] = "Create"
				nodeInfos = append(nodeInfos, secondaryNodeInfo)

				(*call.SdkParam)["NodeInfo"] = nodeInfos

				(*call.SdkParam)["StorageType"] = "LocalSSD"

				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("Result.InstanceId", *resp)
				d.SetId(id.(string))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Running"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	callbacks = append(callbacks, callback)

	// parameters callback
	if _, ok := resourceData.GetOk("parameters"); ok {
		paramCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyDBInstanceParameters",
				ContentType: ve.ContentTypeJson,
				ConvertMode: ve.RequestConvertInConvert,
				Convert: map[string]ve.RequestConvert{
					"parameters": {
						ConvertType: ve.ConvertJsonObjectArray,
					},
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					if len(*call.SdkParam) > 0 {
						(*call.SdkParam)["InstanceId"] = d.Id()
						return true, nil
					}
					return false, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{"Running"},
					Timeout: resourceData.Timeout(schema.TimeoutCreate),
				},
			},
		}
		callbacks = append(callbacks, paramCallback)
	}
	return callbacks
}

func (VolcengineRdsPostgresqlInstanceService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"DBEngineVersion": {
				TargetField: "db_engine_version",
			},
			"EstimationResult": {
				TargetField: "estimation_result",
				KeepDefault: true,
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineRdsPostgresqlInstanceService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	// ModifyDBInstanceName
	if resourceData.HasChange("instance_name") {
		nameCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyDBInstanceName",
				ContentType: ve.ContentTypeJson,
				ConvertMode: ve.RequestConvertInConvert,
				Convert: map[string]ve.RequestConvert{
					"instance_name": {
						TargetField: "InstanceNewName",
					},
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					if len(*call.SdkParam) > 0 {
						(*call.SdkParam)["InstanceId"] = d.Id()
						return true, nil
					}
					return false, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					common, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
					if err != nil {
						return common, err
					}
					return common, nil
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{"Running"},
					Timeout: resourceData.Timeout(schema.TimeoutUpdate),
				},
			},
		}
		callbacks = append(callbacks, nameCallback)
	}

	// ModifyDBInstanceChargeType
	if resourceData.HasChange("charge_info") {
		chargeCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyDBInstanceChargeType",
				ContentType: ve.ContentTypeJson,
				ConvertMode: ve.RequestConvertIgnore,
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					// 文档约束：仅支持将 PostPaid 转为 PrePaid，若希望转回 PostPaid，则会报错
					_, newType := d.GetChange("charge_info.0.charge_type")
					if newType == "PostPaid" {
						return false, fmt.Errorf("only support convert PostPaid to PrePaid")
					}
					(*call.SdkParam)["InstanceId"] = d.Id()
					(*call.SdkParam)["ChargeType"] = newType
					(*call.SdkParam)["PeriodUnit"] = d.Get("charge_info.0.period_unit")
					(*call.SdkParam)["Period"] = d.Get("charge_info.0.period")
					(*call.SdkParam)["AutoRenew"] = d.Get("charge_info.0.auto_renew")
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
				Refresh: &ve.StateRefresh{Target: []string{"Running"}, Timeout: resourceData.Timeout(schema.TimeoutUpdate)},
			},
		}
		callbacks = append(callbacks, chargeCallback)
	}

	// ModifyDBInstanceAvailabilityZone
	if resourceData.HasChange("zone_migrations") || resourceData.HasChange("secondary_zone_id") {
		zoneCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyDBInstanceAvailabilityZone",
				ContentType: ve.ContentTypeJson,
				ConvertMode: ve.RequestConvertIgnore,
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					if d.HasChange("secondary_zone_id") {
						if _, ok := d.GetOk("zone_migrations"); !ok {
							return false, fmt.Errorf("the zone_migrations field is needed to migrate the secondary node")
						}
					}
					(*call.SdkParam)["InstanceId"] = d.Id()

					instance, err := s.ReadResource(d, d.Id())
					if err != nil {
						return false, err
					}
					var primaryZone, secondaryZone string
					if nodeArr, ok := instance["Nodes"].([]interface{}); ok {
						for _, node := range nodeArr {
							if nodeMap, ok1 := node.(map[string]interface{}); ok1 {
								if nodeMap["NodeType"] == "Primary" {
									if z, ok2 := nodeMap["ZoneId"].(string); ok2 {
										primaryZone = z
									}
								} else if nodeMap["NodeType"] == "Secondary" {
									if z, ok2 := nodeMap["ZoneId"].(string); ok2 {
										secondaryZone = z
									}
								}
							}
						}
					}
					if primaryZone != "" && secondaryZone != "" && primaryZone != secondaryZone {
						return false, fmt.Errorf("Cross-AZ instance migration is not supported, currently, Primary vs Secondary: %s vs %s", primaryZone, secondaryZone)
					}
					if v, ok := d.GetOk("zone_migrations"); ok {
						arr := v.([]interface{})
						nodeInfo := make([]map[string]interface{}, 0)
						for _, item := range arr {
							m, ok := item.(map[string]interface{})
							if !ok {
								continue
							}
							if t, ok := m["node_type"].(string); ok && (t == "Secondary" || t == "ReadOnly") {
								node := map[string]interface{}{
									"NodeId": m["node_id"],
									"ZoneId": m["zone_id"],
								}
								nodeInfo = append(nodeInfo, node)
							}
						}
						if len(nodeInfo) > 0 {
							(*call.SdkParam)["NodeInfo"] = nodeInfo
							return true, nil
						}
					}
					return false, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
				Refresh: &ve.StateRefresh{Target: []string{"Running"}, Timeout: resourceData.Timeout(schema.TimeoutUpdate)},
			},
		}
		callbacks = append(callbacks, zoneCallback)
	}

	// ModifyDBInstanceSpec
	if resourceData.HasChanges("node_spec", "storage_space") || (resourceData.HasChange("estimate_only") && resourceData.Get("estimate_only").(bool)) {

		instanceCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyDBInstanceSpec",
				ContentType: ve.ContentTypeJson,
				ConvertMode: ve.RequestConvertIgnore,
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["InstanceId"] = d.Id()

					if d.HasChange("storage_space") {
						(*call.SdkParam)["StorageType"] = "LocalSSD"
						(*call.SdkParam)["StorageSpace"] = d.Get("storage_space")
					}

					if d.HasChange("node_spec") {
						nodeInfos := make([]interface{}, 0)
						primaryNodeInfo := make(map[string]interface{})
						secondaryNodeInfo := make(map[string]interface{})

						instance, err := s.ReadResource(resourceData, d.Id())
						if err != nil {
							return false, err
						}
						if nodeArr, ok := instance["Nodes"].([]interface{}); ok {
							for _, node := range nodeArr {
								if nodeMap, ok1 := node.(map[string]interface{}); ok1 {
									if nodeMap["NodeType"] == "Primary" {
										primaryNodeInfo["NodeId"] = nodeMap["NodeId"]
									} else if nodeMap["NodeType"] == "Secondary" {
										secondaryNodeInfo["NodeId"] = nodeMap["NodeId"]
									} else if nodeMap["NodeType"] == "ReadOnly" {
										readonlyNodeInfo := make(map[string]interface{})
										readonlyNodeInfo["NodeId"] = nodeMap["NodeId"]
										readonlyNodeInfo["NodeType"] = nodeMap["NodeType"]
										readonlyNodeInfo["NodeSpec"] = nodeMap["NodeSpec"]
										readonlyNodeInfo["ZoneId"] = nodeMap["ZoneId"]
										nodeInfos = append(nodeInfos, readonlyNodeInfo)
									}
								}
							}
						}

						primaryNodeInfo["NodeType"] = "Primary"
						primaryNodeInfo["ZoneId"] = d.Get("primary_zone_id")
						primaryNodeInfo["NodeSpec"] = d.Get("node_spec")
						primaryNodeInfo["NodeOperateType"] = "Modify"
						nodeInfos = append(nodeInfos, primaryNodeInfo)

						secondaryNodeInfo["NodeType"] = "Secondary"
						secondaryNodeInfo["ZoneId"] = d.Get("secondary_zone_id")
						secondaryNodeInfo["NodeSpec"] = d.Get("node_spec")
						secondaryNodeInfo["NodeOperateType"] = "Modify"
						nodeInfos = append(nodeInfos, secondaryNodeInfo)

						(*call.SdkParam)["NodeInfo"] = nodeInfos
					}

					// Temporary mode handling
					if v, ok := d.GetOk("modify_type"); ok && v.(string) == "Temporary" {
						(*call.SdkParam)["ModifyType"] = "Temporary"
						if rt, ok2 := d.GetOk("rollback_time"); ok2 {
							(*call.SdkParam)["RollbackTime"] = rt.(string)
						}
						delete(*call.SdkParam, "StorageSpace")
						delete(*call.SdkParam, "StorageType")
					} else if v, ok := d.GetOk("modify_type"); ok {
						(*call.SdkParam)["ModifyType"] = v
					}
					if v, ok := d.GetOk("estimate_only"); ok && v.(bool) {
						(*call.SdkParam)["EstimateOnly"] = true
					}
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
				Refresh: &ve.StateRefresh{Target: []string{"Running"}, Timeout: resourceData.Timeout(schema.TimeoutUpdate)},
				AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
					if v, ok := d.GetOk("estimate_only"); ok && v.(bool) {
						est, _ := ve.ObtainSdkValue("Result.EstimationResult", *resp)
						if m, ok2 := est.(map[string]interface{}); ok2 {
							plans := []string{}
							effects := []string{}
							if p, okp := m["Plans"].([]interface{}); okp {
								for _, it := range p {
									if s2, ok3 := it.(string); ok3 {
										plans = append(plans, s2)
									}
								}
							}
							if e, oke := m["Effects"].([]interface{}); oke {
								for _, it := range e {
									if s2, ok3 := it.(string); ok3 {
										effects = append(effects, s2)
									}
								}
							}
							_ = d.Set("estimation_result", []map[string]interface{}{{"plans": plans, "effects": effects}})
						}
					}
					return nil
				},
			},
		}
		callbacks = append(callbacks, instanceCallback)
	}

	// ModifyDBInstanceParameters
	if resourceData.HasChange("parameters") {
		modifiedParams, _, _, _ := ve.GetSetDifference("parameters", resourceData, parameterHash, false)

		parameterCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyDBInstanceParameters",
				ContentType: ve.ContentTypeJson,
				ConvertMode: ve.RequestConvertIgnore,
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					if modifiedParams != nil && len(modifiedParams.List()) > 0 {
						(*call.SdkParam)["InstanceId"] = d.Id()
						(*call.SdkParam)["Parameters"] = make([]map[string]interface{}, 0)
						for _, v := range modifiedParams.List() {
							paramMap, ok := v.(map[string]interface{})
							if !ok {
								return false, fmt.Errorf("parameter is not map ")
							}
							(*call.SdkParam)["Parameters"] = append((*call.SdkParam)["Parameters"].([]map[string]interface{}), map[string]interface{}{
								"Name":  paramMap["name"],
								"Value": paramMap["value"],
							})
						}
						return true, nil
					}
					return false, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{"Running"},
					Timeout: resourceData.Timeout(schema.TimeoutUpdate),
				},
			},
		}
		callbacks = append(callbacks, parameterCallback)
	}

	// Tags
	callbacks = s.setResourceTags(resourceData, callbacks)

	return callbacks
}

func (s *VolcengineRdsPostgresqlInstanceService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteDBInstance",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"InstanceId": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 10*time.Minute)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading rds postgre instance on delete %q, %w", d.Id(), callErr))
						}
					}
					_, callErr = call.ExecuteCall(d, client, call)
					if callErr == nil {
						return nil
					}
					return resource.RetryableError(callErr)
				})
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineRdsPostgresqlInstanceService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"db_engine_version": {
				TargetField: "DBEngineVersion",
			},
			"tags": {
				TargetField: "TagFilters",
				ConvertType: ve.ConvertJsonObjectArray,
			},
		},
		NameField:    "InstanceName",
		IdField:      "InstanceId",
		CollectField: "instances",
		ContentType:  ve.ContentTypeJson,
		ResponseConverts: map[string]ve.ResponseConvert{
			"InstanceId": {
				TargetField: "id",
				KeepDefault: true,
			},
			"DBEngineVersion": {
				TargetField: "db_engine_version",
			},
			"IPAddress": {
				TargetField: "ip_address",
			},
			"DNSVisibility": {
				TargetField: "dns_visibility",
			},
			"VCPU": {
				TargetField: "v_cpu",
			},
			"StorageWALUse": {
				TargetField: "storage_wal_use",
			},
		},
	}
}

func (s *VolcengineRdsPostgresqlInstanceService) ReadResourceId(id string) string {
	return id
}

func (s *VolcengineRdsPostgresqlInstanceService) setResourceTags(resourceData *schema.ResourceData, callbacks []ve.Callback) []ve.Callback {
	addedTags, removedTags, _, _ := ve.GetSetDifference("tags", resourceData, ve.TagsHash, false)

	removeCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "RemoveTagsFromResource",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if removedTags != nil && len(removedTags.List()) > 0 {
					(*call.SdkParam)["InstanceIds"] = []string{resourceData.Id()}
					(*call.SdkParam)["TagKeys"] = make([]string, 0)
					for _, tag := range removedTags.List() {
						(*call.SdkParam)["TagKeys"] = append((*call.SdkParam)["TagKeys"].([]string), tag.(map[string]interface{})["key"].(string))
					}
					return true, nil
				}
				return false, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	callbacks = append(callbacks, removeCallback)

	addCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "AddTagsToResource",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if addedTags != nil && len(addedTags.List()) > 0 {
					(*call.SdkParam)["InstanceIds"] = []string{resourceData.Id()}
					(*call.SdkParam)["Tags"] = make([]map[string]interface{}, 0)
					for _, tag := range addedTags.List() {
						(*call.SdkParam)["Tags"] = append((*call.SdkParam)["Tags"].([]map[string]interface{}), tag.(map[string]interface{}))
					}
					return true, nil
				}
				return false, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	callbacks = append(callbacks, addCallback)

	return callbacks
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

func (s *VolcengineRdsPostgresqlInstanceService) ProjectTrn() *ve.ProjectTrn {
	return &ve.ProjectTrn{
		ServiceName:          "rds_postgresql",
		ResourceType:         "instance",
		ProjectResponseField: "ProjectName",
		ProjectSchemaField:   "project_name",
	}
}

func getVPCUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "vpc",
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		Action:      actionName,
	}
}

func (s *VolcengineRdsPostgresqlInstanceService) UnsubscribeInfo(resourceData *schema.ResourceData, resource *schema.Resource) (*ve.UnsubscribeInfo, error) {
	info := ve.UnsubscribeInfo{
		InstanceId: s.ReadResourceId(resourceData.Id()),
	}
	if resourceData.Get("charge_info.0.charge_type").(string) == "PrePaid" {
		info.NeedUnsubscribe = true
		info.Products = []string{"RDS for PostgreSQL"}
	}
	return &info, nil
}
