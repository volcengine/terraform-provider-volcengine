package describe_trace

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
		resp *map[string]interface{}
		ok   bool
	)

	// Check required parameter
	traceInstanceId, ok := m["trace_instance_id"].(string)
	if !ok || traceInstanceId == "" {
		if traceInstanceId, ok = m["TraceInstanceId"].(string); !ok || traceInstanceId == "" {
			return nil, fmt.Errorf("trace_instance_id is required")
		}
	}

	// Check required parameter
	traceId, ok := m["trace_id"].(string)
	if !ok || traceId == "" {
		if traceId, ok = m["TraceId"].(string); !ok || traceId == "" {
			return nil, fmt.Errorf("trace_id is required")
		}
	}

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
