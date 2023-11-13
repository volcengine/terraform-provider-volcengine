package instance_parameter

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The node type of instance parameter, valid value contains `Node`, `Shard`, `ConfigServer`, `Mongos`.",
			},
			"parameter_names": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The parameter names, support fuzzy query, case insensitive.",
			},
			"instance_parameters": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The collection of parameter query.",
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
			"parameters": {
				Type:        schema.TypeList,
				Computed:    true,
				Deprecated:  "This field has been deprecated and it is recommended to use instance_parameters.",
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
	err := ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineMongoDBInstanceParameters())
	if err != nil {
		return err
	}
	condition := map[string]interface{}{
		"InstanceId": d.Get("instance_id"),
	}
	if name, ok := d.GetOk("parameter_names"); ok {
		condition["ParameterNames"] = name
	}
	if role, ok := d.GetOk("parameter_role"); ok {
		condition["ParameterRole"] = role
	}
	results, err := service.ReadAll(condition)
	if err != nil {
		return err
	}
	resource := DataSourceVolcengineMongoDBInstanceParameters()
	dataSourceInfo := service.DatasourceResources(d, resource)
	dataSourceInfo.CollectField = "parameters"
	dataSourceInfo.ResponseConverts = map[string]ve.ResponseConvert{
		"DBEngine": {
			TargetField: "db_engine",
		},
		"DBEngineVersion": {
			TargetField: "db_engine_version",
		},
		"ParameterNames": {
			TargetField: "parameter_name",
		},
	}
	return ve.ResponseToDataSource(d, resource, dataSourceInfo, results)
}
