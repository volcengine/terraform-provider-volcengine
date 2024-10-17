package vedb_mysql_backup

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
VedbMysqlBackup can be imported using the id, e.g.
```
$ terraform import volcengine_vedb_mysql_backup.default resource_id
```

*/

func ResourceVolcengineVedbMysqlBackup() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVedbMysqlBackupCreate,
		Read:   resourceVolcengineVedbMysqlBackupRead,
		Update: resourceVolcengineVedbMysqlBackupUpdate,
		Delete: resourceVolcengineVedbMysqlBackupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
		    // TODO: Add all your arguments and attributes.
			"replace_with_arguments": {
				Type:     schema.TypeString,
				Optional: true,
			},
			// TODO: See setting, getting, flattening, expanding examples below for this complex argument.
			"complex_argument": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sub_field_one": {
							Type:         schema.TypeString,
							Required:     true,
						},
						"sub_field_two": {
							Type:     schema.TypeString,
							Optional: true,
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
