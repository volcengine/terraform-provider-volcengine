package network_load_balancer

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

type VolcengineNlbNetworkLoadBalancerService struct {
	Client *ve.SdkClient
}

func NewNlbNetworkLoadBalancerService(c *ve.SdkClient) *VolcengineNlbNetworkLoadBalancerService {
	return &VolcengineNlbNetworkLoadBalancerService{
		Client: c,
	}
}

func (s *VolcengineNlbNetworkLoadBalancerService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineNlbNetworkLoadBalancerService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithNextTokenQuery(m, "MaxResults", "NextToken", 100, nil, func(condition map[string]interface{}) ([]interface{}, string, error) {
		action := "DescribeNetworkLoadBalancers"

		bytes, _ := json.Marshal(condition)
		logger.Debug(logger.ReqFormat, action, string(bytes))
		if condition == nil {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), nil)
			if err != nil {
				return nil, "", err
			}
		} else {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &condition)
			if err != nil {
				return nil, "", err
			}
		}
		respBytes, _ := json.Marshal(resp)
		logger.Debug(logger.RespFormat, action, condition, string(respBytes))
		results, err = ve.ObtainSdkValue("Result.NetworkLoadBalancers", *resp)
		if err != nil || results == nil {
			results, err = ve.ObtainSdkValue("Result.LoadBalancers", *resp)
		}
		if err != nil {
			return nil, "", err
		}
		if results == nil {
			results = []interface{}{}
		}

		var nextTokenStr string
		nextToken, _ := ve.ObtainSdkValue("Result.NextToken", *resp)
		if nextToken != nil {
			if s, ok := nextToken.(string); ok {
				nextTokenStr = s
			}
		}

		if data, ok = results.([]interface{}); !ok {
			return nil, "", errors.New("Result.NetworkLoadBalancers or Result.LoadBalancers is not Slice")
		}
		return data, nextTokenStr, err
	})
}

func (s *VolcengineNlbNetworkLoadBalancerService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"LoadBalancerIds.1": id,
	}
	results, err := s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		if tempData, ok := v.(map[string]interface{}); ok {
			if lbId, ok := tempData["LoadBalancerId"].(string); ok && lbId == id {
				return tempData, nil
			}
		}
	}
	return nil, fmt.Errorf("resource not found")
}

func (s *VolcengineNlbNetworkLoadBalancerService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{"Creating", "Configuring", "Provisioning"},
		Target:     target,
		Refresh:    ve.ResourceStateRefreshFunc(resourceData, s.ReadResource, id),
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 5 * time.Second,
	}
}

func (s *VolcengineNlbNetworkLoadBalancerService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action: "CreateNetworkLoadBalancer",
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				param, err := ve.ResourceDateToRequest(d, resource, false, s.createRequestConvert(), ve.RequestConvertAll, ve.ContentTypeDefault)
				if err != nil {
					return nil, err
				}
				delete(param, "Ipv4BandwidthPackageId")
				delete(param, "Ipv6BandwidthPackageId")
				logger.Debug(logger.ReqFormat, call.Action, param)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), &param)
				logger.Debug(logger.AllFormat, call.Action, param, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, err := ve.ObtainSdkValue("Result.LoadBalancerId", *resp)
				if err != nil {
					return err
				}
				if s, ok := id.(string); ok && s != "" {
					d.SetId(s)
				} else {
					return errors.New("Result.LoadBalancerId is not string")
				}
				_, err = s.RefreshResourceState(d, []string{"Active"}, 10*time.Minute, d.Id()).WaitForState()
				if err != nil {
					return err
				}
				return nil
			},
		},
	}
	callbacks := []ve.Callback{callback}
	setResourceTagsCallbacks := ve.SetResourceTags(s.Client, "TagNLBResources", "UntagNLBResources", "nlb", resourceData, getUniversalInfo)
	callbacks = append(callbacks, setResourceTagsCallbacks...)
	return callbacks
}

