package rds_postgresql_backup_download

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRdsPostgresqlBackupDownloads() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRdsPostgresqlBackupDownloadsRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of the PostgreSQL instance.",
			},
			"backup_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the logical backup to be downloaded.",
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
			"downloads": {
				Description: "Download link information (if needed, please trigger the download task first).",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"backup_download_link": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The public network download address of the backup.",
						},
						"inner_backup_download_link": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The inner network download address of the backup.",
						},
						"backup_description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the backup set.",
						},
						"backup_file_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the backup file.",
						},
						"backup_file_size": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The size of the backup file, in Byte.",
						},
						"backup_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the backup.",
						},
						"backup_method": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the backup.",
						},
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the PostgreSQL instance.",
						},
						"link_expired_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Expiration time of the download link, format:yyyy-MM-ddTHH:mm:ss.sssZ (UTC time).",
						},
						"prepare_progress": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The prepare progress of the backup.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineRdsPostgresqlBackupDownloadsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewRdsPostgresqlBackupDownloadService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineRdsPostgresqlBackupDownloads())
}
