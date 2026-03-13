package nlb_listener_health

import (
	"encoding/json"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineNlbListenerHealthService struct {
	Client *ve.SdkClient
}

func NewNlbListenerHealthService(c *ve.SdkClient) *VolcengineNlbListenerHealthService {
	return &VolcengineNlbListenerHealthService{
		Client: c,
	}
}

func (s *VolcengineNlbListenerHealthService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineNlbListenerHealthService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	// The API DescribeNLBListenerHealth returns health info for a specific listener.
	// It's not a standard list query, so we handle it manually.
	action := "DescribeNLBListenerHealth"
	logger.Debug(logger.ReqFormat, action, m)
	resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(action), &m)
	if err != nil {
		return nil, err
	}
	respBytes, _ := json.Marshal(resp)
	logger.Debug(logger.RespFormat, action, m, string(respBytes))

	if resp != nil {
		if result, ok := (*resp)["Result"]; ok {
			data = append(data, result)
		}
	}
	return data, nil
}

func (s *VolcengineNlbListenerHealthService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	return nil, nil
}

func (s *VolcengineNlbListenerHealthService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineNlbListenerHealthService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return nil
}

func (VolcengineNlbListenerHealthService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	return nil
}

func (s *VolcengineNlbListenerHealthService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return nil
}

func (s *VolcengineNlbListenerHealthService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return nil
}

func (s *VolcengineNlbListenerHealthService) DatasourceResources(d *schema.ResourceData, r *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"listener_id": {
				TargetField: "ListenerId",
			},
		},
		IdField:      "ListenerId",
		CollectField: "listener_healths",
		ResponseConverts: map[string]ve.ResponseConvert{
			"ListenerId": {
				TargetField: "listener_id",
			},
			"ServerGroupId": {
				TargetField: "server_group_id",
			},
			"HealthyCount": {
				TargetField: "healthy_count",
			},
			"UnhealthyCount": {
				TargetField: "unhealthy_count",
			},
			"Status": {
				TargetField: "status",
			},
			"Results": {
				TargetField: "results",
				Convert:     transListenerHealthResultsToResponse,
			},
		},
	}
}

func transListenerHealthResultsToResponse(i interface{}) interface{} {
	if i == nil {
		return nil
	}
	list, ok := i.([]interface{})
	if !ok {
		return i
	}
	var res []interface{}
	for _, item := range list {
		m, ok := item.(map[string]interface{})
		if !ok {
			continue
		}
		newMap := make(map[string]interface{})
		if v, ok := m["ServerId"]; ok {
			newMap["server_id"] = v
		}
		if v, ok := m["ServerType"]; ok {
			newMap["server_type"] = v
		}
		if v, ok := m["InstanceId"]; ok {
			newMap["instance_id"] = v
		}
		if v, ok := m["ZoneId"]; ok {
			newMap["zone_id"] = v
		}
		if v, ok := m["Ip"]; ok {
			newMap["ip"] = v
		}
		if v, ok := m["Port"]; ok {
			newMap["port"] = v
		}
		if v, ok := m["Status"]; ok {
			newMap["status"] = v
		}
		res = append(res, newMap)
	}
	return res
}

func (s *VolcengineNlbListenerHealthService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "clb",
		Action:      actionName,
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		RegionType:  ve.Regional,
	}
}
