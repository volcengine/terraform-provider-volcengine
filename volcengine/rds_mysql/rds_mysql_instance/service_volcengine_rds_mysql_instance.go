package rds_mysql_instance

import (
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineRdsMysqlInstanceService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewRdsMysqlInstanceService(c *ve.SdkClient) *VolcengineRdsMysqlInstanceService {
	return &VolcengineRdsMysqlInstanceService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineRdsMysqlInstanceService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineRdsMysqlInstanceService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp        *map[string]interface{}
		results     interface{}
		ok          bool
		rdsInstance map[string]interface{}
	)
	data, err = ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 10, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeDBInstances"
		logger.Debug(logger.ReqFormat, action, condition)
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
		return data, err
	})

	if err != nil {
		return nil, err
	}

	for _, v := range data {
		if rdsInstance, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("Value is not map ")
		} else {
			// query rds instance detail
			instanceDetailInfo, err := s.Client.UniversalClient.DoCall(getUniversalInfo("DescribeDBInstanceDetail"), &map[string]interface{}{
				"InstanceId": rdsInstance["InstanceId"],
			})
			if err != nil {
				logger.Info("DescribeDBInstanceDetail error:", err)
				continue
			}

			// 1. basic info
			basicInfo, err := ve.ObtainSdkValue("Result.BasicInfo", *instanceDetailInfo)
			if err != nil {
				logger.Info("ObtainSdkValue Result.BasicInfo error:", err)
				continue
			}
			if basicInfoMap, ok := basicInfo.(map[string]interface{}); ok {
				rdsInstance["VCpu"] = basicInfoMap["VCPU"]
				rdsInstance["Memory"] = basicInfoMap["Memory"]
				rdsInstance["UpdateTime"] = basicInfoMap["UpdateTime"]
				rdsInstance["BackupUse"] = basicInfoMap["BackupUse"]
				rdsInstance["DataSyncMode"] = basicInfoMap["DataSyncMode"]
			}

			// 2. endpoint info
			endpoints, err := ve.ObtainSdkValue("Result.Endpoints", *instanceDetailInfo)
			if err != nil {
				logger.Info("ObtainSdkValue Result.Endpoints error:", err)
				continue
			}
			for _, v1 := range endpoints.([]interface{}) {
				if endpoint, ok := v1.(map[string]interface{}); ok {
					endpoint["Addresses"] = convertAddressInfo(endpoint["Addresses"])
					endpoint["NodeWeight"] = endpoint["ReadOnlyNodeWeight"]
					delete(endpoint, "ReadOnlyNodeWeight")
				}
			}
			rdsInstance["Endpoints"] = endpoints

			// 3. node info
			nodes, err := ve.ObtainSdkValue("Result.Nodes", *instanceDetailInfo)
			if err != nil {
				logger.Info("ObtainSdkValue Result.Nodes error:", err)
				continue
			}
			for _, v2 := range nodes.([]interface{}) {
				if node, ok := v2.(map[string]interface{}); ok {
					node["VCpu"] = node["VCPU"]
					delete(node, "VCPU")
				}
			}
			rdsInstance["Nodes"] = nodes

			// query rds instance allow list ids
			allowListInfo, err := s.Client.UniversalClient.DoCall(getUniversalInfo("DescribeAllowLists"), &map[string]interface{}{
				"InstanceId": rdsInstance["InstanceId"],
				"RegionId":   s.Client.Region,
			})
			if err != nil {
				logger.Info("DescribeAllowLists error:", err)
				continue
			}

			allowLists, err := ve.ObtainSdkValue("Result.AllowLists", *allowListInfo)
			if err != nil {
				logger.Info("ObtainSdkValue Result.AllowLists error:", err)
				continue
			}
			if allowLists == nil {
				allowLists = []interface{}{}
			}
			allowListsArr, ok := allowLists.([]interface{})
			if !ok {
				logger.Info(" Result.AllowLists is not slice")
				continue
			}
			allowListIds := make([]interface{}, 0)
			for _, allowList := range allowListsArr {
				allowListMap, ok := allowList.(map[string]interface{})
				if !ok {
					logger.Info(" AllowList is not map")
					continue
				}
				allowListIds = append(allowListIds, allowListMap["AllowListId"])
			}
			rdsInstance["AllowListIds"] = allowListIds

			dbProxyConfig, err := s.Client.UniversalClient.DoCall(
				getUniversalInfo("DescribeDBProxyConfig"),
				&map[string]interface{}{
					"InstanceId": rdsInstance["InstanceId"],
				})
			if err != nil {
				logger.Info("DescribeDBProxyConfig error:", err)
				continue
			}
			proxyConfig, err := ve.ObtainSdkValue("Result", *dbProxyConfig)
			if err != nil {
				logger.Info("ObtainSdkValue Result error:", err)
				continue
			}
			proxyMap := proxyConfig.(map[string]interface{})
			rdsInstance["ConnectionPoolType"] = proxyMap["ConnectionPoolType"]
			rdsInstance["BinlogDump"] = proxyMap["BinlogDump"]
			rdsInstance["GlobalReadOnly"] = proxyMap["GlobalReadOnly"]
			rdsInstance["DBProxyStatus"] = proxyMap["DBProxyStatus"]
			rdsInstance["CheckModifyDBProxyAllowed"] = proxyMap["CheckModifyDBProxyAllowed"]
			rdsInstance["FeatureStates"] = proxyMap["FeatureStates"]
		}
	}

	return data, err
}

