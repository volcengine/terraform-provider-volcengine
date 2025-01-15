package rds_mysql_backup

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRdsMysqlBackups() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRdsMysqlBackupsRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the instance.",
			},
			"backup_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the backup.",
			},
			"backup_start_time": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The start time of the backup.",
			},
			"backup_end_time": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The end time of the backup.",
			},
			"backup_type": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "Backup method, value: " +
					"Full: Full backup under physical backup type or library table backup under logical backup type. " +
					"Increment: Incremental backup under physical backup type. " +
					"DumpAll: Full database backup under logical backup type. " +
					"Description: There is no default value. When this field is not passed, " +
					"all backups of all methods under the query conditions limited by other fields are returned.",
			},
			"backup_status": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "Backup status, values: " +
					"Success: Success. " +
					"Failed: Failed. " +
					"Running: In progress. " +
					"Description: There is no default value. When this field is not passed, " +
					"all backups in all states under the query conditions limited by other fields are returned.",
			},
			"create_type": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "Creator of backup. Values: " +
					"System: System. " +
					"User: User. " +
					"Description: There is no default value. " +
					"When this field is not passed, " +
					"all types of backups under the query conditions limited by other fields are returned.",
			},
			"backup_method": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "Backup type, value: " +
					"Physical: Physical backup. Default value. " +
					"Logical: Logical backup. " +
					"Description: There is no default value. " +
					"When this field is not passed, backups of all states under the query conditions limited by other fields are returned.",
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
			"backups": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"backup_end_time": {
							Type:     schema.TypeString,
							Computed: true,
							Description: "The end time of backup, " +
								"in the format of yyyy-MM-ddTHH:mm:ss.sssZ (UTC time).",
						},
						"backup_file_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Backup file name.",
						},
						"backup_file_size": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Backup file size, in bytes.",
						},
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the backup.",
						},
						"backup_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the backup.",
						},
						"backup_method": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Backup type, value: Physical: Physical backup. Logical: Logical backup.",
						},
						"backup_region": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The region where the backup is located.",
						},
						"backup_start_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The start time of backup, in the format of yyyy-MM-ddTHH:mm:ss.sssZ (UTC time).",
						},
						"backup_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Backup status, values: Success. Failed. Running.",
						},
						"backup_type": {
							Type:     schema.TypeString,
							Computed: true,
							Description: "Backup method, values:\n" +
								"Full: Full backup under physical backup type or library table backup under logical backup type." +
								"\nIncrement: Incremental backup under physical backup type (created by the system)." +
								"\nDumpAll: Full database backup under logical backup type.\nDescription:" +
								"\nThere is no default value. When this field is not passed, " +
								"all types of backups under the query conditions limited by other fields are returned.",
						},
						"consistent_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The time point of a consistent snapshot is in the format of yyyy-MM-ddTHH:mm:ss.sssZ (UTC time).",
						},
						"create_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Creator of backup. Values: System. User.",
						},
						"db_table_infos": {
							Type:     schema.TypeList,
							Computed: true,
							Description: "The database table information contained in the backup set can include up to 10,000 tables." +
								"\nExplanation:\nWhen the database is empty, this field is not returned.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"database": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Database name.",
									},
									"tables": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Description: "Table names.",
									},
								},
							},
						},
						"download_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Download status. Values:\nNotDownload: Not downloaded.\nSuccess: Downloaded.\nFailed: Download failed.\nRunning: Downloading.",
						},
						"expired_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Expired time of backup, in the format of yyyy-MM-ddTHH:mm:ss.sssZ (UTC time).",
						},
						"is_encrypted": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Is the data backup encrypted? Value:\ntrue: Encrypted.\nfalse: Not encrypted.",
						},
						"is_expired": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the backup has expired. Value:\ntrue: Expired.\nfalse: Not expired.",
						},
						"error_message": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Error message.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineRdsMysqlBackupsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewRdsMysqlBackupService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineRdsMysqlBackups())
}
