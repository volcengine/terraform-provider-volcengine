package waf_bot_analyse_protect_rule

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineWafBotAnalyseProtectRules() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineWafBotAnalyseProtectRulesRead,
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
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the project to which your domain names belong.",
			},
			"host": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Website domain names that require the setting of protection rules.",
			},
			"bot_space": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Bot protection rule type.",
			},
			"rule_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Unique identification of rules.",
			},
			"path": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Protective path.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the rule.",
			},
			"data": {
				Description: "The details of the Bot rules.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"path": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The requested path.",
						},
						"total_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The total number of statistical protection rules under the current domain name.",
						},
						"enable_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of statistical protection rules enabled under the current domain name.\n",
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
												"statistical_type": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Statistical content method.",
												},
												"statistical_duration": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The duration of the statistics.",
												},
												"single_threshold": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The maximum number of ips for the same statistical object.",
												},
												"single_proportion": {
													Type:        schema.TypeFloat,
													Computed:    true,
													Description: "The IP proportion of the same statistical object.",
												},
												"rule_priority": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Rule execution priority.",
												},
												"path": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Request path.",
												},
												"name": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The name of rule.",
												},
												"host": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The domain name where the protection rule is located.",
												},
												"field": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "statistical object.",
												},
												"exemption_time": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Exemption time.",
												},
												"enable": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Whether to enable the rules.",
												},
												"effect_time": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Limit the duration.",
												},
												"action_type": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "perform the action.",
												},
												"action_after_verification": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Perform actions after human-machine verification /JS challenges.",
												},
												"path_threshold": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Threshold of path access times.",
												},
												"id": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Rule unique identifier.",
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
												"pass_ratio": {
													Type:        schema.TypeFloat,
													Computed:    true,
													Description: "JS challenge/human-machine verification pass rate.",
												},
												"accurate_group_priority": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Advanced condition priority.",
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

func dataSourceVolcengineWafBotAnalyseProtectRulesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewWafBotAnalyseProtectRuleService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineWafBotAnalyseProtectRules())
}
