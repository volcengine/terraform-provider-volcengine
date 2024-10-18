package vedb_mysql_account

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
VedbMysqlAccount can be imported using the instance id and account name, e.g.
```
$ terraform import volcengine_vedb_mysql_account.default vedbm-r3xq0zdl****:testuser

```

*/

func ResourceVolcengineVedbMysqlAccount() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVedbMysqlAccountCreate,
		Read:   resourceVolcengineVedbMysqlAccountRead,
		Update: resourceVolcengineVedbMysqlAccountUpdate,
		Delete: resourceVolcengineVedbMysqlAccountDelete,
		Importer: &schema.ResourceImporter{
			State: veDBMysqlAccountImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				ForceNew:    true,
				Required:    true,
				Description: "The id of the instance.",
			},
			"account_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: "Database account name. " +
					"The account name must meet the following requirements:\n " +
					"The name is unique and within 2 to 32 characters in length.\n " +
					"Consists of lowercase letters, numbers, or underscores (_).\n " +
					"Starts with a lowercase letter and ends with a letter or number.\n " +
					"The name cannot contain certain prohibited words. " +
					"For detailed information, please refer to prohibited keywords. " +
					"And certain reserved words such as root, admin, etc. cannot be used.",
			},
			"account_password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
				Description: "Password of database account. " +
					"The account password must meet the following requirements:\n " +
					"It can only contain upper and lower case letters, numbers and the following special characters _#!@$%^&*()+=-. " +
					"\nIt must be within 8 to 32 characters in length.\n " +
					"It must contain at least three of upper case letters, lower case letters, numbers or special characters.",
			},
			"account_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: "Database account type. Values: " +
					"\nSuper: High-privilege account. " +
					"Only one high-privilege account can be created for an instance." +
					" It has all permissions for all databases under this instance and can manage all ordinary accounts and databases. " +
					"\nNormal: Multiple ordinary accounts can be created for an instance. " +
					"Specific database permissions need to be manually granted to ordinary accounts.",
			},
			"account_privileges": {
				Type:     schema.TypeSet,
				Optional: true,
				Set:      veDBMysqlAccountPrivilegeHash,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if d.Get("account_type") == "Super" {
						return true
					}
					return false
				},
				Description: "Database permission information. " +
					"When the value of AccountType is Super, this parameter does not need to be passed." +
					" High-privilege accounts by default have all permissions for all databases under this instance. " +
					"When the value of AccountType is Normal, " +
					"it is recommended to pass this parameter to grant specified permissions for specified databases to ordinary accounts. " +
					"If not set, this account does not have any permissions for any database.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"db_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Database name requiring authorization.",
						},
						"account_privilege": {
							Type:     schema.TypeString,
							Required: true,
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
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: veDBMysqlAccountPrivilegeStrDiffSuppress,
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

func resourceVolcengineVedbMysqlAccountCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVedbMysqlAccountService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineVedbMysqlAccount())
	if err != nil {
		return fmt.Errorf("error on creating vedb_mysql_account %q, %s", d.Id(), err)
	}
	return resourceVolcengineVedbMysqlAccountRead(d, meta)
}

func resourceVolcengineVedbMysqlAccountRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVedbMysqlAccountService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineVedbMysqlAccount())
	if err != nil {
		return fmt.Errorf("error on reading vedb_mysql_account %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineVedbMysqlAccountUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVedbMysqlAccountService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineVedbMysqlAccount())
	if err != nil {
		return fmt.Errorf("error on updating vedb_mysql_account %q, %s", d.Id(), err)
	}
	return resourceVolcengineVedbMysqlAccountRead(d, meta)
}

func resourceVolcengineVedbMysqlAccountDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVedbMysqlAccountService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineVedbMysqlAccount())
	if err != nil {
		return fmt.Errorf("error on deleting vedb_mysql_account %q, %s", d.Id(), err)
	}
	return err
}
