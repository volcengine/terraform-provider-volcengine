package vedb_mysql_database

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
VedbMysqlDatabase can be imported using the instance id and database name, e.g.
```
$ terraform import volcengine_vedb_mysql_database.default vedbm-r3xq0zdl****:testdb

```

*/

func ResourceVolcengineVedbMysqlDatabase() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVedbMysqlDatabaseCreate,
		Read:   resourceVolcengineVedbMysqlDatabaseRead,
		Update: resourceVolcengineVedbMysqlDatabaseUpdate,
		Delete: resourceVolcengineVedbMysqlDatabaseDelete,
		Importer: &schema.ResourceImporter{
			State: vedbMysqlDatabaseImporter,
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
			"db_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: "The name of the database. Naming rules:\n " +
					"Unique name. Start with a lowercase letter and end with a letter or number. " +
					"The length is within 2 to 64 characters.\n " +
					"Consist of lowercase letters, numbers, underscores (_), or hyphens (-).\n " +
					"The name cannot contain certain reserved words.",
			},
			"character_set_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "Database character set: utf8mb4 (default), utf8, latin1, ascii.",
			},
		},
	}
	return resource
}

func resourceVolcengineVedbMysqlDatabaseCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVedbMysqlDatabaseService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineVedbMysqlDatabase())
	if err != nil {
		return fmt.Errorf("error on creating vedb_mysql_database %q, %s", d.Id(), err)
	}
	return resourceVolcengineVedbMysqlDatabaseRead(d, meta)
}

func resourceVolcengineVedbMysqlDatabaseRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVedbMysqlDatabaseService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineVedbMysqlDatabase())
	if err != nil {
		return fmt.Errorf("error on reading vedb_mysql_database %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineVedbMysqlDatabaseUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVedbMysqlDatabaseService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineVedbMysqlDatabase())
	if err != nil {
		return fmt.Errorf("error on updating vedb_mysql_database %q, %s", d.Id(), err)
	}
	return resourceVolcengineVedbMysqlDatabaseRead(d, meta)
}

func resourceVolcengineVedbMysqlDatabaseDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVedbMysqlDatabaseService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineVedbMysqlDatabase())
	if err != nil {
		return fmt.Errorf("error on deleting vedb_mysql_database %q, %s", d.Id(), err)
	}
	return err
}

var vedbMysqlDatabaseImporter = func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
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
}
