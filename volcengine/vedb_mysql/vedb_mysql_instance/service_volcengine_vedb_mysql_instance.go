package vedb_mysql_instance

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineVedbMysqlInstanceService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewVedbMysqlInstanceService(c *ve.SdkClient) *VolcengineVedbMysqlInstanceService {
	return &VolcengineVedbMysqlInstanceService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineVedbMysqlInstanceService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineVedbMysqlInstanceService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeDBInstances"

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

		results, err = ve.ObtainSdkValue("Result.Instances", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.Instances is not Slice")
		}
		// 拆charge信息出来
		// append detail
		for _, v := range data {
			var (
				detailErr    error
				endpointInfo interface{}
				detailInfo   interface{}
			)
			instance := v.(map[string]interface{})
			charge, ok := instance["ChargeDetail"]
			if !ok {
				continue
			}
			if chargeMap, ok := charge.(map[string]interface{}); ok {
				instance["ChargeType"] = chargeMap["ChargeType"]
				instance["ChargeStatus"] = chargeMap["ChargeStatus"]
				instance["OverdueReclaimTime"] = chargeMap["OverdueReclaimTime"]
				instance["OverdueTime"] = chargeMap["OverdueTime"]
				instance["AutoRenew"] = chargeMap["AutoRenew"]
				instance["ChargeStartTime"] = chargeMap["ChargeStartTime"]
				instance["ChargeEndTime"] = chargeMap["ChargeEndTime"]
			}
			if nodes, ok := instance["Nodes"]; ok {
				if nodeList, ok := nodes.([]interface{}); ok && len(nodeList) > 0 {
					nodeNum := len(nodeList)
					nodeSpec := nodeList[0].(map[string]interface{})["NodeSpec"]
					instance["NodeNumber"] = nodeNum
					instance["NodeSpec"] = nodeSpec
				}
			}

			action = "DescribeDBInstanceDetail"
			req := map[string]interface{}{
				"InstanceId": instance["InstanceId"],
			}
			resp, detailErr = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &req)
			if detailErr != nil {
				logger.Info("DescribeDBInstanceDetail error:", detailErr)
				continue
			}
			respBytes, _ = json.Marshal(resp)
			logger.Debug(logger.RespFormat, action, req, string(respBytes))

			// append endpoint info
			endpointInfo, detailErr = ve.ObtainSdkValue("Result.Endpoints", *resp)
			if detailErr != nil {
				logger.Info("ObtainSdkValue Result.Endpoints error:", detailErr)
				continue
			}
			if infos, ok := endpointInfo.([]interface{}); ok {
				instance["Endpoints"] = infos
			} else {
				// 接口返回nil
				instance["Endpoints"] = []interface{}{}
			}

			detailInfo, detailErr = ve.ObtainSdkValue("Result.InstanceDetail", *resp)
			if detailErr != nil {
				logger.Info("ObtainSdkValue Result.InstanceDetail error:", detailErr)
				continue
			}
			if infos, ok := detailInfo.(map[string]interface{}); ok {
				instance["InstanceDetail"] = infos
			} else {
				// 接口返回nil
				instance["InstanceDetail"] = map[string]interface{}{}
			}
		}

		return data, err
	})
}

func (s *VolcengineVedbMysqlInstanceService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"InstanceId": id,
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
		return data, fmt.Errorf("vedb_mysql_instance %s not exist ", id)
	}
	// 接口有问题，实例running后一段时间还是查不到subnet
	if subnet, ok := data["SubnetId"]; !ok || subnet == nil || subnet.(string) == "" {
		subnetId, ok := resourceData.GetOk("subnet_id")
		if !ok {
			data["SubnetId"] = ""
		} else {
			data["SubnetId"] = subnetId.(string)
		}
	}
	return data, err
}

