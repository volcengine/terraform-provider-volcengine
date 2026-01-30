package listener_health_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/clb/listener_health"
)

const testAccVolcengineListenerHealthsDatasourceConfig = `
data "volcengine_listener_healths" "foo" {
  listener_id = "lsn-13f1749gp981s3n6nu5c4209a"
}
`

func TestAccVolcengineListenerHealthsDatasource_Basic(t *testing.T) {
	resourceName := "data.volcengine_listener_healths.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &listener_health.VolcengineListenerHealthService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers: volcengine.GetTestAccProviders(),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineListenerHealthsDatasourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(acc.ResourceId, "health_info.#"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "total_count"),
				),
			},
		},
	})
}
