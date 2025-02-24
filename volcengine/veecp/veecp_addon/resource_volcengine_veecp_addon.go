package veecp_addon

import (
	"fmt"
	"strings"
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

Notice
Some kind of VeecpAddon can not be removed from volcengine, and it will make a forbidden error when try to destroy.
If you want to remove it from terraform state, please use command
```
$ terraform state rm volcengine_veecp_addon.${name}
```

*/

func ResourceVolcengineVeecpAddon() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVeecpAddonCreate,
		Read:   resourceVolcengineVeecpAddonRead,
		Update: resourceVolcengineVeecpAddonUpdate,
		Delete: resourceVolcengineVeecpAddonDelete,
		CustomizeDiff: func(diff *schema.ResourceDiff, i interface{}) error {
			if diff.HasChange("config") {
				if n, ok := diff.Get("name").(string); ok && !checkSupportUpdate(n) {
					return diff.ForceNew("config")
				}
			}
			return nil
		},
		Importer: &schema.ResourceImporter{
			State: func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
				items := strings.Split(data.Id(), ":")
				if len(items) != 2 {
					return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
				}
				if err := data.Set("cluster_id", items[0]); err != nil {
					return []*schema.ResourceData{data}, err
				}
				if err := data.Set("name", items[1]); err != nil {
					return []*schema.ResourceData{data}, err
				}
				return []*schema.ResourceData{data}, nil
			},
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
				Description: "The cluster id of the addon.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the addon.",
			},
			"version": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				//ForceNew:    true,
				Description: "The version info of the cluster.",
			},
			"deploy_mode": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The deploy mode.",
			},
			"deploy_node_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The deploy node type.",
			},
			"config": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Description: "The config info of addon. " +
					"Please notice that `ingress-nginx` component prohibits updating config, can only works on the web console.",
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
