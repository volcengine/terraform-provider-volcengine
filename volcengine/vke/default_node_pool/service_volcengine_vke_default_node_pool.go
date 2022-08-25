package default_node_pool

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vke/node"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vke/node_pool"
)

type VolcengineDefaultNodePoolService struct {
	Client          *ve.SdkClient
	Dispatcher      *ve.Dispatcher
	nodePoolService *node_pool.VolcengineNodePoolService
	nodeService     *node.VolcengineVkeNodeService
}

func NewDefaultNodePoolService(c *ve.SdkClient) *VolcengineDefaultNodePoolService {
	return &VolcengineDefaultNodePoolService{
		Client:          c,
		Dispatcher:      &ve.Dispatcher{},
		nodePoolService: node_pool.NewNodePoolService(c),
		nodeService:     node.NewVolcengineVkeNodeService(c),
	}
}

const (
	DefaultNodePoolName = "vke-default-nodepool"
)

func (s *VolcengineDefaultNodePoolService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineDefaultNodePoolService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	return data, err
}

func (s *VolcengineDefaultNodePoolService) ReadResource(resourceData *schema.ResourceData, nodePoolId string) (data map[string]interface{}, err error) {
	var (
		nodes []interface{}
	)
	if nodePoolId == "" {
		nodePoolId = s.ReadResourceId(resourceData.Id())
	}
	data, err = s.nodePoolService.ReadResource(resourceData, nodePoolId)
	if err != nil {
		return data, err
	}

	// 只能导入默认节点池，不是默认节点池直接报错
	if data["Name"].(string) != DefaultNodePoolName {
		return nil, fmt.Errorf("only the default node pool is supported")
	}

	data["NodeConfig"].(map[string]interface{})["Security"].(map[string]interface{})["Login"].(map[string]interface{})["Password"] =
		resourceData.Get("node_config.0.security.0.login.0.password")

	nodes, err = s.nodeService.ReadResources(map[string]interface{}{
		"Filter": map[string]interface{}{
			"NodePoolIds": []string{nodePoolId},
		},
	})
	if err != nil {
		return data, err
	}

	instanceMap := make(map[string]string)
	instances := resourceData.Get("instances").(*schema.Set)
	for _, ins := range instances.List() {
		instancesId, _ := ve.ObtainSdkValue("instance_id", ins)
		imageId, _ := ve.ObtainSdkValue("image_id", ins)
		instanceMap[instancesId.(string)] = imageId.(string)
	}
	var ins []interface{}
	for _, n := range nodes {
		instancesId, _ := ve.ObtainSdkValue("InstanceId", n)
		if v, ok := instanceMap[instancesId.(string)]; ok {
			if v == "" {
				n.(map[string]interface{})["ImageId"] = ""
			}
			n.(map[string]interface{})["Phase"], _ = ve.ObtainSdkValue("Status.Phase", n)
			ins = append(ins, n)
		}
	}

	data["Instances"] = ins

	return data, err
}

func (s *VolcengineDefaultNodePoolService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				nodes      []interface{}
				status     interface{}
				statuses   []string
				failStates []string
			)
			instanceMap := make(map[string]bool)
			instances := resourceData.Get("instances").(*schema.Set)
			for _, ins := range instances.List() {
				instancesId, _ := ve.ObtainSdkValue("instance_id", ins)
				instanceMap[instancesId.(string)] = true
			}

			failStates = append(failStates, "Failed")
			nodes, err = s.nodeService.ReadResources(map[string]interface{}{
				"Filter": map[string]interface{}{
					"NodePoolIds": []string{id},
				},
			})
			if err != nil {
				return nil, "", err
			}
			for _, n := range nodes {
				instancesId, _ := ve.ObtainSdkValue("InstanceId", n)
				sts, _ := ve.ObtainSdkValue("Status.Phase", n)
				if _, ok := instanceMap[instancesId.(string)]; ok {
					statuses = append(statuses, sts.(string))
				}
			}
			for _, v := range failStates {
				for _, v1 := range statuses {
					if v == v1 {
						return nil, "", fmt.Errorf("node status error, status:%s", status.(string))
					}
				}
			}

			for _, v := range statuses {
				if v != "Running" && v != "Stopped" && v != "Failed" {
					return nodes, v, err
				}
			}

			return nodes, "Running", err
		},
	}
}

