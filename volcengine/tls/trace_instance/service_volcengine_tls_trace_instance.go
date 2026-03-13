package trace_instance

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

type VolcengineTlsTraceInstanceService struct {
	Client *ve.SdkClient
}

func (v *VolcengineTlsTraceInstanceService) GetClient() *ve.SdkClient {
	return v.Client
}

func (v *VolcengineTlsTraceInstanceService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 20, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeTraceInstances"
		logger.Debug(logger.ReqFormat, action, condition)
		resp, err = v.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
			ContentType: ve.Default,
			HttpMethod:  ve.GET,
			Path:        []string{action},
			Client:      v.Client.BypassSvcClient.NewTlsClient(),
		}, &condition)
		if err != nil {
			return nil, err
		}
		logger.Debug(logger.RespFormat, action, condition, resp)

		results, err = ve.ObtainSdkValue("RESPONSE.TraceInstances", *resp)
		if err != nil {
			return nil, err
		}
		if results == nil {
			results = []interface{}{}
		}

		if data, ok := results.([]interface{}); ok {
			return data, nil
		}
		return nil, fmt.Errorf("results is not []interface{}")
	})
}

func (v *VolcengineTlsTraceInstanceService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	if id == "" {
		id = resourceData.Id()
	}

	req := map[string]interface{}{
		"TraceInstanceId": id,
	}

	if projectId, ok := resourceData.GetOk("project_id"); ok {
		req["ProjectId"] = projectId
	}

	results, err := v.ReadResources(req)
	if err != nil {
		return nil, err
	}

	for _, result := range results {
		if m, ok := result.(map[string]interface{}); ok {
			if tid, ok := m["TraceInstanceId"].(string); ok && tid == id {
				return m, nil
			}
		}
	}

	return nil, fmt.Errorf("tls trace instance %s is not exist", id)
}

func (v *VolcengineTlsTraceInstanceService) RefreshResourceState(data *schema.ResourceData, states []string, duration time.Duration, s string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (v *VolcengineTlsTraceInstanceService) WithResourceResponseHandlers(m map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		// Special handling for BackendConfig
		var bcRaw interface{}
		for k, val := range m {
			if strings.EqualFold(k, "BackendConfig") {
				bcRaw = val
				break
			}
		}

		if bcRaw != nil {
			if bc, ok := bcRaw.(map[string]interface{}); ok {
				newBc := make(map[string]interface{})
				for k, val := range bc {
					lk := strings.ToLower(k)
					switch lk {
					case "ttl":
						newBc["ttl"] = val
					case "hotttl", "hot_ttl":
						newBc["hot_ttl"] = val
					case "coldttl", "cold_ttl":
						newBc["cold_ttl"] = val
					case "archivettl", "archive_ttl":
						newBc["archive_ttl"] = val
					case "enablehotttl", "enable_hot_ttl":
						newBc["enable_hot_ttl"] = val
					case "autosplit", "auto_split":
						newBc["auto_split"] = val
					case "maxsplitpartitions", "max_split_partitions":
						newBc["max_split_partitions"] = val
					}
				}
				m["BackendConfig"] = []interface{}{newBc}
			}
		}

		return m, map[string]ve.ResponseConvert{
			"TraceInstanceId": {
				TargetField: "trace_instance_id",
			},
			"TraceInstanceName": {
				TargetField: "trace_instance_name",
			},
			"Description": {
				TargetField: "description",
			},
			"ProjectId": {
				TargetField: "project_id",
			},
			"BackendConfig": {
				TargetField: "backend_config",
			},
			"CreateTime": {
				TargetField: "create_time",
			},
			"UpdateTime": {
				TargetField: "update_time",
			},
		}, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (v *VolcengineTlsTraceInstanceService) CreateResource(data *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateTraceInstance",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"project_id": {
					TargetField: "ProjectId",
					ConvertType: ve.ConvertDefault,
				},
				"trace_instance_name": {
					TargetField: "TraceInstanceName",
					ConvertType: ve.ConvertDefault,
				},
				"description": {
					TargetField: "Description",
					ConvertType: ve.ConvertDefault,
				},
				"backend_config": {
					TargetField: "BackendConfig",
					ConvertType: ve.ConvertJsonObject,
					ForceGet:    true,
					NextLevelConvert: map[string]ve.RequestConvert{
						"ttl": {
							TargetField: "Ttl",
							ForceGet:    true,
						},
						"hot_ttl": {
							TargetField: "HotTtl",
							ForceGet:    true,
						},
						"cold_ttl": {
							TargetField: "ColdTtl",
							ForceGet:    true,
						},
						"archive_ttl": {
							TargetField: "ArchiveTtl",
							ForceGet:    true,
						},
						"enable_hot_ttl": {
							TargetField: "EnableHotTtl",
							ForceGet:    true,
						},
						"auto_split": {
							TargetField: "AutoSplit",
							ForceGet:    true,
						},
						"max_split_partitions": {
							TargetField: "MaxSplitPartitions",
							ForceGet:    true,
						},
					},
				},
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := v.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.POST,
					Path:        []string{call.Action},
					Client:      v.Client.BypassSvcClient.NewTlsClient(),
				}, call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam, resp)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, err := ve.ObtainSdkValue("RESPONSE.TraceInstanceId", *resp)
				if err != nil {
					return err
				}
				if s, ok := id.(string); ok {
					d.SetId(s)
				} else {
					return errors.New("TraceInstanceId is not string")
				}
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (v *VolcengineTlsTraceInstanceService) ModifyResource(data *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyTraceInstance",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"trace_instance_id": {
					TargetField: "TraceInstanceId",
					ConvertType: ve.ConvertDefault,
				},
				"trace_instance_name": {
					TargetField: "TraceInstanceName",
					ConvertType: ve.ConvertDefault,
				},
				"description": {
					TargetField: "Description",
					ConvertType: ve.ConvertDefault,
				},
				"backend_config": {
					TargetField: "BackendConfig",
					ConvertType: ve.ConvertJsonObject,
					ForceGet:    true,
					NextLevelConvert: map[string]ve.RequestConvert{
						"ttl": {
							TargetField: "Ttl",
							ForceGet:    true,
						},
						"hot_ttl": {
							TargetField: "HotTtl",
							ForceGet:    true,
						},
						"cold_ttl": {
							TargetField: "ColdTtl",
							ForceGet:    true,
						},
						"archive_ttl": {
							TargetField: "ArchiveTtl",
							ForceGet:    true,
						},
						"enable_hot_ttl": {
							TargetField: "EnableHotTtl",
							ForceGet:    true,
						},
						"auto_split": {
							TargetField: "AutoSplit",
							ForceGet:    true,
						},
						"max_split_partitions": {
							TargetField: "MaxSplitPartitions",
							ForceGet:    true,
						},
					},
				},
			},
			BeforeCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (bool, error) {
				(*call.SdkParam)["TraceInstanceId"] = d.Id()
				return true, nil
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := v.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.PUT,
					Path:        []string{call.Action},
					Client:      v.Client.BypassSvcClient.NewTlsClient(),
				}, call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam, resp)
				return resp, err
			},
		},
	}
	return []ve.Callback{callback}
}

