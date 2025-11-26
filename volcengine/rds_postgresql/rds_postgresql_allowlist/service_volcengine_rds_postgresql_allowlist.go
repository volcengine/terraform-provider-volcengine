package rds_postgresql_allowlist

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineRdsPostgresqlAllowlistService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewRdsPostgresqlAllowlistService(c *ve.SdkClient) *VolcengineRdsPostgresqlAllowlistService {
	return &VolcengineRdsPostgresqlAllowlistService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineRdsPostgresqlAllowlistService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineRdsPostgresqlAllowlistService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp        *map[string]interface{}
		results     interface{}
		ok          bool
		allowListId string
	)
	if id, exist := m["AllowListId"]; exist {
		allowListId = id.(string)
	}

	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 20, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		if condition != nil {
			condition["RegionId"] = s.Client.Region
		}

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
		results, err = ve.ObtainSdkValue("Result.AllowLists", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.AllowLists is not slice ")
		}

		for _, ele := range data {
			allowList, ok := ele.(map[string]interface{})
			if !ok {
				return data, fmt.Errorf("The value of the Result.AllowLists is not map ")
			}

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
				instances, err := ve.ObtainSdkValue("Result.AssociatedInstances", *resp)
				if err != nil {
					return data, err
				}
				allowList["AssociatedInstances"] = instances
				associatedNum, _ := ve.ObtainSdkValue("Result.AssociatedInstanceNum", *resp)
				if associatedNum != nil {
					allowList["AssociatedInstanceNum"] = associatedNum
				}
				allowListIp, err := ve.ObtainSdkValue("Result.AllowList", *resp)
				if err != nil {
					return data, err
				}
				if allowListIp != nil {
					allowList["AllowList"] = strings.Split(allowListIp.(string), ",")
				} else {
					allowList["AllowList"] = []string{}
				}
				userAllowList, err := ve.ObtainSdkValue("Result.UserAllowList", *resp)
				if err != nil {
					return data, err
				}
				if userAllowList != nil {
					allowList["UserAllowList"] = strings.Split(userAllowList.(string), ",")
				} else {
					allowList["UserAllowList"] = []string{}
				}
				bindInfos, _ := ve.ObtainSdkValue("Result.SecurityGroupBindInfos", *resp)
				if bindInfos != nil {
					allowList["SecurityGroupBindInfos"] = bindInfos
				}
			}
		}
		return data, err
	})
}

