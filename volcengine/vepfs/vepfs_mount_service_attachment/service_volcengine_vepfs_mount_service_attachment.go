package vepfs_mount_service_attachment

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

type VolcengineVepfsMountServiceAttachmentService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewVepfsMountServiceAttachmentService(c *ve.SdkClient) *VolcengineVepfsMountServiceAttachmentService {
	return &VolcengineVepfsMountServiceAttachmentService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineVepfsMountServiceAttachmentService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineVepfsMountServiceAttachmentService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		newCondition map[string]interface{}
		resp         *map[string]interface{}
		results      interface{}
		ok           bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeMountServices"

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
		results, err = ve.ObtainSdkValue("Result.MountServices", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.MountServices is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineVepfsMountServiceAttachmentService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	ids := strings.Split(id, ":")
	if len(ids) != 2 {
		return data, fmt.Errorf("invalid mount service attachment Id: %s", id)
	}

	req := map[string]interface{}{
		"Filters": map[string]interface{}{
			"MountServiceId": ids[0],
			"FileSystemId":   ids[1],
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
		return data, fmt.Errorf("vepfs_mount_service %s not exist ", id)
	}

	attached := false
	if fileSystems, exist := data["AttachFileSystems"]; exist {
		fileSystemArr, ok := fileSystems.([]interface{})
		if !ok {
			return data, fmt.Errorf(" mount serive fileSystems is not slice ")
		}
		for _, v := range fileSystemArr {
			fileSystemMap, ok := v.(map[string]interface{})
			if !ok {
				return data, fmt.Errorf(" mount serive fileSystems Value is not map ")
			}
			if fileSystemMap["FileSystemId"] == ids[1] {
				data["CustomerPath"] = fileSystemMap["CustomerPath"]
				data["AttachStatus"] = fileSystemMap["Status"]
				attached = true
				break
			}
		}
	} else {
		data["AttachFileSystems"] = []interface{}{}
	}
	if !attached {
		return data, fmt.Errorf("vepfs_mount_service_attachment %s not exist ", id)
	}

	return data, err
}

func (s *VolcengineVepfsMountServiceAttachmentService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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
			failStates = append(failStates, "AttachError")
			failStates = append(failStates, "DetachError")
			d, err = s.ReadResource(resourceData, id)
			if err != nil {
				return nil, "", err
			}
			status, err = ve.ObtainSdkValue("AttachStatus", d)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("vepfs_mount_service_attachment status error, status: %s", status.(string))
				}
			}
			return d, status.(string), err
		},
	}
}

func (VolcengineVepfsMountServiceAttachmentService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineVepfsMountServiceAttachmentService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "AttachMountServiceToSelfFileSystem",
			ConvertMode: ve.RequestConvertInConvert,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"mount_service_id": {
					TargetField: "MountServiceId",
				},
				"file_system_id": {
					TargetField: "FileSystemId",
				},
				"customer_path": {
					TargetField: "CustomerPath",
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				mountServiceId := d.Get("mount_service_id")
				fileSystemId := d.Get("file_system_id")
				d.SetId(mountServiceId.(string) + ":" + fileSystemId.(string))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Attached"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
			// 必须顺序执行，否则并发失败
			LockId: func(d *schema.ResourceData) string {
				return "lock-VePFS-attachment"
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineVepfsMountServiceAttachmentService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineVepfsMountServiceAttachmentService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DetachMountServiceFromSelfFileSystem",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				ids := strings.Split(d.Id(), ":")
				if len(ids) != 2 {
					return false, fmt.Errorf("invalid mount service attachment Id: %s", d.Id())
				}

				(*call.SdkParam)["MountServiceId"] = ids[0]
				(*call.SdkParam)["FileSystemId"] = ids[1]
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
							return resource.NonRetryableError(fmt.Errorf("error on reading vepfs mount service attachment on delete %q, %w", d.Id(), callErr))
						}
					}
					_, callErr = call.ExecuteCall(d, client, call)
					if callErr == nil {
						return nil
					}
					return resource.RetryableError(callErr)
				})
			},
			// 必须顺序执行，否则并发失败
			LockId: func(d *schema.ResourceData) string {
				return "lock-VePFS-attachment"
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineVepfsMountServiceAttachmentService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineVepfsMountServiceAttachmentService) ReadResourceId(id string) string {
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
