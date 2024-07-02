package transit_router_vpc_attachment

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineTransitRouterVpcAttachments() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineTransitRouterVpcAttachmentsRead,
		Schema: map[string]*schema.Schema{
			"transit_router_attachment_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of Transit Router Attachment ids.",
			},
			"transit_router_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of transit router.",
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of vpc.",
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
			"attachments": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"transit_router_attachment_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the transit router attachment.",
						},
						"transit_router_attachment_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the transit router attachment.",
						},
						"transit_router_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the transit router.",
						},
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of vpc.",
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
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the transit router.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description info.",
						},
						"auto_publish_route_enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to auto publish route of the transit router to vpc instance.",
						},
						"tags": ve.TagsSchemaComputed(),
						"attach_points": {
							Description: "The collection of attach points.",
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"zone_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The ID of zone.",
									},
									"subnet_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The ID of subnet.",
									},
									"network_interface_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The ID of network interface.",
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

func dataSourceVolcengineTransitRouterVpcAttachmentsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineTransitRouterVpcAttachments())
}
