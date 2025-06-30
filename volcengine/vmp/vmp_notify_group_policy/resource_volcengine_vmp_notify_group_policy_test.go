package vmp_notify_group_policy_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vmp/vmp_notify_group_policy"
)

const testAccVolcengineVmpNotifyGroupPolicyCreateConfig = `
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
`

const testAccVolcengineVmpNotifyGroupPolicyUpdateConfig = `
resource "volcengine_vmp_notify_group_policy" "foo" {
  name = "acc-test-2"
  description = "acc-test-2"
  levels {
    level = "P0"
    group_by = ["__rule__"]
    group_wait = "35"
    group_interval = "30"
    repeat_interval = "30"
  }
  levels {
    level = "P1"
    group_by = ["__rule__"]
    group_wait = "30"
    group_interval = "30"
    repeat_interval = "30"
  }
  levels {
    level = "P2"
    group_by = ["__rule__"]
    group_wait = "40"
    group_interval = "45"
    repeat_interval = "30"
  }
}
`

func TestAccVolcengineVmpNotifyGroupPolicyResource_Basic(t *testing.T) {
	resourceName := "volcengine_vmp_notify_group_policy.foo"

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
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineVmpNotifyGroupPolicyCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test-1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "levels.#", "3"),
					resource.TestCheckResourceAttr(acc.ResourceId, "name", "acc-test-1"),
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

func TestAccVolcengineVmpNotifyGroupPolicyResource_Update(t *testing.T) {
	resourceName := "volcengine_vmp_notify_group_policy.foo"

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
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineVmpNotifyGroupPolicyCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test-1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "levels.#", "3"),
					resource.TestCheckResourceAttr(acc.ResourceId, "name", "acc-test-1"),
				),
			},
			{
				Config: testAccVolcengineVmpNotifyGroupPolicyUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test-2"),
					resource.TestCheckResourceAttr(acc.ResourceId, "levels.#", "3"),
					resource.TestCheckResourceAttr(acc.ResourceId, "name", "acc-test-2"),
				),
			},
			{
				Config:             testAccVolcengineVmpNotifyGroupPolicyUpdateConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // 修改之后，不应该再产生diff
			},
		},
	})
}
