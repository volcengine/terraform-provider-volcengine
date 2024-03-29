package allow_list_associate

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
mongodb allow list associate can be imported using the instanceId:allowListId, e.g.
```
$ terraform import volcengine_mongodb_allow_list_associate.default mongo-replica-e405f8e2****:acl-d1fd76693bd54e658912e7337d5b****
```

*/

func ResourceVolcengineMongodbAllowListAssociate() *schema.Resource {
	resource := &schema.Resource{
		Read:   resourceVolcengineMongodbAllowListAssociateRead,
		Create: resourceVolcengineMongodbAllowListAssociateCreate,
		Delete: resourceVolcengineMongodbAllowListAssociateDelete,
		Importer: &schema.ResourceImporter{
			State: mongodbAllowListAssociateImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
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

func resourceVolcengineMongodbAllowListAssociateRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewMongodbAllowListAssociateService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(service, d, ResourceVolcengineMongodbAllowListAssociate())
	if err != nil {
		return fmt.Errorf("error on reading mongodb allow list association %v, %v", d.Id(), err)
	}
	return err
}

func resourceVolcengineMongodbAllowListAssociateCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewMongodbAllowListAssociateService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(service, d, ResourceVolcengineMongodbAllowListAssociate())
	if err != nil {
		return fmt.Errorf("error on creating mongodb allow list association %v, %v", d.Id(), err)
	}
	return err
}

func resourceVolcengineMongodbAllowListAssociateDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewMongodbAllowListAssociateService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineMongodbAllowListAssociate())
	if err != nil {
		return fmt.Errorf("error on deleting mongodb allow list association %v, %v", d.Id(), err)
	}
	return err
}
