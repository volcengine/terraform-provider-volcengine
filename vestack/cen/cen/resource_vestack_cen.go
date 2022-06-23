package cen

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-vestack/common"
)

/*

Import
Cen can be imported using the id, e.g.
```
$ terraform import vestack_cen.default cen-7qthudw0ll6jmc****
```

*/

func ResourceVestackCen() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVestackCenCreate,
		Read:   resourceVestackCenRead,
		Update: resourceVestackCenUpdate,
		Delete: resourceVestackCenDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"cen_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of the cen.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The description of the cen.",
			},
		},
	}
	s := DataSourceVestackCens().Schema["cens"].Elem.(*schema.Resource).Schema
	delete(s, "id")
	ve.MergeDateSourceToResource(s, &resource.Schema)
	return resource
}

func resourceVestackCenCreate(d *schema.ResourceData, meta interface{}) (err error) {
	cenService := NewCenService(meta.(*ve.SdkClient))
	err = cenService.Dispatcher.Create(cenService, d, ResourceVestackCen())
	if err != nil {
		return fmt.Errorf("error on creating cen  %q, %s", d.Id(), err)
	}
	return resourceVestackCenRead(d, meta)
}

func resourceVestackCenRead(d *schema.ResourceData, meta interface{}) (err error) {
	cenService := NewCenService(meta.(*ve.SdkClient))
	err = cenService.Dispatcher.Read(cenService, d, ResourceVestackCen())
	if err != nil {
		return fmt.Errorf("error on reading cen %q, %s", d.Id(), err)
	}
	return err
}

func resourceVestackCenUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	cenService := NewCenService(meta.(*ve.SdkClient))
	err = cenService.Dispatcher.Update(cenService, d, ResourceVestackCen())
	if err != nil {
		return fmt.Errorf("error on updating cen %q, %s", d.Id(), err)
	}
	return resourceVestackCenRead(d, meta)
}

func resourceVestackCenDelete(d *schema.ResourceData, meta interface{}) (err error) {
	cenService := NewCenService(meta.(*ve.SdkClient))
	err = cenService.Dispatcher.Delete(cenService, d, ResourceVestackCen())
	if err != nil {
		return fmt.Errorf("error on deleting cen %q, %s", d.Id(), err)
	}
	return err
}
