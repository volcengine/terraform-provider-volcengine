package rds_postgresql_instance_task

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineRdsPostgresqlInstanceTaskService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewRdsPostgresqlInstanceTaskService(c *ve.SdkClient) *VolcengineRdsPostgresqlInstanceTaskService {
	return &VolcengineRdsPostgresqlInstanceTaskService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineRdsPostgresqlInstanceTaskService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineRdsPostgresqlInstanceTaskService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeTasks"

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
		results, err = ve.ObtainSdkValue("Result.TaskInfos", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.TaskInfos is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineRdsPostgresqlInstanceTaskService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	return data, err
}

func (s *VolcengineRdsPostgresqlInstanceTaskService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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

func (s *VolcengineRdsPostgresqlInstanceTaskService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{}
	return []ve.Callback{callback}
}

func (VolcengineRdsPostgresqlInstanceTaskService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineRdsPostgresqlInstanceTaskService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{}
	return []ve.Callback{callback}
}

func (s *VolcengineRdsPostgresqlInstanceTaskService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{}
	return []ve.Callback{callback}
}

func (s *VolcengineRdsPostgresqlInstanceTaskService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"instance_id":         {TargetField: "InstanceId"},
			"task_id":             {TargetField: "TaskId"},
			"task_action":         {TargetField: "TaskAction"},
			"creation_start_time": {TargetField: "CreationStartTime"},
			"creation_end_time":   {TargetField: "CreationEndTime"},
			"task_status": {
				TargetField: "TaskStatus",
				ConvertType: ve.ConvertJsonArray,
			},
			"project_name": {TargetField: "ProjectName"},
		},
		NameField:    "TaskAction",
		IdField:      "TaskId",
		CollectField: "task_infos",
		ResponseConverts: map[string]ve.ResponseConvert{
			"CostTimeMS":               {TargetField: "cost_time_ms"},
			"CreateTime":               {TargetField: "create_time"},
			"FinishTime":               {TargetField: "finish_time"},
			"InstanceId":               {TargetField: "instance_id"},
			"ProjectName":              {TargetField: "project_name"},
			"Region":                   {TargetField: "region"},
			"TaskAction":               {TargetField: "task_action"},
			"TaskId":                   {TargetField: "task_id", KeepDefault: true},
			"TaskParams":               {TargetField: "task_params"},
			"TaskStatus":               {TargetField: "task_status"},
			"ScheduledSwitchEndTime":   {TargetField: "scheduled_switch_end_time"},
			"ScheduledSwitchStartTime": {TargetField: "scheduled_switch_start_time"},
		},
	}
}

func (s *VolcengineRdsPostgresqlInstanceTaskService) ReadResourceId(id string) string {
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
