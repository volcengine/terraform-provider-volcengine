package ecs_instance

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-vestack/common"
	"github.com/volcengine/terraform-provider-vestack/logger"
)

type VestackEcsService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewEcsService(c *ve.SdkClient) *VestackEcsService {
	return &VestackEcsService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VestackEcsService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VestackEcsService) ReadResources(condition map[string]interface{}) ([]interface{}, error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithNextTokenQuery(condition, "MaxResults", "NextToken", 20, nil, func(m map[string]interface{}) (data []interface{}, next string, err error) {
		ecs := s.Client.EcsClient
		action := "DescribeInstances"
		logger.Debug(logger.ReqFormat, action, condition)
		if condition == nil {
			resp, err = ecs.DescribeInstancesCommon(nil)
			if err != nil {
				return data, next, err
			}
		} else {
			resp, err = ecs.DescribeInstancesCommon(&condition)
			if err != nil {
				return data, next, err
			}
		}
		logger.Debug(logger.RespFormat, action, condition, *resp)

		results, err = ve.ObtainSdkValue("Result.Instances", *resp)
		if err != nil {
			return data, next, err
		}
		nextToken, err := ve.ObtainSdkValue("Result.NextToken", *resp)
		if err != nil {
			return data, next, err
		}
		next = nextToken.(string)
		if results == nil {
			results = []interface{}{}
		}

		if data, ok = results.([]interface{}); !ok {
			return data, next, errors.New("Result.Instances is not Slice")
		}
		return data, next, err
	})
}

func (s *VestackEcsService) ReadResource(resourceData *schema.ResourceData, instanceId string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if instanceId == "" {
		instanceId = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"InstanceIds.1": instanceId,
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		if data, ok = v.(map[string]interface{}); !ok {
			return data, fmt.Errorf("Value is not map ")
		}
	}

	if len(data) == 0 {
		return data, fmt.Errorf("Ecs Instance %s not exist ", instanceId)
	}
	return data, nil
}