func (s *VolcengineVedbMysqlInstanceService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      5 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				d          map[string]interface{}
				status     interface{}
				failStates []string
			)
			failStates = append(failStates, "CreateFailed", "Unavailable")
			if err = resource.Retry(20*time.Minute, func() *resource.RetryError {
				d, err = s.ReadResource(resourceData, id)
				if err != nil {
					if ve.ResourceNotFoundError(err) {
						return resource.RetryableError(err)
					} else {
						return resource.NonRetryableError(err)
					}
				}
				return nil
			}); err != nil {
				return nil, "", err
			}
			status, err = ve.ObtainSdkValue("InstanceStatus", d)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("vedb_mysql_instance status error, status: %s", status.(string))
				}
			}
			return d, status.(string), err
		},
	}
}

func (s *VolcengineVedbMysqlInstanceService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateDBInstance",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"db_engine_version": {
					TargetField: "DBEngineVersion",
				},
				"tags": {
					TargetField: "Tags",
					ConvertType: ve.ConvertJsonObjectArray,
				},
				"db_minor_version": {
					TargetField: "DBMinorVersion",
				},
				"db_time_zone": {
					TargetField: "DBTimeZone",
				},
				"pre_paid_storage_in_gb": {
					TargetField: "PrePaidStorageInGB",
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				var (
					subnets []interface{}
					results interface{}
					ok      bool
				)
				// add vpc id
				subnetId := d.Get("subnet_id")
				req := map[string]interface{}{
					"SubnetIds.1": subnetId,
				}
				action := "DescribeSubnets"
				resp, err := s.Client.UniversalClient.DoCall(getVPCUniversalInfo(action), &req)
				if err != nil {
					return false, err
				}
				results, err = ve.ObtainSdkValue("Result.Subnets", *resp)
				if err != nil {
					return false, err
				}
				if results == nil {
					results = []interface{}{}
				}
				if subnets, ok = results.([]interface{}); !ok {
					return false, errors.New("Result.Subnets is not Slice")
				}
				if len(subnets) == 0 {
					return false, fmt.Errorf("subnet %s not exist", subnetId.(string))
				}
				vpcId := subnets[0].(map[string]interface{})["VpcId"]
				zoneId := subnets[0].(map[string]interface{})["ZoneId"]

				(*call.SdkParam)["VpcId"] = vpcId
				(*call.SdkParam)["ZoneIds"] = zoneId

				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("Result.InstanceId", *resp)
				d.SetId(id.(string))
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

func (VolcengineVedbMysqlInstanceService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"DBEngineVersion": {
				TargetField: "db_engine_version",
			},
			"DBMinorVersion": {
				TargetField: "db_minor_version",
			},
			"DBTimeZone": {
				TargetField: "db_time_zone",
			},
			"PrePaidStorageInGB": {
				TargetField: "pre_paid_storage_in_gb",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineVedbMysqlInstanceService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback
	if resourceData.HasChanges("instance_name") {
		nameCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyDBInstanceName",
				ConvertMode: ve.RequestConvertInConvert,
				ContentType: ve.ContentTypeJson,
				Convert: map[string]ve.RequestConvert{
					"instance_name": {
						TargetField: "InstanceNewName",
					},
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["InstanceId"] = d.Id()
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
					logger.Debug(logger.RespFormat, call.Action, resp, err)
					return resp, err
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{"Running"},
					Timeout: resourceData.Timeout(schema.TimeoutCreate),
				},
			},
		}
		callbacks = append(callbacks, nameCallback)
	}

	if resourceData.HasChanges("node_spec", "node_number", "pre_paid_storage_in_gb") {
		specCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyDBInstanceSpec",
				ConvertMode: ve.RequestConvertInConvert,
				ContentType: ve.ContentTypeJson,
				Convert: map[string]ve.RequestConvert{
					"node_spec": {
						TargetField: "NodeSpec",
						ForceGet:    true,
					},
					"node_number": {
						TargetField: "NodeNumber",
						ForceGet:    true,
					},
					"pre_paid_storage_in_gb": {
						TargetField: "PrePaidStorageInGB",
					},
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["InstanceId"] = d.Id()
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
					logger.Debug(logger.RespFormat, call.Action, resp, err)
					return resp, err
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{"Running"},
					Timeout: resourceData.Timeout(schema.TimeoutCreate),
				},
			},
		}
		callbacks = append(callbacks, specCallback)
	}

	if resourceData.HasChanges("charge_type", "storage_charge_type", "auto_renew",
		"period_unit", "period") {
		chargeCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyDBInstanceChargeType",
				ConvertMode: ve.RequestConvertInConvert,
				ContentType: ve.ContentTypeJson,
				Convert: map[string]ve.RequestConvert{
					"charge_type": {
						TargetField: "ChargeType",
						ForceGet:    true,
					},
					"storage_charge_type": {
						TargetField: "StorageChargeType",
					},
					"auto_renew": {
						TargetField: "AutoRenew",
						ForceGet:    true,
					},
					"period_unit": {
						TargetField: "PeriodUnit",
						ForceGet:    true,
					},
					"period": {
						TargetField: "Period",
						ForceGet:    true,
					},
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["InstanceId"] = d.Id()
					// 不支持包年包月转按量计费，让API报错吧，不管了
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
					logger.Debug(logger.RespFormat, call.Action, resp, err)
					return resp, err
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{"Running"},
					Timeout: resourceData.Timeout(schema.TimeoutCreate),
				},
			},
		}
		callbacks = append(callbacks, chargeCallback)
	}

	// tag
	callbacks = s.setResourceTags(resourceData, callbacks)
	return callbacks
}

