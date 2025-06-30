package kafka_allow_list_associate

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
KafkaAllowListAssociate can be imported using the id, e.g.
```
$ terraform import volcengine_kafka_allow_list_associate.default kafka-cnitzqgn**:acl-d1fd76693bd54e658912e7337d5b****
```

*/

func ResourceVolcengineKafkaAllowListAssociate() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineKafkaAllowListAssociateCreate,
		Read:   resourceVolcengineKafkaAllowListAssociateRead,
		Delete: resourceVolcengineKafkaAllowListAssociateDelete,
		Importer: &schema.ResourceImporter{
			State: importAllowListAssociate,
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
				Description: "The id of the kafka instance.",
			},
			"allow_list_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the allow list.",
			},
		},
	}
	return resource
}

func resourceVolcengineKafkaAllowListAssociateCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKafkaAllowListAssociateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineKafkaAllowListAssociate())
	if err != nil {
		return fmt.Errorf("error on creating kafka_allow_list_associate %q, %s", d.Id(), err)
	}
	return resourceVolcengineKafkaAllowListAssociateRead(d, meta)
}

func resourceVolcengineKafkaAllowListAssociateRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKafkaAllowListAssociateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineKafkaAllowListAssociate())
	if err != nil {
		return fmt.Errorf("error on reading kafka_allow_list_associate %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineKafkaAllowListAssociateDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKafkaAllowListAssociateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineKafkaAllowListAssociate())
	if err != nil {
		return fmt.Errorf("error on deleting kafka_allow_list_associate %q, %s", d.Id(), err)
	}
	return err
}

func importAllowListAssociate(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
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
