package vmp_silence_policy

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
VmpSilencePolicy can be imported using the id, e.g.
```
$ terraform import volcengine_vmp_silence_policy.default resource_id
```

*/

func ResourceVolcengineVmpSilencePolicy() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVmpSilencePolicyCreate,
		Read:   resourceVolcengineVmpSilencePolicyRead,
		Update: resourceVolcengineVmpSilencePolicyUpdate,
		Delete: resourceVolcengineVmpSilencePolicyDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the silence policy.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the silence policy.",
			},
			"metric_label_matchers": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 32,
				Description: "Alarm event Label matcher, allowing a maximum of 32 entries, with an \"OR\" relationship between different entries. " +
					"Different label_matchers in the Matcher follow an \"AND\" relationship.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"matchers": {
							Type:        schema.TypeSet,
							Required:    true,
							MaxItems:    24,
							Description: "Label matcher. Among them, each LabelMatcher array can contain a maximum of 24 items.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"label": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Label.",
									},
									"value": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Label value.",
									},
									"operator": {
										Type:     schema.TypeString,
										Optional: true,
										ValidateFunc: validation.StringInSlice([]string{
											"Equal",
											"NotEqual",
											"RegexpEqual",
											"RegexpNotEqual",
										}, false),
										Description: "Operator. The optional values are as follows: Equal, NotEqual, RegexpEqual, RegexpNotEqual.",
									},
								},
							},
						},
					},
				},
			},
			"time_range_matchers": {
				Type:     schema.TypeList,
				Required: true,
				Description: "Alarm silence time. Case 1: Always effective. When the switch is turned on, the matched alarm events are always silenced, and only the location needs to be set. " +
					"Case 2: Periodic effective. When the switch is turned on, the matched alarm events are silenced periodically, and the location and periodic_date need to be set. " +
					"Case 3: Custom effective. When the switch is turned on, the matched alarm events are silenced in the specified time range, and the location and date need to be set.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"date": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Silence time range, like 2025-01-02 15:04~2025-01-03 14:04.",
						},
						"location": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Timezone, e.g. Asia/Shanghai.",
						},
						"periodic_date": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "The cycle of alarm silence. It is used to configure alarm silence that takes effect periodically.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"time": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Time periods, e.g. 20:00~21:12,22:00~23:12. A maximum of 4 time periods can be configured.",
									},
									"weekday": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Weekdays, e.g. 1,3,5. A maximum of 7 time periods can be configured.",
									},
									"day_of_month": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Days of month, e.g. 2~3. A maximum of 10 time periods can be configured.",
									},
								},
							},
						},
					},
				},
			},
			// computed fields
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The create time of the silence policy.",
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The update time of the silence policy.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the silence policy.",
			},
			"source": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The source of the silence policy.",
			},
			"auto_delete_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The auto delete time of the silence policy.",
			},
		},
	}
	return resource
}

// resourceVolcengineVmpSilencePolicyCreate 创建静默策略
func resourceVolcengineVmpSilencePolicyCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVmpSilencePolicyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineVmpSilencePolicy())
	if err != nil {
		return fmt.Errorf("error on creating vmp_silence_policy %q, %s", d.Id(), err)
	}
	return resourceVolcengineVmpSilencePolicyRead(d, meta)
}

// resourceVolcengineVmpSilencePolicyRead 读取静默策略
func resourceVolcengineVmpSilencePolicyRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVmpSilencePolicyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineVmpSilencePolicy())
	if err != nil {
		return fmt.Errorf("error on reading vmp_silence_policy %q, %s", d.Id(), err)
	}
	return err
}

// resourceVolcengineVmpSilencePolicyUpdate 更新静默策略
func resourceVolcengineVmpSilencePolicyUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVmpSilencePolicyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineVmpSilencePolicy())
	if err != nil {
		return fmt.Errorf("error on updating vmp_silence_policy %q, %s", d.Id(), err)
	}
	return resourceVolcengineVmpSilencePolicyRead(d, meta)
}

// resourceVolcengineVmpSilencePolicyDelete 删除静默策略
func resourceVolcengineVmpSilencePolicyDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVmpSilencePolicyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineVmpSilencePolicy())
	if err != nil {
		return fmt.Errorf("error on deleting vmp_silence_policy %q, %s", d.Id(), err)
	}
	return err
}
