package alb_listener_health

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineAlbListenerHealthService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewAlbListenerHealthService(c *ve.SdkClient) *VolcengineAlbListenerHealthService {
	return &VolcengineAlbListenerHealthService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineAlbListenerHealthService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineAlbListenerHealthService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp *map[string]interface{}
		ok   bool
	)
	action := "DescribeListenerHealth"

	bytes, _ := json.Marshal(m)
	logger.Debug(logger.ReqFormat, action, string(bytes))

	resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &m)
	if err != nil {
		return data, err
	}

	respBytes, _ := json.Marshal(resp)
	logger.Debug(logger.RespFormat, action, m, string(respBytes))

	// Extract listeners from response
	listeners, err := ve.ObtainSdkValue("Result.Listeners", *resp)
	if err != nil {
		return data, err
	}

	if listeners == nil {
		return []interface{}{}, nil
	}

	if data, ok = listeners.([]interface{}); !ok {
		return data, errors.New("Result.Listeners is not Slice")
	}

	return data, err
}

func (s *VolcengineAlbListenerHealthService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	// For data source, we don't need to implement single resource read
	return nil, nil
}

func (s *VolcengineAlbListenerHealthService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (VolcengineAlbListenerHealthService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineAlbListenerHealthService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineAlbListenerHealthService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineAlbListenerHealthService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineAlbListenerHealthService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"listener_ids": {
				TargetField: "ListenerIds",
				ConvertType: ve.ConvertWithN,
			},
			"only_un_healthy": {
				TargetField: "OnlyUnHealthy",
			},
			"project_name": {
				TargetField: "ProjectName",
			},
		},
		IdField:      "ListenerId",
		CollectField: "listeners",
		ResponseConverts: map[string]ve.ResponseConvert{
			"Results": {
				TargetField: "backend_servers",
			},
		},
	}
}

func (s *VolcengineAlbListenerHealthService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "alb",
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		Action:      actionName,
	}
}
