package transit_router_bandwidth_package

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineTRBandwidthPackageService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewTRBandwidthPackageService(c *ve.SdkClient) *VolcengineTRBandwidthPackageService {
	return &VolcengineTRBandwidthPackageService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineTRBandwidthPackageService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineTRBandwidthPackageService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 20, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeTransitRouterBandwidthPackages"
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
		logger.Debug(logger.RespFormat, action, resp)
		results, err = ve.ObtainSdkValue("Result.TransitRouterBandwidthPackages", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.TransitRouterBandwidthPackages is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineTRBandwidthPackageService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}

	req := map[string]interface{}{
		"TransitRouterBandwidthPackageIds.1": id,
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
		return data, fmt.Errorf("TransitRouter bandwidth package %s not exist ", id)
	}

	// 查询计费信息
	params := map[string]interface{}{
		"TransitRouterBandwidthPackageIds.1": id,
	}
	action := "DescribeTransitRouterBandwidthPackagesBilling"
	logger.Debug(logger.ReqFormat, action, params)
	billingRes, err := s.Client.UniversalClient.DoCall(getUniversalInfo(action), &params)
	if err != nil {
		return data, err
	}
	logger.Debug(logger.AllFormat, action, params, *billingRes)

	tmpRes, err := ve.ObtainSdkValue("Result.TransitRouterBandwidthPackages", *billingRes)
	if err != nil {
		return data, err
	}
	if tmpRes == nil {
		return data, errors.New("DescribeTransitRouterBandwidthPackagesBilling Result is nil")
	}
	tmpData, ok := tmpRes.([]interface{})
	if !ok {
		return data, errors.New("Result.TransitRouterBandwidthPackages is not Slice")
	}
	if len(tmpData) == 0 {
		return data, fmt.Errorf("TransitRouterBandwidthPackagesBilling is %s not exist ", id)
	}
	data["RenewType"] = tmpData[0].(map[string]interface{})["RenewType"]
	data["RemainRenewTimes"] = tmpData[0].(map[string]interface{})["RemainRenewTimes"]

	return data, err
}

func (s *VolcengineTRBandwidthPackageService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				data   map[string]interface{}
				status interface{}
			)

			if err = resource.Retry(10*time.Minute, func() *resource.RetryError {
				data, err = s.ReadResource(resourceData, id)
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

			status, err = ve.ObtainSdkValue("Status", data)
			if err != nil {
				return nil, "", err
			}
			return data, status.(string), err
		},
	}

}

