package veecp_edge_node

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
	"github.com/volcengine/terraform-provider-volcengine/volcengine/veecp/veecp_edge_node_pool"
)

type VolcengineVeecpNodeService struct {
	Client          *ve.SdkClient
	Dispatcher      *ve.Dispatcher
	nodePoolService *veecp_edge_node_pool.VolcengineVeecpNodePoolService
}

func NewVeecpNodeService(c *ve.SdkClient) *VolcengineVeecpNodeService {
	return &VolcengineVeecpNodeService{
		Client:          c,
		Dispatcher:      &ve.Dispatcher{},
		nodePoolService: veecp_edge_node_pool.NewVeecpNodePoolService(c),
	}
}

func (s *VolcengineVeecpNodeService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineVeecpNodeService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {

		action := "ListEdgeNodes"

		//if filter, filterExist := condition["Filter"]; filterExist {
		//	if statuses, exist := filter.(map[string]interface{})["Statuses"]; exist {
		//		for index, status := range statuses.([]interface{}) {
		//			if ty, ex := status.(map[string]interface{})["ConditionsType"]; ex {
		//				condition["Filter"].(map[string]interface{})["Statuses"].([]interface{})[index].(map[string]interface{})["Conditions.Type"] = ty
		//				delete(condition["Filter"].(map[string]interface{})["Statuses"].([]interface{})[index].(map[string]interface{}), "ConditionsType")
		//			}
		//		}
		//	}
		//}

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

func (s *VolcengineVeecpNodeService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"Filter": map[string]interface{}{
			"Ids": []string{id},
		},
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
		return data, fmt.Errorf("veecp_node %s not exist ", id)
	}
	// name会被修改，这里只能从tf中获取
	data["Name"] = resourceData.Get("name")
	return data, err
}

func (s *VolcengineVeecpNodeService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				d          map[string]interface{}
				status     interface{}
				failStates []string
			)
			failStates = append(failStates, "Failed")
			d, err = s.ReadResource(resourceData, id)
			if err != nil {
				return nil, "", err
			}
			status, err = ve.ObtainSdkValue("Status.Phase", d)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("veecp_node status error, status: %s", status.(string))
				}
			}
			return d, status.(string), err
		},
	}
}

func (s *VolcengineVeecpNodeService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateEdgeNode",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"auto_complete_config": {
					TargetField: "AutoCompleteConfig",
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"enable": {
							ForceGet: true,
						},
						"machine_auth": {
							TargetField: "MachineAuth",
							ConvertType: ve.ConvertJsonObject,
							NextLevelConvert: map[string]ve.RequestConvert{
								"ssh_port": {
									TargetField: "SSHPort",
								},
							},
						},
						"direct_add_instances": {
							TargetField: "DirectAddInstances",
							ConvertType: ve.ConvertJsonObjectArray,
						},
					},
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				autoCompleteConfig := d.Get("auto_complete_config").([]interface{})
				if !autoCompleteConfig[0].(map[string]interface{})["enable"].(bool) {
					return false, fmt.Errorf("AutoCompleteConfig.Enable must be true")
				}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam, *resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				time.Sleep(time.Second * 5)
				tmpIds, _ := ve.ObtainSdkValue("Result.Ids", *resp)
				ids := tmpIds.([]interface{})
				d.SetId(ids[0].(string))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Running", "Failed"},
				Timeout: 2 * time.Hour,
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("cluster_id").(string)
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineVeecpNodeService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"CreateClientToken": {
				TargetField: "client_token",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineVeecpNodeService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{}
	return []ve.Callback{callback}
}

func (s *VolcengineVeecpNodeService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteEdgeNodes",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"ClusterId":                resourceData.Get("cluster_id"),
				"NodePoolId":               resourceData.Get("node_pool_id"),
				"Ids.1":                    resourceData.Id(),
				"CascadingDeleteResources": []string{"Ecs"},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				nodePool, err := s.nodePoolService.ReadResources(map[string]interface{}{
					"Filter": map[string]interface{}{
						"Ids": []string{d.Get("node_pool_id").(string)},
					},
				})
				if err != nil {
					return false, err
				}
				if len(nodePool) == 0 {
					return false, fmt.Errorf("node pool not found")
				}
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
							return resource.NonRetryableError(fmt.Errorf("error on reading veecp edge node on delete %q, %w", d.Id(), callErr))
						}
					}
					_, callErr = call.ExecuteCall(d, client, call)
					if callErr == nil {
						return nil
					}
					return resource.RetryableError(callErr)
				})
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				err := ve.CheckResourceUtilRemoved(d, s.ReadResource, 10*time.Minute)
				time.Sleep(10 * time.Second)
				return err
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("cluster_id").(string)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineVeecpNodeService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
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
			"zone_ids": {
				TargetField: "Filter.ZoneIds",
				ConvertType: ve.ConvertJsonArray,
			},
			"statuses": {
				TargetField: "Filter.Statuses",
				ConvertType: ve.ConvertJsonObjectArray,
				NextLevelConvert: map[string]ve.RequestConvert{
					"phase": {
						TargetField: "Phase",
					},
					"edge_node_status_condition_type": {
						TargetField: "EdgeNodeStatusConditionType",
					},
				},
			},
			"ips": {
				TargetField: "Filter.Ips",
				ConvertType: ve.ConvertJsonArray,
			},
			"need_bootstrap_script": {
				TargetField: "Filter.NeedBootstrapScript",
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

func (s *VolcengineVeecpNodeService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "veecp_openapi",
		Version:     "2021-03-03",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}
