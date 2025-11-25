package tos_bucket_replication

import (
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineTosBucketReplicationService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewTosBucketReplicationService(c *ve.SdkClient) *VolcengineTosBucketReplicationService {
	return &VolcengineTosBucketReplicationService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineTosBucketReplicationService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineTosBucketReplicationService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	return data, err
}

func (s *VolcengineTosBucketReplicationService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	tos := s.Client.BypassSvcClient
	var (
		ok bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}

	action := "GetBucketReplication"
	logger.Debug(logger.ReqFormat, action, id)
	resp, err := tos.DoBypassSvcCall(ve.BypassSvcInfo{
		HttpMethod: ve.GET,
		Domain:     id,
		UrlParam: map[string]string{
			"replication": "",
			"progress":    "",
		},
	}, nil)
	if err != nil {
		return data, err
	}
	logger.Debug(logger.RespFormat, action, resp, err)
	if data, ok = (*resp)[ve.BypassResponse].(map[string]interface{}); !ok {
		return data, errors.New("GetBucketReplication Resp is not map")
	}
	if len(data) == 0 {
		return data, fmt.Errorf("tos_bucket_replication %s not exist ", id)
	}

	data["BucketName"] = id
	if v, exist := data["Rules"]; exist {
		if rules, ok := v.([]interface{}); ok {
			for _, rule := range rules {
				if rule1, ok := rule.(map[string]interface{}); ok {
					if v1, exist1 := rule1["Destination"]; exist1 {
						if destination, ok1 := v1.(map[string]interface{}); ok1 {
							rule1["Destination"] = []interface{}{destination}
						}
					}
					if v1, exist1 := rule1["AccessControlTranslation"]; exist1 {
						if accessControlTranslation, ok1 := v1.(map[string]interface{}); ok1 {
							rule1["AccessControlTranslation"] = []interface{}{accessControlTranslation}
						}
					}
				}
			}
		}
	}

	return data, err
}

func (s *VolcengineTosBucketReplicationService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineTosBucketReplicationService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"Role": {
				TargetField: "role",
			},
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

func (s *VolcengineTosBucketReplicationService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	callback := s.createOrUpdateReplication(resourceData, resource, false)
	callbacks = append(callbacks, callback)

	return callbacks
}

func (s *VolcengineTosBucketReplicationService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	callback := s.createOrUpdateReplication(resourceData, resource, true)
	callbacks = append(callbacks, callback)

	return callbacks
}

func (s *VolcengineTosBucketReplicationService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteBucketReplication",
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
						"replication": "",
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
							return resource.NonRetryableError(fmt.Errorf("error on reading tos bucket replication on delete %q, %w", s.ReadResourceId(d.Id()), callErr))
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

func (s *VolcengineTosBucketReplicationService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineTosBucketReplicationService) createOrUpdateReplication(resourceData *schema.ResourceData, resource *schema.Resource, isUpdate bool) ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			ServiceCategory: ve.ServiceBypass,
			Action:          "PutBucketReplication",
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
				"role": {
					ConvertType: ve.ConvertDefault,
					TargetField: "Role",
					ForceGet:    isUpdate,
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
						"status": {
							ConvertType: ve.ConvertDefault,
							TargetField: "Status",
							ForceGet:    isUpdate,
						},
						"prefix_set": {
							ConvertType: ve.ConvertJsonArray,
							TargetField: "PrefixSet",
							ForceGet:    isUpdate,
						},
						"destination": {
							ConvertType: ve.ConvertJsonObject,
							TargetField: "Destination",
							ForceGet:    isUpdate,
							NextLevelConvert: map[string]ve.RequestConvert{
								"bucket": {
									ConvertType: ve.ConvertDefault,
									TargetField: "Bucket",
									ForceGet:    isUpdate,
								},
								"location": {
									ConvertType: ve.ConvertDefault,
									TargetField: "Location",
									ForceGet:    isUpdate,
								},
								"storage_class": {
									ConvertType: ve.ConvertDefault,
									TargetField: "StorageClass",
									ForceGet:    isUpdate,
								},
								"storage_class_inherit_directive": {
									ConvertType: ve.ConvertDefault,
									TargetField: "StorageClassInheritDirective",
									ForceGet:    isUpdate,
								},
							},
						},
						"historical_object_replication": {
							ConvertType: ve.ConvertDefault,
							TargetField: "HistoricalObjectReplication",
							ForceGet:    isUpdate,
						},
						"access_control_translation": {
							ConvertType: ve.ConvertJsonObject,
							TargetField: "AccessControlTranslation",
							ForceGet:    isUpdate,
							NextLevelConvert: map[string]ve.RequestConvert{
								"owner": {
									ConvertType: ve.ConvertDefault,
									TargetField: "Owner",
									ForceGet:    isUpdate,
								},
							},
						},
						"transfer_type": {
							ConvertType: ve.ConvertDefault,
							TargetField: "TransferType",
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
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.PUT,
					Domain:      (*call.SdkParam)[ve.BypassDomain].(string),
					UrlParam: map[string]string{
						"replication": "",
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

func (s *VolcengineTosBucketReplicationService) ReadResourceId(id string) string {
	return id
}
