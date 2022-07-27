package instance

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

type VolcengineESCloudInstanceService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewESCloudInstanceService(c *ve.SdkClient) *VolcengineESCloudInstanceService {
	return &VolcengineESCloudInstanceService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineESCloudInstanceService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineESCloudInstanceService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeInstances"
		// 重新组织 Filter 的格式
		if filter, filterExist := condition["Filters"]; filterExist {
			newFilter := make([]interface{}, 0)
			for k, v := range filter.(map[string]interface{}) {
				newFilter = append(newFilter, map[string]interface{}{
					"Name":   k,
					"Values": v,
				})
			}
			condition["Filters"] = newFilter
		}
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

		// get instance node info
		for index, ele := range data {
			ins := ele.(map[string]interface{})
			con := &map[string]interface{}{
				"InstanceId": ins["InstanceId"],
			}
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo("DescribeInstanceNodes"), con)
			if err != nil {
				return data, err
			}
			respBytes, _ = json.Marshal(resp)
			logger.Debug(logger.RespFormat, "DescribeInstanceNodes", con, string(respBytes))
			results, err = ve.ObtainSdkValue("Result.Nodes", *resp)
			if err != nil {
				return data, err
			}
			if results == nil {
				results = []interface{}{}
			}
			data[index].(map[string]interface{})["Nodes"] = results

			// 插件系统只有在 Running 状态下才存在
			if ins["Status"] != "Running" {
				continue
			}
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo("DescribeInstancePlugins"), con)
			if err != nil {
				return data, err
			}
			respBytes, _ = json.Marshal(resp)
			logger.Debug(logger.RespFormat, "DescribeInstancePlugins", con, string(respBytes))
			results, err = ve.ObtainSdkValue("Result.InstancePlugins", *resp)
			if err != nil {
				return data, err
			}
			if results == nil {
				results = []interface{}{}
			}
			data[index].(map[string]interface{})["Plugins"] = results
		}
		return data, err
	})
}

func (s *VolcengineESCloudInstanceService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"Filters": map[string]interface{}{
			"InstanceId": []string{id},
		},
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
		return data, fmt.Errorf("Instance %s not exist ", id)
	}
	// Fixme: 临时解决方案
	if data["MaintenanceTime"] != "" {
		data["InstanceConfiguration"].(map[string]interface{})["MaintenanceTime"] = data["MaintenanceTime"]
	}
	if data["MaintenanceDay"] != nil {
		data["InstanceConfiguration"].(map[string]interface{})["MaintenanceDay"] = data["MaintenanceDay"]
	}
	if resourceData.Get("instance_configuration.0.admin_password") != "" {
		data["InstanceConfiguration"].(map[string]interface{})["AdminPassword"] = resourceData.Get("instance_configuration.0.admin_password")
	}
	if resourceData.Get("instance_configuration.0.configuration_code") != "" {
		data["InstanceConfiguration"].(map[string]interface{})["ConfigurationCode"] = resourceData.Get("instance_configuration.0.configuration_code")
	}
	if resourceData.Get("instance_configuration.0.node_specs_assigns") != nil {
		data["InstanceConfiguration"].(map[string]interface{})["NodeSpecsAssigns"] = resourceData.Get("instance_configuration.0.node_specs_assigns")
	}
	return data, err
}

func (s *VolcengineESCloudInstanceService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,

		Refresh: func() (result interface{}, state string, err error) {
			var (
				demo       map[string]interface{}
				status     interface{}
				failStates []string
			)
			failStates = append(failStates, "Failed")
			demo, err = s.ReadResource(resourceData, id)
			if err != nil {
				return nil, "", err
			}
			logger.Debug("Refresh ESCloud status resp:%v", "ReadResource", demo)
			status, err = ve.ObtainSdkValue("Status", demo)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("ESCloud instance status error,status %s", status.(string))
				}
			}
			return demo, status.(string), err
		},
	}
}

