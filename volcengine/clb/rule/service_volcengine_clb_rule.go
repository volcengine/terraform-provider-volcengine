package rule

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/clb/clb"
)

type VolcengineRuleService struct {
	Client *ve.SdkClient
}

func NewRuleService(c *ve.SdkClient) *VolcengineRuleService {
	return &VolcengineRuleService{
		Client: c,
	}
}

func (s *VolcengineRuleService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineRuleService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithSimpleQuery(condition, func(m map[string]interface{}) ([]interface{}, error) {
		action := "DescribeRules"
		logger.Debug(logger.ReqFormat, action, condition)
		// 检查 RuleIds 是否存在
		idsMap := make(map[string]bool)
		if ids, ok := condition["RuleIds"]; ok {
			var values []interface{}
			switch _ids := ids.(type) {
			case *schema.Set:
				values = _ids.List() // from datasource
			default:
				values = _ids.([]interface{}) // from resource_read
			}
			for _, value := range values {
				if value == nil {
					continue
				}
				idsMap[strings.Trim(value.(string), " ")] = true
			}
			delete(condition, "RuleIds")
		}

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
		results, err = ve.ObtainSdkValue("Result.Rules", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.Rules is not Slice")
		}

		if len(idsMap) == 0 {
			return data, nil
		}
		// checkIds
		var res []interface{}
		for _, ele := range data {
			if _, ok := idsMap[ele.(map[string]interface{})["RuleId"].(string)]; ok {
				res = append(res, ele)
			}
		}
		return res, err
	})
}

func (s *VolcengineRuleService) ReadResource(resourceData *schema.ResourceData, ruleId string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	listenerId, ok := resourceData.GetOk("listener_id")
	if !ok {
		return nil, fmt.Errorf("non ListenerId")
	}
	req := map[string]interface{}{
		"ListenerId": listenerId.(string),
		"RuleIds":    []interface{}{ruleId},
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
		return data, fmt.Errorf("Rule %s not exist ", ruleId)
	}
	return data, err
}

func (s *VolcengineRuleService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh:    nil,
	}
}

