package import_task

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineImportTasks() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineImportTasksRead,
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
			"task_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Import the task ID of the data to be queried.",
			},
			"task_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Import the task name of the data to be queried.",
			},
			"project_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specify the log item ID for querying the data import tasks under the specified log item.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specify the name of the log item for querying the data import tasks under the specified log item. Support fuzzy query..",
			},
			"topic_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specify the log topic ID for querying the data import tasks related to this log topic.",
			},
			"iam_project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specify the IAM project name to query the data import tasks under the specified IAM project.",
			},
			"topic_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specify the name of the log topic for querying the data import tasks related to this log topic. Support fuzzy query.",
			},
			"source_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specify the import type for querying the data import tasks related to this import type.",
			},
			"status": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specify the status of the import task.",
			},
			"task_info": {
				Description: "Data import task list.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"task_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Import the task ID of the data to be queried.",
						},
						"status": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The status of the data import task.",
						},
						"topic_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specify the log topic ID for querying the data import tasks related to this log topic.",
						},
						"task_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Import the task name of the data to be queried.",
						},
						"project_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specify the log item ID for querying the data import tasks under the specified log item.",
						},
						"topic_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specify the name of the log topic for querying the data import tasks related to this log topic. Support fuzzy query.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The creation time of the data import task.",
						},
						"source_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specify the import type for querying the data import tasks related to this import type.",
						},
						"target_info": {
							Type:        schema.TypeList,
							Computed:    true,
							MaxItems:    1,
							Description: "The output information of the data import task.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"region": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Regional ID.",
									},
									"log_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Specify the log parsing type when importing.",
									},
									"log_sample": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Log sample.",
									},
									"extract_rule": {
										Type:        schema.TypeList,
										Computed:    true,
										MaxItems:    1,
										Description: "Log extraction rules.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"keys": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "List of log field names (Keys).",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"quote": {
													Type:     schema.TypeString,
													Computed: true,
													Description: "Reference symbol. " +
														"The content wrapped by the reference will not be separated but will be parsed into a complete field." +
														" It is valid if and only if the LogType is delimiter_log.",
												},
												"time_key": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The field name of the log time field.",
												},
												"time_zone": {
													Type:     schema.TypeString,
													Computed: true,
													Description: "Time zone, supporting both machine time zone (default) and custom time zone. " +
														"Among them, the custom time zone supports GMT and UTC.",
												},
												"delimiter": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Log delimiter.",
												},
												"begin_regex": {
													Type:     schema.TypeString,
													Computed: true,
													Description: "The regular expression used to identify the first line in each log, " +
														"and its matching part will serve as the beginning of the log.",
												},
												"time_format": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The parsing format of the time field.",
												},
												"skip_line_count": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The number of log lines skipped.",
												},
												"un_match_log_key": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "When uploading a log that failed to parse, the key name of the parse failed log.",
												},
												"time_extract_regex": {
													Type:     schema.TypeString,
													Computed: true,
													Description: "A regular expression for extracting time, " +
														"used to extract the time value in the TimeKey field and parse it into the corresponding collection time.",
												},
												"un_match_up_load_switch": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Whether to upload the logs of failed parsing.",
												},
											},
										},
									},
								},
							},
						},
						"description": {
							Computed:    true,
							Type:        schema.TypeString,
							Description: "Data import task description.",
						},
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specify the name of the log item for querying the data import tasks under the specified log item. Support fuzzy query..",
						},
						"import_source_info": {
							Type:        schema.TypeList,
							Computed:    true,
							MaxItems:    1,
							Description: "The source information of the data import task.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tos_source_info": {
										Type:        schema.TypeList,
										Computed:    true,
										MaxItems:    1,
										Description: "TOS imports source information.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"bucket": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The TOS bucket where the log file is located.",
												},
												"prefix": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The path of the file to be imported in the TOS bucket.",
												},
												"region": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The region where the TOS bucket is located. Support cross-regional data import.",
												},
												"compress_type": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The compression mode of data in the TOS bucket.",
												},
											},
										},
									},
									"kafka_source_info": {
										Type:        schema.TypeList,
										Computed:    true,
										MaxItems:    1,
										Description: "TOS imports source information.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"host": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The service addresses corresponding to different types of Kafka clusters are different.",
												},
												"group": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Kafka consumer group.",
												},
												"topic": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Kafka Topic name.",
												},
												"encode": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The encoding format of the data.",
												},
												"password": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The Kafka SASL user password used for identity authentication.",
												},
												"protocol": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Secure Transport protocol.",
												},
												"username": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The Kafka SASL username used for identity authentication.",
												},
												"mechanism": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Password authentication mechanism.",
												},
												"instance_id": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "When you are using the Volcano Engine Message Queue Kafka version, it should be set to the Kafka instance ID.",
												},
												"is_need_auth": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Whether to enable authentication.",
												},
												"initial_offset": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The starting position of data import.",
												},
												"time_source_default": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Specify the log time.",
												},
											},
										},
									},
								},
							},
						},
						"task_statistics": {
							Type:        schema.TypeList,
							Computed:    true,
							MaxItems:    1,
							Description: "The progress of the data import task.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"total": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The total number of resources that have been listed.",
									},
									"failed": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The number of resources that failed to import.",
									},
									"skipped": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Skip the number of imported resources.",
									},
									"not_exist": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The number of non-existent resources.",
									},
									"bytes_total": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The total number of resource bytes that have been listed.",
									},
									"task_status": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Import the status of the task.",
									},
									"transferred": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The number of imported resources.",
									},
									"bytes_transferred": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The number of imported bytes.",
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

func dataSourceVolcengineImportTasksRead(d *schema.ResourceData, meta interface{}) error {
	service := NewImportTaskService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineImportTasks())
}
