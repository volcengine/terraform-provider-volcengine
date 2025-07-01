package waf_custom_page

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineWafCustomPages() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineWafCustomPagesRead,
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
			"host": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The domain names that need to be viewed.",
			},
			"rule_tag": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Unique identification of the rules.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the project to which your domain names belong.",
			},
			"data": {
				Description: "Details of the rules.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"advanced": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Whether to configure advanced conditions.",
						},
						"body": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The layout content of the response page.",
						},
						"client_ip": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Fill in ALL, which means this rule will take effect on all IP addresses.",
						},
						"code": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Custom HTTP code returned when the request is blocked. Required if PageMode=0 or 1.",
						},
						"content_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The layout template of the response page. Required if PageMode=0 or 1.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Rule description.",
						},
						"enable": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Whether to enable the rule.",
						},
						"group_id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The ID of the advanced conditional rule group.",
						},
						"header": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Request header information.",
						},
						"host": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Domain name to be protected.",
						},
						"id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The ID of rule.",
						},
						"isolation_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of Region.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Rule name.",
						},
						"page_mode": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The layout template of the response page.",
						},
						"policy": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Action to be taken on requests that match the rule.",
						},
						"redirect_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The path where users should be redirected.",
						},
						"rule_tag": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Unique identification of the rules.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Rule update time.",
						},
						"url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Match the path.",
						},
						"accurate": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Computed:    true,
							Description: "Advanced conditions.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"accurate_rules": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Details of advanced conditions.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"http_obj": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The HTTP object to be added to the advanced conditions.",
												},
												"obj_type": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The matching field for HTTP objects.",
												},
												"opretar": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The logical operator for the condition.",
												},
												"property": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Operate the properties of the http object.",
												},
												"value_string": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The value to be matched.",
												},
											},
										},
									},
									"logic": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The logical relationship of advanced conditions.",
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

func dataSourceVolcengineWafCustomPagesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewWafCustomPageService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineWafCustomPages())
}
