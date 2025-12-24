package rds_postgresql_parameter_template

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRdsPostgresqlParameterTemplates() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRdsPostgresqlParameterTemplatesRead,
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
			"template_category": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"DBEngine"}, false),
				Description:  "Classification of parameter templates. The current value can only be DBEngine.",
			},
			"template_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"PostgreSQL"}, false),
				Description:  "The type of the parameter template. The current value can only be PostgreSQL.",
			},
			"template_type_version": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "PostgreSQL compatible versions. The current value can only be PostgreSQL_11/12/13/14/15/16/17.",
			},
			"template_source": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice([]string{"User"}, false),
				Description:  "The source of the parameter template. The current value can only be User.",
			},
			"template_infos": {
				Description: "Parameter template list.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Account ID.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Creation time of the parameter template. The format is yyyy-MM-ddTHH:mm:ss.sssZ (UTC time).",
						},
						"need_restart": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates whether the parameter template change requires a restart.",
						},
						"parameter_num": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of parameters in the parameter template.",
						},
						"template_category": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Classification of parameter templates. The current value can only be DBEngine.",
						},
						"template_desc": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Description information of the parameter template.",
						},
						"template_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Parameter template ID.",
						},
						"template_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Parameter template name.",
						},
						"template_source": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The source of the parameter template. The current value can only be User.",
						},
						"template_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the parameter template. The current value can only be PostgreSQL.",
						},
						"template_type_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "PostgreSQL compatible versions. The current value can only be PostgreSQL_11/12/13/14/15/16/17.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Update time of the parameter template. The format is yyyy-MM-ddTHH:mm:ss.sssZ (UTC time).",
						},
						"template_params": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Parameter configuration of the parameter template.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"checking_code": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The value range of the parameter.",
									},
									"default_value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Parameter default value. Refers to the default value provided in the default template corresponding to this instance.",
									},
									"description": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The description of the parameter in English.",
									},
									"description_zh": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The description of the parameter in Chinese.",
									},
									"force_restart": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Indicates whether a restart is required after the parameter is modified.",
									},
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the parameter.",
									},
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type of the parameter.",
									},
									"value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The current value of the parameter.",
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

func dataSourceVolcengineRdsPostgresqlParameterTemplatesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewRdsPostgresqlParameterTemplateService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineRdsPostgresqlParameterTemplates())
}
