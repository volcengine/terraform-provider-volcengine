package kms_region

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineKmsRegions() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineKmsRegionsRead,
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
				Description: "The supported regions.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"region_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The region ID.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineKmsRegionsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewKmsRegionService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineKmsRegions())
}
