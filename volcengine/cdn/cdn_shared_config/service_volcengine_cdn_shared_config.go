package cdn_shared_config

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

type VolcengineCdnSharedConfigService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewCdnSharedConfigService(c *ve.SdkClient) *VolcengineCdnSharedConfigService {
	return &VolcengineCdnSharedConfigService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineCdnSharedConfigService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineCdnSharedConfigService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNum", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListSharedConfig"

		bytes, _ := json.Marshal(condition)
		logger.Debug(logger.ReqFormat, action, string(bytes))
		if condition == nil {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), nil)
			if err != nil {
				return data, err
			}
		} else {
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &condition)
			if err != nil {
				return data, err
			}
		}
		respBytes, _ := json.Marshal(resp)
		logger.Debug(logger.RespFormat, action, condition, string(respBytes))
		results, err = ve.ObtainSdkValue("Result.ConfigData", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.ConfigData is not Slice")
		}
		for index, d := range data {
			config := d.(map[string]interface{})
			query := map[string]interface{}{
				"ConfigName": config["ConfigName"],
			}
			action = "DescribeSharedConfig"
			logger.Debug(logger.ReqFormat, action, query)
			resp, err = s.Client.UniversalClient.DoCall(getUniversalInfo(action), &query)
			if err != nil {
				return data, err
			}
			logger.Debug(logger.RespFormat, action, query, *resp)
			allowIp, err := ve.ObtainSdkValue("Result.AllowIpAccessRule", *resp)
			if err != nil {
				return data, err
			}
			data[index].(map[string]interface{})["AllowIpAccessRule"] = allowIp
			denyIp, err := ve.ObtainSdkValue("Result.DenyIpAccessRule", *resp)
			if err != nil {
				return data, err
			}
			data[index].(map[string]interface{})["DenyIpAccessRule"] = denyIp
			allowReferer, err := ve.ObtainSdkValue("Result.AllowRefererAccessRule", *resp)
			if err != nil {
				return data, err
			}
			if allowReferer != nil {
				allowReferer.(map[string]interface{})["CommonType"] = []interface{}{allowReferer.(map[string]interface{})["CommonType"]}
			}
			data[index].(map[string]interface{})["AllowRefererAccessRule"] = allowReferer
			denyReferer, err := ve.ObtainSdkValue("Result.DenyRefererAccessRule", *resp)
			if err != nil {
				return data, err
			}
			if denyReferer != nil {
				denyReferer.(map[string]interface{})["CommonType"] = []interface{}{denyReferer.(map[string]interface{})["CommonType"]}
			}
			data[index].(map[string]interface{})["DenyRefererAccessRule"] = denyReferer
			common, err := ve.ObtainSdkValue("Result.CommonMatchList", *resp)
			if err != nil {
				return data, err
			}
			if common != nil {
				common.(map[string]interface{})["CommonType"] = []interface{}{common.(map[string]interface{})["CommonType"]}
			}
			data[index].(map[string]interface{})["CommonMatchList"] = common
		}
		return data, err
	})
}

func (s *VolcengineCdnSharedConfigService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"ConfigName": id,
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
		return data, fmt.Errorf("cdn_shared_config %s not exist ", id)
	}
	return data, err
}

func (s *VolcengineCdnSharedConfigService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
	}
}

