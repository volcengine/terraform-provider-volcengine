package node

import (
	"fmt"
	"time"

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
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
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
				Description: "The flag of keep instance name, the value is `true` or `false`.",
			},
			"additional_container_storage_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				ForceNew:    true,
				Description: "The flag of additional container storage enable, the value is `true` or `false`.",
			},
			"image_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The ImageId of NodeConfig.",
			},
			"initialize_script": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The initializeScript of Node.",
			},
			"kubernetes_config": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The KubernetesConfig of Node.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"labels": {
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Description: "The Key of Labels.",
									},
									"value": {
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Description: "The Value of Labels.",
									},
								},
							},
							Description: "The Labels of KubernetesConfig.",
						},
						"taints": {
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"key": {
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Description: "The Key of Taints.",
									},
									"value": {
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Description: "The Value of Taints.",
									},
									"effect": {
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Description: "The Effect of Taints, the value can be `NoSchedule` or `NoExecute` or `PreferNoSchedule`.",
									},
								},
							},
							Description: "The Taints of KubernetesConfig.",
						},
						"cordon": {
							Type:        schema.TypeBool,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Description: "The Cordon of KubernetesConfig.",
						},
					},
				},
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
	err = ve.DefaultDispatcher().Create(nodeService, d, ResourceVolcengineVkeNode())
	if err != nil {
		return fmt.Errorf("error on creating vke node  %q, %s", d.Id(), err)
	}
	return resourceVolcengineVkeNodeRead(d, meta)
}

func resourceVolcengineVkeNodeRead(d *schema.ResourceData, meta interface{}) (err error) {
	nodeService := NewVolcengineVkeNodeService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(nodeService, d, ResourceVolcengineVkeNode())
	if err != nil {
		return fmt.Errorf("error on reading vke node %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineVkeNodeUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	nodeService := NewVolcengineVkeNodeService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(nodeService, d, ResourceVolcengineVkeNode())
	if err != nil {
		return fmt.Errorf("error on updating vke node  %q, %s", d.Id(), err)
	}
	return resourceVolcengineVkeNodeRead(d, meta)
}

func resourceVolcengineVkeNodeDelete(d *schema.ResourceData, meta interface{}) (err error) {
	nodeService := NewVolcengineVkeNodeService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(nodeService, d, ResourceVolcengineVkeNode())
	if err != nil {
		return fmt.Errorf("error on deleting vke node %q, %s", d.Id(), err)
	}
	return err
}
