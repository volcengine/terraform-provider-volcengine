package instance

import (
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineMongoDBInstanceService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewMongoDBInstanceService(c *ve.SdkClient) *VolcengineMongoDBInstanceService {
	return &VolcengineMongoDBInstanceService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineMongoDBInstanceService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineMongoDBInstanceService) readInstanceDetails(id string) (instance interface{}, err error) {
	var (
		resp *map[string]interface{}
	)
	action := "DescribeDBInstanceDetail"
	cond := map[string]interface{}{
		"InstanceId": id,
	}
	logger.Debug(logger.RespFormat, action, cond)
	resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &cond)
	if err != nil {
		return instance, err
	}
	logger.Debug(logger.RespFormat, action, resp)

	instance, err = ve.ObtainSdkValue("Result.DBInstance", *resp) //TODO：文档与实际测试不符，由reflect.Typeof(instance)可知返回只有一个instance，不是数组
	if err != nil {
		return instance, err
	}

	return instance, err
}

func (s *VolcengineMongoDBInstanceService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
	)
	withoutDetail, containWithoutDetail := condition["WithoutDetail"]
	if !containWithoutDetail {
		withoutDetail = false
	}
	return ve.WithPageNumberQuery(condition, "PageSize", "PageNumber", 20, 1, func(m map[string]interface{}) ([]interface{}, error) {
		action := "DescribeDBInstances"
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
		logger.Debug(logger.RespFormat, action, condition, *resp)
		results, err = ve.ObtainSdkValue("Result.DBInstances", *resp)
		if err != nil {
			logger.DebugInfo("ve.ObtainSdkValue return :%v", err)
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		instances, ok := results.([]interface{})
		if !ok {
			return data, fmt.Errorf("DescribeDBInstances responsed instances is not a slice")
		}

		for _, ele := range instances {
			ins := ele.(map[string]interface{})
			instanceId := ins["InstanceId"].(string)

			// do not get detail when refresh status
			if withoutDetail.(bool) {
				data = append(data, ins)
				continue
			}

			detail, err := s.readInstanceDetails(instanceId)
			if err != nil {
				logger.DebugInfo("read instance %s detail failed,err:%v.", instanceId, err)
				data = append(data, ele)
				continue
			}
			ins["ConfigServers"] = detail.(map[string]interface{})["ConfigServers"]
			ins["Nodes"] = detail.(map[string]interface{})["Nodes"]
			ins["Mongos"] = detail.(map[string]interface{})["Mongos"]
			ins["Shards"] = detail.(map[string]interface{})["Shards"]

			logger.DebugInfo("ins:   %v", ins)

			data = append(data, ins)
		}
		return data, nil
	})
}

func (s *VolcengineMongoDBInstanceService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}

	req := map[string]interface{}{
		"InstanceId":    id,
		"withoutDetail": true,
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
		return data, fmt.Errorf("instance %s is not exist", id)
	}
	return data, err
}

func (s *VolcengineMongoDBInstanceService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Delay:      1 * time.Second,
		Pending:    []string{},
		Target:     target,
		Timeout:    timeout,
		MinTimeout: 1 * time.Second,

		Refresh: func() (result interface{}, state string, err error) {
			var (
				instance   map[string]interface{}
				status     interface{}
				failStates []string
			)
			failStates = append(failStates, "CreateFailed", "Failed") //TODO:check fail statues.

			logger.DebugInfo("start refresh :%s", id)
			instance, err = s.ReadResource(resourceData, id)
			if err != nil {
				return nil, "", err
			}
			logger.DebugInfo("Refresh instance status resp: %v", instance)

			status, err = ve.ObtainSdkValue("InstanceStatus", instance)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("instance status error,status %s", status.(string))
				}
			}
			logger.DebugInfo("refresh status:%v", status)
			return instance, status.(string), err
		},
	}
}

func (s *VolcengineMongoDBInstanceService) WithResourceResponseHandlers(instance map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return instance, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineMongoDBInstanceService) getVpcIdAndZoneIdBySubnet(subnetId string) (vpcId, zoneId string, err error) {
	// describe subnet
	req := map[string]interface{}{
		"SubnetIds.1": subnetId,
	}
	action := "DescribeSubnets"
	resp, err := s.Client.VpcClient.DescribeSubnetsCommon(&req)
	if err != nil {
		return "", "", err
	}
	logger.Debug(logger.RespFormat, action, req, *resp)
	results, err := ve.ObtainSdkValue("Result.Subnets", *resp)
	if err != nil {
		return "", "", err
	}
	if results == nil {
		results = []interface{}{}
	}
	subnets, ok := results.([]interface{})
	if !ok {
		return "", "", errors.New("Result.Subnets is not Slice")
	}
	if len(subnets) == 0 {
		return "", "", fmt.Errorf("subnet %s not exist", subnetId)
	}
	vpcId = subnets[0].(map[string]interface{})["VpcId"].(string)
	zoneId = subnets[0].(map[string]interface{})["ZoneId"].(string)
	return vpcId, zoneId, nil
}

