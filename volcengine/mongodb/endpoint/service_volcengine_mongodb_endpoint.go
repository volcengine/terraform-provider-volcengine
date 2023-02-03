package endpoint

import (
	"errors"
	"fmt"
	"time"

	mongodbInstance "github.com/volcengine/terraform-provider-volcengine/volcengine/mongodb/instance"

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

func (s *VolcengineMongoDBEndpointService) GetEndpoint(instanceId, endpointId, networkType, objectId string) (endpoint map[string]interface{}, err error) {
	req := map[string]interface{}{
		"InstanceId": instanceId,
	}
	results, err := s.ReadResources(req)
	if err != nil {
		return nil, err
	}
	for _, v := range results {
		dbEndpoint, ok := v.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("dbEndpoint value is not map")
		}
		if endpointId != "" && endpointId == dbEndpoint["EndpointId"].(string) {
			return dbEndpoint, nil
		}
		if networkType == dbEndpoint["NetworkType"].(string) { // private or public
			if objectId == "" || (objectId != "" && objectId == dbEndpoint["ObjectId"].(string)) { //endpointType is ReplicaSet if objectId is empty
				return dbEndpoint, nil
			}
		}
	}
	return nil, fmt.Errorf("the endpoint not found")
}

func (s *VolcengineMongoDBEndpointService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
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
	instanceId := resourceData.Get("instance_id").(string)
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateDBEndpoint",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"eip_ids": {
					ConvertType: ve.ConvertJsonArray,
				},
				"mongos_node_ids": {
					ConvertType: ve.ConvertJsonArray,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				networkType := d.Get("network_type")
				eipIds := d.Get("eip_ids")
				if networkType != nil && networkType.(string) == "Public" {
					if eipIds == nil {
						return false, fmt.Errorf("eip_ids is required when network_type is 'Public'")
					}
				}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				mongodbInstance.NewMongoDBInstanceService(s.Client): {
					Target:     []string{"Running"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					ResourceId: instanceId,
				},
			},
			LockId: func(d *schema.ResourceData) string {
				logger.Debug("lock instance id:%s", instanceId, "")
				return instanceId
			},
		},
	}
	obtainEndpointIdCallback := ve.Callback{
		Call: ve.SdkCall{
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				networkType := "Private"
				if a, ok := d.GetOkExists("network_type"); ok {
					networkType = a.(string)
				}
				objectId := ""
				if a, ok := d.GetOkExists("object_id"); ok {
					objectId = a.(string)
				}
				endpoint, err := s.GetEndpoint(instanceId, "", networkType, objectId)
				if err != nil {
					return nil, err
				}
				endpointId := endpoint["EndpointId"].(string)
				d.Set("endpoint_id", endpointId)
				d.SetId(fmt.Sprintf("%s:%s", instanceId, endpointId))
				return nil, nil
			},
		},
	}

	return []ve.Callback{callback, obtainEndpointIdCallback}
}

func (s *VolcengineMongoDBEndpointService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (s *VolcengineMongoDBEndpointService) RemoveResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	instanceId := resourceData.Get("instance_id").(string)
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteDBEndpoint",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				// DescribeDBEndpoint : obtain EndpointId
				(*call.SdkParam)["InstanceId"] = instanceId
				(*call.SdkParam)["EndpointId"] = d.Get("endpoint_id")
				if d.Get("mongos_node_id") != nil {
					(*call.SdkParam)["MongosNodeIds"] = []interface{}{d.Get("mongos_node_id")}
				}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				mongodbInstance.NewMongoDBInstanceService(s.Client): {
					Target:     []string{"Running"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					ResourceId: instanceId,
				},
			},
			LockId: func(d *schema.ResourceData) string {
				return instanceId
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
