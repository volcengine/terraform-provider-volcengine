package etl_task

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineEtlTasks() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineEtlTasksRead,
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
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Specify the log item ID for querying the data processing tasks under the specified log item.",
			},
			"project_name": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Specify the name of the log item for querying the data processing tasks under the specified log item. Support fuzzy query.",
			},
			"iam_project_name": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Specify the IAM project name to query the data processing tasks under the specified IAM project.",
			},
			"source_topic_id": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Specify the log topic ID for querying the data processing tasks related to this log topic.",
			},
			"source_topic_name": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Specify the name of the log topic for querying the data processing tasks related to this log topic. Support fuzzy matching.",
			},
			"task_id": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The ID of the processing task.",
			},
			"task_name": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "The name of the processing task.",
			},
			"status": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "Specify the processing task status for querying data processing tasks in this status.",
			},
			"tasks": {
				Description: "Detailed information of the processing task.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Computed:    true,
							Type:        schema.TypeString,
							Description: "The name of the processing task.",
						},
						"enable": {
							Computed:    true,
							Type:        schema.TypeBool,
							Description: "The running status of the processing task.",
						},
						"script": {
							Computed:    true,
							Type:        schema.TypeString,
							Description: "Processing rules.",
						},
						"task_id": {
							Computed:    true,
							Type:        schema.TypeString,
							Description: "The ID of the processing task.",
						},
						"to_time": {
							Computed:    true,
							Type:        schema.TypeInt,
							Description: "The end time of the data to be processed.",
						},
						"dsl_type": {
							Computed:    true,
							Type:        schema.TypeString,
							Description: "DSL type, fixed as NORMAL.",
						},
						"from_time": {
							Computed:    true,
							Type:        schema.TypeInt,
							Description: "The start time of the data to be processed.",
						},
						"task_type": {
							Computed:    true,
							Type:        schema.TypeString,
							Description: "The task type is fixed as Resident.",
						},
						"etl_status": {
							Computed:    true,
							Type:        schema.TypeString,
							Description: "Task scheduling status.",
						},
						"project_id": {
							Computed:    true,
							Type:        schema.TypeString,
							Description: "Specify the log item ID for querying the data processing tasks under the specified log item.",
						},
						"create_time": {
							Computed:    true,
							Type:        schema.TypeString,
							Description: "Processing task creation time.",
						},
						"modify_time": {
							Computed:    true,
							Type:        schema.TypeString,
							Description: "The most recent modification time of the processing task.",
						},
						"description": {
							Computed:    true,
							Type:        schema.TypeString,
							Description: "A simple description of the processing task.",
						},
						"project_name": {
							Computed:    true,
							Type:        schema.TypeString,
							Description: "Specify the name of the log item for querying the data processing tasks under the specified log item. Support fuzzy query.",
						},
						"source_topic_id": {
							Computed:    true,
							Type:        schema.TypeString,
							Description: "The log topic ID where the log to be processed is located.",
						},
						"last_enable_time": {
							Computed:    true,
							Type:        schema.TypeString,
							Description: "Recent startup time.",
						},
						"source_topic_name": {
							Computed:    true,
							Type:        schema.TypeString,
							Description: "The name of the log topic where the log to be processed is located.",
						},
						"target_resources": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Output the relevant information of the target.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"alias": {
										Type:     schema.TypeString,
										Computed: true,
										Description: "Customize the name of the output target, " +
											"which needs to be used to refer to the output target in the data processing rules.",
									},
									"topic_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Log topics used for storing processed logs.",
									},
									"project_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The log item ID used for storing the processed logs.",
									},
									"topic_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the log topic used for storing the processed logs.",
									},
									"project_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the log item used for storing the processed logs.",
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

func dataSourceVolcengineEtlTasksRead(d *schema.ResourceData, meta interface{}) error {
	service := NewEtlTaskService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineEtlTasks())
}
