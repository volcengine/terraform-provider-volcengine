package vmp_notify_template

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineVmpNotifyTemplates() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineVmpNotifyTemplatesRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of IDs.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of notify template. This field support fuzzy query.",
			},
			"channel": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The channel of notify template. Valid values: `LarkBotWebhook`, `DingTalkBotWebhook`, `WeComBotWebhook`.",
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

			"notify_templates": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of notify template.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of notify template.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of notify template.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of notify template.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of notify template.",
						},
						"channel": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The channel of notify template. Valid values: `LarkBotWebhook`, `DingTalkBotWebhook`, `WeComBotWebhook`.",
						},
						"active": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The active notify template info.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"title": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The title of notify template.",
									},
									"content": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The content of notify template.",
									},
								},
							},
						},
						"resolved": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The resolved notify template info.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"title": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The title of notify template.",
									},
									"content": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The content of notify template.",
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

func dataSourceVolcengineVmpNotifyTemplatesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewVmpNotifyTemplateService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineVmpNotifyTemplates())
}
