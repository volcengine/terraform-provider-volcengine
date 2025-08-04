package scaling_group

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

type VolcengineScalingGroupService struct {
	Client *ve.SdkClient
}

func NewScalingGroupService(c *ve.SdkClient) *VolcengineScalingGroupService {
	return &VolcengineScalingGroupService{
		Client: c,
	}
}

func (s *VolcengineScalingGroupService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineScalingGroupService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
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
		logger.Debug(logger.RespFormat, action, action, *resp)
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

func (s *VolcengineScalingGroupService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
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

func (s *VolcengineScalingGroupService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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

func (VolcengineScalingGroupService) WithResourceResponseHandlers(scalingGroup map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return scalingGroup, map[string]ve.ResponseConvert{
			"MultiAZPolicy": {
				TargetField: "multi_az_policy",
			},
			"DBInstanceIds": {
				TargetField: "db_instance_ids",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}

}

func (s *VolcengineScalingGroupService) CreateResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback
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
						if _, ok := data.GetOkExists("desire_instance_number"); !ok {
							return -1
						}
						return i
					},
				},
				"multi_az_policy": {
					TargetField: "MultiAZPolicy",
					ConvertType: ve.ConvertDefault,
				},
				"tags": {
					TargetField: "Tags",
					ConvertType: ve.ConvertListN,
				},
				"db_instance_ids": {
					TargetField: "DBInstanceIds",
					ConvertType: ve.ConvertWithN,
				},
				"launch_template_overrides": {
					ConvertType: ve.ConvertListN,
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
	callbacks = append(callbacks, callback)

	// 伸缩配置来源为 LaunchTemplate, 自动启动伸缩组
	if id, ok := resourceData.GetOk("launch_template_id"); ok && id != "" {
		enableCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "EnableScalingGroup",
				ConvertMode: ve.RequestConvertIgnore,
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["ScalingGroupId"] = d.Id()
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
			},
		}
		callbacks = append(callbacks, enableCallback)
	}

	if v, ok := resourceData.GetOk("wait_for_capacity_timeout"); ok {
		if timeout, _ := time.ParseDuration(v.(string)); timeout > 0 {
			minInstanceNumber := resourceData.Get("min_instance_number")
			desireInstanceNumber := resourceData.Get("desire_instance_number")
			startTime := time.Now().UTC()
			waitCallback := ve.Callback{
				Call: ve.SdkCall{
					Action:      "EnableScalingGroup",
					ConvertMode: ve.RequestConvertIgnore,
					ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
						compare := func(totalInstanceNumber int) error {
							minInstanceNumber := minInstanceNumber.(int)
							if desireInstanceNumber.(int) > 0 {
								minInstanceNumber = desireInstanceNumber.(int)
							}

							if totalInstanceNumber < minInstanceNumber {
								return fmt.Errorf("Expected to have %v instances in the auto scaling group, currently there are %v instances. ", minInstanceNumber, totalInstanceNumber)
							}
							return nil
						}

						if err := s.waitScalingGroupCapacityStateConf(resourceData, timeout, resourceData.Id(), startTime, resourceData.Get("ignore_failed_scaling_activities").(bool), compare); err != nil {
							return nil, err
						}
						return nil, nil
					},
				},
			}
			callbacks = append(callbacks, waitCallback)
		}
	}

	return callbacks
}

func (s *VolcengineScalingGroupService) waitScalingGroupCapacityStateConf(resourceData *schema.ResourceData, timeout time.Duration, groupId string, startTime time.Time, ignoreFailedScalingActivities bool, compare func(int) error) error {
	stateConf := &resource.StateChangeConf{
		Target:  []string{"Satisfied"},
		Timeout: timeout,
		Refresh: s.waitScalingGroupCapacitySatisfied(resourceData, groupId, startTime, ignoreFailedScalingActivities, compare),
	}

	_, err := stateConf.WaitForState()
	if err != nil {
		return err
	}
	return nil
}

