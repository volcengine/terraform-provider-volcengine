package vpn_gateway

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-vestack/common"
	"github.com/volcengine/terraform-provider-vestack/logger"
)

type VestackVpnGatewayService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewVpnGatewayService(c *ve.SdkClient) *VestackVpnGatewayService {
	return &VestackVpnGatewayService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VestackVpnGatewayService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VestackVpnGatewayService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
		nameSet = make(map[string]bool)
	)
	if _, ok = m["VpnGatewayNames.1"]; ok {
		i := 1
		for {
			filed := fmt.Sprintf("VpnGatewayNames.%d", i)
			tmpName, ok := m[filed]
			if !ok {
				break
			}
			nameSet[tmpName.(string)] = true
			i++
			delete(m, filed)
		}
	}
	gateways, err := ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 20, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		vpnClient := s.Client.VpnClient
		action := "DescribeVpnGateways"
		logger.Debug(logger.ReqFormat, action, condition)
		if condition == nil {
			resp, err = vpnClient.DescribeVpnGatewaysCommon(nil)
			if err != nil {
				return data, err
			}
		} else {
			resp, err = vpnClient.DescribeVpnGatewaysCommon(&condition)
			if err != nil {
				return data, err
			}
		}
		logger.Debug(logger.RespFormat, action, resp)
		results, err = ve.ObtainSdkValue("Result.VpnGateways", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.VpnGateways is not Slice")
		}
		return data, err
	})
	if err != nil || len(nameSet) == 0 {
		return gateways, err
	}

	res := make([]interface{}, 0)
	for _, gateway := range gateways {
		if !nameSet[gateway.(map[string]interface{})["VpnGatewayName"].(string)] {
			continue
		}
		res = append(res, gateway)
	}
	return res, nil
}

func (s *VestackVpnGatewayService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"VpnGatewayIds.1": id,
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
		return data, fmt.Errorf("VpnGateway %s not exist ", id)
	}

	// 计费信息
	params := &map[string]interface{}{
		"VpnGatewayIds.1": id,
	}
	billingRes, err := s.Client.VpnClient.DescribeVpnGatewaysBillingCommon(params)
	logger.Debug(logger.AllFormat, "DescribeVpnGatewaysBilling", params, billingRes, err)
	if err != nil {
		return data, err
	}
	tmpRes, err := ve.ObtainSdkValue("Result.VpnGateways", *billingRes)
	if err != nil {
		return data, err
	}
	if tmpRes == nil {
		results = []interface{}{}
	}
	tmpData, ok := tmpRes.([]interface{})
	if !ok {
		return data, errors.New("Result.VpnGateways is not Slice")
	}
	if len(tmpData) == 0 {
		return data, fmt.Errorf("VpnGatewaysBilling %s not exist ", id)
	}
	data["RenewType"] = renewTypeResponseConvert(tmpData[0].(map[string]interface{})["RenewType"])
	data["RemainRenewTimes"] = int(tmpData[0].(map[string]interface{})["RemainRenewTimes"].(float64))

	return data, err
}

func (s *VestackVpnGatewayService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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
					return nil, "", fmt.Errorf("VpnGateway  status  error, status:%s", status.(string))
				}
			}
			//注意 返回的第一个参数不能为空 否则会一直等下去
			return demo, status.(string), err
		},
	}

}

