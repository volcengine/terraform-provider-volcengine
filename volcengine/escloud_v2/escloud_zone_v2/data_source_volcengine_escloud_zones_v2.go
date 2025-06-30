package escloud_zone_v2

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineEscloudZoneV2s() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineEscloudZoneV2sRead,
		Schema: map[string]*schema.Schema{
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
			"zones": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"region_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The region ID of zone.",
						},
						"zone_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of zone.",
						},
						"zone_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of zone.",
						},
						"zone_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of zone.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineEscloudZoneV2sRead(d *schema.ResourceData, meta interface{}) error {
	service := NewEscloudZoneV2Service(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineEscloudZoneV2s())
}
