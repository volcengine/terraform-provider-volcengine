package node

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-vestack/common"
	"github.com/volcengine/terraform-provider-vestack/logger"
)

type VestackVkeNodeService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewVestackVkeNodeService(c *ve.SdkClient) *VestackVkeNodeService {
	return &VestackVkeNodeService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

const (
	nodeAdd    = "Add"
	nodeRemove = "Remove"
	nodeWait   = "Waiting"
)

func (s *VestackVkeNodeService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VestackVkeNodeService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListNodes"
		// 单独适配VKE Conditions.Type 字段，该字段 API 表示不规范
		if filter, filterExist := condition["Filter"]; filterExist {
			if statuses, exist := filter.(map[string]interface{})["Statuses"]; exist {
				for index, status := range statuses.([]interface{}) {
					if ty, ex := status.(map[string]interface{})["ConditionsType"]; ex {
						condition["Filter"].(map[string]interface{})["Statuses"].([]interface{})[index].(map[string]interface{})["Conditions.Type"] = ty
						delete(condition["Filter"].(map[string]interface{})["Statuses"].([]interface{})[index].(map[string]interface{}), "ConditionsType")
					}
				}
			}
		}

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
		results, err = ve.ObtainSdkValue("Result.Items", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.Items is not Slice")
		}
		return data, err
	})
}

func (s *VestackVkeNodeService) ReadResource(resourceData *schema.ResourceData, tmpId string) (data map[string]interface{}, err error) {
	var (
		results    []interface{}
		ok         bool
		res        = make(map[string]interface{})
		clusterSet = make(map[string]bool)
		lossIds    = make([]string, 0)
	)
	if tmpId == "" {
		tmpId = s.ReadResourceId(resourceData.Id())
	}
	ids := strings.Split(tmpId, ":")

	req := map[string]interface{}{
		"Filter": map[string]interface{}{
			"Ids": ids,
		},
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}

	for _, v := range results {
		if data, ok = v.(map[string]interface{}); !ok {
			return res, errors.New("Value is not map ")
		}
		res[data["Id"].(string)] = v
		clusterSet[data["ClusterId"].(string)] = true
	}
	if len(clusterSet) > 1 {
		return res, errors.New("all nodes should be the same cluster")
	}

	for _, id := range ids {
		if _, ok = res[id]; !ok {
			lossIds = append(lossIds, id)
		}
	}
	if len(lossIds) > 0 {
		return res, fmt.Errorf("nodes not exist: %s", strings.Join(lossIds, ","))
	}

	return res, err
}

func (s *VestackVkeNodeService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				demo       map[string]interface{}
				status     interface{}
				failStates []string
				failIds    = make([]string, 0)
				waitIds    = make([]string, 0)
			)
			failStates = append(failStates, "Failed")
			demo, err = s.ReadResource(resourceData, id)
			if err != nil {
				// 删除时特殊处理一下
				if ve.ResourceNotFoundError(err) && target[0] == nodeRemove {
					if len(demo) == 0 {
						return demo, nodeRemove, nil
					}
					for tmpId, _ := range demo {
						failIds = append(failIds, tmpId)
					}
					idStr := strings.Join(failIds, ":")
					resourceData.SetId(idStr)
					return nil, nodeWait, nil
				}
				return nil, "", err
			}
			for tmpId, tmp := range demo {
				data := tmp.(map[string]interface{})
				status, err = ve.ObtainSdkValue("Status.Phase", data)
				if err != nil {
					return nil, "", err
				}
				if target[0] == nodeAdd {
					if status == "Failed" {
						failIds = append(failIds, tmpId)
					}
				}
				if status == "Deleting" || status == "Creating" || status == "Updating" {
					waitIds = append(waitIds, tmpId)
				}
			}
			if len(waitIds) > 0 {
				return demo, nodeWait, nil
			}
			if nodeAdd == target[0] {
				if len(failIds) == 0 {
					return demo, nodeAdd, nil
				} else {
					return demo, nodeAdd, fmt.Errorf("create some nodes failed: %s", strings.Join(failIds, ","))
				}
			} else if nodeRemove == target[0] {
				return demo, nodeWait, nil
			}
			//注意 返回的第一个参数不能为空 否则会一直等下去
			return demo, "", fmt.Errorf("not support target status %s", target[0])
		},
	}
}

