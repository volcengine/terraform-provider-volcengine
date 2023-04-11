package vpc_endpoint_service_resource

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/privatelink/vpc_endpoint_service"
)

type VolcengineService struct {
	Client *ve.SdkClient
}

func NewService(c *ve.SdkClient) *VolcengineService {
	return &VolcengineService{
		Client: c,
	}
}

func (s *VolcengineService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	return nil, err
}

func (s *VolcengineService) describeResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 20, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeVpcEndpointServiceResources"
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
		logger.Debug(logger.RespFormat, action, *resp)
		results, err = ve.ObtainSdkValue("Result.Resources", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.Resources is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}

	ids := strings.Split(id, ":")
	serviceId := ids[0]
	resourceId := ids[1]

	results, err = s.describeResources(map[string]interface{}{
		"ServiceId": serviceId,
	})
	if err != nil {
		return data, err
	}
	if len(results) == 0 {
		return data, fmt.Errorf("Vpc endpoint service resource %s not exist ", id)
	}

	for _, ele := range results {
		if ele.(map[string]interface{})["ResourceId"] == resourceId {
			return map[string]interface{}{
				"ServiceId":  serviceId,
				"ResourceId": resourceId,
			}, nil
		}
	}
	return data, fmt.Errorf("resource does not associate target service. service_id: %s, resource_id: %s", serviceId, resourceId)
}

func (s *VolcengineService) WithResourceResponseHandlers(nodePool map[string]interface{}) []ve.ResourceResponseHandler {
	return nil
}

func (s *VolcengineService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (s *VolcengineService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "AttachResourceToVpcEndpointService",
			ConvertMode: ve.RequestConvertAll,
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam, resp)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				d.SetId(fmt.Sprint((*call.SdkParam)["ServiceId"], ":", (*call.SdkParam)["ResourceId"]))
				return nil
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				vpc_endpoint_service.NewService(s.Client): {
					Target:     []string{"Available"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					ResourceId: resourceData.Get("service_id").(string),
				},
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("service_id").(string)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback
	return callbacks
}

func (s *VolcengineService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	ids := strings.Split(s.ReadResourceId(resourceData.Id()), ":")
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DetachResourceFromVpcEndpointService",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"ServiceId":  ids[0],
				"ResourceId": ids[1],
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam, resp)
				return resp, err
			},
			LockId: func(d *schema.ResourceData) string {
				return d.Get("service_id").(string)
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				vpc_endpoint_service.NewService(s.Client): {
					Target:     []string{"Available"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					ResourceId: resourceData.Get("service_id").(string),
				},
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{}
}

func (s *VolcengineService) ReadResourceId(id string) string {
	return id
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
