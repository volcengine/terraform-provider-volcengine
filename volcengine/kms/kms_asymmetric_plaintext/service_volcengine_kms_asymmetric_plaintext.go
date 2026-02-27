package kms_asymmetric_plaintext

import (
	"encoding/json"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineKmsAsymmetricPlaintextService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewKmsAsymmetricPlaintextService(c *ve.SdkClient) *VolcengineKmsAsymmetricPlaintextService {
	return &VolcengineKmsAsymmetricPlaintextService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineKmsAsymmetricPlaintextService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineKmsAsymmetricPlaintextService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp *map[string]interface{}
	)
	return ve.WithSimpleQuery(m, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "AsymmetricDecrypt"
		// 安全考虑，不打印请求中的密文信息，避免将密文信息写入本地日志
		logParam := make(map[string]interface{}, len(condition))
		for k, v := range condition {
			if k == "CiphertextBlob" {
				logParam[k] = "******"
			} else {
				logParam[k] = v
			}
		}
		bytes, _ := json.Marshal(logParam)
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
		// 安全考虑，不打印响应中的明文信息
		// respBytes, _ := json.Marshal(resp)
		// logger.Debug(logger.RespFormat, action, condition, string(respBytes))
		result := make(map[string]interface{})
		plaintext, _ := ve.ObtainSdkValue("Result.Plaintext", *resp)
		if plaintext != nil {
			result["Plaintext"] = plaintext
		}
		data = append(data, result)
		return data, err
	})
}

func (s *VolcengineKmsAsymmetricPlaintextService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	return data, err
}

func (s *VolcengineKmsAsymmetricPlaintextService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineKmsAsymmetricPlaintextService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (VolcengineKmsAsymmetricPlaintextService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineKmsAsymmetricPlaintextService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineKmsAsymmetricPlaintextService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineKmsAsymmetricPlaintextService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"key_id": {
				TargetField: "KeyID",
			},
		},
		CollectField:     "plaintext_info",
		ResponseConverts: map[string]ve.ResponseConvert{},
	}
}

func (s *VolcengineKmsAsymmetricPlaintextService) ReadResourceId(id string) string {
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
