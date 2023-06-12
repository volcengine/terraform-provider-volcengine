package allow_list

import (
	"errors"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

const (
	ActionDeleteAllowList = "DeleteAllowList"
	ActionCreateAllowList = "CreateAllowList"
	ActionModifyAllowList = "ModifyAllowList"
)

type VolcengineRedisAllowListService struct {
	Client *ve.SdkClient
}

func NewRedisAllowListService(c *ve.SdkClient) *VolcengineRedisAllowListService {
	return &VolcengineRedisAllowListService{
		Client: c,
	}
}

func (s *VolcengineRedisAllowListService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineRedisAllowListService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithSimpleQuery(m, func(condition map[string]interface{}) ([]interface{}, error) {
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
		logger.Debug(logger.RespFormat, action, resp)
		results, err = ve.ObtainSdkValue("Result.AllowLists", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.AllowLists is not Slice")
		}

		for index, element := range data {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo("DescribeAllowListDetail"), &map[string]interface{}{
				"AllowListId": element.(map[string]interface{})["AllowListId"],
			})
			if err != nil {
				return data, err
			}
			respResult, err := ve.ObtainSdkValue("Result", *resp)
			if err != nil {
				return data, err
			}
			logger.Debug(logger.RespFormat, "DescribeAllowListDetail", *resp)
			// 多个地址间用英文逗号（,）隔开
			ips := respResult.(map[string]interface{})["AllowList"]
			data[index].(map[string]interface{})["AllowList"] = strings.Split(ips.(string), ",")
			data[index].(map[string]interface{})["AssociatedInstances"] = respResult.(map[string]interface{})["AssociatedInstances"]
		}
		return data, err
	})
}

func (s *VolcengineRedisAllowListService) ReadResource(resourceData *schema.ResourceData, allowlistId string) (data map[string]interface{}, err error) {
	if allowlistId == "" {
		allowlistId = s.ReadResourceId(resourceData.Id())
	}
	resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo("DescribeAllowListDetail"), &map[string]interface{}{
		"AllowListId": allowlistId,
	})
	if err != nil {
		return data, err
	}
	respResult, err := ve.ObtainSdkValue("Result", *resp)
	if err != nil {
		return data, err
	}
	// 组合数据，DescribeAllowLists 无法查询出来
	data = respResult.(map[string]interface{})
	ips := respResult.(map[string]interface{})["AllowList"]
	data["AllowList"] = strings.Split(ips.(string), ",")
	data["AllowListIPNum"] = len(strings.Split(ips.(string), ","))
	data["AssociatedInstanceNum"] = len(data["AssociatedInstances"].([]interface{}))
	return data, err
}

func (s *VolcengineRedisAllowListService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (VolcengineRedisAllowListService) WithResourceResponseHandlers(allowlist map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return allowlist, map[string]ve.ResponseConvert{
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

func (s *VolcengineRedisAllowListService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      ActionCreateAllowList,
			ContentType: ve.ContentTypeJson,
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"allow_list": {
					Convert: func(data *schema.ResourceData, i interface{}) interface{} {
						var res []string
						for _, ele := range i.(*schema.Set).List() {
							res = append(res, ele.(string))
						}
						return strings.Join(res, ",")
					},
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, *call.SdkParam)
				output, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, *call.SdkParam, *output)
				return output, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, err := ve.ObtainSdkValue("Result.AllowListId", *resp)
				if err != nil {
					return err
				}
				d.SetId(id.(string))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineRedisAllowListService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      ActionModifyAllowList,
			ContentType: ve.ContentTypeJson,
			ConvertMode: ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"allow_list_name": {
					ForceGet: true, // 必须传递
				},
				"allow_list_desc": {
					TargetField: "AllowListDesc",
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["AllowListId"] = d.Id()
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				if d.HasChange("allow_list") {
					data, err := s.ReadResource(d, d.Id())
					if err != nil {
						return nil, err
					}
					(*call.SdkParam)["ApplyInstanceNum"] = data["AssociatedInstanceNum"]

					var res []string
					for _, ele := range d.Get("allow_list").(*schema.Set).List() {
						res = append(res, ele.(string))
					}
					(*call.SdkParam)["AllowList"] = strings.Join(res, ",")
				}

				logger.Debug(logger.ReqFormat, call.Action, *call.SdkParam)
				output, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, *call.SdkParam, *output)
				return output, err
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineRedisAllowListService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      ActionDeleteAllowList,
			ContentType: ve.ContentTypeJson,
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["AllowListId"] = resourceData.Id()
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, *call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineRedisAllowListService) DatasourceResources(data *schema.ResourceData, resource2 *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		ContentType:  ve.ContentTypeJson,
		NameField:    "AllowListName",
		IdField:      "AllowListId",
		CollectField: "allow_lists",
		ResponseConverts: map[string]ve.ResponseConvert{
			"VPC": {
				TargetField: "vpc",
			},
			"AllowListIPNum": {
				TargetField: "allow_list_ip_num",
			},
		},
	}
}

func (s *VolcengineRedisAllowListService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "Redis",
		Version:     "2020-12-07",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}
