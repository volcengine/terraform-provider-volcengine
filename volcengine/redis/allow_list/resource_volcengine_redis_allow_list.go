package allow_list

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Redis AllowList can be imported using the id, e.g.
```
$ terraform import volcengine_redis_allow_list.default acl-cn03wk541s55c376xxxx
```

*/

func ResourceVolcengineRedisAllowList() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineRedisAllowListCreate,
		Read:   resourceVolcengineRedisAllowListRead,
		Update: resourceVolcengineRedisAllowListUpdate,
		Delete: resourceVolcengineRedisAllowListDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"allow_list_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of allow list.",
			},
			"allow_list_desc": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of allow list.",
			},
			"allow_list": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "Ip list of allow list.",
			},
		},
	}
	ve.MergeDateSourceToResource(DataSourceVolcengineRedisAllowLists().Schema["allow_lists"].Elem.(*schema.Resource).Schema, &resource.Schema)
	return resource
}

func resourceVolcengineRedisAllowListCreate(d *schema.ResourceData, meta interface{}) (err error) {
	redisAllowListService := NewRedisAllowListService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(redisAllowListService, d, ResourceVolcengineRedisAllowList())
	if err != nil {
		return fmt.Errorf("error on creating redis allowlist %v, %v", d.Id(), err)
	}
	return resourceVolcengineRedisAllowListRead(d, meta)
}

func resourceVolcengineRedisAllowListUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	redisAllowListService := NewRedisAllowListService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(redisAllowListService, d, ResourceVolcengineRedisAllowList())
	if err != nil {
		return fmt.Errorf("error on updating redis allowlist  %q, %s", d.Id(), err)
	}
	return resourceVolcengineRedisAllowListRead(d, meta)
}

func resourceVolcengineRedisAllowListDelete(d *schema.ResourceData, meta interface{}) (err error) {
	redisAllowListService := NewRedisAllowListService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(redisAllowListService, d, ResourceVolcengineRedisAllowList())
	if err != nil {
		return fmt.Errorf("error on deleting redis allowlist %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineRedisAllowListRead(d *schema.ResourceData, meta interface{}) (err error) {
	redisAllowListService := NewRedisAllowListService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(redisAllowListService, d, ResourceVolcengineRedisAllowList())
	if err != nil {
		return fmt.Errorf("error on reading redis allowlist %q,%s", d.Id(), err)
	}
	return err
}
