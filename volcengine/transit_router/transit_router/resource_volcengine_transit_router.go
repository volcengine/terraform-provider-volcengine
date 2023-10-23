package transit_router

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
TransitRouter can be imported using the id, e.g.
```
$ terraform import volcengine_transit_router.default tr-2d6fr7mzya2gw58ozfes5g2oh
```

*/

func ResourceVolcengineTransitRouter() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineTransitRouterCreate,
		Read:   resourceVolcengineTransitRouterRead,
		Update: resourceVolcengineTransitRouterUpdate,
		Delete: resourceVolcengineTransitRouterDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"transit_router_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of the transit router.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the transit router.",
			},
		},
	}
	dataSource := DataSourceVolcengineTransitRouters().Schema["transit_routers"].Elem.(*schema.Resource).Schema
	delete(dataSource, "id")
	ve.MergeDateSourceToResource(dataSource, &resource.Schema)
	return resource
}

func resourceVolcengineTransitRouterCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineTransitRouter())
	if err != nil {
		return fmt.Errorf("error on creating transit router %q, %s", d.Id(), err)
	}
	return resourceVolcengineTransitRouterRead(d, meta)
}

func resourceVolcengineTransitRouterRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineTransitRouter())
	if err != nil {
		return fmt.Errorf("error on reading transit router %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineTransitRouterUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineTransitRouter())
	if err != nil {
		return fmt.Errorf("error on updating transit router %q, %s", d.Id(), err)
	}
	return resourceVolcengineTransitRouterRead(d, meta)
}

func resourceVolcengineTransitRouterDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineTransitRouter())
	if err != nil {
		return fmt.Errorf("error on deleting transit router %q, %s", d.Id(), err)
	}
	return err
}
