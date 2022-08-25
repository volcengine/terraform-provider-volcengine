package allow_list_associate

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func ResourceVolcengineMongodbAllowListAssociate() *schema.Resource {
	resource := &schema.Resource{
		Read:   resourceVolcengineMongodbAllowListAssociateRead,
		Create: resourceVolcengineMongodbAllowListAssociateCreate,
		Delete: resourceVolcengineMongodbAllowListAssociateDelete,
		Importer: &schema.ResourceImporter{
			State: mongodbAllowListAssociateImporter,
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
	err = service.Dispatcher.Read(service, d, ResourceVolcengineMongodbAllowListAssociate())
	if err != nil {
		return fmt.Errorf("error on reading association %v, %v", d.Id(), err)
	}
	return err
}

func resourceVolcengineMongodbAllowListAssociateCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewMongodbAllowListAssociateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineMongodbAllowListAssociate())
	if err != nil {
		return fmt.Errorf("error on creating mongodb allow list association %v, %v", d.Id(), err)
	}
	return err
}

func resourceVolcengineMongodbAllowListAssociateDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewMongodbAllowListAssociateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineMongodbAllowListAssociate())
	if err != nil {
		return fmt.Errorf("error on deleting mongodb allow list association %v, %v", d.Id(), err)
	}
	return err
}
