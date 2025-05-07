package veecp_batch_edge_machine

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
VeecpBatchEdgeMachine can be imported using the id, e.g.
```
$ terraform import volcengine_veecp_batch_edge_machine.default resource_id
```

*/

func ResourceVolcengineVeecpBatchEdgeMachine() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVeecpBatchEdgeMachineCreate,
		Read:   resourceVolcengineVeecpBatchEdgeMachineRead,
		Update: resourceVolcengineVeecpBatchEdgeMachineUpdate,
		Delete: resourceVolcengineVeecpBatchEdgeMachineDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The cluster id.",
			},
			"node_pool_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The node pool id.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the node.",
			},
			"ttl_hours": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "Effective hours of the managed script are counted from the creation time.",
			},
			"expiration_date": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Expiration date of the managed script, UTC time point, in seconds. If the expiration time is set, TTLHours will be ignored.",
			},
			"client_token": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The client token.",
			},
		},
	}
	return resource
}

func resourceVolcengineVeecpBatchEdgeMachineCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVeecpBatchEdgeMachineService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineVeecpBatchEdgeMachine())
	if err != nil {
		return fmt.Errorf("error on creating veecp_batch_edge_machine %q, %s", d.Id(), err)
	}
	return resourceVolcengineVeecpBatchEdgeMachineRead(d, meta)
}

func resourceVolcengineVeecpBatchEdgeMachineRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVeecpBatchEdgeMachineService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineVeecpBatchEdgeMachine())
	if err != nil {
		return fmt.Errorf("error on reading veecp_batch_edge_machine %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineVeecpBatchEdgeMachineUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVeecpBatchEdgeMachineService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineVeecpBatchEdgeMachine())
	if err != nil {
		return fmt.Errorf("error on updating veecp_batch_edge_machine %q, %s", d.Id(), err)
	}
	return resourceVolcengineVeecpBatchEdgeMachineRead(d, meta)
}

func resourceVolcengineVeecpBatchEdgeMachineDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVeecpBatchEdgeMachineService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineVeecpBatchEdgeMachine())
	if err != nil {
		return fmt.Errorf("error on deleting veecp_batch_edge_machine %q, %s", d.Id(), err)
	}
	return err
}
