package instance_types

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineInstanceTypes() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineInstanceTypesRead,
		Schema: map[string]*schema.Schema{
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of instance types query.",
			},
			"instance_type_configs": {
				Description: "The collection of instance types query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of instance.",
						},
						"instance_type_family": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type family of instance.",
						},
						"instance_type_family_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of instance type family.",
						},
						"gpu_spec": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The gpu spec of instance.",
						},
						"cpu": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The cpu of instance type.",
						},
						"gpu": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The gpu of instance type.",
						},
						"memory": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The memory of instance type.",
						},
						"storage": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The config of storage.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"local_storage_category": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The local storage category.",
									},
									"local_storage_capacity": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The capacity of local storage.",
									},
									"local_storage_amount": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The amount of local storage.",
									},
									"local_storage_unit": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unit of local storage.",
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

func dataSourceVolcengineInstanceTypesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewInstanceTypeService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineInstanceTypes())
}
