package ipv6_gateway_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/ipv6_gateway"
)

const testAccIpv6GatewayConfig = `
	data "volcengine_zones" "foo"{
	}

	data "volcengine_images" "foo" {
	  os_type = "Linux"
	  visibility = "public"
	  instance_type_id = "ecs.g1.large"
	}

	resource "volcengine_vpc" "foo" {
	  vpc_name   = "acc-test-vpc"
	  cidr_block = "172.16.0.0/16"
	  enable_ipv6 = true
	}

	resource "volcengine_subnet" "foo" {
	  subnet_name = "acc-test-subnet"
	  cidr_block = "172.16.0.0/24"
	  zone_id = "${data.volcengine_zones.foo.zones[0].id}"
	  vpc_id = "${volcengine_vpc.foo.id}"
	  ipv6_cidr_block = 1
	}

	resource "volcengine_security_group" "foo" {
	  vpc_id = "${volcengine_vpc.foo.id}"
	  security_group_name = "acc-test-security-group"
	}

	resource "volcengine_vpc_ipv6_gateway" "foo" {
	  vpc_id = "${volcengine_vpc.foo.id}"
	  name = "acc-test-1"
	  description = "test"
	}

	data "volcengine_vpc_ipv6_gateways" "foo" {
		ids = ["${volcengine_vpc_ipv6_gateway.foo.id}"]
	}
`

func TestAccVolcengineIpv6GatewayDataSource_Basic(t *testing.T) {
	resourceName := "data.volcengine_vpc_ipv6_gateways.foo"
	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &ipv6_gateway.VolcengineIpv6GatewayService{},
	}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers: volcengine.GetTestAccProviders(),
		Steps: []resource.TestStep{
			{
				Config: testAccIpv6GatewayConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(acc.ResourceId, "ipv6_gateways.#", "1"),
				),
			},
		},
	})
}
