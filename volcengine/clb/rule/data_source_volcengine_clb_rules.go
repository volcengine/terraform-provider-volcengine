package rule

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRules() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRulesRead,
		Schema: map[string]*schema.Schema{
			"listener_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The Id of listener.",
			},
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of Rule IDs.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"tags": ve.TagsSchema(),
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
						"tags": ve.TagsSchemaComputed(),
						"action_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The action type of Rule. values: `Forward`, `Redirect`.",
						},
						"redirect_config": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"protocol": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The redirect protocol.",
									},
									"host": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The redirect host.",
									},
									"path": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The redirect path.",
									},
									"port": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The redirect port.",
									},
									"status_code": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The redirect status code.",
									},
								},
							},
							Description: "The redirect configuration. When `action_type` is `Redirect`, this parameter is returned.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineRulesRead(d *schema.ResourceData, meta interface{}) error {
	ruleService := NewRuleService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(ruleService, d, DataSourceVolcengineRules())
}
