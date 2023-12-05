package organization_service_control_policy_enabler

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineOrganizationServiceControlPolicyEnablerService struct {
	Client *ve.SdkClient
}

func (s *VolcengineOrganizationServiceControlPolicyEnablerService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	return nil, nil
}

func (s *VolcengineOrganizationServiceControlPolicyEnablerService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	universalClient := s.Client.UniversalClient
	action := "GetServiceControlPolicyEnablement"
	resp, err := universalClient.DoCall(getUniversalInfo(action), &map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	logger.Debug(logger.ReqFormat, action, *resp)
	status, err := ve.ObtainSdkValue("Result.Status", *resp)
	if err != nil {
		return nil, err
	}

	if status != "Enabled" {
		return data, fmt.Errorf(" Organization Service Control Policy is not Enabled")
	}
	return map[string]interface{}{
		"Status": status,
	}, nil
}

func (s *VolcengineOrganizationServiceControlPolicyEnablerService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (s *VolcengineOrganizationServiceControlPolicyEnablerService) WithResourceResponseHandlers(m map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return m, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineOrganizationServiceControlPolicyEnablerService) CreateResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "EnableServiceControlPolicy",
			ConvertMode: ve.RequestConvertIgnore,
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				// 先检查是否开启
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo("GetServiceControlPolicyEnablement"), &map[string]interface{}{})
				if err != nil {
					return nil, err
				}
				logger.Debug(logger.ReqFormat, "GetServiceControlPolicyEnablement", *resp)
				status, err := ve.ObtainSdkValue("Result.Status", *resp)
				if err != nil {
					return nil, err
				}
				if status == "Enabled" {
					return nil, nil // 不需要再重复开启了
				}

				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				d.SetId("organization:service_control_policy_enable")
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineOrganizationServiceControlPolicyEnablerService) ModifyResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineOrganizationServiceControlPolicyEnablerService) RemoveResource(data *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DisableServiceControlPolicy",
			ConvertMode: ve.RequestConvertIgnore,
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineOrganizationServiceControlPolicyEnablerService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineOrganizationServiceControlPolicyEnablerService) ReadResourceId(id string) string {
	return id
}

func NewService(client *ve.SdkClient) *VolcengineOrganizationServiceControlPolicyEnablerService {
	return &VolcengineOrganizationServiceControlPolicyEnablerService{
		Client: client,
	}
}

func (s *VolcengineOrganizationServiceControlPolicyEnablerService) GetClient() *ve.SdkClient {
	return s.Client
}

func getUniversalInfo(action string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "organization",
		Action:      action,
		Version:     "2022-01-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
	}
}
