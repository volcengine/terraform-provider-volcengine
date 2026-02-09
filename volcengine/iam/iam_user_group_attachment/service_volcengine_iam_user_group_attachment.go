package iam_user_group_attachment

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

type VolcengineIamUserGroupAttachmentService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewIamUserGroupAttachmentService(c *ve.SdkClient) *VolcengineIamUserGroupAttachmentService {
	return &VolcengineIamUserGroupAttachmentService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineIamUserGroupAttachmentService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineIamUserGroupAttachmentService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageOffsetQuery(m, "Limit", "Offset", 100, 0, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListGroupsForUser"
		if _, ok := condition["UserGroupName"]; ok {
			action = "ListUsersForGroup"
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
		if action == "ListUsersForGroup" {
			results, err = ve.ObtainSdkValue("Result.Users", *resp)
		} else {
			results, err = ve.ObtainSdkValue("Result.UserGroups", *resp)
		}
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			if action == "ListUsersForGroup" {
				return data, errors.New("Result.Users is not Slice")
			}
			return data, errors.New("Result.UserGroups is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineIamUserGroupAttachmentService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results  []interface{}
		tempData map[string]interface{}
		ok       bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	ids := strings.Split(id, ":")
	req := map[string]interface{}{
		"UserName": ids[1],
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		if tempData, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("Value is not map ")
		} else {
			if tempData["UserGroupName"].(string) == ids[0] {
				data = tempData
			}
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("iam_user_group_attachment %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineIamUserGroupAttachmentService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineIamUserGroupAttachmentService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "AddUserToGroup",
			ConvertMode: ve.RequestConvertAll,
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				d.SetId(fmt.Sprint((*call.SdkParam)["UserGroupName"], ":", (*call.SdkParam)["UserName"]))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineIamUserGroupAttachmentService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineIamUserGroupAttachmentService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineIamUserGroupAttachmentService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	ids := strings.Split(resourceData.Id(), ":")
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "RemoveUserFromGroup",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"UserGroupName": ids[0],
				"UserName":      ids[1],
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

func (s *VolcengineIamUserGroupAttachmentService) DatasourceResources(d *schema.ResourceData, r *schema.Resource) ve.DataSourceInfo {
	if _, ok := d.GetOk("user_group_name"); ok {
		return ve.DataSourceInfo{
			RequestConverts: map[string]ve.RequestConvert{
				"user_group_name": {
					TargetField: "UserGroupName",
				},
			},
			NameField:    "UserName",
			IdField:      "UserName",
			CollectField: "users",
			ResponseConverts: map[string]ve.ResponseConvert{
				"Id": {
					TargetField: "user_id",
				},
				"UserName": {
					TargetField: "user_name",
				},
				"DisplayName": {
					TargetField: "display_name",
				},
				"Description": {
					TargetField: "description",
				},
				"JoinDate": {
					TargetField: "join_date",
				},
			},
		}
	}
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"user_name": {
				TargetField: "UserName",
			},
			"query": {
				TargetField: "Query",
			},
		},
		NameField:    "UserGroupName",
		IdField:      "UserGroupName",
		CollectField: "user_groups",
		ResponseConverts: map[string]ve.ResponseConvert{
			"UserGroupID": {
				TargetField: "user_group_id",
			},
			"UserGroupName": {
				TargetField: "user_group_name",
			},
			"DisplayName": {
				TargetField: "display_name",
			},
			"Description": {
				TargetField: "description",
			},
			"JoinDate": {
				TargetField: "join_date",
			},
		},
	}
}

func (s *VolcengineIamUserGroupAttachmentService) ReadResourceId(id string) string {
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
