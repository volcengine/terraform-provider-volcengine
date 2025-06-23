package rds_mysql_account

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
RDS mysql account can be imported using the instance_id:account_name, e.g.
```
$ terraform import volcengine_rds_mysql_account.default mysql-42b38c769c4b:test
```

*/

func ResourceVolcengineRdsMysqlAccount() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineRdsMysqlAccountCreate,
		Read:   resourceVolcengineRdsMysqlAccountRead,
		Update: resourceVolcengineRdsMysqlAccountUpdate,
		Delete: resourceVolcengineRdsMysqlAccountDelete,
		Importer: &schema.ResourceImporter{
			State: rdsMysqlAccountImporter,
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
			"account_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Database account name. The rules are as follows:\nUnique name.\nStart with a letter and end with a letter or number.\nConsists of lowercase letters, numbers, or underscores (_).\nThe length is 2~32 characters.\nThe [keyword list](https://www.volcengine.com/docs/6313/66162) is disabled for database accounts, and certain reserved words, including root, admin, etc., cannot be used.",
			},
			"account_password": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "The password of the database account.\nIllustrate:\nCannot start with `!` or `@`.\nThe length is 8~32 characters.\nIt consists of any three of uppercase letters, lowercase letters, numbers, and special characters.\nThe special characters are `!@#$%^*()_+-=`. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
			},
			"account_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Database account type, value:\nSuper: A high-privilege account. Only one database account can be created for an instance.\nNormal: An account with ordinary privileges.",
			},
			"account_desc": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Account information description. The length should not exceed 256 characters.",
			},
			"host": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Description: "Specify the IP address for the account to access the database. " +
					"The default value is %. " +
					"If the Host is specified as %, the account is allowed to access the database from any IP address. " +
					"Wildcards are supported for setting the IP address range that can access the database. " +
					"For example, if the Host is specified as 192.10.10.%, it means the account can access the database from IP addresses between 192.10.10.0 and 192.10.10.255. " +
					"The specified Host needs to be added to the whitelist bound to the instance, otherwise the instance cannot be accessed normally. " +
					"The ModifyAllowList interface can be called to add the Host to the whitelist. " +
					"Note: If the created account type is a high-privilege account, the host IP can only be specified as %. " +
					"That is, when the value of AccountType is Super, the value of Host can only be %.",
			},
			"account_privileges": {
				Type:     schema.TypeSet,
				Optional: true,
				Set:      RdsMysqlAccountPrivilegeHash,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Get("account_type").(string) == "Super"
				},
				Description: "The privilege information of account.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"db_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The name of database.",
						},
						"account_privilege": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The privilege type of the account.",
						},
						"account_privilege_detail": {
							Type:     schema.TypeString,
							Optional: true,
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
							Description: "The privilege detail of the account.",
						},
					},
				},
			},
			"table_column_privileges": {
				Type:        schema.TypeSet,
				Optional:    true,
				ForceNew:    true,
				Description: "Settings for table column permissions of the account.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"db_name": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "Settings for table column permissions of the account.",
						},
						"table_privileges": {
							Type:        schema.TypeSet,
							Optional:    true,
							ForceNew:    true,
							Description: "Table permission information of the account.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"table_name": {
										Type:        schema.TypeString,
										Required:    true,
										ForceNew:    true,
										Description: "The name of the table for setting permissions on the account.",
									},
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
										Description: "Table privileges of the account.",
									},
								},
							},
						},
						"column_privileges": {
							Type:        schema.TypeSet,
							Optional:    true,
							ForceNew:    true,
							Description: "Column permission information of the account.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"column_name": {
										Type:        schema.TypeString,
										Required:    true,
										ForceNew:    true,
										Description: "The name of the column for setting permissions on the account.",
									},
									"table_name": {
										Type:        schema.TypeString,
										Required:    true,
										ForceNew:    true,
										Description: "The name of the table for setting permissions on the account.",
									},
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
										Description: "Table privileges of the account.",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceVolcengineRdsMysqlAccountCreate(d *schema.ResourceData, meta interface{}) (err error) {
	rdsAccountService := NewRdsMysqlAccountService(meta.(*volc.SdkClient))
	err = rdsAccountService.Dispatcher.Create(rdsAccountService, d, ResourceVolcengineRdsMysqlAccount())
	if err != nil {
		return fmt.Errorf("error on creating rds mysql account %q, %w", d.Id(), err)
	}
	return resourceVolcengineRdsMysqlAccountRead(d, meta)
}

func resourceVolcengineRdsMysqlAccountRead(d *schema.ResourceData, meta interface{}) (err error) {
	rdsAccountService := NewRdsMysqlAccountService(meta.(*volc.SdkClient))
	err = rdsAccountService.Dispatcher.Read(rdsAccountService, d, ResourceVolcengineRdsMysqlAccount())
	if err != nil {
		return fmt.Errorf("error on reading rds mysql account %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineRdsMysqlAccountUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsMysqlAccountService(meta.(*volc.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineRdsMysqlAccount())
	if err != nil {
		return fmt.Errorf("error on updating rds mysql account  %q, %s", d.Id(), err)
	}
	return resourceVolcengineRdsMysqlAccountRead(d, meta)
}

func resourceVolcengineRdsMysqlAccountDelete(d *schema.ResourceData, meta interface{}) (err error) {
	rdsAccountService := NewRdsMysqlAccountService(meta.(*volc.SdkClient))
	err = rdsAccountService.Dispatcher.Delete(rdsAccountService, d, ResourceVolcengineRdsMysqlAccount())
	if err != nil {
		return fmt.Errorf("error on deleting rds mysql account %q, %w", d.Id(), err)
	}
	return err
}
