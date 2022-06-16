package vpn_connection

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-vestack/common"
)

func DataSourceVestackVpnConnections() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVestackVpnConnectionsRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of VPN connection ids.",
			},
			"vpn_connection_names": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of VPN connection names.",
			},
			"vpn_gateway_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "An ID of VPN gateway.",
			},
			"customer_gateway_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "An ID of customer gateway.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of VPN connection.",
			},

			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},

			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of VPN connection query.",
			},
			"vpn_connections": {
				Description: "The collection of VPN connection query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the VPN connection.",
						},
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The account ID of the VPN connection.",
						},
						"vpn_connection_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the VPN connection.",
						},
						"creation_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of VPN connection.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of VPN connection.",
						},
						"vpn_connection_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the VPN connection.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the VPN connection.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the VPN connection.",
						},
						"vpn_gateway_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the vpn gateway.",
						},
						"customer_gateway_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the customer gateway.",
						},
						"local_subnet": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The local subnet of the VPN connection.",
						},
						"remote_subnet": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The remote subnet of the VPN connection.",
						},
						"dpd_action": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The dpd action of the VPN connection.",
						},
						"connect_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The connect status of the VPN connection.",
						},
						"nat_traversal": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "The nat traversal of the VPN connection.",
						},

						// ike config
						"ike_config_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The version of the ike config of the VPN connection.",
						},
						"ike_config_psk": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The psk of the ike config of the VPN connection.",
						},
						"ike_config_dh_group": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The dk group of the ike config of the VPN connection.",
						},
						"ike_config_mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The mode of the ike config of the VPN connection.",
						},
						"ike_config_enc_alg": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The enc alg of the ike config of the VPN connection.",
						},
						"ike_config_auth_alg": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The auth alg of the ike config of the VPN connection.",
						},
						"ike_config_local_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The local_id of the ike config of the VPN connection.",
						},
						"ike_config_remote_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The remote id of the ike config of the VPN connection.",
						},
						"ike_config_lifetime": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The lifetime of the ike config of the VPN connection.",
						},

						// ipsec config
						"ipsec_config_auth_alg": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The auth alg of the ipsec config of the VPN connection.",
						},
						"ipsec_config_enc_alg": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The enc alg of the ipsec config of the VPN connection.",
						},
						"ipsec_config_dh_group": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The dh group of the ipsec config of the VPN connection.",
						},
						"ipsec_config_lifetime": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The lifetime of the ike config of the VPN connection.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVestackVpnConnectionsRead(d *schema.ResourceData, meta interface{}) error {
	vpnConnectionService := NewVpnConnectionService(meta.(*ve.SdkClient))
	return vpnConnectionService.Dispatcher.Data(vpnConnectionService, d, DataSourceVestackVpnConnections())
}
