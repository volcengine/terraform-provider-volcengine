package ecs_deployment_set_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/ecs_deployment_set"
	"testing"
)

const testAccVolcengineEcsDeploymentSetsDatasourceConfig = `
resource "volcengine_ecs_deployment_set" "foo" {
    deployment_set_name = "acc-test-ecs-ds-${count.index}"
	description = "acc-test"
    granularity = "switch"
    strategy = "Availability"
	count = 3
}

data "volcengine_ecs_deployment_sets" "foo"{
    granularity = "switch"
    ids = volcengine_ecs_deployment_set.foo[*].id
}
`

func TestAccVolcengineEcsDeploymentSetsDatasource_Basic(t *testing.T) {
	resourceName := "data.volcengine_ecs_deployment_sets.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &ecs_deployment_set.VolcengineEcsDeploymentSetService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers: volcengine.GetTestAccProviders(),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineEcsDeploymentSetsDatasourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(acc.ResourceId, "deployment_sets.#", "3"),
				),
			},
		},
	})
}