func (s *VolcengineRdsPostgresqlAllowlistService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
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
		return data, fmt.Errorf("rds_postgresql_allowlist %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineRdsPostgresqlAllowlistService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (VolcengineRdsPostgresqlAllowlistService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"VPC": {
				TargetField: "vpc",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineRdsPostgresqlAllowlistService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	// 如果设置了InstanceIds，则走合并白名单的流程
	if v, ok := resourceData.GetOk("instance_ids"); ok {
		callback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "UnifyNewAllowList",
				ConvertMode: ve.RequestConvertAll,
				ContentType: ve.ContentTypeJson,
				Convert: map[string]ve.RequestConvert{
					"allow_list_name":           {TargetField: "AllowListName"},
					"allow_list_desc":           {TargetField: "AllowListDesc"},
					"instance_ids":              {Ignore: true},
					"allow_list":                {Ignore: true},
					"user_allow_list":           {Ignore: true},
					"security_group_bind_infos": {Ignore: true},
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					set := v.(*schema.Set)
					var ids []string
					for _, id := range set.List() {
						ids = append(ids, id.(string))
					}
					(*call.SdkParam)["InstanceIds"] = ids
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
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
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateAllowList",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"allow_list_name":     {TargetField: "AllowListName"},
				"allow_list_desc":     {TargetField: "AllowListDesc"},
				"allow_list_type":     {TargetField: "AllowListType"},
				"allow_list_category": {TargetField: "AllowListCategory"},
				"security_group_bind_infos": {
					TargetField: "SecurityGroupBindInfos",
					ConvertType: ve.ConvertJsonObjectArray,
					NextLevelConvert: map[string]ve.RequestConvert{
						"security_group_id":   {TargetField: "SecurityGroupId"},
						"bind_mode":           {TargetField: "BindMode"},
						"ip_list":             {TargetField: "IpList", ConvertType: ve.ConvertJsonArray},
						"security_group_name": {TargetField: "SecurityGroupName"},
					},
				},
				"user_allow_list": {Ignore: true},
				"allow_list":      {Ignore: true},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				var allowStrings []string
				if v, ok := d.GetOk("allow_list"); ok {
					allowListsSet := v.(*schema.Set)
					for _, v := range allowListsSet.List() {
						allowStrings = append(allowStrings, v.(string))
					}
					if len(allowStrings) > 0 {
						allowLists := strings.Join(allowStrings, ",")
						(*call.SdkParam)["AllowList"] = allowLists
					}
				}
				if v, ok := d.GetOk("user_allow_list"); ok {
					s := v.(*schema.Set)
					var items []string
					for _, vv := range s.List() {
						items = append(items, vv.(string))
					}
					if len(items) > 0 {
						(*call.SdkParam)["UserAllowList"] = strings.Join(items, ",")
					}
				}
				logger.Debug(logger.ReqFormat, call.Action+" FullRequest", *call.SdkParam)
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
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

func (s *VolcengineRdsPostgresqlAllowlistService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
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
				},
				"allow_list_category": {TargetField: "AllowListCategory"},
				"security_group_bind_infos": {
					TargetField: "SecurityGroupBindInfos",
					ConvertType: ve.ConvertJsonObjectArray,
					NextLevelConvert: map[string]ve.RequestConvert{
						"security_group_id":   {TargetField: "SecurityGroupId"},
						"bind_mode":           {TargetField: "BindMode"},
						"ip_list":             {TargetField: "IpList", ConvertType: ve.ConvertJsonArray},
						"security_group_name": {TargetField: "SecurityGroupName"},
					},
				},
				"user_allow_list":       {Ignore: true},
				"update_security_group": {TargetField: "UpdateSecurityGroup"},
				"allow_list":            {Ignore: true},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				var allowStrings []string
				if v, ok := d.GetOk("allow_list"); ok {
					allowListsSet := v.(*schema.Set)
					for _, v := range allowListsSet.List() {
						allowStrings = append(allowStrings, v.(string))
					}
					if len(allowStrings) > 0 {
						allowLists := strings.Join(allowStrings, ",")
						logger.Debug(logger.ReqFormat, call.Action, call.SdkParam, allowLists)
						(*call.SdkParam)["AllowList"] = allowLists
					}
				}
				if v, ok := d.GetOk("user_allow_list"); ok {
					s := v.(*schema.Set)
					var items []string
					for _, vv := range s.List() {
						items = append(items, vv.(string))
					}
					if len(items) > 0 {
						(*call.SdkParam)["UserAllowList"] = strings.Join(items, ",")
					}
				}
				(*call.SdkParam)["ApplyInstanceNum"] = d.Get("associated_instance_num")
				(*call.SdkParam)["ModifyMode"] = "Cover"
				(*call.SdkParam)["AllowListId"] = d.Id()
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

func (s *VolcengineRdsPostgresqlAllowlistService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
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
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading rds postgre allow list on delete %q, %w", d.Id(), callErr))
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

func (s *VolcengineRdsPostgresqlAllowlistService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"allow_list_category": {
				TargetField: "AllowListCategory",
			},
			"allow_list_desc": {
				TargetField: "AllowListDesc",
			},
			"allow_list_id": {
				TargetField: "AllowListId",
			},
			"allow_list_name": {
				TargetField: "AllowListName",
			},
			"ip_address": {
				TargetField: "IPAddress",
			},
		},
		NameField:    "AllowListName",
		IdField:      "AllowListId",
		CollectField: "postgresql_allow_lists",
		ResponseConverts: map[string]ve.ResponseConvert{
			"AllowListId": {
				TargetField: "id",
				KeepDefault: true,
			},
			"AllowListIPNum": {
				TargetField: "allow_list_ip_num",
			},
			"VPC": {
				TargetField: "vpc",
			},
		},
	}
}

func (s *VolcengineRdsPostgresqlAllowlistService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "rds_postgresql",
		Version:     "2022-01-01",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}
