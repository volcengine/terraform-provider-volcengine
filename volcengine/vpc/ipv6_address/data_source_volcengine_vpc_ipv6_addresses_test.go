package ipv6_address_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/ipv6_address"
)

const testAccV6AddressDatasourceConfig = `
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

	resource "volcengine_ecs_instance" "foo" {
	  image_id = "${data.volcengine_images.foo.images[0].image_id}"
	  instance_type = "ecs.g1.large"
	  instance_name = "acc-test-ecs-name"
	  password = "93f0cb0614Aab12"
	  instance_charge_type = "PostPaid"
	  system_volume_type = "ESSD_PL0"
	  system_volume_size = 40
	  subnet_id = volcengine_subnet.foo.id
	  security_group_ids = [volcengine_security_group.foo.id]
	  ipv6_address_count = 1
	}

	data "volcengine_vpc_ipv6_addresses" "foo"{
	  associated_instance_id = "${volcengine_ecs_instance.foo.id}"
	}
`

func TestAccVolcengineV6AddressDatasource_Basic(t *testing.T) {
	resourceName := "data.volcengine_vpc_ipv6_addresses.foo"
	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &ipv6_address.VolcengineIpv6AddressService{},
	}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers: volcengine.GetTestAccProviders(),
		Steps: []resource.TestStep{
			{
				Config: testAccV6AddressDatasourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(acc.ResourceId, "ipv6_addresses.#", "1"),
				),
			},
		},
	})
}
