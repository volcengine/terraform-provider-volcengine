package cloud_identity_permission_set_assignment

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
CloudIdentityPermissionSetAssignment can be imported using the permission_set_id:target_id:principal_type:principal_id, e.g.
```
$ terraform import volcengine_cloud_identity_permission_set_assignment.default resource_id
```

*/

func ResourceVolcengineCloudIdentityPermissionSetAssignment() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineCloudIdentityPermissionSetAssignmentCreate,
		Read:   resourceVolcengineCloudIdentityPermissionSetAssignmentRead,
		Delete: resourceVolcengineCloudIdentityPermissionSetAssignmentDelete,
		Importer: &schema.ResourceImporter{
			State: permissionSetAssignmentImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"permission_set_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the cloud identity permission set.",
			},
			"target_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The target account id of the cloud identity permission set assignment.",
			},
			"principal_type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"User", "Group"}, false),
				Description:  "The principal type of the cloud identity permission set. Valid values: `User`, `Group`.",
			},
			"principal_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The principal id of the cloud identity permission set. When the `principal_type` is `User`, this field is specified to `UserId`. When the `principal_type` is `Group`, this field is specified to `GroupId`.",
			},
		},
	}
	return resource
}

func resourceVolcengineCloudIdentityPermissionSetAssignmentCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCloudIdentityPermissionSetAssignmentService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineCloudIdentityPermissionSetAssignment())
	if err != nil {
		return fmt.Errorf("error on creating cloud_identity_permission_set_assignment %q, %s", d.Id(), err)
	}
	return resourceVolcengineCloudIdentityPermissionSetAssignmentRead(d, meta)
}

func resourceVolcengineCloudIdentityPermissionSetAssignmentRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCloudIdentityPermissionSetAssignmentService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineCloudIdentityPermissionSetAssignment())
	if err != nil {
		return fmt.Errorf("error on reading cloud_identity_permission_set_assignment %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineCloudIdentityPermissionSetAssignmentDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewCloudIdentityPermissionSetAssignmentService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineCloudIdentityPermissionSetAssignment())
	if err != nil {
		return fmt.Errorf("error on deleting cloud_identity_permission_set_assignment %q, %s", d.Id(), err)
	}
	return err
}

var permissionSetAssignmentImporter = func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(data.Id(), ":")
	if len(items) != 4 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
	}
	if err := data.Set("permission_set_id", items[0]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	if err := data.Set("target_id", items[1]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	if err := data.Set("principal_type", items[2]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	if err := data.Set("principal_id", items[3]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}
