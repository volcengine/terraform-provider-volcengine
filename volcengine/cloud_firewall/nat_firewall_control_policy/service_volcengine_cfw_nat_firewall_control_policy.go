package nat_firewall_control_policy

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

type VolcengineNatFirewallControlPolicyService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewNatFirewallControlPolicyService(c *ve.SdkClient) *VolcengineNatFirewallControlPolicyService {
	return &VolcengineNatFirewallControlPolicyService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineNatFirewallControlPolicyService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineNatFirewallControlPolicyService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeNatFirewallControlPolicy"

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

func (s *VolcengineNatFirewallControlPolicyService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	ids := strings.Split(id, ":")
	if len(ids) != 3 {
		return data, fmt.Errorf("Invalid nat firewall control policy id: %s ", id)
	}

	req := map[string]interface{}{
		"Direction":     ids[0],
		"NatFirewallId": ids[1],
		"RuleId": []interface{}{
			ids[2],
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
		return data, fmt.Errorf("nat_firewall_control_policy %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineNatFirewallControlPolicyService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (VolcengineNatFirewallControlPolicyService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineNatFirewallControlPolicyService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "AddNatFirewallControlPolicy",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"priority": {
					TargetField: "Prio",
				},
				"repeat_days": {
					TargetField: "RepeatDays",
					ConvertType: ve.ConvertJsonArray,
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				ruleId, _ := ve.ObtainSdkValue("Result.RuleId", *resp)
				direction := d.Get("direction").(string)
				natFirewallId := d.Get("nat_firewall_id").(string)
				d.SetId(direction + ":" + natFirewallId + ":" + ruleId.(string))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineNatFirewallControlPolicyService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyNatFirewallControlPolicy",
			ConvertMode: ve.RequestConvertInConvert,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"action": {
					TargetField: "Action",
					ForceGet:    true,
				},
				"destination_type": {
					TargetField: "DestinationType",
					ForceGet:    true,
				},
				"destination": {
					TargetField: "Destination",
					ForceGet:    true,
				},
				"proto": {
					TargetField: "Proto",
					ForceGet:    true,
				},
				"source_type": {
					TargetField: "SourceType",
					ForceGet:    true,
				},
				"source": {
					TargetField: "Source",
					ForceGet:    true,
				},
				"description": {
					TargetField: "Description",
					ForceGet:    true,
				},
				"dest_port_type": {
					TargetField: "DestPortType",
					ForceGet:    true,
				},
				"dest_port": {
					TargetField: "DestPort",
					ForceGet:    true,
				},
				"repeat_type": {
					TargetField: "RepeatType",
					ForceGet:    true,
				},
				"repeat_start_time": {
					TargetField: "RepeatStartTime",
					ForceGet:    true,
				},
				"repeat_end_time": {
					TargetField: "RepeatEndTime",
					ForceGet:    true,
				},
				"repeat_days": {
					TargetField: "RepeatDays",
					ConvertType: ve.ConvertJsonArray,
					ForceGet:    true,
				},
				"start_time": {
					TargetField: "StartTime",
					ForceGet:    true,
				},
				"end_time": {
					TargetField: "EndTime",
					ForceGet:    true,
				},
				"prio": {
					Ignore: true,
				},
				"status": {
					TargetField: "Status",
					ForceGet:    true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				ids := strings.Split(d.Id(), ":")
				if len(ids) != 3 {
					return false, fmt.Errorf("Invalid nat firewall control policy id: %s ", d.Id())
				}
				(*call.SdkParam)["Direction"] = ids[0]
				(*call.SdkParam)["NatFirewallId"] = ids[1]
				(*call.SdkParam)["RuleId"] = ids[2]
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
	callbacks = append(callbacks, callback)

	return callbacks
}

func (s *VolcengineNatFirewallControlPolicyService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteNatFirewallControlPolicy",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				ids := strings.Split(d.Id(), ":")
				if len(ids) != 3 {
					return false, fmt.Errorf("Invalid nat firewall control policy id: %s ", d.Id())
				}
				(*call.SdkParam)["Direction"] = ids[0]
				(*call.SdkParam)["NatFirewallId"] = ids[1]
				(*call.SdkParam)["RuleId"] = ids[2]
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
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
							return resource.NonRetryableError(fmt.Errorf("error on reading nat firewall control policy on delete %q, %w", d.Id(), callErr))
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

func (s *VolcengineNatFirewallControlPolicyService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"rule_id": {
				TargetField: "RuleId",
				ConvertType: ve.ConvertJsonArray,
			},
			"action": {
				TargetField: "Action",
				ConvertType: ve.ConvertJsonArray,
			},
			"repeat_type": {
				TargetField: "RepeatType",
				ConvertType: ve.ConvertJsonArray,
			},
			"proto": {
				TargetField: "Proto",
				ConvertType: ve.ConvertJsonArray,
			},
			"status": {
				TargetField: "Status",
				ConvertType: ve.ConvertJsonArray,
			},
			"dest_port": {
				TargetField: "DestPort",
				ConvertType: ve.ConvertJsonArray,
			},
			"destination": {
				TargetField: "Destination",
				ConvertType: ve.ConvertJsonArray,
			},
			"source": {
				TargetField: "Source",
				ConvertType: ve.ConvertJsonArray,
			},
		},
		IdField:      "RuleId",
		CollectField: "nat_firewall_control_policies",
		ContentType:  ve.ContentTypeJson,
		ResponseConverts: map[string]ve.ResponseConvert{
			"RuleId": {
				TargetField: "id",
				KeepDefault: true,
			},
		},
	}
}

func (s *VolcengineNatFirewallControlPolicyService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "fw_center",
		Version:     "2021-09-06",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}
