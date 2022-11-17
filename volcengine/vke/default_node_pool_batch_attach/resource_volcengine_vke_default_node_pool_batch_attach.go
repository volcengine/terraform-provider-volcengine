package default_node_pool_batch_attach

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vke/default_node_pool"
)

/*

The resource not support import

*/

func ResourceVolcengineDefaultNodePoolBatchAttach() *schema.Resource {
	m := map[string]*schema.Schema{
		"default_node_pool_id": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "The default NodePool ID.",
		},
	}
	ve.MergeDateSourceToResource(default_node_pool.ResourceVolcengineDefaultNodePool().Schema, &m)
	m["kubernetes_config"].Optional = false
	m["kubernetes_config"].Computed = true

	m["node_config"].Optional = false
	m["node_config"].Required = false
	m["node_config"].Computed = true

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
	err = nodePoolService.Dispatcher.Create(nodePoolService, d, ResourceVolcengineDefaultNodePoolBatchAttach())
	if err != nil {
		return fmt.Errorf("error on creating DefaultNodePoolBatchAttach  %q, %w", d.Id(), err)
	}
	return resourceVolcengineDefaultNodePoolBatchAttachRead(d, meta)
}

func resourceVolcengineDefaultNodePoolBatchAttachRead(d *schema.ResourceData, meta interface{}) (err error) {
	nodePoolService := NewVolcengineVkeDefaultNodePoolBatchAttachService(meta.(*ve.SdkClient))
	err = nodePoolService.Dispatcher.Read(nodePoolService, d, ResourceVolcengineDefaultNodePoolBatchAttach())
	if err != nil {
		return fmt.Errorf("error on reading DefaultNodePoolBatchAttach %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineDefaultNodePoolBatchAttachUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	nodePoolService := NewVolcengineVkeDefaultNodePoolBatchAttachService(meta.(*ve.SdkClient))
	err = nodePoolService.Dispatcher.Update(nodePoolService, d, ResourceVolcengineDefaultNodePoolBatchAttach())
	if err != nil {
		return fmt.Errorf("error on updating DefaultNodePoolBatchAttach  %q, %w", d.Id(), err)
	}
	return resourceVolcengineDefaultNodePoolBatchAttachRead(d, meta)
}

func resourceVolcengineNodePoolBatchAttachDelete(d *schema.ResourceData, meta interface{}) (err error) {
	nodePoolService := NewVolcengineVkeDefaultNodePoolBatchAttachService(meta.(*ve.SdkClient))
	err = nodePoolService.Dispatcher.Delete(nodePoolService, d, ResourceVolcengineDefaultNodePoolBatchAttach())
	if err != nil {
		return fmt.Errorf("error on deleting DefaultNodePoolBatchAttach %q, %w", d.Id(), err)
	}
	return err
}