func (s *VolcengineMongoDBInstanceService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateDBInstance",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"db_engine": {
					TargetField: "DBEngine",
				},
				"db_engine_version": {
					TargetField: "DBEngineVersion",
				},
				"storage_space_gb": {
					TargetField: "StorageSpaceGB",
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				// describe subnet
				subnetId := d.Get("subnet_id")
				vpcId, zoneId, err := s.getVpcIdAndZoneIdBySubnet(subnetId.(string))
				if err != nil {
					return false, fmt.Errorf("get vpc ID by subnet ID %s failed", subnetId)
				}
				// check custom
				if vpcIdCustom, ok := (*call.SdkParam)["VpcId"]; ok {
					if vpcIdCustom.(string) != vpcId {
						return false, fmt.Errorf("vpc ID does not match")
					}
				}
				if zoneIdCustom, ok := (*call.SdkParam)["ZoneId"]; ok {
					if zoneIdCustom.(string) != zoneId {
						return false, fmt.Errorf("zone ID does not match")
					}
				}
				(*call.SdkParam)["VpcId"] = vpcId
				(*call.SdkParam)["ZoneId"] = zoneId
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam, resp)
				id, _ := ve.ObtainSdkValue("Result.InstanceId", *resp)
				d.SetId(id.(string))
				time.Sleep(time.Second * 5) //如果创建之后立即refresh，DescribeDBInstances会查找不到这个实例直接返回错误..
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Running"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineMongoDBInstanceService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callbacks := make([]ve.Callback, 0)

	if resourceData.HasChange("instance_name") {
		callback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyDBInstanceName",
				ConvertMode: ve.RequestConvertIgnore,
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["InstanceId"] = d.Id()
					(*call.SdkParam)["InstanceNewName"] = d.Get("instance_name")
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
			},
		}
		callbacks = append(callbacks, callback)
	}

	if resourceData.HasChange("instance_type") || resourceData.HasChange("node_spec") ||
		resourceData.HasChange("mongos_node_spec") || resourceData.HasChange("shard_number") ||
		resourceData.HasChange("mongos_node_number") || resourceData.HasChange("storage_space_gb") {
		callback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyDBInstanceSpec",
				ConvertMode: ve.RequestConvertIgnore,
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["InstanceId"] = d.Id()
					(*call.SdkParam)["InstanceType"] = d.Get("instance_type")
					if resourceData.HasChange("node_spec") {
						(*call.SdkParam)["NodeSpec"] = d.Get("node_spec")
					}
					if resourceData.HasChange("mongos_node_spec") {
						(*call.SdkParam)["MongosNodeSpec"] = d.Get("mongos_node_spec")
					}
					if resourceData.HasChange("shard_number") {
						(*call.SdkParam)["ShardNumber"] = d.Get("shard_number")
					}
					if resourceData.HasChange("mongos_node_number") {
						(*call.SdkParam)["MongosNodeNumber"] = d.Get("mongos_node_number")
					}
					if resourceData.HasChange("storage_space_gb") {
						(*call.SdkParam)["StorageSpaceGB"] = d.Get("storage_space_gb")
					}
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
			},
		}
		callbacks = append(callbacks, callback)
	}

	if resourceData.HasChange("charge_type") {
		callback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyDBInstanceChargeType",
				ConvertMode: ve.RequestConvertIgnore,
				ContentType: ve.ContentTypeJson,
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["InstanceIds"] = []interface{}{d.Id()}
					chargeType := d.Get("charge_type")
					if chargeType.(string) != "Prepaid" {
						return false, fmt.Errorf("only supports PostPaid to PrePaid currently")
					}
					(*call.SdkParam)["ChargeType"] = chargeType
					(*call.SdkParam)["PeriodUnit"] = d.Get("period_unit")
					(*call.SdkParam)["Period"] = d.Get("period")
					(*call.SdkParam)["AutoRenew"] = d.Get("auto_renew")
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{"Running"},
					Timeout: resourceData.Timeout(schema.TimeoutCreate),
				},
			},
		}
		callbacks = append(callbacks, callback)
	}

	return callbacks
}

func (s *VolcengineMongoDBInstanceService) RemoveResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteDBInstance",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["InstanceId"] = d.Id()
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

func (s *VolcengineMongoDBInstanceService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"db_engine": {
				TargetField: "DBEngine",
			},
			"db_engine_version": {
				TargetField: "DBEngineVersion",
			},
		},
		IdField:      "InstanceId",
		NameField:    "InstanceName",
		CollectField: "instances",
		ContentType:  ve.ContentTypeJson,
		ResponseConverts: map[string]ve.ResponseConvert{
			"DBEngine": {
				TargetField: "db_engine",
			},
			"DBEngineVersion": {
				TargetField: "db_engine_version",
			},
			"DBEngineVersionStr": {
				TargetField: "db_engine_version_str",
			},
			"TotalMemoryGB": {
				TargetField: "total_memory_gb",
			},
			"TotalvCPU": {
				TargetField: "total_vcpu",
			},
			"UsedMemoryGB": {
				TargetField: "used_memory_gb",
			},
			"UsecvCPU": {
				TargetField: "used_vcpu",
			},
			"TotalStorageGB": {
				TargetField: "total_storage_gb",
			},
			"UsedStorageGB": {
				TargetField: "used_storage_gb",
			},
		},
	}
}

func (s *VolcengineMongoDBInstanceService) ReadResourceId(id string) string {
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
