package vpn_gateway_route

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
VpnGatewayRoute can be imported using the id, e.g.
```
$ terraform import volcengine_vpn_gateway_route.default vgr-3tex2c6c0v844c****
```

*/

func ResourceVolcengineVpnGatewayRoute() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVpnGatewayRouteCreate,
		Read:   resourceVolcengineVpnGatewayRouteRead,
		Delete: resourceVolcengineVpnGatewayRouteDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
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
	dataSource := DataSourceVolcengineVpnGatewayRoutes().Schema["vpn_gateway_routes"].Elem.(*schema.Resource).Schema
	delete(dataSource, "id")
	ve.MergeDateSourceToResource(dataSource, &resource.Schema)
	return resource
}

func resourceVolcengineVpnGatewayRouteCreate(d *schema.ResourceData, meta interface{}) (err error) {
	routeService := NewVpnGatewayRouteService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(routeService, d, ResourceVolcengineVpnGatewayRoute())
	if err != nil {
		return fmt.Errorf("error on creating Vpn Gateway route %q, %s", d.Id(), err)
	}
	return resourceVolcengineVpnGatewayRouteRead(d, meta)
}

func resourceVolcengineVpnGatewayRouteRead(d *schema.ResourceData, meta interface{}) (err error) {
	routeService := NewVpnGatewayRouteService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(routeService, d, ResourceVolcengineVpnGatewayRoute())
	if err != nil {
		return fmt.Errorf("error on reading Vpn Gateway route %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineVpnGatewayRouteDelete(d *schema.ResourceData, meta interface{}) (err error) {
	routeService := NewVpnGatewayRouteService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(routeService, d, ResourceVolcengineVpnGatewayRoute())
	if err != nil {
		return fmt.Errorf("error on deleting Vpn Gateway route %q, %s", d.Id(), err)
	}
	return err
}
