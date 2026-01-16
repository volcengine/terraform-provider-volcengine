package alb

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineAlbService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewAlbService(c *ve.SdkClient) *VolcengineAlbService {
	return &VolcengineAlbService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineAlbService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineAlbService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	data, err = ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeLoadBalancers"

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
		results, err = ve.ObtainSdkValue("Result.LoadBalancers", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.LoadBalancers is not Slice")
		}
		data, err = removeSystemTags(data)
		return data, err
	})
	if err != nil {
		return data, err
	}

	for _, value := range data {
		alb, ok := value.(map[string]interface{})
		if !ok {
			return data, fmt.Errorf("Alb is not map ")
		}

		detailAction := "DescribeLoadBalancerAttributes"
		req := map[string]interface{}{
			"LoadBalancerId": alb["LoadBalancerId"],
		}
		logger.Debug(logger.ReqFormat, detailAction, req)
		detailResp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(detailAction), &req)
		if err != nil {
			return data, err
		}
		logger.Debug(logger.RespFormat, detailAction, *detailResp)

		listeners, err := ve.ObtainSdkValue("Result.Listeners", *detailResp)
		if err != nil {
			return data, err
		}
		alb["Listeners"] = listeners

		accessLog, err := ve.ObtainSdkValue("Result.AccessLog", *detailResp)
		if err != nil {
			return data, err
		}
		alb["AccessLog"] = accessLog

		tlsAccessLog, err := ve.ObtainSdkValue("Result.TLSAccessLog", *detailResp)
		if err != nil {
			return data, err
		}
		alb["TLSAccessLog"] = tlsAccessLog

		healthLog, err := ve.ObtainSdkValue("Result.HealthLog", *detailResp)
		if err != nil {
			return data, err
		}
		alb["HealthLog"] = healthLog

		enabled, err := ve.ObtainSdkValue("Result.Enabled", *detailResp)
		if err != nil {
			return data, err
		}
		alb["Enabled"] = enabled

		globalAccelerators, err := ve.ObtainSdkValue("Result.GlobalAccelerators", *detailResp)
		if err != nil {
			return data, err
		}
		alb["GlobalAccelerators"] = globalAccelerators
	}

	return data, err
}

func (s *VolcengineAlbService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"LoadBalancerIds.1": id,
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
		return data, fmt.Errorf("alb %s not exist ", id)
	}

	var subnetIds []interface{}
	zoneMappings, ok := data["ZoneMappings"].([]interface{})
	if !ok {
		return data, fmt.Errorf("ZoneMappings is not slice ")
	}
	for _, v := range zoneMappings {
		zoneMap, ok := v.(map[string]interface{})
		if !ok {
			return data, fmt.Errorf("Zone is not map ")
		}
		subnetIds = append(subnetIds, zoneMap["SubnetId"])
		addresses, ok := zoneMap["LoadBalancerAddresses"].([]interface{})
		if !ok || len(addresses) == 0 {
			return data, fmt.Errorf("LoadBalancerAddresses is not slice ")
		}
		addressMap, ok := addresses[0].(map[string]interface{})
		if !ok {
			return data, fmt.Errorf("LoadBalancerAddresse is not map ")
		}

		if _, exist := addressMap["Eip"]; exist && addressMap["Eip"] != nil {
			eip, ok := addressMap["Eip"].(map[string]interface{})
			if !ok {
				return data, fmt.Errorf("Eip of LoadBalancerAddresse is not map ")
			}
			data["EipBillingConfig"] = eip
		}
		if _, exist := addressMap["Ipv6Eip"]; exist && addressMap["Ipv6Eip"] != nil {
			ipv6Eip, ok := addressMap["Ipv6Eip"].(map[string]interface{})
			if !ok {
				return data, fmt.Errorf("Ipv6Eip of LoadBalancerAddresse is not map ")
			}
			data["Ipv6EipBillingConfig"] = ipv6Eip
		}
	}
	data["SubnetIds"] = subnetIds

	return data, err
}

