package default_node_pool_batch_attach

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vke/default_node_pool"
)

/*

The resource not support import

*/

func ResourceVolcengineDefaultNodePoolBatchAttach() *schema.Resource {
	m := map[string]*schema.Schema{
		"cluster_id": default_node_pool.ResourceVolcengineDefaultNodePool().Schema["cluster_id"],
		"default_node_pool_id": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "The default NodePool ID.",
		},
		"instances": default_node_pool.ResourceVolcengineDefaultNodePool().Schema["instances"],
		"kubernetes_config": {
			Type:     schema.TypeList,
			MaxItems: 1,
			Optional: true,
			ForceNew: true,
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
									Required:    true,
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
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"key": {
									Type:        schema.TypeString,
									Required:    true,
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
									Type:     schema.TypeString,
									Optional: true,
									ForceNew: true,
									ValidateFunc: validation.StringInSlice([]string{
										"NoSchedule",
										"NoExecute",
										"PreferNoSchedule",
									}, false),
									Description: "The Effect of Taints. The value can be one of the following: `NoSchedule`, `NoExecute`, `PreferNoSchedule`, default value is `NoSchedule`.",
								},
							},
						},
						Description: "The Taints of KubernetesConfig.",
					},
					"cordon": {
						Type:        schema.TypeBool,
						Optional:    true,
						ForceNew:    true,
						Description: "The Cordon of KubernetesConfig.",
					},
				},
			},
			Description: "The KubernetesConfig of NodeConfig. Please note that this field is the configuration of the node. The same key is subject to the config of the node pool. Different keys take effect together.",
		},
	}
	ve.MergeDateSourceToResource(default_node_pool.ResourceVolcengineDefaultNodePool().Schema, &m)

	// logger.Debug(logger.RespFormat, "ATTACH_TEST", m)

	return &schema.Resource{
		Create: resourceVolcengineDefaultNodePoolBatchAttachCreate,
		Update: resourceVolcengineDefaultNodePoolBatchAttachUpdate,
		Read:   resourceVolcengineDefaultNodePoolBatchAttachUpdate,
		Delete: resourceVolcengineNodePoolBatchAttachDelete,
		Importer: &schema.ResourceImporter{
			State: func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
				return nil, fmt.Errorf("The resource not support import ")
			},
		},
		Schema: m,
	}
}

func resourceVolcengineDefaultNodePoolBatchAttachCreate(d *schema.ResourceData, meta interface{}) (err error) {
	nodePoolService := NewVolcengineVkeDefaultNodePoolBatchAttachService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(nodePoolService, d, ResourceVolcengineDefaultNodePoolBatchAttach())
	if err != nil {
		return fmt.Errorf("error on creating DefaultNodePoolBatchAttach  %q, %w", d.Id(), err)
	}
	return resourceVolcengineDefaultNodePoolBatchAttachRead(d, meta)
}

func resourceVolcengineDefaultNodePoolBatchAttachRead(d *schema.ResourceData, meta interface{}) (err error) {
	nodePoolService := NewVolcengineVkeDefaultNodePoolBatchAttachService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(nodePoolService, d, ResourceVolcengineDefaultNodePoolBatchAttach())
	if err != nil {
		return fmt.Errorf("error on reading DefaultNodePoolBatchAttach %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineDefaultNodePoolBatchAttachUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	nodePoolService := NewVolcengineVkeDefaultNodePoolBatchAttachService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(nodePoolService, d, ResourceVolcengineDefaultNodePoolBatchAttach())
	if err != nil {
		return fmt.Errorf("error on updating DefaultNodePoolBatchAttach  %q, %w", d.Id(), err)
	}
	return resourceVolcengineDefaultNodePoolBatchAttachRead(d, meta)
}

func resourceVolcengineNodePoolBatchAttachDelete(d *schema.ResourceData, meta interface{}) (err error) {
	nodePoolService := NewVolcengineVkeDefaultNodePoolBatchAttachService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(nodePoolService, d, ResourceVolcengineDefaultNodePoolBatchAttach())
	if err != nil {
		return fmt.Errorf("error on deleting DefaultNodePoolBatchAttach %q, %w", d.Id(), err)
	}
	return err
}
