package tos_bucket_transfer_acceleration

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
TosBucketTransferAcceleration can be imported using the bucketName, e.g.
```
$ terraform import volcengine_tos_bucket_transfer_acceleration.default bucket_name
```

*/

func ResourceVolcengineTosBucketTransferAcceleration() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineTosBucketTransferAccelerationCreate,
		Read:   resourceVolcengineTosBucketTransferAccelerationRead,
		Delete: resourceVolcengineTosBucketTransferAccelerationDelete,
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

func resourceVolcengineTosBucketTransferAccelerationCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTosBucketTransferAccelerationService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineTosBucketTransferAcceleration())
	if err != nil {
		return fmt.Errorf("error on creating tos_bucket_transfer_acceleration %q, %s", d.Id(), err)
	}
	return resourceVolcengineTosBucketTransferAccelerationRead(d, meta)
}

func resourceVolcengineTosBucketTransferAccelerationRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTosBucketTransferAccelerationService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineTosBucketTransferAcceleration())
	if err != nil {
		return fmt.Errorf("error on reading tos_bucket_transfer_acceleration %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineTosBucketTransferAccelerationUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTosBucketTransferAccelerationService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineTosBucketTransferAcceleration())
	if err != nil {
		return fmt.Errorf("error on updating tos_bucket_transfer_acceleration %q, %s", d.Id(), err)
	}
	return resourceVolcengineTosBucketTransferAccelerationRead(d, meta)
}

func resourceVolcengineTosBucketTransferAccelerationDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTosBucketTransferAccelerationService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineTosBucketTransferAcceleration())
	if err != nil {
		return fmt.Errorf("error on deleting tos_bucket_transfer_acceleration %q, %s", d.Id(), err)
	}
	return nil
}
