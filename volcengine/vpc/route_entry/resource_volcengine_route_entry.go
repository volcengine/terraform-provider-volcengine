package route_entry

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Route entry can be imported using the route_table_id:route_entry_id, e.g.
```
$ terraform import volcengine_route_entry.default vtb-274e19skkuhog7fap8u4i8ird:rte-274e1g9ei4k5c7fap8sp974fq
```

*/

func ResourceVolcengineRouteEntry() *schema.Resource {
	return &schema.Resource{
		Delete: resourceVolcengineRouteEntryDelete,
		Create: resourceVolcengineRouteEntryCreate,
		Read:   resourceVolcengineRouteEntryRead,
		Update: resourceVolcengineRouteEntryUpdate,
		Importer: &schema.ResourceImporter{
			State: importRouteEntry,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
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
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The type of the next hop, Optional choice contains `Instance`, `HaVip`, `NetworkInterface`, `NatGW`, `VpnGW`, `TransitRouter`.",
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

func resourceVolcengineRouteEntryCreate(d *schema.ResourceData, meta interface{}) error {
	routeEntryService := NewRouteEntryService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Create(routeEntryService, d, ResourceVolcengineRouteEntry()); err != nil {
		return fmt.Errorf("error on creating route entry  %q, %w", d.Id(), err)
	}
	return resourceVolcengineRouteEntryRead(d, meta)
}

func resourceVolcengineRouteEntryRead(d *schema.ResourceData, meta interface{}) error {
	routeEntryService := NewRouteEntryService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Read(routeEntryService, d, ResourceVolcengineRouteEntry()); err != nil {
		return fmt.Errorf("error on reading route entry %q, %w", d.Id(), err)
	}
	return nil
}

func resourceVolcengineRouteEntryUpdate(d *schema.ResourceData, meta interface{}) error {
	routeEntryService := NewRouteEntryService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Update(routeEntryService, d, ResourceVolcengineRouteEntry()); err != nil {
		return fmt.Errorf("error on updating route entry %q, %w", d.Id(), err)
	}
	return resourceVolcengineRouteEntryRead(d, meta)
}

func resourceVolcengineRouteEntryDelete(d *schema.ResourceData, meta interface{}) error {
	routeEntryService := NewRouteEntryService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Delete(routeEntryService, d, ResourceVolcengineRouteEntry()); err != nil {
		return fmt.Errorf("error on deleting route entry %q, %w", d.Id(), err)
	}
	return nil
}
