package alarm_webhook_integration

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineTlsAlarmWebhookIntegrationService struct {
	Client *ve.SdkClient
}

func NewVolcengineTlsAlarmWebhookIntegrationService(c *ve.SdkClient) *VolcengineTlsAlarmWebhookIntegrationService {
	return &VolcengineTlsAlarmWebhookIntegrationService{
		Client: c,
	}
}

func convertWebhookHeaders(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	list, ok := v.([]interface{})
	if !ok {
		return nil
	}
	headers := make(map[string]interface{})
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		if !ok {
			continue
		}
		key, kOk := m["key"].(string)
		val, vOk := m["value"].(string)
		if kOk && vOk {
			headers[key] = val
		}
	}
	return headers
}

func (v *VolcengineTlsAlarmWebhookIntegrationService) GetClient() *ve.SdkClient {
	return v.Client
}

func (v *VolcengineTlsAlarmWebhookIntegrationService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 20, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeAlarmWebhookIntegrations"
		logger.Debug(logger.ReqFormat, action, condition)
		resp, err = v.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
			ContentType: ve.ApplicationJSON,
			HttpMethod:  ve.GET,
			Path:        []string{action},
			Client:      v.Client.BypassSvcClient.NewTlsClient(),
		}, &m)
		if err != nil {
			return nil, err
		}
		logger.Debug(logger.RespFormat, action, resp, err)
		results, err = ve.ObtainSdkValue("RESPONSE.WebhookIntegrations", *resp)
		if err != nil {
			// If key is missing, treat as empty list
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			// 如果 results 不是 []interface{}，说明返回的结构不符合预期，可能是空列表被解析成了其他类型
			// 但在这里，我们假设它应该是一个列表。如果它是 nil，我们在上面已经初始化为空列表了。
			return nil, fmt.Errorf("WebhookIntegrations is not Slice, got %T", results)
		}
		return data, err
	})
}

func (v *VolcengineTlsAlarmWebhookIntegrationService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = v.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"WebhookID": id,
	}
	results, err = v.ReadResources(req)
	if err != nil {
		return data, err
	}
	// 如果结果为空列表，直接返回 nil
	if len(results) == 0 {
		return nil, nil
	}
	for _, r := range results {
		if data, ok = r.(map[string]interface{}); !ok {
			return data, fmt.Errorf("value is not map")
		}
		if val, ok := data["WebhookID"]; ok && val == id {
			return data, nil
		}
	}
	return nil, nil
}

func (v *VolcengineTlsAlarmWebhookIntegrationService) RefreshResourceState(data *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (v *VolcengineTlsAlarmWebhookIntegrationService) WithResourceResponseHandlers(m map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return m, map[string]ve.ResponseConvert{
			"WebhookID": {
				TargetField: "id",
			},
			"WebhookName": {
				TargetField: "webhook_name",
			},
			"WebhookUrl": {
				TargetField: "webhook_url",
			},
			"WebhookType": {
				TargetField: "webhook_type",
			},
			"WebhookMethod": {
				TargetField: "webhook_method",
			},
			"WebhookSecret": {
				TargetField: "webhook_secret",
			},
			"WebhookHeaders": {
				TargetField: "webhook_headers",
			},
			"CreateTime": {
				TargetField: "create_time",
			},
			"ModifyTime": {
				TargetField: "modify_time",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (v *VolcengineTlsAlarmWebhookIntegrationService) CreateResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	createAlarmWebhookIntegrationCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateAlarmWebhookIntegration",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"webhook_name": {
					TargetField: "WebhookName",
				},
				"webhook_url": {
					TargetField: "WebhookUrl",
				},
				"webhook_type": {
					TargetField: "WebhookType",
				},
				"webhook_method": {
					TargetField: "WebhookMethod",
				},
				"webhook_secret": {
					TargetField: "WebhookSecret",
				},
				"webhook_headers": {
					TargetField: "WebhookHeaders",
					ConvertType: ve.ConvertJsonObjectArray,
					NextLevelConvert: map[string]ve.RequestConvert{
						"key": {
							TargetField: "key",
						},
						"value": {
							TargetField: "value",
						},
					},
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
				id, _ := ve.ObtainSdkValue("RESPONSE.AlarmWebhookIntegrationId", *resp)
				d.SetId(id.(string))
				return nil
			},
		},
	}
	return []ve.Callback{createAlarmWebhookIntegrationCallback}
}

func (v *VolcengineTlsAlarmWebhookIntegrationService) ModifyResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	modifyAlarmWebhookIntegrationCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyAlarmWebhookIntegration",
			ConvertMode: ve.RequestConvertInConvert,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"webhook_name": {
					TargetField: "WebhookName",
				},
				"webhook_url": {
					TargetField: "WebhookUrl",
				},
				"webhook_type": {
					TargetField: "WebhookType",
				},
				"webhook_method": {
					TargetField: "WebhookMethod",
				},
				"webhook_secret": {
					TargetField: "WebhookSecret",
				},
				"webhook_headers": {
					TargetField: "WebhookHeaders",
					ConvertType: ve.ConvertJsonObjectArray,
					NextLevelConvert: map[string]ve.RequestConvert{
						"key": {
							TargetField: "key",
						},
						"value": {
							TargetField: "value",
						},
					},
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
				(*call.SdkParam)["WebhookID"] = d.Id()
				return true, nil
			},
		},
	}
	return []ve.Callback{modifyAlarmWebhookIntegrationCallback}
}

func (v *VolcengineTlsAlarmWebhookIntegrationService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	removeAlarmWebhookIntegrationCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteAlarmWebhookIntegration",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["WebhookID"] = d.Id()
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.DELETE,
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
				return ve.CheckResourceUtilRemoved(d, v.ReadResource, 3*time.Minute)
			},
		},
	}
	return []ve.Callback{removeAlarmWebhookIntegrationCallback}
}

func (v *VolcengineTlsAlarmWebhookIntegrationService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		NameField:    "WebhookName",
		IdField:      "webhook_id",
		CollectField: "integrations",
		RequestConverts: map[string]ve.RequestConvert{
			"webhook_name": {
				TargetField: "WebhookName",
			},
			"webhook_type": {
				TargetField: "WebhookType",
			},
			"webhook_id": {
				TargetField: "WebhookID",
			},
		},
		ResponseConverts: map[string]ve.ResponseConvert{
			"WebhookID": {
				TargetField: "webhook_id",
			},
			"WebhookName": {
				TargetField: "webhook_name",
			},
			"WebhookUrl": {
				TargetField: "webhook_url",
			},
			"WebhookType": {
				TargetField: "webhook_type",
			},
			"WebhookMethod": {
				TargetField: "webhook_method",
			},
			"WebhookSecret": {
				TargetField: "webhook_secret",
			},
			"WebhookHeaders": {
				TargetField: "webhook_headers",
			},
			"CreateTime": {
				TargetField: "create_time",
			},
			"ModifyTime": {
				TargetField: "modify_time",
			},
		},
	}
}

func (v *VolcengineTlsAlarmWebhookIntegrationService) ReadResourceId(id string) string {
	return id
}

func (v *VolcengineTlsAlarmWebhookIntegrationService) ProjectTrn() *ve.ProjectTrn {
	return &ve.ProjectTrn{
		ServiceName:          "tls",
		ResourceType:         "alarmwebhookintegration",
		ProjectSchemaField:   "iam_project_name",
		ProjectResponseField: "IamProjectName",
	}
}
