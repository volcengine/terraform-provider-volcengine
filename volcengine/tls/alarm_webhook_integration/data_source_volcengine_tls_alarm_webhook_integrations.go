package alarm_webhook_integration

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineTlsAlarmWebhookIntegrations() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineTlsAlarmWebhookIntegrationsRead,
		Schema: map[string]*schema.Schema{
			"webhook_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The ID of the alarm webhook integration.",
			},
			"webhook_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the webhook integration. Fuzzy matching is supported.",
			},
			"webhook_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The type of the webhook integration.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of alarm webhook integrations.",
			},
			"integrations": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of alarm webhook integrations.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"webhook_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the alarm webhook integration.",
						},
						"webhook_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the webhook integration.",
						},
						"webhook_url": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The URL of the webhook.",
						},
						"webhook_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the webhook.",
						},
						"webhook_method": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The method of the webhook.",
						},
						"webhook_secret": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The secret of the webhook.",
						},
						"webhook_headers": {
							Type: schema.TypeList,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The key of the header.",
									},
									"value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The value of the header.",
									},
								},
							},
							Computed:    true,
							Description: "The headers of the webhook.",
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
				},
			},
		},
	}
}

func dataSourceVolcengineTlsAlarmWebhookIntegrationsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewVolcengineTlsAlarmWebhookIntegrationService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineTlsAlarmWebhookIntegrations())
}
