package rds_postgresql_data_backup

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
RdsPostgresqlDataBackup can be imported using the id, e.g.
```
$ terraform import volcengine_rds_postgresql_data_backup.default resource_id
```

*/

func ResourceVolcengineRdsPostgresqlDataBackup() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineRdsPostgresqlDataBackupCreate,
		Read:   resourceVolcengineRdsPostgresqlDataBackupRead,
		Delete: resourceVolcengineRdsPostgresqlDataBackupDelete,
		Importer: &schema.ResourceImporter{State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
			parts := strings.Split(d.Id(), ":")
			if len(parts) < 2 {
				return []*schema.ResourceData{d}, fmt.Errorf("import id must be 'instance_id:backup_id'")
			}
			_ = d.Set("instance_id", parts[0])
			_ = d.Set("backup_id", parts[1])
			return []*schema.ResourceData{d}, nil
		}},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the PostgreSQL instance.",
			},
			"backup_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The id of the backup.",
			},
			"backup_scope": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"Instance", "Database"}, false),
				Description:  "The scope of the backup: Instance, Database.",
			},
			"backup_method": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"Physical", "Logical"}, false),
				Description: "The method of the backup: Physical, Logical." +
					"When the value of backup_scope is Database, the value of backup_method can only be Logical.",
			},
			"backup_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"Full", "Increment"}, false),
				Description: "The backup type of the backup: Full(default), Increment. " +
					"Do not set this parameter when backup_method is Logical; otherwise, the creation will fail.",
			},
			"backup_description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The description of the backup set.",
			},
			"backup_meta": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"db_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The name of the database.",
						},
					},
				},
				Description: "Specify the database that needs to be backed up. " +
					"This parameter can only be set when the value of backup_scope is Database.",
			},
		},
	}
	return resource
}

func resourceVolcengineRdsPostgresqlDataBackupCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlDataBackupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineRdsPostgresqlDataBackup())
	if err != nil {
		return fmt.Errorf("error on creating rds_postgresql_data_backup %q, %s", d.Id(), err)
	}
	return resourceVolcengineRdsPostgresqlDataBackupRead(d, meta)
}

func resourceVolcengineRdsPostgresqlDataBackupRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlDataBackupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineRdsPostgresqlDataBackup())
	if err != nil {
		if ve.ResourceNotFoundError(err) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("error on reading rds_postgresql_data_backup %q, %s", d.Id(), err)
	}
	return nil
}

func resourceVolcengineRdsPostgresqlDataBackupDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlDataBackupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineRdsPostgresqlDataBackup())
	if err != nil {
		return fmt.Errorf("error on deleting rds_postgresql_data_backup %q, %s", d.Id(), err)
	}
	return err
}
