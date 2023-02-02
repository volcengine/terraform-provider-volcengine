package instance_parameter

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineMongoDBInstanceParameters() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineMongoDBInstanceParametersRead,
		Schema: map[string]*schema.Schema{
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of mongodb instance parameter query.",
			},
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The instance ID to query.",
			},
			"parameter_role": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "The node type of instance parameter,valid value contains `Node`,`Shard`,`ConfigServer`,`Mongos`.",
				ValidateFunc: validation.StringInSlice([]string{"Node", "Shard", "ConfigServer", "Mongos"}, false),
			},
			"parameter_names": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The parameter names,support fuzzy query, case insensitive.",
			},
			"parameters": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Computed:    true,
				Description: "The collection of parameter query.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"db_engine": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The database engine.",
						},
						"db_engine_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The database engine version.",
						},
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The instance ID.",
						},
						"total": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The total parameters queried.",
						},
						"instance_parameters": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The list of parameters.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"checking_code": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The checking code of parameter.",
									},
									"force_modify": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether the parameter supports modifying.",
									},
									"force_restart": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Does the new parameter value need to restart the instance to take effect after modification.",
									},
									"parameter_default_value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The default value of parameter.",
									},
									"parameter_description": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The description of parameter.",
									},
									"parameter_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of parameter.",
									},
									"parameter_role": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The node type to which the parameter belongs.",
									},
									"parameter_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type of parameter value.",
									},
									"parameter_value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The value of parameter.",
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

func dataSourceVolcengineMongoDBInstanceParametersRead(d *schema.ResourceData, meta interface{}) error {
	service := NewMongoDBInstanceParameterService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineMongoDBInstanceParameters())
}
