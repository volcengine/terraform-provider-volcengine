package cloud_monitor_event_rule

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
CloudMonitorEventRule can be imported using the id, e.g.
```
$ terraform import volcengine_cloud_monitor_event_rule.default rule_id
```

*/

func ResourceVolcengineCloudMonitorEventRule() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineCloudMonitorEventRuleCreate,
		Read:   resourceVolcengineCloudMonitorEventRuleRead,
		Update: resourceVolcengineCloudMonitorEventRuleUpdate,
		Delete: resourceVolcengineCloudMonitorEventRuleDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"rule_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the rule.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the rule.",
			},
			"event_source": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Event source.",
			},
			"event_type": {
				Optional:    true,
				Type:        schema.TypeSet,
				Set:         schema.HashString,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Event type.",
			},
			"status": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "enable",
				Description: "Rule status. `enable`: enable rule(default), `disable`: disable rule.",
			},
			"level": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Severity of alarm rules. Value can be `notice`, `warning`, `critical`.",
			},
			"filter_pattern": {
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				Description: "Filter mode, also known as event matching rules. Custom matching rules are not currently supported.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Required:    true,
							Type:        schema.TypeSet,
							Set:         schema.HashString,
							Description: "The list of corresponding event types in pattern matching, currently set to match any.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"source": {
							Type:        schema.TypeString,
							Required:    true,
							ForceNew:    true,
							Description: "Event source corresponding to pattern matching.",
						},
					},
				},
			},
			"effective_time": {
				Type:        schema.TypeList,
				Description: "The rule takes effect at a certain time and will only be effective during this period.",
				Required:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"start_time": {
							Required:    true,
							Type:        schema.TypeString,
							Description: "Start time for rule activation.",
						},
						"end_time": {
							Required:    true,
							Type:        schema.TypeString,
							Description: "End time for rule activation.",
						},
					},
				},
			},
			"contact_methods": {
				Required: true,
				Type:     schema.TypeSet,
				Set:      schema.HashString,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Description: "Alarm notification methods. Valid value: `Phone`, `Email`, `SMS`, " +
					"`Webhook`: Alarm callback, `TLS`: Log Service, `MQ`: Message Queue Kafka.",
			},
			"contact_group_ids": {
				Optional: true,
				Type:     schema.TypeSet,
				Set:      schema.HashString,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if methods, ok := d.GetOk("contact_methods"); ok {
						methodsArr := methods.(*schema.Set).List()
						if contains("Phone", methodsArr) || contains("Email", methodsArr) ||
							contains("SMS", methodsArr) {
							return false
						}
					}
					return true
				},
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "When the alarm notification method is phone, SMS, or email, the triggered alarm contact group ID.",
			},
			"endpoint": {
				Optional: true,
				Type:     schema.TypeString,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if methods, ok := d.GetOk("contact_methods"); ok {
						methodArr := methods.(*schema.Set).List()
						if contains("Webhook", methodArr) {
							return false
						}
					}
					return true
				},
				Description: "When the alarm notification method is alarm callback, it triggers the callback address.",
			},
			"tls_target": {
				Optional: true,
				Type:     schema.TypeSet,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if methods, ok := d.GetOk("contact_methods"); ok {
						methodArr := methods.(*schema.Set).List()
						if contains("TLS", methodArr) {
							return false
						}
					}
					return true
				},
				Description: "The alarm method for log service triggers the configuration of the log service.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"region_name_en": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The English region name.",
						},
						"region_name_cn": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The Chinese region name.",
						},
						"project_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The project name.",
						},
						"project_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The project id.",
						},
						"topic_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The topic id.",
						},
					},
				},
			},
			"message_queue": {
				Optional: true,
				Type:     schema.TypeSet,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if methods, ok := d.GetOk("contact_methods"); ok {
						methodArr := methods.(*schema.Set).List()
						if contains("MQ", methodArr) {
							return false
						}
					}
					return true
				},
				Description: "The triggered message queue when the alarm notification method is Kafka message queue.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The kafka instance id.",
						},
						"region": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The region.",
						},
						"topic": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The topic name.",
						},
						"type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The message queue type, only support kafka now.",
						},
						"vpc_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The vpc id.",
						},
					},
				},
			},
		},
	}
	return resource
}

func resourceVolcengineCloudMonitorEventRuleCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCloudMonitorEventRuleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineCloudMonitorEventRule())
	if err != nil {
		return fmt.Errorf("error on creating cloud_monitor_event_rule %q, %s", d.Id(), err)
	}
	return resourceVolcengineCloudMonitorEventRuleRead(d, meta)
}

func resourceVolcengineCloudMonitorEventRuleRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCloudMonitorEventRuleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineCloudMonitorEventRule())
	if err != nil {
		return fmt.Errorf("error on reading cloud_monitor_event_rule %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineCloudMonitorEventRuleUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCloudMonitorEventRuleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineCloudMonitorEventRule())
	if err != nil {
		return fmt.Errorf("error on updating cloud_monitor_event_rule %q, %s", d.Id(), err)
	}
	return resourceVolcengineCloudMonitorEventRuleRead(d, meta)
}

func resourceVolcengineCloudMonitorEventRuleDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCloudMonitorEventRuleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineCloudMonitorEventRule())
	if err != nil {
		return fmt.Errorf("error on deleting cloud_monitor_event_rule %q, %s", d.Id(), err)
	}
	return err
}

func contains(target string, arr []interface{}) bool {
	for _, v := range arr {
		if target == v.(string) {
			return true
		}
	}
	return false
}
