package ssl_vpn_server

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
SSL VPN server can be imported using the id, e.g.
```
$ terraform import volcengine_ssl_vpn_server.default vss-zm55pqtvk17oq32zd****
```

*/

func ResourceVolcengineSslVpnServer() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineSslVpnServerCreate,
		Read:   resourceVolcengineSslVpnServerRead,
		Update: resourceVolcengineSslVpnServerUpdate,
		Delete: resourceVolcengineSslVpnServerDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"vpn_gateway_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The vpn gateway id.",
			},
			"local_subnets": {
				Type:     schema.TypeSet,
				Required: true,
				ForceNew: true,
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The local network segment of the SSL server. The local network segment is the address segment that the client accesses through the SSL VPN connection.",
			},
			"client_ip_pool": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "SSL client network segment.",
			},
			"ssl_vpn_server_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of the SSL server.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The description of the ssl server.",
			},
			"protocol": {
				Type:         schema.TypeString,
				Default:      "UDP",
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"UDP", "TCP"}, false),
				Description:  "The protocol used by the SSL server. Valid values are `TCP`, `UDP`. Default Value: `UDP`.",
			},
			"cipher": {
				Type:         schema.TypeString,
				Default:      "AES-128-CBC",
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"AES-128-CBC", "AES-192-CBC", "AES-256-CBC", "None"}, false),
				Description:  "The encryption algorithm of the SSL server.\nValues:\n`AES-128-CBC` (default)\n`AES-192-CBC`\n`AES-256-CBC`\n`None` (do not use encryption).",
			},
			"auth": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				Default:      "SHA1",
				ValidateFunc: validation.StringInSlice([]string{"SHA1", "MD5", "None"}, false),
				Description:  "The authentication algorithm of the SSL server.\nValues:\n`SHA1` (default)\n`MD5`\n`None` (do not use encryption).",
			},
			"compress": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Default:     false,
				Description: "Whether to compress the transmitted data. The default value is false.",
			},
			"ssl_vpn_server_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The id of the ssl vpn server.",
			},
		},
	}
	return resource
}

func resourceVolcengineSslVpnServerCreate(d *schema.ResourceData, meta interface{}) (err error) {
	SslVpnServerService := NewSslVpnServerService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(SslVpnServerService, d, ResourceVolcengineSslVpnServer())
	if err != nil {
		return fmt.Errorf("error on creating SSL Vpn Server %q, %s", d.Id(), err)
	}
	return resourceVolcengineSslVpnServerRead(d, meta)
}

func resourceVolcengineSslVpnServerRead(d *schema.ResourceData, meta interface{}) (err error) {
	SslVpnServerService := NewSslVpnServerService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(SslVpnServerService, d, ResourceVolcengineSslVpnServer())
	if err != nil {
		return fmt.Errorf("error on reading SSL Vpn Server %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineSslVpnServerUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	SslVpnServerService := NewSslVpnServerService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(SslVpnServerService, d, ResourceVolcengineSslVpnServer())
	if err != nil {
		return fmt.Errorf("error on updating SSL Vpn Server %q, %s", d.Id(), err)
	}
	return resourceVolcengineSslVpnServerRead(d, meta)
}

func resourceVolcengineSslVpnServerDelete(d *schema.ResourceData, meta interface{}) (err error) {
	SslVpnServerService := NewSslVpnServerService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(SslVpnServerService, d, ResourceVolcengineSslVpnServer())
	if err != nil {
		return fmt.Errorf("error on deleting SSL Vpn Server %q, %s", d.Id(), err)
	}
	return err
}
