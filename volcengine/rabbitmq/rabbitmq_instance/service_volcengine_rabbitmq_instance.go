package rabbitmq_instance

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineRabbitmqInstanceService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewRabbitmqInstanceService(c *ve.SdkClient) *VolcengineRabbitmqInstanceService {
	return &VolcengineRabbitmqInstanceService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineRabbitmqInstanceService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineRabbitmqInstanceService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
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
		data, ok = results.([]interface{})
		if !ok {
			return data, fmt.Errorf(" Result.InstancesInfo is not slice")
		}

		for _, element := range data {
			instanceMap, ok := element.(map[string]interface{})
			if !ok {
				return data, fmt.Errorf(" InstancesInfo is not map ")
			}

			// 获取 InstanceDetail 信息
			action := "DescribeInstanceDetail"
			req := map[string]interface{}{
				"InstanceId": instanceMap["InstanceId"],
			}
			logger.Debug(logger.ReqFormat, action, req)
			detail, err := s.Client.UniversalClient.DoCall(getUniversalInfo(action), &req)
			if err != nil {
				return data, err
			}
			logger.Debug(logger.RespFormat, action, req, *detail)

			endpoints, err := ve.ObtainSdkValue("Result.Endpoints", *detail)
			if err != nil {
				return data, err
			}
			endpointArr, ok := endpoints.([]interface{})
			if !ok {
				return data, fmt.Errorf(" Result.Endpoints is not slice ")
			}
			instanceMap["Endpoints"] = endpointArr

			userName, err := ve.ObtainSdkValue("Result.BasicInstanceInfo.InitUserName", *detail)
			if err != nil {
				return data, err
			}
			instanceMap["InitUserName"] = userName

			privateDns, err := ve.ObtainSdkValue("Result.BasicInstanceInfo.ApplyPrivateDNSToPublic", *detail)
			if err != nil {
				return data, err
			}
			instanceMap["ApplyPrivateDNSToPublic"] = privateDns
		}
		return data, err
	})

}

func (s *VolcengineRabbitmqInstanceService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
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
		instanceMap := make(map[string]interface{})
		if instanceMap, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("Value is not map ")
		}
		if instanceMap["InstanceId"] == id {
			data = instanceMap
			break
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("rabbitmq_instance %s not exist ", id)
	}

	if zoneId, ok := data["ZoneId"]; ok {
		zoneIds := strings.Split(zoneId.(string), ",")
		data["ZoneIds"] = zoneIds
	}

	data["UserName"] = data["InitUserName"]
	data["ChargeInfo"] = data["ChargeDetail"]

	return data, err
}

func (s *VolcengineRabbitmqInstanceService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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
				resp       *map[string]interface{}
				data       []interface{}
				results    interface{}
				ok         bool
			)
			failStates = append(failStates, "Failed", "CreateFailed")
			req := map[string]interface{}{
				"InstanceId": id,
			}
			// instance status 不为 Running 时，不能调用 DescribeInstanceDetail
			data, err = ve.WithPageNumberQuery(req, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
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
				data, ok = results.([]interface{})
				if !ok {
					return []interface{}{}, fmt.Errorf(" Result.InstancesInfo is not slice")
				}
				return data, err
			})
			if err != nil {
				return nil, "", err
			}
			for _, v := range data {
				instanceMap := make(map[string]interface{})
				if instanceMap, ok = v.(map[string]interface{}); !ok {
					return d, "", errors.New("Value is not map ")
				}
				if instanceMap["InstanceId"] == id {
					d = instanceMap
					break
				}
			}
			if len(d) == 0 {
				return d, "", fmt.Errorf("RefreshResourceState: rabbitmq_instance %s not exist ", id)
			}

			status, err = ve.ObtainSdkValue("InstanceStatus", d)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("rabbitmq_instance status error, status: %s", status.(string))
				}
			}

			// 等待 project 绑定
			projectName, err := ve.ObtainSdkValue("ProjectName", d)
			if err != nil {
				return nil, "", err
			}
			pjName, exist := resourceData.GetOkExists("project_name")
			if !exist {
				return d, status.(string), err
			}
			if projectName != pjName.(string) {
				return nil, "", nil
			}

			return d, status.(string), err
		},
	}
}

