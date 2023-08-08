package instance

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineESCloudInstanceService struct {
	Client *ve.SdkClient
}

func NewESCloudInstanceService(c *ve.SdkClient) *VolcengineESCloudInstanceService {
	return &VolcengineESCloudInstanceService{
		Client: c,
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
	if val := data["MaintenanceTime"]; val != "" {
		data["InstanceConfiguration"].(map[string]interface{})["MaintenanceTime"] = val
	}
	if val := data["MaintenanceDay"]; val != nil {
		data["InstanceConfiguration"].(map[string]interface{})["MaintenanceDay"] = val
	}
	if val := resourceData.Get("instance_configuration.0.admin_password"); val != "" {
		data["InstanceConfiguration"].(map[string]interface{})["AdminPassword"] = val
	}
	if val, ok := resourceData.GetOkExists("instance_configuration.0.force_restart_after_scale"); ok {
		data["InstanceConfiguration"].(map[string]interface{})["ForceRestartAfterScale"] = val
	}
	if subnet, ok := data["InstanceConfiguration"].(map[string]interface{})["Subnet"]; ok {
		data["InstanceConfiguration"].(map[string]interface{})["SubnetId"] = subnet.(map[string]interface{})["SubnetId"]
	}

	if nodeAssigns, ok := data["InstanceConfiguration"].(map[string]interface{})["NodeSpecsAssigns"]; ok {
		finalAssigns := make([]interface{}, 0)
		for _, nodeAssign := range nodeAssigns.([]interface{}) {
			if nodeType := nodeAssign.(map[string]interface{})["Type"]; nodeType == "Master" || nodeType == "Hot" || nodeType == "Kibana" {
				finalAssigns = append(finalAssigns, nodeAssign)
			}
		}
		data["InstanceConfiguration"].(map[string]interface{})["NodeSpecsAssigns"] = finalAssigns
	}

	// 查询 configuration_code
	action := "DescribeNodeAvailableSpecs"
	con := &map[string]interface{}{
		"InstanceId": id,
	}
	logger.Debug(logger.ReqFormat, action, con)
	resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(action), con)
	if err != nil {
		return data, err
	}
	logger.Debug(logger.RespFormat, action, resp)
	configurationCode, err := ve.ObtainSdkValue("Result.ConfigurationCode", *resp)
	if err != nil {
		return data, err
	}
	if configurationCode == nil {
		configurationCode = ""
	}
	data["InstanceConfiguration"].(map[string]interface{})["ConfigurationCode"] = configurationCode

	// 确保查询后 state 文件里的顺序和用户填写的顺序一致
	assigns := resourceData.Get("instance_configuration.0.node_specs_assigns")
	if assigns != nil && len(assigns.([]interface{})) > 0 {
		data["InstanceConfiguration"].(map[string]interface{})["NodeSpecsAssigns"] = assigns
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
			failStates = append(failStates, "CreateFailed", "Error")
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

func preCheckNodeSpec(call ve.SdkCall) error {
	// check node number
	var enablePureMaster bool
	if enablePureMasterValue, ok := (*call.SdkParam)["InstanceConfiguration.EnablePureMaster"]; ok {
		enablePureMaster = enablePureMasterValue.(bool)
	} else {
		enablePureMaster = false
	}

	/**
	打平拆分
	InstanceConfiguration.NodeSpecsAssigns.1.StorageSpecName:es.volume.essd.pl0
	InstanceConfiguration.NodeSpecsAssigns.1.Type:Master
	*/
	nodeConfigs := map[string]map[string]interface{}{}
	for key, value := range *call.SdkParam {
		if (strings.Contains(key, "InstanceConfiguration.NodeSpecsAssigns")) && (value == "Master" || value == "Hot" || value == "Kibana") {
			if _, exist := nodeConfigs[value.(string)]; exist {
				return fmt.Errorf("repeated node configs: %s", value)
			}
			slices := strings.Split(key, ".")
			prefix := strings.TrimSuffix(key, slices[len(slices)-1])

			var number int
			if v, ok := (*call.SdkParam)[prefix+"Number"]; ok {
				number = v.(int)
			}
			nodeConfigs[value.(string)] = map[string]interface{}{
				"StorageSpecName":  (*call.SdkParam)[prefix+"StorageSpecName"],
				"StorageSize":      (*call.SdkParam)[prefix+"StorageSize"],
				"ResourceSpecName": (*call.SdkParam)[prefix+"ResourceSpecName"],
				"Number":           number,
			}
		}
	}
	if len(nodeConfigs) != 3 {
		return fmt.Errorf(" Master, Hot or Kibana NodeSpecsAssigns should be configured.")
	}

	if enablePureMaster {
		// MasterNodeNumber指定的为专属主节点个数，并且取值固定为3。
		// HotNodeNumber指定独立数据节点个数，取值为1-50。此时MasterNode和HotNode计算存储配置可以不一致。
		if nodeConfigs["Master"]["Number"] != 3 {
			return fmt.Errorf(" Master node number muster be 3 if enable_pure_master is true.")
		}
		if nodeConfigs["Hot"]["Number"].(int) < 1 || nodeConfigs["Hot"]["Number"].(int) > 50 {
			return fmt.Errorf(" Hot node number muster range in 1-50 if enable_pure_master is true.")
		}
	} else {
		// MasterNodeNumber=1，HotNodeNumber必需为0，此时代表创建一个单节点的ES实例。
		// MasterNodeNumber=3，HotNodeNumber可选值为0-47，此时MasterNode和HotNode计算存储配置必须一致。
		if nodeConfigs["Master"]["Number"] == 1 {
			if nodeConfigs["Hot"]["Number"] != 0 {
				return fmt.Errorf(" Hot node number muster 0 if enable_pure_master is false and master node number is 1.")
			}
		} else if nodeConfigs["Master"]["Number"] == 3 {
			if nodeConfigs["Hot"]["Number"].(int) < 0 || nodeConfigs["Hot"]["Number"].(int) > 47 {
				return fmt.Errorf(" Hot node number muster range in 0-47 if enable_pure_master is false and master node number is 3.")
			}

			if nodeConfigs["Master"]["ResourceSpecName"] != nodeConfigs["Hot"]["ResourceSpecName"] ||
				nodeConfigs["Master"]["StorageSpecName"] != nodeConfigs["Hot"]["StorageSpecName"] ||
				nodeConfigs["Master"]["StorageSize"] != nodeConfigs["Hot"]["StorageSize"] {
				return fmt.Errorf(" Hot and Master node spec shoud be same if enable_pure_master is false and master node number is 3.")
			}
		} else {
			return fmt.Errorf(" Master node number muster be 1 or 3 if enable_pure_master is false.")
		}
	}
	return nil
}

func diffCheckNodeSpec(d *schema.ResourceData) error {
	oldVal, newVal := d.GetChange("instance_configuration.0.node_specs_assigns")
	oldNodeConfigs := transListToMap(oldVal.([]interface{}))
	newNodeConfigs := transListToMap(newVal.([]interface{}))
	if !reflect.DeepEqual(oldNodeConfigs["Kibana"], newNodeConfigs["Kibana"]) {
		return fmt.Errorf(" Kibana NodeSpecsAssign should not be modified.")
	}
	return nil
}

func (s *VolcengineESCloudInstanceService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateInstance",
			ContentType: ve.ContentTypeJson,
			ConvertMode: ve.RequestConvertAll,
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
					},
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				var (
					results interface{}
					subnets []interface{}
					vpcs    []interface{}
					ok      bool
				)
				// check specs
				if err := preCheckNodeSpec(call); err != nil {
					return false, err
				}

				// check region
				regionId := *(s.Client.ClbClient.Config.Region)
				if regionCustom, ok := (*call.SdkParam)["InstanceConfiguration.RegionId"]; ok {
					if regionId != regionCustom.(string) {
						return false, fmt.Errorf("region does not match")
					}
				}

				// describe subnet
				subnetId := (*call.SdkParam)["InstanceConfiguration.SubnetId"]
				req := map[string]interface{}{
					"SubnetIds.1": subnetId,
				}
				action := "DescribeSubnets"
				resp, err := s.Client.VpcClient.DescribeSubnetsCommon(&req)
				if err != nil {
					return false, err
				}
				logger.Debug(logger.RespFormat, action, req, *resp)
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
				subnetName := subnets[0].(map[string]interface{})["SubnetName"]
				vpcId := subnets[0].(map[string]interface{})["VpcId"]
				zoneId := subnets[0].(map[string]interface{})["ZoneId"]

				//check zone
				if zoneCustom, ok := (*call.SdkParam)["InstanceConfiguration.ZoneId"]; ok {
					if zoneCustom.(string) != zoneId {
						return false, fmt.Errorf("zone does not match")
					}
				}

				// describe vpc
				req = map[string]interface{}{
					"VpcIds.1": vpcId,
				}
				action = "DescribeVpcs"
				resp, err = s.Client.VpcClient.DescribeVpcsCommon(&req)
				if err != nil {
					return false, err
				}
				logger.Debug(logger.RespFormat, action, req, *resp)
				results, err = ve.ObtainSdkValue("Result.Vpcs", *resp)
				if err != nil {
					return false, err
				}
				if results == nil {
					results = []interface{}{}
				}
				if vpcs, ok = results.([]interface{}); !ok {
					return false, errors.New("Result.Vpcs is not Slice")
				}
				if len(vpcs) == 0 {
					return false, fmt.Errorf("vpc %s not exist", subnetId.(string))
				}
				vpcName := vpcs[0].(map[string]interface{})["VpcName"]
				(*call.SdkParam)["InstanceConfiguration.VPC"] = map[string]interface{}{
					"VpcId":   vpcId,
					"VpcName": vpcName,
				}
				(*call.SdkParam)["InstanceConfiguration.Subnet"] = map[string]interface{}{
					"SubnetId":   subnetId,
					"SubnetName": subnetName,
				}
				(*call.SdkParam)["InstanceConfiguration.RegionId"] = regionId
				(*call.SdkParam)["InstanceConfiguration.ZoneId"] = zoneId
				(*call.SdkParam)["ClientToken"] = uuid.New().String()

				logger.DebugInfo("sdk param:%v", *call.SdkParam)
				return true, nil
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
		maintenanceDay := resourceData.Get("instance_configuration.0.maintenance_day").(*schema.Set).List()

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
		userName := resourceData.Get("instance_configuration.0.admin_user_name")

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
					Timeout: resourceData.Timeout(schema.TimeoutUpdate),
				},
			},
		}
		callbacks = append(callbacks, resetAdminPasswdCallback)
	}

	if resourceData.HasChange("instance_configuration.0.node_specs_assigns") {
		logger.DebugInfo("NodeSpecsAssigns:%v", resourceData.Get("instance_configuration.0.node_specs_assigns"))
		scaleCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ScaleInstance",
				ContentType: ve.ContentTypeJson,
				ConvertMode: ve.RequestConvertInConvert,
				Convert: map[string]ve.RequestConvert{
					"instance_configuration": {
						ConvertType: ve.ConvertJsonObject,
						NextLevelConvert: map[string]ve.RequestConvert{
							"node_specs_assigns": {
								ConvertType: ve.ConvertJsonObjectArray,
							},
							"force_restart_after_scale": {
								Ignore: true,
							},
						},
					},
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					totalNodeNumber := 0
					var nodeSpecsAssigns []map[string]interface{}
					nodeSpecs := d.Get("instance_configuration.0.node_specs_assigns").([]interface{})
					// 不允许修改 Kibana，如果修改了 Kibana 就给用户报错
					err := diffCheckNodeSpec(d)
					if err != nil {
						return false, err
					}
					for i, node := range nodeSpecs {
						if node.(map[string]interface{})["type"].(string) == "Master" {
							if d.HasChange("instance_configuration.0.node_specs_assigns." + strconv.Itoa(i) + ".number") {
								return false, fmt.Errorf("master node number is can not be modified")
							}
						}
						if node.(map[string]interface{})["type"].(string) == "Master" || node.(map[string]interface{})["type"].(string) == "Hot" {
							nodeNumber := d.Get("instance_configuration.0.node_specs_assigns." + strconv.Itoa(i) + ".number")
							totalNodeNumber += nodeNumber.(int)
						}
						// 不能传递 Kibana 的信息
						if node.(map[string]interface{})["type"].(string) == "Kibana" {
							continue
						}
						nodeSpecsAssigns = append(nodeSpecsAssigns, map[string]interface{}{
							"Type":             node.(map[string]interface{})["type"],
							"Number":           node.(map[string]interface{})["number"],
							"ResourceSpecName": node.(map[string]interface{})["resource_spec_name"],
							"StorageSpecName":  node.(map[string]interface{})["storage_spec_name"],
							"StorageSize":      node.(map[string]interface{})["storage_size"],
						})
					}
					if totalNodeNumber == 1 {
						return false, fmt.Errorf("single-node cluster does not allow scale")
					}
					// check node specs
					(*call.SdkParam)["InstanceConfiguration.EnablePureMaster"] = d.Get("instance_configuration.0.enable_pure_master")
					err = preCheckNodeSpec(call)
					if err != nil {
						return false, err
					}

					uid, err := uuid.NewUUID()
					if err != nil {
						return false, fmt.Errorf("generate ClientToken failed ")
					}
					(*call.SdkParam)["ClientToken"] = uid.String()
					(*call.SdkParam)["InstanceId"] = d.Id()
					(*call.SdkParam)["NodeSpecsAssigns"] = nodeSpecsAssigns
					(*call.SdkParam)["ConfigurationCode"] = d.Get("instance_configuration.0.configuration_code")
					(*call.SdkParam)["Force"] = d.Get("instance_configuration.0.force_restart_after_scale")
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					delete(*call.SdkParam, "InstanceConfiguration")
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{"Running"},
					Timeout: resourceData.Timeout(schema.TimeoutUpdate),
				},
			},
		}
		callbacks = append(callbacks, scaleCallback)
	}

	if resourceData.HasChange("instance_configuration.0.project_name") {
		callbacks = s.modifyProject(resourceData, callbacks)
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
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading ESCloud instance on delete %q, %w", d.Id(), callErr))
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

func (s *VolcengineESCloudInstanceService) modifyProject(resourceData *schema.ResourceData, callbacks []ve.Callback) []ve.Callback {
	modifyProjectCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "MoveProjectResource",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["TargetProjectName"] = d.Get("instance_configuration.0.project_name")
				if (*call.SdkParam)["TargetProjectName"] == nil || (*call.SdkParam)["TargetProjectName"] == "" {
					return false, fmt.Errorf("Could set ProjectName to empty ")
				}
				//获取用户ID
				input := map[string]interface{}{
					"ProjectName": (*call.SdkParam)["TargetProjectName"],
				}
				logger.Debug(logger.ReqFormat, "GetProject", input)
				out, err := s.Client.UniversalClient.DoCall(getIamUniversalInfo("GetProject"), &input)
				if err != nil {
					return false, err
				}
				accountId, err := ve.ObtainSdkValue("Result.AccountID", *out)
				if err != nil {
					return false, err
				}
				trnStr := fmt.Sprintf("trn:ESCloud:%s:%d:instance/%s", s.Client.Region, int(accountId.(float64)), resourceData.Id())
				(*call.SdkParam)["ResourceTrn.1"] = trnStr
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getIamUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				var (
					instance = make(map[string]interface{})
					err      error
				)
				// 通过 retry 确保 project_name 已成功修改
				err = resource.Retry(15*time.Minute, func() *resource.RetryError {
					instance, err = s.ReadResource(d, d.Id())
					if err != nil {
						if ve.ResourceNotFoundError(err) {
							return resource.RetryableError(err)
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading ESCloud instance %q", d.Id()))
						}
					}
					projectName, err := ve.ObtainSdkValue("InstanceConfiguration.ProjectName", instance)
					if err != nil {
						return resource.RetryableError(err)
					}
					if projectName.(string) != d.Get("instance_configuration.0.project_name").(string) {
						return resource.RetryableError(fmt.Errorf("ESCloud instance is still in updating project name"))
					}
					return nil
				})
				return err
			},
		},
	}
	callbacks = append(callbacks, modifyProjectCallback)

	return callbacks
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

func getIamUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "iam",
		Version:     "2021-08-01",
		HttpMethod:  ve.GET,
		Action:      actionName,
	}
}