func (s *VestackEcsService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      5 * time.Second,
		MinTimeout: 5 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				data       map[string]interface{}
				status     interface{}
				failStates []string
			)
			failStates = append(failStates, "ERROR")

			if err = resource.Retry(20*time.Minute, func() *resource.RetryError {
				data, err = s.ReadResource(resourceData, id)
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

			status, err = ve.ObtainSdkValue("Status", data)
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

func (s *VestackEcsService) WithResourceResponseHandlers(ecs map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		var (
			ebsErr              error
			userDataErr         error
			networkInterfaceErr error
			errorStr            string
			wg                  sync.WaitGroup
			syncMap             sync.Map
		)

		//计算period
		if ecs["InstanceChargeType"].(string) == "PrePaid" {
			ct, _ := time.Parse("2006-01-02T15:04:05", ecs["CreatedAt"].(string)[0:strings.Index(ecs["CreatedAt"].(string), "+")])
			et, _ := time.Parse("2006-01-02T15:04:05", ecs["ExpiredAt"].(string)[0:strings.Index(ecs["ExpiredAt"].(string), "+")])
			y := et.Year() - ct.Year()
			m := et.Month() - ct.Month()
			ecs["Period"] = y*12 + int(m)
		}

		wg.Add(3)
		instanceId := ecs["InstanceId"]
		//read ebs data
		go func() {
			defer func() {
				if _err := recover(); _err != nil {
					logger.Debug(logger.ReqFormat, "DescribeVolumes", _err)
				}
				wg.Done()
			}()
			temp := map[string]interface{}{
				"InstanceId": ecs["InstanceId"],
			}
			_, ebsErr = s.readEbsVolumes([]interface{}{temp})
			if ebsErr != nil {
				return
			}
			syncMap.Store("Volumes", temp["Volumes"])
		}()
		//read user_data
		go func() {
			defer func() {
				if _err := recover(); _err != nil {
					logger.Debug(logger.ReqFormat, "DescribeUserData", _err)
				}
				ve.Release()
				wg.Done()
			}()
			ve.Acquire()
			var (
				userDataParam *map[string]interface{}
				userDataResp  *map[string]interface{}
				userData      interface{}
			)
			userDataParam = &map[string]interface{}{
				"InstanceId": instanceId,
			}
			userDataResp, userDataErr = s.Client.EcsClient.DescribeUserDataCommon(userDataParam)
			if userDataErr != nil {
				return
			}
			userData, userDataErr = ve.ObtainSdkValue("Result.UserData", *userDataResp)
			if userDataErr != nil {
				return
			}
			syncMap.Store("UserData", userData)
		}()
		//read network_interfaces
		go func() {
			defer func() {
				if _err := recover(); _err != nil {
					logger.Debug(logger.ReqFormat, "DescribeNetworkInterfaces", _err)
				}
				ve.Release()
				wg.Done()
			}()
			ve.Acquire()
			var (
				networkInterfaceParam *map[string]interface{}
				networkInterfaceResp  *map[string]interface{}
				networkInterface      interface{}
			)
			networkInterfaceParam = &map[string]interface{}{
				"InstanceId": instanceId,
			}
			networkInterfaceResp, networkInterfaceErr = s.Client.VpcClient.DescribeNetworkInterfacesCommon(networkInterfaceParam)
			if networkInterfaceErr != nil {
				return
			}
			networkInterface, networkInterfaceErr = ve.ObtainSdkValue("Result.NetworkInterfaceSets", *networkInterfaceResp)
			if networkInterfaceErr != nil {
				return
			}
			syncMap.Store("NetworkInterfaces", networkInterface)
		}()
		wg.Wait()
		//error processed
		if ebsErr != nil {
			errorStr = errorStr + ebsErr.Error() + ";"
		}
		if userDataErr != nil {
			errorStr = errorStr + userDataErr.Error() + ";"
		}
		if networkInterfaceErr != nil {
			errorStr = errorStr + networkInterfaceErr.Error() + ";"
		}
		if len(errorStr) > 0 {
			return ecs, s.CommonResponseConvert(), fmt.Errorf(errorStr)
		}
		//clean something
		delete(ecs, "Volumes")
		delete(ecs, "UserData")
		delete(ecs, "NetworkInterfaces")
		//merge extra data
		syncMap.Range(func(key, value interface{}) bool {
			ecs[key.(string)] = value
			return true
		})
		//split primary vif and secondary vif
		if networkInterfaces, ok1 := ecs["NetworkInterfaces"].([]interface{}); ok1 {
			var dataNetworkInterfaces []interface{}
			for _, vif := range networkInterfaces {
				if v1, ok2 := vif.(map[string]interface{}); ok2 {
					if v1["Type"] == "primary" {
						ecs["SubnetId"] = v1["SubnetId"]
						ecs["SecurityGroupIds"] = v1["SecurityGroupIds"]
						ecs["NetworkInterfaceId"] = v1["NetworkInterfaceId"]
					} else {
						dataNetworkInterfaces = append(dataNetworkInterfaces, vif)
					}
				}
			}
			if len(dataNetworkInterfaces) > 0 {
				ecs["SecondaryNetworkInterfaces"] = dataNetworkInterfaces
			}
		}

		//split System volume and Data volume
		if volumes, ok1 := ecs["Volumes"].([]interface{}); ok1 {
			var dataVolumes []interface{}
			for _, volume := range volumes {
				if v1, ok2 := volume.(map[string]interface{}); ok2 {
					if v1["Kind"] == "system" {
						ecs["SystemVolumeType"] = v1["VolumeType"]
						ecs["SystemVolumeSize"] = v1["Size"]
						ecs["SystemVolumeId"] = v1["VolumeId"]
					} else {
						dataVolumes = append(dataVolumes, volume)
					}
				}
			}
			if len(dataVolumes) > 0 {
				ecs["DataVolumes"] = dataVolumes
			}
		}
		return ecs, s.CommonResponseConvert(), nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VestackEcsService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "RunInstances",
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"system_volume_type": {
					ConvertType: ve.ConvertDefault,
					TargetField: "Volumes.1.VolumeType",
				},
				"system_volume_size": {
					ConvertType: ve.ConvertDefault,
					TargetField: "Volumes.1.Size",
				},
				"subnet_id": {
					ConvertType: ve.ConvertDefault,
					TargetField: "NetworkInterfaces.1.SubnetId",
				},
				"security_group_ids": {
					ConvertType: ve.ConvertWithN,
					TargetField: "NetworkInterfaces.1.SecurityGroupIds",
				},
				"data_volumes": {
					ConvertType: ve.ConvertListN,
					TargetField: "Volumes",
					StartIndex:  1,
				},
				"secondary_network_interfaces": {
					ConvertType: ve.ConvertListN,
					TargetField: "NetworkInterfaces",
					NextLevelConvert: map[string]ve.RequestConvert{
						"security_group_ids": {
							ConvertType: ve.ConvertWithN,
						},
					},
					StartIndex: 1,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["ClientToken"] = uuid.New().String()
				(*call.SdkParam)["Volumes.1.DeleteWithInstance"] = true
				(*call.SdkParam)["Count"] = 1
				if (*call.SdkParam)["InstanceChargeType"] == "PrePaid" {
					if (*call.SdkParam)["Period"] == nil || (*call.SdkParam)["Period"].(int) < 1 {
						return false, fmt.Errorf("Instance Charge Type is PrePaid.Must set Period more than 1. ")
					}
					(*call.SdkParam)["PeriodUnit"] = "Month"
				}
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				//创建ECS
				return s.Client.EcsClient.RunInstancesCommon(call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				//注意 获取内容 这个地方不能是指针 需要转一次
				id, _ := ve.ObtainSdkValue("Result.InstanceIds.0", *resp)
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam, id)
				d.SetId(id.(string))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"RUNNING"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VestackEcsService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) (callbacks []ve.Callback) {
	var (
		passwordChange bool
	)
	modifyInstanceSpec := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyInstanceSpec",
			ConvertMode: ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"instance_type": {
					ConvertType: ve.ConvertDefault,
				},
			},
			RequestIdField: "InstanceId",
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if len(*call.SdkParam) > 1 {
					return true, nil
				}
				return false, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				(*call.SdkParam)["ClientToken"] = uuid.New().String()
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				//修改实例规格
				return s.Client.EcsClient.ModifyInstanceSpecCommon(call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"RUNNING", "STOPPED"},
				Timeout: resourceData.Timeout(schema.TimeoutUpdate),
			},
		},
	}
	callbacks = append(callbacks, modifyInstanceSpec)
	modifyInstanceAttribute := ve.Callback{
		Call: ve.SdkCall{
			Action:         "ModifyInstanceAttribute",
			ConvertMode:    ve.RequestConvertInConvert,
			RequestIdField: "InstanceId",
			Convert: map[string]ve.RequestConvert{
				"password": {
					ConvertType: ve.ConvertDefault,
				},
				"user_data": {
					ConvertType: ve.ConvertDefault,
				},
				"instance_name": {
					ConvertType: ve.ConvertDefault,
				},
				"description": {
					ConvertType: ve.ConvertDefault,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				//if image changed ,password change in replaceSystemVolume,not here
				if _, ok := (*call.SdkParam)["Password"]; ok && d.HasChange("image_id") {
					delete(*call.SdkParam, "Password")
				}
				if len(*call.SdkParam) > 1 {
					return true, nil
				}
				return false, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				(*call.SdkParam)["ClientToken"] = uuid.New().String()
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				//修改实例属性
				return s.Client.EcsClient.ModifyInstanceAttributeCommon(call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				if _, ok := (*call.SdkParam)["Password"]; ok {
					passwordChange = true
				}
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"RUNNING", "STOPPED"},
				Timeout: resourceData.Timeout(schema.TimeoutUpdate),
			},
		},
	}
	callbacks = append(callbacks, modifyInstanceAttribute)

	modifyInstanceChargeType := ve.Callback{
		Call: ve.SdkCall{
			Action:         "ModifyInstanceChargeType",
			ConvertMode:    ve.RequestConvertInConvert,
			RequestIdField: "InstanceIds.1",
			Convert: map[string]ve.RequestConvert{
				"instance_charge_type": {
					ConvertType: ve.ConvertDefault,
				},
				"include_data_volumes": {
					ConvertType: ve.ConvertDefault,
					ForceGet:    true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if len(*call.SdkParam) > 2 {
					(*call.SdkParam)["AutoPay"] = true
					if (*call.SdkParam)["InstanceChargeType"].(string) == "PostPaid" {
						//后付费
						return true, nil
					} else {
						//预付费
						period := d.Get("period")
						if period.(int) <= 0 {
							return false, fmt.Errorf("period must set and more than 0 ")
						}
						(*call.SdkParam)["Period"] = period
						//(*call.SdkParam)["PeriodUnit"] = d.Get("period_unit")
						(*call.SdkParam)["PeriodUnit"] = "Month"
						return true, nil
					}
				}
				return false, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				(*call.SdkParam)["ClientToken"] = uuid.New().String()
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				//修改实例计费方式
				return s.Client.EcsClient.ModifyInstanceChargeTypeCommon(call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"RUNNING", "STOPPED"},
				Timeout: resourceData.Timeout(schema.TimeoutUpdate),
			},
		},
	}

	callbacks = append(callbacks, modifyInstanceChargeType)

	//primary vif sg change
	if resourceData.HasChange("security_group_ids") {
		modifyNetworkInterfaceAttributes := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyNetworkInterfaceAttributes",
				ConvertMode: ve.RequestConvertInConvert,
				Convert: map[string]ve.RequestConvert{
					"network_interface_id": {
						ConvertType: ve.ConvertDefault,
						ForceGet:    true,
					},
					"security_group_ids": {
						ConvertType: ve.ConvertWithN,
						ForceGet:    true,
					},
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
					return s.Client.VpcClient.ModifyNetworkInterfaceAttributesCommon(call.SdkParam)
				},
				AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
					return nil
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{"RUNNING", "STOPPED"},
					Timeout: resourceData.Timeout(schema.TimeoutUpdate),
				},
			},
		}
		callbacks = append(callbacks, modifyNetworkInterfaceAttributes)
	}
	//system_volume change
	extendVolume := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ExtendVolume",
			ConvertMode: ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"system_volume_id": {
					ConvertType: ve.ConvertDefault,
					TargetField: "VolumeId",
					ForceGet:    true,
				},
				"system_volume_size": {
					ConvertType: ve.ConvertDefault,
					TargetField: "NewSize",
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if len(*call.SdkParam) == 2 {
					o, n := d.GetChange("system_volume_size")
					if o.(int) > n.(int) {
						return false, fmt.Errorf("SystemVolumeSize only support extend. ")
					}
					return true, nil
				}
				return false, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.EbsClient.ExtendVolumeCommon(call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"RUNNING", "STOPPED"},
				Timeout: resourceData.Timeout(schema.TimeoutUpdate),
			},
		},
	}
	callbacks = append(callbacks, extendVolume)
	//image change
	if resourceData.HasChange("image_id") {
		//先需要关机才能重装系统
		stopInstance := s.StartOrStopInstanceCallback(resourceData, true)
		callbacks = append(callbacks, stopInstance)
		replaceSystemVolume := ve.Callback{
			Call: ve.SdkCall{
				Action:         "ReplaceSystemVolume",
				ConvertMode:    ve.RequestConvertInConvert,
				RequestIdField: "InstanceId",
				Convert: map[string]ve.RequestConvert{
					"image_id": {
						ConvertType: ve.ConvertDefault,
					},
					"system_volume_size": {
						ConvertType: ve.ConvertDefault,
						ForceGet:    true,
					},
					"key_pair_name": {
						ConvertType: ve.ConvertDefault,
						ForceGet:    true,
					},
					"password": {
						ConvertType: ve.ConvertDefault,
						ForceGet:    true,
					},
					"user_data": {
						ConvertType: ve.ConvertDefault,
						ForceGet:    true,
					},
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
					return s.Client.EcsClient.ReplaceSystemVolumeCommon(call.SdkParam)
				},
				AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
					return nil
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{"RUNNING"},
					Timeout: resourceData.Timeout(schema.TimeoutUpdate),
				},
			},
		}
		callbacks = append(callbacks, replaceSystemVolume)
	}

	if !resourceData.HasChange("instance_charge_type") && resourceData.Get("instance_charge_type").(string) == "PrePaid" {
		//只有当没执行实例状态变更才生效并且是预付费
		renewInstance := ve.Callback{
			Call: ve.SdkCall{
				Action:         "RenewInstance",
				ConvertMode:    ve.RequestConvertInConvert,
				RequestIdField: "InstanceId",
				Convert: map[string]ve.RequestConvert{
					"period": {
						ConvertType: ve.ConvertDefault,
						Convert: func(data *schema.ResourceData, i interface{}) interface{} {
							o, n := data.GetChange("period")
							return n.(int) - o.(int)
						},
					},
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					if len(*call.SdkParam) > 1 {
						if (*call.SdkParam)["Period"].(int) <= 0 {
							return false, fmt.Errorf("period must set and more than 0 ")
						}
						//(*call.SdkParam)["PeriodUnit"] = d.Get("period_unit")
						(*call.SdkParam)["PeriodUnit"] = "Month"
						return true, nil
					}
					return false, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					(*call.SdkParam)["ClientToken"] = uuid.New().String()
					logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
					//续费实例
					return s.Client.EcsClient.RenewInstanceCommon(call.SdkParam)
				},
				AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
					return nil
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{"RUNNING", "STOPPED"},
					Timeout: resourceData.Timeout(schema.TimeoutUpdate),
				},
			},
		}
		callbacks = append(callbacks, renewInstance)
	}

	//总体重启 如果需要
	var stopped bool

	stopInstance := s.StartOrStopInstanceCallback(resourceData, true)
	stopInstance.Call.BeforeCall = func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
		instance, err := s.ReadResource(resourceData, resourceData.Id())
		if err != nil {
			return false, err
		}
		status, err := ve.ObtainSdkValue("Status", instance)
		if err != nil {
			return false, err
		}
		if status.(string) == "RUNNING" {
			return passwordChange, nil
		}
		return false, nil
	}
	stopInstance.Call.AfterCall = func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
		stopped = true
		return nil
	}
	callbacks = append(callbacks, stopInstance)

	startInstance := s.StartOrStopInstanceCallback(resourceData, false)
	startInstance.Call.BeforeCall = func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
		return stopped && passwordChange, nil
	}
	callbacks = append(callbacks, startInstance)

	return callbacks
}

