package kms_asymmetric_ciphertext

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineKmsAsymmetricCiphertextService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewKmsAsymmetricCiphertextService(c *ve.SdkClient) *VolcengineKmsAsymmetricCiphertextService {
	return &VolcengineKmsAsymmetricCiphertextService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineKmsAsymmetricCiphertextService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineKmsAsymmetricCiphertextService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp *map[string]interface{}
	)
	return ve.WithSimpleQuery(m, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "AsymmetricEncrypt"

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
		ciphertext, _ := ve.ObtainSdkValue("Result.CiphertextBlob", *resp)
		if ciphertext != nil {
			result["CiphertextBlob"] = ciphertext
		}
		data = append(data, result)
		return data, err
	})
}

func (s *VolcengineKmsAsymmetricCiphertextService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	return data, err
}

func (s *VolcengineKmsAsymmetricCiphertextService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				d          map[string]interface{}
				status     interface{}
				failStates []string
			)
			failStates = append(failStates, "Failed")
			d, err = s.ReadResource(resourceData, id)
			if err != nil {
				return nil, "", err
			}
			status, err = ve.ObtainSdkValue("Status", d)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("kms_asymmetric_ciphertext status error, status: %s", status.(string))
				}
			}
			return d, status.(string), err
		},
	}
}

func (s *VolcengineKmsAsymmetricCiphertextService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "AsymmetricEncrypt",
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"keyring_name": {
					TargetField: "KeyringName",
				},
				"key_name": {
					TargetField: "KeyName",
				},
				"key_id": {
					TargetField: "KeyID",
				},
				"plaintext": {
					TargetField: "Plaintext",
				},
				"algorithm": {
					TargetField: "Algorithm",
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				var keyId string
				if v, ok := d.GetOk("key_id"); ok {
					keyId = v.(string)
				} else {
					keyId = fmt.Sprintf("%s:%s", d.Get("key_name").(string), d.Get("keyring_name").(string))
				}
				d.SetId(keyId)
				ciphertext, _ := ve.ObtainSdkValue("Result.CiphertextBlob", *resp)
				if ciphertext != nil {
					d.Set("ciphertext_blob", ciphertext)
				}
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineKmsAsymmetricCiphertextService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineKmsAsymmetricCiphertextService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineKmsAsymmetricCiphertextService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineKmsAsymmetricCiphertextService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"key_id": {
				TargetField: "KeyID",
			},
		},
		CollectField:     "ciphertext_info",
		ResponseConverts: map[string]ve.ResponseConvert{},
	}
}

func (s *VolcengineKmsAsymmetricCiphertextService) ReadResourceId(id string) string {
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
