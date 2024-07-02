package ecs_instance_state

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
State Instance can be imported using the id, e.g.
```
$ terraform import volcengine_ecs_instance_state.default state:i-mizl7m1kqccg5smt1bdpijuj
```

*/

func ResourceVolcengineEcsInstanceState() *schema.Resource {
	return &schema.Resource{
		Delete: resourceVolcengineEcsInstanceStateDelete,
		Create: resourceVolcengineEcsInstanceStateCreate,
		Read:   resourceVolcengineEcsInstanceStateRead,
		Update: resourceVolcengineEcsInstanceStateUpdate,
		Importer: &schema.ResourceImporter{
			State: ecsInstanceStateImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(1 * time.Hour),
			Update: schema.DefaultTimeout(1 * time.Hour),
			Delete: schema.DefaultTimeout(1 * time.Hour),
		},
		Schema: map[string]*schema.Schema{
			"action": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"Start", "Stop", "ForceStop"}, false),
				Description: "Start or Stop of Instance Action, the value can be `Start`, `Stop` or `ForceStop`. \n" +
					"If the target status of the action is consistent with the current status of the instance, the action will not actually be executed.",
			},
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Id of Instance.",
			},
			"stopped_mode": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// 如开机行为，该字段修改忽略
					return d.Get("action").(string) == "Start"
				},
				Description: "Stop Mode of Instance, the value can be `KeepCharging` or `StopCharging`.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Status of Instance.",
			},
		},
	}
}

func resourceVolcengineEcsInstanceStateCreate(d *schema.ResourceData, meta interface{}) error {
	instanceStateService := NewInstanceStateService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Create(instanceStateService, d, ResourceVolcengineEcsInstanceState()); err != nil {
		return fmt.Errorf("error on creating instance state %q, %w", d.Id(), err)
	}
	return resourceVolcengineEcsInstanceStateRead(d, meta)
}

func resourceVolcengineEcsInstanceStateRead(d *schema.ResourceData, meta interface{}) error {
	instanceStateService := NewInstanceStateService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Read(instanceStateService, d, ResourceVolcengineEcsInstanceState()); err != nil {
		return fmt.Errorf("error on reading instance state %q, %w", d.Id(), err)
	}
	return nil
}

func resourceVolcengineEcsInstanceStateUpdate(d *schema.ResourceData, meta interface{}) error {
	instanceStateService := NewInstanceStateService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Update(instanceStateService, d, ResourceVolcengineEcsInstanceState()); err != nil {
		return fmt.Errorf("error on updating instance state %q, %w", d.Id(), err)
	}
	return resourceVolcengineEcsInstanceStateRead(d, meta)
}

func resourceVolcengineEcsInstanceStateDelete(d *schema.ResourceData, meta interface{}) error {
	instanceStateService := NewInstanceStateService(meta.(*ve.SdkClient))
	if err := ve.DefaultDispatcher().Delete(instanceStateService, d, ResourceVolcengineEcsInstanceState()); err != nil {
		return fmt.Errorf("error on deleting instance state %q, %w", d.Id(), err)
	}
	return nil
}
