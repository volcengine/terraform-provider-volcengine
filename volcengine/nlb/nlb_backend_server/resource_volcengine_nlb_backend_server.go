package nlb_backend_server

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*
Import
NlbBackendServers can be imported using the NLB server group ID, e.g.
```
$ terraform import volcengine_nlb_backend_servers.foo rsp-2d6g5cxxx
```
*/
func ResourceVolcengineNlbBackendServers() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineNlbBackendServersCreate,
		Read:   resourceVolcengineNlbBackendServersRead,
		Update: resourceVolcengineNlbBackendServersUpdate,
		Delete: resourceVolcengineNlbBackendServersDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"server_group_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the server group.",
			},
			"backend_servers": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "The list of backend servers.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The type of the backend server. Valid values: `ecs`, `eni`, `ip`.",
						},
						"instance_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The instance id of the backend server.",
						},
						"ip": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The ip of the backend server.",
						},
						"port": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "The port of the backend server.",
						},
						"weight": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     100,
							Description: "The weight of the backend server.",
						},
						"description": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The description of the backend server.",
						},
						"zone_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The zone id of the backend server.",
						},
						"server_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The server id of the backend server.",
						},
					},
				},
			},
		},
	}
}

func resourceVolcengineNlbBackendServersCreate(d *schema.ResourceData, meta interface{}) error {
	service := NewNlbBackendServerService(meta.(*ve.SdkClient))
	err := ve.DefaultDispatcher().Create(service, d, ResourceVolcengineNlbBackendServers())
	if err != nil {
		return err
	}
	return nil
}

func resourceVolcengineNlbBackendServersRead(d *schema.ResourceData, meta interface{}) error {
	service := NewNlbBackendServerService(meta.(*ve.SdkClient))
	err := ve.DefaultDispatcher().Read(service, d, ResourceVolcengineNlbBackendServers())
	if err != nil {
		return err
	}
	return nil
}

func resourceVolcengineNlbBackendServersUpdate(d *schema.ResourceData, meta interface{}) error {
	service := NewNlbBackendServerService(meta.(*ve.SdkClient))
	err := ve.DefaultDispatcher().Update(service, d, ResourceVolcengineNlbBackendServers())
	if err != nil {
		return err
	}
	return nil
}

func resourceVolcengineNlbBackendServersDelete(d *schema.ResourceData, meta interface{}) error {
	service := NewNlbBackendServerService(meta.(*ve.SdkClient))
	err := ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineNlbBackendServers())
	if err != nil {
		return err
	}
	return nil
}
