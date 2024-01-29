package sqlserver_backup

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Rds Mssql Backup can be imported using the id, e.g.
```
$ terraform import volcengine_rds_mssql_backup.default instanceId:backupId
```

*/

func ResourceVolcengineRdsMssqlBackup() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineRdsMssqlBackupCreate,
		Read:   resourceVolcengineRdsMssqlBackupRead,
		Delete: resourceVolcengineRdsMssqlBackupDelete,
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
			Create: schema.DefaultTimeout(60 * time.Minute),
			Delete: schema.DefaultTimeout(60 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the instance.",
			},
			"backup_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "Backup type. Currently only supports full backup, with a value of Full (default).",
			},
			"backup_meta": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Backup repository information. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"db_name": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "The name of the database.",
						},
					},
				},
			},
			"backup_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the backup.",
			},
		},
	}
	return resource
}

func resourceVolcengineRdsMssqlBackupCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsMssqlBackupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineRdsMssqlBackup())
	if err != nil {
		return fmt.Errorf("error on creating rds_mssql_backup %q, %s", d.Id(), err)
	}
	return resourceVolcengineRdsMssqlBackupRead(d, meta)
}

func resourceVolcengineRdsMssqlBackupRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsMssqlBackupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineRdsMssqlBackup())
	if err != nil {
		return fmt.Errorf("error on reading rds_mssql_backup %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineRdsMssqlBackupDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsMssqlBackupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineRdsMssqlBackup())
	if err != nil {
		return fmt.Errorf("error on deleting rds_mssql_backup %q, %s", d.Id(), err)
	}
	return err
}
