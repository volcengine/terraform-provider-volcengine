package rabbitmq_zone

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRabbitmqZones() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRabbitmqZonesRead,
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
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of zone.",
						},
						"status": {
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

func dataSourceVolcengineRabbitmqZonesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewRabbitmqZoneService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineRabbitmqZones())
}
