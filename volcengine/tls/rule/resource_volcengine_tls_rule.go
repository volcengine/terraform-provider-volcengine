package rule

import (
	"bytes"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
tls rule can be imported using the id, e.g.
```
$ terraform import volcengine_tls_rule.default fa************
```

*/

func ResourceVolcengineTlsRule() *schema.Resource {
	return &schema.Resource{
		Read:   resourceVolcengineTlsRuleRead,
		Create: resourceVolcengineTlsRuleCreate,
		Delete: resourceVolcengineTlsRuleDelete,
		Update: resourceVolcengineTlsRuleUpdate,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"topic_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "The ID of the log topic to which the collection configuration belongs.",
			},
			"rule_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the collection configuration.",
			},
			"paths": {
				Type:     schema.TypeSet,
				Optional: true,
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				MaxItems:    10,
				Description: "Collection path list.",
			},
			"log_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "minimalist_log",
				Description: "The log type. The value can be one of the following: `minimalist_log`, `json_log`, `delimiter_log`, `multiline_log`, `fullregex_log`.",
			},
			"log_sample": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The sample of the log.",
			},
			"input_type": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     0,
				Description: "The type of the collection configuration. Validate value can be `0`(host log file), `1`(K8s container standard output) and `2`(Log files in the K8s container).",
			},
			"exclude_paths": {
				Type:        schema.TypeSet,
				Optional:    true,
				Set:         tlsRuleHash("type", "value"),
				Description: "Collect the blacklist list.",
				MaxItems:    10,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Collection path type. The path type can be `File` or `Path`.",
						},
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Collection path.",
						},
					},
				},
			},
			"extract_rule": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Description: "The extract rule.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"delimiter": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The delimiter of the log.",
						},
						"begin_regex": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The first log line needs to match the regular expression.",
						},
						"log_regex": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The entire log needs to match the regular expression.",
						},
						"keys": {
							Type:        schema.TypeSet,
							Optional:    true,
							Computed:    true,
							Set:         schema.HashString,
							Description: "A list of log field names (Key).",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"time_key": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The field name of the log time field.",
						},
						"time_format": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Parsing format of the time field.",
						},
						"filter_key_regex": {
							Type:        schema.TypeSet,
							Optional:    true,
							Set:         tlsRuleHash("key", "regex"),
							Description: "The filter key list.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The name of the filter key.",
									},
									"regex": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The log content of the filter field needs to match the regular expression.",
									},
								},
							},
						},
						"un_match_up_load_switch": {
							Type:        schema.TypeBool,
							Optional:    true,
							Computed:    true,
							Description: "Whether to upload the log of parsing failure.",
						},
						"un_match_log_key": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "When uploading the failed log, the key name of the failed log.",
						},
						"quote": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The quote symbol.",
						},
						"time_zone": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The time zone.",
						},
						"log_template": {
							Type:        schema.TypeSet,
							Optional:    true,
							MaxItems:    1,
							Set:         tlsRuleHash("type", "format"),
							Description: "Automatically extract log fields according to the specified log template.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"type": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The type of the log template.",
									},
									"format": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Log template content.",
									},
								},
							},
						},
					},
				},
			},
			"user_define_rule": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "User-defined collection rules.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_raw_log": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether to upload raw logs.",
						},
						"fields": {
							Type:        schema.TypeMap,
							Optional:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Description: "Add constant fields to logs.",
						},
						"tail_files": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "LogCollector collection strategy, which specifies whether LogCollector collects incremental logs or full logs. The default is false, which means to collect all logs.",
						},
						"parse_path_rule": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "Rules for parsing collection paths. After the rules are set, the fields in the collection path will be extracted through the regular expressions specified in the rules, and added to the log data as metadata.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"path_sample": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Sample capture path for a real scene.",
									},
									"regex": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Regular expression for extracting path fields. It must match the collection path sample, otherwise it cannot be extracted successfully.",
									},
									"keys": {
										Type:        schema.TypeSet,
										Optional:    true,
										Set:         schema.HashString,
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
							Optional:    true,
							MaxItems:    1,
							Description: "Rules for routing log partitions. Setting this parameter indicates that the HashKey routing shard mode is used when collecting logs, and Log Service will write the data to the shard containing the specified Key value.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"hash_key": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The HashKey of the log group is used to specify the partition (shard) to be written to by the current log group.",
									},
								},
							},
						},
						"plugin": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "Plugin configuration. After the plugin configuration is enabled, one or more LogCollector processor plugins can be added to parse logs with complex or variable structures.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"processors": {
										Type:        schema.TypeSet,
										Required:    true,
										Set:         schema.HashString,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Description: "LogCollector plugin.",
									},
								},
							},
						},
						"advanced": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "LogCollector extension configuration.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"close_inactive": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The wait time to release the log file handle. When the log file has not written a new log for more than the specified time, release the handle of the log file.",
									},
									"close_removed": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "After the log file is removed, whether to release the handle of the log file. The default is false.",
									},
									"close_renamed": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "After the log file is renamed, whether to release the handle of the log file. The default is false.",
									},
									// warning
									"close_eof": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether to release the log file handle after reading to the end of the log file. The default is false.",
									},
									"close_timeout": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The maximum length of time that LogCollector monitors log files. The unit is seconds, and the default is 0 seconds, which means that there is no limit to the length of time LogCollector monitors log files.",
									},
								},
							},
						},
					},
				},
			},
			"container_rule": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Container collection rules.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"stream": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The collection mode.",
						},
						"container_name_regex": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the container to be collected.",
						},
						"include_container_label_regex": {
							Type:        schema.TypeMap,
							Optional:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Description: "The container label whitelist specifies the containers to be collected through the container label. If the whitelist is not enabled, all containers are specified to be collected.",
						},
						"exclude_container_label_regex": {
							Type:        schema.TypeMap,
							Optional:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Description: "The container Label blacklist is used to specify the range of containers not to be collected.",
						},
						"include_container_env_regex": {
							Type:        schema.TypeMap,
							Optional:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Description: "The container environment variable whitelist specifies the container to be collected through the container environment variable. If the whitelist is not enabled, it means that all containers are specified to be collected.",
						},
						"exclude_container_env_regex": {
							Type:        schema.TypeMap,
							Optional:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Description: "The container environment variable blacklist is used to specify the range of containers not to be collected.",
						},
						"env_tag": {
							Type:        schema.TypeMap,
							Optional:    true,
							Elem:        &schema.Schema{Type: schema.TypeString},
							Description: "Whether to add environment variables as log tags to raw log data.",
						},
						"kubernetes_rule": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "Collection rules for Kubernetes containers.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"namespace_name_regex": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The name of the Kubernetes Namespace to be collected. If no Namespace name is specified, all containers will be collected. Namespace names support regular matching.",
									},
									"workload_type": {
										Type:     schema.TypeString,
										Optional: true,
										Description: "Specify the containers to be collected by the type of workload, only one type can be selected. " +
											"When no type is specified, it means all types of containers are collected. The supported types of workloads are:\n" +
											"Deployment: stateless workload.\nStatefulSet: stateful workload.\n" +
											"DaemonSet: daemon process.\nJob: task.\nCronJob: scheduled task.",
									},
									"workload_name_regex": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specify the container to be collected by the name of the workload. When no workload name is specified, all containers are collected. The workload name supports regular matching.",
									},
									"include_pod_label_regex": {
										Type:        schema.TypeMap,
										Optional:    true,
										Description: "The Pod Label whitelist is used to specify containers to be collected. When the Pod Label whitelist is not enabled, it means that all containers are collected.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"exclude_pod_label_regex": {
										Type:        schema.TypeMap,
										Optional:    true,
										Description: "Specify the containers not to be collected through the Pod Label blacklist, and not enable means to collect all containers.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"pod_name_regex": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The Pod name is used to specify the container to be collected. When no Pod name is specified, it means to collect all containers.",
									},
									"label_tag": {
										Type:        schema.TypeMap,
										Optional:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Description: "Whether to add Kubernetes Label as a log label to the original log data.",
									},
									"annotation_tag": {
										Type:        schema.TypeMap,
										Optional:    true,
										Elem:        &schema.Schema{Type: schema.TypeString},
										Description: "Whether to add Kubernetes Annotation as a log tag to the raw log data.",
									},
								},
							},
						},
					},
				},
			},
			"rule_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The id of the rule.",
			},
		},
	}
}

