package vepfs_fileset

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mitchellh/copystructure"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineVepfsFilesetService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewVepfsFilesetService(c *ve.SdkClient) *VolcengineVepfsFilesetService {
	return &VolcengineVepfsFilesetService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineVepfsFilesetService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineVepfsFilesetService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		newCondition map[string]interface{}
		resp         *map[string]interface{}
		results      interface{}
		ok           bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeFilesets"

		deepCopyValue, err := copystructure.Copy(condition)
		if err != nil {
			return data, fmt.Errorf(" DeepCopy condition error: %v ", err)
		}
		if newCondition, ok = deepCopyValue.(map[string]interface{}); !ok {
			return data, fmt.Errorf(" DeepCopy condition error: newCondition is not map ")
		}

		// 处理 Filters
		if filters, exists := condition["Filters"]; exists {
			filtersMap, ok := filters.(map[string]interface{})
			if !ok {
				return data, fmt.Errorf(" Filters is not map ")
			}
			// 处理 Filters.Status，逗号分离
			if statuses, exists := filtersMap["Status"]; exists {
				statusArr, ok := statuses.([]interface{})
				if !ok {
					return data, fmt.Errorf(" Filters.Status is not slice ")
				}
				newStatus := make([]string, 0)
				for _, status := range statusArr {
					newStatus = append(newStatus, status.(string))
				}
				newCondition["Filters"].(map[string]interface{})["Status"] = strings.Join(newStatus, ",")
			}

			newFilters := make([]interface{}, 0)
			for key, value := range newCondition["Filters"].(map[string]interface{}) {
				newFilters = append(newFilters, map[string]interface{}{
					"Key":   key,
					"Value": value,
				})
			}
			newCondition["Filters"] = newFilters
		}

		bytes, _ := json.Marshal(newCondition)
		logger.Debug(logger.ReqFormat, action, string(bytes))
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
		results, err = ve.ObtainSdkValue("Result.Filesets", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.Filesets is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineVepfsFilesetService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	ids := strings.Split(id, ":")
	if len(ids) != 2 {
		return data, fmt.Errorf("invalid vepfs fileset Id: %s", id)
	}

	req := map[string]interface{}{
		"FileSystemId": ids[0],
		"Filters": map[string]interface{}{
			"FilesetId": ids[1],
		},
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		if data, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("Value is not map ")
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("vepfs_fileset %s not exist ", id)
	}

	// 特殊处理 FilesetPath
	if filePath, exist := data["FilesetPath"]; exist {
		newPath := filePath.(string) + "/"
		data["FilesetPath"] = newPath
	}

	return data, err
}

func (s *VolcengineVepfsFilesetService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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
			failStates = append(failStates, "Error")
			failStates = append(failStates, "CreateError")
			failStates = append(failStates, "UpdateError")
			failStates = append(failStates, "DeleteError")
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
					return nil, "", fmt.Errorf("vepfs_fileset status error, status: %s", status.(string))
				}
			}
			return d, status.(string), err
		},
	}
}

func (VolcengineVepfsFilesetService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"IOPSQos": {
				TargetField: "max_iops",
			},
			"BandwidthQos": {
				TargetField: "max_bandwidth",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineVepfsFilesetService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateFileset",
			ConvertMode: ve.RequestConvertInConvert,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"file_system_id": {
					TargetField: "FileSystemId",
				},
				"fileset_name": {
					TargetField: "FilesetName",
				},
				"fileset_path": {
					TargetField: "FilesetPath",
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				filesetId, _ := ve.ObtainSdkValue("Result.FilesetId", *resp)
				fileSystemId := d.Get("file_system_id")
				d.SetId(fileSystemId.(string) + ":" + filesetId.(string))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Running"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	callbacks = append(callbacks, callback)

	// 设置 Qos
	_, ok1 := resourceData.GetOk("max_iops")
	_, ok2 := resourceData.GetOk("max_bandwidth")
	if ok1 || ok2 {
		qosCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "SetFilesetQos",
				ConvertMode: ve.RequestConvertInConvert,
				ContentType: ve.ContentTypeJson,
				Convert: map[string]ve.RequestConvert{
					"max_iops": {
						TargetField: "MaxIops",
					},
					"max_bandwidth": {
						TargetField: "MaxBandwidth",
					},
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					ids := strings.Split(d.Id(), ":")
					if len(ids) != 2 {
						return false, fmt.Errorf("invalid vepfs fileset Id: %s", d.Id())
					}

					(*call.SdkParam)["FileSystemId"] = ids[0]
					(*call.SdkParam)["FilesetId"] = ids[1]
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
					resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
					logger.Debug(logger.RespFormat, call.Action, resp, err)
					return resp, err
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{"Running"},
					Timeout: resourceData.Timeout(schema.TimeoutCreate),
				},
			},
		}
		callbacks = append(callbacks, qosCallback)
	}

	// 设置 Quota
	_, ok3 := resourceData.GetOk("file_limit")
	_, ok4 := resourceData.GetOk("capacity_limit")
	if ok3 || ok4 {
		quotaCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "SetFilesetQuota",
				ConvertMode: ve.RequestConvertInConvert,
				ContentType: ve.ContentTypeJson,
				Convert: map[string]ve.RequestConvert{
					"file_limit": {
						TargetField: "FileLimit",
					},
					"capacity_limit": {
						TargetField: "CapacityLimit",
					},
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					ids := strings.Split(d.Id(), ":")
					if len(ids) != 2 {
						return false, fmt.Errorf("invalid vepfs fileset Id: %s", d.Id())
					}

					(*call.SdkParam)["FileSystemId"] = ids[0]
					(*call.SdkParam)["FilesetId"] = ids[1]
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
					resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
					logger.Debug(logger.RespFormat, call.Action, resp, err)
					return resp, err
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{"Running"},
					Timeout: resourceData.Timeout(schema.TimeoutCreate),
				},
			},
		}
		callbacks = append(callbacks, quotaCallback)
	}

	return callbacks
}