func (VolcengineRuleService) WithResourceResponseHandlers(rule map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return rule, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineRuleService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var clbId string
	var err error

	// 验证action_type和server_group_id的组合
	actionType := resourceData.Get("action_type").(string)
	if actionType == "" {
		actionType = "Forward"
	}

	if actionType == "Forward" {
		serverGroupId := resourceData.Get("server_group_id").(string)
		if serverGroupId == "" {
			return []ve.Callback{{
				Err: fmt.Errorf("server_group_id is required when action_type is Forward"),
			}}
		}

		// 查询 LoadBalancerId
		clbId, err = s.queryLoadBalancerId(serverGroupId)
		if err != nil {
			return []ve.Callback{{
				Err: err,
			}}
		}
	} else {
		// 对于Redirect类型，使用 listener_id 查询 LoadBalancerId，在规则创建期间，需要锁定 clb
		clbId, err = s.queryLoadBalancerIdByListenerId(resourceData.Get("listener_id").(string))
		if err != nil {
			return []ve.Callback{{
				Err: err,
			}}
		}
	}

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateRules",
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"domain": {
					TargetField: "Rules.1.Domain",
				},
				"url": {
					TargetField: "Rules.1.Url",
				},
				"server_group_id": {
					TargetField: "Rules.1.ServerGroupId",
				},
				"description": {
					TargetField: "Rules.1.Description",
				},
				"action_type": {
					TargetField: "Rules.1.ActionType",
				},
				"tags": {
					TargetField: "Rules.1.Tags",
					ConvertType: ve.ConvertListN,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				actionType := d.Get("action_type").(string)
				if actionType == "" {
					actionType = "Forward"
				}

				// 如果是重定向类型，添加重定向配置
				if actionType == "Redirect" {
					if redirectConfig, ok := d.GetOk("redirect_config"); ok {
						configs := redirectConfig.([]interface{})
						if len(configs) > 0 && configs[0] != nil {
							config := configs[0].(map[string]interface{})

							if protocol, ok := config["protocol"].(string); ok && protocol != "" {
								(*call.SdkParam)["Rules.1.RedirectConfig.Protocol"] = protocol
							}
							if host, ok := config["host"].(string); ok && host != "" {
								(*call.SdkParam)["Rules.1.RedirectConfig.Host"] = host
							}
							if path, ok := config["path"].(string); ok && path != "" {
								(*call.SdkParam)["Rules.1.RedirectConfig.Path"] = path
							}
							if port, ok := config["port"].(string); ok && port != "" {
								(*call.SdkParam)["Rules.1.RedirectConfig.Port"] = port
							}
							if statusCode, ok := config["status_code"].(string); ok && statusCode != "" {
								(*call.SdkParam)["Rules.1.RedirectConfig.StatusCode"] = statusCode
							}
						}
					} else {
						return false, fmt.Errorf("redirect_config is required when action_type is Redirect")
					}
				}

				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				ids, _ := ve.ObtainSdkValue("Result.RuleIds", *resp)
				d.SetId(ids.([]interface{})[0].(string))
				return nil
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				clb.NewClbService(s.Client): {
					Target:     []string{"Active", "Inactive"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					ResourceId: clbId,
				},
			},
			LockId: func(d *schema.ResourceData) string {
				return clbId
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineRuleService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback
	var clbId string
	var err error

	actionType := resourceData.Get("action_type").(string)
	if actionType == "" {
		actionType = "Forward"
	}

	if actionType == "Forward" {
		serverGroupId := resourceData.Get("server_group_id").(string)
		if serverGroupId == "" {
			return []ve.Callback{{
				Err: fmt.Errorf("server_group_id is required when action_type is Forward"),
			}}
		}

		// 查询 LoadBalancerId
		clbId, err = s.queryLoadBalancerId(serverGroupId)
		if err != nil {
			return []ve.Callback{{
				Err: err,
			}}
		}
	} else {
		// 对于Redirect类型，使用 listener_id 查询 Load	BalancerId，在规则修改期间，需要锁定 clb
		clbId, err = s.queryLoadBalancerIdByListenerId(resourceData.Get("listener_id").(string))
		if err != nil {
			return []ve.Callback{{
				Err: err,
			}}
		}
	}

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyRules",
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"listener_id": {
					TargetField: "ListenerId",
					ForceGet:    true,
				},
				"server_group_id": {
					TargetField: "Rules.1.ServerGroupId",
				},
				"description": {
					TargetField: "Rules.1.Description",
				},
				"action_type": {
					TargetField: "Rules.1.ActionType",
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["Rules.1.RuleId"] = d.Id()
				actionType := d.Get("action_type").(string)

				// 如果是重定向类型，添加重定向配置
				if actionType == "Redirect" {
					if redirectConfig, ok := d.GetOk("redirect_config"); ok {
						configs := redirectConfig.([]interface{})
						if len(configs) > 0 && configs[0] != nil {
							config := configs[0].(map[string]interface{})

							if protocol, ok := config["protocol"].(string); ok && protocol != "" {
								(*call.SdkParam)["Rules.1.RedirectConfig.Protocol"] = protocol
							}
							if host, ok := config["host"].(string); ok && host != "" {
								(*call.SdkParam)["Rules.1.RedirectConfig.Host"] = host
							}
							if path, ok := config["path"].(string); ok && path != "" {
								(*call.SdkParam)["Rules.1.RedirectConfig.Path"] = path
							}
							if port, ok := config["port"].(string); ok && port != "" {
								(*call.SdkParam)["Rules.1.RedirectConfig.Port"] = port
							}
							if statusCode, ok := config["status_code"].(string); ok && statusCode != "" {
								(*call.SdkParam)["Rules.1.RedirectConfig.StatusCode"] = statusCode
							}
						}
					} else {
						return false, fmt.Errorf("redirect_config is required when action_type is Redirect")
					}
				}

				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				clb.NewClbService(s.Client): {
					Target:     []string{"Active", "Inactive"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					ResourceId: clbId,
				},
			},
			LockId: func(d *schema.ResourceData) string {
				return clbId
			},
		},
	}
	callbacks = append(callbacks, callback)

	// 更新 Tags
	setResourceTagsCallbacks := ve.SetResourceTags(s.Client, "TagResources", "UntagResources", "rule", resourceData, getUniversalInfo)
	callbacks = append(callbacks, setResourceTagsCallbacks...)

	return callbacks
}

func (s *VolcengineRuleService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	var clbId string
	var err error

	actionType := resourceData.Get("action_type").(string)
	if actionType == "" {
		actionType = "Forward"
	}

	// 查询 LoadBalancerId
	if actionType == "Forward" {
		clbId, err = s.queryLoadBalancerId(resourceData.Get("server_group_id").(string))
		if err != nil {
			return []ve.Callback{
				{
					Err: err,
				},
			}
		}
	} else {
		// 对于Redirect类型，使用 listener_id 查询 LoadBalancerId
		clbId, err = s.queryLoadBalancerIdByListenerId(resourceData.Get("listener_id").(string))
		if err != nil {
			return []ve.Callback{
				{
					Err: err,
				},
			}
		}
	}

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteRules",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"RuleIds.1":  resourceData.Id(),
				"ListenerId": resourceData.Get("listener_id"),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on  reading vpc on delete %q, %w", d.Id(), callErr))
						}
					}
					_, callErr = call.ExecuteCall(d, client, call)
					if callErr == nil {
						return nil
					}
					return resource.RetryableError(callErr)
				})
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				clb.NewClbService(s.Client): {
					Target:     []string{"Active", "Inactive"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					ResourceId: clbId,
				},
			},
			LockId: func(d *schema.ResourceData) string {
				return clbId
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineRuleService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "RuleIds",
			},
			"tags": {
				TargetField: "TagFilters",
				ConvertType: ve.ConvertListN,
				NextLevelConvert: map[string]ve.RequestConvert{
					"value": {
						TargetField: "Values.1",
					},
				},
			},
		},
		IdField:      "RuleId",
		CollectField: "rules",
		ResponseConverts: map[string]ve.ResponseConvert{
			"RuleId": {
				TargetField: "id",
				KeepDefault: true,
			},
		},
	}
}