func (v *VolcengineTlsTraceInstanceService) RemoveResource(data *schema.ResourceData, r *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "DeleteTraceInstance",
			ConvertMode: ve.RequestConvertIgnore,
			ContentType: ve.ContentTypeJson,
			SdkParam: &map[string]interface{}{
				"TraceInstanceId": data.Id(),
			},
			ExecuteCall: func(d *schema.ResourceData, client *ve.SdkClient, call ve.SdkCall) (*map[string]interface{}, error) {
				logger.Debug(logger.ReqFormat, call.Action, call.SdkParam)
				resp, err := v.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
					ContentType: ve.ApplicationJSON,
					HttpMethod:  ve.DELETE,
					Path:        []string{call.Action},
					Client:      v.Client.BypassSvcClient.NewTlsClient(),
				}, call.SdkParam)
				logger.Debug(logger.RespFormat, call.Action, call.SdkParam, resp)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (v *VolcengineTlsTraceInstanceService) DatasourceResources(data *schema.ResourceData, r *schema.Resource) ve.DataSourceInfo {
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"page_number":         {TargetField: "PageNumber"},
			"page_size":           {TargetField: "PageSize"},
			"project_id":          {TargetField: "ProjectId"},
			"project_name":        {TargetField: "ProjectName"},
			"iam_project_name":    {TargetField: "IamProjectName"},
			"trace_instance_name": {TargetField: "TraceInstanceName"},
			"trace_instance_id":   {TargetField: "TraceInstanceId"},
			"status":              {TargetField: "Status"},
			"cs_account_channel":  {TargetField: "CsAccountChannel"},
		},
		CollectField: "trace_instances",
		IdField:      "TraceInstanceId",
		NameField:    "TraceInstanceName",
		ContentType:  ve.ContentTypeJson,
	}
}

func (v *VolcengineTlsTraceInstanceService) ReadResourceId(s string) string {
	return s
}

func NewTlsTraceInstanceService(c *ve.SdkClient) *VolcengineTlsTraceInstanceService {
	return &VolcengineTlsTraceInstanceService{
		Client: c,
	}
}
