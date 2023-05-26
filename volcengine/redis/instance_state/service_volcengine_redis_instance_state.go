package instance_state

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/redis/instance"
)

type VolcengineRedisInstanceStateService struct {
	Client *ve.SdkClient
}

func (v *VolcengineRedisInstanceStateService) GetClient() *ve.SdkClient {
	return v.Client
}

func (v *VolcengineRedisInstanceStateService) ReadResources(m map[string]interface{}) ([]interface{}, error) {
	return nil, nil
}

func (v *VolcengineRedisInstanceStateService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	if id == "" {
		id = v.ReadResourceId(resourceData.Id())
	}
	ids := strings.Split(id, ":")
	if len(ids) != 2 {
		return data, errors.New("id err")
	}
	data, err = instance.NewRedisDbInstanceService(v.Client).ReadResource(resourceData, ids[1])
	return data, err
}

func (v *VolcengineRedisInstanceStateService) RefreshResourceState(data *schema.ResourceData, strings []string, duration time.Duration, s string) *resource.StateChangeConf {
	return nil
}

func (v *VolcengineRedisInstanceStateService) WithResourceResponseHandlers(m map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return m, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (v *VolcengineRedisInstanceStateService) CreateResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	action := ""
	if data.Get("action").(string) == "Restart" {
		action = "RestartDBInstance"
	}
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      action,
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"action": {
					Ignore: true,
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return v.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				instanceId := d.Get("instance_id").(string)
				logger.Debug(logger.RespFormat, call.Action, instanceId)
				d.SetId(fmt.Sprintf("state:%s", instanceId))
				return nil
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				instance.NewRedisDbInstanceService(v.Client): {
					Target:     []string{"Running"},
					Timeout:    data.Timeout(schema.TimeoutCreate),
					ResourceId: data.Get("instance_id").(string),
				},
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("instance_id").(string)
			},
		},
	}
	return []ve.Callback{callback}
}

func (v *VolcengineRedisInstanceStateService) ModifyResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return nil
}

func (v *VolcengineRedisInstanceStateService) RemoveResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return nil
}

func (v *VolcengineRedisInstanceStateService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (v *VolcengineRedisInstanceStateService) ReadResourceId(s string) string {
	return s
}

func NewRedisInstanceStateService(c *ve.SdkClient) *VolcengineRedisInstanceStateService {
	return &VolcengineRedisInstanceStateService{
		Client: c,
	}
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "Redis",
		Action:      actionName,
		Version:     "2020-12-07",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
	}
}