func (s *VolcengineVepfsFilesetService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	if resourceData.HasChange("fileset_name") {
		callback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "UpdateFileset",
				ConvertMode: ve.RequestConvertInConvert,
				ContentType: ve.ContentTypeJson,
				Convert: map[string]ve.RequestConvert{
					"fileset_name": {
						TargetField: "FilesetName",
					},
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					ids := strings.Split(d.Id(), ":")
					if len(ids) != 2 {
						return false, fmt.Errorf("invalid vepfs fileset Id: %s", d.Id())
					}

					(*call.SdkParam)["FileSystemId"] = ids[0]
					(*call.SdkParam)["FilesetId"] = ids[1]
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
					resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
					logger.Debug(logger.RespFormat, call.Action, resp, err)
					return resp, err
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{"Running"},
					Timeout: resourceData.Timeout(schema.TimeoutUpdate),
				},
			},
		}
		callbacks = append(callbacks, callback)
	}

	if resourceData.HasChanges("max_iops", "max_bandwidth") {
		qosCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "SetFilesetQos",
				ConvertMode: ve.RequestConvertInConvert,
				ContentType: ve.ContentTypeJson,
				Convert: map[string]ve.RequestConvert{
					"max_iops": {
						TargetField: "MaxIops",
					},
					"max_bandwidth": {
						TargetField: "MaxBandwidth",
					},
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					ids := strings.Split(d.Id(), ":")
					if len(ids) != 2 {
						return false, fmt.Errorf("invalid vepfs fileset Id: %s", d.Id())
					}

					(*call.SdkParam)["FileSystemId"] = ids[0]
					(*call.SdkParam)["FilesetId"] = ids[1]
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
					resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
					logger.Debug(logger.RespFormat, call.Action, resp, err)
					return resp, err
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{"Running"},
					Timeout: resourceData.Timeout(schema.TimeoutUpdate),
				},
			},
		}
		callbacks = append(callbacks, qosCallback)
	}

	if resourceData.HasChanges("file_limit", "capacity_limit") {
		quotaCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "SetFilesetQuota",
				ConvertMode: ve.RequestConvertInConvert,
				ContentType: ve.ContentTypeJson,
				Convert: map[string]ve.RequestConvert{
					"file_limit": {
						TargetField: "FileLimit",
					},
					"capacity_limit": {
						TargetField: "CapacityLimit",
					},
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					ids := strings.Split(d.Id(), ":")
					if len(ids) != 2 {
						return false, fmt.Errorf("invalid vepfs fileset Id: %s", d.Id())
					}

					(*call.SdkParam)["FileSystemId"] = ids[0]
					(*call.SdkParam)["FilesetId"] = ids[1]
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
					resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
					logger.Debug(logger.RespFormat, call.Action, resp, err)
					return resp, err
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{"Running"},
					Timeout: resourceData.Timeout(schema.TimeoutUpdate),
				},
			},
		}
		callbacks = append(callbacks, quotaCallback)
	}

	return callbacks
}

func (s *VolcengineVepfsFilesetService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteFileset",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				ids := strings.Split(d.Id(), ":")
				if len(ids) != 2 {
					return false, fmt.Errorf("invalid vepfs fileset Id: %s", d.Id())
				}

				(*call.SdkParam)["FileSystemId"] = ids[0]
				(*call.SdkParam)["FilesetId"] = ids[1]
				return true, nil
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
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading vepfs fileset on delete %q, %w", d.Id(), callErr))
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

func (s *VolcengineVepfsFilesetService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"status": {
				TargetField: "Filters.Status",
				ConvertType: ve.ConvertJsonArray,
			},
			"file_system_id": {
				TargetField: "FileSystemId",
			},
			"fileset_id": {
				TargetField: "Filters.FilesetId",
			},
			"fileset_name": {
				TargetField: "Filters.FilesetName",
			},
			"fileset_path": {
				TargetField: "Filters.FilesetPath",
			},
		},
		NameField:    "FilesetName",
		IdField:      "FilesetId",
		CollectField: "filesets",
		ContentType:  ve.ContentTypeJson,
		ResponseConverts: map[string]ve.ResponseConvert{
			"FilesetId": {
				TargetField: "id",
				KeepDefault: true,
			},
			"IOPSQos": {
				TargetField: "iops_qos",
			},
		},
	}
}

func (s *VolcengineVepfsFilesetService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "vepfs",
		Version:     "2022-01-01",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}
