package eip_address

import (
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineEipAddressService struct {
	Client *ve.SdkClient
}

func NewEipAddressService(c *ve.SdkClient) *VolcengineEipAddressService {
	return &VolcengineEipAddressService{
		Client: c,
	}
}

func (s *VolcengineEipAddressService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineEipAddressService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 20, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		vpcClient := s.Client.VpcClient
		action := "DescribeEipAddresses"
		logger.Debug(logger.ReqFormat, action, condition)
		if condition == nil {
			resp, err = vpcClient.DescribeEipAddressesCommon(nil)
			if err != nil {
				return data, err
			}
		} else {
			resp, err = vpcClient.DescribeEipAddressesCommon(&condition)
			if err != nil {
				return data, err
			}
		}

		results, err = ve.ObtainSdkValue("Result.EipAddresses", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.EipAddresses is not Slice")
		}
		data, err = RemoveSystemTags(data)
		return data, err
	})
}

func (s *VolcengineEipAddressService) ReadResource(resourceData *schema.ResourceData, allocationId string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if allocationId == "" {
		allocationId = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"AllocationIds.1": allocationId,
	}
	results, err = s.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, v := range results {
		if data, ok = v.(map[string]interface{}); !ok {
			return data, errors.New("value is not map")
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("eip address %s not exist ", allocationId)
	}
	if data["BillingType"].(float64) == 1 {
		return data, fmt.Errorf("not support PrePaid eip address")
	}
	return data, err
}

func (s *VolcengineEipAddressService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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
			status, err = ve.ObtainSdkValue("Status", demo)
			if err != nil {
				return nil, "", err
			}
			for _, v := range failStates {
				if v == status.(string) {
					return nil, "", fmt.Errorf("eip address status error, status:%s", status.(string))
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

func (VolcengineEipAddressService) WithResourceResponseHandlers(eip map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return eip, map[string]ve.ResponseConvert{
			"BillingType": {
				TargetField: "billing_type",
				Convert:     billingTypeResponseConvert,
			},
			"ISP": {
				TargetField: "isp",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}

}

func (s *VolcengineEipAddressService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "AllocateEipAddress",
			ConvertMode: ve.RequestConvertAll,
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.VpcClient.AllocateEipAddressCommon(call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("Result.AllocationId", *resp)
				d.SetId(id.(string))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Available"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
			Convert: map[string]ve.RequestConvert{
				"billing_type": {
					TargetField: "BillingType",
					Convert:     billingTypeRequestConvert,
				},
				"isp": {
					TargetField: "ISP",
				},
				"tags": {
					TargetField: "Tags",
					ConvertType: ve.ConvertListN,
				},
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineEipAddressService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	//project
	projectCallback := ve.NewProjectService(s.Client).ModifyProjectOld(ve.ProjectTrn{
		ResourceType: "eip",
		ResourceID:   resourceData.Id(),
		ServiceName:  "vpc",
	}, resourceData, resource, "project_name",
		&ve.StateRefresh{
			Target:  []string{"Available"},
			Timeout: resourceData.Timeout(schema.TimeoutCreate),
		})

	callbacks = append(callbacks, projectCallback...)

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyEipAddressAttributes",
			ConvertMode: ve.RequestConvertAll,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if len(*call.SdkParam) > 0 {
					(*call.SdkParam)["AllocationId"] = d.Id()
					delete(*call.SdkParam, "Tags")
					return true, nil
				}
				return false, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.VpcClient.ModifyEipAddressAttributesCommon(call.SdkParam)
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Available"},
				Timeout: resourceData.Timeout(schema.TimeoutUpdate),
			},
			Convert: map[string]ve.RequestConvert{
				"billing_type": {
					Ignore: true,
				},
				"isp": {
					Ignore: true,
				},
			},
		},
	}

	callbacks = append(callbacks, callback)

	if resourceData.HasChange("billing_type") {
		chargeTypeCall := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ConvertEipAddressBillingType",
				ConvertMode: ve.RequestConvertInConvert,
				Convert: map[string]ve.RequestConvert{
					"billing_type": {
						Convert: billingTypeRequestConvert,
					},
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					if len(*call.SdkParam) > 0 {
						(*call.SdkParam)["AllocationId"] = d.Id()
						return true, nil
					}
					return false, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{"Available"},
					Timeout: resourceData.Timeout(schema.TimeoutUpdate),
				},
			},
		}
		callbacks = append(callbacks, chargeTypeCall)
	}

	// 更新Tags
	setResourceTagsCallbacks := ve.SetResourceTags(s.Client, "TagResources", "UntagResources", "eip", resourceData, getUniversalInfo)
	callbacks = append(callbacks, setResourceTagsCallbacks...)

	return callbacks
}

func (s *VolcengineEipAddressService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ReleaseEipAddress",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"AllocationId": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.VpcClient.ReleaseEipAddressCommon(call.SdkParam)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading eip address on delete %q, %w", d.Id(), callErr))
						}
					}
					_, callErr = call.ExecuteCall(d, client, call)
					if callErr == nil {
						return nil
					}
					return resource.RetryableError(callErr)
				})
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 3*time.Minute)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineEipAddressService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "AllocationIds",
				ConvertType: ve.ConvertWithN,
			},
			"eip_addresses": {
				TargetField: "EipAddresses",
				ConvertType: ve.ConvertWithN,
			},
			"isp": {
				TargetField: "ISP",
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
		NameField:    "Name",
		IdField:      "AllocationId",
		CollectField: "addresses",
		ResponseConverts: map[string]ve.ResponseConvert{
			"AllocationId": {
				TargetField: "id",
				KeepDefault: true,
			},
			"ISP": {
				TargetField: "isp",
			},
			"BillingType": {
				TargetField: "billing_type",
				Convert:     billingTypeResponseConvert,
			},
		},
	}
}

func (s *VolcengineEipAddressService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "vpc",
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		Action:      actionName,
	}
}
