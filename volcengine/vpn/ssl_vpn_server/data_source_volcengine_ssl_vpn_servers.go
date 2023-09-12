package ssl_vpn_server

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineSslVpnServers() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineSslVpnServersRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The ids list.",
			},
			"vpn_gateway_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the vpn gateway.",
			},
			"ssl_vpn_server_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the ssl vpn server.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of SSL VPN server query.",
			},
			"ssl_vpn_servers": {
				Type:        schema.TypeList,
				Description: "List of SSL VPN servers.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The SSL VPN server id.",
						},
						"vpn_gateway_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The vpn gateway id.",
						},
						"local_subnets": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The local network segment of the SSL server. The local network segment is the address segment that the client accesses through the SSL VPN connection.",
						},
						"client_ip_pool": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "SSL client network segment.",
						},
						"ssl_vpn_server_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the SSL server.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the ssl server.",
						},
						"protocol": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The protocol used by the SSL server. Valid values are `TCP`, `UDP`. Default Value: `UDP`.",
						},
						"cipher": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The encryption algorithm of the SSL server.\nValues:\n`AES-128-CBC` (default)\n`AES-192-CBC`\n`AES-256-CBC`\n`None` (do not use encryption).",
						},
						"auth": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The authentication algorithm of the SSL server.\nValues:\n`SHA1` (default)\n`MD5`\n`None` (do not use encryption).",
						},
						"compress": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to compress the transmitted data. The default value is false.",
						},
						"ssl_vpn_server_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the ssl vpn server.",
						},
						"creation_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The creation time.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the ssl vpn server.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineSslVpnServersRead(d *schema.ResourceData, meta interface{}) error {
	service := NewSslVpnServerService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineSslVpnServers())
}
