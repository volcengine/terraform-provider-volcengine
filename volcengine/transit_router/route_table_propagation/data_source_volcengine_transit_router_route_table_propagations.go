package route_table_propagation

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineTransitRouterRouteTablePropagations() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineTransitRouterRouteTablePropagationsRead,
		Schema: map[string]*schema.Schema{
			"transit_router_attachment_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID of the network instance connection.",
			},
			"transit_router_route_table_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the routing table associated with the transit router instance.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},

			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of data query.",
			},
			"propagations": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of route table propagations.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the route table.",
						},
						"transit_router_attachment_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the network instance connection.",
						},
						"transit_router_route_table_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the routing table associated with the transit router instance.",
						},
						"creation_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The creation time of the route table propagation.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineTransitRouterRouteTablePropagationsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewTRRouteTablePropagationService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineTransitRouterRouteTablePropagations())
}
