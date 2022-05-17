package ecs_instance_state

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-vestack/common"
)

/*

Import
State Instance can be imported using the id, e.g.
```
$ terraform import vestack_ecs_instance_state.default state:i-mizl7m1kqccg5smt1bdpijuj
```

*/

func ResourceVestackEcsInstanceState() *schema.Resource {
	return &schema.Resource{
		Delete: resourceVestackEcsInstanceStateDelete,
		Create: resourceVestackEcsInstanceStateCreate,
		Read:   resourceVestackEcsInstanceStateRead,
		Update: resourceVestackEcsInstanceStateUpdate,
		Importer: &schema.ResourceImporter{
			State: ecsInstanceStateImporter,
		},
		Schema: map[string]*schema.Schema{
			"action": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"Start", "Stop"}, false),
				Description:  "Start or Stop of Instance Action.",
			},
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Id of Instance.",
			},
			"stopped_mode": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "KeepCharging",
				ValidateFunc: validation.StringInSlice([]string{"KeepCharging", "StopCharging"}, false),
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if d.Get("action").(string) == "Stop" {
						return true
					}
					return false
				},
				Description: "Stop Mode of Instance.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Status of Instance.",
			},
		},
	}
}

func resourceVestackEcsInstanceStateCreate(d *schema.ResourceData, meta interface{}) error {
	instanceStateService := NewInstanceStateService(meta.(*ve.SdkClient))
	if err := instanceStateService.Dispatcher.Create(instanceStateService, d, ResourceVestackEcsInstanceState()); err != nil {
		return fmt.Errorf("error on creating instance state %q, %w", d.Id(), err)
	}
	return resourceVestackEcsInstanceStateRead(d, meta)
}

func resourceVestackEcsInstanceStateRead(d *schema.ResourceData, meta interface{}) error {
	instanceStateService := NewInstanceStateService(meta.(*ve.SdkClient))
	if err := instanceStateService.Dispatcher.Read(instanceStateService, d, ResourceVestackEcsInstanceState()); err != nil {
		return fmt.Errorf("error on reading instance state %q, %w", d.Id(), err)
	}
	return nil
}

func resourceVestackEcsInstanceStateUpdate(d *schema.ResourceData, meta interface{}) error {
	instanceStateService := NewInstanceStateService(meta.(*ve.SdkClient))
	if err := instanceStateService.Dispatcher.Update(instanceStateService, d, ResourceVestackEcsInstanceState()); err != nil {
		return fmt.Errorf("error on updating instance state %q, %w", d.Id(), err)
	}
	return resourceVestackEcsInstanceStateRead(d, meta)
}

func resourceVestackEcsInstanceStateDelete(d *schema.ResourceData, meta interface{}) error {
	instanceStateService := NewInstanceStateService(meta.(*ve.SdkClient))
	if err := instanceStateService.Dispatcher.Delete(instanceStateService, d, ResourceVestackEcsInstanceState()); err != nil {
		return fmt.Errorf("error on deleting instance state %q, %w", d.Id(), err)
	}
	return nil
}
