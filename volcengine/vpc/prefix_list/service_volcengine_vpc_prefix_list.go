package prefix_list

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

type VolcengineVpcPrefixListService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewVpcPrefixListService(c *ve.SdkClient) *VolcengineVpcPrefixListService {
	return &VolcengineVpcPrefixListService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineVpcPrefixListService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineVpcPrefixListService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribePrefixLists"
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
		results, err = ve.ObtainSdkValue("Result.PrefixLists", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.Items is not Slice")
		}
		for index, ele := range data {
			prefixList := ele.(map[string]interface{})
			prefixId := prefixList["PrefixListId"].(string)
			query := map[string]interface{}{
				"PrefixListId": prefixId,
			}
			action = "DescribePrefixListAssociations"
			logger.Debug(logger.ReqFormat, action, string(bytes))
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &query)
			if err != nil {
				return data, err
			}
			logger.Debug(logger.RespFormat, action, condition, string(respBytes))
			prefixListAssociations, err := ve.ObtainSdkValue("Result.PrefixListAssociations", *resp)
			if err != nil {
				return data, err
			}
			data[index].(map[string]interface{})["PrefixListAssociations"] = prefixListAssociations
			action = "DescribePrefixListEntries"
			logger.Debug(logger.ReqFormat, action, string(bytes))
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &query)
			if err != nil {
				return data, err
			}
			logger.Debug(logger.RespFormat, action, condition, string(respBytes))
			prefixListEntries, err := ve.ObtainSdkValue("Result.PrefixListEntries", *resp)
			if err != nil {
				return data, err
			}
			data[index].(map[string]interface{})["PrefixListEntries"] = prefixListEntries
		}
		return data, err
	})
}

func (s *VolcengineVpcPrefixListService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"PrefixListIds.1": id,
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
		return data, fmt.Errorf("vpc_prefix_list %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineVpcPrefixListService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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
			if err = resource.Retry(20*time.Minute, func() *resource.RetryError {
				d, err = s.ReadResource(resourceData, id)
				if err != nil {
					if ve.ResourceNotFoundError(err) {
						return resource.RetryableError(err)
					} else {
						return resource.NonRetryableError(err)
					}
				}
				return nil
			}); err != nil {
				return nil, "", err
			}
			status, err = ve.ObtainSdkValue("Status", d)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("vpc_prefix_list status error, status: %s", status.(string))
				}
			}
			return d, status.(string), err
		},
	}
}

func (s *VolcengineVpcPrefixListService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreatePrefixList",
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"prefix_list_entries": {
					TargetField: "PrefixListEntries",
					ConvertType: ve.ConvertListN,
				},
				"tags": {
					TargetField: "Tags",
					ConvertType: ve.ConvertListN,
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("Result.PrefixListId", *resp)
				d.SetId(id.(string))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Available"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineVpcPrefixListService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineVpcPrefixListService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyPrefixList",
			ConvertMode: ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"prefix_list_name": {
					ConvertType: ve.ConvertDefault,
				},
				"description": {
					ConvertType: ve.ConvertDefault,
				},
				"max_entries": {
					ConvertType: ve.ConvertDefault,
				},
				"prefix_list_entries": {
					Ignore: true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["PrefixListId"] = d.Id()
				addEntries, removeEntries, _, _ := ve.GetSetDifference("prefix_list_entries", d, entriesHash, false)
				if addEntries != nil && addEntries.Len() > 0 {
					for index, entry := range addEntries.List() {
						if cidr, ok := entry.(map[string]interface{})["cidr"]; ok {
							(*call.SdkParam)["AddPrefixListEntries."+strconv.Itoa(index+1)+".Cidr"] = cidr
						}
						if description, ok := entry.(map[string]interface{})["description"]; ok {
							(*call.SdkParam)["AddPrefixListEntries."+strconv.Itoa(index+1)+".Description"] = description
						}
					}
				}
				if removeEntries != nil && removeEntries.Len() > 0 {
					for index, entry := range removeEntries.List() {
						if cidr, ok := entry.(map[string]interface{})["cidr"]; ok {
							(*call.SdkParam)["RemovePrefixListEntries."+strconv.Itoa(index+1)+".Cidr"] = cidr
						}
					}
				}
				delete(*call.SdkParam, "Tags")
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Available"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	callbacks = append(callbacks, callback)
	return callbacks
}

func (s *VolcengineVpcPrefixListService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeletePrefixList",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"PrefixListId": resourceData.Id(),
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

func (s *VolcengineVpcPrefixListService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "PrefixListIds",
				ConvertType: ve.ConvertWithN,
			},
			"tag_filters": {
				TargetField: "TagFilters",
				ConvertType: ve.ConvertListN,
				NextLevelConvert: map[string]ve.RequestConvert{
					"values": {
						ConvertType: ve.ConvertWithN,
					},
				},
			},
		},
		NameField:    "PrefixListName",
		IdField:      "PrefixListId",
		CollectField: "prefix_lists",
		ResponseConverts: map[string]ve.ResponseConvert{
			"PrefixListId": {
				TargetField: "id",
				KeepDefault: true,
			},
		},
	}
}

func (s *VolcengineVpcPrefixListService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "vpc",
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		Action:      actionName,
	}
}
