package rds_postgresql_instance_failover_log

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineRdsPostgresqlInstanceFailoverLogService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewRdsPostgresqlInstanceFailoverLogService(c *ve.SdkClient) *VolcengineRdsPostgresqlInstanceFailoverLogService {
	return &VolcengineRdsPostgresqlInstanceFailoverLogService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineRdsPostgresqlInstanceFailoverLogService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineRdsPostgresqlInstanceFailoverLogService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	maxResults := 500
	if v, exist := m["Limit"]; exist {
		switch t := v.(type) {
		case int:
			if t > 0 {
				maxResults = t
			}
		case int64:
			if t > 0 {
				maxResults = int(t)
			}
		case float64:
			if t > 0 {
				maxResults = int(t)
			}
		}
	}
	action := "DescribeFailoverLogs"
	return ve.WithNextTokenQuery(m, "Limit", "Context", maxResults, nil, func(condition map[string]interface{}) ([]interface{}, string, error) {
		bytes, _ := json.Marshal(condition)
		logger.Debug(logger.ReqFormat, action, string(bytes))
		if condition == nil {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), nil)
			if err != nil {
				return data, "", err
			}
		} else {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &condition)
			if err != nil {
				return data, "", err
			}
		}
		respBytes, _ := json.Marshal(resp)
		logger.Debug(logger.RespFormat, action, condition, string(respBytes))
		results, err = ve.ObtainSdkValue("Result.FailoverLogs", *resp)
		if err != nil {
			return data, "", err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, "", errors.New("Result.FailoverLogs is not Slice")
		}
		nextToken := ""
		if v, e := ve.ObtainSdkValue("Result.Context", *resp); e == nil && v != nil {
			if s, ok := v.(string); ok {
				nextToken = s
			}
		}
		return data, nextToken, err
	})
}

func (s *VolcengineRdsPostgresqlInstanceFailoverLogService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	return data, err
}

func (s *VolcengineRdsPostgresqlInstanceFailoverLogService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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

func (s *VolcengineRdsPostgresqlInstanceFailoverLogService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (VolcengineRdsPostgresqlInstanceFailoverLogService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineRdsPostgresqlInstanceFailoverLogService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineRdsPostgresqlInstanceFailoverLogService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineRdsPostgresqlInstanceFailoverLogService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"instance_id":      {TargetField: "InstanceId"},
			"query_start_time": {TargetField: "QueryStartTime"},
			"query_end_time":   {TargetField: "QueryEndTime"},
			"limit":            {TargetField: "Limit"},
			"context":          {TargetField: "Context"},
		},
		CollectField: "failover_logs",
		ResponseConverts: map[string]ve.ResponseConvert{
			"FailoverTime":    {TargetField: "failover_time"},
			"FailoverType":    {TargetField: "failover_type"},
			"NewMasterNodeId": {TargetField: "new_master_node_id"},
			"OldMasterNodeId": {TargetField: "old_master_node_id"},
		},
	}
}

func (s *VolcengineRdsPostgresqlInstanceFailoverLogService) ReadResourceId(id string) string {
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