func (VolcengineAlbService) WithResourceResponseHandlers(alb map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return alb, map[string]ve.ResponseConvert{
			"ISP": {
				TargetField: "isp",
			},
			"EipBillingType": {
				TargetField: "eip_billing_type",
				Convert: func(i interface{}) interface{} {
					if i == nil {
						return nil
					}
					billingType := i.(float64)
					switch billingType {
					case 2:
						return "PostPaidByBandwidth"
					case 3:
						return "PostPaidByTraffic"
					}
					return ""
				},
			},
			"BillingType": {
				TargetField: "billing_type",
				Convert: func(i interface{}) interface{} {
					if i == nil {
						return nil
					}
					billingType := i.(float64)
					switch billingType {
					case 2:
						return "PostPaidByBandwidth"
					case 3:
						return "PostPaidByTraffic"
					}
					return ""
				},
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineAlbService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				d          map[string]interface{}
				status     interface{}
				failStates []string
			)
			failStates = append(failStates, "CreateFailed")
			d, err = s.ReadResource(resourceData, id)
			if err != nil {
				return nil, "", err
			}
			status, err = ve.ObtainSdkValue("Status", d)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("alb status error, status: %s", status.(string))
				}
			}
			return d, status.(string), err
		},
	}
}

func (s *VolcengineAlbService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	// 检查是否指定了source_load_balancer_id，如果指定则使用CloneLoadBalancer API
	if sourceId, ok := resourceData.Get("source_load_balancer_id").(string); ok && sourceId != "" {
		return s.CreateCloneResource(resourceData, resource)
	}
	var callbacks []ve.Callback
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateLoadBalancer",
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"subnet_ids": {
					Ignore: true,
				},
				"allocation_ids": {
					Ignore: true,
				},
				"eip_billing_config": {
					TargetField: "EipBillingConfig",
					ConvertType: ve.ConvertListUnique,
					NextLevelConvert: map[string]ve.RequestConvert{
						"isp": {
							TargetField: "ISP",
						},
					},
				},
				"ipv6_eip_billing_config": {
					TargetField: "Ipv6EipBillingConfig",
					ConvertType: ve.ConvertListUnique,
					NextLevelConvert: map[string]ve.RequestConvert{
						"isp": {
							TargetField: "ISP",
						},
					},
				},
				"tags": {
					TargetField: "Tags",
					ConvertType: ve.ConvertListN,
				},
				// CreateLoadBalancer不支持的参数，在创建时忽略
				"waf_protection_enabled": {
					Ignore: true,
				},
				"waf_instance_id": {
					Ignore: true,
				},
				"waf_protected_domain": {
					Ignore: true,
				},
				"global_accelerator": {
					Ignore: true,
				},
				"proxy_protocol_enabled": {
					Ignore: true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["RegionId"] = s.Client.Region
				(*call.SdkParam)["LoadBalancerBillingType"] = 1

				subnetIds, ok := d.Get("subnet_ids").(*schema.Set)
				if !ok {
					return false, fmt.Errorf("SubnetIds is not set")
				}
				vpcId, zoneIds, err := s.getVpcIdAndZoneIdBySubnets(subnetIds.List())
				if err != nil {
					return false, err
				}
				if subnetIds.Len() != len(zoneIds) {
					return false, fmt.Errorf("The length of subnetIds and zoneIds are not equal ")
				}
				for index, subnetId := range subnetIds.List() {
					(*call.SdkParam)["ZoneMappings."+strconv.Itoa(index+1)+".SubnetId"] = subnetId.(string)
					(*call.SdkParam)["ZoneMappings."+strconv.Itoa(index+1)+".ZoneId"] = zoneIds[subnetId.(string)]
				}
				(*call.SdkParam)["VpcId"] = vpcId

				// private 类型不传 eip_billing_config / ipv6_eip_billing_config
				if (*call.SdkParam)["Type"] == "private" {
					delete(*call.SdkParam, "EipBillingConfig.ISP")
					delete(*call.SdkParam, "EipBillingConfig.EipBillingType")
					delete(*call.SdkParam, "EipBillingConfig.Bandwidth")
					delete(*call.SdkParam, "Ipv6EipBillingConfig.ISP")
					delete(*call.SdkParam, "Ipv6EipBillingConfig.BillingType")
					delete(*call.SdkParam, "Ipv6EipBillingConfig.Bandwidth")
				}

				if eipBillingType, exist := (*call.SdkParam)["EipBillingConfig.EipBillingType"]; exist {
					ty := 0
					switch eipBillingType.(string) {
					case "PostPaidByBandwidth":
						ty = 2
					case "PostPaidByTraffic":
						ty = 3
					}
					(*call.SdkParam)["EipBillingConfig.EipBillingType"] = ty
				}
				if ipv6BillingType, exist := (*call.SdkParam)["Ipv6EipBillingConfig.BillingType"]; exist {
					ty := 0
					switch ipv6BillingType.(string) {
					case "PostPaidByBandwidth":
						ty = 2
					case "PostPaidByTraffic":
						ty = 3
					}
					(*call.SdkParam)["Ipv6EipBillingConfig.BillingType"] = ty
				}

				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("Result.LoadBalancerId", *resp)
				d.SetId(id.(string))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Active"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	callbacks = append(callbacks, callback)
	// 检查创建后是否需要调用ModifyLoadBalancerAttributes来配置不支持的参数
	if s.hasPostCreationAttributes(resourceData) {
		modifyCallback := s.createPostCreationModifyCallback(resourceData)
		callbacks = append(callbacks, modifyCallback)
	}
	return callbacks
}

func (s *VolcengineAlbService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	// 检查是否需要变更ALB实例网络类型
	if resourceData.HasChange("type") {
		callback := s.ModifyLoadBalancerType(resourceData, resource)
		callbacks = append(callbacks, callback)
	}

	// 检查是否需要变更可用区
	if resourceData.HasChange("subnet_ids") {
		callback := s.ModifyLoadBalancerZones(resourceData, resource)
		callbacks = append(callbacks, callback)
	}

	// 更新基本属性
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyLoadBalancerAttributes",
			ConvertMode: ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"load_balancer_name": {
					TargetField: "LoadBalancerName",
				},
				"description": {
					TargetField: "Description",
				},
				"delete_protection": {
					TargetField: "DeleteProtection",
				},
				"modification_protection_status": {
					TargetField: "ModificationProtectionStatus",
				},
				"modification_protection_reason": {
					TargetField: "ModificationProtectionReason",
				},
				"waf_protection_enabled": {
					TargetField: "WafProtectionEnabled",
				},
				"waf_instance_id": {
					TargetField: "WafInstanceId",
				},
				"waf_protected_domain": {
					TargetField: "WafProtectedDomain",
				},
				"proxy_protocol_enabled": {
					TargetField: "ProxyProtocolEnabled",
				},
				"global_accelerator": {
					TargetField: "GlobalAccelerator",
					ConvertType: ve.ConvertListUnique,
					NextLevelConvert: map[string]ve.RequestConvert{
						"accelerator_id": {
							TargetField: "AcceleratorId",
						},
						"accelerator_listener_id": {
							TargetField: "AcceleratorListenerId",
						},
						"endpoint_group_id": {
							TargetField: "EndpointGroupId",
						},
						"weight": {
							TargetField: "Weight",
						},
					},
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if len(*call.SdkParam) > 0 {
					(*call.SdkParam)["LoadBalancerId"] = d.Id()
					return true, nil
				}
				return false, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Active"},
				Timeout: resourceData.Timeout(schema.TimeoutUpdate),
			},
		},
	}
	callbacks = append(callbacks, callback)

	// 更新Tags
	setResourceTagsCallbacks := ve.SetResourceTags(s.Client, "TagResources", "UntagResources", "loadbalancer", resourceData, getUniversalInfo)
	callbacks = append(callbacks, setResourceTagsCallbacks...)

	return callbacks
}

