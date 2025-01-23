package object

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
TOS Object can be imported using the id, e.g.
```
$ terraform import volcengine_tos_object.default bucketName:objectName
```

*/

func ResourceVolcengineTosObject() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineTosObjectCreate,
		Read:   resourceVolcengineTosObjectRead,
		Update: resourceVolcengineTosObjectUpdate,
		Delete: resourceVolcengineTosObjectDelete,
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
		CustomizeDiff: func(diff *schema.ResourceDiff, i interface{}) (err error) {
			if diff.Id() != "" && diff.HasChange("file_path") && !diff.Get("enable_version").(bool) {
				return diff.ForceNew("file_path")
			}
			return err
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
				Type:     schema.TypeString,
				Optional: true,
				//ForceNew:    true,
				Description:  "The file path for upload. Only one of `file_path,content` can be specified.",
				ExactlyOneOf: []string{"file_path", "content"},
			},
			"content_md5": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The file md5 sum (32-bit hexadecimal string) for upload.",
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
					"default",
				}, false),
				Default:     "private",
				Description: "The public acl control of object. Valid value is private|public-read|public-read-write|authenticated-read|bucket-owner-read|default. `default` means to enable the default inheritance bucket ACL function for the object.",
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
			"version_ids": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set:         schema.HashString,
				Description: "The version ids of the object if exist.",
			},
			"is_default": {
				Type: schema.TypeBool,
				//Optional:    true,
				Computed:    true,
				Description: "Whether to enable the default inheritance bucket ACL function for the object.",
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
			"enable_version": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "The flag of enable tos version.",
			},
			"content": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				ExactlyOneOf: []string{"file_path", "content"},
				Description:  "The content of the TOS Object when content type is json or text and xml. Only one of `file_path,content` can be specified.",
			},
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Tos Bucket Tags.",
				Set:         TagsHash,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The Key of Tags.",
						},
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The Value of Tags.",
						},
					},
				},
			},
		},
	}
	return resource
}

func resourceVolcengineTosObjectCreate(d *schema.ResourceData, meta interface{}) (err error) {
	tosBucketService := NewTosObjectService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(tosBucketService, d, ResourceVolcengineTosObject())
	if err != nil {
		return fmt.Errorf("error on creating tos object %q, %s", d.Id(), err)
	}
	return resourceVolcengineTosObjectRead(d, meta)
}

func resourceVolcengineTosObjectRead(d *schema.ResourceData, meta interface{}) (err error) {
	tosBucketService := NewTosObjectService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(tosBucketService, d, ResourceVolcengineTosObject())
	if err != nil {
		return fmt.Errorf("error on reading tos object %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineTosObjectUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	tosBucketService := NewTosObjectService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(tosBucketService, d, ResourceVolcengineTosObject())
	if err != nil {
		return fmt.Errorf("error on updating tos object  %q, %s", d.Id(), err)
	}
	return resourceVolcengineTosObjectRead(d, meta)
}

func resourceVolcengineTosObjectDelete(d *schema.ResourceData, meta interface{}) (err error) {
	tosBucketService := NewTosObjectService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(tosBucketService, d, ResourceVolcengineTosObject())
	if err != nil {
		return fmt.Errorf("error on deleting tos object %q, %s", d.Id(), err)
	}
	return err
}

var TagsHash = func(v interface{}) int {
	if v == nil {
		return hashcode.String("")
	}
	m := v.(map[string]interface{})
	var (
		buf bytes.Buffer
	)
	buf.WriteString(fmt.Sprintf("%v#%v", m["key"], m["value"]))
	return hashcode.String(buf.String())
}
