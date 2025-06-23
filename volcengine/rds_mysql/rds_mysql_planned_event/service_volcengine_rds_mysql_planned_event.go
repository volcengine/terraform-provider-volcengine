package rds_mysql_planned_event

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineRdsMysqlPlannedEventService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewRdsMysqlPlannedEventService(c *ve.SdkClient) *VolcengineRdsMysqlPlannedEventService {
	return &VolcengineRdsMysqlPlannedEventService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineRdsMysqlPlannedEventService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineRdsMysqlPlannedEventService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithSimpleQuery(m, func(condition map[string]interface{}) ([]interface{}, error) {
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

func (s *VolcengineRdsMysqlPlannedEventService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {

	return data, err
}

func (s *VolcengineRdsMysqlPlannedEventService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
	}
}

func (s *VolcengineRdsMysqlPlannedEventService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{}
	return []ve.Callback{callback}
}

func (VolcengineRdsMysqlPlannedEventService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineRdsMysqlPlannedEventService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{}
	return []ve.Callback{callback}
}

func (s *VolcengineRdsMysqlPlannedEventService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{}
	return []ve.Callback{callback}
}

func (s *VolcengineRdsMysqlPlannedEventService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"status": {
				TargetField: "Status",
				ConvertType: ve.ConvertJsonArray,
			},
			"event_type": {
				TargetField: "EventType",
				ConvertType: ve.ConvertJsonArray,
			},
		},
		ContentType:  ve.ContentTypeJson,
		NameField:    "EventName",
		IdField:      "EventID",
		CollectField: "planned_events",
		ResponseConverts: map[string]ve.ResponseConvert{
			"EventID": {
				TargetField: "event_id",
			},
			"DBEngine": {
				TargetField: "db_engine",
			},
		},
	}
}

func (s *VolcengineRdsMysqlPlannedEventService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "rds_mysql",
		Version:     "2022-01-01",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}
