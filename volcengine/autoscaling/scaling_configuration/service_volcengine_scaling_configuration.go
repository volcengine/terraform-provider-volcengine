package scaling_configuration

import (
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineScalingConfigurationService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewScalingConfigurationService(c *ve.SdkClient) *VolcengineScalingConfigurationService {
	return &VolcengineScalingConfigurationService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineScalingConfigurationService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineScalingConfigurationService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 20, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		universalClient := s.Client.UniversalClient
		action := "DescribeScalingConfigurations"
		logger.Debug(logger.ReqFormat, action, condition)
		if condition == nil {
			resp, err = universalClient.DoCall(getUniversalInfo(action), nil)
			if err != nil {
				return data, err
			}
		} else {
			resp, err = universalClient.DoCall(getUniversalInfo(action), &condition)
			if err != nil {
				return data, err
			}
		}
		logger.Debug(logger.RespFormat, action, condition, resp)
		results, err = ve.ObtainSdkValue("Result.ScalingConfigurations", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.ScalingConfigurations is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineScalingConfigurationService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"ScalingConfigurationIds.1": id,
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
		return data, fmt.Errorf("ScalingConfiguration %s not exist ", id)
	}

	// 查看伸缩组状态
	resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo("DescribeScalingGroups"),
		&map[string]interface{}{"ScalingGroupIds.1": data["ScalingGroupId"]})
	if err != nil {
		return data, fmt.Errorf("describe scaling group err: %s", err.Error())
	}
	groups, err := ve.ObtainSdkValue("Result.ScalingGroups", *resp)
	if groups == nil || len(groups.([]interface{})) == 0 {
		return data, fmt.Errorf("scaling group %s not found", data["ScalingGroupId"])
	}
	data["ScalingGroupStatus"] = groups.([]interface{})[0].(map[string]interface{})["LifecycleState"].(string)
	data["ActiveScalingConfigurationId"] = groups.([]interface{})[0].(map[string]interface{})["ActiveScalingConfigurationId"].(string)

	return data, err
}

func (s *VolcengineScalingConfigurationService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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
			failStates = append(failStates, "Error")
			demo, err = s.ReadResource(resourceData, id)
			if err != nil {
				return nil, "", err
			}
			status, err = ve.ObtainSdkValue("LifecycleState", demo)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("ScalingConfiguration LifecycleState error, status:%s", status.(string))
				}
			}
			//注意 返回的第一个参数不能为空 否则会一直等下去
			return demo, status.(string), err
		},
	}

}

func (VolcengineScalingConfigurationService) WithResourceResponseHandlers(scalingConfiguration map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		scalingConfiguration["active"] = scalingConfiguration["LifecycleState"].(string) == "Active"
		scalingConfiguration["enable"] = !(scalingConfiguration["ScalingGroupStatus"].(string) == "InActive" ||
			scalingConfiguration["ScalingGroupStatus"].(string) == "Deleting")
		scalingConfiguration["substitute"] = scalingConfiguration["ActiveScalingConfigurationId"]
		return scalingConfiguration, nil, nil
	}

	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineScalingConfigurationService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	createConfigCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateScalingConfiguration",
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"volumes": {
					ConvertType: ve.ConvertListN,
					ForceGet:    true,
				},
				"security_group_ids": {
					ConvertType: ve.ConvertWithN,
				},
				"instance_types": {
					ConvertType: ve.ConvertWithN,
				},
				"eip_bandwidth": {
					TargetField: "Eip.Bandwidth",
				},
				"eip_isp": {
					TargetField: "Eip.ISP",
				},
				"eip_billing_type": {
					TargetField: "Eip.BillingType",
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				//注意 获取内容 这个地方不能是指针 需要转一次
				id, _ := ve.ObtainSdkValue("Result.ScalingConfigurationId", *resp)
				d.SetId(id.(string))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"InActive", "Active"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	callbacks = append(callbacks, createConfigCallback)

	// 启用伸缩组
	if resourceData.Get("active") != nil && resourceData.Get("active").(bool) {
		callbacks = append(callbacks, ve.Callback{
			Call: ve.SdkCall{
				Action:      "EnableScalingConfiguration",
				ConvertMode: ve.RequestConvertInConvert,
				Convert: map[string]ve.RequestConvert{
					"scaling_group_id": {
						ConvertType: ve.ConvertDefault,
					},
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["ScalingConfigurationId"] = d.Id()
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{"Active"},
					Timeout: resourceData.Timeout(schema.TimeoutCreate),
				},
			},
		})
	}

	// 使能伸缩组
	if resourceData.Get("enable") != nil {
		callbacks = append(callbacks, s.enableOrDisableGroupCallback(resourceData))
	}

	return callbacks

}

