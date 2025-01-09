package control_policy

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
ControlPolicy can be imported using the direction:rule_id, e.g.
```
$ terraform import volcengine_control_policy.default resource_id
```

*/

func ResourceVolcengineControlPolicy() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineControlPolicyCreate,
		Read:   resourceVolcengineControlPolicyRead,
		Update: resourceVolcengineControlPolicyUpdate,
		Delete: resourceVolcengineControlPolicyDelete,
		Importer: &schema.ResourceImporter{
			State: controlPolicyImporter,
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
				Description: "The action of the control policy. Valid values: `accept`, `deny`, `monitor`.",
			},
			"direction": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The direction of the control policy. Valid values: `in`, `out`.",
			},
			"destination_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The destination type of the control policy. Valid values: `net`, `group`, `location`, `domain`.",
			},
			"destination": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The destination of the control policy.",
			},
			"proto": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The proto of the control policy. Valid values: `TCP`, `ICMP`, `UDP`, `ANY`. When the destination_type is `domain`, The proto must be `TCP`.",
			},
			"source_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The source type of the control policy. Valid values: `net`, `group`, `location`.",
			},
			"source": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The source of the control policy.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the control policy.",
			},
			"dest_port_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The dest port type of the control policy. Valid values: `port`, `group`.",
			},
			"dest_port": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The dest port of the control policy.",
			},
			"repeat_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The repeat type of the control policy. Valid values: `Permanent`, `Once`, `Daily`, `Weekly`, `Monthly`.",
			},
			"repeat_start_time": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "The repeat start time of the control policy. Accurate to the minute, in the format of hh: mm. For example: 12:00.\n " +
					"When the value of repeat_type is one of `Daily`, `Weekly`, `Monthly`, this field is required.",
			},
			"repeat_end_time": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "The repeat end time of the control policy. Accurate to the minute, in the format of hh: mm. For example: 12:00.\n " +
					"When the value of repeat_type is one of `Daily`, `Weekly`, `Monthly`, this field is required.",
			},
			"repeat_days": {
				Type:     schema.TypeSet,
				Optional: true,
				Set:      schema.HashInt,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Description: "The repeat days of the control policy. When the value of repeat_type is one of `Weekly`, `Monthly`, this field is required.\n " +
					"When the repeat_type is `Weekly`, the valid value range is 0~6.\n " +
					"When the repeat_type is `Monthly`, the valid value range is 1~31.",
			},
			"start_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Description: "The start time of the control policy. Unix timestamp, fields need to be precise to 23:59:00 of the set date.\n " +
					"When the value of repeat_type is one of `Once`, `Daily`, `Weekly`, `Monthly`, this field is required.",
			},
			"end_time": {
				Type:     schema.TypeInt,
				Optional: true,
				Description: "The end time of the control policy. Unix timestamp, fields need to be precise to 23:59:00 of the set date.\n " +
					"When the value of repeat_type is one of `Once`, `Daily`, `Weekly`, `Monthly`, this field is required.",
			},
			"priority": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
				ForceNew: true,
				Description: "The priority of the control policy. Default is 0. This field is only effective when creating a control policy." +
					"0 means lowest priority, 1 means highest priority. The priority increases in order from 1, with lower priority indicating higher priority.",
			},
			"status": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Whether to enable the control policy. Default is false.",
			},

			// computed fields
			"rule_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The rule id of the control policy.",
			},
			"prio": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The priority of the control policy.",
			},
			"hit_cnt": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The hit count of the control policy.",
			},
			"use_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The use count of the control policy.",
			},
			"account_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The account id of the control policy.",
			},
			"is_effected": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the control policy is effected.",
			},
			"update_time": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The update time of the control policy.",
			},
			"effect_status": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The effect status of the control policy. 1: Not yet effective, 2: Issued in progress, 3: Effective.",
			},
		},
	}
	return resource
}

func resourceVolcengineControlPolicyCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewControlPolicyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineControlPolicy())
	if err != nil {
		return fmt.Errorf("error on creating control_policy %q, %s", d.Id(), err)
	}
	return resourceVolcengineControlPolicyRead(d, meta)
}

func resourceVolcengineControlPolicyRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewControlPolicyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineControlPolicy())
	if err != nil {
		return fmt.Errorf("error on reading control_policy %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineControlPolicyUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewControlPolicyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineControlPolicy())
	if err != nil {
		return fmt.Errorf("error on updating control_policy %q, %s", d.Id(), err)
	}
	return resourceVolcengineControlPolicyRead(d, meta)
}

func resourceVolcengineControlPolicyDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewControlPolicyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineControlPolicy())
	if err != nil {
		return fmt.Errorf("error on deleting control_policy %q, %s", d.Id(), err)
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
