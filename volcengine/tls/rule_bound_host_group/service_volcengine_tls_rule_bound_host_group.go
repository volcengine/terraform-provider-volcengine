package rule_bound_host_group

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineTlsRuleBoundHostGroupService struct {
	Client *ve.SdkClient
}

func NewTlsRuleBoundHostGroupService(c *ve.SdkClient) *VolcengineTlsRuleBoundHostGroupService {
	return &VolcengineTlsRuleBoundHostGroupService{
		Client: c,
	}
}

func (v *VolcengineTlsRuleBoundHostGroupService) GetClient() *ve.SdkClient {
	return v.Client
}

func (v *VolcengineTlsRuleBoundHostGroupService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 20, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeBoundHostGroups"
		logger.Debug(logger.ReqFormat, action, condition)
		resp, err = v.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
			ContentType: ve.Default,
			HttpMethod:  ve.GET,
			Path:        []string{action},
			Client:      v.Client.BypassSvcClient.NewTlsClient(),
		}, &m)
		logger.Debug(logger.RespFormat, action, resp)
		results, err = ve.ObtainSdkValue("RESPONSE.HostGroupInfos", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, fmt.Errorf("Result.HostGroupInfos is not Slice")
		}
		return data, err
	})
}

func (v *VolcengineTlsRuleBoundHostGroupService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	// 这个资源是表示 rule 和 host group 的绑定关系，没有单一的 Read 接口
	// 通过 DescribeBoundHostGroups 查询指定 RuleId 下绑定的 HostGroup
	// 但由于 Terraform 的 id 通常是 rule_id:host_group_id，这里需要适配

	// 这里我们通过 rule_id 查询所有的 host_groups，然后检查 id 指定的 host_group 是否在列表中
	ruleId := resourceData.Get("rule_id").(string)
	hostGroupId := resourceData.Get("host_group_id").(string)

	if ruleId == "" || hostGroupId == "" {
		// 如果从 resourceData 获取不到，尝试从 id 解析
		// 假设 id 格式为 ruleId:hostGroupId
		// 但更标准的做法是依赖 resourceData 中的 required 字段
		return nil, nil
	}

	req := map[string]interface{}{
		"RuleId": ruleId,
	}

	results, err := v.ReadResources(req)
	if err != nil {
		return nil, err
	}

	for _, item := range results {
		if m, ok := item.(map[string]interface{}); ok {
			if hgId, ok := m["HostGroupId"].(string); ok && hgId == hostGroupId {
				m["RuleId"] = ruleId
				return m, nil
			}
		}
	}

	// Not found
	return nil, nil
}

func (v *VolcengineTlsRuleBoundHostGroupService) RefreshResourceState(resourceData *schema.ResourceData, strings []string, duration time.Duration, s string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (v *VolcengineTlsRuleBoundHostGroupService) WithResourceResponseHandlers(m map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		convert := map[string]ve.ResponseConvert{
			"RuleId": {
				TargetField: "rule_id",
			},
			"HostGroupId": {
				TargetField: "host_group_id",
			},
		}
		return m, convert, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (v *VolcengineTlsRuleBoundHostGroupService) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ApplyRuleToHostGroups",
			ContentType: ve.ContentTypeJson,
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"rule_id": {
					TargetField: "RuleId",
				},
				"host_group_id": {
					Ignore: true,
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				if v, ok := d.GetOk("rule_id"); ok {
					(*call.SdkParam)["RuleId"] = v.(string)
				}
				// 单个 ID 转数组
				if hgId, ok := d.GetOk("host_group_id"); ok {
					(*call.SdkParam)["HostGroupIds"] = []string{hgId.(string)}
				}
				return v.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.PUT,
					Path:        []string{call.Action},
					Client:      v.Client.BypassSvcClient.NewTlsClient(),
				}, call.SdkParam)
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				ruleId := d.Get("rule_id").(string)
				hostGroupId := d.Get("host_group_id").(string)
				d.SetId(fmt.Sprintf("%s:%s", ruleId, hostGroupId))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (v *VolcengineTlsRuleBoundHostGroupService) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	// 不支持更新，ForceNew
	return []ve.Callback{}
}

func (v *VolcengineTlsRuleBoundHostGroupService) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteRuleFromHostGroups",
			ContentType: ve.ContentTypeJson,
			ConvertMode: ve.RequestConvertAll,
			Convert: map[string]ve.RequestConvert{
				"rule_id": {
					TargetField: "RuleId",
				},
				"host_group_id": {
					Ignore: true,
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				if v, ok := d.GetOk("rule_id"); ok {
					(*call.SdkParam)["RuleId"] = v.(string)
				}
				if hgId, ok := d.GetOk("host_group_id"); ok {
					(*call.SdkParam)["HostGroupIds"] = []string{hgId.(string)}
				}
				return v.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.PUT,
					Path:        []string{call.Action},
					Client:      v.Client.BypassSvcClient.NewTlsClient(),
				}, call.SdkParam)
			},
			CallError: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall, baseErr error) error {
				return resource.Retry(15*time.Minute, func() *resource.RetryError {
					// 1. 先尝试删除
					_, callErr := call.ExecuteCall(d, client, call)
					if callErr == nil {
						return nil
					}

					// 2. 如果删除失败，检查资源是否还存在
					data, readErr := v.ReadResource(d, "")
					if readErr != nil {
						// 如果读取也报错，这里简单处理，认为是不可重试错误，或者可以根据错误类型判断
						// 但通常 Describe 接口不会报错除非服务挂了
						return resource.NonRetryableError(fmt.Errorf("error on reading tls rule bond host group on delete %q, %w", d.Id(), readErr))
					}
					if len(data) == 0 {
						// 资源不存在，说明删除其实成功了
						return nil
					}

					return resource.RetryableError(callErr)
				})
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				d.SetId("")
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (v *VolcengineTlsRuleBoundHostGroupService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		CollectField: "host_groups",   // DescribeBoundHostGroups 返回的列表字段
		IdField:      "HostGroupId",   // 机器组 ID
		NameField:    "HostGroupName", // 机器组名称
	}
}

func (v *VolcengineTlsRuleBoundHostGroupService) ReadResourceId(s string) string {
	return s
}