func (VestackVkeNodeService) WithResourceResponseHandlers(nodes map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		ids := make([]string, 0)
		instanceIds := make([]string, 0)
		clusterId := ""
		for id, node := range nodes {
			ids = append(ids, id)
			instanceIds = append(instanceIds, (node.(map[string]interface{}))["InstanceId"].(string))
			clusterId = (node.(map[string]interface{}))["ClusterId"].(string)
		}
		nodes["NodeIds"] = ids
		nodes["InstanceIds"] = instanceIds
		nodes["ClusterId"] = clusterId
		return nodes, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}

}

func (s *VestackVkeNodeService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := s.addNodes(resourceData.Get("instance_ids").(*schema.Set).List(), resourceData)
	return []ve.Callback{callback}
}

func (s *VestackVkeNodeService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callbacks := make([]ve.Callback, 0)
	add, remove, _, _ := ve.GetSetDifference("instance_ids", resourceData, schema.HashString, false)
	if remove != nil && len(remove.List()) > 0 {
		callbacks = append(callbacks, s.removeNodes(remove.List(), resourceData)...)
	}
	if add != nil && len(add.List()) > 0 {
		callbacks = append(callbacks, s.addNodes(add.List(), resourceData))
	}
	return callbacks
}

func (s *VestackVkeNodeService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := s.removeNodes(resourceData.Get("instance_ids").(*schema.Set).List(), resourceData)
	return callback
}

func (s *VestackVkeNodeService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "Filter.Ids",
				ConvertType: ve.ConvertJsonArray,
			},
			"cluster_ids": {
				TargetField: "Filter.ClusterIds",
				ConvertType: ve.ConvertJsonArray,
			},
			"name": {
				TargetField: "Filter.Name",
			},
			"create_client_token": {
				TargetField: "Filter.CreateClientToken",
			},
			"node_pool_ids": {
				TargetField: "Filter.NodePoolIds",
				ConvertType: ve.ConvertJsonArray,
			},
			"statuses": {
				TargetField: "Filter.Statuses",
				ConvertType: ve.ConvertJsonObjectArray,
				NextLevelConvert: map[string]ve.RequestConvert{
					"phase": {
						TargetField: "Phase",
					},
					"conditions_type": {
						TargetField: "ConditionsType",
					},
				},
			},
		},
		ContentType:  ve.ContentTypeJson,
		NameField:    "Name",
		IdField:      "Id",
		CollectField: "nodes",
		ResponseConverts: map[string]ve.ResponseConvert{
			"Id": {
				TargetField: "id",
				KeepDefault: true,
			},
			"Status.Phase": {
				TargetField: "phase",
			},
			"Status.Conditions": {
				TargetField: "condition_types",
				Convert: func(i interface{}) interface{} {
					var results []interface{}
					if dd, ok := i.([]interface{}); ok {
						for _, _data := range dd {
							results = append(results, _data.(map[string]interface{})["Type"])
						}
					}
					return results
				},
			},
		},
	}
}

func (s *VestackVkeNodeService) ReadResourceId(id string) string {
	return id
}

func (s *VestackVkeNodeService) addNodes(instanceIds []interface{}, resourceData *schema.ResourceData) ve.Callback {
	return ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateNodes",
			ConvertMode: ve.RequestConvertInConvert,
			ContentType: ve.ContentTypeJson,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if len(instanceIds) < 1 {
					return false, nil
				}
				for i, id := range instanceIds {
					(*call.SdkParam)[fmt.Sprintf("InstanceIds.%d", i+1)] = id
				}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				tmpIds, _ := ve.ObtainSdkValue("Result.Ids", *resp)
				ids := make([]string, 0)
				if len(resourceData.Id()) > 0 {
					ids = strings.Split(resourceData.Id(), ":")
				}
				for _, id := range tmpIds.([]interface{}) {
					ids = append(ids, id.(string))
				}
				d.SetId(strings.Join(ids, ":"))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{nodeAdd},
				Timeout: 2 * time.Hour,
			},
			Convert: map[string]ve.RequestConvert{
				"client_token": {
					ForceGet:    true,
					TargetField: "ClientToken",
				},
				"cluster_id": {
					ForceGet:    true,
					TargetField: "ClusterId",
				},
				"keep_instance_name": {
					ForceGet:    true,
					TargetField: "KeepInstanceName",
				},
				"additional_container_storage_enabled": {
					ForceGet:    true,
					TargetField: "AdditionalContainerStorageEnabled",
				},
				"container_storage_path": {
					ForceGet:    true,
					TargetField: "ContainerStoragePath",
				},
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("cluster_id").(string)
			},
		},
	}
}

