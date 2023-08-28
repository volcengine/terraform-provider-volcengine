package cloudfs_namespace

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
CloudfsNamespace can be imported using the FsName:NsId, e.g.
```
$ terraform import volcengine_cloudfs_namespace.default tfname:1801439850948****
```

*/

func ResourceVolcengineCloudfsNamespace() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineCloudfsNamespaceCreate,
		Read:   resourceVolcengineCloudfsNamespaceRead,
		Delete: resourceVolcengineCloudfsNamespaceDelete,
		Importer: &schema.ResourceImporter{
			State: func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
				items := strings.Split(data.Id(), ":")
				if len(items) != 2 {
					return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
				}
				if err := data.Set("fs_name", items[0]); err != nil {
					return []*schema.ResourceData{data}, err
				}
				if err := data.Set("ns_id", items[1]); err != nil {
					return []*schema.ResourceData{data}, err
				}
				return []*schema.ResourceData{data}, nil
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"fs_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of file system.",
			},
			"tos_bucket": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of tos bucket.",
			},
			"tos_prefix": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The tos prefix. Must not start with /, but must end with /, such as prefix/. When it is empty, it means the root path.",
			},
			"tos_account_id": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
				Description: "When a data lake scenario instance chooses to associate a bucket under another account, you need to set the ID of the account. " +
					"When importing resources, this attribute will not be imported. " +
					"If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
			},
			"tos_ak": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: "The tos ak. When the data lake scenario chooses to associate buckets under other accounts, need to set the Access Key ID of the account. " +
					"When importing resources, this attribute will not be imported. " +
					"If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
			},
			"tos_sk": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: "The tos sk. When the data lake scenario chooses to associate buckets under other accounts, need to set the Secret Access Key of the account. " +
					"When importing resources, this attribute will not be imported. " +
					"If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
			},
			"read_only": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				ForceNew:    true,
				Description: "Whether the namespace is read-only.",
			},

			"ns_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The id of namespace.",
			},
			"is_my_bucket": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the tos bucket is your own bucket.",
			},
			"service_managed": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the namespace is the official service for volcengine.",
			},
			"created_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The creation time of the namespace.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the namespace.",
			},
		},
	}
	return resource
}

func resourceVolcengineCloudfsNamespaceCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCloudfsNamespaceService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(service, d, ResourceVolcengineCloudfsNamespace())
	if err != nil {
		return fmt.Errorf("error on creating cloudfs namespace %q, %s", d.Id(), err)
	}
	return resourceVolcengineCloudfsNamespaceRead(d, meta)
}

func resourceVolcengineCloudfsNamespaceRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCloudfsNamespaceService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(service, d, ResourceVolcengineCloudfsNamespace())
	if err != nil {
		return fmt.Errorf("error on reading cloudfs namespace %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineCloudfsNamespaceUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCloudfsNamespaceService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(service, d, ResourceVolcengineCloudfsNamespace())
	if err != nil {
		return fmt.Errorf("error on updating cloudfs namespace %q, %s", d.Id(), err)
	}
	return resourceVolcengineCloudfsNamespaceRead(d, meta)
}

func resourceVolcengineCloudfsNamespaceDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCloudfsNamespaceService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineCloudfsNamespace())
	if err != nil {
		return fmt.Errorf("error on deleting cloudfs namespace %q, %s", d.Id(), err)
	}
	return err
}
