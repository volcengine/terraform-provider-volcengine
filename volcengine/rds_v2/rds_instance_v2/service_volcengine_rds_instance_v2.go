package rds_instance_v2

import (
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	volc "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineRdsInstanceService struct {
	Client     *volc.SdkClient
	Dispatcher *volc.Dispatcher
}

func NewRdsInstanceService(c *volc.SdkClient) *VolcengineRdsInstanceService {
	return &VolcengineRdsInstanceService{
		Client:     c,
		Dispatcher: &volc.Dispatcher{},
	}
}

func (s *VolcengineRdsInstanceService) GetClient() *volc.SdkClient {
	return s.Client
}

func (s *VolcengineRdsInstanceService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp        *map[string]interface{}
		results     interface{}
		ok          bool
		rdsInstance map[string]interface{}
	)
	data, err = volc.WithPageNumberQuery(m, "PageSize", "PageNumber", 20, 1, func(condition map[string]interface{}) ([]interface{}, error) {
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

		results, err = volc.ObtainSdkValue("Result.InstancesInfo", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.Datas is not Slice")
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
			// query rds connection info
			instanceDetailInfo, err := s.Client.UniversalClient.DoCall(getUniversalInfo("DescribeDBInstanceDetail"), &map[string]interface{}{
				"InstanceId": rdsInstance["InstanceId"],
			})
			if err != nil {
				logger.Info("DescribeDBInstanceDetail error:", err)
				continue
			}

			// 1. node info
			nodeDetailInfo, err := volc.ObtainSdkValue("Result.NodeDetailInfo", *instanceDetailInfo)
			if err != nil {
				logger.Info("ObtainSdkValue Result.NodeDetailInfo error:", err)
				continue
			}
			rdsInstance["NodeDetailInfo"] = nodeDetailInfo
			rdsInstance["NodeInfo"] = nodeDetailInfo

			// 2. connection info
			connectionInfo, err := volc.ObtainSdkValue("Result.ConnectionInfo", *instanceDetailInfo)
			if err != nil {
				return data, err
			}

			rdsInstance["ConnectionInfo"] = convertConnectionInfo(connectionInfo)
		}
	}

	return data, err
}

func convertConnectionInfo(connectionInfo interface{}) interface{} {
	if connectionInfo == nil {
		return nil
	}
	if connectionInfoArr, ok := connectionInfo.([]interface{}); ok {
		for _, v := range connectionInfoArr {
			if vv, ok := v.(map[string]interface{}); ok {
				addresses := vv["Address"].([]interface{})
				for _, address := range addresses {
					if addressMap, ok := address.(map[string]interface{}); ok {
						addressMap["IpAddress"] = addressMap["IPAddress"]
					}
				}
			}
		}
	}

	return connectionInfo
}

func (s *VolcengineRdsInstanceService) ReadResource(resourceData *schema.ResourceData, rdsInstanceId string) (data map[string]interface{}, err error) {
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

	return data, err
}

func (s *VolcengineRdsInstanceService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
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
			demo, err = s.ReadResource(resourceData, id)
			if err != nil {
				return nil, "", err
			}
			status, err = volc.ObtainSdkValue("InstanceStatus", demo)
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

func (*VolcengineRdsInstanceService) WithResourceResponseHandlers(rdsInstance map[string]interface{}) []volc.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]volc.ResponseConvert, error) {
		return rdsInstance, map[string]volc.ResponseConvert{
			"InstanceId": {
				TargetField: "id",
				KeepDefault: true,
			},
			"DBEngine": {
				TargetField: "db_engine",
			},
			"DBEngineVersion": {
				TargetField: "db_engine_version",
			},
			"ChargeDetail": {
				TargetField: "charge_info",
			},
		}, nil
	}
	return []volc.ResourceResponseHandler{handler}

}

