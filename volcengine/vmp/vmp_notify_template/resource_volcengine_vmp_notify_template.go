package vmp_notify_template

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
VmpNotifyTemplate can be imported using the id, e.g.
```
$ terraform import volcengine_vmp_notify_template.default resource_id
```

*/

func ResourceVolcengineVmpNotifyTemplate() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVmpNotifyTemplateCreate,
		Read:   resourceVolcengineVmpNotifyTemplateRead,
		Update: resourceVolcengineVmpNotifyTemplateUpdate,
		Delete: resourceVolcengineVmpNotifyTemplateDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of notify template.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of notify template.",
			},
			"channel": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The channel of notify template. Valid values: `LarkBotWebhook`, `DingTalkBotWebhook`, `WeComBotWebhook`.",
			},
			"active": {
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				Description: "The active notify template info.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"title": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The title of notify template.",
						},
						"content": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The content of notify template.",
						},
					},
				},
			},
			"resolved": {
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				Description: "The resolved notify template info.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"title": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The title of notify template.",
						},
						"content": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The content of notify template.",
						},
					},
				},
			},

			// computed fields
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The create time of notify template.",
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The update time of notify template.",
			},
		},
	}
	return resource
}

func resourceVolcengineVmpNotifyTemplateCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVmpNotifyTemplateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineVmpNotifyTemplate())
	if err != nil {
		return fmt.Errorf("error on creating vmp_notify_template %q, %s", d.Id(), err)
	}
	return resourceVolcengineVmpNotifyTemplateRead(d, meta)
}

func resourceVolcengineVmpNotifyTemplateRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVmpNotifyTemplateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineVmpNotifyTemplate())
	if err != nil {
		return fmt.Errorf("error on reading vmp_notify_template %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineVmpNotifyTemplateUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVmpNotifyTemplateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineVmpNotifyTemplate())
	if err != nil {
		return fmt.Errorf("error on updating vmp_notify_template %q, %s", d.Id(), err)
	}
	return resourceVolcengineVmpNotifyTemplateRead(d, meta)
}

func resourceVolcengineVmpNotifyTemplateDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVmpNotifyTemplateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineVmpNotifyTemplate())
	if err != nil {
		return fmt.Errorf("error on deleting vmp_notify_template %q, %s", d.Id(), err)
	}
	return err
}
