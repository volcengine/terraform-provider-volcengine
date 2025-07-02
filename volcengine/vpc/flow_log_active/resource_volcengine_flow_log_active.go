package flow_log_active

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
FlowLogActive can be imported using the id, e.g.
```
$ terraform import volcengine_flow_log_active.default resource_id
```

*/

func ResourceVolcengineFlowLogActive() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineFlowLogActiveCreate,
		Read:   resourceVolcengineFlowLogActiveRead,
		Delete: resourceVolcengineFlowLogActiveDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"flow_log_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of flow log.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of flow log.",
			},
		},
	}
	return resource
}

func resourceVolcengineFlowLogActiveCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewFlowLogActiveService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineFlowLogActive())
	if err != nil {
		return fmt.Errorf("error on creating flow_log_active %q, %s", d.Id(), err)
	}
	return resourceVolcengineFlowLogActiveRead(d, meta)
}

func resourceVolcengineFlowLogActiveRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewFlowLogActiveService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineFlowLogActive())
	if err != nil {
		return fmt.Errorf("error on reading flow_log_active %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineFlowLogActiveUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewFlowLogActiveService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineFlowLogActive())
	if err != nil {
		return fmt.Errorf("error on updating flow_log_active %q, %s", d.Id(), err)
	}
	return resourceVolcengineFlowLogActiveRead(d, meta)
}

func resourceVolcengineFlowLogActiveDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewFlowLogActiveService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineFlowLogActive())
	if err != nil {
		return fmt.Errorf("error on deleting flow_log_active %q, %s", d.Id(), err)
	}
	return err
}
