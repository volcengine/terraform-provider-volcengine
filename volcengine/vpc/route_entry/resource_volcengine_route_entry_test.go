package route_entry_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/route_entry"
	"testing"
)

const testAccRouteEntryForCreate = `
data "volcengine_zones" "foo"{
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc-rn"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet-rn"
  cidr_block = "172.16.0.0/24"
  zone_id = "${data.volcengine_zones.foo.zones[0].id}"
  vpc_id = "${volcengine_vpc.foo.id}"
}

resource "volcengine_nat_gateway" "foo" {
  vpc_id = "${volcengine_vpc.foo.id}"
  subnet_id = "${volcengine_subnet.foo.id}"
  spec = "Small"
  nat_gateway_name = "acc-test-nat-rn"
}

resource "volcengine_route_table" "foo" {
  vpc_id = "${volcengine_vpc.foo.id}"
  route_table_name = "acc-test-route-table"
}

resource "volcengine_route_entry" "foo" {
  route_table_id = "${volcengine_route_table.foo.id}"
  destination_cidr_block = "172.16.1.0/24"
  next_hop_type = "NatGW"
  next_hop_id = "${volcengine_nat_gateway.foo.id}"
  route_entry_name = "acc-test-route-entry"
}
`

func TestAccVolcengineRouteEntryResource_Basic(t *testing.T) {
	resourceName := "volcengine_route_entry.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &route_entry.VolcengineRouteEntryService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccRouteEntryForCreate,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "route_entry_name", "acc-test-route-entry"),
					// compute status
					resource.TestCheckResourceAttr(acc.ResourceId, "status", "Available"),
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

const testAccRouteEntryForUpdate = `
data "volcengine_zones" "foo"{
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc-rn"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet-rn"
  cidr_block = "172.16.0.0/24"
  zone_id = "${data.volcengine_zones.foo.zones[0].id}"
  vpc_id = "${volcengine_vpc.foo.id}"
}

resource "volcengine_nat_gateway" "foo" {
  vpc_id = "${volcengine_vpc.foo.id}"
  subnet_id = "${volcengine_subnet.foo.id}"
  spec = "Small"
  nat_gateway_name = "acc-test-nat-rn"
}

resource "volcengine_route_table" "foo" {
  vpc_id = "${volcengine_vpc.foo.id}"
  route_table_name = "acc-test-route-table"
}

resource "volcengine_route_entry" "foo" {
  route_table_id = "${volcengine_route_table.foo.id}"
  destination_cidr_block = "172.16.1.0/24"
  next_hop_type = "NatGW"
  next_hop_id = "${volcengine_nat_gateway.foo.id}"
  route_entry_name = "acc-test-route-entry-new"
  description = "tfdesc"
}
`

func TestAccVolcengineRouteEntryResource_Update(t *testing.T) {
	resourceName := "volcengine_route_entry.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &route_entry.VolcengineRouteEntryService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccRouteEntryForCreate,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "route_entry_name", "acc-test-route-entry"),
					// compute status
					resource.TestCheckResourceAttr(acc.ResourceId, "status", "Available"),
				),
			},
			{
				Config: testAccRouteEntryForUpdate,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "tfdesc"),
					resource.TestCheckResourceAttr(acc.ResourceId, "route_entry_name", "acc-test-route-entry-new"),
					// compute status
					resource.TestCheckResourceAttr(acc.ResourceId, "status", "Available"),
				),
			},
			{
				Config:             testAccRouteEntryForUpdate,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // 修改之后，不应该再产生diff
			},
		},
	})
}
