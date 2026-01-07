package alarm_content_template

import (
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineTlsAlarmContentTemplateService struct {
	Client *ve.SdkClient
}

func NewTlsAlarmContentTemplateService(c *ve.SdkClient) *VolcengineTlsAlarmContentTemplateService {
	return &VolcengineTlsAlarmContentTemplateService{
		Client: c,
	}
}

func (v *VolcengineTlsAlarmContentTemplateService) GetClient() *ve.SdkClient {
	return v.Client
}

func (v *VolcengineTlsAlarmContentTemplateService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 20, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeAlarmContentTemplates"
		logger.Debug(logger.ReqFormat, action, condition)
		resp, err = v.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
			ContentType: ve.Default,
			HttpMethod:  ve.GET,
			Path:        []string{action},
			Client:      v.Client.BypassSvcClient.NewTlsClient(),
		}, &m)
		logger.Debug(logger.RespFormat, action, resp)
		results, err = ve.ObtainSdkValue("RESPONSE.AlarmContentTemplates", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, fmt.Errorf("AlarmContentTemplates is not Slice")
		}
		return data, err
	})
}

func (v *VolcengineTlsAlarmContentTemplateService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = v.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"AlarmContentTemplateId": id,
	}
	results, err = v.ReadResources(req)
	if err != nil {
		return data, err
	}
	logger.Debug(logger.RespFormat, "=====DescribeAlarmContentTemplates", req, results)
	for _, r := range results {
		if data, ok = r.(map[string]interface{}); !ok {
			return data, errors.New("value is not map")
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("tls alarm content template %s not exist ", id)
	}
	return data, err
}

func (v *VolcengineTlsAlarmContentTemplateService) RefreshResourceState(data *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (v *VolcengineTlsAlarmContentTemplateService) WithResourceResponseHandlers(m map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return m, map[string]ve.ResponseConvert{
			"AlarmContentTemplateId": {
				TargetField: "alarm_content_template_id",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (v *VolcengineTlsAlarmContentTemplateService) CreateResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	createAlarmContentTemplateCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateAlarmContentTemplate",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"alarm_content_template_name": {
					TargetField: "AlarmContentTemplateName",
				},
				"sms": {
					TargetField: "Sms",
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"content": {
							TargetField: "Content",
						},
						"locale": {
							TargetField: "Locale",
						},
					},
				},
				"vms": {
					TargetField: "Vms",
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"content": {
							TargetField: "Content",
						},
						"locale": {
							TargetField: "Locale",
						},
					},
				},
				"lark": {
					TargetField: "Lark",
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"content": {
							TargetField: "Content",
						},
						"locale": {
							TargetField: "Locale",
						},
						"title": {
							TargetField: "Title",
						},
					},
				},
				"email": {
					TargetField: "Email",
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"content": {
							TargetField: "Content",
						},
						"locale": {
							TargetField: "Locale",
						},
						"subject": {
							TargetField: "Subject",
						},
					},
				},
				"wechat": {
					TargetField: "WeChat",
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"content": {
							TargetField: "Content",
						},
						"locale": {
							TargetField: "Locale",
						},
					},
				},
				"webhook": {
					TargetField: "Webhook",
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"content": {
							TargetField: "Content",
						},
					},
				},
				"ding_talk": {
					TargetField: "DingTalk",
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"content": {
							TargetField: "Content",
						},
						"locale": {
							TargetField: "Locale",
						},
						"title": {
							TargetField: "Title",
						},
					},
				},
				"need_valid_content": {
					TargetField: "NeedValidContent",
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.POST,
					Path:        []string{call.Action},
					Client:      client.BypassSvcClient.NewTlsClient(),
				}, call.SdkParam)
				if err != nil {
					return nil, err
				}
				logger.Debug(logger.RespFormat, call.Action, resp)
				return resp, nil
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("RESPONSE.AlarmContentTemplateId", *resp)
				d.SetId(id.(string))
				d.Set("alarm_content_template_id", id.(string))
				return nil
			},
		},
	}
	return []ve.Callback{createAlarmContentTemplateCallback}
}

