package alb_customized_cfg

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineAlbCustomizedCfgs() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineAlbCustomizedCfgsRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of CustomizedCfg IDs.",
			},
			"listener_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The id of the listener.",
			},
			"customized_cfg_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the CustomizedCfg.",
			},
			"tags": ve.TagsSchema(),
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of CustomizedCfg.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The project name of the CustomizedCfg.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of CustomizedCfg query.",
			},
			"cfgs": {
				Description: "The collection of CustomizedCfg query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of CustomizedCfg.",
						},
						"customized_cfg_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of CustomizedCfg.",
						},
						"customized_cfg_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of CustomizedCfg.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of CustomizedCfg.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of CustomizedCfg.",
						},
						"customized_cfg_content": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The content of CustomizedCfg.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of CustomizedCfg.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of CustomizedCfg.",
						},
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The project name of CustomizedCfg.",
						},
						"tags": ve.TagsSchemaComputed(),
						"listeners": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The listeners of CustomizedCfg.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"listener_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The ID of Listener.",
									},
									"listener_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Name of Listener.",
									},
									"port": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The port info of listener.",
									},
									"protocol": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The protocol info of listener.",
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

func dataSourceVolcengineAlbCustomizedCfgsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewAlbCustomizedCfgService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineAlbCustomizedCfgs())
}
