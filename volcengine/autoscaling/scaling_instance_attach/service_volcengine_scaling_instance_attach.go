package scaling_instance_attach

import (
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineScalingInstanceAttachService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewScalingInstanceAttachService(c *ve.SdkClient) *VolcengineScalingInstanceAttachService {
	return &VolcengineScalingInstanceAttachService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineScalingInstanceAttachService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineScalingInstanceAttachService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 50, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		universalClient := s.Client.UniversalClient
		action := "DescribeScalingInstances"
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
		results, err = ve.ObtainSdkValue("Result.ScalingInstances", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.ScalingInstances is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineScalingInstanceAttachService) ReadResource(resourceData *schema.ResourceData, id string) (res map[string]interface{}, err error) {
	var (
		results     []interface{}
		data        = make(map[string]interface{})
		instanceIds = make([]string, 0)
	)
	if len(id) == 0 {
		id = resourceData.Id()
	}

	// 查询伸缩组下所有实例id
	results, err = s.ReadResources(map[string]interface{}{"ScalingGroupId": id})
	if err != nil {
		return data, err
	}
	for _, v := range results {
		tmpData, ok := v.(map[string]interface{})
		if !ok {
			return data, errors.New("Value is not map ")
		}
		instanceIds = append(instanceIds, tmpData["InstanceId"].(string))
	}

	// 查看伸缩组状态
	resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo("DescribeScalingGroups"),
		&map[string]interface{}{"ScalingGroupIds.1": id})
	if err != nil {
		return data, fmt.Errorf("describe scaling group err: %s", err.Error())
	}
	groups, err := ve.ObtainSdkValue("Result.ScalingGroups", *resp)
	if groups == nil || len(groups.([]interface{})) == 0 {
		return data, fmt.Errorf("scaling group %s not found", id)
	}

	data["InstanceIds"] = instanceIds
	data["Status"] = groups.([]interface{})[0].(map[string]interface{})["LifecycleState"].(string)

	return data, nil
}

func (s *VolcengineScalingInstanceAttachService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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
			status, err = ve.ObtainSdkValue("Status", demo)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("scaling group Status error, status:%s", status.(string))
				}
			}
			//注意 返回的第一个参数不能为空 否则会一直等下去
			return demo, status.(string), err
		},
	}

}

func (VolcengineScalingInstanceAttachService) WithResourceResponseHandlers(scalingGroup map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return scalingGroup, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}

}

func (s *VolcengineScalingInstanceAttachService) CreateResource(d *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	instanceIds := d.Get("instance_ids").(*schema.Set)
	return s.attachInstances(d.Get("scaling_group_id").(string), convertSliceInterfaceToString(instanceIds.List()), d.Timeout(schema.TimeoutUpdate))
}

func (s *VolcengineScalingInstanceAttachService) ModifyResource(d *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var (
		callbacks      []ve.Callback
		scalingGroupId = d.Get("scaling_group_id").(string)
	)

	attachIds, removeIds, _, _ := ve.GetSetDifference("instance_ids", d, schema.HashString, false)
	if removeIds != nil && removeIds.Len() > 0 {
		callbacks = append(callbacks, s.removeInstances(scalingGroupId, convertSliceInterfaceToString(removeIds.List()), d.Timeout(schema.TimeoutUpdate))...)
	}

	if attachIds != nil && attachIds.Len() > 0 {
		callbacks = append(callbacks, s.attachInstances(scalingGroupId, convertSliceInterfaceToString(attachIds.List()), d.Timeout(schema.TimeoutUpdate))...)
	}

	return callbacks
}

func (s *VolcengineScalingInstanceAttachService) RemoveResource(d *schema.ResourceData, r *schema.Resource) []ve.Callback {
	instanceIds := d.Get("instance_ids").(*schema.Set)
	return s.removeInstances(d.Get("scaling_group_id").(string), convertSliceInterfaceToString(instanceIds.List()), d.Timeout(schema.TimeoutDelete))
}

func (s *VolcengineScalingInstanceAttachService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineScalingInstanceAttachService) ReadResourceId(id string) string {
	return id
}

func (s *VolcengineScalingInstanceAttachService) attachInstances(groupId string, instanceIds []string, timeout time.Duration) []ve.Callback {
	callbacks := make([]ve.Callback, 0)
	if len(instanceIds) == 0 {
		return callbacks
	}
	for i := 0; i < int(math.Ceil(float64(len(instanceIds))/float64(20))); i++ {
		max := (i + 1) * 20
		if max > len(instanceIds) {
			max = len(instanceIds)
		}
		callbacks = append(callbacks, func(ids []string) ve.Callback {
			return ve.Callback{
				Call: ve.SdkCall{
					Action:      "AttachInstances",
					ConvertMode: ve.RequestConvertIgnore,
					BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
						if len(ids) < 0 {
							return false, nil
						}
						logger.Debug(logger.RespFormat, call.Action, ids)
						param := formatInstanceIdsRequest(ids)
						param["ScalingGroupId"] = groupId
						*call.SdkParam = param
						return true, nil
					},
					ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
						logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
						common, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
						if err != nil {
							return common, err
						}
						time.Sleep(10 * time.Second) // attach以后需要等一下
						return common, nil
					},
					AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
						d.SetId(d.Get("scaling_group_id").(string))
						return nil
					},
					Refresh: &ve.StateRefresh{
						Target:  []string{"Active"},
						Timeout: timeout,
					},
				},
			}
		}(instanceIds[i*20:max]))
	}
	return callbacks
}

func (s *VolcengineScalingInstanceAttachService) removeInstances(groupId string, instanceIds []string, timeout time.Duration) []ve.Callback {
	callbacks := make([]ve.Callback, 0)
	if len(instanceIds) == 0 {
		return callbacks
	}
	for i := 0; i < int(math.Ceil(float64(len(instanceIds))/float64(20))); i++ {
		max := (i + 1) * 20
		if max > len(instanceIds) {
			max = len(instanceIds)
		}
		callbacks = append(callbacks, func(ids []string) ve.Callback {
			return ve.Callback{
				Call: ve.SdkCall{
					Action:      "RemoveInstances",
					ConvertMode: ve.RequestConvertIgnore,
					BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
						if len(ids) < 0 {
							return false, nil
						}
						param := formatInstanceIdsRequest(ids)
						param["ScalingGroupId"] = groupId
						*call.SdkParam = param
						return true, nil
					},
					ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
						logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
						common, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
						if err != nil {
							return common, err
						}
						time.Sleep(10 * time.Second) // remove以后需要等一下
						return common, nil
					},
					Refresh: &ve.StateRefresh{
						Target:  []string{"Active"},
						Timeout: timeout,
					},
				},
			}
		}(instanceIds[i*20:max]))
	}
	return callbacks
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
