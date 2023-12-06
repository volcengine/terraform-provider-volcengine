package bandwidth_package

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

type VolcengineBandwidthPackageService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewBandwidthPackageService(c *ve.SdkClient) *VolcengineBandwidthPackageService {
	return &VolcengineBandwidthPackageService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineBandwidthPackageService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineBandwidthPackageService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		next    string
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithNextTokenQuery(m, "MaxResults", "NextToken", 20, nil, func(condition map[string]interface{}) ([]interface{}, string, error) {
		action := "DescribeBandwidthPackages"
		bytes, _ := json.Marshal(condition)
		logger.Debug(logger.ReqFormat, action, string(bytes))
		if condition == nil {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), nil)
			if err != nil {
				return data, next, err
			}
		} else {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &condition)
			if err != nil {
				return data, next, err
			}
		}
		respBytes, _ := json.Marshal(resp)
		logger.Debug(logger.RespFormat, action, condition, string(respBytes))
		results, err = ve.ObtainSdkValue("Result.BandwidthPackages", *resp)
		if err != nil {
			return data, next, err
		}
		nextToken, err := ve.ObtainSdkValue("Result.NextToken", *resp)
		if err != nil {
			return data, next, err
		}
		next = nextToken.(string)
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, next, errors.New("Result.BandwidthPackages is not Slice")
		}
		return data, next, err
	})
}

func (s *VolcengineBandwidthPackageService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"BandwidthPackageIds.1": id,
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
		return data, fmt.Errorf("bandwidth_package %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineBandwidthPackageService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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
			failStates = append(failStates, "Failed")
			d, err = s.ReadResource(resourceData, id)
			if err != nil {
				return nil, "", err
			}
			status, err = ve.ObtainSdkValue("Status", d)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("bandwidth_package status error, status: %s", status.(string))
				}
			}
			return d, status.(string), err
		},
	}
}

func (s *VolcengineBandwidthPackageService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateBandwidthPackage",
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"billing_type": {
					TargetField: "BillingType",
					Convert:     billingTypeRequestConvert,
				},
				"tags": {
					TargetField: "Tags",
					ConvertType: ve.ConvertListN,
				},
				"security_protection_types": {
					TargetField: "SecurityProtectionTypes",
					ConvertType: ve.ConvertWithN,
				},
				"isp": {
					TargetField: "ISP",
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("Result.BandwidthPackageId", *resp)
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

func (VolcengineBandwidthPackageService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"BillingType": {
				TargetField: "billing_type",
				Convert:     billingTypeResponseConvert,
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineBandwidthPackageService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyBandwidthPackageAttributes",
			ConvertMode: ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"bandwidth_package_name": {
					TargetField: "BandwidthPackageName",
				},
				"description": {
					TargetField: "Description",
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["BandwidthPackageId"] = d.Id()
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Available"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	callbacks = append(callbacks, callback)
	if resourceData.HasChange("bandwidth") {
		bandwidthCallback := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyBandwidthPackageSpec",
				ConvertMode: ve.RequestConvertInConvert,
				Convert: map[string]ve.RequestConvert{
					"bandwidth": {
						TargetField: "Bandwidth",
					},
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["BandwidthPackageId"] = d.Id()
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
					logger.Debug(logger.RespFormat, call.Action, resp, err)
					return resp, err
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{"Available"},
					Timeout: resourceData.Timeout(schema.TimeoutCreate),
				},
			},
		}
		callbacks = append(callbacks, bandwidthCallback)
	}
	setResourceTagsCallbacks := ve.SetResourceTags(s.Client, "TagResources", "UntagResources", "bandwidthpackage", resourceData, getUniversalInfo)
	callbacks = append(callbacks, setResourceTagsCallbacks...)
	return callbacks
}

func (s *VolcengineBandwidthPackageService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteBandwidthPackage",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"BandwidthPackageId": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 5*time.Minute)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineBandwidthPackageService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "BandwidthPackageIds",
				ConvertType: ve.ConvertWithN,
			},
			"isp": {
				TargetField: "ISP",
			},
			"tag_filters": {
				TargetField: "TagFilters",
				ConvertType: ve.ConvertListN,
				NextLevelConvert: map[string]ve.RequestConvert{
					"key": {
						TargetField: "Key",
					},
					"values": {
						TargetField: "Values",
						ConvertType: ve.ConvertWithN,
					},
				},
			},
		},
		NameField:    "BandwidthPackageName",
		IdField:      "BandwidthPackageId",
		CollectField: "packages",
		ResponseConverts: map[string]ve.ResponseConvert{
			"BandwidthPackageId": {
				TargetField: "id",
				KeepDefault: true,
			},
			"BillingType": {
				TargetField: "billing_type",
				Convert:     billingTypeResponseConvert,
			},
			"ISP": {
				TargetField: "isp",
			},
		},
	}
}

func (s *VolcengineBandwidthPackageService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "vpc",
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		Action:      actionName,
	}
}

func (s *VolcengineBandwidthPackageService) ProjectTrn() *ve.ProjectTrn {
	return &ve.ProjectTrn{
		ServiceName:          "vpc",
		ResourceType:         "bandwidthpackage",
		ProjectResponseField: "ProjectName",
		ProjectSchemaField:   "project_name",
	}
}

func (s *VolcengineBandwidthPackageService) UnsubscribeInfo(resourceData *schema.ResourceData, resource *schema.Resource) (*ve.UnsubscribeInfo, error) {
	info := ve.UnsubscribeInfo{
		InstanceId: s.ReadResourceId(resourceData.Id()),
	}
	if resourceData.Get("billing_type").(string) == "PrePaid" {
		info.NeedUnsubscribe = true
		info.Products = []string{"BandwidthPackage"}
	}
	return &info, nil
}
