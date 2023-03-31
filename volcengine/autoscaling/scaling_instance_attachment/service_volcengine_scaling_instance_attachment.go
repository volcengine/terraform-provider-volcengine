package scaling_instance_attachment

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

type VolcengineScalingInstanceAttachmentService struct {
	Client *ve.SdkClient
}

func NewScalingInstanceAttachmentService(c *ve.SdkClient) *VolcengineScalingInstanceAttachmentService {
	return &VolcengineScalingInstanceAttachmentService{
		Client: c,
	}
}

func (s *VolcengineScalingInstanceAttachmentService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineScalingInstanceAttachmentService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
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
		logger.Debug(logger.RespFormat, action, m, resp, condition)
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

func (s *VolcengineScalingInstanceAttachmentService) ReadResource(resourceData *schema.ResourceData, id string) (res map[string]interface{}, err error) {
	var (
		results    []interface{}
		data       = make(map[string]interface{})
		instanceId string
		status     string
	)
	if len(id) == 0 {
		id = resourceData.Id()
	}
	ids := strings.Split(id, ":")
	// 查询伸缩组下所有实例id
	results, err = s.ReadResources(map[string]interface{}{
		"ScalingGroupId": ids[0],
		"InstanceIds.1":  ids[1],
	})
	if err != nil {
		return data, err
	}
	if len(results) == 0 {
		return data, errors.New("instance not found")
	}
	tempData, ok := results[0].(map[string]interface{})
	if !ok {
		return data, errors.New("value is not map")
	}
	instanceId = tempData["InstanceId"].(string)
	status = tempData["Status"].(string)
	data["InstanceId"] = instanceId
	data["Status"] = status

	return data, nil
}

func (s *VolcengineScalingInstanceAttachmentService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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
					return nil, "", fmt.Errorf("instance status error, status:%s", status.(string))
				}
			}
			//注意 返回的第一个参数不能为空 否则会一直等下去
			return demo, status.(string), err
		},
	}

}

func (VolcengineScalingInstanceAttachmentService) WithResourceResponseHandlers(scalingGroup map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return scalingGroup, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}

}

func (s *VolcengineScalingInstanceAttachmentService) CreateResource(d *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	instanceId := d.Get("instance_id").(string)
	return s.attachInstances(d.Get("scaling_group_id").(string), instanceId, d.Timeout(schema.TimeoutUpdate))
}

func (s *VolcengineScalingInstanceAttachmentService) ModifyResource(d *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineScalingInstanceAttachmentService) RemoveResource(d *schema.ResourceData, r *schema.Resource) []ve.Callback {
	instanceId := d.Get("instance_id").(string)
	deleteType := d.Get("delete_type").(string)
	detachOption := d.Get("detach_option").(string)
	return s.removeInstances(d.Get("scaling_group_id").(string), instanceId, d.Timeout(schema.TimeoutDelete), deleteType, detachOption)
}

func (s *VolcengineScalingInstanceAttachmentService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineScalingInstanceAttachmentService) ReadResourceId(id string) string {
	return id
}

func (s *VolcengineScalingInstanceAttachmentService) attachInstances(groupId string, instanceId string, timeout time.Duration) []ve.Callback {
	callbacks := make([]ve.Callback, 0)
	attachCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "AttachInstances",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				logger.Debug(logger.RespFormat, call.Action, instanceId)
				param := formatInstanceIdsRequest(instanceId)
				param["ScalingGroupId"] = groupId
				if entrusted, ok := d.GetOk("entrusted"); ok {
					param["Entrusted"] = entrusted
				}
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
				d.SetId(fmt.Sprint((*call.SdkParam)["ScalingGroupId"], ":", (*call.SdkParam)["InstanceIds.1"]))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"InService", "Protected"},
				Timeout: timeout,
			},
		},
	}
	callbacks = append(callbacks, attachCallback)
	return callbacks
}

func (s *VolcengineScalingInstanceAttachmentService) removeInstances(groupId string, instanceId string, timeout time.Duration, deleteType, detachOption string) []ve.Callback {
	var action string
	if deleteType == "Detach" {
		action = "DetachInstances"
	} else {
		// 默认remove
		action = "RemoveInstances"
	}
	if detachOption != "none" {
		detachOption = "both"
	}
	callbacks := make([]ve.Callback, 0)
	removeCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      action,
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				param := formatInstanceIdsRequest(instanceId)
				param["ScalingGroupId"] = groupId
				if action == "DetachInstances" {
					param["DetachOption"] = detachOption
				}
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
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, d.Id())
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading scaling instance on delete #{d.Id()}, #{callErr}"))
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
	callbacks = append(callbacks, removeCallback)
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
