package rule_applier

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineTlsRuleApplierService struct {
	Client *ve.SdkClient
}

func (v *VolcengineTlsRuleApplierService) GetClient() *ve.SdkClient {
	return v.Client
}

func (v *VolcengineTlsRuleApplierService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp        *map[string]interface{}
		results     interface{}
		transResult []interface{}
		ok          bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 20, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeHostGroupRules"
		logger.Debug(logger.ReqFormat, action, condition)
		resp, err = v.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
			ContentType: ve.Default,
			HttpMethod:  ve.GET,
			Path:        []string{action},
			Client:      v.Client.BypassSvcClient.NewTlsClient(),
		}, &m)
		logger.Debug(logger.RespFormat, action, resp)
		results, err = ve.ObtainSdkValue("RESPONSE.RuleInfos", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.RuleInfos is not Slice")
		}
		for _, d := range data {
			dataMap, ok := d.(map[string]interface{})
			if !ok {
				return data, errors.New("value is not map")
			}
			userRuleMap, ok := dataMap["UserDefineRule"].(map[string]interface{})
			if !ok {
				return data, errors.New("value is not map")
			}
			plugin, ok := userRuleMap["Plugin"].(map[string]interface{})
			if !ok {
				if plugin == nil {
					continue
				}
				return data, errors.New("value is not map")
			}
			if len(plugin) == 0 {
				continue
			}
			// 接口中 processors 为小写
			processors, ok := plugin["processors"]
			if !ok {
				continue
			}
			logger.DebugInfo("plugin ori : ", processors)
			p, _ := json.Marshal(processors)
			pStr := string(p)
			strings.ReplaceAll(pStr, "\\", "")
			newStr := ""
			if len(pStr) > 1 {
				newStr = pStr[1 : len(pStr)-1]
			}
			logger.DebugInfo("plugin marshal : ", newStr)
			plugin["processors"] = newStr
			userRuleMap["Plugin"] = plugin
			dataMap["UserDefineRule"] = userRuleMap
			transResult = append(transResult, dataMap)
		}
		return data, err
	})
}

func (v *VolcengineTlsRuleApplierService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
	)
	if id == "" {
		id = v.ReadResourceId(resourceData.Id())
	}
	ids := strings.Split(id, ":")
	if len(ids) != 2 {
		return data, errors.New("invalid id")
	}
	req := map[string]interface{}{
		"HostGroupId": ids[1],
	}
	results, err = v.ReadResources(req)
	if err != nil {
		return data, err
	}
	for _, r := range results {
		temp, ok := r.(map[string]interface{})
		if !ok {
			return data, fmt.Errorf("value is not map")
		}
		if temp["RuleId"].(string) == ids[0] {
			data = temp
		}
	}
	if len(data) == 0 {
		return data, fmt.Errorf("tls rule apply %s not exist", id)
	}
	return dataMapTransToList(data), nil
}

func (v *VolcengineTlsRuleApplierService) RefreshResourceState(data *schema.ResourceData, strings []string, duration time.Duration, s string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (v *VolcengineTlsRuleApplierService) WithResourceResponseHandlers(m map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return m, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (v *VolcengineTlsRuleApplierService) CreateResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ApplyRuleToHostGroups",
			ContentType: ve.ContentTypeJson,
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"host_group_id": {
					Ignore: true,
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				hostGroupIds := make([]string, 0)
				hostGroupIds = append(hostGroupIds, d.Get("host_group_id").(string))
				(*call.SdkParam)["HostGroupIds"] = hostGroupIds
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return v.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.PUT,
					Path:        []string{call.Action},
					Client:      v.Client.BypassSvcClient.NewTlsClient(),
				}, call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				ruleId := (*call.SdkParam)["RuleId"].(string)
				hostGroupId := d.Get("host_group_id").(string)
				d.SetId(fmt.Sprint(ruleId, ":", hostGroupId))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (v *VolcengineTlsRuleApplierService) ModifyResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return nil
}

func (v *VolcengineTlsRuleApplierService) RemoveResource(data *schema.ResourceData, r *schema.Resource) []ve.Callback {
	ids := strings.Split(data.Id(), ":")
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteRuleFromHostGroups",
			ContentType: ve.ContentTypeJson,
			ConvertMode: ve.RequestConvertIgnore,
			Convert: map[string]ve.RequestConvert{
				"host_group_id": {
					Ignore: true,
				},
			},
			SdkParam: &map[string]interface{}{
				"RuleId": ids[0],
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				hostGroupIds := make([]string, 0)
				hostGroupIds = append(hostGroupIds, ids[1])
				(*call.SdkParam)["HostGroupIds"] = hostGroupIds
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				return v.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.PUT,
					Path:        []string{call.Action},
					Client:      v.Client.BypassSvcClient.NewTlsClient(),
				}, call.SdkParam)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				//出现错误后重试
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					_, callErr := v.ReadResource(d, "")
					if callErr != nil {
						if ve.ResourceNotFoundError(callErr) {
							return nil
						} else {
							return resource.NonRetryableError(fmt.Errorf("error on reading tls rule apply on delete %q, %w", d.Id(), callErr))
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
				return ve.CheckResourceUtilRemoved(d, v.ReadResource, 3*time.Minute)
			},
		},
	}
	return []ve.Callback{callback}
}

func (v *VolcengineTlsRuleApplierService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		CollectField: "rules",
		ResponseConverts: map[string]ve.ResponseConvert{
			"close_eof": {
				TargetField: "Close_EOF",
			},
		},
	}
}

func (v *VolcengineTlsRuleApplierService) ReadResourceId(s string) string {
	return s
}

func NewTlsRuleApplierService(client *ve.SdkClient) *VolcengineTlsRuleApplierService {
	return &VolcengineTlsRuleApplierService{
		Client: client,
	}
}

// parse_path_rule, shard_hash_key, plugin, advanced, kubernetes_rule map 转 list
func dataMapTransToList(data map[string]interface{}) map[string]interface{} {
	userDefineRule, ok := data["UserDefineRule"].(map[string]interface{})
	if ok && userDefineRule != nil {
		parsePathRule, ok := userDefineRule["ParsePathRule"]
		if ok {
			userDefineRule["ParsePathRule"] = []interface{}{parsePathRule}
		}
		shardHashKey, ok := userDefineRule["ShardHashKey"]
		if ok {
			userDefineRule["ShardHashKey"] = []interface{}{shardHashKey}
		}
		plugin, ok := userDefineRule["Plugin"]
		if ok {
			userDefineRule["Plugin"] = []interface{}{plugin}
		}
		advanced, ok := userDefineRule["Advanced"]
		if ok {
			userDefineRule["Advanced"] = []interface{}{advanced}
		}
	}
	data["UserDefineRule"] = userDefineRule
	containerRule, ok := data["ContainerRule"].(map[string]interface{})
	if ok && containerRule != nil {
		kubernetesRule, ok := containerRule["KubernetesRule"]
		if ok {
			containerRule["KubernetesRule"] = []interface{}{kubernetesRule}
		}
	}
	data["ContainerRule"] = []interface{}{containerRule}
	return data
}
