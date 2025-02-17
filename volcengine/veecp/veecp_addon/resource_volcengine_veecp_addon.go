package veecp_addon

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
VeecpAddon can be imported using the id, e.g.
```
$ terraform import volcengine_veecp_addon.default resource_id
```

*/

func ResourceVolcengineVeecpAddon() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVeecpAddonCreate,
		Read:   resourceVolcengineVeecpAddonRead,
		Update: resourceVolcengineVeecpAddonUpdate,
		Delete: resourceVolcengineVeecpAddonDelete,
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

func resourceVolcengineVeecpAddonCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVeecpAddonService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineVeecpAddon())
	if err != nil {
		return fmt.Errorf("error on creating veecp_addon %q, %s", d.Id(), err)
	}
	return resourceVolcengineVeecpAddonRead(d, meta)
}

func resourceVolcengineVeecpAddonRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVeecpAddonService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineVeecpAddon())
	if err != nil {
		return fmt.Errorf("error on reading veecp_addon %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineVeecpAddonUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVeecpAddonService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineVeecpAddon())
	if err != nil {
		return fmt.Errorf("error on updating veecp_addon %q, %s", d.Id(), err)
	}
	return resourceVolcengineVeecpAddonRead(d, meta)
}

func resourceVolcengineVeecpAddonDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVeecpAddonService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineVeecpAddon())
	if err != nil {
		return fmt.Errorf("error on deleting veecp_addon %q, %s", d.Id(), err)
	}
	return err
}
