package subnet_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/subnet"
	"testing"
)

const testAccSubnetDatasourceConfig = `
data "volcengine_zones" "foo"{
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet"
  cidr_block = "172.16.0.0/24"
  zone_id = "${data.volcengine_zones.foo.zones[0].id}"
  vpc_id = "${volcengine_vpc.foo.id}"
}

resource "volcengine_subnet" "foo1" {
  subnet_name = "acc-test-subnet1"
  cidr_block = "172.16.1.0/24"
  zone_id = "${data.volcengine_zones.foo.zones[0].id}"
  vpc_id = "${volcengine_vpc.foo.id}"
}

data "volcengine_subnets" "foo"{
  ids = ["${volcengine_subnet.foo.id}", "${volcengine_subnet.foo1.id}"]
}
`

func TestAccVolcengineSubnetDatasource_Basic(t *testing.T) {
	resourceName := "data.volcengine_subnets.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &subnet.VolcengineSubnetService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers: volcengine.GetTestAccProviders(),
		Steps: []resource.TestStep{
			{
				Config: testAccSubnetDatasourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(acc.ResourceId, "subnets.#", "2"),
				),
			},
		},
	})
}