// ModifyLoadBalancerType API实现 - 变更ALB实例网络类型
func (s *VolcengineAlbService) ModifyLoadBalancerType(resourceData *schema.ResourceData, resource *schema.Resource) ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyLoadBalancerType",
			ConvertMode: ve.RequestConvertIgnore,
			Convert: map[string]ve.RequestConvert{
				"subnet_ids": {
					Ignore: true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				// IPv4&IPv6双栈ALB实例不支持变更网络类型。
				if d.Get("address_ip_version").(string) == "DualStack" {
					return false, fmt.Errorf("DualStack ALB instances do not support changing the network type.")
				}

				_, newType := d.GetChange("type")
				(*call.SdkParam)["LoadBalancerId"] = d.Id()
				(*call.SdkParam)["Type"] = newType

				// 如果是public类型，需要处理EIP相关配置
				if newType == "public" {
					subnetIds, ok := d.Get("subnet_ids").(*schema.Set)
					var publicIps *schema.Set
					if allocationIDs, exist := d.Get("allocation_ids").(*schema.Set); !exist {
						return false, fmt.Errorf("allocation_ids is required for public type")
					} else {
						publicIps = allocationIDs
					}
					if subnetIds.Len() != publicIps.Len() {
						return false, fmt.Errorf("subnet_ids and allocation_id must have the same length")
					}
					if ok && subnetIds.Len() > 0 {
						_, zoneIds, err := s.getVpcIdAndZoneIdBySubnets(subnetIds.List())
						if err != nil {
							return false, err
						}
						publicIPsSlice := publicIps.List()
						for index, subnetId := range subnetIds.List() {
							(*call.SdkParam)["ZoneMappings."+strconv.Itoa(index+1)+".ZoneId"] = zoneIds[subnetId.(string)]
							(*call.SdkParam)["ZoneMappings."+strconv.Itoa(index+1)+".AllocationId"] = publicIPsSlice[index].(string)
						}
					}
				}

				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Active"},
				Timeout: resourceData.Timeout(schema.TimeoutUpdate),
			},
		},
	}
	return callback
}

