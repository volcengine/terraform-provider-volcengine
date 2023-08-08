package iam_login_profile_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/iam/iam_login_profile"
)

const testAccVolcengineIamLoginProfileCreateConfig = `
resource "volcengine_iam_user" "foo" {
  	user_name = "acc-test-user"
  	description = "acc-test"
  	display_name = "name"
}

resource "volcengine_iam_login_profile" "foo" {
    user_name = "${volcengine_iam_user.foo.user_name}"
  	password = "93f0cb0614Aab12"
  	login_allowed = true
	password_reset_required = false
}
`

func TestAccVolcengineIamLoginProfileResource_Basic(t *testing.T) {
	resourceName := "volcengine_iam_login_profile.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &iam_login_profile.VolcengineIamLoginProfileService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineIamLoginProfileCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "login_allowed", "true"),
					resource.TestCheckResourceAttr(acc.ResourceId, "password", "93f0cb0614Aab12"),
					resource.TestCheckResourceAttr(acc.ResourceId, "password_reset_required", "false"),
					resource.TestCheckResourceAttr(acc.ResourceId, "user_name", "acc-test-user"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"password"},
			},
		},
	})
}

const testAccVolcengineIamLoginProfileUpdateConfig = `
resource "volcengine_iam_user" "foo" {
  	user_name = "acc-test-user"
  	description = "acc-test"
  	display_name = "name"
}

resource "volcengine_iam_login_profile" "foo" {
    user_name = "${volcengine_iam_user.foo.user_name}"
  	password = "93f0cb0614Aab12177"
  	login_allowed = false
	password_reset_required = true
}
`

func TestAccVolcengineIamLoginProfileResource_Update(t *testing.T) {
	resourceName := "volcengine_iam_login_profile.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &iam_login_profile.VolcengineIamLoginProfileService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineIamLoginProfileCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "login_allowed", "true"),
					resource.TestCheckResourceAttr(acc.ResourceId, "password", "93f0cb0614Aab12"),
					resource.TestCheckResourceAttr(acc.ResourceId, "password_reset_required", "false"),
					resource.TestCheckResourceAttr(acc.ResourceId, "user_name", "acc-test-user"),
				),
			},
			{
				Config: testAccVolcengineIamLoginProfileUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "login_allowed", "false"),
					resource.TestCheckResourceAttr(acc.ResourceId, "password", "93f0cb0614Aab12177"),
					resource.TestCheckResourceAttr(acc.ResourceId, "password_reset_required", "true"),
					resource.TestCheckResourceAttr(acc.ResourceId, "user_name", "acc-test-user"),
				),
			},
			{
				Config:             testAccVolcengineIamLoginProfileUpdateConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // 修改之后，不应该再产生diff
			},
		},
	})
}