func (s *VolcengineRuleService) ReadResourceId(id string) string {
	return id
}

func (s *VolcengineRuleService) queryLoadBalancerId(serverGroupId string) (string, error) {
	// 查询 LoadBalancerId
	action := "DescribeServerGroupAttributes"
	serverGroupResp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(action), &map[string]interface{}{
		"ServerGroupId": serverGroupId,
	})
	if err != nil {
		return "", err
	}
	clbId, err := ve.ObtainSdkValue("Result.LoadBalancerId", *serverGroupResp)
	if err != nil {
		return "", err
	}
	return clbId.(string), nil
}

func (s *VolcengineRuleService) queryLoadBalancerIdByListenerId(listenerId string) (string, error) {
	// 使用 listener_id 查询 LoadBalancerId
	// 原因：当 action_type 为 Redirect 时，server_group_id 不再是 Required 参数，而 listener_id 是唯一的
	action := "DescribeListenerAttributes"
	listenerResp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(action), &map[string]interface{}{
		"ListenerId": listenerId,
	})
	if err != nil {
		return "", err
	}
	clbId, err := ve.ObtainSdkValue("Result.LoadBalancerId", *listenerResp)
	if err != nil {
		return "", err
	}
	return clbId.(string), nil
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "clb",
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		Action:      actionName,
	}
}
