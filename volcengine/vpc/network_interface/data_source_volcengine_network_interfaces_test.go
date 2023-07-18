package network_interface_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/network_interface"
)

const testAccVolcengineNetworkInterfacesDatasourceConfig = `
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

resource "volcengine_security_group" "foo" {
  security_group_name = "acc-test-sg"
  vpc_id = "${volcengine_vpc.foo.id}"
}

resource "volcengine_network_interface" "foo" {
  network_interface_name = "acc-test-eni-${count.index}"
  subnet_id = "${volcengine_subnet.foo.id}"
  security_group_ids = ["${volcengine_security_group.foo.id}"]
  count = 3
}

data "volcengine_network_interfaces" "foo"{
    ids = volcengine_network_interface.foo[*].id
}
`

func TestAccVolcengineNetworkInterfacesDatasource_Basic(t *testing.T) {
	resourceName := "data.volcengine_network_interfaces.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &network_interface.VolcengineNetworkInterfaceService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers: volcengine.GetTestAccProviders(),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineNetworkInterfacesDatasourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(acc.ResourceId, "network_interfaces.#", "3"),
				),
			},
		},
	})
}
