package ecs_instance

import (
	"encoding/base64"
	"errors"
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/ecs_deployment_set_associate"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/eip/eip_address"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/subnet"
)

type VolcengineEcsService struct {
	Client        *ve.SdkClient
	SubnetService *subnet.VolcengineSubnetService
}

func NewEcsService(c *ve.SdkClient) *VolcengineEcsService {
	return &VolcengineEcsService{
		Client:        c,
		SubnetService: subnet.NewSubnetService(c),
	}
}

func (s *VolcengineEcsService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineEcsService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	var (
		resp               *map[string]interface{}
		results            interface{}
		next               string
		ok                 bool
		ecsInstance        map[string]interface{}
		networkInterfaces  []interface{}
		networkInterfaceId string
	)
	data, err = ve.WithNextTokenQuery(condition, "MaxResults", "NextToken", 20, nil, func(m map[string]interface{}) ([]interface{}, string, error) {
		action := "DescribeInstances"
		logger.Debug(logger.ReqFormat, action, condition)
		if condition == nil {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), nil)
			if err != nil {
				return data, next, err
			}
		} else {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &condition)
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
		data, err = RemoveSystemTags(data)
		return data, next, err
	})

	if err != nil {
		return nil, err
	}

	for _, v := range data {
		if ecsInstance, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("Value is not map ")
		} else {
			// query primary network interface info of the ecs instance
			if networkInterfaces, ok = ecsInstance["NetworkInterfaces"].([]interface{}); !ok {
				return data, errors.New("Instances.NetworkInterfaces is not Slice")
			}
			for _, networkInterface := range networkInterfaces {
				if networkInterfaceMap, ok := networkInterface.(map[string]interface{}); ok &&
					networkInterfaceMap["Type"] == "primary" {
					networkInterfaceId = networkInterfaceMap["NetworkInterfaceId"].(string)
				}
			}

			action := "DescribeNetworkInterfaces"
			req := map[string]interface{}{
				"NetworkInterfaceIds.1": networkInterfaceId,
			}
			logger.Debug(logger.ReqFormat, action, req)
			res, err := s.Client.UniversalClient.DoCall(getVpcUniversalInfo(action), &req)
			if err != nil {
				logger.Info("DescribeNetworkInterfaces error:", err)
				continue
			}
			logger.Debug(logger.RespFormat, action, condition, *res)

			networkInterfaceInfos, err := ve.ObtainSdkValue("Result.NetworkInterfaceSets", *res)
			if err != nil {
				logger.Info("ObtainSdkValue Result.NetworkInterfaceSets error:", err)
				continue
			}
			if ipv6Sets, ok := networkInterfaceInfos.([]interface{})[0].(map[string]interface{})["IPv6Sets"].([]interface{}); ok {
				ecsInstance["Ipv6Addresses"] = ipv6Sets
				ecsInstance["Ipv6AddressCount"] = len(ipv6Sets)
			}
		}
	}

	return data, err
}

func (s *VolcengineEcsService) ReadResource(resourceData *schema.ResourceData, instanceId string) (data map[string]interface{}, err error) {
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

	if numa := resourceData.Get("cpu_options.0.numa_per_socket"); numa != 0 {
		if v, exist := data["CpuOptions"]; exist {
			cpuOptions, ok := v.(map[string]interface{})
			if !ok {
				return data, fmt.Errorf("CpuOptions is not map ")
			}
			cpuOptions["NumaPerSocket"] = numa
		}
	}

	if eipId := resourceData.Get("eip_id"); eipId != "" {
		if v, exist := data["EipAddress"]; exist && v != nil {
			eipMap, ok := v.(map[string]interface{})
			if !ok {
				return data, fmt.Errorf("DescribeInstances EipAddress is not map")
			}
			if id, ok := eipMap["AllocationId"]; ok && eipId.(string) != id.(string) {
				return data, fmt.Errorf("The eip id of the instance is mismatched, specified id: %s, assigned id: %s ", eipId, id)
			}
		}
	}

	return data, nil
}

