package alarm_content_template

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineTlsAlarmContentTemplates() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineTlsAlarmContentTemplatesRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:        schema.TypeSet,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Optional:    true,
				Set:         schema.HashString,
				Description: "A list of alarm content template IDs.",
			},
			"alarm_content_template_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the alarm content template. Fuzzy matching is supported.",
			},
			"alarm_content_template_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the alarm content template.",
			},
			"order_field": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The order field.",
			},
			"asc": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether to ascend.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of alarm content templates.",
			},
			"templates": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of alarm content templates.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"alarm_content_template_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the alarm content template.",
						},
						"alarm_content_template_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the alarm content template.",
						},
						"content": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The content of the alarm content template.",
						},
						"sms": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The sms content of the alarm content template.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"content": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The content of the sms content template.",
									},
									"locale": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The locale of the sms content template.",
									},
								},
							},
						},
						"vms": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The vms content of the alarm content template.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"content": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The content of the vms content template.",
									},
									"locale": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The locale of the vms content template.",
									},
								},
							},
						},
						"lark": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The lark content of the alarm content template.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"content": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The content of the lark content template.",
									},
									"locale": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The locale of the lark content template.",
									},
									"title": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The title of the lark content template.",
									},
								},
							},
						},
						"email": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The email content of the alarm content template.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"content": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The content of the email content template.",
									},
									"locale": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The locale of the email content template.",
									},
									"subject": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The subject of the email content template.",
									},
								},
							},
						},
						"wechat": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The wechat content of the alarm content template.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"content": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The content of the wechat content template.",
									},
									"locale": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The locale of the wechat content template.",
									},
								},
							},
						},
						"webhook": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The webhook content of the alarm content template.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"content": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The content of the webhook content template.",
									},
								},
							},
						},
						"ding_talk": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The ding_talk content of the alarm content template.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"content": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The content of the ding_talk content template.",
									},
									"locale": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The locale of the ding_talk content template.",
									},
									"title": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The title of the ding_talk content template.",
									},
								},
							},
						},
						"is_default": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the alarm content template is default.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the alarm content template.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the alarm content template.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The creation time of the alarm content template.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of the alarm content template.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineTlsAlarmContentTemplatesRead(d *schema.ResourceData, meta interface{}) error {
	alarmContentTemplateService := NewTlsAlarmContentTemplateService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(alarmContentTemplateService, d, DataSourceVolcengineTlsAlarmContentTemplates())
}
