package iam_user_group_policy_attachment

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
IamUserGroupPolicyAttachment can be imported using the user group name and policy name, e.g.
```
$ terraform import volcengine_iam_user_group_policy_attachment.default userGroupName:policyName
```

*/

func ResourceVolcengineIamUserGroupPolicyAttachment() *schema.Resource {
	resource := &schema.Resource{
		Create: resourceVolcengineIamUserGroupPolicyAttachmentCreate,
		Read:   resourceVolcengineIamUserGroupPolicyAttachmentRead,
		Delete: resourceVolcengineIamUserGroupPolicyAttachmentDelete,
		Importer: &schema.ResourceImporter{
			State: func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
				items := strings.Split(data.Id(), ":")
				if len(items) != 2 {
					return []*schema.ResourceData{data}, fmt.Errorf("import id must split with ':'")
				}
				if err := data.Set("user_group_name", items[0]); err != nil {
					return []*schema.ResourceData{data}, err
				}
				if err := data.Set("policy_name", items[1]); err != nil {
					return []*schema.ResourceData{data}, err
				}
				return []*schema.ResourceData{data}, nil
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"user_group_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The user group name.",
			},
			"policy_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The policy name.",
			},
			"policy_type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Strategy types, System strategy, Custom strategy.",
			},
		},
	}
	return resource
}

func resourceVolcengineIamUserGroupPolicyAttachmentCreate(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewIamUserGroupPolicyAttachmentService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Create(service, d, ResourceVolcengineIamUserGroupPolicyAttachment())
	if err != nil {
		return fmt.Errorf("error on creating iam_user_group_policy_attachment %q, %s", d.Id(), err)
	}
	return resourceVolcengineIamUserGroupPolicyAttachmentRead(d, meta)
}

func resourceVolcengineIamUserGroupPolicyAttachmentRead(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewIamUserGroupPolicyAttachmentService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Read(service, d, ResourceVolcengineIamUserGroupPolicyAttachment())
	if err != nil {
		return fmt.Errorf("error on reading iam_user_group_policy_attachment %q, %s", d.Id(), err)
	}
	return err
}

func resourceVolcengineIamUserGroupPolicyAttachmentDelete(d *schema.ResourceData, meta interface{}) (err error) {
	service := NewIamUserGroupPolicyAttachmentService(meta.(*ve.SdkClient))
	err = service.Dispatcher.Delete(service, d, ResourceVolcengineIamUserGroupPolicyAttachment())
	if err != nil {
		return fmt.Errorf("error on deleting iam_user_group_policy_attachment %q, %s", d.Id(), err)
	}
	return err
}
