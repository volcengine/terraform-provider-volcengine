package nas_file_system

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
NasFileSystem can be imported using the id, e.g.
```
$ terraform import volcengine_nas_file_system.default enas-cnbjd3879745****
```

*/

func ResourceVolcengineNasFileSystem() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineNasFileSystemCreate,
		Read:   resourceVolcengineNasFileSystemRead,
		Update: resourceVolcengineNasFileSystemUpdate,
		Delete: resourceVolcengineNasFileSystemDelete,
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
				Description: "The name of the nas file system.",
			},
			"zone_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The zone id of the nas file system.",
			},
			"capacity": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The capacity of the nas file system. Unit: GiB.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The description of the nas file system.",
			},
			"snapshot_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The snapshot id when creating the nas file system. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The project name of the nas file system.",
			},
			"tags": ve.TagsSchema(),

			"file_system_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of the nas file system.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the nas file system.",
			},
			"charge_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The charge type of the nas file system.",
			},
			"protocol_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The protocol type of the nas file system.",
			},
			"storage_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The storage type of the nas file system.",
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The create time of the nas file system.",
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The update time of the nas file system.",
			},
			"region_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The region id of the nas file system.",
			},
			"zone_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The zone name of the nas file system.",
			},
			"version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The version of the nas file system.",
			},
			"snapshot_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The snapshot count of the nas file system.",
			},
		},
	}
	return resource
}

func resourceVolcengineNasFileSystemCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewNasFileSystemService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Create(service, d, ResourceVolcengineNasFileSystem())
	if err != nil {
		return fmt.Errorf("error on creating nas file system %q, %s", d.Id(), err)
	}
	return resourceVolcengineNasFileSystemRead(d, meta)
}

func resourceVolcengineNasFileSystemRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewNasFileSystemService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Read(service, d, ResourceVolcengineNasFileSystem())
	if err != nil {
		return fmt.Errorf("error on reading nas file system %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineNasFileSystemUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewNasFileSystemService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Update(service, d, ResourceVolcengineNasFileSystem())
	if err != nil {
		return fmt.Errorf("error on updating nas file system %q, %s", d.Id(), err)
	}
	return resourceVolcengineNasFileSystemRead(d, meta)
}

func resourceVolcengineNasFileSystemDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewNasFileSystemService(meta.(*ve.SdkClient))
	err = ve.DefaultDispatcher().Delete(service, d, ResourceVolcengineNasFileSystem())
	if err != nil {
		return fmt.Errorf("error on deleting nas file system %q, %s", d.Id(), err)
	}
	return err
}
