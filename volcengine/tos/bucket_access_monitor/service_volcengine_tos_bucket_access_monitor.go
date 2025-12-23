package tos_bucket_access_monitor

import (
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineTosBucketAccessMonitorService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewTosBucketAccessMonitorService(c *ve.SdkClient) *VolcengineTosBucketAccessMonitorService {
	return &VolcengineTosBucketAccessMonitorService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineTosBucketAccessMonitorService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineTosBucketAccessMonitorService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	return data, err
}

func (s *VolcengineTosBucketAccessMonitorService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	tos := s.Client.BypassSvcClient
	var (
		ok bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}

	action := "GetBucketAccessMonitor"
	logger.Debug(logger.ReqFormat, action, id)
	resp, err := tos.DoBypassSvcCall(ve.BypassSvcInfo{
		HttpMethod: ve.GET,
		Domain:     id,
		UrlParam: map[string]string{
			"accessmonitor": "",
		},
	}, nil)
	if err != nil {
		return data, err
	}
	logger.Debug(logger.RespFormat, action, resp, err)
	if data, ok = (*resp)[ve.BypassResponse].(map[string]interface{}); !ok {
		return data, errors.New("GetBucketAccessMonitor Resp is not map")
	}
	if len(data) == 0 {
		return data, fmt.Errorf("tos_bucket_access_monitor %s not exist ", id)
	}

	data["BucketName"] = id

	return data, err
}

func (s *VolcengineTosBucketAccessMonitorService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineTosBucketAccessMonitorService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineTosBucketAccessMonitorService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	callback := s.createOrUpdateAccessMonitor(resourceData, resource, false)
	callbacks = append(callbacks, callback)

	return callbacks
}

func (s *VolcengineTosBucketAccessMonitorService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	callback := s.createOrUpdateAccessMonitor(resourceData, resource, true)
	callbacks = append(callbacks, callback)

	return callbacks
}

func (s *VolcengineTosBucketAccessMonitorService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteBucketAccessMonitor",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["BucketName"] = d.Id()
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				param := map[string]interface{}{
					"Status": "Disabled",
				}
				resp, err := s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.PUT,
					Domain:      (*call.SdkParam)["BucketName"].(string),
					UrlParam: map[string]string{
						"accessmonitor": "",
					},
				}, &param)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				return resource.Retry(5*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading tos bucket access monitor on delete %q, %w", s.ReadResourceId(d.Id()), callErr))
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

func (s *VolcengineTosBucketAccessMonitorService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineTosBucketAccessMonitorService) createOrUpdateAccessMonitor(resourceData *schema.ResourceData, resource *schema.Resource, isUpdate bool) ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			ServiceCategory: ve.ServiceBypass,
			Action:          "PutBucketAccessMonitor",
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
				param := map[string]interface{}{
					"Status": "Enabled",
				}
				resp, err := s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.PUT,
					Domain:      (*call.SdkParam)[ve.BypassDomain].(string),
					UrlParam: map[string]string{
						"accessmonitor": "",
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

func (s *VolcengineTosBucketAccessMonitorService) ReadResourceId(id string) string {
	return id
}
