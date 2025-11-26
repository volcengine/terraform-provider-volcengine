package rds_postgresql_backup_policy

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRdsPostgresqlBackupPolicys() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRdsPostgresqlBackupPolicysRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of the PostgreSQL instance.",
			},
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
			"backup_policy": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"backup_retention_period": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The backup retention period.",
						},
						"data_incr_backup_periods": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The backup increment data backup periods.",
						},
						"full_backup_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The time when the backup task is executed. Format: HH:mmZ-HH:mmZ (UTC time).",
						},
						"full_backup_period": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The full backup period.",
						},
						"hourly_incr_backup_enable": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to enable the high-frequency backup function.",
						},
						"increment_backup_frequency": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The frequency of increment backup.",
						},
						"wal_log_space_limit_enable": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Status of the local remaining available space protection function.",
						},
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the PostgreSQL instance.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineRdsPostgresqlBackupPolicysRead(d *schema.ResourceData, meta interface{}) error {
	service := NewRdsPostgresqlBackupPolicyService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineRdsPostgresqlBackupPolicys())
}
