package iam_policy_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/iam/iam_policy"
)

const testAccVolcengineIamPoliciesDatasourceConfig = `
resource "volcengine_iam_policy" "foo1" {
    policy_name = "acc-test-policy1"
	description = "acc-test"
	policy_document = "{\"Statement\":[{\"Effect\":\"Allow\",\"Action\":[\"auto_scaling:DescribeScalingGroups\"],\"Resource\":[\"*\"]}]}"
}

resource "volcengine_iam_policy" "foo2" {
    policy_name = "acc-test-policy2"
	description = "acc-test"
	policy_document = "{\"Statement\":[{\"Effect\":\"Allow\",\"Action\":[\"auto_scaling:DescribeScalingConfigurations\"],\"Resource\":[\"*\"]}]}"
}

data "volcengine_iam_policies" "foo"{
    query = "${volcengine_iam_policy.foo1.description}"
}
`

func TestAccVolcengineIamPoliciesDatasource_Basic(t *testing.T) {
	resourceName := "data.volcengine_iam_policies.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &iam_policy.VolcengineIamPolicyService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers: volcengine.GetTestAccProviders(),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineIamPoliciesDatasourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(acc.ResourceId, "policies.#", "2"),
				),
			},
		},
	})
}
