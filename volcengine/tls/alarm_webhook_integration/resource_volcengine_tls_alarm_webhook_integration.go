package alarm_webhook_integration

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
tls alarm webhook integration can be imported using the alarm_webhook_integration_id, e.g.
```
$ terraform import volcengine_tls_alarm_webhook_integration.default alarm-webhook-integration-123456
```

*/

func ResourceVolcengineTlsAlarmWebhookIntegration() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineTlsAlarmWebhookIntegrationCreate,
		Read:   resourceVolcengineTlsAlarmWebhookIntegrationRead,
		Update: resourceVolcengineTlsAlarmWebhookIntegrationUpdate,
		Delete: resourceVolcengineTlsAlarmWebhookIntegrationDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"webhook_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the webhook integration.",
			},
			"webhook_url": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The URL of the webhook.",
			},
			"webhook_type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The type of the webhook integration.",
			},
			"webhook_method": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The method of the webhook.",
			},
			"webhook_secret": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "The secret of the webhook.",
			},
			"webhook_headers": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The headers of the webhook.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The key of the header.",
						},
						"value": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The value of the header.",
						},
					},
				},
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The creation time of the webhook integration.",
			},
			"modify_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The update time of the webhook integration.",
			},
		},
	}
}

func resourceVolcengineTlsAlarmWebhookIntegrationCreate(d *schema.ResourceData, meta interface{}) error {
	service := NewVolcengineTlsAlarmWebhookIntegrationService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Create(service, d, ResourceVolcengineTlsAlarmWebhookIntegration())
}

func resourceVolcengineTlsAlarmWebhookIntegrationRead(d *schema.ResourceData, meta interface{}) error {
	service := NewVolcengineTlsAlarmWebhookIntegrationService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Read(service, d, ResourceVolcengineTlsAlarmWebhookIntegration())
}

func resourceVolcengineTlsAlarmWebhookIntegrationUpdate(d *schema.ResourceData, meta interface{}) error {
	service := NewVolcengineTlsAlarmWebhookIntegrationService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Update(service, d, ResourceVolcengineTlsAlarmWebhookIntegration())
}

func resourceVolcengineTlsAlarmWebhookIntegrationDelete(d *schema.ResourceData, meta interface{}) error {
	service := NewVolcengineTlsAlarmWebhookIntegrationService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineTlsAlarmWebhookIntegration())
}