func (s *VolcengineNlbNetworkLoadBalancerService) createRequestConvert() map[string]ve.RequestConvert {
	return map[string]ve.RequestConvert{
		"load_balancer_name": {
			TargetField: "LoadBalancerName",
		},
		"description": {
			TargetField: "Description",
		},
		"vpc_id": {
			TargetField: "VpcId",
		},
		"region": {
			TargetField: "RegionId",
		},
		"network_type": {
			TargetField: "NetworkType",
		},
		"ip_address_version": {
			TargetField: "IpAddressVersion",
		},
		"cross_zone_enabled": {
			TargetField: "CrossZoneEnabled",
		},
		"project_name": {
			TargetField: "ProjectName",
		},
		"modification_protection_status": {
			TargetField: "ModificationProtectionStatus",
		},
		"modification_protection_reason": {
			TargetField: "ModificationProtectionReason",
		},
		"ipv4_bandwidth_package_id": {
			TargetField: "Ipv4BandwidthPackageId",
			Convert: func(d *schema.ResourceData, i interface{}) interface{} {
				return i
			},
		},
		"ipv6_bandwidth_package_id": {
			TargetField: "Ipv6BandwidthPackageId",
			Convert: func(d *schema.ResourceData, i interface{}) interface{} {
				return i
			},
		},
		"security_group_ids": {
			TargetField: "SecurityGroupIds",
			ConvertType: ve.ConvertWithN,
		},
		"zone_mappings": {
			TargetField: "ZoneMappings",
			ConvertType: ve.ConvertListN,
			NextLevelConvert: map[string]ve.RequestConvert{
				"zone_id": {
					TargetField: "ZoneId",
				},
				"subnet_id": {
					TargetField: "SubnetId",
				},
				"ipv4_address": {
					TargetField: "Ipv4Address",
				},
				"ipv6_address": {
					TargetField: "Ipv6Address",
				},
				"eip_id": {
					TargetField: "EipId",
				},
			},
		},
		"tags": {
			TargetField: "Tags",
			ConvertType: ve.ConvertListN,
			Convert:     transTagsToRequest,
		},
	}
}

func transTagsToRequest(d *schema.ResourceData, i interface{}) interface{} {
	if i == nil {
		return nil
	}
	var tags []map[string]interface{}
	if set, ok := i.(*schema.Set); ok {
		for _, v := range set.List() {
			if m, ok := v.(map[string]interface{}); ok {
				tag := make(map[string]interface{})
				if key, ok := m["key"].(string); ok {
					tag["Key"] = key
				}
				if value, ok := m["value"].(string); ok {
					tag["Value"] = value
				}
				tags = append(tags, tag)
			}
		}
	}
	return tags
}

func transZoneMappingsToResponse(i interface{}) interface{} {
	if i == nil {
		return nil
	}
	list, ok := i.([]interface{})
	if !ok {
		return i
	}
	var res []interface{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		if !ok {
			res = append(res, item)
			continue
		}
		newMap := make(map[string]interface{})
		for k, v := range m {
			newKey := ve.HumpToDownLine(k)
			newMap[newKey] = v
			if k == "EipId" {
				newMap["ipv4_eip_id"] = v
				newMap["eip_id"] = v
			}
			if k == "EipAddress" {
				newMap["ipv4_eip_address"] = v
			}
		}
		res = append(res, newMap)
	}
	return res
}

