package control_policy_priority

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
ControlPolicyPriority can be imported using the direction:rule_id, e.g.
```
$ terraform import volcengine_control_policy_priority.default resource_id
```

*/

func ResourceVolcengineControlPolicyPriority() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineControlPolicyPriorityCreate,
		Read:   resourceVolcengineControlPolicyPriorityRead,
		Update: resourceVolcengineControlPolicyPriorityUpdate,
		Delete: resourceVolcengineControlPolicyPriorityDelete,
		Importer: &schema.ResourceImporter{
			State: controlPolicyImporter,
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
				Description: "The direction of the control policy. Valid values: `in`, `out`.",
			},
			"rule_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The rule id of the control policy.",
			},
			"new_prio": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The new priority of the control policy. The priority increases in order from 1, with lower priority indicating higher priority.",
			},
			"prio": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The priority of the control policy.",
			},
		},
	}
	return resource
}

func resourceVolcengineControlPolicyPriorityCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewControlPolicyPriorityService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineControlPolicyPriority())
	if err != nil {
		return fmt.Errorf("error on creating control_policy_priority %q, %s", d.Id(), err)
	}
	return resourceVolcengineControlPolicyPriorityRead(d, meta)
}

func resourceVolcengineControlPolicyPriorityRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewControlPolicyPriorityService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineControlPolicyPriority())
	if err != nil {
		return fmt.Errorf("error on reading control_policy_priority %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineControlPolicyPriorityUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewControlPolicyPriorityService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineControlPolicyPriority())
	if err != nil {
		return fmt.Errorf("error on updating control_policy_priority %q, %s", d.Id(), err)
	}
	return resourceVolcengineControlPolicyPriorityRead(d, meta)
}

func resourceVolcengineControlPolicyPriorityDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewControlPolicyPriorityService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineControlPolicyPriority())
	if err != nil {
		return fmt.Errorf("error on deleting control_policy_priority %q, %s", d.Id(), err)
	}
	return err
}

var controlPolicyImporter = func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(data.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
	}
	if err := data.Set("direction", items[0]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	if err := data.Set("rule_id", items[1]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}
