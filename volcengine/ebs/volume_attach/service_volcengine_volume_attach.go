package volume_attach

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ebs/volume"
)

type VolcengineVolumeAttachService struct {
	Client        *ve.SdkClient
	volumeService *volume.VolcengineVolumeService
}

func NewVolumeAttachService(c *ve.SdkClient) *VolcengineVolumeAttachService {
	return &VolcengineVolumeAttachService{
		Client:        c,
		volumeService: volume.NewVolumeService(c),
	}
}

func (s *VolcengineVolumeAttachService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineVolumeAttachService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(condition, "PageSize", "PageNumber", 20, 1, func(m map[string]interface{}) ([]interface{}, error) {
		action := "DescribeVolumes"
		logger.Debug(logger.ReqFormat, action, condition)
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

func (s *VolcengineVolumeAttachService) ReadResource(resourceData *schema.ResourceData, volumeAttachId string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if volumeAttachId == "" {
		volumeAttachId = s.ReadResourceId(resourceData.Id())
	}

	parts := strings.Split(volumeAttachId, ":")
	req := map[string]interface{}{
		"VolumeIds.1": parts[0],
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
		return data, fmt.Errorf("volume_attach %s not exist ", volumeAttachId)
	}
	// 检查实例是否已经绑定了
	if len(data["InstanceId"].(string)) == 0 {
		return data, fmt.Errorf("volume %s does not associate instances", parts[0])
	}
	if data["InstanceId"] != parts[1] {
		return data, fmt.Errorf("volume %s does not associate instance. attached_instance_id %s, target_instance_id %s",
			parts[0], data["InstanceId"], parts[1])
	}
	return data, err
}

func (s *VolcengineVolumeAttachService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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
			failStates = append(failStates, "error")
			demo, err = s.ReadResource(resourceData, id)
			if err != nil && !strings.Contains(err.Error(), "not associate") {
				return nil, "", err
			}
			status, err = ve.ObtainSdkValue("Status", demo)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("volume status error, status:%s", status.(string))
				}
			}
			return demo, status.(string), err
		},
	}
}

func (VolcengineVolumeAttachService) WithResourceResponseHandlers(volume map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return volume, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineVolumeAttachService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "AttachVolume",
			ConvertMode: ve.RequestConvertAll,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				_, exist := d.GetOkExists("delete_with_instance")
				if exist {
					return true, nil
				} else {
					volumeId := resourceData.Get("volume_id")
					volume, err := s.volumeService.ReadResource(resourceData, volumeId.(string))
					if err != nil {
						return false, err
					}
					deleteWithInstance, ok := volume["DeleteWithInstance"]
					if !ok {
						return false, fmt.Errorf(" DeleteWithInstance is not exist in volume ")
					}
					(*call.SdkParam)["DeleteWithInstance"] = deleteWithInstance
					return true, nil
				}
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				d.SetId(fmt.Sprint((*call.SdkParam)["VolumeId"], ":", (*call.SdkParam)["InstanceId"]))
				return nil
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				volume.NewVolumeService(s.Client): {
					Target:     []string{"attached"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					ResourceId: resourceData.Get("volume_id").(string),
				},
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineVolumeAttachService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	strs := strings.Split(resourceData.Id(), ":")
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DetachVolume",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"VolumeId":   strs[0],
				"instanceId": strs[1],
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading volume on delete %q, %w", d.Id(), callErr))
						}
					}
					_, callErr = call.ExecuteCall(d, client, call)
					if callErr == nil {
						return nil
					}
					return resource.RetryableError(callErr)
				})
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				volume.NewVolumeService(s.Client): {
					Target:     []string{"available"},
					Timeout:    resourceData.Timeout(schema.TimeoutDelete),
					ResourceId: resourceData.Get("volume_id").(string),
				},
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineVolumeAttachService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineVolumeAttachService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func importVolumeAttach(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	var err error
	items := strings.Split(d.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{d}, fmt.Errorf("import id must be of the form VolumeId:instanceId")
	}
	err = d.Set("volume_id", items[0])
	if err != nil {
		return []*schema.ResourceData{d}, err
	}
	err = d.Set("instance_id", items[1])
	if err != nil {
		return []*schema.ResourceData{d}, err
	}

	return []*schema.ResourceData{d}, nil
}

func (s *VolcengineVolumeAttachService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "storage_ebs",
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		Action:      actionName,
	}
}
