package cr_namespace

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
CR namespace can be imported using the registry:name, e.g.
```
$ terraform import volcengine_cr_namespace.default cr-basic:namespace-1
```

*/

func crNamespaceImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(d.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{d}, fmt.Errorf("the format of import id must be 'registry:namespace'")
	}
	if err := d.Set("registry", items[0]); err != nil {
		return []*schema.ResourceData{d}, err
	}
	if err := d.Set("name", items[1]); err != nil {
		return []*schema.ResourceData{d}, err
	}
	return []*schema.ResourceData{d}, nil
}

func ResourceVolcengineCrNamespace() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineCrNamespaceCreate,
		Read:   resourceVolcengineCrNamespaceRead,
		Update: resourceVolcengineCrNamespaceUpdate,
		Delete: resourceVolcengineCrNamespaceDelete,
		Importer: &schema.ResourceImporter{
			State: crNamespaceImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"registry": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The registry name.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of CrNamespace.",
			},
		},
	}
	dataSource := DataSourceVolcengineCrNamespaces().Schema["namespaces"].Elem.(*schema.Resource).Schema
	ve.MergeDateSourceToResource(dataSource, &resource.Schema)
	return resource
}

func resourceVolcengineCrNamespaceCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCrNamespaceService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineCrNamespace())
	if err != nil {
		return fmt.Errorf("error on creating CrNamespace %q,%s", d.Id(), err)
	}
	return resourceVolcengineCrNamespaceRead(d, meta)
}

func resourceVolcengineCrNamespaceUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCrNamespaceService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineCrNamespace())
	if err != nil {
		return fmt.Errorf("error on updating CrNamespace  %q, %s", d.Id(), err)
	}
	return resourceVolcengineCrNamespaceRead(d, meta)
}

func resourceVolcengineCrNamespaceDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCrNamespaceService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineCrNamespace())
	if err != nil {
		return fmt.Errorf("error on deleting CrNamespace %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineCrNamespaceRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCrNamespaceService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineCrNamespace())
	if err != nil {
		return fmt.Errorf("error on reading CrNamespace %q,%s", d.Id(), err)
	}
	return err
}
