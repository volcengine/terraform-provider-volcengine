package kms_keyring

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mitchellh/copystructure"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineKmsKeyringService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

type filter struct {
	Key    string   `json:"Key"`
	Values []string `json:"Values"`
}

func NewKmsKeyringService(c *ve.SdkClient) *VolcengineKmsKeyringService {
	return &VolcengineKmsKeyringService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineKmsKeyringService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineKmsKeyringService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		filters      []interface{}
		newCondition map[string]interface{}
		resp         *map[string]interface{}
		results      interface{}
		ok           bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "CurrentPage", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeKeyrings"

		deepCopyValue, err := copystructure.Copy(condition)
		if err != nil {
			return data, fmt.Errorf(" DeepCopy condition error: %v ", err)
		}
		if newCondition, ok = deepCopyValue.(map[string]interface{}); !ok {
			return data, fmt.Errorf(" DeepCopy condition error: newCondition is not map ")
		}

		if keyringName, exists := condition["KeyringName"]; exists {
			keyringNameSlice := make([]string, 0)
			keyringNameInter, ok := keyringName.([]interface{})
			if !ok {
				return data, fmt.Errorf(" KeyringName is not slice ")
			}
			for _, v := range keyringNameInter {
				if v == nil {
					keyringNameSlice = append(keyringNameSlice, "")
				} else {
					keyringNameSlice = append(keyringNameSlice, v.(string))
				}
			}
			keyringNameFilter := filter{
				Key:    "KeyringName",
				Values: keyringNameSlice,
			}
			filters = append(filters, keyringNameFilter)
			delete(newCondition, "KeyringName")
		}

		if keyringType, exists := condition["KeyringType"]; exists {
			keyringTypeSlice := make([]string, 0)
			keyringTypeInter, ok := keyringType.([]interface{})
			if !ok {
				return data, fmt.Errorf(" KeyringType is not slice ")
			}
			for _, v := range keyringTypeInter {
				if v == nil {
					keyringTypeSlice = append(keyringTypeSlice, "")
				} else {
					keyringTypeSlice = append(keyringTypeSlice, v.(string))
				}
			}
			keyringTypeFilter := filter{
				Key:    "KeyringType",
				Values: keyringTypeSlice,
			}
			filters = append(filters, keyringTypeFilter)
			delete(newCondition, "KeyringType")
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
		if newCondition == nil {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), nil)
			if err != nil {
				return data, err
			}
		} else {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &newCondition)
			if err != nil {
				return data, err
			}
		}
		respBytes, _ := json.Marshal(resp)
		logger.Debug(logger.RespFormat, action, newCondition, string(respBytes))
		results, err = ve.ObtainSdkValue("Result.Keyrings", *resp)
		if err != nil {
			return data, err
		}
		logger.Debug(logger.RespFormat, "describe kms keyring results", results)
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.Keyrings is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineKmsKeyringService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		ok   bool
		resp *map[string]interface{}
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"KeyringID": id,
	}
	action := "QueryKeyring"
	resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &req)
	if err != nil {
		return data, err
	}
	logger.Debug(logger.RespFormat, action, req, *resp)

	keyring, err := ve.ObtainSdkValue("Result.Keyring", *resp)
	if err != nil {
		return data, err
	}
	data, ok = keyring.(map[string]interface{})
	if !ok {
		return data, fmt.Errorf(" Result is not Map ")
	}

	if len(data) == 0 {
		return data, fmt.Errorf("kms keyrings %s not exist", id)
	}

	return data, err
}

func (s *VolcengineKmsKeyringService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineKmsKeyringService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateKeyring",
			ConvertMode: ve.RequestConvertAll,
			Convert:     map[string]ve.RequestConvert{},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("Result.Keyring.ID", *resp)
				d.SetId(id.(string))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineKmsKeyringService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"UID": {
				TargetField: "uid",
			},
			"TRN": {
				TargetField: "trn",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineKmsKeyringService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateKeyring",
			ConvertMode: ve.RequestConvertAll,
			Convert:     map[string]ve.RequestConvert{},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["KeyringID"] = d.Id()
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

func (s *VolcengineKmsKeyringService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteKeyring",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"KeyringID": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
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
							return resource.NonRetryableError(fmt.Errorf("error on reading keyring on delete %q, %w", d.Id(), callErr))
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

func (s *VolcengineKmsKeyringService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{},
		NameField:       "Name",
		IdField:         "ID",
		CollectField:    "keyrings",
		ResponseConverts: map[string]ve.ResponseConvert{
			"ID": {
				TargetField: "id",
			},
			"UID": {
				TargetField: "uid",
			},
			"TRN": {
				TargetField: "trn",
			},
		},
	}
}

func (s *VolcengineKmsKeyringService) ReadResourceId(id string) string {
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

func (s *VolcengineKmsKeyringService) ProjectTrn() *ve.ProjectTrn {
	return &ve.ProjectTrn{
		ServiceName:          "kms",
		ResourceType:         "keyrings",
		ProjectResponseField: "ProjectName",
		ProjectSchemaField:   "project_name",
	}
}
