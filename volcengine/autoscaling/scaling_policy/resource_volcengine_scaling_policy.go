package scaling_policy

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
ScalingPolicy can be imported using the ScalingGroupId:ScalingPolicyId, e.g.
```
$ terraform import volcengine_scaling_policy.default scg-yblfbfhy7agh9zn72iaz:sp-yblf9l4fvcl8j1prohsp
```

*/

func ResourceVolcengineScalingPolicy() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineScalingPolicyCreate,
		Read:   resourceVolcengineScalingPolicyRead,
		Update: resourceVolcengineScalingPolicyUpdate,
		Delete: resourceVolcengineScalingPolicyDelete,
		Importer: &schema.ResourceImporter{
			State: scalingPolicyImporter,
		},
		Schema: map[string]*schema.Schema{
			"active": {
				Type:        schema.TypeBool,
				Computed:    true,
				Optional:    true,
				Description: "The active flag of the scaling policy. [Warning] the scaling policy can be active only when the scaling group be active otherwise will fail.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the scaling policy. Valid values: Active, InActive.",
			},
			"scaling_group_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the scaling group to which the scaling policy belongs.",
			},
			"scaling_policy_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The id of the scaling policy.",
			},
			"scaling_policy_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the scaling policy.",
			},
			"scaling_policy_type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"Scheduled", "Recurrence", "Alarm"}, false),
				Description:  "The type of scaling policy. Valid values: Scheduled, Recurrence, Alarm.",
			},
			"adjustment_type": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"QuantityChangeInCapacity", "PercentChangeInCapacity", "TotalCapacity"}, false),
				Description:  "The adjustment type of the scaling policy. Valid values: QuantityChangeInCapacity, PercentChangeInCapacity, TotalCapacity.",
			},
			"adjustment_value": {
				Type:     schema.TypeInt,
				Required: true,
				Description: "The adjustment value of the scaling policy. When the value of the `AdjustmentType` parameter is `QuantityChangeInCapacity`: -100 ~ 100, 0 is not allowed, unit: piece. " +
					"When the value of the `AdjustmentType` parameter is `PercentChangeInCapacity`: -100 ~ 10000, 0 is not allowed, unit: %. " +
					"When the value of the `AdjustmentType` parameter is `TotalCapacity`: the default is 0 to 100, unit: piece.",
			},
			"cooldown": {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.IntBetween(0, 86400),
				Description:  "The cooldown of the scaling policy. Default value is the cooldown time of the scaling group. Value: 0~86400, unit: second, if left blank, the cooling time of the scaling group will be used by default.",
			},
			"scheduled_policy_launch_time": {
				Type:             schema.TypeString,
				Optional:         true,
				Computed:         true,
				ValidateFunc:     timeValidation,
				DiffSuppressFunc: policyDiffSuppressFunc("Recurrence", "Scheduled"),
				Description: "The launch time of the scheduled policy of the scaling policy.\n" +
					"When the value of `ScalingPolicyType` is `Scheduled`, it means that the trigger time of the scheduled task must be greater than the current time.\n" +
					"When the value of `ScalingPolicyType` is `Recurrence`: If `ScheduledPolicy.RecurrenceType` is not specified, it means to execute only once according to the date and time specified here.\n" +
					"If `ScheduledPolicy.RecurrenceType` is specified, it indicates the start time of the periodic task. Only the time within 90 days from the date of creation/modification is supported.\n" +
					"When the value of `ScalingPolicyType` is `Alarm`, this parameter is invalid.",
			},
			"scheduled_policy_recurrence_end_time": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateFunc:     timeValidation,
				DiffSuppressFunc: policyDiffSuppressFunc("Recurrence"),
				Description: "The recurrence end time of the scheduled policy of the scaling policy. Valid and required when `ScalingPolicyType` is `Recurrence`. " +
					"If not configured, it will default to the day/week/month after this moment according to the recurrence period (ScheduledPolicy.RecurrenceType).",
			},
			"scheduled_policy_recurrence_type": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: policyDiffSuppressFunc("Recurrence"),
				ValidateFunc:     validation.StringInSlice([]string{"Daily", "Weekly", "Monthly", "Cron"}, false),
				Description:      "The recurrence type the scheduled policy of the scaling policy. Valid values: Daily, Weekly, Monthly, Cron.",
			},
			"scheduled_policy_recurrence_value": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: policyDiffSuppressFunc("Recurrence"),
				Description: "The recurrence value the scheduled policy of the scaling policy. Valid and required when `ScalingPolicyType` is `Recurrence`. " +
					"When the value of the ScheduledPolicy.RecurrenceType parameter is Daily, only one value can be filled in, ranging from 1 to 31.\n" +
					"When the value of the ScheduledPolicy.RecurrenceType parameter is Weekly, you can enter multiple values separated by commas (,). The values from Monday to Sunday are: 1,2,3,4,5,6,7.\n" +
					"When the value of the ScheduledPolicy.RecurrenceType parameter is Monthly, the format is A-B. The value ranges of A and B are both 1~31, and B must be greater than or equal to A.\n" +
					"When the value of the ScheduledPolicy.RecurrenceType parameter is Cron, it means UTC+8 time, supports 5-field expressions of minutes, hours, days, months, and weeks, " +
					"and supports wildcard English commas (,), English question marks (?), and conjunctions ( -), asterisk (*), pound sign (#), slash (/), L, and W.",
			},
			"alarm_policy_rule_type": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "Static",
				ValidateFunc:     validation.StringInSlice([]string{"Static"}, false),
				DiffSuppressFunc: policyDiffSuppressFunc("Alarm"),
				Description:      "The rule type of the alarm policy of the scaling policy. Valid value: Static. It is only valid and required when the value of `ScalingPolicyType` is `Alarm`.",
			},
			"alarm_policy_evaluation_count": {
				Type:             schema.TypeInt,
				Optional:         true,
				DiffSuppressFunc: policyDiffSuppressFunc("Alarm"),
				ValidateFunc:     validation.IntBetween(1, 180),
				Description:      "The evaluation count of the alarm policy of the scaling policy. It is only valid and required when the value of `ScalingPolicyType` is `Alarm`.",
			},
			"alarm_policy_condition_metric_name": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: policyDiffSuppressFunc("Alarm"),
				ValidateFunc: validation.StringInSlice([]string{
					"CpuTotal_Max",
					"CpuTotal_Min",
					"CpuTotal_Avg",
					"MemoryUsedUtilization_Max",
					"MemoryUsedUtilization_Min",
					"MemoryUsedUtilization_Avg",
					"Instance_CpuBusy_Max",
					"Instance_CpuBusy_Min",
					"Instance_CpuBusy_Avg",
				}, false),
				Description: "The metric name of the alarm policy condition of the scaling policy. Valid values: CpuTotal_Max, CpuTotal_Min, CpuTotal_Avg, MemoryUsedUtilization_Max, MemoryUsedUtilization_Min, MemoryUsedUtilization_Avg, Instance_CpuBusy_Max, Instance_CpuBusy_Min, Instance_CpuBusy_Avg.",
			},
			"alarm_policy_condition_metric_unit": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: policyDiffSuppressFunc("Alarm"),
				ValidateFunc:     validation.StringInSlice([]string{"Percent"}, false),
				Description:      "The comparison operator of the alarm policy condition of the scaling policy. It is only valid and required when the value of `ScalingPolicyType` is `Alarm`.",
			},
			"alarm_policy_condition_comparison_operator": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: policyDiffSuppressFunc("Alarm"),
				ValidateFunc:     validation.StringInSlice([]string{">", "<", "="}, false),
				Description:      "The comparison operator of the alarm policy condition of the scaling policy. Valid values: `>`, `<`, `=`. It is only valid and required when the value of `ScalingPolicyType` is `Alarm`.",
			},
			"alarm_policy_condition_threshold": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: policyDiffSuppressFunc("Alarm"),
				Description:      "The threshold of the alarm policy condition of the scaling policy. It is only valid and required when the value of `ScalingPolicyType` is `Alarm`.",
			},
		},
	}
	return resource
}

