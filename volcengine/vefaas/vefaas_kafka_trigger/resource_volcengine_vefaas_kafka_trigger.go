package vefaas_kafka_trigger

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
VefaasKafkaTrigger can be imported using the id, e.g.
```
$ terraform import volcengine_vefaas_kafka_trigger.default resource_id
```

*/

func ResourceVolcengineVefaasKafkaTrigger() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVefaasKafkaTriggerCreate,
		Read:   resourceVolcengineVefaasKafkaTriggerRead,
		Update: resourceVolcengineVefaasKafkaTriggerUpdate,
		Delete: resourceVolcengineVefaasKafkaTriggerDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"function_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of Function.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the Kafka trigger.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The description of the Kafka trigger.",
			},
			"mq_instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The instance ID of Message queue Kafka.",
			},
			"topic_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The Topic name of the message queue Kafka instance.",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Whether to enable triggers at the same time as creating them.",
			},
			"starting_position": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "Specify the location where the messages in the Topic start to be consumed.",
			},
			"maximum_retry_attempts": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "The maximum number of retries when a function has a runtime error.",
			},
			"kafka_credentials": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				ForceNew: true,
				Description: "Kafka identity authentication. " +
					"When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"password": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The SASL/PLAIN user password set when creating a Kafka instance.",
						},
						"username": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The SASL/PLAIN user name set when creating a Kafka instance.",
						},
						"mechanism": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Kafka authentication mechanism.",
						},
					},
				},
			},
			"batch_size": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "The number of messages per batch consumed by the trigger in bulk.",
			},
			"batch_flush_duration_milliseconds": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "The maximum waiting time for batch consumption of triggers.",
			},
			"consumer_group": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The consumer group name of the message queue Kafka instance.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of Kafka trigger.",
			},
			"creation_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The creation time of the Kafka trigger.",
			},
			"last_update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The last update time of the Kafka trigger.",
			},
		},
	}
	return resource
}

func resourceVolcengineVefaasKafkaTriggerCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVefaasKafkaTriggerService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineVefaasKafkaTrigger())
	if err != nil {
		return fmt.Errorf("error on creating vefaas_kafka_trigger %q, %s", d.Id(), err)
	}
	return resourceVolcengineVefaasKafkaTriggerRead(d, meta)
}

func resourceVolcengineVefaasKafkaTriggerRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVefaasKafkaTriggerService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineVefaasKafkaTrigger())
	if err != nil {
		return fmt.Errorf("error on reading vefaas_kafka_trigger %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineVefaasKafkaTriggerUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVefaasKafkaTriggerService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineVefaasKafkaTrigger())
	if err != nil {
		return fmt.Errorf("error on updating vefaas_kafka_trigger %q, %s", d.Id(), err)
	}
	return resourceVolcengineVefaasKafkaTriggerRead(d, meta)
}

func resourceVolcengineVefaasKafkaTriggerDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVefaasKafkaTriggerService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineVefaasKafkaTrigger())
	if err != nil {
		return fmt.Errorf("error on deleting vefaas_kafka_trigger %q, %s", d.Id(), err)
	}
	return err
}