func (s *VestackEcsService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteInstance",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"InstanceId": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				//删除ECS
				return s.Client.EcsClient.DeleteInstanceCommon(call.SdkParam)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					ecs, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on  reading ecs on delete %q, %w", d.Id(), callErr))
						}
					}

					if ecs["InstanceChargeType"] == "PrePaid" {
						return resource.NonRetryableError(fmt.Errorf("PrePaid instance charge type not support remove,Please change instance charge type to PostPaid. "))
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

func (s *VestackEcsService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "InstanceIds",
				ConvertType: ve.ConvertWithN,
			},
		},
		NameField:        "InstanceName",
		IdField:          "InstanceId",
		CollectField:     "instances",
		ResponseConverts: s.CommonResponseConvert(),
		ExtraData: func(sourceData []interface{}) (extraData []interface{}, err error) {
			return s.readEbsVolumes(sourceData)
		},
	}
}

func (s *VestackEcsService) CommonResponseConvert() map[string]ve.ResponseConvert {
	return map[string]ve.ResponseConvert{
		"Id": {
			TargetField: "instance_id",
		},
		"InstanceTypeId": {
			TargetField: "instance_type",
		},
		"InstanceType": {
			Ignore: true,
		},
		"SystemVolumeSize": {
			TargetField: "system_volume_size",
			Convert: func(i interface{}) interface{} {
				size, _ := strconv.Atoi(i.(string))
				return size
			},
		},
		"DataVolumes": {
			TargetField: "data_volumes",
			Convert: func(i interface{}) interface{} {
				var results []interface{}
				if dd, ok := i.([]interface{}); ok {
					for _, _data := range dd {
						if v, ok1 := _data.(map[string]interface{}); ok1 {
							if reflect.TypeOf(v["Size"]).Kind() == reflect.String {
								v["Size"], _ = strconv.Atoi(v["Size"].(string))
							}
							results = append(results, v)
						}
					}
				}
				return results
			},
		},
		"Volumes": {
			TargetField: "volumes",
			Convert: func(i interface{}) interface{} {
				var results []interface{}
				if dd, ok := i.([]interface{}); ok {
					for _, _data := range dd {
						if v, ok1 := _data.(map[string]interface{}); ok1 {
							if reflect.TypeOf(v["Size"]).Kind() == reflect.String {
								v["Size"], _ = strconv.Atoi(v["Size"].(string))
							}
							results = append(results, v)
						}
					}
				}
				return results
			},
		},
	}
}

