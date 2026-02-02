package kms_key

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

type VolcengineKmsKeyService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

type filter struct {
	Key    string   `json:"Key"`
	Values []string `json:"Values"`
}

func NewKmsKeyService(c *ve.SdkClient) *VolcengineKmsKeyService {
	return &VolcengineKmsKeyService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineKmsKeyService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineKmsKeyService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		filters      []interface{}
		newCondition map[string]interface{}
		resp         *map[string]interface{}
		results      interface{}
		ok           bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "CurrentPage", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeKeys"

		deepCopyValue, err := copystructure.Copy(condition)
		if err != nil {
			return data, fmt.Errorf(" DeepCopy condition error: %v ", err)
		}
		if newCondition, ok = deepCopyValue.(map[string]interface{}); !ok {
			return data, fmt.Errorf(" DeepCopy condition error: newCondition is not map ")
		}

		if keyName, exists := condition["KeyName"]; exists {
			keyNameSlice := make([]string, 0)
			keyNameInter, ok := keyName.([]interface{})
			if !ok {
				return data, fmt.Errorf(" key name is not slice ")
			}
			for _, v := range keyNameInter {
				if v == nil {
					keyNameSlice = append(keyNameSlice, "")
				} else {
					keyNameSlice = append(keyNameSlice, v.(string))
				}
			}
			keyNameFilter := filter{
				Key:    "KeyName",
				Values: keyNameSlice,
			}
			filters = append(filters, keyNameFilter)
			delete(newCondition, "KeyName")
		}

		if keySpec, exists := condition["KeySpec"]; exists {
			keySpecSlice := make([]string, 0)
			keySpecInter, ok := keySpec.([]interface{})
			if !ok {
				return data, fmt.Errorf(" key spec is not slice ")
			}
			for _, v := range keySpecInter {
				if v == nil {
					keySpecSlice = append(keySpecSlice, "")
				} else {
					keySpecSlice = append(keySpecSlice, v.(string))
				}
			}
			keySpecFilter := filter{
				Key:    "KeySpec",
				Values: keySpecSlice,
			}
			filters = append(filters, keySpecFilter)
			delete(newCondition, "KeySpec")
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

		if keyState, exists := condition["KeyState"]; exists {
			keyStateSlice := make([]string, 0)
			keyStateInter, ok := keyState.([]interface{})
			if !ok {
				return data, fmt.Errorf(" key state is not slice ")
			}
			for _, v := range keyStateInter {
				if v == nil {
					keyStateSlice = append(keyStateSlice, "")
				} else {
					keyStateSlice = append(keyStateSlice, v.(string))
				}
			}
			keyStateFilter := filter{
				Key:    "KeyState",
				Values: keyStateSlice,
			}
			filters = append(filters, keyStateFilter)
			delete(newCondition, "KeyState")
		}

		if keyUsage, exists := condition["KeyUsage"]; exists {
			keyUsageSlice := make([]string, 0)
			keyUsageInter, ok := keyUsage.([]interface{})
			if !ok {
				return data, fmt.Errorf(" key usage is not slice ")
			}
			for _, v := range keyUsageInter {
				if v == nil {
					keyUsageSlice = append(keyUsageSlice, "")
				} else {
					keyUsageSlice = append(keyUsageSlice, v.(string))
				}
			}
			keyUsageFilter := filter{
				Key:    "KeyUsage",
				Values: keyUsageSlice,
			}
			filters = append(filters, keyUsageFilter)
			delete(newCondition, "KeyUsage")
		}

		if protectionLevel, exists := condition["ProtectionLevel"]; exists {
			protectionLevelSlice := make([]string, 0)
			protectionLevelInter, ok := protectionLevel.([]interface{})
			if !ok {
				return data, fmt.Errorf(" protectionLevel is not slice ")
			}
			for _, v := range protectionLevelInter {
				if v == nil {
					protectionLevelSlice = append(protectionLevelSlice, "")
				} else {
					protectionLevelSlice = append(protectionLevelSlice, v.(string))
				}
			}
			protectionLevelFilter := filter{
				Key:    "ProtectionLevel",
				Values: protectionLevelSlice,
			}
			filters = append(filters, protectionLevelFilter)
			delete(newCondition, "ProtectionLevel")
		}

		if rotateState, exists := condition["RotateState"]; exists {
			rotateStateSlice := make([]string, 0)
			rotateStateInter, ok := rotateState.([]interface{})
			if !ok {
				return data, fmt.Errorf(" rotateState is not slice ")
			}
			for _, v := range rotateStateInter {
				if v == nil {
					rotateStateSlice = append(rotateStateSlice, "")
				} else {
					rotateStateSlice = append(rotateStateSlice, v.(string))
				}
			}
			rotateStateFilter := filter{
				Key:    "RotateState",
				Values: rotateStateSlice,
			}
			filters = append(filters, rotateStateFilter)
			delete(newCondition, "RotateState")
		}

		if origin, exists := condition["Origin"]; exists {
			originSlice := make([]string, 0)
			originInter, ok := origin.([]interface{})
			if !ok {
				return data, fmt.Errorf(" origin is not slice ")
			}
			for _, v := range originInter {
				if v == nil {
					originSlice = append(originSlice, "")
				} else {
					originSlice = append(originSlice, v.(string))
				}
			}
			originFilter := filter{
				Key:    "Origin",
				Values: originSlice,
			}
			filters = append(filters, originFilter)
			delete(newCondition, "Origin")
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
		if newCondition == nil {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalPostInfo(action), nil)
			if err != nil {
				return data, err
			}
		} else {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalPostInfo(action), &newCondition)
			if err != nil {
				return data, err
			}
		}
		respBytes, _ := json.Marshal(resp)
		logger.Debug(logger.RespFormat, action, newCondition, string(respBytes))
		results, err = ve.ObtainSdkValue("Result.Keys", *resp)
		if err != nil {
			return data, err
		}
		logger.Debug(logger.RespFormat, "describe kms keys results", results)
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.Keys is not Slice")
		}

		logger.Debug(logger.RespFormat, "result data is", results)

		return data, err
	})
}

