package vpn_gateway_route

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-vestack/common"
)

/*

Import
VpnGatewayRoute can be imported using the id, e.g.
```
$ terraform import vestack_vpn_gateway_route.default vgr-3tex2c6c0v844c****
```

*/

func ResourceVestackVpnGatewayRoute() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVestackVpnGatewayRouteCreate,
		Read:   resourceVestackVpnGatewayRouteRead,
		Delete: resourceVestackVpnGatewayRouteDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"vpn_gateway_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the VPN gateway of the VPN gateway route.",
			},
			"destination_cidr_block": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.IsCIDR,
				Description:  "The destination cidr block of the VPN gateway route.",
			},
			"next_hop_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The next hop id of the VPN gateway route.",
			},
		},
	}
	dataSource := DataSourceVestackVpnGatewayRoutes().Schema["vpn_gateway_routes"].Elem.(*schema.Resource).Schema
	delete(dataSource, "id")
	ve.MergeDateSourceToResource(dataSource, &resource.Schema)
	return resource
}

func resourceVestackVpnGatewayRouteCreate(d *schema.ResourceData, meta interface{}) (err error) {
	routeService := NewVpnGatewayRouteService(meta.(*ve.SdkClient))
	err = routeService.Dispatcher.Create(routeService, d, ResourceVestackVpnGatewayRoute())
	if err != nil {
		return fmt.Errorf("error on creating Vpn Gateway route %q, %s", d.Id(), err)
	}
	return resourceVestackVpnGatewayRouteRead(d, meta)
}

func resourceVestackVpnGatewayRouteRead(d *schema.ResourceData, meta interface{}) (err error) {
	routeService := NewVpnGatewayRouteService(meta.(*ve.SdkClient))
	err = routeService.Dispatcher.Read(routeService, d, ResourceVestackVpnGatewayRoute())
	if err != nil {
		return fmt.Errorf("error on reading Vpn Gateway route %q, %s", d.Id(), err)
	}
	return err
}

func resourceVestackVpnGatewayRouteDelete(d *schema.ResourceData, meta interface{}) (err error) {
	routeService := NewVpnGatewayRouteService(meta.(*ve.SdkClient))
	err = routeService.Dispatcher.Delete(routeService, d, ResourceVestackVpnGatewayRoute())
	if err != nil {
		return fmt.Errorf("error on deleting Vpn Gateway route %q, %s", d.Id(), err)
	}
	return err
}
