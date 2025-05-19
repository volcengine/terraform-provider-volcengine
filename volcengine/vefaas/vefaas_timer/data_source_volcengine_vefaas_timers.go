package vefaas_timer

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineVefaasTimers() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineVefaasTimersRead,
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
			"function_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of Function.",
			},
			"items": {
				Description: "The list of timer trigger.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the Timer trigger.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the Timer trigger.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The category of the Timer trigger.",
						},
						"enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the Timer trigger is enabled.",
						},
						"function_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of Function.",
						},
						"account_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The ID of account.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the Timer trigger.",
						},
						"creation_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The creation time of the Timer trigger.",
						},
						"image_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The image version of the Timer trigger.",
						},
						"detailed_config": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The details of trigger configuration.",
						},
						"last_update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The last update time of the Timer trigger.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineVefaasTimersRead(d *schema.ResourceData, meta interface{}) error {
	service := NewVefaasTimerService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineVefaasTimers())
}
