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
				Default:     "Full",
				Description: "The type of backup. The value can be Full or Inc.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// 在更新时，timestamp 没发生变化，忽略变化
					if d.Id() != "" && !d.HasChange("time_point") {
						return true
					}
					return false
				},
			},
			"time_point": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Time point of backup, for example: 2021-11-09T06:07:26Z. Use lifecycle and ignore_changes in import.",
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
