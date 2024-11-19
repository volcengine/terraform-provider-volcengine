package host_group

import (
	"errors"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type Service struct {
	Client *ve.SdkClient
}

func NewService(c *ve.SdkClient) *Service {
	return &Service{
		Client: c,
	}
}

func (s *Service) GetClient() *ve.SdkClient {
	return s.Client
}

func (s *Service) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 20, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeHostGroups"
		logger.Debug(logger.ReqFormat, action, condition)
		resp, err = s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
			ContentType: ve.Default,
			HttpMethod:  ve.GET,
			Path:        []string{action},
			Client:      s.Client.BypassSvcClient.NewTlsClient(),
		}, &condition)
		logger.Debug(logger.RespFormat, action, condition, *resp)
		results, err = ve.ObtainSdkValue("RESPONSE.HostGroupHostsRulesInfos", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}
		if data, ok = results.([]interface{}); !ok {
			return data, errors.New("Result.HostGroupHostsRulesInfo is not Slice")
		}
		var res []interface{}
		for _, ele := range data {
			group := ele.(map[string]interface{})["HostGroupInfo"]
			if hosts, exist := ele.(map[string]interface{})["HostInfos"]; exist {
				var ips []string
				for _, host := range hosts.([]interface{}) {
					ips = append(ips, host.(map[string]interface{})["Ip"].(string))
				}
				group.(map[string]interface{})["HostIpList"] = ips
			}
			res = append(res, group)
		}
		return res, nil
	})
}

func (s *Service) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		results []interface{}
		ok      bool
	)
	if id == "" {
		id = s.ReadResourceId(resourceData.Id())
	}
	req := map[string]interface{}{
		"HostGroupId": id,
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
		return data, fmt.Errorf("Host Group %s not exist ", id)
	}
	return data, err
}

func (s *Service) RefreshResourceState(resourceData *schema.ResourceData, target []string, timeout time.Duration, id string) *resource.StateChangeConf {
	return nil
}

func (Service) WithResourceResponseHandlers(data map[string]interface{}) []ve.ResourceResponseHandler {
	return []ve.ResourceResponseHandler{}

}

func (s *Service) CreateResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateHostGroup",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"host_ip_list": {
					ConvertType: ve.ConvertJsonArray,
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.POST,
					Path:        []string{call.Action},
					Client:      s.Client.BypassSvcClient.NewTlsClient(),
				}, call.SdkParam)
				if err != nil {
					return nil, err
				}
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam, *resp)
				return resp, nil
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("RESPONSE.HostGroupId", *resp)
				d.SetId(id.(string))
				return nil
			},
		},
	}
	return []ve.Callback{callback}

}

func (s *Service) ModifyResource(resourceData *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyHostGroup",
			ConvertMode: ve.RequestConvertInConvert,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"host_group_name": {
					TargetField: "HostGroupName",
				},
				"auto_update": {
					TargetField: "AutoUpdate",
				},
				"update_start_time": {
					TargetField: "UpdateStartTime",
				},
				"update_end_time": {
					TargetField: "UpdateEndTime",
				},
				"service_logging": {
					TargetField: "ServiceLogging",
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["HostGroupId"] = d.Id()

				if d.HasChanges("host_ip_list", "host_identifier", "host_group_type") {
					(*call.SdkParam)["HostGroupType"] = d.Get("host_group_type")
					if d.Get("host_group_type").(string) == "IP" {
						(*call.SdkParam)["HostIpList"] = d.Get("host_ip_list").(*schema.Set).List()
					} else {
						(*call.SdkParam)["HostIdentifier"] = d.Get("host_identifier")
					}
				}
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.PUT,
					Path:        []string{call.Action},
					Client:      s.Client.BypassSvcClient.NewTlsClient(),
				}, call.SdkParam)
				if err != nil {
					return nil, err
				}
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam, *resp)
				return resp, nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *Service) RemoveResource(resourceData *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteHostGroup",
			ConvertMode: ve.RequestConvertIgnore,
			SdkParam: &map[string]interface{}{
				"HostGroupId": resourceData.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := s.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.DELETE,
					Path:        []string{call.Action},
					Client:      s.Client.BypassSvcClient.NewTlsClient(),
				}, call.SdkParam)
				if err != nil {
					return nil, err
				}
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam, *resp)
				return resp, nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (s *Service) DatasourceResources(*schema.ResourceData, *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		NameField:    "HostGroupName",
		IdField:      "HostGroupId",
		CollectField: "infos",
	}
}

func (s *Service) ReadResourceId(id string) string {
	return id
}

func (s *Service) ProjectTrn() *ve.ProjectTrn {
	return &ve.ProjectTrn{
		ServiceName:          "tls",
		ResourceType:         "hostgroup",
		ProjectSchemaField:   "iam_project_name",
		ProjectResponseField: "IamProjectName",
	}
}
