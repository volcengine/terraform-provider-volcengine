package server_group_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/clb/server_group"
	"testing"
)

const testAccVolcengineServerGroupCreateConfig = `
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

resource "volcengine_clb" "foo" {
  type = "public"
  subnet_id = "${volcengine_subnet.foo.id}"
  load_balancer_spec = "small_1"
  description = "acc0Demo"
  load_balancer_name = "acc-test-create"
  eip_billing_config {
    isp = "BGP"
    eip_billing_type = "PostPaidByBandwidth"
    bandwidth = 1
  }
}

resource "volcengine_server_group" "foo" {
  load_balancer_id = "${volcengine_clb.foo.id}"
  server_group_name = "acc-test-create"
  description = "hello demo11"
}

`

const testAccVolcengineServerGroupUpdateConfig = `
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

resource "volcengine_clb" "foo" {
  type = "public"
  subnet_id = "${volcengine_subnet.foo.id}"
  load_balancer_spec = "small_1"
  description = "acc0Demo"
  load_balancer_name = "acc-test-create"
  eip_billing_config {
    isp = "BGP"
    eip_billing_type = "PostPaidByBandwidth"
    bandwidth = 1
  }
}

resource "volcengine_server_group" "foo" {
  load_balancer_id = "${volcengine_clb.foo.id}"
  server_group_name = "acc-test-create1"
  description = "acc hello demo11"
}

`

func TestAccVolcengineServerGroupResource_Basic(t *testing.T) {
	resourceName := "volcengine_server_group.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &server_group.VolcengineServerGroupService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineServerGroupCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "hello demo11"),
					resource.TestCheckResourceAttr(acc.ResourceId, "server_group_name", "acc-test-create"),
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

func TestAccVolcengineServerGroupResource_Update(t *testing.T) {
	resourceName := "volcengine_server_group.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &server_group.VolcengineServerGroupService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineServerGroupCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "hello demo11"),
					resource.TestCheckResourceAttr(acc.ResourceId, "server_group_name", "acc-test-create"),
				),
			},
			{
				Config: testAccVolcengineServerGroupUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc hello demo11"),
					resource.TestCheckResourceAttr(acc.ResourceId, "server_group_name", "acc-test-create1"),
				),
			},
			{
				Config:             testAccVolcengineServerGroupUpdateConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // 修改之后，不应该再产生diff
			},
		},
	})
}

const testAccVolcengineServerGroupCreateConfigIpv6 = `
data "volcengine_zones" "foo"{
}

resource "volcengine_vpc" "vpc_ipv6" {
  vpc_name = "acc-test-vpc-ipv6"
  cidr_block = "172.16.0.0/16"
  enable_ipv6 = true
}

resource "volcengine_subnet" "subnet_ipv6" {
  subnet_name = "acc-test-subnet-ipv6"
  cidr_block = "172.16.0.0/24"
  zone_id = data.volcengine_zones.foo.zones[1].id
  vpc_id = volcengine_vpc.vpc_ipv6.id
  ipv6_cidr_block = 1
}

resource "volcengine_clb" "private_clb_ipv6" {
  type = "private"
  subnet_id = volcengine_subnet.subnet_ipv6.id
  load_balancer_name = "acc-test-clb-ipv6"
  load_balancer_spec = "small_1"
  description = "acc-test-demo"
  project_name = "default"
  address_ip_version = "DualStack"
  tags {
    key = "k1"
    value = "v1"
  }
}

resource "volcengine_server_group" "foo" {
  load_balancer_id = "${volcengine_clb.private_clb_ipv6.id}"
  server_group_name = "acc-test-sg-ipv6"
  description = "acc-test"
  address_ip_version = "ipv6"
}
`

func TestAccVolcengineServerGroupResource_CreateIpv6(t *testing.T) {
	resourceName := "volcengine_server_group.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &server_group.VolcengineServerGroupService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineServerGroupCreateConfigIpv6,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "server_group_name", "acc-test-sg-ipv6"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test"),
					resource.TestCheckResourceAttr(acc.ResourceId, "address_ip_version", "ipv6"),
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
