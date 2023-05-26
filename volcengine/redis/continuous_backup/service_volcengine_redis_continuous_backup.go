package continuous_backup

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/redis/instance"
)

type VolcengineRedisContinuousBackupService struct {
	Client *ve.SdkClient
}

func NewRedisContinuousBackupService(c *ve.SdkClient) *VolcengineRedisContinuousBackupService {
	return &VolcengineRedisContinuousBackupService{
		Client: c,
	}
}

func (s *VolcengineRedisContinuousBackupService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineRedisContinuousBackupService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	return nil, nil
}

func (s *VolcengineRedisContinuousBackupService) ReadResource(resourceData *schema.ResourceData, tmpId string) (data map[string]interface{}, err error) {
	var (
		ids []string
	)
	if tmpId == "" {
		tmpId = s.ReadResourceId(resourceData.Id())
	}
	// 无法读取出来
	ids = strings.Split(tmpId, ":")
	if len(ids) != 2 {
		return data, fmt.Errorf("invalid continuous backup id")
	}
	return data, nil
}

func (s *VolcengineRedisContinuousBackupService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Delay:      time.Second,
		Pending:    []string{},
		Target:     target,
		Timeout:    timeout,
		MinTimeout: time.Second,
		Refresh:    nil,
	}
}

func (s *VolcengineRedisContinuousBackupService) WithResourceResponseHandlers(m map[string]interface{}) []ve.ResourceResponseHandler {
	return []ve.ResourceResponseHandler{}
}

func (s *VolcengineRedisContinuousBackupService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "StartContinuousBackup",
			ContentType: ve.ContentTypeJson,
			ConvertMode: ve.RequestConvertAll,
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				(*call.SdkParam)["ClientToken"] = uuid.New().String()
				logger.Debug(logger.ReqFormat, call.Action, *call.SdkParam)
				output, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, *call.SdkParam, *output)
				return output, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				d.SetId(fmt.Sprintf("continuous:%s", d.Get("instance_id")))
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

func (s *VolcengineRedisContinuousBackupService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineRedisContinuousBackupService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "StopContinuousBackup",
			ContentType: ve.ContentTypeJson,
			ConvertMode: ve.RequestConvertIgnore,
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				(*call.SdkParam)["InstanceId"] = d.Get("instance_id").(string)
				(*call.SdkParam)["ClientToken"] = uuid.New().String()
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
	return []ve.Callback{callback}
}

func (s *VolcengineRedisContinuousBackupService) DatasourceResources(data *schema.ResourceData, resource2 *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineRedisContinuousBackupService) ReadResourceId(id string) string {
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
