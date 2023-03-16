package eip_address

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Eip address can be imported using the id, e.g.
```
$ terraform import volcengine_eip_address.default eip-274oj9a8rs9a87fap8sf9515b
```

*/

func ResourceVolcengineEipAddress() *schema.Resource {
	return &schema.Resource{
		Delete: resourceVolcengineEipAddressDelete,
		Create: resourceVolcengineEipAddressCreate,
		Read:   resourceVolcengineEipAddressRead,
		Update: resourceVolcengineEipAddressUpdate,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"billing_type": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"PostPaidByBandwidth", "PostPaidByTraffic"}, false),
				Description:  "The billing type of the EIP Address. And optional choice contains `PostPaidByBandwidth` or `PostPaidByTraffic`.",
			},
			"bandwidth": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "The peek bandwidth of the EIP, the value range in 1~500 for PostPaidByBandwidth, and 1~200 for PostPaidByTraffic.",
			},
			"isp": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				Description:  "The ISP of the EIP, the value can be `BGP` or `ChinaMobile` or `ChinaUnicom` or `ChinaTelecom`.",
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
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ProjectName of the EIP.",
			},
			"tags": ve.TagsSchema(),
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

func resourceVolcengineEipAddressCreate(d *schema.ResourceData, meta interface{}) error {
	eipAddressService := NewEipAddressService(meta.(*ve.SdkClient))
	if err := eipAddressService.Dispatcher.Create(eipAddressService, d, ResourceVolcengineEipAddress()); err != nil {
		return fmt.Errorf("error on creating eip address  %q, %w", d.Id(), err)
	}
	return resourceVolcengineEipAddressRead(d, meta)
}

func resourceVolcengineEipAddressRead(d *schema.ResourceData, meta interface{}) error {
	eipAddressService := NewEipAddressService(meta.(*ve.SdkClient))
	if err := eipAddressService.Dispatcher.Read(eipAddressService, d, ResourceVolcengineEipAddress()); err != nil {
		return fmt.Errorf("error on reading  eip address %q, %w", d.Id(), err)
	}
	return nil
}

func resourceVolcengineEipAddressUpdate(d *schema.ResourceData, meta interface{}) error {
	eipAddressService := NewEipAddressService(meta.(*ve.SdkClient))
	if err := eipAddressService.Dispatcher.Update(eipAddressService, d, ResourceVolcengineEipAddress()); err != nil {
		return fmt.Errorf("error on updating  eip address %q, %w", d.Id(), err)
	}
	return resourceVolcengineEipAddressRead(d, meta)
}

func resourceVolcengineEipAddressDelete(d *schema.ResourceData, meta interface{}) error {
	eipAddressService := NewEipAddressService(meta.(*ve.SdkClient))
	if err := eipAddressService.Dispatcher.Delete(eipAddressService, d, ResourceVolcengineEipAddress()); err != nil {
		return fmt.Errorf("error on deleting  eip address %q, %w", d.Id(), err)
	}
	return nil
}
