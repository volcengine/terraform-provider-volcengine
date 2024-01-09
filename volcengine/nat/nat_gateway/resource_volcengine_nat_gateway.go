package nat_gateway

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
NatGateway can be imported using the id, e.g.
```
$ terraform import volcengine_nat_gateway.default ngw-vv3t043k05sm****
```

*/

func ResourceVolcengineNatGateway() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineNatGatewayCreate,
		Read:   resourceVolcengineNatGatewayRead,
		Update: resourceVolcengineNatGatewayUpdate,
		Delete: resourceVolcengineNatGatewayDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"vpc_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the VPC.",
			},
			"subnet_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the Subnet.",
			},
			"spec": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The specification of the NatGateway. Optional choice contains `Small`(default), `Medium`, `Large` or leave blank.",
			},
			"nat_gateway_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the NatGateway.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the NatGateway.",
			},
			"billing_type": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Default:     "PostPaid",
				Description: "The billing type of the NatGateway, the value is `PostPaid` or `PrePaid`.",
			},
			//"period_unit": {
			//	Type:     schema.TypeString,
			//	Optional: true,
			//	Default:  "Month",
			//	ForceNew: true,
			//	ValidateFunc: validation.StringInSlice([]string{
			//		"Month", "Year",
			//	}, false),
			//	DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
			//		return !(d.Get("billing_type").(string) == "PrePaid")
			//	},
			//	Description: "The period unit of the NatGateway. Optional choice contains `Month` or `Year`. Default is `Month`." +
			//		"This field is only effective when creating a PrePaid NatGateway. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
			//},
			"period": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  12,
				ForceNew: true,
				ValidateFunc: validation.Any(
					validation.IntBetween(1, 9),
					validation.IntInSlice([]int{12, 24, 36})),
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return !(d.Get("billing_type").(string) == "PrePaid")
				},
				Description: "The period of the NatGateway, the valid value range in 1~9 or 12 or 24 or 36. Default value is 12. The period unit defaults to `Month`." +
					"This field is only effective when creating a PrePaid NatGateway. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
			},
			"tags": ve.TagsSchema(),
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The ProjectName of the NatGateway.",
			},
		},
	}
}

func resourceVolcengineNatGatewayCreate(d *schema.ResourceData, meta interface{}) (err error) {
	natGatewayService := NewNatGatewayService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(natGatewayService, d, ResourceVolcengineNatGateway())
	if err != nil {
		return fmt.Errorf("error on creating nat gateway  %q, %w", d.Id(), err)
	}
	return resourceVolcengineNatGatewayRead(d, meta)
}

func resourceVolcengineNatGatewayRead(d *schema.ResourceData, meta interface{}) (err error) {
	natGatewayService := NewNatGatewayService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(natGatewayService, d, ResourceVolcengineNatGateway())
	if err != nil {
		return fmt.Errorf("error on reading nat gateway %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineNatGatewayUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	natGatewayService := NewNatGatewayService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(natGatewayService, d, ResourceVolcengineNatGateway())
	if err != nil {
		return fmt.Errorf("error on updating nat gateway  %q, %w", d.Id(), err)
	}
	return resourceVolcengineNatGatewayRead(d, meta)
}

func resourceVolcengineNatGatewayDelete(d *schema.ResourceData, meta interface{}) (err error) {
	natGatewayService := NewNatGatewayService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(natGatewayService, d, ResourceVolcengineNatGateway())
	if err != nil {
		return fmt.Errorf("error on deleting nat gateway %q, %w", d.Id(), err)
	}
	return err
}
