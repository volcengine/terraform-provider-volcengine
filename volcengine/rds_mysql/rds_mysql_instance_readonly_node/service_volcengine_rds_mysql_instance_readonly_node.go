package rds_mysql_instance_readonly_node

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/rds_mysql/rds_mysql_instance"
)

type VolcengineRdsMysqlInstanceReadonlyNodeService struct {
	Client             *ve.SdkClient
	Dispatcher         *ve.Dispatcher
	RdsInstanceService *rds_mysql_instance.VolcengineRdsMysqlInstanceService
}

func NewRdsMysqlInstanceReadonlyNodeService(c *ve.SdkClient) *VolcengineRdsMysqlInstanceReadonlyNodeService {
	return &VolcengineRdsMysqlInstanceReadonlyNodeService{
		Client:             c,
		Dispatcher:         &ve.Dispatcher{},
		RdsInstanceService: rds_mysql_instance.NewRdsMysqlInstanceService(c),
	}
}

func (s *VolcengineRdsMysqlInstanceReadonlyNodeService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineRdsMysqlInstanceReadonlyNodeService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	return data, err
}

func (s *VolcengineRdsMysqlInstanceReadonlyNodeService) ReadResource(resourceData *schema.ResourceData, rdsInstanceNodeId string) (data map[string]interface{}, err error) {
	if rdsInstanceNodeId == "" {
		rdsInstanceNodeId = resourceData.Id()
	}

	ids := strings.Split(rdsInstanceNodeId, ":")
	if len(ids) != 2 {
		return map[string]interface{}{}, fmt.Errorf("invalid rdsInstanceNodeId: %s", rdsInstanceNodeId)
	}

	instanceId := ids[0]
	nodeId := ids[1]

	result, err := s.RdsInstanceService.ReadResource(resourceData, instanceId)
	if err != nil {
		return result, err
	}
	if len(result) == 0 {
		return result, fmt.Errorf("Rds instance %s not exist ", instanceId)
	}

	if nodeArr, ok := result["Nodes"].([]interface{}); ok {
		for _, node := range nodeArr {
			if nodeMap, ok1 := node.(map[string]interface{}); ok1 {
				if nodeMap["NodeId"] == nodeId {
					data = nodeMap
				}
			}
		}
	}
	data["NodeId"] = nodeId

	return data, err
}

func (s *VolcengineRdsMysqlInstanceReadonlyNodeService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (*VolcengineRdsMysqlInstanceReadonlyNodeService) WithResourceResponseHandlers(rdsInstance map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return rdsInstance, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}

}

