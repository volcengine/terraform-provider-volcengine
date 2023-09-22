package ssl_vpn_client_cert

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
SSL VPN Client Cert can be imported using the id, e.g.
```
$ terraform import volcengine_ssl_vpn_client_cert.default vsc-2d6b7gjrzc2yo58ozfcx2****
```

*/

func ResourceVolcengineSslClientCertServer() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineSslVpnClientCertCreate,
		Read:   resourceVolcengineSslVpnClientCertRead,
		Update: resourceVolcengineSslVpnClientCertUpdate,
		Delete: resourceVolcengineSslVpnClientCertDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"ssl_vpn_server_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the ssl vpn server.",
			},
			"ssl_vpn_client_cert_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of the ssl vpn client cert.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The description of the ssl vpn client cert.",
			},

			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the ssl vpn client.",
			},
			"certificate_status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the ssl vpn client cert.",
			},
			"creation_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The creation time of the ssl vpn client cert.",
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The update time of the ssl vpn client cert.",
			},
			"expired_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The expired time of the ssl vpn client cert.",
			},
			"ca_certificate": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CA certificate.",
			},
			"client_certificate": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The client certificate.",
			},
			"client_key": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The key of the ssl vpn client.",
			},
			"open_vpn_client_config": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The config of the open vpn client.",
			},
		},
	}
	return resource
}

func resourceVolcengineSslVpnClientCertCreate(d *schema.ResourceData, meta interface{}) (err error) {
	SslVpnClientCertService := NewSslVpnClientCertService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(SslVpnClientCertService, d, ResourceVolcengineSslClientCertServer())
	if err != nil {
		return fmt.Errorf("error on creating SSL Vpn Client Cert %q, %s", d.Id(), err)
	}
	return resourceVolcengineSslVpnClientCertRead(d, meta)
}

func resourceVolcengineSslVpnClientCertRead(d *schema.ResourceData, meta interface{}) (err error) {
	SslVpnClientCertService := NewSslVpnClientCertService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(SslVpnClientCertService, d, ResourceVolcengineSslClientCertServer())
	if err != nil {
		return fmt.Errorf("error on reading SSL Vpn Client Cert %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineSslVpnClientCertUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	SslVpnClientCertService := NewSslVpnClientCertService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(SslVpnClientCertService, d, ResourceVolcengineSslClientCertServer())
	if err != nil {
		return fmt.Errorf("error on updating SSL Vpn Client Cert %q, %s", d.Id(), err)
	}
	return resourceVolcengineSslVpnClientCertRead(d, meta)
}

func resourceVolcengineSslVpnClientCertDelete(d *schema.ResourceData, meta interface{}) (err error) {
	SslVpnClientCertService := NewSslVpnClientCertService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(SslVpnClientCertService, d, ResourceVolcengineSslClientCertServer())
	if err != nil {
		return fmt.Errorf("error on deleting SSL Vpn Client Cert %q, %s", d.Id(), err)
	}
	return err
}
