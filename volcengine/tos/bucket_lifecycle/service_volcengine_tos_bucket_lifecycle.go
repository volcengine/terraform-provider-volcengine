package tos_bucket_lifecycle

import (
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineTosBucketLifecycleService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewTosBucketLifecycleService(c *ve.SdkClient) *VolcengineTosBucketLifecycleService {
	return &VolcengineTosBucketLifecycleService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineTosBucketLifecycleService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineTosBucketLifecycleService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	return data, err
}

func (s *VolcengineTosBucketLifecycleService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	tos := s.Client.BypassSvcClient
	var (
		ok bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}

	action := "GetBucketLifecycle"
	logger.Debug(logger.ReqFormat, action, id)
	resp, err := tos.DoBypassSvcCall(ve.BypassSvcInfo{
		HttpMethod: ve.GET,
		Domain:     id,
		UrlParam: map[string]string{
			"lifecycle": "",
		},
	}, nil)
	if err != nil {
		return data, err
	}
	logger.Debug(logger.RespFormat, action, resp, err)
	if data, ok = (*resp)[ve.BypassResponse].(map[string]interface{}); !ok {
		return data, errors.New("GetBucketLifecycle Resp is not map")
	}
	if len(data) == 0 {
		return data, fmt.Errorf("tos_bucket_lifecycle %s not exist ", id)
	}

	data["BucketName"] = id
	if v, exist := data["Rules"]; exist {
		if rules, ok := v.([]interface{}); ok {
			for _, rule := range rules {
				if rule1, ok := rule.(map[string]interface{}); ok {
					if v1, exist1 := rule1["Filter"]; exist1 {
						if filter, ok1 := v1.(map[string]interface{}); ok1 {
							rule1["Filter"] = []interface{}{filter}
						}
					}
					if v1, exist1 := rule1["Expiration"]; exist1 {
						if expiration, ok1 := v1.(map[string]interface{}); ok1 {
							rule1["Expiration"] = []interface{}{expiration}
						}
					}
					if v1, exist1 := rule1["NoncurrentVersionExpiration"]; exist1 {
						if noncurrentVersionExpiration, ok1 := v1.(map[string]interface{}); ok1 {
							rule1["NoncurrentVersionExpiration"] = []interface{}{noncurrentVersionExpiration}
						}
					}
					if v1, exist1 := rule1["Transitions"]; exist1 {
						if transitions, ok1 := v1.(map[string]interface{}); ok1 {
							rule1["Transitions"] = []interface{}{transitions}
						}
					}
					if v1, exist1 := rule1["AbortIncompleteMultipartUpload"]; exist1 {
						if abortIncompleteMultipartUpload, ok1 := v1.(map[string]interface{}); ok1 {
							rule1["AbortIncompleteMultipartUpload"] = []interface{}{abortIncompleteMultipartUpload}
						}
					}
				}
			}
		}
	}

	return data, err
}

func (s *VolcengineTosBucketLifecycleService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineTosBucketLifecycleService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"Rules": {
				TargetField: "rules",
			},
			"ID": {
				TargetField: "id",
			},
			"NoncurrentVersionExpiration": {
				TargetField: "non_current_version_expiration",
			},
			"NoncurrentVersionTransitions": {
				TargetField: "non_current_version_transitions",
			},
			"NoncurrentDays": {
				TargetField: "non_current_days",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineTosBucketLifecycleService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	callback := s.createOrUpdateLifecycle(resourceData, resource, false)
	callbacks = append(callbacks, callback)

	return callbacks
}

func (s *VolcengineTosBucketLifecycleService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	callback := s.createOrUpdateLifecycle(resourceData, resource, true)
	callbacks = append(callbacks, callback)

	return callbacks
}

func (s *VolcengineTosBucketLifecycleService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteBucketLifecycle",
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
						"lifecycle": "",
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
							return resource.NonRetryableError(fmt.Errorf("error on reading tos bucket lifecycle on delete %q, %w", s.ReadResourceId(d.Id()), callErr))
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

func (s *VolcengineTosBucketLifecycleService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineTosBucketLifecycleService) createOrUpdateLifecycle(resourceData *schema.ResourceData, resource *schema.Resource, isUpdate bool) ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			ServiceCategory: ve.ServiceBypass,
			Action:          "PutBucketLifecycle",
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
						"prefix": {
							ConvertType: ve.ConvertDefault,
							TargetField: "Prefix",
							ForceGet:    isUpdate,
						},
						"status": {
							ConvertType: ve.ConvertDefault,
							TargetField: "Status",
							ForceGet:    isUpdate,
						},
						"expiration": {
							ConvertType: ve.ConvertJsonObject,
							TargetField: "Expiration",
							ForceGet:    isUpdate,
							NextLevelConvert: map[string]ve.RequestConvert{
								"days": {
									ConvertType: ve.ConvertDefault,
									TargetField: "Days",
									ForceGet:    isUpdate,
								},
								"date": {
									ConvertType: ve.ConvertDefault,
									TargetField: "Date",
									ForceGet:    isUpdate,
								},
							},
						},
						"transitions": {
							ConvertType: ve.ConvertJsonObjectArray,
							TargetField: "Transitions",
							ForceGet:    isUpdate,
							NextLevelConvert: map[string]ve.RequestConvert{
								"days": {
									ConvertType: ve.ConvertDefault,
									TargetField: "Days",
									ForceGet:    isUpdate,
								},
								"date": {
									ConvertType: ve.ConvertDefault,
									TargetField: "Date",
									ForceGet:    isUpdate,
								},
								"storage_class": {
									ConvertType: ve.ConvertDefault,
									TargetField: "StorageClass",
									ForceGet:    isUpdate,
								},
							},
						},
						"non_current_version_expiration": {
							ConvertType: ve.ConvertJsonObject,
							TargetField: "NonCurrentVersionExpiration",
							ForceGet:    isUpdate,
							NextLevelConvert: map[string]ve.RequestConvert{
								"non_current_days": {
									ConvertType: ve.ConvertDefault,
									TargetField: "NonCurrentDays",
									ForceGet:    isUpdate,
								},
							},
						},
						"non_current_version_transitions": {
							ConvertType: ve.ConvertJsonObjectArray,
							TargetField: "NonCurrentVersionTransitions",
							ForceGet:    isUpdate,
							NextLevelConvert: map[string]ve.RequestConvert{
								"non_current_days": {
									ConvertType: ve.ConvertDefault,
									TargetField: "NonCurrentDays",
									ForceGet:    isUpdate,
								},
								"storage_class": {
									ConvertType: ve.ConvertDefault,
									TargetField: "StorageClass",
									ForceGet:    isUpdate,
								},
							},
						},
						"tags": {
							ConvertType: ve.ConvertJsonObjectArray,
							TargetField: "Tags",
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
						"filter": {
							ConvertType: ve.ConvertJsonObject,
							TargetField: "Filter",
							ForceGet:    isUpdate,
							NextLevelConvert: map[string]ve.RequestConvert{
								"object_size_greater_than": {
									ConvertType: ve.ConvertDefault,
									TargetField: "ObjectSizeGreaterThan",
									ForceGet:    isUpdate,
								},
								"object_size_less_than": {
									ConvertType: ve.ConvertDefault,
									TargetField: "ObjectSizeLessThan",
									ForceGet:    isUpdate,
								},
							},
						},
						"abort_incomplete_multipart_upload": {
							ConvertType: ve.ConvertJsonObject,
							TargetField: "AbortIncompleteMultipartUpload",
							ForceGet:    isUpdate,
							NextLevelConvert: map[string]ve.RequestConvert{
								"days_after_initiation": {
									ConvertType: ve.ConvertDefault,
									TargetField: "DaysAfterInitiation",
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
						"lifecycle": "",
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

func (s *VolcengineTosBucketLifecycleService) ReadResourceId(id string) string {
	return id
}
