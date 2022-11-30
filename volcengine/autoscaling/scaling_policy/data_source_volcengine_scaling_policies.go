package scaling_policy

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineScalingPolicies() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineScalingPoliciesRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of scaling policy ids.",
			},
			"scaling_policy_names": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of scaling policy names.",
			},
			"scaling_group_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "An id of the scaling group to which the scaling policy belongs.",
			},
			"scaling_policy_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"Scheduled", "Recurrence", "Manual", "Alarm"}, false),
				Description:  "A type of scaling policy. Valid values: Scheduled, Recurrence, Manual, Alarm.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of scaling policy.",
			},

			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of scaling policy query.",
			},
			"scaling_policies": {
				Description: "The collection of scaling policy query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the scaling policy.",
						},
						"scaling_policy_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the scaling policy.",
						},
						"scaling_group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the scaling group to which the scaling policy belongs.",
						},
						"scaling_policy_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the scaling policy.",
						},
						"scaling_policy_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the scaling policy.",
						},
						"adjustment_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The adjustment type of the scaling policy.",
						},
						"adjustment_value": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The adjustment value of the scaling policy.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the scaling policy.",
						},
						"cooldown": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The cooldown of the scaling policy.",
						},
						"scheduled_policy_launch_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The launch time of the scheduled policy of the scaling policy.",
						},
						"scheduled_policy_recurrence_start_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The recurrence start time of the scheduled policy of the scaling policy.",
						},
						"scheduled_policy_recurrence_end_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The recurrence end time of the scheduled policy of the scaling policy.",
						},
						"scheduled_policy_recurrence_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The recurrence type of the scheduled policy of the scaling policy.",
						},
						"scheduled_policy_recurrence_value": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The recurrence value of the scheduled policy of the scaling policy.",
						},
						"alarm_policy_rule_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The rule type of the alarm policy of the scaling policy.",
						},
						"alarm_policy_evaluation_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The evaluation count of the alarm policy of the scaling policy.",
						},
						"alarm_policy_condition_metric_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The metric name of the alarm policy condition of the scaling policy.",
						},
						"alarm_policy_condition_metric_unit": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The comparison operator of the alarm policy condition of the scaling policy.",
						},
						"alarm_policy_condition_comparison_operator": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The comparison operator of the alarm policy condition of the scaling policy.",
						},
						"alarm_policy_condition_threshold": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The threshold of the alarm policy condition of the scaling policy.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineScalingPoliciesRead(d *schema.ResourceData, meta interface{}) error {
	scalingPolicyService := NewScalingPolicyService(meta.(*ve.SdkClient))
	return scalingPolicyService.Dispatcher.Data(scalingPolicyService, d, DataSourceVolcengineScalingPolicies())
}
