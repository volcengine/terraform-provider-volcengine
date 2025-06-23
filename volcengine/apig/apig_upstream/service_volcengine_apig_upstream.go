package apig_upstream

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

type VolcengineApigUpstreamService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewApigUpstreamService(c *ve.SdkClient) *VolcengineApigUpstreamService {
	return &VolcengineApigUpstreamService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineApigUpstreamService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineApigUpstreamService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListUpstreams"

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

func (s *VolcengineApigUpstreamService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
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
		return data, fmt.Errorf("apig_upstream %s not exist ", id)
	}

	if upstreamSpec, exist := data["UpstreamSpec"]; exist {
		if upstreamSpecMap, ok := upstreamSpec.(map[string]interface{}); ok {
			if v, exist := upstreamSpecMap["VeFaas"]; exist {
				if veFaas, ok := v.(map[string]interface{}); ok {
					upstreamSpecMap["VeFaas"] = []interface{}{veFaas}
				}
			}
			if v, exist := upstreamSpecMap["K8SService"]; exist {
				if k8SService, ok := v.(map[string]interface{}); ok {
					upstreamSpecMap["K8SService"] = []interface{}{k8SService}
				}
			}
			if v, exist := upstreamSpecMap["Domain"]; exist {
				if domain, ok := v.(map[string]interface{}); ok {
					upstreamSpecMap["Domain"] = []interface{}{domain}
				}
			}
			if v, exist := upstreamSpecMap["NacosService"]; exist {
				if nacosService, ok := v.(map[string]interface{}); ok {
					upstreamSpecMap["NacosService"] = []interface{}{nacosService}
				}
			}
			if v, exist := upstreamSpecMap["VeMLP"]; exist {
				if veMLP, ok := v.(map[string]interface{}); ok {
					if v1, exist1 := veMLP["K8SService"]; exist1 {
						if k8SService, ok1 := v1.(map[string]interface{}); ok1 {
							if v2, exist2 := k8SService["ClusterInfo"]; exist2 {
								if clusterInfo, ok2 := v2.(map[string]interface{}); ok2 {
									k8SService["ClusterInfo"] = []interface{}{clusterInfo}
								}
							}
							veMLP["K8SService"] = []interface{}{k8SService}
						}
					}
					upstreamSpecMap["VeMLP"] = []interface{}{veMLP}
				}
			}
			if v, exist := upstreamSpecMap["AIProvider"]; exist {
				if aiProvider, ok := v.(map[string]interface{}); ok {
					if v1, exist1 := aiProvider["CustomModelService"]; exist1 {
						if customModelService, ok1 := v1.(map[string]interface{}); ok1 {
							aiProvider["CustomModelService"] = []interface{}{customModelService}
						}
					}
					upstreamSpecMap["AIProvider"] = []interface{}{aiProvider}
				}
			}
		}
	}

	if loadBalancerSettings, exist := data["LoadBalancerSettings"]; exist {
		if loadBalancerSettingsMap, ok := loadBalancerSettings.(map[string]interface{}); ok {
			if v, exist := loadBalancerSettingsMap["ConsistentHashLB"]; exist {
				if consistentHashLB, ok := v.(map[string]interface{}); ok {
					if v1, exist1 := consistentHashLB["HTTPCookie"]; exist1 {
						if httpCookie, ok1 := v1.(map[string]interface{}); ok1 {
							consistentHashLB["HTTPCookie"] = []interface{}{httpCookie}
						}
					}
					loadBalancerSettingsMap["ConsistentHashLB"] = []interface{}{consistentHashLB}
				}
			}
		}
	}

	if _, exist := data["VersionDetails"]; !exist {
		data["VersionDetails"] = []interface{}{}
	}

	return data, err
}

