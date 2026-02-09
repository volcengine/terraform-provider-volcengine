package iam_login_profile

import (
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineIamLoginProfileService struct {
	Client *ve.SdkClient
}

func NewIamLoginProfileService(c *ve.SdkClient) *VolcengineIamLoginProfileService {
	return &VolcengineIamLoginProfileService{
		Client: c,
	}
}

func (s *VolcengineIamLoginProfileService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineIamLoginProfileService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		result interface{}
	)
	action := "GetLoginProfile"
	logger.Debug(logger.ReqFormat, action, m)
	resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(action), &m)
	if err != nil {
		return data, err
	}
	logger.Debug(logger.RespFormat, action, resp)
	result, err = ve.ObtainSdkValue("Result.LoginProfile", *resp)
	if err != nil {
		return data, err
	}
	if dataMap, ok := result.(map[string]interface{}); ok {
		delete(dataMap, "Password")
		data = append(data, dataMap)
	} else {
		return data, errors.New("Value is not map ")
	}

	return data, nil
}

func (s *VolcengineIamLoginProfileService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		result interface{}
		ok     bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	condition := map[string]interface{}{"UserName": id}
	action := "GetLoginProfile"
	logger.Debug(logger.ReqFormat, action, condition)
	resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(action), &condition)
	if err != nil {
		return data, err
	}
	logger.Debug(logger.RespFormat, action, resp)
	result, err = ve.ObtainSdkValue("Result.LoginProfile", *resp)
	if err != nil {
		return data, err
	}
	if data, ok = result.(map[string]interface{}); !ok {
		return data, errors.New("Value is not map ")
	}
	if len(data) == 0 {
		return data, fmt.Errorf("login profile %s not exist ", id)
	}

	return data, err
}

func (s *VolcengineIamLoginProfileService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (VolcengineIamLoginProfileService) WithResourceResponseHandlers(v map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		delete(v, "Password")
		return v, map[string]ve.ResponseConvert{}, nil
	}
	return []ve.ResourceResponseHandler{handler}

}

func (s *VolcengineIamLoginProfileService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateLoginProfile",
			ConvertMode: ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"user_name": {
					ForceGet: true,
				},
				"password": {
					ForceGet: true,
				},
				"login_allowed": {
					TargetField: "LoginAllowed",
				},
				"password_reset_required": {
					TargetField: "PasswordResetRequired",
				},
				"safe_auth_flag": {
					TargetField: "SafeAuthFlag",
				},
				"safe_auth_type": {
					TargetField: "SafeAuthType",
				},
				"safe_auth_exempt_required": {
					TargetField: "SafeAuthExemptRequired",
				},
				"safe_auth_exempt_unit": {
					TargetField: "SafeAuthExemptUnit",
				},
				"safe_auth_exempt_duration": {
					TargetField: "SafeAuthExemptDuration",
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam, resp)
				time.Sleep(5 * time.Second)
				d.SetId(d.Get("user_name").(string))
				return nil
			},
		},
	}

	return []ve.Callback{callback}
}

func (s *VolcengineIamLoginProfileService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateLoginProfile",
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"password": {
					ForceGet:    true,
					TargetField: "Password",
				},
				"login_allowed": {
					ForceGet:    true,
					TargetField: "LoginAllowed",
				},
				"password_reset_required": {
					ForceGet:    true,
					TargetField: "PasswordResetRequired",
				},
				"safe_auth_flag": {
					ForceGet:    true,
					TargetField: "SafeAuthFlag",
				},
				"safe_auth_type": {
					ForceGet:    true,
					TargetField: "SafeAuthType",
				},
				"safe_auth_exempt_required": {
					ForceGet:    true,
					TargetField: "SafeAuthExemptRequired",
				},
				"safe_auth_exempt_unit": {
					ForceGet:    true,
					TargetField: "SafeAuthExemptUnit",
				},
				"safe_auth_exempt_duration": {
					ForceGet:    true,
					TargetField: "SafeAuthExemptDuration",
				},
			},
			RequestIdField: "UserName",
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				time.Sleep(5 * time.Second)
				return resp, err
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineIamLoginProfileService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:         "DeleteLoginProfile",
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

func (s *VolcengineIamLoginProfileService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"user_name": {
				TargetField: "UserName",
			},
		},
		NameField:    "UserName",
		IdField:      "UserName",
		CollectField: "login_profiles",
		ResponseConverts: map[string]ve.ResponseConvert{
			"SafeAuthFlag": {
				TargetField: "safe_auth_flag",
			},
			"SafeAuthType": {
				TargetField: "safe_auth_type",
			},
			"SafeAuthExemptRequired": {
				TargetField: "safe_auth_exempt_required",
			},
			"SafeAuthExemptUnit": {
				TargetField: "safe_auth_exempt_unit",
			},
			"SafeAuthExemptDuration": {
				TargetField: "safe_auth_exempt_duration",
			},
		},
	}
}

func (s *VolcengineIamLoginProfileService) ReadResourceId(id string) string {
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
