package vedb_mysql_backup

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
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vedb_mysql/vedb_mysql_instance"
)

type VolcengineVedbMysqlBackupService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewVedbMysqlBackupService(c *ve.SdkClient) *VolcengineVedbMysqlBackupService {
	return &VolcengineVedbMysqlBackupService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineVedbMysqlBackupService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineVedbMysqlBackupService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
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
		results, err = ve.ObtainSdkValue("Result.BackupsInfo", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.BackupsInfo is not Slice")
		}
		for _, v := range data {
			backup := v.(map[string]interface{})
			instanceId := condition["InstanceId"]
			action = "DescribeBackupPolicy"
			req := map[string]interface{}{
				"InstanceId": instanceId,
			}
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &req)
			if err != nil {
				logger.Info("DescribeBackupPolicy error : ", err)
				continue
			}
			respBytes, _ = json.Marshal(resp)
			logger.Debug(logger.RespFormat, action, req, string(respBytes))

			result, err := ve.ObtainSdkValue("Result", *resp)
			if err != nil {
				logger.Info("ObtainSdkValue Result error:", err)
				continue
			}
			backup["BackupPolicy"] = result
		}
		return data, err
	})
}

func (s *VolcengineVedbMysqlBackupService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		tmpData map[string]interface{}
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
		if tmpData, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("Value is not map ")
		} else if tmpData["BackupId"].(string) == ids[1] {
			data = tmpData
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("vedb_mysql_backup %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineVedbMysqlBackupService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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
			status, err = ve.ObtainSdkValue("BackupStatus", d)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("vedb_mysql_backup status error, status: %s", status.(string))
				}
			}
			return d, status.(string), err
		},
	}
}

func (s *VolcengineVedbMysqlBackupService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateBackup",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"backup_policy": {
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
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				vedb_mysql_instance.NewVedbMysqlInstanceService(s.Client): {
					Target:     []string{"Running"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					ResourceId: resourceData.Get("instance_id").(string),
				},
			},
		},
	}
	callbacks = append(callbacks, callback)
	policy, ok := resourceData.GetOk("backup_policy")
	if ok {
		policyCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyBackupPolicy",
				ConvertMode: ve.RequestConvertIgnore,
				ContentType: ve.ContentTypeJson,
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					policyList, ok := policy.([]interface{})
					if !ok {
						return false, fmt.Errorf("policy is not a list")
					}
					p := policyList[0].(map[string]interface{})
					(*call.SdkParam)["InstanceId"] = d.Get("instance_id").(string)
					(*call.SdkParam)["BackupTime"] = p["backup_time"]
					(*call.SdkParam)["FullBackupPeriod"] = p["full_backup_period"]
					(*call.SdkParam)["BackupRetentionPeriod"] = p["backup_retention_period"]
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
					resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
					logger.Debug(logger.RespFormat, call.Action, resp, err)
					return resp, err
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{"Success"},
					Timeout: resourceData.Timeout(schema.TimeoutCreate),
				},
				LockId: func(d *schema.ResourceData) string {
					return d.Get("instance_id").(string)
				},
				ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
					vedb_mysql_instance.NewVedbMysqlInstanceService(s.Client): {
						Target:     []string{"Running"},
						Timeout:    resourceData.Timeout(schema.TimeoutCreate),
						ResourceId: resourceData.Get("instance_id").(string),
					},
				},
			},
		}
		callbacks = append(callbacks, policyCallback)
	}
	return callbacks
}

func (VolcengineVedbMysqlBackupService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineVedbMysqlBackupService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyBackupPolicy",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				policy, ok := resourceData.GetOk("backup_policy")
				if !ok {
					return false, nil
				}
				policyList, ok := policy.([]interface{})
				if !ok {
					return false, fmt.Errorf("policy is not a list")
				}
				p := policyList[0].(map[string]interface{})
				(*call.SdkParam)["InstanceId"] = d.Get("instance_id").(string)
				(*call.SdkParam)["BackupTime"] = p["backup_time"]
				(*call.SdkParam)["FullBackupPeriod"] = p["full_backup_period"]
				(*call.SdkParam)["BackupRetentionPeriod"] = p["backup_retention_period"]
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Success"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("instance_id").(string)
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				vedb_mysql_instance.NewVedbMysqlInstanceService(s.Client): {
					Target:     []string{"Running"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					ResourceId: resourceData.Get("instance_id").(string),
				},
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineVedbMysqlBackupService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	ids := strings.Split(resourceData.Id(), ":")
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteBackup",
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

func (s *VolcengineVedbMysqlBackupService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		IdField:      "BackupId",
		CollectField: "backups",
		ResponseConverts: map[string]ve.ResponseConvert{
			"BackupId": {
				TargetField: "id",
				KeepDefault: true,
			},
		},
	}
}

func (s *VolcengineVedbMysqlBackupService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "vedbm",
		Version:     "2022-01-01",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}
