package rds_postgresql_account

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
RDS postgresql account can be imported using the instance_id:account_name, e.g.
```
$ terraform import volcengine_rds_postgresql_account.default postgres-ca7b7019****:accountName
```

*/

func ResourceVolcengineRdsPostgresqlAccount() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineRdsPostgresqlAccountCreate,
		Read:   resourceVolcengineRdsPostgresqlAccountRead,
		Update: resourceVolcengineRdsPostgresqlAccountUpdate,
		Delete: resourceVolcengineRdsPostgresqlAccountDelete,
		Importer: &schema.ResourceImporter{
			State: func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
				items := strings.Split(data.Id(), ":")
				if len(items) != 2 {
					return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
				}
				if err := data.Set("instance_id", items[0]); err != nil {
					return []*schema.ResourceData{data}, err
				}
				if err := data.Set("account_name", items[1]); err != nil {
					return []*schema.ResourceData{data}, err
				}
				return []*schema.ResourceData{data}, nil
			},
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
				Description: "Database account name.",
			},
			"account_password": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "The password of the database account. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
			},
			"account_type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				Description:  "Database account type, value:\nSuper: A high-privilege account. Only one database account can be created for an instance.\nNormal: An account with ordinary privileges.",
				ValidateFunc: validation.StringInSlice([]string{"Super", "Normal"}, false),
			},
			"account_status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the database account.",
			},
			"account_privileges": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					oldS := strings.Split(old, ",")
					sort.Strings(oldS)
					newS := strings.Split(new, ",")
					sort.Strings(newS)
					return reflect.DeepEqual(oldS, newS)
				},
				Description: "The privilege information of account. " +
					"When the account type is a super account, there is no need to pass in this parameter, and all privileges are supported by default. " +
					"When the account type is a normal account, this parameter can be passed in, " +
					"the default values are Login and Inherit.",
			},
		},
	}
}

func resourceVolcengineRdsPostgresqlAccountCreate(d *schema.ResourceData, meta interface{}) (err error) {
	rdsAccountService := NewRdsPostgresqlAccountService(meta.(*volc.SdkClient))
	err = rdsAccountService.Dispatcher.Create(rdsAccountService, d, ResourceVolcengineRdsPostgresqlAccount())
	if err != nil {
		return fmt.Errorf("error on creating rds postgresql account %q, %w", d.Id(), err)
	}
	return resourceVolcengineRdsPostgresqlAccountRead(d, meta)
}

func resourceVolcengineRdsPostgresqlAccountRead(d *schema.ResourceData, meta interface{}) (err error) {
	rdsAccountService := NewRdsPostgresqlAccountService(meta.(*volc.SdkClient))
	err = rdsAccountService.Dispatcher.Read(rdsAccountService, d, ResourceVolcengineRdsPostgresqlAccount())
	if err != nil {
		return fmt.Errorf("error on reading rds postgresql account %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineRdsPostgresqlAccountUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRdsPostgresqlAccountService(meta.(*volc.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineRdsPostgresqlAccount())
	if err != nil {
		return fmt.Errorf("error on updating rds postgresql account  %q, %s", d.Id(), err)
	}
	return resourceVolcengineRdsPostgresqlAccountRead(d, meta)
}

func resourceVolcengineRdsPostgresqlAccountDelete(d *schema.ResourceData, meta interface{}) (err error) {
	rdsAccountService := NewRdsPostgresqlAccountService(meta.(*volc.SdkClient))
	err = rdsAccountService.Dispatcher.Delete(rdsAccountService, d, ResourceVolcengineRdsPostgresqlAccount())
	if err != nil {
		return fmt.Errorf("error on deleting rds postgresql account %q, %w", d.Id(), err)
	}
	return err
}
