package veecp_cluster

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"sort"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineVeecpClusterService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewVeecpClusterService(c *ve.SdkClient) *VolcengineVeecpClusterService {
	return &VolcengineVeecpClusterService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineVeecpClusterService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineVeecpClusterService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)

	if filter, filterExist := condition["Filter"]; filterExist {
		if podsConfig, exist := filter.(map[string]interface{})["PodsConfig"]; exist {
			if podNetworkMode, ex := podsConfig.(map[string]interface{})["PodNetworkMode"]; ex {
				condition["Filter"].(map[string]interface{})["PodsConfig.PodNetworkMode"] = podNetworkMode
				delete(condition["Filter"].(map[string]interface{}), "PodsConfig")
			}
		}
	}

	// 适配 Conditions.Type 字段
	if filter, filterExist := condition["Filter"]; filterExist {
		if statuses, exist := filter.(map[string]interface{})["Statuses"]; exist {
			for index, status := range statuses.([]interface{}) {
				if ty, ex := status.(map[string]interface{})["ConditionsType"]; ex {
					condition["Filter"].(map[string]interface{})["Statuses"].([]interface{})[index].(map[string]interface{})["Conditions.Type"] = ty
					delete(condition["Filter"].(map[string]interface{})["Statuses"].([]interface{})[index].(map[string]interface{}), "ConditionsType")
				}
			}
		}
	}

	data, err = ve.WithPageNumberQuery(condition, "PageSize", "PageNumber", 100, 1, func(m map[string]interface{}) ([]interface{}, error) {
		action := "ListClusters"

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
		// TODO: replace result items
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
		data, err = removeSystemTags(data)
		return data, err
	})

	// get extra data
	for _, d := range data {
		if cluster, ok := d.(map[string]interface{}); ok {
			// 1. cluster status
			status, err := ve.ObtainSdkValue("Status.Phase", cluster)
			if err != nil {
				logger.Info("Get cluster status failed, cluster: %+v, err: %s", cluster, err.Error())
				return data, err
			}
			if !clusterReadyStatuses[status.(string)] {
				logger.Info("Cluster not ready, cluster: %+v", cluster)
				continue
			}

			// 2. get kubeconfig and eip allocation id
			publicAccess, err := ve.ObtainSdkValue("ClusterConfig.ApiServerPublicAccessEnabled", cluster)
			if err != nil {
				logger.Info("Get cluster public access failed, cluster: %+v, err: %s", cluster, err.Error())
				return data, err
			}
			publicIp, err := ve.ObtainSdkValue("ClusterConfig.ApiServerEndpoints.PublicIp.Ipv4", cluster)
			if err != nil || publicIp == "" {
				logger.Info("Get cluster public ip error or public ip is empty, cluster: %+v, err: %v", cluster, err)
			} else {
				if publicAccess, ok := publicAccess.(bool); ok && publicAccess {
					// get eip data
					action := "DescribeEipAddresses"
					req := map[string]interface{}{
						"EipAddresses.1": publicIp,
					}
					eipAddressResp, err := s.Client.UniversalClient.DoCall(getVpcUniversalInfo(action), &req)
					if err != nil {
						return data, err
					}
					eipAddresses, err := ve.ObtainSdkValue("Result.EipAddresses", *eipAddressResp)
					if err != nil {
						return data, err
					}

					if eipAddresses, ok := eipAddresses.([]interface{}); !ok {
						return data, errors.New("Result.EipAddresses is not Slice")
					} else if len(eipAddresses) == 0 {
						return data, fmt.Errorf("Eip %s not found ", publicIp)
					} else {
						// get eip allocation id
						cluster["EipAllocationId"] = eipAddresses[0].(map[string]interface{})["AllocationId"].(string)

						// get eip bandwidth, billing_type, isp
						if clusterConfig, exist := cluster["ClusterConfig"]; exist {
							if apiServerPublicAccessConfig, exist := clusterConfig.(map[string]interface{})["ApiServerPublicAccessConfig"]; exist {
								if publicAccessNetworkConfig, exist := apiServerPublicAccessConfig.(map[string]interface{})["PublicAccessNetworkConfig"]; exist {
									if eipConfig, ok := publicAccessNetworkConfig.(map[string]interface{}); ok {
										eipConfig["BillingType"] = eipAddresses[0].(map[string]interface{})["BillingType"]
										eipConfig["Bandwidth"] = eipAddresses[0].(map[string]interface{})["Bandwidth"]
										eipConfig["Isp"] = eipAddresses[0].(map[string]interface{})["ISP"]
									}
								}
							}
						}
					}
				}
			}
			//clusterId := cluster["Id"].(string)
			//privateKubeconfigResp, err := s.getKubeconfig(clusterId, "Private")
			//if err != nil {
			//	logger.Info("Get private kubeconfig error, cluster: %+v, err: %s", cluster, err.Error())
			//	return data, err
			//}
			//
			//kubeconfigs, err := ve.ObtainSdkValue("Result.Items", *privateKubeconfigResp)
			//if err != nil {
			//	return data, err
			//}
			//if len(kubeconfigs.([]interface{})) > 0 {
			//	cluster["KubeconfigPrivate"] = kubeconfigs.([]interface{})[0].(map[string]interface{})["Kubeconfig"]
			//}
		}
	}

	return data, err

}

