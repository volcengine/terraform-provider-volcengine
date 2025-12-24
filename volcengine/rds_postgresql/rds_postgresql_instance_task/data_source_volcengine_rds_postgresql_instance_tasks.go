package rds_postgresql_instance_task

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRdsPostgresqlInstanceTasks() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRdsPostgresqlInstanceTasksRead,
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
				Optional:    true,
				Description: "The id of the PostgreSQL instance.",
			},
			"task_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Task ID. Note: One of TaskId or task time (creation_start_time and creation_end_time) must be specified.",
			},
			"task_action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Task action.",
			},
			"creation_start_time": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Task start time. Format: yyyy-MM-ddTHH:mm:ssZ (UTC). Note: One of TaskId or task time (creation_start_time and creation_end_time) must be specified.",
			},
			"creation_end_time": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Task end time. Format: yyyy-MM-ddTHH:mm:ssZ (UTC). Note: The maximum interval between creation_start_time and creation_end_time cannot exceed 7 days.",
			},
			"task_status": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.StringInSlice([]string{"Canceled", "WaitStart", "WaitSwitch", "Running", "Running_BeforeSwitch", "Running_Switching", "Running_AfterSwitch", "Success", "Failed", "Timeout", "Rollbacking", "RollbackFailed", "Paused"}, false),
				},
				Set:         schema.HashString,
				Description: "Task status. Values: Canceled, WaitStart, WaitSwitch, Running, Running_BeforeSwitch, Running_Switching, Running_AfterSwitch, Success, Failed, Timeout, Rollbacking, RollbackFailed, Paused.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Project name.",
			},
			"task_infos": {
				Description: "Task list.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cost_time_ms": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Task execution time in milliseconds.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Task creation time. Format: yyyy-MM-ddTHH:mm:ssZ (UTC).",
						},
						"finish_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Task finish time. Format: yyyy-MM-ddTHH:mm:ssZ (UTC).",
						},
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Instance ID.",
						},
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Project name.",
						},
						"region": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Region.",
						},
						"task_action": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Task action.",
						},
						"task_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Task ID.",
						},
						"task_params": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Task parameters in JSON string.",
						},
						"task_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Task status.",
						},
						"scheduled_switch_end_time": {
							Type:     schema.TypeString,
							Computed: true,
							Description: "The scheduled end time for the switch. The time format is yyyy-MM-ddTHH:mm:ssZ (UTC time). " +
								"Note: This field will only be returned for tasks in the \"Waiting to Start\", \"Waiting to Execute\", or \"Waiting to Switch\" states.",
						},
						"scheduled_switch_start_time": {
							Type:     schema.TypeString,
							Computed: true,
							Description: "The start time of the scheduled switch. The time format is yyyy-MM-ddTHH:mm:ssZ (UTC time). " +
								"Note: This field will only be returned for tasks in the \"Waiting to Start\", \"Waiting to Execute\", or \"Waiting to Switch\" states.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineRdsPostgresqlInstanceTasksRead(d *schema.ResourceData, meta interface{}) error {
	service := NewRdsPostgresqlInstanceTaskService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineRdsPostgresqlInstanceTasks())
}
