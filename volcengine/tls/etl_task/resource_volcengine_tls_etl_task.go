package etl_task

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
EtlTask can be imported using the id, e.g.
```
$ terraform import volcengine_etl_task.default resource_id
```

*/

func ResourceVolcengineEtlTask() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineEtlTaskCreate,
		Read:   resourceVolcengineEtlTaskRead,
		Update: resourceVolcengineEtlTaskUpdate,
		Delete: resourceVolcengineEtlTaskDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"dsl_type": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "DSL type, fixed as NORMAL. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
			},
			"enable": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeBool,
				Description: "Whether to enable the data processing task.",
			},
			"description": {
				Optional:    true,
				Type:        schema.TypeString,
				Description: "A simple description of the data processing task.",
			},
			"from_time": {
				Optional:    true,
				ForceNew:    true,
				Type:        schema.TypeInt,
				Description: "The start time of the data to be processed.",
			},
			"name": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "The name of the processing task.",
			},
			"script": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Processing rules.",
			},
			"source_topic_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "The log topic where the log to be processed is located.",
			},
			"target_resources": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "Output the relevant information of the target.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"alias": {
							Type:     schema.TypeString,
							Required: true,
							Description: "Customize the name of the output target, " +
								"which needs to be used to refer to the output target in the data processing rules.",
						},
						"topic_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Log topics used for storing processed logs.",
						},
						"role_trn": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Cross-account authorized character names.",
						},
					},
				},
			},
			"task_type": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "The task type is fixed as Resident.",
			},
			"to_time": {
				Optional:    true,
				ForceNew:    true,
				Type:        schema.TypeInt,
				Description: "The end time of the data to be processed.",
			},
		},
	}
	return resource
}

func resourceVolcengineEtlTaskCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEtlTaskService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineEtlTask())
	if err != nil {
		return fmt.Errorf("error on creating etl_task %q, %s", d.Id(), err)
	}
	return resourceVolcengineEtlTaskRead(d, meta)
}

func resourceVolcengineEtlTaskRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEtlTaskService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineEtlTask())
	if err != nil {
		return fmt.Errorf("error on reading etl_task %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineEtlTaskUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEtlTaskService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineEtlTask())
	if err != nil {
		return fmt.Errorf("error on updating etl_task %q, %s", d.Id(), err)
	}
	return resourceVolcengineEtlTaskRead(d, meta)
}

func resourceVolcengineEtlTaskDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEtlTaskService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineEtlTask())
	if err != nil {
		return fmt.Errorf("error on deleting etl_task %q, %s", d.Id(), err)
	}
	return err
}