// ModifyLoadBalancerZones API实现 - 变更ALB实例可用区
func (s *VolcengineAlbService) ModifyLoadBalancerZones(resourceData *schema.ResourceData, resource *schema.Resource) ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyLoadBalancerZones",
			ConvertMode: ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"subnet_ids": {
					Ignore: true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["LoadBalancerId"] = d.Id()

				subnetIds, ok := d.Get("subnet_ids").(*schema.Set)
				if !ok || subnetIds.Len() == 0 {
					return false, fmt.Errorf("subnet_ids is required for zone modification")
				}

				_, zoneIds, err := s.getVpcIdAndZoneIdBySubnets(subnetIds.List())
				if err != nil {
					return false, err
				}
				if subnetIds.Len() != len(zoneIds) {
					return false, fmt.Errorf("The length of subnetIds and zoneIds are not equal")
				}

				for index, subnetId := range subnetIds.List() {
					(*call.SdkParam)["ZoneMappings."+strconv.Itoa(index+1)+".SubnetId"] = subnetId.(string)
					(*call.SdkParam)["ZoneMappings."+strconv.Itoa(index+1)+".ZoneId"] = zoneIds[subnetId.(string)]
				}

				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Active"},
				Timeout: resourceData.Timeout(schema.TimeoutUpdate),
			},
		},
	}
	return callback
}

func (s *VolcengineAlbService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteLoadBalancer",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"LoadBalancerId": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 5*time.Minute)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				if protection, ok := d.Get("delete_protection").(string); ok && protection == "on" {
					// 开启删除保护，直接返回接口报错
					return baseErr
				}

				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on  reading alb on delete %q, %w", d.Id(), callErr))
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

