package transit_router_peer_attachment

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineTransitRouterPeerAttachments() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineTransitRouterPeerAttachmentsRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of IDs.",
			},
			"transit_router_attachment_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of transit router peer attachment.",
			},
			"transit_router_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of local transit router.",
			},
			"peer_transit_router_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of peer transit router.",
			},
			"peer_transit_router_region_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The region id of peer transit router.",
			},
			"tags": ve.TagsSchema(),
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of Resource.",
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
			"transit_router_attachments": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the transit router peer attachment.",
						},
						"transit_router_attachment_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the transit router peer attachment.",
						},
						"transit_router_attachment_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the transit router peer attachment.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the transit router peer attachment.",
						},
						"transit_router_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the local transit router.",
						},
						"peer_transit_router_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the peer transit router.",
						},
						"peer_transit_router_region_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The region id of the peer transit router.",
						},
						"transit_router_route_table_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The route table id of the transit router peer attachment.",
						},
						"transit_router_bandwidth_package_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The bandwidth package id of the transit router peer attachment.",
						},
						"bandwidth": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The bandwidth of the transit router peer attachment.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the transit router peer attachment.",
						},
						"creation_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The creation time of the transit router peer attachment.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of the transit router peer attachment.",
						},
						"tags": ve.TagsSchemaComputed(),
					},
				},
			},
		},
	}
}

func dataSourceVolcengineTransitRouterPeerAttachmentsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewTransitRouterPeerAttachmentService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineTransitRouterPeerAttachments())
}
