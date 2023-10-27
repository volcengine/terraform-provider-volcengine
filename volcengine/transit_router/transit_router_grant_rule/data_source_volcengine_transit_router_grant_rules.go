package transit_router_grant_rule

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineTransitRouterGrantRules() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineTransitRouterGrantRulesRead,
		Schema: map[string]*schema.Schema{
			"transit_router_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of the transit router.",
			},
			"grant_account_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the grant account.",
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
						"grant_account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the grant account.",
						},
						"transit_router_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the transaction router.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the rule.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the rule.",
						},
						"creation_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The creation time of the rule.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of the rule.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineTransitRouterGrantRulesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewTransitRouterGrantRuleService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineTransitRouterGrantRules())
}
