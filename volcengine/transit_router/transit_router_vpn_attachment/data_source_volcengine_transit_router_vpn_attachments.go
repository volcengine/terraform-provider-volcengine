package transit_router_vpn_attachment

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineTransitRouterVpnAttachments() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineTransitRouterVpnAttachmentsRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The ID list of the VPN attachment.",
			},
			"transit_router_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of the transit router.",
			},
			"vpn_connection_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID of the IPSec connection.",
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
						"vpn_connection_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the IPSec connection.",
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
						"zone_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the availability zone.",
						},
						"tags": ve.TagsSchemaComputed(),
					},
				},
			},
		},
	}
}

func dataSourceVolcengineTransitRouterVpnAttachmentsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewTRVpnAttachmentService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineTransitRouterVpnAttachments())
}
