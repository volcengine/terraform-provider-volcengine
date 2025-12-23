package tos_bucket_object_lock_configuration

import (
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineTosBucketObjectLockConfigurationService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewTosBucketObjectLockConfigurationService(c *ve.SdkClient) *VolcengineTosBucketObjectLockConfigurationService {
	return &VolcengineTosBucketObjectLockConfigurationService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineTosBucketObjectLockConfigurationService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineTosBucketObjectLockConfigurationService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	return data, err
}

func (s *VolcengineTosBucketObjectLockConfigurationService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	tos := s.Client.BypassSvcClient
	var (
		ok bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}

	action := "GetBucketObjectLockConfiguration"
	logger.Debug(logger.ReqFormat, action, id)
	resp, err := tos.DoBypassSvcCall(ve.BypassSvcInfo{
		HttpMethod: ve.GET,
		Domain:     id,
		UrlParam: map[string]string{
			"object-lock": "",
		},
	}, nil)
	if err != nil {
		return data, err
	}
	logger.Debug(logger.RespFormat, action, resp, err)
	if data, ok = (*resp)[ve.BypassResponse].(map[string]interface{}); !ok {
		return data, errors.New("GetBucketObjectLockConfiguration Resp is not map")
	}
	if len(data) == 0 {
		return data, fmt.Errorf("tos_bucket_object_lock_configuration %s not exist ", id)
	}

	data["BucketName"] = id
	if v, exist := data["Rule"]; exist {
		if rule, ok := v.(map[string]interface{}); ok {
			if v1, exist1 := rule["DefaultRetention"]; exist1 {
				if defaultRetention, ok1 := v1.(map[string]interface{}); ok1 {
					rule["DefaultRetention"] = []interface{}{defaultRetention}
				}
			}
		}
	}

	return data, err
}

func (s *VolcengineTosBucketObjectLockConfigurationService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineTosBucketObjectLockConfigurationService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"ObjectLockConfiguration": {
				TargetField: "object_lock_configuration",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineTosBucketObjectLockConfigurationService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	callback := s.createOrUpdateObjectLockConfiguration(resourceData, resource, false)
	callbacks = append(callbacks, callback)

	return callbacks
}

func (s *VolcengineTosBucketObjectLockConfigurationService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	callback := s.createOrUpdateObjectLockConfiguration(resourceData, resource, true)
	callbacks = append(callbacks, callback)

	return callbacks
}

func (s *VolcengineTosBucketObjectLockConfigurationService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	logger.Debug(logger.ReqFormat, "RemoveResource", "Remove VolcengineTosBucketObjectLockConfigurationService dose not have callback")
	return []ve.Callback{}
}

func (s *VolcengineTosBucketObjectLockConfigurationService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineTosBucketObjectLockConfigurationService) createOrUpdateObjectLockConfiguration(resourceData *schema.ResourceData, resource *schema.Resource, isUpdate bool) ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			ServiceCategory: ve.ServiceBypass,
			Action:          "PutBucketObjectLockConfiguration",
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
				"rule": {
					ConvertType: ve.ConvertJsonObject,
					TargetField: "Rule",
					ForceGet:    isUpdate,
					NextLevelConvert: map[string]ve.RequestConvert{
						"default_retention": {
							ConvertType: ve.ConvertJsonObject,
							TargetField: "DefaultRetention",
							ForceGet:    isUpdate,
							NextLevelConvert: map[string]ve.RequestConvert{
								"mode": {
									ConvertType: ve.ConvertDefault,
									TargetField: "Mode",
									ForceGet:    isUpdate,
								},
								"days": {
									ConvertType: ve.ConvertDefault,
									TargetField: "Days",
									ForceGet:    isUpdate,
								},
								"years": {
									ConvertType: ve.ConvertDefault,
									TargetField: "Years",
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
				param["ObjectLockEnabled"] = "Enabled"
				resp, err := s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.PUT,
					Domain:      (*call.SdkParam)[ve.BypassDomain].(string),
					UrlParam: map[string]string{
						"object-lock": "",
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

func (s *VolcengineTosBucketObjectLockConfigurationService) ReadResourceId(id string) string {
	return id
}
