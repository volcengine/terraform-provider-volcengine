package nat_gateway

import (
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineNatGatewayService struct {
	Client     *ve.SdkClient
}

func NewNatGatewayService(c *ve.SdkClient) *VolcengineNatGatewayService {
	return &VolcengineNatGatewayService{
		Client:     c,
	}
}

func (s *VolcengineNatGatewayService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineNatGatewayService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(condition, "PageSize", "PageNumber", 20, 1, func(m map[string]interface{}) ([]interface{}, error) {
		natGateway := s.Client.NatClient
		action := "DescribeNatGateways"
		logger.Debug(logger.ReqFormat, action, condition)
		if condition == nil {
			resp, err = natGateway.DescribeNatGatewaysCommon(nil)
			if err != nil {
				return data, err
			}
		} else {
			resp, err = natGateway.DescribeNatGatewaysCommon(&condition)
			if err != nil {
				return data, err
			}
		}
		logger.Debug(logger.RespFormat, "testDescribeNatGateways", condition, *resp)

		results, err = ve.ObtainSdkValue("Result.NatGateways", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.NatGateways is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineNatGatewayService) ReadResource(resourceData *schema.ResourceData, natGatewayId string) (data map[string]interface{}, err error) {
	var (
		results     []interface{}
		ok          bool
		billingType interface{}
	)
	if natGatewayId == "" {
		natGatewayId = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"NatGatewayIds.1": natGatewayId,
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
		return data, fmt.Errorf("NatGateway %s not exist ", natGatewayId)
	}

	billingType, err = ve.ObtainSdkValue("BillingType", data)
	if err != nil {
		return data, err
	}

	if billingType, ok = billingType.(float64); !ok {
		return data, errors.New("BillingType is invalid")
	}

	if billingType.(float64) == 1 {
		// prepaid not support
		return data, fmt.Errorf("Prepaid nat gateway %s is not supported ", natGatewayId)
	}

	return data, err
}

func (s *VolcengineNatGatewayService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      3 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				demo   map[string]interface{}
				status interface{}
			)

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

			//注意 返回的第一个参数不能为空 否则会一直等下去
			return demo, status.(string), err
		},
	}

}

func (VolcengineNatGatewayService) WithResourceResponseHandlers(natGateway map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return natGateway, map[string]ve.ResponseConvert{
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
						return "PostPaid"
					}
					return i
				},
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}

}

func (s *VolcengineNatGatewayService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateNatGateway",
			ConvertMode: ve.RequestConvertAll,
			LockId: func(d *schema.ResourceData) string {
				return d.Get("vpc_id").(string)
			},
			Convert: map[string]ve.RequestConvert{
				"billing_type": {
					TargetField: "BillingType",
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
						}
						return i
					},
				},
				"tags": {
					TargetField: "Tags",
					ConvertType: ve.ConvertListN,
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				//创建natGateway
				return s.Client.NatClient.CreateNatGatewayCommon(call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				//注意 获取内容 这个地方不能是指针 需要转一次
				id, _ := ve.ObtainSdkValue("Result.NatGatewayId", *resp)
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

func (s *VolcengineNatGatewayService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	var callbacks []ve.Callback

	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyNatGatewayAttributes",
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"billing_type": {
					TargetField: "BillingType",
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
						}
						return i
					},
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["NatGatewayId"] = d.Id()
				delete(*call.SdkParam, "Tags")
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				//修改natGateway属性
				return s.Client.NatClient.ModifyNatGatewayAttributesCommon(call.SdkParam)
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Available"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	callbacks = append(callbacks, callback)

	// 更新Tags
	setResourceTagsCallbacks := ve.SetResourceTags(s.Client, "TagResources", "UntagResources", "ngw", resourceData, getUniversalInfo)
	callbacks = append(callbacks, setResourceTagsCallbacks...)

	return callbacks
}

func (s *VolcengineNatGatewayService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	id := resourceData.Id()
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteNatGateway",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"NatGatewayId": id,
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				//删除NatGateway
				return s.Client.NatClient.DeleteNatGatewayCommon(call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				//由于异步删除问题 这里补充一个轮询查询(临时解决方案)
				return resource.Retry(3*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, id)
					//能查询成功代表还在删除中，重试
					if callErr == nil {
						return resource.RetryableError(fmt.Errorf("Nat still in remove "))
					} else {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(callErr)
						}
					}
				})
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on  reading nat gateway on delete %q, %w", d.Id(), callErr))
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

func (s *VolcengineNatGatewayService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "NatGatewayIds",
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
		NameField:    "NatGatewayName",
		IdField:      "NatGatewayId",
		CollectField: "nat_gateways",
		ResponseConverts: map[string]ve.ResponseConvert{
			"NatGatewayId": {
				TargetField: "id",
				KeepDefault: true,
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
						return "PostPaid"
					}
					return i
				},
			},
		},
	}
}

func (s *VolcengineNatGatewayService) ReadResourceId(id string) string {
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