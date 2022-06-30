package route_table_associate

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Route table associate address can be imported using the route_table_id:subnet_id, e.g.
```
$ terraform import volcengine_route_table_associate.default vtb-2fdzao4h726f45******:subnet-2fdzaou4liw3k5oxruv******
```

*/

func ResourceVolcengineRouteTableAssociate() *schema.Resource {
	return &schema.Resource{
		Delete: resourceVolcengineRouteTableAssociateDelete,
		Create: resourceVolcengineRouteTableAssociateCreate,
		Read:   resourceVolcengineRouteTableAssociateRead,
		Importer: &schema.ResourceImporter{
			State: routeTableAssociateImporter,
		},
		Schema: map[string]*schema.Schema{
			"route_table_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the route table.",
			},
			"subnet_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the subnet.",
			},
		},
	}
}

func resourceVolcengineRouteTableAssociateCreate(d *schema.ResourceData, meta interface{}) error {
	routeTableAssociateService := NewRouteTableAssociateService(meta.(*ve.SdkClient))
	if err := routeTableAssociateService.Dispatcher.Create(routeTableAssociateService, d, ResourceVolcengineRouteTableAssociate()); err != nil {
		return fmt.Errorf("error on creating route table associate %q, %w", d.Id(), err)
	}
	return resourceVolcengineRouteTableAssociateRead(d, meta)
}

func resourceVolcengineRouteTableAssociateRead(d *schema.ResourceData, meta interface{}) error {
	routeTableAssociateService := NewRouteTableAssociateService(meta.(*ve.SdkClient))
	if err := routeTableAssociateService.Dispatcher.Read(routeTableAssociateService, d, ResourceVolcengineRouteTableAssociate()); err != nil {
		return fmt.Errorf("error on reading  route table associate %q, %w", d.Id(), err)
	}
	return nil
}

func resourceVolcengineRouteTableAssociateDelete(d *schema.ResourceData, meta interface{}) error {
	routeTableAssociateService := NewRouteTableAssociateService(meta.(*ve.SdkClient))
	if err := routeTableAssociateService.Dispatcher.Delete(routeTableAssociateService, d, ResourceVolcengineRouteTableAssociate()); err != nil {
		return fmt.Errorf("error on deleting  route table associate %q, %w", d.Id(), err)
	}
	return nil
}
