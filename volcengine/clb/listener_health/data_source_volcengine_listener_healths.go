package listener_health

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineListenerHealths() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineListenerHealthsRead,
		Schema: map[string]*schema.Schema{
			"listener_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the listener.",
			},
			"only_un_healthy": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether to return only unhealthy backend servers. Valid values: `true`, `false`.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of backend servers.",
			},
			"health_info": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The health info of backend servers.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"un_healthy_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The count of unhealthy backend servers.",
						},
						"listener_status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The health check status of the listener. Valid values: `Active`, `Error`, `Disabled`, `NoTarget`.",
						},
						"results": {
							Description: "The backend server health status results.",
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"server_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The backend server ID.",
									},
									"server_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The backend server type. Valid values: `ecs`, `eni`.",
									},
									"instance_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The ECS instance or ENI ID.",
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
										Description: "The number of forwarding rules associated with the backend server. TCP/UDP listeners return 0.",
									},
									"server_group_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The server group ID that the backend server belongs to.",
									},
									"status": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The health status of the backend server. Valid values: `Up`, `Down`.",
									},
									"updated_at": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The last update time of the backend server.",
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

func dataSourceVolcengineListenerHealthsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewListenerHealthService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineListenerHealths())
}
