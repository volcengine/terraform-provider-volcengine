package waf_system_bot

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
WafSystemBot can be imported using the id, e.g.
```
$ terraform import volcengine_waf_system_bot.default BotType:Host
```

*/

func ResourceVolcengineWafSystemBot() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineWafSystemBotCreate,
		Read:   resourceVolcengineWafSystemBotRead,
		Update: resourceVolcengineWafSystemBotUpdate,
		Delete: resourceVolcengineWafSystemBotDelete,
		Importer: &schema.ResourceImporter{
			State: wafSystemBotImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"bot_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of bot.",
			},
			"host": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Domain name information.",
			},
			"enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Whether to enable bot.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Name of the affiliated project resource.",
			},
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The execution action of the Bot.",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The description of the Bot.",
			},
			"rule_tag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the Bot rule.",
			},
		},
	}
	return resource
}

func resourceVolcengineWafSystemBotCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewWafSystemBotService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineWafSystemBot())
	if err != nil {
		return fmt.Errorf("error on creating waf_system_bot %q, %s", d.Id(), err)
	}
	return resourceVolcengineWafSystemBotRead(d, meta)
}

func resourceVolcengineWafSystemBotRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewWafSystemBotService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineWafSystemBot())
	if err != nil {
		return fmt.Errorf("error on reading waf_system_bot %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineWafSystemBotUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewWafSystemBotService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineWafSystemBot())
	if err != nil {
		return fmt.Errorf("error on updating waf_system_bot %q, %s", d.Id(), err)
	}
	return resourceVolcengineWafSystemBotRead(d, meta)
}

func resourceVolcengineWafSystemBotDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewWafSystemBotService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineWafSystemBot())
	if err != nil {
		return fmt.Errorf("error on deleting waf_system_bot %q, %s", d.Id(), err)
	}
	return err
}
