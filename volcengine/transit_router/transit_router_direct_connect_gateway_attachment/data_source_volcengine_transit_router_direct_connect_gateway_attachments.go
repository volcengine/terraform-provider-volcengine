package transit_router_direct_connect_gateway_attachment

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineTransitRouterDirectConnectGatewayAttachments() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineTransitRouterDirectConnectGatewayAttachmentsRead,
		Schema: map[string]*schema.Schema{
			"transit_router_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of the transit router.",
			},
			"transit_router_attachment_ids": {
				Type:     schema.TypeSet,
				Set:      schema.HashString,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "ID of the network instance connection.",
			},
			"direct_connect_gateway_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ID of the direct connection gateway.",
			},
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
							Description: "The status of the network instance connection.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description info.",
						},
						"direct_connect_gateway_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The direct connect gateway id.",
						},
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The account id.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineTransitRouterDirectConnectGatewayAttachmentsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewTransitRouterDirectConnectGatewayAttachmentService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineTransitRouterDirectConnectGatewayAttachments())
}
