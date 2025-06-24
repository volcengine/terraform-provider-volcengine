package vmp_notify_group_policy_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vmp/vmp_notify_group_policy"
)

const testAccVolcengineVmpNotifyGroupPoliciesDatasourceConfig = `
resource "volcengine_vmp_notify_group_policy" "foo" {
  name = "acc-test-1"
  description = "acc-test-1"
  levels {
    level = "P2"
    group_by = ["__rule__"]
    group_wait = "35"
    group_interval = "30"
    repeat_interval = "30"
  }
  levels {
    level = "P0"
    group_by = ["__rule__"]
    group_wait = "30"
    group_interval = "30"
    repeat_interval = "30"
  }
  levels {
    level = "P1"
    group_by = ["__rule__"]
    group_wait = "40"
    group_interval = "45"
    repeat_interval = "30"
  }
}

data "volcengine_vmp_notify_group_policies" "foo"{
    ids = [volcengine_vmp_notify_group_policy.foo.id]
}
`

func TestAccVolcengineVmpNotifyGroupPoliciesDatasource_Basic(t *testing.T) {
	resourceName := "data.volcengine_vmp_notify_group_policies.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return vmp_notify_group_policy.NewService(client)
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers: volcengine.GetTestAccProviders(),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineVmpNotifyGroupPoliciesDatasourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(acc.ResourceId, "notify_policies.#", "1"),
				),
			},
		},
	})
}
