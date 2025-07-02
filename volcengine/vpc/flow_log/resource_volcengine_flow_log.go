package flow_log

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
FlowLog can be imported using the id, e.g.
```
$ terraform import volcengine_flow_log.default resource_id
```

*/

func ResourceVolcengineFlowLog() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineFlowLogCreate,
		Read:   resourceVolcengineFlowLogRead,
		Update: resourceVolcengineFlowLogUpdate,
		Delete: resourceVolcengineFlowLogDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"flow_log_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of flow log.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of flow log.",
			},
			"resource_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The type of resource. Valid values: `vpc`, `subnet`, `eni`.",
			},
			"resource_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of resource.",
			},
			"traffic_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The type of traffic. Valid values: `All`, `Allow`, `Drop`.",
			},
			"log_project_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: "The name of log project. If there is no corresponding log project with the name, a new log project will be created. \n" +
					"When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
			},
			"log_topic_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: "The name of log topic. If there is no corresponding log topic with the name, a new log topic will be created. \n" +
					"When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
			},
			"aggregation_interval": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The aggregation interval of flow log. Unit: minute. Valid values: `1`, `5`, `10`.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The project name of flow log.",
			},
			"tags": ve.TagsSchema(),

			// computed fields
			"log_project_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of log project.",
			},
			"log_topic_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of log topic.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of flow log. Values: `Active`, `Pending`, `Inactive`, `Creating`, `Deleting`.",
			},
			"business_status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The business status of flow log.",
			},
			"lock_reason": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The reason why flow log is locked.",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The created time of flow log.",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The updated time of flow log.",
			},
		},
	}
	return resource
}

func resourceVolcengineFlowLogCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewFlowLogService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineFlowLog())
	if err != nil {
		return fmt.Errorf("error on creating flow_log %q, %s", d.Id(), err)
	}
	return resourceVolcengineFlowLogRead(d, meta)
}

func resourceVolcengineFlowLogRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewFlowLogService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineFlowLog())
	if err != nil {
		return fmt.Errorf("error on reading flow_log %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineFlowLogUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewFlowLogService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineFlowLog())
	if err != nil {
		return fmt.Errorf("error on updating flow_log %q, %s", d.Id(), err)
	}
	return resourceVolcengineFlowLogRead(d, meta)
}

func resourceVolcengineFlowLogDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewFlowLogService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineFlowLog())
	if err != nil {
		return fmt.Errorf("error on deleting flow_log %q, %s", d.Id(), err)
	}
	return err
}
