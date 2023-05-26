package allow_list_associate

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Redis AllowList Association can be imported using the instanceId:allowListId, e.g.
```
$ terraform import volcengine_redis_allow_list_associate.default redis-asdljioeixxxx:acl-cn03wk541s55c376xxxx
```
*/

func ResourceVolcengineRedisAllowListAssociate() *schema.Resource {
	resource := &schema.Resource{
		Read:   resourceVolcengineRedisAllowListAssociateRead,
		Create: resourceVolcengineRedisAllowListAssociateCreate,
		Delete: resourceVolcengineRedisAllowListAssociateDelete,
		Importer: &schema.ResourceImporter{
			State: redisAllowListAssociateImporter,
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Id of instance to associate.",
			},
			"allow_list_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Id of allow list to associate.",
			},
		},
	}
	return resource
}

func resourceVolcengineRedisAllowListAssociateRead(d *schema.ResourceData, meta interface{}) (err error) {
	redisAllowListAssociateService := NewRedisAllowListAssociateService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(redisAllowListAssociateService, d, ResourceVolcengineRedisAllowListAssociate())
	if err != nil {
		return fmt.Errorf("error on reading association %v, %v", d.Id(), err)
	}
	return err
}

func resourceVolcengineRedisAllowListAssociateCreate(d *schema.ResourceData, meta interface{}) (err error) {
	redisAllowListAssociateService := NewRedisAllowListAssociateService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(redisAllowListAssociateService, d, ResourceVolcengineRedisAllowListAssociate())
	if err != nil {
		return fmt.Errorf("error on creating redis allow list association %v, %v", d.Id(), err)
	}
	return err
}

func resourceVolcengineRedisAllowListAssociateDelete(d *schema.ResourceData, meta interface{}) (err error) {
	redisAllowListAssociateService := NewRedisAllowListAssociateService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(redisAllowListAssociateService, d, ResourceVolcengineRedisAllowListAssociate())
	if err != nil {
		return fmt.Errorf("error on deleting redis allow list association %v, %v", d.Id(), err)
	}
	return err
}
