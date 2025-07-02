package waf_domain

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineWafDomainService struct {
	Client     *ve.SdkClient
	Dispatcher *ve.Dispatcher
}

func NewWafDomainService(c *ve.SdkClient) *VolcengineWafDomainService {
	return &VolcengineWafDomainService{
		Client:     c,
		Dispatcher: &ve.Dispatcher{},
	}
}

func (s *VolcengineWafDomainService) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *VolcengineWafDomainService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 100, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "ListDomain"

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
		results, err = ve.ObtainSdkValue("Result.Data", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.Data is not Slice")
		}
		return data, err
	})
}

func (s *VolcengineWafDomainService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"Domain":        id,
		"AccurateQuery": 1,
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
		return data, fmt.Errorf("waf_domain %s not exist ", id)
	}

	var protocols []string
	protocolsString, ok := data["Protocols"].(string)
	if !ok {
		return data, fmt.Errorf("Protocols %s is not string ", data["Protocols"])
	}

	if protocolsString == "" {
		data["Protocols"] = protocols
	} else {
		result := strings.Split(protocolsString, ",")
		protocols = append(protocols, result...)
		data["Protocols"] = protocols
	}

	if data["CustomHeader"] != nil {

		customHeaderString, ok := data["CustomHeader"].([]interface{})
		if !ok {
			return data, fmt.Errorf("CustomHeader %s is not []interface{} ", data["CustomHeader"])
		}
		for _, v := range customHeaderString {
			vString, ok := v.(string)
			if !ok {
				return data, fmt.Errorf("CustomHeader %s is not string ", v)
			}
			if vString == "" {
				delete(data, "CustomHeader")
			}
			break
		}
	}

	if data["DefenceMode"] != nil {
		data["DefenceModeComputed"] = data["DefenceMode"]
		delete(data, "DefenceMode")
	}

	return data, err
}

func (s *VolcengineWafDomainService) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return &resource.StateChangeConf{
		Pending:    []string{},
		Delay:      1 * time.Second,
		MinTimeout: 1 * time.Second,
		Target:     target,
		Timeout:    timeout,
		Refresh: func() (result interface{}, state string, err error) {
			var (
				d          map[string]interface{}
				status     interface{}
				failStates []string
			)
			failStates = append(failStates, "3")

			if err = resource.Retry(5*time.Minute, func() *resource.RetryError {
				d, err = s.ReadResource(resourceData, id)
				if err != nil {
					if ve.ResourceNotFoundError(err) {
						return resource.RetryableError(err)
					} else {
						return resource.NonRetryableError(err)
					}
				}

				status, err = ve.ObtainSdkValue("Status", d)
				logger.Debug(logger.ReqFormat, "waf domain status is %s", status)

				if err != nil {
					return resource.NonRetryableError(fmt.Errorf("get sdk status value error %s", err))
				}
				statusInt, ok := status.(float64)
				if !ok {
					return resource.NonRetryableError(fmt.Errorf("status is not int type %s", status))
				}
				statusString := strconv.Itoa(int(statusInt))

				for _, v := range failStates {
					if v == statusString {
						logger.Debug(logger.ReqFormat, "waf domain statusString is %s", statusString)
						return resource.NonRetryableError(fmt.Errorf("waf domain status error, status: %s", statusString))
					}
				}

				if statusString == "2" || statusString == "5" {
					return resource.RetryableError(fmt.Errorf("waf domain status is %s, retry", statusString))
				}
				return nil
			}); err != nil {
				return nil, "", err
			}

			return d, "status by retry", nil
		},
	}
}