func (s *VolcengineScalingConfigurationService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	// 修改伸缩配置
	modifyConfigurationCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyScalingConfiguration",
			ConvertMode: ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"scaling_configuration_name": {
					ConvertType: ve.ConvertDefault,
				},
				"image_id": {
					ConvertType: ve.ConvertDefault,
				},
				"instance_types": {
					ConvertType: ve.ConvertWithN,
				},
				"instance_name": {
					ConvertType: ve.ConvertDefault,
				},
				"instance_description": {
					ConvertType: ve.ConvertDefault,
				},
				"host_name": {
					ConvertType: ve.ConvertDefault,
				},
				"password": {
					ConvertType: ve.ConvertDefault,
				},
				"key_pair_name": {
					ConvertType: ve.ConvertDefault,
				},
				"key_pair_id": {
					ConvertType: ve.ConvertDefault,
				},
				"security_enhancement_strategy": {
					ConvertType: ve.ConvertDefault,
				},
				"user_data": {
					ConvertType: ve.ConvertDefault,
				},
				"volumes": {
					ConvertType: ve.ConvertListN,
				},
				"security_group_ids": {
					ConvertType: ve.ConvertWithN,
				},
				"eip_bandwidth": {
					TargetField: "Eip.Bandwidth",
				},
				"eip_isp": {
					TargetField: "Eip.ISP",
				},
				"eip_billing_type": {
					TargetField: "Eip.BillingType",
				},
			},
			RequestIdField: "ScalingConfigurationId",
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if len(*call.SdkParam) < 2 {
					return false, nil
				}
				if d.HasChange("eip_bandwidth") || d.HasChange("eip_isp") || d.HasChange("eip_billing_type") {
					(*call.SdkParam)["Eip.Bandwidth"] = d.Get("eip_bandwidth")
					(*call.SdkParam)["Eip.ISP"] = d.Get("eip_isp")
					(*call.SdkParam)["Eip.BillingType"] = d.Get("eip_billing_type")
				}
				if d.HasChange("volumes") {
					for i, ele := range d.Get("volumes").([]interface{}) {
						volume := ele.(map[string]interface{})
						(*call.SdkParam)[fmt.Sprintf("Volumes.%d.DeleteWithInstance", i+1)] = volume["delete_with_instance"]
						(*call.SdkParam)[fmt.Sprintf("Volumes.%d.Size", i+1)] = volume["size"]
						(*call.SdkParam)[fmt.Sprintf("Volumes.%d.VolumeType", i+1)] = volume["volume_type"]
					}
				}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	callbacks = append(callbacks, modifyConfigurationCallback)

	// 使能伸缩配置
	if resourceData.HasChange("active") && resourceData.Get("active").(bool) { // 使用当前伸缩配置
		callbacks = append(callbacks, s.enableConfigurationCallback(resourceData.Id(), resourceData))
	} else if resourceData.Get("active") != nil &&
		!resourceData.Get("active").(bool) &&
		(resourceData.HasChange("active") || resourceData.HasChange("substitute")) { // 使用其他伸缩配置
		substituteId := resourceData.Get("substitute")
		if substituteId == nil || len(substituteId.(string)) == 0 || substituteId.(string) == resourceData.Id() {
			callbacks = append(callbacks, ve.Callback{
				Call: ve.SdkCall{
					ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
						return nil, errors.New("scaling group need a substitute configuration rather than the current scaling configuration")
					},
				},
			})
		} else {
			callbacks = append(callbacks, s.enableConfigurationCallback(substituteId.(string), resourceData))
		}
	}

	// 使能伸缩组
	if resourceData.HasChange("enable") {
		callbacks = append(callbacks, s.enableOrDisableGroupCallback(resourceData))
	}

	return callbacks
}

func (s *VolcengineScalingConfigurationService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteScalingConfiguration",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"ScalingConfigurationId": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading ScalingConfiguration on delete %q, %w", d.Id(), callErr))
						}
					}
					_, callErr = call.ExecuteCall(d, client, call)
					if callErr == nil {
						return nil
					}
					return resource.NonRetryableError(callErr)
				})
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineScalingConfigurationService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "ScalingConfigurationIds",
				ConvertType: ve.ConvertWithN,
			},
			"scaling_configuration_names": {
				TargetField: "ScalingConfigurationNames",
				ConvertType: ve.ConvertWithN,
			},
		},
		NameField:    "ScalingConfigurationName",
		IdField:      "ScalingConfigurationId",
		CollectField: "scaling_configurations",
		ResponseConverts: map[string]ve.ResponseConvert{
			"ScalingConfigurationId": {
				TargetField: "id",
				KeepDefault: true,
			},
			"Eip.Bandwidth": {
				TargetField: "eip_bandwidth",
			},
			"Eip.ISP": {
				TargetField: "eip_isp",
			},
			"Eip.BillingType": {
				TargetField: "eip_billing_type",
			},
		},
	}
}

func (s *VolcengineScalingConfigurationService) ReadResourceId(id string) string {
	return id
}

func (s *VolcengineScalingConfigurationService) enableConfigurationCallback(enableConfigId string, d *schema.ResourceData) ve.Callback {
	status := "Active"
	configId, groupId := d.Id(), d.Get("scaling_group_id").(string)
	if configId != enableConfigId {
		status = "InActive"
	}
	return ve.Callback{
		Call: ve.SdkCall{
			Action:      "EnableScalingConfiguration",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"ScalingConfigurationId": enableConfigId,
				"ScalingGroupId":         groupId,
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{status},
				Timeout: d.Timeout(schema.TimeoutCreate),
			},
		},
	}
}

func (s *VolcengineScalingConfigurationService) enableOrDisableGroupCallback(resourceData *schema.ResourceData) ve.Callback {
	groupId, enable := resourceData.Get("scaling_group_id").(string), resourceData.Get("enable").(bool)
	param := &map[string]interface{}{
		"ScalingGroupId": groupId,
	}
	if enable {
		return ve.Callback{
			Call: ve.SdkCall{
				Action:      "EnableScalingGroup",
				ConvertMode: ve.RequestConvertIgnore,
				SdkParam:    param,
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
			},
		}
	} else {
		return ve.Callback{
			Call: ve.SdkCall{
				Action:      "DisableScalingGroup",
				ConvertMode: ve.RequestConvertIgnore,
				SdkParam:    param,
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
			},
		}
	}
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "auto_scaling",
		Action:      actionName,
		Version:     "2020-01-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
	}
}
