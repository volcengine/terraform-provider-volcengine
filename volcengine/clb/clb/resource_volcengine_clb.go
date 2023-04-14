package clb

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
CLB can be imported using the id, e.g.
```
$ terraform import volcengine_clb.default clb-273y2ok6ets007fap8txvf6us
```

*/

func ResourceVolcengineClb() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineClbCreate,
		Read:   resourceVolcengineClbRead,
		Update: resourceVolcengineClbUpdate,
		Delete: resourceVolcengineClbDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(1 * time.Hour),
			Update: schema.DefaultTimeout(1 * time.Hour),
			Delete: schema.DefaultTimeout(1 * time.Hour),
		},
		Schema: map[string]*schema.Schema{
			"region_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The region of the request.",
			},
			"type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				Description:  "The type of the CLB. And optional choice contains `public` or `private`.",
				ValidateFunc: validation.StringInSlice([]string{"public", "private"}, false),
			},
			"load_balancer_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of the CLB.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the CLB.",
			},
			"vpc_id": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
				Description: "The id of the VPC.",
			},
			"subnet_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the Subnet.",
			},
			"eni_address": {
				Type:        schema.TypeString,
				Computed:    true,
				ForceNew:    true,
				Optional:    true,
				Description: "The eni address of the CLB.",
			},
			"modification_protection_status": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "The status of the console modification protection, the value can be `NonProtection` or `ConsoleProtection`.",
				ValidateFunc: validation.StringInSlice([]string{"NonProtection", "ConsoleProtection"}, false),
			},
			"modification_protection_reason": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The reason of the console modification protection.",
			},
			"load_balancer_spec": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The specification of the CLB, the value can be `small_1`, `small_2`, `medium_1`, `medium_2`, `large_1`, `large_2`.",
			},
			"load_balancer_billing_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ForceNew:     true,
				Description:  "The billing type of the CLB, the value can be `PostPaid`.",
				ValidateFunc: validation.StringInSlice([]string{"PostPaid"}, false),
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ProjectName of the CLB.",
			},
			"tags": ve.TagsSchema(),
			"master_zone_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The master zone ID of the CLB.",
			},
			"slave_zone_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The slave zone ID of the CLB.",
			},
			//"period_unit": {
			//	Type:         schema.TypeString,
			//	Optional:     true,
			//	Description:  "The period unit of PrePaid billing type.",
			//	ValidateFunc: validation.StringInSlice([]string{"Month", "Year"}, false),
			//	DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
			//		return d.Id() != ""
			//	},
			//},
			//"period": {
			//	Type:        schema.TypeInt,
			//	Optional:    true,
			//	Description: "The period of PrePaid billing type.",
			//	DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
			//		return d.Id() != ""
			//	},
			//},
		},
	}
}

func resourceVolcengineClbCreate(d *schema.ResourceData, meta interface{}) (err error) {
	clbService := NewClbService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(clbService, d, ResourceVolcengineClb())
	if err != nil {
		return fmt.Errorf("error on creating clb  %q, %w", d.Id(), err)
	}
	return resourceVolcengineClbRead(d, meta)
}

func resourceVolcengineClbRead(d *schema.ResourceData, meta interface{}) (err error) {
	clbService := NewClbService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(clbService, d, ResourceVolcengineClb())
	if err != nil {
		return fmt.Errorf("error on reading clb %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineClbUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	clbService := NewClbService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(clbService, d, ResourceVolcengineClb())
	if err != nil {
		return fmt.Errorf("error on updating clb  %q, %w", d.Id(), err)
	}
	return resourceVolcengineClbRead(d, meta)
}

func resourceVolcengineClbDelete(d *schema.ResourceData, meta interface{}) (err error) {
	clbService := NewClbService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(clbService, d, ResourceVolcengineClb())
	if err != nil {
		return fmt.Errorf("error on deleting clb %q, %w", d.Id(), err)
	}
	return err
}
