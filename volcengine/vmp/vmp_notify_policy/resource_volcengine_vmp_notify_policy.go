package vmp_notify_policy

import (
	"bytes"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
VMP Notify Policy can be imported using the id, e.g.
```
$ terraform import volcengine_vmp_notify_policy.default 60dde3ca-951c-4c05-8777-e5a7caa07ad6
```

*/

func ResourceVolcengineVmpNotifyPolicy() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVmpNotifyPolicyCreate,
		Read:   resourceVolcengineVmpNotifyPolicyRead,
		Update: resourceVolcengineVmpNotifyPolicyUpdate,
		Delete: resourceVolcengineVmpNotifyPolicyDelete,
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
				Description: "The name of the notify policy.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the notify policy.",
			},
			"channel_notify_template_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The channel notify template for the alarm notification policy.",
			},
			"levels": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "The levels of the notify policy.",
				Set:         levelsHash,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"level": {
							Type:     schema.TypeString,
							Required: true,
							ValidateFunc: validation.StringInSlice([]string{
								"P0",
								"P1",
								"P2",
							}, false),
							Description: "The level of the policy, the value can be one of the following: `P0`, `P1`, `P2`.",
						},
						"contact_group_ids": {
							Type:     schema.TypeSet,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set:         schema.HashString,
							Description: "The contact group for the alarm notification policy.",
						},
						"channels": {
							Type:     schema.TypeSet,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set:         schema.HashString,
							Description: "The alarm notification method of the alarm notification policy, the optional value can be `Email`, `Webhook`, `LarkBotWebhook`, `DingTalkBotWebhook`, `WeComBotWebhook`.",
						},
						"resolved_channels": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set:         schema.HashString,
							Description: "The resolved alarm notification method of the alarm notification policy, the optional value can be `Email`, `Webhook`, `LarkBotWebhook`, `DingTalkBotWebhook`, `WeComBotWebhook`.",
						},
					},
				},
			},
		},
	}
	return resource
}

func resourceVolcengineVmpNotifyPolicyCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(service, d, ResourceVolcengineVmpNotifyPolicy())
	if err != nil {
		return fmt.Errorf("error on creating NotifyPolicy %q, %s", d.Id(), err)
	}
	return resourceVolcengineVmpNotifyPolicyRead(d, meta)
}

func resourceVolcengineVmpNotifyPolicyRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(service, d, ResourceVolcengineVmpNotifyPolicy())
	if err != nil {
		return fmt.Errorf("error on reading NotifyPolicy %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineVmpNotifyPolicyUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(service, d, ResourceVolcengineVmpNotifyPolicy())
	if err != nil {
		return fmt.Errorf("error on updating NotifyPolicy %q, %s", d.Id(), err)
	}
	return resourceVolcengineVmpNotifyPolicyRead(d, meta)
}

func resourceVolcengineVmpNotifyPolicyDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineVmpNotifyPolicy())
	if err != nil {
		return fmt.Errorf("error on deleting NotifyPolicy %q, %s", d.Id(), err)
	}
	return err
}

func levelsHash(v interface{}) int {
	if v == nil {
		return hashcode.String("")
	}
	m := v.(map[string]interface{})
	var (
		buf bytes.Buffer
	)
	buf.WriteString(fmt.Sprintf("%v:", m["level"]))
	buf.WriteString(fmt.Sprintf("%v:", m["contact_group_ids"]))
	buf.WriteString(fmt.Sprintf("%v:", m["channels"]))
	return hashcode.String(buf.String())
}
