package kms_secret_rotate

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineKmsSecretRotateService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewKmsSecretRotateService(c *ve.SdkClient) *VolcengineKmsSecretRotateService {
	return &VolcengineKmsSecretRotateService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineKmsSecretRotateService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineKmsSecretRotateService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	return nil, nil
}

func (s *VolcengineKmsSecretRotateService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	return nil, nil
}

func (s *VolcengineKmsSecretRotateService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			return nil, "", err
		},
	}
}

func (s *VolcengineKmsSecretRotateService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "RotateSecret",
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"secret_name": {
					TargetField: "SecretName",
				},
				"version_name": {
					TargetField: "VersionName",
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				d.SetId(d.Get("secret_name").(string))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineKmsSecretRotateService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	return nil
}

func (s *VolcengineKmsSecretRotateService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	// 当 VersionName 发生改变时，也认为是要执行手动轮转
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "RotateSecret",
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"secret_name": {
					TargetField: "SecretName",
					ForceGet:    true,
				},
				"version_name": {
					TargetField: "VersionName",
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				d.SetId(d.Get("secret_name").(string))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineKmsSecretRotateService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return nil
}

func (s *VolcengineKmsSecretRotateService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineKmsSecretRotateService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "kms",
		Version:     "2021-02-18",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		Action:      actionName,
	}
}
