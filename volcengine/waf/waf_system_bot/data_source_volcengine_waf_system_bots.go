package waf_system_bot

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineWafSystemBots() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineWafSystemBotsRead,
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
				Description: "Domain name information.",
			},
			"data": {
				Description: "Host the Bot configuration information.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of Bot.",
						},
						"action": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The execution action of the Bot.",
						},
						"enable": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Whether to enable Bot.",
						},
						"rule_tag": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The rule ID corresponding to Bot.",
						},
						"bot_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of Bot.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineWafSystemBotsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewWafSystemBotService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineWafSystemBots())
}
