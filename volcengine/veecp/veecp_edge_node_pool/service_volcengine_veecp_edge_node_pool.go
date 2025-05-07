package veecp_edge_node_pool

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/security_group"
)

type VolcengineVeecpNodePoolService struct {
	Client               *ve.SdkClient
	Dispatcher           *ve.Dispatcher
	securityGroupService *security_group.VolcengineSecurityGroupService
}

func NewVeecpNodePoolService(c *ve.SdkClient) *VolcengineVeecpNodePoolService {
	return &VolcengineVeecpNodePoolService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineVeecpNodePoolService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineVeecpNodePoolService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)

	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		if enabled, exist := condition["AutoScalingEnabled"]; exist {
			if _, filterExist := condition["Filter"]; !filterExist {
				condition["Filter"] = make(map[string]interface{})
			}
			condition["Filter"].(map[string]interface{})["AutoScaling.Enabled"] = enabled
			delete(condition, "AutoScalingEnabled")
		}

		if filter, filterExist := condition["Filter"]; filterExist {
			if clusterId, clusterIdExist := filter.(map[string]interface{})["ClusterId"]; clusterIdExist {
				if clusterIds, clusterIdsExist := filter.(map[string]interface{})["ClusterIds"]; clusterIdsExist {
					appendFlag := true
					for _, id := range clusterIds.([]interface{}) {
						if id == clusterId {
							appendFlag = false
						}
					}
					if appendFlag {
						condition["Filter"].(map[string]interface{})["ClusterIds"] = append(condition["Filter"].(map[string]interface{})["ClusterIds"].([]interface{}), clusterId)
					}
				} else {
					condition["Filter"].(map[string]interface{})["ClusterIds"] = []interface{}{clusterId}
				}
				delete(condition["Filter"].(map[string]interface{}), "ClusterId")
			}
		}

		action := "ListEdgeNodePools"
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

		respBytes, _ := json.Marshal(resp)
		logger.Debug(logger.RespFormat, action, condition, string(respBytes))
		results, err = ve.ObtainSdkValue("Result.Items", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.Items is not Slice")
		}

		for _, d := range data {
			dMap := d.(map[string]interface{})
			if elasticConfig, ok := dMap["ElasticConfig"]; ok {
				if eMap, ok := elasticConfig.(map[string]interface{}); ok {
					eMap["AutoScaleConfig"] = []interface{}{eMap["AutoScaleConfig"]}
					eMap["InstanceArea"] = []interface{}{eMap["InstanceArea"]}
				} else {
					return nil, fmt.Errorf("ElasticConfig is not map")
				}
			}
		}
		return data, err
	})
}

func (s *VolcengineVeecpNodePoolService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}

	req := map[string]interface{}{
		"Filter": map[string]interface{}{
			"Ids": []string{id},
		},
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
		return data, fmt.Errorf("veecp_edge_node_pool %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineVeecpNodePoolService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
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
			status = d["Status"].(map[string]interface{})["Phase"]
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("veecp_edge_node_pool status error, status: %s", status.(string))
				}
			}
			return d, status.(string), err
		},
	}
}

func (s *VolcengineVeecpNodePoolService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateEdgeNodePool",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"cluster_id": {
					TargetField: "ClusterId",
				},
				"client_token": {
					TargetField: "ClientToken",
				},
				"name": {
					TargetField: "Name",
				},
				"kubernetes_config": {
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"labels": {
							ConvertType: ve.ConvertJsonObjectArray,
						},
						"taints": {
							ConvertType: ve.ConvertJsonObjectArray,
						},
						//"cordon": {
						//	ConvertType: ve.ConvertJsonObject,
						//},
						//"name_prefix": {
						//	ConvertType: ve.ConvertJsonObject,
						//},
					},
				},
				"elastic_config": {
					TargetField: "ElasticConfig",
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"auto_scale_config": {
							ConvertType: ve.ConvertJsonObject,
							NextLevelConvert: map[string]ve.RequestConvert{
								"min_replicas": {
									TargetField: "MinReplicas",
									ForceGet:    true,
								},
								"max_replicas": {
									TargetField: "MaxReplicas",
									ForceGet:    true,
								},
								"desired_replicas": {
									TargetField: "DesiredReplicas",
									ForceGet:    true,
								},
								"priority": {
									TargetField: "Priority",
									ForceGet:    true,
								},
								"enabled": {
									TargetField: "Enabled",
									ForceGet:    true,
								},
							},
						},
						"instance_area": {
							ConvertType: ve.ConvertJsonObject,
						},
					},
				},
				"billing_configs": {
					ConvertType: ve.ConvertJsonObject,
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("Result.Id", *resp)
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

	return callbacks
}

