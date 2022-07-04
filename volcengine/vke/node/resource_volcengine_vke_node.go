package node

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
VKE node can be imported using the node id, e.g.
```
$ terraform import volcengine_vke_node.default nc5t5epmrsf****
```

*/

func ResourceVolcengineVkeNode() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVkeNodeCreate,
		Read:   resourceVolcengineVkeNodeRead,
		Update: resourceVolcengineVkeNodeUpdate,
		Delete: resourceVolcengineVkeNodeDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
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
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The instance id.",
			},
			"keep_instance_name": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					return len(d.Id()) != 0
				},
				Description: "The flag of keep instance name.",
			},
			"additional_container_storage_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				ForceNew:    true,
				Description: "The flag of additional container storage enable.",
			},
			"container_storage_path": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					flag := d.Get("additional_container_storage_enabled")
					return flag == nil || !flag.(bool)
				},
				Description: "The container storage path.",
			},
			"node_pool_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The node pool id.",
			},
		},
	}
	return resource
}

func resourceVolcengineVkeNodeCreate(d *schema.ResourceData, meta interface{}) (err error) {
	nodeService := NewVolcengineVkeNodeService(meta.(*ve.SdkClient))
	err = nodeService.Dispatcher.Create(nodeService, d, ResourceVolcengineVkeNode())
	if err != nil {
		return fmt.Errorf("error on creating vke node  %q, %s", d.Id(), err)
	}
	return resourceVolcengineVkeNodeRead(d, meta)
}

func resourceVolcengineVkeNodeRead(d *schema.ResourceData, meta interface{}) (err error) {
	nodeService := NewVolcengineVkeNodeService(meta.(*ve.SdkClient))
	err = nodeService.Dispatcher.Read(nodeService, d, ResourceVolcengineVkeNode())
	if err != nil {
		return fmt.Errorf("error on reading vke node %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineVkeNodeUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	nodeService := NewVolcengineVkeNodeService(meta.(*ve.SdkClient))
	err = nodeService.Dispatcher.Update(nodeService, d, ResourceVolcengineVkeNode())
	if err != nil {
		return fmt.Errorf("error on updating vke node  %q, %s", d.Id(), err)
	}
	return resourceVolcengineVkeNodeRead(d, meta)
}

func resourceVolcengineVkeNodeDelete(d *schema.ResourceData, meta interface{}) (err error) {
	nodeService := NewVolcengineVkeNodeService(meta.(*ve.SdkClient))
	err = nodeService.Dispatcher.Delete(nodeService, d, ResourceVolcengineVkeNode())
	if err != nil {
		return fmt.Errorf("error on deleting vke node %q, %s", d.Id(), err)
	}
	return err
}
