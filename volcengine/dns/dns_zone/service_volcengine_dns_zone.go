package dns_zone

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

type VolcengineZoneService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewZoneService(c *ve.SdkClient) *VolcengineZoneService {
	return &VolcengineZoneService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineZoneService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineZoneService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListZones"
		logger.Debug(logger.ReqFormat, action, condition)
		if filter, filterExist := condition["TagFilters"]; filterExist {
			for index, tag := range filter.([]interface{}) {
				var (
					Key    string
					Values []string
				)
				for k, v := range tag.(map[string]interface{}) {
					if k == "Key" {
						Key = v.(string)
					}

					if k == "Values" {
						ValuesInter := v.([]interface{})
						for _, value := range ValuesInter {
							Values = append(Values, value.(string))
						}
					}
				}
				tagFilterMap := struct {
					Key    string
					Values []string
				}{
					Key:    Key,
					Values: Values,
				}

				tagFilterMapBytes, _ := json.Marshal(tagFilterMap)
				logger.Debug(logger.RespFormat, action, condition, string(tagFilterMapBytes))

				condition[fmt.Sprintf("TagFilters.%d", index+1)] = string(tagFilterMapBytes)
				delete(condition, "TagFilters")
			}
		}
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
		results, err = ve.ObtainSdkValue("Result.Zones", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.Zones is not Slice")
		}

		for _, ele := range data {
			zone, ok := ele.(map[string]interface{})
			if !ok {
				return data, fmt.Errorf(" zone is not Map ")
			}

			zone["ZoneId"] = strconv.Itoa(int(zone["ZID"].(float64)))

			logger.Debug(logger.ReqFormat, "ZoneId", zone["ZoneId"])

			// 查询域名详细信息
			action := "QueryZone"
			req := map[string]interface{}{
				"ZID": int(zone["ZID"].(float64)),
			}
			logger.Debug(logger.ReqFormat, action, req)

			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &req)
			if err != nil {
				return data, err
			}
			logger.Debug(logger.RespFormat, action, req, *resp)
			queryZone, err := ve.ObtainSdkValue("Result", *resp)
			if err != nil {
				return data, err
			}
			qZMap, ok := queryZone.(map[string]interface{})
			if !ok {
				return data, fmt.Errorf(" Result is not Map ")
			}
			for k, v := range qZMap {
				zone[k] = v
			}
		}
		logger.Debug(logger.RespFormat, "QueryZone data is ", data)
		return data, err
	})
}

func (s *VolcengineZoneService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		resp *map[string]interface{}
		//results []interface{}
		//ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}

	zid, err := strconv.Atoi(id)
	if err != nil {
		return data, fmt.Errorf(" ZID cannot convert to int ")
	}

	req := map[string]interface{}{
		"ZID": zid,
	}
	action := "QueryZone"
	resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &req)
	if err != nil {
		return data, err
	}
	logger.Debug(logger.RespFormat, action, req, *resp)

	queryZone, err := ve.ObtainSdkValue("Result", *resp)
	if err != nil {
		return data, err
	}
	data, ok := queryZone.(map[string]interface{})
	if !ok {
		return data, fmt.Errorf(" Result is not Map ")
	}

	if len(data) == 0 {
		return data, fmt.Errorf("zone %s not exist", id)
	}

	data["ZID"] = zid

	return data, err
}

func (s *VolcengineZoneService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineZoneService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateZone",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"tags": {
					TargetField: "Tags",
					ConvertType: ve.ConvertJsonObjectArray,
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getPostUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("Result.ZID", *resp)
				d.SetId(strconv.Itoa(int(id.(float64))))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineZoneService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"ZID": {
				TargetField: "zid",
			},
			"InstanceID": {
				TargetField: "instance_id",
			},
			"Remark": {
				TargetField: "remark",
			},
			"AllocateDNSServerList": {
				TargetField: "allocate_dns_server_list",
			},
			"IsNSCorrect": {
				TargetField: "is_ns_correct",
			},
			"RealDNSServerList": {
				TargetField: "real_dns_server_list",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineZoneService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callbacks := make([]ve.Callback, 0)

	if resourceData.HasChange("remark") {
		callback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "UpdateZone",
				ConvertMode: ve.RequestConvertInConvert,
				ContentType: ve.ContentTypeJson,
				Convert: map[string]ve.RequestConvert{
					"remark": {
						TargetField: "Remark",
					},
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					zid, err := strconv.Atoi(d.Id())
					if err != nil {
						return false, fmt.Errorf(" ZID cannot convert to int ")
					}
					(*call.SdkParam)["ZID"] = zid
					if resourceData.HasChange("remark") {
						(*call.SdkParam)["Remark"] = d.Get("remark")
					}
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					resp, err := s.Client.UniversalClient.DoCall(getPostUniversalInfo(call.Action), call.SdkParam)
					logger.Debug(logger.RespFormat, call.Action, resp, err)
					return resp, err
				},
			},
		}
		callbacks = append(callbacks, callback)
	}

	// Tags
	callbacks = s.setResourceTags(resourceData, callbacks)

	return callbacks
}

