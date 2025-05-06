package ha_vip

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
HaVip can be imported using the id, e.g.
```
$ terraform import volcengine_ha_vip.default havip-2byzv8icq1b7k2dx0eegb****
```

*/

func ResourceVolcengineHaVip() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineHaVipCreate,
		Read:   resourceVolcengineHaVipRead,
		Update: resourceVolcengineHaVipUpdate,
		Delete: resourceVolcengineHaVipDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"subnet_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The subnet id of the Ha Vip.",
			},
			"ha_vip_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of the Ha Vip.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The description of the Ha Vip.",
			},
			"ip_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The ip address of the Ha Vip.",
			},

			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the Ha Vip.",
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The vpc id of the Ha Vip.",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The create time of the Ha Vip.",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The update time of the Ha Vip.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The project name of the Ha Vip.",
			},
			"tags": ve.TagsSchema(),
			"master_instance_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The master instance id of the Ha Vip.",
			},
			"associated_eip_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The associated eip id of the Ha Vip.",
			},
			"associated_eip_address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The associated eip address of the Ha Vip.",
			},
			"associated_instance_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The associated instance type of the Ha Vip.",
			},
			"associated_instance_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The associated instance ids of the Ha Vip.",
			},
		},
	}
	return resource
}

func resourceVolcengineHaVipCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewHaVipService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineHaVip())
	if err != nil {
		return fmt.Errorf("error on creating ha_vip %q, %s", d.Id(), err)
	}
	return resourceVolcengineHaVipRead(d, meta)
}

func resourceVolcengineHaVipRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewHaVipService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineHaVip())
	if err != nil {
		return fmt.Errorf("error on reading ha_vip %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineHaVipUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewHaVipService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineHaVip())
	if err != nil {
		return fmt.Errorf("error on updating ha_vip %q, %s", d.Id(), err)
	}
	return resourceVolcengineHaVipRead(d, meta)
}

func resourceVolcengineHaVipDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewHaVipService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineHaVip())
	if err != nil {
		return fmt.Errorf("error on deleting ha_vip %q, %s", d.Id(), err)
	}
	return err
}
