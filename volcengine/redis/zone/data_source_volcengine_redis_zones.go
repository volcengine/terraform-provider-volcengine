package zone

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRedisZones() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRedisZonesRead,
		Schema: map[string]*schema.Schema{
			"region_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Id of Region.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of zone query.",
			},
			"zones": {
				Description: "The collection of zone query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the zone.",
						},
						"zone_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the zone.",
						},
						"zone_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the zone.",
						},
						"zone_status": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The status of the zone.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineRedisZonesRead(d *schema.ResourceData, meta interface{}) error {
	redisZoneService := NewZoneService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(redisZoneService, d, DataSourceVolcengineRedisZones())
}
