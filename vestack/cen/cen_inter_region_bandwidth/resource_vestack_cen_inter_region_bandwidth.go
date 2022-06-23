package cen_inter_region_bandwidth

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-vestack/common"
)

/*

Import
CenInterRegionBandwidth can be imported using the id, e.g.
```
$ terraform import vestack_cen_inter_region_bandwidth.default cirb-3tex2x1cwd4c6c0v****
```

*/

func ResourceVestackCenInterRegionBandwidth() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVestackCenInterRegionBandwidthCreate,
		Read:   resourceVestackCenInterRegionBandwidthRead,
		Update: resourceVestackCenInterRegionBandwidthUpdate,
		Delete: resourceVestackCenInterRegionBandwidthDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
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
		},
	}
	s := DataSourceVestackCenInterRegionBandwidths().Schema["inter_region_bandwidths"].Elem.(*schema.Resource).Schema
	delete(s, "id")
	ve.MergeDateSourceToResource(s, &resource.Schema)
	return resource
}

func resourceVestackCenInterRegionBandwidthCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCenInterRegionBandwidthService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVestackCenInterRegionBandwidth())
	if err != nil {
		return fmt.Errorf("error on creating cen inter region bandwidth %q, %s", d.Id(), err)
	}
	return resourceVestackCenInterRegionBandwidthRead(d, meta)
}

func resourceVestackCenInterRegionBandwidthRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCenInterRegionBandwidthService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVestackCenInterRegionBandwidth())
	if err != nil {
		return fmt.Errorf("error on reading cen inter region bandwidth %q, %s", d.Id(), err)
	}
	return err
}

func resourceVestackCenInterRegionBandwidthUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCenInterRegionBandwidthService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVestackCenInterRegionBandwidth())
	if err != nil {
		return fmt.Errorf("error on updating cen inter region bandwidth %q, %s", d.Id(), err)
	}
	return resourceVestackCenInterRegionBandwidthRead(d, meta)
}

func resourceVestackCenInterRegionBandwidthDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCenInterRegionBandwidthService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVestackCenInterRegionBandwidth())
	if err != nil {
		return fmt.Errorf("error on deleting cen inter region bandwidth %q, %s", d.Id(), err)
	}
	return err
}
