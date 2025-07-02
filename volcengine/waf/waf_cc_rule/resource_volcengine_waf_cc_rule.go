package waf_cc_rule

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
WafCcRule can be imported using the id, e.g.
```
$ terraform import volcengine_waf_cc_rule.default resource_id:Host
```

*/

func ResourceVolcengineWafCcRule() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineWafCcRuleCreate,
		Read:   resourceVolcengineWafCcRuleRead,
		Update: resourceVolcengineWafCcRuleUpdate,
		Delete: resourceVolcengineWafCcRuleDelete,
		Importer: &schema.ResourceImporter{
			State: wafCcRuleImporter,
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
				Description: "The name of cc rule.",
			},
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The website request path that needs protection.",
			},
			"advanced_enable": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether to enable advanced conditions.",
			},
			"field": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "statistical object.",
			},
			"single_threshold": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The threshold of the number of times each statistical object accesses the request path.",
			},
			"path_threshold": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The threshold of the total number of times the request path is accessed.",
			},
			"count_time": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The statistical period of the strategy.",
			},
			"cc_type": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The actions performed on subsequent requests after meeting the statistical conditions.",
			},
			"effect_time": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Limit the duration, that is, the effective duration of the action.",
			},
			"rule_priority": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Rule execution priority.",
			},
			"enable": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Whether to enable the rules.",
			},
			"host": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Website domain names that require the setting of protection rules.",
			},
			"accurate_group": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Advanced conditions.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"accurate_rules": {
							Type:        schema.TypeSet,
							Required:    true,
							Description: "Details of advanced conditions.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"http_obj": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The HTTP object to be added to the advanced conditions.",
									},
									"obj_type": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "The matching field for HTTP objects.",
									},
									"opretar": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "The logical operator for the condition.",
									},
									"property": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "Operate the properties of the http object.",
									},
									"value_string": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "The value to be matched.",
									},
								},
							},
						},
						"logic": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "The logical relationship of advanced conditions.",
						},
					},
				},
			},
			"exemption_time": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "Strategy exemption time.",
			},
			"cron_enable": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "Whether to set the cycle to take effect.",
			},
			"cron_confs": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Details of the periodic loop configuration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"crontab": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The weekly cycle days and cycle time periods.",
						},
						"path_threshold": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "The threshold of the number of requests for path access.",
						},
						"single_threshold": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "The threshold of the number of visits to each statistical object.",
						},
					},
				},
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Name of the affiliated project resource.",
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
	}
	return resource
}

func resourceVolcengineWafCcRuleCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewWafCcRuleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineWafCcRule())
	if err != nil {
		return fmt.Errorf("error on creating waf_cc_rule %q, %s", d.Id(), err)
	}
	return resourceVolcengineWafCcRuleRead(d, meta)
}

func resourceVolcengineWafCcRuleRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewWafCcRuleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineWafCcRule())
	if err != nil {
		return fmt.Errorf("error on reading waf_cc_rule %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineWafCcRuleUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewWafCcRuleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineWafCcRule())
	if err != nil {
		return fmt.Errorf("error on updating waf_cc_rule %q, %s", d.Id(), err)
	}
	return resourceVolcengineWafCcRuleRead(d, meta)
}

func resourceVolcengineWafCcRuleDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewWafCcRuleService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineWafCcRule())
	if err != nil {
		return fmt.Errorf("error on deleting waf_cc_rule %q, %s", d.Id(), err)
	}
	return err
}