//func (s *VolcengineVeecpClusterService) getKubeconfig(clusterId, accessType string) (*map[string]interface{}, error) {
//	kubeconfigReq := &map[string]interface{}{
//		"Filter": map[string]interface{}{
//			"ClusterIds": []string{clusterId},
//			"Types":      []string{accessType},
//		},
//	}
//	logger.Debug(logger.ReqFormat, "ListKubeconfigs", kubeconfigReq)
//	return s.Client.UniversalClient.DoCall(getUniversalInfo("ListKubeconfigs"), kubeconfigReq)
//}

func (s *VolcengineVeecpClusterService) ReadResource(resourceData *schema.ResourceData, clusterId string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if clusterId == "" {
		clusterId = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"Filter": map[string]interface{}{
			"Ids": []string{clusterId},
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
		return data, fmt.Errorf("Veecp Cluster %s not exist ", clusterId)
	}

	return data, err
}

func (s *VolcengineVeecpClusterService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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
			status, err = ve.ObtainSdkValue("Status.Phase", d)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("veecp_cluster status error, status: %s", status.(string))
				}
			}
			return d, status.(string), err
		},
	}
}

func (s *VolcengineVeecpClusterService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateCluster",
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"cluster_config": {
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"subnet_ids": {
							ConvertType: ve.ConvertJsonArray,
						},
						"api_server_public_access_config": {
							ConvertType: ve.ConvertJsonObject,
							NextLevelConvert: map[string]ve.RequestConvert{
								"public_access_network_config": {
									ConvertType: ve.ConvertJsonObject,
								},
							},
						},
					},
				},
				"pods_config": {
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"flannel_config": {
							ConvertType: ve.ConvertJsonObject,
							NextLevelConvert: map[string]ve.RequestConvert{
								"pod_cidrs": {
									ConvertType: ve.ConvertJsonArray,
								},
							},
						},
						"vpc_cni_config": {
							ConvertType: ve.ConvertJsonObject,
							NextLevelConvert: map[string]ve.RequestConvert{
								"subnet_ids": {
									ConvertType: ve.ConvertJsonArray,
								},
								"vpc_id": {
									Ignore: true,
								},
							},
						},
					},
				},
				"services_config": {
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"service_cidrsv4": {
							ConvertType: ve.ConvertJsonArray,
						},
					},
				},
				"tags": {
					TargetField: "Tags",
					ConvertType: ve.ConvertJsonObjectArray,
				},
				"logging_config": {
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"log_setups": {
							ConvertType: ve.ConvertJsonObjectArray,
						},
					},
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if billingType, ok := (*call.SdkParam)["ClusterConfig.ApiServerPublicAccessConfig.PublicAccessNetworkConfig.BillingType"]; ok {
					realBillingType := billingTypeRequestConvert(d, billingType)
					(*call.SdkParam)["ClusterConfig.ApiServerPublicAccessConfig.PublicAccessNetworkConfig.BillingType"] = realBillingType
				}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				//创建cluster
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				//注意 获取内容 这个地方不能是指针 需要转一次
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

