package ipv6_gateway_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/ipv6_gateway"
)

const testAccVpcIpv6GatewayCreate = `
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
`

const testAccVpcIpv6GatewayUpdate = `
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
	  name = "acc-test-2"
	  description = "test update"
	}
`

func TestAccVpcIpv6GatewayResource_Basic(t *testing.T) {
	resourceName := "volcengine_vpc_ipv6_gateway.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &ipv6_gateway.VolcengineIpv6GatewayService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVpcIpv6GatewayCreate,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "name", "acc-test-1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "test"),
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

func TestAccVpcIpv6GatewayResource_Update(t *testing.T) {
	resourceName := "volcengine_vpc_ipv6_gateway.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &ipv6_gateway.VolcengineIpv6GatewayService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVpcIpv6GatewayCreate,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "name", "acc-test-1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "test"),
				),
			},
			{
				Config: testAccVpcIpv6GatewayUpdate,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "name", "acc-test-2"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "test update"),
				),
			},
			{
				Config:             testAccVpcIpv6GatewayUpdate,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}
