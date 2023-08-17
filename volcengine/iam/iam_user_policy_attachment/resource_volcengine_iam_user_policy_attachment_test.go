package iam_user_policy_attachment_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/iam/iam_user_policy_attachment"
	"testing"
)

const testAccVolcengineIamUserPolicyAttachmentCreateConfig = `
resource "volcengine_iam_user" "foo" {
  user_name = "acc-test-user"
  description = "acc test"
  display_name = "name"
}
resource "volcengine_iam_policy" "foo" {
    policy_name = "acc-test-policy"
	description = "acc-test"
	policy_document = "{\"Statement\":[{\"Effect\":\"Allow\",\"Action\":[\"auto_scaling:DescribeScalingGroups\"],\"Resource\":[\"*\"]}]}"
}
resource "volcengine_iam_user_policy_attachment" "foo" {
    policy_name = volcengine_iam_policy.foo.policy_name
    policy_type = "Custom"
    user_name = volcengine_iam_user.foo.user_name
}
`

func TestAccVolcengineIamUserPolicyAttachmentResource_Basic(t *testing.T) {
	resourceName := "volcengine_iam_user_policy_attachment.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &iam_user_policy_attachment.VolcengineIamUserPolicyAttachmentService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineIamUserPolicyAttachmentCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "policy_type", "Custom"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
