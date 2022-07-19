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
$ terraform import volcengine_rds_database.default mysql-42b38c769c4b:dbname
```

*/

func ResourceVolcengineRdsDatabase() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineRdsDatabaseCreate,
		Read:   resourceVolcengineRdsDatabaseRead,
		Delete: resourceVolcengineRdsDatabaseDelete,
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
				Description: "The name of the RDS database.",
			},
			"character_set_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The character set of the RDS database.",
			},
		},
	}
}

func resourceVolcengineRdsDatabaseCreate(d *schema.ResourceData, meta interface{}) (err error) {
	databaseService := NewRdsDatabaseService(meta.(*volc.SdkClient))
	err = databaseService.Dispatcher.Create(databaseService, d, ResourceVolcengineRdsDatabase())
	if err != nil {
		return fmt.Errorf("error on creating database %q, %w", d.Id(), err)
	}
	return resourceVolcengineRdsDatabaseRead(d, meta)
}

func resourceVolcengineRdsDatabaseRead(d *schema.ResourceData, meta interface{}) (err error) {
	databaseService := NewRdsDatabaseService(meta.(*volc.SdkClient))
	err = databaseService.Dispatcher.Read(databaseService, d, ResourceVolcengineRdsDatabase())
	if err != nil {
		return fmt.Errorf("error on reading database %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineRdsDatabaseDelete(d *schema.ResourceData, meta interface{}) (err error) {
	databaseService := NewRdsDatabaseService(meta.(*volc.SdkClient))
	err = databaseService.Dispatcher.Delete(databaseService, d, ResourceVolcengineRdsDatabase())
	if err != nil {
		return fmt.Errorf("error on deleting database %q, %w", d.Id(), err)
	}
	return err
}
