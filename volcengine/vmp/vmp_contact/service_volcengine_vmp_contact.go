package vmp_contact

import (
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineVmpContactService struct {
	Client *ve.SdkClient
}

func NewService(c *ve.SdkClient) *VolcengineVmpContactService {
	return &VolcengineVmpContactService{
		Client: c,
	}
}

func (s *VolcengineVmpContactService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineVmpContactService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 50, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListContacts"
		logger.Debug(logger.ReqFormat, action, condition)
		if condition == nil {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), nil)
			if err != nil {
				return data, err
			}
		} else {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &condition)
			if err != nil {
				return data, err
			}
		}
		logger.Debug(logger.RespFormat, action, resp)
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

		return data, err
	})
}

func (s *VolcengineVmpContactService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"Filter": map[string]interface{}{
			"Ids": []interface{}{id},
		},
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		if data, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("Value is not map ")
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("Contact %s not exist ", id)
	}

	if _, exist := data["ContactGroupIds"]; !exist {
		data["ContactGroupIds"] = []interface{}{} // 添加空list
	}

	return data, err
}

func (s *VolcengineVmpContactService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (VolcengineVmpContactService) WithResourceResponseHandlers(data map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		for _, key := range []string{"Webhook", "LarkBotWebhook", "DingTalkBotWebhook", "WeComBotWebhook"} {
			if v, exist := data[key]; exist {
				if address, exist := v.(map[string]interface{})["Address"]; exist {
					if address.(string) == "" {
						delete(data, key)
					}
				}
			}
		}
		phoneNumber, ok := data["PhoneNumber"]
		if ok {
			number := phoneNumber.(map[string]interface{})["Number"]
			if len(number.(string)) == 0 {
				delete(data, "PhoneNumber")
			}
		}
		return data, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}

}

func (s *VolcengineVmpContactService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateContact",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"webhook": {
					ConvertType: ve.ConvertJsonObject,
				},
				"lark_bot_webhook": {
					ConvertType: ve.ConvertJsonObject,
				},
				"ding_talk_bot_webhook": {
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"at_mobiles": {
							ConvertType: ve.ConvertJsonArray,
						},
						"at_user_ids": {
							ConvertType: ve.ConvertJsonArray,
						},
					},
				},
				"we_com_bot_webhook": {
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"at_user_ids": {
							ConvertType: ve.ConvertJsonArray,
						},
					},
				},
				"phone_number": {
					ConvertType: ve.ConvertJsonObject,
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
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

func (s *VolcengineVmpContactService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateContact",
			ConvertMode: ve.RequestConvertInConvert,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"name": {
					TargetField: "Name",
				},
				"email": {
					TargetField: "Email",
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["Id"] = d.Id()

				// 必须全部传递
				(*call.SdkParam)["Webhook"] = map[string]interface{}{
					"Address": d.Get("webhook.0.address"),
					"Token":   d.Get("webhook.0.token"),
				}
				(*call.SdkParam)["LarkBotWebhook"] = map[string]interface{}{
					"Address":   d.Get("lark_bot_webhook.0.address"),
					"SecretKey": d.Get("lark_bot_webhook.0.secret_key"),
				}
				(*call.SdkParam)["DingTalkBotWebhook"] = map[string]interface{}{
					"Address":   d.Get("ding_talk_bot_webhook.0.address"),
					"SecretKey": d.Get("ding_talk_bot_webhook.0.secret_key"),
					"AtMobiles": d.Get("ding_talk_bot_webhook.0.at_mobiles"),
					"AtUserIds": d.Get("ding_talk_bot_webhook.0.at_user_ids"),
				}
				(*call.SdkParam)["WeComBotWebhook"] = map[string]interface{}{
					"Address":   d.Get("we_com_bot_webhook.0.address"),
					"AtUserIds": d.Get("we_com_bot_webhook.0.at_user_ids"),
				}
				(*call.SdkParam)["PhoneNumber"] = map[string]interface{}{
					"CountryCode": d.Get("phone_number.0.country_code"),
					"Number":      d.Get("phone_number.0.number"),
				}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
		},
	}
	callbacks = append(callbacks, callback)
	return callbacks
}

func (s *VolcengineVmpContactService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteContacts",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"Ids": []interface{}{resourceData.Id()},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 5*time.Minute)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(5*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading vmp contact on delete %q, %w", d.Id(), callErr))
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

func (s *VolcengineVmpContactService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "Filter.Ids",
				ConvertType: ve.ConvertJsonArray,
			},
			"name": {
				TargetField: "Filter.Name",
			},
			"email": {
				TargetField: "Filter.Email",
			},
		},
		ContentType:  ve.ContentTypeJson,
		NameField:    "Name",
		IdField:      "Id",
		CollectField: "contacts",
	}
}

func (s *VolcengineVmpContactService) ReadResourceId(id string) string {
	return id
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