func (s *VolcengineVedbMysqlInstanceService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteDBInstance",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"InstanceId": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 10*time.Minute)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading vedb mysql instance on delete %q, %w", d.Id(), callErr))
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

func (s *VolcengineVedbMysqlInstanceService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"db_engine_version": {
				TargetField: "DBEngineVersion",
			},
			"tags": {
				TargetField: "TagFilters",
				ConvertType: ve.ConvertJsonObjectArray,
			},
		},
		NameField:    "InstanceName",
		IdField:      "InstanceId",
		CollectField: "instances",
		ContentType:  ve.ContentTypeJson,
		ResponseConverts: map[string]ve.ResponseConvert{
			"InstanceId": {
				TargetField: "id",
				KeepDefault: true,
			},
			"DBEngineVersion": {
				TargetField: "db_engine_version",
			},
			"PrePaidStorageInGB": {
				TargetField: "pre_paid_storage_in_gb",
			},
			"StorageUsedGiB": {
				TargetField: "storage_used_gib",
			},
			"vCPU": {
				TargetField: "v_cpu",
			},
		},
	}
}

func (s *VolcengineVedbMysqlInstanceService) ReadResourceId(id string) string {
	return id
}

func getVPCUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "vpc",
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		Action:      actionName,
	}
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "vedbm",
		Version:     "2022-01-01",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}

func (s *VolcengineVedbMysqlInstanceService) ProjectTrn() *ve.ProjectTrn {
	return &ve.ProjectTrn{
		ServiceName:          "vedbm",
		ResourceType:         "instance",
		ProjectResponseField: "ProjectName",
		ProjectSchemaField:   "project_name",
	}
}

func (s *VolcengineVedbMysqlInstanceService) UnsubscribeInfo(resourceData *schema.ResourceData, resource *schema.Resource) (*ve.UnsubscribeInfo, error) {
	info := ve.UnsubscribeInfo{
		InstanceId: s.ReadResourceId(resourceData.Id()),
	}
	if resourceData.Get("charge_type").(string) == "PrePaid" {
		info.NeedUnsubscribe = true
		info.Products = []string{"veDB for MySQL"}
	}
	return &info, nil
}

func (s *VolcengineVedbMysqlInstanceService) setResourceTags(resourceData *schema.ResourceData, callbacks []ve.Callback) []ve.Callback {
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
