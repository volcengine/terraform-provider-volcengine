package nat_firewall_control_policy

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineNatFirewallControlPolicys() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineNatFirewallControlPolicysRead,
		Schema: map[string]*schema.Schema{
			"nat_firewall_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The nat firewall id of the nat firewall control policy.",
			},
			"direction": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The direction of nat firewall control policy. Valid values: `in`, `out`.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the nat firewall control policy. This field support fuzzy query.",
			},
			"rule_id": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The rule id of the nat firewall control policy. This field support fuzzy query.",
			},
			"dest_port": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The dest port of the nat firewall control policy.",
			},
			"destination": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The destination of the nat firewall control policy. This field support fuzzy query.",
			},
			"source": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The source of the nat firewall control policy. This field support fuzzy query.",
			},
			"action": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The action list of the nat firewall control policy. Valid values: `accept`, `deny`, `monitor`.",
			},
			"repeat_type": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The repeat type of the nat firewall control policy. Valid values: `Permanent`, `Once`, `Daily`, `Weekly`, `Monthly`.",
			},
			"proto": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The proto list of the nat firewall control policy. Valid values: `TCP`, `ICMP`, `UDP`, `ANY`. When the destination_type is `domain`, The proto must be `TCP`.",
			},
			"status": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeBool,
				},
				Description: "The enable status list of the nat firewall control policy.",
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
			"nat_firewall_control_policies": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the nat firewall control policy.",
						},
						"rule_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the nat firewall control policy.",
						},
						"action": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The action of the nat firewall control policy.",
						},
						"nat_firewall_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the nat firewall.",
						},
						"nat_firewall_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the nat firewall.",
						},
						"direction": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The direction of the nat firewall control policy.",
						},
						"destination_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The destination type of the nat firewall control policy.",
						},
						"destination": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The destination of the nat firewall control policy.",
						},
						"destination_group_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The destination group type of the nat firewall control policy.",
						},
						"destination_cidr_list": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The destination cidr list of the nat firewall control policy.",
						},
						"destination_group_list": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The destination group list of the nat firewall control policy.",
						},
						"proto": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The proto of the nat firewall control policy.",
						},
						"source_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The source type of the nat firewall control policy.",
						},
						"source": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The source of the nat firewall control policy.",
						},
						"source_group_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The source group type of the nat firewall control policy.",
						},
						"source_cidr_list": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The source cidr list of the nat firewall control policy.",
						},
						"source_group_list": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The source group list of the nat firewall control policy.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the nat firewall control policy.",
						},
						"dest_port_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The dest port type of the nat firewall control policy.",
						},
						"dest_port": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The dest port of the nat firewall control policy.",
						},
						"dest_port_group_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The dest port group type of the nat firewall control policy.",
						},
						"dest_port_list": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The dest port list of the nat firewall control policy.",
						},
						"dest_port_group_list": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The dest port group list of the nat firewall control policy.",
						},
						"repeat_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The repeat type of the nat firewall control policy.",
						},
						"repeat_start_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The repeat start time of the nat firewall control policy.",
						},
						"repeat_end_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The repeat end time of the nat firewall control policy.",
						},
						"repeat_days": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
							Description: "The repeat days of the nat firewall control policy.",
						},
						"start_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The start time of the nat firewall control policy. Unix timestamp.",
						},
						"end_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The end time of the nat firewall control policy. Unix timestamp.",
						},
						"prio": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The priority of the nat firewall control policy.",
						},
						"status": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to enable the nat firewall control policy.",
						},
						"hit_cnt": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The hit count of the nat firewall control policy.",
						},
						"use_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The use count of the nat firewall control policy.",
						},
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The account id of the nat firewall control policy.",
						},
						"is_effected": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the nat firewall control policy is effected.",
						},
						"update_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The update time of the nat firewall control policy.",
						},
						"effect_status": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The effect status of the nat firewall control policy. 1: Not yet effective, 2: Issued in progress, 3: Effective.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineNatFirewallControlPolicysRead(d *schema.ResourceData, meta interface{}) error {
	service := NewNatFirewallControlPolicyService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineNatFirewallControlPolicys())
}
