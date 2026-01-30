package vmp_integration_task_enable

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineVmpIntegrationTaskEnableService struct {
	Client *ve.SdkClient
}

func NewService(c *ve.SdkClient) *VolcengineVmpIntegrationTaskEnableService {
	return &VolcengineVmpIntegrationTaskEnableService{
		Client: c,
	}
}

func (s *VolcengineVmpIntegrationTaskEnableService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineVmpIntegrationTaskEnableService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	// This resource does not support query operations
	return nil, fmt.Errorf("vmp_integration_task_enable does not support query operations")
}

func (s *VolcengineVmpIntegrationTaskEnableService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	// This resource does not support read operations
	// Return minimal data to satisfy the interface
	return map[string]interface{}{
		"Id": id,
	}, nil
}

func (s *VolcengineVmpIntegrationTaskEnableService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (VolcengineVmpIntegrationTaskEnableService) WithResourceResponseHandlers(data map[string]interface{}) []ve.ResourceResponseHandler {
	return nil
}

func (s *VolcengineVmpIntegrationTaskEnableService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "EnableIntegrationTasks",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"task_ids": {
					TargetField: "Ids",
					ConvertType: ve.ConvertJsonArray,
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				// Use task_ids as the resource ID since this is an enable operation
				taskIds := d.Get("task_ids").(*schema.Set).List()
				if len(taskIds) > 0 {
					// Create a composite ID from all task IDs
					id := ""
					for i, taskId := range taskIds {
						if i > 0 {
							id += ","
						}
						id += taskId.(string)
					}
					d.SetId(id)
				}
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineVmpIntegrationTaskEnableService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	// This resource does not support update operations
	return nil
}

func (s *VolcengineVmpIntegrationTaskEnableService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DisableIntegrationTasks",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				// Parse task IDs from the resource ID
				taskIds := []string{}
				if d.Id() != "" {
					// Split the composite ID to get individual task IDs
					id := d.Id()
					taskIds = []string{id}
					// If ID contains commas, split by comma
					for i := 0; i < len(id); i++ {
						if id[i] == ',' {
							taskIds = []string{}
							start := 0
							for j := 0; j <= len(id); j++ {
								if j == len(id) || id[j] == ',' {
									if j > start {
										taskIds = append(taskIds, id[start:j])
									}
									start = j + 1
								}
							}
							break
						}
					}
				}
				(*call.SdkParam)["Ids"] = taskIds
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineVmpIntegrationTaskEnableService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	// This resource does not support data source operations
	return ve.DataSourceInfo{}
}

func (s *VolcengineVmpIntegrationTaskEnableService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "vmp",
		Version:     "2021-03-03",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}