func (VolcengineVeecpClusterService) WithResourceResponseHandlers(cluster map[string]interface{}) []ve.ResourceResponseHandler {
	if clusterConfig, ok := cluster["ClusterConfig"]; ok {
		if apiServerPublicAccessConfig, ok := clusterConfig.(map[string]interface{})["ApiServerPublicAccessConfig"]; ok {
			if publicAccessNetworkConfig, ok := apiServerPublicAccessConfig.(map[string]interface{})["PublicAccessNetworkConfig"]; ok {
				apiServerPublicAccessConfig.(map[string]interface{})["PublicAccessNetworkConfig"] = []interface{}{publicAccessNetworkConfig}
			}
			clusterConfig.(map[string]interface{})["ApiServerPublicAccessConfig"] = []interface{}{apiServerPublicAccessConfig}
		}
	}

	if podsConfig, ok := cluster["PodsConfig"]; ok {
		if flannelConfig, ok := podsConfig.(map[string]interface{})["FlannelConfig"]; ok {
			podsConfig.(map[string]interface{})["FlannelConfig"] = []interface{}{flannelConfig}
		}
		if vpcCniConfig, ok := podsConfig.(map[string]interface{})["VpcCniConfig"]; ok {
			podsConfig.(map[string]interface{})["VpcCniConfig"] = []interface{}{vpcCniConfig}
		}
	}

	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return cluster, map[string]ve.ResponseConvert{
			"ClusterConfig": {
				TargetField: "cluster_config",
				Convert: func(i interface{}) interface{} {
					clusterConfig := i.(map[string]interface{})

					if apiServerPublicAccessConfig, ok := clusterConfig["ApiServerPublicAccessConfig"].([]interface{}); !ok {
						return i
					} else if publicAccessNetworkConfig, ok := apiServerPublicAccessConfig[0].(map[string]interface{})["PublicAccessNetworkConfig"].([]interface{}); !ok {
						return i
					} else {
						billingType := publicAccessNetworkConfig[0].(map[string]interface{})["BillingType"]
						if billingType == nil {
							return i
						}
						publicAccessNetworkConfig[0].(map[string]interface{})["BillingType"] = billingTypeResponseConvert(billingType)
					}

					return i
				},
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineVeecpClusterService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	callbacks = s.setResourceTags(resourceData, "Cluster", callbacks)

	return callbacks
}

func (s *VolcengineVeecpClusterService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteCluster",
			ContentType: ve.ContentTypeJson,
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["Id"] = resourceData.Id()
				// TODO： 检查这里填什么
				(*call.SdkParam)["CascadingDeleteResources"] = []string{"All"}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				//删除Cluster
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 10*time.Minute)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				if protection, ok := d.Get("delete_protection_enabled").(bool); ok && protection {
					// 开启集群保护，直接返回失败
					return baseErr
				}

				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on  reading cluster on delete %q, %w", d.Id(), callErr))
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

func (s *VolcengineVeecpClusterService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "Filter.Ids",
				ConvertType: ve.ConvertJsonArray,
			},
			"delete_protection_enabled": {
				TargetField: "Filter.DeleteProtectionEnabled",
			},
			"name": {
				TargetField: "Filter.Name",
			},
			"pods_config_pod_network_mode": {
				TargetField: "Filter.PodsConfig.PodNetworkMode",
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
		ContentType:  ve.ContentTypeJson,
		NameField:    "Name",
		IdField:      "Id",
		CollectField: "clusters",
		ResponseConverts: map[string]ve.ResponseConvert{
			"ClusterConfig": {
				TargetField: "cluster_config",
				Convert: func(i interface{}) interface{} {
					realBillingType, err := ve.ObtainSdkValue("ApiServerPublicAccessConfig.PublicAccessNetworkConfig.BillingType", i)
					if err != nil {
						return i
					}
					billingType := billingTypeResponseConvert(realBillingType)

					if clusterConfig, ok := i.(map[string]interface{}); !ok {
						return i
					} else if apiServerPublicAccessConfig, ok := clusterConfig["ApiServerPublicAccessConfig"].(map[string]interface{}); !ok {
						return i
					} else if publicAccessNetworkConfig, ok := apiServerPublicAccessConfig["PublicAccessNetworkConfig"].(map[string]interface{}); !ok {
						return i
					} else {
						publicAccessNetworkConfig["BillingType"] = billingType
					}

					return i
				},
			},
		},
	}
}

func (s *VolcengineVeecpClusterService) ReadResourceId(id string) string {
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

func getVpcUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "vpc",
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		Action:      actionName,
	}
}

