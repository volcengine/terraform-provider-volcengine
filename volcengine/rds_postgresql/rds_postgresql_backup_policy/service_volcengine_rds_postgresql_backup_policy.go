package rds_postgresql_backup_policy

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineRdsPostgresqlBackupPolicyService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewRdsPostgresqlBackupPolicyService(c *ve.SdkClient) *VolcengineRdsPostgresqlBackupPolicyService {
	return &VolcengineRdsPostgresqlBackupPolicyService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineRdsPostgresqlBackupPolicyService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineRdsPostgresqlBackupPolicyService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
	)
	return ve.WithSimpleQuery(m, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeBackupPolicy"

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
		results, err = ve.ObtainSdkValue("Result", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = map[string]interface{}{}
		}
		if item, ok := results.(map[string]interface{}); ok {
			return []interface{}{item}, nil
		}
		return []interface{}{}, nil
	})
}

func (s *VolcengineRdsPostgresqlBackupPolicyService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	// 以实例Id作为资源id
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"InstanceId": id,
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
		return data, fmt.Errorf("rds_postgresql_backup_policy %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineRdsPostgresqlBackupPolicyService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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

func (s *VolcengineRdsPostgresqlBackupPolicyService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	// 不支持创建
	return []ve.Callback{}
}

func (VolcengineRdsPostgresqlBackupPolicyService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"BackupRetentionPeriod":    {TargetField: "backup_retention_period"},
			"DataIncrBackupPeriods":    {TargetField: "data_incr_backup_periods"},
			"FullBackupTime":           {TargetField: "full_backup_time"},
			"FullBackupPeriod":         {TargetField: "full_backup_period"},
			"HourlyIncrBackupEnable":   {TargetField: "hourly_incr_backup_enable"},
			"IncrementBackupFrequency": {TargetField: "increment_backup_frequency"},
			"InstanceId":               {TargetField: "instance_id"},
			"WALLogSpaceLimitEnable":   {TargetField: "wal_log_space_limit_enable"},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineRdsPostgresqlBackupPolicyService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyBackupPolicy",
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"backup_retention_period":    {TargetField: "BackupRetentionPeriod"},
				"data_incr_backup_periods":   {TargetField: "DataIncrBackupPeriods"},
				"full_backup_time":           {TargetField: "FullBackupTime"},
				"full_backup_period":         {TargetField: "FullBackupPeriod"},
				"hourly_incr_backup_enable":  {TargetField: "HourlyIncrBackupEnable"},
				"increment_backup_frequency": {TargetField: "IncrementBackupFrequency"},
				"instance_id":                {TargetField: "InstanceId"},
				"wal_log_space_limit_enable": {TargetField: "WALLogSpaceLimitEnable"},
			},
			SdkParam: &map[string]interface{}{},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				// Ensure InstanceId is provided
				if v, ok := (*call.SdkParam)["InstanceId"]; !ok || v == nil || v == "" {
					inst := d.Get("instance_id").(string)
					if inst == "" {
						inst = d.Id()
					}
					(*call.SdkParam)["InstanceId"] = inst
				}
				if v, ok := d.Get("backup_retention_period").(string); ok {
					(*call.SdkParam)["BackupRetentionPeriod"] = v
				}
				if v, ok := d.Get("full_backup_time").(string); ok {
					(*call.SdkParam)["FullBackupTime"] = v
				}
				if v, ok := d.Get("data_incr_backup_periods").(string); ok {
					(*call.SdkParam)["DataIncrBackupPeriods"] = v
				}
				if v, ok := d.Get("full_backup_period").(string); ok {
					(*call.SdkParam)["FullBackupPeriod"] = v
				}
				if v, ok := d.Get("increment_backup_frequency").(string); ok {
					(*call.SdkParam)["IncrementBackupFrequency"] = v
				}
				if v, ok := d.Get("hourly_incr_backup_enable").(string); ok {
					(*call.SdkParam)["HourlyIncrBackupEnable"] = v
				}
				if v, ok := d.Get("wal_log_space_limit_enable").(string); ok {
					(*call.SdkParam)["WALLogSpaceLimitEnable"] = v
				}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineRdsPostgresqlBackupPolicyService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineRdsPostgresqlBackupPolicyService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"instance_id":                {TargetField: "InstanceId"},
			"wal_log_space_limit_enable": {TargetField: "WALLogSpaceLimitEnable"},
			"backup_retention_period":    {TargetField: "BackupRetentionPeriod"},
			"data_incr_backup_periods":   {TargetField: "DataIncrBackupPeriods"},
			"full_backup_time":           {TargetField: "FullBackupTime"},
			"full_backup_period":         {TargetField: "FullBackupPeriod"},
			"hourly_incr_backup_enable":  {TargetField: "HourlyIncrBackupEnable"},
			"increment_backup_frequency": {TargetField: "IncrementBackupFrequency"},
		},
		IdField:      "InstanceId",
		CollectField: "backup_policy",
		ResponseConverts: map[string]ve.ResponseConvert{
			"InstanceId":               {TargetField: "instance_id", KeepDefault: true},
			"BackupRetentionPeriod":    {TargetField: "backup_retention_period"},
			"DataIncrBackupPeriods":    {TargetField: "data_incr_backup_periods"},
			"FullBackupTime":           {TargetField: "full_backup_time"},
			"FullBackupPeriod":         {TargetField: "full_backup_period"},
			"HourlyIncrBackupEnable":   {TargetField: "hourly_incr_backup_enable"},
			"IncrementBackupFrequency": {TargetField: "increment_backup_frequency"},
			"WALLogSpaceLimitEnable":   {TargetField: "wal_log_space_limit_enable"},
		},
	}
}

func (s *VolcengineRdsPostgresqlBackupPolicyService) ReadResourceId(id string) string {
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
