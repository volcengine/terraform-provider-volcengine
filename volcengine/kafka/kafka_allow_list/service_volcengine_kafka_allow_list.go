package kafka_allow_list

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineKafkaAllowListService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewKafkaAllowListService(c *ve.SdkClient) *VolcengineKafkaAllowListService {
	return &VolcengineKafkaAllowListService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineKafkaAllowListService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineKafkaAllowListService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithSimpleQuery(m, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeAllowLists"

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
		for _, allowList := range data {
			allowListMap, ok := allowList.(map[string]interface{})
			if !ok {
				continue
			}
			action = "DescribeAllowListDetail"
			req := map[string]interface{}{
				"AllowListId": allowListMap["AllowListId"],
			}
			logger.Debug(logger.ReqFormat, action, req)
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &req)
			if err != nil {
				return data, err
			}
			al, err := ve.ObtainSdkValue("Result.AllowList", *resp)
			if err != nil {
				continue
			}
			alStr, ok := al.(string)
			if ok {
				allowListMap["AllowList"] = strings.Split(alStr, ",")
			}
			instances, err := ve.ObtainSdkValue("Result.AssociatedInstances", *resp)
			if err != nil {
				continue
			}
			allowListMap["AssociatedInstances"] = instances
		}
		return data, err
	})
}

func (s *VolcengineKafkaAllowListService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results  []interface{}
		tempData map[string]interface{}
		ok       bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"RegionId": s.Client.Region,
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		if tempData, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("Value is not map ")
		} else {
			alId, ok := tempData["AllowListId"].(string)
			if ok && alId == id {
				data = tempData
				break
			}
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("kafka_allow_list %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineKafkaAllowListService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
	}
}

func (s *VolcengineKafkaAllowListService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
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
				allowlist, ok := d.GetOk("allow_list")
				if ok {
					alSet, ok := allowlist.(*schema.Set)
					if ok {
						alList := alSet.List()
						var alStrs []string
						for _, al := range alList {
							alStr, ok := al.(string)
							if ok {
								alStrs = append(alStrs, alStr)
							}
						}
						alStr := strings.Join(alStrs, ",")
						(*call.SdkParam)["AllowList"] = alStr
					}
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
				id, _ := ve.ObtainSdkValue("Result.AllowListId", *resp)
				d.SetId(id.(string))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineKafkaAllowListService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineKafkaAllowListService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
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
					ForceGet:    true,
				},
				"allow_list": {
					Ignore: true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["AllowListId"] = d.Id()
				resp, err := s.ReadResource(d, d.Id())
				if err != nil {
					return false, err
				}
				if d.HasChange("allow_list") {
					num := resp["AssociatedInstanceNum"].(float64)
					(*call.SdkParam)["ApplyInstanceNum"] = int(num)
					allowListsSet := d.Get("allow_list")
					alSet, ok := allowListsSet.(*schema.Set)
					if ok {
						alList := alSet.List()
						var alStrs []string
						for _, al := range alList {
							alStr, ok := al.(string)
							if ok {
								alStrs = append(alStrs, alStr)
							}
						}
						alStr := strings.Join(alStrs, ",")
						(*call.SdkParam)["AllowList"] = alStr
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
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineKafkaAllowListService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteAllowList",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"AllowListId": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 5*time.Minute)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineKafkaAllowListService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		ContentType:  ve.ContentTypeJson,
		NameField:    "AllowListName",
		IdField:      "AllowListId",
		CollectField: "allow_lists",
	}
}

func (s *VolcengineKafkaAllowListService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "Kafka",
		Version:     "2022-05-01",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}
