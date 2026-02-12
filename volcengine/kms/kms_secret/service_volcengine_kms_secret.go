package kms_secret

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/mitchellh/copystructure"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineKmsSecretService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

type filter struct {
	Key    string   `json:"Key"`
	Values []string `json:"Values"`
}

func NewKmsSecretService(c *ve.SdkClient) *VolcengineKmsSecretService {
	return &VolcengineKmsSecretService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineKmsSecretService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineKmsSecretService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		filters      []interface{}
		newCondition map[string]interface{}
		resp         *map[string]interface{}
		results      interface{}
		ok           bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "CurrentPage", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeSecrets"

		deepCopyValue, err := copystructure.Copy(condition)
		if err != nil {
			return data, fmt.Errorf(" DeepCopy condition error: %v ", err)
		}
		if newCondition, ok = deepCopyValue.(map[string]interface{}); !ok {
			return data, fmt.Errorf(" DeepCopy condition error: newCondition is not map ")
		}

		if secretName, exists := condition["SecretName"]; exists {
			secretNameSlice := make([]string, 0)
			secretNameInter, ok := secretName.([]interface{})
			if !ok {
				return data, fmt.Errorf(" SecretName is not slice ")
			}
			for _, v := range secretNameInter {
				if v == nil {
					secretNameSlice = append(secretNameSlice, "")
				} else {
					secretNameSlice = append(secretNameSlice, v.(string))
				}
			}
			secretNameFilter := filter{
				Key:    "SecretName",
				Values: secretNameSlice,
			}
			filters = append(filters, secretNameFilter)
			delete(newCondition, "SecretName")
		}

		if trn, exists := condition["Trn"]; exists {
			trnSlice := make([]string, 0)
			trnInter, ok := trn.([]interface{})
			if !ok {
				return data, fmt.Errorf(" trn is not slice ")
			}
			for _, v := range trnInter {
				if v == nil {
					trnSlice = append(trnSlice, "")
				} else {
					trnSlice = append(trnSlice, v.(string))
				}
			}
			trnFilter := filter{
				Key:    "Trn",
				Values: trnSlice,
			}
			filters = append(filters, trnFilter)
			delete(newCondition, "Trn")
		}

		if secretType, exists := condition["SecretType"]; exists {
			secretTypeSlice := make([]string, 0)
			secretTypeInter, ok := secretType.([]interface{})
			if !ok {
				return data, fmt.Errorf(" SecretType is not slice ")
			}
			for _, v := range secretTypeInter {
				if v == nil {
					secretTypeSlice = append(secretTypeSlice, "")
				} else {
					secretTypeSlice = append(secretTypeSlice, v.(string))
				}
			}
			secretTypeFilter := filter{
				Key:    "SecretType",
				Values: secretTypeSlice,
			}
			filters = append(filters, secretTypeFilter)
			delete(newCondition, "SecretType")
		}

		if secretState, exists := condition["SecretState"]; exists {
			secretStateSlice := make([]string, 0)
			secretStateInter, ok := secretState.([]interface{})
			if !ok {
				return data, fmt.Errorf(" SecretState is not slice ")
			}
			for _, v := range secretStateInter {
				if v == nil {
					secretStateSlice = append(secretStateSlice, "")
				} else {
					secretStateSlice = append(secretStateSlice, v.(string))
				}
			}
			secretStateFilter := filter{
				Key:    "SecretState",
				Values: secretStateSlice,
			}
			filters = append(filters, secretStateFilter)
			delete(newCondition, "SecretState")
		}

		if managedState, exists := condition["ManagedState"]; exists {
			managedStateSlice := make([]string, 0)
			managedStateInter, ok := managedState.([]interface{})
			if !ok {
				return data, fmt.Errorf(" ManagedState is not slice ")
			}
			for _, v := range managedStateInter {
				if v == nil {
					managedStateSlice = append(managedStateSlice, "")
				} else {
					managedStateSlice = append(managedStateSlice, v.(string))
				}
			}
			managedStateFilter := filter{
				Key:    "ManagedState",
				Values: managedStateSlice,
			}
			filters = append(filters, managedStateFilter)
			delete(newCondition, "ManagedState")
		}

		if owningService, exists := condition["OwningService"]; exists {
			owningServiceSlice := make([]string, 0)
			owningServiceInter, ok := owningService.([]interface{})
			if !ok {
				return data, fmt.Errorf(" OwningService is not slice ")
			}
			for _, v := range owningServiceInter {
				if v == nil {
					owningServiceSlice = append(owningServiceSlice, "")
				} else {
					owningServiceSlice = append(owningServiceSlice, v.(string))
				}
			}
			owningServiceFilter := filter{
				Key:    "OwningService",
				Values: owningServiceSlice,
			}
			filters = append(filters, owningServiceFilter)
			delete(newCondition, "OwningService")
		}

		if rotationState, exists := condition["RotationState"]; exists {
			rotationStateSlice := make([]string, 0)
			rotationStateInter, ok := rotationState.([]interface{})
			if !ok {
				return data, fmt.Errorf(" RotationState is not slice ")
			}
			for _, v := range rotationStateInter {
				if v == nil {
					rotationStateSlice = append(rotationStateSlice, "")
				} else {
					rotationStateSlice = append(rotationStateSlice, v.(string))
				}
			}
			rotationStateFilter := filter{
				Key:    "RotationState",
				Values: rotationStateSlice,
			}
			filters = append(filters, rotationStateFilter)
			delete(newCondition, "RotationState")
		}

		if description, exists := condition["Description"]; exists {
			descriptionSlice := make([]string, 0)
			descriptionInter, ok := description.([]interface{})
			if !ok {
				return data, fmt.Errorf(" description is not slice ")
			}
			for _, v := range descriptionInter {
				if v == nil {
					descriptionSlice = append(descriptionSlice, "")
				} else {
					descriptionSlice = append(descriptionSlice, v.(string))
				}
			}
			descriptionFilter := filter{
				Key:    "Description",
				Values: descriptionSlice,
			}
			filters = append(filters, descriptionFilter)
			delete(newCondition, "Description")
		}

		if creationDateRange, exists := condition["CreationDateRange"]; exists {
			creationDateRangeSlice := make([]string, 0)
			creationDateRangeInter, ok := creationDateRange.([]interface{})
			if !ok {
				return data, fmt.Errorf(" CreationDateRange is not slice ")
			}
			for _, v := range creationDateRangeInter {
				if v == nil {
					creationDateRangeSlice = append(creationDateRangeSlice, "")
				} else {
					creationDateRangeSlice = append(creationDateRangeSlice, v.(string))
				}
			}
			creationDateRangeFilter := filter{
				Key:    "CreationDateRange",
				Values: creationDateRangeSlice,
			}
			filters = append(filters, creationDateRangeFilter)
			delete(newCondition, "CreationDateRange")
		}

		if updateDateRange, exists := condition["UpdateDateRange"]; exists {
			updateDateRangeSlice := make([]string, 0)
			updateDateRangeInter, ok := updateDateRange.([]interface{})
			if !ok {
				return data, fmt.Errorf(" UpdateDateRange is not slice ")
			}
			for _, v := range updateDateRangeInter {
				if v == nil {
					updateDateRangeSlice = append(updateDateRangeSlice, "")
				} else {
					updateDateRangeSlice = append(updateDateRangeSlice, v.(string))
				}
			}
			updateDateRangeFilter := filter{
				Key:    "UpdateDateRange",
				Values: updateDateRangeSlice,
			}
			filters = append(filters, updateDateRangeFilter)
			delete(newCondition, "UpdateDateRange")
		}

		if len(filters) > 0 {
			filtersBytes, _ := json.Marshal(filters)
			newCondition["Filters"] = string(filtersBytes)
		}

		logger.Debug(logger.ReqFormat, action, &newCondition)

		bytes, _ := json.Marshal(newCondition)
		logger.Debug(logger.ReqFormat, action, string(bytes))
		if condition == nil {
			resp, err = s.Client.UniversalClient.DoCall(getPostUniversalInfo(action), nil)
			if err != nil {
				return data, err
			}
		} else {
			resp, err = s.Client.UniversalClient.DoCall(getPostUniversalInfo(action), &newCondition)
			if err != nil {
				return data, err
			}
		}
		respBytes, _ := json.Marshal(resp)
		logger.Debug(logger.RespFormat, action, newCondition, string(respBytes))
		results, err = ve.ObtainSdkValue("Result.Secrets", *resp)
		if err != nil {
			return data, err
		}
		logger.Debug(logger.RespFormat, "describe kms secrets results", results)
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.Secrets is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineKmsSecretService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		ok                 bool
		resp               *map[string]interface{}
		getSecretValueResp *map[string]interface{}
		secretValueData    map[string]interface{}
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

	keyring, err := ve.ObtainSdkValue("Result.Secret", *resp)
	if err != nil {
		return data, err
	}
	data, ok = keyring.(map[string]interface{})
	if !ok {
		return data, fmt.Errorf(" Result is not Map ")
	}

	if len(data) == 0 {
		return data, fmt.Errorf("kms Secret name %s not exist", resourceData.Get("secret_name"))
	}
	// 转换一下RotationInterval返回
	rotationInterval := data["RotationInterval"].(float64)
	rotationIntervalDay := fmt.Sprintf("%dd", int(rotationInterval)/86400)
	data["RotationInterval"] = rotationIntervalDay

	data["Uuid"] = data["ID"]
	data["State"] = data["SecretState"]

	getSecretValueReq := map[string]interface{}{
		"SecretName": id,
	}
	getSecretValueAction := "GetSecretValue"
	getSecretValueResp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(getSecretValueAction), &getSecretValueReq)
	if err != nil {
		return data, err
	}
	logger.Debug(logger.RespFormat, getSecretValueAction, getSecretValueReq, *getSecretValueResp)

	secretValue, err := ve.ObtainSdkValue("Result", *getSecretValueResp)
	if err != nil {
		return data, err
	}

	secretValueData, ok = secretValue.(map[string]interface{})
	if !ok {
		return data, fmt.Errorf(" Result is not Map ")
	}

	if len(secretValueData) == 0 {
		return data, fmt.Errorf("secretValue %s not exist", resourceData.Get("secret_name"))
	}
	data["SecretValue"] = secretValueData["SecretValue"]

	return data, err
}

