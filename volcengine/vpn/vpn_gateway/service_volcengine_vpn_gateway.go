package vpn_gateway

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineVpnGatewayService struct {
	Client *ve.SdkClient
}

func NewVpnGatewayService(c *ve.SdkClient) *VolcengineVpnGatewayService {
	return &VolcengineVpnGatewayService{
		Client: c,
	}
}

func (s *VolcengineVpnGatewayService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineVpnGatewayService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
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
		universalClient := s.Client.UniversalClient
		action := "DescribeVpnGateways"
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

func (s *VolcengineVpnGatewayService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
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
	data["RenewType"] = tmpData[0].(map[string]interface{})["RenewType"]
	data["RemainRenewTimes"] = int(tmpData[0].(map[string]interface{})["RemainRenewTimes"].(float64))

	return data, err
}

func (s *VolcengineVpnGatewayService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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

func (VolcengineVpnGatewayService) WithResourceResponseHandlers(v map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		//if v["BillingType"].(float64) == 1 {
		//	ct, _ := time.Parse("2006-01-02T15:04:05", v["CreationTime"].(string)[0:strings.Index(v["CreationTime"].(string), "+")])
		//	et, _ := time.Parse("2006-01-02T15:04:05", v["ExpiredTime"].(string)[0:strings.Index(v["ExpiredTime"].(string), "+")])
		//	y := et.Year() - ct.Year()
		//	m := et.Month() - ct.Month()
		//	v["Period"] = y*12 + int(m)
		//}
		return v, map[string]ve.ResponseConvert{
			"BillingType": {
				TargetField: "billing_type",
				Convert:     billingTypeResponseConvert,
			},
			"RenewType": {
				TargetField: "renew_type",
				Convert:     renewTypeResponseConvert,
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}

}

func (s *VolcengineVpnGatewayService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callbacks := make([]ve.Callback, 0)

	// 创建VpnGateway
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
				"billing_type": {
					TargetField: "BillingType",
					Convert:     billingTypeRequestConvert,
				},
				"tags": {
					TargetField: "Tags",
					ConvertType: ve.ConvertListN,
				},
				"project_name": {
					ConvertType: ve.ConvertDefault,
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
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

	return callbacks

}

func (s *VolcengineVpnGatewayService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callbacks := make([]ve.Callback, 0)

	// 修改vpnGateway
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
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	callbacks = append(callbacks, modifyCallback)

	// 续费时长
	if resourceData.Get("renew_type").(string) == "ManualRenew" &&
		(resourceData.HasChange("period") || resourceData.HasChange("period_unit")) {
		renewVpnGateway := ve.Callback{
			Call: ve.SdkCall{
				Action:      "RenewVpnGateway",
				ConvertMode: ve.RequestConvertInConvert,
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["VpnGatewayId"] = d.Id()
					(*call.SdkParam)["PeriodUnit"] = "Month"
					// 计算差值
					month := 0
					o, n := d.GetChange("period")
					oldUnit, newUnit := d.GetChange("period_unit")
					// 默认为Month，不填的统一设置为Month进行计算
					if len(oldUnit.(string)) == 0 {
						oldUnit = "Month"
					}
					if len(newUnit.(string)) == 0 {
						newUnit = "Month"
					}
					if oldUnit.(string) == "Month" && newUnit.(string) == "Year" {
						// 月改年
						month = n.(int)*12 - o.(int)
					} else if oldUnit.(string) == "Year" && newUnit.(string) == "Month" {
						// 年改月
						month = n.(int) - o.(int)*12
					} else if oldUnit == "Year" && newUnit == "Year" {
						// 单位没改则period必然发生了改变，直接计算差值
						// 需要区分单位是年还是月
						month = n.(int)*12 - o.(int)*12
					} else {
						// 不填改月 或者 一直是月没改变
						month = n.(int) - o.(int)
					}
					// 如果差值小于等于0则直接报错
					if month <= 0 {
						return false, fmt.Errorf("The difference modified by Period must be greater than 0. ")
					}
					(*call.SdkParam)["Period"] = month
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
				AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
					return nil
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{"Available"},
					Timeout: resourceData.Timeout(schema.TimeoutUpdate),
				},
			},
		}
		callbacks = append(callbacks, renewVpnGateway)
	}

	return callbacks
}

func (s *VolcengineVpnGatewayService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteVpnGateway",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"VpnGatewayId": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				// todo 打印前台提示日志
				log.Println("[WARN] Terraform will unsubscribe the resource.")
				//return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
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

func (s *VolcengineVpnGatewayService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
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

func (s *VolcengineVpnGatewayService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "vpn",
		Action:      actionName,
		Version:     "2020-04-01",
		HttpMethod:  ve.GET,
		ContentType: ve.Default,
	}
}

func (s *VolcengineVpnGatewayService) ProjectTrn() *ve.ProjectTrn {
	return &ve.ProjectTrn{
		ServiceName:          "vpn",
		ResourceType:         "vpngateway",
		ProjectResponseField: "ProjectName",
		ProjectSchemaField:   "project_name",
	}
}

func (s *VolcengineVpnGatewayService) UnsubscribeInfo(resourceData *schema.ResourceData, resource *schema.Resource) (*ve.UnsubscribeInfo, error) {
	info := ve.UnsubscribeInfo{
		InstanceId: s.ReadResourceId(resourceData.Id()),
	}
	info.NeedUnsubscribe = true
	info.Products = []string{"VPN"}
	return &info, nil
}
