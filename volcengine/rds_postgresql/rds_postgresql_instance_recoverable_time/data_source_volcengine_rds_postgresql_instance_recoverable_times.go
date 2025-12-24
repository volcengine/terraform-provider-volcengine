package rds_postgresql_instance_recoverable_time

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRdsPostgresqlInstanceRecoverableTimes() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRdsPostgresqlInstanceRecoverableTimesRead,
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
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of the Postgresql instance.",
			},
			"recoverable_time_info": {
				Description: "The earliest and latest recoverable times of the instance (UTC time). " +
					"If it is empty, it indicates that the instance is currently unrecoverable.",
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"earliest_recoverable_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The earliest recoverable time of the instance (UTC time).",
						},
						"latest_recoverable_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The latest recoverable time of the instance (UTC time).",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineRdsPostgresqlInstanceRecoverableTimesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewRdsPostgresqlInstanceRecoverableTimeService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineRdsPostgresqlInstanceRecoverableTimes())
}
