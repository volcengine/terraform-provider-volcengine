package ecs_hpc_cluster

import (
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineEcsHpcClusterService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewEcsHpcClusterService(c *ve.SdkClient) *VolcengineEcsHpcClusterService {
	return &VolcengineEcsHpcClusterService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineEcsHpcClusterService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineEcsHpcClusterService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithNextTokenQuery(m, "MaxResults", "NextToken", 20, nil, func(condition map[string]interface{}) (data []interface{}, next string, err error) {
		action := "DescribeHpcClusters"
		logger.Debug(logger.ReqFormat, action, condition)
		if condition == nil {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), nil)
			if err != nil {
				return data, next, err
			}
		} else {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &condition)
			if err != nil {
				return data, next, err
			}
		}
		logger.Debug(logger.RespFormat, action, condition, *resp)

		results, err = ve.ObtainSdkValue("Result.HpcClusters", *resp)
		if err != nil {
			return data, next, err
		}
		nextToken, err := ve.ObtainSdkValue("Result.NextToken", *resp)
		if err != nil {
			return data, next, err
		}
		next = nextToken.(string)
		if results == nil {
			results = []interface{}{}
		}

		if _, ok = results.([]interface{}); !ok {
			return data, next, errors.New("Result.HpcClusters is not Slice")
		}

		return results.([]interface{}), next, err
	})
}

func (s *VolcengineEcsHpcClusterService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		result interface{}
		ok     bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}

	action := "DescribeHpcCluster"
	req := map[string]interface{}{
		"HpcClusterId": id,
	}
	logger.Debug(logger.ReqFormat, action, req)
	resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(action), &req)
	if err != nil {
		return data, err
	}
	logger.Debug(logger.RespFormat, action, req, *resp)

	result, err = ve.ObtainSdkValue("Result", *resp)
	if err != nil {
		return data, err
	}
	if data, ok = result.(map[string]interface{}); !ok {
		return data, errors.New("Value is not map ")
	}
	if len(data) == 0 {
		return data, fmt.Errorf("ecs_hpc_cluster %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineEcsHpcClusterService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (VolcengineEcsHpcClusterService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineEcsHpcClusterService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateHpcCluster",
			ConvertMode: ve.RequestConvertAll,
			Convert:     map[string]ve.RequestConvert{},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("Result.HpcClusterId", *resp)
				d.SetId(id.(string))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineEcsHpcClusterService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyHpcClusterAttribute",
			ConvertMode: ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"name": {
					TargetField: "HpcClusterName",
				},
				"description": {
					TargetField: "Description",
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if len(*call.SdkParam) > 0 {
					(*call.SdkParam)["HpcClusterId"] = d.Id()
					return true, nil
				}
				return false, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineEcsHpcClusterService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteHpcCluster",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"HpcClusterId": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 5*time.Minute)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading hpc cluster on delete %q, %w", d.Id(), callErr))
						}
					}
					_, callErr = call.ExecuteCall(d, client, call)
					if callErr == nil {
						return nil
					}
					return resource.RetryableError(callErr)
				})
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineEcsHpcClusterService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		NameField:    "Name",
		IdField:      "HpcClusterId",
		CollectField: "hpc_clusters",
		ResponseConverts: map[string]ve.ResponseConvert{
			"HpcClusterId": {
				TargetField: "id",
				KeepDefault: true,
			},
		},
	}
}

func (s *VolcengineEcsHpcClusterService) ReadResourceId(id string) string {
	return id
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
