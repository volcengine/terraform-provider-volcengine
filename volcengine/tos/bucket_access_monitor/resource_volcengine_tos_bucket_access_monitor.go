package tos_bucket_access_monitor

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
TosBucketAccessMonitor can be imported using the bucketName, e.g.
```
$ terraform import volcengine_tos_bucket_access_monitor.default bucket_name
```

*/

func ResourceVolcengineTosBucketAccessMonitor() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineTosBucketAccessMonitorCreate,
		Read:   resourceVolcengineTosBucketAccessMonitorRead,
		Delete: resourceVolcengineTosBucketAccessMonitorDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"bucket_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the TOS bucket.",
			},
		},
	}
	return resource
}

func resourceVolcengineTosBucketAccessMonitorCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTosBucketAccessMonitorService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineTosBucketAccessMonitor())
	if err != nil {
		return fmt.Errorf("error on creating tos_bucket_access_monitor %q, %s", d.Id(), err)
	}
	return resourceVolcengineTosBucketAccessMonitorRead(d, meta)
}

func resourceVolcengineTosBucketAccessMonitorRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTosBucketAccessMonitorService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineTosBucketAccessMonitor())
	if err != nil {
		return fmt.Errorf("error on reading tos_bucket_access_monitor %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineTosBucketAccessMonitorUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTosBucketAccessMonitorService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineTosBucketAccessMonitor())
	if err != nil {
		return fmt.Errorf("error on updating tos_bucket_access_monitor %q, %s", d.Id(), err)
	}
	return resourceVolcengineTosBucketAccessMonitorRead(d, meta)
}

func resourceVolcengineTosBucketAccessMonitorDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTosBucketAccessMonitorService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineTosBucketAccessMonitor())
	if err != nil {
		return fmt.Errorf("error on deleting tos_bucket_access_monitor %q, %s", d.Id(), err)
	}
	return nil
}
