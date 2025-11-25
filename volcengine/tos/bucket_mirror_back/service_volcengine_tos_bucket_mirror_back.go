package tos_bucket_mirror_back

import (
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineTosBucketMirrorBackService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewTosBucketMirrorBackService(c *ve.SdkClient) *VolcengineTosBucketMirrorBackService {
	return &VolcengineTosBucketMirrorBackService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineTosBucketMirrorBackService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineTosBucketMirrorBackService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	return data, err
}

func (s *VolcengineTosBucketMirrorBackService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	tos := s.Client.BypassSvcClient
	var (
		ok bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}

	action := "GetBucketMirrorBack"
	logger.Debug(logger.ReqFormat, action, id)
	resp, err := tos.DoBypassSvcCall(ve.BypassSvcInfo{
		HttpMethod: ve.GET,
		Domain:     id,
		UrlParam: map[string]string{
			"mirror": "",
		},
	}, nil)
	if err != nil {
		return data, err
	}
	logger.Debug(logger.RespFormat, action, resp, err)
	if data, ok = (*resp)[ve.BypassResponse].(map[string]interface{}); !ok {
		return data, errors.New("GetBucketMirrorBack Resp is not map")
	}
	if len(data) == 0 {
		return data, fmt.Errorf("tos_bucket_mirror_back %s not exist ", id)
	}

	data["BucketName"] = id
	if v, exist := data["Rules"]; exist {
		if rules, ok := v.([]interface{}); ok {
			for _, rule := range rules {
				if rule1, ok := rule.(map[string]interface{}); ok {
					if v1, exist1 := rule1["Condition"]; exist1 {
						if condition, ok1 := v1.(map[string]interface{}); ok1 {
							rule1["Condition"] = []interface{}{condition}
						}
					}
					if v1, exist1 := rule1["Redirect"]; exist1 {
						if redirect, ok1 := v1.(map[string]interface{}); ok1 {
							if v2, exist2 := redirect["PublicSource"]; exist2 {
								if publicSource, ok2 := v2.(map[string]interface{}); ok2 {
									if v3, exist3 := publicSource["SourceEndpoint"]; exist3 {
										if sourceEndpoint, ok3 := v3.(map[string]interface{}); ok3 {
											publicSource["SourceEndpoint"] = []interface{}{sourceEndpoint}
										}
									}
									redirect["PublicSource"] = []interface{}{publicSource}
								}
							}
							if v2, exist2 := redirect["MirrorHeader"]; exist2 {
								if mirrorHeader, ok2 := v2.(map[string]interface{}); ok2 {
									redirect["MirrorHeader"] = []interface{}{mirrorHeader}
								}
							}
							if v2, exist2 := redirect["FetchHeaderToMetaDataRules"]; exist2 {
								if fetchHeaderToMetaDataRules, ok2 := v2.(map[string]interface{}); ok2 {
									redirect["FetchHeaderToMetaDataRules"] = []interface{}{fetchHeaderToMetaDataRules}
								}
							}
							if v2, exist2 := redirect["Transform"]; exist2 {
								if transform, ok2 := v2.(map[string]interface{}); ok2 {
									if v3, exist3 := transform["ReplaceKeyPrefix"]; exist3 {
										if replaceKeyPrefix, ok3 := v3.(map[string]interface{}); ok3 {
											transform["ReplaceKeyPrefix"] = []interface{}{replaceKeyPrefix}
										}
									}
									redirect["Transform"] = []interface{}{transform}
								}
							}
							rule1["Redirect"] = []interface{}{redirect}
						}
					}
				}
			}
		}
	}

	return data, err
}

func (s *VolcengineTosBucketMirrorBackService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineTosBucketMirrorBackService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"Rules": {
				TargetField: "rules",
			},
			"ID": {
				TargetField: "id",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineTosBucketMirrorBackService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	callback := s.createOrUpdateMirrorBack(resourceData, resource, false)
	callbacks = append(callbacks, callback)

	return callbacks
}

func (s *VolcengineTosBucketMirrorBackService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	callback := s.createOrUpdateMirrorBack(resourceData, resource, true)
	callbacks = append(callbacks, callback)

	return callbacks
}

func (s *VolcengineTosBucketMirrorBackService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteBucketMirrorBack",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["BucketName"] = d.Id()
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.DELETE,
					Domain:      (*call.SdkParam)["BucketName"].(string),
					UrlParam: map[string]string{
						"mirror": "",
					},
				}, nil)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 5*time.Minute)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				return resource.Retry(5*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading tos bucket mirror_back on delete %q, %w", s.ReadResourceId(d.Id()), callErr))
						}
					}
					_, callErr = call.ExecuteCall(d, client, call)
					if callErr == nil {
						return nil
					}
					return resource.RetryableError(callErr)
				})
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("bucket_name").(string)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineTosBucketMirrorBackService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineTosBucketMirrorBackService) createOrUpdateMirrorBack(resourceData *schema.ResourceData, resource *schema.Resource, isUpdate bool) ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			ServiceCategory: ve.ServiceBypass,
			Action:          "PutBucketMirrorBack",
			ConvertMode:     ve.RequestConvertInConvert,
			ContentType:     ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"bucket_name": {
					ConvertType: ve.ConvertDefault,
					TargetField: "BucketName",
					SpecialParam: &ve.SpecialParam{
						Type: ve.DomainParam,
					},
					ForceGet: isUpdate,
				},
				"rules": {
					ConvertType: ve.ConvertJsonObjectArray,
					TargetField: "Rules",
					ForceGet:    isUpdate,
					NextLevelConvert: map[string]ve.RequestConvert{
						"id": {
							ConvertType: ve.ConvertDefault,
							TargetField: "ID",
							ForceGet:    isUpdate,
						},
						"redirect": {
							ConvertType: ve.ConvertJsonObject,
							TargetField: "Redirect",
							ForceGet:    isUpdate,
							NextLevelConvert: map[string]ve.RequestConvert{
								"redirect_type": {
									ConvertType: ve.ConvertDefault,
									TargetField: "RedirectType",
									ForceGet:    isUpdate,
								},
								"fetch_source_on_redirect": {
									ConvertType: ve.ConvertDefault,
									TargetField: "FetchSourceOnRedirect",
									ForceGet:    isUpdate,
								},
								"fetch_source_on_redirect_with_query": {
									ConvertType: ve.ConvertDefault,
									TargetField: "FetchSourceOnRedirectWithQuery",
									ForceGet:    isUpdate,
								},
								"pass_query": {
									ConvertType: ve.ConvertDefault,
									TargetField: "PassQuery",
									ForceGet:    isUpdate,
								},
								"follow_redirect": {
									ConvertType: ve.ConvertDefault,
									TargetField: "FollowRedirect",
									ForceGet:    isUpdate,
								},
								"mirror_header": {
									ConvertType: ve.ConvertJsonObject,
									TargetField: "MirrorHeader",
									ForceGet:    isUpdate,
									NextLevelConvert: map[string]ve.RequestConvert{
										"pass_all": {
											ConvertType: ve.ConvertDefault,
											TargetField: "PassAll",
											ForceGet:    isUpdate,
										},
										"pass": {
											ConvertType: ve.ConvertJsonArray,
											TargetField: "Pass",
											ForceGet:    isUpdate,
										},
										"remove": {
											ConvertType: ve.ConvertJsonArray,
											TargetField: "Remove",
											ForceGet:    isUpdate,
										},
										"set": {
											ConvertType: ve.ConvertJsonObjectArray,
											TargetField: "Set",
											ForceGet:    isUpdate,
											NextLevelConvert: map[string]ve.RequestConvert{
												"key": {
													ConvertType: ve.ConvertDefault,
													TargetField: "Key",
													ForceGet:    isUpdate,
												},
												"value": {
													ConvertType: ve.ConvertDefault,
													TargetField: "Value",
													ForceGet:    isUpdate,
												},
											},
										},
									},
								},
								"public_source": {
									ConvertType: ve.ConvertJsonObject,
									TargetField: "PublicSource",
									ForceGet:    isUpdate,
									NextLevelConvert: map[string]ve.RequestConvert{
										"source_endpoint": {
											ConvertType: ve.ConvertJsonObject,
											TargetField: "SourceEndpoint",
											ForceGet:    isUpdate,
											NextLevelConvert: map[string]ve.RequestConvert{
												"primary": {
													ConvertType: ve.ConvertJsonArray,
													TargetField: "Primary",
													ForceGet:    isUpdate,
												},
												"follower": {
													ConvertType: ve.ConvertJsonArray,
													TargetField: "Follower",
													ForceGet:    isUpdate,
												},
											},
										},
										"fixed_endpoint": {
											ConvertType: ve.ConvertDefault,
											TargetField: "FixedEndpoint",
											ForceGet:    isUpdate,
										},
									},
								},
								"private_source": {
									ConvertType: ve.ConvertJsonObject,
									TargetField: "PrivateSource",
									ForceGet:    isUpdate,
									NextLevelConvert: map[string]ve.RequestConvert{
										"source_endpoint": {
											ConvertType: ve.ConvertJsonObject,
											TargetField: "SourceEndpoint",
											ForceGet:    isUpdate,
											NextLevelConvert: map[string]ve.RequestConvert{
												"primary": {
													ConvertType: ve.ConvertJsonObjectArray,
													TargetField: "Primary",
													ForceGet:    isUpdate,
													NextLevelConvert: map[string]ve.RequestConvert{
														"endpoint": {
															ConvertType: ve.ConvertDefault,
															TargetField: "Endpoint",
															ForceGet:    isUpdate,
														},
														"bucket_name": {
															ConvertType: ve.ConvertDefault,
															TargetField: "BucketName",
															ForceGet:    isUpdate,
														},
														"credential_provider": {
															ConvertType: ve.ConvertJsonObject,
															TargetField: "CredentialProvider",
															ForceGet:    isUpdate,
															NextLevelConvert: map[string]ve.RequestConvert{
																"role": {
																	ConvertType: ve.ConvertDefault,
																	TargetField: "Role",
																	ForceGet:    isUpdate,
																},
															},
														},
													},
												},
												"follower": {
													ConvertType: ve.ConvertJsonObjectArray,
													TargetField: "Follower",
													ForceGet:    isUpdate,
													NextLevelConvert: map[string]ve.RequestConvert{
														"endpoint": {
															ConvertType: ve.ConvertDefault,
															TargetField: "Endpoint",
															ForceGet:    isUpdate,
														},
														"bucket_name": {
															ConvertType: ve.ConvertDefault,
															TargetField: "BucketName",
															ForceGet:    isUpdate,
														},
														"credential_provider": {
															ConvertType: ve.ConvertJsonObject,
															TargetField: "CredentialProvider",
															ForceGet:    isUpdate,
															NextLevelConvert: map[string]ve.RequestConvert{
																"role": {
																	ConvertType: ve.ConvertDefault,
																	TargetField: "Role",
																	ForceGet:    isUpdate,
																},
															},
														},
													},
												},
											},
										},
									},
								},
								"transform": {
									ConvertType: ve.ConvertJsonObject,
									TargetField: "Transform",
									ForceGet:    isUpdate,
									NextLevelConvert: map[string]ve.RequestConvert{
										"with_key_prefix": {
											ConvertType: ve.ConvertDefault,
											TargetField: "WithKeyPrefix",
											ForceGet:    isUpdate,
										},
										"with_key_suffix": {
											ConvertType: ve.ConvertDefault,
											TargetField: "WithKeySuffix",
											ForceGet:    isUpdate,
										},
										"replace_key_prefix": {
											ConvertType: ve.ConvertJsonObject,
											TargetField: "ReplaceKeyPrefix",
											ForceGet:    isUpdate,
											NextLevelConvert: map[string]ve.RequestConvert{
												"key_prefix": {
													ConvertType: ve.ConvertDefault,
													TargetField: "KeyPrefix",
													ForceGet:    isUpdate,
												},
												"replace_with": {
													ConvertType: ve.ConvertDefault,
													TargetField: "ReplaceWith",
													ForceGet:    isUpdate,
												},
											},
										},
										"remove_key_prefix": {
											ConvertType: ve.ConvertDefault,
											TargetField: "RemoveKeyPrefix",
											ForceGet:    isUpdate,
										},
									},
								},
							},
						},
						"condition": {
							ConvertType: ve.ConvertJsonObject,
							TargetField: "Condition",
							ForceGet:    isUpdate,
							NextLevelConvert: map[string]ve.RequestConvert{
								"http_code": {
									ConvertType: ve.ConvertDefault,
									TargetField: "HttpCode",
									ForceGet:    isUpdate,
								},
								"key_prefix": {
									ConvertType: ve.ConvertDefault,
									TargetField: "KeyPrefix",
									ForceGet:    isUpdate,
								},
								"key_suffix": {
									ConvertType: ve.ConvertDefault,
									TargetField: "KeySuffix",
									ForceGet:    isUpdate,
								},
								"allow_host": {
									ConvertType: ve.ConvertJsonArray,
									TargetField: "AllowHost",
									ForceGet:    isUpdate,
								},
								"http_method": {
									ConvertType: ve.ConvertJsonArray,
									TargetField: "HttpMethod",
									ForceGet:    isUpdate,
								},
							},
						},
					},
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				var sourceParam map[string]interface{}
				sourceParam, err := ve.SortAndStartTransJson((*call.SdkParam)[ve.BypassParam].(map[string]interface{}))
				if err != nil {
					return false, err
				}
				(*call.SdkParam)[ve.BypassParam] = sourceParam

				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				param := (*call.SdkParam)[ve.BypassParam].(map[string]interface{})
				resp, err := s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.PUT,
					Domain:      (*call.SdkParam)[ve.BypassDomain].(string),
					UrlParam: map[string]string{
						"mirror": "",
					},
				}, &param)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				d.SetId((*call.SdkParam)[ve.BypassDomain].(string))
				return nil
			},
		},
	}

	return callback
}

func (s *VolcengineTosBucketMirrorBackService) ReadResourceId(id string) string {
	return id
}
