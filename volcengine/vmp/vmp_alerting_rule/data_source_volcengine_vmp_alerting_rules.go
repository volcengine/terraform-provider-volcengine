package vmp_alerting_rule

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineVmpAlertingRules() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineVmpAlertingRulesRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of vmp alerting rule IDs.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of vmp alerting rule. This field support fuzzy query.",
			},
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The type of vmp alerting rule. Valid values: `vmp/PromQL`.",
			},
			"workspace_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The workspace id of vmp alerting rule.",
			},
			"notify_policy_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of notify policy IDs.",
			},
			"notify_group_policy_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of notify group policy IDs.",
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"Running", "Disabled",
				}, false),
				Description: "The status of vmp alerting rule. Valid values: `Running`, `Disabled`.",
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
			"alerting_rules": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the vmp alerting rule.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the vmp alerting rule.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the vmp alerting rule.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the vmp alerting rule.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of the vmp alerting rule.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the vmp alerting rule.",
						},
						"notify_policy_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The notify policy id of the vmp alerting rule.",
						},
						"notify_group_policy_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The notify group policy id of the vmp alerting rule.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the vmp alerting rule.",
						},
						"annotations": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "The annotations of the vmp alerting rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the annotation.",
									},
									"value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The value of the annotation.",
									},
								},
							},
						},
						"labels": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "The labels of the vmp alerting rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the label.",
									},
									"value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The value of the label.",
									},
								},
							},
						},
						"query": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The alerting query of the vmp alerting rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"workspace_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The id of the workspace.",
									},
									"prom_ql": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The prom ql of query.",
									},
								},
							},
						},
						"levels": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The alerting levels of the vmp alerting rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"level": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The level of the vmp alerting rule.",
									},
									"for": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The duration of the alerting rule.",
									},
									"comparator": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The comparator of the vmp alerting rule.",
									},
									"threshold": {
										Type:        schema.TypeFloat,
										Computed:    true,
										Description: "The threshold of the vmp alerting rule.",
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

func dataSourceVolcengineVmpAlertingRulesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewVmpAlertingRuleService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineVmpAlertingRules())
}
