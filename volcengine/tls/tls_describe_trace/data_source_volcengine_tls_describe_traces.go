package tls_describe_trace

// Update
import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineTlsDescribeTraces() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineTlsDescribeTracesRead,
		Schema: map[string]*schema.Schema{
			"trace_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Trace ID.",
			},
			"trace_instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Trace instance ID.",
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
						"spans": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The collection of spans.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"trace_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Trace ID.",
									},
									"span_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Span ID.",
									},
									"parent_span_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Parent Span ID.",
									},
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Span name.",
									},
									"kind": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Span type.",
									},
									"start_time": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Span start time.",
									},
									"end_time": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Span end time.",
									},
									"status": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Span status.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"code": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Status code.",
												},
												"message": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Error message.",
												},
											},
										},
									},
									"attributes": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Span attributes.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"key": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Attribute key.",
												},
												"value": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Attribute value.",
												},
											},
										},
									},
									"events": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Span events.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Event name.",
												},
												"timestamp": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Event timestamp.",
												},
												"attributes": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Event attributes.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"key": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Attribute key.",
															},
															"value": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Attribute value.",
															},
														},
													},
												},
											},
										},
									},
									"links": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Span links.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"trace_id": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Trace ID.",
												},
												"span_id": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Span ID.",
												},
												"trace_state": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Trace state.",
												},
												"attributes": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Link attributes.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"key": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Attribute key.",
															},
															"value": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Attribute value.",
															},
														},
													},
												},
											},
										},
									},
									"trace_state": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Trace state.",
									},
									"resource": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Resource information.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"attributes": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Resource attributes.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"key": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Attribute key.",
															},
															"value": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Attribute value.",
															},
														},
													},
												},
											},
										},
									},
									"instrumentation_library": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Instrumentation library information.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Library name.",
												},
												"version": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Library version.",
												},
											},
										},
									},
								},
							},
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

func dataSourceVolcengineTlsDescribeTracesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewTlsTraceService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineTlsDescribeTraces())
}
