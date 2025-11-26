package rds_postgresql_instance_backup_wal_log

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRdsPostgresqlInstanceBackupWalLogs() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRdsPostgresqlInstanceBackupWalLogsRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of the PostgreSQL instance.",
			},
			"backup_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the backup.",
			},
			"start_time": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The start time of the query. The format is yyyy-MM-ddTHH:mm:ssZ (UTC time).",
			},
			"end_time": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "The end time of the query. The format is yyyy-MM-ddTHH:mm:ssZ (UTC time). " +
					"Note: The maximum interval between start_time and end_time cannot exceed 7 days.",
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
			"wal_log_backups": {
				Description: "List of WAL log backups.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"backup_file_size": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The size of the WAL log backup file. The unit is bytes (Byte).",
						},
						"backup_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the WAL log backup.",
						},
						"backup_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the WAL log backup.",
						},
						"check_sum": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The checksum in the ETag format using the crc64 algorithm.",
						},
						"download_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The downloadable status of the WAL log backup.",
						},
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The project to which the instance of the WAL log backup belongs.",
						},
						"wal_log_backup_end_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The end time of the WAL log backup, in the format of yyyy-MM-ddTHH:mm:ssZ (UTC time).",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineRdsPostgresqlInstanceBackupWalLogsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewRdsPostgresqlInstanceBackupWalLogService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineRdsPostgresqlInstanceBackupWalLogs())
}
