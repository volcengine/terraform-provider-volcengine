package waf_ip_group

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

type VolcengineWafIpGroupService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewWafIpGroupService(c *ve.SdkClient) *VolcengineWafIpGroupService {
	return &VolcengineWafIpGroupService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineWafIpGroupService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineWafIpGroupService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "Page", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListAllIpGroups"

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
		results, err = ve.ObtainSdkValue("Result.IpGroupList", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.IpGroupList is not Slice")
		}

		for _, ele := range data {
			ipGroup, ok := ele.(map[string]interface{})
			if !ok {
				return data, fmt.Errorf(" ipGroup is not Map ")
			}

			ipGroupId := int(ipGroup["IpGroupId"].(float64))

			ipGroup["IpGroupIdString"] = strconv.Itoa(ipGroupId)

			logger.Debug(logger.ReqFormat, "IpGroupIdString", ipGroup["IpGroupIdString"])

			// 查询域名详细信息
			action := "ListIpGroup"
			req := map[string]interface{}{
				"IpGroupId": ipGroupId,
			}
			logger.Debug(logger.ReqFormat, action, req)

			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &req)
			if err != nil {
				return data, err
			}
			logger.Debug(logger.RespFormat, action, req, *resp)
			listIpGroup, err := ve.ObtainSdkValue("Result", *resp)
			if err != nil {
				return data, err
			}
			listIpGroupMap, ok := listIpGroup.(map[string]interface{})
			if !ok {
				return data, fmt.Errorf(" Result is not Map ")
			}

			ipGroup["IpList"] = listIpGroupMap["IpList"]
		}

		return data, err
	})
}

func (s *VolcengineWafIpGroupService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
		result  map[string]interface{}
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}

	ipGroupId, err := strconv.Atoi(id)
	if err != nil {
		return data, fmt.Errorf(" ipGroupId cannot convert to int ")
	}

	req := map[string]interface{}{
		"TimeOrderBy": "DESC",
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		if data, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("Value is not map ")
		}

		if int(data["IpGroupId"].(float64)) == ipGroupId {
			result = data
			break
		}
	}
	if len(result) == 0 {
		return result, fmt.Errorf("waf_host_group %s not exist ", id)
	}
	return result, err
}

func (s *VolcengineWafIpGroupService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineWafIpGroupService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "AddIpGroup",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"ip_list": {
					TargetField: "IpList",
					ConvertType: ve.ConvertJsonArray,
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("Result.IpGroupId", *resp)
				d.SetId(strconv.Itoa(int(id.(float64))))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineWafIpGroupService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineWafIpGroupService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateIpGroup",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"name": {
					TargetField: "Name",
					ForceGet:    true,
				},
				"add_type": {
					TargetField: "AddType",
					ForceGet:    true,
				},
				"ip_list": {
					TargetField: "IpList",
					ConvertType: ve.ConvertJsonArray,
					ForceGet:    true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				ipGroupId, err := strconv.Atoi(d.Id())
				if err != nil {
					return false, fmt.Errorf(" ipGroupId cannot convert to int ")
				}
				(*call.SdkParam)["IpGroupId"] = ipGroupId
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineWafIpGroupService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteIpGroup",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				ipGroupId, err := strconv.Atoi(d.Id())
				if err != nil {
					return false, fmt.Errorf(" ipGroupId cannot convert to int ")
				}
				(*call.SdkParam)["IpGroupIds"] = []int{ipGroupId}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 5*time.Minute)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(5*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on  reading waf ip group on delete %q, %w", d.Id(), callErr))
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

func (s *VolcengineWafIpGroupService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		NameField:    "Name",
		IdField:      "IpGroupIdString",
		CollectField: "ip_group_list",
		ContentType:  ve.ContentTypeJson,
		ResponseConverts: map[string]ve.ResponseConvert{
			"related_rules": {
				TargetField: "RelatedRules",
			},
		},
	}
}

func (s *VolcengineWafIpGroupService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "waf",
		Version:     "2023-12-25",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
		RegionType:  ve.Global,
	}
}

//
//func (s *VolcengineWafIpGroupService) checkResourceUtilRemoved(d *schema.ResourceData, timeout time.Duration) error {
//	return resource.Retry(timeout, func() *resource.RetryError {
//		ipGroup, _ := s.ReadResource(d, d.Id())
//		logger.Debug(logger.RespFormat, "ipGroup", ipGroup)
//
//		// 能查询成功代表还在删除中，重试
//		ipList, ok := ipGroup["IpList"].([]string)
//		if !ok {
//			return resource.NonRetryableError(fmt.Errorf("ipList is not []string"))
//		}
//		if len(ipList) != 0 {
//			return resource.RetryableError(fmt.Errorf("resource still in removing status "))
//		} else {
//			if len(ipList) == 0 {
//				return nil
//			} else {
//				return resource.NonRetryableError(fmt.Errorf("ipGroup status is not deleted "))
//			}
//		}
//	})
//}
