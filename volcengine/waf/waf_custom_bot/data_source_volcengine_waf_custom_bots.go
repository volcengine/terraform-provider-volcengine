package waf_custom_bot

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineWafCustomBots() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineWafCustomBotsRead,
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
				Optional:    true,
				Description: "The domain names that need to be viewed.",
			},
			"data": {
				Description: "The Details of Custom bot.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bot_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "bot name.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of bot.",
						},
						"action": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The execution action of the Bot.",
						},
						"enable": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Whether to enable bot.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time.",
						},
						"advanced": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Whether to set advanced conditions.",
						},
						"rule_tag": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Rule unique identifier.",
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
						"id": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The actual count bits of the rule unique identifier (corresponding to the RuleTag).",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineWafCustomBotsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewWafCustomBotService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineWafCustomBots())
}