func (s *VolcengineRdsMysqlInstanceService) ReadResource(resourceData *schema.ResourceData, rdsInstanceId string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if rdsInstanceId == "" {
		rdsInstanceId = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"InstanceId": rdsInstanceId,
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		if data, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("Value is not map ")
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("Rds instance %s not exist ", rdsInstanceId)
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

	if parameterSet, ok := resourceData.GetOk("parameters"); ok {
		data["Parameters"] = parameterSet.(*schema.Set).List()
	}

	// DescribeDBInstances 不再返回 MaintenanceWindow 字段，需手动赋值为空数组
	if _, ok := data["MaintenanceWindow"]; !ok {
		if mainWindow, ok := resourceData.GetOk("maintenance_window"); ok {
			windowMap := mainWindow.([]interface{})[0].(map[string]interface{})
			maintenanceWindow := make(map[string]interface{})

			if time, ok := windowMap["maintenance_time"]; ok {
				maintenanceWindow["MaintenanceTime"] = time
			}
			if dayKind, ok := windowMap["day_kind"]; ok {
				maintenanceWindow["DayKind"] = dayKind
			}
			if weekDay, ok := windowMap["day_of_week"]; ok {
				maintenanceWindow["WeekDay"] = weekDay.(*schema.Set).List()
			}
			data["MaintenanceWindow"] = maintenanceWindow
		} else {
			data["MaintenanceWindow"] = []interface{}{}
		}
	}

	data["ChargeInfo"] = data["ChargeDetail"]

	return data, err
}

func (s *VolcengineRdsMysqlInstanceService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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
					return nil, "", fmt.Errorf("Rds instance status error, status:%s ", status.(string))
				}
			}
			//注意 返回的第一个参数不能为空 否则会一直等下去
			return demo, status.(string), err
		},
	}

}