func (VolcengineVeecpNodePoolService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineVeecpNodePoolService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateEdgeNodePoolConfig",
			ConvertMode: ve.RequestConvertInConvert,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"cluster_id": {
					TargetField: "ClusterId",
				},
				"client_token": {
					TargetField: "ClientToken",
				},
				"name": {
					TargetField: "Name",
				},
				"kubernetes_config": {
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"labels": {
							ConvertType: ve.ConvertJsonObjectArray,
							ForceGet:    true,
						},
						"taints": {
							ConvertType: ve.ConvertJsonObjectArray,
							ForceGet:    true,
						},
						//"cordon": {
						//	ConvertType: ve.ConvertJsonObject,
						//},
						//"name_prefix": {
						//	ConvertType: ve.ConvertJsonObject,
						//},
					},
				},
				"elastic_config": {
					TargetField: "ElasticConfig",
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"auto_scale_config": {
							ConvertType: ve.ConvertJsonObject,
							NextLevelConvert: map[string]ve.RequestConvert{
								"min_replicas": {
									TargetField: "MinReplicas",
									ForceGet:    true,
								},
								"max_replicas": {
									TargetField: "MaxReplicas",
									ForceGet:    true,
								},
								"desired_replicas": {
									TargetField: "DesiredReplicas",
									ForceGet:    true,
								},
								"enabled": {
									TargetField: "Enabled",
									ForceGet:    true,
								},
								"priority": {
									TargetField: "Priority",
									ForceGet:    true,
								},
							},
						},
						"instance_area": {
							ConvertType: ve.ConvertJsonObject,
						},
					},
				},
				"billing_configs": {
					ConvertType: ve.ConvertJsonObject,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["Id"] = d.Id()
				(*call.SdkParam)["ClusterId"] = d.Get("cluster_id")
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				// 当列表被删除时，入参添加空列表来置空
				ve.DefaultMapValue(call.SdkParam, "KubernetesConfig", map[string]interface{}{
					"Labels": []interface{}{},
					"Taints": []interface{}{},
				})

				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Running"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineVeecpNodePoolService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteNodePool",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"Id":                       resourceData.Id(),
				"ClusterId":                resourceData.Get("cluster_id"),
				"RetainResources":          []string{},
				"CascadingDeleteResources": []string{"Ecs"},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 5*time.Minute)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineVeecpNodePoolService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "Filter.Ids",
				ConvertType: ve.ConvertJsonArray,
			},
			"statuses": {
				TargetField: "Filter.Statuses",
				ConvertType: ve.ConvertJsonObjectArray,
				NextLevelConvert: map[string]ve.RequestConvert{
					"phase": {
						TargetField: "Phase",
					},
					"conditions_type": {
						TargetField: "ConditionsType",
					},
				},
			},
			"cluster_id": {
				TargetField: "Filter.ClusterId",
			},
			"cluster_ids": {
				TargetField: "Filter.ClusterIds",
				ConvertType: ve.ConvertJsonArray,
			},
			"name": {
				TargetField: "Filter.Name",
			},
			"create_client_token": {
				TargetField: "Filter.CreateClientToken",
			},
			"update_client_token": {
				TargetField: "Filter.UpdateClientToken",
			},
			"node_pool_types": {
				TargetField: "Filter.NodePoolTypes",
				ConvertType: ve.ConvertJsonArray,
			},
		},
		NameField:    "Name",
		IdField:      "Id",
		CollectField: "node_pools",
		ContentType:  ve.ContentTypeJson,
		ResponseConverts: map[string]ve.ResponseConvert{
			"Status.Phase": {
				TargetField: "phase",
			},
			"Status.Conditions": {
				TargetField: "condition_types",
				Convert: func(i interface{}) interface{} {
					var results []interface{}
					if dd, ok := i.([]interface{}); ok {
						for _, _data := range dd {
							results = append(results, _data.(map[string]interface{})["Type"])
						}
					}
					return results
				},
			},
			"KubernetesConfig.Labels": {
				TargetField: "label_content",
				Convert: func(i interface{}) interface{} {
					var results []interface{}
					if dd, ok := i.([]interface{}); ok {
						for _, _data := range dd {
							label := make(map[string]string, 0)
							label["key"] = _data.(map[string]interface{})["Key"].(string)
							label["value"] = _data.(map[string]interface{})["Value"].(string)
							results = append(results, label)
						}
					}
					return results
				},
			},
			"KubernetesConfig.Taints": {
				TargetField: "taint_content",
				Convert: func(i interface{}) interface{} {
					var results []interface{}
					if dd, ok := i.([]interface{}); ok {
						for _, _data := range dd {
							label := make(map[string]string, 0)
							label["key"] = _data.(map[string]interface{})["Key"].(string)
							label["value"] = _data.(map[string]interface{})["Value"].(string)
							label["effect"] = _data.(map[string]interface{})["Effect"].(string)
							results = append(results, label)
						}
					}
					return results
				},
			},
			"NodeStatistics": {
				TargetField: "node_statistics",
				Convert: func(i interface{}) interface{} {
					label := make(map[string]interface{}, 0)
					label["total_count"] = int(i.(map[string]interface{})["TotalCount"].(float64))
					label["creating_count"] = int(i.(map[string]interface{})["CreatingCount"].(float64))
					label["running_count"] = int(i.(map[string]interface{})["RunningCount"].(float64))
					label["updating_count"] = int(i.(map[string]interface{})["UpdatingCount"].(float64))
					label["deleting_count"] = int(i.(map[string]interface{})["DeletingCount"].(float64))
					label["failed_count"] = int(i.(map[string]interface{})["FailedCount"].(float64))
					label["stopped_count"] = int(i.(map[string]interface{})["StoppedCount"].(float64))
					label["stopping_count"] = int(i.(map[string]interface{})["StoppingCount"].(float64))
					label["starting_count"] = int(i.(map[string]interface{})["StartingCount"].(float64))
					return label
				},
			},
		},
	}
}

func (s *VolcengineVeecpNodePoolService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "veecp_openapi",
		Version:     "2021-03-03",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}