func (s *VolcengineRdsInstanceService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []volc.Callback {
	callback := volc.Callback{
		Call: volc.SdkCall{
			Action:      "CreateDBInstance",
			ContentType: volc.ContentTypeJson,
			ConvertMode: volc.RequestConvertAll,
			Convert: map[string]volc.RequestConvert{
				"db_engine_version": {
					TargetField: "DBEngineVersion",
				},
				"db_time_zone": {
					TargetField: "DBTimeZone",
				},
				"db_param_group_id": {
					TargetField: "DBParamGroupId",
				},
				"charge_info": {
					ConvertType: volc.ConvertJsonObject,
				},
				"node_info": {
					ConvertType: volc.ConvertJsonObjectArray,
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				//创建rdsInstance
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *volc.SdkClient, resp *map[string]interface{}, call volc.SdkCall) error {
				//注意 获取内容 这个地方不能是指针 需要转一次
				id, _ := volc.ObtainSdkValue("Result.InstanceId", *resp)
				d.SetId(id.(string))
				return nil
			},
			Refresh: &volc.StateRefresh{
				Target:  []string{"Running"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	return []volc.Callback{callback}

}

func (s *VolcengineRdsInstanceService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []volc.Callback {
	if !resourceData.HasChanges("storage_type", "storage_space", "node_info") {
		return []volc.Callback{}
	}

	targetNodeInfo := make([]map[string]interface{}, 0)

	changePrimary := false
	newPrimaryNodeId := ""
	hasSecondary := false

	hasNodeChange := false

	if resourceData.HasChange("node_info") {
		oldNodes, newNodes := resourceData.GetChange("node_info")
		logger.Info("oldNodes:%v", oldNodes)
		logger.Info("newNodes:%v", newNodes)

		oldNodeMap := make(map[string]map[string]interface{})
		newNodeMap := make(map[string]map[string]interface{})

		oldNodeList := oldNodes.([]interface{})
		for _, v := range oldNodeList {
			node := v.(map[string]interface{})
			oldNodeMap[node["node_id"].(string)] = node
		}

		newNodeList := newNodes.([]interface{})
		for _, v := range newNodeList {
			node := v.(map[string]interface{})
			if node["node_id"] == nil {
				// new node
				continue
			}
			newNodeMap[node["node_id"].(string)] = node
		}

		for _, v := range newNodeList {
			// exist, create, modify
			node := v.(map[string]interface{})
			if node["node_id"] == nil || node["node_id"].(string) == "" {
				// new node
				hasNodeChange = true
				targetNodeInfo = append(targetNodeInfo, map[string]interface{}{
					"ZoneId":          node["zone_id"],
					"NodeSpec":        node["node_spec"],
					"NodeType":        node["node_type"],
					"NodeOperateType": "Create",
				})
				continue
			}

			if oldNode, ok := oldNodeMap[node["node_id"].(string)]; ok {
				// exist or modify
				oldNodeType := oldNode["node_type"].(string)
				oldNodeSpec := oldNode["node_spec"].(string)

				newNodeType := node["node_type"].(string)
				newNodeSpec := node["node_spec"].(string)

				if newNodeType == "Secondary" {
					hasSecondary = true
				}

				if oldNodeType != newNodeType || oldNodeSpec != newNodeSpec {
					// modify
					if oldNodeType == "Primary" || oldNodeType == "Secondary" {
						// 接口不支持主备切换，仅支持配置变更
						if oldNodeSpec != newNodeSpec {
							// 仅变更规格
							hasNodeChange = true
							targetNodeInfo = append(targetNodeInfo, map[string]interface{}{
								"NodeId":          node["node_id"],
								"ZoneId":          node["zone_id"],
								"NodeSpec":        newNodeSpec,
								"NodeType":        oldNodeType,
								"NodeOperateType": "Modify",
							})
						} else {
							// 节点不做变更
							targetNodeInfo = append(targetNodeInfo, map[string]interface{}{
								"NodeId":   node["node_id"],
								"ZoneId":   node["zone_id"],
								"NodeSpec": oldNodeSpec,
								"NodeType": oldNodeType,
							})
						}

						if oldNodeType != newNodeType {
							changePrimary = true
							if newNodeType == "Primary" {
								newPrimaryNodeId = node["node_id"].(string)
							}
						}
					} else {
						// 普通节点变更
						hasNodeChange = true
						targetNodeInfo = append(targetNodeInfo, map[string]interface{}{
							"NodeId":          node["node_id"],
							"ZoneId":          node["zone_id"],
							"NodeSpec":        newNodeSpec,
							"NodeType":        newNodeType,
							"NodeOperateType": "Modify",
						})
					}
				} else {
					// exist
					targetNodeInfo = append(targetNodeInfo, map[string]interface{}{
						"NodeId":   node["node_id"],
						"ZoneId":   node["zone_id"],
						"NodeSpec": newNodeSpec,
						"NodeType": newNodeType,
					})
				}
			}
		}

		for _, v := range oldNodeList {
			// delete
			node := v.(map[string]interface{})
			if _, ok := newNodeMap[node["node_id"].(string)]; !ok {
				hasNodeChange = true
				targetNodeInfo = append(targetNodeInfo, map[string]interface{}{
					"NodeId":          node["node_id"],
					"ZoneId":          node["zone_id"],
					"NodeSpec":        node["node_spec"],
					"NodeType":        node["node_type"],
					"NodeOperateType": "Delete",
				})
			}
		}

		logger.Info("targetNodeInfo:%v", targetNodeInfo)
	} else {
		// node info没改，也要按照exist传入
		nodeInfo := resourceData.Get("node_info").([]interface{})
		for _, v := range nodeInfo {
			nodeMap := v.(map[string]interface{})
			targetNodeInfo = append(targetNodeInfo, map[string]interface{}{
				"NodeId":   nodeMap["node_id"],
				"ZoneId":   nodeMap["zone_id"],
				"NodeSpec": nodeMap["node_spec"],
				"NodeType": nodeMap["node_type"],
			})
		}
	}

	callbacks := make([]volc.Callback, 0)
	if hasNodeChange || resourceData.HasChanges("storage_space", "storage_type") {
		modifySpecCallback := volc.Callback{
			Call: volc.SdkCall{
				Action:      "ModifyDBInstanceSpec",
				ContentType: volc.ContentTypeJson,
				ConvertMode: volc.RequestConvertIgnore,
				SdkParam: &map[string]interface{}{
					"InstanceId": resourceData.Id(),
				},
				BeforeCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (bool, error) {
					(*call.SdkParam)["StorageType"] = d.Get("storage_type")
					if d.HasChange("storage_space") {
						(*call.SdkParam)["StorageSpace"] = d.Get("storage_space")
					}
					(*call.SdkParam)["NodeInfo"] = targetNodeInfo
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
				Refresh: &volc.StateRefresh{
					Target:  []string{"Running"},
					Timeout: resourceData.Timeout(schema.TimeoutCreate),
				},
			},
		}
		callbacks = append(callbacks, modifySpecCallback)
	}

	if changePrimary {
		changePrimaryCallback := volc.Callback{
			Call: volc.SdkCall{
				Action:      "ChangeDBInstanceHAMaster",
				ContentType: volc.ContentTypeJson,
				ConvertMode: volc.RequestConvertIgnore,
				SdkParam: &map[string]interface{}{
					"InstanceId": resourceData.Id(),
				},
				BeforeCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (bool, error) {
					if newPrimaryNodeId == "" {
						return false, errors.New("non primary node")
					}
					if !hasSecondary {
						return false, errors.New("non secondary node")
					}
					(*call.SdkParam)["NodeId"] = newPrimaryNodeId
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getV1UniversalInfo(call.Action), call.SdkParam)
				},
				Refresh: &volc.StateRefresh{
					Target:  []string{"Running"},
					Timeout: resourceData.Timeout(schema.TimeoutCreate),
				},
			},
		}
		callbacks = append(callbacks, changePrimaryCallback)
	}

	return callbacks
}

func (s *VolcengineRdsInstanceService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []volc.Callback {
	callback := volc.Callback{
		Call: volc.SdkCall{
			Action:      "DeleteDBInstance",
			ContentType: volc.ContentTypeJson,
			ConvertMode: volc.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"InstanceId": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				//删除RdsInstance
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			CallError: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if volc.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading rds instance on delete %q, %w", d.Id(), callErr))
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
	return []volc.Callback{callback}
}

func (s *VolcengineRdsInstanceService) DatasourceResources(*schema.ResourceData, *schema.Resource) volc.DataSourceInfo {
	return volc.DataSourceInfo{
		ContentType: volc.ContentTypeJson,
		RequestConverts: map[string]volc.RequestConvert{
			"db_engine_version": {
				TargetField: "DBEngineVersion",
			},
		},
		NameField:    "InstanceName",
		IdField:      "InstanceId",
		CollectField: "rds_instances",
		ResponseConverts: map[string]volc.ResponseConvert{
			"InstanceId": {
				TargetField: "id",
				KeepDefault: true,
			},
			"DBEngine": {
				TargetField: "db_engine",
			},
			"DBEngineVersion": {
				TargetField: "db_engine_version",
			},
		},
	}
}

func (s *VolcengineRdsInstanceService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) volc.UniversalInfo {
	return volc.UniversalInfo{
		ServiceName: "rds_mysql",
		Version:     "2022-01-01",
		HttpMethod:  volc.POST,
		ContentType: volc.ApplicationJSON,
		Action:      actionName,
	}
}

func getV1UniversalInfo(actionName string) volc.UniversalInfo {
	return volc.UniversalInfo{
		ServiceName: "rds_mysql",
		Version:     "2018-01-01",
		HttpMethod:  volc.POST,
		ContentType: volc.ApplicationJSON,
		Action:      actionName,
	}
}
