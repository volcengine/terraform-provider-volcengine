package rds_mssql_instance

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/subnet"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineRdsMssqlInstanceService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewRdsMssqlInstanceService(c *ve.SdkClient) *VolcengineRdsMssqlInstanceService {
	return &VolcengineRdsMssqlInstanceService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineRdsMssqlInstanceService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineRdsMssqlInstanceService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
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
		results, err = ve.ObtainSdkValue("Result.InstancesInfo", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.InstancesInfo is not Slice")
		}

		for _, v := range data {
			rdsInstance, ok := v.(map[string]interface{})
			if !ok {
				return data, fmt.Errorf("Instance is not map ")
			}

			detailAction := "DescribeDBInstanceDetail"
			detailReq := map[string]interface{}{
				"InstanceId": rdsInstance["InstanceId"],
			}
			logger.Debug(logger.ReqFormat, detailAction, detailReq)
			detailInfo, err := s.Client.UniversalClient.DoCall(getUniversalInfo(detailAction), &detailReq)
			if err != nil {
				logger.Info("DescribeDBInstanceDetail error:", err)
				continue
			}
			logger.Debug(logger.RespFormat, detailAction, detailReq, &detailInfo)
			basicInfo, err := ve.ObtainSdkValue("Result.BasicInfo", *detailInfo)
			if err != nil {
				logger.Info("ObtainSdkValue Result.BasicInfo error:", err)
				continue
			}
			basicMap, ok := basicInfo.(map[string]interface{})
			if !ok {
				logger.Info("Result.BasicInfo is not map")
				continue
			}
			for k, v := range basicMap {
				rdsInstance[k] = v
			}

			nodeDetail, err := ve.ObtainSdkValue("Result.NodeDetailInfo", *detailInfo)
			if err != nil {
				logger.Info("ObtainSdkValue Result.NodeDetailInfo error:", err)
				continue
			}
			rdsInstance["NodeDetailInfo"] = nodeDetail

			connection, err := ve.ObtainSdkValue("Result.ConnectionInfo", *detailInfo)
			if err != nil {
				logger.Info("ObtainSdkValue Result.ConnectionInfo error:", err)
				continue
			}
			rdsInstance["ConnectionInfo"] = connection

			action := "DescribeDBInstanceParameters"
			req := map[string]interface{}{
				"InstanceId": rdsInstance["InstanceId"],
			}
			logger.Debug(logger.ReqFormat, action, req)
			parameterInfo, err := s.Client.UniversalClient.DoCall(getUniversalInfo(action), &req)
			if err != nil {
				logger.Info("DescribeDBInstanceParameters error:", err)
				continue
			}
			logger.Debug(logger.RespFormat, action, req, &parameterInfo)

			count, err := ve.ObtainSdkValue("Result.ParameterCount", *parameterInfo)
			if err != nil {
				logger.Info("ObtainSdkValue Result.ParameterCount error:", err)
				continue
			}
			rdsInstance["ParameterCount"] = count

			parameters, err := ve.ObtainSdkValue("Result.InstanceParameters", *parameterInfo)
			if err != nil {
				logger.Info("ObtainSdkValue Result.InstanceParameters error:", err)
				continue
			}
			rdsInstance["Parameters"] = parameters
		}

		return data, err
	})
}

func (s *VolcengineRdsMssqlInstanceService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
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
		} else {
			// 处理 subnet_id
			if subnetIds, ok := data["SubnetId"]; ok {
				ids := strings.Split(subnetIds.(string), ";")
				data["SubnetId"] = ids
			}

			// 回填数据
			if _, ok = data["ChargeDetail"]; ok {
				data["ChargeInfo"] = data["ChargeDetail"]
			}
			if _, ok = data["DBEngineVersion"]; ok {
				data["DbEngineVersion"] = data["DBEngineVersion"]
			}
			if fullBackupPeriod, ok := resourceData.GetOk("full_backup_period"); ok {
				data["FullBackupPeriod"] = fullBackupPeriod
			}
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("sqlserver_instance %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineRdsMssqlInstanceService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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
			failStates = append(failStates, "Error")
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
					return nil, "", fmt.Errorf("sqlserver_instance status error, status: %s", status.(string))
				}
			}
			return d, status.(string), err
		},
	}
}

