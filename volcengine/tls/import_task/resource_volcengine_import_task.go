package import_task

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
ImportTask can be imported using the id, e.g.
```
$ terraform import volcengine_import_task.default resource_id
```

*/

func ResourceVolcengineImportTask() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineImportTaskCreate,
		Read:   resourceVolcengineImportTaskRead,
		Update: resourceVolcengineImportTaskUpdate,
		Delete: resourceVolcengineImportTaskDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"description": {
				Optional:    true,
				Computed:    true,
				Type:        schema.TypeString,
				Description: "Data import task description.",
			},
			"import_source_info": {
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				Description: "The source information of the data import task.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tos_source_info": {
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							MaxItems:    1,
							Description: "TOS imports source information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"bucket": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The TOS bucket where the log file is located.",
									},
									"prefix": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The path of the file to be imported in the TOS bucket.",
									},
									"region": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The region where the TOS bucket is located. Support cross-regional data import.",
									},
									"compress_type": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The compression mode of data in the TOS bucket.",
									},
								},
							},
						},
						"kafka_source_info": {
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							MaxItems:    1,
							Description: "TOS imports source information.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"host": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The service addresses corresponding to different types of Kafka clusters are different.",
									},
									"group": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "Kafka consumer group.",
									},
									"topic": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "Kafka Topic name.",
									},
									"encode": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The encoding format of the data.",
									},
									"password": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The Kafka SASL user password used for identity authentication.",
									},
									"protocol": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "Secure Transport protocol.",
									},
									"username": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The Kafka SASL username used for identity authentication.",
									},
									"mechanism": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "Password authentication mechanism.",
									},
									"instance_id": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "When you are using the Volcano Engine Message Queue Kafka version, it should be set to the Kafka instance ID.",
									},
									"is_need_auth": {
										Type:        schema.TypeBool,
										Optional:    true,
										Computed:    true,
										Description: "Whether to enable authentication.",
									},
									"initial_offset": {
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
										Description: "The starting position of data import.",
									},
									"time_source_default": {
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
										Description: "Specify the log time.",
									},
								},
							},
						},
					},
				},
			},
			"project_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The log project ID used for storing data.",
			},
			"source_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Import the source type.",
			},
			"task_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Data import task name.",
			},
			"topic_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The log topic ID used for storing data.",
			},
			"target_info": {
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				Description: "The output information of the data import task.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"region": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Regional ID.",
						},
						"log_type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Specify the log parsing type when importing.",
						},
						"log_sample": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Log sample.",
						},
						"extract_rule": {
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							MaxItems:    1,
							Description: "Log extraction rules.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"keys": {
										Type:        schema.TypeSet,
										Optional:    true,
										Computed:    true,
										Description: "List of log field names (Keys).",
										Set:         schema.HashString,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"quote": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										Description: "Reference symbol. " +
											"The content wrapped by the reference will not be separated but will be parsed into a complete field." +
											" It is valid if and only if the LogType is delimiter_log.",
									},
									"time_key": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The field name of the log time field.",
									},
									"time_zone": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										Description: "Time zone, supporting both machine time zone (default) and custom time zone. " +
											"Among them, the custom time zone supports GMT and UTC.",
									},
									"delimiter": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "Log delimiter.",
									},
									"begin_regex": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										Description: "The regular expression used to identify the first line in each log, " +
											"and its matching part will serve as the beginning of the log.",
									},
									"time_format": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "The parsing format of the time field.",
									},
									"skip_line_count": {
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
										Description: "The number of log lines skipped.",
									},
									"un_match_log_key": {
										Type:        schema.TypeString,
										Optional:    true,
										Computed:    true,
										Description: "When uploading a log that failed to parse, the key name of the parse failed log.",
									},
									"time_extract_regex": {
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
										Description: "A regular expression for extracting time, " +
											"used to extract the time value in the TimeKey field and parse it into the corresponding collection time.",
									},
									"un_match_up_load_switch": {
										Type:        schema.TypeBool,
										Optional:    true,
										Computed:    true,
										Description: "Whether to upload the logs of failed parsing.",
									},
								},
							},
						},
					},
				},
			},
			"status": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Id() == ""
				},
				Description: "The status of the data import task.",
			},
		},
	}
	return resource
}

func resourceVolcengineImportTaskCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewImportTaskService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineImportTask())
	if err != nil {
		return fmt.Errorf("error on creating import_task %q, %s", d.Id(), err)
	}
	return resourceVolcengineImportTaskRead(d, meta)
}

func resourceVolcengineImportTaskRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewImportTaskService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineImportTask())
	if err != nil {
		return fmt.Errorf("error on reading import_task %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineImportTaskUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewImportTaskService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineImportTask())
	if err != nil {
		return fmt.Errorf("error on updating import_task %q, %s", d.Id(), err)
	}
	return resourceVolcengineImportTaskRead(d, meta)
}

func resourceVolcengineImportTaskDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewImportTaskService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineImportTask())
	if err != nil {
		return fmt.Errorf("error on deleting import_task %q, %s", d.Id(), err)
	}
	return err
}
