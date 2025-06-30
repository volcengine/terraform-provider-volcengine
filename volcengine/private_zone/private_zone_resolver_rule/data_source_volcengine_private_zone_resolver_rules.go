package private_zone_resolver_rule

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcenginePrivateZoneResolverRules() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcenginePrivateZoneResolverRulesRead,
		Schema: map[string]*schema.Schema{
			//"ids": {
			//	Type:     schema.TypeSet,
			//	Optional: true,
			//	Elem: &schema.Schema{
			//		Type: schema.TypeString,
			//	},
			//	Set:         schema.HashString,
			//	Description: "A list of IDs.",
			//},
			//"rule_id": {
			//	Type:        schema.TypeInt,
			//	Optional:    true,
			//	Description: "The ID of the rule.",
			//},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the rule.",
			},
			"zone_name": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "The main domain associated with the forwarding rule. " +
					"For example, if you set this parameter to example.com, " +
					"DNS requests for example.com and all subdomains of example.com will be forwarded.",
			},
			"endpoint_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "ID of the exit terminal node.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The project name of the private zone resolver rule.",
			},
			"tag_filters": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of tag filters.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The key of the tag.",
						},
						"values": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "The values of the tag.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
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
			"rules": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the rule.",
						},
						"rule_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The id of the rule.",
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The created time of the rule.",
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The updated time of the rule.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the rule.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the rule.",
						},
						"zone_name": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The zone name of the rule.",
						},
						"endpoint_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The endpoint ID of the rule.",
						},
						"line": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ISP of the exit IP address of the recursive DNS server.",
						},
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The project name of the rule.",
						},
						"tags": ve.TagsSchemaComputed(),
						"forward_ips": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The IP address and port of the DNS server outside of the VPC.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The IP address of the DNS server outside of the VPC.",
									},
									"port": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The port of the DNS server outside of the VPC.",
									},
								},
							},
						},
						"bind_vpcs": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The id of the bind vpc.",
									},
									"region": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The region of the bind vpc.",
									},
									"region_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The region name of the bind vpc.",
									},
									"account_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The account id of the bind vpc.",
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

func dataSourceVolcenginePrivateZoneResolverRulesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewPrivateZoneResolverRuleService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcenginePrivateZoneResolverRules())
}
