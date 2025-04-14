package rds_mysql_backup

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
RdsMysqlBackup can be imported using the id, e.g.
```
$ terraform import volcengine_rds_mysql_backup.default instanceId:backupId
```

*/

func ResourceVolcengineRdsMysqlBackup() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineRdsMysqlBackupCreate,
		Read:   resourceVolcengineRdsMysqlBackupRead,
		Delete: resourceVolcengineRdsMysqlBackupDelete,
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
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
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
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Description: "Backup method. Value range: Full, " +
					"full backup under physical backup type. Default value. " +
					"DumpAll: full database backup under logical backup type. " +
					"Prerequisite: If you need to create a full database backup of logical backup type, " +
					"that is, when the value of BackupType is DumpAll, " +
					"the backup type should be set to logical backup, " +
					"that is, the value of BackupMethod should be Logical. " +
					"If you need to create a database table backup of logical backup type, " +
					"you do not need to pass in this field. " +
					"You only need to specify the database and table to be backed up in the BackupMeta field.",
			},
			"backup_meta": {
				Type:     schema.TypeSet,
				Optional: true,
				ForceNew: true,
				Description: "When creating a library table backup of logical backup type, " +
					"it is used to specify the library table information to be backed up.\n" +
					"Prerequisite: When the value of BackupMethod is Logical, and the BackupType field is not passed." +
					"\nMutual exclusion situation: When the value of the BackupType field is DumpAll, " +
					"this field is not effective.\nQuantity limit: When creating a specified library table backup, " +
					"the upper limit of the number of libraries is 5000, " +
					"and the upper limit of the number of tables in each library is 5000.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if d.Get("backup_type").(string) == "DumpAll" {
						return true
					}
					if d.Get("backup_method").(string) == "Logical" {
						_, ok := d.GetOk("backup_type")
						if !ok {
							return false
						}
					}
					return true
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"db_name": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "Specify the database that needs to be backed up.",
						},
						"table_names": {
							Type:     schema.TypeSet,
							Optional: true,
							ForceNew: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "Specify the tables to be backed up in the specified database." +
								" When this field is empty, it defaults to full database backup.",
						},
					},
				},
			},
		},
	}
	return resource
}

func resourceVolcengineRdsMysqlBackupCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsMysqlBackupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineRdsMysqlBackup())
	if err != nil {
		return fmt.Errorf("error on creating rds_mysql_backup %q, %s", d.Id(), err)
	}
	return resourceVolcengineRdsMysqlBackupRead(d, meta)
}

func resourceVolcengineRdsMysqlBackupRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsMysqlBackupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineRdsMysqlBackup())
	if err != nil {
		return fmt.Errorf("error on reading rds_mysql_backup %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineRdsMysqlBackupDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsMysqlBackupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineRdsMysqlBackup())
	if err != nil {
		return fmt.Errorf("error on deleting rds_mysql_backup %q, %s", d.Id(), err)
	}
	return err
}
