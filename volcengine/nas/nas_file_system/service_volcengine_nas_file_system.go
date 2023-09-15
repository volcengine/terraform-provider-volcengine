package nas_file_system

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineNasFileSystemService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewNasFileSystemService(c *ve.SdkClient) *VolcengineNasFileSystemService {
	return &VolcengineNasFileSystemService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineNasFileSystemService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineNasFileSystemService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 20, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeFileSystems"

		// 处理 FileSystemIds，逗号分离
		if ids, exists := condition["FileSystemIds"]; exists {
			idsArr, ok := ids.([]interface{})
			if !ok {
				return data, fmt.Errorf(" FileSystemIds is not slice ")
			}
			fileSystemIds := make([]string, 0)
			for _, id := range idsArr {
				fileSystemIds = append(fileSystemIds, id.(string))
			}
			condition["FileSystemIds"] = strings.Join(fileSystemIds, ",")
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
				filtersMap["Status"] = strings.Join(newStatus, ",")
			}

			newFilters := make([]interface{}, 0)
			for key, value := range filtersMap {
				newFilters = append(newFilters, map[string]interface{}{
					"Key":   key,
					"Value": value,
				})
			}
			condition["Filters"] = newFilters
		}

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
		results, err = ve.ObtainSdkValue("Result.FileSystems", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.FileSystems is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineNasFileSystemService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"FileSystemIds": []interface{}{id},
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
		return data, fmt.Errorf("nas file system %s is not exist ", id)
	}

	if capacity, exist := data["Capacity"]; exist {
		capacityMap, ok := capacity.(map[string]interface{})
		if !ok {
			return data, fmt.Errorf(" file system capacity is not map ")
		}
		data["Capacity"] = capacityMap["Total"]
	}

	if snapshotId, ok := resourceData.GetOk("snapshot_id"); ok {
		data["SnapshotId"] = snapshotId
	}
	if tags, exist := data["Tags"]; exist {
		tagsArr, ok := tags.([]interface{})
		if !ok {
			return data, fmt.Errorf(" file system tags is not slice ")
		}
		newTags := make([]interface{}, 0)
		for _, tag := range tagsArr {
			tagMap, ok := tag.(map[string]interface{})
			if !ok {
				return data, fmt.Errorf(" file system tag is not map")
			}
			if tagMap["Type"].(string) == "Custom" {
				newTags = append(newTags, map[string]interface{}{
					"Key":   tagMap["Key"],
					"Value": tagMap["Value"],
				})
			}
		}
		data["Tags"] = newTags
	}

	return data, err
}

func (s *VolcengineNasFileSystemService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				demo       map[string]interface{}
				status     interface{}
				failStates []string
			)
			failStates = append(failStates, "Error")

			if err = resource.Retry(5*time.Minute, func() *resource.RetryError {
				demo, err = s.ReadResource(resourceData, id)
				if err != nil {
					if ve.ResourceNotFoundError(err) {
						return resource.RetryableError(err)
					} else {
						return resource.NonRetryableError(err)
					}
				}
				return nil
			}); err != nil {
				return nil, "", err
			}

			status, err = ve.ObtainSdkValue("Status", demo)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("nas file system status error, status: %s", status.(string))
				}
			}
			//注意 返回的第一个参数不能为空 否则会一直等下去
			return demo, status.(string), err
		},
	}
}

