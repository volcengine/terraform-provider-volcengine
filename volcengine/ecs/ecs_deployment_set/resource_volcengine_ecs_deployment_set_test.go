package ecs_deployment_set_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/ecs_deployment_set"
	"testing"
)

const testAccVolcengineEcsDeploymentSetCreateConfig = `
resource "volcengine_ecs_deployment_set" "foo" {
    deployment_set_name = "acc-test-ecs-ds"
	description = "acc-test"
    granularity = "switch"
    strategy = "Availability"
}
`

func TestAccVolcengineEcsDeploymentSetResource_Basic(t *testing.T) {
	resourceName := "volcengine_ecs_deployment_set.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &ecs_deployment_set.VolcengineEcsDeploymentSetService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineEcsDeploymentSetCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "deployment_set_name", "acc-test-ecs-ds"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test"),
					resource.TestCheckResourceAttr(acc.ResourceId, "granularity", "switch"),
					resource.TestCheckResourceAttr(acc.ResourceId, "strategy", "Availability"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"description"},
			},
		},
	})
}

const testAccVolcengineEcsDeploymentSetUpdateConfig = `
resource "volcengine_ecs_deployment_set" "foo" {
    deployment_set_name = "acc-test-ecs-ds-new"
	description = "acc-test"
    granularity = "switch"
    strategy = "Availability"
}
`

func TestAccVolcengineEcsDeploymentSetResource_Update(t *testing.T) {
	resourceName := "volcengine_ecs_deployment_set.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &ecs_deployment_set.VolcengineEcsDeploymentSetService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineEcsDeploymentSetCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "deployment_set_name", "acc-test-ecs-ds"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test"),
					resource.TestCheckResourceAttr(acc.ResourceId, "granularity", "switch"),
					resource.TestCheckResourceAttr(acc.ResourceId, "strategy", "Availability"),
				),
			},
			{
				Config: testAccVolcengineEcsDeploymentSetUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "deployment_set_name", "acc-test-ecs-ds-new"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test"),
					resource.TestCheckResourceAttr(acc.ResourceId, "granularity", "switch"),
					resource.TestCheckResourceAttr(acc.ResourceId, "strategy", "Availability"),
				),
			},
			{
				Config:             testAccVolcengineEcsDeploymentSetUpdateConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // 修改之后，不应该再产生diff
			},
		},
	})
}
