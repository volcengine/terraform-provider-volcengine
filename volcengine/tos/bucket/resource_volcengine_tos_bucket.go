package bucket

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Tos Bucket can be imported using the id, e.g.
```
$ terraform import volcengine_tos_bucket.default region:bucketName
```

*/

func ResourceVolcengineTosBucket() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineTosBucketCreate,
		Read:   resourceVolcengineTosBucketRead,
		Update: resourceVolcengineTosBucketUpdate,
		Delete: resourceVolcengineTosBucketDelete,
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
					"ARCHIVE_FR",
				}, false),
				Default:     "STANDARD",
				Description: "The storage type of the object.Valid value is STANDARD|IA|ARCHIVE_FR.Default is STANDARD.",
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

func resourceVolcengineTosBucketCreate(d *schema.ResourceData, meta interface{}) (err error) {
	tosBucketService := NewTosBucketService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(tosBucketService, d, ResourceVolcengineTosBucket())
	if err != nil {
		return fmt.Errorf("error on creating tos bucket  %q, %s", d.Id(), err)
	}
	return resourceVolcengineTosBucketRead(d, meta)
}

func resourceVolcengineTosBucketRead(d *schema.ResourceData, meta interface{}) (err error) {
	tosBucketService := NewTosBucketService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(tosBucketService, d, ResourceVolcengineTosBucket())
	if err != nil {
		return fmt.Errorf("error on reading tos bucket %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineTosBucketUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	tosBucketService := NewTosBucketService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(tosBucketService, d, ResourceVolcengineTosBucket())
	if err != nil {
		return fmt.Errorf("error on updating tos bucket  %q, %s", d.Id(), err)
	}
	return resourceVolcengineTosBucketRead(d, meta)
}

func resourceVolcengineTosBucketDelete(d *schema.ResourceData, meta interface{}) (err error) {
	tosBucketService := NewTosBucketService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(tosBucketService, d, ResourceVolcengineTosBucket())
	if err != nil {
		return fmt.Errorf("error on deleting tos bucket %q, %s", d.Id(), err)
	}
	return err
}