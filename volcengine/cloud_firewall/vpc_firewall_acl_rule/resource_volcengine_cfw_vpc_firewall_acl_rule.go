package vpc_firewall_acl_rule

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
VpcFirewallAclRule can be imported using the vpc_firewall_id:rule_id, e.g.
```
$ terraform import volcengine_vpc_firewall_acl_rule.default resource_id
```

*/

func ResourceVolcengineVpcFirewallAclRule() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVpcFirewallAclRuleCreate,
		Read:   resourceVolcengineVpcFirewallAclRuleRead,
		Update: resourceVolcengineVpcFirewallAclRuleUpdate,
		Delete: resourceVolcengineVpcFirewallAclRuleDelete,
		Importer: &schema.ResourceImporter{
			State: vpcFirewallAclRuleImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"action": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The action of the vpc firewall acl rule. Valid values: `accept`, `deny`, `monitor`.",
			},
			"vpc_firewall_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the vpc firewall.",
			},
			"destination_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The destination type of the vpc firewall acl rule. Valid values: `net`, `group`, `location`, `domain`.",
			},
			"destination": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The destination of the vpc firewall acl rule.",
			},
			"proto": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The proto of the vpc firewall acl rule. Valid values: `TCP`, `ICMP`, `UDP`, `ANY`. When the destination_type is `domain`, The proto must be `TCP`.",
			},
			"source_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The source type of the vpc firewall acl rule. Valid values: `net`, `group`.",
			},
			"source": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The source of the vpc firewall acl rule.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the vpc firewall acl rule.",
			},
			"dest_port_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The dest port type of the vpc firewall acl rule. Valid values: `port`, `group`.",
			},
			"dest_port": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The dest port of the vpc firewall acl rule.",
			},
			"repeat_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The repeat type of the vpc firewall acl rule. Valid values: `Permanent`, `Once`, `Daily`, `Weekly`, `Monthly`.",
			},
			"repeat_start_time": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "The repeat start time of the vpc firewall acl rule. Accurate to the minute, in the format of hh: mm. For example: 12:00.\n " +
					"When the value of repeat_type is one of `Daily`, `Weekly`, `Monthly`, this field is required.",
			},
			"repeat_end_time": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "The repeat end time of the vpc firewall acl rule. Accurate to the minute, in the format of hh: mm. For example: 12:00.\n " +
					"When the value of repeat_type is one of `Daily`, `Weekly`, `Monthly`, this field is required.",
			},
			"repeat_days": {
				Type:     schema.TypeSet,
				Optional: true,
				Set:      schema.HashInt,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Description: "The repeat days of the vpc firewall acl rule. When the value of repeat_type is one of `Weekly`, `Monthly`, this field is required.\n " +
					"When the repeat_type is `Weekly`, the valid value range is 0~6.\n " +
					"When the repeat_type is `Monthly`, the valid value range is 1~31.",
			},
			"start_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Description: "The start time of the vpc firewall acl rule. Unix timestamp, fields need to be precise to 23:59:00 of the set date.\n " +
					"When the value of repeat_type is one of `Once`, `Daily`, `Weekly`, `Monthly`, this field is required.",
			},
			"end_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Description: "The end time of the vpc firewall acl rule. Unix timestamp, fields need to be precise to 23:59:00 of the set date.\n " +
					"When the value of repeat_type is one of `Once`, `Daily`, `Weekly`, `Monthly`, this field is required.",
			},
			"priority": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Id() != ""
				},
				Description: "The priority of the vpc firewall acl rule. Default is 0. This field is only effective when creating a control policy." +
					"0 means lowest priority, 1 means highest priority. The priority increases in order from 1, with lower priority indicating higher priority.",
			},
			"status": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Whether to enable the vpc firewall acl rule. Default is false.",
			},

			// computed fields
			"rule_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The rule id of the vpc firewall acl rule.",
			},
			"prio": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The priority of the vpc firewall acl rule.",
			},
			"vpc_firewall_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the vpc firewall.",
			},
			"hit_cnt": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The hit count of the vpc firewall acl rule.",
			},
			"use_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The use count of the vpc firewall acl rule.",
			},
			"account_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The account id of the vpc firewall acl rule.",
			},
			"is_effected": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the vpc firewall acl rule is effected.",
			},
			"update_time": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The update time of the vpc firewall acl rule.",
			},
			"effect_status": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The effect status of the vpc firewall acl rule. 1: Not yet effective, 2: Issued in progress, 3: Effective.",
			},
		},
	}
	return resource
}

func resourceVolcengineVpcFirewallAclRuleCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVpcFirewallAclRuleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineVpcFirewallAclRule())
	if err != nil {
		return fmt.Errorf("error on creating vpc_firewall_acl_rule %q, %s", d.Id(), err)
	}
	return resourceVolcengineVpcFirewallAclRuleRead(d, meta)
}

func resourceVolcengineVpcFirewallAclRuleRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVpcFirewallAclRuleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineVpcFirewallAclRule())
	if err != nil {
		return fmt.Errorf("error on reading vpc_firewall_acl_rule %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineVpcFirewallAclRuleUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVpcFirewallAclRuleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineVpcFirewallAclRule())
	if err != nil {
		return fmt.Errorf("error on updating vpc_firewall_acl_rule %q, %s", d.Id(), err)
	}
	return resourceVolcengineVpcFirewallAclRuleRead(d, meta)
}

func resourceVolcengineVpcFirewallAclRuleDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVpcFirewallAclRuleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineVpcFirewallAclRule())
	if err != nil {
		return fmt.Errorf("error on deleting vpc_firewall_acl_rule %q, %s", d.Id(), err)
	}
	return err
}

var vpcFirewallAclRuleImporter = func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(data.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
	}
	if err := data.Set("vpc_firewall_id", items[0]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	if err := data.Set("rule_id", items[1]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}
