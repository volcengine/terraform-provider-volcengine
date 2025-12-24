package rds_postgresql_data_backup

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

type VolcengineRdsPostgresqlDataBackupService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewRdsPostgresqlDataBackupService(c *ve.SdkClient) *VolcengineRdsPostgresqlDataBackupService {
	return &VolcengineRdsPostgresqlDataBackupService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineRdsPostgresqlDataBackupService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineRdsPostgresqlDataBackupService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeBackups"

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
		results, err = ve.ObtainSdkValue("Result.Backups", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.Backups is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineRdsPostgresqlDataBackupService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	// 资源 ID = InstanceId:BackupId
	ids := strings.Split(id, ":")
	// 为兼容物理备份返回 BackupId=null 的情况，查询不带 BackupId 过滤，靠本地匹配确定目标
	req := map[string]interface{}{
		"InstanceId": ids[0],
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
		bid, _ := m["BackupId"].(string)
		bfn, _ := m["BackupFileName"].(string)
		if bid == target || bfn == target {
			data = m
			break
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("rds_postgresql_data_backup %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineRdsPostgresqlDataBackupService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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
			status, err = ve.ObtainSdkValue("BackupStatus", d)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("rds_postgresql_data_backup status error, status: %s", status.(string))
				}
			}
			return d, status.(string), err
		},
	}
}

