package ecs_invocation_result

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineEcsInvocationResults() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineEcsInvocationResultsRead,
		Schema: map[string]*schema.Schema{
			"invocation_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of ecs invocation.",
			},
			"instance_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of ecs instance.",
			},
			"command_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of ecs command.",
			},
			"invocation_result_status": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
					ValidateFunc: validation.StringInSlice([]string{
						"Pending", "Running", "Success", "Failed", "Timeout",
					}, false),
				},
				Set:         schema.HashString,
				Description: "The list of status of ecs invocation in a single instance. Valid values: `Pending`, `Running`, `Success`, `Failed`, `Timeout`.",
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
			"invocation_results": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the ecs invocation result.",
						},
						"invocation_result_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the ecs invocation result.",
						},
						"invocation_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the ecs invocation.",
						},
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the ecs instance.",
						},
						"command_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the ecs command.",
						},
						"invocation_result_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of ecs invocation in a single instance.",
						},
						"username": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The username of the ecs command.",
						},
						"output": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The base64 encoded output message of the ecs invocation. ",
						},
						"exit_code": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The exit code of the ecs command.",
						},
						"error_code": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The error code of the ecs invocation.",
						},
						"error_message": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The error message of the ecs invocation.",
						},
						"start_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The start time of the ecs invocation in the instance.",
						},
						"end_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The end time of the ecs invocation in the instance.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineEcsInvocationResultsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewEcsInvocationResultService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineEcsInvocationResults())
}
