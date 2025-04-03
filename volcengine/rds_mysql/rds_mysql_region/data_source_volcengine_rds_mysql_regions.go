package rds_mysql_region

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRdsMysqlRegions() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRdsMysqlRegionsRead,
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

func dataSourceVolcengineRdsMysqlRegionsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewRdsMysqlRegionService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineRdsMysqlRegions())
}
