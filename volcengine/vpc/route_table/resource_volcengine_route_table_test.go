package route_table_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/route_table"
	"testing"
)

const testAccRouteTableForCreate = `
data "volcengine_zones" "foo"{
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_route_table" "foo" {
  vpc_id = "${volcengine_vpc.foo.id}"
  route_table_name = "acc-test-route-table"
}
`

func TestAccVolcengineRouteTableResource_Basic(t *testing.T) {
	resourceName := "volcengine_route_table.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &route_table.VolcengineRouteTableService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccRouteTableForCreate,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "route_table_name", "acc-test-route-table"),
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

const testAccRouteTableForUpdate = `
data "volcengine_zones" "foo"{
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_route_table" "foo" {
  vpc_id = "${volcengine_vpc.foo.id}"
  route_table_name = "acc-test-route-table-new"
  description = "tfdesc"
}
`

func TestAccVolcengineRouteTableResource_Update(t *testing.T) {
	resourceName := "volcengine_route_table.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &route_table.VolcengineRouteTableService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccRouteTableForCreate,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "route_table_name", "acc-test-route-table"),
				),
			},
			{
				Config: testAccRouteTableForUpdate,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "tfdesc"),
					resource.TestCheckResourceAttr(acc.ResourceId, "route_table_name", "acc-test-route-table-new"),
				),
			},
			{
				Config:             testAccRouteTableForUpdate,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // 修改之后，不应该再产生diff
			},
		},
	})
}
