package vpn_connection

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
VpnConnection can be imported using the id, e.g.
```
$ terraform import volcengine_vpn_connection.default vgc-3tex2x1cwd4c6c0v****
```

*/

func ResourceVolcengineVpnConnection() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVpnConnectionCreate,
		Read:   resourceVolcengineVpnConnectionRead,
		Update: resourceVolcengineVpnConnectionUpdate,
		Delete: resourceVolcengineVpnConnectionDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"vpn_connection_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of the VPN connection.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The description of the VPN connection.",
			},
			"vpn_gateway_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the vpn gateway.",
			},
			"customer_gateway_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the customer gateway.",
			},
			"local_subnet": {
				Type:     schema.TypeSet,
				Required: true,
				Set:      schema.HashString,
				MinItems: 1,
				MaxItems: 5,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.IsCIDR,
				},
				Description: "The local subnet of the VPN connection.",
			},
			"remote_subnet": {
				Type:     schema.TypeSet,
				Required: true,
				Set:      schema.HashString,
				MinItems: 1,
				MaxItems: 5,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.IsCIDR,
				},
				Description: "The remote subnet of the VPN connection.",
			},
			"dpd_action": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "clear",
				ValidateFunc: validation.StringInSlice([]string{"clear", "none", "hold", "restart"}, false),
				Description:  "The dpd action of the VPN connection.",
			},
			"nat_traversal": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "The nat traversal of the VPN connection.",
			},

			// ike config
			"ike_config_psk": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The psk of the ike config of the VPN connection.",
			},
			"ike_config_version": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "ikev1",
				ValidateFunc: validation.StringInSlice([]string{"ikev1", "ikev2"}, false),
				Description:  "The version of the ike config of the VPN connection.",
			},
			"ike_config_mode": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "main",
				ValidateFunc: validation.StringInSlice([]string{"main", "aggressive"}, false),
				Description:  "The mode of the ike config of the VPN connection.",
			},
			"ike_config_enc_alg": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "aes",
				ValidateFunc: validation.StringInSlice([]string{"aes", "aes192", "aes256", "des", "3des", "sm4"}, false),
				Description:  "The enc alg of the ike config of the VPN connection.",
			},
			"ike_config_auth_alg": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "sha1",
				ValidateFunc: validation.StringInSlice([]string{"sha1", "md5", "sha256", "sha384", "sha512", "sm3"}, false),
				Description:  "The auth alg of the ike config of the VPN connection.",
			},
			"ike_config_dh_group": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "group2",
				ValidateFunc: validation.StringInSlice([]string{"group1", "group2", "group5", "group14"}, false),
				Description:  "The dk group of the ike config of the VPN connection.",
			},
			"ike_config_lifetime": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      86400,
				ValidateFunc: validation.IntBetween(0, 86400),
				Description:  "The lifetime of the ike config of the VPN connection.",
			},
			"ike_config_local_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The local_id of the ike config of the VPN connection.",
			},
			"ike_config_remote_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The remote id of the ike config of the VPN connection.",
			},

			// ipsec config
			"ipsec_config_enc_alg": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "aes",
				ValidateFunc: validation.StringInSlice([]string{"aes", "aes192", "aes256", "des", "3des", "sm4"}, false),
				Description:  "The enc alg of the ipsec config of the VPN connection.",
			},
			"ipsec_config_auth_alg": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "sha1",
				ValidateFunc: validation.StringInSlice([]string{"sha1", "md5", "sha256", "sha384", "sha512", "sm3"}, false),
				Description:  "The auth alg of the ipsec config of the VPN connection.",
			},
			"ipsec_config_dh_group": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "group2",
				ValidateFunc: validation.StringInSlice([]string{"group1", "group2", "group5", "group14", "disable"}, false),
				Description:  "The dh group of the ipsec config of the VPN connection.",
			},
			"ipsec_config_lifetime": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      86400,
				ValidateFunc: validation.IntBetween(0, 86400),
				Description:  "The ipsec config of the ike config of the VPN connection.",
			},
			"negotiate_instantly": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether to initiate negotiation mode immediately.",
			},
		},
	}
	dataSource := DataSourceVolcengineVpnConnections().Schema["vpn_connections"].Elem.(*schema.Resource).Schema
	delete(dataSource, "id")
	ve.MergeDateSourceToResource(dataSource, &resource.Schema)
	return resource
}

func resourceVolcengineVpnConnectionCreate(d *schema.ResourceData, meta interface{}) (err error) {
	connectionService := NewVpnConnectionService(meta.(*ve.SdkClient))
	err = connectionService.Dispatcher.Create(connectionService, d, ResourceVolcengineVpnConnection())
	if err != nil {
		return fmt.Errorf("error on creating Vpn Connections %q, %s", d.Id(), err)
	}
	return resourceVolcengineVpnConnectionRead(d, meta)
}

func resourceVolcengineVpnConnectionRead(d *schema.ResourceData, meta interface{}) (err error) {
	connectionService := NewVpnConnectionService(meta.(*ve.SdkClient))
	err = connectionService.Dispatcher.Read(connectionService, d, ResourceVolcengineVpnConnection())
	if err != nil {
		return fmt.Errorf("error on reading Vpn Connection %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineVpnConnectionUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	connectionService := NewVpnConnectionService(meta.(*ve.SdkClient))
	err = connectionService.Dispatcher.Update(connectionService, d, ResourceVolcengineVpnConnection())
	if err != nil {
		return fmt.Errorf("error on updating Vpn Connection %q, %s", d.Id(), err)
	}
	return resourceVolcengineVpnConnectionRead(d, meta)
}

func resourceVolcengineVpnConnectionDelete(d *schema.ResourceData, meta interface{}) (err error) {
	connectionService := NewVpnConnectionService(meta.(*ve.SdkClient))
	err = connectionService.Dispatcher.Delete(connectionService, d, ResourceVolcengineVpnConnection())
	if err != nil {
		return fmt.Errorf("error on deleting Vpn Connection %q, %s", d.Id(), err)
	}
	return err
}