func (*VolcengineRdsMysqlInstanceService) WithResourceResponseHandlers(rdsInstance map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return rdsInstance, map[string]ve.ResponseConvert{
			"DBEngineVersion": {
				TargetField: "db_engine_version",
			},
			"TimeZone": {
				TargetField: "db_time_zone",
			},
			"NodeCPUUsedPercentage": {
				TargetField: "node_cpu_used_percentage",
			},
			"NodeMemoryUsedPercentage": {
				TargetField: "node_memory_used_percentage",
			},
			"NodeSpaceUsedPercentage": {
				TargetField: "node_space_used_percentage",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}

}

func (s *VolcengineRdsMysqlInstanceService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	instanceCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateDBInstance",
			ContentType: ve.ContentTypeJson,
			ConvertMode: ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"db_engine_version": {
					TargetField: "DBEngineVersion",
				},
				"storage_space": {
					TargetField: "StorageSpace",
				},
				"subnet_id": {
					TargetField: "SubnetId",
				},
				"instance_name": {
					TargetField: "InstanceName",
				},
				"lower_case_table_names": {
					TargetField: "LowerCaseTableNames",
				},
				"db_time_zone": {
					TargetField: "DBTimeZone",
				},
				"charge_info": {
					ConvertType: ve.ConvertJsonObject,
				},
				"project_name": {
					TargetField: "ProjectName",
				},
				"tags": {
					TargetField: "InstanceTags",
					ConvertType: ve.ConvertJsonObjectArray,
				},
				"maintenance_window": {
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

				// 1. NodeInfo
				primaryNodeInfo := make(map[string]interface{})
				primaryNodeInfo["NodeType"] = "Primary"
				primaryNodeInfo["ZoneId"] = d.Get("primary_zone_id")
				primaryNodeInfo["NodeSpec"] = d.Get("node_spec")
				nodeInfos = append(nodeInfos, primaryNodeInfo)

				secondaryNodeInfo := make(map[string]interface{})
				secondaryNodeInfo["NodeType"] = "Secondary"
				secondaryNodeInfo["ZoneId"] = d.Get("secondary_zone_id")
				secondaryNodeInfo["NodeSpec"] = d.Get("node_spec")
				nodeInfos = append(nodeInfos, secondaryNodeInfo)

				// 2. VpcId
				subnetId := d.Get("subnet_id")
				req := map[string]interface{}{
					"SubnetIds.1": subnetId,
				}
				action := "DescribeSubnets"
				resp, err := s.Client.UniversalClient.DoCall(getVpcUniversalInfo(action), &req)
				if err != nil {
					return false, err
				}
				logger.Debug(logger.RespFormat, action, req, *resp)
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

				(*call.SdkParam)["NodeInfo"] = nodeInfos
				(*call.SdkParam)["StorageType"] = "LocalSSD"
				(*call.SdkParam)["VpcId"] = vpcId

				if mainWindow, ok := d.GetOk("maintenance_window"); ok {
					windowMap := mainWindow.([]interface{})[0].(map[string]interface{})
					maintenanceWindow := make(map[string]interface{})

					if time, ok := windowMap["maintenance_time"]; ok {
						maintenanceWindow["MaintenanceTime"] = time
					}
					if dayKind, ok := windowMap["day_kind"]; ok {
						maintenanceWindow["DayKind"] = dayKind
					}
					if weekDay, ok := windowMap["day_of_week"]; ok {
						maintenanceWindow["WeekDay"] = weekDay.(*schema.Set).List()
					}

					(*call.SdkParam)["MaintenanceWindow"] = maintenanceWindow
				}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				//创建rdsInstance
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				//注意 获取内容 这个地方不能是指针 需要转一次
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
	callbacks = append(callbacks, instanceCallback)

	// 关联白名单
	allowListCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "AssociateAllowList",
			ContentType: ve.ContentTypeJson,
			ConvertMode: ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"allow_list_ids": {
					ConvertType: ve.ConvertJsonArray,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if len(*call.SdkParam) > 0 {
					(*call.SdkParam)["InstanceIds"] = []string{d.Id()}
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
	callbacks = append(callbacks, allowListCallback)

	// 关联参数
	parameterCallback := ve.Callback{
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
	callbacks = append(callbacks, parameterCallback)

	if connectionPool, ok := resourceData.GetOk("connection_pool_type"); ok {
		connectionPoolCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyDBProxyConfig",
				ConvertMode: ve.RequestConvertIgnore,
				ContentType: ve.ContentTypeJson,
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					if len(*call.SdkParam) > 0 {
						(*call.SdkParam)["InstanceId"] = d.Id()
						(*call.SdkParam)["ConnectionPoolType"] = connectionPool
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
		callbacks = append(callbacks, connectionPoolCallback)
	}

	return callbacks
}

func (s *VolcengineRdsMysqlInstanceService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	// 1. NodeSpec & StorageSpace
	if resourceData.HasChanges("node_spec", "storage_space") {
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

					return true, nil
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
		callbacks = append(callbacks, instanceCallback)
	}

	// InstanceName
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

	// AllowList
	if resourceData.HasChange("allow_list_ids") {
		addAlIds, removeAlIds, _, _ := ve.GetSetDifference("allow_list_ids", resourceData, schema.HashString, false)

		allowListRemoveCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "DisassociateAllowList",
				ContentType: ve.ContentTypeJson,
				ConvertMode: ve.RequestConvertIgnore,
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					if removeAlIds != nil && len(removeAlIds.List()) > 0 {
						(*call.SdkParam)["InstanceIds"] = []string{d.Id()}
						(*call.SdkParam)["AllowListIds"] = removeAlIds.List()
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
		callbacks = append(callbacks, allowListRemoveCallback)

		allowListAddCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "AssociateAllowList",
				ContentType: ve.ContentTypeJson,
				ConvertMode: ve.RequestConvertIgnore,
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					if addAlIds != nil && len(addAlIds.List()) > 0 {
						(*call.SdkParam)["InstanceIds"] = []string{d.Id()}
						(*call.SdkParam)["AllowListIds"] = addAlIds.List()
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
		callbacks = append(callbacks, allowListAddCallback)
	}

	// Parameters
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
								return false, fmt.Errorf("Parameter is not map ")
							}
							(*call.SdkParam)["Parameters"] = append((*call.SdkParam)["Parameters"].([]map[string]interface{}), map[string]interface{}{
								"ParameterName":  paramMap["parameter_name"],
								"ParameterValue": paramMap["parameter_value"],
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

	// 更新Tags
	callbacks = s.setResourceTags(resourceData, callbacks)

	// MaintenanceWindow
	if resourceData.HasChange("maintenance_window") {
		maintenanceWindowCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyDBInstanceMaintenanceWindow",
				ContentType: ve.ContentTypeJson,
				ConvertMode: ve.RequestConvertIgnore,
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["InstanceId"] = d.Id()
					if mainWindow, ok := d.GetOk("maintenance_window"); ok {
						windowMap := mainWindow.([]interface{})[0].(map[string]interface{})
						maintenanceWindow := make(map[string]interface{})

						if time, ok := windowMap["maintenance_time"]; ok {
							maintenanceWindow["MaintenanceTime"] = time
						}
						if dayKind, ok := windowMap["day_kind"]; ok {
							maintenanceWindow["DayKind"] = dayKind
						}
						if weekDay, ok := windowMap["day_of_week"]; ok {
							maintenanceWindow["WeekDay"] = weekDay.(*schema.Set).List()
						}

						(*call.SdkParam)["MaintenanceWindow"] = maintenanceWindow
					}
					return true, nil
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
		callbacks = append(callbacks, maintenanceWindowCallback)
	}

	if resourceData.HasChange("connection_pool_type") {
		connectionPoolCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyDBProxyConfig",
				ConvertMode: ve.RequestConvertIgnore,
				ContentType: ve.ContentTypeJson,
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					if len(*call.SdkParam) > 0 {
						(*call.SdkParam)["InstanceId"] = d.Id()
						(*call.SdkParam)["ConnectionPoolType"] = d.Get("connection_pool_type")
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
		callbacks = append(callbacks, connectionPoolCallback)
	}

	return callbacks
}

func (s *VolcengineRdsMysqlInstanceService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	// 删除包年包月实例时报错
	if chargeType := resourceData.Get("charge_info.0.charge_type"); chargeType.(string) == "PrePaid" {
		return []ve.Callback{{
			Err: fmt.Errorf("can not delete PrePaid mysql instance"),
		}}
	}

	// 1. Disassociate Allow List
	allowListCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DisassociateAllowList",
			ContentType: ve.ContentTypeJson,
			ConvertMode: ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"allow_list_ids": {
					ConvertType: ve.ConvertJsonArray,
					ForceGet:    true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if len(*call.SdkParam) > 0 {
					(*call.SdkParam)["InstanceIds"] = []string{d.Id()}
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
	callbacks = append(callbacks, allowListCallback)

	// 2. delete instance
	removeCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteDBInstance",
			ContentType: ve.ContentTypeJson,
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"InstanceId": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				//删除RdsInstance
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading rds mysql instance on delete %q, %w", d.Id(), callErr))
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
	callbacks = append(callbacks, removeCallback)

	return callbacks
}

func (s *VolcengineRdsMysqlInstanceService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
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
		CollectField: "rds_mysql_instances",
		ContentType:  ve.ContentTypeJson,
		ResponseConverts: map[string]ve.ResponseConvert{
			"InstanceId": {
				TargetField: "id",
				KeepDefault: true,
			},
			"DBEngineVersion": {
				TargetField: "db_engine_version",
			},
			"NodeCPUUsedPercentage": {
				TargetField: "node_cpu_used_percentage",
			},
			"NodeMemoryUsedPercentage": {
				TargetField: "node_memory_used_percentage",
			},
			"NodeSpaceUsedPercentage": {
				TargetField: "node_space_used_percentage",
			},
			"DBProxyStatus": {
				TargetField: "db_proxy_status",
			},
			//"CheckModifyDBProxyAllowed": {
			//	TargetField: "check_modify_db_proxy_allowed",
			//},
		},
	}
}

func convertAddressInfo(addressesInfo interface{}) interface{} {
	if addressesInfo == nil {
		return nil
	}
	var addresses []interface{}
	if addressInfoArr, ok := addressesInfo.([]interface{}); ok {
		for _, address := range addressInfoArr {
			if addressMap, ok := address.(map[string]interface{}); ok {
				addressMap["IpAddress"] = addressMap["IPAddress"]
				addressMap["DnsVisibility"] = addressMap["DNSVisibility"]
				delete(addressMap, "IPAddress")
				delete(addressMap, "DNSVisibility")
			}
			addresses = append(addresses, address)
		}
	}

	return addresses
}

func (s *VolcengineRdsMysqlInstanceService) ReadResourceId(id string) string {
	return id
}

func (s *VolcengineRdsMysqlInstanceService) setResourceTags(resourceData *schema.ResourceData, callbacks []ve.Callback) []ve.Callback {
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

func (s *VolcengineRdsMysqlInstanceService) ProjectTrn() *ve.ProjectTrn {
	return &ve.ProjectTrn{
		ServiceName:          "rds_mysql",
		ResourceType:         "instance",
		ProjectResponseField: "ProjectName",
		ProjectSchemaField:   "project_name",
	}
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "rds_mysql",
		Version:     "2022-01-01",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}

func getVpcUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "vpc",
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		Action:      actionName,
	}
}
