package transit_router

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineTransitRouters() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineTransitRoutersRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of Transit Router ids.",
			},
			"transit_router_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name info.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ProjectName of the transit router.",
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
				Description: "The total count of query.",
			},
			"transit_routers": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the transit router.",
						},
						"transit_router_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the transit router.",
						},
						"transit_router_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the transit router.",
						},
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of account.",
						},
						"creation_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time.",
						},
						"overdue_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The overdue time.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description info.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the transit router.",
						},
						"business_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The business status of the transit router.",
						},
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ProjectName of the transit router.",
						},
						"tags": ve.TagsSchemaComputed(),
						"transit_router_attachments": {
							Description: "The attachments of transit router.",
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"creation_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The create time.",
									},
									"update_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The update time.",
									},
									"status": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The status of the transit router.",
									},
									"resource_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The id of resource.",
									},
									"resource_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type of resource.",
									},
									"transit_router_route_table_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The id of transit router route table.",
									},
									"transit_router_attachment_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The id of transit router attachment.",
									},
									"transit_router_attachment_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of transit router attachment.",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineTransitRoutersRead(d *schema.ResourceData, meta interface{}) error {
	service := NewService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineTransitRouters())
}
