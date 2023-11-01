package direct_connect_bgp_peer

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineDirectConnectBgpPeers() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineDirectConnectBgpPeersRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of IDs.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of Resource.",
			},
			"bgp_peer_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of bgp peer.",
			},
			"virtual_interface_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of virtual interface.",
			},
			"direct_connect_gateway_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of direct connect gateway.",
			},
			"remote_asn": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The remote asn of bgp peer.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of query.",
			},
			"bgp_peers": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"virtual_interface_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of virtual interface.",
						},
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of account.",
						},
						"auth_key": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The key of auth.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of bgp peer.",
						},
						"session_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The session status of bgp peer.",
						},
						"creation_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of bgp peer.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of bgp peer.",
						},
						"bgp_peer_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of bgp peer.",
						},
						"bgp_peer_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of bgp peer.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Description of bgp peer.",
						},
						"remote_asn": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The remote asn of bgp peer.",
						},
						"local_asn": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The local asn of bgp peer.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineDirectConnectBgpPeersRead(d *schema.ResourceData, meta interface{}) error {
	service := NewDirectConnectBgpPeerService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineDirectConnectBgpPeers())
}
