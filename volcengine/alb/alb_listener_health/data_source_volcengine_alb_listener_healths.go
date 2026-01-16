package alb_listener_health

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineAlbListenerHealths() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineAlbListenerHealthsRead,
		Schema: map[string]*schema.Schema{
			"listener_ids": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of Listener IDs.",
			},
			"only_un_healthy": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether to return only backend servers with abnormal health check status.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The project name of the listener.",
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
			"listeners": {
				Description: "The collection of listener health query.",
				Type:        schema.TypeList,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"listener_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the listener.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the listener. Value: Active, Error, NoTarget, Disabled.",
						},
						"total_backend_server_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The total count of backend servers under the listener.",
						},
						"un_healthy_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The count of backend servers with abnormal health check status.",
						},
						"backend_servers": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The list of backend server health details.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"server_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The ID of the backend server.",
									},
									"instance_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The ID of the ECS instance or ENI.",
									},
									"server_group_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The ID of the backend server group.",
									},
									"server_group_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the backend server group.",
									},
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type of backend server. Value: ecs, eni.",
									},
									"ip": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The IP address of the backend server.",
									},
									"port": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The port of the backend server.",
									},
									"rule_number": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The number of forwarding rules associated with the backend server.",
									},
									"status": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The health status of the backend server. Value: Up, Down.",
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

func dataSourceVolcengineAlbListenerHealthsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewAlbListenerHealthService(meta.(*ve.SdkClient))
	return service.Dispatcher.Data(service, d, DataSourceVolcengineAlbListenerHealths())
}
