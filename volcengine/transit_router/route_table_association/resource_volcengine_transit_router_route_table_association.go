package route_table_association

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
TransitRouterRouteTableAssociation can be imported using the TransitRouterAttachmentId:TransitRouterRouteTableId, e.g.
```
$ terraform import volcengine_transit_router_route_table_association.default tr-attach-13n2l4c****:tr-rt-1i5i8khf9m58gae5kcx6****
```

*/

func ResourceVolcengineTransitRouterRouteTableAssociation() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineTransitRouterRouteTableAssociationCreate,
		Read:   resourceVolcengineTransitRouterRouteTableAssociationRead,
		Delete: resourceVolcengineTransitRouterRouteTableAssociationDelete,
		Importer: &schema.ResourceImporter{
			State: routeTableAssociationImporter,
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

func resourceVolcengineTransitRouterRouteTableAssociationCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTRRouteTableAssociationService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineTransitRouterRouteTableAssociation())
	if err != nil {
		return fmt.Errorf("error on creating TransitRouterRouteTableAssociation service %q, %w", d.Id(), err)
	}
	return resourceVolcengineTransitRouterRouteTableAssociationRead(d, meta)
}

func resourceVolcengineTransitRouterRouteTableAssociationRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTRRouteTableAssociationService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineTransitRouterRouteTableAssociation())
	if err != nil {
		return fmt.Errorf("error on reading TransitRouterRouteTableAssociation service %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineTransitRouterRouteTableAssociationDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTRRouteTableAssociationService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineTransitRouterRouteTableAssociation())
	if err != nil {
		return fmt.Errorf("error on deleting TransitRouterRouteTableAssociation service %q, %w", d.Id(), err)
	}
	return err
}

var routeTableAssociationImporter = func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(data.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
	}
	if err := data.Set("transit_router_attachment_id", items[0]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	if err := data.Set("transit_router_route_table_id", items[1]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}
