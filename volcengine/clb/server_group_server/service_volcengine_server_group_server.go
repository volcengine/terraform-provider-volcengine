package server_group_server

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/clb/clb"
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

func (s *VolcengineServerGroupServerService) ReadResources(condition map[string]interface{}) ([]interface{}, error) {
	var (
		serverIdMap = make(map[string]bool)
		res         = make([]interface{}, 0)
	)
	servers, err := ve.WithSimpleQuery(condition, func(m map[string]interface{}) ([]interface{}, error) {
		var (
			resp    *map[string]interface{}
			err     error
			results interface{}
		)
		clb := s.Client.ClbClient
		action := "DescribeServerGroupAttributes"
		logger.Debug(logger.ReqFormat, action, condition)
		if condition == nil {
			resp, err = clb.DescribeServerGroupAttributesCommon(nil)
			if err != nil {
				return []interface{}{}, err
			}
		} else {
			resp, err = clb.DescribeServerGroupAttributesCommon(&condition)
			if err != nil {
				return []interface{}{}, err
			}
		}

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
		return servers, err
	}

	serverIds := make([]string, 0)
	for k, v := range condition {
		if strings.HasPrefix(k, "ServerIds.") {
			serverIds = append(serverIds, v.(string))
		}
	}

	if len(serverIds) == 0 {
		return servers, nil
	}

	for _, id := range serverIds {
		serverIdMap[strings.Trim(id, " ")] = true
	}

	for _, server := range servers {
		if _, ok := serverIdMap[server.(map[string]interface{})["ServerId"].(string)]; ok {
			res = append(res, server)
		}
	}
	return res, nil
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
		"ServerIds.1":   ids[1],
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
		return data, fmt.Errorf("ServerGroup server %s not exist ", serverGroupServerId)
	}
	return data, err
}

func (s *VolcengineServerGroupServerService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (*VolcengineServerGroupServerService) WithResourceResponseHandlers(serverGroupServer map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return serverGroupServer, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}

}

func (s *VolcengineServerGroupServerService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	clbId, err := s.queryLoadBalancerId(resourceData.Get("server_group_id").(string))
	if err != nil {
		return []ve.Callback{{
			Err: err,
		}}
	}

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action: "AddServerGroupBackendServers",
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["ServerGroupId"] = d.Get("server_group_id")
				(*call.SdkParam)["Servers.1.InstanceId"] = d.Get("instance_id")
				(*call.SdkParam)["Servers.1.Type"] = d.Get("type")
				(*call.SdkParam)["Servers.1.Weight"] = d.Get("weight")
				(*call.SdkParam)["Servers.1.Port"] = d.Get("port")
				(*call.SdkParam)["Servers.1.Description"] = d.Get("description")

				ip := d.Get("ip").(string)
				if ip == "" {
					// query ecs primary private ip
					ip, err = s.getEcsPrimaryPrivateIp(d.Get("instance_id").(string))
					if err != nil {
						return false, err
					}
				}
				(*call.SdkParam)["Servers.1.Ip"] = ip

				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				// 创建 server group server
				return s.Client.ClbClient.AddServerGroupBackendServersCommon(call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				// 注意 获取内容 这个地方不能是指针 需要转一次
				id, _ := ve.ObtainSdkValue("Result.ServerIds.0", *resp)
				d.SetId(fmt.Sprintf("%s:%s", (*call.SdkParam)["ServerGroupId"], id.(string)))
				return nil
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				clb.NewClbService(s.Client): {
					Target:     []string{"Active", "Inactive"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					ResourceId: clbId,
				},
			},
			LockId: func(d *schema.ResourceData) string {
				return clbId
			},
		},
	}
	return []ve.Callback{callback}

}

