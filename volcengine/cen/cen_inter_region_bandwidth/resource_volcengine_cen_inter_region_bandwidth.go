package cen_inter_region_bandwidth

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
CenInterRegionBandwidth can be imported using the id, e.g.
```
$ terraform import volcengine_cen_inter_region_bandwidth.default cirb-3tex2x1cwd4c6c0v****
```

*/

func ResourceVolcengineCenInterRegionBandwidth() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineCenInterRegionBandwidthCreate,
		Read:   resourceVolcengineCenInterRegionBandwidthRead,
		Update: resourceVolcengineCenInterRegionBandwidthUpdate,
		Delete: resourceVolcengineCenInterRegionBandwidthDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"cen_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The cen ID of the cen inter region bandwidth.",
			},
			"local_region_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The local region id of the cen inter region bandwidth.",
			},
			"peer_region_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The peer region id of the cen inter region bandwidth.",
			},
			"bandwidth": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntAtLeast(1),
				Description:  "The bandwidth of the cen inter region bandwidth.",
			},
			"cen_bandwidth_package_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The cen bandwidth package id of the cen inter region bandwidth.",
			},
		},
	}
	s := DataSourceVolcengineCenInterRegionBandwidths().Schema["inter_region_bandwidths"].Elem.(*schema.Resource).Schema
	delete(s, "id")
	ve.MergeDateSourceToResource(s, &resource.Schema)
	return resource
}

func resourceVolcengineCenInterRegionBandwidthCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCenInterRegionBandwidthService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(service, d, ResourceVolcengineCenInterRegionBandwidth())
	if err != nil {
		return fmt.Errorf("error on creating cen inter region bandwidth %q, %s", d.Id(), err)
	}
	return resourceVolcengineCenInterRegionBandwidthRead(d, meta)
}

func resourceVolcengineCenInterRegionBandwidthRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCenInterRegionBandwidthService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(service, d, ResourceVolcengineCenInterRegionBandwidth())
	if err != nil {
		return fmt.Errorf("error on reading cen inter region bandwidth %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineCenInterRegionBandwidthUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCenInterRegionBandwidthService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(service, d, ResourceVolcengineCenInterRegionBandwidth())
	if err != nil {
		return fmt.Errorf("error on updating cen inter region bandwidth %q, %s", d.Id(), err)
	}
	return resourceVolcengineCenInterRegionBandwidthRead(d, meta)
}

func resourceVolcengineCenInterRegionBandwidthDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCenInterRegionBandwidthService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineCenInterRegionBandwidth())
	if err != nil {
		return fmt.Errorf("error on deleting cen inter region bandwidth %q, %s", d.Id(), err)
	}
	return err
}