func (s *VolcengineEcsService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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
			project, err := ve.ObtainSdkValue("ProjectName", data)
			if err != nil {
				return nil, "", err
			}
			if resourceData.Get("project_name") != nil && resourceData.Get("project_name").(string) != "" {
				if project != resourceData.Get("project_name") {
					return data, "", err
				}
			}
			return data, status.(string), err
		},
	}
}

func (s *VolcengineEcsService) WithResourceResponseHandlers(ecs map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		var (
			typeErr             error
			ebsErr              error
			userDataErr         error
			networkInterfaceErr error
			errorStr            string
			wg                  sync.WaitGroup
			syncMap             sync.Map
		)
		// 使用小写的 Hostname
		delete(ecs, "HostName")

		//计算period
		if ecs["InstanceChargeType"].(string) == "PrePaid" {
			ct, _ := time.Parse("2006-01-02T15:04:05", ecs["CreatedAt"].(string)[0:strings.Index(ecs["CreatedAt"].(string), "+")])
			et, _ := time.Parse("2006-01-02T15:04:05", ecs["ExpiredAt"].(string)[0:strings.Index(ecs["ExpiredAt"].(string), "+")])
			y := et.Year() - ct.Year()
			m := et.Month() - ct.Month()
			ecs["Period"] = y*12 + int(m)
		}

		wg.Add(4)
		instanceId := ecs["InstanceId"]
		//read instance type
		go func() {
			defer func() {
				if _err := recover(); _err != nil {
					logger.Debug(logger.ReqFormat, "DescribeInstancesType", _err)
				}
				wg.Done()
			}()
			temp := map[string]interface{}{
				"InstanceTypeId": ecs["InstanceTypeId"],
			}
			_, typeErr = s.readInstanceTypes([]interface{}{temp})
			if typeErr != nil {
				return
			}
			syncMap.Store("GpuDevices", temp["GpuDevices"])
			syncMap.Store("IsGpu", temp["IsGpu"])
		}()
		//read ebs data
		go func() {
			defer func() {
				if _err := recover(); _err != nil {
					logger.Debug(logger.ReqFormat, "DescribeVolumes", _err)
				}
				wg.Done()
			}()
			temp := map[string]interface{}{
				"InstanceId":  ecs["InstanceId"],
				"ProjectName": ecs["ProjectName"],
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
				wg.Done()
			}()
			var (
				userDataParam *map[string]interface{}
				userDataResp  *map[string]interface{}
				userData      interface{}
			)
			userDataParam = &map[string]interface{}{
				"InstanceId": instanceId,
			}
			userDataResp, userDataErr = s.Client.UniversalClient.DoCall(getUniversalInfo("DescribeUserData"), userDataParam)
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
				wg.Done()
			}()
			var (
				networkInterfaceParam *map[string]interface{}
				networkInterfaceResp  *map[string]interface{}
				networkInterface      []interface{}
				networkInterfaces     []interface{}
				ok                    bool
				next                  string
			)

			networkInterfaceParam = &map[string]interface{}{}
			if networkInterfaces, ok = ecs["NetworkInterfaces"].([]interface{}); !ok {
				return
			}
			for index, networkInterface := range networkInterfaces {
				if networkInterfaceMap, ok := networkInterface.(map[string]interface{}); ok {
					(*networkInterfaceParam)[fmt.Sprintf("%s.%d", "NetworkInterfaceIds", index)] = networkInterfaceMap["NetworkInterfaceId"].(string)
				}
			}
			networkInterface, networkInterfaceErr = ve.WithNextTokenQuery(*networkInterfaceParam, "MaxResults", "NextToken", 100, nil, func(condition map[string]interface{}) ([]interface{}, string, error) {
				action := "DescribeNetworkInterfaces"
				logger.Debug(logger.ReqFormat, action, condition)
				networkInterfaceResp, networkInterfaceErr = s.Client.UniversalClient.DoCall(getVpcUniversalInfo(action), &condition)
				if networkInterfaceErr != nil {
					return networkInterface, next, networkInterfaceErr
				}
				logger.Debug(logger.RespFormat, action, condition, *networkInterfaceResp)

				results, networkInterfaceErr := ve.ObtainSdkValue("Result.NetworkInterfaceSets", *networkInterfaceResp)
				if networkInterfaceErr != nil {
					return networkInterface, next, networkInterfaceErr
				}
				nextToken, err := ve.ObtainSdkValue("Result.NextToken", *networkInterfaceResp)
				if err != nil {
					return networkInterface, next, err
				}
				next = nextToken.(string)
				if results == nil {
					results = []interface{}{}
				}

				if networkInterface, ok = results.([]interface{}); !ok {
					return networkInterface, next, errors.New("Result.NetworkInterfaceSets is not Slice")
				}
				return networkInterface, next, networkInterfaceErr
			})
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
						ecs["PrimaryIpAddress"] = v1["PrimaryIpAddress"]
					} else {
						dataNetworkInterfaces = append(dataNetworkInterfaces, vif)
					}
				}
			}
			if len(dataNetworkInterfaces) > 0 {
				ecs["SecondaryNetworkInterfaces"] = dataNetworkInterfaces
			}
		}

		//split System volume and Data volumes
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
				v1 := volumeInfo{
					list: dataVolumes,
				}
				sort.Sort(&v1)
				ecs["DataVolumes"] = v1.list
			}
		}
		return ecs, s.CommonResponseConvert(), nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineEcsService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "RunInstances",
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"zone_id": {
					ConvertType: ve.ConvertDefault,
					ForceGet:    true,
				},
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
				"primary_ip_address": {
					ConvertType: ve.ConvertDefault,
					TargetField: "NetworkInterfaces.1.PrimaryIpAddress",
				},
				"data_volumes": {
					ConvertType: ve.ConvertListN,
					TargetField: "Volumes",
					StartIndex:  1,
					NextLevelConvert: map[string]ve.RequestConvert{
						"delete_with_instance": {
							TargetField: "DeleteWithInstance",
							ForceGet:    true,
						},
					},
				},
				"cpu_options": {
					ConvertType: ve.ConvertListUnique,
					TargetField: "CpuOptions",
					NextLevelConvert: map[string]ve.RequestConvert{
						"threads_per_core": {
							TargetField: "ThreadsPerCore",
						},
						"numa_per_socket": {
							TargetField: "NumaPerSocket",
						},
					},
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
				"user_data": {
					ConvertType: ve.ConvertDefault,
					TargetField: "UserData",
					Convert: func(data *schema.ResourceData, i interface{}) interface{} {
						_, base64DecodeError := base64.StdEncoding.DecodeString(i.(string))
						if base64DecodeError == nil {
							return i.(string)
						} else {
							return base64.StdEncoding.EncodeToString([]byte(i.(string)))
						}
					},
				},
				"tags": {
					TargetField: "Tags",
					ConvertType: ve.ConvertListN,
				},
				"ipv6_address_count": {
					Ignore: true,
				},
				"ipv6_addresses": {
					Ignore: true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["ClientToken"] = uuid.New().String()
				(*call.SdkParam)["Volumes.1.DeleteWithInstance"] = true
				(*call.SdkParam)["Count"] = 1

				if (*call.SdkParam)["InstanceChargeType"] != "PrePaid" {
					delete(*call.SdkParam, "AutoRenew")
					delete(*call.SdkParam, "AutoRenewPeriod")
					delete(*call.SdkParam, "Period")
				}

				if _, ok := (*call.SdkParam)["ZoneId"]; !ok || (*call.SdkParam)["ZoneId"] == "" {
					var (
						vnet map[string]interface{}
						err  error
						zone interface{}
					)
					vnet, err = s.SubnetService.ReadResource(d, (*call.SdkParam)["NetworkInterfaces.1.SubnetId"].(string))
					if err != nil {
						return false, err
					}
					zone, err = ve.ObtainSdkValue("ZoneId", vnet)
					if err != nil {
						return false, err
					}
					(*call.SdkParam)["ZoneId"] = zone
				}

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
				return s.Client.UniversalClient.DoCall(getUniversalInfo("RunInstances"), call.SdkParam)
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
	callbacks = append(callbacks, callback)

	// 分配Ipv6
	ipv6Callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "AssignIpv6Addresses",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				ipv6AddressCount, ok1 := d.GetOk("ipv6_address_count")
				ipv6Addresses, ok2 := d.GetOk("ipv6_addresses")
				if !ok1 && !ok2 {
					return false, nil
				}

				var (
					networkInterfaceId string
					networkInterfaces  []interface{}
					ok                 bool
				)
				ecsInstance, err := s.ReadResource(resourceData, d.Id())
				if err != nil {
					return false, err
				}
				// query primary network interface info of the ecs instance
				if networkInterfaces, ok = ecsInstance["NetworkInterfaces"].([]interface{}); !ok {
					return false, errors.New("Instances.NetworkInterfaces is not Slice")
				}
				for _, networkInterface := range networkInterfaces {
					if networkInterfaceMap, ok := networkInterface.(map[string]interface{}); ok &&
						networkInterfaceMap["Type"] == "primary" {
						networkInterfaceId = networkInterfaceMap["NetworkInterfaceId"].(string)
					}
				}

				(*call.SdkParam)["NetworkInterfaceId"] = networkInterfaceId
				if ok1 {
					(*call.SdkParam)["Ipv6AddressCount"] = ipv6AddressCount.(int)
				} else if ok2 {
					for index, ipv6Address := range ipv6Addresses.(*schema.Set).List() {
						(*call.SdkParam)["Ipv6Address."+strconv.Itoa(index)] = ipv6Address
					}
				}

				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				//分配Ipv6地址
				return s.Client.UniversalClient.DoCall(getVpcUniversalInfo(call.Action), call.SdkParam)
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"RUNNING"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	callbacks = append(callbacks, ipv6Callback)

	// 绑定eip
	eipCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "AssociateEipAddress",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				eipId, ok := d.GetOk("eip_id")
				if !ok {
					return false, nil
				}
				(*call.SdkParam)["AllocationId"] = eipId.(string)
				(*call.SdkParam)["InstanceId"] = d.Id()
				(*call.SdkParam)["InstanceType"] = "EcsInstance"
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				output, err := s.Client.UniversalClient.DoCall(getVpcUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, *call.SdkParam, *output)
				if err != nil {
					d.Set("eip_id", nil)
				}
				return output, err
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"RUNNING"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				eip_address.NewEipAddressService(s.Client): {
					Target:     []string{"Attached"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					ResourceId: resourceData.Get("eip_id").(string),
				},
			},
		},
	}
	callbacks = append(callbacks, eipCallback)

	return callbacks
}

