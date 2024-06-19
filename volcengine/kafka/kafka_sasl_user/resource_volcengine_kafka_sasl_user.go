package kafka_sasl_user

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
KafkaSaslUser can be imported using the kafka_id:username, e.g.
```
$ terraform import volcengine_kafka_sasl_user.default kafka-cnngbnntswg1****:tfuser
```

*/

func ResourceVolcengineKafkaSaslUser() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineKafkaSaslUserCreate,
		Read:   resourceVolcengineKafkaSaslUserRead,
		Update: resourceVolcengineKafkaSaslUserUpdate,
		Delete: resourceVolcengineKafkaSaslUserDelete,
		Importer: &schema.ResourceImporter{
			State: func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
				items := strings.Split(data.Id(), ":")
				if len(items) != 2 {
					return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
				}
				if err := data.Set("user_name", items[1]); err != nil {
					return []*schema.ResourceData{data}, err
				}
				if err := data.Set("instance_id", items[0]); err != nil {
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
				Description: "The id of instance.",
				ForceNew:    true,
			},
			"user_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of user.",
				ForceNew:    true,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of user.",
				ForceNew:    true, // 不支持修改
			},
			"user_password": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The password of user.",
				Sensitive:   true,
				ForceNew:    true,
			},
			"all_authority": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Whether this user has read and write permissions for all topics. Default is true.",
			},
			"password_type": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Default:     "Plain",
				Description: "The type of password. Valid values are `Scram` and `Plain`. Default is `Plain`.",
			},
		},
	}
	return resource
}

func resourceVolcengineKafkaSaslUserCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKafkaSaslUserService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineKafkaSaslUser())
	if err != nil {
		return fmt.Errorf("error on creating kafka_sasl_user %q, %s", d.Id(), err)
	}
	return resourceVolcengineKafkaSaslUserRead(d, meta)
}

func resourceVolcengineKafkaSaslUserRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKafkaSaslUserService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineKafkaSaslUser())
	if err != nil {
		return fmt.Errorf("error on reading kafka_sasl_user %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineKafkaSaslUserUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKafkaSaslUserService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineKafkaSaslUser())
	if err != nil {
		return fmt.Errorf("error on updating kafka_sasl_user %q, %s", d.Id(), err)
	}
	return resourceVolcengineKafkaSaslUserRead(d, meta)
}

func resourceVolcengineKafkaSaslUserDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewKafkaSaslUserService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineKafkaSaslUser())
	if err != nil {
		return fmt.Errorf("error on deleting kafka_sasl_user %q, %s", d.Id(), err)
	}
	return err
}