func (s *VolcengineZoneService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteZone",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				zid, err := strconv.Atoi(d.Id())
				if err != nil {
					return false, fmt.Errorf(" ZID cannot convert to int ")
				}
				(*call.SdkParam)["ZID"] = zid
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getPostUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 5*time.Minute)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on  reading zone on delete %q, %w", d.Id(), callErr))
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

func (s *VolcengineZoneService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"zid": {
				TargetField: "ZID",
			},
			"tags": {
				TargetField: "TagFilters",
				ConvertType: ve.ConvertJsonObjectArray,
				NextLevelConvert: map[string]ve.RequestConvert{
					"key": {
						TargetField: "Key",
					},
					"values": {
						TargetField: "Values",
						ConvertType: ve.ConvertJsonArray,
					},
				},
			},
		},
		NameField:    "ZoneName",
		IdField:      "ZoneId",
		CollectField: "zones",
		ContentType:  ve.ContentTypeJson,
		ResponseConverts: map[string]ve.ResponseConvert{
			"ZoneId": {
				TargetField: "id",
			},
			"InstanceID": {
				TargetField: "instance_id",
			},
			"ZID": {
				TargetField: "zid",
			},
			"AllocateDNSServerList": {
				TargetField: "allocate_dns_server_list",
			},
			"IsNSCorrect": {
				TargetField: "is_ns_correct",
			},
			"RealDNSServerList": {
				TargetField: "real_dns_server_list",
			},
			"Remark": {
				TargetField: "remark",
			},
		},
	}
}

func (s *VolcengineZoneService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "DNS",
		Version:     "2018-08-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		Action:      actionName,
		RegionType:  ve.Global,
	}
}

func getPostUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "DNS",
		Version:     "2018-08-01",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
		RegionType:  ve.Global,
	}
}

func (s *VolcengineZoneService) setResourceTags(resourceData *schema.ResourceData, callbacks []ve.Callback) []ve.Callback {
	addedTags, removedTags, _, _ := ve.GetSetDifference("tags", resourceData, ve.TagsHash, false)

	removeCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UntagResources",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if removedTags != nil && len(removedTags.List()) > 0 {
					(*call.SdkParam)["ResourceIds"] = []string{resourceData.Id()}
					(*call.SdkParam)["ResourceType"] = "zone"
					(*call.SdkParam)["TagKeys"] = make([]string, 0)
					for _, tag := range removedTags.List() {
						(*call.SdkParam)["TagKeys"] = append((*call.SdkParam)["TagKeys"].([]string), tag.(map[string]interface{})["key"].(string))
					}
					return true, nil
				}
				return false, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getPostUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	callbacks = append(callbacks, removeCallback)

	addCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "TagResources",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if addedTags != nil && len(addedTags.List()) > 0 {
					(*call.SdkParam)["ResourceIds"] = []string{resourceData.Id()}
					(*call.SdkParam)["ResourceType"] = "zone"
					(*call.SdkParam)["Tags"] = make([]map[string]interface{}, 0)
					for _, tag := range addedTags.List() {
						(*call.SdkParam)["Tags"] = append((*call.SdkParam)["Tags"].([]map[string]interface{}), tag.(map[string]interface{}))
					}
					return true, nil
				}
				return false, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getPostUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	callbacks = append(callbacks, addCallback)

	return callbacks
}

func (s *VolcengineZoneService) ProjectTrn() *ve.ProjectTrn {
	return &ve.ProjectTrn{
		ServiceName:          "dns",
		ResourceType:         "zone",
		ProjectResponseField: "ProjectName",
		ProjectSchemaField:   "project_name",
	}
}
