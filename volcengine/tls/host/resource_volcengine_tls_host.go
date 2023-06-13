package host

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Tls Host can be imported using the host_group_id:ip, e.g.
```
$ terraform import volcengine_tls_host.default edf051ed-3c46-49:1.1.1.1
```

*/

func ResourceVolcengineTlsHost() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineTlsHostCreate,
		Read:   resourceVolcengineTlsHostRead,
		Delete: resourceVolcengineTlsHostDelete,
		Importer: &schema.ResourceImporter{
			State: func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
				items := strings.Split(data.Id(), ":")
				if len(items) != 2 {
					return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
				}
				if err := data.Set("host_group_id", items[0]); err != nil {
					return []*schema.ResourceData{data}, err
				}
				if err := data.Set("ip", items[1]); err != nil {
					return []*schema.ResourceData{data}, err
				}
				return []*schema.ResourceData{data}, nil
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"host_group_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of host group.",
			},
			"ip": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ip address.",
			},
		},
	}
	return resource
}

func resourceVolcengineTlsHostCreate(d *schema.ResourceData, meta interface{}) (err error) {
	return fmt.Errorf("only allows importing scenes, for removing unhealthy tls host instance")
}

func resourceVolcengineTlsHostRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(service, d, ResourceVolcengineTlsHost())
	if err != nil {
		return fmt.Errorf("error on reading tls host %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineTlsHostDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineTlsHost())
	if err != nil {
		return fmt.Errorf("error on deleting tls host %q, %s", d.Id(), err)
	}
	return err
}
