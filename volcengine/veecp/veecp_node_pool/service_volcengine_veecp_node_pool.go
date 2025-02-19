package veecp_node_pool

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
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
		return data, fmt.Errorf("veecp_node_pool %s not exist ", id)
	}

	if _, ok = data["NodeConfig"]; ok {
		data["NodeConfig"].(map[string]interface{})["Security"].(map[string]interface{})["Login"].(map[string]interface{})["Password"] =
			resourceData.Get("node_config.0.security.0.login.0.password")

		// 安全组过滤默认安全组
		tmpSecurityGroupIds := data["NodeConfig"].(map[string]interface{})["Security"].(map[string]interface{})["SecurityGroupIds"].([]interface{})
		if len(tmpSecurityGroupIds) > 0 {
			// 查询安全组
			securityGroupIdMap := make(map[string]interface{})
			for i, securityGroupId := range tmpSecurityGroupIds {
				securityGroupIdMap[fmt.Sprintf("SecurityGroupIds.%d", i+1)] = securityGroupId
			}
			securityGroups, err := s.securityGroupService.ReadResources(securityGroupIdMap)
			logger.Debug(logger.RespFormat, "DescribeSecurityGroups", securityGroupIdMap, securityGroups)
			if err != nil {
				return nil, err
			}

			// 每个节点池有个默认安全组，名称是${cluster_id}-common, 如果没有配置默认安全组，在这里过滤一下默认安全组
			defaultSecurityGroupName := fmt.Sprintf("%v-common", data["ClusterId"])
			nameMap := make(map[string]string)
			filteredSecurityGroupIds := make([]interface{}, 0)
			defaultCount := 0
			defaultSecurityGroupId := ""
			for _, securityGroup := range securityGroups {
				nameMap[securityGroup.(map[string]interface{})["SecurityGroupId"].(string)] = securityGroup.(map[string]interface{})["SecurityGroupName"].(string)
			}
			for _, securityGroupId := range tmpSecurityGroupIds {
				if nameMap[securityGroupId.(string)] == defaultSecurityGroupName {
					defaultCount++
					defaultSecurityGroupId = securityGroupId.(string)
					continue
				}
				filteredSecurityGroupIds = append(filteredSecurityGroupIds, securityGroupId)
			}
			if defaultCount > 1 {
				return nil, fmt.Errorf("default security group is not unique")
			}

			// 如果用户传了默认安全组id，不需要过滤
			oldSecurityGroupIds := resourceData.Get("node_config.0.security.0.security_group_ids").([]interface{})
			useDefaultSecurityGroupId := false
			for _, securityGroupId := range oldSecurityGroupIds {
				if securityGroupId.(string) == defaultSecurityGroupId {
					useDefaultSecurityGroupId = true
				}
			}
			if !useDefaultSecurityGroupId {
				data["NodeConfig"].(map[string]interface{})["Security"].(map[string]interface{})["SecurityGroupIds"] = filteredSecurityGroupIds
			}

			logger.Debug(logger.RespFormat, "filteredSecurityGroupIds", tmpSecurityGroupIds, filteredSecurityGroupIds)
		}

		if instanceIds, ok := resourceData.GetOk("instance_ids"); ok {
			data["InstanceIds"] = instanceIds.(*schema.Set).List()
		}

		if ecsTags, ok := data["NodeConfig"].(map[string]interface{})["Tags"]; ok {
			data["NodeConfig"].(map[string]interface{})["EcsTags"] = ecsTags
			delete(data["NodeConfig"].(map[string]interface{}), "Tags")
		}

		logger.Debug(logger.RespFormat, "result of ReadResource ", data)
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
					return nil, "", fmt.Errorf("veecp_node_pool status error, status: %s", status.(string))
				}
			}
			return d, status.(string), err
		},
	}
}

func (s *VolcengineVeecpNodePoolService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			// TODO: replace create action
			Action:      "CreateResource",
			ConvertMode: ve.RequestConvertAll,
			Convert:     map[string]ve.RequestConvert{},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				// TODO: replace id fields
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

func (VolcengineVeecpNodePoolService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		if _, ok := d["NodeConfig"]; ok {
			var (
				security     = make([]interface{}, 0)
				systemVolume = make([]interface{}, 0)
				login        = make([]interface{}, 0)
			)

			priSecurity := d["NodeConfig"].(map[string]interface{})["Security"]
			priLogin := priSecurity.(map[string]interface{})["Login"]
			delete(d, "Login")
			login = append(login, priLogin)
			priSecurity.(map[string]interface{})["Login"] = login
			security = append(security, priSecurity)

			delete(d, "Security")
			d["NodeConfig"].(map[string]interface{})["Security"] = security

			priSystemVolume := d["NodeConfig"].(map[string]interface{})["SystemVolume"]
			systemVolume = append(systemVolume, priSystemVolume)
			delete(d, "SystemVolume")
			d["NodeConfig"].(map[string]interface{})["SystemVolume"] = systemVolume

			return d, nil, nil
		}
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineVeecpNodePoolService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			// TODO: replace modify action
			Action:      "ModifyResource",
			ConvertMode: ve.RequestConvertAll,
			Convert:     map[string]ve.RequestConvert{},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["Id"] = d.Id()
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
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

			//TODO: 确认下是否存在删除接口
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
			"tags": {
				TargetField: "Tags",
				ConvertType: ve.ConvertJsonObjectArray,
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
			"NodeConfig.ImageId": {
				TargetField: "image_id",
			},
			"NodeConfig.SystemVolume": {
				TargetField: "system_volume",
				Convert: func(i interface{}) interface{} {
					var results []interface{}
					if i.(map[string]interface{})["Type"] == nil || i.(map[string]interface{})["Size"] == nil {
						return results
					}
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
			"NodeConfig.Security.SecurityGroupIds": {
				TargetField: "security_group_ids",
				Convert: func(i interface{}) interface{} {
					var results []interface{}
					if dd, ok := i.([]interface{}); ok {
						results = dd
					}
					return results
				},
			},
			"NodeConfig.Security.SecurityStrategyEnabled": {
				TargetField: "security_strategy_enabled",
			},
			"NodeConfig.Security.SecurityStrategies": {
				TargetField: "security_strategies",
				Convert: func(i interface{}) interface{} {
					var results []interface{}
					if dd, ok := i.([]interface{}); ok {
						results = dd
					}
					return results
				},
			},
			"NodeConfig.Security.Login.Type": {
				TargetField: "login_type",
			},
			"NodeConfig.Security.Login.SshKeyPairName": {
				TargetField: "login_key_pair_name",
			},
			"NodeConfig.InstanceChargeType": {
				TargetField: "instance_charge_type",
			},
			"NodeConfig.Period": {
				TargetField: "period",
			},
			"NodeConfig.AutoRenew": {
				TargetField: "auto_renew",
			},
			"NodeConfig.AutoRenewPeriod": {
				TargetField: "auto_renew_period",
			},
			"NodeConfig.NamePrefix": {
				TargetField: "name_prefix",
			},
			"NodeConfig.HpcClusterIds": {
				TargetField: "hpc_cluster_ids",
			},
			"NodeConfig.Tags": {
				TargetField: "ecs_tags",
				Convert: func(i interface{}) interface{} {
					var results []interface{}
					if dd, ok := i.([]interface{}); ok {
						for _, data := range dd {
							tag := make(map[string]interface{}, 0)
							tag["key"] = data.(map[string]interface{})["Key"].(string)
							tag["value"] = data.(map[string]interface{})["Value"].(string)
							results = append(results, tag)
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
