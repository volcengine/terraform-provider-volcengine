package iam_user

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineIamUserService struct {
	Client *ve.SdkClient
}

func NewIamUserService(c *ve.SdkClient) *VolcengineIamUserService {
	return &VolcengineIamUserService{
		Client: c,
	}
}

func (s *VolcengineIamUserService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineIamUserService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	cens, err := ve.WithPageOffsetQuery(m, "Limit", "Offset", 100, 0, func(condition map[string]interface{}) ([]interface{}, error) {
		universalClient := s.Client.UniversalClient
		action := "ListUsers"
		logger.Debug(logger.ReqFormat, action, condition)
		if condition == nil {
			resp, err = universalClient.DoCall(getUniversalInfo(action), nil)
			if err != nil {
				return data, err
			}
		} else {
			resp, err = universalClient.DoCall(getUniversalInfo(action), &condition)
			if err != nil {
				return data, err
			}
		}
		logger.Debug(logger.RespFormat, action, resp)
		results, err = ve.ObtainSdkValue("Result.UserMetadata", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.UserMetadata is not Slice")
		}
		data, err = removeSystemTags(data)
		return data, err
	})
	return cens, err
}

func (s *VolcengineIamUserService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"Query": id,
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		m, ok := v.(map[string]interface{})
		if !ok {
			continue
		}
		if name, ok := m["UserName"].(string); ok && name == id {
			data = m
			break
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("user %s not exist ", id)
	}

	return data, err
}

func (s *VolcengineIamUserService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (VolcengineIamUserService) WithResourceResponseHandlers(v map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return v, map[string]ve.ResponseConvert{
			"AccountId": {
				TargetField: "account_id",
				Convert: func(i interface{}) interface{} {
					return strconv.FormatFloat(i.(float64), 'f', 0, 64)
				},
			},
			"Id": {
				TargetField: "user_id",
				Convert: func(i interface{}) interface{} {
					return strconv.FormatFloat(i.(float64), 'f', 0, 64)
				},
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}

}

func (s *VolcengineIamUserService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateUser",
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"tags": {
					TargetField: "Tags",
					ConvertType: ve.ConvertListN,
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam, resp)
				d.SetId(d.Get("user_name").(string))
				return nil
			},
		},
	}

	return []ve.Callback{callback}
}

func (s *VolcengineIamUserService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateUser",
			ConvertMode: ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"user_name": {
					TargetField: "NewUserName",
					ConvertType: ve.ConvertDefault,
				},
				"display_name": {
					TargetField: "NewDisplayName",
					ConvertType: ve.ConvertDefault,
					Convert:     defaultConvert,
				},
				"mobile_phone": {
					TargetField: "NewMobilePhone",
				},
				"email": {
					TargetField: "NewEmail",
					Convert:     defaultConvert,
				},
				"description": {
					TargetField: "NewDescription",
					ConvertType: ve.ConvertDefault,
					Convert:     defaultConvert,
				},
			},
			RequestIdField: "UserName",
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				if d.HasChange("user_name") {
					d.SetId(d.Get("user_name").(string))
				}
				return nil
			},
		},
	}
	callbacks = append(callbacks, callback)
	setResourceTagsCallbacks := s.setResourceTags(resourceData, "User", callbacks)
	return setResourceTagsCallbacks
}

func (s *VolcengineIamUserService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:         "DeleteUser",
			ConvertMode:    ve.RequestConvertIgnore,
			RequestIdField: "UserName",
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}

	return []ve.Callback{callback}
}

func (s *VolcengineIamUserService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"query": {
				TargetField: "Query",
			},
		},
		NameField:    "UserName",
		IdField:      "UserName",
		CollectField: "users",
		ResponseConverts: map[string]ve.ResponseConvert{
			"Id": {
				TargetField: "user_id",
				Convert: func(i interface{}) interface{} {
					return strconv.FormatFloat(i.(float64), 'f', 0, 64)
				},
			},
			"AccountId": {
				TargetField: "account_id",
				Convert: func(i interface{}) interface{} {
					return strconv.FormatFloat(i.(float64), 'f', 0, 64)
				},
			},
			"Tags": {
				TargetField: "tags",
				Convert: func(i interface{}) interface{} {
					if i == nil {
						return nil
					}
					var tags []map[string]interface{}
					if list, ok := i.([]interface{}); ok {
						for _, v := range list {
							if m, ok := v.(map[string]interface{}); ok {
								tag := make(map[string]interface{})
								if key, ok := m["Key"].(string); ok {
									tag["key"] = key
								}
								if value, ok := m["Value"].(string); ok {
									tag["value"] = value
								}
								tags = append(tags, tag)
							}
						}
					}
					return tags
				},
			},
		},
	}
}

func (s *VolcengineIamUserService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "iam",
		Action:      actionName,
		Version:     "2018-01-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		RegionType:  ve.Global,
	}
}

func (s *VolcengineIamUserService) setResourceTags(resourceData *schema.ResourceData, resourceType string, callbacks []ve.Callback) []ve.Callback {
	addedTags, removedTags, _, _ := ve.GetSetDifference("tags", resourceData, ve.TagsHash, false)

	removeCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UntagResources",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if removedTags != nil && len(removedTags.List()) > 0 {
					(*call.SdkParam)["ResourceNames.1"] = resourceData.Id()
					(*call.SdkParam)["ResourceType"] = resourceType
					for index, tag := range removedTags.List() {
						(*call.SdkParam)["TagKeys."+strconv.Itoa(index+1)] = tag.(map[string]interface{})["key"].(string)
					}
					return true, nil
				}
				return false, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	callbacks = append(callbacks, removeCallback)

	addCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "TagResources",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if addedTags != nil && len(addedTags.List()) > 0 {
					(*call.SdkParam)["ResourceNames.1"] = resourceData.Id()
					(*call.SdkParam)["ResourceType"] = resourceType
					for index, tag := range addedTags.List() {
						(*call.SdkParam)["Tags."+strconv.Itoa(index+1)+".Key"] = tag.(map[string]interface{})["key"].(string)
						(*call.SdkParam)["Tags."+strconv.Itoa(index+1)+".Value"] = tag.(map[string]interface{})["value"].(string)
					}
					return true, nil
				}
				return false, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	callbacks = append(callbacks, addCallback)

	return callbacks
}

func removeSystemTags(data []interface{}) ([]interface{}, error) {
	var (
		ok      bool
		result  map[string]interface{}
		results []interface{}
		tags    []interface{}
	)
	for _, d := range data {
		if result, ok = d.(map[string]interface{}); !ok {
			return results, errors.New("The elements in data are not map ")
		}
		tags, ok = result["Tags"].([]interface{})
		if ok {
			tags = ve.FilterSystemTags(tags)
			result["Tags"] = tags
		}
		results = append(results, result)
	}
	return results, nil
}
