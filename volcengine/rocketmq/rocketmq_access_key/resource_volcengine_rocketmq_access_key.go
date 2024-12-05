package rocketmq_access_key

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
RocketmqAccessKey can be imported using the instance_id:access_key, e.g.
```
$ terraform import volcengine_rocketmq_access_key.default resource_id
```

*/

func ResourceVolcengineRocketmqAccessKey() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineRocketmqAccessKeyCreate,
		Read:   resourceVolcengineRocketmqAccessKeyRead,
		Update: resourceVolcengineRocketmqAccessKeyUpdate,
		Delete: resourceVolcengineRocketmqAccessKeyDelete,
		Importer: &schema.ResourceImporter{
			State: importRocketmqAccessKey,
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
			"all_authority": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The default authority of the rocketmq topic. Valid values: `ALL`, `PUB`, `SUB`, `DENY`. Default is `DENY`.",
			},
			"description": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The description of the rocketmq topic. The description is used to effectively distinguish and manage keys. Please use different descriptions for each key.",
			},

			// computed fields
			"access_key": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The access key id of the rocketmq key.",
			},
			"acl_config_json": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The acl config of the rocketmq key.",
			},
			"actived": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "The active status of the rocketmq key.",
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The create time of the rocketmq key.",
			},
			"secret_key": {
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "The secret key of the rocketmq key.",
			},
			"topic_permissions": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The custom authority of the rocketmq key.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"topic_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the rocketmq topic.",
						},
						"permission": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The custom authority for the topic.",
						},
					},
				},
			},
		},
	}
	return resource
}

func resourceVolcengineRocketmqAccessKeyCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRocketmqAccessKeyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineRocketmqAccessKey())
	if err != nil {
		return fmt.Errorf("error on creating rocketmq_access_key %q, %s", d.Id(), err)
	}
	return resourceVolcengineRocketmqAccessKeyRead(d, meta)
}

func resourceVolcengineRocketmqAccessKeyRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRocketmqAccessKeyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineRocketmqAccessKey())
	if err != nil {
		return fmt.Errorf("error on reading rocketmq_access_key %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineRocketmqAccessKeyUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRocketmqAccessKeyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineRocketmqAccessKey())
	if err != nil {
		return fmt.Errorf("error on updating rocketmq_access_key %q, %s", d.Id(), err)
	}
	return resourceVolcengineRocketmqAccessKeyRead(d, meta)
}

func resourceVolcengineRocketmqAccessKeyDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRocketmqAccessKeyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineRocketmqAccessKey())
	if err != nil {
		return fmt.Errorf("error on deleting rocketmq_access_key %q, %s", d.Id(), err)
	}
	return err
}

func importRocketmqAccessKey(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	var err error
	items := strings.Split(data.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must be of the form InstanceId:AccessKey")
	}
	err = data.Set("instance_id", items[0])
	if err != nil {
		return []*schema.ResourceData{data}, err
	}
	err = data.Set("access_key", items[1])
	if err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}
