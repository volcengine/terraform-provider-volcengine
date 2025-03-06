package rabbitmq_instance_plugin

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
RabbitmqInstancePlugin can be imported using the instance_id:plugin_name, e.g.
```
$ terraform import volcengine_rabbitmq_instance_plugin.default resource_id
```

*/

func ResourceVolcengineRabbitmqInstancePlugin() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineRabbitmqInstancePluginCreate,
		Read:   resourceVolcengineRabbitmqInstancePluginRead,
		Delete: resourceVolcengineRabbitmqInstancePluginDelete,
		Importer: &schema.ResourceImporter{
			State: importRabbitmqPlugin,
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
				Description: "The id of the rabbitmq instance..",
			},
			"plugin_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the plugin.",
			},

			// computed fields
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The description of the plugin.",
			},
			"disable_prompt": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The disable prompt of the plugin.",
			},
			"enable_prompt": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The enable prompt of the plugin.",
			},
			"need_reboot_on_change": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Will changing the enabled state of the plugin cause a reboot of the rabbitmq instance.",
			},
			"version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The version of the plugin.",
			},
			"port": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The port of the plugin.",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the plugin is enabled.",
			},
		},
	}
	return resource
}

func resourceVolcengineRabbitmqInstancePluginCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRabbitmqInstancePluginService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineRabbitmqInstancePlugin())
	if err != nil {
		return fmt.Errorf("error on creating rabbitmq_instance_plugin %q, %s", d.Id(), err)
	}
	return resourceVolcengineRabbitmqInstancePluginRead(d, meta)
}

func resourceVolcengineRabbitmqInstancePluginRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRabbitmqInstancePluginService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineRabbitmqInstancePlugin())
	if err != nil {
		return fmt.Errorf("error on reading rabbitmq_instance_plugin %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineRabbitmqInstancePluginUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRabbitmqInstancePluginService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineRabbitmqInstancePlugin())
	if err != nil {
		return fmt.Errorf("error on updating rabbitmq_instance_plugin %q, %s", d.Id(), err)
	}
	return resourceVolcengineRabbitmqInstancePluginRead(d, meta)
}

func resourceVolcengineRabbitmqInstancePluginDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewRabbitmqInstancePluginService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineRabbitmqInstancePlugin())
	if err != nil {
		return fmt.Errorf("error on deleting rabbitmq_instance_plugin %q, %s", d.Id(), err)
	}
	return err
}

func importRabbitmqPlugin(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	var err error
	items := strings.Split(data.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must be of the form InstanceId:Plugin")
	}
	err = data.Set("instance_id", items[0])
	if err != nil {
		return []*schema.ResourceData{data}, err
	}
	err = data.Set("plugin_name", items[1])
	if err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}