func (VestackVpnGatewayService) WithResourceResponseHandlers(v map[string]interface{}) []ve.ResourceResponseHandler {
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

func (s *VestackVpnGatewayService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callbacks := make([]ve.Callback, 0)

	// 创建vpn
	createVpnGateway := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateVpnGateway",
			ConvertMode: ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"bandwidth": {
					ConvertType: ve.ConvertDefault,
				},
				"description": {
					ConvertType: ve.ConvertDefault,
				},
				"period": {
					ConvertType: ve.ConvertDefault,
				},
				"period_unit": {
					ConvertType: ve.ConvertDefault,
				},
				"subnet_id": {
					ConvertType: ve.ConvertDefault,
				},
				"vpc_id": {
					ConvertType: ve.ConvertDefault,
				},
				"vpn_gateway_name": {
					ConvertType: ve.ConvertDefault,
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				(*call.SdkParam)["BillingType"] = 1
				return s.Client.VpnClient.CreateVpnGatewayCommon(call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				//注意 获取内容 这个地方不能是指针 需要转一次
				id, _ := ve.ObtainSdkValue("Result.VpnGatewayId", *resp)
				d.SetId(id.(string))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Available"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	callbacks = append(callbacks, createVpnGateway)

	// 修改计费方式
	if resourceData.Get("renew_type") != nil && resourceData.Get("renew_type").(string) != "ManualRenew" {
		callbacks = append(callbacks, s.setVpnGatewayRenewal())
	}

	return callbacks

}

func (s *VestackVpnGatewayService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callbacks := make([]ve.Callback, 0)

	// 修改vpn
	modifyCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyVpnGatewayAttributes",
			ConvertMode: ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"vpn_gateway_name": {
					ConvertType: ve.ConvertDefault,
				},
				"description": {
					ConvertType: ve.ConvertDefault,
				},
				"bandwidth": {
					ConvertType: ve.ConvertDefault,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if len(*call.SdkParam) < 1 {
					return false, nil
				}
				(*call.SdkParam)["VpnGatewayId"] = d.Id()
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.VpnClient.ModifyVpnGatewayAttributesCommon(call.SdkParam)
			},
		},
	}
	callbacks = append(callbacks, modifyCallback)

	// 修改计费方式
	if resourceData.Get("renew_type") != nil {
		if resourceData.HasChange("renew_type") ||
			(resourceData.Get("renew_type").(string) == "AutoRenew" &&
				(resourceData.HasChange("renew_period") || resourceData.HasChange("remain_renew_times"))) {
			callbacks = append(callbacks, s.setVpnGatewayRenewal())
		}
	}

	return callbacks
}

func (s *VestackVpnGatewayService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteVpnGateway",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"VpnGatewayId": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				//logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				// todo 打印前台提示日志
				log.Println("[WARN] Cannot destroy resource vestack_vpn_gateway. Terraform will remove this resource from the state file, however resources may remain.")
				//return s.Client.VpnClient.DeleteVpnGatewayCommon(call.SdkParam)
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
							return resource.NonRetryableError(fmt.Errorf("error on  reading VpnGateway on delete %q, %w", d.Id(), callErr))
						}
					}
					resp, callErr := call.ExecuteCall(d, client, call)
					logger.Debug(logger.AllFormat, call.Action, call.SdkParam, resp, callErr)
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

func (s *VestackVpnGatewayService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "VpnGatewayIds",
				ConvertType: ve.ConvertWithN,
			},
			"vpn_gateway_names": {
				TargetField: "VpnGatewayNames",
				ConvertType: ve.ConvertWithN,
			},
		},
		NameField:    "VpnGatewayName",
		IdField:      "VpnGatewayId",
		CollectField: "vpn_gateways",
		ResponseConverts: map[string]ve.ResponseConvert{
			"VpnGatewayId": {
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

func (s *VestackVpnGatewayService) ReadResourceId(id string) string {
	return id
}

func (s *VestackVpnGatewayService) setVpnGatewayRenewal() ve.Callback {
	return ve.Callback{
		Call: ve.SdkCall{
			Action:      "SetVpnGatewayRenewal",
			ConvertMode: ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"renew_period": {
					ConvertType: ve.ConvertDefault,
					ForceGet:    true,
				},
				"renew_type": {
					TargetField: "RenewType",
					Convert:     renewTypeRequestConvert,
					ForceGet:    true,
				},
				"remain_renew_times": {
					ConvertType: ve.ConvertDefault,
					ForceGet:    true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if len(*call.SdkParam) < 1 {
					return false, nil
				}
				(*call.SdkParam)["VpnGatewayId"] = d.Id()
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.VpnClient.SetVpnGatewayRenewalCommon(call.SdkParam)
			},
		},
	}
}
