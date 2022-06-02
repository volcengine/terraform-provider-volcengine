package cluster

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-vestack/common"
	"github.com/volcengine/terraform-provider-vestack/logger"
)

type VestackVkeClusterService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewVkeClusterService(c *ve.SdkClient) *VestackVkeClusterService {
	return &VestackVkeClusterService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VestackVkeClusterService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VestackVkeClusterService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
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

	return ve.WithPageNumberQuery(condition, "PageSize", "PageNumber", 20, 1, func(m map[string]interface{}) ([]interface{}, error) {
		vke := s.Client.VkeClient
		action := "ListClusters"
		logger.Debug(logger.ReqFormat, action, condition)
		if condition == nil {
			resp, err = vke.ListClustersCommon(nil)
			if err != nil {
				return data, err
			}
		} else {
			resp, err = vke.ListClustersCommon(&condition)
			if err != nil {
				return data, err
			}
		}

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

func (s *VestackVkeClusterService) ReadResource(resourceData *schema.ResourceData, clusterId string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if clusterId == "" {
		clusterId = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"Filter.Ids": []string{clusterId},
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
		return data, fmt.Errorf("Vke Cluster %s not exist ", clusterId)
	}
	return data, err
}

func (s *VestackVkeClusterService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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
			status, err = ve.ObtainSdkValue("Status.Phase", demo)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf(" Vke Cluster status error, status:%s", status.(string))
				}
			}
			//注意 返回的第一个参数不能为空 否则会一直等下去
			return demo, status.(string), err
		},
	}

}

func (VestackVkeClusterService) WithResourceResponseHandlers(cluster map[string]interface{}) []ve.ResourceResponseHandler {
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

func (s *VestackVkeClusterService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateCluster",
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"cluster_config": {
					ConvertType: ve.ConvertJsonObject,
					ForceGet:    true,
					NextLevelConvert: map[string]ve.RequestConvert{
						"subnet_ids": {
							ConvertType: ve.ConvertJsonArray,
						},
						"api_server_public_access_config": {
							ConvertType: ve.ConvertJsonObject,
							ForceGet:    true,
							NextLevelConvert: map[string]ve.RequestConvert{
								"public_access_network_config": {
									ConvertType: ve.ConvertJsonObject,
									ForceGet:    true,
								},
							},
						},
					},
				},
				"pods_config": {
					ConvertType: ve.ConvertJsonObject,
					ForceGet:    true,
					NextLevelConvert: map[string]ve.RequestConvert{
						"flannel_config": {
							ConvertType: ve.ConvertJsonObject,
							ForceGet:    true,
							NextLevelConvert: map[string]ve.RequestConvert{
								"pod_cidrs": {
									ConvertType: ve.ConvertJsonArray,
								},
							},
						},
						"vpc_cni_config": {
							ConvertType: ve.ConvertJsonObject,
							ForceGet:    true,
							NextLevelConvert: map[string]ve.RequestConvert{
								"subnet_ids": {
									ConvertType: ve.ConvertJsonArray,
								},
							},
						},
					},
				},
				"services_config": {
					ConvertType: ve.ConvertJsonObject,
					ForceGet:    true,
					NextLevelConvert: map[string]ve.RequestConvert{
						"service_cidrsv4": {
							ConvertType: ve.ConvertJsonArray,
						},
					},
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				b, _ := json.Marshal(*call.SdkParam)
				log.Println(string(b))
				//创建cluster
				return s.Client.VkeClient.CreateClusterCommon(call.SdkParam)
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

func (s *VestackVkeClusterService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateClusterConfig",
			ConvertMode: ve.RequestConvertAll,
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
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["Id"] = d.Id()
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				//修改cluster属性
				return s.Client.VkeClient.UpdateClusterConfigCommon(call.SdkParam)
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Running"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VestackVkeClusterService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteCluster",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"Id": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				//删除Cluster
				return s.Client.VkeClient.DeleteClusterCommon(call.SdkParam)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
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

func (s *VestackVkeClusterService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "Filter.Ids",
				ConvertType: ve.ConvertJsonArray,
			},
			"delete_protection_enabled": {
				TargetField: "Filter.DeleteProtectionEnabled",
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

func (s *VestackVkeClusterService) ReadResourceId(id string) string {
	return id
}
