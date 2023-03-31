package scaling_lifecycle_hook

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

type VolcengineScalingLifecycleHookService struct {
	Client *ve.SdkClient
}

func NewScalingLifecycleHookService(c *ve.SdkClient) *VolcengineScalingLifecycleHookService {
	return &VolcengineScalingLifecycleHookService{
		Client: c,
	}
}

func (s *VolcengineScalingLifecycleHookService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineScalingLifecycleHookService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
		nameSet = make(map[string]bool)
	)
	if _, ok = m["LifecycleHookNames.1"]; ok {
		i := 1
		for {
			field := fmt.Sprintf("LifecycleHookNames.%d", i)
			if name, ok := m[field]; ok {
				nameSet[name.(string)] = true
				i = i + 1
				delete(m, field)
			} else {
				break
			}
		}
	}

	hooks, err := ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 20, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		universalClient := s.Client.UniversalClient
		action := "DescribeLifecycleHooks"
		logger.Debug(logger.ReqFormat, action, condition)
		if condition == nil {
			resp, err = universalClient.DoCall(getUniversalInfo(action), nil)
			if err != nil {
				return data, err
			}
		} else {
			resp, err = universalClient.DoCall(getUniversalInfo(action), &condition)
			if err != nil {
				return data, err
			}
		}
		logger.Debug(logger.RespFormat, action, action, resp)
		results, err = ve.ObtainSdkValue("Result.LifecycleHooks", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.LifecycleHooks is not Slice")
		}
		return data, err
	})
	if err != nil {
		return hooks, err
	}
	res := make([]interface{}, 0)
	for _, ele := range data {
		e, ok := ele.(map[string]interface{})
		if !ok {
			continue
		}
		name := e["LifecycleHookName"].(string)
		if len(nameSet) == 0 || nameSet[name] {
			res = append(res, ele)
		}
	}
	return res, nil
}

func (s *VolcengineScalingLifecycleHookService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	ids := strings.Split(id, ":")
	req := map[string]interface{}{
		"LifecycleHookIds.1": ids[1],
		"ScalingGroupId":     ids[0],
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
		return data, fmt.Errorf("ScalingLifecycleHook %s not exist ", ids[1])
	}
	return data, err
}

func (s *VolcengineScalingLifecycleHookService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (VolcengineScalingLifecycleHookService) WithResourceResponseHandlers(scalingGroup map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return scalingGroup, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}

}

func (s *VolcengineScalingLifecycleHookService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateLifecycleHook",
			ConvertMode: ve.RequestConvertAll,
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				//注意 获取内容 这个地方不能是指针 需要转一次
				id, _ := ve.ObtainSdkValue("Result.LifecycleHookId", *resp)
				logger.Debug(logger.RespFormat, call.Action, resourceData.Get("scaling_group_id"))
				d.SetId(fmt.Sprintf("%v:%v", resourceData.Get("scaling_group_id"), id.(string)))
				return nil
			},
		},
	}
	return []ve.Callback{callback}

}

func (s *VolcengineScalingLifecycleHookService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	ids := strings.Split(resourceData.Id(), ":")
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:         "ModifyLifecycleHook",
			ConvertMode:    ve.RequestConvertAll,
			RequestIdField: "LifecycleHookId",
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if len(*call.SdkParam) < 1 {
					return false, nil
				}
				(*call.SdkParam)["LifecycleHookId"] = ids[1]
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineScalingLifecycleHookService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	ids := strings.Split(resourceData.Id(), ":")
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteLifecycleHook",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"LifecycleHookId": ids[1],
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading LifecycleHook on delete %q, %w", d.Id(), callErr))
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

func (s *VolcengineScalingLifecycleHookService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "LifecycleHookIds",
				ConvertType: ve.ConvertWithN,
			},
			"lifecycle_hook_names": {
				TargetField: "LifecycleHookNames",
				ConvertType: ve.ConvertWithN,
			},
		},
		NameField:    "LifecycleHookName",
		IdField:      "LifecycleHookId",
		CollectField: "lifecycle_hooks",
		ResponseConverts: map[string]ve.ResponseConvert{
			"LifecycleHookId": {
				TargetField: "id",
				KeepDefault: true,
			},
		},
	}
}

func (s *VolcengineScalingLifecycleHookService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "auto_scaling",
		Action:      actionName,
		Version:     "2020-01-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
	}
}
