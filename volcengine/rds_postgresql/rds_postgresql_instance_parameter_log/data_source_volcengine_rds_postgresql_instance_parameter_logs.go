package rds_postgresql_instance_parameter_log

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRdsPostgresqlInstanceParameterLogs() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRdsPostgresqlInstanceParameterLogsRead,
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
			"start_time": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The start time of the query. Format: yyyy-MM-ddTHH:mm:ss.sssZ (UTC time).",
			},
			"end_time": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The end time of the query. Format: yyyy-MM-ddTHH:mm:ss.sssZ (UTC time).",
			},
			"parameter_change_logs": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The collection of parameter change logs.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the parameter.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the parameter. Applied: Already in effect. Invalid: Not in effect. Syncing: Being applied, not yet in effect.",
						},
						"new_value": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The new value of the parameter.",
						},
						"old_value": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The old value of the parameter.",
						},
						"modify_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The time when the parameter was last modified. Format: yyyy-MM-ddTHH:mm:ss.sssZ (UTC time).",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineRdsPostgresqlInstanceParameterLogsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewRdsPostgresqlInstanceParameterLogService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineRdsPostgresqlInstanceParameterLogs())
}
