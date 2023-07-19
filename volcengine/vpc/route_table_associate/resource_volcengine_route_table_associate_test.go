package route_table_associate_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/route_table_associate"
	"testing"
)

const testAccRouteTableAssociateForCreate = `
data "volcengine_zones" "foo"{
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc-attach"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet-attach"
  cidr_block = "172.16.0.0/24"
  zone_id = "${data.volcengine_zones.foo.zones[0].id}"
  vpc_id = "${volcengine_vpc.foo.id}"
}

resource "volcengine_route_table" "foo" {
  vpc_id = "${volcengine_vpc.foo.id}"
  route_table_name = "acc-test-route-table-attach"
}

resource "volcengine_route_table_associate" "foo" {
  route_table_id = "${volcengine_route_table.foo.id}"
  subnet_id = "${volcengine_subnet.foo.id}"
}
`

func TestAccVolcengineRouteTableAssociateResource_Basic(t *testing.T) {
	resourceName := "volcengine_route_table_associate.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &route_table_associate.VolcengineRouteTableAssociateService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccRouteTableAssociateForCreate,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
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
