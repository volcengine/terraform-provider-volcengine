package waf_host_group

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineWafHostGroups() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineWafHostGroupsRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of IDs.",
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
			"host_fix": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The domain name information queried.",
			},
			"host_group_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The ID of the domain name group.",
			},
			"list_all": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether to return all domain name groups and their name information, it returns by default.",
			},
			"name_fix": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the domain name group being queried.",
			},
			"rule_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The rule ID associated with domain name groups.",
			},
			"time_order_by": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The list of rule ids associated with the domain name group shows the timing sequence.",
			},
			"host_group_list": {
				Description: "Details of the domain name group list.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Domain name group description.",
						},
						"host_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of domain names contained in the domain name group.",
						},
						"host_group_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The ID of the domain name group.",
						},
						"host_list": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "Domain names that need to be added to this domain name group.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the domain name group.",
						},
						"related_rules": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The list of associated rules.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"rule_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the rule.",
									},
									"rule_tag": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The ID of the rule.",
									},
									"rule_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type of the rule.",
									},
								},
							},
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Domain name group update time.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineWafHostGroupsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewWafHostGroupService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineWafHostGroups())
}
