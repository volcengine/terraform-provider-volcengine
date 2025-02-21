package rabbitmq_instance_plugin

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineRabbitmqInstancePlugins() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineRabbitmqInstancePluginsRead,
		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The id of rabbitmq instance.",
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
			"plugins": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"plugin_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of plugin.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of plugin.",
						},
						"disable_prompt": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The disable prompt of plugin.",
						},
						"enable_prompt": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The enable prompt of plugin.",
						},
						"need_reboot_on_change": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Will changing the enabled state of the plugin cause a reboot of the rabbitmq instance.",
						},
						"version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The version of plugin.",
						},
						"port": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The port of plugin.",
						},
						"enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether plugin is enabled.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineRabbitmqInstancePluginsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewRabbitmqInstancePluginService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineRabbitmqInstancePlugins())
}