func (s *VestackVkeNodeService) removeNodes(instanceIds []interface{}, resourceData *schema.ResourceData) []ve.Callback {
	if len(instanceIds) < 1 {
		return []ve.Callback{}
	}
	callbacks := make([]ve.Callback, 0)
	nodes, err := s.ReadResources(map[string]interface{}{
		"Filter": map[string]interface{}{
			"Ids": resourceData.Get("node_ids"),
		},
	})
	if err != nil {
		return []ve.Callback{
			{
				Call: ve.SdkCall{
					BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
						return false, err
					},
				},
			},
		}
	}

	// 这里需要查出来，知道节点信息和节点池信息
	removeSet := make(map[string]bool)
	nodeIds := make(map[string][]string, 0)
	for _, id := range instanceIds {
		removeSet[id.(string)] = true
	}
	for _, node := range nodes {
		data := node.(map[string]interface{})
		if !removeSet[data["InstanceId"].(string)] {
			continue
		}
		poolId, nodeId := data["NodePoolId"].(string), data["Id"].(string)
		ids, ok := nodeIds[poolId]
		if !ok {
			ids = make([]string, 0)
		}
		ids = append(ids, nodeId)
		nodeIds[poolId] = ids
	}

	// 根据不同的节点池进行删除
	tmpId := resourceData.Id()
	for poolId, ids := range nodeIds {
		callbacks = append(callbacks, ve.Callback{
			Call: ve.SdkCall{
				Action:      "DeleteNodes",
				ConvertMode: ve.RequestConvertIgnore,
				ContentType: ve.ContentTypeJson,
				SdkParam: &map[string]interface{}{
					"ClusterId":  resourceData.Get("cluster_id"),
					"NodePoolId": poolId,
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					if len(ids) < 1 {
						return false, nil
					}

					if resourceData.Get("cascading_delete_resources") != nil {
						for i, v := range resourceData.Get("cascading_delete_resources").(*schema.Set).List() {
							(*call.SdkParam)[fmt.Sprintf("CascadingDeleteResources.%d", i+1)] = v.(string)
						}
					}
					for i, id := range ids {
						(*call.SdkParam)[fmt.Sprintf("Ids.%d", i+1)] = id
					}
					resourceData.SetId(strings.Join(ids, ":"))
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
				CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
					//出现错误后重试
					return resource.Retry(15*time.Minute, func() *resource.RetryError {
						_, callErr := s.ReadResource(d, "")
						if callErr != nil {
							if ve.ResourceNotFoundError(callErr) && strings.Contains(callErr.Error(), strings.Join(strings.Split(resourceData.Id(), ":"), ",")) {
								return nil
							} else {
								return resource.NonRetryableError(fmt.Errorf("error on reading vke node on delete %q, %w", d.Id(), callErr))
							}
						}
						_, callErr = call.ExecuteCall(d, client, call)
						logger.Debug(logger.RespFormat, call.Action, callErr)
						if callErr == nil {
							return nil
						}
						return resource.RetryableError(callErr)
					})
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{nodeRemove},
					Timeout: 2 * time.Hour,
				},
				LockId: func(d *schema.ResourceData) string {
					return d.Get("cluster_id").(string)
				},
			},
		})
	}

	if len(callbacks) > 0 {
		// 这一步的目的是更新node_ids和instance_ids
		callbacks = append(callbacks, ve.Callback{
			Call: ve.SdkCall{
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					readResource, err := s.ReadResource(resourceData, tmpId)
					if err != nil && !ve.ResourceNotFoundError(err) {
						return nil, err
					}
					ids := make([]string, 0)
					for id := range readResource {
						ids = append(ids, id)
					}
					tmpId = strings.Join(ids, ":")
					d.SetId(tmpId)
					return &readResource, nil
				},
			},
		})
	}
	return callbacks
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "vke",
		Version:     "2022-05-12",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}