func (s *VolcengineESCloudInstanceService) WithResourceResponseHandlers(instance map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return instance, map[string]ve.ResponseConvert{
			"InstanceConfiguration": {
				TargetField: "instance_configuration",
			},
			"VPC": {
				TargetField: "vpc",
				Convert: func(i interface{}) interface{} {
					return []interface{}{i}
				},
			},
			"Subnet": {
				TargetField: "subnet",
				Convert: func(i interface{}) interface{} {
					return []interface{}{i}
				},
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineESCloudInstanceService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateInstance",
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"instance_configuration": {
					ConvertType: ve.ConvertJsonObject,
					NextLevelConvert: map[string]ve.RequestConvert{
						"node_specs_assigns": {
							ConvertType: ve.ConvertJsonObjectArray,
							NextLevelConvert: map[string]ve.RequestConvert{
								"type": {
									ConvertType: ve.ConvertJsonObject,
								},
								"number": {
									ConvertType: ve.ConvertJsonObject,
								},
								"resource_spec_name": {
									ConvertType: ve.ConvertJsonObject,
								},
								"storage_spec_name": {
									ConvertType: ve.ConvertJsonObject,
								},
								"storage_size": {
									ConvertType: ve.ConvertJsonObject,
								},
							},
						},
						"subnet": {
							ConvertType: ve.ConvertJsonObject,
						},
						"vpc": {
							ConvertType: ve.ConvertJsonObject,
							TargetField: "VPC",
						},
					},
				},
			},

			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
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

func (s *VolcengineESCloudInstanceService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	if resourceData.HasChange("instance_configuration.0.instance_name") {
		id := resourceData.Id()
		name := resourceData.Get("instance_configuration.0.instance_name")

		logger.DebugInfo("instance_name changed,new_name:%s", name)

		renameCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "RenameInstance",
				ContentType: ve.ContentTypeJson,
				ConvertMode: ve.RequestConvertIgnore,
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["InstanceId"] = id
					(*call.SdkParam)["NewName"] = name
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo("RenameInstance"), call.SdkParam)
				},
			},
		}
		callbacks = append(callbacks, renameCallback)
	}

	if resourceData.HasChange("instance_configuration.0.maintenance_day") || resourceData.HasChange("instance_configuration.0.maintenance_time") {
		id := resourceData.Id()
		maintenanceTime := resourceData.Get("instance_configuration.0.maintenance_time")
		maintenanceDay := resourceData.Get("instance_configuration.0.maintenance_day")

		logger.DebugInfo("maintenance changed:%v", maintenanceTime)
		logger.DebugInfo("maintenance changed:%v", maintenanceDay)

		modifyMaintenanceSettingCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyMaintenanceSetting",
				ContentType: ve.ContentTypeJson,
				ConvertMode: ve.RequestConvertIgnore,
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["InstanceId"] = id
					(*call.SdkParam)["MaintenanceTime"] = maintenanceTime
					(*call.SdkParam)["MaintenanceDay"] = maintenanceDay
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo("ModifyMaintenanceSetting"), call.SdkParam)
				},
			},
		}
		callbacks = append(callbacks, modifyMaintenanceSettingCallback)
	}

	if resourceData.HasChange("instance_configuration.0.admin_password") {
		id := resourceData.Id()
		password := resourceData.Get("instance_configuration.0.admin_password")
		userName := resourceData.Get("instance_configuration.0.instance_name")

		logger.DebugInfo("Modify admin password of instance %s.", id)

		resetAdminPasswdCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ResetAdminPassword",
				ContentType: ve.ContentTypeJson,
				ConvertMode: ve.RequestConvertIgnore,
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["InstanceId"] = id
					(*call.SdkParam)["UserName"] = userName
					(*call.SdkParam)["NewPassword"] = password
					(*call.SdkParam)["Force"] = false
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo("ResetAdminPassword"), call.SdkParam)
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{"Running"},
					Timeout: resourceData.Timeout(schema.TimeoutCreate),
				},
			},
		}
		callbacks = append(callbacks, resetAdminPasswdCallback)
	}

	return callbacks
}

func (s *VolcengineESCloudInstanceService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ReleaseInstance",
			ContentType: ve.ContentTypeJson,
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["InstanceId"] = resourceData.Id()
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo("ReleaseInstance"), call.SdkParam)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineESCloudInstanceService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "Filters.InstanceId",
				ConvertType: ve.ConvertJsonArray,
			},
			"statuses": {
				TargetField: "Filters.Status",
				ConvertType: ve.ConvertJsonArray,
			},
			"charge_types": {
				TargetField: "Filters.ChargeType",
				ConvertType: ve.ConvertJsonArray,
			},
			"names": {
				TargetField: "Filters.InstanceName",
				ConvertType: ve.ConvertJsonArray,
			},
			"versions": {
				TargetField: "Filters.Version",
				ConvertType: ve.ConvertJsonArray,
			},
			"zone_ids": {
				TargetField: "Filters.ZoneId",
				ConvertType: ve.ConvertJsonArray,
			},
		},
		ContentType:  ve.ContentTypeJson,
		IdField:      "InstanceId",
		CollectField: "instances",
		ResponseConverts: map[string]ve.ResponseConvert{
			"InstanceId": {
				TargetField: "id",
				KeepDefault: true,
			},
			"EnableESPublicNetwork": {
				TargetField: "enable_es_public_network",
			},
			"EnableESPrivateNetwork": {
				TargetField: "enable_es_private_network",
			},
			"ESPublicDomain": {
				TargetField: "es_public_domain",
			},
			"ESPrivateDomain": {
				TargetField: "es_private_domain",
			},
			"ESPrivateEndpoint": {
				TargetField: "es_private_endpoint",
			},
			"ESPublicEndpoint": {
				TargetField: "es_public_endpoint",
			},
			"ESInnerEndpoint": {
				TargetField: "es_inner_endpoint",
			},
			"CPU": {
				TargetField: "cpu",
			},
			"VPC": {
				TargetField: "vpc",
			},
		},
	}
}

func (s *VolcengineESCloudInstanceService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "ESCloud",
		Version:     "2018-01-01",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}
