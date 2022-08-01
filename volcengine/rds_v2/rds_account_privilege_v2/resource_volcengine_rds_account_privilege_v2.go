package rds_account_privilege_v2

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	volc "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
RDS account privilege can be imported using the id, e.g.
```
$ terraform import volcengine_rds_account_privilege_v2.default mysql-42b38c769c4b:account_name
```

*/

func ResourceVolcengineRdsAccountPrivilege() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineRdsAccountPrivilegeCreate,
		Read:   resourceVolcengineRdsAccountPrivilegeRead,
		Update: resourceVolcengineRdsAccountPrivilegeUpdate,
		Delete: resourceVolcengineRdsAccountPrivilegeDelete,
		Importer: &schema.ResourceImporter{
			State: rdsAccountPrivilegeImporter,
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
			"db_privileges": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "The privileges of the account.",
				Set:         RdsAccountPrivilegeHash,
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
						"account_privilege_str": {
							Type:             schema.TypeString,
							Optional:         true,
							Computed:         true,
							Description:      "The privilege string of the account.",
							DiffSuppressFunc: RdsAccountPrivilegeStrDiffSuppress,
						},
					},
				},
			},
		},
	}
}

func resourceVolcengineRdsAccountPrivilegeCreate(d *schema.ResourceData, meta interface{}) (err error) {
	rdsAccountPrivilegeService := NewRdsAccountPrivilegeService(meta.(*volc.SdkClient))
	err = rdsAccountPrivilegeService.Dispatcher.Create(rdsAccountPrivilegeService, d, ResourceVolcengineRdsAccountPrivilege())
	if err != nil {
		return fmt.Errorf("error on creating rds account privilege %q, %w", d.Id(), err)
	}
	return resourceVolcengineRdsAccountPrivilegeRead(d, meta)
}

func resourceVolcengineRdsAccountPrivilegeRead(d *schema.ResourceData, meta interface{}) (err error) {
	rdsAccountPrivilegeService := NewRdsAccountPrivilegeService(meta.(*volc.SdkClient))
	err = rdsAccountPrivilegeService.Dispatcher.Read(rdsAccountPrivilegeService, d, ResourceVolcengineRdsAccountPrivilege())
	if err != nil {
		return fmt.Errorf("error on reading rds account privilege %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineRdsAccountPrivilegeUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	rdsAccountPrivilegeService := NewRdsAccountPrivilegeService(meta.(*volc.SdkClient))
	err = rdsAccountPrivilegeService.Dispatcher.Update(rdsAccountPrivilegeService, d, ResourceVolcengineRdsAccountPrivilege())
	if err != nil {
		return fmt.Errorf("error on updating rds account privilege %q, %w", d.Id(), err)
	}
	return resourceVolcengineRdsAccountPrivilegeRead(d, meta)
}

func resourceVolcengineRdsAccountPrivilegeDelete(d *schema.ResourceData, meta interface{}) (err error) {
	rdsAccountPrivilegeService := NewRdsAccountPrivilegeService(meta.(*volc.SdkClient))
	err = rdsAccountPrivilegeService.Dispatcher.Delete(rdsAccountPrivilegeService, d, ResourceVolcengineRdsAccountPrivilege())
	if err != nil {
		return fmt.Errorf("error on deleting rds account privilege %q, %w", d.Id(), err)
	}
	return err
}
