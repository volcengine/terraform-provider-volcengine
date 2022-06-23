package object

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
TOS Object can be imported using the id, e.g.
```
$ terraform import vestack_tos_object.default bucketName:objectName
```

*/

func ResourceVestackTosObject() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVestackTosObjectCreate,
		Read:   resourceVestackTosObjectRead,
		Update: resourceVestackTosObjectUpdate,
		Delete: resourceVestackTosObjectDelete,
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(2 * time.Minute),
			Update: schema.DefaultTimeout(2 * time.Minute),
			Delete: schema.DefaultTimeout(2 * time.Minute),
		},
		Importer: &schema.ResourceImporter{
			State: func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
				items := strings.Split(data.Id(), ":")
				if len(items) != 2 {
					return []*schema.ResourceData{data}, fmt.Errorf("import id must be of the form bucketName:objectName")
				}
				_ = data.Set("bucket_name", items[0])
				_ = data.Set("object_name", items[1])
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
			"object_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the object.",
			},
			"file_path": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The file path for upload.",
			},
			"encryption": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				ValidateFunc: validation.StringInSlice([]string{
					"AES256",
				}, false),
				Description: "The encryption of the object.Valid value is AES256.",
			},
			"content_type": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "The content type of the object.",
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

func resourceVestackTosObjectCreate(d *schema.ResourceData, meta interface{}) (err error) {
	tosBucketService := NewTosObjectService(meta.(*ve.SdkClient))
	err = tosBucketService.Dispatcher.Create(tosBucketService, d, ResourceVestackTosObject())
	if err != nil {
		return fmt.Errorf("error on creating tos bucket  %q, %s", d.Id(), err)
	}
	return resourceVestackTosObjectRead(d, meta)
}

func resourceVestackTosObjectRead(d *schema.ResourceData, meta interface{}) (err error) {
	tosBucketService := NewTosObjectService(meta.(*ve.SdkClient))
	err = tosBucketService.Dispatcher.Read(tosBucketService, d, ResourceVestackTosObject())
	if err != nil {
		return fmt.Errorf("error on reading tos bucket %q, %s", d.Id(), err)
	}
	return err
}

func resourceVestackTosObjectUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	tosBucketService := NewTosObjectService(meta.(*ve.SdkClient))
	err = tosBucketService.Dispatcher.Update(tosBucketService, d, ResourceVestackTosObject())
	if err != nil {
		return fmt.Errorf("error on updating tos bucket  %q, %s", d.Id(), err)
	}
	return resourceVestackTosObjectRead(d, meta)
}

func resourceVestackTosObjectDelete(d *schema.ResourceData, meta interface{}) (err error) {
	tosBucketService := NewTosObjectService(meta.(*ve.SdkClient))
	err = tosBucketService.Dispatcher.Delete(tosBucketService, d, ResourceVestackTosObject())
	if err != nil {
		return fmt.Errorf("error on deleting tos bucket %q, %s", d.Id(), err)
	}
	return err
}
