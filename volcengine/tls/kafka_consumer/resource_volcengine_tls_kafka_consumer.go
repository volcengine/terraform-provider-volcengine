package kafka_consumer

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Tls Kafka Consumer can be imported using the kafka:topic_id, e.g.
```
$ terraform import volcengine_tls_kafka_consumer.default kafka:edf051ed-3c46-49ba-9339-bea628fedc15
```

*/

func ResourceVolcengineTlsKafkaConsumer() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineTlsKafkaConsumerCreate,
		Read:   resourceVolcengineTlsKafkaConsumerRead,
		Delete: resourceVolcengineTlsKafkaConsumerDelete,
		Importer: &schema.ResourceImporter{
			State: func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
				items := strings.Split(data.Id(), ":")
				if len(items) != 2 {
					return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
				}
				if err := data.Set("topic_id", items[1]); err != nil {
					return []*schema.ResourceData{data}, err
				}
				return []*schema.ResourceData{data}, nil
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"topic_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of topic.",
			},
			"allow_consume": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether allow consume.",
			},
			"consume_topic": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The topic of consume.",
			},
		},
	}
	return resource
}

func resourceVolcengineTlsKafkaConsumerCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(service, d, ResourceVolcengineTlsKafkaConsumer())
	if err != nil {
		return fmt.Errorf("error on creating tls kafka consumer %q, %s", d.Id(), err)
	}
	return resourceVolcengineTlsKafkaConsumerRead(d, meta)
}

func resourceVolcengineTlsKafkaConsumerRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(service, d, ResourceVolcengineTlsKafkaConsumer())
	if err != nil {
		return fmt.Errorf("error on reading tls kafka consumer %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineTlsKafkaConsumerDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineTlsKafkaConsumer())
	if err != nil {
		return fmt.Errorf("error on deleting tls kafka consumer %q, %s", d.Id(), err)
	}
	return err
}
