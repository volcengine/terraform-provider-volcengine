package instance_parameter_log

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineMongoDBInstanceParameterLogParameters() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineMongoDBInstanceParameterLogParametersRead,
		Schema: map[string]*schema.Schema{
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of mongodb instance parameter log query.",
			},
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The instance ID to query.",
			},
			"start_time": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The start time to query.",
			},
			"end_time": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The end time to query.",
			},
			"parameter_change_logs": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Computed:    true,
				Description: "The collection of parameter change log query.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"modify_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The modifying time of parameter.",
						},
						"new_parameter_value": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The new parameter value.",
						},
						"old_parameter_value": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The old parameter value.",
						},
						"parameter_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The parameter name.",
						},
						"parameter_role": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The node type to which the parameter belongs.",
						},
						"parameter_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of parameter change.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineMongoDBInstanceParameterLogParametersRead(d *schema.ResourceData, meta interface{}) error {
	service := NewMongoDBInstanceParameterLogParameterService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineMongoDBInstanceParameterLogParameters())
}
