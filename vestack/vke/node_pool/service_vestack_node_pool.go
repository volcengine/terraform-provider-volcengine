package node_pool

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-vestack/common"
	"github.com/volcengine/terraform-provider-vestack/logger"
	"strconv"
	"time"
)

type VestackNodePoolService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewNodePoolService(c *ve.SdkClient) *VestackNodePoolService {
	return &VestackNodePoolService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VestackNodePoolService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VestackNodePoolService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)

	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 20, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		// adapt vke api
		if enabled, exist := condition["AutoScalingEnabled"]; exist {
			if _, filterExist := condition["Filter"]; !filterExist {
				condition["Filter"] = make(map[string]interface{})
			}
			condition["Filter"].(map[string]interface{})["AutoScaling.Enabled"] = enabled
			delete(condition, "AutoScalingEnabled")
		}

		vkeClient := s.Client.VkeClient
		action := "ListNodePools"
		logger.Debug(logger.ReqFormat, action, condition)
		if condition == nil {
			resp, err = vkeClient.ListNodePoolsCommon(nil)
			if err != nil {
				return data, err
			}
		} else {
			resp, err = vkeClient.ListNodePoolsCommon(&condition)
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

		return data, err
	})
}

func (s *VestackNodePoolService) ReadResource(resourceData *schema.ResourceData, nodePoolId string) (data map[string]interface{}, err error) {
	var (
		result *map[string]interface{}
	)
	if nodePoolId == "" {
		nodePoolId = s.ReadResourceId(resourceData.Id())
	}

	vkeClient := s.Client.VkeClient
	nodeId := []string{nodePoolId}
	condition := make(map[string]interface{}, 0)
	condition["Filter"] = map[string]interface{}{
		"Ids": nodeId,
	}

	logger.Debug(logger.RespFormat, "ReadResource ", condition)
	result, err = vkeClient.ListNodePoolsCommon(&condition)
	logger.Debug(logger.RespFormat, "ReadResource ", result)

	if err != nil {
		return data, err
	}
	if result == nil {
		return data, fmt.Errorf("NodePool %s not exist ", nodePoolId)
	}

	return *result, err
}

func (s *VestackNodePoolService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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
			failStates = append(failStates, "Failed")
			demo, err = s.ReadResource(resourceData, id)
			if err != nil {
				return nil, "", err
			}
			//这里vke是Status是一个Object，取Phase字段判断是否失败
			status = demo["Result"].(map[string]interface{})["Items"].([]interface{})[0].(map[string]interface{})["Status"].(map[string]interface{})["Phase"]
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("node pool status error, status:%s", status.(string))
				}
			}
			return demo, status.(string), err
		},
	}

}

func (VestackNodePoolService) WithResourceResponseHandlers(nodePool map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return nodePool, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}

}

func (s *VestackNodePoolService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateNodePool",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"node_config": {
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"instance_type_ids": {
							ConvertType: ve.ConvertJsonArray,
						},
						"subnet_ids": {
							ConvertType: ve.ConvertJsonArray,
						},
						"security": {
							ConvertType: ve.ConvertJsonObject,
							NextLevelConvert: map[string]ve.RequestConvert{
								"login": {
									ConvertType: ve.ConvertJsonObject,
									NextLevelConvert: map[string]ve.RequestConvert{
										"password": {
											ConvertType: ve.ConvertJsonObject,
										},
										"ssh_key_pair_name": {
											ConvertType: ve.ConvertJsonObject,
										},
									},
								},
							},
						},
					},
				},
			},

			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.VkeClient.CreateNodePoolCommon(call.SdkParam)
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
	return []ve.Callback{callback}
}

func (s *VestackNodePoolService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateNodePoolConfig",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["Id"] = d.Id()
				(*call.SdkParam)["ClusterId"] = d.Get("cluster_id")
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.VkeClient.UpdateNodePoolConfigCommon(call.SdkParam)
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Running"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VestackNodePoolService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteNodePool",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"Id":        resourceData.Id(),
				"ClusterId": resourceData.Get("cluster_id"),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.VkeClient.DeleteNodePoolCommon(call.SdkParam)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VestackNodePoolService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
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
			"name": {
				TargetField: "Filter.Name",
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
			"AutoScaling.Enabled": {
				TargetField: "enabled",
			},
			"AutoScaling.DesiredReplicas": {
				TargetField: "desired_replicas",
			},
			"AutoScaling.MinReplicas": {
				TargetField: "min_replicas",
			},
			"AutoScaling.MaxReplicas": {
				TargetField: "max_replicas",
			},
			"AutoScaling.Priority": {
				TargetField: "priority",
			},
			"KubernetesConfig.Cordon": {
				TargetField: "cordon",
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
			"NodeConfig.InitializeScript": {
				TargetField: "initialize_script",
			},
			"NodeConfig.AdditionalContainerStorageEnabled": {
				TargetField: "additional_container_storage_enabled",
			},
			"NodeConfig.InstanceTypeIds": {
				TargetField: "instance_type_ids",
				Convert: func(i interface{}) interface{} {
					var results []interface{}
					if dd, ok := i.([]interface{}); ok {
						results = dd
					}
					return results
				},
			},
			"NodeConfig.SubnetIds": {
				TargetField: "subnet_ids",
				Convert: func(i interface{}) interface{} {
					var results []interface{}
					if dd, ok := i.([]interface{}); ok {
						results = dd
					}
					return results
				},
			},
			"NodeConfig.SystemVolume": {
				TargetField: "system_volume",
				Convert: func(i interface{}) interface{} {
					var results []interface{}
					volume := make(map[string]interface{}, 0)
					volume["type"] = i.(map[string]interface{})["Type"].(string)
					volume["size"] = strconv.FormatFloat(i.(map[string]interface{})["Size"].(float64), 'g', 5, 32)
					results = append(results, volume)
					return results
				},
			},
			"NodeConfig.DataVolumes": {
				TargetField: "data_volumes",
				Convert: func(i interface{}) interface{} {
					var results []interface{}
					if dd, ok := i.([]interface{}); ok {
						for _, _data := range dd {
							volume := make(map[string]interface{}, 0)
							volume["size"] = strconv.FormatFloat(_data.(map[string]interface{})["Size"].(float64), 'g', 5, 32)
							volume["type"] = _data.(map[string]interface{})["Type"].(string)
							results = append(results, volume)
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
					return label
				},
			},
		},
	}
}

func (s *VestackNodePoolService) ReadResourceId(id string) string {
	return id
}
