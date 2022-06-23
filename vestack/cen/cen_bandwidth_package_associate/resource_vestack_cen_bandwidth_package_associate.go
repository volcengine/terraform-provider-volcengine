package cen_bandwidth_package_associate

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-vestack/common"
)

/*

Import
Cen bandwidth package associate can be imported using the CenBandwidthPackageId:CenId, e.g.
```
$ terraform import vestack_cen_bandwidth_package_associate.default cbp-4c2zaavbvh5fx****:cen-7qthudw0ll6jmc****
```

*/

func ResourceVestackCenBandwidthPackageAssociate() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVestackCenBandwidthPackageAssociateCreate,
		Read:   resourceVestackCenBandwidthPackageAssociateRead,
		Delete: resourceVestackCenBandwidthPackageAssociateDelete,
		Importer: &schema.ResourceImporter{
			State: cenGrantInstanceImporter,
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

func resourceVestackCenBandwidthPackageAssociateCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCenBandwidthPackageAssociateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVestackCenBandwidthPackageAssociate())
	if err != nil {
		return fmt.Errorf("error on creating cen bandwidth package associate %q, %s", d.Id(), err)
	}
	return resourceVestackCenBandwidthPackageAssociateRead(d, meta)
}

func resourceVestackCenBandwidthPackageAssociateRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCenBandwidthPackageAssociateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVestackCenBandwidthPackageAssociate())
	if err != nil {
		return fmt.Errorf("error on reading cen bandwidth package associate %q, %s", d.Id(), err)
	}
	return err
}

func resourceVestackCenBandwidthPackageAssociateDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCenBandwidthPackageAssociateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVestackCenBandwidthPackageAssociate())
	if err != nil {
		return fmt.Errorf("error on deleting cen bandwidth package associate %q, %s", d.Id(), err)
	}
	return err
}
