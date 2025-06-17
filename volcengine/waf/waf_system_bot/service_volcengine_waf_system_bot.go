package waf_system_bot

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineWafSystemBotService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewWafSystemBotService(c *ve.SdkClient) *VolcengineWafSystemBotService {
	return &VolcengineWafSystemBotService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineWafSystemBotService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineWafSystemBotService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListSystemBotConfig"

		bytes, _ := json.Marshal(condition)
		logger.Debug(logger.ReqFormat, action, string(bytes))
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
		respBytes, _ := json.Marshal(resp)
		logger.Debug(logger.RespFormat, action, condition, string(respBytes))
		results, err = ve.ObtainSdkValue("Result.Data", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.Data is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineWafSystemBotService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
		result  map[string]interface{}
	)
	req := map[string]interface{}{
		"Host": resourceData.Get("host"),
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		if data, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("Value is not map ")
		}
		if data["BotType"] == resourceData.Get("bot_type") {
			result = data
			break
		}
	}
	if len(result) == 0 {
		return result, fmt.Errorf("waf_system_bot %s not exist ", id)
	}

	return result, err
}

func (s *VolcengineWafSystemBotService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineWafSystemBotService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateSystemBotConfig",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert:     map[string]ve.RequestConvert{},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				d.SetId(fmt.Sprintf("%s:%s", d.Get("bot_type"), d.Get("host")))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineWafSystemBotService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineWafSystemBotService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateSystemBotConfig",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"project_name": {
					TargetField: "ProjectName",
					ForceGet:    true,
				},
				"action": {
					TargetField: "Action",
					ForceGet:    true,
				},
				"enable": {
					TargetField: "Enable",
					ForceGet:    true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				parts := strings.Split(d.Id(), ":")
				if len(parts) != 2 {
					return false, fmt.Errorf("format of waf system bot resource id is invalid,%s", d.Id())
				}
				botType := parts[0]
				host := parts[1]
				(*call.SdkParam)["BotType"] = botType
				(*call.SdkParam)["Host"] = host
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineWafSystemBotService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	//callback := ve.Callback{
	//	Call: ve.SdkCall{
	//		Action:      "UpdateSystemBotConfig",
	//		ConvertMode: ve.RequestConvertIgnore,
	//		ContentType: ve.ContentTypeJson,
	//		BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
	//			parts := strings.Split(d.Id(), ":")
	//			if len(parts) != 2 {
	//				return false, fmt.Errorf("format of waf system bot resource id is invalid,%s", d.Id())
	//			}
	//			botType := parts[0]
	//			host := parts[1]
	//			(*call.SdkParam)["BotType"] = botType
	//			(*call.SdkParam)["Host"] = host
	//			(*call.SdkParam)["Action"] = d.Get("action")
	//			(*call.SdkParam)["Enable"] = 0
	//			return true, nil
	//		},
	//		ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
	//			logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
	//			return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
	//		},
	//		AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
	//			return s.checkResourceUtilRemoved(d, 5*time.Minute)
	//		},
	//		CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
	//			//出现错误后重试
	//			return resource.Retry(5*time.Minute, func() *resource.RetryError {
	//				_, callErr := s.ReadResource(d, "")
	//				if callErr != nil {
	//					if ve.ResourceNotFoundError(callErr) {
	//						return nil
	//					} else {
	//						return resource.NonRetryableError(fmt.Errorf("error on  reading waf custom page on delete %q, %w", d.Id(), callErr))
	//					}
	//				}
	//				_, callErr = call.ExecuteCall(d, client, call)
	//				if callErr == nil {
	//					return nil
	//				}
	//				return resource.RetryableError(callErr)
	//			})
	//		},
	//	},
	//}
	logger.Debug(logger.ReqFormat, "RemoveResource", "Remove only from tf management")
	return []ve.Callback{}
}

func (s *VolcengineWafSystemBotService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		NameField:        "BotType",
		IdField:          "RuleTag",
		CollectField:     "data",
		ContentType:      ve.ContentTypeJson,
		ResponseConverts: map[string]ve.ResponseConvert{},
	}
}

func (s *VolcengineWafSystemBotService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "waf",
		Version:     "2023-12-25",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
		RegionType:  ve.Global,
	}
}

func (s *VolcengineWafSystemBotService) checkResourceUtilRemoved(d *schema.ResourceData, timeout time.Duration) error {
	return resource.Retry(timeout, func() *resource.RetryError {
		systemBotConfig, _ := s.ReadResource(d, d.Id())
		logger.Debug(logger.RespFormat, "systemBotConfig", systemBotConfig)

		// 能查询成功代表还在删除中，重试
		systemBotConfigInt, ok := systemBotConfig["Enable"].(float64)
		if !ok {
			return resource.NonRetryableError(fmt.Errorf("enable is not float64"))
		}
		if int(systemBotConfigInt) == 1 {
			return resource.RetryableError(fmt.Errorf("resource still in removing status "))
		} else {
			if int(systemBotConfigInt) == 0 {
				return nil
			} else {
				return resource.NonRetryableError(fmt.Errorf("system bot status is not disable "))
			}
		}
	})
}
