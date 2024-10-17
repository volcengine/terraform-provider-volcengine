package vedb_mysql_database

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
			"databases_privileges": {
				Type:     schema.TypeSet,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_name": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "Account name that requires authorization.",
						},
						"account_privilege": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
							Description: "Authorization database privilege types: " +
								"\nReadWrite: Read and write privilege.\n " +
								"ReadOnly: Read-only privilege.\n " +
								"DDLOnly: Only DDL privilege.\n " +
								"DMLOnly: Only DML privilege.\n " +
								"Custom: Custom privilege.",
						},
						// 看下非custom的情况下能不能传detail，如果接口允许传那么得在before call里拦截一下，要不闭环不了
						/*
							在 DescribeDatabases 接口中作为返回参数时，无论 AccountPrivilege 取什么值，都返回该权限类型所包含的 SQL 操作权限详情。
						*/
						"account_privilege_detail": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
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
							Description: "The specific SQL operation permissions contained in the permission type are separated by English commas (,) between multiple strings.\n " +
								"When used as a request parameter in the CreateDatabase interface, " +
								"when the AccountPrivilege value is Custom, this parameter is required. " +
								"Value range (multiple selections allowed): SELECT, INSERT, UPDATE, DELETE," +
								" CREATE, DROP, REFERENCES, INDEX, ALTER, CREATE TEMPORARY TABLES, LOCK TABLES, " +
								"EXECUTE, CREATE VIEW, SHOW VIEW, CREATE ROUTINE, ALTER ROUTINE, EVENT, TRIGGER. " +
								"When used as a return parameter in the DescribeDatabases interface, " +
								"regardless of the value of AccountPrivilege, the details of the SQL operation permissions contained in this permission type are returned. " +
								"For the specific SQL operation permissions contained in each permission type, " +
								"please refer to the account permission list.",
						},
					},
				},
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