func (s *VolcengineKmsKeyService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		ok   bool
		resp *map[string]interface{}
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"KeyID": id,
	}
	action := "DescribeKey"
	resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &req)
	if err != nil {
		return data, err
	}
	logger.Debug(logger.RespFormat, action, req, *resp)

	keyring, err := ve.ObtainSdkValue("Result.Key", *resp)
	if err != nil {
		return data, err
	}
	data, ok = keyring.(map[string]interface{})
	if !ok {
		return data, fmt.Errorf(" Result is not Map ")
	}

	if len(data) == 0 {
		return data, fmt.Errorf("kms key %s not exist", id)
	}

	data["State"] = data["KeyState"]

	return data, err
}

func (s *VolcengineKmsKeyService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineKmsKeyService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateKey",
			ContentType: ve.ContentTypeJson,
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"tags": {
					TargetField: "Tags",
					ConvertType: ve.ConvertJsonObjectArray,
				},
				"custom_key_store_id": {
					TargetField: "CustomKeyStoreID",
				},
				"xks_key_id": {
					TargetField: "XksKeyID",
				},
				"rotate_state": {
					Ignore: true,
				},
				"rotate_interval": {
					Ignore: true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if rotateState, ok := d.GetOk("rotate_state"); ok {
					rotate := rotateState.(string)
					if rotate == "Enable" {
						if rotateInterval, ok := d.GetOk("rotate_interval"); ok {
							rotateInterval := rotateInterval.(int)
							(*call.SdkParam)["RotateInterval"] = rotateInterval
						} else {
							return false, fmt.Errorf("rotate_interval is required when rotate_state is Enable")
						}
					}
				}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalPostInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("Result.Key.ID", *resp)
				d.SetId(id.(string))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineKmsKeyService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	if multiRegionConfiguration, ok := d["MultiRegionConfiguration"]; ok {
		if primaryKey, ok := multiRegionConfiguration.(map[string]interface{})["PrimaryKey"]; ok {
			multiRegionConfiguration.(map[string]interface{})["PrimaryKey"] = []interface{}{primaryKey}
		}
		if replicaKeys, ok := multiRegionConfiguration.(map[string]interface{})["ReplicaKeys"]; ok {
			multiRegionConfiguration.(map[string]interface{})["ReplicaKeys"] = []interface{}{replicaKeys}
		}
	} else {
		d["MultiRegionConfiguration"] = interface{}(map[string]interface{}{})
	}

	if rotationState, ok := d["RotationState"].(string); ok {
		if rotationState == "Disable" {
			delete(d, "RotateInterval")
		}
	}

	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineKmsKeyService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback
	// 支持修改 key_name tags descryption
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateKey",
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"tags": {
					Ignore: true,
				},
				"description": {
					TargetField: "Description",
				},
				"rotate_state": {
					Ignore: true,
				},
				"rotate_interval": {
					Ignore: true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["KeyID"] = d.Id()
				(*call.SdkParam)["NewKeyName"] = d.Get("key_name")
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Id()
			},
		},
	}
	callbacks = append(callbacks, callback)

	// 处理 轮转 状态的更新
	// Disable --> Enable: 调用 EnableKeyRotation
	// Enable --> Disable: 调用 DisableKeyRotation
	// Enable --> Enable + 轮转间隔更新: 调用 EnableKeyRotation 更新间隔
	if resourceData.HasChanges("rotate_state", "rotate_interval") {
		oldState, newState := resourceData.GetChange("rotate_state")
		fmt.Println(oldState, newState)
		// Disable --> Enable, 或者 Enable-->Enable + 间隔更新, 或者 空-->Enable: 调用 EnableKeyRotation
		if newState.(string) == "Enable" {
			callback := ve.Callback{
				Call: ve.SdkCall{
					Action:      "EnableKeyRotation",
					ConvertMode: ve.RequestConvertIgnore,
					Convert:     map[string]ve.RequestConvert{},
					BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
						if rotateInterval, ok := d.GetOk("rotate_interval"); ok {
							rotateInterval := rotateInterval.(int)
							(*call.SdkParam)["RotateInterval"] = rotateInterval
						} else {
							return false, fmt.Errorf("rotate_interval is required when rotate_state is Enable")
						}
						// 兼容 多地域 密钥，使用 keyring_name + key_name
						// (*call.SdkParam)["KeyID"] = d.Id()
						(*call.SdkParam)["KeyringName"] = d.Get("keyring_name")
						(*call.SdkParam)["KeyName"] = d.Get("key_name")
						return true, nil
					},
					ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
						logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
						resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
						logger.Debug(logger.RespFormat, call.Action, resp, err)
						return resp, err
					},
					LockId: func(d *schema.ResourceData) string {
						return d.Id()
					},
				},
			}
			callbacks = append(callbacks, callback)
		} else if (oldState.(string) == "Enable" && newState.(string) == "Disable") || newState.(string) == "" {
			// Enable --> Disable 或者 Enable--> 空: 调用 DisableKeyRotation
			callback := ve.Callback{
				Call: ve.SdkCall{
					Action:      "DisableKeyRotation",
					ConvertMode: ve.RequestConvertInConvert,
					Convert: map[string]ve.RequestConvert{
						"keyring_name": {
							TargetField: "KeyringName",
							ForceGet:    true,
						},
						"key_name": {
							TargetField: "KeyName",
							ForceGet:    true,
						},
					},
					ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
						logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
						resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
						logger.Debug(logger.RespFormat, call.Action, resp, err)
						return resp, err
					},
					LockId: func(d *schema.ResourceData) string {
						return d.Id()
					},
				},
			}
			callbacks = append(callbacks, callback)
		}
	}

	// 更新Tags
	setResourceTagsCallbacks := s.setResourceTags(resourceData, "keys")
	callbacks = append(callbacks, setResourceTagsCallbacks...)

	return callbacks
}

