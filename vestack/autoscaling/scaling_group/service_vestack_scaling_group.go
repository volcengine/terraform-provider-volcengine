package scaling_group

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-vestack/common"
	"github.com/volcengine/terraform-provider-vestack/logger"
)

type VestackScalingGroupService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewScalingGroupService(c *ve.SdkClient) *VestackScalingGroupService {
	return &VestackScalingGroupService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VestackScalingGroupService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VestackScalingGroupService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
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

func (s *VestackScalingGroupService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
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

func (s *VestackScalingGroupService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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

func (VestackScalingGroupService) WithResourceResponseHandlers(scalingGroup map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return scalingGroup, map[string]ve.ResponseConvert{
			"MultiAZPolicy": {
				TargetField: "multi_az_policy",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}

}

func (s *VestackScalingGroupService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateScalingGroup",
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"subnet_ids": {
					ConvertType: ve.ConvertWithN,
				},
				"server_group_attributes": {
					ConvertType: ve.ConvertListN,
				},
				"db_instance_ids": {
					TargetField: "DBInstanceIds",
					ConvertType: ve.ConvertWithN,
				},
				"min_instance_number": {
					TargetField: "MinInstanceNumber",
					// 如果为0时，需要这样转一下，要不然不会出现在请求参数
					Convert: func(data *schema.ResourceData, i interface{}) interface{} {
						return i
					},
				},
				"max_instance_number": {
					TargetField: "MaxInstanceNumber",
					// 如果为0时，需要这样转一下，要不然不会出现在请求参数
					Convert: func(data *schema.ResourceData, i interface{}) interface{} {
						return i
					},
				},
				"desire_instance_number": {
					TargetField: "DesireInstanceNumber",
					// 如果为0时，需要这样转一下，要不然不会出现在请求参数
					Convert: func(data *schema.ResourceData, i interface{}) interface{} {
						return i
					},
				},
				"multi_az_policy": {
					TargetField: "MultiAZPolicy",
					ConvertType: ve.ConvertDefault,
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				//注意 获取内容 这个地方不能是指针 需要转一次
				id, _ := ve.ObtainSdkValue("Result.ScalingGroupId", *resp)
				d.SetId(id.(string))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"InActive"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	return []ve.Callback{callback}

}

func (s *VestackScalingGroupService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	// 修改伸缩组
	modifyGroupCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:         "ModifyScalingGroup",
			ConvertMode:    ve.RequestConvertInConvert,
			RequestIdField: "ScalingGroupId",
			Convert: map[string]ve.RequestConvert{
				"scaling_group_name": {
					ConvertType: ve.ConvertDefault,
				},
				"min_instance_number": {
					ConvertType: ve.ConvertDefault,
					Convert: func(data *schema.ResourceData, i interface{}) interface{} {
						return i
					},
				},
				"max_instance_number": {
					ConvertType: ve.ConvertDefault,
					Convert: func(data *schema.ResourceData, i interface{}) interface{} {
						return i
					},
				},
				"subnet_ids": {
					ConvertType: ve.ConvertWithN,
				},
				"desire_instance_number": {
					ConvertType: ve.ConvertDefault,
					Convert: func(data *schema.ResourceData, i interface{}) interface{} {
						return i
					},
				},
				"instance_terminate_policy": {
					ConvertType: ve.ConvertDefault,
				},
				"default_cooldown": {
					ConvertType: ve.ConvertDefault,
				},
				// 暂不支持修改
				//"multi_az_policy": {
				//	TargetField: "MultiAZPolicy",
				//},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if len(*call.SdkParam) < 2 {
					return false, nil
				}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	callbacks = append(callbacks, modifyGroupCallback)

	// 修改RDS属性
	rdsAdd, rdsRemove, _, _ := ve.GetSetDifference("db_instance_ids", resourceData, schema.HashString, false)
	removeDBInstanceCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DetachDBInstances",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if rdsRemove != nil && len(rdsRemove.List()) > 0 {
					(*call.SdkParam)["ScalingGroupId"] = d.Id()
					(*call.SdkParam)["ForceDetach"] = true
					for index, dbId := range rdsRemove.List() {
						(*call.SdkParam)["DBInstanceIds."+strconv.Itoa(index+1)] = dbId.(string)
					}
					return true, nil
				}
				return false, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	addDBInstanceCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "AttachDBInstances",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if rdsAdd != nil && len(rdsAdd.List()) > 0 {
					(*call.SdkParam)["ScalingGroupId"] = d.Id()
					(*call.SdkParam)["ForceAttach"] = true
					for index, dbId := range rdsAdd.List() {
						(*call.SdkParam)["DBInstanceIds."+strconv.Itoa(index+1)] = dbId.(string)
					}
					return true, nil
				}
				return false, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	callbacks = append(callbacks, removeDBInstanceCallback, addDBInstanceCallback)

	return callbacks
}

func (s *VestackScalingGroupService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteScalingGroup",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"ScalingGroupId": resourceData.Id(),
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
							return resource.NonRetryableError(fmt.Errorf("error on reading ScalingGroup on delete %q, %w", d.Id(), callErr))
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

func (s *VestackScalingGroupService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "ScalingGroupIds",
				ConvertType: ve.ConvertWithN,
			},
			"scaling_group_names": {
				TargetField: "ScalingGroupNames",
				ConvertType: ve.ConvertWithN,
			},
			"multi_az_policy": {
				TargetField: "MultiAZPolicy",
			},
		},
		NameField:    "ScalingGroupName",
		IdField:      "ScalingGroupId",
		CollectField: "scaling_groups",
		ResponseConverts: map[string]ve.ResponseConvert{
			"ScalingGroupId": {
				TargetField: "id",
				KeepDefault: true,
			},
			"DBInstanceIds": {
				TargetField: "db_instance_ids",
			},
			"MultiAZPolicy": {
				TargetField: "multi_az_policy",
			},
		},
	}
}

func (s *VestackScalingGroupService) ReadResourceId(id string) string {
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