func (VolcengineNasFileSystemService) WithResourceResponseHandlers(data map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return data, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineNasFileSystemService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateFileSystem",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"tags": {
					TargetField: "Tags",
					ConvertType: ve.ConvertJsonObjectArray,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["ChargeType"] = "PayAsYouGo"
				(*call.SdkParam)["FileSystemType"] = "Extreme"
				(*call.SdkParam)["ProtocolType"] = "NFS"
				(*call.SdkParam)["ClientToken"] = uuid.New().String()
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				// 将 Tags 的属性赋值为 Custom
				if tags, exist := (*call.SdkParam)["Tags"]; exist {
					tagsArr, ok := tags.([]interface{})
					if !ok {
						return nil, fmt.Errorf(" file system tags is not slice ")
					}
					for _, tag := range tagsArr {
						tagMap, ok := tag.(map[string]interface{})
						if !ok {
							return nil, fmt.Errorf(" file system tag is not map")
						}
						tagMap["Type"] = "Custom"
					}
				}

				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("Result.FileSystemId", *resp)
				d.SetId(id.(string))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Running"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineNasFileSystemService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	if resourceData.HasChanges("file_system_name", "description", "project_name", "tags") {
		callback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "UpdateFileSystem",
				ConvertMode: ve.RequestConvertInConvert,
				ContentType: ve.ContentTypeJson,
				Convert: map[string]ve.RequestConvert{
					"file_system_name": {
						TargetField: "FileSystemName",
					},
					"description": {
						TargetField: "Description",
					},
					"project_name": {
						TargetField: "ProjectName",
					},
					"tags": {
						TargetField: "Tags",
						ForceGet:    true,
						ConvertType: ve.ConvertJsonObjectArray,
					},
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					if len(*call.SdkParam) > 0 {
						(*call.SdkParam)["FileSystemId"] = d.Id()
						return true, nil
					}
					return false, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					if d.HasChange("tags") {
						if tags, exists := (*call.SdkParam)["Tags"]; exists {
							tagsArr, ok := tags.([]interface{})
							if !ok {
								return nil, fmt.Errorf(" file system tags is not slice ")
							}
							for _, tag := range tagsArr {
								tagMap, ok := tag.(map[string]interface{})
								if !ok {
									return nil, fmt.Errorf(" file system tag is not map")
								}
								tagMap["Type"] = "Custom"
							}
						} else {
							// 当 Tags 被删除时，入参添加空列表来置空
							(*call.SdkParam)["Tags"] = []interface{}{}
						}
					}

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

	if resourceData.HasChange("capacity") {
		callback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ExpandFileSystem",
				ConvertMode: ve.RequestConvertInConvert,
				ContentType: ve.ContentTypeJson,
				Convert: map[string]ve.RequestConvert{
					"capacity": {
						TargetField: "Capacity",
					},
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					if len(*call.SdkParam) > 0 {
						(*call.SdkParam)["FileSystemId"] = d.Id()
						return true, nil
					}
					return false, nil
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

	return callbacks
}

func (s *VolcengineNasFileSystemService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteFileSystem",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeDefault,
			SdkParam: &map[string]interface{}{
				"FileSystemId": resourceData.Id(),
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
							return resource.NonRetryableError(fmt.Errorf("error on  reading nas file system on delete %q, %w", d.Id(), callErr))
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

func (s *VolcengineNasFileSystemService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "FileSystemIds",
				ConvertType: ve.ConvertJsonArray,
			},
			"status": {
				TargetField: "Filters.Status",
				ConvertType: ve.ConvertJsonArray,
			},
			"file_system_name": {
				TargetField: "Filters.FileSystemName",
			},
			"zone_id": {
				TargetField: "Filters.ZoneId",
			},
			"protocol_type": {
				TargetField: "Filters.ProtocolType",
			},
			"storage_type": {
				TargetField: "Filters.StorageType",
			},
			"charge_type": {
				TargetField: "Filters.ChargeType",
			},
			"permission_id": {
				TargetField: "Filters.PermissionId",
			},
			"mount_point_id": {
				TargetField: "Filters.MountPointId",
			},
			"tags": {
				TargetField: "TagFilters",
				ConvertType: ve.ConvertJsonObjectArray,
			},
		},
		NameField:    "FileSystemName",
		IdField:      "FileSystemId",
		CollectField: "file_systems",
		ContentType:  ve.ContentTypeJson,
		ResponseConverts: map[string]ve.ResponseConvert{
			"FileSystemId": {
				TargetField: "id",
				KeepDefault: true,
			},
		},
	}
}

func (s *VolcengineNasFileSystemService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "FileNAS",
		Version:     "2022-01-01",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}
