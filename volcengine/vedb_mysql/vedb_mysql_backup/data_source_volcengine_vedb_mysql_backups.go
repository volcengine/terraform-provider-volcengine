package vedb_mysql_backup

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineVedbMysqlBackups() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineVedbMysqlBackupsRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of the instance.",
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
			"backup_status": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The status of the backup.",
			},
			"backup_method": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Backup method. Currently, only physical backup is supported. The value is Physical.",
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
						"consistent_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The time point of consistent backup, in the format: yyyy-MM-ddTHH:mm:ssZ (UTC time).",
						},
						"backup_policy": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Data backup strategy for instances.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"instance_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The id of the instance.",
									},
									"backup_time": {
										Type:     schema.TypeString,
										Computed: true,
										Description: "The time for executing the backup task. " +
											"The interval window is two hours. Format: HH:mmZ-HH:mmZ (UTC time).",
									},
									"full_backup_period": {
										Type:     schema.TypeString,
										Computed: true,
										Description: "Full backup period. " +
											"Multiple values are separated by English commas (,). " +
											"Values:\nMonday: Monday.\nTuesday: Tuesday.\nWednesday: Wednesday.\nThursday: Thursday.\nFriday: Friday.\nSaturday: Saturday.\nSunday: Sunday.",
									},
									"backup_retention_period": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Data backup retention period, value: 7 to 30 days.",
									},
									"continue_backup": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether to enable continuous backup. The value is fixed as true.",
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

func dataSourceVolcengineVedbMysqlBackupsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewVedbMysqlBackupService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineVedbMysqlBackups())
}
