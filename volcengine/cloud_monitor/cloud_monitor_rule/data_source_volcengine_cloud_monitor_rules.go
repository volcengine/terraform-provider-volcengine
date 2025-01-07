package cloud_monitor_rule

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineCloudMonitorRules() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineCloudMonitorRulesRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:           schema.HashString,
				ConflictsWith: []string{"rule_name", "alert_state", "namespace", "level", "enable_state"},
				Description:   "A list of cloud monitor ids.",
			},
			"rule_name": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"ids"},
				Description:   "The name of the cloud monitor rule. This field support fuzzy query.",
			},
			"alert_state": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:           schema.HashString,
				ConflictsWith: []string{"ids"},
				Description:   "The alert state of the cloud monitor rule. Valid values: `altering`, `normal`.",
			},
			"namespace": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:           schema.HashString,
				ConflictsWith: []string{"ids"},
				Description:   "The namespace of the cloud monitor rule.",
			},
			"level": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:           schema.HashString,
				ConflictsWith: []string{"ids"},
				Description:   "The level of the cloud monitor rule. Valid values: `critical`, `warning`, `notice`.",
			},
			"enable_state": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:           schema.HashString,
				ConflictsWith: []string{"ids"},
				Description:   "The enable state of the cloud monitor rule. Valid values: `enable`, `disable`.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of Resource.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of query.",
			},
			"rules": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the cloud monitor rule.",
						},
						"rule_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the cloud monitor rule.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the cloud monitor rule.",
						},
						"namespace": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The namespace of the cloud monitor rule.",
						},
						"sub_namespace": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The sub namespace of the cloud monitor rule.",
						},
						"web_hook": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The web hook of the cloud monitor rule.",
						},
						"webhook_ids": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The webhook id list of the cloud monitor rule.",
						},
						"alert_state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The alert state of the cloud monitor rule.",
						},
						"level": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The level of the cloud monitor rule.",
						},
						"enable_state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The enable state of the cloud monitor rule.",
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
						"effect_start_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The effect start time of the cloud monitor rule.",
						},
						"effect_end_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The effect end time of the cloud monitor rule.",
						},
						"evaluation_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The evaluation count of the cloud monitor rule.",
						},
						"silence_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The silence time of the cloud monitor rule. Unit in minutes.",
						},
						"multiple_conditions": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to enable the multiple conditions function of the cloud monitor rule.",
						},
						"condition_operator": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The condition operator of the cloud monitor rule. Valid values: `&&`, `||`.",
						},
						"regions": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The region id of the cloud monitor rule.",
						},
						"contact_group_ids": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The contact group ids of the cloud monitor rule.",
						},
						"alert_methods": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The alert methods of the cloud monitor rule.",
						},
						"conditions": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The conditions of the cloud monitor rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"metric_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The metric name of the cloud monitor rule.",
									},
									"metric_unit": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The metric unit of the cloud monitor rule.",
									},
									"statistics": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The statistics of the cloud monitor rule.",
									},
									"comparison_operator": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The comparison operation of the cloud monitor rule.",
									},
									"threshold": {
										Type:        schema.TypeString,
										Computed:    true,
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
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The original dimensions of the cloud monitor rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The key of the dimension.",
									},
									"value": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
										Description: "The value of the dimension.",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineCloudMonitorRulesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewCloudMonitorRuleService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineCloudMonitorRules())
}
