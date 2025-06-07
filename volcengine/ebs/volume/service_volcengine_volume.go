package volume

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	re "github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineVolumeService struct {
	Client *ve.SdkClient
}

func NewVolumeService(c *ve.SdkClient) *VolcengineVolumeService {
	return &VolcengineVolumeService{
		Client: c,
	}
}

func (s *VolcengineVolumeService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineVolumeService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(condition, "PageSize", "PageNumber", 20, 1, func(m map[string]interface{}) ([]interface{}, error) {
		action := "DescribeVolumes"
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

		results, err = ve.ObtainSdkValue("Result.Volumes", *resp)
		if err != nil {
			return data, err
		}
		logger.Debug(logger.ReqFormat, action, results)
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.Volumes is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineVolumeService) ReadResource(resourceData *schema.ResourceData, volumeId string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if volumeId == "" {
		volumeId = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"VolumeIds.1": volumeId,
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
		return data, fmt.Errorf("volume %s not exist ", volumeId)
	}

	payType, ok := data["PayType"]
	if !ok {
		return data, fmt.Errorf(" PayType of volume is not exist ")
	}
	if payType.(string) == "post" {
		data["VolumeChargeType"] = "PostPaid"
	} else if payType.(string) == "pre" {
		data["VolumeChargeType"] = "PrePaid"
	}

	if extraPerformance, exist := data["ExtraPerformance"]; exist {
		extraPerformanceMap, ok := extraPerformance.(map[string]interface{})
		if !ok {
			return data, fmt.Errorf("The ExtraPerformance of volume is not map ")
		}
		data["ExtraPerformanceTypeId"] = extraPerformanceMap["ExtraPerformanceTypeId"]
		data["ExtraPerformanceIops"] = extraPerformanceMap["IOPS"]
		data["ExtraPerformanceThroughputMb"] = extraPerformanceMap["Throughput"]
	}

	return data, err
}

func (s *VolcengineVolumeService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				ebs        map[string]interface{}
				status     interface{}
				failStates []string
			)
			failStates = append(failStates, "error")

			if err = resource.Retry(20*time.Minute, func() *resource.RetryError {
				ebs, err = s.ReadResource(resourceData, id)
				if err != nil {
					if ve.ResourceNotFoundError(err) {
						return resource.RetryableError(err)
					} else {
						return resource.NonRetryableError(err)
					}
				}
				return nil
			}); err != nil {
				return nil, "", err
			}

			status, err = ve.ObtainSdkValue("Status", ebs)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("volume status error, status:%s", status.(string))
				}
			}
			return ebs, status.(string), err
		},
	}
}

