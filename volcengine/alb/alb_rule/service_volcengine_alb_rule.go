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
	"github.com/volcengine/terraform-provider-volcengine/volcengine/alb/alb"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/alb/alb_listener"
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
	listenerId := resourceData.Get("listener_id").(string)
	listener, _ := alb_listener.NewAlbListenerService(s.Client).ReadResource(resourceData, listenerId)
	loadBalancerId := listener["LoadBalancerId"].(string)
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
				"priority": {
					TargetField: "Rules.1.Priority",
				},
				"sticky_session_enabled": {
					TargetField: "Rules.1.ForwardGroupConfig.StickySessionEnabled",
				},
				"sticky_session_timeout": {
					TargetField: "Rules.1.ForwardGroupConfig.StickySessionTimeout",
				},
				"server_group_tuples": {
					TargetField: "Rules.1.ForwardGroupConfig.ServerGroupTuples",
					ConvertType: ve.ConvertListN,
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
				// 框架支持嵌套转换
				"rule_conditions": {
					TargetField: "Rules.1.RuleConditions",
					ConvertType: ve.ConvertListN,
					NextLevelConvert: map[string]ve.RequestConvert{
						"type": {
							TargetField: "Type",
						},
						"host_config": {
							TargetField: "HostConfig",
							ConvertType: ve.ConvertListUnique,
							NextLevelConvert: map[string]ve.RequestConvert{
								"values": {
									TargetField: "Values",
									ConvertType: ve.ConvertWithN,
								},
							},
						},
						"path_config": {
							TargetField: "PathConfig",
							ConvertType: ve.ConvertListUnique,
							NextLevelConvert: map[string]ve.RequestConvert{
								"values": {
									TargetField: "Values",
									ConvertType: ve.ConvertWithN,
								},
							},
						},
						"header_config": {
							TargetField: "HeaderConfig",
							ConvertType: ve.ConvertListUnique,
							NextLevelConvert: map[string]ve.RequestConvert{
								"values": {
									TargetField: "Values",
									ConvertType: ve.ConvertWithN,
								},
							},
						},
						"method_config": {
							TargetField: "MethodConfig",
							ConvertType: ve.ConvertListUnique,
							NextLevelConvert: map[string]ve.RequestConvert{
								"values": {
									TargetField: "Values",
									ConvertType: ve.ConvertWithN,
								},
							},
						},
						"query_string_config": {
							TargetField: "QueryStringConfig",
							ConvertType: ve.ConvertListUnique,
							NextLevelConvert: map[string]ve.RequestConvert{
								"values": {
									TargetField: "Values",
									ConvertType: ve.ConvertListN,
								},
							},
						},
					},
				},
				"rule_actions": {
					TargetField: "Rules.1.RuleActions",
					ConvertType: ve.ConvertListN,
					NextLevelConvert: map[string]ve.RequestConvert{
						"traffic_limit_config": {
							TargetField: "TrafficLimitConfig",
							ConvertType: ve.ConvertListUnique,
							NextLevelConvert: map[string]ve.RequestConvert{
								"qps": {
									TargetField: "QPS",
								},
							},
						},
						"forward_group_config": {
							TargetField: "ForwardGroupConfig",
							ConvertType: ve.ConvertListUnique,
							NextLevelConvert: map[string]ve.RequestConvert{
								"server_group_sticky_session": {
									TargetField: "ServerGroupStickySession",
									ConvertType: ve.ConvertListUnique,
								},
								"server_group_tuples": {
									TargetField: "ServerGroupTuples",
									ConvertType: ve.ConvertListN,
								},
							},
						},
						"redirect_config": {
							TargetField: "RedirectConfig",
							ConvertType: ve.ConvertListUnique,
						},
						"rewrite_config": {
							TargetField: "RewriteConfig",
							ConvertType: ve.ConvertListUnique,
						},
						"fixed_response_config": {
							TargetField: "FixedResponseConfig",
							ConvertType: ve.ConvertListUnique,
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
			LockId: func(d *schema.ResourceData) string {
				return loadBalancerId
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				ids, _ := ve.ObtainSdkValue("Result.RuleIds", *resp)
				if len(ids.([]interface{})) < 1 {
					return fmt.Errorf("rule id not found")
				}
				ruleId := ids.([]interface{})[0].(string)
				d.SetId(fmt.Sprintf("%v:%v", listenerId, ruleId))
				return nil
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				alb.NewAlbService(s.Client): {
					Target:     []string{"Active", "Inactive"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					ResourceId: loadBalancerId,
				},
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineAlbRuleService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	convert := func(v interface{}) interface{} {
		if v == nil {
			return nil
		}
		if _, ok := v.([]interface{}); ok {
			return v
		}
		return []interface{}{v}
	}
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		converts := map[string]ve.ResponseConvert{
			"HostConfig":               {Convert: convert},
			"PathConfig":               {Convert: convert},
			"HeaderConfig":             {Convert: convert},
			"MethodConfig":             {Convert: convert},
			"QueryStringConfig":        {Convert: convert},
			"TrafficLimitConfig":       {Convert: convert},
			"ForwardGroupConfig":       {Convert: convert},
			"RedirectConfig":           {Convert: convert},
			"RewriteConfig":            {Convert: convert},
			"FixedResponseConfig":      {Convert: convert},
			"ServerGroupStickySession": {Convert: convert},
		}
		return d, converts, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineAlbRuleService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	listenerId := resourceData.Get("listener_id").(string)
	listener, _ := alb_listener.NewAlbListenerService(s.Client).ReadResource(resourceData, listenerId)
	loadBalancerId := listener["LoadBalancerId"].(string)
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
				"priority": {
					TargetField: "Rules.1.Priority",
				},
				"sticky_session_enabled": {
					TargetField: "Rules.1.ForwardGroupConfig.StickySessionEnabled",
				},
				"sticky_session_timeout": {
					TargetField: "Rules.1.ForwardGroupConfig.StickySessionTimeout",
				},
				"server_group_tuples": {
					TargetField: "Rules.1.ForwardGroupConfig.ServerGroupTuples",
					ConvertType: ve.ConvertListN,
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
				"rule_conditions": {
					ForceGet:    true,
					TargetField: "Rules.1.RuleConditions",
					ConvertType: ve.ConvertListN,
					NextLevelConvert: map[string]ve.RequestConvert{
						"type": {
							TargetField: "Type",
							ForceGet:    true,
						},
						"host_config": {
							TargetField: "HostConfig",
							ForceGet:    true,
							ConvertType: ve.ConvertListUnique,
							NextLevelConvert: map[string]ve.RequestConvert{
								"values": {
									TargetField: "Values",
									ConvertType: ve.ConvertWithN,
								},
							},
						},
						"path_config": {
							TargetField: "PathConfig",
							ForceGet:    true,
							ConvertType: ve.ConvertListUnique,
							NextLevelConvert: map[string]ve.RequestConvert{
								"values": {
									TargetField: "Values",
									ConvertType: ve.ConvertWithN,
								},
							},
						},
						"header_config": {
							TargetField: "HeaderConfig",
							ForceGet:    true,
							ConvertType: ve.ConvertListUnique,
							NextLevelConvert: map[string]ve.RequestConvert{
								"values": {
									TargetField: "Values",
									ConvertType: ve.ConvertWithN,
								},
							},
						},
						"method_config": {
							TargetField: "MethodConfig",
							ForceGet:    true,
							ConvertType: ve.ConvertListUnique,
							NextLevelConvert: map[string]ve.RequestConvert{
								"values": {
									TargetField: "Values",
									ConvertType: ve.ConvertWithN,
								},
							},
						},
						"query_string_config": {
							TargetField: "QueryStringConfig",
							ForceGet:    true,
							ConvertType: ve.ConvertListUnique,
							NextLevelConvert: map[string]ve.RequestConvert{
								"values": {
									TargetField: "Values",
									ConvertType: ve.ConvertListN,
								},
							},
						},
					},
				},
				"rule_actions": {
					TargetField: "Rules.1.RuleActions",
					ForceGet:    true,
					ConvertType: ve.ConvertListN,
					NextLevelConvert: map[string]ve.RequestConvert{
						"traffic_limit_config": {
							ForceGet:    true,
							TargetField: "TrafficLimitConfig",
							ConvertType: ve.ConvertListUnique,
							NextLevelConvert: map[string]ve.RequestConvert{
								"qps": {
									TargetField: "QPS",
								},
							},
						},
						"forward_group_config": {
							TargetField: "ForwardGroupConfig",
							ForceGet:    true,
							ConvertType: ve.ConvertListUnique,
							NextLevelConvert: map[string]ve.RequestConvert{
								"server_group_sticky_session": {
									TargetField: "ServerGroupStickySession",
									ConvertType: ve.ConvertListUnique,
								},
								"server_group_tuples": {
									TargetField: "ServerGroupTuples",
									ConvertType: ve.ConvertListN,
								},
							},
						},
						"redirect_config": {
							TargetField: "RedirectConfig",
							ForceGet:    true,
							ConvertType: ve.ConvertListUnique,
						},
						"rewrite_config": {
							TargetField: "RewriteConfig",
							ForceGet:    true,
							ConvertType: ve.ConvertListUnique,
						},
						"fixed_response_config": {
							TargetField: "FixedResponseConfig",
							ForceGet:    true,
							ConvertType: ve.ConvertListUnique,
						},
					},
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
				return loadBalancerId
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				alb.NewAlbService(s.Client): {
					Target:     []string{"Active", "Inactive"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					ResourceId: loadBalancerId,
				},
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineAlbRuleService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	listenerId := resourceData.Get("listener_id").(string)
	listener, _ := alb_listener.NewAlbListenerService(s.Client).ReadResource(resourceData, listenerId)
	loadBalancerId := listener["LoadBalancerId"].(string)
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteRules",
			ConvertMode: ve.RequestConvertIgnore,
			LockId: func(d *schema.ResourceData) string {
				return loadBalancerId
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
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				alb.NewAlbService(s.Client): {
					Target:     []string{"Active", "Inactive"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					ResourceId: loadBalancerId,
				},
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
			"RuleActions.TrafficLimitConfig.QPS": {
				TargetField: "qps",
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
