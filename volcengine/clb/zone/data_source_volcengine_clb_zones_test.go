package zone_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/clb/zone"
	"testing"
)

const testAccVolcengineClbZonesDatasourceConfig = `
data "volcengine_clb_zones" "foo"{
}
`

func TestAccVolcengineClbZonesDatasource_Basic(t *testing.T) {
	resourceName := "data.volcengine_clb_zones.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &zone.VolcengineClbZoneService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers: volcengine.GetTestAccProviders(),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineClbZonesDatasourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(acc.ResourceId, "master_zones.#", "1"),
				),
			},
		},
	})
}
