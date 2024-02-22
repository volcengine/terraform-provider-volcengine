package cdn_domain_enabler

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
CdnDomainEnabler can be imported using the domain, e.g.
```
$ terraform import volcengine_cdn_domain_enabler.default enabler:www.volcengine.com
```

*/

func ResourceVolcengineCdnDomainEnabler() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineCdnDomainEnablerCreate,
		Read:   resourceVolcengineCdnDomainEnablerRead,
		Delete: resourceVolcengineCdnDomainEnablerDelete,
		Importer: &schema.ResourceImporter{
			State: enablerImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"domain": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Specify an acceleration domain name to enable.",
			},
		},
	}
	return resource
}

func resourceVolcengineCdnDomainEnablerCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCdnDomainEnablerService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineCdnDomainEnabler())
	if err != nil {
		return fmt.Errorf("error on creating cdn_domain_enabler %q, %s", d.Id(), err)
	}
	return resourceVolcengineCdnDomainEnablerRead(d, meta)
}

func resourceVolcengineCdnDomainEnablerRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCdnDomainEnablerService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineCdnDomainEnabler())
	if err != nil {
		return fmt.Errorf("error on reading cdn_domain_enabler %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineCdnDomainEnablerDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCdnDomainEnablerService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineCdnDomainEnabler())
	if err != nil {
		return fmt.Errorf("error on deleting cdn_domain_enabler %q, %s", d.Id(), err)
	}
	return err
}

func enablerImporter(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(data.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
	}
	if err := data.Set("domain", items[1]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}