func (VolcengineRabbitmqInstanceService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"ApplyPrivateDNSToPublic": {
				TargetField: "apply_private_dns_to_public",
			},
		}, nil

	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineRabbitmqInstanceService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateInstance",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				//"apply_private_dns_to_public": {
				//	TargetField: "ApplyPrivateDNSToPublic",
				//},
				"charge_info": {
					TargetField: "ChargeInfo",
					ConvertType: ve.ConvertJsonObject,
				},
				"tags": {
					ConvertType: ve.ConvertJsonObjectArray,
				},
				"zone_ids": {
					Ignore: true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				// describe vpc id
				subnetId := (*call.SdkParam)["SubnetId"]
				vpcId, _, err := s.getVpcIdAndZoneIdBySubnet(subnetId.(string))
				if err != nil {
					return false, nil
				}
				(*call.SdkParam)["VpcId"] = vpcId

				zoneIdsArr := d.Get("zone_ids").(*schema.Set).List()
				zoneIds := make([]string, 0)
				for _, id := range zoneIdsArr {
					zoneIds = append(zoneIds, id.(string))
				}
				zoneId := strings.Join(zoneIds, ",")
				(*call.SdkParam)["ZoneId"] = zoneId

				(*call.SdkParam)["ClientToken"] = uuid.New().String()
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

func (s *VolcengineRabbitmqInstanceService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	if resourceData.HasChanges("instance_name", "instance_description") {
		callback := ve.Callback{
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
					if len(*call.SdkParam) > 0 {
						(*call.SdkParam)["InstanceId"] = d.Id()
						return true, nil
					}
					return false, nil
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
		}
		callbacks = append(callbacks, callback)
	}

	if resourceData.HasChanges("compute_spec", "storage_space") {
		specCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyInstanceSpec",
				ConvertMode: ve.RequestConvertInConvert,
				ContentType: ve.ContentTypeJson,
				Convert: map[string]ve.RequestConvert{
					"compute_spec": {
						TargetField: "ComputeSpec",
						ForceGet:    true,
					},
					"storage_space": {
						TargetField: "StorageSpace",
						ForceGet:    true,
					},
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					if len(*call.SdkParam) > 0 {
						(*call.SdkParam)["InstanceId"] = d.Id()
						(*call.SdkParam)["ClientToken"] = uuid.New().String()
						return true, nil
					}
					return false, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
					logger.Debug(logger.RespFormat, call.Action, resp, err)
					return resp, err
				},
				AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
					time.Sleep(5 * time.Second)
					return nil
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{"Running"},
					Timeout: resourceData.Timeout(schema.TimeoutUpdate),
				},
			},
		}
		callbacks = append(callbacks, specCallback)
	}

	if resourceData.HasChange("charge_info.0.charge_type") {
		chargeCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyInstanceChargeType",
				ConvertMode: ve.RequestConvertInConvert,
				ContentType: ve.ContentTypeJson,
				Convert: map[string]ve.RequestConvert{
					"charge_info": {
						TargetField: "ChargeInfo",
						ConvertType: ve.ConvertJsonObject,
						NextLevelConvert: map[string]ve.RequestConvert{
							"charge_type": {
								TargetField: "ChargeType",
								ForceGet:    true,
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
					},
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					old, _ := d.GetChange("charge_info.0.charge_type")
					if old == "PrePaid" {
						return false, fmt.Errorf("The operation is not permitted due to the instance charge type is prepaid. ")
					}
					if len(*call.SdkParam) > 0 {
						(*call.SdkParam)["InstanceId"] = d.Id()
						(*call.SdkParam)["ClientToken"] = uuid.New().String()
						return true, nil
					}
					return false, nil
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
		}
		callbacks = append(callbacks, chargeCallback)
	}

	if resourceData.HasChanges("user_name", "user_password") {
		passwordCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyUserPassword",
				ConvertMode: ve.RequestConvertInConvert,
				ContentType: ve.ContentTypeJson,
				Convert: map[string]ve.RequestConvert{
					"user_name": {
						TargetField: "UserName",
						ForceGet:    true,
					},
					"user_password": {
						TargetField: "Password",
						ForceGet:    true,
					},
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					if len(*call.SdkParam) > 0 {
						(*call.SdkParam)["InstanceId"] = d.Id()
						return true, nil
					}
					return false, nil
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
		}
		callbacks = append(callbacks, passwordCallback)

		// restart rabbitmq instance
		//restartCallback := s.restartInstanceCallback(resourceData)
		//callbacks = append(callbacks, restartCallback)
	}

	// Tags
	callbacks = s.setResourceTags(resourceData, callbacks)

	return callbacks
}

func (s *VolcengineRabbitmqInstanceService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
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
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 5*time.Minute)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(5*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading rabbitmq instance on delete %q, %w", d.Id(), callErr))
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

func (s *VolcengineRabbitmqInstanceService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"tags": {
				TargetField: "TagFilters",
				ConvertType: ve.ConvertJsonObjectArray,
			},
		},
		NameField:    "InstanceName",
		IdField:      "InstanceId",
		CollectField: "rabbitmq_instances",
		ContentType:  ve.ContentTypeJson,
		ResponseConverts: map[string]ve.ResponseConvert{
			"InstanceId": {
				TargetField: "id",
				KeepDefault: true,
			},
			"ApplyPrivateDNSToPublic": {
				TargetField: "apply_private_dns_to_public",
			},
		},
	}
}

func (s *VolcengineRabbitmqInstanceService) ReadResourceId(id string) string {
	return id
}

func (s *VolcengineRabbitmqInstanceService) setResourceTags(resourceData *schema.ResourceData, callbacks []ve.Callback) []ve.Callback {
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

func (s *VolcengineRabbitmqInstanceService) getVpcIdAndZoneIdBySubnet(subnetId string) (vpcId, zoneId string, err error) {
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

func (s *VolcengineRabbitmqInstanceService) restartInstanceCallback(resourceData *schema.ResourceData) ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "RestartInstance",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"InstanceId": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Running"},
				Timeout: resourceData.Timeout(schema.TimeoutUpdate),
			},
		},
	}
	return callback
}

func (s *VolcengineRabbitmqInstanceService) ProjectTrn() *ve.ProjectTrn {
	return &ve.ProjectTrn{
		ServiceName:          "RabbitMQ",
		ResourceType:         "instance",
		ProjectResponseField: "ProjectName",
		ProjectSchemaField:   "project_name",
	}
}

func (s *VolcengineRabbitmqInstanceService) UnsubscribeInfo(resourceData *schema.ResourceData, resource *schema.Resource) (*ve.UnsubscribeInfo, error) {
	info := ve.UnsubscribeInfo{
		InstanceId: s.ReadResourceId(resourceData.Id()),
	}
	if resourceData.Get("charge_info.0.charge_type").(string) == "PrePaid" {
		info.NeedUnsubscribe = true
		info.Products = []string{"Message_Queue_for_RabbitMQ"}
	}
	return &info, nil
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "RabbitMQ",
		Version:     "2022-01-01",
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
