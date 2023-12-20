package ecs_iam_role_attachment

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
IamRoleAttachment can be imported using the iam_role_name:instance_id, e.g.
```
$ terraform import volcengine_iam_role_attachment.default role_name:instance_id
```

*/

func ResourceVolcengineIamRoleAttachment() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineIamRoleAttachmentCreate,
		Read:   resourceVolcengineIamRoleAttachmentRead,
		Delete: resourceVolcengineIamRoleAttachmentDelete,
		Importer: &schema.ResourceImporter{
			State: iamRoleAttachmentImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"iam_role_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the iam role.",
			},
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The id of the ecs instance.",
			},
		},
	}
	return resource
}

func resourceVolcengineIamRoleAttachmentCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewIamRoleAttachmentService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineIamRoleAttachment())
	if err != nil {
		return fmt.Errorf("error on creating iam_role_attachment %q, %s", d.Id(), err)
	}
	return resourceVolcengineIamRoleAttachmentRead(d, meta)
}

func resourceVolcengineIamRoleAttachmentRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewIamRoleAttachmentService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineIamRoleAttachment())
	if err != nil {
		return fmt.Errorf("error on reading iam_role_attachment %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineIamRoleAttachmentDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewIamRoleAttachmentService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineIamRoleAttachment())
	if err != nil {
		return fmt.Errorf("error on deleting iam_role_attachment %q, %s", d.Id(), err)
	}
	return err
}

var iamRoleAttachmentImporter = func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
	items := strings.Split(data.Id(), ":")
	if len(items) != 2 {
		return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
	}
	if err := data.Set("iam_role_name", items[0]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	if err := data.Set("instance_id", items[1]); err != nil {
		return []*schema.ResourceData{data}, err
	}
	return []*schema.ResourceData{data}, nil
}
