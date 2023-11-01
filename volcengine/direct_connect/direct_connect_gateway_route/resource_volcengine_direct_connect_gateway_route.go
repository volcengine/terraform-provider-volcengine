package direct_connect_gateway_route

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
DirectConnectGatewayRoute can be imported using the id, e.g.
```
$ terraform import volcengine_direct_connect_gateway_route.default resource_id
```

*/

func ResourceVolcengineDirectConnectGatewayRoute() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineDirectConnectGatewayRouteCreate,
		Read:   resourceVolcengineDirectConnectGatewayRouteRead,
		Delete: resourceVolcengineDirectConnectGatewayRouteDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"direct_connect_gateway_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of direct connect gateway.",
			},
			"destination_cidr_block": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The cidr block.",
			},
			"next_hop_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of next hop.",
			},
		},
	}
	dataSource := DataSourceVolcengineDirectConnectGatewayRoutes().Schema["direct_connect_gateway_routes"].Elem.(*schema.Resource).Schema
	ve.MergeDateSourceToResource(dataSource, &resource.Schema)
	return resource
}

func resourceVolcengineDirectConnectGatewayRouteCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewDirectConnectGatewayRouteService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineDirectConnectGatewayRoute())
	if err != nil {
		return fmt.Errorf("error on creating direct_connect_gateway_route %q, %s", d.Id(), err)
	}
	return resourceVolcengineDirectConnectGatewayRouteRead(d, meta)
}

func resourceVolcengineDirectConnectGatewayRouteRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewDirectConnectGatewayRouteService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineDirectConnectGatewayRoute())
	if err != nil {
		return fmt.Errorf("error on reading direct_connect_gateway_route %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineDirectConnectGatewayRouteDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewDirectConnectGatewayRouteService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineDirectConnectGatewayRoute())
	if err != nil {
		return fmt.Errorf("error on deleting direct_connect_gateway_route %q, %s", d.Id(), err)
	}
	return err
}
