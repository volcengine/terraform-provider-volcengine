package clb

import (
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineClbService struct {
	Client *ve.SdkClient
}

func NewClbService(c *ve.SdkClient) *VolcengineClbService {
	return &VolcengineClbService{
		Client: c,
	}
}

func (s *VolcengineClbService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineClbService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	data, err = ve.WithPageNumberQuery(condition, "PageSize", "PageNumber", 20, 1, func(m map[string]interface{}) ([]interface{}, error) {
		action := "DescribeLoadBalancers"
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

		results, err = ve.ObtainSdkValue("Result.LoadBalancers", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.LoadBalancers is not Slice")
		}
		return data, err
	})
	if err != nil {
		return data, err
	}

	for _, value := range data {
		clb, ok := value.(map[string]interface{})
		if !ok {
			return data, fmt.Errorf(" Clb is not map ")
		}

		eipAction := "DescribeLoadBalancerAttributes"
		eipReq := map[string]interface{}{
			"LoadBalancerId": clb["LoadBalancerId"],
		}
		logger.Debug(logger.ReqFormat, eipAction, eipReq)
		eipResp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(eipAction), &eipReq)
		if err != nil {
			return data, err
		}
		logger.Debug(logger.RespFormat, eipAction, *eipResp)

		eipConfig, err := ve.ObtainSdkValue("Result.Eip", *eipResp)
		if err != nil {
			return data, err
		}
		clb["EipBillingConfig"] = eipConfig

		ipv6EipConfig, err := ve.ObtainSdkValue("Result.Ipv6AddressBandwidth", *eipResp)
		if err != nil {
			return data, err
		}
		clb["Ipv6AddressBandwidth"] = ipv6EipConfig

		// `PostPaid` 实例不需查询续费相关信息
		if billingType := clb["LoadBalancerBillingType"]; billingType == 2.0 {
			continue
		}
		billingAction := "DescribeLoadBalancersBilling"
		billingReq := map[string]interface{}{
			"LoadBalancerIds.1": clb["LoadBalancerId"],
		}
		logger.Debug(logger.ReqFormat, billingAction, billingReq)
		billingResp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(billingAction), &billingReq)
		if err != nil {
			return data, err
		}
		logger.Debug(logger.RespFormat, billingAction, *billingResp)

		billingConfigs, err := ve.ObtainSdkValue("Result.LoadBalancerBillingConfigs", *billingResp)
		if err != nil {
			return data, err
		}
		if billingConfigs == nil {
			return data, fmt.Errorf(" DescribeLoadBalancersBilling error ")
		}
		configs, ok := billingConfigs.([]interface{})
		if !ok {
			return data, fmt.Errorf(" Result.LoadBalancerBillingConfigs is not slice ")
		}
		if len(configs) == 0 {
			return data, fmt.Errorf("LoadBalancerBilling of the clb instance %s is not exist ", clb["LoadBalancerId"])
		}
		config, ok := configs[0].(map[string]interface{})
		if !ok {
			return data, fmt.Errorf(" BillingConfigs is not map ")
		}
		for k, v := range config {
			clb[k] = v
		}
	}

	return data, err
}

func (s *VolcengineClbService) ReadResource(resourceData *schema.ResourceData, clbId string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if clbId == "" {
		clbId = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"LoadBalancerIds.1": clbId,
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
		return data, fmt.Errorf("Clb %s not exist ", clbId)
	}

	data["RegionId"] = s.Client.Region

	return data, err
}

func (s *VolcengineClbService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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
			failStates = append(failStates, "CreateFailed")
			demo, err = s.ReadResource(resourceData, id)
			if err != nil {
				return nil, "", err
			}
			status, err = ve.ObtainSdkValue("Status", demo)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("Clb  status  error, status:%s", status.(string))
				}
			}
			//注意 返回的第一个参数不能为空 否则会一直等下去
			return demo, status.(string), err
		},
	}

}

