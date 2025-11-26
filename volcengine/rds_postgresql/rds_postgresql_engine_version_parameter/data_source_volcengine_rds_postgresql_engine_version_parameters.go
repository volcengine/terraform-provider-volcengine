package rds_postgresql_engine_version_parameter

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRdsPostgresqlEngineVersionParameters() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRdsPostgresqlEngineVersionParametersRead,
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
			"db_engine": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"PostgreSQL"}, false),
				Description:  "The type of the parameter template. The value can only be PostgreSQL.",
			},
			"db_engine_version": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"PostgreSQL_11", "PostgreSQL_12", "PostgreSQL_13", "PostgreSQL_14", "PostgreSQL_15", "PostgreSQL_16", "PostgreSQL_17"}, false),
				Description: "The database engine version of the RDS PostgreSQL instance. " +
					"Valid value: PostgreSQL_11, PostgreSQL_12, PostgreSQL_13, PostgreSQL_14, PostgreSQL_15, PostgreSQL_16, PostgreSQL_17.",
			},
			"db_engine_version_parameters": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"db_engine_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The database engine version of the RDS PostgreSQL instance.",
						},
						"parameter_count": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The number of parameters that users can set under the specified database engine version.",
						},
						"parameters": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The collection of parameters that users can set under the specified database engine version.",
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
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineRdsPostgresqlEngineVersionParametersRead(d *schema.ResourceData, meta interface{}) error {
	service := NewRdsPostgresqlEngineVersionParameterService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineRdsPostgresqlEngineVersionParameters())
}
