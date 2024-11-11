package ebs_auto_snapshot_policy

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
EbsAutoSnapshotPolicy can be imported using the id, e.g.
```
$ terraform import volcengine_ebs_auto_snapshot_policy.default resource_id
```

*/

func ResourceVolcengineEbsAutoSnapshotPolicy() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineEbsAutoSnapshotPolicyCreate,
		Read:   resourceVolcengineEbsAutoSnapshotPolicyRead,
		Update: resourceVolcengineEbsAutoSnapshotPolicyUpdate,
		Delete: resourceVolcengineEbsAutoSnapshotPolicyDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"auto_snapshot_policy_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the auto snapshot policy.",
			},
			"time_points": {
				Type:     schema.TypeSet,
				Required: true,
				Set:      schema.HashString,
				MinItems: 1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "The creation time points of the auto snapshot policy. The value range is `0~23`, representing a total of 24 time points from 00:00 to 23:00, for example, 1 represents 01:00.",
			},
			"retention_days": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The retention days of the auto snapshot. Valid values: -1 and 1~65536. `-1` means permanently preserving the snapshot.",
			},
			"repeat_weekdays": {
				Type:     schema.TypeSet,
				Optional: true,
				Set:      schema.HashString,
				MinItems: 1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				ExactlyOneOf: []string{"repeat_weekdays", "repeat_days"},
				Description: "The date of creating snapshot repeatedly by week. The value range is `1-7`, for example, 1 represents Monday. " +
					"Only one of `repeat_weekdays, repeat_days` can be specified.",
			},
			"repeat_days": {
				Type:     schema.TypeInt,
				Optional: true,
				Description: "Create snapshots repeatedly on a daily basis, with intervals of a certain number of days between each snapshot. The value range is `1-30`. " +
					"Only one of `repeat_weekdays, repeat_days` can be specified.",
			},
			"project_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The project name of the auto snapshot policy.",
			},
			"tags": ve.TagsSchema(),

			// computed fields
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of the auto snapshot policy.",
			},
			"volume_nums": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of volumes associated with the auto snapshot policy.",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The creation time of the auto snapshot policy.",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The updated time of the auto snapshot policy.",
			},
		},
	}
	return resource
}

func resourceVolcengineEbsAutoSnapshotPolicyCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEbsAutoSnapshotPolicyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineEbsAutoSnapshotPolicy())
	if err != nil {
		return fmt.Errorf("error on creating ebs_auto_snapshot_policy %q, %s", d.Id(), err)
	}
	return resourceVolcengineEbsAutoSnapshotPolicyRead(d, meta)
}

func resourceVolcengineEbsAutoSnapshotPolicyRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEbsAutoSnapshotPolicyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineEbsAutoSnapshotPolicy())
	if err != nil {
		return fmt.Errorf("error on reading ebs_auto_snapshot_policy %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineEbsAutoSnapshotPolicyUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEbsAutoSnapshotPolicyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineEbsAutoSnapshotPolicy())
	if err != nil {
		return fmt.Errorf("error on updating ebs_auto_snapshot_policy %q, %s", d.Id(), err)
	}
	return resourceVolcengineEbsAutoSnapshotPolicyRead(d, meta)
}

func resourceVolcengineEbsAutoSnapshotPolicyDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewEbsAutoSnapshotPolicyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineEbsAutoSnapshotPolicy())
	if err != nil {
		return fmt.Errorf("error on deleting ebs_auto_snapshot_policy %q, %s", d.Id(), err)
	}
	return err
}
