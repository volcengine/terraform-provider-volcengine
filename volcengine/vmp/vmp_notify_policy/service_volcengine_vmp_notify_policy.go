package vmp_notify_policy

import (
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineVmpNotifyPolicyService struct {
	Client *ve.SdkClient
}

func (v *VolcengineVmpNotifyPolicyService) GetClient() *ve.SdkClient {
	return v.Client
}

func (v *VolcengineVmpNotifyPolicyService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 10, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListNotifyPolicies"
		logger.Debug(logger.ReqFormat, action, condition)
		if condition == nil {
			resp, err = v.Client.UniversalClient.DoCall(getUniversalInfo(action), nil)
			if err != nil {
				return data, err
			}
		} else {
			resp, err = v.Client.UniversalClient.DoCall(getUniversalInfo(action), &condition)
			if err != nil {
				return data, err
			}
		}
		logger.Debug(logger.ReqFormat, action, *resp)
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
		return data, nil
	})
}

func (v *VolcengineVmpNotifyPolicyService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = v.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"Filter": map[string]interface{}{
			"Ids": []interface{}{id},
		},
	}
	results, err = v.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, r := range results {
		if data, ok = r.(map[string]interface{}); !ok {
			return data, errors.New("Value is not map ")
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("Notify Policy %s not exist ", id)
	}
	return data, err
}

func (v *VolcengineVmpNotifyPolicyService) RefreshResourceState(data *schema.ResourceData, strings []string, duration time.Duration, s string) *resource.StateChangeConf {
	return nil
}

func (v *VolcengineVmpNotifyPolicyService) WithResourceResponseHandlers(m map[string]interface{}) []ve.ResourceResponseHandler {
	return []ve.ResourceResponseHandler{}
}

func (v *VolcengineVmpNotifyPolicyService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateNotifyPolicy",
			ContentType: ve.ContentTypeJson,
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"channel_notify_template_ids": {
					TargetField: "ChannelNotifyTemplateIds",
					ConvertType: ve.ConvertJsonArray,
				},
				"levels": {
					ConvertType: ve.ConvertJsonObjectArray,
					NextLevelConvert: map[string]ve.RequestConvert{
						"contact_group_ids": {
							ConvertType: ve.ConvertJsonArray,
						},
						"channels": {
							ConvertType: ve.ConvertJsonArray,
						},
						"resolved_channels": {
							ConvertType: ve.ConvertJsonArray,
						},
					},
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return v.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
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

func (v *VolcengineVmpNotifyPolicyService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateNotifyPolicy",
			ContentType: ve.ContentTypeJson,
			ConvertMode: ve.RequestConvertInConvert,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["Id"] = d.Id()
				return true, nil
			},
			Convert: map[string]ve.RequestConvert{
				"name": {
					TargetField: "Name",
				},
				"description": {
					TargetField: "Description",
				},
				"channel_notify_template_ids": {
					TargetField: "ChannelNotifyTemplateIds",
					ConvertType: ve.ConvertJsonArray,
					ForceGet:    true,
				},
				"levels": {
					TargetField: "Levels",
					ForceGet:    true,
					ConvertType: ve.ConvertJsonObjectArray,
					NextLevelConvert: map[string]ve.RequestConvert{
						"level": {
							ForceGet: true,
						},
						"contact_group_ids": {
							ConvertType: ve.ConvertJsonArray,
							ForceGet:    true,
						},
						"channels": {
							ConvertType: ve.ConvertJsonArray,
							ForceGet:    true,
						},
						"resolved_channels": {
							ConvertType: ve.ConvertJsonArray,
							ForceGet:    true,
						},
					},
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return v.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	return []ve.Callback{callback}
}

func (v *VolcengineVmpNotifyPolicyService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteNotifyPolicy",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"Id": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return v.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, v.ReadResource, 5*time.Minute)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(5*time.Minute, func() *resource.RetryError {
					_, callErr := v.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading vmp notify policy on delete %q, %w", d.Id(), callErr))
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

func (v *VolcengineVmpNotifyPolicyService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "Filter.Ids",
				ConvertType: ve.ConvertJsonArray,
			},
			"name": {
				TargetField: "Filter.Name",
			},
			"contact_group_ids": {
				TargetField: "Filter.ContactGroupIds",
				ConvertType: ve.ConvertJsonArray,
			},
			"channel_notify_template_ids": {
				TargetField: "Filter.ChannelNotifyTemplateIds",
				ConvertType: ve.ConvertJsonArray,
			},
		},
		ContentType:  ve.ContentTypeJson,
		NameField:    "Name",
		IdField:      "Id",
		CollectField: "notify_policies",
	}
}

func (v *VolcengineVmpNotifyPolicyService) ReadResourceId(s string) string {
	return s
}

func NewService(c *ve.SdkClient) *VolcengineVmpNotifyPolicyService {
	return &VolcengineVmpNotifyPolicyService{
		Client: c,
	}
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "vmp",
		Version:     "2021-03-03",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}