func (s *VolcengineEcsService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) (callbacks []ve.Callback) {
	var (
		passwordChange bool
		flag           bool
	)

	if resourceData.HasChange("password") && !resourceData.HasChange("image_id") {
		passwordChange = true
	}

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
					delete(*call.SdkParam, "Tags")
					return true, nil
				}
				return false, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				(*call.SdkParam)["ClientToken"] = uuid.New().String()
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				//修改实例属性
				return s.Client.UniversalClient.DoCall(getUniversalInfo("ModifyInstanceAttribute"), call.SdkParam)
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
				"auto_renew": {
					ConvertType: ve.ConvertDefault,
					ForceGet:    true,
				},
				"auto_renew_period": {
					ConvertType: ve.ConvertDefault,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if len(*call.SdkParam) > 3 {
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
				return s.Client.UniversalClient.DoCall(getUniversalInfo("ModifyInstanceChargeType"), call.SdkParam)
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
					return s.Client.UniversalClient.DoCall(getVpcUniversalInfo("ModifyNetworkInterfaceAttributes"), call.SdkParam)
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
				return s.Client.UniversalClient.DoCall(getEbsUniversalInfo(call.Action), call.SdkParam)
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
					return s.Client.UniversalClient.DoCall(getUniversalInfo("RenewInstance"), call.SdkParam)
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
	//only password changed need stop
	if passwordChange {
		stopInstance := s.StartOrStopInstanceCallback(resourceData, true, &flag)
		callbacks = append(callbacks, stopInstance)
	}
	//instance_type
	if resourceData.HasChange("instance_type") {
		//need stop before ModifyInstanceSpec

		stopInstance := s.StartOrStopInstanceCallback(resourceData, true, &flag)
		callbacks = append(callbacks, stopInstance)

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
					return s.Client.UniversalClient.DoCall(getUniversalInfo("ModifyInstanceSpec"), call.SdkParam)
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
	}
	//image change
	if resourceData.HasChange("image_id") {
		//need stop before ReplaceSystemVolume
		stopInstance := s.StartOrStopInstanceCallback(resourceData, true, &flag)
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
					"keep_image_credential": {
						ConvertType: ve.ConvertDefault,
						ForceGet:    true,
					},
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					keyPairName, exist := d.GetOkExists("key_pair_name")
					if !exist || keyPairName == "" {
						delete(*call.SdkParam, "KeyPairName")
					}
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo("ReplaceSystemVolume"), call.SdkParam)
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

	if resourceData.HasChange("deployment_set_id") {
		stopInstance := s.StartOrStopInstanceCallback(resourceData, true, &flag)
		callbacks = append(callbacks, stopInstance)
		deploymentSet := ve.Callback{
			Call: ve.SdkCall{
				Action:         "ModifyInstanceDeployment",
				ConvertMode:    ve.RequestConvertInConvert,
				Convert:        map[string]ve.RequestConvert{},
				RequestIdField: "InstanceId",
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["DeploymentSetId"] = resourceData.Get("deployment_set_id")
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
				AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
					return nil
				},
			},
		}
		refresh := map[ve.ResourceService]*ve.StateRefresh{
			ecs_deployment_set_associate.NewEcsDeploymentSetAssociateService(s.Client): {
				Target:     []string{"success"},
				Timeout:    resourceData.Timeout(schema.TimeoutCreate),
				ResourceId: resourceData.Get("deployment_set_id").(string) + ":" + resourceData.Id(),
			},
		}

		if resourceData.Get("deployment_set_id").(string) != "" {
			deploymentSet.Call.ExtraRefresh = refresh
		}

		callbacks = append(callbacks, deploymentSet)
	}

	startInstance := s.StartOrStopInstanceCallback(resourceData, false, &flag)
	callbacks = append(callbacks, startInstance)

	// 更新Tags
	setResourceTagsCallbacks := ve.SetResourceTags(s.Client, "CreateTags", "DeleteTags", "instance", resourceData, getUniversalInfo)
	callbacks = append(callbacks, setResourceTagsCallbacks...)

	return callbacks
}

