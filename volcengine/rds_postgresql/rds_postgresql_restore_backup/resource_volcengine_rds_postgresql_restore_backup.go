package rds_postgresql_restore_backup

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
RdsPostgresqlRestoreBackup can be imported using the id, e.g.
```
$ terraform import volcengine_rds_postgresql_restore_backup.default resource_id
```

*/

func ResourceVolcengineRdsPostgresqlRestoreBackup() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineRdsPostgresqlRestoreBackupCreate,
		Read:   resourceVolcengineRdsPostgresqlRestoreBackupRead,
		Update: resourceVolcengineRdsPostgresqlRestoreBackupUpdate,
		Delete: resourceVolcengineRdsPostgresqlRestoreBackupDelete,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			// 用于将备份恢复到已有实例
			"backup_id": {
				Type:     schema.TypeString,
				Required: true,
				Description: "The backup ID used for restore." +
					"Only supports restoring data to an existing instance through logical backup.",
			},
			"source_db_instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the instance to which the backup belongs.",
			},
			"target_db_instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the target instance for restore.",
			},
			"target_db_instance_account": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The account used as the Owner of the newly restored database in the target instance.",
			},
			"databases": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "Information of the database to be restored.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"db_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Original database name.",
						},
						"new_db_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "New database name.",
						},
					},
				},
			},
		},
	}
	return resource
}

func resourceVolcengineRdsPostgresqlRestoreBackupCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlRestoreBackupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineRdsPostgresqlRestoreBackup())
	if err != nil {
		return fmt.Errorf("error on creating rds_postgresql_restore_backup %q, %s", d.Id(), err)
	}
	return resourceVolcengineRdsPostgresqlRestoreBackupRead(d, meta)
}

func resourceVolcengineRdsPostgresqlRestoreBackupRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlRestoreBackupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineRdsPostgresqlRestoreBackup())
	if err != nil {
		return fmt.Errorf("error on reading rds_postgresql_restore_backup %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineRdsPostgresqlRestoreBackupUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlRestoreBackupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineRdsPostgresqlRestoreBackup())
	if err != nil {
		return fmt.Errorf("error on updating rds_postgresql_restore_backup %q, %s", d.Id(), err)
	}
	return resourceVolcengineRdsPostgresqlRestoreBackupRead(d, meta)
}

func resourceVolcengineRdsPostgresqlRestoreBackupDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlRestoreBackupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineRdsPostgresqlRestoreBackup())
	if err != nil {
		return fmt.Errorf("error on deleting rds_postgresql_restore_backup %q, %s", d.Id(), err)
	}
	return err
}