func (VolcengineTRBandwidthPackageService) WithResourceResponseHandlers(v map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		if v["BillingType"].(float64) == 1 {
			var (
				ct time.Time
				et time.Time
			)
			if strings.Contains(v["CreationTime"].(string), "+") {
				ct, _ = time.Parse("2006-01-02T15:04:05", v["CreationTime"].(string)[0:strings.Index(v["CreationTime"].(string), "+")])
			} else {
				ct, _ = time.Parse("2006-01-02 15:04:05", v["CreationTime"].(string))
			}
			if strings.Contains(v["ExpiredTime"].(string), "+") {
				et, _ = time.Parse("2006-01-02T15:04:05", v["ExpiredTime"].(string)[0:strings.Index(v["ExpiredTime"].(string), "+")])
			} else {
				et, _ = time.Parse("2006-01-02 15:04:05", v["ExpiredTime"].(string))
			}
			y := et.Year() - ct.Year()
			m := et.Month() - ct.Month()
			v["Period"] = y*12 + int(m)
		}
		return v, map[string]ve.ResponseConvert{
			"BillingType": {
				TargetField: "billing_type",
				Convert:     billingTypeResponseConvert,
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}

}

func (s *VolcengineTRBandwidthPackageService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateTransitRouterBandwidthPackage",
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				//"billing_type": {
				//	TargetField: "BillingType",
				//	Convert:     billingTypeRequestConvert,
				//},
				"tags": {
					TargetField: "Tags",
					ConvertType: ve.ConvertListN,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["LocalGeographicRegionSetId"] = "China"
				(*call.SdkParam)["PeerGeographicRegionSetId"] = "China"
				(*call.SdkParam)["BillingType"] = 1
				(*call.SdkParam)["PeriodUnit"] = "Month"
				(*call.SdkParam)["ClientToken"] = uuid.New().String()
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("Result.TransitRouterBandwidthPackageId", *resp)
				d.SetId(id.(string))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Available"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	return []ve.Callback{callback}

}

func (s *VolcengineTRBandwidthPackageService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callbacks := make([]ve.Callback, 0)

	modifyCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyTransitRouterBandwidthPackageAttributes",
			ConvertMode: ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"transit_router_bandwidth_package_name": {
					TargetField: "TransitRouterBandwidthPackageName",
				},
				"description": {
					TargetField: "Description",
				},
				"bandwidth": {
					TargetField: "Bandwidth",
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if len(*call.SdkParam) == 0 {
					return false, nil
				}
				(*call.SdkParam)["TransitRouterBandwidthPackageId"] = d.Id()
				delete(*call.SdkParam, "Tags")
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Available"},
				Timeout: resourceData.Timeout(schema.TimeoutUpdate),
			},
		},
	}
	callbacks = append(callbacks, modifyCallback)

	// 续费方式
	if resourceData.HasChanges("renew_type", "renew_period", "remain_renew_times") {
		renewCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "SetTransitRouterBandwidthPackageRenewal",
				ConvertMode: ve.RequestConvertInConvert,
				Convert: map[string]ve.RequestConvert{
					"renew_type": {
						TargetField: "RenewType",
						ConvertType: ve.ConvertDefault,
						ForceGet:    true,
					},
					"renew_period": {
						TargetField: "RenewPeriod",
						ConvertType: ve.ConvertDefault,
					},
					"remain_renew_times": {
						TargetField: "RemainRenewTimes",
						ConvertType: ve.ConvertDefault,
						ForceGet:    true,
					},
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					if len(*call.SdkParam) > 0 {
						(*call.SdkParam)["TransitRouterBandwidthPackageId"] = d.Id()
						delete(*call.SdkParam, "Tags")
						return true, nil
					}
					return false, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{"Available"},
					Timeout: resourceData.Timeout(schema.TimeoutUpdate),
				},
			},
		}
		callbacks = append(callbacks, renewCallback)
	}

	// 手动续费时长
	if resourceData.Get("renew_type").(string) == "Manual" && resourceData.HasChange("period") {
		periodCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "RenewTransitRouterBandwidthPackage",
				ConvertMode: ve.RequestConvertInConvert,
				Convert: map[string]ve.RequestConvert{
					"period": {
						ConvertType: ve.ConvertDefault,
						Convert: func(data *schema.ResourceData, i interface{}) interface{} {
							o, n := data.GetChange("period")
							return n.(int) - o.(int)
						},
					},
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					if len(*call.SdkParam) > 0 {
						if (*call.SdkParam)["Period"].(int) <= 0 {
							return false, fmt.Errorf(" period must be set and can only be increased ")
						}
						(*call.SdkParam)["PeriodUnit"] = "Month"
						(*call.SdkParam)["TransitRouterBandwidthPackageId"] = d.Id()
						delete(*call.SdkParam, "Tags")
						return true, nil
					}
					return false, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{"Available"},
					Timeout: resourceData.Timeout(schema.TimeoutUpdate),
				},
			},
		}
		callbacks = append(callbacks, periodCallback)
	}

	// 更新Tags
	setResourceTagsCallbacks := ve.SetResourceTags(s.Client, "TagResources", "UntagResources", "transitrouterbandwidthpackage", resourceData, getUniversalInfo)
	callbacks = append(callbacks, setResourceTagsCallbacks...)

	return callbacks
}

func (s *VolcengineTRBandwidthPackageService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteTransitRouterBandwidthPackage",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"TransitRouterBandwidthPackageId": resourceData.Get("transit_router_bandwidth_package_id"),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading transit router bandwidth package on delete %q, %w", d.Id(), callErr))
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

func (s *VolcengineTRBandwidthPackageService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "TransitRouterBandwidthPackageIds",
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
		NameField:    "TransitRouterBandwidthPackageName",
		IdField:      "TransitRouterBandwidthPackageId",
		CollectField: "bandwidth_packages",
		ResponseConverts: map[string]ve.ResponseConvert{
			"TransitRouterBandwidthPackageId": {
				TargetField: "id",
				KeepDefault: true,
			},
			"BillingType": {
				TargetField: "billing_type",
				Convert:     billingTypeResponseConvert,
			},
		},
	}
}

func (s *VolcengineTRBandwidthPackageService) ReadResourceId(id string) string {
	return id
}

func (s *VolcengineTRBandwidthPackageService) UnsubscribeInfo(resourceData *schema.ResourceData, resource *schema.Resource) (*ve.UnsubscribeInfo, error) {
	info := ve.UnsubscribeInfo{
		InstanceId: s.ReadResourceId(resourceData.Id()),
	}
	info.NeedUnsubscribe = true
	info.Products = []string{"TransitRouter_InterRegionBandwidth"}
	return &info, nil
}

func (s *VolcengineTRBandwidthPackageService) ProjectTrn() *ve.ProjectTrn {
	return &ve.ProjectTrn{
		ServiceName:          "transitrouter",
		ResourceType:         "transitrouterbandwidthpackage",
		ProjectResponseField: "ProjectName",
		ProjectSchemaField:   "project_name",
	}
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "transitrouter",
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		Action:      actionName,
	}
}
