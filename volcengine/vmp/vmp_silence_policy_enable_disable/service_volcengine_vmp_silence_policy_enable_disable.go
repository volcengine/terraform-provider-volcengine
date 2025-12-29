package vmp_silence_policy_enable_disable

import (
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineVmpSilencePolicyEnableDisableService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewVmpSilencePolicyEnableDisableService(c *ve.SdkClient) *VolcengineVmpSilencePolicyEnableDisableService {
	return &VolcengineVmpSilencePolicyEnableDisableService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

// GetClient 返回 SDK 客户端
func (s *VolcengineVmpSilencePolicyEnableDisableService) GetClient() *ve.SdkClient {
	return s.Client
}

// ReadResources 分页查询静默策略列表（用于校验状态）
func (s *VolcengineVmpSilencePolicyEnableDisableService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
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

// ReadResource 根据 ids 过滤查询静默策略并返回详情（用于 Read 校验）
func (s *VolcengineVmpSilencePolicyEnableDisableService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	ids := []interface{}{}
	if v, ok := resourceData.GetOk("ids"); ok {
		for _, ele := range v.(*schema.Set).List() {
			ids = append(ids, ele)
		}
	}
	req := map[string]interface{}{
		"Filter": map[string]interface{}{
			"Ids": ids,
		},
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	if len(results) == 0 {
		return data, fmt.Errorf("vmp_silence_policy_enable_disable target ids not exist")
	}
	if data, ok = results[0].(map[string]interface{}); !ok {
		return data, errors.New("Value is not map ")
	}
	// 写入 Ids 以在 Read 流程中保持状态与配置一致，避免 ForceNew 误替换
	idsList := []string{}
	if v, ok := resourceData.GetOk("ids"); ok {
		for _, ele := range v.(*schema.Set).List() {
			idsList = append(idsList, ele.(string))
		}
	}
	if len(idsList) > 0 {
		data["Ids"] = idsList
	}
	return data, err
}

// RefreshResourceState 资源状态刷新（静默策略无需状态轮询）
func (s *VolcengineVmpSilencePolicyEnableDisableService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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
					return nil, "", fmt.Errorf("vmp_silence_policy_enable_disable status error, status: %s", status.(string))
				}
			}
			return d, status.(string), err
		},
	}
}

// CreateResource 启用静默策略（批量）
func (s *VolcengineVmpSilencePolicyEnableDisableService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "EnableSilencePolicies",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"ids": {
					TargetField: "Ids",
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
				if resp != nil {
					if unsuccess, err := ve.ObtainSdkValue("Result.UnsuccessfulItems", *resp); err == nil && unsuccess != nil {
						if items, ok := unsuccess.([]interface{}); ok && len(items) > 0 {
							var errMsgs []string
							for _, item := range items {
								if itemMap, ok := item.(map[string]interface{}); ok {
									id := itemMap["Id"]
									errorObj := itemMap["Error"]
									errMsgs = append(errMsgs, fmt.Sprintf("Id: %v, Error: %v", id, errorObj))
								}
							}
							return fmt.Errorf("EnableSilencePolicies partially or fully failed: %s", strings.Join(errMsgs, "; "))
						}
					}

					if success, err := ve.ObtainSdkValue("Result.SuccessfulItems", *resp); err == nil && success != nil {
						if items, ok := success.([]interface{}); ok {
							idsList := make([]string, 0, len(items))
							for _, item := range items {
								if id, ok := item.(string); ok {
									idsList = append(idsList, id)
								}
							}
							sort.Strings(idsList)
							d.SetId(strings.Join(idsList, ","))
						}
					}
				}
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

// WithResourceResponseHandlers 资源读回填的响应转换（默认映射）
func (VolcengineVmpSilencePolicyEnableDisableService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"Ids": {
				TargetField: "ids",
				KeepDefault: true,
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

// ModifyResource 当前资源不支持更新操作
func (s *VolcengineVmpSilencePolicyEnableDisableService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

// RemoveResource 禁用静默策略（批量）
func (s *VolcengineVmpSilencePolicyEnableDisableService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DisableSilencePolicies",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"ids": {
					TargetField: "Ids",
					ConvertType: ve.ConvertJsonArray,
				},
			},
			SdkParam: &map[string]interface{}{},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				ids := []interface{}{}
				if v, ok := d.GetOk("ids"); ok {
					for _, ele := range v.(*schema.Set).List() {
						ids = append(ids, ele)
					}
				}
				if len(ids) == 0 {
					// 兼容导入或历史资源：从资源 ID 中解析 ids（逗号分隔）
					if rid := d.Id(); rid != "" {
						parts := strings.Split(rid, ",")
						for _, p := range parts {
							p = strings.TrimSpace(p)
							if p != "" {
								ids = append(ids, p)
							}
						}
					}
				}
				(*call.SdkParam)["Ids"] = ids
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				if resp != nil {
					if unsuccess, err := ve.ObtainSdkValue("Result.UnsuccessfulItems", *resp); err == nil && unsuccess != nil {
						if items, ok := unsuccess.([]interface{}); ok && len(items) > 0 {
							var errMsgs []string
							for _, item := range items {
								if itemMap, ok := item.(map[string]interface{}); ok {
									id := itemMap["Id"]
									errorObj := itemMap["Error"]
									errMsgs = append(errMsgs, fmt.Sprintf("Id: %v, Error: %v", id, errorObj))
								}
							}
							return fmt.Errorf("DisableSilencePolicies partially or fully failed: %s", strings.Join(errMsgs, "; "))
						}
					}
				}
				// 删除后，资源即被视为移除
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineVmpSilencePolicyEnableDisableService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

// ReadResourceId 返回资源唯一标识
func (s *VolcengineVmpSilencePolicyEnableDisableService) ReadResourceId(id string) string {
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
