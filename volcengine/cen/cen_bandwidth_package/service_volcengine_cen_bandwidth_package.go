package cen_bandwidth_package

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineCenBandwidthPackageService struct {
	Client *ve.SdkClient
}

func NewCenBandwidthPackageService(c *ve.SdkClient) *VolcengineCenBandwidthPackageService {
	return &VolcengineCenBandwidthPackageService{
		Client: c,
	}
}

func (s *VolcengineCenBandwidthPackageService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineCenBandwidthPackageService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
		nameSet = make(map[string]bool)
	)
	if _, ok = m["CenBandwidthPackageNames.1"]; ok {
		i := 1
		for {
			filed := fmt.Sprintf("CenBandwidthPackageNames.%d", i)
			tmpName, ok := m[filed]
			if !ok {
				break
			}
			nameSet[tmpName.(string)] = true
			i++
			delete(m, filed)
		}
	}
	packages, err := ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 20, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		universalClient := s.Client.UniversalClient
		action := "DescribeCenBandwidthPackages"
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
		logger.Debug(logger.RespFormat, action, resp)
		results, err = ve.ObtainSdkValue("Result.CenBandwidthPackages", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.CenBandwidthPackages is not Slice")
		}
		return data, err
	})
	if err != nil || len(nameSet) == 0 {
		return packages, err
	}

	res := make([]interface{}, 0)
	for _, v := range packages {
		if !nameSet[v.(map[string]interface{})["CenBandwidthPackageName"].(string)] {
			continue
		}
		res = append(res, v)
	}
	return res, nil
}

func (s *VolcengineCenBandwidthPackageService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"CenBandwidthPackageIds.1": id,
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
		return data, fmt.Errorf("cen bandwidth package %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineCenBandwidthPackageService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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

			if err = resource.Retry(20*time.Minute, func() *resource.RetryError {
				demo, err = s.ReadResource(resourceData, id)
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

			status, err = ve.ObtainSdkValue("Status", demo)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("cen status error, status:%s", status.(string))
				}
			}
			project, err := ve.ObtainSdkValue("ProjectName", demo)
			if err != nil {
				return nil, "", err
			}
			if resourceData.Get("project_name") != nil && resourceData.Get("project_name").(string) != "" {
				if project != resourceData.Get("project_name") {
					return demo, "", err
				}
			}
			//注意 返回的第一个参数不能为空 否则会一直等下去
			return demo, status.(string), err
		},
	}

}

func (VolcengineCenBandwidthPackageService) WithResourceResponseHandlers(v map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return v, map[string]ve.ResponseConvert{
			"BillingType": {
				TargetField: "billing_type",
				Convert:     billingTypeResponseConvert,
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}

}

func (s *VolcengineCenBandwidthPackageService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateCenBandwidthPackage",
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
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["ClientToken"] = uuid.New().String()
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				//注意 获取内容 这个地方不能是指针 需要转一次
				id, _ := ve.ObtainSdkValue("Result.CenBandwidthPackageId", *resp)
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

func (s *VolcengineCenBandwidthPackageService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyCenBandwidthPackageAttributes",
			ConvertMode: ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"bandwidth": {
					ConvertType: ve.ConvertDefault,
				},
				"cen_bandwidth_package_name": {
					ConvertType: ve.ConvertDefault,
				},
				"description": {
					ConvertType: ve.ConvertDefault,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if len(*call.SdkParam) == 0 {
					return false, nil
				}
				(*call.SdkParam)["CenBandwidthPackageId"] = d.Id()
				delete(*call.SdkParam, "Tags")
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Available", "InUse"},
				Timeout: resourceData.Timeout(schema.TimeoutUpdate),
			},
		},
	}
	callbacks = append(callbacks, callback)
	setResourceTagsCallbacks := ve.SetResourceTags(s.Client, "TagResources", "UntagResources", "cenbandwidthpackage", resourceData, getUniversalInfo)
	callbacks = append(callbacks, setResourceTagsCallbacks...)
	if resourceData.HasChange("project_name") {
		projectCallback := s.ModifyProject(resourceData)
		callbacks = append(callbacks, projectCallback...)
	}
	return callbacks
}

func (s *VolcengineCenBandwidthPackageService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteCenBandwidthPackage",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"CenBandwidthPackageId": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				return nil, nil
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading cen bandwidth package on delete %q, %w", d.Id(), callErr))
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

func (s *VolcengineCenBandwidthPackageService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "CenBandwidthPackageIds",
				ConvertType: ve.ConvertWithN,
			},
			"cen_bandwidth_package_names": {
				TargetField: "CenBandwidthPackageNames",
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
		NameField:    "CenBandwidthPackageName",
		IdField:      "CenBandwidthPackageId",
		CollectField: "bandwidth_packages",
		ResponseConverts: map[string]ve.ResponseConvert{
			"CenBandwidthPackageId": {
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

func (s *VolcengineCenBandwidthPackageService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "cen",
		Action:      actionName,
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
	}
}

func (s *VolcengineCenBandwidthPackageService) ModifyProject(resourceData *schema.ResourceData) []ve.Callback {
	var call []ve.Callback
	id := s.ReadResourceId(resourceData.Id())
	if resourceData.HasChange("project_name") {
		modifyProject := ve.Callback{
			Call: ve.SdkCall{
				Action:      "MoveProjectResource",
				ConvertMode: ve.RequestConvertInConvert,
				Convert: map[string]ve.RequestConvert{
					"project_name": {
						ConvertType: ve.ConvertDefault,
						TargetField: "TargetProjectName",
					},
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					if (*call.SdkParam)["TargetProjectName"] == nil || (*call.SdkParam)["TargetProjectName"] == "" {
						return false, fmt.Errorf("Could set ProjectName to empty ")
					}
					//获取用户ID
					input := map[string]interface{}{
						"ProjectName": (*call.SdkParam)["TargetProjectName"],
					}
					logger.Debug(logger.ReqFormat, "GetProject", input)
					out, err := s.Client.UniversalClient.DoCall(s.getIAMUniversalInfo("GetProject"), &input)
					if err != nil {
						return false, err
					}
					accountId, err := ve.ObtainSdkValue("Result.AccountID", *out)
					if err != nil {
						return false, err
					}
					trnStr := fmt.Sprintf("trn:%s:%s:%d:%s/%s", "cen", "", int(accountId.(float64)),
						"cenbandwidthpackage", id)
					(*call.SdkParam)["ResourceTrn.1"] = trnStr
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(s.getIAMUniversalInfo(call.Action), call.SdkParam)
				},
			},
		}
		call = append(call, modifyProject)
	}
	return call
}

func (s *VolcengineCenBandwidthPackageService) getIAMUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "iam",
		Action:      actionName,
		Version:     "2018-01-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
		RegionType:  ve.Global,
	}
}

func (s *VolcengineCenBandwidthPackageService) UnsubscribeInfo(resourceData *schema.ResourceData, resource *schema.Resource) (*ve.UnsubscribeInfo, error) {
	info := ve.UnsubscribeInfo{
		InstanceId: s.ReadResourceId(resourceData.Id()),
	}
	if resourceData.Get("billing_type").(string) == "PrePaid" {
		info.NeedUnsubscribe = true
		info.Products = []string{"CEN"}
	}
	return &info, nil
}