func (s *VolcengineDefaultNodePoolService) WithResourceResponseHandlers(nodePool map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		var (
			security = make([]interface{}, 0)
			login    = make([]interface{}, 0)
		)

		priSecurity := nodePool["NodeConfig"].(map[string]interface{})["Security"]
		priLogin := priSecurity.(map[string]interface{})["Login"]
		delete(nodePool, "Login")
		login = append(login, priLogin)
		priSecurity.(map[string]interface{})["Login"] = login
		security = append(security, priSecurity)

		delete(nodePool, "Security")
		nodePool["NodeConfig"].(map[string]interface{})["Security"] = security

		return nodePool, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineDefaultNodePoolService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var calls []ve.Callback
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateDefaultNodePool",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"instances": {
					Ignore: true,
				},
				"kubernetes_config": {
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"labels": {
							ConvertType: ve.ConvertJsonArray,
						},
						"taints": {
							ConvertType: ve.ConvertJsonArray,
						},
					},
				},
				"node_config": {
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"security": {
							ConvertType: ve.ConvertJsonObject,
							NextLevelConvert: map[string]ve.RequestConvert{
								"login": {
									ConvertType: ve.ConvertJsonObject,
								},
								"security_group_ids": {
									ConvertType: ve.ConvertJsonArray,
								},
								"security_strategies": {
									ConvertType: ve.ConvertJsonArray,
								},
							},
						},
					},
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				(*call.SdkParam)["ClientToken"] = "default-nodeService-pool-" + uuid.New().String()
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("Result.Id", *resp)
				d.SetId(id.(string))
				return nil
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				node_pool.NewNodePoolService(s.Client): {
					Target:  []string{"Running"},
					Timeout: resourceData.Timeout(schema.TimeoutCreate),
				},
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("cluster_id").(string)
			},
		},
	}

	calls = append(calls, callback)

	if _, ok := resourceData.GetOk("instances"); ok {
		calls = s.processNodeInstances(resourceData, calls)
	}

	return calls
}

func (s *VolcengineDefaultNodePoolService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var calls []ve.Callback
	// 先修改节点池配置
	calls = append(calls, s.nodePoolService.ModifyResource(resourceData, resource)...)
	//修改实例
	if resourceData.HasChange("instances") {
		calls = s.processNodeInstances(resourceData, calls)
	}
	return calls
}

