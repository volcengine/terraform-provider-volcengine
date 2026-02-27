package nlb_listener

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineNlbListenerService struct {
	Client *ve.SdkClient
}

func NewNlbListenerService(c *ve.SdkClient) *VolcengineNlbListenerService {
	return &VolcengineNlbListenerService{
		Client: c,
	}
}

func (s *VolcengineNlbListenerService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineNlbListenerService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	return ve.WithNextTokenQuery(m, "MaxResults", "NextToken", 100, nil, func(condition map[string]interface{}) ([]interface{}, string, error) {
		action := "DescribeNLBListeners"
		bytes, _ := json.Marshal(condition)
		logger.Debug(logger.ReqFormat, action, string(bytes))
		resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(action), &condition)
		if err != nil {
			return nil, "", err
		}
		respBytes, _ := json.Marshal(resp)
		logger.Debug(logger.RespFormat, action, condition, string(respBytes))
		results, err := ve.ObtainSdkValue("Result.Listeners", *resp)
		if err != nil {
			return nil, "", err
		}
		if results == nil {
			results = []interface{}{}
		}

		var nextTokenStr string
		nextToken, _ := ve.ObtainSdkValue("Result.NextToken", *resp)
		if nextToken != nil {
			if s, ok := nextToken.(string); ok {
				nextTokenStr = s
			}
		}

		if data, ok := results.([]interface{}); ok {
			return data, nextTokenStr, nil
		}
		return nil, "", errors.New("Result.Listeners is not Slice")
	})
}

func (s *VolcengineNlbListenerService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	if id == "" {
		id = resourceData.Id()
	}
	req := map[string]interface{}{
		"ListenerIds.1": id,
	}
	results, err := s.ReadResources(req)
	if err != nil {
		return nil, err
	}
	for _, v := range results {
		if tempData, ok := v.(map[string]interface{}); ok {
			if lId, ok := tempData["ListenerId"].(string); ok && lId == id {
				return tempData, nil
			}
		}
	}
	return nil, fmt.Errorf("resource not found")
}

func (s *VolcengineNlbListenerService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{"Creating"},
		Target:     target,
		Refresh:    ve.ResourceStateRefreshFunc(resourceData, s.ReadResource, id),
		Timeout:    timeout,
		Delay:      5 * time.Second,
		MinTimeout: 2 * time.Second,
	}
}

func (s *VolcengineNlbListenerService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action: "CreateNLBListener",
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				param, err := ve.ResourceDateToRequest(d, resource, false, s.createRequestConvert(), ve.RequestConvertAll, ve.ContentTypeDefault)
				if err != nil {
					return nil, err
				}
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), &param)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, err := ve.ObtainSdkValue("Result.ListenerId", *resp)
				if err != nil {
					return err
				}
				if s, ok := id.(string); ok && s != "" {
					d.SetId(s)
				} else {
					return errors.New("Result.ListenerId is not string")
				}
				_, err = s.RefreshResourceState(d, []string{"Active"}, 10*time.Minute, d.Id()).WaitForState()
				return err
			},
		},
	}
	callbacks := []ve.Callback{callback}
	setResourceTagsCallbacks := ve.SetResourceTags(s.Client, "TagNLBResources", "UntagNLBResources", "nlb_listener", resourceData, getUniversalInfo)
	callbacks = append(callbacks, setResourceTagsCallbacks...)
	return callbacks
}

func (s *VolcengineNlbListenerService) createRequestConvert() map[string]ve.RequestConvert {
	return map[string]ve.RequestConvert{
		"load_balancer_id": {
			TargetField: "LoadBalancerId",
		},
		"protocol": {
			TargetField: "Protocol",
		},
		"port": {
			TargetField: "Port",
			ForceGet:    true,
		},
		"server_group_id": {
			TargetField: "ServerGroupId",
		},
		"listener_name": {
			TargetField: "ListenerName",
		},
		"description": {
			TargetField: "Description",
		},
		"connection_timeout": {
			TargetField: "ConnectionTimeout",
		},
		"enabled": {
			TargetField: "Enabled",
		},
		"start_port": {
			TargetField: "StartPort",
		},
		"end_port": {
			TargetField: "EndPort",
		},
		"tags": {
			TargetField: "Tags",
			Convert:     transTagsToRequest,
		},
	}
}

func (s *VolcengineNlbListenerService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	// 1. Modify attributes
	if resourceData.HasChanges("listener_name", "description", "server_group_id", "connection_timeout", "enabled") {
		callbacks = append(callbacks, ve.Callback{
			Call: ve.SdkCall{
				Action: "ModifyNLBListenerAttributes",
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					param, err := ve.ResourceDateToRequest(d, resource, true, s.createRequestConvert(), ve.RequestConvertInConvert, ve.ContentTypeDefault)
					if err != nil {
						return nil, err
					}
					param["ListenerId"] = d.Id()
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), &param)
				},
			},
		})
	}

	setResourceTagsCallbacks := ve.SetResourceTags(s.Client, "TagNLBResources", "UntagNLBResources", "nlb_listener", resourceData, getUniversalInfo)
	callbacks = append(callbacks, setResourceTagsCallbacks...)

	return callbacks
}

