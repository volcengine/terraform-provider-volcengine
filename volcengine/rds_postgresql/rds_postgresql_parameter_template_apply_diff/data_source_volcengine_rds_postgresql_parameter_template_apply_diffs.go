package rds_postgresql_parameter_template_apply_diff

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRdsPostgresqlParameterTemplateApplyDiffs() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRdsPostgresqlParameterTemplateApplyDiffsRead,
		Schema: map[string]*schema.Schema{
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
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of the PostgreSQL instance.",
			},
			"template_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of the template.",
			},
			"parameters": {
				Description: "Changes in instance parameters after applying the specified parameter template.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the parameter.",
						},
						"new_value": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The running value defined for this parameter in the parameter template.",
						},
						"old_value": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The current running value of this parameter in the instance.",
						},
						"restart": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates whether a restart is required after the parameter is modified.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineRdsPostgresqlParameterTemplateApplyDiffsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewRdsPostgresqlParameterTemplateApplyDiffService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineRdsPostgresqlParameterTemplateApplyDiffs())
}
