package kafka_instance

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

type VolcengineKafkaInstanceService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewKafkaInstanceService(c *ve.SdkClient) *VolcengineKafkaInstanceService {
	return &VolcengineKafkaInstanceService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineKafkaInstanceService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineKafkaInstanceService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
	)
	if v, ok := condition["Tags"]; ok {
		if len(v.(map[string]interface{})) == 0 {
			delete(condition, "Tags")
		}
	}
	return ve.WithPageNumberQuery(condition, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeInstances"

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

		for _, element := range results.([]interface{}) {
			instance := element.(map[string]interface{})
			// 拆开 ChargeDetail
			chargeInfo := instance["ChargeDetail"].(map[string]interface{})
			for k, v := range chargeInfo {
				instance[k] = v
			}
			delete(instance, "ChargeDetail")

			// update tags
			if v, ok := instance["Tags"]; ok {
				var tags []interface{}
				for k, v := range v.(map[string]interface{}) {
					tags = append(tags, map[string]interface{}{
						"Key":   k,
						"Value": v,
					})
				}
				instance["Tags"] = tags
			}

			// 获取 InstanceDetail 信息
			req := map[string]interface{}{
				"InstanceId": instance["InstanceId"],
			}
			logger.Debug(logger.ReqFormat, "DescribeInstanceDetail", req)
			detail, err := s.Client.UniversalClient.DoCall(getUniversalInfo("DescribeInstanceDetail"), &req)
			if err != nil {
				return data, err
			}
			logger.Debug(logger.RespFormat, "DescribeInstanceDetail", req, *detail)
			connection, err := ve.ObtainSdkValue("Result.ConnectionInfo", *detail)
			if err != nil {
				return data, err
			}
			instance["ConnectionInfo"] = connection
			params, err := ve.ObtainSdkValue("Result.Parameters", *detail)
			if err != nil {
				return data, err
			}
			paramsMap := make(map[string]interface{})
			if err = json.Unmarshal([]byte(params.(string)), &paramsMap); err != nil {
				return data, err
			}
			var paramsList []interface{}
			for k, v := range paramsMap {
				paramsList = append(paramsList, map[string]interface{}{
					"ParameterName":  k,
					"ParameterValue": v,
				})
			}
			instance["Parameters"] = paramsList
		}
		return results.([]interface{}), err
	})
}

func (s *VolcengineKafkaInstanceService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
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
		return data, fmt.Errorf("kafka_instance %s not exist ", id)
	}
	// parameters 会有默认参数，防止不一致产生
	delete(data, "Parameters")
	if parameterSet, ok := resourceData.GetOk("parameters"); ok {
		if set, ok := parameterSet.(*schema.Set); ok {
			data["Parameters"] = set.List()
		}
	}
	return data, err
}

func (s *VolcengineKafkaInstanceService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending: []string{},
		// 15s后才能查询 ChargeInfo
		Delay:      15 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				d          map[string]interface{}
				status     interface{}
				failStates []string
			)
			failStates = append(failStates, "CreateFailed", "Error", "Fail", "Failed")

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

			d, err = s.ReadResource(resourceData, id)
			if err != nil {
				return nil, "", err
			}
			status, err = ve.ObtainSdkValue("InstanceStatus", d)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("kafka_instance status error, status: %s", status.(string))
				}
			}
			return d, status.(string), err
		},
	}
}