func (s *VestackEcsService) StartOrStopInstanceCallback(resourceData *schema.ResourceData, isStop bool) ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"InstanceId": resourceData.Id(),
			},
		},
	}
	if isStop {
		callback.Call.Action = "StopInstance"
		callback.Call.ExecuteCall = func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
			logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
			return s.Client.EcsClient.StopInstanceCommon(call.SdkParam)
		}
		callback.Call.Refresh = &ve.StateRefresh{
			Target:  []string{"STOPPED"},
			Timeout: resourceData.Timeout(schema.TimeoutUpdate),
		}
	} else {
		callback.Call.Action = "StartInstance"
		callback.Call.ExecuteCall = func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
			logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
			return s.Client.EcsClient.StartInstanceCommon(call.SdkParam)
		}
		callback.Call.Refresh = &ve.StateRefresh{
			Target:  []string{"RUNNING"},
			Timeout: resourceData.Timeout(schema.TimeoutUpdate),
		}
	}
	return callback
}

func (s *VestackEcsService) ReadResourceId(id string) string {
	return id
}

func (s *VestackEcsService) readEbsVolumes(sourceData []interface{}) (extraData []interface{}, err error) {
	//merge ebs
	var (
		wg      sync.WaitGroup
		syncMap sync.Map
	)
	if len(sourceData) == 0 {
		return sourceData, err
	}
	wg.Add(len(sourceData))
	for _, data := range sourceData {
		instance := data
		var (
			instanceId interface{}
			action     string
			resp       *map[string]interface{}
			results    interface{}
			_err       error
		)
		go func() {
			defer func() {
				if e := recover(); e != nil {
					logger.Debug(logger.ReqFormat, action, e)
				}
				ve.Release()
				wg.Done()
			}()
			ve.Acquire()

			instanceId, _err = ve.ObtainSdkValue("InstanceId", instance)
			if _err != nil {
				syncMap.Store(instanceId, err)
				return
			}
			action = "DescribeVolumes"
			logger.Debug(logger.ReqFormat, action, instanceId)
			volumeCondition := map[string]interface{}{
				"InstanceId": instanceId,
			}
			logger.Debug(logger.ReqFormat, action, volumeCondition)
			resp, _err = s.Client.EbsClient.DescribeVolumesCommon(&volumeCondition)
			if _err != nil {
				syncMap.Store(instanceId, err)
				return
			}
			logger.Debug(logger.RespFormat, action, volumeCondition, *resp)
			results, _err = ve.ObtainSdkValue("Result.Volumes", *resp)
			if _err != nil {
				syncMap.Store(instanceId, err)
				return
			}
			syncMap.Store(instanceId, results)
		}()
	}
	wg.Wait()
	var errorStr string
	for _, instance := range sourceData {
		var (
			instanceId interface{}
		)
		instanceId, err = ve.ObtainSdkValue("InstanceId", instance)
		if err != nil {
			return
		}
		if v, ok := syncMap.Load(instanceId); ok {
			if e1, ok1 := v.(error); ok1 {
				errorStr = errorStr + e1.Error() + ";"
			}
			instance.(map[string]interface{})["Volumes"] = v
		}
		extraData = append(extraData, instance)
	}
	if len(errorStr) > 0 {
		return extraData, fmt.Errorf(errorStr)
	}
	return extraData, err
}
