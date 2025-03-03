package vpc_firewall_acl_rule_priority

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
VpcFirewallAclRulePriority can be imported using the vpc_firewall_id:rule_id, e.g.
```
$ terraform import volcengine_vpc_firewall_acl_rule_priority.default resource_id
```

*/

func ResourceVolcengineVpcFirewallAclRulePriority() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVpcFirewallAclRulePriorityCreate,
		Read:   resourceVolcengineVpcFirewallAclRulePriorityRead,
		Update: resourceVolcengineVpcFirewallAclRulePriorityUpdate,
		Delete: resourceVolcengineVpcFirewallAclRulePriorityDelete,
		Importer: &schema.ResourceImporter{
			State: vpcFirewallAclRuleImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"vpc_firewall_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the vpc firewall.",
			},
			"rule_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The rule id of the vpc firewall acl rule.",
			},
			"new_prio": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The new priority of the vpc firewall acl rule. The priority increases in order from 1, with lower priority indicating higher priority.",
			},
			"prio": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The priority of the vpc firewall acl rule.",
			},
		},
	}
	return resource
}

func resourceVolcengineVpcFirewallAclRulePriorityCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVpcFirewallAclRulePriorityService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineVpcFirewallAclRulePriority())
	if err != nil {
		return fmt.Errorf("error on creating vpc_firewall_acl_rule_priority %q, %s", d.Id(), err)
	}
	return resourceVolcengineVpcFirewallAclRulePriorityRead(d, meta)
}

func resourceVolcengineVpcFirewallAclRulePriorityRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVpcFirewallAclRulePriorityService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineVpcFirewallAclRulePriority())
	if err != nil {
		return fmt.Errorf("error on reading vpc_firewall_acl_rule_priority %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineVpcFirewallAclRulePriorityUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVpcFirewallAclRulePriorityService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineVpcFirewallAclRulePriority())
	if err != nil {
		return fmt.Errorf("error on updating vpc_firewall_acl_rule_priority %q, %s", d.Id(), err)
	}
	return resourceVolcengineVpcFirewallAclRulePriorityRead(d, meta)
}

func resourceVolcengineVpcFirewallAclRulePriorityDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVpcFirewallAclRulePriorityService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineVpcFirewallAclRulePriority())
	if err != nil {
		return fmt.Errorf("error on deleting vpc_firewall_acl_rule_priority %q, %s", d.Id(), err)
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
