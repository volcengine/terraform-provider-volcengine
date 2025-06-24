package slow_log

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineSlowLogs() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineSlowLogsRead,
		Schema: map[string]*schema.Schema{
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of Resource.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of query.",
			},
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of Instance.",
			},
			"node_ids": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The node ID of the slow log needs to be queried.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"slow_log_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The types of slow logs.",
			},
			"query_start_time": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Query the start time in the format of yyyy-MM-ddTHH:mm:ssZ (UTC).",
			},
			"query_end_time": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Query the end time in the format of yyyy-MM-ddTHH:mm:ssZ (UTC).",
			},
			"db_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The database where the slow log is located.",
			},
			"context": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The context of the query results for slow log recording is used when more slow log records need to be loaded.",
			},
			"slow_query": {
				Description: "The Details of the slow log.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of Instance.",
						},
						"node_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The node ID to which the slow log belongs.",
						},
						"db_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the database.",
						},
						"user_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the account.",
						},
						"execution_start_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The execution start time of the slow query statement is in the format of yyyy-MM-ddTHH:mm: ssZ (UTC).",
						},
						"query_text": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Slow query statement.",
						},
						"query_times": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Slow query statement execution time, unit: microseconds (us).",
						},
						"host_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The address of the client that issues the slow query request.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineSlowLogsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewSlowLogService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineSlowLogs())
}
