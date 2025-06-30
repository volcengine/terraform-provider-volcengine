package rabbitmq_region

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRabbitmqRegions() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRabbitmqRegionsRead,
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
			"regions": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"region_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of region.",
						},
						"region_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of region.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of region.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of region.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineRabbitmqRegionsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewRabbitmqRegionService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineRabbitmqRegions())
}
