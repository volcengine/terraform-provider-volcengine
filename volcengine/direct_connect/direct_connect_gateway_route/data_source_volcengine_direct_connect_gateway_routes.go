package direct_connect_gateway_route

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineDirectConnectGatewayRoutes() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineDirectConnectGatewayRoutesRead,
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
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"direct_connect_gateway_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of direct connect gateway.",
			},
			"destination_cidr_block": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The cidr block.",
			},
			"next_hop_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of next hop.",
			},
			"next_hop_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The type of next hop.",
			},
			"route_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The type of route. The value can be BGP or CEN or Static.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of query.",
			},
			"direct_connect_gateway_routes": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of account.",
						},
						"direct_connect_gateway_route_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of direct connect gateway route.",
						},
						"next_hop_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of next hop.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status info.",
						},
						"creation_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time.",
						},
						"route_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of route.",
						},
						"next_hop_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of next hop.",
						},
						"direct_connect_gateway_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of direct connect gateway.",
						},
						"destination_cidr_block": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The cidr block.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineDirectConnectGatewayRoutesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewDirectConnectGatewayRouteService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineDirectConnectGatewayRoutes())
}
