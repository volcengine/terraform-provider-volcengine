package rocketmq_allow_list_associate

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
RocketmqAllowListAssociate can be imported using the instance_id:allow_list_id, e.g.
```
$ terraform import volcengine_rocketmq_allow_list_associate.default resource_id
```

*/

func ResourceVolcengineRocketmqAllowListAssociate() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineRocketmqAllowListAssociateCreate,
		Read:   resourceVolcengineRocketmqAllowListAssociateRead,
		Delete: resourceVolcengineRocketmqAllowListAssociateDelete,
		Importer: &schema.ResourceImporter{
			State: allowListAssociateImporter,
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
				Description: "The id of the rocketmq instance.",
			},
			"allow_list_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the rocketmq allow list.",
			},
		},
	}
	return resource
}

func resourceVolcengineRocketmqAllowListAssociateCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRocketmqAllowListAssociateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineRocketmqAllowListAssociate())
	if err != nil {
		return fmt.Errorf("error on creating rocketmq_allow_list_associate %q, %s", d.Id(), err)
	}
	return resourceVolcengineRocketmqAllowListAssociateRead(d, meta)
}

func resourceVolcengineRocketmqAllowListAssociateRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRocketmqAllowListAssociateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineRocketmqAllowListAssociate())
	if err != nil {
		return fmt.Errorf("error on reading rocketmq_allow_list_associate %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineRocketmqAllowListAssociateUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRocketmqAllowListAssociateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineRocketmqAllowListAssociate())
	if err != nil {
		return fmt.Errorf("error on updating rocketmq_allow_list_associate %q, %s", d.Id(), err)
	}
	return resourceVolcengineRocketmqAllowListAssociateRead(d, meta)
}

func resourceVolcengineRocketmqAllowListAssociateDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRocketmqAllowListAssociateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineRocketmqAllowListAssociate())
	if err != nil {
		return fmt.Errorf("error on deleting rocketmq_allow_list_associate %q, %s", d.Id(), err)
	}
	return err
}

func allowListAssociateImporter(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	var err error
	items := strings.Split(data.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must be of the form InstanceId:AllowListId")
	}
	err = data.Set("instance_id", items[0])
	if err != nil {
		return []*schema.ResourceData{data}, err
	}
	err = data.Set("allow_list_id", items[1])
	if err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}
