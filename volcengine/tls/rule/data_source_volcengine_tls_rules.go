package rule

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineTlsRules() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineTlsRulesRead,
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The project id.",
			},
			"rule_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The rule id.",
			},
			"rule_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The rule name.",
			},
			"topic_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The topic id.",
			},
			"topic_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The topic name.",
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
			"rules": {
				Description: "The rules list.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"topic_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The topic id.",
						},
						"topic_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The topic name.",
						},
						"rule_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The rule id.",
						},
						"rule_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The rule name.",
						},
						"paths": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "Collection path list.",
						},
						"log_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The log type.",
						},
						"extract_rule": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The extract rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"delimiter": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The delimiter of the log.",
									},
									"begin_regex": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The first log line needs to match the regular expression.",
									},
									"log_regex": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The entire log needs to match the regular expression.",
									},
									"keys": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "A list of log field names (Key).",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"time_key": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The field name of the log time field.",
									},
									"time_format": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Parsing format of the time field.",
									},
									"filter_key_regex": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The filter key list.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"key": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The name of the filter key.",
												},
												"regex": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The log content of the filter field needs to match the regular expression.",
												},
											},
										},
									},
									"un_match_up_load_switch": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether to upload the log of parsing failure.",
									},
									"un_match_log_key": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "When uploading the failed log, the key name of the failed log.",
									},
									"log_template": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Automatically extract log fields according to the specified log template.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"type": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The type of the log template.",
												},
												"format": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Log template content.",
												},
											},
										},
									},
								},
							},
						},
						"exclude_paths": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Collect the blacklist list.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Collection path type. The path type can be `File` or `Path`.",
									},
									"value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Collection path.",
									},
								},
							},
						},
						"log_sample": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Log sample.",
						},
						"user_define_rule": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "User-defined collection rules.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enable_raw_log": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether to upload raw logs.",
									},
									"fields": {
										Type:        schema.TypeMap,
										Computed:    true,
										Elem:        schema.TypeString,
										Description: "Add constant fields to logs.",
									},
									"tail_files": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "LogCollector collection strategy, which specifies whether LogCollector collects incremental logs or full logs. The default is false, which means to collect all logs.",
									},
									"parse_path_rule": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Rules for parsing collection paths. After the rules are set, the fields in the collection path will be extracted through the regular expressions specified in the rules, and added to the log data as metadata.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"path_sample": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Sample capture path for a real scene.",
												},
												"regex": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Regular expression for extracting path fields. It must match the collection path sample, otherwise it cannot be extracted successfully.",
												},
												"keys": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "A list of field names. Log Service will parse the path sample (PathSample) into multiple fields according to the regular expression (Regex), and Keys is used to specify the field name of each field.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"shard_hash_key": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Rules for routing log partitions. Setting this parameter indicates that the HashKey routing shard mode is used when collecting logs, and Log Service will write the data to the shard containing the specified Key value.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"hash_key": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The HashKey of the log group is used to specify the partition (shard) to be written to by the current log group.",
												},
											},
										},
									},
									"plugin": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Plugin configuration. After the plugin configuration is enabled, one or more LogCollector processor plugins can be added to parse logs with complex or variable structures.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												// warning
												"processors": {
													Type:        schema.TypeList,
													Computed:    true,
													Elem:        &schema.Schema{Type: schema.TypeString},
													Description: "LogCollector plugin.",
												},
											},
										},
									},
									"advanced": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "LogCollector extension configuration.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"close_inactive": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The wait time to release the log file handle. When the log file has not written a new log for more than the specified time, release the handle of the log file.",
												},
												"close_removed": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "After the log file is removed, whether to release the handle of the log file. The default is false.",
												},
												"close_renamed": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "After the log file is renamed, whether to release the handle of the log file. The default is false.",
												},
												// warning
												"close_eof": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Whether to release the log file handle after reading to the end of the log file. The default is false.",
												},
												"close_timeout": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The maximum length of time that LogCollector monitors log files. The unit is seconds, and the default is 0 seconds, which means that there is no limit to the length of time LogCollector monitors log files.",
												},
											},
										},
									},
								},
							},
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The creation time.",
						},
						"modify_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The modification time.",
						},
						"input_type": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The collection type.",
						},
						"container_rule": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Container collection rules.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"stream": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The collection mode.",
									},
									"container_name_regex": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the container to be collected.",
									},
									"include_container_label_regex": {
										Type:        schema.TypeMap,
										Computed:    true,
										Elem:        schema.TypeString,
										Description: "The container label whitelist specifies the containers to be collected through the container label. If the whitelist is not enabled, all containers are specified to be collected.",
									},
									"exclude_container_label_regex": {
										Type:        schema.TypeMap,
										Computed:    true,
										Elem:        schema.TypeString,
										Description: "The container Label blacklist is used to specify the range of containers not to be collected.",
									},
									"include_container_env_regex": {
										Type:        schema.TypeMap,
										Computed:    true,
										Elem:        schema.TypeString,
										Description: "The container environment variable whitelist specifies the container to be collected through the container environment variable. If the whitelist is not enabled, it means that all containers are specified to be collected.",
									},
									"exclude_container_env_regex": {
										Type:        schema.TypeMap,
										Computed:    true,
										Elem:        schema.TypeString,
										Description: "The container environment variable blacklist is used to specify the range of containers not to be collected.",
									},
									"env_tag": {
										Type:        schema.TypeMap,
										Computed:    true,
										Elem:        schema.TypeString,
										Description: "Whether to add environment variables as log tags to raw log data.",
									},
									"kubernetes_rule": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Collection rules for Kubernetes containers.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"namespace_name_regex": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The name of the Kubernetes Namespace to be collected. If no Namespace name is specified, all containers will be collected. Namespace names support regular matching.",
												},
												"workload_type": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specify the container to be collected by the type of workload. Only one type can be selected. When no type is specified, it means to collect all types of containers.",
												},
												"workload_name_regex": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specify the container to be collected by the name of the workload. When no workload name is specified, all containers are collected. The workload name supports regular matching.",
												},
												"include_pod_label_regex": {
													Type:        schema.TypeMap,
													Computed:    true,
													Description: "The Pod Label whitelist is used to specify containers to be collected. When the Pod Label whitelist is not enabled, it means that all containers are collected.",
													Elem:        schema.TypeString,
												},
												"exclude_pod_label_regex": {
													Type:        schema.TypeMap,
													Computed:    true,
													Description: "Specify the containers not to be collected through the Pod Label blacklist, and not enable means to collect all containers.",
													Elem:        schema.TypeString,
												},
												"pod_name_regex": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The Pod name is used to specify the container to be collected. When no Pod name is specified, it means to collect all containers.",
												},
												"label_tag": {
													Type:        schema.TypeMap,
													Computed:    true,
													Elem:        schema.TypeString,
													Description: "Whether to add Kubernetes Label as a log label to the original log data.",
												},
												"annotation_tag": {
													Type:        schema.TypeMap,
													Computed:    true,
													Elem:        schema.TypeString,
													Description: "Whether to add Kubernetes Annotation as a log tag to the raw log data.",
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

func dataSourceVolcengineTlsRulesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewTlsRuleService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineTlsRules())
}
