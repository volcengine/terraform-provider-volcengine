package eip_associate

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-vestack/common"
)

/*

Import
Eip associate can be imported using the eip allocation_id:instance_id, e.g.
```
$ terraform import vestack_eip_associate.default eip-274oj9a8rs9a87fap8sf9515b:i-cm9t9ug9lggu79yr5tcw
```

*/

func ResourceVestackEipAssociate() *schema.Resource {
	return &schema.Resource{
		Delete: resourceVestackEipAssociateDelete,
		Create: resourceVestackEipAssociateCreate,
		Read:   resourceVestackEipAssociateRead,
		Importer: &schema.ResourceImporter{
			State: eipAssociateImporter,
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
				Description: "The type of the associated instance.",
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

func resourceVestackEipAssociateCreate(d *schema.ResourceData, meta interface{}) error {
	eipAssociateService := NewEipAssociateService(meta.(*ve.SdkClient))
	if err := eipAssociateService.Dispatcher.Create(eipAssociateService, d, ResourceVestackEipAssociate()); err != nil {
		return fmt.Errorf("error on creating eip associate %q, %w", d.Id(), err)
	}
	return resourceVestackEipAssociateRead(d, meta)
}

func resourceVestackEipAssociateRead(d *schema.ResourceData, meta interface{}) error {
	eipAssociateService := NewEipAssociateService(meta.(*ve.SdkClient))
	if err := eipAssociateService.Dispatcher.Read(eipAssociateService, d, ResourceVestackEipAssociate()); err != nil {
		return fmt.Errorf("error on reading  eip associate %q, %w", d.Id(), err)
	}
	return nil
}

func resourceVestackEipAssociateDelete(d *schema.ResourceData, meta interface{}) error {
	eipAssociateService := NewEipAssociateService(meta.(*ve.SdkClient))
	if err := eipAssociateService.Dispatcher.Delete(eipAssociateService, d, ResourceVestackEipAssociate()); err != nil {
		return fmt.Errorf("error on deleting  eip associate %q, %w", d.Id(), err)
	}
	return nil
}
