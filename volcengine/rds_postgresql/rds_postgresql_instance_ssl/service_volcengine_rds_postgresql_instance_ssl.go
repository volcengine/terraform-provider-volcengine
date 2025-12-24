package rds_postgresql_instance_ssl

import (
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
	rdsPgInstance "github.com/volcengine/terraform-provider-volcengine/volcengine/rds_postgresql/rds_postgresql_instance"
)

type VolcengineRdsPostgresqlInstanceSslService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewRdsPostgresqlInstanceSslService(c *ve.SdkClient) *VolcengineRdsPostgresqlInstanceSslService {
	return &VolcengineRdsPostgresqlInstanceSslService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineRdsPostgresqlInstanceSslService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineRdsPostgresqlInstanceSslService) ReadResources(condition map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
	)
	return ve.WithSimpleQuery(condition, func(m map[string]interface{}) ([]interface{}, error) {
		action := "DescribeDBInstanceSSL"
		logger.Debug(logger.ReqFormat, action, condition)
		if m == nil {
			return []interface{}{}, nil
		}
		if idsRaw, ok := m["Ids"]; ok {
			var idList []string
			if set, ok := idsRaw.(*schema.Set); ok {
				for _, it := range set.List() {
					idList = append(idList, fmt.Sprint(it))
				}
			} else if arr, ok := idsRaw.([]interface{}); ok {
				for _, it := range arr {
					idList = append(idList, fmt.Sprint(it))
				}
			} else if arr, ok := idsRaw.([]string); ok {
				idList = arr
			}
			for _, v := range idList {
				req := map[string]interface{}{"InstanceId": v}
				resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &req)
				if err != nil {
					return data, err
				}
				results, err = ve.ObtainSdkValue("Result", *resp)
				if err != nil {
					return data, err
				}
				if results == nil {
					results = map[string]interface{}{}
				}
				if rm, ok3 := results.(map[string]interface{}); ok3 {
					rm["InstanceId"] = v
					if dl, ok := m["DownloadCertificate"]; ok && dl.(bool) {
						creq := map[string]interface{}{"InstanceId": v}
						resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo("DownloadSSLCertificate"), &creq)
						if err != nil {
							return data, err
						}
						cert, err := ve.ObtainSdkValue("Result.Certificate", *resp)
						if err != nil {
							return data, err
						}
						rm["Certificate"] = cert
					}
					data = append(data, rm)
				} else {
					data = append(data, results)
				}
			}
			return data, nil
		}
		resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &m)
		if err != nil {
			return data, err
		}
		results, err = ve.ObtainSdkValue("Result", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = map[string]interface{}{}
		}
		if rm, ok3 := results.(map[string]interface{}); ok3 {
			if iid, ok4 := m["InstanceId"]; ok4 {
				rm["InstanceId"] = iid
			}
			if dl, ok := m["DownloadCertificate"]; ok && dl.(bool) {
				creq := map[string]interface{}{"InstanceId": rm["InstanceId"]}
				resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo("DownloadSSLCertificate"), &creq)
				if err != nil {
					return data, err
				}
				cert, err := ve.ObtainSdkValue("Result.Certificate", *resp)
				if err != nil {
					return data, err
				}
				rm["Certificate"] = cert
			}
		}
		return []interface{}{results}, nil
	})
}

func (s *VolcengineRdsPostgresqlInstanceSslService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"InstanceId": id,
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
		return data, fmt.Errorf("rds_postgresql_instance_ssl %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineRdsPostgresqlInstanceSslService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			return map[string]interface{}{}, "", nil
		},
	}
}

func (s *VolcengineRdsPostgresqlInstanceSslService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	instanceId := resourceData.Get("instance_id").(string)
	// Create 仅用于开启 SSL；如需强制加密，追加第二次独立调用；原因是 ssl_enable 和 force_encryption 互斥
	first := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyDBInstanceSSL",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam:    &map[string]interface{}{},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["InstanceId"] = instanceId
				if v, ok := d.GetOk("ssl_enable"); ok && !v.(bool) {
					return false, fmt.Errorf("ssl_enable must be true on create")
				}
				(*call.SdkParam)["SSLEnable"] = true
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				d.SetId(fmt.Sprintf("%s", instanceId))
				return nil
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				rdsPgInstance.NewRdsPostgresqlInstanceService(s.Client): {
					Target:     []string{"Running"},
					Timeout:    resourceData.Timeout(schema.TimeoutCreate),
					ResourceId: instanceId,
				},
			},
			LockId: func(d *schema.ResourceData) string {
				return instanceId
			},
		},
	}
	var callbacks []ve.Callback
	callbacks = append(callbacks, first)
	if v, ok := resourceData.GetOk("force_encryption"); ok {
		second := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyDBInstanceSSL",
				ConvertMode: ve.RequestConvertIgnore,
				ContentType: ve.ContentTypeJson,
				SdkParam:    &map[string]interface{}{},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["InstanceId"] = instanceId
					(*call.SdkParam)["ForceEncryption"] = v.(bool)
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
				ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
					rdsPgInstance.NewRdsPostgresqlInstanceService(s.Client): {
						Target:     []string{"Running"},
						Timeout:    resourceData.Timeout(schema.TimeoutCreate),
						ResourceId: instanceId,
					},
				},
				LockId: func(d *schema.ResourceData) string { return instanceId },
			},
		}
		callbacks = append(callbacks, second)
	}
	return callbacks
}