func (s *VolcengineWafDomainService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateDomain",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"protocol_ports": {
					ConvertType: ve.ConvertJsonObject,
					TargetField: "ProtocolPorts",
					NextLevelConvert: map[string]ve.RequestConvert{
						"http": {
							TargetField: "HTTP",
							ConvertType: ve.ConvertJsonArray,
						},
						"https": {
							TargetField: "HTTPS",
							ConvertType: ve.ConvertJsonArray,
						},
					},
				},
				"enable_http2": {
					TargetField: "EnableHTTP2",
				},
				"enable_ipv6": {
					TargetField: "EnableIPv6",
				},
				"certificate_id": {
					TargetField: "CertificateID",
				},
				"tls_enable": {
					TargetField: "TLSEnable",
				},
				"ssl_protocols": {
					TargetField: "SSLProtocols",
					ConvertType: ve.ConvertJsonArray,
				},
				"ssl_ciphers": {
					TargetField: "SSLCiphers",
					ConvertType: ve.ConvertJsonArray,
				},
				"lb_algorithm": {
					TargetField: "LBAlgorithm",
				},
				"vpc_id": {
					TargetField: "VpcID",
				},
				"protocols": {
					TargetField: "Protocols",
					ConvertType: ve.ConvertJsonArray,
				},
				"backend_groups": {
					ConvertType: ve.ConvertJsonObjectArray,
					TargetField: "BackendGroups",
					NextLevelConvert: map[string]ve.RequestConvert{
						"access_port": {
							TargetField: "AccessPort",
							ConvertType: ve.ConvertJsonArray,
						},
						"backends": {
							ConvertType: ve.ConvertJsonObjectArray,
							TargetField: "Backends",
							NextLevelConvert: map[string]ve.RequestConvert{
								"ip": {
									TargetField: "IP",
								},
							},
						},
					},
				},
				"client_ip_location": {
					TargetField: "ClientIPLocation",
				},
				"cloud_access_config": {
					ConvertType: ve.ConvertJsonObjectArray,
					TargetField: "CloudAccessConfig",
					NextLevelConvert: map[string]ve.RequestConvert{
						"instance_id": {
							TargetField: "InstanceID",
						},
						"listener_id": {
							TargetField: "ListenerID",
						},
						"lost_association_from_alb": {
							TargetField: "LostAssociationFromALB",
						},
					},
				},
				"redirect_https": {
					TargetField: "RedirectHTTPS",
				},
				"volc_certificate_id": {
					TargetField: "VolcCertificateID",
				},
				"custom_sni": {
					TargetField: "CustomSNI",
				},
				"enable_sni": {
					TargetField: "EnableSNI",
				},
				"llm_available": {
					TargetField: "LLMAvailable",
				},
				"tls_fields_config": {
					ConvertType: ve.ConvertJsonObject,
					TargetField: "TLSFieldsConfig",
					NextLevelConvert: map[string]ve.RequestConvert{
						"headers_config": {
							ConvertType: ve.ConvertJsonObject,
							TargetField: "HeadersConfig",
							NextLevelConvert: map[string]ve.RequestConvert{
								"enable": {
									TargetField: "Enable",
								},
								"excluded_key_list": {
									TargetField: "ExcludedKeyList",
								},
								"statistical_key_list": {
									TargetField: "StatisticalKeyList",
								},
							},
						},
					},
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["Region"] = client.Region
				tLSEnable, ok := d.Get("tls_enable").(int)
				if !ok {
					return false, errors.New("TLSEnable is not int")
				}
				if tLSEnable == 0 {
					(*call.SdkParam)["TLSEnable"] = 0
				}
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("Result.Domain", *resp)
				d.SetId(id.(string))
				return nil
			},
			Refresh: &ve.StateRefresh{
				Target:  []string{"status by retry"},
				Timeout: resourceData.Timeout(schema.TimeoutCreate),
			},
		},
	}
	return []ve.Callback{callback}
}

