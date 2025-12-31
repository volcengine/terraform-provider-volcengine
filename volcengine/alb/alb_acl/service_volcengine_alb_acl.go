package alb_acl

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

type VolcengineAclService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewAclService(c *ve.SdkClient) *VolcengineAclService {
	return &VolcengineAclService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineAclService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineAclService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(condition, "PageSize", "PageNumber", 20, 1, func(m map[string]interface{}) ([]interface{}, error) {
		action := "DescribeAcls"
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
		logger.Debug(logger.RespFormat, action, condition, *resp)

		results, err = ve.ObtainSdkValue("Result.Acls", *resp)
		if err != nil {
			return data, err
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.Acls is not Slice")
		}

		for index, ele := range data {
			acl := ele.(map[string]interface{})
			query := map[string]interface{}{
				"AclId": acl["AclId"],
			}

			logger.Debug(logger.ReqFormat, "DescribeAclAttributes", query)
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo("DescribeAclAttributes"), &query)
			if err != nil {
				return data, err
			}
			logger.Debug(logger.RespFormat, "DescribeAclAttributes", query, *resp)

			ls, err := ve.ObtainSdkValue("Result.Listeners", *resp)
			if err != nil {
				return data, err
			}
			ae, err := ve.ObtainSdkValue("Result.AclEntries", *resp)
			if err != nil {
				return data, err
			}
			data[index].(map[string]interface{})["Listeners"] = ls
			data[index].(map[string]interface{})["AclEntries"] = ae
		}
		return data, err
	})
}

func (s *VolcengineAclService) ReadResource(resourceData *schema.ResourceData, aclId string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if aclId == "" {
		aclId = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"AclIds.1": aclId,
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
		return data, fmt.Errorf("acl %s not exist ", aclId)
	}
	return data, err
}

func (s *VolcengineAclService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (VolcengineAclService) WithResourceResponseHandlers(acl map[string]interface{}) []ve.ResourceResponseHandler {
	return []ve.ResourceResponseHandler{}

}

func (s *VolcengineAclService) CreateResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateAcl",
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"acl_entries": {
					Ignore: true,
				},
				"tags": {
					ConvertType: ve.ConvertListN,
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				time.Sleep(2 * time.Second)
				id, _ := ve.ObtainSdkValue("Result.AclId", *resp)
				d.SetId(id.(string))
				return nil
			},
		},
	}

	callbacks = append(callbacks, callback)
	//规则创建
	entryCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "AddAclEntries",
			ConvertMode: ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"acl_entries": {
					ConvertType: ve.ConvertListN,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				if len(*call.SdkParam) > 0 {
					(*call.SdkParam)["AclId"] = d.Id()
					return true, nil
				}
				return false, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				time.Sleep(2 * time.Second) // 均为异步操作，无法获取 Status
				return nil
			},
		},
	}
	callbacks = append(callbacks, entryCallback)
	return callbacks

}

func (s *VolcengineAclService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyAclAttributes",
			ConvertMode: ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"acl_name": {
					TargetField: "AclName",
				},
				"description": {
					TargetField: "Description",
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if len(*call.SdkParam) == 0 {
					return false, nil
				}
				(*call.SdkParam)["AclId"] = d.Id()
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	callbacks = append(callbacks, callback)

	//规则修改
	add, remove, _, _ := ve.GetSetDifference("acl_entries", resourceData, AclEntryHash, false)

	entryRemoveCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "RemoveAclEntries",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if remove != nil && len(remove.List()) > 0 {
					(*call.SdkParam)["AclId"] = d.Id()
					for index, entry := range remove.List() {
						(*call.SdkParam)["Entries."+strconv.Itoa(index+1)] = entry.(map[string]interface{})["entry"].(string)
					}
					return true, nil
				}
				return false, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				time.Sleep(2 * time.Second)
				return nil
			},
		},
	}
	callbacks = append(callbacks, entryRemoveCallback)

	entryAddCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "AddAclEntries",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if add != nil && len(add.List()) > 0 {
					(*call.SdkParam)["AclId"] = d.Id()
					for index, entry := range add.List() {
						(*call.SdkParam)["AclEntries."+strconv.Itoa(index+1)+"."+"Entry"] = entry.(map[string]interface{})["entry"].(string)
						(*call.SdkParam)["AclEntries."+strconv.Itoa(index+1)+"."+"Description"] = entry.(map[string]interface{})["description"].(string)
					}
					return true, nil
				}
				return false, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				time.Sleep(2 * time.Second)
				return nil
			},
		},
	}
	callbacks = append(callbacks, entryAddCallback)

	//更新 tags
	setResourceTagsCallbacks := ve.SetResourceTags(s.Client, "TagResources", "UntagResources", "acl", resourceData, getUniversalInfo)
	callbacks = append(callbacks, setResourceTagsCallbacks...)

	return callbacks
}

func (s *VolcengineAclService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteAcl",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"AclId": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineAclService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "AclIds",
				ConvertType: ve.ConvertWithN,
			},
			"tags": {
				TargetField: "TagFilters",
				ConvertType: ve.ConvertListN,
				NextLevelConvert: map[string]ve.RequestConvert{
					"value": {
						TargetField: "Values.1",
					},
				},
			},
		},
		NameField:    "AclName",
		IdField:      "AclId",
		CollectField: "acls",
		ResponseConverts: map[string]ve.ResponseConvert{
			"AclId": {
				TargetField: "id",
				KeepDefault: true,
			},
		},
	}
}

func (s *VolcengineAclService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "alb",
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		Action:      actionName,
	}
}
