package rds_postgresql_instance_parameter

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRdsPostgresqlInstanceParameters() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRdsPostgresqlInstanceParametersRead,
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
				Description: "The ID of the PostgreSQL instance.",
			},
			"parameter_name": {
				Type:     schema.TypeString,
				Optional: true,
				Description: "The name of the parameter, supports fuzzy query. " +
					"If no value is passed or a null value is passed, all parameters under the specified instance will be queried.",
			},
			"instance_parameters": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the PostgreSQL instance.",
						},
						"db_engine_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The version of the PostgreSQL engine.",
						},
						"parameter_count": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The total count of parameters.",
						},
						"none_kernel_parameters": {
							Type:       schema.TypeList,
							Computed:   true,
							Deprecated: "The current parameter configuration of the instance (non-kernel parameters).",
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
						"parameters": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The current parameter configuration of the instance (kernel parameters).",
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

func dataSourceVolcengineRdsPostgresqlInstanceParametersRead(d *schema.ResourceData, meta interface{}) error {
	service := NewRdsPostgresqlInstanceParameterService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineRdsPostgresqlInstanceParameters())
}
