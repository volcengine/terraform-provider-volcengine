package vpn_gateway_route

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineVpnGatewayRoutes() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineVpnGatewayRoutesRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of VPN gateway route ids.",
			},
			"vpn_gateway_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "An ID of VPN gateway.",
			},
			"next_hop_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "An ID of next hop.",
			},
			"destination_cidr_block": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsCIDR,
				Description:  "A destination cidr block.",
			},

			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of VPN gateway route query.",
			},
			"vpn_gateway_routes": {
				Description: "The collection of VPN gateway route query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the VPN gateway route.",
						},
						"vpn_gateway_route_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the VPN gateway route.",
						},
						"vpn_gateway_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the VPN gateway of the VPN gateway route.",
						},
						"destination_cidr_block": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The destination cidr block of the VPN gateway route.",
						},
						"next_hop_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The next hop id of the VPN gateway route.",
						},
						"creation_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of VPN gateway route.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of VPN gateway route.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the VPN gateway route.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineVpnGatewayRoutesRead(d *schema.ResourceData, meta interface{}) error {
	routeService := NewVpnGatewayRouteService(meta.(*ve.SdkClient))
	return routeService.Dispatcher.Data(routeService, d, DataSourceVolcengineVpnGatewayRoutes())
}
