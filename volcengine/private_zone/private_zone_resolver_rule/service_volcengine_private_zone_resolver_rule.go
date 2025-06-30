package private_zone_resolver_rule

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcenginePrivateZoneResolverRuleService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewPrivateZoneResolverRuleService(c *ve.SdkClient) *VolcenginePrivateZoneResolverRuleService {
	return &VolcenginePrivateZoneResolverRuleService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcenginePrivateZoneResolverRuleService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcenginePrivateZoneResolverRuleService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	data, err = ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListResolverRules"

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

		results, err = ve.ObtainSdkValue("Result.Rules", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.Items is not Slice")
		}
		return data, err
	})
	if err != nil {
		return data, err
	}

	for _, v := range data {
		rule, ok := v.(map[string]interface{})
		if !ok {
			return data, errors.New("Value is not map ")
		}
		rule["RID"] = rule["ID"]
		delete(rule, "ID")
		rule["RuleId"] = strconv.Itoa(int(rule["RID"].(float64)))
		action := "QueryResolverRule"
		req := map[string]interface{}{
			"RuleID": rule["RID"],
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
		rule["BindVpcs"] = bindVpcs

		if zone, ok := rule["ZoneName"]; ok {
			zoneSet := make([]string, 0)
			arr := strings.Split(zone.(string), ",")
			if len(arr) == 1 {
				// 接口自动加的.
				if strings.HasSuffix(arr[0], ".") {
					zoneSet = append(zoneSet, arr[0][:len(arr[0])-1])
				} else {
					zoneSet = arr
				}
			} else {
				zoneSet = arr
			}
			rule["ZoneName"] = zoneSet
		}
	}
	return data, nil
}

func (s *VolcenginePrivateZoneResolverRuleService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}

	ruleId, err := strconv.Atoi(id)
	if err != nil {
		return data, errors.New("rule id is not a integer")
	}

	req := map[string]interface{}{
		//"Id": id,
	}

	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}

	for _, v := range results {
		tmpData, ok := v.(map[string]interface{})
		if !ok {
			return data, errors.New("Value is not map ")
		} else if ruleId == int(tmpData["RID"].(float64)) {
			data = tmpData
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("private_zone_resolver_rule %s not exist ", id)
	}
	return data, err
}

func (s *VolcenginePrivateZoneResolverRuleService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				d          map[string]interface{}
				status     interface{}
				failStates []string
			)
			failStates = append(failStates, "Failed")
			d, err = s.ReadResource(resourceData, id)
			if err != nil {
				return nil, "", err
			}
			status, err = ve.ObtainSdkValue("Status", d)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("private_zone_resolver_rule status error, status: %s", status.(string))
				}
			}
			return d, status.(string), err
		},
	}
}

func (s *VolcenginePrivateZoneResolverRuleService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateResolverRule",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"vpcs": {
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
				"forward_ips": {
					TargetField: "ForwardIPs",
					ConvertType: ve.ConvertJsonObjectArray,
					NextLevelConvert: map[string]ve.RequestConvert{
						"ip": {
							TargetField: "IP",
						},
					},
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
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

				zoneNames, ok := d.GetOk("zone_name")
				if ok {
					zoneNamesSet, ok := zoneNames.(*schema.Set)
					if ok {
						arr := make([]string, 0)
						for _, zone := range zoneNamesSet.List() {
							arr = append(arr, zone.(string))
						}
						(*call.SdkParam)["ZoneName"] = strings.Join(arr, ",")
					}
				}

				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getPostUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("Result.RuleID", *resp)
				d.SetId(strconv.Itoa(int(id.(float64))))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcenginePrivateZoneResolverRuleService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"ID": {
				TargetField: "vpc_id",
			},
			"BindVpcs": {
				TargetField: "vpcs",
			},
			"EndpointID": {
				TargetField: "endpoint_id",
			},
			"ForwardIPs": {
				TargetField: "forward_ips",
			},
			"IP": {
				TargetField: "ip",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcenginePrivateZoneResolverRuleService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateResolverRule",
			ConvertMode: ve.RequestConvertInConvert,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"name": {
					TargetField: "Name",
				},
				"line": {
					TargetField: "Line",
				},
				"forward_ips": {
					Ignore: true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				id, err := strconv.Atoi(d.Id())
				if err != nil {
					return false, err
				}
				(*call.SdkParam)["RuleID"] = id

				if d.HasChange("forward_ips") {
					ips, ok := d.GetOk("forward_ips")
					if ok {
						var results []interface{}
						for _, ip := range ips.(*schema.Set).List() {
							ipMap := ip.(map[string]interface{})
							result := make(map[string]interface{})
							if _, ok = ipMap["ip"]; ok {
								result["IP"] = ipMap["ip"]
							}
							if _, ok = ipMap["port"]; ok {
								result["Port"] = ipMap["port"]
							}
							if len(result) > 0 {
								results = append(results, result)
							}
						}
						(*call.SdkParam)["ForwardIPs"] = results
					}
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

	if resourceData.HasChange("vpcs") {
		bindVpcCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "BindRuleVPC",
				ConvertMode: ve.RequestConvertInConvert,
				ContentType: ve.ContentTypeJson,
				Convert: map[string]ve.RequestConvert{
					"vpcs": {
						Ignore: true,
					},
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					id, err := strconv.Atoi(d.Id())
					if err != nil {
						return false, err
					}
					(*call.SdkParam)["RuleID"] = id

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

func (s *VolcenginePrivateZoneResolverRuleService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteResolverRule",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				id, err := strconv.Atoi(d.Id())
				if err != nil {
					return false, err
				}
				(*call.SdkParam)["RuleID"] = id
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
							return resource.NonRetryableError(fmt.Errorf("error on reading private zone resolver rule on delete %q, %w", d.Id(), callErr))
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

func (s *VolcenginePrivateZoneResolverRuleService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"endpoint_id": {
				TargetField: "EndpointID",
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
		NameField:    "Name",
		IdField:      "RuleId",
		CollectField: "rules",
		ContentType:  ve.ContentTypeJson,
		ResponseConverts: map[string]ve.ResponseConvert{
			"RuleId": {
				TargetField: "id",
			},
			"RID": {
				TargetField: "rule_id",
			},
			"EndpointID": {
				TargetField: "endpoint_id",
			},
			"ForwardIPs": {
				TargetField: "forward_ips",
			},
			"IP": {
				TargetField: "ip",
			},
		},
	}
}

func (s *VolcenginePrivateZoneResolverRuleService) ReadResourceId(id string) string {
	return id
}

func (s *VolcenginePrivateZoneResolverRuleService) setResourceTags(resourceData *schema.ResourceData, callbacks []ve.Callback) []ve.Callback {
	addedTags, removedTags, _, _ := ve.GetSetDifference("tags", resourceData, ve.TagsHash, false)

	removeCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UntagResources",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if removedTags != nil && len(removedTags.List()) > 0 {
					(*call.SdkParam)["ResourceIds"] = []string{resourceData.Id()}
					(*call.SdkParam)["ResourceType"] = "rule"
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
					(*call.SdkParam)["ResourceType"] = "rule"
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

func (s *VolcenginePrivateZoneResolverRuleService) ProjectTrn() *ve.ProjectTrn {
	return &ve.ProjectTrn{
		ServiceName:          "private_zone",
		ResourceType:         "rule",
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
