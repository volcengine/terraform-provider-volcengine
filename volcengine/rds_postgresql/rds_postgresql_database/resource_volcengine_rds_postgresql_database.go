package rds_postgresql_database

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	volc "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Database can be imported using the instanceId:dbName, e.g.
```
$ terraform import volcengine_rds_postgresql_database.default postgres-ca7b7019****:dbname
```

*/

func ResourceVolcengineRdsPostgresqlDatabase() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineRdsPostgresqlDatabaseCreate,
		Read:   resourceVolcengineRdsPostgresqlDatabaseRead,
		Delete: resourceVolcengineRdsPostgresqlDatabaseDelete,
		Importer: &schema.ResourceImporter{
			State: func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
				items := strings.Split(data.Id(), ":")
				if len(items) != 2 {
					return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
				}
				if err := data.Set("instance_id", items[0]); err != nil {
					return []*schema.ResourceData{data}, err
				}
				if err := data.Set("db_name", items[1]); err != nil {
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
				Description: "The ID of the RDS instance.",
			},
			"db_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of database.",
			},
			"character_set_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "Database character set. Currently supported character sets include: utf8, latin1, ascii. Default is utf8.",
			},
			"collate": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The collate of database. Sorting rules. Value range: C (default), C.UTF-8, en_US.utf8, zh_CN.utf8 and POSIX.",
			},
			"c_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "Character classification. Value range: C (default), C.UTF-8, en_US.utf8, zh_CN.utf8, and POSIX.",
			},
			"owner": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The owner of database.",
			},
			"db_status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the RDS database.",
			},
		},
	}
}

func resourceVolcengineRdsPostgresqlDatabaseCreate(d *schema.ResourceData, meta interface{}) (err error) {
	databaseService := NewRdsPostgresqlDatabaseService(meta.(*volc.SdkClient))
	err = databaseService.Dispatcher.Create(databaseService, d, ResourceVolcengineRdsPostgresqlDatabase())
	if err != nil {
		return fmt.Errorf("error on creating postgresql database %q, %w", d.Id(), err)
	}
	return resourceVolcengineRdsPostgresqlDatabaseRead(d, meta)
}

func resourceVolcengineRdsPostgresqlDatabaseRead(d *schema.ResourceData, meta interface{}) (err error) {
	databaseService := NewRdsPostgresqlDatabaseService(meta.(*volc.SdkClient))
	err = databaseService.Dispatcher.Read(databaseService, d, ResourceVolcengineRdsPostgresqlDatabase())
	if err != nil {
		return fmt.Errorf("error on reading postgresql database %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineRdsPostgresqlDatabaseDelete(d *schema.ResourceData, meta interface{}) (err error) {
	databaseService := NewRdsPostgresqlDatabaseService(meta.(*volc.SdkClient))
	err = databaseService.Dispatcher.Delete(databaseService, d, ResourceVolcengineRdsPostgresqlDatabase())
	if err != nil {
		return fmt.Errorf("error on deleting postgresql database %q, %w", d.Id(), err)
	}
	return err
}
