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
				ValidateFunc: validation.StringInSlice([]string{"PrePaid", "PostPaidByBandwidth", "PostPaidByTraffic"}, false),
				Description:  "The billing type of the EIP Address. And optional choice contains `PostPaidByBandwidth` or `PostPaidByTraffic` or `PrePaid`.",
			},
			//"period_unit": {
			//	Type:     schema.TypeString,
			//	Optional: true,
			//	Default:  "Month",
			//	ValidateFunc: validation.StringInSlice([]string{
			//		"Month", "Year",
			//	}, false),
			//	DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
			//		// 创建时，只有付费类型为 PrePaid 时生效
			//		if d.Id() == "" {
			//			if d.Get("billing_type").(string) == "PrePaid" {
			//				return false
			//			}
			//		} else { // 修改时，只有付费类型由按量付费转为 PrePaid 时生效
			//			if d.HasChange("billing_type") && d.Get("billing_type").(string) == "PrePaid" {
			//				return false
			//			}
			//		}
			//		return true
			//	},
			//	Description: "The period unit of the EIP Address. Optional choice contains `Month` or `Year`. Default is `Month`." +
			//		"This field is only effective when creating a PrePaid Eip or changing the billing_type from PostPaid to PrePaid.",
			//},
			"period": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  12,
				ValidateFunc: validation.Any(
					validation.IntBetween(1, 9),
					validation.IntInSlice([]int{12, 36})),
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// 创建时，只有付费类型为 PrePaid 时生效
					if d.Id() == "" {
						if d.Get("billing_type").(string) == "PrePaid" {
							return false
						}
					} else { // 修改时，只有付费类型由按量付费转为 PrePaid 时生效
						if d.HasChange("billing_type") && d.Get("billing_type").(string) == "PrePaid" {
							return false
						}
					}
					return true
				},
				Description: "The period of the EIP Address, the valid value range in 1~9 or 12 or 36. Default value is 12. The period unit defaults to `Month`." +
					"This field is only effective when creating a PrePaid Eip or changing the billing_type from PostPaid to PrePaid.",
			},
			"bandwidth": {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.IntBetween(1, 500),
				Description:  "The peek bandwidth of the EIP, the value range in 1~500 for PostPaidByBandwidth, and 1~200 for PostPaidByTraffic.",
			},
			"isp": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				Description:  "The ISP of the EIP, the value can be `BGP` or `ChinaMobile` or `ChinaUnicom` or `ChinaTelecom` or `SingleLine_BGP` or `Static_BGP`.",
				ValidateFunc: validation.StringInSlice([]string{"BGP", "ChinaMobile", "ChinaUnicom", "ChinaTelecom", "SingleLine_BGP", "Static_BGP"}, false),
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
			"overdue_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The overdue time of the EIP.",
			},
			"deleted_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The deleted time of the EIP.",
			},
			"expired_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The expired time of the EIP.",
			},
		},
	}
}

func resourceVolcengineEipAddressCreate(d *schema.ResourceData, meta interface{}) error {
	eipAddressService := NewEipAddressService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Create(eipAddressService, d, ResourceVolcengineEipAddress()); err != nil {
		return fmt.Errorf("error on creating eip address  %q, %w", d.Id(), err)
	}
	return resourceVolcengineEipAddressRead(d, meta)
}

func resourceVolcengineEipAddressRead(d *schema.ResourceData, meta interface{}) error {
	eipAddressService := NewEipAddressService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Read(eipAddressService, d, ResourceVolcengineEipAddress()); err != nil {
		return fmt.Errorf("error on reading  eip address %q, %w", d.Id(), err)
	}
	return nil
}

func resourceVolcengineEipAddressUpdate(d *schema.ResourceData, meta interface{}) error {
	eipAddressService := NewEipAddressService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Update(eipAddressService, d, ResourceVolcengineEipAddress()); err != nil {
		return fmt.Errorf("error on updating  eip address %q, %w", d.Id(), err)
	}
	return resourceVolcengineEipAddressRead(d, meta)
}

func resourceVolcengineEipAddressDelete(d *schema.ResourceData, meta interface{}) error {
	eipAddressService := NewEipAddressService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Delete(eipAddressService, d, ResourceVolcengineEipAddress()); err != nil {
		return fmt.Errorf("error on deleting  eip address %q, %w", d.Id(), err)
	}
	return nil
}
