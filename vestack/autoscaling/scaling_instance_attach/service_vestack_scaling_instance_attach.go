package scaling_instance_attach

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-vestack/common"
	"github.com/volcengine/terraform-provider-vestack/logger"
)

type VestackScalingInstanceAttachService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewScalingInstanceAttachService(c *ve.SdkClient) *VestackScalingInstanceAttachService {
	return &VestackScalingInstanceAttachService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VestackScalingInstanceAttachService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VestackScalingInstanceAttachService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 50, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		autoScalingClient := s.Client.AutoScalingClient
		action := "DescribeScalingInstances"
		logger.Debug(logger.ReqFormat, action, condition)
		if condition == nil {
			resp, err = autoScalingClient.DescribeScalingInstancesCommon(nil)
			if err != nil {
				return data, err
			}
		} else {
			resp, err = autoScalingClient.DescribeScalingInstancesCommon(&condition)
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

func (s *VestackScalingInstanceAttachService) ReadResource(resourceData *schema.ResourceData, id string) (res map[string]interface{}, err error) {
	var (
		results []interface{}
		lossIds []string
		data    = make(map[string]interface{})
	)
	if len(id) == 0 {
		id = resourceData.Id()
	}
	ids := strings.Split(id, ":")
	req := formatInstanceIdsRequest(ids[1:])
	req["ScalingGroupId"] = ids[0]
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		tmpData, ok := v.(map[string]interface{})
		if !ok {
			return data, errors.New("Value is not map ")
		}
		data[tmpData["InstanceId"].(string)] = tmpData
	}
	for _, tmpId := range ids[1:] {
		if _, ok := data[tmpId]; !ok {
			lossIds = append(lossIds, tmpId)
		}
	}
	if len(lossIds) > 0 {
		return data, fmt.Errorf("instance not found: %s", strings.Join(lossIds, ","))
	}
	return data, nil
}

func (s *VestackScalingInstanceAttachService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				demo          map[string]interface{}
				targetStatus  = target[0]
				mutableStatus = map[string]bool{
					"Init":     true,
					"Pending":  true,
					"Removing": true,
				}
				existIdMap        = make(map[string]bool)
				existIds, lossIds []string
			)
			ids := strings.Split(id, ":")[1:]

			// 查看伸缩组状态
			resp, err := s.Client.AutoScalingClient.DescribeScalingGroupsCommon(&map[string]interface{}{"ScalingGroupId": ids[0]})
			if err != nil {
				return resp, "Error", err
			}
			groups, err := ve.ObtainSdkValue("Result.ScalingGroups", *resp)
			if err != nil || groups == nil {
				return resp, "Error", errors.New("invalid scaling group id")
			}
			g := groups.([]interface{})
			if len(g) == 0 {
				return resp, "Error", errors.New("invalid scaling group id")
			}
			if g[0].(map[string]interface{})["LifecycleState"].(string) == "Locked" {
				return g[0], "Locked", errors.New("scaling group has a active activity")
			}

			// 查看伸缩实例
			demo, err = s.ReadResource(resourceData, id)
			if err != nil && !strings.Contains(err.Error(), "instance not found") {
				return nil, "", err
			}
			for instanceId, data := range demo {
				tmpStatus, ok := data.(map[string]interface{})["Status"].(string)
				if !ok {
					return demo, "", errors.New("fail to get instance status")
				}
				if _, ok = mutableStatus[tmpStatus]; ok {
					return nil, "", fmt.Errorf("instance %s is in mutable state: %s", instanceId, tmpStatus)
				}
				existIdMap[instanceId] = true
			}
			for _, tmpId := range ids {
				if existIdMap[tmpId] {
					existIds = append(existIds, tmpId)
				} else {
					lossIds = append(lossIds, tmpId)
				}
			}
			if targetStatus == "Attached" {
				if len(lossIds) == 0 {
					return demo, targetStatus, nil
				}
				return demo, "Error", fmt.Errorf("%s attach fail", strings.Join(lossIds, "、"))
			} else if targetStatus == "Removed" {
				if len(existIds) == 0 {
					return demo, targetStatus, nil
				}
				return demo, "Error", fmt.Errorf("%s remove fail", strings.Join(existIds, "、"))
			}

			//注意 返回的第一个参数不能为空 否则会一直等下去
			return demo, "Error", fmt.Errorf("dont support target status: %s", targetStatus)
		},
	}

}

