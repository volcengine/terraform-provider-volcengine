package network_interface

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

type VolcengineNetworkInterfaceService struct {
	Client *ve.SdkClient
}

func NewNetworkInterfaceService(c *ve.SdkClient) *VolcengineNetworkInterfaceService {
	return &VolcengineNetworkInterfaceService{
		Client: c,
	}
}

func (s *VolcengineNetworkInterfaceService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineNetworkInterfaceService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 20, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeNetworkInterfaces"
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
		logger.Debug(logger.RespFormat, action, *resp)
		results, err = ve.ObtainSdkValue("Result.NetworkInterfaceSets", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.NetworkInterfaceSets is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineNetworkInterfaceService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"NetworkInterfaceIds.1": id,
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		if data, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("value is not map")
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("network_interface %s not exist ", id)
	}
	privateIpAddress := make([]string, 0)
	if privateIpMap, ok := data["PrivateIpSets"].(map[string]interface{}); ok {
		if privateIpSets, ok := privateIpMap["PrivateIpSet"].([]interface{}); ok {
			for _, p := range privateIpSets {
				if pMap, ok := p.(map[string]interface{}); ok {
					isPrimary := pMap["Primary"].(bool)
					ip := pMap["PrivateIpAddress"].(string)
					if !isPrimary {
						privateIpAddress = append(privateIpAddress, ip)
					}
				}
			}
		}
	}
	data["PrivateIpAddress"] = privateIpAddress
	data["SecondaryPrivateIpAddressCount"] = len(privateIpAddress)

	if ipv6Sets, ok := data["IPv6Sets"].([]interface{}); ok {
		data["Ipv6Addresses"] = ipv6Sets
		data["Ipv6AddressCount"] = len(ipv6Sets)
	}

	return data, err
}

func (s *VolcengineNetworkInterfaceService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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
			status, err = ve.ObtainSdkValue("Status", demo)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("network_interface status error, status:%s", status.(string))
				}
			}
			//注意 返回的第一个参数不能为空 否则会一直等下去
			return demo, status.(string), err
		},
	}
}

func (VolcengineNetworkInterfaceService) WithResourceResponseHandlers(networkInterface map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return networkInterface, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}

}

