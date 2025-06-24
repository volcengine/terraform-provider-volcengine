package kafka_allow_list

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
KafkaAllowList can be imported using the id, e.g.
```
$ terraform import volcengine_kafka_allow_list.default resource_id
```

*/

func ResourceVolcengineKafkaAllowList() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineKafkaAllowListCreate,
		Read:   resourceVolcengineKafkaAllowListRead,
		Update: resourceVolcengineKafkaAllowListUpdate,
		Delete: resourceVolcengineKafkaAllowListDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"allow_list_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the allow list.",
			},
			"allow_list_desc": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the allow list.",
			},
			"allow_list": {
				Type:     schema.TypeSet,
				Required: true,
				Description: "Whitelist rule list. " +
					"Supports specifying as IP addresses or IP network segments. " +
					"Each whitelist can be configured with a maximum of 300 IP addresses or network segments.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
	return resource
}

func resourceVolcengineKafkaAllowListCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKafkaAllowListService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineKafkaAllowList())
	if err != nil {
		return fmt.Errorf("error on creating kafka_allow_list %q, %s", d.Id(), err)
	}
	return resourceVolcengineKafkaAllowListRead(d, meta)
}

func resourceVolcengineKafkaAllowListRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKafkaAllowListService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineKafkaAllowList())
	if err != nil {
		return fmt.Errorf("error on reading kafka_allow_list %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineKafkaAllowListUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKafkaAllowListService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineKafkaAllowList())
	if err != nil {
		return fmt.Errorf("error on updating kafka_allow_list %q, %s", d.Id(), err)
	}
	return resourceVolcengineKafkaAllowListRead(d, meta)
}

func resourceVolcengineKafkaAllowListDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKafkaAllowListService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineKafkaAllowList())
	if err != nil {
		return fmt.Errorf("error on deleting kafka_allow_list %q, %s", d.Id(), err)
	}
	return err
}