func (VolcengineClbService) WithResourceResponseHandlers(clb map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return clb, map[string]ve.ResponseConvert{
			"LoadBalancerBillingType": {
				TargetField: "load_balancer_billing_type",
				Convert: func(i interface{}) interface{} {
					if i == nil {
						return nil
					}
					billingType := i.(float64)
					switch billingType {
					case 1:
						return "PrePaid"
					case 2:
						return "PostPaid"
					case 3:
						return "PostPaidByLCU"
					}
					return i
				},
			},
			"RenewType": {
				TargetField: "renew_type",
				Convert: func(i interface{}) interface{} {
					if i == nil {
						return nil
					}
					renewType := i.(float64)
					switch renewType {
					case 1:
						return "ManualRenew"
					case 2:
						return "AutoRenew"
					case 3:
						return "NoneRenew"
					}
					return i
				},
			},
			"EipID": {
				TargetField: "eip_id",
			},
			"ISP": {
				TargetField: "isp",
			},
			"EipBillingType": {
				TargetField: "eip_billing_type",
				Convert: func(i interface{}) interface{} {
					if i == nil {
						return nil
					}
					billingType := i.(float64)
					switch billingType {
					case 1:
						return "PrePaid"
					case 2:
						return "PostPaidByBandwidth"
					case 3:
						return "PostPaidByTraffic"
					}
					return ""
				},
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}

}

func (s *VolcengineClbService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateLoadBalancer",
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"load_balancer_billing_type": {
					TargetField: "LoadBalancerBillingType",
					Convert: func(data *schema.ResourceData, i interface{}) interface{} {
						if i == nil {
							return nil
						}
						billingType := i.(string)
						switch billingType {
						case "PrePaid":
							return 1
						case "PostPaid":
							return 2
						case "PostPaidByLCU":
							return 3
						}
						return i
					},
				},
				"eip_billing_config": {
					TargetField: "EipBillingConfig",
					ConvertType: ve.ConvertListUnique,
					NextLevelConvert: map[string]ve.RequestConvert{
						"isp": {
							TargetField: "ISP",
						},
					},
				},
				"tags": {
					TargetField: "Tags",
					ConvertType: ve.ConvertListN,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if regionId, ok := (*call.SdkParam)["RegionId"]; !ok {
					(*call.SdkParam)["RegionId"] = s.Client.Region
				} else if regionId.(string) != s.Client.Region {
					return false, fmt.Errorf("region_id is not equal to provider region config(%s)", s.Client.Region)
				}

				// private 类型不传 eip_billing_config
				if (*call.SdkParam)["Type"] == "private" {
					delete(*call.SdkParam, "EipBillingConfig.ISP")
					delete(*call.SdkParam, "EipBillingConfig.EipBillingType")
					delete(*call.SdkParam, "EipBillingConfig.Bandwidth")
				}
				if eipBillingType, exist := (*call.SdkParam)["EipBillingConfig.EipBillingType"]; exist {
					ty := 0
					switch eipBillingType.(string) {
					case "PrePaid":
						ty = 1
					case "PostPaidByBandwidth":
						ty = 2
					case "PostPaidByTraffic":
						ty = 3
					}
					(*call.SdkParam)["EipBillingConfig.EipBillingType"] = ty
				}

				// PeriodUnit 默认传 Month
				if (*call.SdkParam)["LoadBalancerBillingType"] == 1 {
					(*call.SdkParam)["PeriodUnit"] = "Month"
				}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				//创建clb
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				//注意 获取内容 这个地方不能是指针 需要转一次
				id, _ := ve.ObtainSdkValue("Result.LoadBalancerId", *resp)
				d.SetId(id.(string))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Active"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	return []ve.Callback{callback}

}

func (s *VolcengineClbService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	attributesCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyLoadBalancerAttributes",
			ConvertMode: ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"load_balancer_name": {
					TargetField: "LoadBalancerName",
				},
				"description": {
					TargetField: "Description",
				},
				"modification_protection_status": {
					TargetField: "ModificationProtectionStatus",
				},
				"modification_protection_reason": {
					TargetField: "ModificationProtectionReason",
				},
				"load_balancer_spec": {
					TargetField: "LoadBalancerSpec",
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				oldType, _ := d.GetChange("load_balancer_billing_type")
				if oldType == "PostPaidByLCU" {
					delete(*call.SdkParam, "LoadBalancerSpec")
				}
				if len(*call.SdkParam) > 0 {
					(*call.SdkParam)["LoadBalancerId"] = d.Id()
					return true, nil
				}
				return false, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				//修改clb属性
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Active"},
				Timeout: resourceData.Timeout(schema.TimeoutUpdate),
			},
		},
	}
	callbacks = append(callbacks, attributesCallback)

	if resourceData.HasChange("load_balancer_billing_type") {
		billingTypeCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ConvertLoadBalancerBillingType",
				ConvertMode: ve.RequestConvertInConvert,
				Convert: map[string]ve.RequestConvert{
					"load_balancer_billing_type": {
						TargetField: "LoadBalancerBillingType",
						Convert: func(data *schema.ResourceData, i interface{}) interface{} {
							if i == nil {
								return nil
							}
							billingType := i.(string)
							switch billingType {
							case "PrePaid":
								return 1
							case "PostPaid":
								return 2
							case "PostPaidByLCU":
								return 3
							}
							return i
						},
					},
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					if len(*call.SdkParam) > 0 {
						(*call.SdkParam)["LoadBalancerId"] = d.Id()
						oldType, newType := d.GetChange("load_balancer_billing_type")
						if oldType == "PostPaidByLCU" && newType == "PostPaid" {
							(*call.SdkParam)["LoadBalancerSpec"] = d.Get("load_balancer_spec")
						}

						if (*call.SdkParam)["LoadBalancerBillingType"].(int) == 1 {
							// PeriodUnit 默认传 Month
							(*call.SdkParam)["PeriodUnit"] = "Month"
							(*call.SdkParam)["Period"] = d.Get("period")
						}
						return true, nil
					}
					return false, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
					//修改 clb 计费类型
					resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
					time.Sleep(10 * time.Second)
					return resp, err
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{"Active"},
					Timeout: resourceData.Timeout(schema.TimeoutUpdate),
				},
			},
		}
		callbacks = append(callbacks, billingTypeCallback)
	} else if resourceData.Get("renew_type").(string) == "ManualRenew" && resourceData.HasChange("period") {
		renewCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "RenewLoadBalancer",
				ConvertMode: ve.RequestConvertInConvert,
				Convert: map[string]ve.RequestConvert{
					"period": {
						TargetField: "Period",
						Convert: func(data *schema.ResourceData, i interface{}) interface{} {
							o, n := data.GetChange("period")
							return n.(int) - o.(int)
						},
					},
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					if len(*call.SdkParam) > 0 {
						if (*call.SdkParam)["Period"].(int) <= 0 {
							return false, fmt.Errorf("period can only be enlarged ")
						}

						// PeriodUnit 默认传 Month
						(*call.SdkParam)["PeriodUnit"] = "Month"
						(*call.SdkParam)["LoadBalancerId"] = d.Id()
						return true, nil
					}
					return false, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
				AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
					return nil
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{"Active"},
					Timeout: resourceData.Timeout(schema.TimeoutUpdate),
				},
			},
		}
		callbacks = append(callbacks, renewCallback)
	}

	// 更新Tags
	setResourceTagsCallbacks := ve.SetResourceTags(s.Client, "TagResources", "UntagResources", "CLB", resourceData, getUniversalInfo)
	callbacks = append(callbacks, setResourceTagsCallbacks...)

	return callbacks
}

