package waf_cc_rule

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineWafCcRules() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineWafCcRulesRead,
		Schema: map[string]*schema.Schema{
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
			"cc_type": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Set:         schema.HashInt,
				Description: "The actions performed on subsequent requests after meeting the statistical conditions.",
			},
			"host": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Website domain names that require the setting of protection rules.",
			},
			"path_order_by": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The list shows the order.",
			},
			"rule_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Search by rule name in a fuzzy manner.",
			},
			"rule_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Search precisely according to the rule ID.",
			},
			"url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Fuzzy search by the requested path.",
			},
			"data": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The requested path.",
						},
						"enable_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The total number of enabled rules within the rule group.",
						},
						"insert_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The creation time of the rule group.",
						},
						"total_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The total number of rules within the rule group.",
						},
						"rule_group": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Details of the rule group.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"group": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Computed:    true,
										Description: "Rule group information.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The ID of Rule group.",
												},
												"accurate_group_priority": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "After the rule creation is completed, the priority of the automatically generated rule group.",
												},
												"logic": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "In the rule group, the high-level conditional operation relationships corresponding to each rule.",
												},
												"accurate_rules": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Request characteristic information of the rule group.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"http_obj": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Custom object.",
															},
															"obj_type": {
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "matching field.",
															},
															"opretar": {
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The logical operator for the condition.",
															},
															"property": {
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "Operate the properties of the http object.",
															},
															"value_string": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The value to be matched.",
															},
														},
													},
												},
											},
										},
									},
									"rules": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specific rule information within the rule group.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The ID of Rule group.",
												},
												"name": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The Name of Rule group.",
												},
												"host": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Protected website domain names.",
												},
												"url": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Request path.",
												},
												"enable": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Whether the rule is enabled.",
												},
												"exemption_time": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Strategy exemption time.",
												},
												"cron_enable": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Whether to set the cycle to take effect.",
												},
												"single_threshold": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The threshold of the number of visits to each statistical object.",
												},
												"path_threshold": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The threshold of the number of requests for path access.",
												},
												"accurate_group_priority": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "After the rule creation is completed, the priority of the automatically generated rule group.",
												},
												"rule_priority": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Rule execution priority.",
												},
												"cc_type": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The actions performed on subsequent requests after meeting the statistical conditions.",
												},
												"field": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "statistical object.",
												},
												"count_time": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The statistical period of the strategy.",
												},
												"effect_time": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Limit the duration, that is, the effective duration of the action.",
												},
												"cron_confs": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Details of the periodic loop configuration.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"crontab": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The weekly cycle days and cycle time periods.",
															},
															"path_threshold": {
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The threshold of the number of requests for path access.",
															},
															"single_threshold": {
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The threshold of the number of visits to each statistical object.",
															},
														},
													},
												},
												"accurate_group": {
													Type:        schema.TypeList,
													MaxItems:    1,
													Computed:    true,
													Description: "Advanced conditions.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"id": {
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The ID of Rule group.",
															},
															"accurate_group_priority": {
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "After the rule creation is completed, the priority of the automatically generated rule group.",
															},
															"logic": {
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "In the rule group, the high-level conditional operation relationships corresponding to each rule.",
															},
															"accurate_rules": {
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Request characteristic information of the rule group.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"http_obj": {
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "Custom object.",
																		},
																		"obj_type": {
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "matching field.",
																		},
																		"opretar": {
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "The logical operator for the condition.",
																		},
																		"property": {
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Operate the properties of the http object.",
																		},
																		"value_string": {
																			Type:        schema.TypeString,
																			Computed:    true,
																			Description: "The value to be matched.",
																		},
																	},
																},
															},
														},
													},
												},
												"rule_tag": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Rule label, that is, the complete rule ID.",
												},
												"update_time": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Rule update time.",
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

func dataSourceVolcengineWafCcRulesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewWafCcRuleService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineWafCcRules())
}
