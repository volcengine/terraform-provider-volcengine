package kms_re_encrypt

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineKmsReEncryptService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewKmsReEncryptService(c *ve.SdkClient) *VolcengineKmsReEncryptService {
	return &VolcengineKmsReEncryptService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineKmsReEncryptService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineKmsReEncryptService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp *map[string]interface{}
	)
	return ve.WithSimpleQuery(m, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ReEncrypt"

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

func (s *VolcengineKmsReEncryptService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	return data, err
}

func (s *VolcengineKmsReEncryptService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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
					return nil, "", fmt.Errorf("kms_re_encrypt status error, status: %s", status.(string))
				}
			}
			return d, status.(string), err
		},
	}
}

func (s *VolcengineKmsReEncryptService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ReEncrypt",
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"new_keyring_name": {
					TargetField: "NewKeyringName",
				},
				"new_key_name": {
					TargetField: "NewKeyName",
				},
				"new_key_id": {
					TargetField: "NewKeyID",
				},
				"source_ciphertext_blob": {
					TargetField: "CiphertextBlob",
				},
				"old_encryption_context": {
					TargetField: "OldEncryptionContext",
				},
				"new_encryption_context": {
					TargetField: "NewEncryptionContext",
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				ciphertext, _ := ve.ObtainSdkValue("Result.CiphertextBlob", *resp)
				if ciphertext != nil {
					_ = d.Set("ciphertext_blob", ciphertext)
				}
				var keyId string
				if v, ok := d.GetOk("new_key_id"); ok {
					keyId = v.(string)
				} else {
					keyId = fmt.Sprintf("%s:%s", d.Get("new_key_name").(string), d.Get("new_keyring_name").(string))
				}
				d.SetId(keyId)
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineKmsReEncryptService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineKmsReEncryptService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineKmsReEncryptService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineKmsReEncryptService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"new_key_id": {
				TargetField: "NewKeyID",
			},
			"source_ciphertext_blob": {
				TargetField: "CiphertextBlob",
			},
		},
		CollectField:     "ciphertext_info",
		ResponseConverts: map[string]ve.ResponseConvert{},
	}
}

func (s *VolcengineKmsReEncryptService) ReadResourceId(id string) string {
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
