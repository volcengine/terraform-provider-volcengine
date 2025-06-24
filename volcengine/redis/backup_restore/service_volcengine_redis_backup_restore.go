package backup_restore

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/redis/instance"
)

type VolcengineRedisBackupRestoreService struct {
	Client *ve.SdkClient
}

const (
	ActionRestoreDBInstance = "RestoreDBInstance"
)

func NewRedisBackupRestoreService(c *ve.SdkClient) *VolcengineRedisBackupRestoreService {
	return &VolcengineRedisBackupRestoreService{
		Client: c,
	}
}

func (s *VolcengineRedisBackupRestoreService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineRedisBackupRestoreService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	return nil, nil
}

func (s *VolcengineRedisBackupRestoreService) ReadResource(resourceData *schema.ResourceData, tmpId string) (data map[string]interface{}, err error) {
	var (
		ids []string
	)
	if tmpId == "" {
		tmpId = s.ReadResourceId(resourceData.Id())
	}
	// 无法读取出来
	ids = strings.Split(tmpId, ":")
	if len(ids) != 2 {
		return data, fmt.Errorf("invalid resource id")
	}
	return data, nil
}

func (s *VolcengineRedisBackupRestoreService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Delay:      time.Second,
		Pending:    []string{},
		Target:     target,
		Timeout:    timeout,
		MinTimeout: time.Second,
		Refresh:    nil,
	}
}

func (s *VolcengineRedisBackupRestoreService) WithResourceResponseHandlers(m map[string]interface{}) []ve.ResourceResponseHandler {
	return []ve.ResourceResponseHandler{}
}

func (s *VolcengineRedisBackupRestoreService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      ActionRestoreDBInstance,
			ContentType: ve.ContentTypeJson,
			Convert:     map[string]ve.RequestConvert{},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, *call.SdkParam)
				output, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, *call.SdkParam, *output)
				return output, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				d.SetId(fmt.Sprintf("restore:%s", d.Get("instance_id")))
				return nil
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("instance_id").(string)
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				instance.NewRedisDbInstanceService(s.Client): {
					Target:     []string{"Running"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					ResourceId: resourceData.Get("instance_id").(string),
				},
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineRedisBackupRestoreService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) (callbacks []ve.Callback) {
	if resourceData.HasChanges("time_point", "backup_point_id") {
		callback := ve.Callback{
			Call: ve.SdkCall{
				Action:      ActionRestoreDBInstance,
				ContentType: ve.ContentTypeJson,
				ConvertMode: ve.RequestConvertInConvert,
				Convert: map[string]ve.RequestConvert{
					"instance_id": {
						ForceGet: true,
					},
					"backup_type": {
						ForceGet: true,
					},
					"time_point": {
						TargetField: "TimePoint",
					},
					"backup_point_id": {
						TargetField: "BackupPointId",
					},
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, *call.SdkParam)
					output, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
					logger.Debug(logger.RespFormat, call.Action, *call.SdkParam, *output)
					return output, err
				},
				LockId: func(d *schema.ResourceData) string {
					return d.Get("instance_id").(string)
				},
				ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
					instance.NewRedisDbInstanceService(s.Client): {
						Target:     []string{"Running"},
						Timeout:    resourceData.Timeout(schema.TimeoutCreate),
						ResourceId: resourceData.Get("instance_id").(string),
					},
				},
			},
		}
		callbacks = append(callbacks, callback)
	}
	return callbacks
}

func (s *VolcengineRedisBackupRestoreService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineRedisBackupRestoreService) DatasourceResources(data *schema.ResourceData, resource2 *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		ContentType: ve.ContentTypeJson,
	}
}

func (s *VolcengineRedisBackupRestoreService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "Redis",
		Version:     "2020-12-07",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}
