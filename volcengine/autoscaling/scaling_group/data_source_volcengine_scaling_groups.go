package scaling_group

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineScalingGroups() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineScalingGroupsRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of scaling group ids.",
			},
			"scaling_group_names": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of scaling group names.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of scaling group.",
			},

			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of scaling group query.",
			},
			"scaling_groups": {
				Description: "The collection of scaling group query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the scaling group.",
						},
						"scaling_group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the scaling group.",
						},
						"scaling_group_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the scaling group.",
						},
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The VPC id of the scaling group.",
						},
						"subnet_ids": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The list of the subnet id to which the ENI is connected.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"default_cooldown": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The default cooldown interval of the scaling group.",
						},
						"lifecycle_state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The lifecycle state of the scaling group.",
						},
						"active_scaling_configuration_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The scaling configuration id which used by the scaling group.",
						},
						"desire_instance_number": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The desire instance number of the scaling group.",
						},
						"min_instance_number": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The min instance number of the scaling group.",
						},
						"max_instance_number": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The max instance number of the scaling group.",
						},
						"total_instance_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The total instance count of the scaling group.",
						},
						"instance_terminate_policy": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The instance terminate policy of the scaling group.",
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the scaling group.",
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the scaling group.",
						},
						"multi_az_policy": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The multi az policy of the scaling group. Valid values: PRIORITY, BALANCE.",
						},
						"db_instance_ids": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The list of db instance ids.",
						},
						"server_group_attributes": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"load_balancer_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The load balancer id.",
									},
									"port": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The port receiving request of the server group.",
									},
									"server_group_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The server group id.",
									},
									"weight": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The weight of the instance.",
									},
								},
							},
							Description: "The list of server group attributes.",
						},
						"launch_template_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the launch template bound to the scaling group.",
						},
						"launch_template_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The version of the launch template bound to the scaling group.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineScalingGroupsRead(d *schema.ResourceData, meta interface{}) error {
	scalingGroupService := NewScalingGroupService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(scalingGroupService, d, DataSourceVolcengineScalingGroups())
}