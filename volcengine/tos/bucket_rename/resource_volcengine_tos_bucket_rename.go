package tos_bucket_rename

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
TosBucketRename can be imported using the bucketName, e.g.
```
$ terraform import volcengine_tos_bucket_rename.default bucket_name
```

*/

func ResourceVolcengineTosBucketRename() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineTosBucketRenameCreate,
		Read:   resourceVolcengineTosBucketRenameRead,
		Delete: resourceVolcengineTosBucketRenameDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				bucketName := d.Id()
				if err := d.Set("bucket_name", bucketName); err != nil {
					return nil, err
				}
				d.SetId(bucketName)
				return []*schema.ResourceData{d}, nil
			},
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
				Description: "The name of the TOS bucket to configure rename functionality for.",
			},
		},
	}
	return resource
}

func resourceVolcengineTosBucketRenameCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTosBucketRenameService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineTosBucketRename())
	if err != nil {
		return fmt.Errorf("error on creating tos_bucket_rename %q, %s", d.Id(), err)
	}
	return resourceVolcengineTosBucketRenameRead(d, meta)
}

func resourceVolcengineTosBucketRenameRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTosBucketRenameService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineTosBucketRename())
	if err != nil {
		return fmt.Errorf("error on reading tos_bucket_rename %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineTosBucketRenameUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTosBucketRenameService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineTosBucketRename())
	if err != nil {
		return fmt.Errorf("error on updating tos_bucket_rename %q, %s", d.Id(), err)
	}
	return resourceVolcengineTosBucketRenameRead(d, meta)
}

func resourceVolcengineTosBucketRenameDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewTosBucketRenameService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineTosBucketRename())
	if err != nil {
		return fmt.Errorf("error on deleting tos_bucket_rename %q, %s", d.Id(), err)
	}
	return nil
}
