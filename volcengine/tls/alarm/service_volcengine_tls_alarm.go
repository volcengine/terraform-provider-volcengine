package alarm

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineTlsAlarmService struct {
	Client *ve.SdkClient
}

func (v *VolcengineTlsAlarmService) GetClient() *ve.SdkClient {
	return v.Client
}

func (v *VolcengineTlsAlarmService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 20, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeAlarms"
		logger.Debug(logger.ReqFormat, action, condition)
		resp, err = v.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
			ContentType: ve.Default,
			HttpMethod:  ve.GET,
			Path:        []string{action},
			Client:      v.Client.BypassSvcClient.NewTlsClient(),
		}, &m)
		logger.Debug(logger.RespFormat, action, resp)
		results, err = ve.ObtainSdkValue("RESPONSE.Alarms", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.Alarms is not Slice")
		}
		return data, err
	})
}

func (v *VolcengineTlsAlarmService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = v.ReadResourceId(resourceData.Id())
	}
	ids := strings.Split(id, ":")
	if len(ids) != 2 {
		return nil, errors.New("invalid id")
	}
	req := map[string]interface{}{
		"ProjectId": ids[0],
		"AlarmId":   ids[1],
	}
	results, err = v.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, r := range results {
		if data, ok = r.(map[string]interface{}); !ok {
			return data, errors.New("value is not map")
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("tls alarm %s not exist ", id)
	}
	groups := data["AlarmNotifyGroup"]
	if groups != nil {
		groupIds := make([]string, 0)
		if _, ok = groups.([]interface{}); !ok {
			return data, fmt.Errorf("groups value is not slice")
		}
		for _, group := range groups.([]interface{}) {
			groupMap, ok := group.(map[string]interface{})
			if !ok {
				return data, fmt.Errorf("group value is not map")
			}
			groupIds = append(groupIds, groupMap["AlarmNotifyGroupId"].(string))
		}
		data["AlarmNotifyGroup"] = groupIds
	}
	logger.Debug(logger.ReqFormat, "ReadResource", data)
	return data, err
}

func (v *VolcengineTlsAlarmService) RefreshResourceState(data *schema.ResourceData, strings []string, duration time.Duration, s string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (v *VolcengineTlsAlarmService) WithResourceResponseHandlers(m map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return m, map[string]ve.ResponseConvert{
			"SMS": {
				TargetField: "sms",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (v *VolcengineTlsAlarmService) CreateResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateAlarm",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"alarm_notify_group": {
					ConvertType: ve.ConvertJsonArray,
				},
				"query_request": {
					ConvertType: ve.ConvertJsonObjectArray,
				},
				"request_cycle": {
					ConvertType: ve.ConvertJsonObject,
				},
				"alarm_period_detail": {
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"sms": {
							TargetField: "SMS",
						},
					},
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return v.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.POST,
					Path:        []string{call.Action},
					Client:      v.Client.BypassSvcClient.NewTlsClient(),
				}, call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("RESPONSE.AlarmId", *resp)
				d.SetId(fmt.Sprint(d.Get("project_id").(string), ":", id.(string)))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (v *VolcengineTlsAlarmService) ModifyResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyAlarm",
			ConvertMode: ve.RequestConvertInConvert,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"alarm_name": {
					ConvertType: ve.ConvertDefault,
				},
				"status": {
					ConvertType: ve.ConvertDefault,
				},
				"condition": {
					ConvertType: ve.ConvertDefault,
				},
				"alarm_period": {
					ConvertType: ve.ConvertDefault,
				},
				"trigger_period": {
					ConvertType: ve.ConvertDefault,
				},
				"query_request": {
					ConvertType: ve.ConvertJsonObjectArray,
					NextLevelConvert: map[string]ve.RequestConvert{
						"topic_id": {
							ConvertType: ve.ConvertDefault,
							ForceGet:    true,
						},
						"query": {
							ConvertType: ve.ConvertDefault,
							ForceGet:    true,
						},
						"number": {
							ConvertType: ve.ConvertDefault,
							ForceGet:    true,
						},
						"start_time_offset": {
							ConvertType: ve.ConvertDefault,
							ForceGet:    true,
						},
						"end_time_offset": {
							ConvertType: ve.ConvertDefault,
							ForceGet:    true,
						},
					},
				},
				"request_cycle": {
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"type": {
							ConvertType: ve.ConvertDefault,
							ForceGet:    true,
						},
						"time": {
							ConvertType: ve.ConvertDefault,
							ForceGet:    true,
						},
					},
				},
				"alarm_period_detail": {
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"sms": {
							TargetField: "SMS",
							ForceGet:    true,
						},
						"phone": {
							ConvertType: ve.ConvertDefault,
							ForceGet:    true,
						},
						"email": {
							ConvertType: ve.ConvertDefault,
							ForceGet:    true,
						},
						"general_webhook": {
							ConvertType: ve.ConvertDefault,
							ForceGet:    true,
						},
					},
				},
				"user_define_msg": {
					ConvertType: ve.ConvertDefault,
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return v.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.PUT,
					Path:        []string{call.Action},
					Client:      v.Client.BypassSvcClient.NewTlsClient(),
				}, call.SdkParam)
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				ids := strings.Split(d.Id(), ":")
				(*call.SdkParam)["AlarmId"] = ids[1]
				return true, nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (v *VolcengineTlsAlarmService) RemoveResource(data *schema.ResourceData, re *schema.Resource) []ve.Callback {
	ids := strings.Split(data.Id(), ":")
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteAlarm",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"AlarmId": ids[1],
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return v.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.DELETE,
					Path:        []string{call.Action},
					Client:      v.Client.BypassSvcClient.NewTlsClient(),
				}, call.SdkParam)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := v.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading tls alarm on delete %q, %w", d.Id(), callErr))
						}
					}
					_, callErr = call.ExecuteCall(d, client, call)
					if callErr == nil {
						return nil
					}
					return resource.RetryableError(callErr)
				})
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, v.ReadResource, 3*time.Minute)
			},
		},
	}
	return []ve.Callback{callback}
}

func (v *VolcengineTlsAlarmService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		CollectField: "alarms",
		NameField:    "AlarmName",
		IdField:      "AlarmId",
		ResponseConverts: map[string]ve.ResponseConvert{
			"SMS": {
				TargetField: "sms",
			},
		},
	}
}

func (v *VolcengineTlsAlarmService) ReadResourceId(s string) string {
	return s
}

func NewVolcengineTlsAlarmService(client *ve.SdkClient) *VolcengineTlsAlarmService {
	return &VolcengineTlsAlarmService{
		Client: client,
	}
}