func (v *VolcengineTlsAlarmContentTemplateService) ModifyResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	modifyAlarmContentTemplateCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyAlarmContentTemplate",
			ConvertMode: ve.RequestConvertInConvert,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"alarm_content_template_name": {
					TargetField: "AlarmContentTemplateName",
				},
				"sms": {
					TargetField: "Sms",
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"content": {
							TargetField: "Content",
						},
						"locale": {
							TargetField: "Locale",
						},
					},
				},
				"vms": {
					TargetField: "Vms",
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"content": {
							TargetField: "Content",
						},
						"locale": {
							TargetField: "Locale",
						},
					},
				},
				"lark": {
					TargetField: "Lark",
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"content": {
							TargetField: "Content",
						},
						"locale": {
							TargetField: "Locale",
						},
						"title": {
							TargetField: "Title",
						},
					},
				},
				"email": {
					TargetField: "Email",
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"content": {
							TargetField: "Content",
						},
						"locale": {
							TargetField: "Locale",
						},
						"subject": {
							TargetField: "Subject",
						},
					},
				},
				"wechat": {
					TargetField: "WeChat",
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"content": {
							TargetField: "Content",
						},
						"locale": {
							TargetField: "Locale",
						},
					},
				},
				"webhook": {
					TargetField: "Webhook",
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"content": {
							TargetField: "Content",
						},
					},
				},
				"ding_talk": {
					TargetField: "DingTalk",
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"content": {
							TargetField: "Content",
						},
						"locale": {
							TargetField: "Locale",
						},
						"title": {
							TargetField: "Title",
						},
					},
				},
				"need_valid_content": {
					TargetField: "NeedValidContent",
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.PUT,
					Path:        []string{call.Action},
					Client:      client.BypassSvcClient.NewTlsClient(),
				}, call.SdkParam)
				if err != nil {
					return nil, err
				}
				logger.Debug(logger.RespFormat, call.Action, resp)
				return resp, nil
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["AlarmContentTemplateId"] = d.Id()
				return true, nil
			},
		},
	}
	return []ve.Callback{modifyAlarmContentTemplateCallback}
}

func (v *VolcengineTlsAlarmContentTemplateService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	logger.Debug(logger.RespFormat, resourceData.Id(), "DeleteAlarmContentTemplate")

	removeAlarmContentTemplateCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteAlarmContentTemplate",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"AlarmContentTemplateId": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.DELETE,
					Path:        []string{call.Action},
					Client:      client.BypassSvcClient.NewTlsClient(),
				}, call.SdkParam)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					data, callErr := v.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading tls alarm content template on delete %q, %w", d.Id(), callErr))
						}
					}
					if len(data) == 0 {
						return nil
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
	return []ve.Callback{removeAlarmContentTemplateCallback}
}

func (v *VolcengineTlsAlarmContentTemplateService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		NameField:    "AlarmContentTemplateName",
		IdField:      "AlarmContentTemplateId",
		CollectField: "templates",
		ResponseConverts: map[string]ve.ResponseConvert{
			"Sms": {
				TargetField: "sms",
			},
			"Vms": {
				TargetField: "vms",
			},
			"Lark": {
				TargetField: "lark",
			},
			"Email": {
				TargetField: "email",
			},
			"WeChat": {
				TargetField: "wechat",
			},
			"Webhook": {
				TargetField: "webhook",
			},
			"DingTalk": {
				TargetField: "ding_talk",
			},
		},
	}
}

func (v *VolcengineTlsAlarmContentTemplateService) ReadResourceId(id string) string {
	return id
}

func (v *VolcengineTlsAlarmContentTemplateService) ProjectTrn() *ve.ProjectTrn {
	return &ve.ProjectTrn{
		ServiceName:          "tls",
		ResourceType:         "alarmcontenttemplate",
		ProjectSchemaField:   "iam_project_name",
		ProjectResponseField: "IamProjectName",
	}
}