func (s *VolcengineRdsPostgresqlDataBackupService) CreateResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateBackup",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"instance_id":        {TargetField: "InstanceId"},
				"backup_meta":        {Ignore: true},
				"backup_scope":       {TargetField: "BackupScope"},
				"backup_method":      {TargetField: "BackupMethod"},
				"backup_type":        {TargetField: "BackupType"},
				"backup_description": {TargetField: "BackupDescription"},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if scope, ok := d.GetOk("backup_scope"); ok {
					(*call.SdkParam)["BackupScope"] = scope
				}
				if method, ok := d.GetOk("backup_method"); ok {
					(*call.SdkParam)["BackupMethod"] = method
				}
				if btype, ok := d.GetOk("backup_type"); ok {
					(*call.SdkParam)["BackupType"] = btype
				}
				if desc, ok := d.GetOk("backup_description"); ok {
					(*call.SdkParam)["BackupDescription"] = desc
				}
				if metaVal, ok := d.GetOk("backup_meta"); ok {
					list := make([]interface{}, 0)
					for _, m := range metaVal.([]interface{}) {
						mm := m.(map[string]interface{})
						item := map[string]interface{}{
							"DBName": mm["db_name"],
						}
						list = append(list, item)
					}
					(*call.SdkParam)["BackupMeta"] = list
				}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				instanceId := (*call.SdkParam)["InstanceId"].(string)
				// 在 BackupMethod 的取值为 Logical 时有该返回值。当 BackupMethod 的取值为 Physical 时，返回为 null。
				backupId, _ := ve.ObtainSdkValue("Result.BackupId", *resp)
				if backupIdStr, ok := backupId.(string); ok && backupIdStr != "" {
					d.SetId(fmt.Sprintf("%s:%s", instanceId, backupIdStr))
					return nil
				}
				// fallback：等待 DescribeBackups 可见后按时间窗口+用户创建+描述匹配选最新一条
				method, _ := (*call.SdkParam)["BackupMethod"].(string)
				scope, _ := (*call.SdkParam)["BackupScope"].(string)
				desc := ""
				if v, ok := d.GetOk("backup_description"); ok {
					desc = v.(string)
				}
				start := time.Now().UTC().Add(-10 * time.Minute).Format("2006-01-02T15:04:05.000Z")
				end := time.Now().UTC().Add(10 * time.Minute).Format("2006-01-02T15:04:05.000Z")
				var foundId string
				var foundFile string
				deadline := time.Now().Add(5 * time.Minute)
				for {
					req := map[string]interface{}{
						"InstanceId":      instanceId,
						"BackupStartTime": start,
						"BackupEndTime":   end,
						"CreateType":      "User",
					}
					if method != "" {
						req["BackupMethod"] = method
					}
					if scope != "" {
						req["BackupScope"] = scope
					}
					if desc != "" {
						req["BackupDescription"] = desc
					}
					list, err := s.ReadResources(req)
					if err == nil {
						var latestStart string
						for _, it := range list {
							m, _ := it.(map[string]interface{})
							if m == nil {
								continue
							}
							bs, _ := m["BackupStartTime"].(string)
							if latestStart == "" || bs > latestStart {
								latestStart = bs
								if bid, _ := m["BackupId"].(string); bid != "" {
									foundId = bid
								} else if bfn, _ := m["BackupFileName"].(string); bfn != "" {
									foundFile = bfn
								}
							}
						}
						if foundId != "" || foundFile != "" {
							break
						}
					}
					if time.Now().After(deadline) {
						break
					}
					time.Sleep(10 * time.Second)
				}
				if foundId == "" && foundFile == "" {
					return fmt.Errorf("create backup succeeded but BackupId missing in response and fallback query did not find a match")
				}
				if foundId != "" {
					d.SetId(fmt.Sprintf("%s:%s", instanceId, foundId))
					return nil
				}
				d.SetId(fmt.Sprintf("%s:%s", instanceId, foundFile))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Success"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
			LockId: func(d *schema.ResourceData) string { return d.Get("instance_id").(string) },
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineRdsPostgresqlDataBackupService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"BackupId":          {TargetField: "backup_id"},
			"BackupMethod":      {TargetField: "backup_method"},
			"BackupType":        {TargetField: "backup_type"},
			"BackupStatus":      {TargetField: "backup_status"},
			"BackupStartTime":   {TargetField: "backup_start_time"},
			"BackupEndTime":     {TargetField: "backup_end_time"},
			"BackupDescription": {TargetField: "backup_description"},
			"DownloadStatus":    {TargetField: "download_status"},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineRdsPostgresqlDataBackupService) ModifyResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineRdsPostgresqlDataBackupService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteBackup",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			SdkParam:    &map[string]interface{}{},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				// 当前仅支持删除逻辑备份
				if d.Get("backup_method").(string) == "Physical" {
					return false, fmt.Errorf("only logical backups can be deleted")
				}
				(*call.SdkParam)["InstanceId"] = d.Get("instance_id").(string)
				bid := d.Get("backup_id").(string)
				if bid == "" {
					ids := strings.Split(d.Id(), ":")
					req := map[string]interface{}{
						"InstanceId": ids[0],
					}
					list, err := s.ReadResources(req)
					if err == nil {
						for _, it := range list {
							m, _ := it.(map[string]interface{})
							if m == nil {
								continue
							}
							bfn, _ := m["BackupFileName"].(string)
							if bfn == ids[1] {
								bid, _ = m["BackupId"].(string)
								break
							}
						}
					}
					if bid == "" {
						bid = ids[1]
					}
				}
				(*call.SdkParam)["BackupId"] = bid
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 5*time.Minute)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineRdsPostgresqlDataBackupService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"instance_id":          {TargetField: "InstanceId"},
			"backup_id":            {TargetField: "BackupId"},
			"backup_start_time":    {TargetField: "BackupStartTime"},
			"backup_end_time":      {TargetField: "BackupEndTime"},
			"backup_status":        {TargetField: "BackupStatus"},
			"backup_type":          {TargetField: "BackupType"},
			"backup_method":        {TargetField: "BackupMethod"},
			"backup_database_name": {TargetField: "BackupDatabaseName"},
		},
		IdField:      "BackupId",
		CollectField: "backups",
		ResponseConverts: map[string]ve.ResponseConvert{
			"BackupId": {TargetField: "backup_id", KeepDefault: true},
		},
	}
}

func (s *VolcengineRdsPostgresqlDataBackupService) ReadResourceId(id string) string {
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
