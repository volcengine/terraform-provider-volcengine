package nlb_zone

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineNlbZones() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineNlbZonesRead,
		Schema: map[string]*schema.Schema{
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"zones": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of zones.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"zone_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the zone.",
						},
					},
				},
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of zones.",
			},
		},
	}
}

func dataSourceVolcengineNlbZonesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewNlbZoneService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineNlbZones())
}
