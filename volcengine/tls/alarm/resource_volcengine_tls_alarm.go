package alarm

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
tls alarm can be imported using the id and project id, e.g.
```
$ terraform import volcengine_tls_alarm.default projectId:fc************
```

*/

func ResourceVolcengineTlsAlarm() *schema.Resource {
	return &schema.Resource{
		Read:   resourceVolcengineTlsAlarmRead,
		Create: resourceVolcengineTlsAlarmCreate,
		Update: resourceVolcengineTlsAlarmUpdate,
		Delete: resourceVolcengineTlsAlarmDelete,
		Importer: &schema.ResourceImporter{
			State: importTlsAlarmApply,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"alarm_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the alarm.",
			},
			"project_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The project id.",
			},
			"status": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Whether to enable the alert policy. The default value is true, that is, on.",
			},
			"trigger_period": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      1,
				ValidateFunc: validation.IntBetween(1, 10),
				Description:  "Continuous cycle. The alarm will be issued after the trigger condition is continuously met for TriggerPeriod periods; the minimum value is 1, the maximum value is 10, and the default value is 1.",
			},
			"alarm_period": {
				Type:         schema.TypeInt,
				Optional:     true,
				ValidateFunc: validation.IntBetween(10, 1440),
				ExactlyOneOf: []string{"alarm_period_detail", "alarm_period"},
				Description:  "Period for sending alarm notifications. When the number of continuous alarm triggers reaches the specified limit (TriggerPeriod), Log Service will send alarm notifications according to the specified period.",
			},
			"alarm_notify_group": {
				Type:        schema.TypeSet,
				Required:    true,
				ForceNew:    true,
				Set:         schema.HashString,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "List of notification groups corresponding to the alarm.",
			},
			"user_define_msg": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Customize the alarm notification content.",
			},
			"query_request": {
				Type:     schema.TypeSet,
				Required: true,
				MaxItems: 3,
				MinItems: 1,
				Set: func(i interface{}) int {
					if i == nil {
						return hashcode.String("")
					}
					m := i.(map[string]interface{})
					var (
						buf bytes.Buffer
					)
					buf.WriteString(fmt.Sprintf("%v#%v#%v#%v#%v", m["topic_id"], m["query"], m["number"], m["start_time_offset"], m["end_time_offset"]))
					return hashcode.String(buf.String())
				},
				Description: "Search and analyze sentences, 1~3 can be configured.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"topic_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The id of the topic.",
						},
						"query": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Query statement, the maximum supported length is 1024.",
						},
						"number": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Alarm object sequence number; increments from 1.",
						},
						"start_time_offset": {
							Type:         schema.TypeInt,
							Required:     true,
							ValidateFunc: validation.IntBetween(-1440, 0),
							Description:  "The start time of the query range is relative to the current historical time, in minutes. The value is non-positive, the maximum value is 0, and the minimum value is -1440.",
						},
						"end_time_offset": {
							Type:         schema.TypeInt,
							Required:     true,
							ValidateFunc: validation.IntBetween(-1440, 0),
							Description:  "The end time of the query range is relative to the current historical time. The unit is minutes. The value is not positive and must be greater than StartTimeOffset. The maximum value is 0 and the minimum value is -1440.",
						},
					},
				},
			},
			"condition": {
				Required:    true,
				Type:        schema.TypeString,
				Description: "Alarm trigger condition.",
			},
			"request_cycle": {
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				Description: "The execution period of the alarm task.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringInSlice([]string{"Period", "Fixed"}, false),
							Description:  "Execution cycle type.",
						},
						"time": {
							Type:         schema.TypeInt,
							Required:     true,
							ValidateFunc: validation.IntBetween(1, 1440),
							Description:  "The cycle of alarm task execution, or the time point of periodic execution. The unit is minutes, and the value range is 1~1440.",
						},
					},
				},
			},
			"alarm_period_detail": {
				Type:         schema.TypeList,
				Optional:     true,
				MaxItems:     1,
				ExactlyOneOf: []string{"alarm_period_detail", "alarm_period"},
				Description:  "Period for sending alarm notifications. When the number of continuous alarm triggers reaches the specified limit (TriggerPeriod), Log Service will send alarm notifications according to the specified period.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"sms": {
							Type:         schema.TypeInt,
							Required:     true,
							ValidateFunc: validation.IntBetween(10, 1440),
							Description:  "SMS alarm cycle, the unit is minutes, and the value range is 10~1440.",
						},
						"phone": {
							Type:         schema.TypeInt,
							Required:     true,
							ValidateFunc: validation.IntBetween(10, 1440),
							Description:  "Telephone alarm cycle, the unit is minutes, and the value range is 10~1440.",
						},
						"email": {
							Type:         schema.TypeInt,
							Required:     true,
							ValidateFunc: validation.IntBetween(1, 1440),
							Description:  "Email alarm period, the unit is minutes, and the value range is 1~1440.",
						},
						"general_webhook": {
							Type:         schema.TypeInt,
							Required:     true,
							ValidateFunc: validation.IntBetween(1, 1440),
							Description:  "Customize the webhook alarm period, the unit is minutes, and the value range is 1~1440.",
						},
					},
				},
			},
			"alarm_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The alarm id.",
			},
		},
	}
}

func resourceVolcengineTlsAlarmCreate(d *schema.ResourceData, meta interface{}) error {
	TlsAlarmService := NewVolcengineTlsAlarmService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Create(TlsAlarmService, d, ResourceVolcengineTlsAlarm()); err != nil {
		return fmt.Errorf("error on creating tls Alarm  %q, %w", d.Id(), err)
	}
	return resourceVolcengineTlsAlarmRead(d, meta)
}

func resourceVolcengineTlsAlarmRead(d *schema.ResourceData, meta interface{}) error {
	TlsAlarmService := NewVolcengineTlsAlarmService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Read(TlsAlarmService, d, ResourceVolcengineTlsAlarm()); err != nil {
		return fmt.Errorf("error on reading tls Alarm %q, %w", d.Id(), err)
	}
	return nil
}

func resourceVolcengineTlsAlarmUpdate(d *schema.ResourceData, meta interface{}) error {
	TlsAlarmService := NewVolcengineTlsAlarmService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Update(TlsAlarmService, d, ResourceVolcengineTlsAlarm()); err != nil {
		return fmt.Errorf("error on updating tls Alarm %q, %w", d.Id(), err)
	}
	return resourceVolcengineTlsAlarmRead(d, meta)
}

func resourceVolcengineTlsAlarmDelete(d *schema.ResourceData, meta interface{}) error {
	TlsAlarmService := NewVolcengineTlsAlarmService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Delete(TlsAlarmService, d, ResourceVolcengineTlsAlarm()); err != nil {
		return fmt.Errorf("error on deleting tls Alarm %q, %w", d.Id(), err)
	}
	return nil
}

func importTlsAlarmApply(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	var err error
	items := strings.Split(data.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must be of the form project_id:alarm_id")
	}
	err = data.Set("project_id", items[0])
	if err != nil {
		return []*schema.ResourceData{data}, err
	}
	err = data.Set("alarm_id", items[1])
	if err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}
