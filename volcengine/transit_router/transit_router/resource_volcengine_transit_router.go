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
			"asn": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "The asn of the transit router. Valid value range in 64512-65534 and 4200000000-4294967294. Default is 64512.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The ProjectName of the transit router.",
			},
			"tags": ve.TagsSchema(),
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
