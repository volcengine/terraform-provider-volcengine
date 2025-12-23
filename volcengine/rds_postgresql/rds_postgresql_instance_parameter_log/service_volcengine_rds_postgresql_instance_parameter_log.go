package rds_postgresql_instance_parameter_log

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineRdsPostgresqlInstanceParameterLogService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewRdsPostgresqlInstanceParameterLogService(c *ve.SdkClient) *VolcengineRdsPostgresqlInstanceParameterLogService {
	return &VolcengineRdsPostgresqlInstanceParameterLogService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineRdsPostgresqlInstanceParameterLogService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineRdsPostgresqlInstanceParameterLogService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeDBInstanceParametersLog"

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

		results, err = ve.ObtainSdkValue("Result.ParameterChangeLogs", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.ParameterChangeLogs is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineRdsPostgresqlInstanceParameterLogService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		exist   bool
	)
	req := map[string]interface{}{}
	if v, ok := resourceData.GetOk("instance_id"); ok {
		req["InstanceId"] = v
	}
	if v, ok := resourceData.GetOk("start_time"); ok {
		req["StartTime"] = v
	}
	if v, ok := resourceData.GetOk("end_time"); ok {
		req["EndTime"] = v
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		if data, exist = v.(map[string]interface{}); !exist {
			return data, errors.New("value is not map ")
		} else {
			break
		}
	}
	return data, err
}

func (s *VolcengineRdsPostgresqlInstanceParameterLogService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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

func (s *VolcengineRdsPostgresqlInstanceParameterLogService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (VolcengineRdsPostgresqlInstanceParameterLogService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineRdsPostgresqlInstanceParameterLogService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineRdsPostgresqlInstanceParameterLogService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineRdsPostgresqlInstanceParameterLogService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"instance_id": {
				TargetField: "InstanceId",
			},
			"start_time": {
				TargetField: "StartTime",
			},
			"end_time": {
				TargetField: "EndTime",
			},
		},
		CollectField: "parameter_change_logs",
	}
}

func (s *VolcengineRdsPostgresqlInstanceParameterLogService) ReadResourceId(id string) string {
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