func removeSystemTags(data []interface{}) ([]interface{}, error) {
	var (
		ok      bool
		result  map[string]interface{}
		results []interface{}
		tags    []interface{}
	)
	for _, d := range data {
		if result, ok = d.(map[string]interface{}); !ok {
			return results, errors.New("The elements in data are not map ")
		}
		tags, ok = result["Tags"].([]interface{})
		if ok {
			tags = ve.FilterSystemTags(tags)
			result["Tags"] = tags
		}
		results = append(results, result)
	}
	return results, nil
}

var (
	// cluster ready status
	clusterReadyStatuses = map[string]bool{
		"Running":  true,
		"Updating": true,
		"Stopped":  true,
		"Starting": true,
	}
)

func validateLogSetups(d *schema.ResourceData) error {
	if d.HasChange("logging_config.0.log_setups") {
		oldSet, newSet := d.GetChange("logging_config.0.log_setups")
		logger.DebugInfo("set get change", oldSet, newSet)
		oldTypeArr := make([]string, 0)
		newTypeArr := make([]string, 0)
		// 取到old和new的去重log type数组
		for _, o := range oldSet.(*schema.Set).List() {
			if oMap, ok := o.(map[string]interface{}); ok {
				if !ContainsInSlice(oldTypeArr, oMap["log_type"].(string)) {
					oldTypeArr = append(oldTypeArr, oMap["log_type"].(string))
				}
			}
		}
		for _, n := range newSet.(*schema.Set).List() {
			if nMap, ok := n.(map[string]interface{}); ok {
				if !ContainsInSlice(newTypeArr, nMap["log_type"].(string)) {
					newTypeArr = append(newTypeArr, nMap["log_type"].(string))
				}
			}
		}
		/*
			1. old数组长度大，必出现了减少 报错
			2. old数组长度小，需判断old所有type是否都在new中，如有缺失，报错
			3. old和new长度相等，需判断old和new完全相等
		*/
		if len(oldTypeArr) > len(newTypeArr) {
			return fmt.Errorf("logging setups can only be modified and added, and cannot be deleted")
		}
		if len(oldTypeArr) < len(newTypeArr) {
			for _, o := range oldTypeArr {
				if !ContainsInSlice(newTypeArr, o) {
					return fmt.Errorf("logging setups can only be modified and added, and cannot be deleted")
				}
			}
		} else {
			sort.Strings(newTypeArr)
			sort.Strings(oldTypeArr)
			if !reflect.DeepEqual(oldTypeArr, newTypeArr) {
				return fmt.Errorf("logging setups can only be modified and added, and cannot be deleted")
			}
		}
	}
	return nil
}

func ContainsInSlice(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

var billingTypeRequestConvert = func(data *schema.ResourceData, old interface{}) interface{} {
	ty := 0
	switch old.(string) {
	case "PrePaid":
		ty = 1
	case "PostPaidByBandwidth":
		ty = 2
	case "PostPaidByTraffic":
		ty = 3
	}
	return ty
}

var billingTypeResponseConvert = func(i interface{}) interface{} {
	var ty string
	switch i.(float64) {
	case 1:
		ty = "PrePaid"
	case 2:
		ty = "PostPaidByBandwidth"
	case 3:
		ty = "PostPaidByTraffic"
	default:
		ty = fmt.Sprintf("%v", i)
	}
	return ty
}

func (s *VolcengineVeecpClusterService) setResourceTags(resourceData *schema.ResourceData, resourceType string, callbacks []ve.Callback) []ve.Callback {
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
