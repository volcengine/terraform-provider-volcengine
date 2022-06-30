package iam_role_policy_attachment

import (
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-vestack/common"
	"github.com/volcengine/terraform-provider-vestack/logger"
	"strings"
	"time"
)

type VestackIamRolePolicyAttachmentService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewIamRolePolicyAttachmentService(c *ve.SdkClient) *VestackIamRolePolicyAttachmentService {
	return &VestackIamRolePolicyAttachmentService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VestackIamRolePolicyAttachmentService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VestackIamRolePolicyAttachmentService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	action := "ListAttachedRolePolicies"
	logger.Debug(logger.ReqFormat, action, m)
	if m == nil {
		resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), nil)
		if err != nil {
			return data, err
		}
	} else {
		resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &m)
		if err != nil {
			return data, err
		}
	}

	logger.Debug(logger.RespFormat, action, m, *resp)

	results, err = ve.ObtainSdkValue("Result.AttachedPolicyMetadata", *resp)
	if err != nil {
		return data, err
	}
	if results == nil {
		results = []interface{}{}
	}
	if data, ok = results.([]interface{}); !ok {
		return data, errors.New("Result.AttachedPolicyMetadata is not Slice")
	}
	return data, err
}

func (s *VestackIamRolePolicyAttachmentService) ReadResource(resourceData *schema.ResourceData, roleId string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if roleId == "" {
		roleId = s.ReadResourceId(resourceData.Id())
	}
	ids := strings.Split(roleId, ":")
	if len(ids) != 3 {
		return data, fmt.Errorf("import id is invalid")
	}
	req := map[string]interface{}{
		"RoleName": ids[0],
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		if data, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("value is not map")
		} else if ids[1] == data["PolicyName"].(string) && ids[2] == data["PolicyType"].(string) {
			return data, err
		}
	}
	return data, fmt.Errorf("Role policy attachment %s not exist ", roleId)
}

func (s *VestackIamRolePolicyAttachmentService) RefreshResourceState(data *schema.ResourceData, strings []string, duration time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (s *VestackIamRolePolicyAttachmentService) WithResourceResponseHandlers(rolePolicyAttachment map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return rolePolicyAttachment, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VestackIamRolePolicyAttachmentService) CreateResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	createIamRolePolicyAttachmentCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "AttachRolePolicy",
			ConvertMode: ve.RequestConvertAll,
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				d.SetId(fmt.Sprintf("%s:%s:%s", d.Get("role_name").(string),
					d.Get("policy_name").(string), d.Get("policy_type").(string)))
				return nil
			},
		},
	}
	return []ve.Callback{createIamRolePolicyAttachmentCallback}
}

func (s *VestackIamRolePolicyAttachmentService) ModifyResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VestackIamRolePolicyAttachmentService) RemoveResource(data *schema.ResourceData, r *schema.Resource) []ve.Callback {
	deleteRoleCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DetachRolePolicy",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				ids := strings.Split(d.Id(), ":")
				(*call.SdkParam)["RoleName"] = ids[0]
				(*call.SdkParam)["PolicyName"] = ids[1]
				(*call.SdkParam)["PolicyType"] = ids[2]
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading iam role on delete %q, %w", d.Id(), callErr))
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
	return []ve.Callback{deleteRoleCallback}
}

func (s *VestackIamRolePolicyAttachmentService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VestackIamRolePolicyAttachmentService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "iam",
		Version:     "2018-01-01",
		HttpMethod:  ve.GET,
		Action:      actionName,
	}
}