// CloneLoadBalancer API实现
func (s *VolcengineAlbService) CreateCloneResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CloneLoadBalancer",
			ConvertMode: ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"source_load_balancer_id": {
					TargetField: "LoadBalancerId",
				},
				"load_balancer_name": {
					TargetField: "LoadBalancerName",
				},
				"description": {
					TargetField: "Description",
				},
				"subnet_ids": {
					Ignore: true,
				},
				"delete_protection": {
					TargetField: "DeleteProtection",
				},
				"project_name": {
					TargetField: "ProjectName",
				},
				"eip_billing_config": {
					TargetField: "EipBillingConfig",
					ConvertType: ve.ConvertListUnique,
					NextLevelConvert: map[string]ve.RequestConvert{
						"isp": {
							TargetField: "ISP",
						},
					},
				},
				"ipv6_eip_billing_config": {
					TargetField: "Ipv6EipBillingConfig",
					ConvertType: ve.ConvertListUnique,
					NextLevelConvert: map[string]ve.RequestConvert{
						"isp": {
							TargetField: "ISP",
						},
					},
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				sourceId := d.Get("source_load_balancer_id").(string)
				if sourceId == "" {
					return false, fmt.Errorf("source_load_balancer_id is required for cloning")
				}

				(*call.SdkParam)["RegionId"] = s.Client.Region
				subnetIds, ok := d.Get("subnet_ids").(*schema.Set)
				if ok && subnetIds.Len() > 0 {
					_, zoneIds, err := s.getVpcIdAndZoneIdBySubnets(subnetIds.List())
					if err != nil {
						return false, err
					}
					if subnetIds.Len() != len(zoneIds) {
						return false, fmt.Errorf("The length of subnetIds and zoneIds are not equal")
					}
					for index, subnetId := range subnetIds.List() {
						(*call.SdkParam)["ZoneMappings."+strconv.Itoa(index+1)+".SubnetId"] = subnetId.(string)
						(*call.SdkParam)["ZoneMappings."+strconv.Itoa(index+1)+".ZoneId"] = zoneIds[subnetId.(string)]
					}
				}

				if (*call.SdkParam)["Type"] == "private" {
					delete(*call.SdkParam, "EipBillingConfig.ISP")
					delete(*call.SdkParam, "EipBillingConfig.EipBillingType")
					delete(*call.SdkParam, "EipBillingConfig.Bandwidth")
					delete(*call.SdkParam, "Ipv6EipBillingConfig.ISP")
					delete(*call.SdkParam, "Ipv6EipBillingConfig.BillingType")
					delete(*call.SdkParam, "Ipv6EipBillingConfig.Bandwidth")
				}

				if eipBillingType, exist := (*call.SdkParam)["EipBillingConfig.EipBillingType"]; exist {
					ty := 0
					switch eipBillingType.(string) {
					case "PostPaidByBandwidth":
						ty = 2
					case "PostPaidByTraffic":
						ty = 3
					}
					(*call.SdkParam)["EipBillingConfig.EipBillingType"] = ty
				}
				if ipv6BillingType, exist := (*call.SdkParam)["Ipv6EipBillingConfig.BillingType"]; exist {
					ty := 0
					switch ipv6BillingType.(string) {
					case "PostPaidByBandwidth":
						ty = 2
					case "PostPaidByTraffic":
						ty = 3
					}
					(*call.SdkParam)["Ipv6EipBillingConfig.BillingType"] = ty
				}

				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("Result.LoadBalancerId", *resp)
				d.SetId(id.(string))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Active"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	callbacks = append(callbacks, callback)

	// 检查创建后是否需要调用ModifyLoadBalancerAttributes来配置不支持的参数
	if s.hasPostCreationAttributes(resourceData) {
		modifyCallback := s.createPostCreationModifyCallback(resourceData)
		callbacks = append(callbacks, modifyCallback)
	}

	return callbacks
}

// 检查是否有创建后需要配置的属性
func (s *VolcengineAlbService) hasPostCreationAttributes(d *schema.ResourceData) bool {
	if wafEnabled, ok := d.GetOk("waf_protection_enabled"); ok && wafEnabled != "off" {
		return true
	}
	if wafInstanceId, ok := d.GetOk("waf_instance_id"); ok && wafInstanceId != "" {
		return true
	}
	if wafDomain, ok := d.GetOk("waf_protected_domain"); ok && wafDomain != "" {
		return true
	}
	if globalAccel, ok := d.GetOk("global_accelerator"); ok && globalAccel != nil {
		if globalAccelList, ok := globalAccel.([]interface{}); ok && len(globalAccelList) > 0 {
			return true
		}
	}
	if proxyProtocol, ok := d.GetOk("proxy_protocol_enabled"); ok && proxyProtocol != "off" {
		return true
	}
	return false
}

// 创建实例后的Modify回调
func (s *VolcengineAlbService) createPostCreationModifyCallback(d *schema.ResourceData) ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyLoadBalancerAttributes",
			ConvertMode: ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"waf_protection_enabled": {
					TargetField: "WafProtectionEnabled",
				},
				"waf_instance_id": {
					TargetField: "WafInstanceId",
				},
				"waf_protected_domain": {
					TargetField: "WafProtectedDomain",
				},
				"global_accelerator": {
					TargetField: "GlobalAccelerator",
					ConvertType: ve.ConvertListUnique,
					NextLevelConvert: map[string]ve.RequestConvert{
						"accelerator_id": {
							TargetField: "AcceleratorId",
						},
						"accelerator_listener_id": {
							TargetField: "AcceleratorListenerId",
						},
						"endpoint_group_id": {
							TargetField: "EndpointGroupId",
						},
						"weight": {
							TargetField: "Weight",
						},
					},
				},
				"proxy_protocol_enabled": {
					TargetField: "ProxyProtocolEnabled",
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				// 设置LoadBalancerId
				(*call.SdkParam)["LoadBalancerId"] = d.Id()
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Active"},
				Timeout: d.Timeout(schema.TimeoutCreate),
			},
		},
	}
	return callback
}

