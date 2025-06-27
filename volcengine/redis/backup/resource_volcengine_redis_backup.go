package backup

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Redis Backup can be imported using the instanceId:backupId, e.g.
```
$ terraform import volcengine_redis_backup.default redis-cn02aqusft7ws****:b-cn02xmmrp751i9cdzcphjmk4****
```

*/

func ResourceVolcengineRedisBackup() *schema.Resource {
	resource := &schema.Resource{
		Read:   resourceVolcengineRedisBackupRead,
		Create: resourceVolcengineRedisBackupCreate,
		Delete: resourceVolcengineRedisBackupDelete,
		Importer: &schema.ResourceImporter{
			State: func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
				items := strings.Split(data.Id(), ":")
				if len(items) != 2 {
					return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
				}
				if err := data.Set("instance_id", items[0]); err != nil {
					return []*schema.ResourceData{data}, err
				}
				if err := data.Set("backup_point_id", items[1]); err != nil {
					return []*schema.ResourceData{data}, err
				}
				return []*schema.ResourceData{data}, nil
			},
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Id of instance to create backup.",
			},
			"backup_point_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "Set the backup name for the manually created backup.",
			},
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
			"instance_info": {
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
						"deletion_protection": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the deletion protection function of the instance.",
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
						"shard_capacity": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Capacity of shard.",
						},
						"shard_number": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of shards in the instance.",
						},
						"total_capacity": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Total capacity of instance.",
						},
						"zone_ids": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of id of zone.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The private network ID of the instance.",
						},
					},
				},
			},
			"ttl": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Backup retention days.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Project name of instance.",
			},
			"backup_point_download_urls": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The download address information of the backup file to which the current backup point belongs.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"shard_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The shard ID where the RDB file is located.",
						},
						"rdb_file_size": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "RDB file size, unit: Byte.",
						},
						"public_download_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The public network download address for RDB files.",
						},
						"private_download_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The private network download address for RDB files.",
						},
					},
				},
			},
		},
	}
	return resource
}

func resourceVolcengineRedisBackupCreate(d *schema.ResourceData, meta interface{}) (err error) {
	redisBackupService := NewRedisBackupService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(redisBackupService, d, ResourceVolcengineRedisBackup())
	if err != nil {
		return fmt.Errorf("error on creating redis backup %v, %v", d.Id(), err)
	}
	return resourceVolcengineRedisBackupRead(d, meta)
}

func resourceVolcengineRedisBackupDelete(d *schema.ResourceData, meta interface{}) (err error) {
	return nil
}

func resourceVolcengineRedisBackupRead(d *schema.ResourceData, meta interface{}) (err error) {
	redisBackupService := NewRedisBackupService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(redisBackupService, d, ResourceVolcengineRedisBackup())
	if err != nil {
		return fmt.Errorf("error on reading redis backup %q,%s", d.Id(), err)
	}
	return err
}
