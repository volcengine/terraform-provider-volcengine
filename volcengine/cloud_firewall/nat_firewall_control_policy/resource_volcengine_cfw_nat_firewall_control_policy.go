package nat_firewall_control_policy

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
NatFirewallControlPolicy can be imported using the direction_nat_firewall_id:rule_id, e.g.
```
$ terraform import volcengine_nat_firewall_control_policy.default resource_id
```

*/

func ResourceVolcengineNatFirewallControlPolicy() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineNatFirewallControlPolicyCreate,
		Read:   resourceVolcengineNatFirewallControlPolicyRead,
		Update: resourceVolcengineNatFirewallControlPolicyUpdate,
		Delete: resourceVolcengineNatFirewallControlPolicyDelete,
		Importer: &schema.ResourceImporter{
			State: natFirewallControlPolicyImporter,
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
				Description: "The action of the nat firewall control policy. Valid values: `accept`, `deny`, `monitor`.",
			},
			"direction": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The direction of the nat firewall control policy. Valid values: `in`, `out`.",
			},
			"nat_firewall_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the nat firewall.",
			},
			"destination_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The destination type of the nat firewall control policy. Valid values: `net`, `group`, `location`, `domain`.",
			},
			"destination": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The destination of the nat firewall control policy.",
			},
			"proto": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The proto of the nat firewall control policy. Valid values: `TCP`, `ICMP`, `UDP`, `ANY`. When the destination_type is `domain`, The proto must be `TCP`.",
			},
			"source_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The source type of the nat firewall control policy. Valid values: `net`, `group`.",
			},
			"source": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The source of the nat firewall control policy.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the nat firewall control policy.",
			},
			"dest_port_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The dest port type of the nat firewall control policy. Valid values: `port`, `group`.",
			},
			"dest_port": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The dest port of the nat firewall control policy.",
			},
			"repeat_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The repeat type of the nat firewall control policy. Valid values: `Permanent`, `Once`, `Daily`, `Weekly`, `Monthly`.",
			},
			"repeat_start_time": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "The repeat start time of the nat firewall control policy. Accurate to the minute, in the format of hh: mm. For example: 12:00.\n " +
					"When the value of repeat_type is one of `Daily`, `Weekly`, `Monthly`, this field is required.",
			},
			"repeat_end_time": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "The repeat end time of the nat firewall control policy. Accurate to the minute, in the format of hh: mm. For example: 12:00.\n " +
					"When the value of repeat_type is one of `Daily`, `Weekly`, `Monthly`, this field is required.",
			},
			"repeat_days": {
				Type:     schema.TypeSet,
				Optional: true,
				Set:      schema.HashInt,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Description: "The repeat days of the nat firewall control policy. When the value of repeat_type is one of `Weekly`, `Monthly`, this field is required.\n " +
					"When the repeat_type is `Weekly`, the valid value range is 0~6.\n " +
					"When the repeat_type is `Monthly`, the valid value range is 1~31.",
			},
			"start_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Description: "The start time of the nat firewall control policy. Unix timestamp, fields need to be precise to 23:59:00 of the set date.\n " +
					"When the value of repeat_type is one of `Once`, `Daily`, `Weekly`, `Monthly`, this field is required.",
			},
			"end_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Description: "The end time of the nat firewall control policy. Unix timestamp, fields need to be precise to 23:59:00 of the set date.\n " +
					"When the value of repeat_type is one of `Once`, `Daily`, `Weekly`, `Monthly`, this field is required.",
			},
			"priority": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return d.Id() != ""
				},
				Description: "The priority of the nat firewall control policy. Default is 0. This field is only effective when creating a control policy." +
					"0 means lowest priority, 1 means highest priority. The priority increases in order from 1, with lower priority indicating higher priority.",
			},
			"status": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Whether to enable the nat firewall control policy. Default is false.",
			},

			// computed fields
			"rule_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The rule id of the nat firewall control policy.",
			},
			"prio": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The priority of the nat firewall control policy.",
			},
			"nat_firewall_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the nat firewall.",
			},
			"hit_cnt": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The hit count of the nat firewall control policy.",
			},
			"use_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The use count of the nat firewall control policy.",
			},
			"account_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The account id of the nat firewall control policy.",
			},
			"is_effected": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the nat firewall control policy is effected.",
			},
			"update_time": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The update time of the nat firewall control policy.",
			},
			"effect_status": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The effect status of the nat firewall control policy. 1: Not yet effective, 2: Issued in progress, 3: Effective.",
			},
		},
	}
	return resource
}

func resourceVolcengineNatFirewallControlPolicyCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewNatFirewallControlPolicyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineNatFirewallControlPolicy())
	if err != nil {
		return fmt.Errorf("error on creating nat_firewall_control_policy %q, %s", d.Id(), err)
	}
	return resourceVolcengineNatFirewallControlPolicyRead(d, meta)
}

func resourceVolcengineNatFirewallControlPolicyRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewNatFirewallControlPolicyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineNatFirewallControlPolicy())
	if err != nil {
		return fmt.Errorf("error on reading nat_firewall_control_policy %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineNatFirewallControlPolicyUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewNatFirewallControlPolicyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineNatFirewallControlPolicy())
	if err != nil {
		return fmt.Errorf("error on updating nat_firewall_control_policy %q, %s", d.Id(), err)
	}
	return resourceVolcengineNatFirewallControlPolicyRead(d, meta)
}

func resourceVolcengineNatFirewallControlPolicyDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewNatFirewallControlPolicyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineNatFirewallControlPolicy())
	if err != nil {
		return fmt.Errorf("error on deleting nat_firewall_control_policy %q, %s", d.Id(), err)
	}
	return err
}

var natFirewallControlPolicyImporter = func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(data.Id(), ":")
	if len(items) != 3 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
	}
	if err := data.Set("direction", items[0]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	if err := data.Set("nat_firewall_id", items[1]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	if err := data.Set("rule_id", items[2]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}
