package vpn_connection

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-vestack/common"
	"github.com/volcengine/terraform-provider-vestack/logger"
)

type VestackVpnConnectionService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewVpnConnectionService(c *ve.SdkClient) *VestackVpnConnectionService {
	return &VestackVpnConnectionService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VestackVpnConnectionService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VestackVpnConnectionService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
		nameSet = make(map[string]bool)
	)
	if _, ok = m["VpnConnectionNames.1"]; ok {
		i := 1
		for {
			filed := fmt.Sprintf("VpnConnectionNames.%d", i)
			tmpName, ok := m[filed]
			if !ok {
				break
			}
			nameSet[tmpName.(string)] = true
			i++
			delete(m, filed)
		}
	}
	connections, err := ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 20, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		universalClient := s.Client.UniversalClient
		action := "DescribeVpnConnections"
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
		results, err = ve.ObtainSdkValue("Result.VpnConnections", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.VpnConnections is not Slice")
		}
		return data, err
	})
	if err != nil || len(nameSet) == 0 {
		return connections, err
	}

	res := make([]interface{}, 0)
	for _, connection := range connections {
		if !nameSet[connection.(map[string]interface{})["VpnConnectionName"].(string)] {
			continue
		}
		res = append(res, connection)
	}
	return res, nil
}

func (s *VestackVpnConnectionService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"VpnConnectionIds.1": id,
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
		return data, fmt.Errorf("VpnConnection %s not exist ", id)
	}
	return data, err
}

func (s *VestackVpnConnectionService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
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
					return nil, "", fmt.Errorf("VpnConnection  status  error, status:%s", status.(string))
				}
			}
			//注意 返回的第一个参数不能为空 否则会一直等下去
			return demo, status.(string), err
		},
	}

}

