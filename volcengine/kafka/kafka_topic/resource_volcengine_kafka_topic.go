package kafka_topic

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
KafkaTopic can be imported using the instance_id:topic_name, e.g.
```
$ terraform import volcengine_kafka_topic.default kafka-cnoeeapetf4s****:topic
```

*/

func ResourceVolcengineKafkaTopic() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineKafkaTopicCreate,
		Read:   resourceVolcengineKafkaTopicRead,
		Update: resourceVolcengineKafkaTopicUpdate,
		Delete: resourceVolcengineKafkaTopicDelete,
		Importer: &schema.ResourceImporter{
			State: kafkaTopicImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The instance id of the kafka topic.",
			},
			"topic_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the kafka topic.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the kafka topic.",
			},
			"partition_number": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntBetween(1, 300),
				Description:  "The number of partition in kafka topic. The value range in 1-300. This field can only be adjusted up but not down.",
			},
			"replica_number": {
				Type:         schema.TypeInt,
				Optional:     true,
				ForceNew:     true, // 不支持修改
				Default:      3,
				ValidateFunc: validation.IntInSlice([]int{2, 3}),
				Description:  "The number of replica in kafka topic. The value can be 2 or 3. Default is 3.",
			},
			"parameters": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Description: "The parameters of the kafka topic.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"min_insync_replica_number": {
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
							Description: "The min number of sync replica. The default value is the replica number minus 1.",
						},
						"message_max_byte": {
							Type:         schema.TypeInt,
							Optional:     true,
							Default:      10,
							ValidateFunc: validation.IntBetween(1, 12),
							Description:  "The max byte of message. Unit: MB. Valid values: 1-12. Default is 10.",
						},
						"log_retention_hours": {
							Type:         schema.TypeInt,
							Optional:     true,
							Default:      72,
							ValidateFunc: validation.IntBetween(0, 2160),
							Description:  "The retention hours of log. Unit: hour. Valid values: 0-2160. Default is 72.",
						},
					},
				},
			},
			"all_authority": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Whether the kafka topic is configured to be accessible by all users. Default: true.",
			},
			"access_policies": {
				Type:        schema.TypeSet,
				Optional:    true,
				Set:         kafkaAccessPolicyHash,
				Description: "The access policies info of the kafka topic. This field only valid when the value of the AllAuthority is false.",
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Get("all_authority").(bool)
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"user_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The name of SASL user.",
						},
						"access_policy": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"PubSub", "Pub", "Sub"}, false),
							Description:  "The access policy of SASL user. Valid values: `PubSub`, `Pub`, `Sub`.",
						},
					},
				},
			},
		},
	}
	return resource
}

func resourceVolcengineKafkaTopicCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKafkaTopicService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(service, d, ResourceVolcengineKafkaTopic())
	if err != nil {
		return fmt.Errorf("error on creating kafka_topic %q, %s", d.Id(), err)
	}
	return resourceVolcengineKafkaTopicRead(d, meta)
}

func resourceVolcengineKafkaTopicRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKafkaTopicService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(service, d, ResourceVolcengineKafkaTopic())
	if err != nil {
		return fmt.Errorf("error on reading kafka_topic %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineKafkaTopicUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKafkaTopicService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(service, d, ResourceVolcengineKafkaTopic())
	if err != nil {
		return fmt.Errorf("error on updating kafka_topic %q, %s", d.Id(), err)
	}
	return resourceVolcengineKafkaTopicRead(d, meta)
}

func resourceVolcengineKafkaTopicDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKafkaTopicService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineKafkaTopic())
	if err != nil {
		return fmt.Errorf("error on deleting kafka_topic %q, %s", d.Id(), err)
	}
	return err
}
