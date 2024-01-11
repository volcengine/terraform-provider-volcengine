package cloud_monitor_contact

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
CloudMonitor Contact can be imported using the id, e.g.
```
$ terraform import volcengine_cloud_monitor_contact.default 145258255725730****
```

*/

func ResourceVolcengineCloudMonitorContact() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineCloudMonitorContactCreate,
		Read:   resourceVolcengineCloudMonitorContactRead,
		Update: resourceVolcengineCloudMonitorContactUpdate,
		Delete: resourceVolcengineCloudMonitorContactDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of contact.",
			},
			"email": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The email of contact.",
			},
			"phone": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The phone of contact.",
			},
		},
	}
}

func resourceVolcengineCloudMonitorContactCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(service, d, ResourceVolcengineCloudMonitorContact())
	if err != nil {
		return fmt.Errorf("error on creating Contact %q, %w", d.Id(), err)
	}
	return resourceVolcengineCloudMonitorContactRead(d, meta)
}

func resourceVolcengineCloudMonitorContactRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(service, d, ResourceVolcengineCloudMonitorContact())
	if err != nil {
		return fmt.Errorf("error on reading Contact %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineCloudMonitorContactUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(service, d, ResourceVolcengineCloudMonitorContact())
	if err != nil {
		return fmt.Errorf("error on updating Contact %q, %w", d.Id(), err)
	}
	return resourceVolcengineCloudMonitorContactRead(d, meta)
}

func resourceVolcengineCloudMonitorContactDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineCloudMonitorContact())
	if err != nil {
		return fmt.Errorf("error on deleting Contact %q, %w", d.Id(), err)
	}
	return err
}
