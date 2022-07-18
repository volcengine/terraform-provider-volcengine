package rds_account

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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
		Update: resourceVolcengineRdsAccountUpdate,
		Delete: resourceVolcengineRdsAccountDelete,
		Importer: &schema.ResourceImporter{
			State: rdsAccountImporter,
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The ID of the RDS instance.",
			},
			"account_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the database account.",
			},
			"account_password": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The password of the database account.",
			},
			"account_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The type of the database account.",
			},
		},
	}
}

func resourceVolcengineRdsAccountCreate(d *schema.ResourceData, meta interface{}) (err error) {
	rdsAccountService := NewRdsAccountService(meta.(*volc.SdkClient))
	err = rdsAccountService.Dispatcher.Create(rdsAccountService, d, ResourceVolcengineRdsAccount())
	if err != nil {
		return fmt.Errorf("error on creating database account %q, %w", d.Id(), err)
	}
	return resourceVolcengineRdsAccountRead(d, meta)
}

func resourceVolcengineRdsAccountRead(d *schema.ResourceData, meta interface{}) (err error) {
	rdsAccountService := NewRdsAccountService(meta.(*volc.SdkClient))
	err = rdsAccountService.Dispatcher.Read(rdsAccountService, d, ResourceVolcengineRdsAccount())
	if err != nil {
		return fmt.Errorf("error on reading database account %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineRdsAccountUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	rdsAccountService := NewRdsAccountService(meta.(*volc.SdkClient))
	err = rdsAccountService.Dispatcher.Update(rdsAccountService, d, ResourceVolcengineRdsAccount())
	if err != nil {
		return fmt.Errorf("error on updating database account %q, %w", d.Id(), err)
	}
	return resourceVolcengineRdsAccountRead(d, meta)
}

func resourceVolcengineRdsAccountDelete(d *schema.ResourceData, meta interface{}) (err error) {
	rdsAccountService := NewRdsAccountService(meta.(*volc.SdkClient))
	err = rdsAccountService.Dispatcher.Delete(rdsAccountService, d, ResourceVolcengineRdsAccount())
	if err != nil {
		return fmt.Errorf("error on deleting database account %q, %w", d.Id(), err)
	}
	return err
}
