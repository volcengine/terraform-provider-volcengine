package waf_ip_group

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineWafIpGroups() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineWafIpGroupsRead,
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
			"rule_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Query the association rule ID.",
			},
			"time_order_by": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The arrangement order of the address group.",
			},
			"ip": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The address or address segment of the query.",
			},
			"ip_group_list": {
				Description: "Address group list information.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the ip group.",
						},
						"ip_group_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The ID of the ip group.",
						},
						"ip_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of IP addresses within the address group.",
						},
						"ip_list": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The IP address to be added.",
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
									"host": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The information of the protected domain names associated with the rules.",
									},
								},
							},
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ip group update time.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineWafIpGroupsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewWafIpGroupService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineWafIpGroups())
}
