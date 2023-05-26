package pitr_time_period

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRedisPitrTimeWindows() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRedisPitrTimeWindowsRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The ids of the instances.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of redis instances time window query.",
			},
			"periods": {
				Description: "The list of time windows.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"start_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The recoverable start time (in UTC time) supported when restoring data by point in time.",
						},
						"end_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Recoverable end time (UTC time) supported when restoring data by point in time.",
						},
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The instance id.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineRedisPitrTimeWindowsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewVolcengineRedisPitrTimeWindowService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineRedisPitrTimeWindows())
}
