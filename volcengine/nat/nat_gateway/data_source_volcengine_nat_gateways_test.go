package nat_gateway_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/nat/nat_gateway"
	"testing"
)

const testAccVolcengineNatGatewaysDatasourceConfig = `
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

resource "volcengine_nat_gateway" "foo" {
	vpc_id = "${volcengine_vpc.foo.id}"
    subnet_id = "${volcengine_subnet.foo.id}"
	spec = "Small"
	nat_gateway_name = "acc-test-ng-${count.index}"
	description = "acc-test"
	billing_type = "PostPaid"
	project_name = "default"
	tags {
		key = "k1"
		value = "v1"
	}
	count =3 
}

data "volcengine_nat_gateways" "foo"{
    ids = volcengine_nat_gateway.foo[*].id
}
`

func TestAccVolcengineNatGatewaysDatasource_Basic(t *testing.T) {
	resourceName := "data.volcengine_nat_gateways.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &nat_gateway.VolcengineNatGatewayService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers: volcengine.GetTestAccProviders(),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineNatGatewaysDatasourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(acc.ResourceId, "nat_gateways.#", "3"),
				),
			},
		},
	})
}
