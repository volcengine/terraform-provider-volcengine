package search_trace

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/logger"
)

type VolcengineTlsTraceService struct {
	Client *ve.SdkClient
}

func (v *VolcengineTlsTraceService) GetClient() *ve.SdkClient {
	return v.Client
}

func (v *VolcengineTlsTraceService) ReadResources(m map[string]interface{}) (data []interface{}, err error) {
	var (
		resp       *map[string]interface{}
		results    interface{}
		traceInfos []interface{}
		ok         bool
	)

	action := "SearchTraces"
	logger.Debug(logger.ReqFormat, action, m)

	resp, err = v.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
		ContentType: ve.ApplicationJSON,
		HttpMethod:  ve.POST,
		Path:        []string{action},
		Client:      v.Client.BypassSvcClient.NewTlsClient(),
	}, &m)

	if err != nil {
		return nil, err
	}
	logger.Debug(logger.RespFormat, action, m, *resp)

	results, err = ve.ObtainSdkValue("RESPONSE.TraceInfos", *resp)
	if err != nil {
		return nil, err
	}
	if results == nil {
		results = []interface{}{}
	}

	traceInfos, ok = results.([]interface{})
	if !ok {
		return nil, fmt.Errorf("results is not []interface{}")
	}
	return traceInfos, nil
}

func (v *VolcengineTlsTraceService) ReadResource(resourceData *schema.ResourceData, id string) (data map[string]interface{}, err error) {
	return nil, nil
}

func (v *VolcengineTlsTraceService) RefreshResourceState(data *schema.ResourceData, strings []string, duration time.Duration, s string) *resource.StateChangeConf {
	return &resource.StateChangeConf{}
}

func (v *VolcengineTlsTraceService) WithResourceResponseHandlers(m map[string]interface{}) []ve.ResourceResponseHandler {
	handler := func() (map[string]interface{}, map[string]ve.ResponseConvert, error) {
		return m, nil, nil
	}
	return []ve.ResourceResponseHandler{handler}
}

func (v *VolcengineTlsTraceService) CreateResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (v *VolcengineTlsTraceService) ModifyResource(data *schema.ResourceData, resource *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (v *VolcengineTlsTraceService) RemoveResource(data *schema.ResourceData, r *schema.Resource) []ve.Callback {
	return []ve.Callback{}
}

func (v *VolcengineTlsTraceService) DatasourceResources(data *schema.ResourceData, resource *schema.Resource) ve.DataSourceInfo {
	responseConverts := map[string]ve.ResponseConvert{
		"TraceId": {
			TargetField: "trace_id",
		},
		"ServiceName": {
			TargetField: "service_name",
		},
		"OperationName": {
			TargetField: "operation_name",
		},
		"StartTime": {
			TargetField: "start_time",
		},
		"EndTime": {
			TargetField: "end_time",
		},
		"Duration": {
			TargetField: "duration",
		},
		"StatusCode": {
			TargetField: "status_code",
		},
		"Attributes": {
			TargetField: "attributes",
		},
	}

	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"trace_instance_id": {
				TargetField: "TraceInstanceId",
			},
			"query": {
				TargetField: "Query",
				ConvertType: ve.ConvertJsonObject,
				NextLevelConvert: map[string]ve.RequestConvert{
					"asc":            {TargetField: "Asc"},
					"kind":           {TargetField: "Kind"},
					"order":          {TargetField: "Order"},
					"trace_id":       {TargetField: "TraceId"},
					"status_code":    {TargetField: "StatusCode"},
					"duration_max":   {TargetField: "DurationMax"},
					"duration_min":   {TargetField: "DurationMin"},
					"service_name":   {TargetField: "ServiceName"},
					"operation_name": {TargetField: "OperationName"},
					"start_time_min": {TargetField: "StartTimeMin"},
					"start_time_max": {TargetField: "StartTimeMax"},
					"limit":          {TargetField: "Limit"},
					"offset":         {TargetField: "Offset"},
					"attributes": {
						TargetField: "Attributes",
						Convert: func(d *schema.ResourceData, v interface{}) interface{} {
							if list, ok := v.([]interface{}); ok {
								attrMap := make(map[string]interface{})
								for _, item := range list {
									if m, ok := item.(map[string]interface{}); ok {
										attrMap[m["key"].(string)] = m["value"]
									}
								}
								return attrMap
							}
							return v
						},
					},
				},
			},
		},
		ResponseConverts: responseConverts,
		CollectField:     "traces",
		IdField:          "TraceId",
		NameField:        "TraceId",
		ContentType:      ve.ContentTypeJson,
	}
}

func (v *VolcengineTlsTraceService) ReadResourceId(id string) string {
	return id
}

func NewTlsTraceService(c *ve.SdkClient) *VolcengineTlsTraceService {
	return &VolcengineTlsTraceService{
		Client: c,
	}
}
