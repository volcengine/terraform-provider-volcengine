package vmp_alert

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineVmpAlerts() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineVmpAlertsRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of vmp alert IDs.",
			},
			"current_phase": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"Pending", "Active", "Resolved", "Disabled",
				}, false),
				Description: "The status of vmp alert. Valid values: `Pending`, `Active`, `Resolved`, `Disabled`.",
			},
			"level": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"P0", "P1", "P2",
				}, false),
				Description: "The level of vmp alert. Valid values: `P0`, `P1`, `P2`.",
			},
			"alerting_rule_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of alerting rule IDs.",
			},
			"desc": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether to use descending sorting.",
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
			"alerts": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the vmp alert.",
						},
						"alerting_rule_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the vmp alerting rule.",
						},
						"initial_alert_timestamp": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The start time of the vmp alert. Format: RFC3339.",
						},
						"last_alert_timestamp": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The last time of the vmp alert. Format: RFC3339.",
						},
						"resolve_alert_timestamp": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The end time of the vmp alert. Format: RFC3339.",
						},
						"current_phase": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the vmp alert.",
						},
						"current_level": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The current level of the vmp alert.",
						},
						"alerting_rule_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the vmp alerting rule.",
						},
						"alerting_rule_query": {
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
							Description: "The alerting levels of the vmp alert.",
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
						"resource": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The alerting resource of the vmp alert.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"labels": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The labels of alerting resource.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"key": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The key of the label.",
												},
												"value": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The value of the label.",
												},
											},
										},
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

func dataSourceVolcengineVmpAlertsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewVmpAlertService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineVmpAlerts())
}
