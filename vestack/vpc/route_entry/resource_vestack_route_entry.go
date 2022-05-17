package route_entry

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-vestack/common"
)

/*

Import
Route entry can be imported using the route_table_id:route_entry_id, e.g.
```
$ terraform import vestack_route_entry.default vtb-274e19skkuhog7fap8u4i8ird:rte-274e1g9ei4k5c7fap8sp974fq
```

*/

func ResourceVestackRouteEntry() *schema.Resource {
	return &schema.Resource{
		Delete: resourceVestackRouteEntryDelete,
		Create: resourceVestackRouteEntryCreate,
		Read:   resourceVestackRouteEntryRead,
		Update: resourceVestackRouteEntryUpdate,
		Importer: &schema.ResourceImporter{
			State: importRouteEntry,
		},
		Schema: map[string]*schema.Schema{
			"route_table_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the route table.",
			},
			"route_entry_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The id of the route entry.",
			},
			"destination_cidr_block": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The destination CIDR block of the route entry.",
			},
			"next_hop_type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				Description:  "The type of the next hop.",
				ValidateFunc: validation.StringInSlice([]string{"Instance", "NetworkInterface", "NatGW", "VpnGW"}, false),
			},
			"next_hop_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the next hop.",
			},
			"route_entry_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the route entry.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the route entry.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The description of the route entry.",
			},
		},
	}
}

func resourceVestackRouteEntryCreate(d *schema.ResourceData, meta interface{}) error {
	routeEntryService := NewRouteEntryService(meta.(*ve.SdkClient))
	if err := routeEntryService.Dispatcher.Create(routeEntryService, d, ResourceVestackRouteEntry()); err != nil {
		return fmt.Errorf("error on creating route entry  %q, %w", d.Id(), err)
	}
	return resourceVestackRouteEntryRead(d, meta)
}

func resourceVestackRouteEntryRead(d *schema.ResourceData, meta interface{}) error {
	routeEntryService := NewRouteEntryService(meta.(*ve.SdkClient))
	if err := routeEntryService.Dispatcher.Read(routeEntryService, d, ResourceVestackRouteEntry()); err != nil {
		return fmt.Errorf("error on reading route entry %q, %w", d.Id(), err)
	}
	return nil
}

func resourceVestackRouteEntryUpdate(d *schema.ResourceData, meta interface{}) error {
	routeEntryService := NewRouteEntryService(meta.(*ve.SdkClient))
	if err := routeEntryService.Dispatcher.Update(routeEntryService, d, ResourceVestackRouteEntry()); err != nil {
		return fmt.Errorf("error on updating route entry %q, %w", d.Id(), err)
	}
	return resourceVestackRouteEntryRead(d, meta)
}

func resourceVestackRouteEntryDelete(d *schema.ResourceData, meta interface{}) error {
	routeEntryService := NewRouteEntryService(meta.(*ve.SdkClient))
	if err := routeEntryService.Dispatcher.Delete(routeEntryService, d, ResourceVestackRouteEntry()); err != nil {
		return fmt.Errorf("error on deleting route entry %q, %w", d.Id(), err)
	}
	return nil
}
