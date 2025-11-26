package rds_postgresql_replication_slot

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

type VolcengineRdsPostgresqlReplicationSlotService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewRdsPostgresqlReplicationSlotService(c *ve.SdkClient) *VolcengineRdsPostgresqlReplicationSlotService {
	return &VolcengineRdsPostgresqlReplicationSlotService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineRdsPostgresqlReplicationSlotService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineRdsPostgresqlReplicationSlotService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithSimpleQuery(m, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeSlots"

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
		results, err = ve.ObtainSdkValue("Result.ReplicationSlots", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.ReplicationSlots is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineRdsPostgresqlReplicationSlotService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	// 资源 ID=InstanceId:SlotName
	ids := strings.Split(id, ":")
	req := map[string]interface{}{
		"InstanceId": ids[0],
		"SlotName":   ids[1],
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	target := ids[1]
	for _, v := range results {
		m, _ := v.(map[string]interface{})
		if m == nil {
			continue
		}
		nameVal, _ := m["SlotName"].(string)
		if nameVal == target {
			data = m
			break
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("rds_postgresql_replication_slot %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineRdsPostgresqlReplicationSlotService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      5 * time.Second,
		MinTimeout: 5 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				d          map[string]interface{}
				status     interface{}
				failStates []string
			)
			failStates = append(failStates, "Failed")
			d, err = s.ReadResource(resourceData, id)
			if err != nil {
				return nil, "", err
			}
			status, err = ve.ObtainSdkValue("SlotStatus", d)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("rds_postgresql_replication_slot status error, status: %s", status.(string))
				}
			}
			return d, fmt.Sprintf("%v", status), err
		},
	}
}

func (s *VolcengineRdsPostgresqlReplicationSlotService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (VolcengineRdsPostgresqlReplicationSlotService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"SlotName":   {TargetField: "slot_name"},
			"Plugin":     {TargetField: "plugin"},
			"SlotType":   {TargetField: "slot_type"},
			"SlotStatus": {TargetField: "slot_status"},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineRdsPostgresqlReplicationSlotService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineRdsPostgresqlReplicationSlotService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteSlot",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam:    &map[string]interface{}{},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["InstanceId"] = d.Get("instance_id").(string)
				(*call.SdkParam)["SlotName"] = d.Get("slot_name").(string)
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return resource.Retry(2*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, d.Id())
					if callErr == nil {
						return resource.RetryableError(fmt.Errorf("resource still in removing status "))
					}
					if ve.ResourceNotFoundError(callErr) {
						return nil
					}
					// 当接口不返回明确 not found 时，进一步判断列表里是否已不存在
					ids := strings.Split(d.Id(), ":")
					req := map[string]interface{}{
						"InstanceId": ids[0],
						"SlotName":   ids[1],
					}
					list, err := s.ReadResources(req)
					if err == nil && len(list) == 0 {
						return nil
					}
					return resource.NonRetryableError(callErr)
				})
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineRdsPostgresqlReplicationSlotService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"instance_id": {TargetField: "InstanceId"},
			"slot_name":   {TargetField: "SlotName"},
			"slot_type":   {TargetField: "SlotType"},
			"plugin":      {TargetField: "Plugin"},
			"data_base":   {TargetField: "Database"},
			"temporary":   {TargetField: "Temporary"},
			"slot_status": {TargetField: "SlotStatus"},
			"ip_address":  {TargetField: "IPAddress"},
		},
		IdField:      "SlotName",
		CollectField: "replication_slots",
		ResponseConverts: map[string]ve.ResponseConvert{
			"SlotName":   {TargetField: "slot_name", KeepDefault: true},
			"Plugin":     {TargetField: "plugin"},
			"SlotType":   {TargetField: "slot_type"},
			"SlotStatus": {TargetField: "slot_status"},
			"Database":   {TargetField: "data_base"},
			"Temporary":  {TargetField: "temporary"},
			"WalDelay":   {TargetField: "wal_delay"},
			"IPAddress":  {TargetField: "ip_address"},
		},
	}
}

func (s *VolcengineRdsPostgresqlReplicationSlotService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "rds_postgresql",
		Version:     "2022-01-01",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}
