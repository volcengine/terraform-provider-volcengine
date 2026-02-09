package iam_security_config

import (
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineIamSecurityConfigService struct {
	Client *ve.SdkClient
}

func NewIamSecurityConfigService(c *ve.SdkClient) *VolcengineIamSecurityConfigService {
	return &VolcengineIamSecurityConfigService{
		Client: c,
	}
}

func (s *VolcengineIamSecurityConfigService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineIamSecurityConfigService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		result interface{}
	)
	action := "GetSecurityConfig"
	logger.Debug(logger.ReqFormat, action, m)
	resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(action), &m)
	if err != nil {
		return data, err
	}
	logger.Debug(logger.RespFormat, action, resp)
	result, err = ve.ObtainSdkValue("Result", *resp)
	if err != nil {
		return data, err
	}
	if dataMap, ok := result.(map[string]interface{}); ok {
		// API returns UserID, SafeAuthType, etc. We need to inject UserName because API doesn't return it in Result (only in request)
		// Wait, ReadResources is usually for List. GetSecurityConfig is singular.
		// If m contains UserName, we can inject it back.
		if userName, ok := m["UserName"]; ok {
			dataMap["UserName"] = userName
		}
		data = append(data, dataMap)
	} else {
		return data, errors.New("Value is not map ")
	}

	return data, nil
}

func (s *VolcengineIamSecurityConfigService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		result interface{}
		ok     bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	condition := map[string]interface{}{"UserName": id}
	action := "GetSecurityConfig"
	logger.Debug(logger.ReqFormat, action, condition)
	resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(action), &condition)
	if err != nil {
		return data, err
	}
	logger.Debug(logger.RespFormat, action, resp)
	result, err = ve.ObtainSdkValue("Result", *resp)
	if err != nil {
		return data, err
	}
	if data, ok = result.(map[string]interface{}); !ok {
		return data, errors.New("Value is not map ")
	}
	if len(data) == 0 {
		return data, fmt.Errorf("security config for user %s not exist ", id)
	}

	return data, err
}

func (s *VolcengineIamSecurityConfigService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (VolcengineIamSecurityConfigService) WithResourceResponseHandlers(v map[string]interface{}) []ve.ResourceResponseHandler {
	return []ve.ResourceResponseHandler{}
}

func (s *VolcengineIamSecurityConfigService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "SetSecurityConfig",
			ConvertMode: ve.RequestConvertAll,
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				d.SetId(d.Get("user_name").(string))
				return nil
			},
		},
	}

	return []ve.Callback{callback}
}

func (s *VolcengineIamSecurityConfigService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:         "SetSecurityConfig",
			ConvertMode:    ve.RequestConvertAll,
			RequestIdField: "UserName",
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineIamSecurityConfigService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	// No delete API available.
	return []ve.Callback{}
}

func (s *VolcengineIamSecurityConfigService) DatasourceResources(resourceData *schema.ResourceData, r *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"user_name": {
				TargetField: "UserName",
			},
		},
		NameField:    "UserName",
		IdField:      "UserName",
		CollectField: "security_configs",
		ResponseConverts: map[string]ve.ResponseConvert{
			"UserID": {
				TargetField: "user_id",
			},
			"SafeAuthType": {
				TargetField: "safe_auth_type",
			},
			"SafeAuthExemptDuration": {
				TargetField: "safe_auth_exempt_duration",
			},
			"SafeAuthClose": {
				TargetField: "safe_auth_close",
			},
		},
	}
}

func (s *VolcengineIamSecurityConfigService) ReadResourceId(id string) string {
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