func (s *VolcengineNetworkInterfaceService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateNetworkInterface",
			ConvertMode: ve.RequestConvertAll,
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.VpcClient.CreateNetworkInterfaceCommon(call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("Result.NetworkInterfaceId", *resp)
				d.SetId(id.(string))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Available"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
			Convert: map[string]ve.RequestConvert{
				"security_group_ids": {
					TargetField: "SecurityGroupIds",
					ConvertType: ve.ConvertWithN,
				},
				"private_ip_address": {
					TargetField: "PrivateIpAddress",
					ConvertType: ve.ConvertWithN,
				},
				"ipv6_addresses": {
					TargetField: "Ipv6Address",
					ConvertType: ve.ConvertWithN,
				},
				"tags": {
					TargetField: "Tags",
					ConvertType: ve.ConvertListN,
				},
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineNetworkInterfaceService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyNetworkInterfaceAttributes",
			ConvertMode: ve.RequestConvertAll,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["NetworkInterfaceId"] = d.Id()
				delete(*call.SdkParam, "Tags")
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.VpcClient.ModifyNetworkInterfaceAttributesCommon(call.SdkParam)
			},
			Convert: map[string]ve.RequestConvert{
				"security_group_ids": {
					TargetField: "SecurityGroupIds",
					ConvertType: ve.ConvertWithN,
				},
				"private_ip_address": {
					Ignore: true,
				},
				"secondary_private_ip_address_count": {
					Ignore: true,
				},
			},
		},
	}
	callbacks = append(callbacks, callback)

	// 检查private_ip_address改变
	if resourceData.HasChange("private_ip_address") {
		add, remove, _, _ := ve.GetSetDifference("private_ip_address", resourceData, schema.HashString, false)
		if remove.Len() > 0 {
			callback = ve.Callback{
				Call: ve.SdkCall{
					Action:      "UnassignPrivateIpAddresses",
					ConvertMode: ve.RequestConvertInConvert,
					BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
						(*call.SdkParam)["NetworkInterfaceId"] = d.Id()
						for index, r := range remove.List() {
							(*call.SdkParam)["PrivateIpAddress."+strconv.Itoa(index+1)] = r
						}
						return true, nil
					},
					Convert: map[string]ve.RequestConvert{
						"private_ip_address": {
							Ignore: true,
						},
						"secondary_private_ip_address_count": {
							Ignore: true,
						},
					},
					ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
						logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
						return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
					},
				},
			}
			callbacks = append(callbacks, callback)
		}
		if add.Len() > 0 {
			callback = ve.Callback{
				Call: ve.SdkCall{
					Action:      "AssignPrivateIpAddresses",
					ConvertMode: ve.RequestConvertInConvert,
					BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
						(*call.SdkParam)["NetworkInterfaceId"] = d.Id()
						for index, r := range add.List() {
							(*call.SdkParam)["PrivateIpAddress."+strconv.Itoa(index+1)] = r
						}
						return true, nil
					},
					Convert: map[string]ve.RequestConvert{
						"private_ip_address": {
							Ignore: true,
						},
						"secondary_private_ip_address_count": {
							Ignore: true,
						},
					},
					ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
						logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
						return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
					},
				},
			}
			callbacks = append(callbacks, callback)
		}
	}
	// 检查secondary_private_ip_address_count改变
	if resourceData.HasChange("secondary_private_ip_address_count") {
		privateIpAddress := resourceData.Get("private_ip_address").(*schema.Set).List()
		oldCount, newCount := resourceData.GetChange("secondary_private_ip_address_count")
		if oldCount != nil && newCount != nil && newCount != len(privateIpAddress) {
			diff := newCount.(int) - oldCount.(int)
			if diff > 0 {
				callback = ve.Callback{
					Call: ve.SdkCall{
						Action:      "AssignPrivateIpAddresses",
						ConvertMode: ve.RequestConvertInConvert,
						BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
							(*call.SdkParam)["NetworkInterfaceId"] = d.Id()
							(*call.SdkParam)["SecondaryPrivateIpAddressCount"] = diff
							return true, nil
						},
						Convert: map[string]ve.RequestConvert{
							"private_ip_address": {
								Ignore: true,
							},
							"secondary_private_ip_address_count": {
								Ignore: true,
							},
						},
						ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
							logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
							return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
						},
					},
				}
				callbacks = append(callbacks, callback)
			} else {
				diff *= -1
				removeIpAddress := privateIpAddress[:diff]
				callback = ve.Callback{
					Call: ve.SdkCall{
						Action:      "UnassignPrivateIpAddresses",
						ConvertMode: ve.RequestConvertInConvert,
						BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
							(*call.SdkParam)["NetworkInterfaceId"] = d.Id()
							for index, r := range removeIpAddress {
								(*call.SdkParam)["PrivateIpAddress."+strconv.Itoa(index+1)] = r
							}
							return true, nil
						},
						Convert: map[string]ve.RequestConvert{
							"private_ip_address": {
								Ignore: true,
							},
							"secondary_private_ip_address_count": {
								Ignore: true,
							},
						},
						ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
							logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
							return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
						},
					},
				}
				callbacks = append(callbacks, callback)
			}
		}
	}

	// 检查 ipv6_addresses 改变
	if resourceData.HasChange("ipv6_addresses") {
		add, remove, _, _ := ve.GetSetDifference("ipv6_addresses", resourceData, schema.HashString, false)
		if remove.Len() > 0 {
			callback = ve.Callback{
				Call: ve.SdkCall{
					Action:      "UnassignIpv6Addresses",
					ConvertMode: ve.RequestConvertInConvert,
					BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
						(*call.SdkParam)["NetworkInterfaceId"] = d.Id()
						for index, r := range remove.List() {
							(*call.SdkParam)["Ipv6Address."+strconv.Itoa(index+1)] = r
						}
						return true, nil
					},
					Convert: map[string]ve.RequestConvert{
						"ipv6_addresses": {
							Ignore: true,
						},
						"ipv6_address_count": {
							Ignore: true,
						},
					},
					ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
						logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
						return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
					},
				},
			}
			callbacks = append(callbacks, callback)
		}
		if add.Len() > 0 {
			callback = ve.Callback{
				Call: ve.SdkCall{
					Action:      "AssignIpv6Addresses",
					ConvertMode: ve.RequestConvertInConvert,
					BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
						(*call.SdkParam)["NetworkInterfaceId"] = d.Id()
						for index, r := range add.List() {
							(*call.SdkParam)["Ipv6Address."+strconv.Itoa(index+1)] = r
						}
						return true, nil
					},
					Convert: map[string]ve.RequestConvert{
						"ipv6_addresses": {
							Ignore: true,
						},
						"ipv6_address_count": {
							Ignore: true,
						},
					},
					ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
						logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
						return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
					},
				},
			}
			callbacks = append(callbacks, callback)
		}
	}
	// 检查 ipv6_address_count 改变
	if resourceData.HasChange("ipv6_address_count") {
		ipv6Addresses := resourceData.Get("ipv6_addresses").(*schema.Set).List()
		oldCount, newCount := resourceData.GetChange("ipv6_address_count")
		if oldCount != nil && newCount != nil && newCount != len(ipv6Addresses) {
			diff := newCount.(int) - oldCount.(int)
			if diff > 0 {
				callback = ve.Callback{
					Call: ve.SdkCall{
						Action:      "AssignIpv6Addresses",
						ConvertMode: ve.RequestConvertInConvert,
						BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
							(*call.SdkParam)["NetworkInterfaceId"] = d.Id()
							(*call.SdkParam)["Ipv6AddressCount"] = diff
							return true, nil
						},
						Convert: map[string]ve.RequestConvert{
							"ipv6_addresses": {
								Ignore: true,
							},
							"ipv6_address_count": {
								Ignore: true,
							},
						},
						ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
							logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
							return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
						},
					},
				}
				callbacks = append(callbacks, callback)
			} else {
				diff *= -1
				removeIpAddress := ipv6Addresses[:diff]
				callback = ve.Callback{
					Call: ve.SdkCall{
						Action:      "UnassignIpv6Addresses",
						ConvertMode: ve.RequestConvertInConvert,
						BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
							(*call.SdkParam)["NetworkInterfaceId"] = d.Id()
							for index, r := range removeIpAddress {
								(*call.SdkParam)["Ipv6Address."+strconv.Itoa(index+1)] = r
							}
							return true, nil
						},
						Convert: map[string]ve.RequestConvert{
							"ipv6_addresses": {
								Ignore: true,
							},
							"ipv6_address_count": {
								Ignore: true,
							},
						},
						ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
							logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
							return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
						},
					},
				}
				callbacks = append(callbacks, callback)
			}
		}
	}

	// 更新Tags
	setResourceTagsCallbacks := ve.SetResourceTags(s.Client, "TagResources", "UntagResources", "eni", resourceData, getUniversalInfo)
	callbacks = append(callbacks, setResourceTagsCallbacks...)

	return callbacks
}

