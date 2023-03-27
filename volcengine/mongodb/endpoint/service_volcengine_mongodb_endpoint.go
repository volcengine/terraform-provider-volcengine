package endpoint

import (
	"errors"
	"fmt"
	"strings"
	"time"

	mongodbInstance "github.com/volcengine/terraform-provider-volcengine/volcengine/mongodb/instance"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineMongoDBEndpointService struct {
	Client *ve.SdkClient
}

func NewMongoDBEndpointService(c *ve.SdkClient) *VolcengineMongoDBEndpointService {
	return &VolcengineMongoDBEndpointService{
		Client: c,
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
		instanceId   = ""
		endpointId   = ""
		objectId     = ""
		tempObjectId = ""
		networkType  = ""
	)

	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	parts := strings.Split(id, ":")
	if len(parts) != 2 {
		return data, fmt.Errorf("format of mongodb endpoint resource id is invalid,%s", id)
	}
	instanceId = parts[0]
	endpointId = parts[1]

	req := map[string]interface{}{
		"InstanceId": instanceId,
	}
	results, err := s.ReadResources(req)
	if err != nil {
		return nil, err
	}

	var targetEndpoint map[string]interface{}
	if a, ok := resourceData.GetOkExists("network_type"); ok {
		networkType = a.(string)
	}
	if a, ok := resourceData.GetOkExists("object_id"); ok {
		objectId = a.(string)
	}
	for _, v := range results {
		dbEndpoint, ok := v.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("dbEndpoint value is not map")
		}
		eId, err := ve.ObtainSdkValue("EndpointId", dbEndpoint)
		if err != nil {
			return nil, err
		}
		logger.DebugInfo("---- endpointId:%s,eid:%s", endpointId, eId)
		if endpointId != "" { // check by EndpointId
			if endpointId == eId.(string) {
				logger.DebugInfo("get endpoint of endpointId:%s", endpointId)
				targetEndpoint = dbEndpoint
				break
			}
		} else { // check by NetworkType and ObjectId
			nType, err := ve.ObtainSdkValue("NetworkType", dbEndpoint)
			if err != nil {
				return data, err
			}
			oId, err := ve.ObtainSdkValue("ObjectId", dbEndpoint)
			if err != nil {
				return data, err
			}
			if oId != nil {
				tempObjectId = oId.(string)
			}
			if networkType == nType.(string) { // Private or Public
				if objectId == "" || objectId == tempObjectId { //endpointType is ReplicaSet if objectId is empty
					logger.DebugInfo("get mongodb endpoint by  network type and object id %s", networkType, objectId)
					targetEndpoint = dbEndpoint
					break
				}
			}
		}
	}

	if targetEndpoint == nil {
		return data, fmt.Errorf("mongodb endpoint not found")
	}

	nodeIds := make([]string, 0)
	eipIds := make([]string, 0)
	addresses, err := ve.ObtainSdkValue("DBAddresses", targetEndpoint)
	if err != nil {
		return data, err
	}
	endpointType, err := ve.ObtainSdkValue("EndpointType", targetEndpoint)
	if err != nil {
		return data, err
	}
	nType, err := ve.ObtainSdkValue("NetworkType", targetEndpoint)
	if err != nil {
		return data, err
	}
	for _, address := range addresses.([]interface{}) {
		logger.DebugInfo("address %v :", address)
		if nodeId, ok := address.(map[string]interface{})["NodeId"]; ok && nodeId.(string) != "" &&
			endpointType == "Mongos" && nType == "Public" {
			nodeIds = append(nodeIds, nodeId.(string))
		}
		if eipId, ok := address.(map[string]interface{})["EipId"]; ok && eipId.(string) != "" {
			eipIds = append(eipIds, eipId.(string))
		}
	}
	targetEndpoint["MongosNodeIds"] = nodeIds
	targetEndpoint["EipIds"] = eipIds

	return targetEndpoint, err
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
				// 在 LockId 执行后再进行已有 Endpoint 信息的查询
				endpoint, err := s.ReadResource(d, fmt.Sprintf("%s:", instanceId))
				if err != nil && !strings.Contains(err.Error(), "mongodb endpoint not found") {
					return nil, err
				} else if len(endpoint) != 0 {
					return nil, fmt.Errorf("the instance already contains this endpoint, and duplicate creation is not allowed")
				}

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
				endpoint, err := s.ReadResource(d, fmt.Sprintf("%s:", instanceId))
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
			ConvertMode: ve.RequestConvertInConvert,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"mongos_node_ids": {
					ConvertType: ve.ConvertJsonArray,
					ForceGet:    true,
				},
			},
			SdkParam: &map[string]interface{}{
				"InstanceId": instanceId,
				"EndpointId": resourceData.Get("endpoint_id").(string),
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
