package vpn_gateway_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpn/vpn_gateway"
	"testing"
)

const testAccVolcengineVpnGatewayCreateConfig = `
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

resource "volcengine_vpn_gateway" "foo" {
  vpc_id = "${volcengine_vpc.foo.id}"
  subnet_id = "${volcengine_subnet.foo.id}"
  bandwidth = 20
  vpn_gateway_name = "acc-test"
  description = "acc-test"
  period = 2
  project_name = "default"
}

`

const testAccVolcengineVpnGatewayUpdateConfig = `
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

resource "volcengine_vpn_gateway" "foo" {
  vpc_id = "${volcengine_vpc.foo.id}"
  subnet_id = "${volcengine_subnet.foo.id}"
  bandwidth = 50
  vpn_gateway_name = "acc-test1"
  description = "acc-test1"
  period = 2
  project_name = "default"
}

`

func TestAccVolcengineVpnGatewayResource_Basic(t *testing.T) {
	resourceName := "volcengine_vpn_gateway.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &vpn_gateway.VolcengineVpnGatewayService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineVpnGatewayCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "bandwidth", "20"),
					resource.TestCheckResourceAttr(acc.ResourceId, "billing_type", "PrePaid"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test"),
					resource.TestCheckResourceAttr(acc.ResourceId, "period", "2"),
					resource.TestCheckResourceAttr(acc.ResourceId, "project_name", "default"),
					resource.TestCheckResourceAttr(acc.ResourceId, "vpn_gateway_name", "acc-test"),
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

func TestAccVolcengineVpnGatewayResource_Update(t *testing.T) {
	resourceName := "volcengine_vpn_gateway.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &vpn_gateway.VolcengineVpnGatewayService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineVpnGatewayCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "bandwidth", "20"),
					resource.TestCheckResourceAttr(acc.ResourceId, "billing_type", "PrePaid"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test"),
					resource.TestCheckResourceAttr(acc.ResourceId, "period", "2"),
					resource.TestCheckResourceAttr(acc.ResourceId, "project_name", "default"),
					resource.TestCheckResourceAttr(acc.ResourceId, "vpn_gateway_name", "acc-test"),
				),
			},
			{
				Config: testAccVolcengineVpnGatewayUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "bandwidth", "50"),
					resource.TestCheckResourceAttr(acc.ResourceId, "billing_type", "PrePaid"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "period", "2"),
					resource.TestCheckResourceAttr(acc.ResourceId, "project_name", "default"),
					resource.TestCheckResourceAttr(acc.ResourceId, "vpn_gateway_name", "acc-test1"),
				),
			},
			{
				Config:             testAccVolcengineVpnGatewayUpdateConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // 修改之后，不应该再产生diff
			},
		},
	})
}
