package cen_bandwidth_package_associate

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Cen bandwidth package associate can be imported using the CenBandwidthPackageId:CenId, e.g.
```
$ terraform import volcengine_cen_bandwidth_package_associate.default cbp-4c2zaavbvh5fx****:cen-7qthudw0ll6jmc****
```

*/

func ResourceVolcengineCenBandwidthPackageAssociate() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineCenBandwidthPackageAssociateCreate,
		Read:   resourceVolcengineCenBandwidthPackageAssociateRead,
		Delete: resourceVolcengineCenBandwidthPackageAssociateDelete,
		Importer: &schema.ResourceImporter{
			State: cenGrantInstanceImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"cen_bandwidth_package_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the cen bandwidth package.",
			},
			"cen_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the cen.",
			},
		},
	}
	return resource
}

func resourceVolcengineCenBandwidthPackageAssociateCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCenBandwidthPackageAssociateService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(service, d, ResourceVolcengineCenBandwidthPackageAssociate())
	if err != nil {
		return fmt.Errorf("error on creating cen bandwidth package associate %q, %s", d.Id(), err)
	}
	return resourceVolcengineCenBandwidthPackageAssociateRead(d, meta)
}

func resourceVolcengineCenBandwidthPackageAssociateRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCenBandwidthPackageAssociateService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(service, d, ResourceVolcengineCenBandwidthPackageAssociate())
	if err != nil {
		return fmt.Errorf("error on reading cen bandwidth package associate %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineCenBandwidthPackageAssociateDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCenBandwidthPackageAssociateService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineCenBandwidthPackageAssociate())
	if err != nil {
		return fmt.Errorf("error on deleting cen bandwidth package associate %q, %s", d.Id(), err)
	}
	return err
}