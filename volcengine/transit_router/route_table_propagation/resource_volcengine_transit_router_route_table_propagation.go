package route_table_propagation

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
TransitRouterRouteTablePropagation can be imported using the propagation:TransitRouterAttachmentId:TransitRouterRouteTableId, e.g.
```
$ terraform import volcengine_transit_router_route_table_propagation.default propagation:tr-attach-13n2l4c****:tr-rt-1i5i8khf9m58gae5kcx6****
```

*/

func ResourceVolcengineTransitRouterRouteTablePropagation() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineTransitRouterRouteTablePropagationCreate,
		Read:   resourceVolcengineTransitRouterRouteTablePropagationRead,
		Delete: resourceVolcengineTransitRouterRouteTablePropagationDelete,
		Importer: &schema.ResourceImporter{
			State: routeTablePropagationImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"transit_router_attachment_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the network instance connection.",
			},
			"transit_router_route_table_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the routing table associated with the transit router instance.",
			},
		},
	}
	return resource
}

func resourceVolcengineTransitRouterRouteTablePropagationCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTRRouteTablePropagationService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineTransitRouterRouteTablePropagation())
	if err != nil {
		return fmt.Errorf("error on creating TransitRouterRouteTablePropagation service %q, %w", d.Id(), err)
	}
	return resourceVolcengineTransitRouterRouteTablePropagationRead(d, meta)
}

func resourceVolcengineTransitRouterRouteTablePropagationRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTRRouteTablePropagationService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineTransitRouterRouteTablePropagation())
	if err != nil {
		return fmt.Errorf("error on reading TransitRouterRouteTablePropagation service %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineTransitRouterRouteTablePropagationDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTRRouteTablePropagationService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineTransitRouterRouteTablePropagation())
	if err != nil {
		return fmt.Errorf("error on deleting TransitRouterRouteTablePropagation service %q, %w", d.Id(), err)
	}
	return err
}

var routeTablePropagationImporter = func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(data.Id(), ":")
	if len(items) != 3 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
	}
	if err := data.Set("transit_router_attachment_id", items[1]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	if err := data.Set("transit_router_route_table_id", items[2]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}
