package cen_inter_region_bandwidth

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineCenInterRegionBandwidths() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineCenInterRegionBandwidthsRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of cen inter region bandwidth IDs.",
			},

			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of cen inter region bandwidth query.",
			},
			"inter_region_bandwidths": {
				Description: "The collection of cen inter region bandwidth query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the cen inter region bandwidth.",
						},
						"inter_region_bandwidth_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the cen inter region bandwidth.",
						},
						"creation_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the cen inter region bandwidth.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of the cen inter region bandwidth.",
						},
						"cen_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The cen ID of the cen inter region bandwidth.",
						},
						"local_region_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The local region id of the cen inter region bandwidth.",
						},
						"peer_region_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The peer region id of the cen inter region bandwidth.",
						},
						"bandwidth": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The bandwidth of the cen inter region bandwidth.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the cen inter region bandwidth.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineCenInterRegionBandwidthsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewCenInterRegionBandwidthService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineCenInterRegionBandwidths())
}
