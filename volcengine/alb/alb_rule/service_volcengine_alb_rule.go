package alb_rule

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

type VolcengineAlbRuleService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewAlbRuleService(c *ve.SdkClient) *VolcengineAlbRuleService {
	return &VolcengineAlbRuleService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineAlbRuleService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineAlbRuleService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	return ve.WithSimpleQuery(condition, func(m map[string]interface{}) ([]interface{}, error) {
		var (
			resp    *map[string]interface{}
			results interface{}
			ok      bool
		)
		action := "DescribeRules"
		logger.Debug(logger.ReqFormat, action, condition)
		if condition == nil {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), nil)
		} else {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &condition)
		}
		if err != nil {
			return nil, err
		}
		logger.Debug(logger.RespFormat, action, *resp)
		results, err = ve.ObtainSdkValue("Result.Rules", *resp)
		if err != nil {
			return []interface{}{}, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.Rules is not Slice")
		} else {
			return data, err
		}
	})
}

func (s *VolcengineAlbRuleService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		temp    map[string]interface{}
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	ids := strings.Split(id, ":")
	req := map[string]interface{}{
		"ListenerId": ids[0],
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		if temp, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("Value is not map ")
		} else if temp["RuleId"].(string) == ids[1] {
			data = temp
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("alb_rule %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineAlbRuleService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineAlbRuleService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateRules",
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"listener_id": {
					TargetField: "ListenerId",
				},
				"domain": {
					TargetField: "Rules.1.Domain",
				},
				"url": {
					TargetField: "Rules.1.Url",
				},
				"rule_action": {
					TargetField: "Rules.1.RuleAction",
				},
				"server_group_id": {
					TargetField: "Rules.1.ServerGroupId",
				},
				"description": {
					TargetField: "Rules.1.Description",
				},
				"traffic_limit_enabled": {
					TargetField: "Rules.1.TrafficLimitEnabled",
				},
				"traffic_limit_qps": {
					TargetField: "Rules.1.TrafficLimitQPS",
				},
				"rewrite_enabled": {
					TargetField: "Rules.1.RewriteEnabled",
				},
				"rewrite_config": {
					TargetField: "Rules.1.RewriteConfig",
					ConvertType: ve.ConvertListUnique,
				},
				"redirect_config": {
					TargetField: "Rules.1.RedirectConfig",
					ConvertType: ve.ConvertListUnique,
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("listener_id").(string)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				ids, _ := ve.ObtainSdkValue("Result.RuleIds", *resp)
				if len(ids.([]interface{})) < 1 {
					return fmt.Errorf("rule id not found")
				}
				ruleId := ids.([]interface{})[0].(string)
				listenerId := d.Get("listener_id").(string)
				d.SetId(fmt.Sprintf("%v:%v", listenerId, ruleId))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineAlbRuleService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineAlbRuleService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyRules",
			ConvertMode: ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"server_group_id": {
					TargetField: "Rules.1.ServerGroupId",
				},
				"description": {
					TargetField: "Rules.1.Description",
				},
				"traffic_limit_enabled": {
					TargetField: "Rules.1.TrafficLimitEnabled",
				},
				"traffic_limit_qps": {
					TargetField: "Rules.1.TrafficLimitQPS",
				},
				"rewrite_enabled": {
					TargetField: "Rules.1.RewriteEnabled",
				},
				"rewrite_config": {
					TargetField: "Rules.1.RewriteConfig",
					ConvertType: ve.ConvertListUnique,
					ForceGet:    true,
				},
				"redirect_config": {
					TargetField: "Rules.1.RedirectConfig",
					ConvertType: ve.ConvertListUnique,
					ForceGet:    true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				ids := strings.Split(d.Id(), ":")
				(*call.SdkParam)["ListenerId"] = ids[0]
				(*call.SdkParam)["Rules.1.RuleId"] = ids[1]
				ruleAction, ok := d.GetOk("rule_action")
				/*
					1. ruleAction = Redirect，则redirect_config必传
					2. 若ruleAction没写，则serverGroupId必传
				*/
				if ok {
					(*call.SdkParam)["Rules.1.RuleAction"] = ruleAction
					_, ok = d.GetOk("redirect_config")
					if ruleAction.(string) == "Redirect" && !ok {
						return false, fmt.Errorf("redirect_config is required when rule_action is Redirect")
					}
				}
				return true, nil
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("listener_id").(string)
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

func (s *VolcengineAlbRuleService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteRules",
			ConvertMode: ve.RequestConvertIgnore,
			LockId: func(d *schema.ResourceData) string {
				return d.Get("listener_id").(string)
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				ids := strings.Split(d.Id(), ":")
				(*call.SdkParam)["ListenerId"] = ids[0]
				(*call.SdkParam)["RuleIds.1"] = ids[1]
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 5*time.Minute)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineAlbRuleService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		IdField:      "RuleId",
		CollectField: "rules",
		ResponseConverts: map[string]ve.ResponseConvert{
			"RuleId": {
				TargetField: "id",
				KeepDefault: true,
			},
			"TrafficLimitQPS": {
				TargetField: "traffic_limit_qps",
			},
		},
	}
}

func (s *VolcengineAlbRuleService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "alb",
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		Action:      actionName,
	}
}
