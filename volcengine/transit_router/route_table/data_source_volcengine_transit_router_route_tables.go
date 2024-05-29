package route_table

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineTransitRouterRouteTables() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineTransitRouterRouteTablesRead,
		Schema: map[string]*schema.Schema{
			"transit_router_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of the transit router.",
			},
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				MaxItems:    100,
				MinItems:    1,
				Set:         schema.HashString,
				Description: "The ids of the transit router route table.",
			},
			"transit_router_route_table_type": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"System",
					"Custom",
				}, false),
				Description: "The type of the route table. The value can be System or Custom.",
			},
			"tags": ve.TagsSchema(),
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

			"route_tables": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of route tables query.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"creation_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The creation time of the route table.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the route table.",
						},
						"transit_router_route_table_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the route table.",
						},
						"transit_router_route_table_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the route table.",
						},
						"transit_router_route_table_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of route table.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of the route table.",
						},
						"tags": ve.TagsSchemaComputed(),
					},
				},
			},
		},
	}
}

func dataSourceVolcengineTransitRouterRouteTablesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewTRRouteTableService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineTransitRouterRouteTables())
}
