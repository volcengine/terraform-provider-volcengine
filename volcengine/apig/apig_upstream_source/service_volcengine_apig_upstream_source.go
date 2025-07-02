package apig_upstream_source

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

type VolcengineApigUpstreamSourceService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewApigUpstreamSourceService(c *ve.SdkClient) *VolcengineApigUpstreamSourceService {
	return &VolcengineApigUpstreamSourceService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineApigUpstreamSourceService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineApigUpstreamSourceService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListUpstreamSources"

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

func (s *VolcengineApigUpstreamSourceService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		result interface{}
		ok     bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}

	action := "GetUpstreamSource"
	req := map[string]interface{}{
		"Id": id,
	}
	logger.Debug(logger.ReqFormat, action, req)
	resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(action), &req)
	if err != nil {
		return data, err
	}
	logger.Debug(logger.RespFormat, action, req, *resp)

	result, err = ve.ObtainSdkValue("Result.UpstreamSource", *resp)
	if err != nil {
		return data, err
	}

	if data, ok = result.(map[string]interface{}); !ok {
		return data, errors.New("Value is not map ")
	}
	if len(data) == 0 {
		return data, fmt.Errorf("apig_upstream_source %s not exist ", id)
	}

	if sourceSpec, ok := data["SourceSpec"]; ok {
		if sourceSpecMap, ok := sourceSpec.(map[string]interface{}); ok {
			if k8sSource, ok := sourceSpecMap["K8SSource"]; ok {
				if k8sSourceMap, ok := k8sSource.(map[string]interface{}); ok {
					sourceSpecMap["K8SSource"] = []interface{}{k8sSourceMap}
				}
			}
			if nacosSource, ok := sourceSpecMap["NacosSource"]; ok {
				if nacosSourceMap, ok := nacosSource.(map[string]interface{}); ok {
					if authConfig, ok := resourceData.GetOk("source_spec.0.nacos_source.0.auth_config"); ok {
						nacosSourceMap["AuthConfig"] = authConfig
					}
					sourceSpecMap["NacosSource"] = []interface{}{nacosSourceMap}
				}
			}
		}
	}

	return data, err
}

func (s *VolcengineApigUpstreamSourceService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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
			failStates = append(failStates, "SyncedFailed")
			d, err = s.ReadResource(resourceData, id)
			if err != nil {
				return nil, "", err
			}
			status, err = ve.ObtainSdkValue("Status", d)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("apig_upstream_source status error, status: %s", status.(string))
				}
			}
			return d, status.(string), err
		},
	}
}

func (VolcengineApigUpstreamSourceService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"K8SSource": {
				TargetField: "k8s_source",
			},
			"GRPCPort": {
				TargetField: "grpc_port",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineApigUpstreamSourceService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateUpstreamSource",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"source_spec": {
					TargetField: "SourceSpec",
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"k8s_source": {
							TargetField: "K8SSource",
							ConvertType: ve.ConvertJsonObject,
						},
						"nacos_source": {
							TargetField: "NacosSource",
							ConvertType: ve.ConvertJsonObject,
							NextLevelConvert: map[string]ve.RequestConvert{
								"auth_config": {
									TargetField: "AuthConfig",
									ConvertType: ve.ConvertJsonObject,
									NextLevelConvert: map[string]ve.RequestConvert{
										"basic": {
											TargetField: "Basic",
											ConvertType: ve.ConvertJsonObject,
										},
									},
								},
							},
						},
					},
				},
				"ingress_settings": {
					TargetField: "IngressSettings",
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"enable_ingress": {
							TargetField: "EnableIngress",
							ForceGet:    true,
						},
						"enable_all_ingress_classes": {
							TargetField: "EnableAllIngressClasses",
							ForceGet:    true,
						},
						"enable_ingress_without_ingress_class": {
							TargetField: "EnableIngressWithoutIngressClass",
							ForceGet:    true,
						},
						"enable_all_namespaces": {
							TargetField: "EnableAllNamespaces",
							ForceGet:    true,
						},
						"update_status": {
							TargetField: "UpdateStatus",
							ForceGet:    true,
						},
						"ingress_classes": {
							TargetField: "IngressClasses",
							ConvertType: ve.ConvertJsonArray,
						},
						"watch_namespaces": {
							TargetField: "WatchNamespaces",
							ConvertType: ve.ConvertJsonArray,
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
			Refresh: &ve.StateRefresh{
				Target:  []string{"SyncedSucceed"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineApigUpstreamSourceService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateUpstreamSource",
			ConvertMode: ve.RequestConvertInConvert,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"comments": {
					TargetField: "Comments",
					ForceGet:    true,
				},
				"ingress_settings": {
					TargetField: "IngressSettings",
					ConvertType: ve.ConvertJsonObject,
					ForceGet:    true,
					NextLevelConvert: map[string]ve.RequestConvert{
						"enable_ingress": {
							TargetField: "EnableIngress",
							ForceGet:    true,
						},
						"enable_all_ingress_classes": {
							TargetField: "EnableAllIngressClasses",
							ForceGet:    true,
						},
						"enable_ingress_without_ingress_class": {
							TargetField: "EnableIngressWithoutIngressClass",
							ForceGet:    true,
						},
						"enable_all_namespaces": {
							TargetField: "EnableAllNamespaces",
							ForceGet:    true,
						},
						"update_status": {
							TargetField: "UpdateStatus",
							ForceGet:    true,
						},
						"ingress_classes": {
							TargetField: "IngressClasses",
							ConvertType: ve.ConvertJsonArray,
						},
						"watch_namespaces": {
							TargetField: "WatchNamespaces",
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
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"SyncedSucceed"},
				Timeout: resourceData.Timeout(schema.TimeoutUpdate),
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineApigUpstreamSourceService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteUpstreamSource",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"Id": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
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
							return resource.NonRetryableError(fmt.Errorf("error on reading apig upstream source on delete %q, %w", d.Id(), callErr))
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

func (s *VolcengineApigUpstreamSourceService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"name": {
				TargetField: "Filter.Name",
			},
			"source_type": {
				TargetField: "Filter.SourceType",
			},
			"status": {
				TargetField: "Filter.Status",
			},
			"enable_ingress": {
				TargetField: "IngressSettingsFilter.EnableIngress",
			},
		},
		IdField:      "Id",
		CollectField: "upstream_sources",
		ContentType:  ve.ContentTypeJson,
		ResponseConverts: map[string]ve.ResponseConvert{
			"K8SSource": {
				TargetField: "k8s_source",
			},
			"GRPCPort": {
				TargetField: "grpc_port",
			},
		},
	}
}

func (s *VolcengineApigUpstreamSourceService) ReadResourceId(id string) string {
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