func (s *VolcengineNetworkInterfaceService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteNetworkInterface",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"NetworkInterfaceId": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.VpcClient.DeleteNetworkInterfaceCommon(call.SdkParam)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading network interface on delete %q, %w", d.Id(), callErr))
						}
					}
					_, callErr = call.ExecuteCall(d, client, call)
					if callErr == nil {
						return nil
					}
					return resource.RetryableError(callErr)
				})
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 3*time.Minute)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineNetworkInterfaceService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "NetworkInterfaceIds",
				ConvertType: ve.ConvertWithN,
			},
			"primary_ip_addresses": {
				TargetField: "PrimaryIpAddresses",
				ConvertType: ve.ConvertWithN,
			},
			"private_ip_addresses": {
				TargetField: "PrivateIpAddresses",
				ConvertType: ve.ConvertWithN,
			},
			"network_interface_ids": {
				TargetField: "NetworkInterfaceIds",
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
		NameField:    "NetworkInterfaceName",
		IdField:      "NetworkInterfaceId",
		CollectField: "network_interfaces",
		ResponseConverts: map[string]ve.ResponseConvert{
			"NetworkInterfaceId": {
				TargetField: "id",
				KeepDefault: true,
			},
			"AssociatedElasticIp.AllocationId": {
				TargetField: "associated_elastic_ip_id",
			},
			"AssociatedElasticIp.EipAddress": {
				TargetField: "associated_elastic_ip_address",
			},
		},
	}
}

func (s *VolcengineNetworkInterfaceService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "vpc",
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		Action:      actionName,
	}
}

func (s *VolcengineNetworkInterfaceService) ProjectTrn() *ve.ProjectTrn {
	return &ve.ProjectTrn{
		ServiceName:          "vpc",
		ResourceType:         "eni",
		ProjectResponseField: "ProjectName",
		ProjectSchemaField:   "project_name",
	}
}
