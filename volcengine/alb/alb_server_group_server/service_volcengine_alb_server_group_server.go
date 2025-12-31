package alb_server_group_server

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
	"github.com/volcengine/terraform-provider-volcengine/volcengine/alb/alb_server_group"
)

type VolcengineServerGroupServerService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewServerGroupServerService(c *ve.SdkClient) *VolcengineServerGroupServerService {
	return &VolcengineServerGroupServerService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineServerGroupServerService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineServerGroupServerService) ReadResources(m map[string]interface{}) ([]interface{}, error) {
	var (
		resp    *map[string]interface{}
		err     error
		results interface{}
	)
	servers, err := ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 20, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeServerGroupBackendServers"
		bytes, _ := json.Marshal(condition)
		logger.Debug(logger.ReqFormat, action, string(bytes))
		if condition == nil {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), nil)
			if err != nil {
				return nil, err
			}
		} else {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &condition)
			if err != nil {
				return nil, err
			}
		}
		respBytes, _ := json.Marshal(resp)
		logger.Debug(logger.RespFormat, action, condition, string(respBytes))

		results, err = ve.ObtainSdkValue("Result.Servers", *resp)
		if err != nil {
			return []interface{}{}, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok := results.([]interface{}); !ok {
			return data, errors.New("Result.Servers is not Slice")
		} else {
			return data, err
		}
	})
	if err != nil {
		return nil, err
	}
	return servers, nil
}

func (s *VolcengineServerGroupServerService) ReadResource(resourceData *schema.ResourceData, serverGroupServerId string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if serverGroupServerId == "" {
		serverGroupServerId = resourceData.Id()
	}
	ids := strings.Split(serverGroupServerId, ":")
	req := map[string]interface{}{
		"ServerGroupId": ids[0],
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		if data, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("Value is not map ")
		}
		// 找到对应的 server id
		if v.(map[string]interface{})["ServerId"] == ids[1] {
			return v.(map[string]interface{}), nil
		}
	}
	return data, fmt.Errorf("ServerGroup server %s not exist ", serverGroupServerId)
}

func (s *VolcengineServerGroupServerService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (*VolcengineServerGroupServerService) WithResourceResponseHandlers(serverGroupServer map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return serverGroupServer, map[string]ve.ResponseConvert{
			"ServerId": {
				TargetField: "id",
				KeepDefault: true,
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineServerGroupServerService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "AddServerGroupBackendServers",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["ServerGroupId"] = d.Get("server_group_id")
				(*call.SdkParam)["Servers.1.InstanceId"] = d.Get("instance_id")
				(*call.SdkParam)["Servers.1.Type"] = d.Get("type")
				(*call.SdkParam)["Servers.1.Weight"] = d.Get("weight")
				(*call.SdkParam)["Servers.1.Port"] = d.Get("port")
				(*call.SdkParam)["Servers.1.Description"] = d.Get("description")
				(*call.SdkParam)["Servers.1.Ip"] = d.Get("ip")
				if v, ok := d.GetOk("remote_enabled"); ok {
					(*call.SdkParam)["Servers.1.RemoteEnabled"] = v
				}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("Result.ServerIds.0", *resp)
				d.SetId(fmt.Sprintf("%s:%s", (*call.SdkParam)["ServerGroupId"], id.(string)))
				return nil
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("server_group_id").(string)
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				alb_server_group.NewAlbServerGroupService(s.Client): {
					Target:     []string{"Active"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					ResourceId: resourceData.Get("server_group_id").(string),
				},
			},
		},
	}
	return []ve.Callback{callback}

}

func (s *VolcengineServerGroupServerService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	ids := strings.Split(resourceData.Id(), ":")
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyServerGroupBackendServers",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["ServerGroupId"] = ids[0]
				(*call.SdkParam)["Servers.1.ServerId"] = ids[1]
				(*call.SdkParam)["Servers.1.Weight"] = d.Get("weight")
				(*call.SdkParam)["Servers.1.Port"] = d.Get("port")
				(*call.SdkParam)["Servers.1.Description"] = d.Get("description")
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("server_group_id").(string)
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				alb_server_group.NewAlbServerGroupService(s.Client): {
					Target:     []string{"Active"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					ResourceId: resourceData.Get("server_group_id").(string),
				},
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineServerGroupServerService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	ids := strings.Split(resourceData.Id(), ":")

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "RemoveServerGroupBackendServers",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"ServerGroupId": ids[0],
				"ServerIds.1":   ids[1],
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				//删除 Server Group
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("server_group_id").(string)
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				alb_server_group.NewAlbServerGroupService(s.Client): {
					Target:     []string{"Active"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					ResourceId: resourceData.Get("server_group_id").(string),
				},
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineServerGroupServerService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"instance_ids": {
				TargetField: "InstanceIds",
				ConvertType: ve.ConvertWithN,
			},
			"ips": {
				TargetField: "Ips",
				ConvertType: ve.ConvertWithN,
			},
		},
		NameField:    "ServerId",
		IdField:      "ServerId",
		CollectField: "servers",
		ResponseConverts: map[string]ve.ResponseConvert{
			"ServerId": {
				TargetField: "id",
				KeepDefault: true,
			},
		},
	}
}

func (s *VolcengineServerGroupServerService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "alb",
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		Action:      actionName,
	}
}
