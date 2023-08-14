package iam_user_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/iam/iam_user"
	"testing"
)

const testAccVolcengineIamUserCreateConfig = `
resource "volcengine_iam_user" "foo" {
  user_name = "acc-test-user"
  description = "acc test"
  display_name = "name"
}
`

const testAccVolcengineIamUserUpdateConfig = `
resource "volcengine_iam_user" "foo" {
    description = "acc test update"
    display_name = "name2"
    email = "xxx@163.com"
    user_name = "acc-test-user2"
}
`

func TestAccVolcengineIamUserResource_Basic(t *testing.T) {
	resourceName := "volcengine_iam_user.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &iam_user.VolcengineIamUserService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineIamUserCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc test"),
					resource.TestCheckResourceAttr(acc.ResourceId, "display_name", "name"),
					resource.TestCheckResourceAttr(acc.ResourceId, "email", ""),
					resource.TestCheckResourceAttr(acc.ResourceId, "mobile_phone", ""),
					resource.TestCheckResourceAttr(acc.ResourceId, "user_name", "acc-test-user"),
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

func TestAccVolcengineIamUserResource_Update(t *testing.T) {
	resourceName := "volcengine_iam_user.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &iam_user.VolcengineIamUserService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineIamUserCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc test"),
					resource.TestCheckResourceAttr(acc.ResourceId, "display_name", "name"),
					resource.TestCheckResourceAttr(acc.ResourceId, "email", ""),
					resource.TestCheckResourceAttr(acc.ResourceId, "mobile_phone", ""),
					resource.TestCheckResourceAttr(acc.ResourceId, "user_name", "acc-test-user"),
				),
			},
			{
				Config: testAccVolcengineIamUserUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc test update"),
					resource.TestCheckResourceAttr(acc.ResourceId, "display_name", "name2"),
					resource.TestCheckResourceAttr(acc.ResourceId, "email", "xxx@163.com"),
					resource.TestCheckResourceAttr(acc.ResourceId, "mobile_phone", ""),
					resource.TestCheckResourceAttr(acc.ResourceId, "user_name", "acc-test-user2"),
				),
			},
			{
				Config:             testAccVolcengineIamUserUpdateConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // 修改之后，不应该再产生diff
			},
		},
	})
}
