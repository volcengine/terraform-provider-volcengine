package ecs_command

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineEcsCommands() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineEcsCommandsRead,
		Schema: map[string]*schema.Schema{
			"command_provider": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The provider of public command. When this field is not specified, query for custom commands.",
			},
			"command_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of ecs command.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of ecs command. This field support fuzzy query.",
			},
			"type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The type of ecs command. Valid values: `Shell`.",
			},
			"order": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The order of ecs command query result.",
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
			"commands": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the ecs command.",
						},
						"command_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the ecs command.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the ecs command.",
						},
						"created_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the ecs command.",
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of the ecs command.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the ecs command.",
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
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the ecs command.",
						},
						"command_provider": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The provider of the public command.",
						},
						"command_content": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The base64 encoded content of the ecs command.",
						},
						"invocation_times": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The invocation times of the ecs command. Public commands do not display the invocation times.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineEcsCommandsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewEcsCommandService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineEcsCommands())
}
