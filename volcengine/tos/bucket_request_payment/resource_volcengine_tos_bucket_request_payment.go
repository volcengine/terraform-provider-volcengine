package tos_bucket_request_payment

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
TosBucketRequestPayment can be imported using the bucketName, e.g.
```
$ terraform import volcengine_tos_bucket_request_payment.default bucket_name
```

*/

func ResourceVolcengineTosBucketRequestPayment() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineTosBucketRequestPaymentCreate,
		Read:   resourceVolcengineTosBucketRequestPaymentRead,
		Delete: resourceVolcengineTosBucketRequestPaymentDelete,
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

func resourceVolcengineTosBucketRequestPaymentCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTosBucketRequestPaymentService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineTosBucketRequestPayment())
	if err != nil {
		return fmt.Errorf("error on creating tos_bucket_request_payment %q, %s", d.Id(), err)
	}
	return resourceVolcengineTosBucketRequestPaymentRead(d, meta)
}

func resourceVolcengineTosBucketRequestPaymentRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTosBucketRequestPaymentService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineTosBucketRequestPayment())
	if err != nil {
		return fmt.Errorf("error on reading tos_bucket_request_payment %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineTosBucketRequestPaymentDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTosBucketRequestPaymentService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineTosBucketRequestPayment())
	if err != nil {
		return fmt.Errorf("error on deleting tos_bucket_request_payment %q, %s", d.Id(), err)
	}
	return nil
}
