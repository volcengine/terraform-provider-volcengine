package instance

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

type VolcengineMongoDBInstanceService struct {
	Client *ve.SdkClient
}

func NewMongoDBInstanceService(c *ve.SdkClient) *VolcengineMongoDBInstanceService {
	return &VolcengineMongoDBInstanceService{
		Client: c,
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

	instance, err = ve.ObtainSdkValue("Result.DBInstance", *resp)
	if err != nil {
		return instance, err
	}

	return instance, err
}

func (s *VolcengineMongoDBInstanceService) readSSLDetails(id string) (ssl interface{}, err error) {
	var (
		resp *map[string]interface{}
	)
	action := "DescribeDBInstanceSSL"
	cond := map[string]interface{}{
		"InstanceId": id,
	}
	logger.Debug(logger.RespFormat, action, cond)
	resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &cond)
	if err != nil {
		return ssl, err
	}
	logger.Debug(logger.RespFormat, action, resp)

	return ve.ObtainSdkValue("Result", *resp)
}

func (s *VolcengineMongoDBInstanceService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
	)
	withoutDetail, ok := condition["WithoutDetail"]
	if !ok {
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
			return data, fmt.Errorf("DescribeDBInstances response instances is not a slice")
		}

		for _, ele := range instances {
			ins := ele.(map[string]interface{})
			instanceId, err := ve.ObtainSdkValue("InstanceId", ele)
			if err != nil {
				return data, err
			}
			// do not get detail when refresh status
			if withoutDetail.(bool) {
				data = append(data, ins)
				continue
			}

			detail, err := s.readInstanceDetails(instanceId.(string))
			if err != nil {
				logger.DebugInfo("read instance %s detail failed,err:%v.", instanceId, err)
				data = append(data, ele)
				continue
			}
			ssl, err := s.readSSLDetails(instanceId.(string))
			if err != nil {
				logger.DebugInfo("read instance ssl information of %s failed,err:%v.", instanceId, err)
				data = append(data, ele)
				continue
			}
			ConfigServers, err := ve.ObtainSdkValue("ConfigServers", detail)
			if err != nil {
				return data, err
			}
			Nodes, err := ve.ObtainSdkValue("Nodes", detail)
			if err != nil {
				return data, err
			}
			Mongos, err := ve.ObtainSdkValue("Mongos", detail)
			if err != nil {
				return data, err
			}
			Shards, err := ve.ObtainSdkValue("Shards", detail)
			if err != nil {
				return data, err
			}
			SSLEnable, err := ve.ObtainSdkValue("SSLEnable", ssl)
			if err != nil {
				return data, err
			}
			SSLIsValid, err := ve.ObtainSdkValue("SSLIsValid", ssl)
			if err != nil {
				return data, err
			}
			SSLExpiredTime, err := ve.ObtainSdkValue("SSLExpiredTime", ssl)
			if err != nil {
				return data, err
			}

			ins["ConfigServers"] = ConfigServers
			ins["Nodes"] = Nodes
			ins["Mongos"] = Mongos
			ins["Shards"] = Shards
			ins["SSLEnable"] = SSLEnable
			ins["SSLIsValid"] = SSLIsValid
			ins["SSLExpiredTime"] = SSLExpiredTime
			data = append(data, ins)
		}
		return data, nil
	})
}

func (s *VolcengineMongoDBInstanceService) readResource(resourceData *schema.ResourceData, id string, withoutDetail bool) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}

	req := map[string]interface{}{
		"InstanceId":    id,
		"WithoutDetail": withoutDetail,
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

	if zoneId, ok := data["ZoneId"]; ok {
		zoneIds := strings.Split(zoneId.(string), ",")
		data["ZoneIds"] = zoneIds
	}

	if nodeZoneSet, ok := resourceData.GetOk("node_availability_zone"); ok {
		data["NodeAvailabilityZone"] = nodeZoneSet.(*schema.Set).List()
	}

	if withoutDetail {
		return data, nil
	}
	instanceType, _ := ve.ObtainSdkValue("InstanceType", data)
	if instanceType.(string) == "ReplicaSet" {
		n, err := ve.ObtainSdkValue("Nodes", data)
		if err != nil || n == nil {
			data["NodeNumber"] = 0
		} else {
			nodes := n.([]interface{})
			data["NodeNumber"] = len(nodes)
			data["NodeSpec"] = nodes[0].(map[string]interface{})["NodeSpec"]
			data["StorageSpaceGb"] = nodes[0].(map[string]interface{})["TotalStorageGB"]
		}
	} else if instanceType.(string) == "ShardedCluster" {
		m, err := ve.ObtainSdkValue("Mongos", data)
		if err != nil || m == nil {
			data["MongosNodeNumber"] = 0
		} else {
			mongos := m.([]interface{})
			data["MongosNodeNumber"] = len(mongos)
			data["MongosNodeSpec"] = mongos[0].(map[string]interface{})["NodeSpec"]
		}
		s, err := ve.ObtainSdkValue("Shards", data)
		if err != nil || s == nil {
			data["ShardNumber"] = 0
			data["StorageSpaceGb"] = 0
		} else {
			shards := s.([]interface{})
			data["ShardNumber"] = len(shards)
			if tmp, ok := shards[0].(map[string]interface{})["Nodes"]; ok {
				nodes := tmp.([]interface{})
				data["StorageSpaceGb"] = nodes[0].(map[string]interface{})["TotalStorageGB"]
				data["NodeSpec"] = nodes[0].(map[string]interface{})["NodeSpec"]
				data["NodeNumber"] = len(nodes)
			}
		}
	}
	return data, err
}