func (VolcengineWafDomainService) WithResourceResponseHandlers(d map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return d, map[string]ve.ResponseConvert{
			"AdvancedDefenseIP": {
				TargetField: "advanced_defense_ip",
			},
			"AdvancedDefenseIPv6": {
				TargetField: "advanced_defense_ipv6",
			},
			"CertificateID": {
				TargetField: "certificate_id",
			},
			"LBAlgorithm": {
				TargetField: "lb_algorithm",
			},
			"InstanceID": {
				TargetField: "instance_id",
			},
			"ListenerID": {
				TargetField: "listener_id",
			},
			"VpcID": {
				TargetField: "vpc_id",
			},
			"HTTP": {
				TargetField: "http",
			},
			"HTTPS": {
				TargetField: "https",
			},
			"EnableHTTP2": {
				TargetField: "enable_http2",
			},
			"EnableIPv6": {
				TargetField: "enable_ipv6",
			},
			"ClientIPLocation": {
				TargetField: "client_ip_location",
			},
			"TLSEnable": {
				TargetField: "tls_enable",
			},
			"SSLProtocols": {
				TargetField: "ssl_protocols",
			},
			"SSLCiphers": {
				TargetField: "ssl_ciphers",
			},
			"AutoCCEnable": {
				TargetField: "auto_cc_enable",
			},
			"IP": {
				TargetField: "ip",
			},
			"LostAssociationFromALB": {
				TargetField: "lost_association_from_alb",
			},
			"VolcCertificateID": {
				TargetField: "volc_certificate_id",
			},
			"CustomSNI": {
				TargetField: "custom_sni",
			},
			"EnableSNI": {
				TargetField: "enable_sni",
			},
			"LLMAvailable": {
				TargetField: "llm_available",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (s *VolcengineWafDomainService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) (callbacks []ve.Callback) {

	if resourceData.HasChanges("vpc_id", "tls_enable", "ssl_protocols", "ssl_ciphers", "redirect_https", "public_real_server", "proxy_write_time_out",
		"proxy_retry", "proxy_read_time_out", "proxy_keep_alive_time_out", "proxy_keep_alive", "proxy_connect_time_out", "proxy_config",
		"protocol_follow", "lb_algorithm", "keep_alive_time_out", "keep_alive_request", "enable_ipv6", "enable_http2",
		"client_max_body_size", "certificate_id", "client_ip_location", "protocols", "protocol_ports", "backend_groups",
		"cloud_access_config", "custom_header", "volc_certificate_id", "certificate_platform", "custom_sni",
		"enable_custom_redirect", "enable_sni", "llm_available", "tls_fields_config") {
		modifyDomain := ve.Callback{
			Call: ve.SdkCall{
				Action:      "UpdateDomain",
				ConvertMode: ve.RequestConvertInConvert,
				ContentType: ve.ContentTypeJson,
				Convert: map[string]ve.RequestConvert{
					"access_mode": {
						TargetField: "AccessMode",
						ForceGet:    true,
					},
					"vpc_id": {
						TargetField: "VpcID",
						ForceGet:    true,
					},
					"tls_enable": {
						TargetField: "TLSEnable",
						ForceGet:    true,
					},
					"ssl_protocols": {
						TargetField: "SSLProtocols",
						ConvertType: ve.ConvertJsonArray,
						ForceGet:    true,
					},
					"ssl_ciphers": {
						TargetField: "SSLCiphers",
						ConvertType: ve.ConvertJsonArray,
						ForceGet:    true,
					},
					"redirect_https": {
						TargetField: "RedirectHTTPS",
						ForceGet:    true,
					},
					"public_real_server": {
						TargetField: "PublicRealServer",
						ForceGet:    true,
					},
					"proxy_write_time_out": {
						TargetField: "ProxyWriteTimeOut",
						ForceGet:    true,
					},
					"proxy_retry": {
						TargetField: "ProxyRetry",
						ForceGet:    true,
					},
					"proxy_read_time_out": {
						TargetField: "ProxyReadTimeOut",
						ForceGet:    true,
					},
					"proxy_keep_alive_time_out": {
						TargetField: "ProxyKeepAliveTimeOut",
						ForceGet:    true,
					},
					"proxy_keep_alive": {
						TargetField: "ProxyKeepAlive",
						ForceGet:    true,
					},
					"proxy_connect_time_out": {
						TargetField: "ProxyConnectTimeOut",
						ForceGet:    true,
					},
					"proxy_config": {
						TargetField: "ProxyConfig",
						ForceGet:    true,
					},
					"protocol_follow": {
						TargetField: "ProtocolFollow",
						ForceGet:    true,
					},
					"lb_algorithm": {
						TargetField: "LBAlgorithm",
						ForceGet:    true,
					},
					"keep_alive_time_out": {
						TargetField: "KeepAliveTimeOut",
						ForceGet:    true,
					},
					"keep_alive_request": {
						TargetField: "KeepAliveRequest",
						ForceGet:    true,
					},
					"enable_http2": {
						TargetField: "EnableHTTP2",
						ForceGet:    true,
					},
					"enable_ipv6": {
						TargetField: "EnableIPv6",
						ForceGet:    true,
					},
					"client_max_body_size": {
						TargetField: "ClientMaxBodySize",
						ForceGet:    true,
					},
					"client_ip_location": {
						TargetField: "ClientIPLocation",
						ForceGet:    true,
					},
					"certificate_id": {
						TargetField: "CertificateID",
						ForceGet:    true,
					},
					"protocols": {
						TargetField: "Protocols",
						ConvertType: ve.ConvertJsonArray,
						ForceGet:    true,
					},
					"protocol_ports": {
						ConvertType: ve.ConvertJsonObject,
						TargetField: "ProtocolPorts",
						ForceGet:    true,
						NextLevelConvert: map[string]ve.RequestConvert{
							"http": {
								TargetField: "HTTP",
								ConvertType: ve.ConvertJsonArray,
								ForceGet:    true,
							},
							"https": {
								TargetField: "HTTPS",
								ForceGet:    true,
								ConvertType: ve.ConvertJsonArray,
							},
						},
					},
					"backend_groups": {
						ForceGet:    true,
						ConvertType: ve.ConvertJsonObjectArray,
						TargetField: "BackendGroups",
						NextLevelConvert: map[string]ve.RequestConvert{
							"access_port": {
								TargetField: "AccessPort",
								ForceGet:    true,
								ConvertType: ve.ConvertJsonArray,
							},
							"backends": {
								ConvertType: ve.ConvertJsonObjectArray,
								TargetField: "Backends",
								ForceGet:    true,
								NextLevelConvert: map[string]ve.RequestConvert{
									"ip": {
										TargetField: "IP",
										ForceGet:    true,
									},
									"protocol": {
										TargetField: "Protocol",
										ForceGet:    true,
									},
									"port": {
										TargetField: "Port",
										ForceGet:    true,
									},
									"weight": {
										TargetField: "Weight",
										ForceGet:    true,
									},
								},
							},
							"name": {
								TargetField: "Name",
								ForceGet:    true,
							},
						},
					},
					"cloud_access_config": {
						ConvertType: ve.ConvertJsonObjectArray,
						ForceGet:    true,
						TargetField: "CloudAccessConfig",
						NextLevelConvert: map[string]ve.RequestConvert{
							"instance_id": {
								TargetField: "InstanceID",
								ForceGet:    true,
							},
							"listener_id": {
								TargetField: "ListenerID",
								ForceGet:    true,
							},
							"lost_association_from_alb": {
								TargetField: "LostAssociationFromALB",
								ForceGet:    true,
							},
							"access_protocol": {
								TargetField: "AccessProtocol",
								ForceGet:    true,
							},
							"instance_name": {
								TargetField: "InstanceName",
								ForceGet:    true,
							},
							"port": {
								TargetField: "Port",
								ForceGet:    true,
							},
							"protocol": {
								TargetField: "Protocol",
								ForceGet:    true,
							},
						},
					},
					"custom_header": {
						TargetField: "CustomHeader",
						ConvertType: ve.ConvertJsonArray,
						ForceGet:    true,
					},
					"volc_certificate_id": {
						TargetField: "VolcCertificateID",
						ForceGet:    true,
					},
					"certificate_platform": {
						TargetField: "CertificatePlatform",
						ForceGet:    true,
					},
					"custom_sni": {
						TargetField: "CustomSNI",
						ForceGet:    true,
					},
					"enable_custom_redirect": {
						TargetField: "EnableCustomRedirect",
						ForceGet:    true,
					},
					"enable_sni": {
						TargetField: "EnableSNI",
						ForceGet:    true,
					},
					"llm_available": {
						TargetField: "LLMAvailable",
						ForceGet:    true,
					},
					"tls_fields_config": {
						ConvertType: ve.ConvertJsonObject,
						ForceGet:    true,
						TargetField: "TLSFieldsConfig",
						NextLevelConvert: map[string]ve.RequestConvert{
							"headers_config": {
								ConvertType: ve.ConvertJsonObject,
								ForceGet:    true,
								TargetField: "HeadersConfig",
								NextLevelConvert: map[string]ve.RequestConvert{
									"enable": {
										TargetField: "Enable",
										ForceGet:    true,
									},
									"excluded_key_list": {
										TargetField: "ExcludedKeyList",
										ForceGet:    true,
									},
									"statistical_key_list": {
										TargetField: "StatisticalKeyList",
										ForceGet:    true,
									},
								},
							},
						},
					},
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["Region"] = client.Region
					(*call.SdkParam)["Domain"] = d.Id()
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
					resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo("UpdateDomain"), call.SdkParam)
					logger.Debug(logger.RespFormat, call.Action, resp, err)
					return resp, err
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{"status by retry"},
					Timeout: resourceData.Timeout(schema.TimeoutUpdate),
				},
			},
		}
		callbacks = append(callbacks, modifyDomain)
	}

	if resourceData.HasChanges("bot_repeat_enable", "bot_dytoken_enable", "auto_cc_enable", "bot_sequence_enable",
		"bot_sequence_default_action", "bot_frequency_enable", "waf_enable", "cc_enable", "white_enable",
		"black_ip_enable", "black_lct_enable", "waf_white_req_enable", "white_field_enable", "custom_rsp_enable",
		"system_bot_enable", "custom_bot_enable", "api_enable", "tamper_proof_enable", "dlp_enable") {
		modifyUpdateWafServiceControl := ve.Callback{
			Call: ve.SdkCall{
				Action:      "UpdateWafServiceControl",
				ConvertMode: ve.RequestConvertInConvert,
				ContentType: ve.ContentTypeJson,
				Convert: map[string]ve.RequestConvert{
					"auto_cc_enable": {
						TargetField: "AutoCCEnable",
					},
					"tls_enable": {
						TargetField: "TLSEnable",
					},
					"bot_repeat_enable": {
						TargetField: "BotRepeatEnable",
					},
					"bot_dytoken_enable": {
						TargetField: "BotDytokenEnable",
					},
					"bot_sequence_enable": {
						TargetField: "BotSequenceEnable",
					},
					"bot_sequence_default_action": {
						TargetField: "BotSequenceDefaultAction",
					},
					"bot_frequency_enable": {
						TargetField: "BotFrequencyEnable",
					},
					"waf_enable": {
						TargetField: "WafEnable",
					},
					"cc_enable": {
						TargetField: "CcEnable",
					},
					"white_enable": {
						TargetField: "WhiteEnable",
					},
					"black_ip_enable": {
						TargetField: "BlackIpEnable",
					},
					"black_lct_enable": {
						TargetField: "BlackLctEnable",
					},
					"waf_white_req_enable": {
						TargetField: "WafWhiteReqEnable",
					},
					"white_field_enable": {
						TargetField: "WhiteFieldEnable",
					},
					"custom_rsp_enable": {
						TargetField: "CustomRspEnable",
					},
					"system_bot_enable": {
						TargetField: "SystemBotEnable",
					},
					"custom_bot_enable": {
						TargetField: "CustomBotEnable",
					},
					"api_enable": {
						TargetField: "ApiEnable",
					},
					"tamper_proof_enable": {
						TargetField: "TamperProofEnable",
					},
					"dlp_enable": {
						TargetField: "DlpEnable",
					},
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["Region"] = client.Region
					(*call.SdkParam)["Host"] = d.Id()
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
					resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo("UpdateWafServiceControl"), call.SdkParam)
					logger.Debug(logger.RespFormat, call.Action, resp, err)
					return resp, err
				},
				AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
					return nil
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{"status by retry"},
					Timeout: resourceData.Timeout(schema.TimeoutUpdate),
				},
			},
		}
		callbacks = append(callbacks, modifyUpdateWafServiceControl)
	}

	if resourceData.HasChanges("extra_defence_mode_lb_instance", "defence_mode") {

		modifyServiceDefenceMode := ve.Callback{
			Call: ve.SdkCall{
				Action:      "ModifyServiceDefenceMode",
				ConvertMode: ve.RequestConvertInConvert,
				ContentType: ve.ContentTypeJson,
				Convert: map[string]ve.RequestConvert{
					"extra_defence_mode_lb_instance": {
						ConvertType: ve.ConvertJsonObjectArray,
						TargetField: "ExtraDefenceModeLBInstance",
						NextLevelConvert: map[string]ve.RequestConvert{
							"defence_mode": {
								TargetField: "DefenceMode",
							},
							"instance_id": {
								TargetField: "InstanceID",
							},
						},
					},
					"defence_mode": {
						TargetField: "DefenceMode",
						ForceGet:    true,
					},
				},
				BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
					(*call.SdkParam)["Host"] = d.Id()
					(*call.SdkParam)["Region"] = client.Region
					return true, nil
				},
				ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
					logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
					resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo("ModifyServiceDefenceMode"), call.SdkParam)
					logger.Debug(logger.RespFormat, call.Action, resp, err)
					return resp, err
				},
				AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
					return nil
				},
				Refresh: &ve.StateRefresh{
					Target:  []string{"status by retry"},
					Timeout: resourceData.Timeout(schema.TimeoutUpdate),
				},
			},
		}
		callbacks = append(callbacks, modifyServiceDefenceMode)
	}

	return callbacks
}

