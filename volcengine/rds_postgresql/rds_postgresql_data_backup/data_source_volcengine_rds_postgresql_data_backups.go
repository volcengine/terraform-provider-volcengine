package rds_postgresql_data_backup

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRdsPostgresqlDataBackups() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRdsPostgresqlDataBackupsRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the PostgreSQL instance.",
			},
			"backup_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID of the backup.",
			},
			"backup_start_time": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The earliest time when the backup is created, in the format of yyyy-MM-ddTHH:mm:ss.sssZ (UTC time).",
			},
			"backup_end_time": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The latest time when the backup is created, in the format of yyyy-MM-ddTHH:mm:ss.sssZ (UTC time).",
			},
			"backup_status": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"Success", "Failed", "Running"}, false),
				Description:  "The status of the backup: Success, Failed, Running.",
			},
			"backup_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"Full", "Increment"}, false),
				Description:  "The type of the backup: Full, Increment.",
			},
			"download_status": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "The downloadable status of the backup set. " +
					"NotAllowed: download is not supported. NeedToPrepare: the backup set is in place and needs background preparation for backup. " +
					"LinkReady: the backup set is ready for download.",
			},
			"backup_method": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"Physical", "Logical"}, false),
				Description:  "The method of the backup: Physical, Logical.",
			},
			"backup_database_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the database included in the backup set. Only effective when the value of backup_method is Logical.",
			},
			"backup_description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the backup set.",
			},
			"create_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The creation type of the backup: System, User.",
			},
			"backup_scope": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The scope of the backup: Instance, Database.",
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
				Description: "The collection of the query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"backup_data_size": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The original size of the data contained in the backup, in Bytes.",
						},
						"backup_description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the backup set.",
						},
						"backup_end_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The end time of the backup. The time format is yyyy-MM-ddTHH:mm:ss.sssZ (UTC time).",
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
							Description: "The ID of the backup.",
						},
						"backup_method": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The method of the backup: Physical, Logical.",
						},
						"backup_progress": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The progress of the backup. The unit is percentage.",
						},
						"backup_scope": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The scope of the backup: Instance, Database.",
						},
						"backup_start_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The start time of the backup. The time format is yyyy-MM-ddTHH:mm:ss.sssZ (UTC time).",
						},
						"backup_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the backup: Full, Increment.",
						},
						"backup_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the backup: Success, Failed, Running.",
						},
						"create_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The creation type of the backup: System, User.",
						},
						"download_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The downloadable status of the backup set.",
						},
						"backup_meta": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The information about the databases included in the backup.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"db_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the database.",
									},
								},
							},
						},
						// InstanceInfo 是否有必要全字段返回，待商榷
					},
				},
			},
		},
	}
}

func dataSourceVolcengineRdsPostgresqlDataBackupsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewRdsPostgresqlDataBackupService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineRdsPostgresqlDataBackups())
}
