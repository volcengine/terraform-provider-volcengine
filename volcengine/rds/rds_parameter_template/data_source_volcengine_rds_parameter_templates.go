package rds_parameter_template

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	volc "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRdsParameterTemplates() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRdsParameterTemplatesRead,
		Schema: map[string]*schema.Schema{
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of RDS parameter template.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of RDS parameter templates query.",
			},
			"template_category": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Parameter template type, range of values:\nDBEngine - Engine parameters.",
			},
			"template_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Parameter template database type, range of values:\nMySQL - MySQL database.",
			},
			"template_type_version": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Parameter template database version, value range:\nMySQL_Community_5_7 - MySQL 5.7\nMySQL_8_0 - MySQL 8.0.",
			},
			"template_source": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Template source, value range:\nSystem - System\nUser - the user.",
			},
			"rds_parameter_templates": {
				Description: "The collection of RDS parameter templates query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the RDS parameter template.",
						},
						"template_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the RDS parameter template.",
						},
						"template_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the RDS parameter template.",
						},
						"template_desc": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the RDS parameter template.",
						},
						"template_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Parameter template database type, range of values:\nMySQL - MySQL database.",
						},
						"template_type_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Parameter template database version, value range:\nMySQL_Community_5_7 - MySQL 5.7\nMySQL_8_0 - MySQL 8.0.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Creation time.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Update time.",
						},
						"template_params": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Parameters contained in the template.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Parameter name.",
									},
									"default_value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Parameter default value.",
									},
									"running_value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Parameter running value.",
									},
									"restart": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether the modified parameters need to be restarted to take effect.",
									},
									"value_range": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Parameter value range.",
									},
									"description": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Parameter description.",
									},
								},
							},
						},
						"parameter_num": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of parameters the template contains.",
						},
						"need_restart": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the template contains parameters that need to be restarted.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineRdsParameterTemplatesRead(d *schema.ResourceData, meta interface{}) error {
	rdsParameterTemplateService := NewRdsParameterTemplateService(meta.(*volc.SdkClient))
	return volc.DefaultDispatcher().Data(rdsParameterTemplateService, d, DataSourceVolcengineRdsParameterTemplates())
}
