package cloud_monitor_rule

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
CloudMonitorRule can be imported using the id, e.g.
```
$ terraform import volcengine_cloud_monitor_rule.default 174284623567451****
```

*/

func ResourceVolcengineCloudMonitorRule() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineCloudMonitorRuleCreate,
		Read:   resourceVolcengineCloudMonitorRuleRead,
		Update: resourceVolcengineCloudMonitorRuleUpdate,
		Delete: resourceVolcengineCloudMonitorRuleDelete,
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
				Description: "The name of the cloud monitor rule.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The description of the cloud monitor rule.",
			},
			"namespace": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The namespace of the cloud monitor rule.",
			},
			"sub_namespace": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The sub namespace of the cloud monitor rule.",
			},
			"level": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"critical", "warning", "notice"}, false),
				Description:  "The level of the cloud monitor rule. Valid values: `critical`, `warning`, `notice`.",
			},
			"enable_state": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"enable", "disable"}, false),
				Description:  "The enable state of the cloud monitor rule. Valid values: `enable`, `disable`.",
			},
			"evaluation_count": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The evaluation count of the cloud monitor rule.",
			},
			"effect_start_at": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The effect start time of the cloud monitor rule. The expression is `HH:MM`.",
			},
			"effect_end_at": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The effect end time of the cloud monitor rule. The expression is `HH:MM`.",
			},
			"silence_time": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The silence time of the cloud monitor rule. Unit in minutes. Valid values: 5, 30, 60, 180, 360, 720, 1440.",
			},
			"multiple_conditions": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Whether to enable the multiple conditions function of the cloud monitor rule.",
			},
			"condition_operator": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The condition operator of the cloud monitor rule. Valid values: `&&`, `||`.",
			},
			"alert_methods": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The alert methods of the cloud monitor rule. Valid values: `Email`, `Phone`, `SMS`, `Webhook`.",
			},
			"web_hook": {
				Type:         schema.TypeString,
				Optional:     true,
				AtLeastOneOf: []string{"web_hook", "contact_group_ids"},
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if methods, ok := d.GetOk("alert_methods"); ok {
						methodArr := methods.(*schema.Set).List()
						if contains("Webhook", methodArr) {
							return false
						}
					}
					return true
				},
				Description: "The web hook of the cloud monitor rule. When the alert method is `Webhook`, This field must be specified.",
			},
			"contact_group_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:          schema.HashString,
				AtLeastOneOf: []string{"web_hook", "contact_group_ids"},
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if methods, ok := d.GetOk("alert_methods"); ok {
						methodsArr := methods.(*schema.Set).List()
						if contains("Phone", methodsArr) || contains("Email", methodsArr) ||
							contains("SMS", methodsArr) {
							return false
						}
					}
					return true
				},
				Description: "The contact group ids of the cloud monitor rule. When the alert method is `Email`, `SMS`, or `Phone`, This field must be specified.",
			},
			"recovery_notify": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Description: "The recovery notify of the cloud monitor rule.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     true,
							Description: "Whether to enable the recovery notify function.",
						},
					},
				},
			},
			"regions": {
				Type:     schema.TypeSet,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The region ids of the cloud monitor rule. Only one region id can be specified currently.",
			},
			"conditions": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "The conditions of the cloud monitor rule.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"metric_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The metric name of the cloud monitor rule.",
						},
						"metric_unit": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The metric unit of the cloud monitor rule.",
						},
						"statistics": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The statistics of the cloud monitor rule. Valid values: `avg`, `max`, `min`.",
						},
						"comparison_operator": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The comparison operation of the cloud monitor rule. Valid values: `>`, `>=`, `<`, `<=`, `!=`, `=`.",
						},
						"threshold": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The threshold of the cloud monitor rule.",
						},
						"period": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The period of the cloud monitor rule.",
						},
					},
				},
			},
			"original_dimensions": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "The original dimensions of the cloud monitor rule.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The key of the dimension.",
						},
						"value": {
							Type:     schema.TypeSet,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Set:         schema.HashString,
							Description: "The value of the dimension.",
						},
					},
				},
			},

			// computed fields
			"alert_state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The alert state of the cloud monitor rule.",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The created time of the cloud monitor rule.",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The updated time of the cloud monitor rule.",
			},
		},
	}
	return resource
}

func resourceVolcengineCloudMonitorRuleCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCloudMonitorRuleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineCloudMonitorRule())
	if err != nil {
		return fmt.Errorf("error on creating cloud_monitor_rule %q, %s", d.Id(), err)
	}
	return resourceVolcengineCloudMonitorRuleRead(d, meta)
}

func resourceVolcengineCloudMonitorRuleRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCloudMonitorRuleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineCloudMonitorRule())
	if err != nil {
		return fmt.Errorf("error on reading cloud_monitor_rule %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineCloudMonitorRuleUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCloudMonitorRuleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineCloudMonitorRule())
	if err != nil {
		return fmt.Errorf("error on updating cloud_monitor_rule %q, %s", d.Id(), err)
	}
	return resourceVolcengineCloudMonitorRuleRead(d, meta)
}

func resourceVolcengineCloudMonitorRuleDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCloudMonitorRuleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineCloudMonitorRule())
	if err != nil {
		return fmt.Errorf("error on deleting cloud_monitor_rule %q, %s", d.Id(), err)
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