func (s *VolcengineCdnSharedConfigService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "AddSharedConfig",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"allow_ip_access_rule": {
					Ignore: true,
				},
				"deny_ip_access_rule": {
					Ignore: true,
				},
				"project_name": {
					TargetField: "project",
				},
				"allow_referer_access_rule": {
					Ignore: true,
				},
				"deny_referer_access_rule": {
					Ignore: true,
				},
				"common_match_list": {
					Ignore: true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				if allowIp, ok := d.GetOk("allow_ip_access_rule"); ok {
					result := make(map[string]interface{})
					if list, ok := allowIp.([]interface{}); ok && len(list) > 0 {
						ipMap := list[0].(map[string]interface{})
						rules := ipMap["rules"]
						result["Rules"] = rules.(*schema.Set).List()
					}
					(*call.SdkParam)["AllowIpAccessRule"] = result
				}

				if denyIp, ok := d.GetOk("deny_ip_access_rule"); ok {
					result := make(map[string]interface{})
					if list, ok := denyIp.([]interface{}); ok && len(list) > 0 {
						ipMap := list[0].(map[string]interface{})
						rules := ipMap["rules"]
						result["Rules"] = rules.(*schema.Set).List()
					}
					(*call.SdkParam)["DenyIpAccessRule"] = result
				}

				if allowRefererAccessRule, ok := d.GetOk("allow_referer_access_rule"); ok {
					result := make(map[string]interface{})
					if list, ok := allowRefererAccessRule.([]interface{}); ok && len(list) > 0 {
						allowRefererAccessRuleMap := list[0].(map[string]interface{})
						if allowEmpty, ok := allowRefererAccessRuleMap["allow_empty"]; ok {
							result["AllowEmpty"] = allowEmpty
						}
						commonTypeResult := make(map[string]interface{})
						commonType := allowRefererAccessRuleMap["common_type"]
						if commonTypeList, ok := commonType.([]interface{}); ok && len(commonTypeList) > 0 {
							commonTypeMap := commonTypeList[0].(map[string]interface{})
							if ignoreCase, ok := commonTypeMap["ignore_case"]; ok {
								commonTypeResult["IgnoreCase"] = ignoreCase
							}
							rules := commonTypeMap["rules"]
							commonTypeResult["Rules"] = rules.(*schema.Set).List()
							result["CommonType"] = commonTypeResult
						}
					}
					(*call.SdkParam)["AllowRefererAccessRule"] = result
				}

				if denyRefererAccessRule, ok := d.GetOk("deny_referer_access_rule"); ok {
					result := make(map[string]interface{})
					if list, ok := denyRefererAccessRule.([]interface{}); ok && len(list) > 0 {
						denyRefererAccessRuleMap := list[0].(map[string]interface{})
						if allowEmpty, ok := denyRefererAccessRuleMap["allow_empty"]; ok {
							result["AllowEmpty"] = allowEmpty
						}
						commonTypeResult := make(map[string]interface{})
						commonType := denyRefererAccessRuleMap["common_type"]
						if commonTypeList, ok := commonType.([]interface{}); ok && len(commonTypeList) > 0 {
							commonTypeMap := commonTypeList[0].(map[string]interface{})
							if ignoreCase, ok := commonTypeMap["ignore_case"]; ok {
								commonTypeResult["IgnoreCase"] = ignoreCase
							}
							rules := commonTypeMap["rules"]
							commonTypeResult["Rules"] = rules.(*schema.Set).List()
							result["CommonType"] = commonTypeResult
						}
					}
					(*call.SdkParam)["DenyRefererAccessRule"] = result
				}

				if common, ok := d.GetOk("common_match_list"); ok {
					result := make(map[string]interface{})
					if list, ok := common.([]interface{}); ok && len(list) > 0 {
						commonMap := list[0].(map[string]interface{})
						commonTypeResult := make(map[string]interface{})
						commonType := commonMap["common_type"]
						if commonTypeList, ok := commonType.([]interface{}); ok && len(commonTypeList) > 0 {
							commonTypeMap := commonTypeList[0].(map[string]interface{})
							if ignoreCase, ok := commonTypeMap["ignore_case"]; ok {
								commonTypeResult["IgnoreCase"] = ignoreCase
							}
							rules := commonTypeMap["rules"]
							commonTypeResult["Rules"] = rules.(*schema.Set).List()
							result["CommonType"] = commonTypeResult
						}
					}
					(*call.SdkParam)["CommonMatchList"] = result
				}

				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id := d.Get("config_name")
				d.SetId(id.(string))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineCdnSharedConfigService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"Project": {
				TargetField: "project_name",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineCdnSharedConfigService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "UpdateSharedConfig",
			ConvertMode: ve.RequestConvertIgnore,
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["ConfigName"] = d.Id()

				if d.HasChange("allow_ip_access_rule") {
					if allowIp, ok := d.GetOk("allow_ip_access_rule"); ok {
						result := make(map[string]interface{})
						if list, ok := allowIp.([]interface{}); ok && len(list) > 0 {
							ipMap := list[0].(map[string]interface{})
							rules := ipMap["rules"]
							result["Rules"] = rules.(*schema.Set).List()
						}
						(*call.SdkParam)["AllowIpAccessRule"] = result
					} else {
						(*call.SdkParam)["AllowIpAccessRule"] = map[string]interface{}{}
					}
				}

				if d.HasChange("deny_ip_access_rule") {
					if denyIp, ok := d.GetOk("deny_ip_access_rule"); ok {
						result := make(map[string]interface{})
						if list, ok := denyIp.([]interface{}); ok && len(list) > 0 {
							ipMap := list[0].(map[string]interface{})
							rules := ipMap["rules"]
							result["Rules"] = rules.(*schema.Set).List()
						}
						(*call.SdkParam)["DenyIpAccessRule"] = result
					} else {
						(*call.SdkParam)["DenyIpAccessRule"] = map[string]interface{}{}
					}
				}

				if d.HasChange("allow_referer_access_rule") {
					// common type 必传
					if allowRefererAccessRule, ok := d.GetOk("allow_referer_access_rule"); ok {
						result := make(map[string]interface{})
						if list, ok := allowRefererAccessRule.([]interface{}); ok && len(list) > 0 {
							allowRefererAccessRuleMap := list[0].(map[string]interface{})
							if allowEmpty, ok := allowRefererAccessRuleMap["allow_empty"]; ok {
								result["AllowEmpty"] = allowEmpty
							}
							commonTypeResult := make(map[string]interface{})
							commonType := allowRefererAccessRuleMap["common_type"]
							if commonTypeList, ok := commonType.([]interface{}); ok && len(commonTypeList) > 0 {
								commonTypeMap := commonTypeList[0].(map[string]interface{})
								if ignoreCase, ok := commonTypeMap["ignore_case"]; ok {
									commonTypeResult["IgnoreCase"] = ignoreCase
								}
								rules := commonTypeMap["rules"]
								commonTypeResult["Rules"] = rules.(*schema.Set).List()
								result["CommonType"] = commonTypeResult
							}
						}
						(*call.SdkParam)["AllowRefererAccessRule"] = result
					} else {
						(*call.SdkParam)["AllowRefererAccessRule"] = map[string]interface{}{}
					}
				}

				if d.HasChange("deny_referer_access_rule") {
					// common type 必传
					if denyRefererAccessRule, ok := d.GetOk("deny_referer_access_rule"); ok {
						result := make(map[string]interface{})
						if list, ok := denyRefererAccessRule.([]interface{}); ok && len(list) > 0 {
							denyRefererAccessRuleMap := list[0].(map[string]interface{})
							if allowEmpty, ok := denyRefererAccessRuleMap["allow_empty"]; ok {
								result["AllowEmpty"] = allowEmpty
							}
							commonTypeResult := make(map[string]interface{})
							commonType := denyRefererAccessRuleMap["common_type"]
							if commonTypeList, ok := commonType.([]interface{}); ok && len(commonTypeList) > 0 {
								commonTypeMap := commonTypeList[0].(map[string]interface{})
								if ignoreCase, ok := commonTypeMap["ignore_case"]; ok {
									commonTypeResult["IgnoreCase"] = ignoreCase
								}
								rules := commonTypeMap["rules"]
								commonTypeResult["Rules"] = rules.(*schema.Set).List()
								result["CommonType"] = commonTypeResult
							}
						}
						(*call.SdkParam)["DenyRefererAccessRule"] = result
					} else {
						(*call.SdkParam)["DenyRefererAccessRule"] = map[string]interface{}{}
					}
				}

				if d.HasChange("common_match_list") {
					// common type 必传
					if common, ok := d.GetOk("common_match_list"); ok {
						result := make(map[string]interface{})
						if list, ok := common.([]interface{}); ok && len(list) > 0 {
							commonMap := list[0].(map[string]interface{})
							commonTypeResult := make(map[string]interface{})
							commonType := commonMap["common_type"]
							if commonTypeList, ok := commonType.([]interface{}); ok && len(commonTypeList) > 0 {
								commonTypeMap := commonTypeList[0].(map[string]interface{})
								if ignoreCase, ok := commonTypeMap["ignore_case"]; ok {
									commonTypeResult["IgnoreCase"] = ignoreCase
								}
								rules := commonTypeMap["rules"]
								commonTypeResult["Rules"] = rules.(*schema.Set).List()
								result["CommonType"] = commonTypeResult
							}
						}
						(*call.SdkParam)["CommonMatchList"] = result
					} else {
						(*call.SdkParam)["CommonMatchList"] = map[string]interface{}{}
					}
				}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *VolcengineCdnSharedConfigService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteSharedConfig",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"ConfigName": resourceData.Id(),
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

func (s *VolcengineCdnSharedConfigService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"project_name": {
				TargetField: "Project",
			},
			"config_type_list": {
				TargetField: "ConfigTypeList",
				ConvertType: ve.ConvertListN,
			},
		},
		NameField:    "ConfigName",
		IdField:      "ConfigName",
		ContentType:  ve.ContentTypeJson,
		CollectField: "config_data",
		ResponseConverts: map[string]ve.ResponseConvert{
			"Project": {
				TargetField: "project_name",
			},
		},
	}
}

func (s *VolcengineCdnSharedConfigService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "CDN",
		Version:     "2021-03-01",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
	}
}
