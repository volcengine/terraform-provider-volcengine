package veecp_node_pool

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
VeecpNodePool can be imported using the id, e.g.
```
$ terraform import volcengine_veecp_node_pool.default resource_id
```

*/

func ResourceVolcengineVeecpNodePool() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVeecpNodePoolCreate,
		Read:   resourceVolcengineVeecpNodePoolRead,
		Update: resourceVolcengineVeecpNodePoolUpdate,
		Delete: resourceVolcengineVeecpNodePoolDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
		    // TODO: Add all your arguments and attributes.
			"replace_with_arguments": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// TODO: See setting, getting, flattening, expanding examples below for this complex argument.
			"complex_argument": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sub_field_one": {
							Type:         schema.TypeString,
							Required:     true,
						},
						"sub_field_two": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
	return resource
}

func resourceVolcengineVeecpNodePoolCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVeecpNodePoolService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineVeecpNodePool())
	if err != nil {
		return fmt.Errorf("error on creating veecp_node_pool %q, %s", d.Id(), err)
	}
	return resourceVolcengineVeecpNodePoolRead(d, meta)
}

func resourceVolcengineVeecpNodePoolRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVeecpNodePoolService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineVeecpNodePool())
	if err != nil {
		return fmt.Errorf("error on reading veecp_node_pool %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineVeecpNodePoolUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVeecpNodePoolService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineVeecpNodePool())
	if err != nil {
		return fmt.Errorf("error on updating veecp_node_pool %q, %s", d.Id(), err)
	}
	return resourceVolcengineVeecpNodePoolRead(d, meta)
}

func resourceVolcengineVeecpNodePoolDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVeecpNodePoolService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineVeecpNodePool())
	if err != nil {
		return fmt.Errorf("error on deleting veecp_node_pool %q, %s", d.Id(), err)
	}
	return err
}
