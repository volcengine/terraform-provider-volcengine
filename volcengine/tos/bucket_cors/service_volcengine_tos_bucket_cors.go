package tos_bucket_cors

import (
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineTosBucketCorsService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewTosBucketCorsService(c *ve.SdkClient) *VolcengineTosBucketCorsService {
	return &VolcengineTosBucketCorsService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineTosBucketCorsService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineTosBucketCorsService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	return data, err
}

func (s *VolcengineTosBucketCorsService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	tos := s.Client.BypassSvcClient
	var (
		ok bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}

	action := "GetBucketCORS"
	logger.Debug(logger.ReqFormat, action, id)
	resp, err := tos.DoBypassSvcCall(ve.BypassSvcInfo{
		HttpMethod: ve.GET,
		Domain:     id,
		UrlParam: map[string]string{
			"cors": "",
		},
	}, nil)
	if err != nil {
		return data, err
	}
	logger.Debug(logger.RespFormat, action, resp, err)
	if data, ok = (*resp)[ve.BypassResponse].(map[string]interface{}); !ok {
		return data, errors.New("GetBucketCORS Resp is not map ")
	}
	if len(data) == 0 {
		return data, fmt.Errorf("tos_bucket_cors %s not exist ", id)
	}

	data["BucketName"] = id

	return data, err
}

func (s *VolcengineTosBucketCorsService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (VolcengineTosBucketCorsService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"CORSRules": {
				TargetField: "cors_rules",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineTosBucketCorsService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	callback := s.createOrUpdateCors(resourceData, resource, false)
	callbacks = append(callbacks, callback)

	return callbacks
}

func (s *VolcengineTosBucketCorsService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	callback := s.createOrUpdateCors(resourceData, resource, true)
	callbacks = append(callbacks, callback)

	return callbacks
}

func (s *VolcengineTosBucketCorsService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteBucketCORS",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["BucketName"] = d.Id()
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.DELETE,
					Domain:      (*call.SdkParam)["BucketName"].(string),
					UrlParam: map[string]string{
						"cors": "",
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
							return resource.NonRetryableError(fmt.Errorf("error on reading tos bucket cors on delete %q, %w", s.ReadResourceId(d.Id()), callErr))
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

func (s *VolcengineTosBucketCorsService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineTosBucketCorsService) createOrUpdateCors(resourceData *schema.ResourceData, resource *schema.Resource, isUpdate bool) ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			ServiceCategory: ve.ServiceBypass,
			Action:          "PutBucketCORS",
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
				"cors_rules": {
					ConvertType: ve.ConvertJsonObjectArray,
					TargetField: "CORSRules",
					ForceGet:    isUpdate,
					NextLevelConvert: map[string]ve.RequestConvert{
						"allowed_origins": {
							ConvertType: ve.ConvertJsonArray,
							TargetField: "AllowedOrigins",
							ForceGet:    isUpdate,
						},
						"allowed_methods": {
							ConvertType: ve.ConvertJsonArray,
							TargetField: "AllowedMethods",
							ForceGet:    isUpdate,
						},
						"allowed_headers": {
							ConvertType: ve.ConvertJsonArray,
							TargetField: "AllowedHeaders",
							ForceGet:    isUpdate,
						},
						"expose_headers": {
							ConvertType: ve.ConvertJsonArray,
							TargetField: "ExposeHeaders",
							ForceGet:    isUpdate,
						},
						"max_age_seconds": {
							ConvertType: ve.ConvertDefault,
							TargetField: "MaxAgeSeconds",
							ForceGet:    isUpdate,
						},
						"response_vary": {
							ConvertType: ve.ConvertDefault,
							TargetField: "ResponseVary",
							ForceGet:    isUpdate,
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
					HttpMethod:  ve.PUT,
					ContentType: ve.ApplicationJSON,
					Domain:      (*call.SdkParam)[ve.BypassDomain].(string),
					UrlParam: map[string]string{
						"cors": "",
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

func (s *VolcengineTosBucketCorsService) ReadResourceId(id string) string {
	return id
}
