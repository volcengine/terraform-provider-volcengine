package ecs_invocation

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineEcsInvocations() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineEcsInvocationsRead,
		Schema: map[string]*schema.Schema{
			"invocation_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of ecs invocation.",
			},
			"invocation_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of ecs invocation. This field support fuzzy query.",
			},
			"command_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of ecs command.",
			},
			"command_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of ecs command. This field support fuzzy query.",
			},
			"command_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The type of ecs command. Valid values: `Shell`.",
			},
			"repeat_mode": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"Once",
					"Rate",
					"Fixed",
				}, false),
				Description: "The repeat mode of ecs invocation. Valid values: `Once`, `Rate`, `Fixed`.",
			},
			"invocation_status": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
					ValidateFunc: validation.StringInSlice([]string{
						"Pending", "Scheduled", "Running", "Success",
						"Failed", "Stopped", "PartialFailed", "Finished",
					}, false),
				},
				Set:         schema.HashString,
				Description: "The list of status of ecs invocation. Valid values: `Pending`, `Scheduled`, `Running`, `Success`, `Failed`, `Stopped`, `PartialFailed`, `Finished`.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of Resource.",
			},
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
			"invocations": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the ecs invocation.",
						},
						"invocation_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the ecs invocation.",
						},
						"invocation_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the ecs invocation.",
						},
						"invocation_description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the ecs invocation.",
						},
						"command_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the ecs command.",
						},
						"command_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the ecs command.",
						},
						"command_description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the ecs command.",
						},
						"command_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the ecs command.",
						},
						"command_provider": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The provider of the ecs command.",
						},
						"invocation_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the ecs invocation.",
						},
						"command_content": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The base64 encoded content of the ecs command.",
						},
						"timeout": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The timeout of the ecs command.",
						},
						"working_dir": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The working directory of the ecs command.",
						},
						"username": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The username of the ecs command.",
						},
						"start_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The start time of the ecs invocation.",
						},
						"end_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The end time of the ecs invocation.",
						},
						"instance_number": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The instance number of the ecs invocation.",
						},
						"instance_ids": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The list of ECS instance IDs.",
						},
						"repeat_mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The repeat mode of the ecs invocation.",
						},
						"frequency": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The frequency of the ecs invocation.",
						},
						"launch_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The launch time of the ecs invocation.",
						},
						"recurrence_end_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The recurrence end time of the ecs invocation.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineEcsInvocationsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewEcsInvocationService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineEcsInvocations())
}