func (s *VolcengineAlbService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "LoadBalancerIds",
				ConvertType: ve.ConvertWithN,
			},
			"project": {
				TargetField: "ProjectName",
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
		NameField:    "LoadBalancerName",
		IdField:      "LoadBalancerId",
		CollectField: "albs",
		ResponseConverts: map[string]ve.ResponseConvert{
			"LoadBalancerId": {
				TargetField: "id",
				KeepDefault: true,
			},
			"ISP": {
				TargetField: "isp",
			},
			"EipBillingType": {
				TargetField: "eip_billing_type",
				Convert: func(i interface{}) interface{} {
					if i == nil {
						return nil
					}
					billingType := i.(float64)
					switch billingType {
					case 2:
						return "PostPaidByBandwidth"
					case 3:
						return "PostPaidByTraffic"
					}
					return ""
				},
			},
			"BillingType": {
				TargetField: "billing_type",
				Convert: func(i interface{}) interface{} {
					if i == nil {
						return nil
					}
					billingType := i.(float64)
					switch billingType {
					case 2:
						return "PostPaidByBandwidth"
					case 3:
						return "PostPaidByTraffic"
					}
					return ""
				},
			},
		},
	}
}

func (s *VolcengineAlbService) ReadResourceId(id string) string {
	return id
}

func (s *VolcengineAlbService) ProjectTrn() *ve.ProjectTrn {
	return &ve.ProjectTrn{
		ServiceName:          "alb",
		ResourceType:         "loadbalancer",
		ProjectResponseField: "ProjectName",
		ProjectSchemaField:   "project_name",
	}
}

func (s *VolcengineAlbService) getVpcIdAndZoneIdBySubnets(subnetIds []interface{}) (string, map[string]string, error) {
	// describe subnet
	req := make(map[string]interface{})
	for index, subnetId := range subnetIds {
		req["SubnetIds."+strconv.Itoa(index+1)] = subnetId.(string)
	}
	action := "DescribeSubnets"
	resp, err := s.Client.UniversalClient.DoCall(getVpcUniversalInfo(action), &req)
	if err != nil {
		return "", map[string]string{}, err
	}
	logger.Debug(logger.RespFormat, action, req, *resp)
	results, err := ve.ObtainSdkValue("Result.Subnets", *resp)
	if err != nil {
		return "", map[string]string{}, err
	}
	if results == nil {
		results = []interface{}{}
	}
	subnets, ok := results.([]interface{})
	if !ok {
		return "", map[string]string{}, errors.New("Result.Subnets is not Slice")
	}
	if len(subnets) == 0 {
		return "", map[string]string{}, fmt.Errorf("subnets %v not exist", subnetIds)
	}
	zoneIds := make(map[string]string, 0)
	for _, v := range subnets {
		subnet, ok := v.(map[string]interface{})
		if !ok {
			return "", map[string]string{}, fmt.Errorf("Result.Subnet is not map")
		}
		zoneIds[subnet["SubnetId"].(string)] = subnet["ZoneId"].(string)
	}
	vpcId := subnets[0].(map[string]interface{})["VpcId"].(string)
	return vpcId, zoneIds, nil
}

func removeSystemTags(data []interface{}) ([]interface{}, error) {
	var (
		ok      bool
		result  map[string]interface{}
		results []interface{}
		tags    []interface{}
	)
	for _, d := range data {
		if result, ok = d.(map[string]interface{}); !ok {
			return results, errors.New("The elements in data are not map ")
		}
		tags, ok = result["Tags"].([]interface{})
		if ok {
			tags = ve.FilterSystemTags(tags)
			result["Tags"] = tags
		}
		results = append(results, result)
	}
	return results, nil
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "alb",
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
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