func (s *VolcengineDefaultNodePoolService) RemoveResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteNodePool",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"Id":        resourceData.Id(),
				"ClusterId": resourceData.Get("cluster_id"),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 5*time.Minute)
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineDefaultNodePoolService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineDefaultNodePoolService) ReadResourceId(id string) string {
	return id
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

func (s *VolcengineDefaultNodePoolService) processNodeInstances(resourceData *schema.ResourceData, calls []ve.Callback) []ve.Callback {
	add, remove, _, _ := ve.GetSetDifference("instances", resourceData, defaultNodePoolNodeHash, false)
	logger.Debug(logger.RespFormat, "processNodeInstancesAdd", add)
	logger.Debug(logger.RespFormat, "processNodeInstancesRemove", remove)
	newNode := make(map[string][]string)
	var delNode []string
	if add != nil {
		for _, v := range add.List() {
			m := v.(map[string]interface{})
			key := strconv.FormatBool(m["keep_instance_name"].(bool)) + ":" + strconv.FormatBool(m["additional_container_storage_enabled"].(bool)) + ":" +
				m["image_id"].(string) + ":" + m["container_storage_path"].(string)
			if _, ok1 := newNode[key]; !ok1 {
				newNode[key] = []string{}
			}
			newNode[key] = append(newNode[key], m["instance_id"].(string))
		}
	}
	if remove != nil {
		for _, v := range remove.List() {
			m := v.(map[string]interface{})
			delNode = append(delNode, m["id"].(string))
		}
	}

	// 删除节点
	for i := 0; i < len(delNode)/100+1; i++ {
		start := i * 100
		end := (i + 1) * 100
		if end > len(delNode) {
			end = len(delNode)
		}
		if end <= start {
			break
		}
		calls = append(calls, func(nodeIds []string, clusterId, nodePoolId string) ve.Callback {
			return ve.Callback{
				Call: ve.SdkCall{
					Action:      "DeleteNodes",
					ConvertMode: ve.RequestConvertIgnore,
					ContentType: ve.ContentTypeJson,
					SdkParam: &map[string]interface{}{
						"ClusterId":  clusterId,
						"NodePoolId": nodePoolId,
					},
					BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
						if len(nodeIds) < 1 {
							return false, nil
						}
						for index, id := range nodeIds {
							(*call.SdkParam)[fmt.Sprintf("Ids.%d", index+1)] = id
						}
						return true, nil
					},
					ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
						logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
						resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
						logger.Debug(logger.RespFormat, call.Action, resp, err)
						return resp, err
					},
					Refresh: &ve.StateRefresh{
						Target:  []string{"Running"},
						Timeout: resourceData.Timeout(schema.TimeoutCreate),
					},
					ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
						node_pool.NewNodePoolService(s.Client): {
							Target:  []string{"Running"},
							Timeout: resourceData.Timeout(schema.TimeoutCreate),
						},
					},
					LockId: func(d *schema.ResourceData) string {
						return d.Get("cluster_id").(string)
					},
				},
			}
		}(delNode[start:end], resourceData.Get("cluster_id").(string), resourceData.Id()))
	}

	// 新增加节点
	for k, v := range newNode {
		nodeCall := ve.Callback{
			Call: ve.SdkCall{
				Action:      "CreateNodes",
				ConvertMode: ve.RequestConvertIgnore,
				ContentType: ve.ContentTypeJson,
				SdkParam: &map[string]interface{}{
					"Key":   k,
					"Value": v,
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["ClientToken"] = "default-nodeService-pool-" + uuid.New().String()
					(*call.SdkParam)["ClusterId"] = d.Get("cluster_id")
					for i, v1 := range (*call.SdkParam)["Value"].([]string) {
						(*call.SdkParam)["InstanceIds."+strconv.Itoa(i+1)] = v1
					}
					ks := strings.Split((*call.SdkParam)["Key"].(string), ":")
					keepInstanceName, _ := strconv.ParseBool(ks[0])
					additionalContainerStorageEnabled, _ := strconv.ParseBool(ks[1])
					(*call.SdkParam)["KeepInstanceName"] = keepInstanceName
					(*call.SdkParam)["AdditionalContainerStorageEnabled"] = additionalContainerStorageEnabled
					if ks[2] != "" {
						(*call.SdkParam)["ImageId"] = ks[2]
					}
					if len(ks) == 4 && ks[3] != "" {
						(*call.SdkParam)["ContainerStoragePath"] = ks[3]
					} else if len(ks) > 4 {
						(*call.SdkParam)["ContainerStoragePath"] = (*call.SdkParam)["Key"].(string)[len(ks[0])+1+len(ks[1])+1+len(ks[2])+1:]
					}

					delete(*call.SdkParam, "Key")
					delete(*call.SdkParam, "Value")

					logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
					resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
					logger.Debug(logger.RespFormat, call.Action, resp, err)
					return resp, err
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{"Running"},
					Timeout: resourceData.Timeout(schema.TimeoutCreate),
				},
				ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
					node_pool.NewNodePoolService(s.Client): {
						Target:  []string{"Running"},
						Timeout: resourceData.Timeout(schema.TimeoutCreate),
					},
				},
				LockId: func(d *schema.ResourceData) string {
					return d.Get("cluster_id").(string)
				},
			},
		}

		calls = append(calls, nodeCall)
	}
	return calls
}
