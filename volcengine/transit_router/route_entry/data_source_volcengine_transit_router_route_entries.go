package route_entry

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineTransitRouterRouteEntries() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineTransitRouterRouteEntriesRead,
		Schema: map[string]*schema.Schema{
			"transit_router_route_table_id": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "The id of the route table.",
			},
			"destination_cidr_block": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The target network segment of the route entry.",
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"Available",
					"Creating",
					"Pending",
					"Deleting",
					"Conflicted",
				}, false),
				Description: "The status of the route entry.",
			},
			"transit_router_route_entry_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the route entry.",
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
				Description: "The ids of the transit router route entry.",
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

			"entries": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of route entries.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Description of the transit router route entry.",
						},
						"destination_cidr_block": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The target network segment of the route entry.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the route entry.",
						},
						"transit_router_route_entry_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the route entry.",
						},
						"transit_router_route_entry_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the route entry.",
						},
						"transit_router_route_entry_next_hop_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The next hop type of the routing entry. The value can be Attachment or BlackHole.",
						},
						"transit_router_route_entry_next_hop_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The next hot id of the routing entry.",
						},
						"transit_router_route_entry_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the route entry.",
						},
						"creation_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The creation time of the route entry.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of the route entry.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineTransitRouterRouteEntriesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewTREntryService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineTransitRouterRouteEntries())
}
