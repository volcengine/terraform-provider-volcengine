package traffic_mirror_target

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
TrafficMirrorTarget can be imported using the id, e.g.
```
$ terraform import volcengine_traffic_mirror_target.default resource_id
```

*/

func ResourceVolcengineTrafficMirrorTarget() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineTrafficMirrorTargetCreate,
		Read:   resourceVolcengineTrafficMirrorTargetRead,
		Update: resourceVolcengineTrafficMirrorTargetUpdate,
		Delete: resourceVolcengineTrafficMirrorTargetDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"instance_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The instance type of traffic mirror target. Valid values: `NetworkInterface`, `ClbInstance`.",
			},
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The instance id of traffic mirror target.",
			},
			"traffic_mirror_target_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of traffic mirror target.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of traffic mirror target.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The project name of traffic mirror target.",
			},
			"tags": ve.TagsSchema(),

			// computed fields
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of traffic mirror target.",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The create time of traffic mirror target.",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The update time of traffic mirror target.",
			},
		},
	}
	return resource
}

func resourceVolcengineTrafficMirrorTargetCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTrafficMirrorTargetService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineTrafficMirrorTarget())
	if err != nil {
		return fmt.Errorf("error on creating traffic_mirror_target %q, %s", d.Id(), err)
	}
	return resourceVolcengineTrafficMirrorTargetRead(d, meta)
}

func resourceVolcengineTrafficMirrorTargetRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTrafficMirrorTargetService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineTrafficMirrorTarget())
	if err != nil {
		return fmt.Errorf("error on reading traffic_mirror_target %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineTrafficMirrorTargetUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTrafficMirrorTargetService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineTrafficMirrorTarget())
	if err != nil {
		return fmt.Errorf("error on updating traffic_mirror_target %q, %s", d.Id(), err)
	}
	return resourceVolcengineTrafficMirrorTargetRead(d, meta)
}

func resourceVolcengineTrafficMirrorTargetDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTrafficMirrorTargetService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineTrafficMirrorTarget())
	if err != nil {
		return fmt.Errorf("error on deleting traffic_mirror_target %q, %s", d.Id(), err)
	}
	return err
}