func (s *VolcengineScalingGroupService) waitScalingGroupCapacitySatisfied(resourceData *schema.ResourceData, groupId string, startTime time.Time, ignoreFailedScalingActivities bool, compare func(int) error) resource.StateRefreshFunc {
	return func() (result interface{}, state string, err error) {
		var (
			resp       *map[string]interface{}
			results    interface{}
			ok         bool
			activities []interface{}
			instances  []interface{}
		)
		if !ignoreFailedScalingActivities {
			condition := map[string]interface{}{
				"ScalingGroupId": resourceData.Id(),
				"StartTime":      startTime.Format("2006-01-02T15:04Z"),
			}
			activities, err = ve.WithPageNumberQuery(condition, "PageSize", "PageNumber", 50, 1, func(condition map[string]interface{}) ([]interface{}, error) {
				universalClient := s.Client.UniversalClient
				action := "DescribeScalingActivities"
				logger.Debug(logger.ReqFormat, action, condition)
				if condition == nil {
					resp, err = universalClient.DoCall(getUniversalInfo(action), nil)
					if err != nil {
						return activities, err
					}
				} else {
					resp, err = universalClient.DoCall(getUniversalInfo(action), &condition)
					if err != nil {
						return activities, err
					}
				}
				logger.Debug(logger.RespFormat, action, resp)
				results, err = ve.ObtainSdkValue("Result.ScalingActivities", *resp)
				if err != nil {
					return activities, err
				}
				if results == nil {
					results = []interface{}{}
				}
				if activities, ok = results.([]interface{}); !ok {
					return activities, errors.New("Result.ScalingActivities is not Slice")
				}
				return activities, err
			})
			if err != nil {
				return nil, "", err
			}

			var errs []error
			for _, v := range activities {
				activity, ok := v.(map[string]interface{})
				if !ok {
					return nil, "", errors.New("ScalingActivity is not Map")
				}
				if activity["StatusCode"] == "Error" || activity["StatusCode"] == "Exception" {
					errs = append(errs, fmt.Errorf("ScalingActivity %s failed, status: %s, msg: %s", activity["ScalingActivityId"], activity["StatusCode"], activity["ResultMsg"]))
				}
			}
			err = errors.Join(errs...)
			if err != nil {
				return nil, "", err
			}
		}

		condition := map[string]interface{}{
			"ScalingGroupId": resourceData.Id(),
		}
		instances, err = ve.WithPageNumberQuery(condition, "PageSize", "PageNumber", 50, 1, func(condition map[string]interface{}) ([]interface{}, error) {
			universalClient := s.Client.UniversalClient
			action := "DescribeScalingInstances"
			logger.Debug(logger.ReqFormat, action, condition)
			if condition == nil {
				resp, err = universalClient.DoCall(getUniversalInfo(action), nil)
				if err != nil {
					return activities, err
				}
			} else {
				resp, err = universalClient.DoCall(getUniversalInfo(action), &condition)
				if err != nil {
					return activities, err
				}
			}
			logger.Debug(logger.RespFormat, action, condition, *resp)
			results, err = ve.ObtainSdkValue("Result.ScalingInstances", *resp)
			if err != nil {
				return activities, err
			}
			if results == nil {
				results = []interface{}{}
			}
			if instances, ok = results.([]interface{}); !ok {
				return activities, errors.New("Result.ScalingInstances is not Slice")
			}
			return instances, err
		})
		if err != nil {
			return nil, "", err
		}

		totalNumber := 0
		for _, v := range instances {
			instance, ok := v.(map[string]interface{})
			if !ok {
				return nil, "", errors.New("ScalingInstance is not Map")
			}
			if instance["Status"] == "InService" {
				totalNumber += 1
			}
		}

		err = compare(totalNumber)
		if err != nil {
			return []interface{}{}, err.Error(), nil
		}
		return []interface{}{}, "Satisfied", nil
	}
}

func (s *VolcengineScalingGroupService) ModifyResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
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
				"multi_az_policy": {
					TargetField: "MultiAZPolicy",
				},
				"launch_template_id": {
					ConvertType: ve.ConvertDefault,
				},
				"launch_template_version": {
					ConvertType: ve.ConvertDefault,
				},
				"launch_template_overrides": {
					ConvertType: ve.ConvertListN,
				},
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
	// serverGroup modify
	attrAdd, attrRemove, _, _ := ve.GetSetDifference("server_group_attributes", resourceData, serverGroupAttributeHash, false)
	removeAttrCallback := ve.Callback{
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
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						return resource.NonRetryableError(fmt.Errorf("error on reading scaling group %q: %w", d.Id(), callErr))
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
	attachAttrCallback := ve.Callback{
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
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						return resource.NonRetryableError(fmt.Errorf("error on reading scaling group %q, %w", d.Id(), callErr))
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
	callbacks = append(callbacks, removeAttrCallback, attachAttrCallback)
	// 更新Tags
	setResourceTagsCallbacks := ve.SetResourceTags(s.Client, "TagResources", "UntagResources", "scalinggroup", resourceData, getUniversalInfo)
	callbacks = append(callbacks, setResourceTagsCallbacks...)

	if resourceData.HasChanges("desire_instance_number", "min_instance_number") {
		if v, ok := resourceData.GetOk("wait_for_capacity_timeout"); ok {
			if timeout, _ := time.ParseDuration(v.(string)); timeout > 0 {
				minInstanceNumber := resourceData.Get("min_instance_number")
				desireInstanceNumber := resourceData.Get("desire_instance_number")
				startTime := time.Now().UTC()
				waitCallback := ve.Callback{
					Call: ve.SdkCall{
						Action:      "EnableScalingGroup",
						ConvertMode: ve.RequestConvertIgnore,
						ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
							compare := func(totalInstanceNumber int) error {
								minInstanceNumber := minInstanceNumber.(int)
								if desireInstanceNumber.(int) > 0 {
									minInstanceNumber = desireInstanceNumber.(int)
								}

								if totalInstanceNumber != minInstanceNumber {
									return fmt.Errorf("Expected to have %v instances in the auto scaling group, currently there are %v instances. ", minInstanceNumber, totalInstanceNumber)
								}
								return nil
							}

							if err := s.waitScalingGroupCapacityStateConf(resourceData, timeout, resourceData.Id(), startTime, resourceData.Get("ignore_failed_scaling_activities").(bool), compare); err != nil {
								return nil, err
							}
							return nil, nil
						},
					},
				}
				callbacks = append(callbacks, waitCallback)
			}
		}
	}

	return callbacks
}

func (s *VolcengineScalingGroupService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
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
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 5*time.Minute)
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

func (s *VolcengineScalingGroupService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
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
		},
		NameField:    "ScalingGroupName",
		IdField:      "ScalingGroupId",
		CollectField: "scaling_groups",
		ResponseConverts: map[string]ve.ResponseConvert{
			"ScalingGroupId": {
				TargetField: "id",
				KeepDefault: true,
			},
			"MultiAZPolicy": {
				TargetField: "multi_az_policy",
			},
			"DBInstanceIds": {
				TargetField: "db_instance_ids",
			},
		},
	}
}

func (s *VolcengineScalingGroupService) ReadResourceId(id string) string {
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

func (s *VolcengineScalingGroupService) ProjectTrn() *ve.ProjectTrn {
	return &ve.ProjectTrn{
		ServiceName:          "auto_scaling",
		ResourceType:         "scalinggroup",
		ProjectResponseField: "ProjectName",
		ProjectSchemaField:   "project_name",
	}
}
