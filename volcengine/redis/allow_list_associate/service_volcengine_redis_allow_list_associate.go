package allow_list_associate

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

type VolcengineRedisAllowListAssociateService struct {
	Client *ve.SdkClient
}

const (
	ActionAssociateAllowList    = "AssociateAllowList"
	ActionDisassociateAllowList = "DisassociateAllowList"
)

func NewRedisAllowListAssociateService(c *ve.SdkClient) *VolcengineRedisAllowListAssociateService {
	return &VolcengineRedisAllowListAssociateService{
		Client: c,
	}
}

func (s *VolcengineRedisAllowListAssociateService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineRedisAllowListAssociateService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	return nil, nil
}

func (s *VolcengineRedisAllowListAssociateService) ReadResource(resourceData *schema.ResourceData, tmpId string) (data map[string]interface{}, err error) {
	var (
		ids     []string
		req     map[string]interface{}
		output  *map[string]interface{}
		results interface{}
		ok      bool
	)
	if tmpId == "" {
		tmpId = s.ReadResourceId(resourceData.Id())
	}
	ids = strings.Split(tmpId, ":")
	if len(ids) != 2 {
		return data, fmt.Errorf("invalid id")
	}
	req = map[string]interface{}{
		"AllowListId": ids[1],
	}

	action := "DescribeAllowListDetail"
	logger.Debug(logger.ReqFormat, action, req)
	output, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &req)
	logger.Debug(logger.RespFormat, action, req, *output)

	if err != nil {
		return data, err
	}
	results, err = ve.ObtainSdkValue("Result", *output)
	if err != nil {
		return data, err
	}
	if data, ok = results.(map[string]interface{}); !ok {
		return data, errors.New("value is not map")
	}
	res := map[string]interface{}{
		"InstanceId":  ids[0],
		"AllowListId": ids[1],
	}

	attached := false
	for _, ins := range data["AssociatedInstances"].([]interface{}) {
		if ins.(map[string]interface{})["InstanceId"].(string) == ids[0] {
			attached = true
			break
		}
	}
	if !attached {
		return nil, fmt.Errorf("not associated instance and allow list. %s", tmpId)
	}
	return res, nil
}

func (s *VolcengineRedisAllowListAssociateService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (s *VolcengineRedisAllowListAssociateService) WithResourceResponseHandlers(association map[string]interface{}) []ve.ResourceResponseHandler {
	return []ve.ResourceResponseHandler{}
}

func (s *VolcengineRedisAllowListAssociateService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      ActionAssociateAllowList,
			ContentType: ve.ContentTypeJson,
			ConvertMode: ve.RequestConvertIgnore,
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {

				req := make(map[string]interface{})
				req["InstanceIds"] = []interface{}{d.Get("instance_id")}
				req["AllowListIds"] = []interface{}{d.Get("allow_list_id")}
				output, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), &req)
				logger.Debug(logger.RespFormat, call.Action, *call.SdkParam, *output)
				return output, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				d.SetId(fmt.Sprintf("%s:%s", d.Get("instance_id"), d.Get("allow_list_id")))
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

func (s *VolcengineRedisAllowListAssociateService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineRedisAllowListAssociateService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      ActionDisassociateAllowList,
			ContentType: ve.ContentTypeJson,
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				id := s.ReadResourceId(d.Id())
				ids := strings.Split(id, ":")
				instanceId := ids[0]
				allowListId := ids[1]
				(*call.SdkParam)["InstanceIds"] = []string{instanceId}
				(*call.SdkParam)["AllowListIds"] = []string{allowListId}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, *call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("instance_id").(string)
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				instance.NewRedisDbInstanceService(s.Client): {
					Target:     []string{"Running"},
					Timeout:    resourceData.Timeout(schema.TimeoutDelete),
					ResourceId: resourceData.Get("instance_id").(string),
				},
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineRedisAllowListAssociateService) DatasourceResources(data *schema.ResourceData, resource2 *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineRedisAllowListAssociateService) ReadResourceId(id string) string {
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