func (s *VolcengineClbService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteLoadBalancer",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"LoadBalancerId": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				//删除Clb
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
							return resource.NonRetryableError(fmt.Errorf("error on  reading clb on delete %q, %w", d.Id(), callErr))
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

func (s *VolcengineClbService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "LoadBalancerIds",
				ConvertType: ve.ConvertWithN,
			},
			"tags": {
				TargetField: "TagFilters",
				ConvertType: ve.ConvertListN,
				NextLevelConvert: map[string]ve.RequestConvert{
					"value": {
						TargetField: "Values.1",
					},
				},
			},
		},
		NameField:    "LoadBalancerName",
		IdField:      "LoadBalancerId",
		CollectField: "clbs",
		ResponseConverts: map[string]ve.ResponseConvert{
			"LoadBalancerId": {
				TargetField: "id",
				KeepDefault: true,
			},
			"EipID": {
				TargetField: "eip_id",
			},
			"EniID": {
				TargetField: "eni_id",
			},
			"LoadBalancerBillingType": {
				TargetField: "load_balancer_billing_type",
				Convert: func(i interface{}) interface{} {
					if i == nil {
						return nil
					}
					billingType := i.(float64)
					switch billingType {
					case 1:
						return "PrePaid"
					case 2:
						return "PostPaid"
					case 3:
						return "PostPaidByLCU"
					}
					return i
				},
			},
			"RenewType": {
				TargetField: "renew_type",
				Convert: func(i interface{}) interface{} {
					if i == nil {
						return nil
					}
					renewType := i.(float64)
					switch renewType {
					case 1:
						return "ManualRenew"
					case 2:
						return "AutoRenew"
					case 3:
						return "NoneRenew"
					}
					return i
				},
			},
			"ISP": {
				TargetField: "isp",
			},
			"EipBillingType": {
				TargetField: "eip_billing_type",
				Convert: func(i interface{}) interface{} {
					if i == nil {
						return nil
					}
					billingType := i.(float64)
					switch billingType {
					case 1:
						return "PrePaid"
					case 2:
						return "PostPaidByBandwidth"
					case 3:
						return "PostPaidByTraffic"
					}
					return ""
				},
			},
			"BillingType": {
				TargetField: "billing_type",
				Convert: func(i interface{}) interface{} {
					if i == nil {
						return nil
					}
					billingType := i.(float64)
					switch billingType {
					case 1:
						return "PrePaid"
					case 2:
						return "PostPaidByBandwidth"
					case 3:
						return "PostPaidByTraffic"
					}
					return ""
				},
			},
		},
	}
}

func (s *VolcengineClbService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "clb",
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		Action:      actionName,
	}
}

func (s *VolcengineClbService) ProjectTrn() *ve.ProjectTrn {
	return &ve.ProjectTrn{
		ServiceName:          "clb",
		ResourceType:         "clb",
		ProjectResponseField: "ProjectName",
		ProjectSchemaField:   "project_name",
	}
}

func (s *VolcengineClbService) UnsubscribeInfo(resourceData *schema.ResourceData, resource *schema.Resource) (*ve.UnsubscribeInfo, error) {
	info := ve.UnsubscribeInfo{
		InstanceId: s.ReadResourceId(resourceData.Id()),
	}
	if resourceData.Get("load_balancer_billing_type") == "PrePaid" {
		info.Products = []string{"CLB"}
		info.NeedUnsubscribe = true
	}
	return &info, nil
}
