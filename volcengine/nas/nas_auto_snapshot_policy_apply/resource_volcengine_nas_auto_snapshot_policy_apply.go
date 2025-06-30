package nas_auto_snapshot_policy_apply

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
NasAutoSnapshotPolicyApply can be imported using the auto_snapshot_policy_id:file_system_id, e.g.
```
$ terraform import volcengine_nas_auto_snapshot_policy_apply.default resource_id
```

*/

func ResourceVolcengineNasAutoSnapshotPolicyApply() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineNasAutoSnapshotPolicyApplyCreate,
		Read:   resourceVolcengineNasAutoSnapshotPolicyApplyRead,
		Delete: resourceVolcengineNasAutoSnapshotPolicyApplyDelete,
		Importer: &schema.ResourceImporter{
			State: importAutoSnapshotPolicyApply,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"auto_snapshot_policy_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of auto snapshot policy.",
			},
			"file_system_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of file system.",
			},
		},
	}
	return resource
}

func resourceVolcengineNasAutoSnapshotPolicyApplyCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewNasAutoSnapshotPolicyApplyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineNasAutoSnapshotPolicyApply())
	if err != nil {
		return fmt.Errorf("error on creating nas_auto_snapshot_policy_apply %q, %s", d.Id(), err)
	}
	return resourceVolcengineNasAutoSnapshotPolicyApplyRead(d, meta)
}

func resourceVolcengineNasAutoSnapshotPolicyApplyRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewNasAutoSnapshotPolicyApplyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineNasAutoSnapshotPolicyApply())
	if err != nil {
		return fmt.Errorf("error on reading nas_auto_snapshot_policy_apply %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineNasAutoSnapshotPolicyApplyUpdate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewNasAutoSnapshotPolicyApplyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Update(service, d, ResourceVolcengineNasAutoSnapshotPolicyApply())
	if err != nil {
		return fmt.Errorf("error on updating nas_auto_snapshot_policy_apply %q, %s", d.Id(), err)
	}
	return resourceVolcengineNasAutoSnapshotPolicyApplyRead(d, meta)
}

func resourceVolcengineNasAutoSnapshotPolicyApplyDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewNasAutoSnapshotPolicyApplyService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineNasAutoSnapshotPolicyApply())
	if err != nil {
		return fmt.Errorf("error on deleting nas_auto_snapshot_policy_apply %q, %s", d.Id(), err)
	}
	return err
}

func importAutoSnapshotPolicyApply(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	var err error
	items := strings.Split(d.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{d}, fmt.Errorf("import id must be of the form auto_snapshot_policy_id:file_system_id")
	}
	err = d.Set("auto_snapshot_policy_id", items[0])
	if err != nil {
		return []*schema.ResourceData{d}, err
	}
	err = d.Set("file_system_id", items[1])
	if err != nil {
		return []*schema.ResourceData{d}, err
	}

	return []*schema.ResourceData{d}, nil
}
