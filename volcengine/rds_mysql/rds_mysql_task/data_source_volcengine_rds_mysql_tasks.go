package rds_mysql_task

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRdsMysqlTasks() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRdsMysqlTasksRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Instance ID.",
			},
			"creation_start_time": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "The start time of the task. " +
					"The time format is yyyy-MM-ddTHH:mm:ssZ (UTC time). " +
					"Instructions: For the two groups of parameters, task time (CreationStartTime and CreationEndTime) and TaskId, one of them must be selected. " +
					"The maximum time interval between the task start time (CreationStartTime) and the task end time (CreationEndTime) cannot exceed 7 days.",
			},
			"creation_end_time": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "The end time of the task. The time format is yyyy-MM-ddTHH:mm:ssZ (UTC time). " +
					"Instructions: For the two groups of parameters, task time (CreationStartTime and CreationEndTime) and TaskId, one of them must be selected. " +
					"The maximum time interval between the task start time (CreationStartTime) and the task end time (CreationEndTime) shall not exceed 7 days.",
			},
			"task_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Task ID. Description: For the two groups of parameters, TaskId and task time (CreationStartTime and CreationEndTime), one of them must be selected.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The project name.",
			},
			"task_action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Task name.",
			},
			"task_type": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "Task type. Values: Web: Console request. " +
					"OpenAPI: OpenAPI request. " +
					"AssumeRole: Role - playing request. " +
					"Other: Other requests.",
			},
			"task_source": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "Task source. Values: " +
					"User: Tenant. " +
					"System: System. " +
					"SystemUser: Internal operation and maintenance. " +
					"UserMaintain: Maintenance operations initiated by system/operation and maintenance administrators and visible to tenants.",
			},
			"task_status": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Task status. The values are as shown in the following list, " +
					"and multiple values can be selected: " +
					"WaitSwitch: Waiting for switching. " +
					"WaitStart: Waiting for execution. " +
					"Canceled: Canceled. " +
					"Stopped: Terminated. " +
					"Running_BeforeSwitch: Running (before switching). " +
					"Timeout: Execution Timeout. " +
					"Success: Execution Success. " +
					"Failed: Execution Failed. " +
					"Running: In Execution. " +
					"Stopping: In Termination.",
			},
			"task_category": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Task Category. The values are as shown in the following list, " +
					"and multiple values can be selected: " +
					"BackupRecoveryManagement: Backup and Recovery Management. " +
					"DatabaseAdminManagement: Database Administration Management. " +
					"DatabaseProxy: Database Proxy. " +
					"HighAvailability: High Availability. " +
					"InstanceAttribute: Instance Attribute. " +
					"InstanceManagement: Instance Management. " +
					"NetworkManagement: Network Management. " +
					"SecurityManagement: Security Management. " +
					"SystemMaintainManagement: System Operation and Maintenance Management. " +
					"VersionUpgrade: Version Upgrade.",
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
			"datas": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The creation time of the task.",
						},
						"finish_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The completion time of the task.",
						},
						"progress": {
							Type:     schema.TypeInt,
							Computed: true,
							Description: "Task progress. The unit is percentage. " +
								"Description: Only tasks with a task status of In Progress, that is, " +
								"tasks with a TaskStatus value of Running, will return the task progress.",
						},
						"scheduled_execute_end_time": {
							Type:     schema.TypeString,
							Computed: true,
							Description: "The deadline for the planned startup. " +
								"The time format is yyyy-MM-ddTHH:mm:ssZ (UTC time). " +
								"Description: This field will only be returned for tasks in the \"Waiting to Start\", \"Waiting to Execute\", or \"Waiting to Switch\" states.",
						},
						"scheduled_switch_end_time": {
							Type:     schema.TypeString,
							Computed: true,
							Description: "The scheduled end time for the switch. " +
								"The time format is yyyy-MM-ddTHH:mm:ssZ (UTC time). " +
								"Description: This field will only be returned for tasks in the \"Waiting to Start\", \"Waiting to Execute\", or \"Waiting to Switch\" states.",
						},
						"scheduled_switch_start_time": {
							Type:     schema.TypeString,
							Computed: true,
							Description: "The start time of the scheduled switch. " +
								"The time format is yyyy-MM-ddTHH:mm:ssZ (UTC time). " +
								"Description: This field is returned only for tasks in the \"Waiting to Start\", \"Waiting to Execute\", or \"Waiting to Switch\" state.",
						},
						"start_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The start time of the task.",
						},
						"task_action": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Task name.",
						},
						"task_category": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Task category.",
						},
						"task_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Task ID.",
						},
						"task_desc": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the task.",
						},
						"task_params": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Task parameters.",
						},
						"task_status": {
							Type:     schema.TypeString,
							Computed: true,
							Description: "Task status. The values are as shown in the following list, " +
								"and multiple values can be selected: " +
								"WaitSwitch: Waiting for switching. " +
								"WaitStart: Waiting for execution. " +
								"Canceled: Canceled. " +
								"Stopped: Terminated. " +
								"Running_BeforeSwitch: Running (before switching). " +
								"Timeout: Execution Timeout. " +
								"Success: Execution Success. " +
								"Failed: Execution Failed. " +
								"Running: In Execution. " +
								"Stopping: In Termination.",
						},
						"task_detail": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Detailed information of the task.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"task_info": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Details of the task.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"create_time": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The creation time of the task.",
												},
												"finish_time": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The completion time of the task.",
												},
												"progress": {
													Type:     schema.TypeInt,
													Computed: true,
													Description: "Task progress. The unit is percentage. " +
														"Description: Only tasks with a task status of In Progress, that is, " +
														"tasks with a TaskStatus value of Running, will return the task progress.",
												},
												"related_instance_infos": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Instances related to the task.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"instance_id": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Instance ID.",
															},
														},
													},
												},
											},
										},
									},
									"check_item_log": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The log of inspection items for the instance major version upgrade.",
									},
									"check_items": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Check results for major version upgrade.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"risk_level": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The risk level of the failed check items. Values:\nNotice: Attention.\nWarning: Warning.\nError: Error.",
												},
												"item_name": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The name of the check item.",
												},
												"description": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The description of the check item.",
												},
												"check_detail": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Details of the failed check items.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"issue": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Problems that caused the failure to pass the check items.",
															},
															"impact": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The impact of the issue that caused the failure of the check item after the upgrade.",
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
						"task_progress": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Progress details.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:     schema.TypeString,
										Computed: true,
										Description: "Step Name. " +
											"Values:\nInstanceInitialization: Task initialization.\nInstanceRecoveryPreparation Instance recovery preparation.\n" +
											"DataBackupImport: Cold backup import.\nLogBackupBinlogAdd: Binlog playback.\nTaskSuccessful: Task success.",
									},
									"step_status": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Step status. Values:\nRunning: In progress.\nSuccess: Successful.\nFailed: Failed.\nUnexecuted: Not executed.",
									},
									"step_extra_info": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specific information of the step.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"type": {
													Type:     schema.TypeString,
													Computed: true,
													Description: "Current stage. " +
														"CostTime: The time taken for the current stage.\n" +
														"CurDataSize: The amount of data imported currently.\n" +
														"CurBinlog: The number of Binlog files being replayed currently.\n" +
														"RemainCostTime: The remaining time taken.\n" +
														"RemainDataSize: The remaining amount of data to be imported. " +
														"RemainBinlog: The number of Binlog files remaining for playback.",
												},
												"unit": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Unit. Values:\nMS: Milliseconds.\nBytes: Bytes.\nFiles: Number of (files).",
												},
												"value": {
													Type:        schema.TypeFloat,
													Computed:    true,
													Description: "The specific value corresponding to the Type field.",
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
		},
	}
}

func dataSourceVolcengineRdsMysqlTasksRead(d *schema.ResourceData, meta interface{}) error {
	service := NewRdsMysqlTaskService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineRdsMysqlTasks())
}
