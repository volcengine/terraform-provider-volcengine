package customer_gateway

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

type VolcengineCustomerGatewayService struct {
	Client *ve.SdkClient
}

func NewCustomerGatewayService(c *ve.SdkClient) *VolcengineCustomerGatewayService {
	return &VolcengineCustomerGatewayService{
		Client: c,
	}
}

func (s *VolcengineCustomerGatewayService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineCustomerGatewayService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
		nameSet = make(map[string]bool)
	)
	if _, ok = m["CustomerGatewayNames.1"]; ok {
		i := 1
		for {
			filed := fmt.Sprintf("CustomerGatewayNames.%d", i)
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
		action := "DescribeCustomerGateways"
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
		results, err = ve.ObtainSdkValue("Result.CustomerGateways", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.CustomerGateways is not Slice")
		}
		return data, err
	})
	if err != nil || len(nameSet) == 0 {
		return gateways, err
	}

	res := make([]interface{}, 0)
	for _, gateway := range gateways {
		if !nameSet[gateway.(map[string]interface{})["CustomerGatewayName"].(string)] {
			continue
		}
		res = append(res, gateway)
	}
	return res, nil
}

func (s *VolcengineCustomerGatewayService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"CustomerGatewayIds.1": id,
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
		return data, fmt.Errorf("CustomerGateway %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineCustomerGatewayService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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
					return nil, "", fmt.Errorf("CustomerGateway status  error, status:%s", status.(string))
				}
			}
			//注意 返回的第一个参数不能为空 否则会一直等下去
			return demo, status.(string), err
		},
	}

}

func (VolcengineCustomerGatewayService) WithResourceResponseHandlers(v map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return v, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}

}

func (s *VolcengineCustomerGatewayService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callbacks := make([]ve.Callback, 0)

	// 创建CustomerGateway
	createCustomerGateway := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateCustomerGateway",
			ConvertMode: ve.RequestConvertAll,
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
				id, _ := ve.ObtainSdkValue("Result.CustomerGatewayId", *resp)
				d.SetId(id.(string))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Available"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	callbacks = append(callbacks, createCustomerGateway)

	return callbacks

}

func (s *VolcengineCustomerGatewayService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callbacks := make([]ve.Callback, 0)

	// 修改CustomerGateway
	modifyCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyCustomerGatewayAttributes",
			ConvertMode: ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"customer_gateway_name": {
					ConvertType: ve.ConvertDefault,
				},
				"description": {
					ConvertType: ve.ConvertDefault,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if len(*call.SdkParam) < 1 {
					return false, nil
				}
				(*call.SdkParam)["CustomerGatewayId"] = d.Id()
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
		},
	}
	callbacks = append(callbacks, modifyCallback)

	return callbacks
}

func (s *VolcengineCustomerGatewayService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteCustomerGateway",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"CustomerGatewayId": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
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
							return resource.NonRetryableError(fmt.Errorf("error on  reading CustomerGateway on delete %q, %w", d.Id(), callErr))
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
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 5*time.Minute)
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineCustomerGatewayService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "CustomerGatewayIds",
				ConvertType: ve.ConvertWithN,
			},
			"customer_gateway_names": {
				TargetField: "CustomerGatewayNames",
				ConvertType: ve.ConvertWithN,
			},
		},
		NameField:    "CustomerGatewayName",
		IdField:      "CustomerGatewayId",
		CollectField: "customer_gateways",
		ResponseConverts: map[string]ve.ResponseConvert{
			"CustomerGatewayId": {
				TargetField: "id",
				KeepDefault: true,
			},
		},
	}
}

func (s *VolcengineCustomerGatewayService) ReadResourceId(id string) string {
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
