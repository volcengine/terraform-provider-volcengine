package vmp_notify_group_policy

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
VMP Notify Group Policy can be imported using the id, e.g.
```
$ terraform import volcengine_vmp_notify_group_policy.default 60dde3ca-951c-4c05-8777-e5a7caa07ad6
```

*/

func ResourceVolcengineVmpNotifyGroupPolicy() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVmpNotifyGroupPolicyCreate,
		Read:   resourceVolcengineVmpNotifyGroupPolicyRead,
		Update: resourceVolcengineVmpNotifyGroupPolicyUpdate,
		Delete: resourceVolcengineVmpNotifyGroupPolicyDelete,
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
				Description: "The name of the notify group policy.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the notify group policy.",
			},
			"levels": {
				Type:        schema.TypeSet,
				Required:    true,
				MaxItems:    3,
				MinItems:    3,
				Set:         levelsHash,
				Description: "The levels of the notify group policy. Levels must be registered in three (`P0`, `P1`, `P2`) aggregation strategies, and `Level` cannot be repeated.",
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
						"group_by": {
							Type:     schema.TypeSet,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
								ValidateFunc: validation.StringInSlice([]string{
									"__rule__",
								}, false),
							},
							Set:         schema.HashString,
							Description: "The aggregate dimension, the value can be `__rule__`.",
						},
						"group_wait": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The wait time. Integer form, unit is second.",
						},
						"group_interval": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The aggregation cycle. Integer form, unit is second.",
						},
						"repeat_interval": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The notification cycle. Integer form, unit is second.",
						},
					},
				},
			},
		},
	}
	return resource
}

func resourceVolcengineVmpNotifyGroupPolicyCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(service, d, ResourceVolcengineVmpNotifyGroupPolicy())
	if err != nil {
		return fmt.Errorf("error on creating NotifyGroupPolicy %q, %s", d.Id(), err)
	}
	return resourceVolcengineVmpNotifyGroupPolicyRead(d, meta)
}

func resourceVolcengineVmpNotifyGroupPolicyRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(service, d, ResourceVolcengineVmpNotifyGroupPolicy())
	if err != nil {
		return fmt.Errorf("error on reading NotifyGroupPolicy %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineVmpNotifyGroupPolicyUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(service, d, ResourceVolcengineVmpNotifyGroupPolicy())
	if err != nil {
		return fmt.Errorf("error on updating NotifyGroupPolicy %q, %s", d.Id(), err)
	}
	return resourceVolcengineVmpNotifyGroupPolicyRead(d, meta)
}

func resourceVolcengineVmpNotifyGroupPolicyDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineVmpNotifyGroupPolicy())
	if err != nil {
		return fmt.Errorf("error on deleting NotifyGroupPolicy %q, %s", d.Id(), err)
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
	buf.WriteString(fmt.Sprintf("%v:", m["group_by"]))
	buf.WriteString(fmt.Sprintf("%v:", m["group_wait"]))
	buf.WriteString(fmt.Sprintf("%v:", m["group_interval"]))
	buf.WriteString(fmt.Sprintf("%v:", m["repeat_interval"]))
	return hashcode.String(buf.String())
}
