package kafka_group

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
KafkaGroup can be imported using the instance_id:group_id, e.g.
```
$ terraform import volcengine_kafka_group.default kafka-****x:groupId
```

*/

func ResourceVolcengineKafkaGroup() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineKafkaGroupCreate,
		Read:   resourceVolcengineKafkaGroupRead,
		Update: resourceVolcengineKafkaGroupUpdate,
		Delete: resourceVolcengineKafkaGroupDelete,
		Importer: &schema.ResourceImporter{
			State: func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
				items := strings.Split(data.Id(), ":")
				if len(items) != 2 {
					return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
				}
				if err := data.Set("instance_id", items[0]); err != nil {
					return []*schema.ResourceData{data}, err
				}
				if err := data.Set("group_id", items[1]); err != nil {
					return []*schema.ResourceData{data}, err
				}
				return []*schema.ResourceData{data}, nil
			},
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
				Description: "The instance id of kafka group.",
			},
			"group_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of kafka group.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The description of kafka group.",
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The state of kafka group.",
			},
		},
	}
	return resource
}

func resourceVolcengineKafkaGroupCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKafkaGroupService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(service, d, ResourceVolcengineKafkaGroup())
	if err != nil {
		return fmt.Errorf("error on creating kafka_group %q, %s", d.Id(), err)
	}
	return resourceVolcengineKafkaGroupRead(d, meta)
}

func resourceVolcengineKafkaGroupRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKafkaGroupService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(service, d, ResourceVolcengineKafkaGroup())
	if err != nil {
		return fmt.Errorf("error on reading kafka_group %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineKafkaGroupUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKafkaGroupService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(service, d, ResourceVolcengineKafkaGroup())
	if err != nil {
		return fmt.Errorf("error on updating kafka_group %q, %s", d.Id(), err)
	}
	return resourceVolcengineKafkaGroupRead(d, meta)
}

func resourceVolcengineKafkaGroupDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKafkaGroupService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineKafkaGroup())
	if err != nil {
		return fmt.Errorf("error on deleting kafka_group %q, %s", d.Id(), err)
	}
	return err
}
