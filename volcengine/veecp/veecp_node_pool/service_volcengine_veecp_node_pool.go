package veecp_node_pool

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
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
		Client:               c,
		Dispatcher:           &ve.Dispatcher{},
		securityGroupService: security_group.NewSecurityGroupService(c),
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
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 20, 1, func(condition map[string]interface{}) ([]interface{}, error) {

		if enabled, exist := condition["AutoScalingEnabled"]; exist {
			if _, filterExist := condition["Filter"]; !filterExist {
				condition["Filter"] = make(map[string]interface{})
			}
			condition["Filter"].(map[string]interface{})["AutoScaling.Enabled"] = enabled
			delete(condition, "AutoScalingEnabled")
		}

		// 单独适配 ClusterId 字段，将 ClusterId 加入 Filter.ClusterIds
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

		action := "ListNodePools"
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

func (s *VolcengineVeecpNodePoolService) ReadResource(resourceData *schema.ResourceData, nodePoolId string) (data map[string]interface{}, err error) {
	var (
		results interface{}
		resp    *map[string]interface{}
		result  map[string]interface{}
		temp    []interface{}
		ok      bool
	)
	if nodePoolId == "" {
		nodePoolId = s.ReadResourceId(resourceData.Id())
	}

	action := "ListNodePools"
	nodeId := []string{nodePoolId}
	condition := make(map[string]interface{}, 0)
	condition["Filter"] = map[string]interface{}{
		"Ids": nodeId,
	}

	logger.Debug(logger.RespFormat, "ReadResource ", condition)
	resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &condition)
	logger.Debug(logger.RespFormat, "ReadResource ", resp)

	if err != nil {
		return data, err
	}
	if resp == nil {
		return data, fmt.Errorf("NodePool %s not exist ", nodePoolId)
	}

	results, err = ve.ObtainSdkValue("Result.Items", *resp)
	if err != nil {
		return data, err
	}
	if results == nil {
		results = []interface{}{}
	}

	if temp, ok = results.([]interface{}); !ok {
		return data, errors.New("Result.Items is not Slice")
	}

	if len(temp) == 0 {
		return data, fmt.Errorf("NodePool %s not exist ", nodePoolId)
	}

	result = temp[0].(map[string]interface{})
	result["NodeConfig"].(map[string]interface{})["Security"].(map[string]interface{})["Login"].(map[string]interface{})["Password"] =
		resourceData.Get("node_config.0.security.0.login.0.password")

	// 安全组过滤默认安全组
	tmpSecurityGroupIds := result["NodeConfig"].(map[string]interface{})["Security"].(map[string]interface{})["SecurityGroupIds"].([]interface{})
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
		defaultSecurityGroupName := fmt.Sprintf("%v-common", result["ClusterId"])
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
			result["NodeConfig"].(map[string]interface{})["Security"].(map[string]interface{})["SecurityGroupIds"] = filteredSecurityGroupIds
		}

		logger.Debug(logger.RespFormat, "filteredSecurityGroupIds", tmpSecurityGroupIds, filteredSecurityGroupIds)
	}

	if instanceIds, ok := resourceData.GetOk("instance_ids"); ok {
		result["InstanceIds"] = instanceIds.(*schema.Set).List()
	}

	if ecsTags, ok := result["NodeConfig"].(map[string]interface{})["Tags"]; ok {
		result["NodeConfig"].(map[string]interface{})["EcsTags"] = ecsTags
		delete(result["NodeConfig"].(map[string]interface{}), "Tags")
	}

	logger.Debug(logger.RespFormat, "result of ReadResource ", result)
	return result, err
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
	var callbacks []ve.Callback

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateNodePool",
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
				"instance_ids": {
					Ignore: true,
				},
				"keep_instance_name": {
					Ignore: true,
				},
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
								"security_group_ids": {
									ConvertType: ve.ConvertJsonArray,
								},
								"security_strategies": {
									ConvertType: ve.ConvertJsonArray,
								},
							},
						},
						"system_volume": {
							Ignore: true,
						},
						"data_volumes": {
							Ignore: true,
						},
						"initialize_script": {
							ConvertType: ve.ConvertJsonObject,
						},
						"additional_container_storage_enabled": {
							ConvertType: ve.ConvertJsonObject,
						},
						"image_id": {
							ConvertType: ve.ConvertJsonObject,
						},
						"instance_charge_type": {
							ConvertType: ve.ConvertJsonObject,
						},
						"period": {
							ConvertType: ve.ConvertJsonObject,
						},
						"auto_renew": {
							ForceGet:    true,
							TargetField: "AutoRenew",
						},
						"auto_renew_period": {
							ConvertType: ve.ConvertJsonObject,
						},
						"name_prefix": {
							ConvertType: ve.ConvertJsonObject,
						},
						"ecs_tags": {
							TargetField: "Tags",
							ConvertType: ve.ConvertJsonObjectArray,
						},
						"hpc_cluster_ids": {
							ConvertType: ve.ConvertJsonArray,
						},
					},
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
						"cordon": {
							ConvertType: ve.ConvertJsonObject,
						},
						"name_prefix": {
							ConvertType: ve.ConvertJsonObject,
						},
					},
				},
				"auto_scaling": {
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"enabled": {
							TargetField: "Enabled",
						},
						"max_replicas": {
							TargetField: "MaxReplicas",
						},
						"min_replicas": {
							TargetField: "MinReplicas",
						},
						"desired_replicas": {
							TargetField: "DesiredReplicas",
						},
						"priority": {
							TargetField: "Priority",
						},
						"subnet_policy": {
							TargetField: "SubnetPolicy",
						},
					},
				},
				"tags": {
					TargetField: "Tags",
					ConvertType: ve.ConvertJsonObjectArray,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if chargeType, ok := (*call.SdkParam)["NodeConfig.InstanceChargeType"]; ok {
					if autoScalingEnabled, ok := (*call.SdkParam)["AutoScaling.Enabled"]; ok {
						if chargeType.(string) == "PrePaid" && autoScalingEnabled.(bool) {
							return false, fmt.Errorf("PrePaid charge type cannot support auto scaling")
						}
					}
				}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				// 手动转data_volumes
				if dataVolumes, ok := d.GetOk("node_config.0.data_volumes"); ok {
					delete((*call.SdkParam)["NodeConfig"].(map[string]interface{}), "DataVolumes")
					volumes := make([]interface{}, 0)
					for index, _ := range dataVolumes.([]interface{}) {
						volume := make(map[string]interface{})
						if v, ok := d.GetOkExists(fmt.Sprintf("node_config.0.data_volumes.%d.type", index)); ok {
							volume["Type"] = v
						}
						if v, ok := d.GetOkExists(fmt.Sprintf("node_config.0.data_volumes.%d.size", index)); ok {
							volume["Size"] = v
						}
						if v, ok := d.GetOkExists(fmt.Sprintf("node_config.0.data_volumes.%d.mount_point", index)); ok {
							volume["MountPoint"] = v
						}
						volumes = append(volumes, volume)
					}
					(*call.SdkParam)["NodeConfig"].(map[string]interface{})["DataVolumes"] = volumes
				}
				// 手动转system_volume
				if _, ok := d.GetOk("node_config.0.system_volume"); ok {
					delete((*call.SdkParam)["NodeConfig"].(map[string]interface{}), "SystemVolume")
					systemVolume := map[string]interface{}{}
					if v, ok := d.GetOkExists("node_config.0.system_volume.0.type"); ok {
						systemVolume["Type"] = v
					}
					if v, ok := d.GetOkExists("node_config.0.system_volume.0.size"); ok {
						systemVolume["Size"] = v
					}
					(*call.SdkParam)["NodeConfig"].(map[string]interface{})["SystemVolume"] = systemVolume
				}
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

	// 添加已有实例到自定义节点池
	nodeCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateNodes",
			ConvertMode: ve.RequestConvertInConvert,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"cluster_id": {
					TargetField: "ClusterId",
				},
				"keep_instance_name": {
					TargetField: "KeepInstanceName",
				},
				"instance_ids": {
					TargetField: "InstanceIds",
					ConvertType: ve.ConvertJsonArray,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if _, ok := d.GetOk("instance_ids"); ok {
					(*call.SdkParam)["NodePoolId"] = d.Id()
					(*call.SdkParam)["ClientToken"] = uuid.New().String()
					return true, nil
				}
				return false, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			//AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
			//	tmpIds, _ := ve.ObtainSdkValue("Result.Ids", *resp)
			//	ids := tmpIds.([]interface{})
			//	d.Set("node_ids", ids)
			//	return nil
			//},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Running"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	callbacks = append(callbacks, nodeCallback)

	return callbacks
}

func (VolcengineVeecpNodePoolService) WithResourceResponseHandlers(nodePool map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		var (
			security     = make([]interface{}, 0)
			systemVolume = make([]interface{}, 0)
			login        = make([]interface{}, 0)
		)

		priSecurity := nodePool["NodeConfig"].(map[string]interface{})["Security"]
		priLogin := priSecurity.(map[string]interface{})["Login"]
		delete(nodePool, "Login")
		login = append(login, priLogin)
		priSecurity.(map[string]interface{})["Login"] = login
		security = append(security, priSecurity)

		delete(nodePool, "Security")
		nodePool["NodeConfig"].(map[string]interface{})["Security"] = security

		priSystemVolume := nodePool["NodeConfig"].(map[string]interface{})["SystemVolume"]
		systemVolume = append(systemVolume, priSystemVolume)
		delete(nodePool, "SystemVolume")
		nodePool["NodeConfig"].(map[string]interface{})["SystemVolume"] = systemVolume

		return nodePool, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineVeecpNodePoolService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateNodePoolConfig",
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
				"instance_ids": {
					Ignore: true,
				},
				"keep_instance_name": {
					Ignore: true,
				},
				"node_config": {
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
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
								"security_group_ids": {
									ConvertType: ve.ConvertJsonArray,
								},
								"security_strategies": {
									ConvertType: ve.ConvertJsonArray,
								},
							},
						},
						"additional_container_storage_enabled": {
							ConvertType: ve.ConvertJsonObject,
						},
						"initialize_script": {
							ConvertType: ve.ConvertJsonObject,
						},
						"subnet_ids": {
							ConvertType: ve.ConvertJsonArray,
						},
						"period": {
							ConvertType: ve.ConvertJsonObject,
						},
						"auto_renew": {
							ForceGet:    true,
							TargetField: "AutoRenew",
						},
						"auto_renew_period": {
							ConvertType: ve.ConvertJsonObject,
						},
						"name_prefix": {
							ConvertType: ve.ConvertJsonObject,
						},
						"ecs_tags": {
							TargetField: "Tags",
							ConvertType: ve.ConvertJsonObjectArray,
						},
						"instance_type_ids": {
							ConvertType: ve.ConvertJsonArray,
						},
						"hpc_cluster_ids": {
							ConvertType: ve.ConvertJsonArray,
						},
						"image_id": {
							ConvertType: ve.ConvertJsonObject,
						},
					},
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
						"cordon": {
							ConvertType: ve.ConvertJsonObject,
						},
						"name_prefix": {
							ConvertType: ve.ConvertJsonObject,
						},
					},
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["Id"] = d.Id()
				(*call.SdkParam)["ClusterId"] = d.Get("cluster_id")

				delete(*call.SdkParam, "Tags")
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				//adapt vke api
				nodeconfig := (*call.SdkParam)["NodeConfig"]
				if nodeconfig != nil {
					security := nodeconfig.(map[string]interface{})["Security"]
					if security != nil {
						login := security.(map[string]interface{})["Login"]
						if login != nil && login.(map[string]interface{})["SshKeyPairName"] != nil && login.(map[string]interface{})["SshKeyPairName"].(string) == "" {
							delete((*call.SdkParam)["NodeConfig"].(map[string]interface{})["Security"].(map[string]interface{})["Login"].(map[string]interface{}), "SshKeyPairName")
						}
						_, exist := security.(map[string]interface{})["SecurityStrategies"]
						if !exist && d.HasChange("node_config.0.security.0.security_strategies") {
							security.(map[string]interface{})["SecurityStrategies"] = []interface{}{}
						}
					}

					if _, ok1 := nodeconfig.(map[string]interface{})["HpcClusterIds"]; ok1 {
						if _, ok2 := nodeconfig.(map[string]interface{})["InstanceTypeIds"]; !ok2 {
							(*call.SdkParam)["NodeConfig"].(map[string]interface{})["InstanceTypeIds"] = make([]interface{}, 0)
							instanceTypeIds := d.Get("node_config.0.instance_type_ids")
							for _, instanceTypeId := range instanceTypeIds.([]interface{}) {
								(*call.SdkParam)["NodeConfig"].(map[string]interface{})["InstanceTypeIds"] = append((*call.SdkParam)["NodeConfig"].(map[string]interface{})["InstanceTypeIds"].([]interface{}), instanceTypeId.(string))
							}
						}
					}
					if d.HasChange("node_config.0.hpc_cluster_ids") {
						ve.DefaultMapValue(call.SdkParam, "NodeConfig", map[string]interface{}{
							"HpcClusterIds": []interface{}{},
						})
					}
				}

				instanceChargeType := d.Get("node_config").([]interface{})[0].(map[string]interface{})["instance_charge_type"].(string)
				if instanceChargeType != "PrePaid" {
					if nodeCfg, ok := (*call.SdkParam)["NodeConfig"]; ok {
						if _, ok := nodeCfg.(map[string]interface{})["AutoRenew"]; ok {
							delete((*call.SdkParam)["NodeConfig"].(map[string]interface{}), "AutoRenew")
						}
					}
				}

				// 当列表被删除时，入参添加空列表来置空
				ve.DefaultMapValue(call.SdkParam, "KubernetesConfig", map[string]interface{}{
					"Labels": []interface{}{},
					"Taints": []interface{}{},
				})

				if d.HasChange("node_config.0.ecs_tags") {
					ve.DefaultMapValue(call.SdkParam, "NodeConfig", map[string]interface{}{
						"Tags": []interface{}{},
					})
				}

				// 手动转数据盘
				if d.HasChange("node_config.0.data_volumes") {
					if dataVolumes, ok := d.GetOk("node_config.0.data_volumes"); ok {
						volumes := make([]interface{}, 0)
						for index, _ := range dataVolumes.([]interface{}) {
							volume := make(map[string]interface{})
							if v, ok := d.GetOkExists(fmt.Sprintf("node_config.0.data_volumes.%d.type", index)); ok {
								volume["Type"] = v
							}
							if v, ok := d.GetOkExists(fmt.Sprintf("node_config.0.data_volumes.%d.size", index)); ok {
								volume["Size"] = v
							}
							if v, ok := d.GetOkExists(fmt.Sprintf("node_config.0.data_volumes.%d.mount_point", index)); ok {
								if v != nil && len(v.(string)) > 0 {
									volume["MountPoint"] = v
								}
							}
							volumes = append(volumes, volume)
						}
						(*call.SdkParam)["NodeConfig"].(map[string]interface{})["DataVolumes"] = volumes
					} else {
						// 用户清空数据盘，传空list
						(*call.SdkParam)["NodeConfig"].(map[string]interface{})["DataVolumes"] = []interface{}{}
					}
				}

				if d.HasChange("node_config.0.system_volume") {
					// 手动转system_volume
					if _, ok := d.GetOk("node_config.0.system_volume"); ok {
						systemVolume := map[string]interface{}{}
						if v, ok := d.GetOkExists("node_config.0.system_volume.0.type"); ok {
							systemVolume["Type"] = v
						}
						if v, ok := d.GetOkExists("node_config.0.system_volume.0.size"); ok {
							systemVolume["Size"] = v
						}
						(*call.SdkParam)["NodeConfig"].(map[string]interface{})["SystemVolume"] = systemVolume
					}
				}

				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Running"},
				Timeout: resourceData.Timeout(schema.TimeoutUpdate),
			},
		},
	}
	callbacks = append(callbacks, callback)

	if resourceData.HasChange("auto_scaling") {
		desiredReplicasCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "UpdateNodePoolConfig",
				ConvertMode: ve.RequestConvertInConvert,
				ContentType: ve.ContentTypeJson,
				Convert: map[string]ve.RequestConvert{
					"auto_scaling": {
						ConvertType: ve.ConvertJsonObject,
						NextLevelConvert: map[string]ve.RequestConvert{
							"enabled": {
								ForceGet:    true,
								TargetField: "Enabled",
							},
							"max_replicas": {
								ForceGet:    true,
								TargetField: "MaxReplicas",
							},
							"min_replicas": {
								ForceGet:    true,
								TargetField: "MinReplicas",
							},
							"desired_replicas": {
								ForceGet:    true,
								TargetField: "DesiredReplicas",
							},
							"priority": {
								ForceGet:    true,
								TargetField: "Priority",
							},
							"subnet_policy": {
								ForceGet:    true,
								TargetField: "SubnetPolicy",
							},
						},
					},
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["Id"] = d.Id()
					(*call.SdkParam)["ClusterId"] = d.Get("cluster_id")
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
					Timeout: resourceData.Timeout(schema.TimeoutUpdate),
				},
				AfterRefresh: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) error {
					result, err := s.ReadResource(d, d.Id())
					if err != nil {
						return err
					}
					nodes, ok := result["NodeStatistics"].(map[string]interface{})
					if !ok {
						return fmt.Errorf("NodeStatistics is not map ")
					}
					if int(nodes["TotalCount"].(float64)) != d.Get("auto_scaling.0.desired_replicas").(int) {
						return fmt.Errorf("The number of nodes in node_pool %s is inconsistent. Suggest obtaining more detailed error message through the Volcengine console. ", d.Id())
					}
					return nil
				},
			},
		}
		callbacks = append(callbacks, desiredReplicasCallback)
	}

	if resourceData.HasChange("instance_ids") {
		callbacks = s.updateNodes(resourceData, callbacks)
	}

	// 更新Tags
	callbacks = s.setResourceTags(resourceData, "NodePool", callbacks)

	return callbacks
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
			"tags": {
				TargetField: "Tags",
				ConvertType: ve.ConvertJsonObjectArray,
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
			"AutoScaling.SubnetPolicy": {
				TargetField: "subnet_policy",
			},
			"KubernetesConfig.NamePrefix": {
				TargetField: "kube_config_name_prefix",
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
							if p, ok := _data.(map[string]interface{})["MountPoint"]; ok { // 可能不存在
								volume["mount_point"] = p.(string)
							}
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

func (s *VolcengineVeecpNodePoolService) setResourceTags(resourceData *schema.ResourceData, resourceType string, callbacks []ve.Callback) []ve.Callback {
	addedTags, removedTags, _, _ := ve.GetSetDifference("tags", resourceData, ve.TagsHash, false)

	removeCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UntagResources",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if removedTags != nil && len(removedTags.List()) > 0 {
					(*call.SdkParam)["ResourceIds"] = []string{resourceData.Id()}
					(*call.SdkParam)["ResourceType"] = resourceType
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
			Action:      "TagResources",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if addedTags != nil && len(addedTags.List()) > 0 {
					(*call.SdkParam)["ResourceIds"] = []string{resourceData.Id()}
					(*call.SdkParam)["ResourceType"] = resourceType
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

func (s *VolcengineVeecpNodePoolService) updateNodes(resourceData *schema.ResourceData, callbacks []ve.Callback) []ve.Callback {
	addedNodes, removedNodes, _, _ := ve.GetSetDifference("instance_ids", resourceData, schema.HashString, false)

	removeCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteNodes",
			ConvertMode: ve.RequestConvertInConvert,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"cluster_id": {
					TargetField: "ClusterId",
					ForceGet:    true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if removedNodes != nil && len(removedNodes.List()) > 0 {
					nodes, err := s.getAllNodeIds(resourceData.Id())
					if err != nil {
						return false, err
					}
					var removeNodeList []string
					for _, v := range nodes {
						nodeMap, ok := v.(map[string]interface{})
						if !ok {
							return false, fmt.Errorf("getAllNodeIds Node is not map")
						}
						for _, instanceId := range removedNodes.List() {
							if nodeMap["InstanceId"] == instanceId {
								removeNodeList = append(removeNodeList, nodeMap["Id"].(string))
							}
						}
					}

					(*call.SdkParam)["NodePoolId"] = resourceData.Id()
					(*call.SdkParam)["Ids"] = removeNodeList
					(*call.SdkParam)["RetainResources"] = []string{"Ecs"}
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
	callbacks = append(callbacks, removeCallback)

	addCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateNodes",
			ConvertMode: ve.RequestConvertInConvert,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"cluster_id": {
					TargetField: "ClusterId",
					ForceGet:    true,
				},
				"keep_instance_name": {
					TargetField: "KeepInstanceName",
					ForceGet:    true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if addedNodes != nil && len(addedNodes.List()) > 0 {
					(*call.SdkParam)["NodePoolId"] = resourceData.Id()
					(*call.SdkParam)["InstanceIds"] = addedNodes.List()
					(*call.SdkParam)["ClientToken"] = uuid.New().String()
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
	callbacks = append(callbacks, addCallback)

	return callbacks
}

func (s *VolcengineVeecpNodePoolService) getAllNodeIds(nodePoolId string) (nodes []interface{}, err error) {
	// describe nodes
	req := map[string]interface{}{
		"Filter": map[string]interface{}{
			"NodePoolIds": []string{nodePoolId},
		},
	}
	action := "ListNodes"
	resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(action), &req)
	if err != nil {
		return nodes, err
	}
	logger.Debug(logger.RespFormat, action, req, *resp)
	results, err := ve.ObtainSdkValue("Result.Items", *resp)
	if err != nil {
		return nodes, err
	}
	if results == nil {
		results = []interface{}{}
	}
	nodes, ok := results.([]interface{})
	if !ok {
		return nodes, errors.New("Result.Items is not Slice")
	}
	return nodes, nil
}