func (VolcengineNlbNetworkLoadBalancerService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"LoadBalancerId": {
				TargetField: "load_balancer_id",
			},
			"LoadBalancerName": {
				TargetField: "load_balancer_name",
			},
			"Description": {
				TargetField: "description",
			},
			"VpcId": {
				TargetField: "vpc_id",
			},
			"RegionId": {
				TargetField: "region",
			},
			"ProjectName": {
				TargetField: "project_name",
			},
			"Ipv4NetworkType": {
				TargetField: "ipv4_network_type",
			},
			"NetworkType": {
				TargetField: "network_type",
			},
			"Ipv6NetworkType": {
				TargetField: "ipv6_network_type",
			},
			"IpAddressVersion": {
				TargetField: "ip_address_version",
			},
			"CrossZoneEnabled": {
				TargetField: "cross_zone_enabled",
			},
			"SecurityGroupIds": {
				TargetField: "security_group_ids",
			},
			"Ipv4BandwidthPackageId": {
				TargetField: "ipv4_bandwidth_package_id",
			},
			"Ipv6BandwidthPackageId": {
				TargetField: "ipv6_bandwidth_package_id",
			},
			"ModificationProtectionStatus": {
				TargetField: "modification_protection_status",
			},
			"ModificationProtectionReason": {
				TargetField: "modification_protection_reason",
			},
			"AccessLog": {
				TargetField: "access_log_config",
			},
			"ZoneMappings": {
				TargetField: "zone_mappings",
				Convert:     transZoneMappingsToResponse,
			},
			"Tags": {
				TargetField: "tags",
				Convert:     ve.TransTagsToResponse,
			},
			"Status": {
				TargetField: "status",
			},
			"DNSName": {
				TargetField: "dns_name",
			},
			"CreateTime": {
				TargetField: "create_time",
			},
			"UpdateTime": {
				TargetField: "update_time",
			},
			"AccountId": {
				TargetField: "account_id",
			},
			"BillingStatus": {
				TargetField: "billing_status",
			},
			"OverdueTime": {
				TargetField: "overdue_time",
			},
			"ReclaimedTime": {
				TargetField: "reclaimed_time",
			},
			"ExpectedOverdueTime": {
				TargetField: "expected_overdue_time",
			},
			"ManagedSecurityGroupId": {
				TargetField: "managed_security_group_id",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineNlbNetworkLoadBalancerService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	if resourceData.HasChanges("load_balancer_name", "description", "cross_zone_enabled", "modification_protection_status", "modification_protection_reason") {
		callbacks = append(callbacks, ve.Callback{
			Call: ve.SdkCall{
				Action: "ModifyNetworkLoadBalancerAttributes",
				SdkParam: &map[string]interface{}{
					"LoadBalancerId":               resourceData.Id(),
					"LoadBalancerName":             resourceData.Get("load_balancer_name"),
					"Description":                  resourceData.Get("description"),
					"CrossZoneEnabled":             resourceData.Get("cross_zone_enabled"),
					"ModificationProtectionStatus": resourceData.Get("modification_protection_status"),
					"ModificationProtectionReason": resourceData.Get("modification_protection_reason"),
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
			},
		})
	}

	if resourceData.HasChange("security_group_ids") {
		callbacks = append(callbacks, ve.Callback{
			Call: ve.SdkCall{
				Action: "ModifyNetworkLoadBalancerSecurityGroups",
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					param, err := ve.ResourceDateToRequest(d, resource, true, map[string]ve.RequestConvert{
						"security_group_ids": s.createRequestConvert()["security_group_ids"],
					}, ve.RequestConvertInConvert, ve.ContentTypeDefault)
					if err != nil {
						return nil, err
					}
					param["LoadBalancerId"] = d.Id()
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), &param)
				},
			},
		})
	}

	if resourceData.HasChanges("ipv4_bandwidth_package_id", "ipv6_bandwidth_package_id") {
		callbacks = append(callbacks, ve.Callback{
			Call: ve.SdkCall{
				Action: "ModifyNetworkLoadBalancerBandwidthPackage",
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					param, err := ve.ResourceDateToRequest(d, resource, true, map[string]ve.RequestConvert{
						"ipv4_bandwidth_package_id": s.createRequestConvert()["ipv4_bandwidth_package_id"],
						"ipv6_bandwidth_package_id": s.createRequestConvert()["ipv6_bandwidth_package_id"],
					}, ve.RequestConvertInConvert, ve.ContentTypeDefault)
					if err != nil {
						return nil, err
					}
					if v, ok := param["Ipv4BandwidthPackageId"]; ok {
						param["BandwidthPackageId"] = v
						delete(param, "Ipv4BandwidthPackageId")
						param["Protocol"] = "ipv4"
					}
					if v, ok := param["Ipv6BandwidthPackageId"]; ok {
						param["BandwidthPackageId"] = v
						delete(param, "Ipv6BandwidthPackageId")
						param["Protocol"] = "ipv6"
					}
					if _, ok := param["BandwidthPackageId"]; !ok {
						if d.HasChange("ipv4_bandwidth_package_id") {
							param["BandwidthPackageId"] = ""
							param["Protocol"] = "ipv4"
						} else if d.HasChange("ipv6_bandwidth_package_id") {
							param["BandwidthPackageId"] = ""
							param["Protocol"] = "ipv6"
						}
					}

					if v, ok := param["BandwidthPackageId"]; ok && v != "" {
						if d.HasChange("ipv4_bandwidth_package_id") {
							o, _ := d.GetChange("ipv4_bandwidth_package_id")
							if val, ok := o.(string); ok && val != "" {
								removeParam := map[string]interface{}{
									"LoadBalancerId":     d.Id(),
									"Operation":          "remove",
									"Protocol":           "ipv4",
									"BandwidthPackageId": val,
								}
								_, err = s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), &removeParam)
								if err != nil {
									return nil, err
								}
								time.Sleep(2 * time.Second)
							}
						}
						if d.HasChange("ipv6_bandwidth_package_id") {
							o, _ := d.GetChange("ipv6_bandwidth_package_id")
							if val, ok := o.(string); ok && val != "" {
								removeParam := map[string]interface{}{
									"LoadBalancerId":     d.Id(),
									"Operation":          "remove",
									"Protocol":           "ipv6",
									"BandwidthPackageId": val,
								}
								_, err = s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), &removeParam)
								if err != nil {
									return nil, err
								}
							}
						}

						param["Operation"] = "add"
					} else {
						param["Operation"] = "remove"
						if d.HasChange("ipv4_bandwidth_package_id") {
							o, _ := d.GetChange("ipv4_bandwidth_package_id")
							if val, ok := o.(string); ok && val != "" {
								param["BandwidthPackageId"] = val
								param["Protocol"] = "ipv4"
							}
						} else if d.HasChange("ipv6_bandwidth_package_id") {
							o, _ := d.GetChange("ipv6_bandwidth_package_id")
							if val, ok := o.(string); ok && val != "" {
								param["BandwidthPackageId"] = val
								param["Protocol"] = "ipv6"
							}
						}
					}
					param["LoadBalancerId"] = d.Id()
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), &param)
				},
			},
		})
	}

	if resourceData.HasChange("access_log_config") {
		v := resourceData.Get("access_log_config")
		config, ok := v.([]interface{})
		if !ok {
			return []ve.Callback{{Err: errors.New("access_log_config is not []interface{}")}}
		}
		param := map[string]interface{}{
			"LoadBalancerId": resourceData.Id(),
		}
		if len(config) > 0 {
			c, ok := config[0].(map[string]interface{})
			if ok {
				param["AccessLogEnabled"] = c["enabled"]
				param["ProjectName"] = c["project_name"]
				param["TopicName"] = c["topic_name"]
			} else {
				param["AccessLogEnabled"] = false
			}
		} else {
			param["AccessLogEnabled"] = false
		}
		callbacks = append(callbacks, ve.Callback{
			Call: ve.SdkCall{
				Action: "ModifyNetworkLoadBalancerAccessLog",
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), &param)
				},
			},
		})
	}

	if resourceData.HasChange("network_type") {
		callbacks = append(callbacks, ve.Callback{
			Call: ve.SdkCall{
				Action: "ModifyNetworkLoadBalancerNetworkType",
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					param, err := ve.ResourceDateToRequest(d, resource, true, map[string]ve.RequestConvert{
						"network_type": s.createRequestConvert()["network_type"],
					}, ve.RequestConvertInConvert, ve.ContentTypeDefault)
					if err != nil {
						return nil, err
					}
					param["LoadBalancerId"] = d.Id()
					param["Protocol"] = d.Get("ip_address_version")
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), &param)
				},
			},
		})
	}

	if resourceData.HasChange("zone_mappings") {
		callbacks = append(callbacks, ve.Callback{
			Call: ve.SdkCall{
				Action: "ModifyNetworkLoadBalancerZones",
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					param := make(map[string]interface{})
					param["LoadBalancerId"] = d.Id()
					oldZ, newZ := d.GetChange("zone_mappings")
					if oldZ == nil {
						oldZ = []interface{}{}
					}
					if newZ == nil {
						newZ = []interface{}{}
					}
					oldZList, ok := oldZ.([]interface{})
					if !ok {
						return nil, errors.New("old zone_mappings is not []interface{}")
					}
					newZList, ok := newZ.([]interface{})
					if !ok {
						return nil, errors.New("new zone_mappings is not []interface{}")
					}
					addZList := make([]interface{}, 0)
					delZList := make([]interface{}, 0)
					for _, n := range newZList {
						nm, ok := n.(map[string]interface{})
						if !ok {
							return nil, errors.New("new zone_mappings item is not map")
						}
						exist := false
						for _, o := range oldZList {
							om, ok := o.(map[string]interface{})
							if !ok {
								return nil, errors.New("old zone_mappings item is not map")
							}
							if nm["zone_id"] == om["zone_id"] {
								exist = true
								break
							}
						}
						if !exist {
							addZList = append(addZList, n)
						}
					}
					for _, o := range oldZList {
						om, ok := o.(map[string]interface{})
						if !ok {
							return nil, errors.New("old zone_mappings item is not map")
						}
						exist := false
						for _, n := range newZList {
							nm, ok := n.(map[string]interface{})
							if !ok {
								return nil, errors.New("new zone_mappings item is not map")
							}
							if nm["zone_id"] == om["zone_id"] {
								exist = true
								break
							}
						}
						if !exist {
							delZList = append(delZList, o)
						}
					}
					if len(addZList) > 0 {
						addReq := make(map[string]interface{})
						addSchema := s.createRequestConvert()["zone_mappings"]
						addSchema.ConvertType = ve.ConvertListN
						ve.Convert(d, "zone_mappings", addZList, addSchema, 0, &addReq, "", false, ve.ContentTypeDefault, "", nil)
						for k, v := range addReq {
							param["Add"+k] = v
						}
					}
					if len(delZList) > 0 {
						for i, v := range delZList {
							vm, ok := v.(map[string]interface{})
							if ok {
								param[fmt.Sprintf("DeleteZoneMappings.%d", i+1)] = vm["zone_id"]
							}
						}
					}
					if len(addZList) > 0 || len(delZList) > 0 {
						return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), &param)
					}
					return nil, nil
				},
			},
		})
	}

	setResourceTagsCallbacks := ve.SetResourceTags(s.Client, "TagNLBResources", "UntagNLBResources", "nlb", resourceData, getUniversalInfo)
	callbacks = append(callbacks, setResourceTagsCallbacks...)

	return callbacks
}

