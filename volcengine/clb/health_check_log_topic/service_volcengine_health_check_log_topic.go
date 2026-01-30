package health_check_log_topic

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
)

type VolcengineHealthCheckLogTopicService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewHealthCheckLogTopicService(c *ve.SdkClient) *VolcengineHealthCheckLogTopicService {
	return &VolcengineHealthCheckLogTopicService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineHealthCheckLogTopicService) GetClient() *ve.SdkClient {
	return s.Client
}

// 查询日志主题绑定的CLB实例列表
func (s *VolcengineHealthCheckLogTopicService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
	)
	return ve.WithSimpleQuery(m, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeHealthCheckLogTopicAttributes"

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
		// 直接取返回的Result字段
		results, err = ve.ObtainSdkValue("Result", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		// 单个结果，需要包装成列表返回
		if resultMap, ok := results.(map[string]interface{}); ok {
			data = []interface{}{resultMap}
		}
		return data, err
	})
}

// 根据LogTopicId查询绑定的CLB实例
func (s *VolcengineHealthCheckLogTopicService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	parts := strings.Split(id, ":")
	if len(parts) != 2 {
		return data, fmt.Errorf("Invalid health check log topic id: %s ", id)
	}
	req := map[string]interface{}{
		"LogTopicId": parts[0],
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
		return data, fmt.Errorf("health_check_log_topic %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineHealthCheckLogTopicService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				d          map[string]interface{}
				failStates []string
			)
			failStates = append(failStates, "Failed")
			d, err = s.ReadResource(resourceData, id)
			if err != nil {
				return nil, "", err
			}
			return d, "Available", err
		},
	}
}

// 创建绑定关系 - AttachHealthCheckLogTopic
func (s *VolcengineHealthCheckLogTopicService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "AttachHealthCheckLogTopic",
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"log_topic_id": {
					TargetField: "LogTopicId",
				},
				"load_balancer_id": {
					TargetField: "LoadBalancerId",
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				// 使用LogTopicId和LoadBalancerId作为唯一标识
				id := fmt.Sprintf("%s:%s", d.Get("log_topic_id").(string), d.Get("load_balancer_id").(string))
				d.SetId(id)
				return nil
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("load_balancer_id").(string)
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineHealthCheckLogTopicService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

// 更新绑定关系 - 实际上健康检查日志主题绑定不支持更新，这里保留空实现
func (s *VolcengineHealthCheckLogTopicService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{}
	return []ve.Callback{callback}
}

// 删除绑定关系 - DetachHealthCheckLogTopic
func (s *VolcengineHealthCheckLogTopicService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DetachHealthCheckLogTopic",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"LoadBalancerId": resourceData.Get("load_balancer_id").(string),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 5*time.Minute)
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("load_balancer_id").(string)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineHealthCheckLogTopicService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts:  map[string]ve.RequestConvert{},
		NameField:        "log_topic_id",
		IdField:          "log_topic_id",
		CollectField:     "health_check_log_topics",
		ResponseConverts: map[string]ve.ResponseConvert{},
	}
}

func (s *VolcengineHealthCheckLogTopicService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "clb",
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		Action:      actionName,
	}
}
