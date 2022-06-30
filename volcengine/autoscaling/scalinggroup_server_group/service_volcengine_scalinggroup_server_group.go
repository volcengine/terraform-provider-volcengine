package scalinggroup_server_group

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineScalinggroupServerGroupService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewScalinggroupServerGroupService(c *ve.SdkClient) *VolcengineScalinggroupServerGroupService {
	return &VolcengineScalinggroupServerGroupService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineScalinggroupServerGroupService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineScalinggroupServerGroupService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 20, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		universalClient := s.Client.UniversalClient
		action := "DescribeScalingGroups"
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
		results, err = ve.ObtainSdkValue("Result.ScalingGroups", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.ScalingGroups is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineScalinggroupServerGroupService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"ScalingGroupIds.1": id,
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
		return data, fmt.Errorf("ScalingGroup %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineScalinggroupServerGroupService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				demo       map[string]interface{}
				status     interface{}
				failStates []string
			)
			failStates = append(failStates, "Error")
			demo, err = s.ReadResource(resourceData, id)
			if err != nil {
				return nil, "", err
			}
			status, err = ve.ObtainSdkValue("LifecycleState", demo)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("ScalingGroup  LifecycleState  error, status:%s", status.(string))
				}
			}
			//注意 返回的第一个参数不能为空 否则会一直等下去
			return demo, status.(string), err
		},
	}
}

func (VolcengineScalinggroupServerGroupService) WithResourceResponseHandlers(scalingGroup map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return scalingGroup, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}

}

func (s *VolcengineScalinggroupServerGroupService) CreateResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "AttachServerGroups",
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"server_group_attributes": {
					ConvertType: ve.ConvertListN,
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, d.Get("scaling_group_id").(string))
					if callErr != nil {
						return resource.NonRetryableError(fmt.Errorf("error on reading ScalingGroup %q, %w", d.Id(), callErr))
					}
					_, callErr = call.ExecuteCall(d, client, call)
					if callErr == nil {
						return nil
					}
					return resource.RetryableError(callErr)
				})
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				//注意 获取内容 这个地方不能是指针 需要转一次
				d.SetId(resourceData.Get("scaling_group_id").(string))
				return nil
			},
		},
	}
	return []ve.Callback{callback}

}

func (s *VolcengineScalinggroupServerGroupService) ModifyResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback
	// 修改负载均衡属性
	attrAdd, attrRemove, _, _ := ve.GetSetDifference("server_group_attributes", resourceData, serverGroupAttributeHash, false)
	removeLoadBalancerAttributeCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DetachServerGroups",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if attrRemove != nil && len(attrRemove.List()) > 0 {
					(*call.SdkParam)["ScalingGroupId"] = d.Id()
					for index, attr := range attrRemove.List() {
						(*call.SdkParam)["ServerGroupAttributes."+strconv.Itoa(index+1)+".ServerGroupId"] =
							attr.(map[string]interface{})["server_group_id"].(string)
					}
					return true, nil
				}
				return false, nil
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						return resource.NonRetryableError(fmt.Errorf("error on reading ScalingGroup %q, %w", d.Id(), callErr))
					}
					_, callErr = call.ExecuteCall(d, client, call)
					if callErr == nil {
						return nil
					}
					return resource.RetryableError(callErr)
				})
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	addLoadBalancerAttributeCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "AttachServerGroups",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if attrAdd != nil && len(attrAdd.List()) > 0 {
					(*call.SdkParam)["ScalingGroupId"] = d.Id()
					for index, attr := range attrAdd.List() {
						(*call.SdkParam)["ServerGroupAttributes."+strconv.Itoa(index+1)+".Port"] =
							attr.(map[string]interface{})["port"].(int)
						(*call.SdkParam)["ServerGroupAttributes."+strconv.Itoa(index+1)+".ServerGroupId"] =
							attr.(map[string]interface{})["server_group_id"].(string)
						(*call.SdkParam)["ServerGroupAttributes."+strconv.Itoa(index+1)+".Weight"] =
							attr.(map[string]interface{})["weight"].(int)
					}
					return true, nil
				}
				return false, nil
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
						return resource.NonRetryableError(fmt.Errorf("error on reading ScalingGroup %q, %w", d.Id(), callErr))
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
	callbacks = append(callbacks, removeLoadBalancerAttributeCallback, addLoadBalancerAttributeCallback)

	return callbacks
}

func (s *VolcengineScalinggroupServerGroupService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DetachServerGroups",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if d.Get("server_group_attributes") != nil && d.Get("server_group_attributes").(*schema.Set).Len() > 0 {
					(*call.SdkParam)["ScalingGroupId"] = d.Id()
					for index, attr := range d.Get("server_group_attributes").(*schema.Set).List() {
						(*call.SdkParam)["ServerGroupAttributes."+strconv.Itoa(index+1)+".ServerGroupId"] =
							attr.(map[string]interface{})["server_group_id"].(string)
					}
					return true, nil
				}
				return false, nil
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
							return resource.NonRetryableError(fmt.Errorf("error on reading ScalingGroup %q, %w", d.Id(), callErr))
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

func (s *VolcengineScalinggroupServerGroupService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineScalinggroupServerGroupService) ReadResourceId(id string) string {
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
