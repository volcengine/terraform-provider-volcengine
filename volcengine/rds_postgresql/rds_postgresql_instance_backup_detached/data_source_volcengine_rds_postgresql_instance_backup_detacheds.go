package rds_postgresql_instance_backup_detached

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRdsPostgresqlInstanceBackupDetacheds() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRdsPostgresqlInstanceBackupDetachedsRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID of the PostgreSQL instance.",
			},
			"backup_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID of the backup.",
			},
			"instance_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the PostgreSQL instance.",
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
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The project to which the instance belongs.",
			},
			"backup_status": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"Success"}, false),
				Description:  "The status of the backup.",
			},
			"backup_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"Full", "Increment"}, false),
				Description:  "The type of the backup.",
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
			// TODO: change this field to the target datasource
			"backups": {
				Description: "List of deleted instance backups.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"backup_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the backup.",
						},
						"backup_file_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the backup file.",
						},
						"backup_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the backup: Success, Failed, Running.",
						},
						"backup_progress": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The progress of the backup. The unit is percentage.",
						},
						"backup_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the backup: Full, Increment.",
						},
						"create_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The creation type of the backup: System, User.",
						},
						"backup_file_size": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The size of the backup file, in Byte.",
						},
						"backup_start_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The start time of the backup. The time format is yyyy-MM-ddTHH:mm:ss.sssZ (UTC time).",
						},
						"backup_end_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The end time of the backup. The time format is yyyy-MM-ddTHH:mm:ss.sssZ (UTC time).",
						},
						// 额外补充实例信息
						"instance_info": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Information about the PostgreSQL instance associated with this backup.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"instance_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The ID of the instance.",
									},
									"instance_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the instance.",
									},
									"instance_status": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The status of the instance.",
									},
									"db_engine_version": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The version of the database engine.",
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

func dataSourceVolcengineRdsPostgresqlInstanceBackupDetachedsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewRdsPostgresqlInstanceBackupDetachedService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineRdsPostgresqlInstanceBackupDetacheds())
}