func (s *VolcengineEcsService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	// 解绑eip
	eipCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DisassociateEipAddress",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				eipId, ok := d.GetOk("eip_id")
				if !ok {
					return false, nil
				}
				(*call.SdkParam)["AllocationId"] = eipId.(string)
				(*call.SdkParam)["InstanceId"] = d.Id()
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getVpcUniversalInfo(call.Action), call.SdkParam)
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"RUNNING"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				eip_address.NewEipAddressService(s.Client): {
					Target:     []string{"Available"},
					Timeout:    resourceData.Timeout(schema.TimeoutDelete),
					ResourceId: resourceData.Get("eip_id").(string),
				},
			},
		},
	}
	callbacks = append(callbacks, eipCallback)

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
				return s.Client.UniversalClient.DoCall(getUniversalInfo("DeleteInstance"), call.SdkParam)
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

					//if ecs["InstanceChargeType"] == "PrePaid" {
					//	return resource.NonRetryableError(fmt.Errorf("PrePaid instance charge type not support remove,Please change instance charge type to PostPaid. "))
					//}

					if ecs["InstanceChargeType"] == "PrePaid" {
						logger.Debug(logger.RespFormat, call.Action, "PrePaid instance charge type not support remove,Only Remove from State")
						return nil
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
	callbacks = append(callbacks, callback)

	return callbacks
}

