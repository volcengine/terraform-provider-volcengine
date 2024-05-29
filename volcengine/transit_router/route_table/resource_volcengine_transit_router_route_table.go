package route_table

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
transit router route table can be imported using the router id and route table id, e.g.
```
$ terraform import volcengine_transit_router_route_table.default tr-2ff4v69tkxji859gp684cm14e:tr-rtb-hy13n2l4c6c0v****
```

*/

func ResourceVolcengineTransitRouterRouteTable() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineTransitRouterRouteTableCreate,
		Read:   resourceVolcengineTransitRouterRouteTableRead,
		Update: resourceVolcengineTransitRouterRouteTableUpdate,
		Delete: resourceVolcengineTransitRouterRouteTableDelete,
		Importer: &schema.ResourceImporter{
			State: routeTableImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"transit_router_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Id of the transit router.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Description of the transit router route table.",
			},
			"transit_router_route_table_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of the route table.",
			},
			"transit_router_route_table_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The id of the route table.",
			},
			"tags": ve.TagsSchema(),
		},
	}
	dataSource := DataSourceVolcengineTransitRouterRouteTables().Schema["route_tables"].Elem.(*schema.Resource).Schema
	ve.MergeDateSourceToResource(dataSource, &resource.Schema)
	return resource
}

func resourceVolcengineTransitRouterRouteTableCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTRRouteTableService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineTransitRouterRouteTable())
	if err != nil {
		return fmt.Errorf("error on creating TransitRouterRouteTable service %q, %w", d.Id(), err)
	}
	return resourceVolcengineTransitRouterRouteTableRead(d, meta)
}

func resourceVolcengineTransitRouterRouteTableRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTRRouteTableService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineTransitRouterRouteTable())
	if err != nil {
		return fmt.Errorf("error on reading TransitRouterRouteTable service %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineTransitRouterRouteTableUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTRRouteTableService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineTransitRouterRouteTable())
	if err != nil {
		return fmt.Errorf("error on updating TransitRouterRouteTable service %q, %w", d.Id(), err)
	}
	return resourceVolcengineTransitRouterRouteTableRead(d, meta)
}

func resourceVolcengineTransitRouterRouteTableDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTRRouteTableService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineTransitRouterRouteTable())
	if err != nil {
		return fmt.Errorf("error on deleting TransitRouterRouteTable service%q, %w", d.Id(), err)
	}
	return err
}

var routeTableImporter = func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(data.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
	}
	if err := data.Set("transit_router_id", items[0]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	if err := data.Set("transit_router_route_table_id", items[1]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}
