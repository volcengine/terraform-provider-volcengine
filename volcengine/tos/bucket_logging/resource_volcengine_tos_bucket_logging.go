package tos_bucket_logging

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
TosBucketLogging can be imported using the bucketName, e.g.
```
$ terraform import volcengine_tos_bucket_logging.default bucket_name
```

*/

func ResourceVolcengineTosBucketLogging() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineTosBucketLoggingCreate,
		Read:   resourceVolcengineTosBucketLoggingRead,
		Update: resourceVolcengineTosBucketLoggingUpdate,
		Delete: resourceVolcengineTosBucketLoggingDelete,
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
			"logging_enabled": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The name of the TOS bucket.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"target_bucket": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the target bucket where the access logs are stored.",
						},
						"target_prefix": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The prefix for the log object keys.",
						},
						"role": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The role that is assumed by TOS to write log objects to the target bucket.",
						},
					},
				},
			},
		},
	}
	return resource
}

func resourceVolcengineTosBucketLoggingCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTosBucketLoggingService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineTosBucketLogging())
	if err != nil {
		return fmt.Errorf("error on creating tos_bucket_logging %q, %s", d.Id(), err)
	}
	return resourceVolcengineTosBucketLoggingRead(d, meta)
}

func resourceVolcengineTosBucketLoggingRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTosBucketLoggingService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineTosBucketLogging())
	if err != nil {
		return fmt.Errorf("error on reading tos_bucket_logging %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineTosBucketLoggingUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTosBucketLoggingService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineTosBucketLogging())
	if err != nil {
		return fmt.Errorf("error on updating tos_bucket_logging %q, %s", d.Id(), err)
	}
	return resourceVolcengineTosBucketLoggingRead(d, meta)
}

func resourceVolcengineTosBucketLoggingDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTosBucketLoggingService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineTosBucketLogging())
	if err != nil {
		return fmt.Errorf("error on deleting tos_bucket_logging %q, %s", d.Id(), err)
	}
	return nil
}
