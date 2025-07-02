package traffic_mirror_session

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
TrafficMirrorSession can be imported using the id, e.g.
```
$ terraform import volcengine_traffic_mirror_session.default resource_id
```

*/

func ResourceVolcengineTrafficMirrorSession() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineTrafficMirrorSessionCreate,
		Read:   resourceVolcengineTrafficMirrorSessionRead,
		Update: resourceVolcengineTrafficMirrorSessionUpdate,
		Delete: resourceVolcengineTrafficMirrorSessionDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"traffic_mirror_session_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of the traffic mirror session.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the traffic mirror session.",
			},
			"network_interface_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of network interface.",
			},
			"traffic_mirror_target_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of traffic mirror target.",
			},
			"traffic_mirror_filter_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of traffic mirror filter.",
			},
			"virtual_network_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "The ID of virtual network.",
			},
			"packet_length": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "The packet length of traffic mirror session.",
			},
			"priority": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The priority of traffic mirror session. Valid values: 1~32766.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The project name of traffic mirror session.",
			},
			"tags": ve.TagsSchema(),

			// computed fields
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of traffic mirror session.",
			},
			"business_status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The business status of traffic mirror session.",
			},
			"lock_reason": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The lock reason of traffic mirror session.",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The create time of traffic mirror session.",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The update time of traffic mirror session.",
			},
		},
	}
	return resource
}

func resourceVolcengineTrafficMirrorSessionCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTrafficMirrorSessionService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineTrafficMirrorSession())
	if err != nil {
		return fmt.Errorf("error on creating traffic_mirror_session %q, %s", d.Id(), err)
	}
	return resourceVolcengineTrafficMirrorSessionRead(d, meta)
}

func resourceVolcengineTrafficMirrorSessionRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTrafficMirrorSessionService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineTrafficMirrorSession())
	if err != nil {
		return fmt.Errorf("error on reading traffic_mirror_session %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineTrafficMirrorSessionUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTrafficMirrorSessionService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineTrafficMirrorSession())
	if err != nil {
		return fmt.Errorf("error on updating traffic_mirror_session %q, %s", d.Id(), err)
	}
	return resourceVolcengineTrafficMirrorSessionRead(d, meta)
}

func resourceVolcengineTrafficMirrorSessionDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTrafficMirrorSessionService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineTrafficMirrorSession())
	if err != nil {
		return fmt.Errorf("error on deleting traffic_mirror_session %q, %s", d.Id(), err)
	}
	return err
}
