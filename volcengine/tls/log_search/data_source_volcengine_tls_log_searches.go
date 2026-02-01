package log_search

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

// DataSourceVolcengineTlsSearchLogs 日志数据源定义
func DataSourceVolcengineTlsSearchLogs() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineTlsSearchLogsRead,
		Schema: map[string]*schema.Schema{
			// SearchLogs parameters
			"topic_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the topic.",
			},
			"query": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The query of the log.",
			},
			"start_time": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The start time of the log.",
			},
			"end_time": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The end time of the log.",
			},
			"limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     100,
				Description: "The limit of the logs.",
			},
			"context": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The context of the log.",
			},
			"sort": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The sort of the log.",
			},
			"highlight": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether to highlight the log.",
			},
			"accurate_query": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether to use accurate query.",
			},

			// Output parameters
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of the logs.",
			},
			"logs": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The collection of query result.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"result_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the query.",
						},
						"hit_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The count of the logs.",
						},
						"list_over": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the list is over.",
						},
						"analysis": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the result is analysis.",
						},
						"limit": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The limit of the logs.",
						},
						"context": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The context of the log.",
						},
						"elapsed_millisecond": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The elapsed time of the query.",
						},
						"analysis_result": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "The analysis result of the query.",
						},
						"highlight": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The highlight of the query.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The key of the highlight.",
									},
									"value": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The value of the highlight.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"logs": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The list of the logs.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"log_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The ID of the log.",
									},
									"content": {
										Type:        schema.TypeMap,
										Computed:    true,
										Description: "The content of the log.",
									},
									"timestamp": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The timestamp of the log.",
									},
									"source": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The source of the log.",
									},
									"filename": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The filename of the log.",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

// dataSourceVolcengineTlsSearchLogsRead 读取日志数据源
func dataSourceVolcengineTlsSearchLogsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineTlsSearchLogs())
}
