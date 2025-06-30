package private_zone

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

type VolcenginePrivateZoneService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewPrivateZoneService(c *ve.SdkClient) *VolcenginePrivateZoneService {
	return &VolcenginePrivateZoneService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcenginePrivateZoneService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcenginePrivateZoneService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	data, err = ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListPrivateZones"

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
		return data, err
	})
	if err != nil {
		return data, err
	}

	for _, v := range data {
		zone, ok := v.(map[string]interface{})
		if !ok {
			return data, fmt.Errorf("Value is not map ")
		}
		zone["ZoneId"] = strconv.Itoa(int(zone["ZID"].(float64)))

		action := "QueryPrivateZone"
		req := map[string]interface{}{
			"ZID": zone["ZID"],
		}
		logger.Debug(logger.ReqFormat, action, req)
		resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &req)
		if err != nil {
			return data, err
		}
		logger.Debug(logger.RespFormat, action, req, resp)

		bindVpcs, err := ve.ObtainSdkValue("Result.BindVPCs", *resp)
		if err != nil {
			return data, err
		}
		zone["BindVpcs"] = bindVpcs
	}

	return data, err
}

func (s *VolcenginePrivateZoneService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	zid, err := strconv.Atoi(id)
	if err != nil {
		return data, fmt.Errorf(" ZID cannot convert to int ")
	}

	req := map[string]interface{}{
		"ZIDs": zid,
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		if data, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("Value is not map ")
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("private_zone %s not exist ", id)
	}
	return data, err
}

func (s *VolcenginePrivateZoneService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (VolcenginePrivateZoneService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"ID": {
				TargetField: "vpc_id",
			},
			"BindVpcs": {
				TargetField: "vpcs",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcenginePrivateZoneService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreatePrivateZone",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"vpcs": {
					Ignore: true,
				},
				"intelligent_mode": {
					Ignore: true,
				},
				"load_balance_mode": {
					Ignore: true,
				},
				"vpc_trns": {
					TargetField: "VpcTrns",
					ConvertType: ve.ConvertJsonArray,
				},
				"tags": {
					TargetField: "Tags",
					ConvertType: ve.ConvertJsonObjectArray,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				intelligentMode := d.Get("intelligent_mode").(bool)
				loadBalanceMode := d.Get("load_balance_mode").(bool)
				lineMode := 0
				if !intelligentMode && loadBalanceMode {
					lineMode = 1
				} else if intelligentMode && !loadBalanceMode {
					lineMode = 2
				} else if intelligentMode && loadBalanceMode {
					lineMode = 3
				}
				(*call.SdkParam)["LineMode"] = lineMode

				vpcSet, ok := d.Get("vpcs").(*schema.Set)
				if !ok {
					return false, fmt.Errorf(" vpcs is not set ")
				}
				bindVpcMap := make(map[string][]string)
				for _, v := range vpcSet.List() {
					var (
						region string
						vpcId  string
					)
					vpcMap, ok := v.(map[string]interface{})
					if !ok {
						return false, fmt.Errorf(" vpcs value is not map ")
					}
					vpcId = vpcMap["vpc_id"].(string)
					if regionId, exist := vpcMap["region"]; exist && regionId != "" {
						region = regionId.(string)
					} else {
						region = client.Region
					}
					bindVpcMap[region] = append(bindVpcMap[region], vpcId)
				}
				(*call.SdkParam)["VPCs"] = bindVpcMap

				return true, nil
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

func (s *VolcenginePrivateZoneService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdatePrivateZone",
			ConvertMode: ve.RequestConvertInConvert,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"recursion_mode": {
					TargetField: "RecursionMode",
				},
				"load_balance_mode": {
					TargetField: "LoadBalance",
				},
				"remark": {
					TargetField: "Remark",
				},
				"vpcs": {
					Ignore: true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if len(*call.SdkParam) > 0 {
					zid, err := strconv.Atoi(d.Id())
					if err != nil {
						return false, fmt.Errorf(" ZID cannot convert to int ")
					}
					(*call.SdkParam)["ZID"] = zid
					return true, nil
				}
				return false, nil
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

	if resourceData.HasChange("vpcs") {
		bindVpcCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "BindVPC",
				ConvertMode: ve.RequestConvertInConvert,
				ContentType: ve.ContentTypeJson,
				Convert: map[string]ve.RequestConvert{
					"vpcs": {
						Ignore: true,
					},
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					zid, err := strconv.Atoi(d.Id())
					if err != nil {
						return false, fmt.Errorf(" ZID cannot convert to int ")
					}
					(*call.SdkParam)["ZID"] = zid

					vpcSet, ok := d.Get("vpcs").(*schema.Set)
					if !ok {
						return false, fmt.Errorf(" vpcs is not set ")
					}
					bindVpcMap := make(map[string][]string)
					for _, v := range vpcSet.List() {
						var (
							region string
							vpcId  string
						)
						vpcMap, ok := v.(map[string]interface{})
						if !ok {
							return false, fmt.Errorf(" vpcs value is not map ")
						}
						vpcId = vpcMap["vpc_id"].(string)
						if regionId, exist := vpcMap["region"]; exist {
							region = regionId.(string)
						} else {
							region = client.Region
						}
						bindVpcMap[region] = append(bindVpcMap[region], vpcId)
					}
					(*call.SdkParam)["VPCs"] = bindVpcMap

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
		callbacks = append(callbacks, bindVpcCallback)
	}

	// Tags
	callbacks = s.setResourceTags(resourceData, callbacks)

	return callbacks
}

func (s *VolcenginePrivateZoneService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeletePrivateZone",
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
							return resource.NonRetryableError(fmt.Errorf("error on reading private zone on delete %q, %w", d.Id(), callErr))
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

func (s *VolcenginePrivateZoneService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"zid": {
				TargetField: "ZIDs",
			},
			"vpc_id": {
				TargetField: "VpcID",
			},
			"tag_filters": {
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
		CollectField: "private_zones",
		ContentType:  ve.ContentTypeJson,
		ResponseConverts: map[string]ve.ResponseConvert{
			"ZoneId": {
				TargetField: "id",
			},
			"ZID": {
				TargetField: "zid",
			},
			"ID": {
				TargetField: "id",
			},
			"AccountID": {
				TargetField: "account_id",
			},
		},
	}
}

func (s *VolcenginePrivateZoneService) ReadResourceId(id string) string {
	return id
}

func (s *VolcenginePrivateZoneService) setResourceTags(resourceData *schema.ResourceData, callbacks []ve.Callback) []ve.Callback {
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

func (s *VolcenginePrivateZoneService) ProjectTrn() *ve.ProjectTrn {
	return &ve.ProjectTrn{
		ServiceName:          "private_zone",
		ResourceType:         "zone",
		ProjectResponseField: "ProjectName",
		ProjectSchemaField:   "project_name",
	}
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "private_zone",
		Version:     "2022-06-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		RegionType:  ve.Global,
		Action:      actionName,
	}
}

func getPostUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "private_zone",
		Version:     "2022-06-01",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		RegionType:  ve.Global,
		Action:      actionName,
	}
}
