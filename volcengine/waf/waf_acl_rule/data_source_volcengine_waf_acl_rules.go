package waf_acl_rule

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineWafAclRules() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineWafAclRulesRead,
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
			"acl_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The types of access control rules.",
			},
			"action": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "Action to be taken on requests that match the rule.",
			},
			"defence_host": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The list of queried domain names.",
			},
			"enable": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
				Set:         schema.HashInt,
				Description: "The enabled status of the rule.",
			},
			"rule_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Rule name, fuzzy search.",
			},
			"rule_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Rule unique identifier, precise search.",
			},
			"time_order_by": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The list shows the timing sequence.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the project to which your domain names belong.",
			},
			"rules": {
				Description: "Details of the rules.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Rule ID.",
						},
						"ip_add_type": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Type of IP address addition.",
						},
						"ip_group_id": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
							Description: "Add the list of address group ids in the address group mode.",
						},
						"ip_location_subregion": {
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "Domestic region code.",
						},
						"accurate_group": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Computed:    true,
							Description: "Advanced conditions.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"accurate_rules": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Details of advanced conditions.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"http_obj": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The HTTP object to be added to the advanced conditions.",
												},
												"obj_type": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The matching field for HTTP objects.",
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
									"logic": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The logical relationship of advanced conditions.",
									},
								},
							},
						},
						"action": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Action to be taken on requests that match the rule.",
						},
						"advanced": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Whether to set advanced conditions.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Rule description.",
						},
						"enable": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Whether to enable the rule.",
						},
						"host_add_type": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Type of domain name addition.",
						},
						"host_group_id": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
							Description: "The ID of the domain group.",
						},
						"host_groups": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The list of domain name groups.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"host_group_id": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The ID of host group.",
									},
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Name of host group.",
									},
								},
							},
						},
						"host_list": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "Single or multiple domain names are supported.",
						},
						"ip_list": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "Single or multiple IP addresses are supported.",
						},
						"ip_location_country": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "Country or region code.",
						},
						"rule_tag": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Rule unique identifier.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Update time of the rule.",
						},
						"url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The path of Matching.",
						},
						"ip_groups": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The list of domain name groups.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_group_id": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The ID of the IP address group.",
									},
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Name of the IP address group.",
									},
								},
							},
						},
						"client_ip": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IP address.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Rule name.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineWafAclRulesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewWafAclRuleService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineWafAclRules())
}
