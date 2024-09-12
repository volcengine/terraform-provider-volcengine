package allowlist

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	volc "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineRdsMysqlAllowListService struct {
	Client     *volc.SdkClient
	Dispatcher *volc.Dispatcher
}

func (s *VolcengineRdsMysqlAllowListService) GetClient() *volc.SdkClient {
	return s.Client
}

func (s *VolcengineRdsMysqlAllowListService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	var (
		resp        *map[string]interface{}
		results     interface{}
		ok          bool
		allowListId string
	)
	return volc.WithSimpleQuery(condition, func(m map[string]interface{}) ([]interface{}, error) {
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
		results, err = volc.ObtainSdkValue("Result.AllowLists", *resp)
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
		for index, ele := range data {
			allowList := ele.(map[string]interface{})

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
				instances, err := volc.ObtainSdkValue("Result.AssociatedInstances", *resp)
				if err != nil {
					return data, err
				}
				data[index].(map[string]interface{})["AssociatedInstances"] = instances
				allowListIp, err := volc.ObtainSdkValue("Result.AllowList", *resp)
				if err != nil {
					return data, err
				}
				allowListIpArr := strings.Split(allowListIp.(string), ",")
				data[index].(map[string]interface{})["AllowList"] = allowListIpArr
			}
		}
		return data, err
	})
}

func (s *VolcengineRdsMysqlAllowListService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
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
		return data, fmt.Errorf("Rds instance %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineRdsMysqlAllowListService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineRdsMysqlAllowListService) WithResourceResponseHandlers(m map[string]interface{}) []volc.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]volc.ResponseConvert, error) {
		return m, map[string]volc.ResponseConvert{}, nil
	}
	return []volc.ResourceResponseHandler{handler}
}

func (s *VolcengineRdsMysqlAllowListService) CreateResource(data *schema.ResourceData, resource *schema.Resource) []volc.Callback {
	callback := volc.Callback{
		Call: volc.SdkCall{
			Action:      "CreateAllowList",
			ConvertMode: volc.RequestConvertAll,
			ContentType: volc.ContentTypeJson,
			Convert: map[string]volc.RequestConvert{
				"allow_list": {
					Ignore: true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (bool, error) {
				var allowStrings []string
				allowListsSet := d.Get("allow_list").(*schema.Set)
				allowLists := allowListsSet.List()
				for _, list := range allowLists {
					allowStrings = append(allowStrings, list.(string))
				}
				lists := strings.Join(allowStrings, ",")
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam, lists)
				(*call.SdkParam)["AllowList"] = lists
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *volc.SdkClient, resp *map[string]interface{}, call volc.SdkCall) error {
				id, _ := volc.ObtainSdkValue("Result.AllowListId", *resp)
				d.SetId(id.(string))
				return nil
			},
		},
	}
	return []volc.Callback{callback}
}

func (s *VolcengineRdsMysqlAllowListService) ModifyResource(data *schema.ResourceData, resource *schema.Resource) []volc.Callback {
	callback := volc.Callback{
		Call: volc.SdkCall{
			Action:      "ModifyAllowList",
			ConvertMode: volc.RequestConvertInConvert,
			ContentType: volc.ContentTypeJson,
			Convert: map[string]volc.RequestConvert{
				"allow_list": {
					Ignore: true,
				},
				"apply_instance_num": {
					Ignore: true,
				},
				"allow_list_desc": {
					ForceGet: true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (bool, error) {
				var allowStrings []string
				// 修改allowList必须传ApplyInstanceNum
				resp, err := s.ReadResource(d, d.Id())
				if err != nil {
					return false, err
				}
				num := resp["AssociatedInstanceNum"].(float64)
				(*call.SdkParam)["ApplyInstanceNum"] = int(num)
				allowListsSet := d.Get("allow_list").(*schema.Set)
				allowLists := allowListsSet.List()
				for _, list := range allowLists {
					allowStrings = append(allowStrings, list.(string))
				}
				lists := strings.Join(allowStrings, ",")
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam, lists)
				(*call.SdkParam)["AllowList"] = lists
				return true, nil
			},
			SdkParam: &map[string]interface{}{
				"AllowListId":   data.Id(),
				"AllowListName": data.Get("allow_list_name").(string),
			},
			ExecuteCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	return []volc.Callback{callback}
}

func (s *VolcengineRdsMysqlAllowListService) RemoveResource(data *schema.ResourceData, resource *schema.Resource) []volc.Callback {
	callback := volc.Callback{
		Call: volc.SdkCall{
			Action:      "DeleteAllowList",
			ConvertMode: volc.RequestConvertIgnore,
			ContentType: volc.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"AllowListId": data.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	return []volc.Callback{callback}
}

func (s *VolcengineRdsMysqlAllowListService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) volc.DataSourceInfo {
	return volc.DataSourceInfo{
		ContentType:  volc.ContentTypeJson,
		NameField:    "AllowListName",
		IdField:      "AllowListId",
		CollectField: "allow_lists",
		ResponseConverts: map[string]volc.ResponseConvert{
			"AllowListIPNum": {
				TargetField: "allow_list_ip_num",
			},
			"VPC": {
				TargetField: "vpc",
			},
		},
	}
}

func (s *VolcengineRdsMysqlAllowListService) ReadResourceId(id string) string {
	return id
}

func NewRdsMysqlAllowListService(client *volc.SdkClient) *VolcengineRdsMysqlAllowListService {
	return &VolcengineRdsMysqlAllowListService{
		Client:     client,
		Dispatcher: &volc.Dispatcher{},
	}
}

func getUniversalInfo(actionName string) volc.UniversalInfo {
	return volc.UniversalInfo{
		ServiceName: "rds_mysql",
		Version:     "2022-01-01",
		HttpMethod:  volc.POST,
		ContentType: volc.ApplicationJSON,
		Action:      actionName,
	}
}
