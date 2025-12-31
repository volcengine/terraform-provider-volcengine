package vmp_alert_sample

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineVmpAlertSamples() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineVmpAlertSamplesRead,
		Schema: map[string]*schema.Schema{
			"alert_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Alert ID to filter samples.",
			},
			"sample_since": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Filter start timestamp (unix).",
			},
			"sample_until": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Filter end timestamp (unix).",
			},
			"limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Limit of samples, default 100, max 500.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of query.",
			},
			"alert_samples": {
				Description: "Alert samples collection.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"alert_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Alert ID.",
						},
						"timestamp": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Alert sample timestamp(unix).",
						},
						"phase": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Alert sample phase.",
						},
						"level": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Alert sample level.",
						},
						"value": {
							Type:        schema.TypeFloat,
							Computed:    true,
							Description: "Alert sample value.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineVmpAlertSamplesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewVmpAlertSampleService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineVmpAlertSamples())
}
