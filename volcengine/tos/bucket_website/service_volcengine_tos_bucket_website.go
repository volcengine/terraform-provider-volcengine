package tos_bucket_website

import (
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineTosBucketWebsiteService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewTosBucketWebsiteService(c *ve.SdkClient) *VolcengineTosBucketWebsiteService {
	return &VolcengineTosBucketWebsiteService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineTosBucketWebsiteService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineTosBucketWebsiteService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	return data, err
}

func (s *VolcengineTosBucketWebsiteService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	tos := s.Client.BypassSvcClient
	var (
		ok bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}

	action := "GetBucketWebsite"
	logger.Debug(logger.ReqFormat, action, id)
	resp, err := tos.DoBypassSvcCall(ve.BypassSvcInfo{
		HttpMethod: ve.GET,
		Domain:     id,
		UrlParam: map[string]string{
			"website": "",
		},
	}, nil)
	if err != nil {
		return data, err
	}
	logger.Debug(logger.RespFormat, action, resp, err)
	if data, ok = (*resp)[ve.BypassResponse].(map[string]interface{}); !ok {
		return data, errors.New("GetBucketWebsite Resp is not map")
	}
	if len(data) == 0 {
		return data, fmt.Errorf("tos_bucket_website %s not exist ", id)
	}

	data["BucketName"] = id

	if v, exist := data["RoutingRules"]; exist {
		if routingRules, ok := v.([]interface{}); ok {
			for _, rule := range routingRules {
				if rule1, ok := rule.(map[string]interface{}); ok {
					if v1, exist1 := rule1["Condition"]; exist1 {
						if condition, ok1 := v1.(map[string]interface{}); ok1 {
							rule1["Condition"] = []interface{}{condition}
						}
					}
					if v1, exist1 := rule1["Redirect"]; exist1 {
						if redirect, ok1 := v1.(map[string]interface{}); ok1 {
							if v2, ok := redirect["ReplaceKeyPrefixWith"]; ok {
								redirect["ReplaceKeyPrefixWith"] = v2
							}
							if v2, ok := redirect["ReplaceKeyWith"]; ok {
								redirect["ReplaceKeyWith"] = v2
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

func (s *VolcengineTosBucketWebsiteService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineTosBucketWebsiteService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"IndexDocument": {
				TargetField: "index_document",
			},
			"ErrorDocument": {
				TargetField: "error_document",
			},
			"RedirectAllRequestsTo": {
				TargetField: "redirect_all_requests_to",
			},
			"RoutingRules": {
				TargetField: "routing_rules",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineTosBucketWebsiteService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	callback := s.createOrUpdateWebsite(resourceData, resource, false)
	callbacks = append(callbacks, callback)

	return callbacks
}

func (s *VolcengineTosBucketWebsiteService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	callback := s.createOrUpdateWebsite(resourceData, resource, true)
	callbacks = append(callbacks, callback)

	return callbacks
}

func (s *VolcengineTosBucketWebsiteService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteBucketWebsite",
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
						"website": "",
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
							return resource.NonRetryableError(fmt.Errorf("error on reading tos bucket website on delete %q, %w", s.ReadResourceId(d.Id()), callErr))
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

func (s *VolcengineTosBucketWebsiteService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineTosBucketWebsiteService) createOrUpdateWebsite(resourceData *schema.ResourceData, resource *schema.Resource, isUpdate bool) ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			ServiceCategory: ve.ServiceBypass,
			Action:          "PutBucketWebsite",
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
				"index_document": {
					ConvertType: ve.ConvertJsonObject,
					TargetField: "IndexDocument",
					ForceGet:    isUpdate,
					NextLevelConvert: map[string]ve.RequestConvert{
						"suffix": {
							ConvertType: ve.ConvertDefault,
							TargetField: "Suffix",
							ForceGet:    isUpdate,
						},
						"support_sub_dir": {
							ConvertType: ve.ConvertDefault,
							TargetField: "SupportSubDir",
							ForceGet:    isUpdate,
						},
					},
				},
				"error_document": {
					ConvertType: ve.ConvertJsonObject,
					TargetField: "ErrorDocument",
					ForceGet:    isUpdate,
					NextLevelConvert: map[string]ve.RequestConvert{
						"key": {
							ConvertType: ve.ConvertDefault,
							TargetField: "Key",
							ForceGet:    isUpdate,
						},
					},
				},
				"redirect_all_requests_to": {
					ConvertType: ve.ConvertJsonObject,
					TargetField: "RedirectAllRequestsTo",
					ForceGet:    isUpdate,
					NextLevelConvert: map[string]ve.RequestConvert{
						"host_name": {
							ConvertType: ve.ConvertDefault,
							TargetField: "HostName",
							ForceGet:    isUpdate,
						},
						"protocol": {
							ConvertType: ve.ConvertDefault,
							TargetField: "Protocol",
							ForceGet:    isUpdate,
						},
					},
				},
				"routing_rules": {
					ConvertType: ve.ConvertJsonObjectArray,
					TargetField: "RoutingRules",
					NextLevelConvert: map[string]ve.RequestConvert{
						"condition": {
							ConvertType: ve.ConvertJsonObject,
							TargetField: "Condition",
							ForceGet:    isUpdate,
							NextLevelConvert: map[string]ve.RequestConvert{
								"key_prefix_equals": {
									ConvertType: ve.ConvertDefault,
									TargetField: "KeyPrefixEquals",
									ForceGet:    isUpdate,
								},
								"http_error_code_returned_equals": {
									ConvertType: ve.ConvertDefault,
									TargetField: "HttpErrorCodeReturnedEquals",
									ForceGet:    isUpdate,
								},
							},
						},
						"redirect": {
							ConvertType: ve.ConvertJsonObject,
							TargetField: "Redirect",
							NextLevelConvert: map[string]ve.RequestConvert{
								"protocol": {
									ConvertType: ve.ConvertDefault,
									TargetField: "Protocol",
									ForceGet:    isUpdate,
								},
								"host_name": {
									ConvertType: ve.ConvertDefault,
									TargetField: "HostName",
									ForceGet:    isUpdate,
								},
								"replace_key_with": {
									ConvertType: ve.ConvertDefault,
									TargetField: "ReplaceKeyWith",
								},
								"replace_key_prefix_with": {
									ConvertType: ve.ConvertDefault,
									TargetField: "ReplaceKeyPrefixWith",
								},
								"http_redirect_code": {
									ConvertType: ve.ConvertDefault,
									TargetField: "HttpRedirectCode",
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
						"website": "",
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

func (s *VolcengineTosBucketWebsiteService) ReadResourceId(id string) string {
	return id
}
