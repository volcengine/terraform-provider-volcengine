package kubeconfig

import (
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
	"time"
)

type VolcengineVkeKubeconfigService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewVkeKubeconfigService(c *ve.SdkClient) *VolcengineVkeKubeconfigService {
	return &VolcengineVkeKubeconfigService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineVkeKubeconfigService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineVkeKubeconfigService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)

	data, err = ve.WithPageNumberQuery(condition, "PageSize", "PageNumber", 10, 1, func(m map[string]interface{}) ([]interface{}, error) {
		action := "ListKubeconfigs"
		logger.Debug(logger.ReqFormat, action, condition)
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
	if err != nil {
		return data, err
	}

	return data, err
}

func (s *VolcengineVkeKubeconfigService) ReadResource(resourceData *schema.ResourceData, kubeconfigId string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if kubeconfigId == "" {
		kubeconfigId = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"Filter": map[string]interface{}{
			"Ids": []string{kubeconfigId},
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
		return data, fmt.Errorf("Vke Kubeconfig %s not exist ", kubeconfigId)
	}

	return data, err
}

func (s *VolcengineVkeKubeconfigService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh:    nil, // no need to refresh resources
	}
}

func (VolcengineVkeKubeconfigService) WithResourceResponseHandlers(kubeconfig map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return kubeconfig, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineVkeKubeconfigService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateKubeconfig",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				//创建Kubeconfig
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				//注意 获取内容 这个地方不能是指针 需要转一次
				id, _ := ve.ObtainSdkValue("Result.Id", *resp)
				d.SetId(id.(string))
				//异步操作 这里需要等5s
				time.Sleep(time.Duration(5) * time.Second)
				return nil
			},
		},
	}
	return []ve.Callback{callback}

}

func (s *VolcengineVkeKubeconfigService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineVkeKubeconfigService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteKubeconfigs",
			ContentType: ve.ContentTypeJson,
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["Ids"] = []string{resourceData.Id()}
				(*call.SdkParam)["ClusterId"] = resourceData.Get("cluster_id")
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				//删除Kubeconfig
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				//异步操作 这里需要等5s
				time.Sleep(time.Duration(5) * time.Second)
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 10*time.Minute)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading kubeconfig on delete %q, %w", d.Id(), callErr))
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

func (s *VolcengineVkeKubeconfigService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
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
			"types": {
				TargetField: "Filter.Types",
				ConvertType: ve.ConvertJsonArray,
			},
		},
		ContentType:  ve.ContentTypeJson,
		IdField:      "Id",
		CollectField: "kubeconfigs",
		ResponseConverts: map[string]ve.ResponseConvert{
			"Id": {
				TargetField: "kubeconfig_id",
				KeepDefault: true,
			},
		},
	}
}

func (s *VolcengineVkeKubeconfigService) ReadResourceId(id string) string {
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