func (s *VolcengineEcsService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "InstanceIds",
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
			"deployment_set_ids": {
				TargetField: "DeploymentSetIds",
				ConvertType: ve.ConvertWithN,
			},
			"eip_addresses": {
				TargetField: "EipAddresses",
				ConvertType: ve.ConvertWithN,
			},
			"ipv6_addresses": {
				TargetField: "Ipv6Addresses",
				ConvertType: ve.ConvertWithN,
			},
			"instance_type_families": {
				TargetField: "InstanceTypeFamilies",
				ConvertType: ve.ConvertWithN,
			},
			"instance_type_ids": {
				TargetField: "InstanceTypeIds",
				ConvertType: ve.ConvertWithN,
			},
		},
		NameField:        "InstanceName",
		IdField:          "InstanceId",
		CollectField:     "instances",
		ResponseConverts: s.CommonResponseConvert(),
		ExtraData: func(sourceData []interface{}) (extraData []interface{}, err error) {
			sourceData, err = s.readInstanceTypes(sourceData)
			if err != nil {
				return extraData, err
			}
			sourceData, err = s.readEbsVolumes(sourceData)
			if err != nil {
				return extraData, err
			}
			return sourceData, err
		},
	}
}

func (s *VolcengineEcsService) CommonResponseConvert() map[string]ve.ResponseConvert {
	return map[string]ve.ResponseConvert{
		"Id": {
			TargetField: "instance_id",
		},
		"Hostname": {
			TargetField: "host_name",
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
		"UserData": {
			TargetField: "user_data",
			Convert: func(i interface{}) interface{} {
				v, base64DecodeError := base64.StdEncoding.DecodeString(i.(string))
				if base64DecodeError != nil {
					v = []byte(i.(string))
				}
				return string(v)
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
		"GpuDevices": {
			TargetField: "gpu_devices",
			Convert: func(i interface{}) interface{} {
				var results []interface{}
				if dd, ok := i.([]interface{}); ok {
					for _, _data := range dd {
						if v, ok1 := _data.(map[string]interface{}); ok1 {
							memorySize, _ := ve.ObtainSdkValue("Memory.Size", v)
							encryptedMemorySize, _ := ve.ObtainSdkValue("Memory.EncryptedSize", v)
							delete(v, "Memory")
							v["MemorySize"] = memorySize
							v["EncryptedMemorySize"] = encryptedMemorySize
							results = append(results, v)
						}
					}
				}
				return results
			},
		},
	}
}

func (s *VolcengineEcsService) StartOrStopInstanceCallback(resourceData *schema.ResourceData, isStop bool, flag *bool) ve.Callback {
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
		callback.Call.BeforeCall = func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
			instance, err := s.ReadResource(resourceData, resourceData.Id())
			if err != nil {
				return false, err
			}
			status, err := ve.ObtainSdkValue("Status", instance)
			if err != nil {
				return false, err
			}
			if status.(string) == "RUNNING" {
				return true, nil
			}
			return false, nil
		}
		callback.Call.ExecuteCall = func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
			logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
			return s.Client.UniversalClient.DoCall(getUniversalInfo("StopInstance"), call.SdkParam)
		}
		callback.Call.AfterCall = func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
			*flag = true
			return nil
		}
		callback.Call.Refresh = &ve.StateRefresh{
			Target:  []string{"STOPPED"},
			Timeout: resourceData.Timeout(schema.TimeoutUpdate),
		}
	} else {
		callback.Call.Action = "StartInstance"
		callback.Call.BeforeCall = func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
			instance, err := s.ReadResource(resourceData, resourceData.Id())
			if err != nil {
				return false, err
			}
			status, err := ve.ObtainSdkValue("Status", instance)
			if err != nil {
				return false, err
			}
			if status.(string) == "RUNNING" {
				return false, nil
			}
			return *flag, nil
		}
		callback.Call.ExecuteCall = func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
			logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
			return s.Client.UniversalClient.DoCall(getUniversalInfo("StartInstance"), call.SdkParam)
		}
		callback.Call.Refresh = &ve.StateRefresh{
			Target:  []string{"RUNNING"},
			Timeout: resourceData.Timeout(schema.TimeoutUpdate),
		}
	}
	return callback
}

