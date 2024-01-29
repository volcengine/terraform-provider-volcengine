package sqlserver_backup

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRdsMssqlBackups() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineMssqlBackupsRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
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
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The type of the backup.",
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
						"backup_start_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The start time of the backup.",
						},
						"backup_end_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The end time of the backup.",
						},
						"backup_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the backup.",
						},
						"backup_file_size": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The size of the backup file.",
						},
						"backup_method": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the backup method.",
						},
						"backup_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the backup.",
						},
						"create_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the backup create.",
						},
						"backup_database_detail": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The detail of the database.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"backup_start_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The start time of the backup.",
									},
									"backup_end_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The end time of the backup.",
									},
									"backup_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type of the backup.",
									},
									"backup_file_size": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The size of the backup file.",
									},
									"backup_file_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the backup file.",
									},
									"database_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the database.",
									},
									"backup_download_link_inner": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Intranet backup download link.",
									},
									"backup_download_link_eip": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "External backup download link.",
									},
									"link_expired_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Download link expiration time.",
									},
									"download_progress": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Backup file preparation progress, unit: %.",
									},
									"download_status": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Download status.",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineMssqlBackupsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewRdsMssqlBackupService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineRdsMssqlBackups())
}
