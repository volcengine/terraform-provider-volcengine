package vmp_alerting_rule

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
VmpAlertingRule can be imported using the id, e.g.
```
$ terraform import volcengine_vmp_alerting_rule.default 5bd29e81-2717-4ac8-a1a6-d76da2b1****
```

*/

func ResourceVolcengineVmpAlertingRule() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVmpAlertingRuleCreate,
		Read:   resourceVolcengineVmpAlertingRuleRead,
		Update: resourceVolcengineVmpAlertingRuleUpdate,
		Delete: resourceVolcengineVmpAlertingRuleDelete,
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
				Description: "The name of the vmp alerting rule.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The description of the vmp alerting rule.",
			},
			"notify_policy_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the notify policy.",
			},
			"notify_group_policy_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of the notify group policy.",
			},
			"query": {
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				Description: "The alerting query of the vmp alerting rule.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"workspace_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The id of the workspace.",
						},
						"prom_ql": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The prom ql of query.",
						},
					},
				},
			},
			"levels": {
				Type:        schema.TypeSet,
				Required:    true,
				MaxItems:    3,
				Set:         alertingRuleLevelsHash,
				Description: "The alerting levels of the vmp alerting rule.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"level": {
							Type:     schema.TypeString,
							Required: true,
							ValidateFunc: validation.StringInSlice([]string{
								"P0", "P1", "P2",
							}, false),
							Description: "The level of the vmp alerting rule. Valid values: `P0`, `P1`, `P2`. The value of this field cannot be duplicate.",
						},
						"for": {
							Type:     schema.TypeString,
							Required: true,
							ValidateFunc: validation.StringInSlice([]string{
								"0s", "1m", "2m", "5m", "10m",
							}, false),
							Description: "The duration of the alerting rule. Valid values: `0s`, `1m`, `2m`, `5m`, `10m`.",
						},
						"comparator": {
							Type:     schema.TypeString,
							Required: true,
							ValidateFunc: validation.StringInSlice([]string{
								">", ">=", "<", "<=", "==", "!=",
							}, false),
							Description: "The comparator of the vmp alerting rule. Valid values: `>`, `>=`, `<`, `<=`, `==`, `!=`.",
						},
						"threshold": {
							Type:        schema.TypeFloat,
							Required:    true,
							Description: "The threshold of the vmp alerting rule.",
						},
					},
				},
			},
			"annotations": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "The annotations of the vmp alerting rule.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The name of the annotation.",
						},
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The value of the annotation.",
						},
					},
				},
			},
			"labels": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "The labels of the vmp alerting rule.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The name of the label.",
						},
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The value of the label.",
						},
					},
				},
			},

			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The create time of the vmp alerting rule.",
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The update time of the vmp alerting rule.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the vmp alerting rule.",
			},
		},
	}
	return resource
}

func alertingRuleLevelsHash(v interface{}) int {
	if v == nil {
		return hashcode.String("")
	}
	m := v.(map[string]interface{})
	var (
		buf bytes.Buffer
	)
	buf.WriteString(fmt.Sprintf("%v:", m["level"]))
	buf.WriteString(fmt.Sprintf("%v:", m["for"]))
	buf.WriteString(fmt.Sprintf("%v:", m["comparator"]))
	buf.WriteString(fmt.Sprintf("%v:", m["threshold"]))
	return hashcode.String(buf.String())
}

func resourceVolcengineVmpAlertingRuleCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVmpAlertingRuleService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(service, d, ResourceVolcengineVmpAlertingRule())
	if err != nil {
		return fmt.Errorf("error on creating vmp alerting rule %q, %s", d.Id(), err)
	}
	return resourceVolcengineVmpAlertingRuleRead(d, meta)
}

func resourceVolcengineVmpAlertingRuleRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVmpAlertingRuleService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(service, d, ResourceVolcengineVmpAlertingRule())
	if err != nil {
		return fmt.Errorf("error on reading vmp alerting rule %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineVmpAlertingRuleUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVmpAlertingRuleService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(service, d, ResourceVolcengineVmpAlertingRule())
	if err != nil {
		return fmt.Errorf("error on updating vmp alerting rule %q, %s", d.Id(), err)
	}
	return resourceVolcengineVmpAlertingRuleRead(d, meta)
}

func resourceVolcengineVmpAlertingRuleDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVmpAlertingRuleService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineVmpAlertingRule())
	if err != nil {
		return fmt.Errorf("error on deleting vmp alerting rule %q, %s", d.Id(), err)
	}
	return err
}