func (s *VolcengineEcsService) ReadResourceId(id string) string {
	return id
}

func (s *VolcengineEcsService) readInstanceTypes(sourceData []interface{}) (extraData []interface{}, err error) {
	//merge instance_type_info
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
			instanceTypeId interface{}
			action         string
			resp           *map[string]interface{}
			results        interface{}
			_err           error
		)
		go func() {
			defer func() {
				if e := recover(); e != nil {
					logger.Debug(logger.ReqFormat, action, e)
				}
				wg.Done()
			}()

			instanceTypeId, _err = ve.ObtainSdkValue("InstanceTypeId", instance)
			if _err != nil {
				syncMap.Store(instanceTypeId, err)
				return
			}
			//if exist continue
			if _, ok := syncMap.Load(instanceTypeId); ok {
				return
			}

			action = "DescribeInstanceTypes"
			logger.Debug(logger.ReqFormat, action, instanceTypeId)
			instanceTypeCondition := map[string]interface{}{
				"InstanceTypeIds.1": instanceTypeId,
			}
			logger.Debug(logger.ReqFormat, action, instanceTypeCondition)
			resp, _err = s.Client.UniversalClient.DoCall(getUniversalInfo("DescribeInstanceTypes"), &instanceTypeCondition)
			if _err != nil {
				syncMap.Store(instanceTypeId, err)
				return
			}
			logger.Debug(logger.RespFormat, action, instanceTypeCondition, *resp)
			results, _err = ve.ObtainSdkValue("Result.InstanceTypes.0", *resp)
			if _err != nil {
				syncMap.Store(instanceTypeId, err)
				return
			}
			syncMap.Store(instanceTypeId, results)
		}()
	}
	wg.Wait()
	var errorStr string
	for _, instance := range sourceData {
		var (
			instanceTypeId interface{}
			gpu            interface{}
			gpuDevices     interface{}
		)
		instanceTypeId, err = ve.ObtainSdkValue("InstanceTypeId", instance)
		if err != nil {
			return
		}
		if v, ok := syncMap.Load(instanceTypeId); ok {
			if e1, ok1 := v.(error); ok1 {
				errorStr = errorStr + e1.Error() + ";"
			}
			gpu, _ = ve.ObtainSdkValue("Gpu", v)
			if gpu != nil {
				gpuDevices, _ = ve.ObtainSdkValue("Gpu.GpuDevices", v)
				instance.(map[string]interface{})["GpuDevices"] = gpuDevices
				instance.(map[string]interface{})["IsGpu"] = true
			} else {
				instance.(map[string]interface{})["GpuDevices"] = []interface{}{}
				instance.(map[string]interface{})["IsGpu"] = false
			}
		}
		extraData = append(extraData, instance)
	}
	if len(errorStr) > 0 {
		return extraData, fmt.Errorf(errorStr)
	}
	return extraData, err
}

