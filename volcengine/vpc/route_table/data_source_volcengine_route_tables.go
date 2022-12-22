package route_table

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRouteTables() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRouteTablesRead,
		Schema: map[string]*schema.Schema{
			"vpc_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "An id of VPC.",
			},
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "A list of route table ids.",
			},
			"route_table_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A name of route table.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ProjectName of the route table.",
			},

			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of route table query.",
			},
			"route_tables": {
				Description: "The collection of route tables.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the route table.",
						},
						"route_table_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the route table.",
						},
						"route_table_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the route table.",
						},
						"route_table_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the route table.",
						},
						"subnet_ids": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The list of the subnet ids to which the entry table associates.",
						},
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the virtual private cloud (VPC) to which the route entry belongs.",
						},
						"vpc_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the virtual private cloud (VPC) to which the route entry belongs.",
						},
						"creation_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the route table.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The last update time of the route table.",
						},
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The account id of the route table creator.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the route table.",
						},
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ProjectName of the route table.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineRouteTablesRead(d *schema.ResourceData, meta interface{}) error {
	routeTableService := NewRouteTableService(meta.(*ve.SdkClient))
	return routeTableService.Dispatcher.Data(routeTableService, d, DataSourceVolcengineRouteTables())
}