func resourceVolcengineScalingPolicyCreate(d *schema.ResourceData, meta interface{}) (err error) {
	scalingPolicyService := NewScalingPolicyService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(scalingPolicyService, d, ResourceVolcengineScalingPolicy())
	if err != nil {
		return fmt.Errorf("error on creating ScalingPolicy %q, %s", d.Id(), err)
	}
	return resourceVolcengineScalingPolicyRead(d, meta)
}

func resourceVolcengineScalingPolicyRead(d *schema.ResourceData, meta interface{}) (err error) {
	scalingPolicyService := NewScalingPolicyService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(scalingPolicyService, d, ResourceVolcengineScalingPolicy())
	if err != nil {
		return fmt.Errorf("error on reading ScalingPolicy %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineScalingPolicyUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	scalingPolicyService := NewScalingPolicyService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(scalingPolicyService, d, ResourceVolcengineScalingPolicy())
	if err != nil {
		return fmt.Errorf("error on updating ScalingPolicy %q, %s", d.Id(), err)
	}
	return resourceVolcengineScalingPolicyRead(d, meta)
}

func resourceVolcengineScalingPolicyDelete(d *schema.ResourceData, meta interface{}) (err error) {
	scalingPolicyService := NewScalingPolicyService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(scalingPolicyService, d, ResourceVolcengineScalingPolicy())
	if err != nil {
		return fmt.Errorf("error on deleting ScalingPolicy %q, %s", d.Id(), err)
	}
	return err
}