func (s *VolcengineKafkaInstanceService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateInstance",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"parameters": {
					ConvertType: ve.ConvertJsonObjectArray,
				},
				"tags": {
					ConvertType: ve.ConvertJsonObjectArray,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				subnetId := (*call.SdkParam)["SubnetId"].(string)
				action := "DescribeSubnetAttributes"
				req := map[string]interface{}{
					"SubnetId": subnetId,
				}
				resp, err := s.Client.UniversalClient.DoCall(getVpcUniversalInfo(action), &req)
				if err != nil {
					return false, err
				}
				logger.Debug(logger.RespFormat, action, req, *resp)
				vpcId, err := ve.ObtainSdkValue("Result.VpcId", *resp)
				if err != nil {
					return false, err
				}
				(*call.SdkParam)["VpcId"] = vpcId
				zoneId, err := ve.ObtainSdkValue("Result.ZoneId", *resp)
				if err != nil {
					return false, err
				}
				(*call.SdkParam)["ZoneId"] = zoneId
				// update charge info
				charge := make(map[string]interface{})
				if (*call.SdkParam)["ChargeType"] == "PrePaid" {
					if (*call.SdkParam)["Period"] == nil || (*call.SdkParam)["Period"].(int) < 1 {
						return false, fmt.Errorf("Instance Charge Type is PrePaid. Must set Period more than 1. ")
					}
					charge["PeriodUnit"] = "Month"
				}
				charge["ChargeType"] = (*call.SdkParam)["ChargeType"]
				delete(*call.SdkParam, "ChargeType")
				if v, ok := (*call.SdkParam)["AutoRenew"]; ok {
					charge["AutoRenew"] = v
					delete(*call.SdkParam, "AutoRenew")
				}
				if v, ok := (*call.SdkParam)["Period"]; ok {
					charge["Period"] = v
					delete(*call.SdkParam, "Period")
				}
				(*call.SdkParam)["ChargeInfo"] = charge
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				// update tags
				if v, ok := (*call.SdkParam)["Tags"]; ok {
					tags := v.([]interface{})
					if len(tags) > 0 {
						temp := make(map[string]interface{})
						for _, ele := range tags {
							e := ele.(map[string]interface{})
							temp[e["Key"].(string)] = e["Value"]
						}
						(*call.SdkParam)["Tags"] = temp
					}
				}
				// update params
				if v, ok := (*call.SdkParam)["Parameters"]; ok {
					params := v.([]interface{})
					if len(params) > 0 {
						temp := make(map[string]interface{})
						for _, ele := range params {
							e := ele.(map[string]interface{})
							temp[e["ParameterName"].(string)] = e["ParameterValue"]
						}
						bytes, err := json.Marshal(&temp)
						if err != nil {
							return nil, err
						}
						(*call.SdkParam)["Parameters"] = string(bytes)
					}
				}

				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.ReqFormat, call.Action, *resp)
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

func (VolcengineKafkaInstanceService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineKafkaInstanceService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var res []ve.Callback
	if resourceData.HasChange("instance_name") || resourceData.HasChange("instance_description") {
		res = append(res, ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyInstanceAttributes",
				ConvertMode: ve.RequestConvertInConvert,
				ContentType: ve.ContentTypeJson,
				Convert: map[string]ve.RequestConvert{
					"instance_name": {
						TargetField: "InstanceName",
					},
					"instance_description": {
						TargetField: "InstanceDescription",
					},
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["InstanceId"] = d.Id()
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, *call.SdkParam)
					resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
					logger.Debug(logger.ReqFormat, call.Action, *resp)
					return resp, err
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{"Running"},
					Timeout: resourceData.Timeout(schema.TimeoutCreate),
				},
			},
		})
	}
	if resourceData.HasChanges("compute_spec", "storage_space", "partition_number") {
		res = append(res, ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyInstanceSpec",
				ConvertMode: ve.RequestConvertInConvert,
				ContentType: ve.ContentTypeJson,
				Convert: map[string]ve.RequestConvert{
					"compute_spec": {
						TargetField: "ComputeSpec",
					},
					"storage_space": {
						TargetField: "StorageSpace",
					},
					"partition_number": {
						TargetField: "PartitionNumber",
					},
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["InstanceId"] = d.Id()
					if d.HasChange("compute_spec") { // 变更实例的计算规格时才需要选择是否再均衡
						if v, ok := d.GetOkExists("need_rebalance"); ok {
							(*call.SdkParam)["NeedRebalance"] = v
						}
						if v, ok := d.GetOkExists("rebalance_time"); ok {
							(*call.SdkParam)["RebalanceTime"] = v
						}
					}
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, *call.SdkParam)
					resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
					logger.Debug(logger.ReqFormat, call.Action, *resp)
					return resp, err
				},
				AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
					time.Sleep(10 * time.Second)
					return nil
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{"Running"},
					Timeout: resourceData.Timeout(schema.TimeoutCreate),
				},
			},
		})
	}

	if resourceData.HasChange("parameters") {
		parameterCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyInstanceParameters",
				ContentType: ve.ContentTypeJson,
				ConvertMode: ve.RequestConvertInConvert,
				Convert: map[string]ve.RequestConvert{
					"parameters": {
						ConvertType: ve.ConvertJsonObjectArray,
						ForceGet:    true,
					},
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					if _, exist := (*call.SdkParam)["Parameters"]; !exist {
						return nil, nil
					}
					params := (*call.SdkParam)["Parameters"].([]interface{})
					if len(params) == 0 {
						return nil, nil
					}
					temp := make(map[string]interface{})
					for _, ele := range params {
						para := ele.(map[string]interface{})
						temp[para["ParameterName"].(string)] = para["ParameterValue"]
					}
					bytes, err := json.Marshal(&temp)
					if err != nil {
						return nil, err
					}
					(*call.SdkParam)["Parameters"] = string(bytes)
					(*call.SdkParam)["InstanceId"] = d.Id()

					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{"Running"},
					Timeout: resourceData.Timeout(schema.TimeoutUpdate),
				},
			},
		}

		res = append(res, parameterCallback)
	}
	if resourceData.HasChanges("charge_type") {
		res = append(res, ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyInstanceChargeType",
				ConvertMode: ve.RequestConvertIgnore,
				ContentType: ve.ContentTypeJson,
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					// 仅支持按量付费转包年包月
					if d.Get("charge_type") == "PostPaid" {
						return false, fmt.Errorf("onny support PostPaid to PrePaid")
					}

					if d.Get("charge_type") == "PrePaid" {
						if d.Get("period") == nil || d.Get("period").(int) < 1 {
							return false, fmt.Errorf("Instance Charge Type is PrePaid. Must set Period more than 1. ")
						}
					}

					(*call.SdkParam)["InstanceId"] = d.Id()
					charge := make(map[string]interface{})
					charge["PeriodUnit"] = "Month"
					charge["AutoRenew"] = d.Get("auto_renew")
					charge["Period"] = d.Get("period")
					charge["ChargeType"] = d.Get("charge_type")
					(*call.SdkParam)["ChargeInfo"] = charge
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
					Timeout: resourceData.Timeout(schema.TimeoutUpdate),
				},
			},
		})
	}

	// 更新Tags
	res = s.setResourceTags(resourceData, res)
	return res
}

