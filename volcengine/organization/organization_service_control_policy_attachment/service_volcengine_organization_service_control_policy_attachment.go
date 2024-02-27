package organization_service_control_policy_attachment

import (
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
	"strings"
	"time"
)

type VolcengineServiceControlPolicyAttachmentService struct {
	Client *ve.SdkClient
}

func NewServiceControlPolicyAttachmentService(c *ve.SdkClient) *VolcengineServiceControlPolicyAttachmentService {
	return &VolcengineServiceControlPolicyAttachmentService{
		Client: c,
	}
}

func (s *VolcengineServiceControlPolicyAttachmentService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineServiceControlPolicyAttachmentService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	action := "ListTargetAttachmentsForServiceControlPolicy"
	logger.Debug(logger.ReqFormat, action, m)
	if m == nil {
		resp, err = s.Client.UniversalClient.DoCall(postUniversalInfo(action), nil)
		if err != nil {
			return data, err
		}
	} else {
		resp, err = s.Client.UniversalClient.DoCall(postUniversalInfo(action), &m)
		if err != nil {
			return data, err
		}
	}

	logger.Debug(logger.RespFormat, action, m, *resp)

	results, err = ve.ObtainSdkValue("Result", *resp)
	if err != nil {
		return data, err
	}
	if results == nil {
		results = []interface{}{}
	}
	if data, ok = results.([]interface{}); !ok {
		return data, errors.New(" Result is not Slice")
	}
	return data, err
}

func (s *VolcengineServiceControlPolicyAttachmentService) ReadResource(resourceData *schema.ResourceData, roleId string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if roleId == "" {
		roleId = s.ReadResourceId(resourceData.Id())
	}
	ids := strings.Split(roleId, ":")
	if len(ids) != 2 {
		return data, fmt.Errorf("import id is invalid")
	}
	req := map[string]interface{}{
		"PolicyID": ids[0],
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		if data, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("value is not map")
		} else if ids[1] == data["TargetID"].(string) {
			return data, err
		}
	}
	return data, fmt.Errorf("service control policy attachment %s not exist ", roleId)
}

func (s *VolcengineServiceControlPolicyAttachmentService) RefreshResourceState(data *schema.ResourceData, strings []string, duration time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (s *VolcengineServiceControlPolicyAttachmentService) WithResourceResponseHandlers(rolePolicyAttachment map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return rolePolicyAttachment, map[string]ve.ResponseConvert{
			"TargetID": {
				TargetField: "target_id",
			},
			"PolicyID": {
				TargetField: "policy_id",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineServiceControlPolicyAttachmentService) CreateResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	createPolicyAttachmentCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "AttachServiceControlPolicy",
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"policy_id": {
					TargetField: "PolicyID",
				},
				"target_id": {
					TargetField: "TargetID",
				},
				"target_type": {
					Convert: func(data *schema.ResourceData, old interface{}) interface{} {
						ty := 0
						switch old.(string) {
						case "OU":
							ty = 1
						case "Account":
							ty = 2
						}
						return ty
					},
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(postUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				d.SetId(fmt.Sprintf("%s:%s", d.Get("policy_id"), d.Get("target_id")))
				return nil
			},
			LockId: func(d *schema.ResourceData) string {
				return "lock-Organization"
			},
		},
	}
	return []ve.Callback{createPolicyAttachmentCallback}
}

func (s *VolcengineServiceControlPolicyAttachmentService) ModifyResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineServiceControlPolicyAttachmentService) RemoveResource(data *schema.ResourceData, r *schema.Resource) []ve.Callback {
	deleteRoleCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DetachServiceControlPolicy",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				ids := strings.Split(d.Id(), ":")
				(*call.SdkParam)["PolicyID"] = ids[0]
				(*call.SdkParam)["TargetID"] = ids[1]
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(postUniversalInfo(call.Action), call.SdkParam)
			},
			LockId: func(d *schema.ResourceData) string {
				return "lock-Organization"
			},
		},
	}
	return []ve.Callback{deleteRoleCallback}
}

func (s *VolcengineServiceControlPolicyAttachmentService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineServiceControlPolicyAttachmentService) ReadResourceId(id string) string {
	return id
}

func postUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "organization",
		Version:     "2022-01-01",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}
