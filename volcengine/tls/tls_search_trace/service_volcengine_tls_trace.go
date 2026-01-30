package tls_describe_trace

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

	// Check required parameter
	traceInstanceId, ok := m["trace_instance_id"].(string)
	if !ok || traceInstanceId == "" {
		if traceInstanceId, ok = m["TraceInstanceId"].(string); !ok || traceInstanceId == "" {
			return nil, fmt.Errorf("trace_instance_id is required")
		}
	}

	// Check if trace_id is provided, if so, call DescribeTrace
	traceId, ok := m["trace_id"].(string)
	if !ok {
		traceId, ok = m["TraceId"].(string)
	}

	if ok && traceId != "" {
		// Call DescribeTrace
		action := "DescribeTrace"
		req := map[string]interface{}{
			"TraceId":         traceId,
			"TraceInstanceId": traceInstanceId,
		}

		resp, err = v.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
			ContentType: ve.ApplicationJSON,
			HttpMethod:  ve.POST,
			Path:        []string{action},
			Client:      v.Client.BypassSvcClient.NewTlsClient(),
		}, &req)

		if err != nil {
			return nil, err
		}
		logger.Debug(logger.RespFormat, action, req, *resp)

		// Get trace from response
		var trace interface{}
		trace, err = ve.ObtainSdkValue("RESPONSE.Trace", *resp)
		if err != nil {
			return nil, err
		}
		if trace == nil {
			return nil, fmt.Errorf("tls trace %s not found", traceId)
		}

		traceMap, ok := trace.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("trace is not map[string]interface{}")
		}

		// Add TraceId to result for consistency
		traceMap["trace_id"] = traceId

		// Return as a list containing the single trace
		return []interface{}{traceMap}, nil
	}

	// Otherwise, call SearchTraces
	action := "SearchTraces"

	// Construct request manually
	req := map[string]interface{}{
		"TraceInstanceId": traceInstanceId,
	}

	if queryRaw, ok := m["Query"].([]interface{}); ok && len(queryRaw) > 0 {
		queryMap := queryRaw[0].(map[string]interface{})
		apiQuery := map[string]interface{}{}

		// Simple fields
		if v, ok := queryMap["asc"]; ok {
			apiQuery["Asc"] = v
		}
		if v, ok := queryMap["kind"]; ok && v != "" {
			apiQuery["Kind"] = v
		}
		if v, ok := queryMap["order"]; ok && v != "" {
			apiQuery["Order"] = v
		}
		if v, ok := queryMap["trace_id"]; ok && v != "" {
			apiQuery["TraceId"] = v
		}
		if v, ok := queryMap["status_code"]; ok && v != "" {
			apiQuery["StatusCode"] = v
		}
		if v, ok := queryMap["duration_max"]; ok {
			if val, ok := v.(int); ok && val > 0 {
				apiQuery["DurationMax"] = v
			}
		}
		if v, ok := queryMap["duration_min"]; ok {
			if val, ok := v.(int); ok && val > 0 {
				apiQuery["DurationMin"] = v
			}
		}
		if v, ok := queryMap["service_name"]; ok && v != "" {
			apiQuery["ServiceName"] = v
		}
		if v, ok := queryMap["operation_name"]; ok && v != "" {
			apiQuery["OperationName"] = v
		}
		if v, ok := queryMap["start_time_min"]; ok {
			if val, ok := v.(int); ok && val > 0 {
				apiQuery["StartTimeMin"] = v
			}
		}
		if v, ok := queryMap["start_time_max"]; ok {
			if val, ok := v.(int); ok && val > 0 {
				apiQuery["StartTimeMax"] = v
			}
		}
		if v, ok := queryMap["limit"]; ok {
			if val, ok := v.(int); ok && val > 0 {
				apiQuery["Limit"] = v
			}
		}
		if v, ok := queryMap["offset"]; ok {
			if val, ok := v.(int); ok && val >= 0 {
				apiQuery["Offset"] = v
			}
		}

		// Handle Attributes (List -> Map)
		if attrs, ok := queryMap["attributes"].([]interface{}); ok && len(attrs) > 0 {
			attrMap := map[string]interface{}{}
			for _, attr := range attrs {
				if kv, ok := attr.(map[string]interface{}); ok {
					if k, ok := kv["key"].(string); ok {
						attrMap[k] = kv["value"]
					}
				}
			}
			apiQuery["Attributes"] = attrMap
		}

		req["Query"] = apiQuery
	}

	resp, err = v.Client.BypassSvcClient.DoBypassSvcCall(ve.BypassSvcInfo{
		ContentType: ve.ApplicationJSON,
		HttpMethod:  ve.POST,
		Path:        []string{action},
		Client:      v.Client.BypassSvcClient.NewTlsClient(),
	}, &req)

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
	// SearchTraces
	if _, ok := resource.Schema["query"]; ok {
		return ve.DataSourceInfo{
			RequestConverts: map[string]ve.RequestConvert{
				"trace_instance_id": {
					TargetField: "TraceInstanceId",
				},
				"query": {
					TargetField: "Query",
					ConvertType: ve.ConvertJsonArray, // Pass through as list of maps
					NextLevelConvert: map[string]ve.RequestConvert{
						"asc":            {TargetField: "asc"},
						"kind":           {TargetField: "kind"},
						"order":          {TargetField: "order"},
						"trace_id":       {TargetField: "trace_id"},
						"status_code":    {TargetField: "status_code"},
						"duration_max":   {TargetField: "duration_max"},
						"duration_min":   {TargetField: "duration_min"},
						"service_name":   {TargetField: "service_name"},
						"operation_name": {TargetField: "operation_name"},
						"start_time_min": {TargetField: "start_time_min"},
						"start_time_max": {TargetField: "start_time_max"},
						"limit":          {TargetField: "limit"},
						"offset":         {TargetField: "offset"},
						"attributes": {
							TargetField: "attributes",
							ConvertType: ve.ConvertJsonArray,
							NextLevelConvert: map[string]ve.RequestConvert{
								"key":   {TargetField: "key"},
								"value": {TargetField: "value"},
							},
						},
					},
				},
			},
			CollectField: "traces",
			IdField:      "TraceId",
			NameField:    "TraceId",
			ContentType:  ve.ContentTypeJson,
		}
	}

	// DescribeTrace
	return ve.DataSourceInfo{
		RequestConverts: map[string]ve.RequestConvert{
			"trace_instance_id": {
				TargetField: "TraceInstanceId",
			},
			"trace_id": {
				TargetField: "TraceId",
			},
		},
		CollectField: "traces",
		IdField:      "TraceId",
		NameField:    "TraceId",
		ContentType:  ve.ContentTypeJson,
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
