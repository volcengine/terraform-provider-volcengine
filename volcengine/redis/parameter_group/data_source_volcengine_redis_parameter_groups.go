package parameter_group

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineParameterGroups() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineParameterGroupsRead,
		Schema: map[string]*schema.Schema{
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
			"engine_version": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Redis database version applicable to the parameter template.",
			},
			"source": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The source of creating the parameter template.",
			},
			"parameter_groups": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The details of the parameter template.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the parameter template.",
						},
						"source": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The source of creating the parameter template.",
						},
						"default": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether it is the default parameter template.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The creation time of the parameter template, in the format of yyyy-MM-ddTHH:mm:ssZ (UTC).",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The last update time of the parameter template, in the format of yyyy-MM-ddTHH:mm:ssZ (UTC).",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description the parameter template.",
						},
						"parameter_num": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of parameters contained in the parameter template.",
						},
						"engine_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The database version applicable to the parameter template.",
						},
						"parameter_group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the parameter template.",
						},
						"parameters": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The list of parameter information contained in the parameter template.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"options": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The optional list of selector type parameters.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"value": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Optional selector type parameters.",
												},
												"description": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The description the Optional parameters.",
												},
											},
										},
									},
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type of the parameter.",
									},
									"unit": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unit of the numerical type parameter.",
									},
									"range": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The value range of numerical type parameters.",
									},
									"need_reboot": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether to restart the instance to take effect after modifying this parameter.",
									},
									"description": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The description the parameter.",
									},
									"current_value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The current running value of the parameter.",
									},
									"param_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of parameter.",
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

func dataSourceVolcengineParameterGroupsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewParameterGroupService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineParameterGroups())
}
