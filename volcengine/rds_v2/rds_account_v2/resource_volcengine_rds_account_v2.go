package rds_account_v2

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	volc "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
RDS account can be imported using the id, e.g.
```
$ terraform import volcengine_rds_account.default mysql-42b38c769c4b:test
```

*/

func ResourceVolcengineRdsAccount() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineRdsAccountCreate,
		Read:   resourceVolcengineRdsAccountRead,
		Delete: resourceVolcengineRdsAccountDelete,
		Importer: &schema.ResourceImporter{
			State: rdsAccountImporter,
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
				ForceNew:    true,
				Description: "The password of the database account.\nillustrate\nCannot start with `!` or `@`.\nThe length is 8~32 characters.\nIt consists of any three of uppercase letters, lowercase letters, numbers, and special characters.\nThe special characters are `!@#$%^*()_+-=`.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Id() != ""
				},
			},
			"account_type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				Description:  "Database account type, value:\nSuper: A high-privilege account. Only one database account can be created for an instance.\nNormal: An account with ordinary privileges.",
				ValidateFunc: validation.StringInSlice([]string{"Super", "Normal"}, false),
			},
		},
	}
}

func resourceVolcengineRdsAccountCreate(d *schema.ResourceData, meta interface{}) (err error) {
	rdsAccountService := NewRdsAccountService(meta.(*volc.SdkClient))
	err = rdsAccountService.Dispatcher.Create(rdsAccountService, d, ResourceVolcengineRdsAccount())
	if err != nil {
		return fmt.Errorf("error on creating rds account %q, %w", d.Id(), err)
	}
	return resourceVolcengineRdsAccountRead(d, meta)
}

func resourceVolcengineRdsAccountRead(d *schema.ResourceData, meta interface{}) (err error) {
	rdsAccountService := NewRdsAccountService(meta.(*volc.SdkClient))
	err = rdsAccountService.Dispatcher.Read(rdsAccountService, d, ResourceVolcengineRdsAccount())
	if err != nil {
		return fmt.Errorf("error on reading rds account %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineRdsAccountDelete(d *schema.ResourceData, meta interface{}) (err error) {
	rdsAccountService := NewRdsAccountService(meta.(*volc.SdkClient))
	err = rdsAccountService.Dispatcher.Delete(rdsAccountService, d, ResourceVolcengineRdsAccount())
	if err != nil {
		return fmt.Errorf("error on deleting rds account %q, %w", d.Id(), err)
	}
	return err
}
