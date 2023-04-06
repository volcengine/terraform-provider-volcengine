package cr_vpc_endpoint

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineCrVpcEndpointService struct {
	Client *ve.SdkClient
}

func (v *VolcengineCrVpcEndpointService) GetClient() *ve.SdkClient {
	return v.Client
}

func (v *VolcengineCrVpcEndpointService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
	)
	action := "GetVpcEndpoint"
	logger.Debug(logger.ReqFormat, action, condition)
	if condition == nil {
		resp, err = v.Client.UniversalClient.DoCall(getUniversalInfo(action), nil)
		if err != nil {
			return data, err
		}
	} else {
		resp, err = v.Client.UniversalClient.DoCall(getUniversalInfo(action), &condition)
		if err != nil {
			return data, err
		}
	}
	logger.Debug(logger.RespFormat, action, resp)
	results, err = ve.ObtainSdkValue("Result", *resp)
	if err != nil {
		return data, err
	}
	if results == nil {
		return data, fmt.Errorf("GetPublicEndpoint return an empty result")
	}
	registry, err := ve.ObtainSdkValue("Result.Registry", *resp)
	if err != nil {
		return data, err
	}
	vpcs, err := ve.ObtainSdkValue("Result.Vpcs", *resp)
	if err != nil {
		return data, err
	}
	endpoints := map[string]interface{}{
		"Registry": registry,
		"Vpcs":     vpcs,
	}
	return []interface{}{endpoints}, err
}

func (v *VolcengineCrVpcEndpointService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	registry := resourceData.Get("registry").(string)
	req := map[string]interface{}{
		"Registry": registry,
	}
	results, err = v.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, r := range results {
		if data, ok = r.(map[string]interface{}); !ok {
			return data, errors.New("GetVpcEndpoint value is not a map")
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("cr vpc endpoint %s is not exist", id)
	}
	return data, err
}

func (v *VolcengineCrVpcEndpointService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (v *VolcengineCrVpcEndpointService) WithResourceResponseHandlers(m map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return m, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (v *VolcengineCrVpcEndpointService) CreateResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateVpcEndpoint",
			ContentType: ve.ContentTypeJson,
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"vpcs": {
					ConvertType: ve.ConvertJsonObjectArray,
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				var err error
				call.SdkParam, err = v.integrateVpcParams(d, call)
				if err != nil {
					return nil, err
				}
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return v.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id := "crVpcEndpoint:" + d.Get("registry").(string)
				d.SetId(id)
				return nil
			},
			AfterRefresh: v.RefreshVpcStatus(),
		},
	}
	return []ve.Callback{callback}
}

func (v *VolcengineCrVpcEndpointService) ModifyResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateVpcEndpoint",
			ContentType: ve.ContentTypeJson,
			ConvertMode: ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"registry": {
					ConvertType: ve.ConvertDefault,
					ForceGet:    true,
				},
				"vpcs": {
					ConvertType: ve.ConvertJsonObjectArray,
					ForceGet:    true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				var (
					ok bool
				)
				logger.DebugInfo("sdk param : ", *call.SdkParam)
				paramMap := *call.SdkParam
				for key, value := range paramMap {
					if strings.Contains(key, "AccountId") {
						value, ok = value.(int)
						if !ok {
							return false, fmt.Errorf("sdk param account id is not a integer")
						}
						// 删除force get导致的用户ID为0
						if value == 0 {
							delete(paramMap, key)
						}
					} else if strings.Contains(key, "Subnet") {
						value, ok = value.(string)
						if !ok {
							return false, fmt.Errorf("sdk param subnet is not a string")
						}
						// 删除force get导致的子网为空
						if len(value.(string)) == 0 {
							delete(paramMap, key)
						}
					} else if strings.Contains(key, "VpcId") {
						value, ok = value.(string)
						if !ok {
							return false, fmt.Errorf("sdk param vpcId is not a string")
						}
						// 删除force get导致的vpcId为空
						if len(value.(string)) == 0 {
							delete(paramMap, key)
						}
					}
				}
				*call.SdkParam = paramMap
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				vpcs := (*call.SdkParam)["Vpcs"]
				if vpcs == nil || len(vpcs.([]interface{})) == 0 {
					(*call.SdkParam)["Vpcs"] = []interface{}{}
				}
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return v.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterRefresh: v.RefreshVpcStatus(),
		},
	}
	return []ve.Callback{callback}
}

func (v *VolcengineCrVpcEndpointService) RemoveResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateVpcEndpoint",
			ContentType: ve.ContentTypeJson,
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["Registry"] = d.Get("registry")
				(*call.SdkParam)["Vpcs"] = []interface{}{}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return v.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterRefresh: v.RefreshVpcStatus(),
		},
	}
	return []ve.Callback{callback}
}

