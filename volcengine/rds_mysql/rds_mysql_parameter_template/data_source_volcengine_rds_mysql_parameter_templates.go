package rds_mysql_parameter_template

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRdsMysqlParameterTemplates() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRdsMysqlParameterTemplatesRead,
		Schema: map[string]*schema.Schema{
			"template_category": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Template category, with a value of DBEngine (database engine parameters).",
			},
			"template_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Database type of parameter template. The default value is Mysql.",
			},
			"template_type_version": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "Database version of parameter template. " +
					"Value range:\nMySQL_5_7: Default value. MySQL 5.7 version.\n" +
					"MySQL_8_0: MySQL 8.0 version.",
			},
			"template_source": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Parameter template source, value range: System. User.",
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
			"templates": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The account ID.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Creation time.",
						},
						"need_restart": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Does the template contain parameters that require restart.",
						},
						"parameter_num": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of parameters contained in the template.",
						},
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The project to which the template belongs.",
						},
						"template_category": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Template category, with a value of DBEngine (database engine parameter).",
						},
						"template_desc": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Parameter template description.",
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
						"template_params": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Parameters contained in the template.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"default_value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Parameter default value.",
									},
									"description": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Parameter description.",
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
										Description: "Instance parameter name.\n" +
											"Description: When using CreateParameterTemplate and ModifyParameterTemplate as request parameters, " +
											"only Name and RunningValue need to be passed in.",
									},
									"restart": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Is it necessary to restart the instance for the changes to take effect.",
									},
									"running_value": {
										Type:     schema.TypeString,
										Computed: true,
										Description: "Parameter running value.\n" +
											"Description: When making requests with CreateParameterTemplate and ModifyParameterTemplate as request parameters," +
											" only Name and RunningValue need to be passed in.",
									},
									"value_range": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Value range of parameters.",
									},
								},
							},
						},
						"template_source": {
							Type:     schema.TypeString,
							Computed: true,
							Description: "The type of parameter template. " +
								"Values:\nSystem: System template." +
								"\nUser: User template.",
						},
						"template_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The database type of the parameter template. The default value is Mysql.",
						},
						"template_type_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Parameter template database version, value range:\n\"MySQL_5_7\": MySQL 5.7 version.\n\"MySQL_8_0\": MySQL 8.0 version.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Modification time of the template.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineRdsMysqlParameterTemplatesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewRdsMysqlParameterTemplateService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineRdsMysqlParameterTemplates())
}
