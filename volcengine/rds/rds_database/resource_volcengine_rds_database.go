package rds_database

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	volc "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Database can be imported using the id, e.g.
```
$ terraform import volcengine_database.default mysql-42b38c769c4b:dbname
```

*/

func ResourceVolcengineDatabase() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineDatabaseCreate,
		Read:   resourceVolcengineDatabaseRead,
		Delete: resourceVolcengineDatabaseDelete,
		Importer: &schema.ResourceImporter{
			State: databaseImporter,
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
				Description: "The name of the database.",
			},
			"character_set_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The character set of the database.",
			},
		},
	}
}

func resourceVolcengineDatabaseCreate(d *schema.ResourceData, meta interface{}) (err error) {
	databaseService := NewDatabaseService(meta.(*volc.SdkClient))
	err = databaseService.Dispatcher.Create(databaseService, d, ResourceVolcengineDatabase())
	if err != nil {
		return fmt.Errorf("error on creating database %q, %w", d.Id(), err)
	}
	return resourceVolcengineDatabaseRead(d, meta)
}

func resourceVolcengineDatabaseRead(d *schema.ResourceData, meta interface{}) (err error) {
	databaseService := NewDatabaseService(meta.(*volc.SdkClient))
	err = databaseService.Dispatcher.Read(databaseService, d, ResourceVolcengineDatabase())
	if err != nil {
		return fmt.Errorf("error on reading database %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineDatabaseDelete(d *schema.ResourceData, meta interface{}) (err error) {
	databaseService := NewDatabaseService(meta.(*volc.SdkClient))
	err = databaseService.Dispatcher.Delete(databaseService, d, ResourceVolcengineDatabase())
	if err != nil {
		return fmt.Errorf("error on deleting database %q, %w", d.Id(), err)
	}
	return err
}