func (s *VolcengineApigUpstreamService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (VolcengineApigUpstreamService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"SimpleLB": {
				TargetField: "simple_lb",
			},
			"ConsistentHashLB": {
				TargetField: "consistent_hash_lb",
			},
			"HTTPCookie": {
				TargetField: "http_cookie",
			},
			"K8SService": {
				TargetField: "k8s_service",
			},
			"IP": {
				TargetField: "ip",
			},
			"FixedIPList": {
				TargetField: "fixed_ip_list",
			},
			"VeMLP": {
				TargetField: "ve_mlp",
			},
			"ServiceURL": {
				TargetField: "service_url",
			},
			"AIProvider": {
				TargetField: "ai_provider",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineApigUpstreamService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateUpstream",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"load_balancer_settings": {
					TargetField: "LoadBalancerSettings",
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"simple_lb": {
							TargetField: "SimpleLB",
						},
						"consistent_hash_lb": {
							TargetField: "ConsistentHashLB",
							ConvertType: ve.ConvertJsonObject,
							NextLevelConvert: map[string]ve.RequestConvert{
								"http_cookie": {
									TargetField: "HTTPCookie",
									ConvertType: ve.ConvertJsonObject,
								},
							},
						},
					},
				},
				"tls_settings": {
					TargetField: "TlsSettings",
					ConvertType: ve.ConvertJsonObject,
				},
				"circuit_breaking_settings": {
					TargetField: "CircuitBreakingSettings",
					ConvertType: ve.ConvertJsonObject,
				},
				"upstream_spec": {
					TargetField: "UpstreamSpec",
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"ve_faas": {
							TargetField: "VeFaas",
							ConvertType: ve.ConvertJsonObject,
						},
						"k8s_service": {
							TargetField: "K8SService",
							ConvertType: ve.ConvertJsonObject,
						},
						"ecs_list": {
							TargetField: "EcsList",
							ConvertType: ve.ConvertJsonObjectArray,
							NextLevelConvert: map[string]ve.RequestConvert{
								"ip": {
									TargetField: "IP",
								},
							},
						},
						"fixed_ip_list": {
							TargetField: "FixedIPList",
							ConvertType: ve.ConvertJsonObjectArray,
							NextLevelConvert: map[string]ve.RequestConvert{
								"ip": {
									TargetField: "IP",
								},
							},
						},
						"domain": {
							TargetField: "Domain",
							ConvertType: ve.ConvertJsonObject,
							NextLevelConvert: map[string]ve.RequestConvert{
								"domain_list": {
									TargetField: "DomainList",
									ConvertType: ve.ConvertJsonObjectArray,
								},
							},
						},
						"nacos_service": {
							TargetField: "NacosService",
							ConvertType: ve.ConvertJsonObject,
						},
						"ve_mlp": {
							TargetField: "VeMLP",
							ConvertType: ve.ConvertJsonObject,
							NextLevelConvert: map[string]ve.RequestConvert{
								"service_url": {
									TargetField: "ServiceURL",
								},
								"k8s_service": {
									TargetField: "K8SService",
									ConvertType: ve.ConvertJsonObject,
									NextLevelConvert: map[string]ve.RequestConvert{
										"cluster_info": {
											TargetField: "ClusterInfo",
											ConvertType: ve.ConvertJsonObject,
										},
									},
								},
							},
						},
						"ai_provider": {
							TargetField: "AIProvider",
							ConvertType: ve.ConvertJsonObject,
							NextLevelConvert: map[string]ve.RequestConvert{
								"custom_model_service": {
									TargetField: "CustomModelService",
									ConvertType: ve.ConvertJsonObject,
								},
								"custom_header_params": {
									TargetField: "CustomHeaderParams",
									ConvertType: ve.ConvertJsonObject,
								},
								"custom_body_params": {
									TargetField: "CustomBodyParams",
									ConvertType: ve.ConvertJsonObject,
								},
							},
						},
					},
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
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineApigUpstreamService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateUpstream",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"name": {
					TargetField: "Name",
					ForceGet:    true,
				},
				"source_type": {
					TargetField: "SourceType",
					ForceGet:    true,
				},
				"comments": {
					TargetField: "Comments",
					ForceGet:    true,
				},
				"protocol": {
					TargetField: "Protocol",
				},
				"load_balancer_settings": {
					TargetField: "LoadBalancerSettings",
					ConvertType: ve.ConvertJsonObject,
					ForceGet:    true,
					NextLevelConvert: map[string]ve.RequestConvert{
						"simple_lb": {
							TargetField: "SimpleLB",
						},
						"consistent_hash_lb": {
							TargetField: "ConsistentHashLB",
							ConvertType: ve.ConvertJsonObject,
							NextLevelConvert: map[string]ve.RequestConvert{
								"http_cookie": {
									TargetField: "HTTPCookie",
									ConvertType: ve.ConvertJsonObject,
								},
							},
						},
					},
				},
				"tls_settings": {
					TargetField: "TlsSettings",
					ConvertType: ve.ConvertJsonObject,
					ForceGet:    true,
				},
				"circuit_breaking_settings": {
					TargetField: "CircuitBreakingSettings",
					ConvertType: ve.ConvertJsonObject,
					ForceGet:    true,
				},
				"upstream_spec": {
					TargetField: "UpstreamSpec",
					ConvertType: ve.ConvertJsonObject,
					ForceGet:    true,
					NextLevelConvert: map[string]ve.RequestConvert{
						"ve_faas": {
							TargetField: "VeFaas",
							ConvertType: ve.ConvertJsonObject,
						},
						"k8s_service": {
							TargetField: "K8SService",
							ConvertType: ve.ConvertJsonObject,
						},
						"ecs_list": {
							TargetField: "EcsList",
							ConvertType: ve.ConvertJsonObjectArray,
							NextLevelConvert: map[string]ve.RequestConvert{
								"ip": {
									TargetField: "IP",
								},
							},
						},
						"fixed_ip_list": {
							TargetField: "FixedIPList",
							ConvertType: ve.ConvertJsonObjectArray,
							NextLevelConvert: map[string]ve.RequestConvert{
								"ip": {
									TargetField: "IP",
								},
							},
						},
						"domain": {
							TargetField: "Domain",
							ConvertType: ve.ConvertJsonObject,
							NextLevelConvert: map[string]ve.RequestConvert{
								"domain_list": {
									TargetField: "DomainList",
									ConvertType: ve.ConvertJsonObjectArray,
								},
							},
						},
						"nacos_service": {
							TargetField: "NacosService",
							ConvertType: ve.ConvertJsonObject,
						},
						"ve_mlp": {
							TargetField: "VeMLP",
							ConvertType: ve.ConvertJsonObject,
							NextLevelConvert: map[string]ve.RequestConvert{
								"service_url": {
									TargetField: "ServiceURL",
								},
								"k8s_service": {
									TargetField: "K8SService",
									ConvertType: ve.ConvertJsonObject,
									NextLevelConvert: map[string]ve.RequestConvert{
										"cluster_info": {
											TargetField: "ClusterInfo",
											ConvertType: ve.ConvertJsonObject,
										},
									},
								},
							},
						},
						"ai_provider": {
							TargetField: "AIProvider",
							ConvertType: ve.ConvertJsonObject,
							NextLevelConvert: map[string]ve.RequestConvert{
								"custom_model_service": {
									TargetField: "CustomModelService",
									ConvertType: ve.ConvertJsonObject,
								},
								"custom_header_params": {
									TargetField: "CustomHeaderParams",
									ConvertType: ve.ConvertJsonObject,
								},
								"custom_body_params": {
									TargetField: "CustomBodyParams",
									ConvertType: ve.ConvertJsonObject,
								},
							},
						},
					},
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["Id"] = d.Id()
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				if v, exist := (*call.SdkParam)["LoadBalancerSettings"]; exist {
					if vMap, ok := v.(map[string]interface{}); ok {
						if lbPolicy, exist := vMap["LbPolicy"]; exist {
							if lbPolicyStr, ok := lbPolicy.(string); ok {
								if lbPolicyStr == "ConsistentHashLB" {
									delete(vMap, "SimpleLB")
								}
							}
						}
						if simpleLb, exist := vMap["SimpleLB"]; exist {
							if simpleLbStr, ok := simpleLb.(string); ok {
								if simpleLbStr != "LEAST_CONN" && simpleLbStr != "ROUND_ROBIN" {
									delete(vMap, "WarmupDuration")
								}
							}
						} else {
							delete(vMap, "WarmupDuration")
						}
					}
				}
				if v, exist := (*call.SdkParam)["TlsSettings"]; exist {
					if vMap, ok := v.(map[string]interface{}); ok {
						if sni, exist := vMap["Sni"]; exist {
							if sniStr, ok := sni.(string); ok {
								if sniStr == "" {
									delete(vMap, "Sni")
								}
							}
						}
					}
				}

				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineApigUpstreamService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteUpstream",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"Id": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 5*time.Minute)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(5*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading apig upstream on delete %q, %w", d.Id(), callErr))
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

func (s *VolcengineApigUpstreamService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "Filter.Ids",
				ConvertType: ve.ConvertJsonArray,
			},
			"name": {
				TargetField: "Filter.Name",
			},
		},
		NameField:    "Name",
		IdField:      "Id",
		CollectField: "upstreams",
		ContentType:  ve.ContentTypeJson,
		ResponseConverts: map[string]ve.ResponseConvert{
			"SimpleLB": {
				TargetField: "simple_lb",
			},
			"ConsistentHashLB": {
				TargetField: "consistent_hash_lb",
			},
			"HTTPCookie": {
				TargetField: "http_cookie",
			},
			"K8SService": {
				TargetField: "k8s_service",
			},
			"IP": {
				TargetField: "ip",
			},
			"FixedIPList": {
				TargetField: "fixed_ip_list",
			},
			"VeMLP": {
				TargetField: "ve_mlp",
			},
			"ServiceURL": {
				TargetField: "service_url",
			},
			"AIProvider": {
				TargetField: "ai_provider",
			},
		},
	}
}

func (s *VolcengineApigUpstreamService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "apig",
		Version:     "2021-03-03",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}
