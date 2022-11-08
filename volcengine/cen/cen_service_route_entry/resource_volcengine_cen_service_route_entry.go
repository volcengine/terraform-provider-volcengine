package cen_service_route_entry

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
CenServiceRouteEntry can be imported using the CenId:DestinationCidrBlock:ServiceRegionId:ServiceVpcId, e.g.
```
$ terraform import volcengine_cen_service_route_entry.default cen-2nim00ybaylts7trquyzt****:100.XX.XX.0/24:cn-beijing:vpc-3rlkeggyn6tc010exd32q****
```

*/

func ResourceVolcengineCenServiceRouteEntry() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineCenServiceRouteEntryCreate,
		Read:   resourceVolcengineCenServiceRouteEntryRead,
		Delete: resourceVolcengineCenServiceRouteEntryDelete,
		Importer: &schema.ResourceImporter{
			State: cenServiceRouteEntryImporter,
		},
		Schema: map[string]*schema.Schema{
			"cen_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The cen ID of the cen service route entry.",
			},
			"destination_cidr_block": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsCIDR,
				Description:  "The destination cidr block of the cen service route entry.",
			},
			"service_region_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The service region id of the cen service route entry.",
			},
			"service_vpc_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The service VPC id of the cen service route entry.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The description of the cen service route entry.",
			},
		},
	}
	s := DataSourceVolcengineCenServiceRouteEntries().Schema["service_route_entries"].Elem.(*schema.Resource).Schema
	ve.MergeDateSourceToResource(s, &resource.Schema)
	return resource
}

func resourceVolcengineCenServiceRouteEntryCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCenServiceRouteEntryService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineCenServiceRouteEntry())
	if err != nil {
		return fmt.Errorf("error on creating cen service route entry %q, %s", d.Id(), err)
	}
	return resourceVolcengineCenServiceRouteEntryRead(d, meta)
}

func resourceVolcengineCenServiceRouteEntryRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCenServiceRouteEntryService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineCenServiceRouteEntry())
	if err != nil {
		return fmt.Errorf("error on reading cen service route entry %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineCenServiceRouteEntryDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCenServiceRouteEntryService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineCenServiceRouteEntry())
	if err != nil {
		return fmt.Errorf("error on deleting cen service route entry %q, %s", d.Id(), err)
	}
	return err
}
