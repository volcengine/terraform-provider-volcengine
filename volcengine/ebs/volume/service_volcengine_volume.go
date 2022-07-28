package volume

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineVolumeService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewVolumeService(c *ve.SdkClient) *VolcengineVolumeService {
	return &VolcengineVolumeService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineVolumeService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineVolumeService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(condition, "PageSize", "PageNumber", 20, 1, func(m map[string]interface{}) ([]interface{}, error) {
		ebs := s.Client.EbsClient
		action := "DescribeVolumes"
		logger.Debug(logger.ReqFormat, action, condition)
		if condition == nil {
			resp, err = ebs.DescribeVolumesCommon(nil)
			if err != nil {
				return data, err
			}
		} else {
			resp, err = ebs.DescribeVolumesCommon(&condition)
			if err != nil {
				return data, err
			}
		}

		results, err = ve.ObtainSdkValue("Result.Volumes", *resp)
		if err != nil {
			return data, err
		}
		logger.Debug(logger.ReqFormat, action, results)
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.Volumes is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineVolumeService) ReadResource(resourceData *schema.ResourceData, volumeId string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if volumeId == "" {
		volumeId = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"VolumeIds.1": volumeId,
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
		return data, fmt.Errorf("volume %s not exist ", volumeId)
	}
	return data, err
}

func (s *VolcengineVolumeService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				ebs        map[string]interface{}
				status     interface{}
				failStates []string
			)
			failStates = append(failStates, "error")

			if err = resource.Retry(20*time.Minute, func() *resource.RetryError {
				ebs, err = s.ReadResource(resourceData, id)
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

			status, err = ve.ObtainSdkValue("Status", ebs)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("volume status error, status:%s", status.(string))
				}
			}
			return ebs, status.(string), err
		},
	}
}

func (VolcengineVolumeService) WithResourceResponseHandlers(volume map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return volume, map[string]ve.ResponseConvert{
			"Size": {
				TargetField: "size",
				Convert:     sizeConvertFunc,
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineVolumeService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateVolume",
			ConvertMode: ve.RequestConvertAll,
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.EbsClient.CreateVolumeCommon(call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("Result.VolumeId", *resp)
				d.SetId(id.(string))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"available"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineVolumeService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback
	if resourceData.HasChanges("volume_name", "description", "delete_with_instance") {
		callbacks = append(callbacks, ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyVolumeAttribute",
				ConvertMode: ve.RequestConvertAll,
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["VolumeId"] = d.Id()
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
					return s.Client.EbsClient.ModifyVolumeAttributeCommon(call.SdkParam)
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{"available", "attached"},
					Timeout: resourceData.Timeout(schema.TimeoutUpdate),
				},
			},
		})
	}

	if resourceData.HasChange("size") { // 调用新的 api
		callbacks = append(callbacks, ve.Callback{
			Call: ve.SdkCall{
				Action:      "ExtendVolume",
				ConvertMode: ve.RequestConvertIgnore,
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["VolumeId"] = d.Id()
					(*call.SdkParam)["NewSize"] = d.Get("size")
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
					return s.Client.EbsClient.ExtendVolumeCommon(call.SdkParam)
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{"available", "attached"},
					Timeout: resourceData.Timeout(schema.TimeoutUpdate),
				},
			},
		})
	}
	return callbacks
}

func (s *VolcengineVolumeService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteVolume",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"VolumeId": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.EbsClient.DeleteVolumeCommon(call.SdkParam)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on  reading vpc on delete %q, %w", d.Id(), callErr))
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

func (s *VolcengineVolumeService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "VolumeIds",
				ConvertType: ve.ConvertWithN,
			},
		},
		NameField:    "VolumeName",
		IdField:      "VolumeId",
		CollectField: "volumes",
		ResponseConverts: map[string]ve.ResponseConvert{
			"VolumeId": {
				TargetField: "id",
				KeepDefault: true,
			},
			"Size": {
				TargetField: "size",
				Convert:     sizeConvertFunc,
			},
		},
	}
}

var sizeConvertFunc = func(i interface{}) interface{} {
	// Notice: the type of filed Size in openapi doc is size, but api return type is string
	size, ok := i.(string)
	if !ok {
		return i
	}
	res, err := strconv.Atoi(size)
	if err != nil {
		logger.Debug(logger.ReqFormat, "sizeConvertFunc", i)
		return i
	}
	return res
}

func (s *VolcengineVolumeService) ReadResourceId(id string) string {
	return id
}
