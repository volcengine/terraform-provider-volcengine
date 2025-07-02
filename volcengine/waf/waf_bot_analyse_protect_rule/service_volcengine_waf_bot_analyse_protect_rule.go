package waf_bot_analyse_protect_rule

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

type VolcengineWafBotAnalyseProtectRuleService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewWafBotAnalyseProtectRuleService(c *ve.SdkClient) *VolcengineWafBotAnalyseProtectRuleService {
	return &VolcengineWafBotAnalyseProtectRuleService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineWafBotAnalyseProtectRuleService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineWafBotAnalyseProtectRuleService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "Page", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListBotAnalyseProtectRule"

		bytes, _ := json.Marshal(condition)
		logger.Debug(logger.ReqFormat, action, string(bytes))
		condition["Region"] = s.Client.Region
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

func (s *VolcengineWafBotAnalyseProtectRuleService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
		ruleTag string
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}

	parts := strings.Split(id, ":")
	if len(parts) != 3 {
		return data, fmt.Errorf("format of waf bot analyse protect rule resource id is invalid,%s", id)
	}
	ruleId := parts[0]
	botSpace := parts[1]
	host := parts[2]

	ruleIdInt, err := strconv.Atoi(ruleId)
	tag := fmt.Sprintf("%012d", ruleIdInt)
	if botSpace == "BotFrequency" {
		ruleTag = "R" + tag
	}

	if botSpace == "BotRepeat" {
		ruleTag = "S" + tag
	}

	req := map[string]interface{}{
		"Host":     host,
		"BotSpace": botSpace,
		"RuleTag":  ruleTag,
		"Region":   s.Client.Region,
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
		return data, fmt.Errorf("waf_bot_analyse_protect_rule %s not exist ", id)
	}

	if ruleGroups, ruleGroupsExist := data["RuleGroup"]; ruleGroupsExist {
		for _, ruleGroup := range ruleGroups.([]interface{}) {
			if rules, rulesExist := ruleGroup.(map[string]interface{})["Rules"]; rulesExist {
				for _, rule := range rules.([]interface{}) {
					if accurateGroup, accurateGroupExist := rule.(map[string]interface{})["AccurateGroup"]; accurateGroupExist {
						rule.(map[string]interface{})["AccurateGroup"] = []interface{}{
							accurateGroup,
						}
					}
				}
			}
		}
	}
	if ruleGroups, ruleGroupsExist := data["RuleGroup"]; ruleGroupsExist {
		for _, ruleGroup := range ruleGroups.([]interface{}) {
			if group, groupExist := ruleGroup.(map[string]interface{})["Group"]; groupExist {
				ruleGroup.(map[string]interface{})["Group"] = []interface{}{
					group,
				}
			}
		}
	}

	if ruleGroups, ruleGroupsExist := data["RuleGroup"]; ruleGroupsExist {
		for _, ruleGroup := range ruleGroups.([]interface{}) {
			if rules, rulesExist := ruleGroup.(map[string]interface{})["Rules"]; rulesExist {
				for _, rule := range rules.([]interface{}) {
					if actionAfterVerification, actionAfterVerificationExist := rule.(map[string]interface{})["ActionAfterVerification"]; actionAfterVerificationExist {
						data["ActionAfterVerification"] = actionAfterVerification
					}
					if actionType, actionTypeExist := rule.(map[string]interface{})["ActionType"]; actionTypeExist {
						data["ActionType"] = actionType
					}
					if effectTime, effectTimeExist := rule.(map[string]interface{})["EffectTime"]; effectTimeExist {
						data["EffectTime"] = effectTime
					}
					if enable, enableExist := rule.(map[string]interface{})["Enable"]; enableExist {
						data["Enable"] = enable
					}
					if exemptionTime, exemptionTimeExist := rule.(map[string]interface{})["ExemptionTime"]; exemptionTimeExist {
						data["ExemptionTime"] = exemptionTime
					}
					if field, fieldExist := rule.(map[string]interface{})["Field"]; fieldExist {
						data["Field"] = field
					}
					if name, nameExist := rule.(map[string]interface{})["Name"]; nameExist {
						data["Name"] = name
					}
					if pathThreshold, pathThresholdExist := rule.(map[string]interface{})["PathThreshold"]; pathThresholdExist {
						data["PathThreshold"] = pathThreshold
					}
					if rulePriority, rulePriorityExist := rule.(map[string]interface{})["RulePriority"]; rulePriorityExist {
						data["RulePriority"] = rulePriority
					}
					if singleProportion, singleProportionExist := rule.(map[string]interface{})["SingleProportion"]; singleProportionExist {
						data["SingleProportion"] = singleProportion
					}
					if singleThreshold, singleThresholdExist := rule.(map[string]interface{})["SingleThreshold"]; singleThresholdExist {
						data["SingleThreshold"] = singleThreshold
					}
					if statisticalDuration, statisticalDurationExist := rule.(map[string]interface{})["StatisticalDuration"]; statisticalDurationExist {
						data["StatisticalDuration"] = statisticalDuration
					}
					if statisticalType, statisticalTypeExist := rule.(map[string]interface{})["StatisticalType"]; statisticalTypeExist {
						data["StatisticalType"] = statisticalType
					}
					if accurateGroup, accurateGroupExist := rule.(map[string]interface{})["AccurateGroup"]; accurateGroupExist {
						data["AccurateGroup"] = accurateGroup
					}
				}
			}
		}
	}

	return data, err
}

