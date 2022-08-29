package iam_role

import (
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineIamRoleService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewIamRoleService(c *ve.SdkClient) *VolcengineIamRoleService {
	return &VolcengineIamRoleService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineIamRoleService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineIamRoleService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageOffsetQuery(m, "Limit", "Offset", 100, 0, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListRoles"
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

		results, err = ve.ObtainSdkValue("Result.RoleMetadata", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.RoleMetadata is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineIamRoleService) ReadResource(resourceData *schema.ResourceData, roleId string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if roleId == "" {
		roleId = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"RoleName": roleId,
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		if data, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("value is not map")
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("Role %s not exist ", roleId)
	}
	return data, err
}

func (s *VolcengineIamRoleService) RefreshResourceState(data *schema.ResourceData, strings []string, duration time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (s *VolcengineIamRoleService) WithResourceResponseHandlers(role map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		role["Id"] = role["RoleName"]
		return role, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineIamRoleService) CreateResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	createIamRoleCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateRole",
			ConvertMode: ve.RequestConvertAll,
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				roleName, err := ve.ObtainSdkValue("Result.Role.RoleName", *resp)
				if err != nil {
					return err
				}
				d.SetId(roleName.(string))
				return nil
			},
		},
	}
	return []ve.Callback{createIamRoleCallback}
}

func (s *VolcengineIamRoleService) ModifyResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	updateRoleCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateRole",
			ConvertMode: ve.RequestConvertAll,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["RoleName"] = d.Get("role_name")
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	return []ve.Callback{updateRoleCallback}
}

func (s *VolcengineIamRoleService) RemoveResource(data *schema.ResourceData, r *schema.Resource) []ve.Callback {
	deleteRoleCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteRole",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["RoleName"] = d.Id()
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

func (s *VolcengineIamRoleService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		ResponseConverts: map[string]ve.ResponseConvert{
			"RoleName": {
				TargetField: "id",
				KeepDefault: true,
			},
		},
		NameField:    "RoleName",
		IdField:      "RoleName",
		CollectField: "roles",
	}
}

func (s *VolcengineIamRoleService) ReadResourceId(id string) string {
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
