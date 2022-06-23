package bucket

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-vestack/common"
)

/*

Import
Tos Bucket can be imported using the id, e.g.
```
$ terraform import vestack_tos_bucket.default region:bucketName
```

*/

func ResourceVestackTosBucket() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVestackTosBucketCreate,
		Read:   resourceVestackTosBucketRead,
		Update: resourceVestackTosBucketUpdate,
		Delete: resourceVestackTosBucketDelete,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(2 * time.Minute),
			Update: schema.DefaultTimeout(2 * time.Minute),
			Delete: schema.DefaultTimeout(2 * time.Minute),
		},
		Importer: &schema.ResourceImporter{
			State: func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
				items := strings.Split(data.Id(), ":")
				if len(items) != 1 {
					return []*schema.ResourceData{data}, fmt.Errorf("import id must be of the form bucketName")
				}
				_ = data.Set("bucket_name", items[0])
				return []*schema.ResourceData{data}, nil
			},
		},
		Schema: map[string]*schema.Schema{
			"bucket_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the bucket.",
			},
			"public_acl": {
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
				Description: "The public acl control of object.Valid value is private|public-read|public-read-write|authenticated-read|bucket-owner-read.",
			},
			"storage_class": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"STANDARD",
					"IA",
				}, false),
				Default:     "STANDARD",
				Description: "The storage type of the object.Valid value is STANDARD|IA.",
			},
			"enable_version": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "The flag of enable tos version.",
			},

			"account_acl": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "The user set of grant full control.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The accountId to control.",
						},
						"acl_type": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "CanonicalUser",
							ValidateFunc: validation.StringInSlice([]string{
								"CanonicalUser",
							}, false),
							Description: "The acl type to control.Valid value is CanonicalUser.",
						},
						"permission": {
							Type:     schema.TypeString,
							Required: true,
							ValidateFunc: validation.StringInSlice([]string{
								"FULL_CONTROL",
								"READ",
								"READ_ACP",
								"WRITE",
								"WRITE_ACP",
							}, false),
							Description: "The permission to control.Valid value is FULL_CONTROL|READ|READ_ACP|WRITE|WRITE_ACP.",
						},
					},
				},
				Set: ve.TosAccountAclHash,
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
