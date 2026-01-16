package listener_health

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineListenerHealthService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewListenerHealthService(c *ve.SdkClient) *VolcengineListenerHealthService {
	return &VolcengineListenerHealthService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineListenerHealthService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineListenerHealthService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp           *map[string]interface{}
		unHealthyCount interface{}
		status         interface{}
	)

	results, err := ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) (data []interface{}, err error) {
		action := "DescribeListenerHealth"
		logger.Debug(logger.ReqFormat, action, m)
		resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &m)
		if err != nil {
			return nil, err
		}
		respBytes, _ := json.Marshal(resp)
		logger.Debug(logger.RespFormat, action, m, string(respBytes))

		// 获取后端服务器健康状态结果
		pageResults, err := ve.ObtainSdkValue("Result.Results", *resp)
		if err != nil {
			return nil, err
		}

		// 同时获取汇总信息
		unHealthyCount, _ = ve.ObtainSdkValue("Result.UnHealthyCount", *resp)
		status, _ = ve.ObtainSdkValue("Result.Status", *resp)

		if pageResults == nil {
			return []interface{}{}, nil
		}
		if slice, ok := pageResults.([]interface{}); ok {
			return slice, nil
		}
		return nil, errors.New("Result.Results is not Slice")
	})

	if err != nil {
		return data, err
	}

	// 构造返回结构，包含汇总信息和服务器列表
	healthInfo := map[string]interface{}{
		"UnHealthyCount": unHealthyCount,
		"ListenerStatus": status,
		"Results":        results,
	}

	return []interface{}{healthInfo}, nil
}

func (s *VolcengineListenerHealthService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	return data, nil
}

func (s *VolcengineListenerHealthService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (s *VolcengineListenerHealthService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{}
	return []ve.Callback{callback}
}

func (VolcengineListenerHealthService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineListenerHealthService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{}
	return []ve.Callback{callback}
}

func (s *VolcengineListenerHealthService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{}
	return []ve.Callback{callback}
}

func (s *VolcengineListenerHealthService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{},
		CollectField:    "health_info",
	}
}

func (s *VolcengineListenerHealthService) ReadResourceId(id string) string {
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