func (s *VolcengineWafBotAnalyseProtectRuleService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineWafBotAnalyseProtectRuleService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateBotAnalyseProtectRule",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"accurate_group": {
					ConvertType: ve.ConvertJsonObject,
					TargetField: "AccurateGroup",
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

				statisticalType, ok := d.Get("statistical_type").(int)
				if !ok {
					return errors.New("statistical_type is not int")
				}

				botSpace := transStatisticalTypeToBotSpace(statisticalType)
				d.SetId(fmt.Sprintf("%s:%s:%s", strconv.Itoa(int(id.(float64))), botSpace, host))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineWafBotAnalyseProtectRuleService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineWafBotAnalyseProtectRuleService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateBotAnalyseProtectRule",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"statistical_duration": {
					TargetField: "StatisticalDuration",
					ForceGet:    true,
				},
				"statistical_type": {
					TargetField: "StatisticalType",
					ForceGet:    true,
				},
				"field": {
					TargetField: "Field",
					ForceGet:    true,
				},
				"accurate_group": {
					ConvertType: ve.ConvertJsonObject,
					ForceGet:    true,
					TargetField: "AccurateGroup",
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
				"action_after_verification": {
					TargetField: "ActionAfterVerification",
					ForceGet:    true,
				},
				"action_type": {
					TargetField: "ActionType",
					ForceGet:    true,
				},
				"effect_time": {
					TargetField: "EffectTime",
					ForceGet:    true,
				},
				"exemption_time": {
					TargetField: "ExemptionTime",
					ForceGet:    true,
				},
				"enable": {
					TargetField: "Enable",
					ForceGet:    true,
				},
				"name": {
					TargetField: "Name",
					ForceGet:    true,
				},
				"path": {
					TargetField: "Path",
					ForceGet:    true,
				},
				"path_threshold": {
					TargetField: "PathThreshold",
					ForceGet:    true,
				},
				"rule_priority": {
					TargetField: "RulePriority",
					ForceGet:    true,
				},
				"single_proportion": {
					TargetField: "SingleProportion",
					ForceGet:    true,
				},
				"single_threshold": {
					TargetField: "SingleThreshold",
					ForceGet:    true,
				},
				"project_name": {
					TargetField: "ProjectName",
					ForceGet:    true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				parts := strings.Split(d.Id(), ":")
				if len(parts) != 3 {
					return false, fmt.Errorf("format of waf bot analyse protect rule resource id is invalid,%s", d.Id())
				}
				id := parts[0]
				host := parts[2]
				ruleId, err := strconv.Atoi(id)
				if err != nil {
					return false, fmt.Errorf(" ruleId cannot convert to int ")
				}
				(*call.SdkParam)["Id"] = ruleId
				(*call.SdkParam)["Host"] = host
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				parts := strings.Split(d.Id(), ":")
				if len(parts) != 3 {
					return fmt.Errorf("format of waf bot analyse protect rule resource id is invalid,%s", d.Id())
				}
				id := parts[0]

				host, ok := d.Get("host").(string)
				if !ok {
					return errors.New("host is not string")
				}

				statisticalType, ok := d.Get("statistical_type").(int)
				if !ok {
					return errors.New("statistical_type is not int")
				}

				botSpace := transStatisticalTypeToBotSpace(statisticalType)
				d.SetId(fmt.Sprintf("%s:%s:%s", id, botSpace, host))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineWafBotAnalyseProtectRuleService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteBotAnalyseProtectRule",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				parts := strings.Split(d.Id(), ":")
				if len(parts) != 3 {
					return false, fmt.Errorf("format of waf bot analyse protect rule resource id is invalid,%s", d.Id())
				}
				id := parts[0]
				ruleId, err := strconv.Atoi(id)
				if err != nil {
					return false, fmt.Errorf(" ruleId cannot convert to int ")
				}
				(*call.SdkParam)["Id"] = ruleId
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
							return resource.NonRetryableError(fmt.Errorf("error on  reading waf acl rule on delete %q, %w", d.Id(), callErr))
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

func (s *VolcengineWafBotAnalyseProtectRuleService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		CollectField: "data",
	}
}

func (s *VolcengineWafBotAnalyseProtectRuleService) ReadResourceId(id string) string {
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

func transStatisticalTypeToBotSpace(statisticalType int) string {
	switch statisticalType {
	case 1:
		return "BotFrequency"
	case 2:
		return "BotRepeat"
	case 3:
		return "BotRepeat"
	default:
		return "bot space not supported"
	}
}
