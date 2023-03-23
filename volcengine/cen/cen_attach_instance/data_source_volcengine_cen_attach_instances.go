package cen_attach_instance

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineCenAttachInstances() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineCenInstanceAttachInstances,
		Schema: map[string]*schema.Schema{
			"cen_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A cen ID.",
			},
			"instance_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "An instance ID.",
			},
			"instance_type": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "An instance type.",
				ValidateFunc: validation.StringInSlice([]string{"VPC", "DCGW"}, false),
			},
			"instance_region_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A region id of instance.",
			},

			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of cen attach instance query.",
			},
			"attach_instances": {
				Description: "The collection of cen attach instance query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cen_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the cen.",
						},
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the instance.",
						},
						"instance_owner_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The owner ID of the instance.",
						},
						"instance_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the instance.",
						},
						"instance_region_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The region id of the instance.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the cen attaching instance.",
						},
						"creation_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the cen attaching instance.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineCenInstanceAttachInstances(d *schema.ResourceData, meta interface{}) error {
	cenAttachInstanceService := NewCenAttachInstanceService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(cenAttachInstanceService, d, DataSourceVolcengineCenAttachInstances())
}