package iam_access_key_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/iam/iam_access_key"
)

const testAccVolcengineIamAccessKeyCreateConfig = `
resource "volcengine_iam_user" "foo" {
  	user_name = "acc-test-user"
  	description = "acc-test"
  	display_name = "name"
}

resource "volcengine_iam_access_key" "foo" {
	user_name = "${volcengine_iam_user.foo.user_name}"
    secret_file = "./sk"
    status = "active"
}
`

func TestAccVolcengineIamAccessKeyResource_Basic(t *testing.T) {
	resourceName := "volcengine_iam_access_key.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &iam_access_key.VolcengineIamAccessKeyService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineIamAccessKeyCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "user_name", "acc-test-user"),
					resource.TestCheckResourceAttr(acc.ResourceId, "secret_file", "./sk"),
					resource.TestCheckResourceAttr(acc.ResourceId, "status", "active"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "create_date"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "secret"),
					resource.TestCheckNoResourceAttr(acc.ResourceId, "pgp_key"),
					resource.TestCheckNoResourceAttr(acc.ResourceId, "encrypted_secret"),
					resource.TestCheckNoResourceAttr(acc.ResourceId, "key_fingerprint"),
				),
			},
		},
	})
}

const testAccVolcengineIamAccessKeyUpdateConfig = `
resource "volcengine_iam_user" "foo" {
  	user_name = "acc-test-user"
  	description = "acc-test"
  	display_name = "name"
}

resource "volcengine_iam_access_key" "foo" {
	user_name = "${volcengine_iam_user.foo.user_name}"
    secret_file = "./sk"
    status = "inactive"
}
`

func TestAccVolcengineIamAccessKeyResource_Update(t *testing.T) {
	resourceName := "volcengine_iam_access_key.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &iam_access_key.VolcengineIamAccessKeyService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineIamAccessKeyCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "user_name", "acc-test-user"),
					resource.TestCheckResourceAttr(acc.ResourceId, "secret_file", "./sk"),
					resource.TestCheckResourceAttr(acc.ResourceId, "status", "active"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "create_date"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "secret"),
					resource.TestCheckNoResourceAttr(acc.ResourceId, "pgp_key"),
					resource.TestCheckNoResourceAttr(acc.ResourceId, "encrypted_secret"),
					resource.TestCheckNoResourceAttr(acc.ResourceId, "key_fingerprint"),
				),
			},
			{
				Config: testAccVolcengineIamAccessKeyUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "user_name", "acc-test-user"),
					resource.TestCheckResourceAttr(acc.ResourceId, "secret_file", "./sk"),
					resource.TestCheckResourceAttr(acc.ResourceId, "status", "inactive"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "create_date"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "secret"),
					resource.TestCheckNoResourceAttr(acc.ResourceId, "pgp_key"),
					resource.TestCheckNoResourceAttr(acc.ResourceId, "encrypted_secret"),
					resource.TestCheckNoResourceAttr(acc.ResourceId, "key_fingerprint"),
				),
			},
			{
				Config:             testAccVolcengineIamAccessKeyUpdateConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // 修改之后，不应该再产生diff
			},
		},
	})
}
