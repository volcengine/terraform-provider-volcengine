package kms_key_material

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineKmsKeyMaterialService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewKmsKeyMaterialService(c *ve.SdkClient) *VolcengineKmsKeyMaterialService {
	return &VolcengineKmsKeyMaterialService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineKmsKeyMaterialService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineKmsKeyMaterialService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp *map[string]interface{}
	)
	return ve.WithSimpleQuery(m, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "GetParametersForImport"

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
		keyringId, _ := ve.ObtainSdkValue("Result.KeyringID", *resp)
		if keyringId != nil {
			result["keyring_id"] = keyringId.(string)
		}

		keyId, _ := ve.ObtainSdkValue("Result.KeyID", *resp)
		if keyId != nil {
			result["key_id"] = keyId.(string)
		}

		publicKey, _ := ve.ObtainSdkValue("Result.PublicKey", *resp)
		if publicKey != nil {
			result["public_key"] = publicKey.(string)
		}

		importToken, _ := ve.ObtainSdkValue("Result.ImportToken", *resp)
		if importToken != nil {
			result["import_token"] = importToken.(string)
		}

		tokenExpireTime, _ := ve.ObtainSdkValue("Result.TokenExpireTime", *resp)
		if tokenExpireTime != nil {
			result["token_expire_time"] = tokenExpireTime.(string)
		}

		return []interface{}{result}, nil
	})
}

func (s *VolcengineKmsKeyMaterialService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	return data, nil
}

func (s *VolcengineKmsKeyMaterialService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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
					return nil, "", fmt.Errorf("kms_key_material status error, status: %s", status.(string))
				}
			}
			return d, status.(string), err
		},
	}
}

func (s *VolcengineKmsKeyMaterialService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ImportKeyMaterial",
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
				"encrypted_key_material": {
					TargetField: "EncryptedKeyMaterial",
				},
				"import_token": {
					TargetField: "ImportToken",
				},
				"expiration_model": {
					TargetField: "ExpirationModel",
				},
				"valid_to": {
					TargetField: "ValidTo",
				},
			},
			// 交给 API 层进行校验
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfoPost(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				var resourceId string
				if v, ok := d.GetOk("key_id"); ok {
					resourceId = v.(string)
				} else {
					resourceId = fmt.Sprintf("%s:%s", d.Get("key_name").(string), d.Get("keyring_name").(string))
				}
				d.SetId(resourceId)
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineKmsKeyMaterialService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"KeyID": {
				TargetField: "key_id",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineKmsKeyMaterialService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{}
	return []ve.Callback{callback}
}

func (s *VolcengineKmsKeyMaterialService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteKeyMaterial",
			ConvertMode: ve.RequestConvertIgnore,
			Convert:     map[string]ve.RequestConvert{},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if keyId, ok := d.GetOk("key_id"); ok {
					(*call.SdkParam)["KeyID"] = keyId.(string)
				}
				if keyName, ok := d.GetOk("key_name"); ok {
					(*call.SdkParam)["KeyName"] = keyName.(string)
				}
				if keyringName, ok := d.GetOk("keyring_name"); ok {
					(*call.SdkParam)["KeyringName"] = keyringName.(string)
				}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineKmsKeyMaterialService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"keyring_name": {
				TargetField: "KeyringName",
			},
			"key_name": {
				TargetField: "KeyName",
			},
			"key_id": {
				TargetField: "KeyID",
			},
			"wrapping_key_spec": {
				TargetField: "WrappingKeySpec",
			},
			"wrapping_algorithm": {
				TargetField: "WrappingAlgorithm",
			},
		},
		NameField:    "keyring_name",
		IdField:      "KeyID",
		CollectField: "import_parameters",
		ResponseConverts: map[string]ve.ResponseConvert{
			"KeyringID": {
				TargetField: "keyring_id",
			},
			"KeyID": {
				TargetField: "key_id",
			},
			"PublicKey": {
				TargetField: "public_key",
			},
			"ImportToken": {
				TargetField: "import_token",
			},
			"TokenExpireTime": {
				TargetField: "token_expire_time",
			},
		},
	}
}

func (s *VolcengineKmsKeyMaterialService) ReadResourceId(id string) string {
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

func getUniversalInfoPost(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "kms",
		Version:     "2021-02-18",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}
