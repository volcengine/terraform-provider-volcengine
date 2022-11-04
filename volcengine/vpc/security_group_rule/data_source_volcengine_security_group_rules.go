package security_group_rule

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineSecurityGroupRules() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineSecurityGroupRulesRead,
		Schema: map[string]*schema.Schema{
			"security_group_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "SecurityGroup ID.",
			},
			"direction": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"ingress",
					"egress",
				}, false),
				Description: "Direction of rule, ingress (inbound) or egress (outbound).",
			},
			"cidr_ip": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Cidr ip of egress/ingress Rule.",
			},
			"protocol": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"tcp",
					"udp",
					"icmp",
					"all",
				}, false),
				Description: "Protocol of the SecurityGroup, the value can be `tcp` or `udp` or `icmp` or `all`.",
			},
			"source_group_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ID of the source security group whose access permission you want to set.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"security_group_rules": {
				Description: "The collection of SecurityGroup query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"direction": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Direction of rule, ingress (inbound) or egress (outbound).",
						},
						"protocol": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Protocol of the SecurityGroup, the value can be `tcp` or `udp` or `icmp` or `all`.",
						},
						"security_group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Id of SecurityGroup.",
						},
						"port_start": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Port start of egress/ingress Rule.",
						},
						"port_end": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Port end of egress/ingress Rule.",
						},
						"cidr_ip": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Cidr ip of egress/ingress Rule.",
						},
						"source_group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the source security group whose access permission you want to set.",
						},
						"policy": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Access strategy.",
						},
						"priority": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Priority of a security group rule.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "description of a group rule.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineSecurityGroupRulesRead(d *schema.ResourceData, meta interface{}) error {
	securityGroupService := NewSecurityGroupRuleService(meta.(*ve.SdkClient))
	return securityGroupService.Dispatcher.Data(securityGroupService, d, DataSourceVolcengineSecurityGroupRules())
}
