package alb_health_check_template

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineAlbHealthCheckTemplates() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineAlbHealthCheckTemplatesRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "The list of health check templates to query.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"health_check_template_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of health check template to query.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of health check template.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The project name to query.",
			},
			"tags": ve.TagsSchema(),
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of health check template query.",
			},
			"health_check_templates": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The collection of health check template query.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The id of the health check template.",
						},
						"health_check_template_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of health check template.",
						},
						"health_check_template_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of health check template.",
						},
						"health_check_port": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The port for health check. 0 means use backend server port for health check, 1-65535 means use the specified port.",
						},
						"health_check_interval": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The interval for performing health checks, the default value is 2, and the value is 1-300.",
						},
						"health_check_timeout": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The timeout of health check response,the default value is 2, and the value is 1-60.",
						},
						"healthy_threshold": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The healthy threshold of the health check, the default is 3, the value is 2-10.",
						},
						"unhealthy_threshold": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The unhealthy threshold of the health check, the default is 3, the value is 2-10.",
						},
						"health_check_method": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The health check method, support `GET` and `HEAD`.",
						},
						"health_check_domain": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The domain name to health check.",
						},
						"health_check_uri": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The uri to health check,default is `/`.",
						},
						"health_check_http_code": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The normal HTTP status code for health check, the default is http_2xx, http_3xx, separated by commas.",
						},
						"health_check_protocol": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The protocol of health check, support HTTP and TCP.",
						},
						"health_check_http_version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The HTTP version of health check.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of health check template.",
						},
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The project name to which the health check template belongs.",
						},
						"tags": ve.TagsSchemaComputed(),
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The creation time of the health check template.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The last update time of the health check template.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineAlbHealthCheckTemplatesRead(d *schema.ResourceData, meta interface{}) error {
	service := NewAlbHealthCheckTemplateService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineAlbHealthCheckTemplates())
}
