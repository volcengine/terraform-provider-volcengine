package financial_relation

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineFinancialRelations() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineFinancialRelationsRead,
		Schema: map[string]*schema.Schema{
			"account_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of sub account IDs.",
			},
			"relation": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of relation. Valid values: `1`, `4`.",
			},
			"status": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of status. Valid values: `100`, `200`, `250`, `300`, `400`, `500`.",
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
			"financial_relations": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"relation_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the financial relation.",
						},
						"major_account_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The id of the major account.",
						},
						"major_account_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the major account.",
						},
						"sub_account_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The id of the sub account.",
						},
						"sub_account_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the sub account.",
						},
						"account_alias": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The display name of the sub account.",
						},
						"relation": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The relation of the financial.",
						},
						"relation_desc": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The relation description of the financial.",
						},
						"filiation": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The filiation of the financial relation.",
						},
						"filiation_desc": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The filiation description of the financial relation.",
						},
						"status": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The status of the financial relation.",
						},
						"status_desc": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status description of the financial relation.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of the financial relation.",
						},
						"auth_info": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The authorization info of the financial relation.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auth_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The auth id of the financial relation.",
									},
									"auth_status": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The auth status of the financial relation.",
									},
									"auth_list": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeInt,
										},
										Description: "The auth list of the financial relation.",
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

func dataSourceVolcengineFinancialRelationsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewFinancialRelationService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineFinancialRelations())
}
