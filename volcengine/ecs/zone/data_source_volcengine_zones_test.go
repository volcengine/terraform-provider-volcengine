package zone_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/zone"
	"testing"
)

const testAccVolcengineZonesDatasourceConfig = `
data "volcengine_zones" "foo"{
    ids = ["cn-beijing-a"]
}
`

func TestAccVolcengineZonesDatasource_Basic(t *testing.T) {
	resourceName := "data.volcengine_zones.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &zone.VolcengineZoneService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers: volcengine.GetTestAccProviders(),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineZonesDatasourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(acc.ResourceId, "zones.#", "1"),
				),
			},
		},
	})
}
