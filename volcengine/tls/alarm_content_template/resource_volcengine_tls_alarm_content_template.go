package alarm_content_template

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
tls alarm content template can be imported using the alarm_content_template_id, e.g.
```
$ terraform import volcengine_tls_alarm_content_template.default alarm-content-template-123456
```

*/

func ResourceVolcengineTlsAlarmContentTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineTlsAlarmContentTemplateCreate,
		Read:   resourceVolcengineTlsAlarmContentTemplateRead,
		Update: resourceVolcengineTlsAlarmContentTemplateUpdate,
		Delete: resourceVolcengineTlsAlarmContentTemplateDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"alarm_content_template_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the alarm content template.",
			},
			"sms": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The sms content of the alarm content template.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"content": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The content of the sms content template.",
						},
						"locale": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The locale of the sms content template.",
						},
					},
				},
			},
			"vms": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The vms content of the alarm content template.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"content": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The content of the vms content template.",
						},
						"locale": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The locale of the vms content template.",
						},
					},
				},
			},
			"lark": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The lark content of the alarm content template.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"content": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The content of the lark content template.",
						},
						"locale": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The locale of the lark content template.",
						},
						"title": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The title of the lark content template.",
						},
					},
				},
			},
			"email": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The email content of the alarm content template.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"content": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The content of the email content template.",
						},
						"locale": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The locale of the email content template.",
						},
						"subject": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The subject of the email content template.",
						},
					},
				},
			},
			"wechat": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The wechat content of the alarm content template.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"content": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The content of the wechat content template.",
						},
						"locale": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The locale of the wechat content template.",
						},
					},
				},
			},
			"webhook": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The webhook content of the alarm content template.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"content": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The content of the webhook content template.",
						},
					},
				},
			},
			"ding_talk": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The ding_talk content of the alarm content template.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"content": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The content of the ding_talk content template.",
						},
						"locale": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The locale of the ding_talk content template.",
						},
						"title": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The title of the ding_talk content template.",
						},
					},
				},
			},
			"need_valid_content": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether to validate the content template.",
			},
			"alarm_content_template_id": {
				Type:        schema.TypeString,
				Computed:    true, // 由后端返回，不可手动设置
				Description: "The ID of the alarm content template.",
			},
		},
	}
}

func resourceVolcengineTlsAlarmContentTemplateCreate(d *schema.ResourceData, meta interface{}) error {
	service := NewTlsAlarmContentTemplateService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Create(service, d, ResourceVolcengineTlsAlarmContentTemplate())
}

func resourceVolcengineTlsAlarmContentTemplateRead(d *schema.ResourceData, meta interface{}) error {
	service := NewTlsAlarmContentTemplateService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Read(service, d, ResourceVolcengineTlsAlarmContentTemplate())
}

func resourceVolcengineTlsAlarmContentTemplateUpdate(d *schema.ResourceData, meta interface{}) error {
	service := NewTlsAlarmContentTemplateService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Update(service, d, ResourceVolcengineTlsAlarmContentTemplate())
}

func resourceVolcengineTlsAlarmContentTemplateDelete(d *schema.ResourceData, meta interface{}) error {
	service := NewTlsAlarmContentTemplateService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineTlsAlarmContentTemplate())
}
