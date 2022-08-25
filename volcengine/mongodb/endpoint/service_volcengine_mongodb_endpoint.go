package endpoint

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

type VolcengineMongoDBEndpointService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewMongoDBEndpointService(c *ve.SdkClient) *VolcengineMongoDBEndpointService {
	return &VolcengineMongoDBEndpointService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineMongoDBEndpointService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineMongoDBEndpointService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	action := "DescribeDBEndpoint"

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
	results, err = ve.ObtainSdkValue("Result.DBEndpoints", *resp)
	if err != nil {
		return data, err
	}
	if results == nil {
		results = []interface{}{}
	}
	if data, ok = results.([]interface{}); !ok {
		return data, errors.New("Result.DBEndpoints is not Slice")
	}
	return data, err
}

func (s *VolcengineMongoDBEndpointService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	parts := strings.Split(id, ":")
	if len(parts) != 2 {
		return data, fmt.Errorf("the format of import id must be 'endpoint:instanceId'")
	}
	req := map[string]interface{}{
		"InstanceId": parts[1],
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		if data, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("value is not map")
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("endpoint of instance %s is not exist", id)
	}
	return data, err
}

func (s *VolcengineMongoDBEndpointService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (s *VolcengineMongoDBEndpointService) WithResourceResponseHandlers(instance map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return instance, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineMongoDBEndpointService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateDBEndpoint",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"mongos_node_ids": {
					ConvertType: ve.ConvertJsonArray,
				},
				"eip_ids": {
					ConvertType: ve.ConvertJsonArray,
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam, resp)
				id := d.Get("instance_id")
				d.SetId("endpoint:" + id.(string))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineMongoDBEndpointService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineMongoDBEndpointService) RemoveResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteDBEndpoint",
			ConvertMode: ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"mongos_node_ids": {
					ConvertType: ve.ConvertJsonArray,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				// DescribeDBEndpoint : obtain EndpointId
				id := d.Id()
				parts := strings.Split(id, ":")
				if len(parts) != 2 {
					return false, fmt.Errorf("the format of import id must be 'endpoint:instanceId'")
				}
				instanceId := parts[1]
				data, err := s.ReadResource(d, instanceId)
				if err != nil {
					return false, fmt.Errorf("get endpoint id failed")
				}
				endpointId := data["EndpointId"]

				(*call.SdkParam)["InstanceId"] = instanceId
				(*call.SdkParam)["EndpointId"] = endpointId
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineMongoDBEndpointService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"db_engine": {
				TargetField: "DBEngine",
			},
			"db_engine_version": {
				TargetField: "DBEngineVersion",
			},
		},
		IdField:      "EndpointId",
		CollectField: "endpoints",
		ContentType:  ve.ContentTypeJson,
		ResponseConverts: map[string]ve.ResponseConvert{
			"DBAddresses": {
				TargetField: "db_addresses",
			},
			"AddressIP": {
				TargetField: "address_ip",
			},
		},
	}
}

func (s *VolcengineMongoDBEndpointService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "mongodb",
		Action:      actionName,
		Version:     "2022-01-01",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
	}
}
