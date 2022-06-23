package vpn_gateway

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-vestack/common"
)

/*

Import
VpnGateway can be imported using the id, e.g.
```
$ terraform import vestack_vpn_gateway.default vgw-273zkshb2qayo7fap8t2****
```

*/

func ResourceVestackVpnGateway() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVestackVpnGatewayCreate,
		Read:   resourceVestackVpnGatewayRead,
		Update: resourceVestackVpnGatewayUpdate,
		Delete: resourceVestackVpnGatewayDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"vpc_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the VPC where you want to create the VPN gateway.",
			},
			"subnet_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the subnet where you want to create the VPN gateway.",
			},
			"bandwidth": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The bandwidth of the VPN gateway.",
			},
			"vpn_gateway_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of the VPN gateway.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The description of the VPN gateway.",
			},
			//"billing_type": {
			//	Type:        schema.TypeString,
			//	Optional:    true,
			//	Computed:    true,
			//	Description: "The BillingType of the VPN gateway.",
			//},
			"period_unit": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "Month",
				ValidateFunc: validation.StringInSlice([]string{"Month", "Year"}, false),
				Description:  "The PeriodUnit of the VPN gateway.  This parameter is only useful when creating vpn gateway.",
			},
			"period": {
				Type:        schema.TypeInt,
				Optional:    true,
				Default:     1,
				Description: "The Period of the VPN gateway. This parameter is only useful when creating vpn gateway.",
			},
			"renew_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "ManualRenew",
				ValidateFunc: validation.StringInSlice([]string{"ManualRenew", "AutoRenew", "NoneRenew"}, false),
				Description:  "The renew type of the VPN gateway.",
			},
			"renew_period": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The renew period of the VPN gateway.",
			},
			"remain_renew_times": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The remain renew times of the VPN gateway.",
			},
		},
	}
	dataSource := DataSourceVestackVpnGateways().Schema["vpn_gateways"].Elem.(*schema.Resource).Schema
	delete(dataSource, "id")
	ve.MergeDateSourceToResource(dataSource, &resource.Schema)
	return resource
}

func resourceVestackVpnGatewayCreate(d *schema.ResourceData, meta interface{}) (err error) {
	vpnGatewayService := NewVpnGatewayService(meta.(*ve.SdkClient))
	err = vpnGatewayService.Dispatcher.Create(vpnGatewayService, d, ResourceVestackVpnGateway())
	if err != nil {
		return fmt.Errorf("error on creating Vpn Gateway %q, %s", d.Id(), err)
	}
	return resourceVestackVpnGatewayRead(d, meta)
}

func resourceVestackVpnGatewayRead(d *schema.ResourceData, meta interface{}) (err error) {
	vpnGatewayService := NewVpnGatewayService(meta.(*ve.SdkClient))
	err = vpnGatewayService.Dispatcher.Read(vpnGatewayService, d, ResourceVestackVpnGateway())
	if err != nil {
		return fmt.Errorf("error on reading Vpn Gateway %q, %s", d.Id(), err)
	}
	return err
}

func resourceVestackVpnGatewayUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	vpnGatewayService := NewVpnGatewayService(meta.(*ve.SdkClient))
	err = vpnGatewayService.Dispatcher.Update(vpnGatewayService, d, ResourceVestackVpnGateway())
	if err != nil {
		return fmt.Errorf("error on updating Vpn Gateway %q, %s", d.Id(), err)
	}
	return resourceVestackVpnGatewayRead(d, meta)
}

func resourceVestackVpnGatewayDelete(d *schema.ResourceData, meta interface{}) (err error) {
	vpnGatewayService := NewVpnGatewayService(meta.(*ve.SdkClient))
	err = vpnGatewayService.Dispatcher.Delete(vpnGatewayService, d, ResourceVestackVpnGateway())
	if err != nil {
		return fmt.Errorf("error on deleting Vpn Gateway %q, %s", d.Id(), err)
	}
	return err
}
