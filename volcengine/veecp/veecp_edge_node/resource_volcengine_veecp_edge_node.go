package veecp_edge_node

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
VeecpNode can be imported using the id, e.g.
```
$ terraform import volcengine_veecp_node.default resource_id
```

*/

func ResourceVolcengineVeecpNode() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVeecpNodeCreate,
		Read:   resourceVolcengineVeecpNodeRead,
		Update: resourceVolcengineVeecpNodeUpdate,
		Delete: resourceVolcengineVeecpNodeDelete,
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
							Type:     schema.TypeString,
							Required: true,
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

func resourceVolcengineVeecpNodeCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVeecpNodeService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineVeecpNode())
	if err != nil {
		return fmt.Errorf("error on creating veecp_node %q, %s", d.Id(), err)
	}
	return resourceVolcengineVeecpNodeRead(d, meta)
}

func resourceVolcengineVeecpNodeRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVeecpNodeService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineVeecpNode())
	if err != nil {
		return fmt.Errorf("error on reading veecp_node %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineVeecpNodeUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVeecpNodeService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineVeecpNode())
	if err != nil {
		return fmt.Errorf("error on updating veecp_node %q, %s", d.Id(), err)
	}
	return resourceVolcengineVeecpNodeRead(d, meta)
}

func resourceVolcengineVeecpNodeDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVeecpNodeService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineVeecpNode())
	if err != nil {
		return fmt.Errorf("error on deleting veecp_node %q, %s", d.Id(), err)
	}
	return err
}
