package iam_role_policy_attachment

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Iam role policy attachment can be imported using the id, e.g.
```
$ terraform import volcengine_iam_role_policy_attachment.default TerraformTestRole:TerraformTestPolicy:Custom
```

*/

func ResourceVolcengineIamRolePolicyAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineIamRolePolicyAttachmentCreate,
		Read:   resourceVolcengineIamRolePolicyAttachmentRead,
		Delete: resourceVolcengineIamRolePolicyAttachmentDelete,
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

func resourceVolcengineIamRolePolicyAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	iamRolePolicyAttachmentService := NewIamRolePolicyAttachmentService(meta.(*ve.SdkClient))
	if err := iamRolePolicyAttachmentService.Dispatcher.Create(iamRolePolicyAttachmentService, d, ResourceVolcengineIamRolePolicyAttachment()); err != nil {
		return fmt.Errorf("error on creating iam role policy attachment %q, %w", d.Id(), err)
	}
	return resourceVolcengineIamRolePolicyAttachmentRead(d, meta)
}

func resourceVolcengineIamRolePolicyAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	iamRolePolicyAttachmentService := NewIamRolePolicyAttachmentService(meta.(*ve.SdkClient))
	if err := iamRolePolicyAttachmentService.Dispatcher.Read(iamRolePolicyAttachmentService, d, ResourceVolcengineIamRolePolicyAttachment()); err != nil {
		return fmt.Errorf("error on reading iam role policy attachment %q, %w", d.Id(), err)
	}
	return nil
}

func resourceVolcengineIamRolePolicyAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	iamRolePolicyAttachmentService := NewIamRolePolicyAttachmentService(meta.(*ve.SdkClient))
	if err := iamRolePolicyAttachmentService.Dispatcher.Delete(iamRolePolicyAttachmentService, d, ResourceVolcengineIamRolePolicyAttachment()); err != nil {
		return fmt.Errorf("error on deleting iam role policy attachment %q, %w", d.Id(), err)
	}
	return nil
}
