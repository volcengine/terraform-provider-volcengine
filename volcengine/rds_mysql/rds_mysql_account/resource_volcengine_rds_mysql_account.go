package rds_mysql_account

import (
	"fmt"
	"reflect"
	"sort"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
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
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				Description:  "Database account type, value:\nSuper: A high-privilege account. Only one database account can be created for an instance.\nNormal: An account with ordinary privileges.",
				ValidateFunc: validation.StringInSlice([]string{"Super", "Normal"}, false),
			},
			"account_privileges": {
				Type:        schema.TypeSet,
				Optional:    true,
				Set:         RdsMysqlAccountPrivilegeHash,
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
