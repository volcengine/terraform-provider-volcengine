package alb_server_group

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineAlbServerGroups() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineAlbServerGroupsRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of Alb server group IDs.",
			},
			"server_group_names": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of Alb server group name.",
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The vpc id of Alb server group.",
			},
			"server_group_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The type of Alb server group. Valid values: `instance`, `ip`.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The project name of Alb server group.",
			},
			"tags": ve.TagsSchema(),
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
			"server_groups": {
				Description: "The collection of query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the Alb server group.",
						},
						"server_group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the Alb server group.",
						},
						"server_group_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the Alb server group.",
						},
						"server_group_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the Alb server group.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the Alb server group.",
						},
						"vpc_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The vpc id of the Alb server group.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the Alb server group.",
						},
						"scheduler": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The scheduler algorithm of the Alb server group.",
						},
						"protocol": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The backend protocol of the Alb server group.",
						},
						"cross_zone_enabled": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Whether to enable cross-zone load balancing for the server group.",
						},
						"server_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The server count of the Alb server group.",
						},
						"project_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The project name of the Alb server group.",
						},
						"create_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The create time of the Alb server group.",
						},
						"update_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The update time of the Alb server group.",
						},
						"ip_address_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ip address type of the server group.",
						},
						"tags": ve.TagsSchemaComputed(),
						"listeners": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
							Description: "The listener information of the Alb server group.",
						},
						"health_check": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The health check config of the Alb server group.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The enable status of health check function.",
									},
									"interval": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The interval executing health check.",
									},
									"timeout": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The response timeout of health check.",
									},
									"healthy_threshold": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The healthy threshold of health check.",
									},
									"unhealthy_threshold": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The unhealthy threshold of health check.",
									},
									"method": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The method of health check.",
									},
									"domain": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The domain of health check.",
									},
									"uri": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The uri of health check.",
									},
									"http_code": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The normal http status code of health check.",
									},
									"protocol": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The protocol of health check.",
									},
									"port": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The port of health check.",
									},
									"http_version": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The http version of health check.",
									},
								},
							},
						},
						"sticky_session_config": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The sticky session config of the Alb server group.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"sticky_session_enabled": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The enable status of sticky session.",
									},
									"sticky_session_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The cookie handle type of the sticky session.",
									},
									"cookie": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The cookie name of the sticky session.",
									},
									"cookie_timeout": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The cookie timeout of the sticky session.",
									},
								},
							},
						},
						"servers": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The server information of the Alb server group.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"server_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The id of the server group server.",
									},
									"instance_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The id of the ecs instance or the network interface.",
									},
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type of the server group server.",
									},
									"remote_enabled": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Whether to enable the remote IP function.",
									},
									"weight": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The weight of the server group server.",
									},
									"ip": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The private ip of the server group server.",
									},
									"port": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The port receiving request of the server group server.",
									},
									"description": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The description of the server group server.",
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

func dataSourceVolcengineAlbServerGroupsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewAlbServerGroupService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineAlbServerGroups())
}
