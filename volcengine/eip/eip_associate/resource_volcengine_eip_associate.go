package eip_associate

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Eip associate can be imported using the eip allocation_id:instance_id, e.g.
```
$ terraform import volcengine_eip_associate.default eip-274oj9a8rs9a87fap8sf9515b:i-cm9t9ug9lggu79yr5tcw
```

*/

func ResourceVolcengineEipAssociate() *schema.Resource {
	return &schema.Resource{
		Delete: resourceVolcengineEipAssociateDelete,
		Create: resourceVolcengineEipAssociateCreate,
		Read:   resourceVolcengineEipAssociateRead,
		Importer: &schema.ResourceImporter{
			State: eipAssociateImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"allocation_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The allocation id of the EIP.",
			},
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The instance id which be associated to the EIP.",
			},
			"instance_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The type of the associated instance,the value is `NAT` or `NetworkInterface` or `ClbInstance` or `EcsInstance` or `HaVip`.",
			},
			"private_ip_address": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "The private IP address of the instance will be associated to the EIP.",
			},
		},
	}
}

func resourceVolcengineEipAssociateCreate(d *schema.ResourceData, meta interface{}) error {
	eipAssociateService := NewEipAssociateService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Create(eipAssociateService, d, ResourceVolcengineEipAssociate()); err != nil {
		return fmt.Errorf("error on creating eip associate %q, %w", d.Id(), err)
	}
	return resourceVolcengineEipAssociateRead(d, meta)
}

func resourceVolcengineEipAssociateRead(d *schema.ResourceData, meta interface{}) error {
	eipAssociateService := NewEipAssociateService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Read(eipAssociateService, d, ResourceVolcengineEipAssociate()); err != nil {
		return fmt.Errorf("error on reading  eip associate %q, %w", d.Id(), err)
	}
	return nil
}

func resourceVolcengineEipAssociateDelete(d *schema.ResourceData, meta interface{}) error {
	eipAssociateService := NewEipAssociateService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Delete(eipAssociateService, d, ResourceVolcengineEipAssociate()); err != nil {
		return fmt.Errorf("error on deleting  eip associate %q, %w", d.Id(), err)
	}
	return nil
}