func (s *VolcengineKmsSecretService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineKmsSecretService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateSecret",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert:     map[string]ve.RequestConvert{},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getPostUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("Result.Secret.SecretName", *resp)
				d.SetId(id.(string))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineKmsSecretService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"UID": {
				TargetField: "uid",
			},
			"Trn": {
				TargetField: "trn",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineKmsSecretService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callbacks := make([]ve.Callback, 0)

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action: "UpdateSecret",
			// UpdateSecret 只支持修改 Description
			ConvertMode: ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"description": {
					TargetField: "Description",
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["SecretName"] = d.Id()
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

	callbacks = append(callbacks, callback)

	if resourceData.HasChanges("automatic_rotation", "rotation_interval") {
		secretRotationCallBack := ve.Callback{
			Call: ve.SdkCall{
				Action:      "UpdateSecretRotationPolicy",
				ConvertMode: ve.RequestConvertIgnore,
				Convert:     map[string]ve.RequestConvert{},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["SecretName"] = resourceData.Id()
					(*call.SdkParam)["AutomaticRotation"] = resourceData.Get("automatic_rotation")
					(*call.SdkParam)["RotationInterval"] = resourceData.Get("rotation_interval")
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
			},
		}
		callbacks = append(callbacks, secretRotationCallBack)
	}

	if resourceData.HasChanges("secret_value", "version_name") {
		// Generic 类型的 Secret 才支持存入新的 secret_value 和 version_name
		secretValueCallBack := ve.Callback{
			Call: ve.SdkCall{
				Action:      "SetSecretValue",
				ConvertMode: ve.RequestConvertIgnore,
				Convert:     map[string]ve.RequestConvert{},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					if resourceData.Get("secret_type").(string) != "Generic" {
						return false, fmt.Errorf("Only Generic type secret support modifying secret_value and version_name.")
					}
					(*call.SdkParam)["SecretName"] = resourceData.Id()
					(*call.SdkParam)["SecretValue"] = resourceData.Get("secret_value")
					(*call.SdkParam)["VersionName"] = resourceData.Get("version_name")
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
			},
		}
		callbacks = append(callbacks, secretValueCallBack)
	}

	return callbacks
}

func (s *VolcengineKmsSecretService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ScheduleSecretDeletion",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam:    &map[string]interface{}{},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				forceDelete := d.Get("force_delete").(bool)
				if !forceDelete {
					if pendingWindowInDays := d.Get("pending_window_in_days").(int); pendingWindowInDays > 0 {
						(*call.SdkParam)["PendingWindowInDays"] = pendingWindowInDays
					} else {
						(*call.SdkParam)["PendingWindowInDays"] = 7
					}
				}
				(*call.SdkParam)["SecretName"] = d.Id()
				(*call.SdkParam)["ForceDelete"] = forceDelete
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				forceDelete := d.Get("force_delete").(bool)
				if !forceDelete {
					return nil
				}
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 5*time.Minute)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(5*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					logger.Debug(logger.RespFormat, call.Action, callErr)
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading secret on delete %q, %w", d.Id(), callErr))
						}
					}
					_, callErr = call.ExecuteCall(d, client, call)
					if callErr == nil {
						return nil
					}
					return resource.RetryableError(callErr)
				})
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineKmsSecretService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{},
		NameField:       "SecretName",
		IdField:         "ID",
		CollectField:    "secrets",
		ResponseConverts: map[string]ve.ResponseConvert{
			"ID": {
				TargetField: "id",
			},
			"UID": {
				TargetField: "uid",
			},
			"Trn": {
				TargetField: "trn",
			},
		},
	}
}

func (s *VolcengineKmsSecretService) ReadResourceId(id string) string {
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

func getPostUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "kms",
		Version:     "2021-02-18",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}

func (s *VolcengineKmsSecretService) ProjectTrn() *ve.ProjectTrn {
	return &ve.ProjectTrn{
		ServiceName:          "kms",
		ResourceType:         "secrets",
		ProjectResponseField: "ProjectName",
		ProjectSchemaField:   "project_name",
	}
}
