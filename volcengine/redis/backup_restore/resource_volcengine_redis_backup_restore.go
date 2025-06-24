package backup_restore

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Redis Backup Restore can be imported using the restore:instanceId, e.g.
```
$ terraform import volcengine_redis_backup_restore.default restore:redis-asdljioeixxxx
```
*/

func ResourceVolcengineRedisBackupRestore() *schema.Resource {
	resource := &schema.Resource{
		Read:   resourceVolcengineRedisBackupRestoreRead,
		Create: resourceVolcengineRedisBackupRestoreCreate,
		Delete: resourceVolcengineRedisBackupRestoreDelete,
		Update: resourceVolcengineRedisBackupRestoreUpdate,
		Importer: &schema.ResourceImporter{
			State: func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
				items := strings.Split(data.Id(), ":")
				if len(items) != 2 {
					return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
				}
				if err := data.Set("instance_id", items[1]); err != nil {
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
				Description: "Id of instance.",
			},
			"backup_type": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Default:     "Full",
				Description: "The type of backup. The value can be Full or Inc.",
			},
			"time_point": {
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"time_point", "backup_point_id"},
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// 在更新时， 如果backup_type为full或空，忽略变化
					if d.Id() != "" && (d.Get("backup_type") == "" || d.Get("backup_type") == "Full") {
						return true
					}
					return false
				},
				Description: "Time point of backup, for example: 2021-11-09T06:07:26Z. Use lifecycle and ignore_changes in import.",
			},
			"backup_point_id": {
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"time_point", "backup_point_id"},
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// 在更新时， 如果backup_type为full或空，忽略变化
					if d.Id() != "" && (d.Get("backup_type") == "Inc") {
						return true
					}
					return false
				},
				Description: "Backup ID, used to specify the backups to be used when restoring by the backup set. " +
					"When choosing to restore by backup set (i.e., BackupType is Full), this parameter is required. " +
					"Use lifecycle and ignore_changes in import.",
			},
		},
	}
	return resource
}

func resourceVolcengineRedisBackupRestoreRead(d *schema.ResourceData, meta interface{}) (err error) {
	redisBackupRestoreService := NewRedisBackupRestoreService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(redisBackupRestoreService, d, ResourceVolcengineRedisBackupRestore())
	if err != nil {
		return fmt.Errorf("error on read restore backup %v, %v", d.Id(), err)
	}
	return nil
}

func resourceVolcengineRedisBackupRestoreCreate(d *schema.ResourceData, meta interface{}) (err error) {
	redisBackupRestoreService := NewRedisBackupRestoreService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(redisBackupRestoreService, d, ResourceVolcengineRedisBackupRestore())
	if err != nil {
		return fmt.Errorf("error on restore backup %v, %v", d.Id(), err)
	}
	return resourceVolcengineRedisBackupRestoreRead(d, meta)
}

func resourceVolcengineRedisBackupRestoreUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	redisBackupRestoreService := NewRedisBackupRestoreService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(redisBackupRestoreService, d, ResourceVolcengineRedisBackupRestore())
	if err != nil {
		return fmt.Errorf("error on update backup %v, %v", d.Id(), err)
	}
	return resourceVolcengineRedisBackupRestoreRead(d, meta)
}

func resourceVolcengineRedisBackupRestoreDelete(d *schema.ResourceData, meta interface{}) (err error) {
	redisBackupRestoreService := NewRedisBackupRestoreService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(redisBackupRestoreService, d, ResourceVolcengineRedisBackupRestore())
	if err != nil {
		return fmt.Errorf("error on delete backup %v, %v", d.Id(), err)
	}
	return nil
}
