package vmp_integration_task_enable

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
VMP Integration Task Enable can be imported using the task ids, e.g.
```
$ terraform import volcengine_vmp_integration_task_enable.default task-id1,task-id2
```

*/

func ResourceVolcengineVmpIntegrationTaskEnable() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVmpIntegrationTaskEnableCreate,
		Read:   resourceVolcengineVmpIntegrationTaskEnableRead,
		Delete: resourceVolcengineVmpIntegrationTaskEnableDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"task_ids": {
				Type:     schema.TypeSet,
				Required: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "A list of integration task IDs to enable.",
			},
		},
	}
	return resource
}

func resourceVolcengineVmpIntegrationTaskEnableCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(service, d, ResourceVolcengineVmpIntegrationTaskEnable())
	if err != nil {
		return fmt.Errorf("error on enabling integration tasks %q, %s", d.Id(), err)
	}
	return resourceVolcengineVmpIntegrationTaskEnableRead(d, meta)
}

func resourceVolcengineVmpIntegrationTaskEnableRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(service, d, ResourceVolcengineVmpIntegrationTaskEnable())
	if err != nil {
		return fmt.Errorf("error on reading integration task enable %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineVmpIntegrationTaskEnableUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	// This resource does not support update operations
	return fmt.Errorf("vmp_integration_task_enable does not support update operations")
}

func resourceVolcengineVmpIntegrationTaskEnableDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineVmpIntegrationTaskEnable())
	if err != nil {
		return fmt.Errorf("error on disabling integration tasks %q, %s", d.Id(), err)
	}
	return err
}
