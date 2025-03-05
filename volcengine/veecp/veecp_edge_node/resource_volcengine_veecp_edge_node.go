package veecp_edge_node

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
VeecpNode can be imported using the id, e.g.
```
$ terraform import volcengine_veecp_node.default resource_id
```

*/

func ResourceVolcengineVeecpNode() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVeecpNodeCreate,
		Read:   resourceVolcengineVeecpNodeRead,
		Delete: resourceVolcengineVeecpNodeDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"client_token": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "The client token.",
			},
			"cluster_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The cluster id.",
			},
			"node_pool_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The node pool id.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of node.",
			},
			"auto_complete_config": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "Machine information to be managed.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable": {
							Type:        schema.TypeBool,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Description: "Enable/Disable automatic management.",
						},
						"address": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Description: "The address of the machine to be managed.",
						},
						"machine_auth": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Description: "Login credentials.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auth_type": {
										Type:        schema.TypeString,
										Required:    true,
										ForceNew:    true,
										Description: "Authentication method. Currently only Password is open.",
									},
									"user": {
										Type:        schema.TypeString,
										Required:    true,
										ForceNew:    true,
										Description: "Login username.",
									},
									"ssh_port": {
										Type:        schema.TypeInt,
										Required:    true,
										ForceNew:    true,
										Description: "SSH port, default 22.",
									},
								},
							},
						},
						"direct_add": {
							Type:        schema.TypeBool,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Description: "Directly managed through the edge computing instance ID. When it is true, there is no need to provide Address. Only DirectAddInstances needs to be provided.",
						},
						"direct_add_instances": {
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Description: "Edge computing instance ID on Volcano Engine.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"instance_identity": {
										Type:        schema.TypeString,
										Required:    true,
										ForceNew:    true,
										Description: "Edge computing instance ID.",
									},
									"cloud_server_identity": {
										Type:        schema.TypeString,
										Required:    true,
										ForceNew:    true,
										Description: "Edge service ID.",
									},
								},
							},
						},
					},
				},
			},
		},
	}
	return resource
}

func resourceVolcengineVeecpNodeCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVeecpNodeService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineVeecpNode())
	if err != nil {
		return fmt.Errorf("error on creating veecp_node %q, %s", d.Id(), err)
	}
	return resourceVolcengineVeecpNodeRead(d, meta)
}

func resourceVolcengineVeecpNodeRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVeecpNodeService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineVeecpNode())
	if err != nil {
		return fmt.Errorf("error on reading veecp_node %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineVeecpNodeDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVeecpNodeService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineVeecpNode())
	if err != nil {
		return fmt.Errorf("error on deleting veecp_node %q, %s", d.Id(), err)
	}
	return err
}
