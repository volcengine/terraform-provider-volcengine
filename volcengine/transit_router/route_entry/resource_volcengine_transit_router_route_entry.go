package route_entry

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
transit router route entry can be imported using the table and entry id, e.g.
```
$ terraform import volcengine_transit_router_route_entry.default tr-rtb-12b7qd3fmzf2817q7y2jkbd55:tr-rte-1i5i8khf9m58gae5kcx6***
```

*/

func ResourceVolcengineTransitRouterRouteEntry() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineTransitRouterRouteEntryCreate,
		Read:   resourceVolcengineTransitRouterRouteEntryRead,
		Update: resourceVolcengineTransitRouterRouteEntryUpdate,
		Delete: resourceVolcengineTransitRouterRouteEntryDelete,
		Importer: &schema.ResourceImporter{
			State: routeEntryImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"transit_router_route_table_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the route table.",
			},
			"destination_cidr_block": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The target network segment of the route entry.",
			},
			"transit_router_route_entry_next_hop_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"Attachment",
					"BlackHole",
				}, false),
				Description: "The next hop type of the routing entry. The value can be Attachment or BlackHole.",
			},
			"transit_router_route_entry_next_hop_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The next hot id of the routing entry. When the parameter TransitRouterRouteEntryNextHopType is Attachment, this parameter must be filled.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Description of the transit router route entry.",
			},
			"transit_router_route_entry_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of the route entry.",
			},
			"transit_router_route_entry_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The id of the route entry.",
			},
		},
	}
	dataSource := DataSourceVolcengineTransitRouterRouteEntries().Schema["entries"].Elem.(*schema.Resource).Schema
	ve.MergeDateSourceToResource(dataSource, &resource.Schema)
	return resource
}

func resourceVolcengineTransitRouterRouteEntryCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTREntryService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineTransitRouterRouteEntry())
	if err != nil {
		return fmt.Errorf("error on creating TransitRouterRouteEntry service %q, %w", d.Id(), err)
	}
	return resourceVolcengineTransitRouterRouteEntryRead(d, meta)
}

func resourceVolcengineTransitRouterRouteEntryRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTREntryService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineTransitRouterRouteEntry())
	if err != nil {
		return fmt.Errorf("error on reading TransitRouterRouteEntry service %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineTransitRouterRouteEntryUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTREntryService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineTransitRouterRouteEntry())
	if err != nil {
		return fmt.Errorf("error on updating TransitRouterRouteEntry service %q, %w", d.Id(), err)
	}
	return resourceVolcengineTransitRouterRouteEntryRead(d, meta)
}

func resourceVolcengineTransitRouterRouteEntryDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTREntryService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineTransitRouterRouteEntry())
	if err != nil {
		return fmt.Errorf("error on deleting TransitRouterRouteEntry service %q, %w", d.Id(), err)
	}
	return err
}

var routeEntryImporter = func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(data.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
	}
	if err := data.Set("transit_router_route_table_id", items[0]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	if err := data.Set("transit_router_route_entry_id", items[1]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}
