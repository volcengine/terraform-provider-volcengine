package kms_asymmetric_verification

import (
	"encoding/json"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineKmsAsymmetricVerificationService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewKmsAsymmetricVerificationService(c *ve.SdkClient) *VolcengineKmsAsymmetricVerificationService {
	return &VolcengineKmsAsymmetricVerificationService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineKmsAsymmetricVerificationService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineKmsAsymmetricVerificationService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp *map[string]interface{}
	)
	return ve.WithSimpleQuery(m, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "AsymmetricVerify"

		bytes, _ := json.Marshal(condition)
		logger.Debug(logger.ReqFormat, action, string(bytes))
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
		respBytes, _ := json.Marshal(resp)
		logger.Debug(logger.RespFormat, action, condition, string(respBytes))
		result := make(map[string]interface{})
		signatureValid, _ := ve.ObtainSdkValue("Result.SignatureValid", *resp)
		if signatureValid != nil {
			result["SignatureValid"] = signatureValid
		}
		data = append(data, result)
		return data, err
	})
}

func (s *VolcengineKmsAsymmetricVerificationService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	return data, err
}

func (s *VolcengineKmsAsymmetricVerificationService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineKmsAsymmetricVerificationService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (VolcengineKmsAsymmetricVerificationService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineKmsAsymmetricVerificationService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineKmsAsymmetricVerificationService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineKmsAsymmetricVerificationService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"key_id": {
				TargetField: "KeyID",
			},
		},
		CollectField:     "verification_info",
		ResponseConverts: map[string]ve.ResponseConvert{},
	}
}

func (s *VolcengineKmsAsymmetricVerificationService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "kms",
		Version:     "2021-02-18",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}
