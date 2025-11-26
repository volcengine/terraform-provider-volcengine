package rds_postgresql_backup_download

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

type VolcengineRdsPostgresqlBackupDownloadService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewRdsPostgresqlBackupDownloadService(c *ve.SdkClient) *VolcengineRdsPostgresqlBackupDownloadService {
	return &VolcengineRdsPostgresqlBackupDownloadService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineRdsPostgresqlBackupDownloadService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineRdsPostgresqlBackupDownloadService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
	)
	return ve.WithSimpleQuery(m, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "GetBackupDownloadLink"

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

func (s *VolcengineRdsPostgresqlBackupDownloadService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	ids := strings.Split(id, ":")
	if len(ids) < 2 {
		return data, fmt.Errorf("rds_postgresql_backup_download id must be 'instance_id:backup_id'")
	}
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
		return data, fmt.Errorf("rds_postgresql_backup_download %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineRdsPostgresqlBackupDownloadService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      10 * time.Second,
		MinTimeout: 10 * time.Second,
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
			status, err = ve.ObtainSdkValue("PrepareProgess", d)
			if err != nil {
				return nil, "", err
			}
			return d, fmt.Sprintf("%v", status), err
		},
	}
}

func (s *VolcengineRdsPostgresqlBackupDownloadService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DownloadBackup",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"instance_id": {TargetField: "InstanceId"},
				"backup_id":   {TargetField: "BackupId"},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				instanceId := (*call.SdkParam)["InstanceId"].(string)
				backupId := (*call.SdkParam)["BackupId"].(string)
				d.SetId(fmt.Sprintf("%s:%s", instanceId, backupId))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"100"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
			LockId: func(d *schema.ResourceData) string { return d.Get("instance_id").(string) },
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineRdsPostgresqlBackupDownloadService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"BackupDescription":       {TargetField: "backup_description"},
			"BackupDownloadLink":      {TargetField: "backup_download_link"},
			"BackupFileName":          {TargetField: "backup_file_name"},
			"BackupFileSize":          {TargetField: "backup_file_size"},
			"BackupId":                {TargetField: "backup_id"},
			"BackupMethod":            {TargetField: "backup_method"},
			"InnerBackupDownloadLink": {TargetField: "inner_backup_download_link"},
			"InstanceId":              {TargetField: "instance_id"},
			"LinkExpiredTime":         {TargetField: "link_expired_time"},
			"PrepareProgess":          {TargetField: "prepare_progress"},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineRdsPostgresqlBackupDownloadService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{}
	return []ve.Callback{callback}
}

func (s *VolcengineRdsPostgresqlBackupDownloadService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{}
	return []ve.Callback{callback}
}

func (s *VolcengineRdsPostgresqlBackupDownloadService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"instance_id": {TargetField: "InstanceId"},
			"backup_id":   {TargetField: "BackupId"},
		},
		IdField:      "BackupId",
		CollectField: "downloads",
		ResponseConverts: map[string]ve.ResponseConvert{
			"BackupId": {TargetField: "backup_id", KeepDefault: true},
		},
		ContentType: ve.ContentTypeJson,
	}
}

func (s *VolcengineRdsPostgresqlBackupDownloadService) ReadResourceId(id string) string {
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
