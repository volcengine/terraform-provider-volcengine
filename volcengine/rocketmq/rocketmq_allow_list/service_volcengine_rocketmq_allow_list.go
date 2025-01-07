package rocketmq_allow_list

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineRocketmqAllowListService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewRocketmqAllowListService(c *ve.SdkClient) *VolcengineRocketmqAllowListService {
	return &VolcengineRocketmqAllowListService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineRocketmqAllowListService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineRocketmqAllowListService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp        *map[string]interface{}
		results     interface{}
		ok          bool
		allowListId string
	)
	return ve.WithSimpleQuery(m, func(condition map[string]interface{}) ([]interface{}, error) {
		if condition != nil {
			condition["RegionId"] = s.Client.Region
		}

		action := "DescribeAllowLists"
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
		results, err = ve.ObtainSdkValue("Result.AllowLists", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.AllowLists is not slice ")
		}

		if id, exist := condition["AllowListId"]; exist {
			allowListId = id.(string)
		}
		for _, ele := range data {
			allowList, ok := ele.(map[string]interface{})
			if !ok {
				return data, fmt.Errorf("The value of the Result.AllowLists is not map ")
			}

			if allowListId == "" || allowListId == allowList["AllowListId"].(string) {
				query := map[string]interface{}{
					"AllowListId": allowList["AllowListId"],
				}
				action = "DescribeAllowListDetail"
				logger.Debug(logger.ReqFormat, action, query)
				resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &query)
				if err != nil {
					return data, err
				}
				logger.Debug(logger.RespFormat, action, query, *resp)
				instances, err := ve.ObtainSdkValue("Result.AssociatedInstances", *resp)
				if err != nil {
					return data, err
				}
				allowList["AssociatedInstances"] = instances
				allowListIp, err := ve.ObtainSdkValue("Result.AllowList", *resp)
				if err != nil {
					return data, err
				}
				allowListIpArr := strings.Split(allowListIp.(string), ",")
				allowList["AllowList"] = allowListIpArr
			}
		}
		return data, err
	})
}

func (s *VolcengineRocketmqAllowListService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"RegionId":    s.Client.Region,
		"AllowListId": id,
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		result, ok := v.(map[string]interface{})
		if !ok {
			return data, errors.New("Value is not map ")
		}
		if result["AllowListId"].(string) == id {
			data = result
			break
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("rocketmq_allow_list %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineRocketmqAllowListService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (VolcengineRocketmqAllowListService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"AllowListIPNum": {
				TargetField: "allow_list_ip_num",
			},
			"VPC": {
				TargetField: "vpc",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineRocketmqAllowListService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateAllowList",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"allow_list": {
					Ignore: true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				var allowStrings []string
				allowListsSet := d.Get("allow_list").(*schema.Set)
				for _, v := range allowListsSet.List() {
					allowStrings = append(allowStrings, v.(string))
				}
				allowLists := strings.Join(allowStrings, ",")

				(*call.SdkParam)["AllowList"] = allowLists
				(*call.SdkParam)["AllowListType"] = "IPv4"
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("Result.AllowListId", *resp)
				d.SetId(id.(string))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineRocketmqAllowListService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyAllowList",
			ConvertMode: ve.RequestConvertInConvert,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"allow_list_name": {
					TargetField: "AllowListName",
					ForceGet:    true,
				},
				"allow_list_desc": {
					TargetField: "AllowListDesc",
				},
				"allow_list": {
					Ignore: true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				var allowStrings []string
				allowListsSet := d.Get("allow_list").(*schema.Set)
				for _, v := range allowListsSet.List() {
					allowStrings = append(allowStrings, v.(string))
				}
				allowLists := strings.Join(allowStrings, ",")
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam, allowLists)
				(*call.SdkParam)["AllowList"] = allowLists

				(*call.SdkParam)["AllowListId"] = d.Id()
				(*call.SdkParam)["ModifyMode"] = "Cover"
				(*call.SdkParam)["ApplyInstanceNum"] = d.Get("associated_instance_num")
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

func (s *VolcengineRocketmqAllowListService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteAllowList",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"AllowListId": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 5*time.Minute)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(1*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading rocketmq allow list on delete %q, %w", d.Id(), callErr))
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

func (s *VolcengineRocketmqAllowListService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		NameField:    "AllowListName",
		IdField:      "AllowListId",
		CollectField: "rocketmq_allow_lists",
		ContentType:  ve.ContentTypeJson,
		ResponseConverts: map[string]ve.ResponseConvert{
			"AllowListId": {
				TargetField: "id",
				KeepDefault: true,
			},
			"AllowListIPNum": {
				TargetField: "allow_list_ip_num",
			},
			"VPC": {
				TargetField: "vpc",
			},
		},
	}
}

func (s *VolcengineRocketmqAllowListService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "RocketMQ",
		Version:     "2023-01-01",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}
