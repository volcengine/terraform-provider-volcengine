package scaling_instance_attach

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-vestack/common"
)

/*

Import
Scaling instance attachment can be imported using the scaling_group_id:instance_id.1:instance_id2..., e.g.
```
$ terraform import vestack_scaling_instance_attach.default scg-mizl7m1kqccg5smt1bdpijuj:i-***:i-***
```

*/

func ResourceVestackScalingInstanceAttach() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVestackScalingInstanceAttachCreate,
		Read:   resourceVestackScalingInstanceAttachRead,
		Update: resourceVestackScalingInstanceAttachUpdate,
		Delete: resourceVestackScalingInstanceAttachDelete,
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
				MinItems:    1,
				MaxItems:    20,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
				Description: "The list of instance ids the scaling group.",
			},
		},
	}
	return resource
}

func resourceVestackScalingInstanceAttachCreate(d *schema.ResourceData, meta interface{}) (err error) {
	scalingInstanceAttachService := NewScalingInstanceAttachService(meta.(*ve.SdkClient))
	err = scalingInstanceAttachService.Dispatcher.Create(scalingInstanceAttachService, d, ResourceVestackScalingInstanceAttach())
	if err != nil {
		return fmt.Errorf("error on creating ScalingInatanceAttach %q, %s", d.Id(), err)
	}
	return resourceVestackScalingInstanceAttachRead(d, meta)
}

func resourceVestackScalingInstanceAttachRead(d *schema.ResourceData, meta interface{}) (err error) {
	scalingInstanceAttachService := NewScalingInstanceAttachService(meta.(*ve.SdkClient))
	err = scalingInstanceAttachService.Dispatcher.Read(scalingInstanceAttachService, d, ResourceVestackScalingInstanceAttach())
	if err != nil {
		return fmt.Errorf("error on reading ScalingInatanceAttach %q, %s", d.Id(), err)
	}
	return err
}

func resourceVestackScalingInstanceAttachUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	scalingInstanceAttachService := NewScalingInstanceAttachService(meta.(*ve.SdkClient))
	err = scalingInstanceAttachService.Dispatcher.Update(scalingInstanceAttachService, d, ResourceVestackScalingInstanceAttach())
	if err != nil {
		return fmt.Errorf("error on updating ScalingInatanceAttach %q, %s", d.Id(), err)
	}
	return resourceVestackScalingInstanceAttachRead(d, meta)
}

func resourceVestackScalingInstanceAttachDelete(d *schema.ResourceData, meta interface{}) (err error) {
	scalingInstanceAttachService := NewScalingInstanceAttachService(meta.(*ve.SdkClient))
	err = scalingInstanceAttachService.Dispatcher.Delete(scalingInstanceAttachService, d, ResourceVestackScalingInstanceAttach())
	if err != nil {
		return fmt.Errorf("error on deleting ScalingInatanceAttach %q, %s", d.Id(), err)
	}
	return err
}
