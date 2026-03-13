package nlb_backend_server

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineNlbBackendServerService struct {
	Client *ve.SdkClient
}

func NewNlbBackendServerService(c *ve.SdkClient) *VolcengineNlbBackendServerService {
	return &VolcengineNlbBackendServerService{
		Client: c,
	}
}

func (s *VolcengineNlbBackendServerService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineNlbBackendServerService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	action := "DescribeNLBServerGroupAttributes"
	logger.Debug(logger.ReqFormat, action, m)
	resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(action), &m)
	if err != nil {
		return nil, err
	}
	results, err := ve.ObtainSdkValue("Result.Servers", *resp)
	if err != nil {
		return nil, err
	}
	if results == nil {
		return []interface{}{}, nil
	}
	if data, ok := results.([]interface{}); ok {
		return data, nil
	}
	return nil, errors.New("Servers is not Slice")
}

func (s *VolcengineNlbBackendServerService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	if id == "" {
		id = resourceData.Id()
	}
	serverGroupId := id

	req := map[string]interface{}{
		"ServerGroupId": serverGroupId,
	}
	results, err := s.ReadResources(req)
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("resource not found")
	}
	return map[string]interface{}{
		"ServerGroupId": serverGroupId,
		"Servers":       results,
	}, nil
}

func (s *VolcengineNlbBackendServerService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Refresh: func() (interface{}, string, error) {
			res, err := s.ReadResource(resourceData, id)
			return res, "", err
		},
		Target:     target,
		Timeout:    timeout,
		Delay:      2 * time.Second,
		MinTimeout: 1 * time.Second,
	}
}

func (s *VolcengineNlbBackendServerService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action: "AddNLBBackendServers",
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				param, err := ve.ResourceDateToRequest(d, resource, false, s.createRequestConvert(), ve.RequestConvertInConvert, ve.ContentTypeDefault)
				if err != nil {
					return nil, err
				}
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), &param)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				v := d.Get("server_group_id")
				serverGroupId, ok := v.(string)
				if !ok {
					return errors.New("server_group_id is not string")
				}
				d.SetId(serverGroupId)
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineNlbBackendServerService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback
	if resourceData.HasChange("backend_servers") {
		oldV, newV := resourceData.GetChange("backend_servers")
		oldList, ok := oldV.([]interface{})
		if !ok {
			return []ve.Callback{{Err: errors.New("old backend_servers is not []interface{}")}}
		}
		newList, ok := newV.([]interface{})
		if !ok {
			return []ve.Callback{{Err: errors.New("new backend_servers is not []interface{}")}}
		}

		// 1. Find servers to remove
		var toRemove []string
		for _, oldItem := range oldList {
			oldMap, ok := oldItem.(map[string]interface{})
			if !ok {
				return []ve.Callback{{Err: errors.New("old backend_servers item is not map")}}
			}
			found := false
			for _, newItem := range newList {
				newMap, ok := newItem.(map[string]interface{})
				if !ok {
					return []ve.Callback{{Err: errors.New("new backend_servers item is not map")}}
				}
				if s.isSameServer(oldMap, newMap) {
					found = true
					break
				}
			}
			if !found {
				if serverId, ok := oldMap["server_id"].(string); ok && serverId != "" {
					toRemove = append(toRemove, serverId)
				}
			}
		}

		if len(toRemove) > 0 {
			callbacks = append(callbacks, ve.Callback{
				Call: ve.SdkCall{
					Action: "RemoveNLBBackendServers",
					ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
						param := map[string]interface{}{
							"ServerGroupId": d.Id(),
						}
						for i, id := range toRemove {
							param[fmt.Sprintf("ServerIds.%d", i+1)] = id
						}
						return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), &param)
					},
				},
			})
		}

		// 2. Find servers to add
		var toAdd []interface{}
		for _, newItem := range newList {
			newMap, ok := newItem.(map[string]interface{})
			if !ok {
				return []ve.Callback{{Err: errors.New("new backend_servers item is not map")}}
			}
			found := false
			for _, oldItem := range oldList {
				oldMap, ok := oldItem.(map[string]interface{})
				if !ok {
					return []ve.Callback{{Err: errors.New("old backend_servers item is not map")}}
				}
				if s.isSameServer(oldMap, newMap) {
					found = true
					break
				}
			}
			if !found {
				toAdd = append(toAdd, newItem)
			}
		}

		if len(toAdd) > 0 {
			callbacks = append(callbacks, ve.Callback{
				Call: ve.SdkCall{
					Action: "AddNLBBackendServers",
					ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
						param := map[string]interface{}{
							"ServerGroupId": d.Id(),
						}
						for i, item := range toAdd {
							m, ok := item.(map[string]interface{})
							if !ok {
								return nil, errors.New("Value is not map ")
							}
							prefix := fmt.Sprintf("Servers.%d.", i+1)
							param[prefix+"Type"] = m["type"]
							param[prefix+"InstanceId"] = m["instance_id"]
							param[prefix+"Ip"] = m["ip"]
							param[prefix+"Port"] = m["port"]
							param[prefix+"Weight"] = m["weight"]
							if v, ok := m["description"].(string); ok && v != "" {
								param[prefix+"Description"] = v
							}
							if v, ok := m["zone_id"].(string); ok && v != "" {
								param[prefix+"ZoneId"] = v
							}
						}
						return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), &param)
					},
				},
			})
		}

		// 3. Find servers to modify
		var toModify []interface{}
		for _, newItem := range newList {
			newMap, ok := newItem.(map[string]interface{})
			if !ok {
				return []ve.Callback{{Err: errors.New("new backend_servers item is not map")}}
			}
			for _, oldItem := range oldList {
				oldMap, ok := oldItem.(map[string]interface{})
				if !ok {
					return []ve.Callback{{Err: errors.New("old backend_servers item is not map")}}
				}
				if s.isSameServer(oldMap, newMap) {
					if !s.isSameAttributes(oldMap, newMap) {
						m := make(map[string]interface{})
						for k, v := range newMap {
							m[k] = v
						}
						m["server_id"] = oldMap["server_id"]
						toModify = append(toModify, m)
					}
					break
				}
			}
		}

		if len(toModify) > 0 {
			callbacks = append(callbacks, ve.Callback{
				Call: ve.SdkCall{
					Action: "ModifyNLBBackendServersAttributes",
					ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
						param := map[string]interface{}{
							"ServerGroupId": d.Id(),
						}
						for i, item := range toModify {
							m, ok := item.(map[string]interface{})
							if !ok {
								return nil, errors.New("Value is not map ")
							}
							prefix := fmt.Sprintf("Servers.%d.", i+1)
							param[prefix+"ServerId"] = m["server_id"]
							param[prefix+"Weight"] = s.toInt(m["weight"])
							if v, ok := m["description"].(string); ok && v != "" {
								param[prefix+"Description"] = v
							}
						}
						return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), &param)
					},
				},
			})
		}
	}
	return callbacks
}

