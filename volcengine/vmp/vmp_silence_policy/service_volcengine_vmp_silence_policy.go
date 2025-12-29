package vmp_silence_policy

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineVmpSilencePolicyService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

// NewVmpSilencePolicyService 创建并返回 VMP 静默策略服务实例
func NewVmpSilencePolicyService(c *ve.SdkClient) *VolcengineVmpSilencePolicyService {
	return &VolcengineVmpSilencePolicyService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

// GetClient 返回 SDK 客户端
func (s *VolcengineVmpSilencePolicyService) GetClient() *ve.SdkClient {
	return s.Client
}

// ReadResources 分页查询静默策略列表
func (s *VolcengineVmpSilencePolicyService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 10, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListSilencePolicies"

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
		// TODO: replace result items
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

// ReadResource 根据 Id 查询并返回静默策略详情
func (s *VolcengineVmpSilencePolicyService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
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
		return data, fmt.Errorf("vmp_silence_policy %s not exist ", id)
	}
	return data, err
}

// RefreshResourceState 资源状态刷新（静默策略无需状态轮询）
func (s *VolcengineVmpSilencePolicyService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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
					return nil, "", fmt.Errorf("vmp_silence_policy status error, status: %s", status.(string))
				}
			}
			return d, status.(string), err
		},
	}
}

// CreateResource 创建静默策略
func (s *VolcengineVmpSilencePolicyService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateSilencePolicy",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"name":        {TargetField: "Name"},
				"description": {TargetField: "Description"},
				"time_range_matchers": {
					TargetField: "TimeRangeMatchers",
					ConvertType: ve.ConvertJsonObjectArray,
					NextLevelConvert: map[string]ve.RequestConvert{
						"date":     {TargetField: "Date"},
						"location": {TargetField: "Location"},
						"periodic_date": {
							TargetField: "PeriodicDate",
							ConvertType: ve.ConvertJsonObject,
							NextLevelConvert: map[string]ve.RequestConvert{
								"time":         {TargetField: "Time"},
								"weekday":      {TargetField: "Weekday"},
								"day_of_month": {TargetField: "DayOfMonth"},
							},
						},
					},
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if v, ok := d.GetOk("metric_label_matchers"); ok {
					lm := expandLabelMatchers(v.([]interface{}))
					(*call.SdkParam)["MetricLabelMatchers"] = lm
				}
				return true, nil
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

// WithResourceResponseHandlers 资源读回填的响应转换（默认映射）
func (VolcengineVmpSilencePolicyService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		// API 返回的 PeriodicDate 是 Object，TF Schema 定义为 List (MaxItems: 1)
		// 需要手动将 Object 包装为 List
		if v, ok := d["TimeRangeMatchers"]; ok {
			if list, ok := v.([]interface{}); ok {
				for _, item := range list {
					if m, ok := item.(map[string]interface{}); ok {
						if pd, ok := m["PeriodicDate"]; ok && pd != nil {
							// 如果已经是 Slice 就不处理（防止重复处理）
							if _, isSlice := pd.([]interface{}); !isSlice {
								m["PeriodicDate"] = []interface{}{pd}
							}
						}
					}
				}
			}
		}

		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

// ModifyResource 更新静默策略
func (s *VolcengineVmpSilencePolicyService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateSilencePolicy",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"name":        {TargetField: "Name"},
				"description": {TargetField: "Description"},
				"time_range_matchers": {
					TargetField: "TimeRangeMatchers",
					ConvertType: ve.ConvertJsonObjectArray,
					NextLevelConvert: map[string]ve.RequestConvert{
						"date":     {TargetField: "Date"},
						"location": {TargetField: "Location"},
						"periodic_date": {
							TargetField: "PeriodicDate",
							ConvertType: ve.ConvertJsonObject,
							NextLevelConvert: map[string]ve.RequestConvert{
								"time":         {TargetField: "Time"},
								"weekday":      {TargetField: "Weekday"},
								"day_of_month": {TargetField: "DayOfMonth"},
							},
						},
					},
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["Id"] = d.Id()
				if v, ok := d.GetOk("metric_label_matchers"); ok {
					lm := expandLabelMatchers(v.([]interface{}))
					(*call.SdkParam)["MetricLabelMatchers"] = lm
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

// RemoveResource 删除静默策略
func (s *VolcengineVmpSilencePolicyService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteSilencePolicies",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"Ids": []interface{}{resourceData.Id()},
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

// DatasourceResources 数据源请求与响应映射
func (s *VolcengineVmpSilencePolicyService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "Filter.Ids",
				ConvertType: ve.ConvertJsonArray,
			},
			"name": {
				TargetField: "Filter.Name",
			},
			"status": {
				TargetField: "Filter.Status",
			},
			"sources": {
				TargetField: "Filter.Sources",
				ConvertType: ve.ConvertJsonArray,
			},
		},
		ContentType:  ve.ContentTypeJson,
		NameField:    "Name",
		IdField:      "Id",
		CollectField: "silence_policies",
	}
}

// ReadResourceId 返回资源唯一标识
func (s *VolcengineVmpSilencePolicyService) ReadResourceId(id string) string {
	return id
}

// getUniversalInfo 统一网关请求信息
func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "vmp",
		Version:     "2021-03-03",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}

// expandLabelMatchers 将 Terraform 结构展开为 API 所需的 Matchers 二维数组
func expandLabelMatchers(groups []interface{}) []interface{} {
	var out []interface{}
	for _, g := range groups {
		if g == nil {
			continue
		}
		gm, ok := g.(map[string]interface{})
		if !ok {
			continue
		}
		var one []interface{}
		if items, ok := gm["matchers"].(*schema.Set); ok {
			for _, m := range items.List() {
				if m == nil {
					continue
				}
				mv, ok := m.(map[string]interface{})
				if !ok {
					continue
				}
				// operator 不一定存在，需要判断是否存在
				if operator, ok := mv["operator"]; ok {
					one = append(one, map[string]interface{}{
						"Label":    mv["label"],
						"Value":    mv["value"],
						"Operator": operator,
					})
				} else {
					one = append(one, map[string]interface{}{
						"Label": mv["label"],
						"Value": mv["value"],
					})
				}
			}
		}
		if len(one) > 0 {
			out = append(out, map[string]interface{}{
				"Matchers": one,
			})
		}
	}
	return out
}
