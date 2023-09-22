package ssl_vpn_client_cert

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineSslVpnClientCerts() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineSslVpnClientCertsRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The ids list of ssl vpn client cert.",
			},
			"ssl_vpn_server_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the ssl vpn server.",
			},
			"ssl_vpn_client_cert_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the ssl vpn client cert.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of ssl vpn client cert.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of ssl vpn client cert query.",
			},
			"ssl_vpn_client_certs": {
				Type:        schema.TypeList,
				Description: "The collection of of ssl vpn client certs.",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the ssl vpn client cert.",
						},
						"ssl_vpn_client_cert_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the ssl vpn client cert.",
						},
						"ssl_vpn_client_cert_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the ssl vpn client cert.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the ssl vpn client cert.",
						},
						"ssl_vpn_server_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the ssl vpn server.",
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
				},
			},
		},
	}
}

func dataSourceVolcengineSslVpnClientCertsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewSslVpnClientCertService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineSslVpnClientCerts())
}
