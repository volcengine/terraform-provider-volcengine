package iam_role_policy_attachment

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-vestack/common"
)

/*

Import
Iam role policy attachment can be imported using the id, e.g.
```
$ terraform import vestack_iam_role_policy_attachment.default TerraformTestRole:TerraformTestPolicy:Custom
```

*/

func ResourceVestackIamRolePolicyAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceVestackIamRolePolicyAttachmentCreate,
		Read:   resourceVestackIamRolePolicyAttachmentRead,
		Delete: resourceVestackIamRolePolicyAttachmentDelete,
		Importer: &schema.ResourceImporter{
			State: iamRolePolicyAttachmentImporter,
		},
		Schema: map[string]*schema.Schema{
			"role_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the Role.",
			},
			"policy_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the Policy.",
			},
			"policy_type": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"System", "Custom"}, false),
				Description:  "The type of the Policy.",
			},
		},
	}
}

func resourceVestackIamRolePolicyAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	iamRolePolicyAttachmentService := NewIamRolePolicyAttachmentService(meta.(*ve.SdkClient))
	if err := iamRolePolicyAttachmentService.Dispatcher.Create(iamRolePolicyAttachmentService, d, ResourceVestackIamRolePolicyAttachment()); err != nil {
		return fmt.Errorf("error on creating iam role policy attachment %q, %w", d.Id(), err)
	}
	return resourceVestackIamRolePolicyAttachmentRead(d, meta)
}

func resourceVestackIamRolePolicyAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	iamRolePolicyAttachmentService := NewIamRolePolicyAttachmentService(meta.(*ve.SdkClient))
	if err := iamRolePolicyAttachmentService.Dispatcher.Read(iamRolePolicyAttachmentService, d, ResourceVestackIamRolePolicyAttachment()); err != nil {
		return fmt.Errorf("error on reading iam role policy attachment %q, %w", d.Id(), err)
	}
	return nil
}

func resourceVestackIamRolePolicyAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	iamRolePolicyAttachmentService := NewIamRolePolicyAttachmentService(meta.(*ve.SdkClient))
	if err := iamRolePolicyAttachmentService.Dispatcher.Delete(iamRolePolicyAttachmentService, d, ResourceVestackIamRolePolicyAttachment()); err != nil {
		return fmt.Errorf("error on deleting iam role policy attachment %q, %w", d.Id(), err)
	}
	return nil
}