func (s *VolcengineKafkaInstanceService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteInstance",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"InstanceId": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineKafkaInstanceService) setResourceTags(resourceData *schema.ResourceData, callbacks []ve.Callback) []ve.Callback {
	addedTags, removedTags, _, _ := ve.GetSetDifference("tags", resourceData, TagsHash, false)

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
						t := tag.(map[string]interface{})
						temp := make(map[string]interface{})
						temp["Key"] = t["key"].(string)
						temp["Value"] = t["value"].(string)
						(*call.SdkParam)["Tags"] = append((*call.SdkParam)["Tags"].([]map[string]interface{}), temp)
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

func (s *VolcengineKafkaInstanceService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			// Map<String, Array of String> 类型
			"tags": {
				Convert: func(data *schema.ResourceData, i interface{}) interface{} {
					tags := i.(*schema.Set).List()
					res := make(map[string]interface{})
					for _, ele := range tags {
						tag := ele.(map[string]interface{})
						res[tag["key"].(string)] = []interface{}{tag["value"]}
					}
					return res
				},
			},
		},
		NameField:    "InstanceName",
		IdField:      "InstanceId",
		CollectField: "instances",
		ResponseConverts: map[string]ve.ResponseConvert{
			"InstanceId": {
				TargetField: "id",
				KeepDefault: true,
			},
		},
	}
}

func (s *VolcengineKafkaInstanceService) ReadResourceId(id string) string {
	return id
}

func (s *VolcengineKafkaInstanceService) ProjectTrn() *ve.ProjectTrn {
	return &ve.ProjectTrn{
		ServiceName:          "Kafka",
		ResourceType:         "instance",
		ProjectResponseField: "ProjectName",
		ProjectSchemaField:   "project_name",
	}
}

func (s *VolcengineKafkaInstanceService) UnsubscribeInfo(resourceData *schema.ResourceData, resource *schema.Resource) (*ve.UnsubscribeInfo, error) {
	info := ve.UnsubscribeInfo{
		InstanceId: s.ReadResourceId(resourceData.Id()),
	}
	if resourceData.Get("charge_type") == "PrePaid" {
		info.Products = []string{"Message_Queue_for_Kafka"}
		info.NeedUnsubscribe = true
	}
	return &info, nil
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "kafka",
		Version:     "2022-05-01",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
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