func (s *VolcengineRdsMysqlInstanceReadonlyNodeService) CreateResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	var (
		callbacks               []ve.Callback
		existingReadOnlyNodeIds = make(map[string]bool)
	)

	nodeCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyDBInstanceSpec",
			ContentType: ve.ContentTypeJson,
			ConvertMode: ve.RequestConvertIgnore,
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (common *map[string]interface{}, err error) {
				// 在LockId执行后再进行已有Node信息的查询
				(*call.SdkParam)["InstanceId"] = d.Get("instance_id").(string)

				nodeInfos := make([]interface{}, 0)
				// 1. 获取当前RdsInstance已有的Node信息
				instance, err := s.RdsInstanceService.ReadResource(resourceData, d.Get("instance_id").(string))
				if err != nil {
					return common, err
				}
				if nodeArr, ok := instance["Nodes"].([]interface{}); ok {
					for _, node := range nodeArr {
						if nodeMap, ok1 := node.(map[string]interface{}); ok1 {
							if nodeMap["NodeType"] == "Primary" {
								primaryNodeInfo := make(map[string]interface{})
								primaryNodeInfo["NodeId"] = nodeMap["NodeId"]
								primaryNodeInfo["NodeType"] = nodeMap["NodeType"]
								primaryNodeInfo["NodeSpec"] = nodeMap["NodeSpec"]
								primaryNodeInfo["ZoneId"] = nodeMap["ZoneId"]
								nodeInfos = append(nodeInfos, primaryNodeInfo)
							} else if nodeMap["NodeType"] == "Secondary" {
								secondaryNodeInfo := make(map[string]interface{})
								secondaryNodeInfo["NodeId"] = nodeMap["NodeId"]
								secondaryNodeInfo["NodeType"] = nodeMap["NodeType"]
								secondaryNodeInfo["NodeSpec"] = nodeMap["NodeSpec"]
								secondaryNodeInfo["ZoneId"] = nodeMap["ZoneId"]
								nodeInfos = append(nodeInfos, secondaryNodeInfo)
							} else if nodeMap["NodeType"] == "ReadOnly" {
								readonlyNodeInfo := make(map[string]interface{})
								readonlyNodeInfo["NodeId"] = nodeMap["NodeId"]
								readonlyNodeInfo["NodeType"] = nodeMap["NodeType"]
								readonlyNodeInfo["NodeSpec"] = nodeMap["NodeSpec"]
								readonlyNodeInfo["ZoneId"] = nodeMap["ZoneId"]
								nodeInfos = append(nodeInfos, readonlyNodeInfo)

								existingReadOnlyNodeIds[readonlyNodeInfo["NodeId"].(string)] = true
							}
						}
					}
				}

				// 2. 新增 readonly node
				newReadonlyNodeInfo := make(map[string]interface{})
				newReadonlyNodeInfo["NodeType"] = "ReadOnly"
				newReadonlyNodeInfo["NodeSpec"] = d.Get("node_spec")
				newReadonlyNodeInfo["ZoneId"] = d.Get("zone_id")
				newReadonlyNodeInfo["NodeOperateType"] = "Create"
				nodeInfos = append(nodeInfos, newReadonlyNodeInfo)

				(*call.SdkParam)["NodeInfo"] = nodeInfos

				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				common, err = s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				if err != nil {
					return common, err
				}
				time.Sleep(10 * time.Second) // 新增只读节点后，需要等一下
				return common, nil
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				var (
					instance = make(map[string]interface{})
					err      error
				)
				// 通过retry确保获取当前新建只读节点的Id
				resource.Retry(15*time.Minute, func() *resource.RetryError {
					instance, err = s.RdsInstanceService.ReadResource(d, d.Get("instance_id").(string))
					if err != nil {
						if ve.ResourceNotFoundError(err) {
							return resource.RetryableError(err)
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading rds instance %q", d.Get("instance_id")))
						}
					}
					status, err := ve.ObtainSdkValue("InstanceStatus", instance)
					if err != nil {
						return resource.RetryableError(err)
					}
					if status.(string) != "Running" {
						return resource.RetryableError(fmt.Errorf("rds instance is still in updating"))
					}
					return nil
				})
				logger.Debug(logger.ReqFormat, "testReadonly", instance)
				var newReadonlyNodeId string
				if nodeArr, ok := instance["Nodes"].([]interface{}); ok {
					for _, node := range nodeArr {
						if nodeMap, ok1 := node.(map[string]interface{}); ok1 {
							logger.Debug(logger.ReqFormat, "existingReadOnlyNodeIds", existingReadOnlyNodeIds)
							if nodeMap["NodeType"] == "ReadOnly" {
								if _, ok := existingReadOnlyNodeIds[nodeMap["NodeId"].(string)]; !ok {
									newReadonlyNodeId = nodeMap["NodeId"].(string)
								}
							}
						}
					}
				}
				// ResourceData中，rds_mysql_instance_readonly_node的Id形式为'instance_id:node_id'
				logger.Debug(logger.ReqFormat, "newReadonlyNodeId", newReadonlyNodeId)
				id := fmt.Sprintf("%s:%s", d.Get("instance_id"), newReadonlyNodeId)
				d.SetId(id)
				return nil
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("instance_id").(string)
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				s.RdsInstanceService: {
					Target:     []string{"Running"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					ResourceId: resourceData.Get("instance_id").(string),
				},
			},
		},
	}
	callbacks = append(callbacks, nodeCallback)

	return callbacks
}

