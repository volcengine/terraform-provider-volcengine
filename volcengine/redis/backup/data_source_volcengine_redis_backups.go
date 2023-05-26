package backup

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRedisBackups() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRedisBackupsRead,
		Schema: map[string]*schema.Schema{
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Id of instance.",
			},
			"start_time": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Query start time.",
			},
			"end_time": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Query end time.",
			},
			"backup_strategy_list": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The list of backup strategy, support AutomatedBackup and ManualBackup.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
					ValidateFunc: validation.StringInSlice([]string{
						"ManualBackup",
						"AutomatedBackup",
					}, false),
				},
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of backup query.",
			},
			"backups": {
				Description: "Information of backups.",
				Computed:    true,
				Type:        schema.TypeList,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"backup_point_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of backup point.",
						},
						"backup_strategy": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Backup strategy.",
						},
						"backup_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Backup type.",
						},
						"end_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "End time of backup.",
						},
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Id of instance.",
						},
						"size": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Size in MiB.",
						},
						"start_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Start time of backup.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Status of backup (Creating/Available/Unavailable/Deleting).",
						},
						"instance_detail": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Information of instance.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"account_id": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Id of account.",
									},
									"arch_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Arch type of instance(Standard/Cluster).",
									},
									"charge_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Charge type of instance(Postpaid/Prepaid).",
									},
									"engine_version": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Engine version of instance.",
									},
									"expired_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Expired time of instance.",
									},
									"instance_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Id of instance.",
									},
									"instance_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Name of instance.",
									},
									"maintenance_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The maintainable period (in UTC) of the instance.",
									},
									"network_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Network type of instance.",
									},
									"project_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Project name of instance.",
									},
									"region_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Id of region.",
									},
									"replicas": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Count of replica in which shard.",
									},
									"server_cpu": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Count of cpu cores of instance.",
									},
									"shard_capacity": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Capacity of shard.",
									},
									"shard_count": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Count of shard.",
									},
									"total_capacity": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Total capacity of instance.",
									},
									"used_capacity": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Capacity used of this instance.",
									},
									"zone_ids": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "List of id of zone.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"vpc_info": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Information of vpc.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Id of vpc.",
												},
												"name": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Name of vpc.",
												},
											},
										},
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

func dataSourceVolcengineRedisBackupsRead(d *schema.ResourceData, meta interface{}) error {
	redisBackupService := NewRedisBackupService(meta.(*ve.SdkClient))
	err := ve.DefaultDispatcher().Data(redisBackupService, d, DataSourceVolcengineRedisBackups())
	return err
}
