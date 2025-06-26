package rds_mysql_backup_policy

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

type VolcengineRdsMysqlBackupPolicyService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewRdsMysqlBackupPolicyService(c *ve.SdkClient) *VolcengineRdsMysqlBackupPolicyService {
	return &VolcengineRdsMysqlBackupPolicyService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineRdsMysqlBackupPolicyService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineRdsMysqlBackupPolicyService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
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
			results = []interface{}{}
		}
		action = "DescribeCrossBackupPolicy"
		req := map[string]interface{}{
			"InstanceId": condition["InstanceId"],
		}
		resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &req)
		if err != nil {
			return data, err
		}
		respBytes, _ = json.Marshal(resp)
		logger.Debug(logger.RespFormat, action, condition, string(respBytes))
		crossPolicy, err := ve.ObtainSdkValue("Result", *resp)
		if err != nil {
			return data, err
		}
		if crossPolicy != nil {
			if _, ok := results.(map[string]interface{}); ok {
				results.(map[string]interface{})["CrossBackupPolicy"] = crossPolicy
			}
		}
		results = []interface{}{results}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.BackupPolicy is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineRdsMysqlBackupPolicyService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	ids := strings.Split(id, ":")
	req := map[string]interface{}{
		"InstanceId": ids[0],
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
		return data, fmt.Errorf("rds_mysql_backup_policy %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineRdsMysqlBackupPolicyService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
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
			status, err = ve.ObtainSdkValue("Status", d)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("rds_mysql_backup_policy status error, status: %s", status.(string))
				}
			}
			return d, status.(string), err
		},
	}
}

func (s *VolcengineRdsMysqlBackupPolicyService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyBackupPolicy",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"data_full_backup_periods": {
					TargetField: "DataFullBackupPeriods",
					ConvertType: ve.ConvertJsonArray,
				},
				"lock_ddl_time": {
					TargetField: "LockDDLTime",
				},
				"data_full_backup_start_utc_hour": {
					TargetField: "DataFullBackupStartUTCHour",
				},
				"data_incr_backup_periods": {
					TargetField: "DataIncrBackupPeriods",
					ConvertType: ve.ConvertJsonArray,
				},
				"cross_backup_policy": {
					Ignore: true,
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				instanceId := d.Get("instance_id").(string)
				d.SetId(instanceId + ":" + "backupPolicy")
				return nil
			},
		},
	}
	callbacks = append(callbacks, callback)
	if _, ok := resourceData.GetOk("cross_backup_policy"); ok {
		crossCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyCrossBackupPolicy",
				ConvertMode: ve.RequestConvertIgnore,
				ContentType: ve.ContentTypeJson,
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["InstanceId"] = d.Get("instance_id")
					if backupEnabled, ok := d.GetOk("cross_backup_policy.0.backup_enabled"); ok {
						(*call.SdkParam)["BackupEnabled"] = backupEnabled
					}
					if backupRegion, ok := d.GetOk("cross_backup_policy.0.cross_backup_region"); ok {
						(*call.SdkParam)["CrossBackupRegion"] = backupRegion
					}
					if logBackupEnabled, ok := d.GetOk("cross_backup_policy.0.log_backup_enabled"); ok {
						(*call.SdkParam)["LogBackupEnabled"] = logBackupEnabled
					}
					if retention, ok := d.GetOk("cross_backup_policy.0.retention"); ok {
						(*call.SdkParam)["Retention"] = retention
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
		callbacks = append(callbacks, crossCallback)
	}
	return callbacks
}

func (VolcengineRdsMysqlBackupPolicyService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"LockDDLTime": {
				TargetField: "lock_ddl_time",
			},
			"DataFullBackupStartUTCHour": {
				TargetField: "data_full_backup_start_utc_hour",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineRdsMysqlBackupPolicyService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyBackupPolicy",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"data_full_backup_periods": {
					TargetField: "DataFullBackupPeriods",
					ConvertType: ve.ConvertJsonArray,
				},
				"data_backup_retention_day": {
					TargetField: "DataBackupRetentionDay",
				},
				"data_full_backup_time": {
					TargetField: "DataFullBackupTime",
				},
				"data_incr_backup_periods": {
					TargetField: "DataIncrBackupPeriods",
					ConvertType: ve.ConvertJsonArray,
				},
				"binlog_file_counts_enable": {
					TargetField: "BinlogFileCountsEnable",
				},
				"binlog_limit_count": {
					TargetField: "BinlogLimitCount",
				},
				"binlog_local_retention_hour": {
					TargetField: "BinlogLocalRetentionHour",
				},
				"binlog_space_limit_enable": {
					TargetField: "BinlogSpaceLimitEnable",
				},
				"binlog_storage_percentage": {
					TargetField: "BinlogStoragePercentage",
				},
				"log_backup_retention_day": {
					TargetField: "LogBackupRetentionDay",
				},
				"log_ddl_time": {
					TargetField: "LogDDLTime",
				},
				"data_full_backup_start_utc_hour": {
					TargetField: "DataFullBackupStartUTCHour",
				},
				"hourly_incr_backup_enable": {
					TargetField: "HourlyIncrBackupEnable",
				},
				"incr_backup_hour_period": {
					TargetField: "IncrBackupHourPeriod",
				},
				"data_backup_encryption_enabled": {
					TargetField: "DataBackupEncryptionEnabled",
				},
				"binlog_backup_encryption_enabled": {
					TargetField: "BinlogBackupEncryptionEnabled",
				},
				"data_keep_policy_after_released": {
					TargetField: "DataKeepPolicyAfterReleased",
				},
				"data_keep_days_after_released": {
					TargetField: "DataKeepDaysAfterReleased",
				},
				"data_backup_all_retention": {
					TargetField: "DataBackupAllRetention",
				},
				"binlog_backup_all_retention": {
					TargetField: "BinlogBackupAllRetention",
				},
				"binlog_backup_enabled": {
					TargetField: "BinlogBackupEnabled",
				},
				"retention_policy_synced": {
					TargetField: "RetentionPolicySynced",
				},
				"cross_backup_policy": {
					Ignore: true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["InstanceId"] = d.Get("instance_id").(string)
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
	callbacks = append(callbacks, callback)
	if resourceData.HasChange("cross_backup_policy") {
		crossCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyCrossBackupPolicy",
				ConvertMode: ve.RequestConvertIgnore,
				ContentType: ve.ContentTypeJson,
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["InstanceId"] = d.Get("instance_id")
					if backupEnabled, ok := d.GetOk("cross_backup_policy.0.backup_enabled"); ok {
						(*call.SdkParam)["BackupEnabled"] = backupEnabled
					}
					if backupRegion, ok := d.GetOk("cross_backup_policy.0.cross_backup_region"); ok {
						(*call.SdkParam)["CrossBackupRegion"] = backupRegion
					}
					if logBackupEnabled, ok := d.GetOk("cross_backup_policy.0.log_backup_enabled"); ok {
						(*call.SdkParam)["LogBackupEnabled"] = logBackupEnabled
					}
					if retention, ok := d.GetOk("cross_backup_policy.0.retention"); ok {
						(*call.SdkParam)["Retention"] = retention
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
		callbacks = append(callbacks, crossCallback)
	}
	return callbacks
}

func (s *VolcengineRdsMysqlBackupPolicyService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineRdsMysqlBackupPolicyService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineRdsMysqlBackupPolicyService) ReadResourceId(id string) string {
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