func (VestackVpnConnectionService) WithResourceResponseHandlers(v map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return v, map[string]ve.ResponseConvert{
			"IkeConfig.Psk": {
				TargetField: "ike_config_psk",
			},
			"IkeConfig.Version": {
				TargetField: "ike_config_version",
			},
			"IkeConfig.Mode": {
				TargetField: "ike_config_mode",
			},
			"IkeConfig.EncAlg": {
				TargetField: "ike_config_enc_alg",
			},
			"IkeConfig.AuthAlg": {
				TargetField: "ike_config_auth_alg",
			},
			"IkeConfig.DhGroup": {
				TargetField: "ike_config_dh_group",
			},
			"IkeConfig.Lifetime": {
				TargetField: "ike_config_lifetime",
			},
			"IkeConfig.LocalId": {
				TargetField: "ike_config_local_id",
			},
			"IkeConfig.RemoteId": {
				TargetField: "ike_config_remote_id",
			},
			"IpsecConfig.EncAlg": {
				TargetField: "ipsec_config_enc_alg",
			},
			"IpsecConfig.AuthAlg": {
				TargetField: "ipsec_config_auth_alg",
			},
			"IpsecConfig.DhGroup": {
				TargetField: "ipsec_config_dh_group",
			},
			"IpsecConfig.Lifetime": {
				TargetField: "ipsec_config_lifetime",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}

}

func (s *VestackVpnConnectionService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callbacks := make([]ve.Callback, 0)

	// 创建vpn
	createVpnConnection := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateVpnConnection",
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"local_subnet": {
					TargetField: "LocalSubnet",
					ConvertType: ve.ConvertWithN,
				},
				"remote_subnet": {
					TargetField: "RemoteSubnet",
					ConvertType: ve.ConvertWithN,
				},
				"ike_config_psk": {
					TargetField: "IkeConfig.Psk",
					ConvertType: ve.ConvertDefault,
				},
				"ike_config_version": {
					TargetField: "IkeConfig.Version",
					ConvertType: ve.ConvertDefault,
				},
				"ike_config_mode": {
					TargetField: "IkeConfig.Mode",
					ConvertType: ve.ConvertDefault,
				},
				"ike_config_enc_alg": {
					TargetField: "IkeConfig.EncAlg",
					ConvertType: ve.ConvertDefault,
				},
				"ike_config_auth_alg": {
					TargetField: "IkeConfig.AuthAlg",
					ConvertType: ve.ConvertDefault,
				},
				"ike_config_dh_group": {
					TargetField: "IkeConfig.DhGroup",
					ConvertType: ve.ConvertDefault,
				},
				"ike_config_lifetime": {
					TargetField: "IkeConfig.Lifetime",
					ConvertType: ve.ConvertDefault,
				},
				"ike_config_local_id": {
					TargetField: "IkeConfig.LocalId",
					ConvertType: ve.ConvertDefault,
				},
				"ike_config_remote_id": {
					TargetField: "IkeConfig.RemoteId",
					ConvertType: ve.ConvertDefault,
				},
				"ipsec_config_enc_alg": {
					TargetField: "IpsecConfig.EncAlg",
					ConvertType: ve.ConvertDefault,
				},
				"ipsec_config_auth_alg": {
					TargetField: "IpsecConfig.AuthAlg",
					ConvertType: ve.ConvertDefault,
				},
				"ipsec_config_dh_group": {
					TargetField: "IpsecConfig.DhGroup",
					ConvertType: ve.ConvertDefault,
				},
				"ipsec_config_lifetime": {
					TargetField: "IpsecConfig.Lifetime",
					ConvertType: ve.ConvertDefault,
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
				id, _ := ve.ObtainSdkValue("Result.VpnConnectionId", *resp)
				d.SetId(id.(string))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"Available"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	callbacks = append(callbacks, createVpnConnection)

	return callbacks

}

func (s *VestackVpnConnectionService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callbacks := make([]ve.Callback, 0)

	// 修改vpn
	modifyCallback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyVpnConnectionAttributes",
			ConvertMode: ve.RequestConvertInConvert,
			Convert: map[string]ve.RequestConvert{
				"vpn_connection_name": {
					ConvertType: ve.ConvertDefault,
				},
				"description": {
					ConvertType: ve.ConvertDefault,
				},
				"local_subnet": {
					TargetField: "LocalSubnet",
					ConvertType: ve.ConvertWithN,
				},
				"remote_subnet": {
					TargetField: "RemoteSubnet",
					ConvertType: ve.ConvertWithN,
				},
				"nat_traversal": {
					ConvertType: ve.ConvertDefault,
				},
				"dpd_action": {
					ConvertType: ve.ConvertDefault,
				},
				"ike_config_psk": {
					TargetField: "IkeConfig.Psk",
					ConvertType: ve.ConvertDefault,
				},
				"ike_config_version": {
					TargetField: "IkeConfig.Version",
					ConvertType: ve.ConvertDefault,
				},
				"ike_config_mode": {
					TargetField: "IkeConfig.Mode",
					ConvertType: ve.ConvertDefault,
				},
				"ike_config_enc_alg": {
					TargetField: "IkeConfig.EncAlg",
					ConvertType: ve.ConvertDefault,
				},
				"ike_config_auth_alg": {
					TargetField: "IkeConfig.AuthAlg",
					ConvertType: ve.ConvertDefault,
				},
				"ike_config_dh_group": {
					TargetField: "IkeConfig.DhGroup",
					ConvertType: ve.ConvertDefault,
				},
				"ike_config_lifetime": {
					TargetField: "IkeConfig.Lifetime",
					ConvertType: ve.ConvertDefault,
				},
				"ike_config_local_id": {
					TargetField: "IkeConfig.LocalId",
					ConvertType: ve.ConvertDefault,
				},
				"ike_config_remote_id": {
					TargetField: "IkeConfig.RemoteId",
					ConvertType: ve.ConvertDefault,
				},
				"ipsec_config_enc_alg": {
					TargetField: "IpsecConfig.EncAlg",
					ConvertType: ve.ConvertDefault,
				},
				"ipsec_config_auth_alg": {
					TargetField: "IpsecConfig.AuthAlg",
					ConvertType: ve.ConvertDefault,
				},
				"ipsec_config_dh_group": {
					TargetField: "IpsecConfig.DhGroup",
					ConvertType: ve.ConvertDefault,
				},
				"ipsec_config_lifetime": {
					TargetField: "IpsecConfig.Lifetime",
					ConvertType: ve.ConvertDefault,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if len(*call.SdkParam) < 1 {
					return false, nil
				}
				(*call.SdkParam)["VpnConnectionId"] = d.Id()
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

func (s *VestackVpnConnectionService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteVpnConnection",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"VpnConnectionId": resourceData.Id(),
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
							return resource.NonRetryableError(fmt.Errorf("error on reading VpnConnection on delete %q, %w", d.Id(), callErr))
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

func (s *VestackVpnConnectionService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids": {
				TargetField: "VpnConnectionIds",
				ConvertType: ve.ConvertWithN,
			},
			"vpn_connection_names": {
				TargetField: "VpnConnectionNames",
				ConvertType: ve.ConvertWithN,
			},
		},
		NameField:    "VpnConnectionName",
		IdField:      "VpnConnectionId",
		CollectField: "vpn_connections",
		ResponseConverts: map[string]ve.ResponseConvert{
			"VpnConnectionId": {
				TargetField: "id",
				KeepDefault: true,
			},
			"IkeConfig.Psk": {
				TargetField: "ike_config_psk",
			},
			"IkeConfig.Version": {
				TargetField: "ike_config_version",
			},
			"IkeConfig.Mode": {
				TargetField: "ike_config_mode",
			},
			"IkeConfig.EncAlg": {
				TargetField: "ike_config_enc_alg",
			},
			"IkeConfig.AuthAlg": {
				TargetField: "ike_config_auth_alg",
			},
			"IkeConfig.DhGroup": {
				TargetField: "ike_config_dh_group",
			},
			"IkeConfig.Lifetime": {
				TargetField: "ike_config_lifetime",
			},
			"IkeConfig.LocalId": {
				TargetField: "ike_config_local_id",
			},
			"IkeConfig.RemoteId": {
				TargetField: "ike_config_remote_id",
			},
			"IpsecConfig.EncAlg": {
				TargetField: "ipsec_config_enc_alg",
			},
			"IpsecConfig.AuthAlg": {
				TargetField: "ipsec_config_auth_alg",
			},
			"IpsecConfig.DhGroup": {
				TargetField: "ipsec_config_dh_group",
			},
			"IpsecConfig.Lifetime": {
				TargetField: "ipsec_config_lifetime",
			},
		},
	}
}

func (s *VestackVpnConnectionService) ReadResourceId(id string) string {
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
