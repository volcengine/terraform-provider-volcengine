package ecs_available_resource_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/ecs_available_resource"
)

const testAccVolcengineAvailableResourcesDatasourceConfig = `
data "volcengine_ecs_available_resources" "foo"{
    destination_resource = "InstanceType"
}
`

func TestAccVolcengineAvailableResourcesDatasource_Basic(t *testing.T) {
	resourceName := "data.volcengine_ecs_available_resources.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return ecs_available_resource.NewEcsAvailableResourceService(client)
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers: volcengine.GetTestAccProviders(),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineAvailableResourcesDatasourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(acc.ResourceId, "available_zones.#", "3"),
				),
			},
		},
	})
}
