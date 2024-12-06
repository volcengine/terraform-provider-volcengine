package vedb_mysql_backup

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
VedbMysqlBackup can be imported using the instance id and backup id, e.g.
```
$ terraform import volcengine_vedb_mysql_backup.default instanceID:backupId
```

*/

func ResourceVolcengineVedbMysqlBackup() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVedbMysqlBackupCreate,
		Read:   resourceVolcengineVedbMysqlBackupRead,
		Update: resourceVolcengineVedbMysqlBackupUpdate,
		Delete: resourceVolcengineVedbMysqlBackupDelete,
		Importer: &schema.ResourceImporter{
			State: func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
				items := strings.Split(data.Id(), ":")
				if len(items) != 2 {
					return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
				}
				if err := data.Set("instance_id", items[0]); err != nil {
					return []*schema.ResourceData{data}, err
				}
				if err := data.Set("backup_id", items[1]); err != nil {
					return []*schema.ResourceData{data}, err
				}
				return []*schema.ResourceData{data}, nil
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(90 * time.Minute),
			Update: schema.DefaultTimeout(90 * time.Minute),
			Delete: schema.DefaultTimeout(90 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the instance.",
			},
			"backup_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The id of the backup.",
			},
			"backup_type": {
				Optional:    true,
				Computed:    true,
				Type:        schema.TypeString,
				ForceNew:    true,
				Description: "Backup type. Currently, only full backup is supported. The value is Full.",
			},
			"backup_method": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "Backup method. Currently, only physical backup is supported. The value is Physical.",
			},
			"backup_policy": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				Description: "Data backup strategy for instances.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"backup_time": {
							Type:     schema.TypeString,
							Required: true,
							Description: "The time for executing the backup task has an interval window of 2 hours and must be an even-hour time. " +
								"Format: HH:mmZ-HH:mmZ (UTC time).",
						},
						"full_backup_period": {
							Type:     schema.TypeString,
							Required: true,
							DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
								if len(old) != len(new) {
									return false
								}
								oldArr := strings.Split(old, ",")
								newArr := strings.Split(new, ",")
								sort.Strings(oldArr)
								sort.Strings(newArr)
								return reflect.DeepEqual(oldArr, newArr)
							},
							Description: "Full backup period. " +
								"It is recommended to select at least 2 days per week for full backup. " +
								"Multiple values are separated by English commas (,). Values: Monday: Monday. Tuesday: Tuesday. Wednesday: Wednesday. Thursday: Thursday. Friday: Friday. Saturday: Saturday. Sunday: Sunday.",
						},
						"backup_retention_period": {
							// 文档写的string，实际返回是int
							// error on reading vedb_mysql_backup "vedbm-ajg6odlhufzc:snap-67175524-a848", backup_policy.0.backup_retention_period: '' expected type 'string', got unconvertible type 'float64'
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Data backup retention period, value: 7 to 30 days.",
						},
					},
				},
			},
		},
	}
	return resource
}

func resourceVolcengineVedbMysqlBackupCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVedbMysqlBackupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineVedbMysqlBackup())
	if err != nil {
		return fmt.Errorf("error on creating vedb_mysql_backup %q, %s", d.Id(), err)
	}
	return resourceVolcengineVedbMysqlBackupRead(d, meta)
}

func resourceVolcengineVedbMysqlBackupRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVedbMysqlBackupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineVedbMysqlBackup())
	if err != nil {
		return fmt.Errorf("error on reading vedb_mysql_backup %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineVedbMysqlBackupUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVedbMysqlBackupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineVedbMysqlBackup())
	if err != nil {
		return fmt.Errorf("error on updating vedb_mysql_backup %q, %s", d.Id(), err)
	}
	return resourceVolcengineVedbMysqlBackupRead(d, meta)
}

func resourceVolcengineVedbMysqlBackupDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVedbMysqlBackupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineVedbMysqlBackup())
	if err != nil {
		return fmt.Errorf("error on deleting vedb_mysql_backup %q, %s", d.Id(), err)
	}
	return err
}
