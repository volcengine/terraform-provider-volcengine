package nas_zone

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineNasZones() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineNasZonesRead,
		Schema: map[string]*schema.Schema{
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
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status info.",
						},
						"sales": {
							Description: "The collection of sales info.",
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"status": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The status info.",
									},
									"protocol_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type of protocol.",
									},
									"storage_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type of storage.",
									},
									"file_system_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type of file system.",
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

func dataSourceVolcengineNasZonesRead(d *schema.ResourceData, meta interface{}) error {
	zoneService := NewZoneService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(zoneService, d, DataSourceVolcengineNasZones())
}