func (s *VolcengineServerGroupServerService) getEcsPrimaryPrivateIp(instanceId string) (string, error) {
	var (
		err     error
		results interface{}
		ok      bool
		data    []interface{}
	)

	resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo("DescribeInstances"), &map[string]interface{}{
		"InstanceIds.1": instanceId,
	})
	if err != nil {
		return "", err
	}

	results, err = ve.ObtainSdkValue("Result.Instances", *resp)
	if err != nil {
		return "", err
	}
	if results == nil {
		results = []interface{}{}
	}
	if data, ok = results.([]interface{}); !ok {
		return "", errors.New("Result.Instances is not Slice")
	}

	if len(data) == 0 {
		return "", fmt.Errorf("Instance %s not exist ", instanceId)
	}

	interfaces, err := ve.ObtainSdkValue("NetworkInterfaces", data[0])
	if err != nil {
		return "", err
	}

	primaryPrivateIp := ""
	if networkInterfaces, ok := interfaces.([]interface{}); !ok {
		return "", errors.New("NetworkInterfaces is not Slice")
	} else {
		for _, v := range networkInterfaces {
			vv := v.(map[string]interface{})
			if vv["Type"].(string) == "primary" {
				primaryPrivateIp = vv["PrimaryIpAddress"].(string)
			}
		}
	}

	if primaryPrivateIp == "" {
		return "", fmt.Errorf("Instance %s primary ip not exist ", instanceId)
	}

	return primaryPrivateIp, nil
}

func (s *VolcengineServerGroupServerService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	ids := strings.Split(resourceData.Id(), ":")

	clbId, err := s.queryLoadBalancerId(ids[0])
	if err != nil {
		return []ve.Callback{{
			Err: err,
		}}
	}

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action: "ModifyServerGroupAttributes",
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["ServerGroupId"] = ids[0]
				(*call.SdkParam)["Servers.1.ServerId"] = ids[1]
				(*call.SdkParam)["Servers.1.Weight"] = d.Get("weight")
				(*call.SdkParam)["Servers.1.Port"] = d.Get("port")
				(*call.SdkParam)["Servers.1.Description"] = d.Get("description")
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				// 修改 server group server 属性
				return s.Client.ClbClient.ModifyServerGroupAttributesCommon(call.SdkParam)
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				clb.NewClbService(s.Client): {
					Target:     []string{"Active", "Inactive"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					ResourceId: clbId,
				},
			},
			LockId: func(d *schema.ResourceData) string {
				return clbId
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineServerGroupServerService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	ids := strings.Split(resourceData.Id(), ":")

	clbId, err := s.queryLoadBalancerId(ids[0])
	if err != nil {
		return []ve.Callback{{
			Err: err,
		}}
	}

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
				return s.Client.ClbClient.RemoveServerGroupBackendServersCommon(call.SdkParam)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on  reading server group server on delete %q, %w", d.Id(), callErr))
						}
					}
					_, callErr = call.ExecuteCall(d, client, call)
					if callErr == nil {
						return nil
					}
					return resource.RetryableError(callErr)
				})
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				clb.NewClbService(s.Client): {
					Target:     []string{"Active", "Inactive"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					ResourceId: clbId,
				},
			},
			LockId: func(d *schema.ResourceData) string {
				return clbId
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineServerGroupServerService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "ServerIds",
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

func (s *VolcengineServerGroupServerService) queryLoadBalancerId(serverGroupId string) (string, error) {
	if serverGroupId == "" {
		return "", fmt.Errorf("server_group_id cannot be empty")
	}

	// 查询 LoadBalancerId
	serverGroupResp, err := s.Client.ClbClient.DescribeServerGroupAttributesCommon(&map[string]interface{}{
		"ServerGroupId": serverGroupId,
	})
	if err != nil {
		return "", err
	}
	clbId, err := ve.ObtainSdkValue("Result.LoadBalancerId", *serverGroupResp)
	if err != nil {
		return "", err
	}
	return clbId.(string), nil
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "ecs",
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		Action:      actionName,
	}
}
