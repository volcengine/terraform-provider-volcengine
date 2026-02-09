package iam_role_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/iam/iam_role"
)

const testAccVolcengineIamRolesDatasourceConfig = `
resource "volcengine_iam_role" "foo1" {
	role_name = "acc-test-role1"
    display_name = "acc-test1"
	description = "acc-test1"
    trust_policy_document = "{\"Statement\":[{\"Effect\":\"Allow\",\"Action\":[\"sts:AssumeRole\"],\"Principal\":{\"Service\":[\"auto_scaling\"]}}]}"
	max_session_duration = 3600
}

resource "volcengine_iam_role" "foo2" {
    role_name = "acc-test-role2"
    display_name = "acc-test2"
	description = "acc-test2"
    trust_policy_document = "{\"Statement\":[{\"Effect\":\"Allow\",\"Action\":[\"sts:AssumeRole\"],\"Principal\":{\"Service\":[\"ecs\"]}}]}"
	max_session_duration = 3600
}

data "volcengine_iam_roles" "foo"{
    query = "acc-test-role"
    depends_on = [volcengine_iam_role.foo1, volcengine_iam_role.foo2]
}
`

func TestAccVolcengineIamRolesDatasource_Basic(t *testing.T) {
	resourceName := "data.volcengine_iam_roles.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &iam_role.VolcengineIamRoleService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers: volcengine.GetTestAccProviders(),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineIamRolesDatasourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(acc.ResourceId, "roles.#", "2"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "roles.0.update_date"),
				),
			},
		},
	})
}
