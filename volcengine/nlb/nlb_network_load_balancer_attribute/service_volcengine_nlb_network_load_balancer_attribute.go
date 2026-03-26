package network_load_balancer_attribute

import (
	"encoding/json"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineNlbNetworkLoadBalancerAttributeService struct {
	Client *ve.SdkClient
}

func NewNlbNetworkLoadBalancerAttributeService(c *ve.SdkClient) *VolcengineNlbNetworkLoadBalancerAttributeService {
	return &VolcengineNlbNetworkLoadBalancerAttributeService{
		Client: c,
	}
}

func (s *VolcengineNlbNetworkLoadBalancerAttributeService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineNlbNetworkLoadBalancerAttributeService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	action := "DescribeNetworkLoadBalancerAttributes"
	logger.Debug(logger.ReqFormat, action, m)
	resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(action), &m)
	if err != nil {
		return nil, err
	}
	respBytes, _ := json.Marshal(resp)
	logger.Debug(logger.RespFormat, action, m, string(respBytes))

	// The API returns a single object in Result, we wrap it in a slice for the dispatcher
	if resp != nil {
		if result, ok := (*resp)["Result"]; ok {
			data = append(data, result)
		}
	}
	return data, nil
}

func (s *VolcengineNlbNetworkLoadBalancerAttributeService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	return nil, nil
}

func (s *VolcengineNlbNetworkLoadBalancerAttributeService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineNlbNetworkLoadBalancerAttributeService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return nil
}

func (VolcengineNlbNetworkLoadBalancerAttributeService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	return nil
}

func (s *VolcengineNlbNetworkLoadBalancerAttributeService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return nil
}

func (s *VolcengineNlbNetworkLoadBalancerAttributeService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return nil
}

func (s *VolcengineNlbNetworkLoadBalancerAttributeService) DatasourceResources(d *schema.ResourceData, r *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"load_balancer_id": {
				TargetField: "LoadBalancerId",
			},
		},
		IdField:      "LoadBalancerId",
		CollectField: "network_load_balancer_attributes",
		ResponseConverts: map[string]ve.ResponseConvert{
			"LoadBalancerId": {
				TargetField: "load_balancer_id",
			},
			"LoadBalancerName": {
				TargetField: "load_balancer_name",
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
			"BillingType": {
				TargetField: "billing_type",
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
			"RegionId": {
				TargetField: "region",
			},
			"NetworkType": {
				TargetField: "network_type",
			},
		},
	}
}

func transZoneMappingsToResponse(v interface{}) interface{} {
	if v == nil {
		return nil
	}
	var zoneMappings []map[string]interface{}
	if list, ok := v.([]interface{}); ok {
		for _, item := range list {
			if m, ok := item.(map[string]interface{}); ok {
				mapping := make(map[string]interface{})
				if val, ok := m["ZoneId"]; ok {
					mapping["zone_id"] = val
				}
				if val, ok := m["SubnetId"]; ok {
					mapping["subnet_id"] = val
				}
				if val, ok := m["LoadBalancerAddresses"]; ok {
					if addrs, ok := val.([]interface{}); ok && len(addrs) > 0 {
						for _, addr := range addrs {
							if addrMap, ok := addr.(map[string]interface{}); ok {
								if eipId, ok := addrMap["EipId"]; ok {
									mapping["ipv4_eip_id"] = eipId
								}
								if eipAddress, ok := addrMap["EipAddress"]; ok {
									mapping["ipv4_eip_address"] = eipAddress
								}
								if eniId, ok := addrMap["EniId"]; ok {
									mapping["eni_id"] = eniId
								}
								if eniAddress, ok := addrMap["EniAddress"]; ok {
									mapping["eni_address"] = eniAddress
								}
								if ipv6Address, ok := addrMap["Ipv6Address"]; ok {
									mapping["ipv6_address"] = ipv6Address
								}
								if ipv4Address, ok := addrMap["Ipv4Address"]; ok {
									mapping["ipv4_address"] = ipv4Address
								}
								if ipv4EipAddress, ok := addrMap["Ipv4EipAddress"]; ok {
									mapping["ipv4_eip_address"] = ipv4EipAddress
								}
								if ipv4HcStatus, ok := addrMap["Ipv4HcStatus"]; ok {
									mapping["ipv4_hc_status"] = ipv4HcStatus
								}
								if ipv4LocalAddresses, ok := addrMap["Ipv4LocalAddresses"]; ok {
									mapping["ipv4_local_addresses"] = ipv4LocalAddresses
								}
								if ipv6HcStatus, ok := addrMap["Ipv6HcStatus"]; ok {
									mapping["ipv6_hc_status"] = ipv6HcStatus
								}
								if ipv6LocalAddresses, ok := addrMap["Ipv6LocalAddresses"]; ok {
									mapping["ipv6_local_addresses"] = ipv6LocalAddresses
								}
							}
						}
					}
				}
				// 兼容扁平结构：如果 LoadBalancerAddresses 不存在或为空，尝试直接从 ZoneMappings 元素中读取字段
				// 某些 API 版本或特定场景下，网络信息可能直接位于 ZoneMappings 层级
				if _, ok := mapping["eni_id"]; !ok {
					if val, ok := m["EniId"]; ok {
						mapping["eni_id"] = val
					}
					if val, ok := m["EniAddress"]; ok {
						mapping["eni_address"] = val
					}
					if val, ok := m["Ipv6Address"]; ok {
						mapping["ipv6_address"] = val
					}
					if val, ok := m["Ipv4Address"]; ok {
						mapping["ipv4_address"] = val
					}
					// 尝试读取 EIP 相关信息
					if val, ok := m["Ipv4EipId"]; ok {
						mapping["ipv4_eip_id"] = val
					}
					if val, ok := m["Ipv4EipAddress"]; ok {
						mapping["ipv4_eip_address"] = val
					}
					// 兼容可能的旧字段名
					if val, ok := m["EipId"]; ok {
						mapping["ipv4_eip_id"] = val
					}
					if val, ok := m["EipAddress"]; ok {
						mapping["ipv4_eip_address"] = val
					}
					if val, ok := m["Ipv4HcStatus"]; ok {
						mapping["ipv4_hc_status"] = val
					}
					if val, ok := m["Ipv4LocalAddresses"]; ok {
						mapping["ipv4_local_addresses"] = val
					}
					if val, ok := m["Ipv6HcStatus"]; ok {
						mapping["ipv6_hc_status"] = val
					}
					if val, ok := m["Ipv6LocalAddresses"]; ok {
						mapping["ipv6_local_addresses"] = val
					}
				}
				zoneMappings = append(zoneMappings, mapping)
			}
		}
	}
	return zoneMappings
}

func (s *VolcengineNlbNetworkLoadBalancerAttributeService) ReadResourceId(id string) string {
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