func (s *VolcengineMongoDBInstanceService) readResourceWithoutDetail(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	return s.readResource(resourceData, id, true)
}

func (s *VolcengineMongoDBInstanceService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	return s.readResource(resourceData, id, false)
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
			failStates = append(failStates, "CreateFailed", "Failed")

			logger.DebugInfo("start refresh :%s", id)
			instance, err = s.readResourceWithoutDetail(resourceData, id)
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

			// 判断下实例的计费类型
			if chargeType, ok := resourceData.GetOk("charge_type"); ok && chargeType == "Prepaid" {
				dataChargeType, err := ve.ObtainSdkValue("ChargeType", instance)
				if err != nil || dataChargeType != "Prepaid" {
					return nil, "", err
				}
			}

			logger.DebugInfo("refresh status:%v", status)
			return instance, status.(string), err
		},
	}
}

func (s *VolcengineMongoDBInstanceService) WithResourceResponseHandlers(instance map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return instance, map[string]ve.ResponseConvert{
			"DBEngine": {
				TargetField: "db_engine",
			},
			"DBEngineVersion": {
				TargetField: "db_engine_version",
			},
			"DBEngineVersionStr": {
				TargetField: "db_engine_version_str",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineMongoDBInstanceService) getVpcIdAndZoneIdBySubnet(subnetId string) (vpcId, zoneId string, err error) {
	// describe subnet
	req := map[string]interface{}{
		"SubnetIds.1": subnetId,
	}
	action := "DescribeSubnets"
	resp, err := s.Client.UniversalClient.DoCall(getVpcUniversalInfo(action), &req)
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
				// "db_engine": {
				// 	TargetField: "DBEngine",
				// },
				"db_engine_version": {
					TargetField: "DBEngineVersion",
				},
				"storage_space_gb": {
					TargetField: "StorageSpaceGB",
				},
				"config_server_node_spec": {
					TargetField: "ConfigServerNodeSpec",
				},
				"config_server_storage_space_gb": {
					TargetField: "ConfigServerStorageSpaceGB",
				},
				"tags": {
					TargetField: "Tags",
					ConvertType: ve.ConvertJsonObjectArray,
				},
				"node_availability_zone": {
					TargetField: "NodeAvailabilityZone",
					ConvertType: ve.ConvertJsonObjectArray,
				},
				"zone_id": {
					Ignore: true,
				},
				"zone_ids": {
					Ignore: true,
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

				var zoneIdsStr string
				zoneIdsArr, ok := d.Get("zone_ids").([]interface{})
				if !ok {
					return false, fmt.Errorf("zone_ids is not slice")
				}
				if len(zoneIdsArr) > 0 {
					zoneIds := make([]string, 0)
					for _, id := range zoneIdsArr {
						zoneIds = append(zoneIds, id.(string))
					}
					zoneIdsStr = strings.Join(zoneIds, ",")
				} else {
					zoneIdsStr = zoneId
				}

				(*call.SdkParam)["VpcId"] = vpcId
				(*call.SdkParam)["ZoneId"] = zoneIdsStr
				// (*call.SdkParam)["DBEngine"] = "MongoDB"
				// (*call.SdkParam)["DBEngineVersion"] = "MongoDB_4_0"
				// (*call.SdkParam)["NodeNumber"] = 3
				// (*call.SdkParam)["SuperAccountName"] = "root"

				if (*call.SdkParam)["InstanceType"] == "ShardedCluster" {
					if _, ok := (*call.SdkParam)["MongosNodeSpec"]; !ok {
						return false, fmt.Errorf("mongos_node_spec must exist for ShardedCluster")
					}
					if _, ok := (*call.SdkParam)["MongosNodeNumber"]; !ok {
						return false, fmt.Errorf("mongos_node_number must exist for ShardedCluster")
					}
					if _, ok := (*call.SdkParam)["ShardNumber"]; !ok {
						return false, fmt.Errorf("shard_number must exist for ShardedCluster")
					}
				}
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
				time.Sleep(time.Second * 10) //如果创建之后立即refresh，DescribeDBInstances会查找不到这个实例直接返回错误..
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
				AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
					time.Sleep(time.Second * 10) //变更之后立即refresh，实例状态还是Running将立即返回..
					return nil
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{"Running"},
					Timeout: resourceData.Timeout(schema.TimeoutCreate),
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
	if resourceData.HasChange("super_account_password") {
		callback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ResetDBAccount",
				ConvertMode: ve.RequestConvertIgnore,
				ContentType: ve.ContentTypeJson,
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["InstanceId"] = d.Id()
					//暂时写死 当前不支持这个字段 只能是root
					(*call.SdkParam)["AccountName"] = "root"
					(*call.SdkParam)["AccountPassword"] = d.Get("super_account_password")
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

	// 更新Tags
	callbacks = s.setResourceTags(resourceData, callbacks)

	return callbacks
}

func (s *VolcengineMongoDBInstanceService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteDBInstance",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"InstanceId": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 15*time.Minute)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on  reading mongodb on delete %q, %w", d.Id(), callErr))
						}
					}
					_, callErr = call.ExecuteCall(d, client, call)
					if callErr == nil {
						return nil
					}
					return resource.RetryableError(callErr)
				})
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
			"tags": {
				TargetField: "Tags",
				ConvertType: ve.ConvertJsonObjectArray,
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
			"UsedvCPU": {
				TargetField: "used_vcpu",
			},
			"TotalStorageGB": {
				TargetField: "total_storage_gb",
			},
			"UsedStorageGB": {
				TargetField: "used_storage_gb",
			},
			"SSLEnable": {
				TargetField: "ssl_enable",
			},
			"SSLIsValid": {
				TargetField: "ssl_is_valid",
			},
			"SSLExpireTime": {
				TargetField: "ssl_expire_time",
			},
		},
	}
}

