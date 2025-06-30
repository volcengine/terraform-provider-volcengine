package vmp_rule

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineVmpRules() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineVmpRulesRead,
		Schema: map[string]*schema.Schema{
			"workspace_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of workspace.",
			},
			"kind": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The kind of rule.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of rule.",
			},
			"status": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The status of rule.",
			},
			"rule_file_names": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The name of rule file.",
			},
			"rule_group_names": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The name of rule group.",
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
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of rule.",
						},
						"rule_file_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of rule file.",
						},
						"rule_group_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of rule group.",
						},
						"kind": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The kind of rule.",
						},
						"expr": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The expr of rule.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of rule.",
						},
						"reason": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The reason of rule.",
						},
						"last_evaluation": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The last evaluation of rule.",
						},
						"labels": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The labels of rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The key of label.",
									},
									"value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The value of label.",
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

func dataSourceVolcengineVmpRulesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineVmpRules())
}
