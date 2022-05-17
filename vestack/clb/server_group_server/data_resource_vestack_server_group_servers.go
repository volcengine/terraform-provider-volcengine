package server_group_server

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-vestack/common"
)

func DataSourceVestackServerGroupServers() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVestackServerGroupServersRead,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The list of ServerGroupServer IDs.",
			},
			"name_regex": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringIsValidRegExp,
				Description:  "A Name Regex of ServerGroupServer.",
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
			"server_group_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the ServerGroup.",
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
							Description: "The type of instance. Optional choice contains `ecs`, `eni`.",
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

func dataSourceVestackServerGroupServersRead(d *schema.ResourceData, meta interface{}) error {
	serverGroupServerService := NewServerGroupServerService(meta.(*ve.SdkClient))
	return serverGroupServerService.Dispatcher.Data(serverGroupServerService, d, DataSourceVestackServerGroupServers())
}
