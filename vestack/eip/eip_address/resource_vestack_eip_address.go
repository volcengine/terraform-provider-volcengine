package eip_address

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-vestack/common"
)

/*

Import
Eip address can be imported using the id, e.g.
```
$ terraform import volcstack_eip_address.default eip-274oj9a8rs9a87fap8sf9515b
```

*/

func ResourceVestackEipAddress() *schema.Resource {
	return &schema.Resource{
		Delete: resourceVestackEipAddressDelete,
		Create: resourceVestackEipAddressCreate,
		Read:   resourceVestackEipAddressRead,
		Update: resourceVestackEipAddressUpdate,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"billing_type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"PostPaidByBandwidth", "PostPaidByTraffic"}, false),
				Description:  "The billing type of the EIP Address. And optional choice contains `PostPaidByBandwidth` or `PostPaidByTraffic`.",
			},
			"bandwidth": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "The peek bandwidth of the EIP.",
			},
			"isp": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				Description:  "The ISP of the EIP.",
				ValidateFunc: validation.StringInSlice([]string{"BGP", "ChinaMobile", "ChinaUnicom", "ChinaTelecom"}, false),
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of the EIP Address.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the EIP.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the EIP.",
			},
			"eip_address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ip address of the EIP.",
			},
		},
	}
}

func resourceVestackEipAddressCreate(d *schema.ResourceData, meta interface{}) error {
	eipAddressService := NewEipAddressService(meta.(*ve.SdkClient))
	if err := eipAddressService.Dispatcher.Create(eipAddressService, d, ResourceVestackEipAddress()); err != nil {
		return fmt.Errorf("error on creating eip address  %q, %w", d.Id(), err)
	}
	return resourceVestackEipAddressRead(d, meta)
}

func resourceVestackEipAddressRead(d *schema.ResourceData, meta interface{}) error {
	eipAddressService := NewEipAddressService(meta.(*ve.SdkClient))
	if err := eipAddressService.Dispatcher.Read(eipAddressService, d, ResourceVestackEipAddress()); err != nil {
		return fmt.Errorf("error on reading  eip address %q, %w", d.Id(), err)
	}
	return nil
}

func resourceVestackEipAddressUpdate(d *schema.ResourceData, meta interface{}) error {
	eipAddressService := NewEipAddressService(meta.(*ve.SdkClient))
	if err := eipAddressService.Dispatcher.Update(eipAddressService, d, ResourceVestackEipAddress()); err != nil {
		return fmt.Errorf("error on updating  eip address %q, %w", d.Id(), err)
	}
	return resourceVestackEipAddressRead(d, meta)
}

func resourceVestackEipAddressDelete(d *schema.ResourceData, meta interface{}) error {
	eipAddressService := NewEipAddressService(meta.(*ve.SdkClient))
	if err := eipAddressService.Dispatcher.Delete(eipAddressService, d, ResourceVestackEipAddress()); err != nil {
		return fmt.Errorf("error on deleting  eip address %q, %w", d.Id(), err)
	}
	return nil
}
