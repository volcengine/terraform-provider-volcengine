package direct_connect_gateway_route

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/direct_connect/direct_connect_gateway"
)

type VolcengineDirectConnectGatewayRouteService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewDirectConnectGatewayRouteService(c *ve.SdkClient) *VolcengineDirectConnectGatewayRouteService {
	return &VolcengineDirectConnectGatewayRouteService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineDirectConnectGatewayRouteService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineDirectConnectGatewayRouteService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeDirectConnectGatewayRoutes"

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
		results, err = ve.ObtainSdkValue("Result.DirectConnectGatewayRoutes", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.DirectConnectGatewayRoutes is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineDirectConnectGatewayRouteService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"DirectConnectGatewayRouteIds.1": id,
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
		return data, fmt.Errorf("direct_connect_gateway_route %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineDirectConnectGatewayRouteService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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
			status, err = ve.ObtainSdkValue("Status", d)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("direct_connect_gateway_route status error, status: %s", status.(string))
				}
			}
			return d, status.(string), err
		},
	}
}

func (s *VolcengineDirectConnectGatewayRouteService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateDirectConnectGatewayRoute",
			ConvertMode: ve.RequestConvertAll,
			Convert:     map[string]ve.RequestConvert{},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("Result.DirectConnectGatewayRouteId", *resp)
				d.SetId(id.(string))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Available", "UnAvailable"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("direct_connect_gateway_id").(string)
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				direct_connect_gateway.NewDirectConnectGatewayService(s.Client): {
					Target:     []string{"Available"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					ResourceId: resourceData.Get("direct_connect_gateway_id").(string),
				},
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineDirectConnectGatewayRouteService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineDirectConnectGatewayRouteService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineDirectConnectGatewayRouteService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteDirectConnectGatewayRoute",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"DirectConnectGatewayRouteId": resourceData.Id(),
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

func (s *VolcengineDirectConnectGatewayRouteService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "DirectConnectGatewayRouteIds",
				ConvertType: ve.ConvertWithN,
			},
		},
		IdField:      "DirectConnectGatewayRouteId",
		CollectField: "direct_connect_gateway_routes",
	}
}

func (s *VolcengineDirectConnectGatewayRouteService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "directconnect",
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		Action:      actionName,
	}
}
