package ecs_invocation

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

type VolcengineEcsInvocationService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewEcsInvocationService(c *ve.SdkClient) *VolcengineEcsInvocationService {
	return &VolcengineEcsInvocationService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineEcsInvocationService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineEcsInvocationService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 20, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeInvocations"
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
		results, err = ve.ObtainSdkValue("Result.Invocations", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.Invocations is not Slice")
		}

		for _, v := range data {
			instanceIds := make([]string, 0)
			invocation, ok := v.(map[string]interface{})
			if !ok {
				return data, fmt.Errorf(" Invocation is not map ")
			}
			action := "DescribeInvocationInstances"
			req := map[string]interface{}{
				"InvocationId": invocation["InvocationId"],
			}
			logger.Debug(logger.ReqFormat, action, req)
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &req)
			logger.Debug(logger.RespFormat, action, req, resp)
			results, err := ve.ObtainSdkValue("Result.InvocationInstances", *resp)
			if err != nil {
				return data, err
			}
			if results == nil {
				results = []interface{}{}
			}
			instances, ok := results.([]interface{})
			if !ok {
				return data, errors.New("Result.InvocationInstances is not Slice")
			}
			if len(instances) == 0 {
				return data, fmt.Errorf("invocation %s does not contain any instances", invocation["InvocationId"])
			}
			for _, v1 := range instances {
				instance, ok := v1.(map[string]interface{})
				if !ok {
					return data, fmt.Errorf(" invocation instance is not map ")
				}
				instanceIds = append(instanceIds, instance["InstanceId"].(string))
			}
			invocation["InstanceIds"] = instanceIds
		}

		return data, err
	})
}

func (s *VolcengineEcsInvocationService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"InvocationId": id,
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
		return data, fmt.Errorf("ecs invocation %s is not exist ", id)
	}

	// 处理 launch_time、recurrence_end_time 传参与查询结果不一致的问题
	if mode := resourceData.Get("repeat_mode"); mode.(string) == "Rate" {
		layout := "2006-01-02T15:04:05Z"
		launchTimeExpr, exist1 := resourceData.GetOkExists("launch_time")
		endTimeExpr, exist2 := resourceData.GetOkExists("recurrence_end_time")
		if !exist1 || !exist2 {
			return data, fmt.Errorf(" launch_time or recurrence_end_time is not exist ")
		}
		launchTime, err := ParseUTCTime(launchTimeExpr.(string))
		if err != nil {
			return data, err
		}
		endTime, err := ParseUTCTime(endTimeExpr.(string))
		if err != nil {
			return data, err
		}
		lt := launchTime.Format(layout)
		et := endTime.Format(layout)
		if lt == data["LaunchTime"].(string) {
			data["LaunchTime"] = launchTimeExpr
		}
		if et == data["RecurrenceEndTime"].(string) {
			data["RecurrenceEndTime"] = endTimeExpr
		}
	}

	return data, err
}

func ParseUTCTime(timeExpr string) (time.Time, error) {
	timeWithoutSecond, err := ParseUTCTimeWithoutSecond(timeExpr)
	if err != nil {
		timeWithSecond, err := ParseUTCTimeWithSecond(timeExpr)
		if err != nil {
			return time.Time{}, err
		} else {
			return timeWithSecond, nil
		}
	} else {
		return timeWithoutSecond, nil
	}
}

func ParseUTCTimeWithoutSecond(timeExpr string) (time.Time, error) {
	t, err := time.Parse("2006-01-02T15:04Z", timeExpr)
	if err != nil {
		return time.Time{}, fmt.Errorf("parse time failed, error: %v, time expr: %v", err, timeExpr)
	}

	return t, nil
}

func ParseUTCTimeWithSecond(timeExpr string) (time.Time, error) {
	t, err := time.Parse("2006-01-02T15:04:05Z", timeExpr)
	if err != nil {
		return time.Time{}, fmt.Errorf("parse time failed, error: %v, time expr: %v", err, timeExpr)
	}

	t = t.Add(time.Duration(t.Second()) * time.Second * -1)

	return t, nil
}

func (s *VolcengineEcsInvocationService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				demo   map[string]interface{}
				status interface{}
			)
			//no failed status.
			demo, err = s.ReadResource(resourceData, id)
			if err != nil {
				return nil, "", err
			}
			status, err = ve.ObtainSdkValue("InvocationStatus", demo)
			if err != nil {
				return nil, "", err
			}
			return demo, status.(string), err
		},
	}
}

func (VolcengineEcsInvocationService) WithResourceResponseHandlers(invocation map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return invocation, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineEcsInvocationService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "InvokeCommand",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeDefault,
			Convert: map[string]ve.RequestConvert{
				"instance_ids": {
					TargetField: "InstanceIds",
					ConvertType: ve.ConvertWithN,
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("Result.InvocationId", *resp)
				d.SetId(id.(string))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineEcsInvocationService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineEcsInvocationService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "StopInvocation",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeDefault,
			SdkParam: &map[string]interface{}{
				"InvocationId": resourceData.Id(),
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				status := d.Get("invocation_status")
				mode := d.Get("repeat_mode")
				if mode.(string) == "Once" || (status.(string) != "Pending" && status.(string) != "Scheduled") {
					return false, nil
				} else {
					(*call.SdkParam)["InvocationId"] = d.Id()
					return true, nil
				}
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Stopped"},
				Timeout: resourceData.Timeout(schema.TimeoutDelete),
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineEcsInvocationService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		NameField:    "InvocationName",
		IdField:      "InvocationId",
		CollectField: "invocations",
		ResponseConverts: map[string]ve.ResponseConvert{
			"InvocationId": {
				TargetField: "id",
				KeepDefault: true,
			},
		},
	}
}

func (s *VolcengineEcsInvocationService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "ecs",
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		Action:      actionName,
	}
}
