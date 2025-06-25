package schedule_sql_task

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineScheduleSqlTasks() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineScheduleSqlTasksRead,
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
			"project_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The log project ID to which the source log topic belongs.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the log item to which the source log topic belongs.",
			},
			"iam_project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IAM log project name.",
			},
			"topic_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Source log topic ID.",
			},
			"source_topic_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Source log topic name.",
			},
			"task_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Timed SQL analysis task ID.",
			},
			"task_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Timed SQL analysis task name.",
			},
			"status": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Timed SQL analysis task status.",
			},
			"tasks": {
				Description: "The List of timed SQL analysis tasks.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"query": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timed SQL analysis tasks are retrieval and analysis statements that are executed regularly.",
						},
						"status": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Whether to start the scheduled SQL analysis task immediately after completing the task configuration.",
						},
						"task_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timed SQL analysis task ID.",
						},
						"task_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timed SQL analysis task name.",
						},
						"dest_region": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The region to which the target log project belongs.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A simple description of the timed SQL analysis task.",
						},
						"dest_topic_id": {
							Computed:    true,
							Type:        schema.TypeString,
							Description: "The target log topic ID used for storing the result data of timed SQL analysis.",
						},
						"request_cycle": {
							Type:        schema.TypeList,
							Computed:    true,
							MaxItems:    1,
							Description: "The scheduling cycle of timed SQL analysis tasks.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"time": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The scheduling cycle or the time point of regular execution (the number of minutes away from 00:00), with a value range of 1 to 1440, and the unit is minutes.",
									},
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type of Scheduling cycle.",
									},
									"cron_tab": {
										Type:     schema.TypeString,
										Computed: true,
										Description: "Cron expression. The log service specifies the timed execution of alarm tasks through the Cron expression. " +
											"The minimum granularity of Cron expressions is minutes, 24 hours. " +
											"For example, 0 18 * * * indicates that an alarm task is executed exactly at 18:00 every day.",
									},
									"cron_time_zone": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "When setting the Type to Cron, the time zone also needs to be set.",
									},
								},
							},
						},
						"dest_project_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The log project ID to which the target log topic belongs.",
						},
						"dest_topic_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the target log topic used for storing the data of the timed SQL analysis results.",
						},
						"process_end_time": {
							Computed:    true,
							Type:        schema.TypeInt,
							Description: "Schedule the end time of the timed SQL analysis task in the format of a second-level timestamp.",
						},
						"create_time_stamp": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Set the creation time of timed SQL analysis tasks.",
						},
						"modify_time_stamp": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The most recent modification time of the scheduled SQL analysis task.",
						},
						"process_sql_delay": {
							Computed:    true,
							Type:        schema.TypeInt,
							Description: "The delay time of each scheduling. The value range is from 0 to 120, and the unit is seconds.",
						},
						"source_project_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The log project ID to which the source log topic belongs.",
						},
						"source_topic_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the source log topic where the original log for timed SQL analysis is located.",
						},
						"process_start_time": {
							Type:     schema.TypeInt,
							Computed: true,
							Description: "The start time of the scheduled SQL task, that is, the start time when the first instance is scheduled. " +
								"The format is a timestamp at the second level.",
						},
						"process_time_window": {
							Computed: true,
							Type:     schema.TypeString,
							Description: "SQL time window, which refers to the time range for log retrieval and analysis " +
								"when a timed SQL analysis task is running, is in a left-closed and right-open format.",
						},
						"source_project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the log item to which the source log topic belongs.",
						},
						"source_topic_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The source log topic ID where the original log for timed SQL analysis is located.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineScheduleSqlTasksRead(d *schema.ResourceData, meta interface{}) error {
	service := NewScheduleSqlTaskService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineScheduleSqlTasks())
}