func (s *VolcengineKmsKeyService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ScheduleKeyDeletion",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if resourceData.Get("pending_window_in_days") != 0 {
					(*call.SdkParam)["KeyID"] = d.Id()
					(*call.SdkParam)["PendingWindowInDays"] = resourceData.Get("pending_window_in_days")
				} else {
					(*call.SdkParam)["KeyID"] = d.Id()
				}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return s.checkResourceUtilRemoved(d, 5*time.Minute)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(5*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					logger.Debug(logger.RespFormat, call.Action, callErr)
					if callErr != nil {
						return resource.NonRetryableError(fmt.Errorf("error on reading key on PendingDelete %q, %w", d.Id(), callErr))
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

func (s *VolcengineKmsKeyService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"tags": {
				TargetField: "TagFilters",
				ConvertType: ve.ConvertJsonObjectArray,
				NextLevelConvert: map[string]ve.RequestConvert{
					"values": {
						TargetField: "Values",
						ConvertType: ve.ConvertJsonArray,
					},
				},
			},
			"keyring_id": {
				TargetField: "KeyringID",
			},
		},
		NameField:    "KeyName",
		IdField:      "ID",
		CollectField: "keys",
		ContentType:  ve.ContentTypeJson,
		ResponseConverts: map[string]ve.ResponseConvert{
			"ID": {
				TargetField: "id",
			},
			"CustomKeyStoreID": {
				TargetField: "custom_key_store_id",
			},
			"XksKeyConfiguration.ID": {
				TargetField: "id",
			},
		},
	}
}

