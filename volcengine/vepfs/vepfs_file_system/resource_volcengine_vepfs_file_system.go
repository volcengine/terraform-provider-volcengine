package vepfs_file_system

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
VepfsFileSystem can be imported using the id, e.g.
```
$ terraform import volcengine_vepfs_file_system.default resource_id
```

*/

func ResourceVolcengineVepfsFileSystem() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineVepfsFileSystemCreate,
		Read:   resourceVolcengineVepfsFileSystemRead,
		Update: resourceVolcengineVepfsFileSystemUpdate,
		Delete: resourceVolcengineVepfsFileSystemDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"file_system_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the vepfs file system.",
			},
			"subnet_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The subnet id of the vepfs file system.",
			},
			"store_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The store type of the vepfs file system. Valid values: `Advance_100`, `Performance` , `Intelligent_Computing`.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description info of the vepfs file system.",
			},
			"capacity": {
				Type:        schema.TypeInt,
				Required:    true, // 实测必须传递
				Description: "The capacity of the vepfs file system.",
			},
			"enable_restripe": {
				Type:     schema.TypeBool,
				Optional: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					// 创建时不存在这个参数，修改时存在这个参数
					return d.Id() == ""
				},
				Description: "Whether to enable data balance after capacity expansion. This filed is valid only when expanding capacity.",
			},
			"project": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The project of the vepfs file system.",
			},
			"tags": ve.TagsSchema(),

			// computed field
			"region_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The id of the region.",
			},
			"zone_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The id of the zone.",
			},
			"zone_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the zone.",
			},
			"bandwidth": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The bandwidth info of the vepfs file system.",
			},
			"version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The version info of the vepfs file system.",
			},
			"account_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The id of the account.",
			},
			"file_system_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of the vepfs file system.",
			},
			"store_type_cn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The store type cn name of the vepfs file system.",
			},
			"protocol_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The protocol type of the vepfs file system.",
			},
			"charge_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The charge type of the vepfs file system.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the vepfs file system.",
			},
			"charge_status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The charge status of the vepfs file system.",
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The create time of the vepfs file system.",
			},
			"last_modify_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The last modify time of the vepfs file system.",
			},
			"expire_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The expire time of the vepfs file system.",
			},
			"stop_service_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The stop service time of the vepfs file system.",
			},
			"free_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The free time of the vepfs file system.",
			},
		},
	}
	return resource
}

func resourceVolcengineVepfsFileSystemCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVepfsFileSystemService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineVepfsFileSystem())
	if err != nil {
		return fmt.Errorf("error on creating vepfs_file_system %q, %s", d.Id(), err)
	}
	return resourceVolcengineVepfsFileSystemRead(d, meta)
}

func resourceVolcengineVepfsFileSystemRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVepfsFileSystemService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineVepfsFileSystem())
	if err != nil {
		return fmt.Errorf("error on reading vepfs_file_system %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineVepfsFileSystemUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVepfsFileSystemService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineVepfsFileSystem())
	if err != nil {
		return fmt.Errorf("error on updating vepfs_file_system %q, %s", d.Id(), err)
	}
	return resourceVolcengineVepfsFileSystemRead(d, meta)
}

func resourceVolcengineVepfsFileSystemDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewVepfsFileSystemService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineVepfsFileSystem())
	if err != nil {
		return fmt.Errorf("error on deleting vepfs_file_system %q, %s", d.Id(), err)
	}
	return err
}