func (VolcengineRdsPostgresqlInstanceSslService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"Address":         {TargetField: "address"},
			"ForceEncryption": {TargetField: "force_encryption"},
			"IsValid":         {TargetField: "is_valid"},
			"SSLEnable":       {TargetField: "ssl_enable"},
			"SSLExpireTime":   {TargetField: "ssl_expire_time"},
			"TLSVersion":      {TargetField: "tls_version"},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineRdsPostgresqlInstanceSslService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	instanceId := s.ReadResourceId(resourceData.Id())
	var callbacks []ve.Callback
	if resourceData.HasChange("ssl_enable") {
		cb := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyDBInstanceSSL",
				ConvertMode: ve.RequestConvertIgnore,
				ContentType: ve.ContentTypeJson,
				SdkParam:    &map[string]interface{}{},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["InstanceId"] = instanceId
					(*call.SdkParam)["SSLEnable"] = d.Get("ssl_enable").(bool)
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
				ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
					rdsPgInstance.NewRdsPostgresqlInstanceService(s.Client): {
						Target:     []string{"Running"},
						Timeout:    resourceData.Timeout(schema.TimeoutUpdate),
						ResourceId: instanceId,
					},
				},
				LockId: func(d *schema.ResourceData) string { return instanceId },
			},
		}
		callbacks = append(callbacks, cb)
	}
	if resourceData.HasChange("force_encryption") {
		cb := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyDBInstanceSSL",
				ConvertMode: ve.RequestConvertIgnore,
				ContentType: ve.ContentTypeJson,
				SdkParam:    &map[string]interface{}{},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["InstanceId"] = instanceId
					var enabled bool
					if d.HasChange("ssl_enable") {
						enabled = d.Get("ssl_enable").(bool)
					} else {
						current, err := s.ReadResource(d, instanceId)
						if err != nil {
							return false, err
						}
						if v, ok := current["SSLEnable"].(bool); ok {
							enabled = v
						}
					}
					if !enabled {
						return false, fmt.Errorf("force_encryption requires ssl enabled")
					}
					(*call.SdkParam)["ForceEncryption"] = d.Get("force_encryption")
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
				ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
					rdsPgInstance.NewRdsPostgresqlInstanceService(s.Client): {
						Target:     []string{"Running"},
						Timeout:    resourceData.Timeout(schema.TimeoutUpdate),
						ResourceId: instanceId,
					},
				},
				LockId: func(d *schema.ResourceData) string { return instanceId },
			},
		}
		callbacks = append(callbacks, cb)
	}
	if resourceData.HasChange("reload_ssl_certificate") && resourceData.Get("reload_ssl_certificate").(bool) {
		cb := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyDBInstanceSSL",
				ConvertMode: ve.RequestConvertIgnore,
				ContentType: ve.ContentTypeJson,
				SdkParam: &map[string]interface{}{
					"InstanceId":           instanceId,
					"ReloadSSLCertificate": true,
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					var enabled bool
					if d.HasChange("ssl_enable") {
						enabled = d.Get("ssl_enable").(bool)
					} else {
						current, err := s.ReadResource(d, instanceId)
						if err != nil {
							return false, err
						}
						if v, ok := current["SSLEnable"].(bool); ok {
							enabled = v
						}
					}
					if !enabled {
						return false, fmt.Errorf("reload_ssl_certificate requires ssl enabled")
					}
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				},
				ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
					rdsPgInstance.NewRdsPostgresqlInstanceService(s.Client): {
						Target:     []string{"Running"},
						Timeout:    resourceData.Timeout(schema.TimeoutUpdate),
						ResourceId: instanceId,
					},
				},
				LockId: func(d *schema.ResourceData) string { return instanceId },
			},
		}
		callbacks = append(callbacks, cb)
	}
	return callbacks
}

func (s *VolcengineRdsPostgresqlInstanceSslService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	instanceId := s.ReadResourceId(resourceData.Id())
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyDBInstanceSSL",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				// 相当于关闭ssl配置
				"InstanceId": instanceId,
				"SSLEnable":  false,
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
			},
			ExtraRefresh: map[ve.ResourceService]*ve.StateRefresh{
				rdsPgInstance.NewRdsPostgresqlInstanceService(s.Client): {
					Target:     []string{"Running"},
					Timeout:    resourceData.Timeout(schema.TimeoutDelete),
					ResourceId: instanceId,
				},
			},
			LockId: func(d *schema.ResourceData) string { return instanceId },
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineRdsPostgresqlInstanceSslService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"ids":                  {TargetField: "Ids"},
			"download_certificate": {TargetField: "DownloadCertificate"},
		},
		CollectField: "ssls",
		ContentType:  ve.ContentTypeJson,
		ResponseConverts: map[string]ve.ResponseConvert{
			"InstanceId":      {TargetField: "instance_id"},
			"SSLEnable":       {TargetField: "ssl_enable"},
			"ForceEncryption": {TargetField: "force_encryption"},
			"IsValid":         {TargetField: "is_valid"},
			"SSLExpireTime":   {TargetField: "ssl_expire_time"},
			"TLSVersion":      {TargetField: "tls_version"},
			"Address":         {TargetField: "address"},
			"Certificate":     {TargetField: "certificate"},
		},
	}
}

func (s *VolcengineRdsPostgresqlInstanceSslService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "rds_postgresql",
		Version:     "2022-01-01",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}