func (s *VolcengineMongoDBInstanceService) ReadResourceId(id string) string {
	return id
}

func (s *VolcengineMongoDBInstanceService) setResourceTags(resourceData *schema.ResourceData, callbacks []ve.Callback) []ve.Callback {
	addedTags, removedTags, _, _ := ve.GetSetDifference("tags", resourceData, ve.TagsHash, false)

	removeCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "RemoveTagsFromResource",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if removedTags != nil && len(removedTags.List()) > 0 {
					(*call.SdkParam)["InstanceIds"] = []string{resourceData.Id()}
					(*call.SdkParam)["TagKeys"] = make([]string, 0)
					for _, tag := range removedTags.List() {
						(*call.SdkParam)["TagKeys"] = append((*call.SdkParam)["TagKeys"].([]string), tag.(map[string]interface{})["key"].(string))
					}
					return true, nil
				}
				return false, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	callbacks = append(callbacks, removeCallback)

	addCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "AddTagsToResource",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if addedTags != nil && len(addedTags.List()) > 0 {
					(*call.SdkParam)["InstanceIds"] = []string{resourceData.Id()}
					(*call.SdkParam)["Tags"] = make([]map[string]interface{}, 0)
					for _, tag := range addedTags.List() {
						(*call.SdkParam)["Tags"] = append((*call.SdkParam)["Tags"].([]map[string]interface{}), tag.(map[string]interface{}))
					}
					return true, nil
				}
				return false, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	callbacks = append(callbacks, addCallback)

	return callbacks
}

func (s *VolcengineMongoDBInstanceService) ProjectTrn() *ve.ProjectTrn {
	return &ve.ProjectTrn{
		ServiceName:          "mongodb",
		ResourceType:         "instance",
		ProjectResponseField: "ProjectName",
		ProjectSchemaField:   "project_name",
	}
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

func getVpcUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "vpc",
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		Action:      actionName,
	}
}
