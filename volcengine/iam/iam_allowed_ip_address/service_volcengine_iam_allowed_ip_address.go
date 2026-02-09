package iam_allowed_ip_address

import (
	"errors"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineIamAllowedIpAddressService struct {
	Client *ve.SdkClient
}

func NewIamAllowedIpAddressService(c *ve.SdkClient) *VolcengineIamAllowedIpAddressService {
	return &VolcengineIamAllowedIpAddressService{
		Client: c,
	}
}

func (s *VolcengineIamAllowedIpAddressService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineIamAllowedIpAddressService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	action := "GetAllowedIPAddresses"
	logger.Debug(logger.ReqFormat, action, m)
	resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(action), &m)
	if err != nil {
		return data, err
	}
	logger.Debug(logger.RespFormat, action, resp)

	// 根据返回结构，Result 是顶级字段
	result, err := ve.ObtainSdkValue("Result", *resp)
	if err != nil {
		return data, err
	}

	if resMap, ok := result.(map[string]interface{}); ok {
		data = append(data, resMap)
	} else {
		return data, errors.New("Result is not map")
	}

	return data, nil
}

func (s *VolcengineIamAllowedIpAddressService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	results, err := s.ReadResources(nil)
	if err != nil {
		return data, err
	}
	if len(results) > 0 {
		return results[0].(map[string]interface{}), nil
	}
	return data, errors.New("Allowed IP Address not found")
}

func (s *VolcengineIamAllowedIpAddressService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (VolcengineIamAllowedIpAddressService) WithResourceResponseHandlers(v map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return v, map[string]ve.ResponseConvert{
			"IPList": {
				TargetField: "ip_list",
				Convert: func(i interface{}) interface{} {
					if i == nil {
						return nil
					}
					var list []interface{}
					if rawList, ok := i.([]interface{}); ok {
						for _, item := range rawList {
							if m, ok := item.(map[string]interface{}); ok {
								list = append(list, map[string]interface{}{
									"ip":          m["IP"],
									"description": m["Description"],
								})
							}
						}
					}
					return list
				},
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineIamAllowedIpAddressService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateAllowedIPAddresses",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"enable_ip_list": {
					TargetField: "EnableIPList",
				},
				"ip_list": {
					TargetField: "IPList",
					ConvertType: ve.ConvertJsonObjectArray,
					NextLevelConvert: map[string]ve.RequestConvert{
						"ip": {
							TargetField: "IP",
						},
						"description": {
							TargetField: "Description",
						},
					},
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				d.SetId("iam_allowed_ip_address")
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineIamAllowedIpAddressService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return s.CreateResource(resourceData, resource)
}

func (s *VolcengineIamAllowedIpAddressService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateAllowedIPAddresses",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["EnableIPList"] = false
				(*call.SdkParam)["IPList"] = []interface{}{}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineIamAllowedIpAddressService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		CollectField: "allowed_ip_addresses",
		ResponseConverts: map[string]ve.ResponseConvert{
			"IPList": {
				TargetField: "ip_list",
				Convert: func(i interface{}) interface{} {
					if i == nil {
						return nil
					}
					var list []interface{}
					if rawList, ok := i.([]interface{}); ok {
						for _, item := range rawList {
							if m, ok := item.(map[string]interface{}); ok {
								list = append(list, map[string]interface{}{
									"ip":          m["IP"],
									"description": m["Description"],
								})
							}
						}
					}
					return list
				},
			},
		},
	}
}

func (s *VolcengineIamAllowedIpAddressService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	method := ve.GET
	contentType := ve.Default
	if actionName == "UpdateAllowedIPAddresses" {
		method = ve.POST
		contentType = ve.ApplicationJSON
	}
	return ve.UniversalInfo{
		ServiceName: "iam",
		Action:      actionName,
		Version:     "2018-01-01",
		HttpMethod:  method,
		ContentType: contentType,
		RegionType:  ve.Global,
	}
}
