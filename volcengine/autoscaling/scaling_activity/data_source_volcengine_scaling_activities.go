package scaling_activity

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineScalingActivities() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineScalingActivitiesRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of Scaling Activity IDs.",
			},
			"scaling_group_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A Id of Scaling Group.",
			},
			"start_time": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: timeValidation,
				Description:  "A start time to start a Scaling Activity.",
			},
			"end_time": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: timeValidation,
				Description:  "An end time to start a Scaling Activity.",
			},
			"status_code": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A status code of Scaling Activity.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of Scaling Activity query.",
			},
			"activities": {
				Description: "The collection of Scaling Activity query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of Scaling Activity.",
						},
						"scaling_activity_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of Scaling Activity.",
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of Scaling Activity.",
						},
						"expected_run_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The expected run time of Scaling Activity.",
						},
						"stopped_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The stopped time of Scaling Activity.",
						},
						"task_category": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The task category of Scaling Activity.",
						},
						"scaling_group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The scaling group Id.",
						},
						"status_code": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Status Code of Scaling Activity.",
						},
						"result_msg": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Result of Scaling Activity.",
						},
						"max_instance_number": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The Max Instance Number.",
						},
						"min_instance_number": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The Min Instance Number.",
						},
						"current_instance_number": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The Current Instance Number.",
						},
						"actual_adjust_instance_number": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The Actual Adjustment Instance Number.",
						},
						"cooldown": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The Cooldown time.",
						},
						"activity_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Actual Type.",
						},
						"related_instances": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"operate_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Operation Type.",
									},
									"status": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Status.",
									},
									"instance_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Instance ID.",
									},
									"message": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The message of Instance.",
									},
								},
							},
							Description: "The related instances.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineScalingActivitiesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewScalingActivityService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineScalingActivities())
}
