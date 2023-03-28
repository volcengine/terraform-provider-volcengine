package scaling_instance

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineScalingInstances() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineScalingInstancesRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of instance ids.",
			},
			"scaling_group_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of the scaling group.",
			},
			"scaling_configuration_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the scaling configuration id.",
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"Init",
					"Pending",
					"Pending:Wait",
					"InService",
					"Error",
					"Removing",
					"Removing:Wait",
					"Stopped",
					"Protected",
				}, false),
				Description: "The status of instances. Valid values: Init, Pending, Pending:Wait, InService, Error, Removing, Removing:Wait, Stopped, Protected.",
			},
			"creation_type": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"AutoCreated",
					"Attached",
				}, false),
				Description: "The creation type of the instances. Valid values: AutoCreated, Attached.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of scaling instances query.",
			},
			"scaling_instances": {
				Description: "The collection of scaling instances.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the scaling instance.",
						},
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the scaling instance.",
						},
						"scaling_group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the scaling group.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of instances.",
						},
						"scaling_policy_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the scaling policy.",
						},
						"scaling_configuration_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the scaling configuration.",
						},
						"entrusted": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to host the instance to a scaling group.",
						},
						"created_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The time when the instance was added to the scaling group.",
						},
						"creation_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The creation type of the instance. Valid values: AutoCreated, Attached.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineScalingInstancesRead(d *schema.ResourceData, meta interface{}) error {
	scalingInstanceService := NewScalingInstanceService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(scalingInstanceService, d, DataSourceVolcengineScalingInstances())
}
