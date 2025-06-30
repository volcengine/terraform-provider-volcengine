package rds_mysql_backup

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

type VolcengineRdsMysqlBackupService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewRdsMysqlBackupService(c *ve.SdkClient) *VolcengineRdsMysqlBackupService {
	return &VolcengineRdsMysqlBackupService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineRdsMysqlBackupService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineRdsMysqlBackupService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
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
		for _, v := range data {
			var (
				itemErr error
			)
			instanceId, ok := condition["InstanceId"]
			if !ok {
				continue
			}
			req := map[string]interface{}{
				"InstanceId": instanceId,
			}
			action = "DescribeBackupStats"
			bytes, _ = json.Marshal(req)
			logger.Debug(logger.ReqFormat, action, string(bytes))
			resp, itemErr = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &req)
			if itemErr != nil {
				continue
			}
			respBytes, _ = json.Marshal(resp)
			logger.Debug(logger.RespFormat, action, req, string(respBytes))
			results, itemErr = ve.ObtainSdkValue("Result.UsageStats", *resp)
			if itemErr != nil {
				continue
			}
			backupMap, ok := v.(map[string]interface{})
			if !ok {
				return data, errors.New("Value is not map ")
			}
			backupMap["UsageStats"] = results
			backupId := backupMap["BackupId"]
			req = map[string]interface{}{
				"InstanceId": instanceId,
				"BackupId":   backupId,
			}
			action = "DescribeBackupDecryptionKey"
			bytes, _ = json.Marshal(req)
			logger.Debug(logger.ReqFormat, action, string(bytes))
			resp, itemErr = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &req)
			if itemErr != nil {
				continue
			}
			respBytes, _ = json.Marshal(resp)
			logger.Debug(logger.RespFormat, action, req, string(respBytes))
			results, itemErr = ve.ObtainSdkValue("Result.DecryptionKey", *resp)
			if itemErr != nil {
				continue
			}
			backupMap["DecryptionKey"] = results
			results, itemErr = ve.ObtainSdkValue("Result.Iv", *resp)
			if itemErr != nil {
				continue
			}
			backupMap["Iv"] = results
		}
		return data, err
	})
}

func (s *VolcengineRdsMysqlBackupService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
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
		"BackupId":   ids[1],
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
		return data, fmt.Errorf("rds_mysql_backup %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineRdsMysqlBackupService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				demo       map[string]interface{}
				status     interface{}
				failStates []string
			)
			failStates = append(failStates, "Error", "Failed")

			// 可能查询不到
			if err = resource.Retry(20*time.Minute, func() *resource.RetryError {
				demo, err = s.ReadResource(resourceData, id)
				if err != nil {
					if ve.ResourceNotFoundError(err) {
						return resource.RetryableError(err)
					} else {
						return resource.NonRetryableError(err)
					}
				}
				return nil
			}); err != nil {
				return nil, "", err
			}

			demo, err = s.ReadResource(resourceData, id)
			if err != nil {
				return nil, "", err
			}
			status, err = ve.ObtainSdkValue("BackupStatus", demo)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("rds_mysql_backup status error, status: %s", status.(string))
				}
			}
			return demo, status.(string), err
		},
	}
}

func (s *VolcengineRdsMysqlBackupService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateBackup",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"backup_meta": {
					Ignore: true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if meta, ok := d.GetOk("backup_meta"); ok {
					metaList := make([]interface{}, 0)
					metaSet := meta.(*schema.Set).List()
					for _, m := range metaSet {
						metaMap := make(map[string]interface{})
						oriMap := m.(map[string]interface{})
						metaMap["DBName"] = oriMap["db_name"]
						if tableNames, ok := oriMap["table_names"]; ok {
							metaMap["TableNames"] = tableNames.(*schema.Set).List()
						}
						metaList = append(metaList, metaMap)
					}
					(*call.SdkParam)["BackupMeta"] = metaList
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
				instanceId := (*call.SdkParam)["InstanceId"]
				backupId, _ := ve.ObtainSdkValue("Result.BackupId", *resp)
				d.SetId(fmt.Sprintf("%s:%s", instanceId, backupId))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Success"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("instance_id").(string)
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineRdsMysqlBackupService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineRdsMysqlBackupService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{}
	return []ve.Callback{callback}
}

func (s *VolcengineRdsMysqlBackupService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	ids := strings.Split(resourceData.Id(), ":")
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteDataBackup",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"InstanceId": ids[0],
				"BackupId":   ids[1],
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

func (s *VolcengineRdsMysqlBackupService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{},
		IdField:         "BackupId",
		CollectField:    "backups",
		ResponseConverts: map[string]ve.ResponseConvert{
			"BackupId": {
				TargetField: "id",
				KeepDefault: true,
			},
			"DBTableInfos": {
				TargetField: "db_table_infos",
			},
		},
	}
}

func (s *VolcengineRdsMysqlBackupService) ReadResourceId(id string) string {
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
