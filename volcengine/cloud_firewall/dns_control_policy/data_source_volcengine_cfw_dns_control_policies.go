package dns_control_policy

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineDnsControlPolicies() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineDnsControlPoliciesRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The rule id list of the dns control policy. This field support fuzzy query.",
			},
			"source": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The source list of the dns control policy. This field support fuzzy query.",
			},
			"destination": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The destination list of the dns control policy. This field support fuzzy query.",
			},
			"status": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeBool,
				},
				Description: "The enable status list of the dns control policy. This field support fuzzy query.",
			},
			"internet_firewall_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The internet firewall id of the dns control policy.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the dns control policy. This field support fuzzy query.",
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
			"dns_control_policies": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the dns control policy.",
						},
						"rule_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the dns control policy.",
						},
						"hit_cnt": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The hit count of the dns control policy.",
						},
						"use_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The use count of the dns control policy.",
						},
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The account id of the dns control policy.",
						},
						"status": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to enable the dns control policy.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the dns control policy.",
						},
						"last_hit_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The last hit time of the dns control policy. Unix timestamp.",
						},
						"destination_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The destination type of the dns control policy.",
						},
						"destination": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The destination of the dns control policy.",
						},
						"destination_group_list": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The destination group list of the dns control policy.",
						},
						"domain_list": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The destination domain list of the dns control policy.",
						},
						"source": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The source vpc list of the dns control policy.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"vpc_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The id of the source vpc.",
									},
									"region": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The region of the source vpc.",
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

func dataSourceVolcengineDnsControlPoliciesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewDnsControlPolicyService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineDnsControlPolicies())
}
