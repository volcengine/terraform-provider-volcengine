package alb_server_group_server

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

func DataSourceVolcengineAlbServerGroupServers() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVolcengineAlbServerGroupServersRead,
		Schema: map[string]*schema.Schema{
			"server_group_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the ServerGroup.",
			},
			"instance_ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
				Description: "A list of instance IDs. When the backend server is ECS, the parameter value is the ID of the ECS. " +
					"When the backend server is a secondary network interface card, the parameter value is the ID of the secondary network interface card.",
			},
			"ips": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of private IP addresses.",
			},
			"output_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name where to save data source results.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The total count of ServerGroupServer query.",
			},
			"servers": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The server list of ServerGroup.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The server id of instance in ServerGroup.",
						},
						"server_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The server id of instance in ServerGroup.",
						},
						"instance_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of ecs instance or the network card bound to ecs instance.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of instance. Optional choice contains `ecs`, `eni`, `ip`.",
						},
						"remote_enabled": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Whether to enable remote IP function. Optional choice contains `on`, `off`.",
						},
						"weight": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The weight of the instance.",
						},
						"ip": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The private ip of the instance.",
						},
						"port": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The port receiving request.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the instance.",
						},
					},
				},
			},
		},
	}
}

func dataSourceVolcengineAlbServerGroupServersRead(d *schema.ResourceData, meta interface{}) error {
	serverGroupServerService := NewServerGroupServerService(meta.(*ve.SdkClient))
	return serverGroupServerService.Dispatcher.Data(serverGroupServerService, d, DataSourceVolcengineAlbServerGroupServers())
}
