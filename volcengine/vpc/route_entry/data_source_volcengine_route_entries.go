package route_entry

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRouteEntries() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRouteEntriesRead,
		Schema: map[string]*schema.Schema{
			"route_table_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "An id of route table.",
			},
			"route_entry_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A type of route entry.",
			},
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "A list of route entry ids.",
			},
			"route_entry_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A name of route entry.",
			},
			"destination_cidr_block": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A destination CIDR block of route entry.",
			},
			"next_hop_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "An id of next hop.",
			},
			"next_hop_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A type of next hop.",
			},

			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of route entry query.",
			},
			"route_entries": {
				Description: "The collection of route tables.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the route entry.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the route entry.",
						},
						"destination_cidr_block": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The destination CIDR block of the route entry.",
						},
						"route_entry_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the route entry.",
						},
						"route_entry_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the route entry.",
						},
						"route_table_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the route table to which the route entry belongs.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the route entry.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the route entry.",
						},
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the virtual private cloud (VPC) to which the route entry belongs.",
						},
						"next_hop_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the next hop.",
						},
						"next_hop_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the next hop.",
						},
						"next_hop_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the next hop.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineRouteEntriesRead(d *schema.ResourceData, meta interface{}) error {
	routeEntryService := NewRouteEntryService(meta.(*ve.SdkClient))
	return routeEntryService.Dispatcher.Data(routeEntryService, d, DataSourceVolcengineRouteEntries())
}
