package iam_user_group_policy_attachment

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

type VolcengineIamUserGroupPolicyAttachmentService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewIamUserGroupPolicyAttachmentService(c *ve.SdkClient) *VolcengineIamUserGroupPolicyAttachmentService {
	return &VolcengineIamUserGroupPolicyAttachmentService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineIamUserGroupPolicyAttachmentService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineIamUserGroupPolicyAttachmentService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithSimpleQuery(condition, func(m map[string]interface{}) ([]interface{}, error) {
		action := "ListAttachedUserGroupPolicies"
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
		results, err = ve.ObtainSdkValue("Result.AttachedPolicyMetadata", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.AttachedPolicyMetadata is not slice")
		}
		return data, err
	})
}

func (s *VolcengineIamUserGroupPolicyAttachmentService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		tempData = map[string]interface{}{}
		results  []interface{}
		ok       bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	ids := strings.Split(id, ":")
	req := map[string]interface{}{
		"UserGroupName": ids[0],
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		if tempData, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("Value is not map ")
		} else if tempData["PolicyName"].(string) == ids[1] {
			data = tempData
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("iam_user_group_policy_attachment %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineIamUserGroupPolicyAttachmentService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineIamUserGroupPolicyAttachmentService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "AttachUserGroupPolicy",
			ConvertMode: ve.RequestConvertAll,
			Convert:     map[string]ve.RequestConvert{},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				d.SetId(fmt.Sprintf("%s:%s", d.Get("user_group_name").(string), d.Get("policy_name").(string)))
				return nil
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("user_group_name").(string)
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineIamUserGroupPolicyAttachmentService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineIamUserGroupPolicyAttachmentService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineIamUserGroupPolicyAttachmentService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	ids := strings.Split(resourceData.Id(), ":")
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DetachUserGroupPolicy",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"UserGroupName": ids[0],
				"PolicyName":    ids[1],
				"PolicyType":    resourceData.Get("policy_type").(string),
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

func (s *VolcengineIamUserGroupPolicyAttachmentService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		NameField:    "PolicyName",
		IdField:      "PolicyTrn",
		CollectField: "policies",
	}
}

func (s *VolcengineIamUserGroupPolicyAttachmentService) ReadResourceId(id string) string {
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
