package traffic_mirror_filter_rule

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
TrafficMirrorFilterRule can be imported using the id, e.g.
```
$ terraform import volcengine_traffic_mirror_filter_rule.default resource_id
```

*/

func ResourceVolcengineTrafficMirrorFilterRule() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineTrafficMirrorFilterRuleCreate,
		Read:   resourceVolcengineTrafficMirrorFilterRuleRead,
		Update: resourceVolcengineTrafficMirrorFilterRuleUpdate,
		Delete: resourceVolcengineTrafficMirrorFilterRuleDelete,
		Importer: &schema.ResourceImporter{
			State: trafficMirrorFilterRuleImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"traffic_mirror_filter_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of traffic mirror filter.",
			},
			"traffic_direction": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The traffic direction of traffic mirror filter rule. Valid values: `ingress`; `egress`.",
			},
			"priority": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "The priority of traffic mirror filter rule. Valid values: 1~1000. Default value is 1.",
			},
			"policy": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The policy of traffic mirror filter rule. Valid values: `accept`, `reject`.",
			},
			"protocol": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The protocol of traffic mirror filter rule. Valid values: `tcp`, `udp`, `icmp`, `all`.",
			},
			"source_cidr_block": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The source cidr block of traffic mirror filter rule.",
			},
			"source_port_range": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Description: "The source port range of traffic mirror filter rule. When the protocol is `all` or `icmp`, the value is `-1/-1`. \n" +
					"When the protocol is `tcp` or `udp`, the value can be `1/200`, `80/80`, which means port 1 to port 200, port 80.",
			},
			"destination_cidr_block": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The destination cidr block of traffic mirror filter rule.",
			},
			"destination_port_range": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Description: "The destination port range of traffic mirror filter rule. When the protocol is `all` or `icmp`, the value is `-1/-1`. \n" +
					"When the protocol is `tcp` or `udp`, the value can be `1/200`, `80/80`, which means port 1 to port 200, port 80.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of traffic mirror filter rule.",
			},

			// computed fields
			"traffic_mirror_filter_rule_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of traffic mirror filter rule.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of traffic mirror filter rule.",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The create time of traffic mirror filter rule.",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The last update time of traffic mirror filter rule.",
			},
		},
	}
	return resource
}

func resourceVolcengineTrafficMirrorFilterRuleCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTrafficMirrorFilterRuleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineTrafficMirrorFilterRule())
	if err != nil {
		return fmt.Errorf("error on creating traffic_mirror_filter_rule %q, %s", d.Id(), err)
	}
	return resourceVolcengineTrafficMirrorFilterRuleRead(d, meta)
}

func resourceVolcengineTrafficMirrorFilterRuleRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTrafficMirrorFilterRuleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineTrafficMirrorFilterRule())
	if err != nil {
		return fmt.Errorf("error on reading traffic_mirror_filter_rule %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineTrafficMirrorFilterRuleUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTrafficMirrorFilterRuleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineTrafficMirrorFilterRule())
	if err != nil {
		return fmt.Errorf("error on updating traffic_mirror_filter_rule %q, %s", d.Id(), err)
	}
	return resourceVolcengineTrafficMirrorFilterRuleRead(d, meta)
}

func resourceVolcengineTrafficMirrorFilterRuleDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTrafficMirrorFilterRuleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineTrafficMirrorFilterRule())
	if err != nil {
		return fmt.Errorf("error on deleting traffic_mirror_filter_rule %q, %s", d.Id(), err)
	}
	return err
}

var trafficMirrorFilterRuleImporter = func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(data.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
	}
	if err := data.Set("traffic_mirror_filter_id", items[0]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	if err := data.Set("traffic_mirror_filter_rule_id", items[1]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}
