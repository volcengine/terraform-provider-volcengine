package cr_endpoint

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
CR endpoints can be imported using the endpoint:registryName, e.g.
```
$ terraform import volcengine_cr_endpoint.default endpoint:cr-basic
```

*/

func crEndpointImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(d.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{d}, fmt.Errorf("the format of import id must start with 'endpoint:',eg: 'endpoint:[registry-1]'")
	}
	if err := d.Set("registry", items[1]); err != nil {
		return []*schema.ResourceData{d}, err
	}
	return []*schema.ResourceData{d}, nil
}

func ResourceVolcengineCrEndpoint() *schema.Resource {
	resource := &schema.Resource{
		Read:   resourceVolcengineCrEndpointRead,
		Create: resourceVolcengineCrEndpointCreate,
		Update: resourceVolcengineCrEndpointUpdate,
		Delete: resourceVolcengineCrEndpointDelete,
		Importer: &schema.ResourceImporter{
			State: crEndpointImporter,
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
				ForceNew:    true,
				Description: "The CrRegistry name.",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether enable public endpoint.",
			},
		},
	}
	dataSource := DataSourceVolcengineCrEndpoints().Schema["endpoints"].Elem.(*schema.Resource).Schema
	ve.MergeDateSourceToResource(dataSource, &resource.Schema)
	return resource
}

func resourceVolcengineCrEndpointCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCrEndpointService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(service, d, ResourceVolcengineCrEndpoint())
	if err != nil {
		return fmt.Errorf("Error on creating CrEndpoint %q,%s", d.Id(), err)
	}
	return resourceVolcengineCrEndpointRead(d, meta)
}

func resourceVolcengineCrEndpointUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCrEndpointService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(service, d, ResourceVolcengineCrEndpoint())
	if err != nil {
		return fmt.Errorf("error on updating CrEndpoint  %q, %s", d.Id(), err)
	}
	return resourceVolcengineCrEndpointRead(d, meta)
}

func resourceVolcengineCrEndpointDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCrEndpointService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineCrEndpoint())
	if err != nil {
		return fmt.Errorf("error on deleting CrEndpoint %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineCrEndpointRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCrEndpointService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(service, d, ResourceVolcengineCrEndpoint())
	if err != nil {
		return fmt.Errorf("Error on reading CrEndpoint %q,%s", d.Id(), err)
	}
	return err
}