func (VolcengineVolumeService) WithResourceResponseHandlers(volume map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return volume, map[string]ve.ResponseConvert{
			"Size": {
				TargetField: "size",
				Convert:     sizeConvertFunc,
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineVolumeService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateVolume",
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"tags": {
					TargetField: "Tags",
					ConvertType: ve.ConvertListN,
				},
				"extra_performance_iops": {
					TargetField: "ExtraPerformanceIOPS",
				},
				"extra_performance_throughput_mb": {
					TargetField: "ExtraPerformanceThroughputMB",
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("Result.VolumeId", *resp)
				d.SetId(id.(string))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"available", "attached"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	callbacks = append(callbacks, callback)

	if resourceData.Get("delete_with_instance").(bool) {
		callbacks = append(callbacks, ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyVolumeAttribute",
				ConvertMode: ve.RequestConvertInConvert,
				Convert: map[string]ve.RequestConvert{
					"delete_with_instance": {
						TargetField: "DeleteWithInstance",
					},
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["VolumeId"] = d.Id()
					delete(*call.SdkParam, "Tags")
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
					logger.Debug(logger.RespFormat, call.Action, resp, err)
					return resp, err
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{"available", "attached"},
					Timeout: resourceData.Timeout(schema.TimeoutUpdate),
				},
			},
		})
	}

	return callbacks
}

func (s *VolcengineVolumeService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback
	if resourceData.HasChanges("volume_name", "description", "delete_with_instance") {
		callbacks = append(callbacks, ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyVolumeAttribute",
				ConvertMode: ve.RequestConvertInConvert,
				Convert: map[string]ve.RequestConvert{
					"volume_name": {
						TargetField: "VolumeName",
						ForceGet:    true,
					},
					"description": {
						TargetField: "Description",
					},
					"delete_with_instance": {
						TargetField: "DeleteWithInstance",
					},
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["VolumeId"] = d.Id()
					delete(*call.SdkParam, "Tags")
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{"available", "attached"},
					Timeout: resourceData.Timeout(schema.TimeoutUpdate),
				},
			},
		})
	}

	if resourceData.HasChange("size") { // 调用新的 api
		callbacks = append(callbacks, ve.Callback{
			Call: ve.SdkCall{
				Action:      "ExtendVolume",
				ConvertMode: ve.RequestConvertIgnore,
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["VolumeId"] = d.Id()
					(*call.SdkParam)["NewSize"] = d.Get("size")
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{"available", "attached"},
					Timeout: resourceData.Timeout(schema.TimeoutUpdate),
				},
			},
		})
	}

	if resourceData.HasChange("volume_charge_type") {
		callbacks = append(callbacks, ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyVolumeChargeType",
				ConvertMode: ve.RequestConvertIgnore,
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					if d.Get("instance_id").(string) == "" {
						return false, errors.New("instance id cannot be empty")
					}

					chargeType := resourceData.Get("volume_charge_type")
					(*call.SdkParam)["VolumeIds.1"] = d.Id()
					(*call.SdkParam)["DiskChargeType"] = chargeType
					(*call.SdkParam)["InstanceId"] = d.Get("instance_id")
					if chargeType == "PrePaid" {
						(*call.SdkParam)["AutoPay"] = true
					}
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
					logger.Debug(logger.RespFormat, call.Action, resp)
					logger.Debug(logger.RespFormat, call.Action, err)
					return resp, err
				},
				CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
					chargeType := resourceData.Get("volume_charge_type")
					if d.Get("instance_id").(string) == "" {
						return errors.New("instance id cannot be empty")
					}
					// retry modifyVolumeChargeType
					return re.Retry(15*time.Minute, func() *re.RetryError {
						data, callErr := s.ReadResource(d, d.Id())
						if callErr != nil {
							return re.NonRetryableError(fmt.Errorf("error on reading volume %q: %w", d.Id(), callErr))
						}
						// 计费方式已经转变成功
						if (chargeType == "PrePaid" && data["PayType"] == "pre") || (chargeType == "PostPaid" && data["PayType"] == "post") {
							return nil
						}
						// 计费方式还没有转换成功，尝试重新转换
						_, callErr = call.ExecuteCall(d, client, call)
						if callErr == nil {
							return nil
						}
						// 按量实例下挂载的云盘不支持按量转包年操作
						if strings.Contains(callErr.Error(), "ErrorInvalidEcsChargeType") {
							return re.NonRetryableError(callErr)
						}
						return re.RetryableError(callErr)
					})
				},
			},
		})
	}

	if resourceData.HasChange("volume_type") {
		callbacks = append(callbacks, ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyVolumeSpec",
				ConvertMode: ve.RequestConvertIgnore,
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["VolumeId"] = d.Id()
					(*call.SdkParam)["TargetVolumeType"] = d.Get("volume_type")
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
					logger.Debug(logger.RespFormat, call.Action, resp)
					return resp, err
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{"available", "attached"},
					Timeout: resourceData.Timeout(schema.TimeoutUpdate),
				},
			},
		})
	}

	if resourceData.HasChanges("extra_performance_type_id", "extra_performance_iops", "extra_performance_throughput_mb") {
		callbacks = append(callbacks, ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyVolumeExtraPerformance",
				ConvertMode: ve.RequestConvertInConvert,
				Convert: map[string]ve.RequestConvert{
					"extra_performance_iops": {
						TargetField: "ExtraPerformanceIOPS",
					},
					"extra_performance_throughput_mb": {
						TargetField: "ExtraPerformanceThroughputMB",
					},
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["VolumeId"] = d.Id()
					(*call.SdkParam)["ExtraPerformanceTypeId"] = d.Get("extra_performance_type_id")
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
					logger.Debug(logger.RespFormat, call.Action, resp)
					return resp, err
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{"available", "attached"},
					Timeout: resourceData.Timeout(schema.TimeoutUpdate),
				},
			},
		})
	}

	// 更新Tags
	setResourceTagsCallbacks := ve.SetResourceTags(s.Client, "CreateTags", "DeleteTags", "volume", resourceData, getUniversalInfo)
	callbacks = append(callbacks, setResourceTagsCallbacks...)

	return callbacks
}

