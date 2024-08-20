package vepfs_mount_service

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

type VolcengineVepfsMountServiceService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewVepfsMountServiceService(c *ve.SdkClient) *VolcengineVepfsMountServiceService {
	return &VolcengineVepfsMountServiceService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineVepfsMountServiceService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineVepfsMountServiceService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
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

func (s *VolcengineVepfsMountServiceService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}

	req := map[string]interface{}{
		"Filters": map[string]interface{}{
			"MountServiceId": id,
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

	if _, exist := data["AttachFileSystems"]; !exist {
		data["AttachFileSystems"] = []interface{}{}
	}

	return data, err
}

func (s *VolcengineVepfsMountServiceService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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
					return nil, "", fmt.Errorf("vepfs_mount_service status error, status: %s", status.(string))
				}
			}
			return d, status.(string), err
		},
	}
}

func (VolcengineVepfsMountServiceService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineVepfsMountServiceService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateMountService",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert:     map[string]ve.RequestConvert{},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				subnetId := d.Get("subnet_id")
				vpcId, zoneId, err := s.getVpcIdAndZoneIdBySubnet(subnetId.(string))
				if err != nil {
					return false, err
				}

				(*call.SdkParam)["VpcId"] = vpcId
				(*call.SdkParam)["ZoneId"] = zoneId
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("Result.MountServiceId", *resp)
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

func (s *VolcengineVepfsMountServiceService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateMountService",
			ConvertMode: ve.RequestConvertInConvert,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"mount_service_name": {
					TargetField: "MountServiceName",
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if len(*call.SdkParam) > 0 {
					(*call.SdkParam)["MountServiceId"] = d.Id()
					return true, nil
				}
				return false, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
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
	return []ve.Callback{callback}
}

func (s *VolcengineVepfsMountServiceService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteMountService",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"MountServiceId": resourceData.Id(),
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
							return resource.NonRetryableError(fmt.Errorf("error on reading vepfs mount service on delete %q, %w", d.Id(), callErr))
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

func (s *VolcengineVepfsMountServiceService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"status": {
				TargetField: "Filters.Status",
				ConvertType: ve.ConvertJsonArray,
			},
			"file_system_id": {
				TargetField: "Filters.FileSystemId",
			},
			"mount_service_id": {
				TargetField: "Filters.MountServiceId",
			},
			"mount_service_name": {
				TargetField: "Filters.MountServiceName",
			},
		},
		NameField:    "MountServiceName",
		IdField:      "MountServiceId",
		CollectField: "mount_services",
		ContentType:  ve.ContentTypeJson,
		ResponseConverts: map[string]ve.ResponseConvert{
			"MountServiceId": {
				TargetField: "id",
				KeepDefault: true,
			},
		},
	}
}

func (s *VolcengineVepfsMountServiceService) ReadResourceId(id string) string {
	return id
}

func (s *VolcengineVepfsMountServiceService) getVpcIdAndZoneIdBySubnet(subnetId string) (vpcId, zoneId string, err error) {
	// describe subnet
	req := map[string]interface{}{
		"SubnetIds.1": subnetId,
	}
	action := "DescribeSubnets"
	resp, err := s.Client.UniversalClient.DoCall(getVpcUniversalInfo(action), &req)
	if err != nil {
		return "", "", err
	}
	logger.Debug(logger.RespFormat, action, req, *resp)
	results, err := ve.ObtainSdkValue("Result.Subnets", *resp)
	if err != nil {
		return "", "", err
	}
	if results == nil {
		results = []interface{}{}
	}
	subnets, ok := results.([]interface{})
	if !ok {
		return "", "", errors.New("Result.Subnets is not Slice")
	}
	if len(subnets) == 0 {
		return "", "", fmt.Errorf("subnet %s not exist", subnetId)
	}
	vpcId = subnets[0].(map[string]interface{})["VpcId"].(string)
	zoneId = subnets[0].(map[string]interface{})["ZoneId"].(string)
	return vpcId, zoneId, nil
}

func (s *VolcengineVepfsMountServiceService) ProjectTrn() *ve.ProjectTrn {
	return &ve.ProjectTrn{
		ServiceName:          "vepfs",
		ResourceType:         "mountservice",
		ProjectResponseField: "Project",
		ProjectSchemaField:   "project",
	}
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

func getVpcUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "vpc",
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		Action:      actionName,
	}
}
