package rocketmq_topic

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
RocketmqTopic can be imported using the instance_id:topic_name, e.g.
```
$ terraform import volcengine_rocketmq_topic.default resource_id
```

*/

func ResourceVolcengineRocketmqTopic() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineRocketmqTopicCreate,
		Read:   resourceVolcengineRocketmqTopicRead,
		Update: resourceVolcengineRocketmqTopicUpdate,
		Delete: resourceVolcengineRocketmqTopicDelete,
		Importer: &schema.ResourceImporter{
			State: importRocketmqTopic,
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
				Description: "The id of rocketmq instance.",
			},
			"topic_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the rocketmq topic.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The description of the rocketmq topic.",
			},
			"queue_number": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "The maximum number of queues for the current topic, which cannot exceed the remaining available queues for the current rocketmq instance.",
			},
			"message_type": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "The type of the message. Valid values: `0`: Regular message, `1`: Transaction message, `2`: Partition order message, `3`: Global sequential message, `4`: Delay message.",
			},
			"access_policies": {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Description: "The access policies of the rocketmq topic. This field can only be added or modified. Deleting this field is invalid.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"access_key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The access key of the rocketmq key.",
						},
						"authority": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The authority of the rocketmq key for the current topic. Valid values: `ALL`, `PUB`, `SUB`, `DENY`. Default is `DENY`.",
						},
					},
				},
			},

			// computed fields
			"queues": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The queues information of the rocketmq topic.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"queue_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the rocketmq queue.",
						},
						"start_offset": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The start offset of the rocketmq queue.",
						},
						"end_offset": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The end offset of the rocketmq queue.",
						},
						"message_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The message count of the rocketmq queue.",
						},
						"last_update_timestamp": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The last update timestamp of the rocketmq queue.",
						},
					},
				},
			},
			"groups": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The groups information of the rocketmq topic.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the rocketmq group.",
						},
						"message_model": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The message model of the rocketmq group.",
						},
						"sub_string": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The sub string of the rocketmq group.",
						},
					},
				},
			},
		},
	}
	return resource
}

func resourceVolcengineRocketmqTopicCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRocketmqTopicService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineRocketmqTopic())
	if err != nil {
		return fmt.Errorf("error on creating rocketmq_topic %q, %s", d.Id(), err)
	}
	return resourceVolcengineRocketmqTopicRead(d, meta)
}

func resourceVolcengineRocketmqTopicRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRocketmqTopicService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineRocketmqTopic())
	if err != nil {
		return fmt.Errorf("error on reading rocketmq_topic %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineRocketmqTopicUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRocketmqTopicService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineRocketmqTopic())
	if err != nil {
		return fmt.Errorf("error on updating rocketmq_topic %q, %s", d.Id(), err)
	}
	return resourceVolcengineRocketmqTopicRead(d, meta)
}

func resourceVolcengineRocketmqTopicDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRocketmqTopicService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineRocketmqTopic())
	if err != nil {
		return fmt.Errorf("error on deleting rocketmq_topic %q, %s", d.Id(), err)
	}
	return err
}

func importRocketmqTopic(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	var err error
	items := strings.Split(data.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must be of the form InstanceId:TopicName")
	}
	err = data.Set("instance_id", items[0])
	if err != nil {
		return []*schema.ResourceData{data}, err
	}
	err = data.Set("topic_name", items[1])
	if err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}
