package account

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	volc "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Redis account can be imported using the instanceId:accountName, e.g.
```
$ terraform import volcengine_redis_account.default redis-42b38c769c4b:test
```

*/

var accountImporter = func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
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
}

func ResourceVolcengineRedisAccount() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineRedisAccountCreate,
		Read:   resourceVolcengineRedisAccountRead,
		Delete: resourceVolcengineRedisAccountDelete,
		Update: resourceVolcengineRedisAccountUpdate,
		Importer: &schema.ResourceImporter{
			State: accountImporter,
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the Redis instance.",
			},
			"account_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Redis account name.",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "The password of the redis account. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the redis account.",
			},
			"role_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Role type, the valid value can be `Administrator`, `ReadWrite`, `ReadOnly`, `NotDangerous`.",
			},
		},
	}
}

func resourceVolcengineRedisAccountCreate(d *schema.ResourceData, meta interface{}) (err error) {
	redisAccountService := NewAccountService(meta.(*volc.SdkClient))
	err = volc.DefaultDispatcher().Create(redisAccountService, d, ResourceVolcengineRedisAccount())
	if err != nil {
		return fmt.Errorf("error on creating redis account %q, %w", d.Id(), err)
	}
	return resourceVolcengineRedisAccountRead(d, meta)
}

func resourceVolcengineRedisAccountRead(d *schema.ResourceData, meta interface{}) (err error) {
	redisAccountService := NewAccountService(meta.(*volc.SdkClient))
	err = volc.DefaultDispatcher().Read(redisAccountService, d, ResourceVolcengineRedisAccount())
	if err != nil {
		return fmt.Errorf("error on reading redis account %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineRedisAccountUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	redisAccountService := NewAccountService(meta.(*volc.SdkClient))
	err = volc.DefaultDispatcher().Update(redisAccountService, d, ResourceVolcengineRedisAccount())
	if err != nil {
		return fmt.Errorf("error on update redis account %q, %w", d.Id(), err)
	}
	return err
}
func resourceVolcengineRedisAccountDelete(d *schema.ResourceData, meta interface{}) (err error) {
	redisAccountService := NewAccountService(meta.(*volc.SdkClient))
	err = volc.DefaultDispatcher().Delete(redisAccountService, d, ResourceVolcengineRedisAccount())
	if err != nil {
		return fmt.Errorf("error on deleting redis account %q, %w", d.Id(), err)
	}
	return err
}
