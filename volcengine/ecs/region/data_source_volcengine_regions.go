package region

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRegions() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengieRegionsRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of region ids.",
			},
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
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the region.",
						},
						"region_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the region.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengieRegionsRead(d *schema.ResourceData, meta interface{}) error {
	regionService := NewRegionService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(regionService, d, DataSourceVolcengineRegions())
}
