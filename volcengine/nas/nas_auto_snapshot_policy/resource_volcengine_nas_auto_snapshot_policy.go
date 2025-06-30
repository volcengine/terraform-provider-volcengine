package nas_auto_snapshot_policy

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
NasAutoSnapshotPolicy can be imported using the id, e.g.
```
$ terraform import volcengine_nas_auto_snapshot_policy.default resource_id
```

*/

func ResourceVolcengineNasAutoSnapshotPolicy() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineNasAutoSnapshotPolicyCreate,
		Read:   resourceVolcengineNasAutoSnapshotPolicyRead,
		Update: resourceVolcengineNasAutoSnapshotPolicyUpdate,
		Delete: resourceVolcengineNasAutoSnapshotPolicyDelete,
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
			"repeat_weekdays": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The repeat weekdays of the auto snapshot policy. Support setting multiple dates, separated by English commas. Valid values: `1` ~ `7`.",
			},
			"time_points": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The time points of the auto snapshot policy. Support setting multiple dates, separated by English commas. Valid values: `0` ~ `23`.",
			},
			"retention_days": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "The retention days of the auto snapshot policy. Valid values: -1(permanent) or 1 ~ 65536. Default is 30.",
			},

			// computed fields
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The status of auto snapshot policy.",
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The create time of auto snapshot policy.",
			},
			"file_system_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The count of file system which auto snapshot policy bind.",
			},
		},
	}
	return resource
}

func resourceVolcengineNasAutoSnapshotPolicyCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewNasAutoSnapshotPolicyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineNasAutoSnapshotPolicy())
	if err != nil {
		return fmt.Errorf("error on creating nas_auto_snapshot_policy %q, %s", d.Id(), err)
	}
	return resourceVolcengineNasAutoSnapshotPolicyRead(d, meta)
}

func resourceVolcengineNasAutoSnapshotPolicyRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewNasAutoSnapshotPolicyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineNasAutoSnapshotPolicy())
	if err != nil {
		return fmt.Errorf("error on reading nas_auto_snapshot_policy %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineNasAutoSnapshotPolicyUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewNasAutoSnapshotPolicyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineNasAutoSnapshotPolicy())
	if err != nil {
		return fmt.Errorf("error on updating nas_auto_snapshot_policy %q, %s", d.Id(), err)
	}
	return resourceVolcengineNasAutoSnapshotPolicyRead(d, meta)
}

func resourceVolcengineNasAutoSnapshotPolicyDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewNasAutoSnapshotPolicyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineNasAutoSnapshotPolicy())
	if err != nil {
		return fmt.Errorf("error on deleting nas_auto_snapshot_policy %q, %s", d.Id(), err)
	}
	return err
}
