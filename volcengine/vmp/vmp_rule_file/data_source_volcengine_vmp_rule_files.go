package vmp_rule_file

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineVmpRuleFiles() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineVmpRuleFilesRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of Rule File IDs.",
			},
			"workspace_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of workspace.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of rule file.",
			},
			"status": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The status of rule file.",
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
			"files": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of rule file.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of rule file.",
						},
						"last_update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The last update time of rule file.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of rule file.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of rule file.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of rule file.",
						},
						"content": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The content of rule file.",
						},
						"rule_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The rule count number of rule file.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineVmpRuleFilesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineVmpRuleFiles())
}
