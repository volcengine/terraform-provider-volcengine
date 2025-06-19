package private_zone_resolver_endpoint

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

type VolcenginePrivateZoneResolverEndpointService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewPrivateZoneResolverEndpointService(c *ve.SdkClient) *VolcenginePrivateZoneResolverEndpointService {
	return &VolcenginePrivateZoneResolverEndpointService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcenginePrivateZoneResolverEndpointService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcenginePrivateZoneResolverEndpointService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListResolverEndpoints"

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
		results, err = ve.ObtainSdkValue("Result.Endpoints", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		logger.Debug(logger.RespFormat, action, condition, results)
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.Endpoints is not Slice")
		}

		for _, v := range data {
			endpoint, ok := v.(map[string]interface{})
			if !ok {
				return data, errors.New("Value is not map ")
			}
			endpoint["EndpointId"] = strconv.Itoa(int(endpoint["ID"].(float64)))
		}

		return data, err
	})
}

func (s *VolcenginePrivateZoneResolverEndpointService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	endpointId, err := strconv.Atoi(id)
	if err != nil {
		return data, errors.New("endpoint id is not a integer")
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
		} else if endpointId == int(tmpData["ID"].(float64)) {
			data = tmpData
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("private_zone_resolver_endpoint %s not exist ", id)
	}
	return data, err
}

func (s *VolcenginePrivateZoneResolverEndpointService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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
			failStates = append(failStates, "Error")
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
					return nil, "", fmt.Errorf("private_zone_resolver_endpoint status error, status: %s", status.(string))
				}
			}
			return d, status.(string), err
		},
	}
}

func (s *VolcenginePrivateZoneResolverEndpointService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateResolverEndpoint",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"vpc_id": {
					TargetField: "VpcID",
				},
				"security_group_id": {
					TargetField: "SecurityGroupID",
				},
				"vpc_trns": {
					TargetField: "VpcTrns",
					ConvertType: ve.ConvertJsonArray,
				},
				"tags": {
					TargetField: "Tags",
					ConvertType: ve.ConvertJsonObjectArray,
				},
				"ip_configs": {
					TargetField: "IpConfigs",
					ConvertType: ve.ConvertJsonObjectArray,
					NextLevelConvert: map[string]ve.RequestConvert{
						"az_id": {
							TargetField: "AzID",
						},
						"subnet_id": {
							TargetField: "SubnetID",
						},
						"ip": {
							TargetField: "IP",
						},
					},
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getPostUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("Result.EndpointID", *resp)
				d.SetId(strconv.Itoa(int(id.(float64))))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Running"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcenginePrivateZoneResolverEndpointService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"VpcID": {
				TargetField: "vpc_id",
			},
			"SecurityGroupID": {
				TargetField: "security_group_id",
			},
			"AzID": {
				TargetField: "az_id",
			},
			"SubnetID": {
				TargetField: "subnet_id",
			},
			"IP": {
				TargetField: "ip",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcenginePrivateZoneResolverEndpointService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateResolverEndpoint",
			ConvertMode: ve.RequestConvertInConvert,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"name": {
					TargetField: "Name",
				},
				"ip_configs": {
					Ignore: true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				id, err := strconv.Atoi(d.Id())
				if err != nil {
					return false, err
				}
				(*call.SdkParam)["EndpointID"] = id

				if d.HasChange("ip_configs") {
					ipConfigs, ok := d.GetOk("ip_configs")
					if ok {
						var results []interface{}
						for _, ipConfig := range ipConfigs.(*schema.Set).List() {
							ipMap := ipConfig.(map[string]interface{})
							result := make(map[string]interface{})
							if _, ok = ipMap["az_id"]; ok {
								result["AzID"] = ipMap["az_id"]
							}
							if _, ok = ipMap["subnet_id"]; ok {
								result["SubnetID"] = ipMap["subnet_id"]
							}
							if _, ok = ipMap["ip"]; ok {
								result["IP"] = ipMap["ip"]
							}
							if len(result) > 0 {
								results = append(results, result)
							}
						}
						(*call.SdkParam)["IpConfigs"] = results
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
			Refresh: &ve.StateRefresh{
				Target:  []string{"Running"},
				Timeout: resourceData.Timeout(schema.TimeoutUpdate),
			},
		},
	}
	callbacks = append(callbacks, callback)

	// Tags
	callbacks = s.setResourceTags(resourceData, callbacks)

	return callbacks
}

func (s *VolcenginePrivateZoneResolverEndpointService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteResolverEndpoint",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				id, err := strconv.Atoi(d.Id())
				if err != nil {
					return false, err
				}
				(*call.SdkParam)["EndpointID"] = id
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
							return resource.NonRetryableError(fmt.Errorf("error on reading private zone resolver endpoint on delete %q, %w", d.Id(), callErr))
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

func (s *VolcenginePrivateZoneResolverEndpointService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
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
		NameField:    "Name",
		IdField:      "EndpointId",
		CollectField: "endpoints",
		ContentType:  ve.ContentTypeJson,
		ResponseConverts: map[string]ve.ResponseConvert{
			"Id": {
				TargetField: "endpoint_id",
			},
			"EndpointId": {
				TargetField: "id",
			},
			"VpcID": {
				TargetField: "vpc_id",
			},
			"SecurityGroupID": {
				TargetField: "security_group_id",
			},
			"AzID": {
				TargetField: "az_id",
			},
			"SubnetID": {
				TargetField: "subnet_id",
			},
			"IP": {
				TargetField: "ip",
			},
		},
	}
}

func (s *VolcenginePrivateZoneResolverEndpointService) ReadResourceId(id string) string {
	return id
}

func (s *VolcenginePrivateZoneResolverEndpointService) setResourceTags(resourceData *schema.ResourceData, callbacks []ve.Callback) []ve.Callback {
	addedTags, removedTags, _, _ := ve.GetSetDifference("tags", resourceData, ve.TagsHash, false)

	removeCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UntagResources",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if removedTags != nil && len(removedTags.List()) > 0 {
					(*call.SdkParam)["ResourceIds"] = []string{resourceData.Id()}
					(*call.SdkParam)["ResourceType"] = "endpoint"
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
					(*call.SdkParam)["ResourceType"] = "endpoint"
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

func (s *VolcenginePrivateZoneResolverEndpointService) ProjectTrn() *ve.ProjectTrn {
	return &ve.ProjectTrn{
		ServiceName:          "private_zone",
		ResourceType:         "endpoint",
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
