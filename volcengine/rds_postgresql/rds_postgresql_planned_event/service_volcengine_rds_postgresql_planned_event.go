package rds_postgresql_planned_event

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineRdsPostgresqlPlannedEventService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewRdsPostgresqlPlannedEventService(c *ve.SdkClient) *VolcengineRdsPostgresqlPlannedEventService {
	return &VolcengineRdsPostgresqlPlannedEventService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineRdsPostgresqlPlannedEventService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineRdsPostgresqlPlannedEventService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribePlannedEvents"

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
		results, err = ve.ObtainSdkValue("Result.PlannedEvents", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.PlannedEvents is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineRdsPostgresqlPlannedEventService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	return data, err
}

func (s *VolcengineRdsPostgresqlPlannedEventService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			return nil, "", err
		},
	}
}

func (s *VolcengineRdsPostgresqlPlannedEventService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{}
	return []ve.Callback{callback}
}

func (VolcengineRdsPostgresqlPlannedEventService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineRdsPostgresqlPlannedEventService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{}
	return []ve.Callback{callback}
}

func (s *VolcengineRdsPostgresqlPlannedEventService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{}
	return []ve.Callback{callback}
}

func (s *VolcengineRdsPostgresqlPlannedEventService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"instance_name": {TargetField: "InstanceName"},
			"instance_id":   {TargetField: "InstanceId"},
			"event_id":      {TargetField: "EventId"},
			"event_type": {
				TargetField: "EventType",
				ConvertType: ve.ConvertJsonArray,
			},
			"status": {
				TargetField: "Status",
				ConvertType: ve.ConvertJsonArray,
			},
			"planned_begin_time_search_range_start":  {TargetField: "PlannedBeginTimeSearchRangeStart"},
			"planned_begin_time_search_range_end":    {TargetField: "PlannedBeginTimeSearchRangeEnd"},
			"planned_switch_time_search_range_start": {TargetField: "PlannedSwitchTimeSearchRangeStart"},
			"planned_switch_time_search_range_end":   {TargetField: "PlannedSwitchTimeSearchRangeEnd"},
		},
		NameField:    "EventType",
		IdField:      "EventID",
		CollectField: "planned_events",
		ResponseConverts: map[string]ve.ResponseConvert{
			"BusinessImpact":         {TargetField: "business_impact"},
			"EventID":                {TargetField: "event_id", KeepDefault: true},
			"EventType":              {TargetField: "event_type"},
			"InstanceId":             {TargetField: "instance_id"},
			"InstanceName":           {TargetField: "instance_name"},
			"MaxDelayTime":           {TargetField: "max_delay_time"},
			"PlannedBeginTime":       {TargetField: "planned_begin_time"},
			"PlannedEventReason":     {TargetField: "planned_event_reason"},
			"PlannedSwitchBeginTime": {TargetField: "planned_switch_begin_time"},
			"PlannedSwitchEndTime":   {TargetField: "planned_switch_end_time"},
			"Region":                 {TargetField: "region"},
			"Status":                 {TargetField: "status"},
		},
	}
}

func (s *VolcengineRdsPostgresqlPlannedEventService) ReadResourceId(id string) string {
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
