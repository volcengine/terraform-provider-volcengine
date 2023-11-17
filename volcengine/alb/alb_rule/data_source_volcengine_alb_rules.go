package alb_rule

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineAlbRules() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineAlbRulesRead,
		Schema: map[string]*schema.Schema{
			"listener_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Id of listener.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of Rule query.",
			},
			"rules": {
				Description: "The collection of Rule query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Id of Rule.",
						},
						"rule_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Id of Rule.",
						},
						"domain": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Domain of Rule.",
						},
						"url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Url of Rule.",
						},
						"rule_action": {
							Type:     schema.TypeString,
							Computed: true,
							Description: "The forwarding rule action, if this parameter is empty, " +
								"forward to server group, if value is `Redirect`, will redirect.",
						},
						"server_group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Id of Server Group.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Description of Rule.",
						},
						"traffic_limit_enabled": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Forwarding rule QPS rate limiting switch:\n on: enable.\noff: disable (default).",
						},
						"traffic_limit_qps": {
							Type:     schema.TypeInt,
							Computed: true,
							Description: "When Rules.N.TrafficLimitEnabled is turned on, this field is required. " +
								"Requests per second. Valid values are between 100 and 100000.",
						},
						"rewrite_enabled": {
							Type:     schema.TypeString,
							Computed: true,
							Description: "Rewrite configuration switch for forwarding rules, only allows configuration and takes effect when RuleAction is empty (i.e., forwarding to server group). " +
								"Only available for whitelist users, please submit an application to experience. " +
								"Supported values are as follows:\non: enable.\noff: disable.",
						},
						"rewrite_config": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The list of rewrite configurations.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"rewrite_path": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Rewrite path.",
									},
								},
							},
						},
						"redirect_config": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Redirect related configuration.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"redirect_domain": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The redirect domain.",
									},
									"redirect_uri": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The redirect URI.",
									},
									"redirect_port": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The redirect port.",
									},
									"redirect_http_code": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The redirect HTTP code,support 301(default), 302, 307, 308.",
									},
									"redirect_protocol": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The redirect protocol,support HTTP,HTTPS(default).",
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

func dataSourceVolcengineAlbRulesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewAlbRuleService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineAlbRules())
}
