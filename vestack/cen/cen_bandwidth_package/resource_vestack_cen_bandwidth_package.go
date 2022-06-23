package cen_bandwidth_package

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-vestack/common"
)

/*

Import
CenBandwidthPackage can be imported using the id, e.g.
```
$ terraform import vestack_cen_bandwidth_package.default cbp-4c2zaavbvh5f42****
```

*/

func ResourceVestackCenBandwidthPackage() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVestackCenBandwidthPackageCreate,
		Read:   resourceVestackCenBandwidthPackageRead,
		Update: resourceVestackCenBandwidthPackageUpdate,
		Delete: resourceVestackCenBandwidthPackageDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"local_geographic_region_set_id": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Default:      "China",
				ValidateFunc: validation.StringInSlice([]string{"China"}, false),
				Description:  "The local geographic region set id of the cen bandwidth package.",
			},
			"peer_geographic_region_set_id": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Default:      "China",
				ValidateFunc: validation.StringInSlice([]string{"China"}, false),
				Description:  "The peer geographic region set id of the cen bandwidth package.",
			},
			"bandwidth": {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.IntBetween(2, 10000),
				Description:  "The bandwidth of the cen bandwidth package.",
			},
			"cen_bandwidth_package_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of the cen bandwidth package.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The description of the cen bandwidth package.",
			},
			"billing_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Default:      "PrePaid",
				ValidateFunc: validation.StringInSlice([]string{"PrePaid"}, false),
				Description:  "The billing type of the cen bandwidth package. Terraform will only remove the PrePaid cen bandwidth package from the state file, not actually remove.",
			},
			"period_unit": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "Month",
				ValidateFunc:     validation.StringInSlice([]string{"Month", "Year"}, false),
				DiffSuppressFunc: periodDiffSuppress,
				Description:      "The period unit of the cen bandwidth package.",
			},
			"period": {
				Type:             schema.TypeInt,
				Optional:         true,
				Default:          1,
				DiffSuppressFunc: periodDiffSuppress,
				Description:      "The period of the cen bandwidth package.",
			},
		},
	}
	s := DataSourceVestackCenBandwidthPackages().Schema["bandwidth_packages"].Elem.(*schema.Resource).Schema
	delete(s, "id")
	ve.MergeDateSourceToResource(s, &resource.Schema)
	return resource
}

func resourceVestackCenBandwidthPackageCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCenBandwidthPackageService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVestackCenBandwidthPackage())
	if err != nil {
		return fmt.Errorf("error on creating cen bandwidth package %q, %s", d.Id(), err)
	}
	return resourceVestackCenBandwidthPackageRead(d, meta)
}

func resourceVestackCenBandwidthPackageRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCenBandwidthPackageService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVestackCenBandwidthPackage())
	if err != nil {
		return fmt.Errorf("error on reading cen bandwidth package %q, %s", d.Id(), err)
	}
	return err
}

func resourceVestackCenBandwidthPackageUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCenBandwidthPackageService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVestackCenBandwidthPackage())
	if err != nil {
		return fmt.Errorf("error on updating cen bandwidth package %q, %s", d.Id(), err)
	}
	return resourceVestackCenBandwidthPackageRead(d, meta)
}

func resourceVestackCenBandwidthPackageDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCenBandwidthPackageService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVestackCenBandwidthPackage())
	if err != nil {
		return fmt.Errorf("error on deleting cen bandwidth package %q, %s", d.Id(), err)
	}
	return err
}
