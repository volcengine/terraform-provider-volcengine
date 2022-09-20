package region

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineESCloudRegions() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRegionsRead,
		Schema: map[string]*schema.Schema{
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of region query.",
			},
			"regions": {
				Description: "The collection of region query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"region_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the region.",
						},
						"region_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of region.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineRegionsRead(d *schema.ResourceData, meta interface{}) error {
	regionService := NewRegionService(meta.(*ve.SdkClient))
	return regionService.Dispatcher.Data(regionService, d, DataSourceVolcengineESCloudRegions())
}
