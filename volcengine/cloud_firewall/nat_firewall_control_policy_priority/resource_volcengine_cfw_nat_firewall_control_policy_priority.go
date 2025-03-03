package nat_firewall_control_policy_priority

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
NatFirewallControlPolicyPriority can be imported using the direction_nat_firewall_id:rule_id, e.g.
```
$ terraform import volcengine_nat_firewall_control_policy_priority.default resource_id
```

*/

func ResourceVolcengineNatFirewallControlPolicyPriority() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineNatFirewallControlPolicyPriorityCreate,
		Read:   resourceVolcengineNatFirewallControlPolicyPriorityRead,
		Update: resourceVolcengineNatFirewallControlPolicyPriorityUpdate,
		Delete: resourceVolcengineNatFirewallControlPolicyPriorityDelete,
		Importer: &schema.ResourceImporter{
			State: natFirewallControlPolicyImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
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
			"rule_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The rule id of the nat firewall control policy.",
			},
			"new_prio": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The new priority of the nat firewall control policy. The priority increases in order from 1, with lower priority indicating higher priority.",
			},
			"prio": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The priority of the nat firewall control policy.",
			},
		},
	}
	return resource
}

func resourceVolcengineNatFirewallControlPolicyPriorityCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewNatFirewallControlPolicyPriorityService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineNatFirewallControlPolicyPriority())
	if err != nil {
		return fmt.Errorf("error on creating nat_firewall_control_policy_priority %q, %s", d.Id(), err)
	}
	return resourceVolcengineNatFirewallControlPolicyPriorityRead(d, meta)
}

func resourceVolcengineNatFirewallControlPolicyPriorityRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewNatFirewallControlPolicyPriorityService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineNatFirewallControlPolicyPriority())
	if err != nil {
		return fmt.Errorf("error on reading nat_firewall_control_policy_priority %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineNatFirewallControlPolicyPriorityUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewNatFirewallControlPolicyPriorityService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineNatFirewallControlPolicyPriority())
	if err != nil {
		return fmt.Errorf("error on updating nat_firewall_control_policy_priority %q, %s", d.Id(), err)
	}
	return resourceVolcengineNatFirewallControlPolicyPriorityRead(d, meta)
}

func resourceVolcengineNatFirewallControlPolicyPriorityDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewNatFirewallControlPolicyPriorityService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineNatFirewallControlPolicyPriority())
	if err != nil {
		return fmt.Errorf("error on deleting nat_firewall_control_policy_priority %q, %s", d.Id(), err)
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
