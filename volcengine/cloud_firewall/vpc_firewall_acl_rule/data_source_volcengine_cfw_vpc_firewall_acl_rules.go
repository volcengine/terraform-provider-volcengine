package vpc_firewall_acl_rule

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineVpcFirewallAclRules() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineVpcFirewallAclRulesRead,
		Schema: map[string]*schema.Schema{
			"vpc_firewall_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The vpc firewall id of the vpc firewall acl rule. Valid values: `in`, `out`.",
			},
			"rule_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The rule id of the vpc firewall acl rule. This field support fuzzy query.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the vpc firewall acl rule. This field support fuzzy query.",
			},
			"destination": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The destination of the vpc firewall acl rule. This field support fuzzy query.",
			},
			"source": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The source of the vpc firewall acl rule. This field support fuzzy query.",
			},
			"action": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The action list of the vpc firewall acl rule. Valid values: `accept`, `deny`, `monitor`.",
			},
			"repeat_type": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The repeat type of the vpc firewall acl rule. Valid values: `Permanent`, `Once`, `Daily`, `Weekly`, `Monthly`.",
			},
			"proto": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The proto list of the vpc firewall acl rule. Valid values: `TCP`, `ICMP`, `UDP`, `ANY`. When the destination_type is `domain`, The proto must be `TCP`.",
			},
			"status": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeBool,
				},
				Description: "The enable status list of the vpc firewall acl rule.",
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
			"vpc_firewall_acl_rules": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the vpc firewall acl rule.",
						},
						"rule_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the vpc firewall acl rule.",
						},
						"action": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The action of the vpc firewall acl rule.",
						},
						"vpc_firewall_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the vpc firewall.",
						},
						"vpc_firewall_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the vpc firewall.",
						},
						"destination_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The destination type of the vpc firewall acl rule.",
						},
						"destination": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The destination of the vpc firewall acl rule.",
						},
						"destination_group_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The destination group type of the vpc firewall acl rule.",
						},
						"destination_cidr_list": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The destination cidr list of the vpc firewall acl rule.",
						},
						"proto": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The proto of the vpc firewall acl rule.",
						},
						"source_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The source type of the vpc firewall acl rule.",
						},
						"source": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The source of the vpc firewall acl rule.",
						},
						"source_group_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The source group type of the vpc firewall acl rule.",
						},
						"source_cidr_list": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The source cidr list of the vpc firewall acl rule.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the vpc firewall acl rule.",
						},
						"dest_port_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The dest port type of the vpc firewall acl rule.",
						},
						"dest_port": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The dest port of the vpc firewall acl rule.",
						},
						"dest_port_group_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The dest port group type of the vpc firewall acl rule.",
						},
						"dest_port_list": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The dest port list of the vpc firewall acl rule.",
						},
						"repeat_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The repeat type of the vpc firewall acl rule.",
						},
						"repeat_start_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The repeat start time of the vpc firewall acl rule.",
						},
						"repeat_end_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The repeat end time of the vpc firewall acl rule.",
						},
						"repeat_days": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
							Description: "The repeat days of the vpc firewall acl rule.",
						},
						"start_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The start time of the vpc firewall acl rule. Unix timestamp.",
						},
						"end_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The end time of the vpc firewall acl rule. Unix timestamp.",
						},
						"prio": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The priority of the vpc firewall acl rule.",
						},
						"status": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to enable the vpc firewall acl rule.",
						},
						"hit_cnt": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The hit count of the vpc firewall acl rule.",
						},
						"use_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The use count of the vpc firewall acl rule.",
						},
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The account id of the vpc firewall acl rule.",
						},
						"is_effected": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the vpc firewall acl rule is effected.",
						},
						"update_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The update time of the vpc firewall acl rule.",
						},
						"effect_status": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The effect status of the vpc firewall acl rule. 1: Not yet effective, 2: Issued in progress, 3: Effective.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineVpcFirewallAclRulesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewVpcFirewallAclRuleService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineVpcFirewallAclRules())
}
