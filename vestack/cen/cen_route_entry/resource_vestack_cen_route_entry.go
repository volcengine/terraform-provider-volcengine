package cen_route_entry

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-vestack/common"
)

/*

Import
CenRouteEntry can be imported using the CenId:DestinationCidrBlock:InstanceId:InstanceType:InstanceRegionId, e.g.
```
$ terraform import vestack_cen_route_entry.default cen-2nim00ybaylts7trquyzt****:100.XX.XX.0/24:vpc-vtbnbb04qw3k2hgi12cv****:VPC:cn-beijing
```

*/

func ResourceVestackCenRouteEntry() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVestackCenRouteEntryCreate,
		Read:   resourceVestackCenRouteEntryRead,
		Delete: resourceVestackCenRouteEntryDelete,
		Importer: &schema.ResourceImporter{
			State: cenRouteEntryImporter,
		},
		Schema: map[string]*schema.Schema{
			"cen_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The cen ID of the cen route entry.",
			},
			"instance_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Default:      "VPC",
				ValidateFunc: validation.StringInSlice([]string{"VPC"}, false),
				Description:  "The instance type of the next hop of the cen route entry.",
			},
			"instance_region_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The instance region id of the next hop of the cen route entry.",
			},
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The instance id of the next hop of the cen route entry.",
			},
			"destination_cidr_block": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The destination cidr block of the cen route entry.",
			},
		},
	}
	s := DataSourceVestackCenRouteEntries().Schema["cen_route_entries"].Elem.(*schema.Resource).Schema
	ve.MergeDateSourceToResource(s, &resource.Schema)
	return resource
}

func resourceVestackCenRouteEntryCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCenRouteEntryService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVestackCenRouteEntry())
	if err != nil {
		return fmt.Errorf("error on creating cen route entry %q, %s", d.Id(), err)
	}
	return resourceVestackCenRouteEntryRead(d, meta)
}

func resourceVestackCenRouteEntryRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCenRouteEntryService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVestackCenRouteEntry())
	if err != nil {
		return fmt.Errorf("error on reading cen route entry %q, %s", d.Id(), err)
	}
	return err
}

func resourceVestackCenRouteEntryDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCenRouteEntryService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVestackCenRouteEntry())
	if err != nil {
		return fmt.Errorf("error on deleting cen route entry %q, %s", d.Id(), err)
	}
	return err
}
