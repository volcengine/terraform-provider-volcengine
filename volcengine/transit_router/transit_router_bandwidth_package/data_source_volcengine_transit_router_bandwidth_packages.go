package transit_router_bandwidth_package

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineTransitRouterBandwidthPackages() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineTransitRouterBandwidthPackagesRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The ID list of the TransitRouter bandwidth package.",
			},
			"transit_router_peer_attachment_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID of the peer attachment.",
			},
			"transit_router_bandwidth_package_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the TransitRouter bandwidth package.",
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
			"bandwidth_packages": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the transit router bandwidth package.",
						},
						"transit_router_bandwidth_package_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the transit router attachment.",
						},
						"transit_router_bandwidth_package_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the transit router bandwidth package.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the transit router bandwidth package.",
						},
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The account id.",
						},
						"local_geographic_region_set_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The local geographic region set ID.",
						},
						"peer_geographic_region_set_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The peer geographic region set ID.",
						},
						"bandwidth": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The bandwidth peak of the transit router bandwidth package. Unit: Mbps.",
						},
						"remaining_bandwidth": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The remaining bandwidth of the transit router bandwidth package. Unit: Mbps.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the transit router bandwidth package.",
						},
						"business_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The business status of the transit router bandwidth package.",
						},
						"billing_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The billing type of the transit router bandwidth package.",
						},
						"creation_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the transit router bandwidth package.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of the transit router bandwidth package.",
						},
						"expired_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The expired time of the transit router bandwidth package.",
						},
						"delete_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The delete time of the transit router bandwidth package.",
						},
						"allocations": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The detailed information on cross regional connections associated with bandwidth packets.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"transit_router_peer_attachment_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The ID of the peer attachment.",
									},
									"allocate_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The delete time of the transit router bandwidth package.",
									},
									"local_region_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The local region id of the transit router.",
									},
									"delete_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The peer region id of the transit router.",
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

func dataSourceVolcengineTransitRouterBandwidthPackagesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewTRBandwidthPackageService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineTransitRouterBandwidthPackages())
}
