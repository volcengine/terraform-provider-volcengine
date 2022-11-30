package iam_user_policy_attachment

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
)

/*

Import
Iam user policy attachment can be imported using the UserName:PolicyName:PolicyType, e.g.
```
$ terraform import volcengine_iam_user_policy_attachment.default TerraformTestUser:TerraformTestPolicy:Custom
```

*/

func ResourceVolcengineIamUserPolicyAttachment() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolcengineIamUserPolicyAttachmentCreate,
		Read:   resourceVolcengineIamUserPolicyAttachmentRead,
		Delete: resourceVolcengineIamUserPolicyAttachmentDelete,
		Importer: &schema.ResourceImporter{
			State: iamUserPolicyAttachmentImporter,
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
		Schema: map[string]*schema.Schema{
			"user_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the user.",
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

func resourceVolcengineIamUserPolicyAttachmentCreate(d *schema.ResourceData, meta interface{}) error {
	iamUserPolicyAttachmentService := NewIamUserPolicyAttachmentService(meta.(*ve.SdkClient))
	if err := iamUserPolicyAttachmentService.Dispatcher.Create(iamUserPolicyAttachmentService, d, ResourceVolcengineIamUserPolicyAttachment()); err != nil {
		return fmt.Errorf("error on creating iam user policy attachment %q, %w", d.Id(), err)
	}
	return resourceVolcengineIamUserPolicyAttachmentRead(d, meta)
}

func resourceVolcengineIamUserPolicyAttachmentRead(d *schema.ResourceData, meta interface{}) error {
	iamUserPolicyAttachmentService := NewIamUserPolicyAttachmentService(meta.(*ve.SdkClient))
	if err := iamUserPolicyAttachmentService.Dispatcher.Read(iamUserPolicyAttachmentService, d, ResourceVolcengineIamUserPolicyAttachment()); err != nil {
		return fmt.Errorf("error on reading iam user policy attachment %q, %w", d.Id(), err)
	}
	return nil
}

func resourceVolcengineIamUserPolicyAttachmentDelete(d *schema.ResourceData, meta interface{}) error {
	iamUserPolicyAttachmentService := NewIamUserPolicyAttachmentService(meta.(*ve.SdkClient))
	if err := iamUserPolicyAttachmentService.Dispatcher.Delete(iamUserPolicyAttachmentService, d, ResourceVolcengineIamUserPolicyAttachment()); err != nil {
		return fmt.Errorf("error on deleting iam user policy attachment %q, %w", d.Id(), err)
	}
	return nil
}