func (s *VolcengineRdsMysqlInstanceReadonlyNodeService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	if resourceData.HasChange("node_spec") {
		nodeCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyDBInstanceSpec",
				ContentType: ve.ContentTypeJson,
				ConvertMode: ve.RequestConvertIgnore,
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (common *map[string]interface{}, err error) {
					// 在 LockId 后再进行已有 Node 信息的查询
					ids := strings.Split(d.Id(), ":")
					if len(ids) != 2 {
						return common, fmt.Errorf("invalid rdsInstanceNodeId: %s", d.Id())
					}
					instanceId := ids[0]
					nodeId := ids[1]
					(*call.SdkParam)["InstanceId"] = instanceId

					nodeInfos := make([]interface{}, 0)
					// 1. 获取当前RdsInstance已有的Node信息
					instance, err := s.RdsInstanceService.ReadResource(resourceData, d.Get("instance_id").(string))
					if err != nil {
						return common, err
					}
					if nodeArr, ok := instance["Nodes"].([]interface{}); ok {
						for _, node := range nodeArr {
							if nodeMap, ok1 := node.(map[string]interface{}); ok1 {
								if nodeMap["NodeType"] == "Primary" {
									primaryNodeInfo := make(map[string]interface{})
									primaryNodeInfo["NodeId"] = nodeMap["NodeId"]
									primaryNodeInfo["NodeType"] = nodeMap["NodeType"]
									primaryNodeInfo["NodeSpec"] = nodeMap["NodeSpec"]
									primaryNodeInfo["ZoneId"] = nodeMap["ZoneId"]
									nodeInfos = append(nodeInfos, primaryNodeInfo)
								} else if nodeMap["NodeType"] == "Secondary" {
									secondaryNodeInfo := make(map[string]interface{})
									secondaryNodeInfo["NodeId"] = nodeMap["NodeId"]
									secondaryNodeInfo["NodeType"] = nodeMap["NodeType"]
									secondaryNodeInfo["NodeSpec"] = nodeMap["NodeSpec"]
									secondaryNodeInfo["ZoneId"] = nodeMap["ZoneId"]
									nodeInfos = append(nodeInfos, secondaryNodeInfo)
								} else if nodeMap["NodeType"] == "ReadOnly" && nodeMap["NodeId"] != nodeId {
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

					// 2. 修改当前 readonly node
					newReadonlyNodeInfo := make(map[string]interface{})
					newReadonlyNodeInfo["NodeId"] = nodeId
					newReadonlyNodeInfo["NodeType"] = "ReadOnly"
					newReadonlyNodeInfo["NodeSpec"] = d.Get("node_spec")
					newReadonlyNodeInfo["ZoneId"] = d.Get("zone_id")
					newReadonlyNodeInfo["NodeOperateType"] = "Modify"
					nodeInfos = append(nodeInfos, newReadonlyNodeInfo)

					(*call.SdkParam)["NodeInfo"] = nodeInfos

					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					common, err = s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
					if err != nil {
						return common, err
					}
					time.Sleep(10 * time.Second) // 修改只读节点后，需要等一下
					return common, nil
				},
				LockId: func(d *schema.ResourceData) string {
					return d.Get("instance_id").(string)
				},
				ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
					s.RdsInstanceService: {
						Target:     []string{"Running"},
						Timeout:    resourceData.Timeout(schema.TimeoutUpdate),
						ResourceId: resourceData.Get("instance_id").(string),
					},
				},
			},
		}
		callbacks = append(callbacks, nodeCallback)
	}

	return callbacks
}

func (s *VolcengineRdsMysqlInstanceReadonlyNodeService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	nodeCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyDBInstanceSpec",
			ContentType: ve.ContentTypeJson,
			ConvertMode: ve.RequestConvertIgnore,
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (common *map[string]interface{}, err error) {
				// 在 LockId 后再进行已有 Node 信息的查询
				ids := strings.Split(d.Id(), ":")
				if len(ids) != 2 {
					return common, fmt.Errorf("invalid rdsInstanceNodeId: %s", d.Id())
				}
				instanceId := ids[0]
				nodeId := ids[1]
				(*call.SdkParam)["InstanceId"] = instanceId

				nodeInfos := make([]interface{}, 0)
				// 1. 获取当前RdsInstance已有的Node信息
				instance, err := s.RdsInstanceService.ReadResource(resourceData, d.Get("instance_id").(string))
				if err != nil {
					return common, err
				}
				if nodeArr, ok := instance["Nodes"].([]interface{}); ok {
					for _, node := range nodeArr {
						if nodeMap, ok1 := node.(map[string]interface{}); ok1 {
							if nodeMap["NodeType"] == "Primary" {
								primaryNodeInfo := make(map[string]interface{})
								primaryNodeInfo["NodeId"] = nodeMap["NodeId"]
								primaryNodeInfo["NodeType"] = nodeMap["NodeType"]
								primaryNodeInfo["NodeSpec"] = nodeMap["NodeSpec"]
								primaryNodeInfo["ZoneId"] = nodeMap["ZoneId"]
								nodeInfos = append(nodeInfos, primaryNodeInfo)
							} else if nodeMap["NodeType"] == "Secondary" {
								secondaryNodeInfo := make(map[string]interface{})
								secondaryNodeInfo["NodeId"] = nodeMap["NodeId"]
								secondaryNodeInfo["NodeType"] = nodeMap["NodeType"]
								secondaryNodeInfo["NodeSpec"] = nodeMap["NodeSpec"]
								secondaryNodeInfo["ZoneId"] = nodeMap["ZoneId"]
								nodeInfos = append(nodeInfos, secondaryNodeInfo)
							} else if nodeMap["NodeType"] == "ReadOnly" && nodeMap["NodeId"] != nodeId {
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

				// 2. 删除 readonly node
				newReadonlyNodeInfo := make(map[string]interface{})
				newReadonlyNodeInfo["NodeId"] = nodeId
				newReadonlyNodeInfo["NodeType"] = "ReadOnly"
				newReadonlyNodeInfo["NodeSpec"] = d.Get("node_spec")
				newReadonlyNodeInfo["ZoneId"] = d.Get("zone_id")
				newReadonlyNodeInfo["NodeOperateType"] = "Delete"
				nodeInfos = append(nodeInfos, newReadonlyNodeInfo)

				(*call.SdkParam)["NodeInfo"] = nodeInfos

				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				common, err = s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				if err != nil {
					return common, err
				}
				time.Sleep(10 * time.Second) // 删除只读节点后，需要等一下
				return common, nil
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("instance_id").(string)
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				s.RdsInstanceService: {
					Target:     []string{"Running"},
					Timeout:    resourceData.Timeout(schema.TimeoutDelete),
					ResourceId: resourceData.Get("instance_id").(string),
				},
			},
		},
	}
	callbacks = append(callbacks, nodeCallback)

	return callbacks
}

func (s *VolcengineRdsMysqlInstanceReadonlyNodeService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineRdsMysqlInstanceReadonlyNodeService) ReadResourceId(id string) string {
	return id
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
