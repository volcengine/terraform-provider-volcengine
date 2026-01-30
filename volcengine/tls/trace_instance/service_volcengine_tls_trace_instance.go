package trace_instance

import (
	"fmt"
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
		ok      bool
	)
	return ve.WithPageNumberQuery(m, "PageSize", "PageNumber", 20, 1, func(condition map[string]interface{}) ([]interface{}, error) {
		action := "DescribeTraceInstances"
		logger.Debug(logger.ReqFormat, action, condition)
		resp, err = v.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
			ContentType: ve.Default,
			HttpMethod:  ve.GET,
			Path:        []string{action},
			Client:      v.Client.BypassSvcClient.NewTlsClient(),
		}, &m)
		if err != nil {
			return data, err
		}
		logger.Debug(logger.RespFormat, action, resp)

		results, err = ve.ObtainSdkValue("RESPONSE.TraceInstances", *resp)
		if err != nil {
			return data, err
		}
		if results == nil {
			results = []interface{}{}
		}

		if data, ok = results.([]interface{}); !ok {
			return data, fmt.Errorf("results is not []interface{}")
		}
		return data, err
	})
}

func (v *VolcengineTlsTraceInstanceService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	var (
		resp    *map[string]interface{}
		results interface{}
		ok      bool
	)

	if id == "" {
		id = resourceData.Id()
	}

	action := "DescribeTraceInstance"
	req := map[string]interface{}{
		"TraceInstanceId": id,
	}
	logger.Debug(logger.ReqFormat, action, req)
	resp, err = v.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
		ContentType: ve.Default,
		HttpMethod:  ve.GET,
		Path:        []string{action},
		Client:      v.Client.BypassSvcClient.NewTlsClient(),
	}, &req)
	if err != nil {
		return data, err
	}
	logger.Debug(logger.RespFormat, action, resp)

	results, err = ve.ObtainSdkValue("RESPONSE", *resp)
	if err != nil {
		return data, err
	}

	if data, ok = results.(map[string]interface{}); !ok {
		return data, fmt.Errorf("response is not map[string]interface{}")
	}
	return data, nil
}

func (v *VolcengineTlsTraceInstanceService) RefreshResourceState(data *schema.ResourceData, strings []string, duration time.Duration, s string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (v *VolcengineTlsTraceInstanceService) WithResourceResponseHandlers(m map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return m, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (v *VolcengineTlsTraceInstanceService) CreateResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "CreateTraceInstance",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"project_id": {
					ConvertType: ve.ConvertDefault,
				},
				"trace_instance_name": {
					ConvertType: ve.ConvertDefault,
				},
				"description": {
					ConvertType: ve.ConvertDefault,
				},
				"backend_config": {
					ConvertType: ve.ConvertJsonObject,
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
				logger.Debug(logger.RespFormat, call.Action, resp)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				id, _ := ve.ObtainSdkValue("RESPONSE.TraceInstanceId", *resp)
				d.SetId(id.(string))
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (v *VolcengineTlsTraceInstanceService) ModifyResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	callback := ve.Callback{
		Call: ve.SdkCall{
			Action:      "ModifyTraceInstance",
			ConvertMode: ve.RequestConvertAll,
			ContentType: ve.ContentTypeJson,
			Convert: map[string]ve.RequestConvert{
				"trace_instance_id": {
					ConvertType: ve.ConvertDefault,
				},
				"trace_instance_name": {
					ConvertType: ve.ConvertDefault,
				},
				"description": {
					ConvertType: ve.ConvertDefault,
				},
				"backend_config": {
					ConvertType: ve.ConvertJsonObject,
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
				logger.Debug(logger.RespFormat, call.Action, resp)
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
				logger.Debug(logger.RespFormat, call.Action, resp)
				return resp, err
			},
			AfterCall: func(d *schema.ResourceData, client *ve.SdkClient, resp *map[string]interface{}, call ve.SdkCall) error {
				return nil
			},
		},
	}
	return []ve.Callback{callback}
}

func (v *VolcengineTlsTraceInstanceService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
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