func (s *VolcengineKmsKeyService) ReadResourceId(id string) string {
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

func getUniversalPostInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "kms",
		Version:     "2021-02-18",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}

func (s *VolcengineKmsKeyService) setResourceTags(resourceData *schema.ResourceData, resourceType string) []ve.Callback {
	var callbacks []ve.Callback
	addedTags, removedTags, _, _ := ve.GetSetDifference("tags", resourceData, ve.TagsHash, false)

	removeCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UntagResources",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if removedTags != nil && len(removedTags.List()) > 0 {
					(*call.SdkParam)["ResourceIds"] = []string{resourceData.Id()}
					(*call.SdkParam)["ResourceType"] = resourceType
					(*call.SdkParam)["TagKeys"] = make([]string, 0)
					for _, tag := range removedTags.List() {
						(*call.SdkParam)["TagKeys"] = append((*call.SdkParam)["TagKeys"].([]string), tag.(map[string]interface{})["key"].(string))
					}
					return true, nil
				}
				return false, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalPostInfo(call.Action), call.SdkParam)
			},
		},
	}
	callbacks = append(callbacks, removeCallback)

	addCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "TagResources",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if addedTags != nil && len(addedTags.List()) > 0 {
					(*call.SdkParam)["ResourceIds"] = []string{resourceData.Id()}
					(*call.SdkParam)["ResourceType"] = resourceType
					(*call.SdkParam)["Tags"] = make([]map[string]interface{}, 0)
					for _, tag := range addedTags.List() {
						(*call.SdkParam)["Tags"] = append((*call.SdkParam)["Tags"].([]map[string]interface{}), tag.(map[string]interface{}))
					}
					return true, nil
				}
				return false, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalPostInfo(call.Action), call.SdkParam)
			},
		},
	}
	callbacks = append(callbacks, addCallback)

	return callbacks
}

func (s *VolcengineKmsKeyService) checkResourceUtilRemoved(d *schema.ResourceData, timeout time.Duration) error {
	return resource.Retry(timeout, func() *resource.RetryError {
		keyStatus, _ := s.ReadResource(d, d.Id())
		// 能查询成功代表还在删除中，重试
		if keyStatus["KeyState"] != "PendingDelete" {
			return resource.RetryableError(fmt.Errorf("resource still in removing status "))
		} else {
			if keyStatus["KeyState"] == "PendingDelete" {
				return nil
			} else {
				return resource.NonRetryableError(fmt.Errorf("kms key status is not PendingDelete "))
			}
		}
	})
}