func (s *VolcengineWafDomainService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteDomain",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"Domain": resourceData.Id(),
				"Region": s.Client.Region,
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam)
				resp, err := s.Client.UniversalClient.DoCall(getUniversalInfo(call.Action), call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, resp, err)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return ve.CheckResourceUtilRemoved(d, s.ReadResource, 5*time.Minute)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(5*time.Minute, func() *resource.RetryError {
					_, callErr := s.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on  reading waf domain on delete %q, %w", d.Id(), callErr))
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

func (s *VolcengineWafDomainService) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		NameField:    "Domain",
		CollectField: "data",
		ContentType:  ve.ContentTypeJson,
		ResponseConverts: map[string]ve.ResponseConvert{
			"AdvancedDefenseIP": {
				TargetField: "advanced_defense_ip",
			},
			"AdvancedDefenseIPv6": {
				TargetField: "advanced_defense_ipv6",
			},
			"CertificateID": {
				TargetField: "certificate_id",
			},
			"LBAlgorithm": {
				TargetField: "lb_algorithm",
			},
			"InstanceID": {
				TargetField: "instance_id",
			},
			"ListenerID": {
				TargetField: "listener_id",
			},
			"VpcID": {
				TargetField: "vpc_id",
			},
			"HTTP": {
				TargetField: "http",
			},
			"HTTPS": {
				TargetField: "https",
			},
			"EnableHTTP2": {
				TargetField: "enable_http2",
			},
			"EnableIPv6": {
				TargetField: "enable_ipv6",
			},
			"ClientIPLocation": {
				TargetField: "client_ip_location",
			},
			"TLSEnable": {
				TargetField: "tls_enable",
			},
			"SSLProtocols": {
				TargetField: "ssl_protocols",
			},
			"SSLCiphers": {
				TargetField: "ssl_ciphers",
			},
			"AutoCCEnable": {
				TargetField: "auto_cc_enable",
			},
			"IP": {
				TargetField: "ip",
			},
		},
	}
}

func (s *VolcengineWafDomainService) ReadResourceId(id string) string {
	return id
}

func getUniversalInfo(actionName string) ve.UniversalInfo {
	return ve.UniversalInfo{
		ServiceName: "waf",
		Version:     "2023-12-25",
		HttpMethod:  ve.POST,
		ContentType: ve.ApplicationJSON,
		Action:      actionName,
		RegionType:  ve.Global,
	}
}

func (s *VolcengineWafDomainService) ProjectTrn() *ve.ProjectTrn {
	return &ve.ProjectTrn{
		ServiceName:          "waf",
		ResourceType:         "domain",
		ProjectResponseField: "ProjectName",
		ProjectSchemaField:   "project_name",
	}
}
