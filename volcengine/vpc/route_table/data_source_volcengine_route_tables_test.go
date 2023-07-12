package route_table_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/route_table"
	"testing"
)

const testAccRouteTableDatasourceConfig = `
data "volcengine_zones" "foo"{
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_route_table" "foo" {
  vpc_id = "${volcengine_vpc.foo.id}"
  route_table_name = "acc-test-route-table"
  count = 3
}

data "volcengine_route_tables" "foo" {
  ids = ["${volcengine_route_table.foo[0].id}", "${volcengine_route_table.foo[1].id}", "${volcengine_route_table.foo[2].id}"]
}
`

func TestAccVolcengineRouteTableDatasource_Basic(t *testing.T) {
	resourceName := "data.volcengine_route_tables.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &route_table.VolcengineRouteTableService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers: volcengine.GetTestAccProviders(),
		Steps: []resource.TestStep{
			{
				Config: testAccRouteTableDatasourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(acc.ResourceId, "route_tables.#", "3"),
				),
			},
		},
	})
}
