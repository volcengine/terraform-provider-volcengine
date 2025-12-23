package rds_postgresql_restore_backup

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineRdsPostgresqlRestoreBackupService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewRdsPostgresqlRestoreBackupService(c *ve.SdkClient) *VolcengineRdsPostgresqlRestoreBackupService {
	return &VolcengineRdsPostgresqlRestoreBackupService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineRdsPostgresqlRestoreBackupService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineRdsPostgresqlRestoreBackupService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	return nil, nil
}

func (s *VolcengineRdsPostgresqlRestoreBackupService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {

	return nil, nil
}

func (s *VolcengineRdsPostgresqlRestoreBackupService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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

func (s *VolcengineRdsPostgresqlRestoreBackupService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "RestoreToExistedInstance",
			ConvertMode: ve.RequestConvertAll,
			Convert:     map[string]ve.RequestConvert{},
			SdkParam:    &map[string]interface{}{},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["BackupId"] = d.Get("backup_id").(string)
				(*call.SdkParam)["SourceDBInstanceId"] = d.Get("source_db_instance_id").(string)
				(*call.SdkParam)["TargetDBInstanceId"] = d.Get("target_db_instance_id").(string)
				(*call.SdkParam)["TargetDBInstanceAccount"] = d.Get("target_db_instance_account").(string)
				raw := d.Get("databases").([]interface{})
				if len(raw) > 50 {
					return false, fmt.Errorf("databases count must be <= 50")
				}
				list := make([]interface{}, 0, len(raw))
				for _, v := range raw {
					m := v.(map[string]interface{})
					item := map[string]interface{}{
						"DBName":    m["db_name"],
						"NewDBName": m["new_db_name"],
					}
					list = append(list, item)
				}
				(*call.SdkParam)["Databases"] = list

				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				d.SetId(call.Action)
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineRdsPostgresqlRestoreBackupService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineRdsPostgresqlRestoreBackupService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{}
	return []ve.Callback{callback}
}

func (s *VolcengineRdsPostgresqlRestoreBackupService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{}
	return []ve.Callback{callback}
}

func (s *VolcengineRdsPostgresqlRestoreBackupService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineRdsPostgresqlRestoreBackupService) ReadResourceId(id string) string {
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
