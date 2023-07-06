package cloudfs_file_system

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
CloudFileSystem can be imported using the FsName, e.g.
```
$ terraform import volcengine_cloudfs_file_system.default tfname
```

*/

func ResourceVolcengineCloudfsFileSystem() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineCloudfsFileSystemCreate,
		Read:   resourceVolcengineCloudfsFileSystemRead,
		Update: resourceVolcengineCloudfsFileSystemUpdate,
		Delete: resourceVolcengineCloudfsFileSystemDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"fs_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of file system.",
			},
			"zone_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of zone.",
			},
			"cache_plan": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"T2", "T4", "DISABLED"}, false),
				Description: "The cache plan. The value can be `DISABLED` or `T2` or `T4`. " +
					"When expanding the cache size, the cache plan should remain the same. For data lakes, cache must be enabled.",
			},
			"cache_capacity_tib": {
				Type:             schema.TypeInt,
				Optional:         true,
				DiffSuppressFunc: diffCache,
				Description:      "The capacity of cache. This parameter is required when cache acceleration is enabled.",
			},
			"subnet_id": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: diffCache,
				Description:      "The id of subnet. This parameter is required when cache acceleration is enabled.",
			},
			"security_group_id": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: diffCache,
				Description:      "The id of security group. This parameter is required when cache acceleration is enabled.",
			},
			"vpc_route_enabled": {
				Type:             schema.TypeBool,
				Optional:         true,
				DiffSuppressFunc: diffCache,
				Description:      "Whether enable all vpc route.",
			},
			"mode": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"HDFS_MODE", "ACC_MODE"}, false),
				Description:  "The mode of file system. The value can be `HDFS_MODE` or `ACC_MODE`.",
			},
			"tos_bucket": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The tos bucket. When importing ACC_MODE resources, this attribute will not be imported.",
			},
			"tos_prefix": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: "The tos prefix. Must not start with /, but must end with /, such as prefix/. When it is empty, it means the root path. " +
					"When importing ACC_MODE resources, this attribute will not be imported.",
			},
			"tos_account_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: "When a data lake scenario instance chooses to associate a bucket under another account, you need to set the ID of the account. " +
					"When importing resources, this attribute will not be imported.",
			},
			"tos_ak": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: "The tos ak. When the data lake scenario chooses to associate buckets under other accounts, need to set the Access Key ID of the account. " +
					"When importing resources, this attribute will not be imported.",
			},
			"tos_sk": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: "The tos sk. When the data lake scenario chooses to associate buckets under other accounts, need to set the Secret Access Key of the account. " +
					"When importing resources, this attribute will not be imported.",
			},
			"read_only": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: "Whether the Namespace created automatically when mounting the TOS Bucket is read-only. " +
					"When importing resources, this attribute will not be imported. " +
					"If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
			},

			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Status of file system.",
			},
			"created_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The creation time.",
			},
			"mount_point": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The point mount.",
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The id of vpc.",
			},
			"access_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The default vpc access id.",
			},
		},
	}
	return resource
}

func resourceVolcengineCloudfsFileSystemCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(service, d, ResourceVolcengineCloudfsFileSystem())
	if err != nil {
		return fmt.Errorf("error on creating file system %q, %s", d.Id(), err)
	}
	return resourceVolcengineCloudfsFileSystemRead(d, meta)
}

func resourceVolcengineCloudfsFileSystemRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(service, d, ResourceVolcengineCloudfsFileSystem())
	if err != nil {
		return fmt.Errorf("error on reading file system %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineCloudfsFileSystemUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(service, d, ResourceVolcengineCloudfsFileSystem())
	if err != nil {
		return fmt.Errorf("error on updating file system %q, %s", d.Id(), err)
	}
	return resourceVolcengineCloudfsFileSystemRead(d, meta)
}

func resourceVolcengineCloudfsFileSystemDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineCloudfsFileSystem())
	if err != nil {
		return fmt.Errorf("error on deleting file system %q, %s", d.Id(), err)
	}
	return err
}

func diffCache(k, old, new string, d *schema.ResourceData) bool {
	// 禁用缓存时，不起作用
	if d.Get("cache_plan").(string) == "DISABLED" {
		return true
	}
	// cache_plan 没有发生变化，只是扩容，忽略变更
	if d.Id() != "" && !d.HasChange("cache_plan") {
		if k == "subnet_id" || k == "security_group_id" {
			return true
		}
	}
	return false
}
