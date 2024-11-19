package transit_router_bandwidth_package

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
TransitRouterBandwidthPackage can be imported using the Id, e.g.
```
$ terraform import volcengine_transit_router_bandwidth_package.default tbp-cd-2felfww0i6pkw59gp68bq****
```

*/

func ResourceVolcengineTransitRouterBandwidthPackage() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineTransitRouterBandwidthPackageCreate,
		Read:   resourceVolcengineTransitRouterBandwidthPackageRead,
		Update: resourceVolcengineTransitRouterBandwidthPackageUpdate,
		Delete: resourceVolcengineTransitRouterBandwidthPackageDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"transit_router_bandwidth_package_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of the transit router bandwidth package.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The description of the transit router bandwidth package.",
			},
			"local_geographic_region_set_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "China",
				Description: "The local geographic region set ID. Valid values: `China`, `Asia`. Default is China.",
			},
			"peer_geographic_region_set_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "China",
				Description: "The peer geographic region set ID. Valid values: `China`, `Asia`. Default is China.",
			},
			"bandwidth": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     2,
				Description: "The bandwidth peak of the transit router bandwidth package. Unit: Mbps. Valid values: 2-10000. Default is 2 Mbps.",
			},
			"period": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  12,
				ValidateFunc: validation.Any(
					validation.IntBetween(1, 9),
					validation.IntInSlice([]int{12, 36})),
				DiffSuppressFunc: transitRouterBandwidthPackageDiffSuppress,
				Description: "The period of the transit router bandwidth package, the valid value range in 1~9 or 12 or 36. Default value is 12. The period unit defaults to `Month`." +
					"The modification of this field only takes effect when the value of the `renew_type` is `Manual`.",
			},
			"renew_type": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "Manual",
				ValidateFunc:     validation.StringInSlice([]string{"Manual", "Auto", "NoRenew"}, false),
				DiffSuppressFunc: transitRouterBandwidthPackageDiffSuppress,
				Description: "The renewal type of the transit router bandwidth package. Valid values: `Manual`, `Auto`, `NoRenew`. Default is `Manual`." +
					"This field is only effective when modifying the bandwidth package.",
			},
			"renew_period": {
				Type:             schema.TypeInt,
				Optional:         true,
				Default:          1,
				ValidateFunc:     validation.IntInSlice([]int{1, 2, 3, 6, 12}),
				DiffSuppressFunc: transitRouterBandwidthPackageDiffSuppress,
				Description: "The auto renewal period of the transit router bandwidth package. Valid values: 1,2,3,6,12. Default value is 1. Unit: Month." +
					"This field is only effective when the value of the `renew_type` is `Auto`. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
			},
			"remain_renew_times": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  -1,
				ValidateFunc: validation.Any(
					validation.IntBetween(1, 100),
					validation.IntInSlice([]int{-1})),
				DiffSuppressFunc: transitRouterBandwidthPackageDiffSuppress,
				Description: "The remaining renewal times of of the transit router bandwidth package. Valid values: -1 or 1~100. Default value is -1, means unlimited renewal." +
					"This field is only effective when the value of the `renew_type` is `Auto`.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The ProjectName of the transit router bandwidth package.",
			},
			"tags": ve.TagsSchema(),

			// computed fields
			"remaining_bandwidth": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The remaining bandwidth of the transit router bandwidth package. Unit: Mbps.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the transit router bandwidth package.",
			},
			"business_status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The business status of the transit router bandwidth package.",
			},
			"creation_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The create time of the transit router bandwidth package.",
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The update time of the transit router bandwidth package.",
			},
			"expired_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The expired time of the transit router bandwidth package.",
			},
			"delete_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The delete time of the transit router bandwidth package.",
			},
			"allocations": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The detailed information on cross regional connections associated with bandwidth packets.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"transit_router_peer_attachment_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the peer attachment.",
						},
						"allocate_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The delete time of the transit router bandwidth package.",
						},
						"local_region_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The local region id of the transit router.",
						},
						"delete_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The peer region id of the transit router.",
						},
					},
				},
			},
		},
	}
	return resource
}

func resourceVolcengineTransitRouterBandwidthPackageCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTRBandwidthPackageService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineTransitRouterBandwidthPackage())
	if err != nil {
		return fmt.Errorf("error on creating transit router bandwidth package %q, %s", d.Id(), err)
	}
	return resourceVolcengineTransitRouterBandwidthPackageRead(d, meta)
}

func resourceVolcengineTransitRouterBandwidthPackageRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTRBandwidthPackageService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineTransitRouterBandwidthPackage())
	if err != nil {
		return fmt.Errorf("error on reading transit router bandwidth package %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineTransitRouterBandwidthPackageUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTRBandwidthPackageService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineTransitRouterBandwidthPackage())
	if err != nil {
		return fmt.Errorf("error on updating transit router bandwidth package %q, %s", d.Id(), err)
	}
	return resourceVolcengineTransitRouterBandwidthPackageRead(d, meta)
}

func resourceVolcengineTransitRouterBandwidthPackageDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTRBandwidthPackageService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineTransitRouterBandwidthPackage())
	if err != nil {
		return fmt.Errorf("error on deleting transit router bandwidth package %q, %s", d.Id(), err)
	}
	return err
}
