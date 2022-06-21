package node

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-vestack/common"
)

/*

Import
VKE node can be imported using the node id, e.g.
```
$ terraform import vestack_vke_node.default nc5t5epmrsf****
```

*/

func ResourceVestackVkeNode() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVestackVkeNodeCreate,
		Read:   resourceVestackVkeNodeRead,
		Update: resourceVestackVkeNodeUpdate,
		Delete: resourceVestackVkeNodeDelete,
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
				Description: "The instance ids.",
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
			"cascading_delete_resources": {
				Type: schema.TypeSet,
				Set:  schema.HashString,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.StringInSlice([]string{"Ecs"}, false),
				},
				Optional:    true,
				Description: "Is cascading delete resource.",
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

func resourceVestackVkeNodeCreate(d *schema.ResourceData, meta interface{}) (err error) {
	nodeService := NewVestackVkeNodeService(meta.(*ve.SdkClient))
	err = nodeService.Dispatcher.Create(nodeService, d, ResourceVestackVkeNode())
	if err != nil {
		return fmt.Errorf("error on creating vke node  %q, %s", d.Id(), err)
	}
	return resourceVestackVkeNodeRead(d, meta)
}

func resourceVestackVkeNodeRead(d *schema.ResourceData, meta interface{}) (err error) {
	nodeService := NewVestackVkeNodeService(meta.(*ve.SdkClient))
	err = nodeService.Dispatcher.Read(nodeService, d, ResourceVestackVkeNode())
	if err != nil {
		return fmt.Errorf("error on reading vke node %q, %s", d.Id(), err)
	}
	return err
}

func resourceVestackVkeNodeUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	nodeService := NewVestackVkeNodeService(meta.(*ve.SdkClient))
	err = nodeService.Dispatcher.Update(nodeService, d, ResourceVestackVkeNode())
	if err != nil {
		return fmt.Errorf("error on updating vke node  %q, %s", d.Id(), err)
	}
	return resourceVestackVkeNodeRead(d, meta)
}

func resourceVestackVkeNodeDelete(d *schema.ResourceData, meta interface{}) (err error) {
	nodeService := NewVestackVkeNodeService(meta.(*ve.SdkClient))
	err = nodeService.Dispatcher.Delete(nodeService, d, ResourceVestackVkeNode())
	if err != nil {
		return fmt.Errorf("error on deleting vke node %q, %s", d.Id(), err)
	}
	return err
}
