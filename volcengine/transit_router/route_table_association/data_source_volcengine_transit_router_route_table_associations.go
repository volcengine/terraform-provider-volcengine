package route_table_association

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineTransitRouterRouteTableAssociations() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineTransitRouterRouteTableAssociationsRead,
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
			"associations": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of route table associations.",
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
					},
				},
			},
		},
	}
}

func dataSourceVolcengineTransitRouterRouteTableAssociationsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewTRRouteTableAssociationService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineTransitRouterRouteTableAssociations())
}
