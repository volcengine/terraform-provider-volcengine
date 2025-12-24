package rds_postgresql_instance_failover_log

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRdsPostgresqlInstanceFailoverLogs() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRdsPostgresqlInstanceFailoverLogsRead,
		Schema: map[string]*schema.Schema{
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
				Description: "The ID of the PostgreSQL instance.",
			},
			"query_start_time": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The start time of the query. Format: yyyy-MM-ddTHH:mmZ (UTC time).",
			},
			"query_end_time": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The end time of the query. Format: yyyy-MM-ddTHH:mmZ (UTC time).",
			},
			"limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The number of records per page. Max: 1000, Min: 1.",
			},
			"failover_logs": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The collection of failover logs.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"failover_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The time when the failover occurred. Format: yyyy-MM-ddTHH:mm:ss.sssZ (UTC time).",
						},
						"failover_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the failover, such as User or System.",
						},
						"new_master_node_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The node ID of the new master after failover.",
						},
						"old_master_node_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The node ID of the old master before failover.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineRdsPostgresqlInstanceFailoverLogsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewRdsPostgresqlInstanceFailoverLogService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineRdsPostgresqlInstanceFailoverLogs())
}