func (VestackScalingInstanceAttachService) WithResourceResponseHandlers(scalingGroup map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return scalingGroup, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}

}

func (s *VestackScalingInstanceAttachService) CreateResource(d *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	instanceIds := d.Get("instance_ids").(*schema.Set)
	return []ve.Callback{s.attachInstances(d.Get("scaling_group_id").(string), convertSliceInterfaceToString(instanceIds.List()), d.Timeout(schema.TimeoutUpdate))}
}

func (s *VestackScalingInstanceAttachService) ModifyResource(d *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var (
		callbacks      []ve.Callback
		scalingGroupId = d.Get("scaling_group_id").(string)
	)

	attachIds, removeIds, _, _ := ve.GetSetDifference("instance_ids", d, schema.HashString, false)
	if removeIds != nil && removeIds.Len() > 0 {
		callbacks = append(callbacks, s.removeInstances(scalingGroupId, convertSliceInterfaceToString(removeIds.List()), d.Timeout(schema.TimeoutUpdate)))
	}

	if attachIds != nil && attachIds.Len() > 0 {
		callbacks = append(callbacks, s.attachInstances(scalingGroupId, convertSliceInterfaceToString(attachIds.List()), d.Timeout(schema.TimeoutUpdate)))
	}

	if len(callbacks) > 0 {
		callbacks = append(callbacks, ve.Callback{
			Call: ve.SdkCall{
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					// 更新resource id
					d.SetId(fmt.Sprintf("%s:%s", d.Get("scaling_group_id"),
						strings.Join(convertSliceInterfaceToString(d.Get("instance_ids").(*schema.Set).List()), ":")))
					return &map[string]interface{}{}, nil
				},
			},
		})
	}

	return callbacks
}

func (s *VestackScalingInstanceAttachService) RemoveResource(d *schema.ResourceData, r *schema.Resource) []ve.Callback {
	instanceIds := d.Get("instance_ids").(*schema.Set)
	return []ve.Callback{s.removeInstances(d.Get("scaling_group_id").(string), convertSliceInterfaceToString(instanceIds.List()), d.Timeout(schema.TimeoutDelete))}
}

func (s *VestackScalingInstanceAttachService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VestackScalingInstanceAttachService) ReadResourceId(id string) string {
	return id
}

func (s *VestackScalingInstanceAttachService) attachInstances(groupId string, instanceIds []string, timeout time.Duration) ve.Callback {
	return ve.Callback{
		Call: ve.SdkCall{
			Action:      "AttachInstances",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if len(instanceIds) < 0 {
					return false, nil
				}
				param := formatInstanceIdsRequest(instanceIds)
				param["ScalingGroupId"] = groupId
				*call.SdkParam = param
				d.SetId(fmt.Sprintf("%s:%s", groupId, strings.Join(instanceIds, ":")))
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				common, err := s.Client.AutoScalingClient.AttachInstancesCommon(call.SdkParam)
				if err != nil {
					return common, err
				}
				time.Sleep(10 * time.Second) // attach以后需要等一下，否则查不到数据
				return common, nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Attached"},
				Timeout: timeout,
			},
		},
	}
}

func (s *VestackScalingInstanceAttachService) removeInstances(groupId string, instanceIds []string, timeout time.Duration) ve.Callback {
	return ve.Callback{
		Call: ve.SdkCall{
			Action:      "RemoveInstances",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if len(instanceIds) < 0 {
					return false, nil
				}
				param := formatInstanceIdsRequest(instanceIds)
				param["ScalingGroupId"] = groupId
				*call.SdkParam = param
				d.SetId(fmt.Sprintf("%s:%s", groupId, strings.Join(instanceIds, ":")))
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				common, err := s.Client.AutoScalingClient.RemoveInstancesCommon(call.SdkParam)
				if err != nil {
					return common, err
				}
				time.Sleep(5 * time.Second)
				d.SetId(fmt.Sprintf("%s:%s", groupId, strings.Join(instanceIds, ":")))
				return common, nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Removed"},
				Timeout: timeout,
			},
		},
	}
}