func (s *VolcengineVolumeService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteVolume",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"VolumeId": resourceData.Id(),
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				volume, err := s.ReadResource(d, d.Id())
				if err != nil {
					return false, err
				}

				// 包年包月云盘和随实例删除云盘，直接移除管理
				chargeType, err := ve.ObtainSdkValue("VolumeChargeType", volume)
				if err != nil {
					return false, err
				}
				deleteWithInstance, err := ve.ObtainSdkValue("DeleteWithInstance", volume)
				if err != nil {
					return false, err
				}
				if chargeType == "PrePaid" && deleteWithInstance.(bool) {
					logger.DebugInfo("The Resource volcengine_volume %s ChargeType is PrePaid and its attribute DeleteWithInstance is true, so it will remove from state.", d.Id())
					return false, nil
				}

				status, err := ve.ObtainSdkValue("Status", volume)
				if err != nil {
					return false, err
				}
				if status != "available" {
					return false, fmt.Errorf(" Only volume with a status of `available` can be deleted. ")
				}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				// 不能删除已挂载云盘
				if strings.Contains(baseErr.Error(), "Only volume with a status of `available` can be deleted.") {
					msg := fmt.Sprintf("error: %s\n msg: %s",
						baseErr.Error(),
						"For volume with a status of `attached`, please use `terraform state rm volcengine_volume.resource_name` command to remove it from terraform state file and management.")
					return fmt.Errorf(msg)
				}
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on  reading vpc on delete %q, %w", d.Id(), callErr))
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

func (s *VolcengineVolumeService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "VolumeIds",
				ConvertType: ve.ConvertWithN,
			},
			"tags": {
				TargetField: "TagFilters",
				ConvertType: ve.ConvertListN,
				NextLevelConvert: map[string]ve.RequestConvert{
					"value": {
						TargetField: "Values.1",
					},
				},
			},
		},
		NameField:    "VolumeName",
		IdField:      "VolumeId",
		CollectField: "volumes",
		ResponseConverts: map[string]ve.ResponseConvert{
			"VolumeId": {
				TargetField: "id",
				KeepDefault: true,
			},
			"Size": {
				TargetField: "size",
				Convert:     sizeConvertFunc,
			},
			"IOPS": {
				TargetField: "iops",
			},
		},
	}
}

var sizeConvertFunc = func(i interface{}) interface{} {
	// Notice: the type of filed Size in openapi doc is size, but api return type is string
	size, ok := i.(string)
	if !ok {
		return i
	}
	res, err := strconv.Atoi(size)
	if err != nil {
		logger.Debug(logger.ReqFormat, "sizeConvertFunc", i)
		return i
	}
	return res
}

func (s *VolcengineVolumeService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "storage_ebs",
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		Action:      actionName,
	}
}

func (s *VolcengineVolumeService) ProjectTrn() *ve.ProjectTrn {
	return &ve.ProjectTrn{
		ServiceName:          "storage_ebs",
		ResourceType:         "volume",
		ProjectResponseField: "ProjectName",
		ProjectSchemaField:   "project_name",
	}
}
