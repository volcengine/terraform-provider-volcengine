package iam_policy_project_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/iam/iam_policy_project"
)

func TestAccVolcengineIamPolicyProjectResource_Basic(t *testing.T) {
	resourceName := "volcengine_iam_policy_project.foo"
	userName := "acc-test-user-project"
	policyName := "acc-test-policy-project"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers: volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(&volcengine.AccTestResource{
			ResourceId: resourceName,
			Svc:        &iam_policy_project.VolcengineIamPolicyProjectService{},
		}),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineIamPolicyProjectResourceConfig(userName, policyName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "project_names.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "project_names.0", "default"),
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

func TestAccVolcengineIamPolicyProjectResource_SystemPolicy(t *testing.T) {
	resourceName := "volcengine_iam_policy_project.foo"
	userName := "acc-test-user-project-sys"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers: volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(&volcengine.AccTestResource{
			ResourceId: resourceName,
			Svc:        &iam_policy_project.VolcengineIamPolicyProjectService{},
		}),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineIamPolicyProjectResourceConfigSystem(userName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "project_names.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "project_names.0", "default"),
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

func testAccVolcengineIamPolicyProjectResourceConfigSystem(userName string) string {
	return fmt.Sprintf(`
resource "volcengine_iam_user" "foo" {
  user_name = "%s"
  display_name = "acc-test"
}

resource "volcengine_iam_policy_project" "foo" {
  principal_type = "User"
  principal_name = volcengine_iam_user.foo.user_name
  policy_type = "System"
  policy_name = "AdministratorAccess"
  project_names = ["default"]
}
`, userName)
}

func testAccVolcengineIamPolicyProjectResourceConfig(userName, policyName string) string {
	return fmt.Sprintf(`
resource "volcengine_iam_user" "foo" {
  user_name = "%s"
  display_name = "acc-test"
}

resource "volcengine_iam_policy" "foo" {
  policy_name = "%s"
  description = "acc-test"
  policy_document = "{\"Statement\":[{\"Effect\":\"Allow\",\"Action\":[\"auto_scaling:DescribeScalingGroups\"],\"Resource\":[\"*\"]}]}"
}

resource "volcengine_iam_policy_project" "foo" {
  principal_type = "User"
  principal_name = volcengine_iam_user.foo.user_name
  policy_type = "Custom"
  policy_name = volcengine_iam_policy.foo.policy_name
  project_names = ["default"]
}
`, userName, policyName)
}