func (v *VolcengineCrVpcEndpointService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"statuses": {
				TargetField: "Filter.Statuses",
				ConvertType: ve.ConvertJsonObjectArray,
			},
		},
		ContentType:  ve.ContentTypeJson,
		CollectField: "endpoints",
	}
}

func (v *VolcengineCrVpcEndpointService) ReadResourceId(id string) string {
	return id
}

func NewCrVpcEndpointService(c *ve.SdkClient) *VolcengineCrVpcEndpointService {
	return &VolcengineCrVpcEndpointService{
		Client: c,
	}
}

func (v *VolcengineCrVpcEndpointService) RefreshVpcStatus() ve.CallFunc {
	return func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) error {
		if err := v.checkVpcStatus(d); err != nil {
			return err
		}
		return nil
	}
}

func (v *VolcengineCrVpcEndpointService) checkVpcStatus(d *schema.ResourceData) error {
	return resource.Retry(10*time.Minute, func() *resource.RetryError {
		var (
			vpcInfos []interface{}
			vpcMap   map[string]interface{}
			status   string
			ok       bool
			statusOk bool
		)
		resp, err := v.ReadResource(d, d.Id())
		if err != nil {
			return resource.NonRetryableError(err)
		}
		vpcs := resp["Vpcs"]
		// 删除时直接return
		if vpcs == nil {
			return nil
		}
		vpcInfos, ok = vpcs.([]interface{})
		if !ok {
			return resource.NonRetryableError(fmt.Errorf("vpcs is not an array"))
		}
		statusOk = true
		for _, vpc := range vpcInfos {
			vpcMap, ok = vpc.(map[string]interface{})
			if !ok {
				return resource.NonRetryableError(fmt.Errorf("vpc is not a map"))
			}
			status, ok = vpcMap["Status"].(string)
			if !ok {
				return resource.NonRetryableError(fmt.Errorf("get vpc status err"))
			}
			if status != "Enabled" {
				statusOk = false
				break
			}
		}
		if !statusOk {
			logger.DebugInfo("still in waiting")
			return resource.RetryableError(fmt.Errorf("vpc endpoint still in waiting"))
		}
		return nil
	})
}

func (v *VolcengineCrVpcEndpointService) integrateVpcParams(d *schema.ResourceData, call ve.SdkCall) (*map[string]interface{}, error) {
	var (
		vpcIdsMap = make(map[string]interface{})
		vpcs      []interface{}
		newVpc    []interface{}
		oldVpc    []interface{}
		ok        bool
	)
	logger.DebugInfo("sdk param : ", *call.SdkParam)
	/*
		0. 预备传入的参数 vpcs
		1. 读取用户已有信息，如空直接返回，拿到用户已有vpc信息，加入去重map，加入vpcs
		2. 取现有tf内定义的vpc信息，如存在VpcId，则比较是否跟map中重复，如不重复，加入vpcs
		3. 覆盖原sdk param vpcs
	*/
	resp, err := v.ReadResource(d, d.Get("registry").(string))
	if err != nil {
		return nil, fmt.Errorf("read resource err")
	}

	old := resp["Vpcs"]
	if old == nil {
		return call.SdkParam, nil
	}
	oldVpc, ok = old.([]interface{})
	if !ok {
		return nil, fmt.Errorf("vpcs is not a slice")
	}
	for _, o := range oldVpc {
		oldVpcMap, ok := o.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("vpc is not a map")
		}
		vpcId, ok := oldVpcMap["VpcId"].(string)
		if !ok {
			return nil, fmt.Errorf("get vpc id err")
		}
		vpcIdsMap[vpcId] = o
		vpcs = append(vpcs, map[string]interface{}{
			"VpcId":     vpcId,
			"SubnetId":  oldVpcMap["SubnetId"].(string),
			"AccountId": (int)(oldVpcMap["AccountId"].(float64)),
		})
	}
	logger.DebugInfo("old vpc map : ", vpcs)

	newVpcInterface := (*call.SdkParam)["Vpcs"]
	if newVpcInterface == nil {
		(*call.SdkParam)["Vpcs"] = vpcs
		return call.SdkParam, nil
	}
	newVpc, ok = newVpcInterface.([]interface{})
	if !ok {
		return nil, fmt.Errorf("vpcs is not a slice")
	}
	for _, n := range newVpc {
		newVpcMap, ok := n.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("vpc is not a map")
		}
		vpcId, ok := newVpcMap["VpcId"].(string)
		if !ok {
			return nil, fmt.Errorf("vpc id is missing")
		}
		// 跟用户已有的重复
		if _, ok := vpcIdsMap[vpcId]; ok {
			continue
		}
		vpcIdsMap[vpcId] = n
		vpcs = append(vpcs, n)
	}

	logger.DebugInfo("sdk param vpcs change : ", vpcs)
	(*call.SdkParam)["Vpcs"] = vpcs
	return call.SdkParam, nil
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "cr",
		Version:     "2022-05-12",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}
