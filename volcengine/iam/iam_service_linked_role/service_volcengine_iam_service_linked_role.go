package iam_service_linked_role

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineIamServiceLinkedRoleService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewIamServiceLinkedRoleService(c *ve.SdkClient) *VolcengineIamServiceLinkedRoleService {
	return &VolcengineIamServiceLinkedRoleService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineIamServiceLinkedRoleService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineIamServiceLinkedRoleService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	return data, err
}

func (s *VolcengineIamServiceLinkedRoleService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		ok bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	ids := strings.Split(id, ":")
	if len(ids) != 2 {
		return data, fmt.Errorf("invalid id %s, expect format service_name:role_name", id)
	}
	serviceName := ids[0]
	roleName := ids[1]

	data = map[string]interface{}{}
	data["ServiceName"] = serviceName

	roleAction := "GetRole"
	roleReq := map[string]interface{}{
		"RoleName": roleName,
	}
	logger.Debug(logger.RespFormat, roleAction, roleReq)
	roleResp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(roleAction), &roleReq)
	logger.Debug(logger.RespFormat, roleAction, roleResp, err)
	if err != nil {
		return data, err
	}
	result, err := ve.ObtainSdkValue("Result.Role", *roleResp)
	if err != nil {
		return data, err
	}
	role, ok := result.(map[string]interface{})
	if !ok {
		return data, fmt.Errorf("Result.Role is not map")
	}
	data["Status"] = role["Status"]
	data["RoleId"] = role["RoleId"]

	if len(data) == 0 {
		return data, fmt.Errorf("iam_service_linked_role %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineIamServiceLinkedRoleService) queryRoleName(resourceData *schema.ResourceData, serviceName string) (string, error) {
	action := "GetServiceLinkedRoleTemplate"
	req := map[string]interface{}{
		"ServiceName": serviceName,
	}
	logger.Debug(logger.RespFormat, action, req)
	resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(action), &req)
	logger.Debug(logger.RespFormat, action, resp, err)
	if err != nil {
		return "", fmt.Errorf("GetServiceLinkedRoleTemplate failed, err: %v", err)
	}
	result, err := ve.ObtainSdkValue("Result.ServiceLinkedRoleTemplate", *resp)
	if err != nil {
		return "", fmt.Errorf("ObtainSdkValue Result.ServiceLinkedRoleTemplate failed, err: %v", err)
	}
	templates, ok := result.([]interface{})
	if !ok {
		return "", fmt.Errorf("Result.ServiceLinkedRoleTemplate is not slice")
	}
	if len(templates) == 0 {
		return "", fmt.Errorf("role name for %s is not exist ", serviceName)
	}
	template, ok := templates[0].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("Result.ServiceLinkedRoleTemplate[0] is not map")
	}
	roleName, ok := template["RoleName"].(string)
	if !ok {
		return "", fmt.Errorf("Result.ServiceLinkedRoleTemplate[0].RoleName is not string")
	}
	return roleName, nil
}

func (s *VolcengineIamServiceLinkedRoleService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (VolcengineIamServiceLinkedRoleService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineIamServiceLinkedRoleService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateServiceLinkedRole",
			ConvertMode: ve.RequestConvertAll,
			Convert:     map[string]ve.RequestConvert{},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				roleName, err := s.queryRoleName(resourceData, d.Get("service_name").(string))
				if err != nil {
					return false, err
				}
				d.Set("role_name", roleName)
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				serviceName := d.Get("service_name").(string)
				roleName := d.Get("role_name").(string)
				d.SetId(serviceName + ":" + roleName)
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineIamServiceLinkedRoleService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineIamServiceLinkedRoleService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteServiceLinkedRole",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				ids := strings.Split(d.Id(), ":")
				if len(ids) != 2 {
					return false, fmt.Errorf("id %s is invalid", d.Id())
				}
				(*call.SdkParam)["RoleName"] = ids[1]
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(5*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading iam service linked role on delete %q, %w", d.Id(), callErr))
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

func (s *VolcengineIamServiceLinkedRoleService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineIamServiceLinkedRoleService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "iam",
		Version:     "2018-01-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		Action:      actionName,
		RegionType:  ve.Global,
	}
}
