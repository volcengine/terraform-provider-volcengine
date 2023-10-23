package shared_transit_router_state

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
SharedTransitRouterState can be imported using the id, e.g.
```
$ terraform import volcengine_shared_transit_router_state.default state:transitRouterId
```

*/

func ResourceVolcengineSharedTransitRouterState() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineSharedTransitRouterStateCreate,
		Read:   resourceVolcengineSharedTransitRouterStateRead,
		Update: resourceVolcengineSharedTransitRouterStateUpdate,
		Delete: resourceVolcengineSharedTransitRouterStateDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"transit_router_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the transit router.",
			},
			"action": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"Accept",
					"Reject",
				}, false),
				Description: "Accept or reject the shared transit router.",
			},
		},
	}
	return resource
}

func resourceVolcengineSharedTransitRouterStateCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewSharedTransitRouterStateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineSharedTransitRouterState())
	if err != nil {
		return fmt.Errorf("error on creating shared_transit_router_state %q, %s", d.Id(), err)
	}
	return resourceVolcengineSharedTransitRouterStateRead(d, meta)
}

func resourceVolcengineSharedTransitRouterStateRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewSharedTransitRouterStateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineSharedTransitRouterState())
	if err != nil {
		return fmt.Errorf("error on reading shared_transit_router_state %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineSharedTransitRouterStateUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewSharedTransitRouterStateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineSharedTransitRouterState())
	if err != nil {
		return fmt.Errorf("error on update shared_transit_router_state %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineSharedTransitRouterStateDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewSharedTransitRouterStateService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineSharedTransitRouterState())
	if err != nil {
		return fmt.Errorf("error on deleting shared_transit_router_state %q, %s", d.Id(), err)
	}
	return err
}

var stateImporter = func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(data.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
	}
	if err := data.Set("transit_router_id", items[1]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}
