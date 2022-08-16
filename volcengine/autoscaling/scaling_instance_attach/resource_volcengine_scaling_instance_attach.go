package scaling_instance_attach

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Scaling instance attachment can be imported using the scaling_group_id, e.g.
```
$ terraform import volcengine_scaling_instance_attach.default scg-mizl7m1kqccg5smt1bdpijuj
```

*/

func ResourceVolcengineScalingInstanceAttach() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineScalingInstanceAttachCreate,
		Read:   resourceVolcengineScalingInstanceAttachRead,
		Update: resourceVolcengineScalingInstanceAttachUpdate,
		Delete: resourceVolcengineScalingInstanceAttachDelete,
		Importer: &schema.ResourceImporter{
			State: scalingInstanceImporter,
		},
		Schema: map[string]*schema.Schema{
			"scaling_group_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the scaling group..",
			},
			"instance_ids": {
				Type:        schema.TypeSet,
				Required:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Description: "The list of instance ids the scaling group.",
			},
		},
	}
	return resource
}

func resourceVolcengineScalingInstanceAttachCreate(d *schema.ResourceData, meta interface{}) (err error) {
	scalingInstanceAttachService := NewScalingInstanceAttachService(meta.(*ve.SdkClient))
	err = scalingInstanceAttachService.Dispatcher.Create(scalingInstanceAttachService, d, ResourceVolcengineScalingInstanceAttach())
	if err != nil {
		return fmt.Errorf("error on creating ScalingInatanceAttach %q, %s", d.Id(), err)
	}
	return resourceVolcengineScalingInstanceAttachRead(d, meta)
}

func resourceVolcengineScalingInstanceAttachRead(d *schema.ResourceData, meta interface{}) (err error) {
	scalingInstanceAttachService := NewScalingInstanceAttachService(meta.(*ve.SdkClient))
	err = scalingInstanceAttachService.Dispatcher.Read(scalingInstanceAttachService, d, ResourceVolcengineScalingInstanceAttach())
	if err != nil {
		return fmt.Errorf("error on reading ScalingInatanceAttach %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineScalingInstanceAttachUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	scalingInstanceAttachService := NewScalingInstanceAttachService(meta.(*ve.SdkClient))
	err = scalingInstanceAttachService.Dispatcher.Update(scalingInstanceAttachService, d, ResourceVolcengineScalingInstanceAttach())
	if err != nil {
		return fmt.Errorf("error on updating ScalingInatanceAttach %q, %s", d.Id(), err)
	}
	return resourceVolcengineScalingInstanceAttachRead(d, meta)
}

func resourceVolcengineScalingInstanceAttachDelete(d *schema.ResourceData, meta interface{}) (err error) {
	scalingInstanceAttachService := NewScalingInstanceAttachService(meta.(*ve.SdkClient))
	err = scalingInstanceAttachService.Dispatcher.Delete(scalingInstanceAttachService, d, ResourceVolcengineScalingInstanceAttach())
	if err != nil {
		return fmt.Errorf("error on deleting ScalingInatanceAttach %q, %s", d.Id(), err)
	}
	return err
}
