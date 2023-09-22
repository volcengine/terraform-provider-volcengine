package nas_snapshot

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/mitchellh/copystructure"
	volc "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineNasSnapshotService struct {
	Client *volc.SdkClient
}

func NewService(c *volc.SdkClient) *VolcengineNasSnapshotService {
	return &VolcengineNasSnapshotService{
		Client: c,
	}
}

func (s *VolcengineNasSnapshotService) GetClient() *volc.SdkClient {
	return s.Client
}

func interfaceSlice2StringSlice(data []interface{}) []string {
	var res []string
	for _, ele := range data {
		res = append(res, ele.(string))
	}
	return res
}

func (s *VolcengineNasSnapshotService) ReadResources(m map[string]interface{}) ([]interface{}, error) {
	return volc.WithPageNumberQuery(m, "PageSize", "PageNumber", 20, 0, func(condition map[string]interface{}) (data []interface{}, err error) {
		var (
			newCondition map[string]interface{}
			resp         *map[string]interface{}
			results      interface{}
			ok           bool
		)

		deepCopyValue, err := copystructure.Copy(condition)
		if err != nil {
			return data, fmt.Errorf(" DeepCopy condition error: %v ", err)
		}
		if newCondition, ok = deepCopyValue.(map[string]interface{}); !ok {
			return data, fmt.Errorf(" DeepCopy condition error: newCondition is not map ")
		}

		// 处理 SnapshotIds，逗号分离
		if v, ok := condition["SnapshotIds"]; ok {
			ids, ok := v.([]interface{})
			if !ok {
				return data, fmt.Errorf(" SnapshotIds is not slice ")
			}
			if len(ids) > 0 {
				newCondition["SnapshotIds"] = strings.Join(interfaceSlice2StringSlice(ids), ",")
			}
		}

		universalClient := s.Client.UniversalClient
		action := "DescribeSnapshots"
		logger.Debug(logger.ReqFormat, action, newCondition)
		if newCondition == nil {
			resp, err = universalClient.DoCall(getUniversalInfo(action), nil)
			if err != nil {
				return data, err
			}
		} else {
			resp, err = universalClient.DoCall(getUniversalInfo(action), &newCondition)
			if err != nil {
				return data, err
			}
		}
		logger.Debug(logger.RespFormat, action, newCondition, *resp)

		results, err = volc.ObtainSdkValue("Result.Snapshots", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.Snapshots is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineNasSnapshotService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	req := map[string]interface{}{
		"SnapshotIds": []interface{}{
			id,
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
		return data, fmt.Errorf("Snapshot %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineNasSnapshotService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				data       map[string]interface{}
				status     interface{}
				failStates []string
			)
			failStates = append(failStates, "Error")
			data, err = s.ReadResource(resourceData, id)
			if err != nil {
				return nil, "", err
			}
			status, err = volc.ObtainSdkValue("Status", data)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("Snapshot status  error, status:%s", status.(string))
				}
			}
			return data, status.(string), err
		},
	}
}

func (VolcengineNasSnapshotService) WithResourceResponseHandlers(rdsAccount map[string]interface{}) []volc.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]volc.ResponseConvert, error) {
		return rdsAccount, nil, nil
	}
	return []volc.ResourceResponseHandler{handler}

}

func (s *VolcengineNasSnapshotService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []volc.Callback {
	callback := volc.Callback{
		Call: volc.SdkCall{
			Action:      "CreateSnapshot",
			ConvertMode: volc.RequestConvertAll,
			ContentType: volc.ContentTypeJson,
			ExecuteCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *volc.SdkClient, resp *map[string]interface{}, call volc.SdkCall) error {
				id, _ := volc.ObtainSdkValue("Result.SnapshotId", *resp)
				d.SetId(id.(string))
				return nil
			},
			Refresh: &volc.StateRefresh{
				Target:  []string{"Accomplished"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("file_system_id").(string)
			},
		},
	}
	return []volc.Callback{callback}
}

func (s *VolcengineNasSnapshotService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []volc.Callback {
	var callbacks []volc.Callback

	callback := volc.Callback{
		Call: volc.SdkCall{
			Action:      "UpdateSnapshot",
			ConvertMode: volc.RequestConvertAll,
			ContentType: volc.ContentTypeJson,
			BeforeCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (bool, error) {
				(*call.SdkParam)["SnapshotId"] = d.Id()
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			Refresh: &volc.StateRefresh{
				Target:  []string{"Accomplished"},
				Timeout: resourceData.Timeout(schema.TimeoutUpdate),
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("file_system_id").(string)
			},
		},
	}
	callbacks = append(callbacks, callback)
	return callbacks
}

func (s *VolcengineNasSnapshotService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []volc.Callback {
	callback := volc.Callback{
		Call: volc.SdkCall{
			Action:      "DeleteSnapshot",
			ContentType: volc.ContentTypeJson,
			ConvertMode: volc.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (bool, error) {
				(*call.SdkParam)["SnapshotId"] = d.Id()
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *volc.SdkClient, call volc.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	return []volc.Callback{callback}
}

func (s *VolcengineNasSnapshotService) DatasourceResources(*schema.ResourceData, *schema.Resource) volc.DataSourceInfo {
	return volc.DataSourceInfo{
		ContentType:  volc.ContentTypeJson,
		NameField:    "SnapshotName",
		CollectField: "snapshots",
		RequestConverts: map[string]volc.RequestConvert{
			"ids": {
				TargetField: "SnapshotIds",
				ConvertType: volc.ConvertJsonArray,
			},
		},
		ResponseConverts: map[string]volc.ResponseConvert{
			"SnapshotId": {
				TargetField: "id",
				KeepDefault: true,
			},
		},
	}
}

func (s *VolcengineNasSnapshotService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) volc.UniversalInfo {
	return volc.UniversalInfo{
		ServiceName: "filenas",
		Version:     "2022-01-01",
		HttpMethod:  volc.POST,
		ContentType: volc.ApplicationJSON,
		Action:      actionName,
	}
}
