package bucket

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-vestack/common"
)

/*

Import
VPC can be imported using the id, e.g.
```
$ terraform import vestack_vpc.default vpc-mizl7m1kqccg5smt1bdpijuj
```

*/

func ResourceVestackTosBucket() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVestackTosBucketCreate,
		Read:   resourceVestackTosBucketRead,
		Update: resourceVestackTosBucketUpdate,
		Delete: resourceVestackTosBucketDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"bucket_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the bucket.",
			},
			"tos_acl": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"private",
					"public-read",
					"public-read-write",
					"authenticated-read",
					"bucket-owner-read",
				}, false),
				Default:     "private",
				Description: "The public acl control of bucket.",
			},
			"tos_storage_class": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"STANDARD",
					"IA",
				}, false),
				Default:     "STANDARD",
				Description: "The storage type of the bucket.",
			},
		},
	}
	return resource
}

func resourceVestackTosBucketCreate(d *schema.ResourceData, meta interface{}) (err error) {
	tosBucketService := NewTosBucketService(meta.(*ve.SdkClient))
	err = tosBucketService.Dispatcher.Create(tosBucketService, d, ResourceVestackTosBucket())
	if err != nil {
		return fmt.Errorf("error on creating tos bucket  %q, %s", d.Id(), err)
	}
	return resourceVestackTosBucketRead(d, meta)
}

func resourceVestackTosBucketRead(d *schema.ResourceData, meta interface{}) (err error) {
	tosBucketService := NewTosBucketService(meta.(*ve.SdkClient))
	err = tosBucketService.Dispatcher.Read(tosBucketService, d, ResourceVestackTosBucket())
	if err != nil {
		return fmt.Errorf("error on reading tos bucket %q, %s", d.Id(), err)
	}
	return err
}

func resourceVestackTosBucketUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	tosBucketService := NewTosBucketService(meta.(*ve.SdkClient))
	err = tosBucketService.Dispatcher.Update(tosBucketService, d, ResourceVestackTosBucket())
	if err != nil {
		return fmt.Errorf("error on updating tos bucket  %q, %s", d.Id(), err)
	}
	return resourceVestackTosBucketRead(d, meta)
}

func resourceVestackTosBucketDelete(d *schema.ResourceData, meta interface{}) (err error) {
	tosBucketService := NewTosBucketService(meta.(*ve.SdkClient))
	err = tosBucketService.Dispatcher.Delete(tosBucketService, d, ResourceVestackTosBucket())
	if err != nil {
		return fmt.Errorf("error on deleting tos bucket %q, %s", d.Id(), err)
	}
	return err
}
