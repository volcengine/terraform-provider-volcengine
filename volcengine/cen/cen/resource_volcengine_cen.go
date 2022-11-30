package cen

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Cen can be imported using the id, e.g.
```
$ terraform import volcengine_cen.default cen-7qthudw0ll6jmc****
```

*/

func ResourceVolcengineCen() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineCenCreate,
		Read:   resourceVolcengineCenRead,
		Update: resourceVolcengineCenUpdate,
		Delete: resourceVolcengineCenDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(1 * time.Hour),
			Update: schema.DefaultTimeout(1 * time.Hour),
			Delete: schema.DefaultTimeout(1 * time.Hour),
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
	s := DataSourceVolcengineCens().Schema["cens"].Elem.(*schema.Resource).Schema
	delete(s, "id")
	ve.MergeDateSourceToResource(s, &resource.Schema)
	return resource
}

func resourceVolcengineCenCreate(d *schema.ResourceData, meta interface{}) (err error) {
	cenService := NewCenService(meta.(*ve.SdkClient))
	err = cenService.Dispatcher.Create(cenService, d, ResourceVolcengineCen())
	if err != nil {
		return fmt.Errorf("error on creating cen  %q, %s", d.Id(), err)
	}
	return resourceVolcengineCenRead(d, meta)
}

func resourceVolcengineCenRead(d *schema.ResourceData, meta interface{}) (err error) {
	cenService := NewCenService(meta.(*ve.SdkClient))
	err = cenService.Dispatcher.Read(cenService, d, ResourceVolcengineCen())
	if err != nil {
		return fmt.Errorf("error on reading cen %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineCenUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	cenService := NewCenService(meta.(*ve.SdkClient))
	err = cenService.Dispatcher.Update(cenService, d, ResourceVolcengineCen())
	if err != nil {
		return fmt.Errorf("error on updating cen %q, %s", d.Id(), err)
	}
	return resourceVolcengineCenRead(d, meta)
}

func resourceVolcengineCenDelete(d *schema.ResourceData, meta interface{}) (err error) {
	cenService := NewCenService(meta.(*ve.SdkClient))
	err = cenService.Dispatcher.Delete(cenService, d, ResourceVolcengineCen())
	if err != nil {
		return fmt.Errorf("error on deleting cen %q, %s", d.Id(), err)
	}
	return err
}
