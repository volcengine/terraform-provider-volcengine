package ebs_snapshot_group

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
EbsSnapshotGroup can be imported using the id, e.g.
```
$ terraform import volcengine_ebs_snapshot_group.default resource_id
```

*/

func ResourceVolcengineEbsSnapshotGroup() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineEbsSnapshotGroupCreate,
		Read:   resourceVolcengineEbsSnapshotGroupRead,
		Update: resourceVolcengineEbsSnapshotGroupUpdate,
		Delete: resourceVolcengineEbsSnapshotGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"volume_ids": {
				Type:     schema.TypeSet,
				Required: true,
				ForceNew: true,
				Set:      schema.HashString,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The volume id of the snapshot group. The status of the volume must be `attached`." +
					"If multiple volumes are specified, they need to be attached to the same ECS instance.",
			},
			"instance_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				ForceNew:    true,
				Description: "The instance id of the snapshot group.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The name of the snapshot group.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The instance id of the snapshot group.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The project name of the snapshot group.",
			},
			"tags": ve.TagsSchema(),

			// computed fields
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the snapshot group.",
			},
			"image_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The image id of the snapshot group.",
			},
			"creation_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The creation time of the snapshot group.",
			},
		},
	}
	return resource
}

func resourceVolcengineEbsSnapshotGroupCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEbsSnapshotGroupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineEbsSnapshotGroup())
	if err != nil {
		return fmt.Errorf("error on creating ebs_snapshot_group %q, %s", d.Id(), err)
	}
	return resourceVolcengineEbsSnapshotGroupRead(d, meta)
}

func resourceVolcengineEbsSnapshotGroupRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEbsSnapshotGroupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineEbsSnapshotGroup())
	if err != nil {
		return fmt.Errorf("error on reading ebs_snapshot_group %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineEbsSnapshotGroupUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEbsSnapshotGroupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineEbsSnapshotGroup())
	if err != nil {
		return fmt.Errorf("error on updating ebs_snapshot_group %q, %s", d.Id(), err)
	}
	return resourceVolcengineEbsSnapshotGroupRead(d, meta)
}

func resourceVolcengineEbsSnapshotGroupDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEbsSnapshotGroupService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineEbsSnapshotGroup())
	if err != nil {
		return fmt.Errorf("error on deleting ebs_snapshot_group %q, %s", d.Id(), err)
	}
	return err
}
