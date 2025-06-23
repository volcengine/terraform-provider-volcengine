package rds_mysql_database

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	volc "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Database can be imported using the instanceId:dbName, e.g.
```
$ terraform import volcengine_rds_mysql_database.default mysql-42b38c769c4b:dbname
```

*/

func ResourceVolcengineRdsMysqlDatabase() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineRdsMysqlDatabaseCreate,
		Read:   resourceVolcengineRdsMysqlDatabaseRead,
		Update: resourceVolcengineRdsMysqlDatabaseUpdate,
		Delete: resourceVolcengineRdsMysqlDatabaseDelete,
		Importer: &schema.ResourceImporter{
			State: databaseImporter,
		},
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
				Optional:    true,
				Default:     "utf8mb4",
				ForceNew:    true,
				Description: "Database character set. Currently supported character sets include: utf8, utf8mb4, latin1, ascii.",
			},
			"database_privileges": {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Set:         RdsMysqlDatabasePrivilegeHash,
				Description: "Authorization database privilege information.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Account name that requires authorization.",
						},
						"host": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							Description: "The IP address of the database that the specified database account can access. " +
								"The default value is %.\nIf the Host is specified as %, the account is allowed to access the database from any IP address. " +
								"If the Host is specified as 192.10.10.%, " +
								"it means the account can access the database from IP addresses between 192.10.10.0 and 192.10.10.255. " +
								"The specified Host needs to be added to the whitelist bound to the instance.",
						},
						"account_privilege": {
							Type:     schema.TypeString,
							Required: true,
							Description: "The types of account permissions granted, with the following values: " +
								"ReadWrite: Read and write permissions. " +
								"ReadOnly: Read-only permissions. " +
								"DDLOnly: Only DDL permissions. " +
								"DMLOnly: Only DML permissions. " +
								"Custom: Custom permissions.",
						},
						"account_privilege_detail": {
							Type:     schema.TypeString,
							Optional: true,
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
							Description: "The specific permissions granted to the account. " +
								"For example, if the AccountPrivileges value is ReadWrite, " +
								"the AccountPrivilegeDetail value can be SELECT, INSERT, UPDATE, " +
								"DELETE, CREATE, DROP, ALTER, INDEX, CREATE VIEW, SHOW VIEW, CREATE ROUTINE, ALTER ROUTINE, EXECUTE, REPLICATION CLIENT," +
								" CREATE TEMPORARY TABLES, LOCK TABLES, CREATE USER, EVENT, TRIGGER, and so on. " +
								"When used as a return result, regardless of whether the value of AccountPrivilege is Custom, the detailed permissions of AccountPrivilege will be displayed. " +
								"Instructions: Multiple strings are separated by English commas (,).",
						},
					},
				},
			},
			"db_desc": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "The description information of the database, with a length not exceeding 256 characters." +
					" This field is optional." +
					" If this field is not set, or if this field is set but the length of the description information is 0, " +
					"then the description information is empty.",
			},
		},
	}
}

func resourceVolcengineRdsMysqlDatabaseCreate(d *schema.ResourceData, meta interface{}) (err error) {
	databaseService := NewRdsMysqlDatabaseService(meta.(*volc.SdkClient))
	err = databaseService.Dispatcher.Create(databaseService, d, ResourceVolcengineRdsMysqlDatabase())
	if err != nil {
		return fmt.Errorf("error on creating database %q, %w", d.Id(), err)
	}
	return resourceVolcengineRdsMysqlDatabaseRead(d, meta)
}

func resourceVolcengineRdsMysqlDatabaseRead(d *schema.ResourceData, meta interface{}) (err error) {
	databaseService := NewRdsMysqlDatabaseService(meta.(*volc.SdkClient))
	err = databaseService.Dispatcher.Read(databaseService, d, ResourceVolcengineRdsMysqlDatabase())
	if err != nil {
		return fmt.Errorf("error on reading database %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineRdsMysqlDatabaseUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	databaseService := NewRdsMysqlDatabaseService(meta.(*volc.SdkClient))
	err = databaseService.Dispatcher.Update(databaseService, d, ResourceVolcengineRdsMysqlDatabase())
	if err != nil {
		return fmt.Errorf("error on updating database %q, %w", d.Id(), err)
	}
	return resourceVolcengineRdsMysqlDatabaseRead(d, meta)
}

func resourceVolcengineRdsMysqlDatabaseDelete(d *schema.ResourceData, meta interface{}) (err error) {
	databaseService := NewRdsMysqlDatabaseService(meta.(*volc.SdkClient))
	err = databaseService.Dispatcher.Delete(databaseService, d, ResourceVolcengineRdsMysqlDatabase())
	if err != nil {
		return fmt.Errorf("error on deleting database %q, %w", d.Id(), err)
	}
	return err
}