func (s *VolcengineEcsService) readEbsVolumes(sourceData []interface{}) (extraData []interface{}, err error) {
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
			instanceId  interface{}
			projectName interface{}
			action      string
			resp        *map[string]interface{}
			results     interface{}
			volumes     []interface{}
			_err        error
		)
		go func() {
			defer func() {
				if e := recover(); e != nil {
					logger.Debug(logger.ReqFormat, action, e)
				}
				wg.Done()
			}()

			instanceId, _err = ve.ObtainSdkValue("InstanceId", instance)
			if _err != nil {
				syncMap.Store(instanceId, _err)
				return
			}
			projectName, _err = ve.ObtainSdkValue("ProjectName", instance)
			if _err != nil {
				syncMap.Store(instanceId, _err)
				return
			}

			// query system volume
			systemVolume, _err := s.describeSystemVolume(instanceId.(string), projectName.(string))
			if _err != nil {
				syncMap.Store(instanceId, _err)
				return
			}
			volumes = append(volumes, systemVolume)

			// query data volumes
			action = "DescribeVolumes"
			logger.Debug(logger.ReqFormat, action, instanceId)
			volumeCondition := map[string]interface{}{
				"InstanceId": instanceId,
				"Kind":       "data",
			}
			logger.Debug(logger.ReqFormat, action, volumeCondition)
			resp, _err = s.Client.UniversalClient.DoCall(getEbsUniversalInfo(action), &volumeCondition)
			if _err != nil {
				if ve.AccessDeniedError(_err) {
					// 权限错误，直接返回
					syncMap.Store(instanceId, volumes)
					return
				} else {
					syncMap.Store(instanceId, _err)
					return
				}
			}
			logger.Debug(logger.RespFormat, action, volumeCondition, *resp)
			results, _err = ve.ObtainSdkValue("Result.Volumes", *resp)
			if _err != nil {
				syncMap.Store(instanceId, _err)
				return
			}
			if results == nil {
				results = []interface{}{}
			}
			dataVolumes, ok := results.([]interface{})
			if !ok {
				syncMap.Store(instanceId, errors.New("Result.Volumes is not Slice"))
				return
			}
			volumes = append(volumes, dataVolumes)

			syncMap.Store(instanceId, volumes)
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

func (s *VolcengineEcsService) describeSystemVolume(instanceId, projectName string) (map[string]interface{}, error) {
	var (
		action       string
		req          *map[string]interface{}
		resp         *map[string]interface{}
		results      interface{}
		systemVolume map[string]interface{}
		err          error
	)

	action = "DescribeVolumes"
	req = &map[string]interface{}{
		"InstanceId":  instanceId,
		"ProjectName": projectName,
		"Kind":        "system",
	}
	logger.Debug(logger.ReqFormat, action, *req)
	resp, err = s.Client.UniversalClient.DoCall(getEbsUniversalInfo("DescribeVolumes"), req)
	if err != nil {
		return systemVolume, err
	}
	logger.Debug(logger.RespFormat, action, *req, *resp)
	results, err = ve.ObtainSdkValue("Result.Volumes", *resp)
	if err != nil {
		return systemVolume, err
	}
	if results == nil {
		results = []interface{}{}
	}
	volumes, ok := results.([]interface{})
	if !ok {
		return systemVolume, errors.New("Result.Volumes is not Slice")
	}
	for _, volume := range volumes {
		if systemVolume, ok = volume.(map[string]interface{}); !ok {
			return systemVolume, errors.New("Volumes Value is not map ")
		}
	}
	if len(systemVolume) == 0 {
		return systemVolume, fmt.Errorf("System Volume of %s is empty ", instanceId)
	}
	return systemVolume, nil
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "ecs",
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		Action:      actionName,
	}
}

func getVpcUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "vpc",
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		Action:      actionName,
	}
}

func getEbsUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "storage_ebs",
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		Action:      actionName,
	}
}

type volumeInfo struct {
	list []interface{}
}

func (v *volumeInfo) Len() int {
	return len(v.list)
}

func (v *volumeInfo) Less(i, j int) bool {
	return v.list[i].(map[string]interface{})["VolumeName"].(string) < v.list[j].(map[string]interface{})["VolumeName"].(string)
}

func (v *volumeInfo) Swap(i, j int) {
	v.list[i], v.list[j] = v.list[j], v.list[i]
}

func (s *VolcengineEcsService) ProjectTrn() *ve.ProjectTrn {
	return &ve.ProjectTrn{
		ServiceName:          "ecs",
		ResourceType:         "instance",
		ProjectResponseField: "ProjectName",
		ProjectSchemaField:   "project_name",
	}
}

func (s *VolcengineEcsService) UnsubscribeInfo(resourceData *schema.ResourceData, resource *schema.Resource) (*ve.UnsubscribeInfo, error) {
	info := ve.UnsubscribeInfo{
		InstanceId: s.ReadResourceId(resourceData.Id()),
	}
	if resourceData.Get("instance_charge_type") == "PrePaid" {
		//查询实例类型的配置
		action := "DescribeInstanceTypes"
		input := map[string]interface{}{
			"InstanceTypeIds.1": resourceData.Get("instance_type"),
		}
		var (
			output *map[string]interface{}
			err    error
			t      interface{}
		)
		output, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &input)
		if err != nil {
			return &info, err
		}
		t, err = ve.ObtainSdkValue("Result.InstanceTypes.0", *output)
		if err != nil {
			return &info, err
		}
		if tt, ok := t.(map[string]interface{}); ok {
			if tt["Gpu"] != nil && tt["Rdma"] != nil {
				info.Products = []string{"HPC_GPU", "ECS", "ECS_BareMetal", "GPU_Server"}
			} else if tt["Gpu"] != nil && tt["Rdma"] == nil {
				info.Products = []string{"GPU_Server", "ECS", "ECS_BareMetal", "HPC_GPU"}
			} else {
				info.Products = []string{"ECS", "ECS_BareMetal", "GPU_Server", "HPC_GPU"}
			}
			info.NeedUnsubscribe = true
		}

	}
	return &info, nil
}
