package ecs_instance_state

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

type VolcengineInstanceStateService struct {
	Client *ve.SdkClient
}

func NewInstanceStateService(c *ve.SdkClient) *VolcengineInstanceStateService {
	return &VolcengineInstanceStateService{
		Client: c,
	}
}

func (s *VolcengineInstanceStateService) GetClient() *ve.SdkClient {
	return s.Client
}

func (VolcengineInstanceStateService) WithResourceResponseHandlers(subnet map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return subnet, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}

}

func (s *VolcengineInstanceStateService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var action string
	targetStatus := []string{"RUNNING"}
	instanceAction := resourceData.Get("action").(string)
	if instanceAction == string(StartAction) {
		action = "StartInstance"
	} else {
		action = "StopInstance"
		targetStatus = []string{"STOPPED"}
	}

	// 根据实例当前状态判断是否执行操作
	update, err := s.describeCurrentStatus(resourceData, targetStatus)
	if err != nil {
		return []ve.Callback{{
			Err: err,
		}}
	}
	if !update {
		resourceData.SetId(fmt.Sprintf("state:%v", resourceData.Get("instance_id")))
		return []ve.Callback{}
	}

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      action,
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"action": {
					Ignore: true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if instanceAction == string(ForceStopAction) {
					(*call.SdkParam)["ForceStop"] = true
				}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				var (
					resp *map[string]interface{}
					err  error
				)
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				if instanceAction == string(StartAction) {
					resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				} else {
					resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				}
				logger.Debug(logger.RespFormat, call.Action, resp)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				instanceId := d.Get("instance_id").(string)
				logger.Debug(logger.RespFormat, call.Action, instanceId)
				d.SetId(fmt.Sprintf("state:%s", instanceId))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  targetStatus,
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}

	return []ve.Callback{callback}
}

func (s *VolcengineInstanceStateService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(condition, "PageSize", "PageNumber", 20, 1, func(m map[string]interface{}) (data []interface{}, err error) {
		action := "DescribeInstances"
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
		logger.Debug(logger.RespFormat, action, resp)

		results, err = ve.ObtainSdkValue("Result.Instances", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.Instances is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineInstanceStateService) ReadResource(resourceData *schema.ResourceData, tmpId string) (data map[string]interface{}, err error) {
	var (
		ok bool
	)
	if tmpId == "" {
		tmpId = resourceData.Id()
	}
	ids := strings.Split(tmpId, ":")
	if len(ids) != 2 {
		return nil, fmt.Errorf("invalid id format. id: %s", tmpId)
	}

	instanceId := ids[1]
	req := map[string]interface{}{
		"InstanceIds.1": instanceId,
	}

	var tempData []interface{}
	if tempData, err = s.ReadResources(req); err != nil {
		return nil, err
	}
	if len(tempData) == 0 {
		return nil, fmt.Errorf("instance %s not exist ", instanceId)
	}
	if data, ok = tempData[0].(map[string]interface{}); !ok {
		return nil, errors.New("Value is not map ")
	}

	if _, ok = resourceData.GetOk("action"); !ok {
		// check status
		status := data["Status"].(string)
		if status == "RUNNING" {
			_ = resourceData.Set("action", "Start")
		} else if status == "STOPPED" {
			_ = resourceData.Set("action", "Stop")
		} else {
			return nil, fmt.Errorf("instance %s status %s is not RUNNING or STOPPED", instanceId, status)
		}
	}
	return data, nil
}

func (s *VolcengineInstanceStateService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				data       map[string]interface{}
				status     interface{}
				failStates []string
			)
			failStates = append(failStates, "ERROR")
			data, err = s.ReadResource(resourceData, id)
			if err != nil {
				return nil, "", err
			}
			status, err = ve.ObtainSdkValue("Status", data)
			logger.Debug(logger.ReqFormat, "DescribeInstances", data)
			logger.Debug(logger.ReqFormat, "DescribeInstances", status)
			logger.Debug(logger.ReqFormat, "DescribeInstances", target)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("Ecs Instance  status  error, status:%s", status.(string))
				}
			}
			return data, status.(string), err
		},
	}
}

func (s *VolcengineInstanceStateService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var action string
	targetStatus := []string{"RUNNING"}
	instanceAction := resourceData.Get("action").(string)
	if instanceAction == string(StartAction) {
		action = "StartInstance"
	} else {
		action = "StopInstance"
		targetStatus = []string{"STOPPED"}
	}

	// 根据实例当前状态判断是否执行操作
	update, err := s.describeCurrentStatus(resourceData, targetStatus)
	if err != nil {
		return []ve.Callback{{
			Err: err,
		}}
	}
	if !update {
		return []ve.Callback{}
	}

	strs := strings.Split(resourceData.Id(), ":")

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      action,
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"action": {
					Ignore: true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["InstanceId"] = strs[1]
				if instanceAction == string(StopAction) || instanceAction == string(ForceStopAction) {
					(*call.SdkParam)["StoppedMode"] = d.Get("stopped_mode")
				}
				if instanceAction == string(ForceStopAction) {
					(*call.SdkParam)["ForceStop"] = true
				}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				if instanceAction == string(StartAction) {
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				} else {
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				}
			},
			Refresh: &ve.StateRefresh{
				Target:  targetStatus,
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineInstanceStateService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineInstanceStateService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineInstanceStateService) ReadResourceId(id string) string {
	return id
}

func (s *VolcengineInstanceStateService) describeCurrentStatus(resourceData *schema.ResourceData, targetStatus []string) (bool, error) {
	instanceId := resourceData.Get("instance_id").(string)
	data, err := s.ReadResource(resourceData, "state:"+instanceId)
	if err != nil {
		return false, err
	}
	status, err := ve.ObtainSdkValue("Status", data)
	if err != nil {
		return false, err
	}
	for _, v := range targetStatus {
		// 目标状态和当前状态相同时，不执行操作
		if v == status.(string) {
			return false, nil
		}
	}
	return true, nil
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "ecs",
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		Action:      actionName,
	}
}