func (s *VolcengineRdsMssqlInstanceService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateDBInstance",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"db_engine_version": {
					TargetField: "DBEngineVersion",
				},
				"charge_info": {
					ConvertType: ve.ConvertJsonObject,
				},
				"tags": {
					ConvertType: ve.ConvertJsonObjectArray,
				},
				"subnet_id": {
					Ignore: true,
				},
				"full_backup_period": {
					Ignore: true,
				},
				"backup_time": {
					Ignore: true,
				},
				"backup_retention_period": {
					Ignore: true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				subnetId := d.Get("subnet_id")
				ids, ok := subnetId.([]interface{})
				if !ok {
					return false, fmt.Errorf("subnet_id is not slice ")
				}
				zoneIds := make([]string, 0)
				subnetIds := make([]string, 0)
				for _, id := range ids {
					resp, err := subnet.NewSubnetService(s.Client).ReadResource(resourceData, id.(string))
					if err != nil {
						return false, err
					}
					(*call.SdkParam)["VpcId"] = resp["VpcId"]
					zoneIds = append(zoneIds, resp["ZoneId"].(string))
					subnetIds = append(subnetIds, id.(string))
				}
				(*call.SdkParam)["ZoneId"] = strings.Join(zoneIds, ";")
				(*call.SdkParam)["SubnetId"] = strings.Join(subnetIds, ";")

				if (*call.SdkParam)["ChargeInfo.0.ChargeType"] == "PrePaid" {
					if (*call.SdkParam)["ChargeInfo.0.Period"] == nil || (*call.SdkParam)["ChargeInfo.0.Period"].(int) < 1 {
						return false, fmt.Errorf("Instance Charge Type is PrePaid. Must set Period more than 1. ")
					}
					(*call.SdkParam)["ChargeInfo.0.PeriodUnit"] = "Month"
				}

				(*call.SdkParam)["DBTimeZone"] = "China Standard Time"
				(*call.SdkParam)["ServerCollation"] = "Chinese_PRC_CI_AS"
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
	callbacks = append(callbacks, callback)

	backTime, timeOk := resourceData.GetOk("backup_time")
	fullBackupPeriod, fullOk := resourceData.GetOk("full_backup_period")
	backPeriod, retentionOk := resourceData.GetOk("backup_retention_period")
	if timeOk || fullOk || retentionOk {
		backupCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyBackupPolicy",
				ContentType: ve.ContentTypeJson,
				ConvertMode: ve.RequestConvertIgnore,
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					err := fmt.Errorf("backup_time, full_backup_period and backup_retention_period are required when set the backup plan. ")
					if retentionOk {
						(*call.SdkParam)["BackupRetentionPeriod"] = backPeriod
					} else {
						return false, err
					}
					if timeOk {
						(*call.SdkParam)["BackupTime"] = backTime
					} else {
						return false, err
					}
					if fullOk {
						var (
							period    string
							periodStr []string
						)
						periodList := fullBackupPeriod.(*schema.Set).List()
						for _, p := range periodList {
							periodStr = append(periodStr, p.(string))
						}
						period = strings.Join(periodStr, ",")
						(*call.SdkParam)["FullBackupPeriod"] = period
					} else {
						return false, err
					}
					(*call.SdkParam)["InstanceId"] = d.Id()
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
		callbacks = append(callbacks, backupCallback)
	}
	return callbacks
}

func (VolcengineRdsMssqlInstanceService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineRdsMssqlInstanceService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback
	if resourceData.HasChange("backup_time") || resourceData.HasChange("full_backup_period") ||
		resourceData.HasChange("backup_retention_period") {
		backupCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyBackupPolicy",
				ContentType: ve.ContentTypeJson,
				ConvertMode: ve.RequestConvertIgnore,
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					err := fmt.Errorf("backup_time, full_backup_period and backup_retention_period are required when set the backup plan. ")
					if backPeriod, ok := d.GetOk("backup_retention_period"); ok {
						(*call.SdkParam)["BackupRetentionPeriod"] = backPeriod
					} else {
						return false, err
					}
					if backTime, ok := d.GetOk("backup_time"); ok {
						(*call.SdkParam)["BackupTime"] = backTime
					} else {
						return false, err
					}
					if fullBackupPeriod, ok := d.GetOk("full_backup_period"); ok {
						var (
							period    string
							periodStr []string
						)
						periodList := fullBackupPeriod.(*schema.Set).List()
						for _, p := range periodList {
							periodStr = append(periodStr, p.(string))
						}
						period = strings.Join(periodStr, ",")
						(*call.SdkParam)["FullBackupPeriod"] = period
					} else {
						return false, err
					}
					(*call.SdkParam)["InstanceId"] = d.Id()
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
		callbacks = append(callbacks, backupCallback)
	}

	// 更新Tags
	callbacks = s.setResourceTags(resourceData, callbacks)

	return callbacks
}

func (s *VolcengineRdsMssqlInstanceService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{}
	return []ve.Callback{callback}
}

func (s *VolcengineRdsMssqlInstanceService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
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
		ContentType:  ve.ContentTypeJson,
		NameField:    "InstanceName",
		IdField:      "InstanceId",
		CollectField: "instances",
		ResponseConverts: map[string]ve.ResponseConvert{
			"InstanceId": {
				TargetField: "id",
				KeepDefault: true,
			},
			"VCPU": {
				TargetField: "v_cpu",
			},
			"NodeIP": {
				TargetField: "node_ip",
			},
			"DBEngineVersion": {
				TargetField: "db_engine_version",
			},
			"DNSVisibility": {
				TargetField: "dns_visibility",
			},
			"IPAddress": {
				TargetField: "ip_address",
			},
		},
	}
}

func (s *VolcengineRdsMssqlInstanceService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "rds_mssql",
		Version:     "2022-01-01",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}

func (s *VolcengineRdsMssqlInstanceService) setResourceTags(resourceData *schema.ResourceData, callbacks []ve.Callback) []ve.Callback {
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

func (s *VolcengineRdsMssqlInstanceService) ProjectTrn() *ve.ProjectTrn {
	return &ve.ProjectTrn{
		ServiceName:          "rds_mssql",
		ResourceType:         "instance",
		ProjectResponseField: "ProjectName",
		ProjectSchemaField:   "project_name",
	}
}
