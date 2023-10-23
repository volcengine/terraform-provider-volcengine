package shared_transit_router_state

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineSharedTransitRouterStateService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewSharedTransitRouterStateService(c *ve.SdkClient) *VolcengineSharedTransitRouterStateService {
	return &VolcengineSharedTransitRouterStateService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineSharedTransitRouterStateService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineSharedTransitRouterStateService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeTransitRouters"
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
		logger.Debug(logger.RespFormat, action, resp)
		results, err = ve.ObtainSdkValue("Result.TransitRouters", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.TransitRouters is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineSharedTransitRouterStateService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	ids := strings.Split(id, ":")
	req := map[string]interface{}{
		"TransitRouterIds.1": ids[1],
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
		if resourceData.Get("action").(string) == "Reject" {
			return data, nil
		}
		return data, fmt.Errorf("TransitRouter %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineSharedTransitRouterStateService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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
				//if ve.ResourceNotFoundError(err) && action == "Reject" {
				//	return d, "Rejected", nil
				//}
				return nil, "", err
			}
			status, err = ve.ObtainSdkValue("GrantStatus", d)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("shared_transit_router_state status error, status: %s", status.(string))
				}
			}
			return d, status.(string), err
		},
	}
}

func (s *VolcengineSharedTransitRouterStateService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callback ve.Callback
	api := ""
	action := resourceData.Get("action").(string)
	targetStatus := []string{""}
	if action == "Accept" {
		api = "AcceptSharedTransitRouter"
		targetStatus = append(targetStatus, "Accepted")
		callback = ve.Callback{
			Call: ve.SdkCall{
				Action:      api,
				ConvertMode: ve.RequestConvertAll,
				Convert: map[string]ve.RequestConvert{
					"action": {
						Ignore: true,
					},
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
					resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
					logger.Debug(logger.RespFormat, call.Action, resp, err)
					return resp, err
				},
				AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
					id := d.Get("transit_router_id").(string)
					d.SetId(fmt.Sprintf("state:%s", id))
					return nil
				},
				Refresh: &ve.StateRefresh{
					Target:  targetStatus,
					Timeout: resourceData.Timeout(schema.TimeoutCreate),
				},
			},
		}
	} else {
		api = "RejectSharedTransitRouter"
		callback = ve.Callback{
			Call: ve.SdkCall{
				Action:      api,
				ConvertMode: ve.RequestConvertAll,
				Convert: map[string]ve.RequestConvert{
					"action": {
						Ignore: true,
					},
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
					resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
					logger.Debug(logger.RespFormat, call.Action, resp, err)
					return resp, err
				},
				AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
					id := d.Get("transit_router_id").(string)
					d.SetId(fmt.Sprintf("state:%s", id))
					return nil
				},
			},
		}
	}
	return []ve.Callback{callback}
}

func (VolcengineSharedTransitRouterStateService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineSharedTransitRouterStateService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callback ve.Callback
	api := ""
	action := resourceData.Get("action").(string)
	targetStatus := []string{""}
	if action == "Accept" {
		api = "AcceptSharedTransitRouter"
		targetStatus = append(targetStatus, "Accepted")
		callback = ve.Callback{
			Call: ve.SdkCall{
				Action:      api,
				ConvertMode: ve.RequestConvertAll,
				Convert: map[string]ve.RequestConvert{
					"action": {
						Ignore: true,
					},
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["TransitRouterId"] = d.Get("transit_router_id")
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
					resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
					logger.Debug(logger.RespFormat, call.Action, resp, err)
					return resp, err
				},
				Refresh: &ve.StateRefresh{
					Target:  targetStatus,
					Timeout: resourceData.Timeout(schema.TimeoutCreate),
				},
			},
		}
	} else {
		api = "RejectSharedTransitRouter"
		callback = ve.Callback{
			Call: ve.SdkCall{
				Action:      api,
				ConvertMode: ve.RequestConvertAll,
				Convert: map[string]ve.RequestConvert{
					"action": {
						Ignore: true,
					},
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["TransitRouterId"] = d.Get("transit_router_id")
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
					resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
					logger.Debug(logger.RespFormat, call.Action, resp, err)
					return resp, err
				},
			},
		}
	}
	return []ve.Callback{callback}
}

func (s *VolcengineSharedTransitRouterStateService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineSharedTransitRouterStateService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineSharedTransitRouterStateService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "transitrouter",
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		Action:      actionName,
	}
}
