package host

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
The TlsHost is not support import.

*/

func ResourceVolcengineTlsHost() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineTlsHostCreate,
		Read:   resourceVolcengineTlsHostRead,
		Delete: resourceVolcengineTlsHostDelete,
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
				Optional:    true,
				ForceNew:    true,
				Description: "The ip address.",
			},
		},
	}
	return resource
}

func resourceVolcengineTlsHostCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(service, d, ResourceVolcengineTlsHost())
	if err != nil {
		return fmt.Errorf("error on creating tls host, %s", err)
	}
	return nil
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