func (s *VolcengineNlbBackendServerService) isSameServer(s1, s2 map[string]interface{}) bool {
	t1 := strings.ToLower(s.toString(s1["type"]))
	t2 := strings.ToLower(s.toString(s2["type"]))
	if t1 != t2 {
		return false
	}
	if s.toInt(s1["port"]) != s.toInt(s2["port"]) {
		return false
	}
	if t1 == "ecs" {
		return s.toString(s1["instance_id"]) == s.toString(s2["instance_id"])
	}
	return s.toString(s1["ip"]) == s.toString(s2["ip"])
}

func (s *VolcengineNlbBackendServerService) isSameAttributes(s1, s2 map[string]interface{}) bool {
	return s.toInt(s1["weight"]) == s.toInt(s2["weight"]) &&
		s.toString(s1["description"]) == s.toString(s2["description"])
}

func (s *VolcengineNlbBackendServerService) toString(i interface{}) string {
	if i == nil {
		return ""
	}
	if v, ok := i.(string); ok {
		return v
	}
	return fmt.Sprintf("%v", i)
}

func (s *VolcengineNlbBackendServerService) toInt(i interface{}) int {
	if i == nil {
		return 0
	}
	switch v := i.(type) {
	case int:
		return v
	case int32:
		return int(v)
	case int64:
		return int(v)
	case float32:
		return int(v)
	case float64:
		return int(v)
	case string:
		res, _ := strconv.Atoi(v)
		return res
	}
	return 0
}

func (s *VolcengineNlbBackendServerService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action: "RemoveNLBBackendServers",
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				serverGroupId := d.Id()
				backendServers, ok := d.Get("backend_servers").([]interface{})
				if !ok {
					return nil, errors.New("backend_servers is not []interface{}")
				}
				param := map[string]interface{}{
					"ServerGroupId": serverGroupId,
				}
				for i, item := range backendServers {
					m, ok := item.(map[string]interface{})
					if !ok {
						return nil, errors.New("Value is not map ")
					}
					if serverId, ok := m["server_id"].(string); ok && serverId != "" {
						param[fmt.Sprintf("ServerIds.%d", i+1)] = serverId
					}
				}
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), &param)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 10*time.Minute)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s VolcengineNlbBackendServerService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"ServerGroupId": {
				TargetField: "server_group_id",
			},
			"Servers": {
				TargetField: "backend_servers",
				Convert: func(i interface{}) interface{} {
					if i == nil {
						return nil
					}
					servers, ok := i.([]interface{})
					if !ok {
						return nil
					}
					var res []interface{}
					for _, v := range servers {
						m, ok := v.(map[string]interface{})
						if !ok {
							continue
						}
						item := make(map[string]interface{})
						item["server_id"] = m["ServerId"]
						item["type"] = m["Type"]
						item["instance_id"] = m["InstanceId"]
						item["ip"] = m["Ip"]
						item["port"] = s.toInt(m["Port"])
						item["weight"] = s.toInt(m["Weight"])
						item["description"] = m["Description"]
						if v, ok := m["ZoneId"]; ok {
							item["zone_id"] = s.toString(v)
						} else if v, ok := m["ZoneID"]; ok {
							item["zone_id"] = s.toString(v)
						}
						res = append(res, item)
					}
					return res
				},
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineNlbBackendServerService) createRequestConvert() map[string]ve.RequestConvert {
	return map[string]ve.RequestConvert{
		"server_group_id": {
			TargetField: "ServerGroupId",
		},
		"backend_servers": {
			TargetField: "Servers",
			ConvertType: ve.ConvertListN,
			NextLevelConvert: map[string]ve.RequestConvert{
				"type":        {TargetField: "Type"},
				"instance_id": {TargetField: "InstanceId"},
				"ip":          {TargetField: "Ip"},
				"port":        {TargetField: "Port"},
				"weight":      {TargetField: "Weight"},
				"description": {TargetField: "Description"},
				"zone_id":     {TargetField: "ZoneId"},
			},
		},
	}
}

func (s *VolcengineNlbBackendServerService) DatasourceResources(d *schema.ResourceData, r *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineNlbBackendServerService) ReadResourceId(id string) string {
	return id
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