func (s *VolcengineNlbNetworkLoadBalancerService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteNetworkLoadBalancer",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"LoadBalancerId": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 10*time.Minute)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineNlbNetworkLoadBalancerService) DatasourceResources(d *schema.ResourceData, r *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"load_balancer_ids": {
				TargetField: "LoadBalancerIds",
				ConvertType: ve.ConvertWithN,
			},
			"vpc_id": {
				TargetField: "VpcId",
			},
			"load_balancer_name": {
				TargetField: "LoadBalancerName",
			},
			"status": {
				TargetField: "Status",
			},
			"ip_address_version": {
				TargetField: "IpAddressVersion",
			},
			"zone_id": {
				TargetField: "ZoneId",
			},
			"project_name": {
				TargetField: "ProjectName",
			},
			"tags": {
				TargetField: "TagFilters",
				Convert:     ve.TransTagFiltersToRequest,
			},
		},
		NameField:    "LoadBalancerName",
		IdField:      "LoadBalancerId",
		CollectField: "network_load_balancers",
		ResponseConverts: map[string]ve.ResponseConvert{
			"LoadBalancerId": {
				TargetField: "load_balancer_id",
			},
			"LoadBalancerName": {
				TargetField: "load_balancer_name",
			},
			"RegionId": {
				TargetField: "region",
			},
			"Description": {
				TargetField: "description",
			},
			"Status": {
				TargetField: "status",
			},
			"CreateTime": {
				TargetField: "create_time",
			},
			"UpdateTime": {
				TargetField: "update_time",
			},
			"VpcId": {
				TargetField: "vpc_id",
			},
			"ProjectName": {
				TargetField: "project_name",
			},
			"Ipv4NetworkType": {
				TargetField: "ipv4_network_type",
			},
			"NetworkType": {
				TargetField: "network_type",
			},
			"Ipv6NetworkType": {
				TargetField: "ipv6_network_type",
			},
			"IpAddressVersion": {
				TargetField: "ip_address_version",
			},
			"CrossZoneEnabled": {
				TargetField: "cross_zone_enabled",
			},
			"DNSName": {
				TargetField: "dns_name",
			},
			"AccountId": {
				TargetField: "account_id",
			},
			"BillingStatus": {
				TargetField: "billing_status",
			},
			"OverdueTime": {
				TargetField: "overdue_time",
			},
			"ReclaimedTime": {
				TargetField: "reclaimed_time",
			},
			"ExpectedOverdueTime": {
				TargetField: "expected_overdue_time",
			},
			"Ipv4BandwidthPackageId": {
				TargetField: "ipv4_bandwidth_package_id",
			},
			"Ipv6BandwidthPackageId": {
				TargetField: "ipv6_bandwidth_package_id",
			},
			"ModificationProtectionStatus": {
				TargetField: "modification_protection_status",
			},
			"ModificationProtectionReason": {
				TargetField: "modification_protection_reason",
			},
			"ManagedSecurityGroupId": {
				TargetField: "managed_security_group_id",
			},
			"SecurityGroupIds": {
				TargetField: "security_group_ids",
			},
			"ZoneMappings": {
				TargetField: "zone_mappings",
				Convert:     transZoneMappingsToResponse,
			},
			"Tags": {
				TargetField: "tags",
				Convert:     ve.TransTagsToResponse,
			},
		},
	}
}

func (s *VolcengineNlbNetworkLoadBalancerService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "clb",
		Action:      actionName,
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
	}
}
