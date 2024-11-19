package ebs_auto_snapshot_policy_attachment

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineEbsAutoSnapshotPolicyAttachmentService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewEbsAutoSnapshotPolicyAttachmentService(c *ve.SdkClient) *VolcengineEbsAutoSnapshotPolicyAttachmentService {
	return &VolcengineEbsAutoSnapshotPolicyAttachmentService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineEbsAutoSnapshotPolicyAttachmentService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineEbsAutoSnapshotPolicyAttachmentService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeAutoSnapshotPolicy"

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

		results, err = ve.ObtainSdkValue("Result.AutoSnapshotPolicies", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.AutoSnapshotPolicies is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineEbsAutoSnapshotPolicyAttachmentService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	ids := strings.Split(id, ":")
	if len(ids) != 2 {
		return data, fmt.Errorf("invalid auto_snapshot_policy_attachment id: %s", id)
	}

	req := map[string]interface{}{
		"AutoSnapshotPolicyIds.1": ids[0],
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
		return data, fmt.Errorf("ebs_auto_snapshot_policy_attachment %s not exist ", id)
	}

	data["VolumeId"] = ids[1]
	return data, err
}

func (s *VolcengineEbsAutoSnapshotPolicyAttachmentService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (VolcengineEbsAutoSnapshotPolicyAttachmentService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineEbsAutoSnapshotPolicyAttachmentService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ApplyAutoSnapshotPolicy",
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"volume_id": {
					TargetField: "VolumeIds.1",
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				policyId := d.Get("auto_snapshot_policy_id").(string)
				volumeId := d.Get("volume_id").(string)
				d.SetId(policyId + ":" + volumeId)
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineEbsAutoSnapshotPolicyAttachmentService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineEbsAutoSnapshotPolicyAttachmentService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CancelAutoSnapshotPolicy",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				ids := strings.Split(d.Id(), ":")
				if len(ids) != 2 {
					return false, fmt.Errorf("invalid auto_snapshot_policy_attachment id: %s", d.Id())
				}
				(*call.SdkParam)["AutoSnapshotPolicyId"] = ids[0]
				(*call.SdkParam)["VolumeIds.1"] = ids[1]
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineEbsAutoSnapshotPolicyAttachmentService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineEbsAutoSnapshotPolicyAttachmentService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "storage_ebs",
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		Action:      actionName,
	}
}
