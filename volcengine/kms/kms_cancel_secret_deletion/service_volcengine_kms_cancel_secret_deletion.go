package kms_cancel_secret_deletion

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineKmsCancelSecretDeletionService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewKmsCancelSecretDeletionService(c *ve.SdkClient) *VolcengineKmsCancelSecretDeletionService {
	return &VolcengineKmsCancelSecretDeletionService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineKmsCancelSecretDeletionService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineKmsCancelSecretDeletionService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	return nil, nil
}

func (s *VolcengineKmsCancelSecretDeletionService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		ok   bool
		resp *map[string]interface{}
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"SecretName": id,
	}
	action := "DescribeSecret"
	resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &req)
	if err != nil {
		return data, err
	}
	logger.Debug(logger.RespFormat, action, req, *resp)

	secret, err := ve.ObtainSdkValue("Result.Secret", *resp)
	if err != nil {
		return data, err
	}
	data, ok = secret.(map[string]interface{})
	if !ok {
		return data, fmt.Errorf(" Result is not Map ")
	}
	if len(data) == 0 {
		return data, fmt.Errorf("kms cancel secret deletion %s not exist", id)
	}
	if _, exists := data["SecretName"]; !exists {
		data["SecretName"] = id
	}
	return data, err
}

func (s *VolcengineKmsCancelSecretDeletionService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineKmsCancelSecretDeletionService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CancelSecretDeletion",
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"secret_name": {
					TargetField: "SecretName",
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				d.SetId(resourceData.Get("secret_name").(string))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineKmsCancelSecretDeletionService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineKmsCancelSecretDeletionService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineKmsCancelSecretDeletionService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineKmsCancelSecretDeletionService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineKmsCancelSecretDeletionService) ReadResourceId(id string) string {
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
