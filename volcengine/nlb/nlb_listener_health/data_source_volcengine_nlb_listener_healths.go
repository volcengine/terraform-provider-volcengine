package nlb_listener_health

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineNlbListenerHealths() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineNlbListenerHealthsRead,
		Schema: map[string]*schema.Schema{
			"listener_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the listener.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"listener_healths": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The collection of listener health query.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"listener_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the listener.",
						},
						"server_group_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the server group.",
						},
						"healthy_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of healthy backend servers.",
						},
						"unhealthy_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of unhealthy backend servers.",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the server group. Valid values: `Active`, `NoTarget`, `Error`, `Disabled`.\n`NoTarget`: The server group does not have any backend servers added, or the server group has closed cross-zone forwarding and there are no backend servers in the zone where the access traffic originates.\n`Error`: There are unhealthy backend servers in the server group.\n`Active`: All backend servers in the server group are healthy.\n`Disabled`: The server group has closed health checks, or the NLB instance associated with the server group is in a stopped state.",
						},
						"results": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The list of backend servers.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"server_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The ID of the backend server.",
									},
									"server_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type of the backend server. Valid values: `ecs`, `eni`, `ip`.\n`ecs`: ECS instance (primary network interface).\n`eni`: Secondary network interface.\n`ip`: IP address.",
									},
									"instance_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The ID or IP address of the backend server.",
									},
									"zone_id": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The zone ID that the backend server receives access traffic from.",
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
									"status": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The health status of the backend server. Valid values: `Up`, `Down`, `Unused`.\n`Up`: Normal.\n`Down`: Abnormal.\n`Unused`: Unused. The NLB instance has closed cross-zone forwarding, and there is no access traffic from the zone of the backend server.",
									},
								},
							},
						},
					},
				},
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of listener health query.",
			},
		},
	}
}

func dataSourceVolcengineNlbListenerHealthsRead(d *schema.ResourceData, meta interface{}) error {
	service := NewNlbListenerHealthService(meta.(*ve.SdkClient))
	return ve.DefaultDispatcher().Data(service, d, DataSourceVolcengineNlbListenerHealths())
}
