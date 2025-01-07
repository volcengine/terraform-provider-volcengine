package rocketmq_group

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
RocketmqGroup can be imported using the instance_id:group_id, e.g.
```
$ terraform import volcengine_rocketmq_group.default resource_id
```

*/

func ResourceVolcengineRocketmqGroup() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineRocketmqGroupCreate,
		Read:   resourceVolcengineRocketmqGroupRead,
		Delete: resourceVolcengineRocketmqGroupDelete,
		Importer: &schema.ResourceImporter{
			State: importRocketmqGroup,
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
				Description: "The id of rocketmq instance.",
			},
			"group_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of rocketmq group.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The description of rocketmq group.",
			},

			// computed fields
			"group_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of the rocketmq group.",
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The create time of the rocketmq group.",
			},
			"is_sub_same": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the subscription relationship of consumer instance groups within the group is consistent.",
			},
			"message_delay_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The message delay time of the rocketmq group. The unit is milliseconds.",
			},
			"message_model": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The message model of the rocketmq group.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the rocketmq group.",
			},
			"total_consume_rate": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The total consume rate of the rocketmq group. The unit is per second.",
			},
			"total_diff": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total amount of unconsumed messages.",
			},
		},
	}
	return resource
}

func resourceVolcengineRocketmqGroupCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRocketmqGroupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineRocketmqGroup())
	if err != nil {
		return fmt.Errorf("error on creating rocketmq_group %q, %s", d.Id(), err)
	}
	return resourceVolcengineRocketmqGroupRead(d, meta)
}

func resourceVolcengineRocketmqGroupRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRocketmqGroupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineRocketmqGroup())
	if err != nil {
		return fmt.Errorf("error on reading rocketmq_group %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineRocketmqGroupUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRocketmqGroupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineRocketmqGroup())
	if err != nil {
		return fmt.Errorf("error on updating rocketmq_group %q, %s", d.Id(), err)
	}
	return resourceVolcengineRocketmqGroupRead(d, meta)
}

func resourceVolcengineRocketmqGroupDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRocketmqGroupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineRocketmqGroup())
	if err != nil {
		return fmt.Errorf("error on deleting rocketmq_group %q, %s", d.Id(), err)
	}
	return err
}

func importRocketmqGroup(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	var err error
	items := strings.Split(data.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must be of the form InstanceId:GroupId")
	}
	err = data.Set("instance_id", items[0])
	if err != nil {
		return []*schema.ResourceData{data}, err
	}
	err = data.Set("group_id", items[1])
	if err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}
