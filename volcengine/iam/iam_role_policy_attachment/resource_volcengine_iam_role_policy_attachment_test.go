package iam_role_policy_attachment_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/iam/iam_role_policy_attachment"
)

const testAccVolcengineIamRolePolicyAttachmentCreateConfig = `
resource "volcengine_iam_role" "foo" {
	role_name = "acc-test-role"
    display_name = "acc-test"
	description = "acc-test"
    trust_policy_document = "{\"Statement\":[{\"Effect\":\"Allow\",\"Action\":[\"sts:AssumeRole\"],\"Principal\":{\"Service\":[\"auto_scaling\"]}}]}"
	max_session_duration = 3600
}

resource "volcengine_iam_policy" "foo" {
    policy_name = "acc-test-policy"
	description = "acc-test"
	policy_document = "{\"Statement\":[{\"Effect\":\"Allow\",\"Action\":[\"auto_scaling:DescribeScalingGroups\"],\"Resource\":[\"*\"]}]}"
}

resource "volcengine_iam_role_policy_attachment" "foo" {
    policy_name = volcengine_iam_policy.foo.policy_name
    policy_type = "Custom"
    role_name = volcengine_iam_role.foo.role_name
}
`

func TestAccVolcengineIamRolePolicyAttachmentResource_Basic(t *testing.T) {
	resourceName := "volcengine_iam_role_policy_attachment.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &iam_role_policy_attachment.VolcengineIamRolePolicyAttachmentService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineIamRolePolicyAttachmentCreateConfig,
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
