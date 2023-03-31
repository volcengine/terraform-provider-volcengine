package rds_database

import (
	"fmt"
	"time"

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
				Description: "Name database.\nillustrate:\nUnique name.\nThe length is 2~64 characters.\nStart with a letter and end with a letter or number.\nConsists of lowercase letters, numbers, and underscores (_) or dashes (-).\nDatabase names are disabled [keywords](https://www.volcengine.com/docs/6313/66162).",
			},
			"character_set_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Database character set. Currently supported character sets include: utf8, utf8mb4, latin1, ascii.",
			},
		},
	}
}

func resourceVolcengineRdsDatabaseCreate(d *schema.ResourceData, meta interface{}) (err error) {
	databaseService := NewRdsDatabaseService(meta.(*volc.SdkClient))
	err = volc.DefaultDispatcher().Create(databaseService, d, ResourceVolcengineRdsDatabase())
	if err != nil {
		return fmt.Errorf("error on creating database %q, %w", d.Id(), err)
	}
	return resourceVolcengineRdsDatabaseRead(d, meta)
}

func resourceVolcengineRdsDatabaseRead(d *schema.ResourceData, meta interface{}) (err error) {
	databaseService := NewRdsDatabaseService(meta.(*volc.SdkClient))
	err = volc.DefaultDispatcher().Read(databaseService, d, ResourceVolcengineRdsDatabase())
	if err != nil {
		return fmt.Errorf("error on reading database %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineRdsDatabaseDelete(d *schema.ResourceData, meta interface{}) (err error) {
	databaseService := NewRdsDatabaseService(meta.(*volc.SdkClient))
	err = volc.DefaultDispatcher().Delete(databaseService, d, ResourceVolcengineRdsDatabase())
	if err != nil {
		return fmt.Errorf("error on deleting database %q, %w", d.Id(), err)
	}
	return err
}
