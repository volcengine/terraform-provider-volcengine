package vefaas_timer

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
VefaasTimer can be imported using the id, e.g.
```
$ terraform import volcengine_vefaas_timer.default FunctionId:Id
```

*/

func ResourceVolcengineVefaasTimer() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVefaasTimerCreate,
		Read:   resourceVolcengineVefaasTimerRead,
		Update: resourceVolcengineVefaasTimerUpdate,
		Delete: resourceVolcengineVefaasTimerDelete,
		Importer: &schema.ResourceImporter{
			State: vefaasTimerImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"function_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of Function.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the Timer trigger.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The description of the Timer trigger.",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Whether the Timer trigger is enabled.",
			},
			"crontab": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Set the timing trigger time of the Timer trigger.",
			},
			"payload": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The Timer trigger sends the content payload of the request.",
			},
			"enable_concurrency": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Whether the Timer trigger allows concurrency.",
			},
			"retries": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "The retry count of the Timer trigger.",
			},
			"creation_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The creation time of the Timer trigger.",
			},
			"last_update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The last update time of the Timer trigger.",
			},
		},
	}
	return resource
}

func resourceVolcengineVefaasTimerCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVefaasTimerService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineVefaasTimer())
	if err != nil {
		return fmt.Errorf("error on creating vefaas_timer %q, %s", d.Id(), err)
	}
	return resourceVolcengineVefaasTimerRead(d, meta)
}

func resourceVolcengineVefaasTimerRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVefaasTimerService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineVefaasTimer())
	if err != nil {
		return fmt.Errorf("error on reading vefaas_timer %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineVefaasTimerUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVefaasTimerService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineVefaasTimer())
	if err != nil {
		return fmt.Errorf("error on updating vefaas_timer %q, %s", d.Id(), err)
	}
	return resourceVolcengineVefaasTimerRead(d, meta)
}

func resourceVolcengineVefaasTimerDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVefaasTimerService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineVefaasTimer())
	if err != nil {
		return fmt.Errorf("error on deleting vefaas_timer %q, %s", d.Id(), err)
	}
	return err
}
