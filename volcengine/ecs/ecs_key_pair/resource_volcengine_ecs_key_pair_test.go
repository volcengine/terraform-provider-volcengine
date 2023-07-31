package ecs_key_pair_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/ecs_key_pair"
	"testing"
)

const testAccVolcengineEcsKeyPairCreateConfig = `
resource "volcengine_ecs_key_pair" "foo" {
  key_pair_name = "acc-test-key-name"
  description ="acc-test"
}
`

const testAccVolcengineEcsKeyPairUpdateConfig = `
resource "volcengine_ecs_key_pair" "foo" {
    description = "acc-test-2"
    key_pair_name = "acc-test-key-name"
}
`

func TestAccVolcengineEcsKeyPairResource_Basic(t *testing.T) {
	resourceName := "volcengine_ecs_key_pair.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &ecs_key_pair.VolcengineEcsKeyPairService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineEcsKeyPairCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test"),
					resource.TestCheckResourceAttr(acc.ResourceId, "key_pair_name", "acc-test-key-name"),
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

func TestAccVolcengineEcsKeyPairResource_Update(t *testing.T) {
	resourceName := "volcengine_ecs_key_pair.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &ecs_key_pair.VolcengineEcsKeyPairService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineEcsKeyPairCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test"),
					resource.TestCheckResourceAttr(acc.ResourceId, "key_pair_name", "acc-test-key-name"),
				),
			},
			{
				Config: testAccVolcengineEcsKeyPairUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test-2"),
					resource.TestCheckResourceAttr(acc.ResourceId, "key_pair_name", "acc-test-key-name"),
				),
			},
			{
				Config:             testAccVolcengineEcsKeyPairUpdateConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // 修改之后，不应该再产生diff
			},
		},
	})
}
