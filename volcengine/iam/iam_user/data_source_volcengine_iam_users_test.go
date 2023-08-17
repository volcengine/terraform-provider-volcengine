package iam_user_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/iam/iam_user"
	"testing"
)

const testAccVolcengineIamUsersDatasourceConfig = `
resource "volcengine_iam_user" "foo" {
  user_name = "acc-test-user"
  description = "acc test"
  display_name = "name"
}
data "volcengine_iam_users" "foo"{
    user_names = [volcengine_iam_user.foo.user_name]
}
`

func TestAccVolcengineIamUsersDatasource_Basic(t *testing.T) {
	resourceName := "data.volcengine_iam_users.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &iam_user.VolcengineIamUserService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers: volcengine.GetTestAccProviders(),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineIamUsersDatasourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(acc.ResourceId, "users.#", "1"),
				),
			},
		},
	})
}