func resourceVolcengineTlsRuleCreate(d *schema.ResourceData, meta interface{}) error {
	TlsRuleService := NewTlsRuleService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Create(TlsRuleService, d, ResourceVolcengineTlsRule()); err != nil {
		return fmt.Errorf("error on creating tls rule  %q, %w", d.Id(), err)
	}
	return resourceVolcengineTlsRuleRead(d, meta)
}

func resourceVolcengineTlsRuleRead(d *schema.ResourceData, meta interface{}) error {
	TlsRuleService := NewTlsRuleService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Read(TlsRuleService, d, ResourceVolcengineTlsRule()); err != nil {
		return fmt.Errorf("error on reading tls rule %q, %w", d.Id(), err)
	}
	return nil
}

func resourceVolcengineTlsRuleUpdate(d *schema.ResourceData, meta interface{}) error {
	TlsRuleService := NewTlsRuleService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Update(TlsRuleService, d, ResourceVolcengineTlsRule()); err != nil {
		return fmt.Errorf("error on updating tls rule %q, %w", d.Id(), err)
	}
	return resourceVolcengineTlsRuleRead(d, meta)
}

func resourceVolcengineTlsRuleDelete(d *schema.ResourceData, meta interface{}) error {
	TlsRuleService := NewTlsRuleService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Delete(TlsRuleService, d, ResourceVolcengineTlsRule()); err != nil {
		return fmt.Errorf("error on deleting tls rule %q, %w", d.Id(), err)
	}
	return nil
}

func tlsRuleHash(key, value string) func(v interface{}) int {
	return func(v interface{}) int {
		if v == nil {
			return 0
		}
		m, ok := v.(map[string]interface{})
		if !ok {
			return hashcode.String(fmt.Sprintf("%v", v))
		}
		var (
			buf bytes.Buffer
		)
		v1 := fmt.Sprintf("%v", m[key])
		v2 := fmt.Sprintf("%v", m[value])
		buf.WriteString(fmt.Sprintf("%v#%v", v1, v2))
		return hashcode.String(buf.String())
	}
}
