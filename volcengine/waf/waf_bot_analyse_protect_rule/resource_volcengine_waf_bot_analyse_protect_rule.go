package waf_bot_analyse_protect_rule

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
WafBotAnalyseProtectRule can be imported using the id, e.g.
```
$ terraform import volcengine_waf_bot_analyse_protect_rule.default resource_id
```

*/

func ResourceVolcengineWafBotAnalyseProtectRule() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineWafBotAnalyseProtectRuleCreate,
		Read:   resourceVolcengineWafBotAnalyseProtectRuleRead,
		Update: resourceVolcengineWafBotAnalyseProtectRuleUpdate,
		Delete: resourceVolcengineWafBotAnalyseProtectRuleDelete,
		Importer: &schema.ResourceImporter{
			State: wafBotAnalyseProtectRuleImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Name of the affiliated project resource.",
			},
			"statistical_type": {
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntInSlice([]int{1, 2, 3}),
				Required:     true,
				Description:  "Statistical content and methods.",
			},
			"statistical_duration": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The duration of statistics.",
			},
			"single_threshold": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The maximum number of ips of the same statistical object is enabled when StatisticalType=2.",
			},
			"single_proportion": {
				Type:        schema.TypeFloat,
				Optional:    true,
				Description: "The IP proportion of the same statistical object needs to be configured when StatisticalType=3.",
			},
			"rule_priority": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Priority of rule effectiveness.",
			},
			"path_threshold": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The path access frequency threshold is enabled when StatisticalType=1.",
			},
			"path": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The requested path.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of rule.",
			},
			"host": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Website domain names that require the setting of protection rules.",
			},
			"field": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Statistical objects, with multiple objects separated by commas.",
			},
			"exemption_time": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Exemption time takes effect when the execution action is human-machine challenge /JS/ Proof of work.",
			},
			"enable": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Whether to enable the rules.",
			},
			"effect_time": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Limit the duration.",
			},
			"action_type": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "perform the action.",
			},
			"action_after_verification": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Perform the action after verification/challenge.",
			},
			"accurate_group": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Advanced conditions.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"logic": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "In the rule group, the high-level conditional operation relationships corresponding to each rule.",
						},
						"accurate_rules": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "Request characteristic information of the rule group.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"http_obj": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Custom object.",
									},
									"obj_type": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "matching field.",
									},
									"opretar": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The logical operator for the condition.",
									},
									"property": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Operate the properties of the http object.",
									},
									"value_string": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The value to be matched.",
									},
								},
							},
						},
					},
				},
			},
			// computed
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
			"enable_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of statistical protection rules enabled under the current domain name.\n",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total number of statistical protection rules under the current domain name.",
			},
		},
	}
	return resource
}

func resourceVolcengineWafBotAnalyseProtectRuleCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewWafBotAnalyseProtectRuleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineWafBotAnalyseProtectRule())
	if err != nil {
		return fmt.Errorf("error on creating waf_bot_analyse_protect_rule %q, %s", d.Id(), err)
	}
	return resourceVolcengineWafBotAnalyseProtectRuleRead(d, meta)
}

func resourceVolcengineWafBotAnalyseProtectRuleRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewWafBotAnalyseProtectRuleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineWafBotAnalyseProtectRule())
	if err != nil {
		return fmt.Errorf("error on reading waf_bot_analyse_protect_rule %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineWafBotAnalyseProtectRuleUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewWafBotAnalyseProtectRuleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineWafBotAnalyseProtectRule())
	if err != nil {
		return fmt.Errorf("error on updating waf_bot_analyse_protect_rule %q, %s", d.Id(), err)
	}
	return resourceVolcengineWafBotAnalyseProtectRuleRead(d, meta)
}

func resourceVolcengineWafBotAnalyseProtectRuleDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewWafBotAnalyseProtectRuleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineWafBotAnalyseProtectRule())
	if err != nil {
		return fmt.Errorf("error on deleting waf_bot_analyse_protect_rule %q, %s", d.Id(), err)
	}
	return err
}
