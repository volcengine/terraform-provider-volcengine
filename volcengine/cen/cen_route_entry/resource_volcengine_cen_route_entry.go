package cen_route_entry

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
CenRouteEntry can be imported using the CenId:DestinationCidrBlock:InstanceId:InstanceType:InstanceRegionId, e.g.
```
$ terraform import volcengine_cen_route_entry.default cen-2nim00ybaylts7trquyzt****:100.XX.XX.0/24:vpc-vtbnbb04qw3k2hgi12cv****:VPC:cn-beijing
```

*/

func ResourceVolcengineCenRouteEntry() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineCenRouteEntryCreate,
		Read:   resourceVolcengineCenRouteEntryRead,
		Delete: resourceVolcengineCenRouteEntryDelete,
		Importer: &schema.ResourceImporter{
			State: cenRouteEntryImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
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
	s := DataSourceVolcengineCenRouteEntries().Schema["cen_route_entries"].Elem.(*schema.Resource).Schema
	ve.MergeDateSourceToResource(s, &resource.Schema)
	return resource
}

func resourceVolcengineCenRouteEntryCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCenRouteEntryService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineCenRouteEntry())
	if err != nil {
		return fmt.Errorf("error on creating cen route entry %q, %s", d.Id(), err)
	}
	return resourceVolcengineCenRouteEntryRead(d, meta)
}

func resourceVolcengineCenRouteEntryRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCenRouteEntryService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineCenRouteEntry())
	if err != nil {
		return fmt.Errorf("error on reading cen route entry %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineCenRouteEntryDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCenRouteEntryService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineCenRouteEntry())
	if err != nil {
		return fmt.Errorf("error on deleting cen route entry %q, %s", d.Id(), err)
	}
	return err
}
