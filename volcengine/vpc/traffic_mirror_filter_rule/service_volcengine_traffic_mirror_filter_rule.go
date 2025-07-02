package traffic_mirror_filter_rule

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

type VolcengineTrafficMirrorFilterRuleService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewTrafficMirrorFilterRuleService(c *ve.SdkClient) *VolcengineTrafficMirrorFilterRuleService {
	return &VolcengineTrafficMirrorFilterRuleService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineTrafficMirrorFilterRuleService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineTrafficMirrorFilterRuleService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		filters []interface{}
		ok      bool
	)
	return ve.WithNextTokenQuery(m, "MaxResults", "NextToken", 100, nil, func(condition map[string]interface{}) (data []interface{}, next string, err error) {
		action := "DescribeTrafficMirrorFilters"

		bytes, _ := json.Marshal(condition)
		logger.Debug(logger.ReqFormat, action, string(bytes))
		if condition == nil {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), nil)
			if err != nil {
				return data, next, err
			}
		} else {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &condition)
			if err != nil {
				return data, next, err
			}
		}
		respBytes, _ := json.Marshal(resp)
		logger.Debug(logger.RespFormat, action, condition, string(respBytes))
		results, err = ve.ObtainSdkValue("Result.TrafficMirrorFilters", *resp)
		if err != nil {
			return data, next, err
		}
		nextToken, err := ve.ObtainSdkValue("Result.NextToken", *resp)
		if err != nil {
			return data, next, err
		}
		next = nextToken.(string)
		if results == nil {
			results = []interface{}{}
		}
		if filters, ok = results.([]interface{}); !ok {
			return data, next, errors.New("Result.TrafficMirrorFilters is not Slice")
		}

		for _, filter := range filters {
			filterMap, ok := filter.(map[string]interface{})
			if !ok {
				return data, next, errors.New("TrafficMirrorFilter is not map")
			}
			if v, ok := filterMap["IngressFilterRules"]; ok {
				if ingressRules, ok := v.([]interface{}); ok {
					data = append(data, ingressRules...)
				}
			}
			if v, ok := filterMap["EgressFilterRules"]; ok {
				if egressRules, ok := v.([]interface{}); ok {
					data = append(data, egressRules...)
				}
			}
		}

		return data, next, err
	})
}

func (s *VolcengineTrafficMirrorFilterRuleService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	ids := strings.Split(id, ":")
	if len(ids) != 2 {
		return data, fmt.Errorf("invald traffic mirror filter rule id: %s", id)
	}

	req := map[string]interface{}{
		"TrafficMirrorFilterIds.1": ids[0],
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		rule := make(map[string]interface{})
		if rule, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("Value is not map ")
		}
		if rule["TrafficMirrorFilterRuleId"] == ids[1] {
			data = rule
			break
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("traffic_mirror_filter_rule %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineTrafficMirrorFilterRuleService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				d          map[string]interface{}
				status     interface{}
				failStates []string
			)
			failStates = append(failStates, "Failed")
			d, err = s.ReadResource(resourceData, id)
			if err != nil {
				return nil, "", err
			}
			status, err = ve.ObtainSdkValue("Status", d)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("traffic_mirror_filter_rule status error, status: %s", status.(string))
				}
			}
			return d, status.(string), err
		},
	}
}

func (VolcengineTrafficMirrorFilterRuleService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineTrafficMirrorFilterRuleService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateTrafficMirrorFilterRule",
			ConvertMode: ve.RequestConvertAll,
			Convert:     map[string]ve.RequestConvert{},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("Result.TrafficMirrorFilterRuleId", *resp)
				filterId := d.Get("traffic_mirror_filter_id").(string)
				d.SetId(filterId + ":" + id.(string))
				return nil
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("traffic_mirror_filter_id").(string)
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Available"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineTrafficMirrorFilterRuleService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyTrafficMirrorFilterRuleAttributes",
			ConvertMode: ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"traffic_direction": {
					TargetField: "TrafficDirection",
				},
				"priority": {
					TargetField: "Priority",
				},
				"policy": {
					TargetField: "Policy",
				},
				"protocol": {
					TargetField: "Protocol",
				},
				"source_cidr_block": {
					TargetField: "SourceCidrBlock",
				},
				"source_port_range": {
					TargetField: "SourcePortRange",
				},
				"destination_cidr_block": {
					TargetField: "DestinationCidrBlock",
				},
				"destination_port_range": {
					TargetField: "DestinationPortRange",
				},
				"description": {
					TargetField: "Description",
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if len(*call.SdkParam) > 0 {
					ids := strings.Split(d.Id(), ":")
					if len(ids) != 2 {
						return false, fmt.Errorf("invalid traffic mirror filter rule id: %s", d.Id())
					}
					(*call.SdkParam)["TrafficMirrorFilterRuleId"] = ids[1]
					return true, nil
				}
				return false, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("traffic_mirror_filter_id").(string)
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Available"},
				Timeout: resourceData.Timeout(schema.TimeoutUpdate),
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineTrafficMirrorFilterRuleService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteTrafficMirrorFilterRule",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				ids := strings.Split(d.Id(), ":")
				if len(ids) != 2 {
					return false, fmt.Errorf("invalid traffic mirror filter rule id: %s", d.Id())
				}
				(*call.SdkParam)["TrafficMirrorFilterRuleId"] = ids[1]
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("traffic_mirror_filter_id").(string)
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
							return resource.NonRetryableError(fmt.Errorf("error on reading traffic mirror filter rule on delete %q, %w", d.Id(), callErr))
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

func (s *VolcengineTrafficMirrorFilterRuleService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"traffic_mirror_filter_ids": {
				TargetField: "TrafficMirrorFilterIds",
				ConvertType: ve.ConvertWithN,
			},
			"traffic_mirror_filter_names": {
				TargetField: "TrafficMirrorFilterNames",
				ConvertType: ve.ConvertWithN,
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
		IdField:      "TrafficMirrorFilterRuleId",
		CollectField: "traffic_mirror_filter_rules",
		ResponseConverts: map[string]ve.ResponseConvert{
			"TrafficMirrorFilterRuleId": {
				TargetField: "id",
				KeepDefault: true,
			},
		},
	}
}

func (s *VolcengineTrafficMirrorFilterRuleService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "vpc",
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		Action:      actionName,
	}
}
