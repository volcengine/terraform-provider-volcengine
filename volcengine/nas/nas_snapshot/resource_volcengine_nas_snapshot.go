package nas_snapshot

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	volc "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Nas Snapshot can be imported using the id, e.g.
```
$ terraform import volcengine_nas_snapshot.default snap-472a716f****
```

*/

func ResourceVolcengineNasSnapshot() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineNasSnapshotCreate,
		Read:   resourceVolcengineNasSnapshotRead,
		Delete: resourceVolcengineNasSnapshotDelete,
		Update: resourceVolcengineNasSnapshotUpdate,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"file_system_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The ID of the file system.",
			},
			"snapshot_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of snapshot.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of snapshot.",
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The create time of snapshot.",
			},
			"progress": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The progress of snapshot.",
			},
			"is_encrypt": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether is encrypt.",
			},
			"file_system_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of file system.",
			},
			"zone_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of zone.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of snapshot.",
			},
			"source_version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The source version info.",
			},
			"source_size": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The size of source.",
			},
			"snapshot_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of snapshot.",
			},
		},
	}
}

func resourceVolcengineNasSnapshotCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*volc.SdkClient))
	err = volc.DefaultDispatcher().Create(service, d, ResourceVolcengineNasSnapshot())
	if err != nil {
		return fmt.Errorf("error on creating nas snapshot %q, %w", d.Id(), err)
	}
	return resourceVolcengineNasSnapshotRead(d, meta)
}

func resourceVolcengineNasSnapshotRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*volc.SdkClient))
	err = volc.DefaultDispatcher().Read(service, d, ResourceVolcengineNasSnapshot())
	if err != nil {
		return fmt.Errorf("error on reading nas snapshot %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineNasSnapshotDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*volc.SdkClient))
	err = volc.DefaultDispatcher().Delete(service, d, ResourceVolcengineNasSnapshot())
	if err != nil {
		return fmt.Errorf("error on deleting nas snapshot %q, %w", d.Id(), err)
	}
	return err
}

func resourceVolcengineNasSnapshotUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewService(meta.(*volc.SdkClient))
	err = volc.DefaultDispatcher().Update(service, d, ResourceVolcengineNasSnapshot())
	if err != nil {
		return fmt.Errorf("error on creating nas snapshot %q, %w", d.Id(), err)
	}
	return resourceVolcengineNasSnapshotRead(d, meta)
}
