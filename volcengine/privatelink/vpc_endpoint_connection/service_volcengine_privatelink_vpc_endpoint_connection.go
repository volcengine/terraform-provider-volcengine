package vpc_endpoint_connection

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

type VolcenginePrivateLinkVpcEndpointConnection struct {
	Client *ve.SdkClient
}

func (v *VolcenginePrivateLinkVpcEndpointConnection) GetClient() *ve.SdkClient {
	return v.Client
}

func (v *VolcenginePrivateLinkVpcEndpointConnection) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 20, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeVpcEndpointConnections"
		logger.Debug(logger.ReqFormat, action, condition)
		if condition == nil {
			resp, err = v.Client.UniversalClient.DoCall(getUniversalInfo(action), nil)
			if err != nil {
				return data, err
			}
		} else {
			resp, err = v.Client.UniversalClient.DoCall(getUniversalInfo(action), &condition)
			if err != nil {
				return data, err
			}
		}
		logger.Debug(logger.RespFormat, action, condition, *resp)
		results, err = ve.ObtainSdkValue("Result.EndpointConnections", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.EndpointConnections is not Slice")
		}
		return data, err
	})
}

func (v *VolcenginePrivateLinkVpcEndpointConnection) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = v.ReadResourceId(resourceData.Id())
	}
	ids := strings.Split(id, ":")
	if len(ids) != 2 {
		return nil, errors.New("vpc endpoint connection id err")
	}
	req := map[string]interface{}{
		"EndpointId": ids[0],
		"ServiceId":  ids[1],
	}
	results, err = v.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, r := range results {
		if data, ok = r.(map[string]interface{}); !ok {
			return data, errors.New("Value is not map ")
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("vpc endpoint connection %s not exist", id)
	}
	return data, nil
}

func (v *VolcenginePrivateLinkVpcEndpointConnection) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				data       map[string]interface{}
				status     interface{}
				failStates []string
			)
			failStates = append(failStates, "Error")
			data, err = v.ReadResource(resourceData, id)
			if err != nil {
				return nil, "", err
			}
			status, err = ve.ObtainSdkValue("ConnectionStatus", data)
			if err != nil {
				return nil, "", err
			}
			for _, f := range failStates {
				if f == status.(string) {
					return nil, "", fmt.Errorf("Vpc endpoint connection status error, status: %s ", status.(string))
				}
			}
			//注意 返回的第一个参数不能为空 否则会一直等下去
			return data, status.(string), err
		},
	}
}

func (v *VolcenginePrivateLinkVpcEndpointConnection) WithResourceResponseHandlers(m map[string]interface{}) []ve.ResourceResponseHandler {
	return nil
}

func (v *VolcenginePrivateLinkVpcEndpointConnection) CreateResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "EnableVpcEndpointConnection",
			ConvertMode: ve.RequestConvertAll,
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return v.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id := fmt.Sprintf("%s:%s", d.Get("endpoint_id"), d.Get("service_id"))
				d.SetId(id)
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Connected"},
				Timeout: data.Timeout(schema.TimeoutCreate),
			},
		},
	}
	return []ve.Callback{callback}
}

func (v *VolcenginePrivateLinkVpcEndpointConnection) ModifyResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return nil
}

func (v *VolcenginePrivateLinkVpcEndpointConnection) RemoveResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DisableVpcEndpointConnection",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"EndpointId": resourceData.Get("endpoint_id"),
				"ServiceId":  resourceData.Get("service_id"),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return v.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Rejected"},
				Timeout: resourceData.Timeout(schema.TimeoutDelete),
			},
		},
	}
	return []ve.Callback{callback}
}

func (v *VolcenginePrivateLinkVpcEndpointConnection) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		CollectField: "connections",
		ResponseConverts: map[string]ve.ResponseConvert{
			"NetworkInterfaceIP": {
				TargetField: "network_interface_ip",
			},
		},
	}
}

func (v *VolcenginePrivateLinkVpcEndpointConnection) ReadResourceId(s string) string {
	return s
}

func NewVpcEndpointConnectionService(c *ve.SdkClient) *VolcenginePrivateLinkVpcEndpointConnection {
	return &VolcenginePrivateLinkVpcEndpointConnection{
		Client: c,
	}
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "privatelink",
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		Action:      actionName,
		ContentType: ve.Default,
	}
}
