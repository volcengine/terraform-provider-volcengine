package consumer_group

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
ConsumerGroup can be imported using the id, e.g.
```
$ terraform import volcengine_consumer_group.default resource_id
```

*/

func ResourceVolcengineConsumerGroup() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineConsumerGroupCreate,
		Read:   resourceVolcengineConsumerGroupRead,
		Update: resourceVolcengineConsumerGroupUpdate,
		Delete: resourceVolcengineConsumerGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"project_id": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "The log project ID to which the consumption group belongs.",
			},
			"topic_id_list": {
				Required: true,
				Type:     schema.TypeSet,
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The list of log topic ids to be consumed by the consumer group.",
			},
			"consumer_group_name": {
				Required:    true,
				ForceNew:    true,
				Type:        schema.TypeString,
				Description: "The name of the consumer group.",
			},
			"heartbeat_ttl": {
				Required:    true,
				Type:        schema.TypeInt,
				Description: "The time of heart rate expiration, measured in seconds, has a value range of 1 to 300.",
			},
			"ordered_consume": {
				Required:    true,
				Type:        schema.TypeBool,
				Description: "Whether to consume in sequence.",
			},
		},
	}
	return resource
}

func resourceVolcengineConsumerGroupCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewConsumerGroupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineConsumerGroup())
	if err != nil {
		return fmt.Errorf("error on creating consumer_group %q, %s", d.Id(), err)
	}
	return resourceVolcengineConsumerGroupRead(d, meta)
}

func resourceVolcengineConsumerGroupRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewConsumerGroupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineConsumerGroup())
	if err != nil {
		return fmt.Errorf("error on reading consumer_group %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineConsumerGroupUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewConsumerGroupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineConsumerGroup())
	if err != nil {
		return fmt.Errorf("error on updating consumer_group %q, %s", d.Id(), err)
	}
	return resourceVolcengineConsumerGroupRead(d, meta)
}

func resourceVolcengineConsumerGroupDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewConsumerGroupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineConsumerGroup())
	if err != nil {
		return fmt.Errorf("error on deleting consumer_group %q, %s", d.Id(), err)
	}
	return err
}
