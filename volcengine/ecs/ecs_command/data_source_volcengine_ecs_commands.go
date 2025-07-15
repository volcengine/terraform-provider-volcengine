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
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The project name of ecs command.",
			},
			"tags": ve.TagsSchema(),
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
						"content_encoding": {
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
							Description: "Whether the command content is base64 encoded. Valid values: `Base64`, `PlainText`. Default is `Base64`.",
						},
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The project name of the ecs command.",
						},
						"tags": ve.TagsSchemaComputed(),
						"enable_parameter": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to enable custom parameter. Default is `false`.",
						},
						"parameter_definitions": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The custom parameter definitions of the ecs command.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the custom parameter.",
									},
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type of the custom parameter. Valid values: `String`, `Digit`.",
									},
									"required": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether the custom parameter is required.",
									},
									"default_value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The default value of the custom parameter.",
									},
									"min_length": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The minimum length of the custom parameter. This field is required when the parameter type is `String`.",
									},
									"max_length": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The maximum length of the custom parameter. This field is required when the parameter type is `String`.",
									},
									"min_value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The minimum value of the custom parameter. This field is required when the parameter type is `Digit`.",
									},
									"max_value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The maximum value of the custom parameter. This field is required when the parameter type is `Digit`.",
									},
									"decimal_precision": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The decimal precision of the custom parameter. This field is required when the parameter type is `Digit`.",
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

func dataSourceVolcengineEcsCommandsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewEcsCommandService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineEcsCommands())
}
