package waf_custom_bot

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineWafCustomBotService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewWafCustomBotService(c *ve.SdkClient) *VolcengineWafCustomBotService {
	return &VolcengineWafCustomBotService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineWafCustomBotService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineWafCustomBotService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "Page", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListCustomBotConfig"

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

func (s *VolcengineWafCustomBotService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	parts := strings.Split(id, ":")
	if len(parts) != 2 {
		return data, fmt.Errorf("format of waf custom page resource id is invalid,%s", id)
	}
	customBotId := parts[0]
	host := parts[1]

	customPageIdInt, err := strconv.Atoi(customBotId)
	tag := fmt.Sprintf("%012d", customPageIdInt)
	ruleTag := "K" + tag

	req := map[string]interface{}{
		"RuleTag": ruleTag,
		"Host":    host,
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
		return data, fmt.Errorf("waf_custom_bot %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineWafCustomBotService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineWafCustomBotService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateCustomBotConfig",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"accurate": {
					ConvertType: ve.ConvertJsonObject,
					TargetField: "Accurate",
					NextLevelConvert: map[string]ve.RequestConvert{
						"accurate_rules": {
							ConvertType: ve.ConvertJsonObjectArray,
							TargetField: "AccurateRules",
							NextLevelConvert: map[string]ve.RequestConvert{
								"http_obj": {
									TargetField: "HttpObj",
								},
								"obj_type": {
									TargetField: "ObjType",
								},
								"opretar": {
									TargetField: "Opretar",
								},
								"property": {
									TargetField: "Property",
								},
								"value_string": {
									TargetField: "ValueString",
								},
							},
						},
						"logic": {
							TargetField: "Logic",
						},
					},
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
				host, ok := d.Get("host").(string)
				if !ok {
					return errors.New("host is not string")
				}
				d.SetId(fmt.Sprintf("%s:%s", strconv.Itoa(int(id.(float64))), host))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineWafCustomBotService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineWafCustomBotService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateCustomBotConfig",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"bot_type": {
					TargetField: "BotType",
					ForceGet:    true,
				},
				"description": {
					TargetField: "Description",
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
				"project_name": {
					TargetField: "ProjectName",
					ForceGet:    true,
				},
				"accurate": {
					ConvertType: ve.ConvertJsonObject,
					ForceGet:    true,
					TargetField: "Accurate",
					NextLevelConvert: map[string]ve.RequestConvert{
						"accurate_rules": {
							ConvertType: ve.ConvertJsonObjectArray,
							ForceGet:    true,
							TargetField: "AccurateRules",
							NextLevelConvert: map[string]ve.RequestConvert{
								"http_obj": {
									TargetField: "HttpObj",
									ForceGet:    true,
								},
								"obj_type": {
									TargetField: "ObjType",
									ForceGet:    true,
								},
								"opretar": {
									TargetField: "Opretar",
									ForceGet:    true,
								},
								"property": {
									TargetField: "Property",
									ForceGet:    true,
								},
								"value_string": {
									TargetField: "ValueString",
									ForceGet:    true,
								},
							},
						},
						"logic": {
							TargetField: "Logic",
							ForceGet:    true,
						},
					},
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				parts := strings.Split(d.Id(), ":")
				if len(parts) != 2 {
					return false, fmt.Errorf("format of waf custom bot resource id is invalid,%s", d.Id())
				}
				id := parts[0]
				host := parts[1]
				customBotId, err := strconv.Atoi(id)
				if err != nil {
					return false, fmt.Errorf(" custom bot id cannot convert to int ")
				}
				(*call.SdkParam)["Host"] = host
				(*call.SdkParam)["Id"] = customBotId
				logic, ok := d.Get("accurate.0.logic").(int)
				if !ok {
					return false, fmt.Errorf("accurate.0.logic cannot convert to int ")
				}

				if logic == 0 {
					delete(*call.SdkParam, "Accurate.Logic")
				}
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

func (s *VolcengineWafCustomBotService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteCustomBotConfig",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				parts := strings.Split(d.Id(), ":")
				if len(parts) != 2 {
					return false, fmt.Errorf("format of waf custom bot resource id is invalid,%s", d.Id())
				}
				id := parts[0]
				host := parts[1]
				customBotId, err := strconv.Atoi(id)
				if err != nil {
					return false, fmt.Errorf(" custom bot id cannot convert to int ")
				}
				(*call.SdkParam)["BotID"] = customBotId
				(*call.SdkParam)["Host"] = host
				return true, nil
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
							return resource.NonRetryableError(fmt.Errorf("error on  reading waf custom page on delete %q, %w", d.Id(), callErr))
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

func (s *VolcengineWafCustomBotService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		NameField:        "BotType",
		IdField:          "RuleTag",
		CollectField:     "data",
		ResponseConverts: map[string]ve.ResponseConvert{},
	}
}

func (s *VolcengineWafCustomBotService) ReadResourceId(id string) string {
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
