package describe_trace

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineTlsSearchTraces() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineTlsSearchTracesRead,
		Schema: map[string]*schema.Schema{
			"trace_instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Trace instance ID.",
			},
			"query": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Query conditions.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"trace_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Trace ID.",
						},
						"asc": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether to sort results in ascending order. true means ascending, false means descending.",
						},
						"kind": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Type of the trace.",
						},
						"order": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Sorting field. Supported fields: Kind, Name, ServiceName, Start, End, Duration, and indexed fields in Attributes.",
						},
						"status_code": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Trace status code, used to filter traces with specific status.",
						},
						"duration_max": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Maximum trace duration in microseconds.",
						},
						"duration_min": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Minimum trace duration in microseconds.",
						},
						"service_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Service name, used to filter traces from specific service.",
						},
						"operation_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Operation name, used to filter traces with specific operation.",
						},
						"start_time_min": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Minimum start time for searching traces, in microsecond timestamp format.",
						},
						"start_time_max": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Maximum start time for searching traces, in microsecond timestamp format.",
						},
						"attributes": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Attributes.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Attribute key.",
									},
									"value": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Attribute value.",
									},
								},
							},
						},
						"limit": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Maximum number of records to return, used for pagination.",
						},
						"offset": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Offset for paginated query.",
						},
					},
				},
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"traces": {
				Description: "The collection of tls trace query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"trace_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Trace ID.",
						},
						"service_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Service name.",
						},
						"operation_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Operation name.",
						},
						"start_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Trace start time in microseconds.",
						},
						"end_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Trace end time in microseconds.",
						},
						"duration": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Trace duration in microseconds.",
						},
						"status_code": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Trace status code.",
						},
						"attributes": {
							Type:     schema.TypeMap,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "Trace attributes.",
						},
					},
				},
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of tls trace query.",
			},
		},
	}
}

func dataSourceVolcengineTlsSearchTracesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewTlsTraceService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineTlsSearchTraces())
}
