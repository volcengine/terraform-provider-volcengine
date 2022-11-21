package default_node_pool_batch_attach

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vke/default_node_pool"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vke/node_pool"
)

type VolcengineVkeDefaultNodePoolBatchAttachService struct {
	Client                 *ve.SdkClient
	Dispatcher             *ve.Dispatcher
	defaultNodePoolService *default_node_pool.VolcengineDefaultNodePoolService
}

func NewVolcengineVkeDefaultNodePoolBatchAttachService(c *ve.SdkClient) *VolcengineVkeDefaultNodePoolBatchAttachService {
	return &VolcengineVkeDefaultNodePoolBatchAttachService{
		Client:                 c,
		Dispatcher:             &ve.Dispatcher{},
		defaultNodePoolService: default_node_pool.NewDefaultNodePoolService(c),
	}
}

func (s *VolcengineVkeDefaultNodePoolBatchAttachService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineVkeDefaultNodePoolBatchAttachService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	return data, err
}

func (s *VolcengineVkeDefaultNodePoolBatchAttachService) ReadResource(resourceData *schema.ResourceData, nodePoolId string) (data map[string]interface{}, err error) {
	if nodePoolId == "" {
		nodePoolId = s.ReadResourceId(resourceData.Id())
	}
	return s.defaultNodePoolService.ReadResource(resourceData, nodePoolId)
}

func (s *VolcengineVkeDefaultNodePoolBatchAttachService) RefreshResourceState(data *schema.ResourceData, strings []string, duration time.Duration, id string) *resource.StateChangeConf {
	return s.defaultNodePoolService.RefreshResourceState(data, strings, duration, id)
}

func (s *VolcengineVkeDefaultNodePoolBatchAttachService) WithResourceResponseHandlers(m map[string]interface{}) []ve.ResourceResponseHandler {
	return s.defaultNodePoolService.WithResourceResponseHandlers(m)
}

func (s *VolcengineVkeDefaultNodePoolBatchAttachService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var calls []ve.Callback
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateDefaultNodePool",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				resp := make(map[string]interface{})
				resp["Id"] = (*call.SdkParam)["DefaultNodePoolId"]
				return &resp, nil
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("Id", *resp)
				d.SetId(id.(string))
				return nil
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				node_pool.NewNodePoolService(s.Client): {
					Target:  []string{"Running"},
					Timeout: resourceData.Timeout(schema.TimeoutCreate),
				},
			},
		},
	}

	calls = append(calls, callback)

	calls = s.defaultNodePoolService.ProcessNodeInstances(resourceData, calls)

	return calls
}

func (s *VolcengineVkeDefaultNodePoolBatchAttachService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var calls []ve.Callback
	calls = s.defaultNodePoolService.ProcessNodeInstances(resourceData, calls)
	return calls
}

func (s *VolcengineVkeDefaultNodePoolBatchAttachService) RemoveResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var calls []ve.Callback
	var delNode []string
	nv := resourceData.Get("instances")
	if nv == nil {
		nv = new(schema.Set)
	}
	remove := nv.(*schema.Set)

	for _, v := range remove.List() {
		m := v.(map[string]interface{})
		delNode = append(delNode, m["id"].(string))
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

	return calls
}

func (s *VolcengineVkeDefaultNodePoolBatchAttachService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineVkeDefaultNodePoolBatchAttachService) ReadResourceId(id string) string {
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