func (s *VolcengineNlbListenerService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteNLBListener",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"ListenerId": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 10*time.Minute)
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineNlbListenerService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"ListenerId": {
				TargetField: "listener_id",
			},
			"ListenerName": {
				TargetField: "listener_name",
			},
			"LoadBalancerId": {
				TargetField: "load_balancer_id",
			},
			"Description": {
				TargetField: "description",
			},
			"Protocol": {
				TargetField: "protocol",
			},
			"Port": {
				TargetField: "port",
			},
			"ServerGroupId": {
				TargetField: "server_group_id",
			},
			"ConnectionTimeout": {
				TargetField: "connection_timeout",
			},
			"Enabled": {
				TargetField: "enabled",
			},
			"StartPort": {
				TargetField: "start_port",
			},
			"EndPort": {
				TargetField: "end_port",
			},
			"Status": {
				TargetField: "status",
			},
			"CreateTime": {
				TargetField: "create_time",
			},
			"UpdateTime": {
				TargetField: "update_time",
			},
			"Tags": {
				TargetField: "tags",
				Convert:     ve.TransTagsToResponse,
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineNlbListenerService) DatasourceResources(d *schema.ResourceData, r *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"load_balancer_id": {
				TargetField: "LoadBalancerId",
			},
			"listener_name": {
				TargetField: "ListenerName",
			},
			"protocol": {
				TargetField: "Protocol",
			},
			"listener_ids": {
				TargetField: "ListenerIds",
				ConvertType: ve.ConvertWithN,
			},
			"tags": {
				TargetField: "TagFilters",
				Convert:     transTagsToRequest,
				Ignore:      true,
			},
		},
		CollectField: "listeners",
		IdField:      "ListenerId",
		ResponseConverts: map[string]ve.ResponseConvert{
			"ListenerId": {
				TargetField: "listener_id",
			},
			"ListenerName": {
				TargetField: "listener_name",
			},
			"LoadBalancerId": {
				TargetField: "load_balancer_id",
			},
			"Description": {
				TargetField: "description",
			},
			"Status": {
				TargetField: "status",
			},
			"Protocol": {
				TargetField: "protocol",
			},
			"Port": {
				TargetField: "port",
			},
			"ServerGroupId": {
				TargetField: "server_group_id",
			},
			"ConnectionTimeout": {
				TargetField: "connection_timeout",
			},
			"CreateTime": {
				TargetField: "create_time",
			},
			"UpdateTime": {
				TargetField: "update_time",
			},
			"Enabled": {
				TargetField: "enabled",
			},
			"StartPort": {
				TargetField: "start_port",
			},
			"EndPort": {
				TargetField: "end_port",
			},
			"Tags": {
				TargetField: "tags",
				Convert:     ve.TransTagsToResponse,
			},
		},
	}
}

func (s *VolcengineNlbListenerService) ReadResourceId(id string) string {
	return id
}

func (s *VolcengineNlbListenerService) getDiff(old, new []interface{}) (toAdd, toRemove []interface{}) {
	oldMap := make(map[string]bool)
	for _, v := range old {
		if s, ok := v.(string); ok {
			oldMap[s] = true
		}
	}
	newMap := make(map[string]bool)
	for _, v := range new {
		if s, ok := v.(string); ok {
			newMap[s] = true
		}
	}

	for _, v := range new {
		if s, ok := v.(string); ok {
			if !oldMap[s] {
				toAdd = append(toAdd, v)
			}
		}
	}
	for _, v := range old {
		if s, ok := v.(string); ok {
			if !newMap[s] {
				toRemove = append(toRemove, v)
			}
		}
	}
	return
}

func (s *VolcengineNlbListenerService) interfaceSliceToStringSlice(v []interface{}) []string {
	res := make([]string, 0, len(v))
	for _, item := range v {
		if s, ok := item.(string); ok {
			res = append(res, s)
		}
	}
	return res
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "clb",
		Action:      actionName,
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		RegionType:  ve.Regional,
	}
}

func transTagsToRequest(d *schema.ResourceData, i interface{}) interface{} {
	if i == nil {
		return nil
	}
	var tags []map[string]interface{}
	if set, ok := i.(*schema.Set); ok {
		if set.Len() == 0 {
			return nil
		}
		for _, v := range set.List() {
			if m, ok := v.(map[string]interface{}); ok {
				tag := make(map[string]interface{})
				if key, ok := m["key"].(string); ok {
					tag["Key"] = key
				}
				if value, ok := m["value"].(string); ok {
					tag["Values.1"] = value
				}
				tags = append(tags, tag)
			}
		}
	}
	return tags
}
