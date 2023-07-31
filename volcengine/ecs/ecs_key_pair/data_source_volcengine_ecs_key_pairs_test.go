package ecs_key_pair_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/ecs_key_pair"
	"testing"
)

const testAccVolcengineEcsKeyPairsDatasourceConfig = `
resource "volcengine_ecs_key_pair" "foo" {
  key_pair_name = "acc-test-key-name"
  description ="acc-test"
}
data "volcengine_ecs_key_pairs" "foo"{
    key_pair_name = "${volcengine_ecs_key_pair.foo.key_pair_name}"
}
`

func TestAccVolcengineEcsKeyPairsDatasource_Basic(t *testing.T) {
	resourceName := "data.volcengine_ecs_key_pairs.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &ecs_key_pair.VolcengineEcsKeyPairService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers: volcengine.GetTestAccProviders(),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineEcsKeyPairsDatasourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(acc.ResourceId, "key_pairs.#", "1"),
				),
			},
		},
	})
}
