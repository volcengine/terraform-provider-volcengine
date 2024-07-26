package private_zone

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
PrivateZone can be imported using the id, e.g.
```
$ terraform import volcengine_private_zone.default resource_id
```

*/

func ResourceVolcenginePrivateZone() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcenginePrivateZoneCreate,
		Read:   resourceVolcenginePrivateZoneRead,
		Update: resourceVolcenginePrivateZoneUpdate,
		Delete: resourceVolcenginePrivateZoneDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"zone_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the private zone.",
			},
			"remark": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The remark of the private zone.",
			},
			"recursion_mode": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether to enable the recursion mode of the private zone.",
			},
			"intelligent_mode": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Whether to enable the intelligent mode of the private zone.",
			},
			"load_balance_mode": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether to enable the load balance mode of the private zone.",
			},
			"vpcs": {
				Type:        schema.TypeSet,
				Required:    true,
				MinItems:    1,
				Description: "The bind vpc object of the private zone. If you want to bind another account's VPC, you need to first use resource volcengine_private_zone_user_vpc_authorization to complete the authorization.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"region": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "The region of the bind vpc. The default value is the region of the default provider config.",
						},
						"vpc_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The id of the bind vpc.",
						},
					},
				},
			},
		},
	}
	return resource
}

func resourceVolcenginePrivateZoneCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewPrivateZoneService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcenginePrivateZone())
	if err != nil {
		return fmt.Errorf("error on creating private_zone %q, %s", d.Id(), err)
	}
	return resourceVolcenginePrivateZoneRead(d, meta)
}

func resourceVolcenginePrivateZoneRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewPrivateZoneService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcenginePrivateZone())
	if err != nil {
		return fmt.Errorf("error on reading private_zone %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcenginePrivateZoneUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewPrivateZoneService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcenginePrivateZone())
	if err != nil {
		return fmt.Errorf("error on updating private_zone %q, %s", d.Id(), err)
	}
	return resourceVolcenginePrivateZoneRead(d, meta)
}

func resourceVolcenginePrivateZoneDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewPrivateZoneService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcenginePrivateZone())
	if err != nil {
		return fmt.Errorf("error on deleting private_zone %q, %s", d.Id(), err)
	}
	return err
}
