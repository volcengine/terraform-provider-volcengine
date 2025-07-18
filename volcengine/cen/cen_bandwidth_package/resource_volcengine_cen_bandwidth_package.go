package cen_bandwidth_package

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
CenBandwidthPackage can be imported using the id, e.g.
```
$ terraform import volcengine_cen_bandwidth_package.default cbp-4c2zaavbvh5f42****
```

*/

func ResourceVolcengineCenBandwidthPackage() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineCenBandwidthPackageCreate,
		Read:   resourceVolcengineCenBandwidthPackageRead,
		Update: resourceVolcengineCenBandwidthPackageUpdate,
		Delete: resourceVolcengineCenBandwidthPackageDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"local_geographic_region_set_id": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Default:      "China",
				ValidateFunc: validation.StringInSlice([]string{"China", "Asia"}, false),
				Description:  "The local geographic region set id of the cen bandwidth package. Valid value: `China`, `Asia`.",
			},
			"peer_geographic_region_set_id": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Default:      "China",
				ValidateFunc: validation.StringInSlice([]string{"China", "Asia"}, false),
				Description:  "The peer geographic region set id of the cen bandwidth package. Valid value: `China`, `Asia`.",
			},
			"line_operator": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  "ChinaUnicom",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Get("local_geographic_region_set_id").(string) == d.Get("peer_geographic_region_set_id").(string)
				},
				Description: "The line operator of the cen bandwidth package. Valid value: `ChinaUnicom`, `ChinaTelecom`. This field is only valid when `local_geographic_region_set_id` and `peer_geographic_region_set_id` are different.",
			},
			"bandwidth": {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.IntBetween(2, 100000),
				Description:  "The bandwidth of the cen bandwidth package. Value: 2~10000.",
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
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Default:     "PrePaid",
				Description: "The billing type of the cen bandwidth package. Only support `PrePaid` and `PayBy95Peak`, default value is `PrePaid`.",
			},
			"period_unit": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "Month",
				ValidateFunc:     validation.StringInSlice([]string{"Month", "Year"}, false),
				DiffSuppressFunc: periodDiffSuppress,
				Description:      "The period unit of the cen bandwidth package. Value: `Month`, `Year`. Default value is `Month`.",
			},
			"period": {
				Type:             schema.TypeInt,
				Optional:         true,
				Default:          1,
				DiffSuppressFunc: periodDiffSuppress,
				Description:      "The period of the cen bandwidth package. Default value is 1.",
			},
			"tags": ve.TagsSchema(),
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The ProjectName of the cen bandwidth package.",
			},
		},
	}
	s := DataSourceVolcengineCenBandwidthPackages().Schema["bandwidth_packages"].Elem.(*schema.Resource).Schema
	delete(s, "id")
	ve.MergeDateSourceToResource(s, &resource.Schema)
	return resource
}

func resourceVolcengineCenBandwidthPackageCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCenBandwidthPackageService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(service, d, ResourceVolcengineCenBandwidthPackage())
	if err != nil {
		return fmt.Errorf("error on creating cen bandwidth package %q, %s", d.Id(), err)
	}
	return resourceVolcengineCenBandwidthPackageRead(d, meta)
}

func resourceVolcengineCenBandwidthPackageRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCenBandwidthPackageService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(service, d, ResourceVolcengineCenBandwidthPackage())
	if err != nil {
		return fmt.Errorf("error on reading cen bandwidth package %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineCenBandwidthPackageUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCenBandwidthPackageService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(service, d, ResourceVolcengineCenBandwidthPackage())
	if err != nil {
		return fmt.Errorf("error on updating cen bandwidth package %q, %s", d.Id(), err)
	}
	return resourceVolcengineCenBandwidthPackageRead(d, meta)
}

func resourceVolcengineCenBandwidthPackageDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCenBandwidthPackageService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineCenBandwidthPackage())
	if err != nil {
		return fmt.Errorf("error on deleting cen bandwidth package %q, %s", d.Id(), err)
	}
	return err
}
