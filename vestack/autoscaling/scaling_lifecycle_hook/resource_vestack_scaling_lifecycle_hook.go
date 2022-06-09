package scaling_lifecycle_hook

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-vestack/common"
)

/*

Import
ScalingLifecycleHook can be imported using the ScalingGroupId:LifecycleHookId, e.g.
```
$ terraform import vestack_scaling_lifecycle_hook.default scg-yblfbfhy7agh9zn72iaz:sgh-ybqholahe4gso0ee88sd
```

*/

func ResourceVestackScalingLifecycleHook() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVestackScalingLifecycleHookCreate,
		Read:   resourceVestackScalingLifecycleHookRead,
		Update: resourceVetackScalingLifecycleHookUpdate,
		Delete: resourceVetackScalingLifecycleHookDelete,
		Importer: &schema.ResourceImporter{
			State: lifecycleHookImporter,
		},
		Schema: map[string]*schema.Schema{
			"scaling_group_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the scaling group.",
			},
			"lifecycle_hook_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the lifecycle hook.",
			},
			"lifecycle_hook_timeout": {
				Type:         schema.TypeInt,
				Required:     true,
				ValidateFunc: validation.IntBetween(30, 21600),
				Description:  "The timeout of the lifecycle hook.",
			},
			"lifecycle_hook_type": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"SCALE_IN", "SCALE_OUT"}, false),
				Description:  "The type of the lifecycle hook.",
			},
			"lifecycle_hook_policy": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"CONTINUE", "REJECT"}, false),
				Description:  "The policy of the lifecycle hook.",
			},
		},
	}
	dataSource := DataSourceVestackScalingLifecycleHooks().Schema["lifecycle_hooks"].Elem.(*schema.Resource).Schema
	delete(dataSource, "id")
	ve.MergeDateSourceToResource(dataSource, &resource.Schema)
	return resource
}

func resourceVestackScalingLifecycleHookCreate(d *schema.ResourceData, meta interface{}) (err error) {
	lifecycleHookService := NewScalingLifecycleHookService(meta.(*ve.SdkClient))
	err = lifecycleHookService.Dispatcher.Create(lifecycleHookService, d, ResourceVestackScalingLifecycleHook())
	if err != nil {
		return fmt.Errorf("error on creating ScalingLifecycleHook %q, %s", d.Id(), err)
	}
	return resourceVestackScalingLifecycleHookRead(d, meta)
}

func resourceVestackScalingLifecycleHookRead(d *schema.ResourceData, meta interface{}) (err error) {
	lifecycleHookService := NewScalingLifecycleHookService(meta.(*ve.SdkClient))
	err = lifecycleHookService.Dispatcher.Read(lifecycleHookService, d, ResourceVestackScalingLifecycleHook())
	if err != nil {
		return fmt.Errorf("error on reading ScalingLifecycleHook %q, %s", d.Id(), err)
	}
	return err
}

func resourceVetackScalingLifecycleHookUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	lifecycleHookService := NewScalingLifecycleHookService(meta.(*ve.SdkClient))
	err = lifecycleHookService.Dispatcher.Update(lifecycleHookService, d, ResourceVestackScalingLifecycleHook())
	if err != nil {
		return fmt.Errorf("error on updating ScalingLifecycleHook %q, %s", d.Id(), err)
	}
	return resourceVestackScalingLifecycleHookRead(d, meta)
}

func resourceVetackScalingLifecycleHookDelete(d *schema.ResourceData, meta interface{}) (err error) {
	lifecycleHookService := NewScalingLifecycleHookService(meta.(*ve.SdkClient))
	err = lifecycleHookService.Dispatcher.Delete(lifecycleHookService, d, ResourceVestackScalingLifecycleHook())
	if err != nil {
		return fmt.Errorf("error on deleting ScalingLifecycleHook %q, %s", d.Id(), err)
	}
	return err
}
