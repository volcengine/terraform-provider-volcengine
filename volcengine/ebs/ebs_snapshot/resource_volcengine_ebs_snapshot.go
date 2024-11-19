package ebs_snapshot

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
EbsSnapshot can be imported using the id, e.g.
```
$ terraform import volcengine_ebs_snapshot.default resource_id
```

*/

func ResourceVolcengineEbsSnapshot() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineEbsSnapshotCreate,
		Read:   resourceVolcengineEbsSnapshotRead,
		Update: resourceVolcengineEbsSnapshotUpdate,
		Delete: resourceVolcengineEbsSnapshotDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"volume_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The volume id to create snapshot.",
			},
			"snapshot_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the snapshot.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The description of the snapshot.",
			},
			"retention_days": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				Description: "The retention days of the snapshot. Valid values: 1~65536. Not specifying this field means permanently preserving the snapshot." +
					"When modifying this field, the retention days only supports extension and not shortening. The value range is N+1~65536, where N is the retention days set during snapshot creation.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The project name of the snapshot.",
			},
			"tags": ve.TagsSchema(),

			// computed fields
			"snapshot_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of the snapshot.",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the snapshot.",
			},
			"volume_kind": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The volume kind of the snapshot.",
			},
			"volume_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The volume name of the snapshot.",
			},
			"volume_size": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The volume size of the snapshot.",
			},
			"volume_status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The volume status of the snapshot.",
			},
			"volume_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The volume type of the snapshot.",
			},
			"zone_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The zone id of the snapshot.",
			},
			"creation_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The creation time of the snapshot.",
			},
		},
	}
	return resource
}

func resourceVolcengineEbsSnapshotCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEbsSnapshotService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineEbsSnapshot())
	if err != nil {
		return fmt.Errorf("error on creating ebs_snapshot %q, %s", d.Id(), err)
	}
	return resourceVolcengineEbsSnapshotRead(d, meta)
}

func resourceVolcengineEbsSnapshotRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEbsSnapshotService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineEbsSnapshot())
	if err != nil {
		return fmt.Errorf("error on reading ebs_snapshot %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineEbsSnapshotUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEbsSnapshotService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineEbsSnapshot())
	if err != nil {
		return fmt.Errorf("error on updating ebs_snapshot %q, %s", d.Id(), err)
	}
	return resourceVolcengineEbsSnapshotRead(d, meta)
}

func resourceVolcengineEbsSnapshotDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEbsSnapshotService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineEbsSnapshot())
	if err != nil {
		return fmt.Errorf("error on deleting ebs_snapshot %q, %s", d.Id(), err)
	}
	return err
}
