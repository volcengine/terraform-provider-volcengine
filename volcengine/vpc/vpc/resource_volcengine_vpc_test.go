package vpc_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/vpc"
	"testing"
)

const testAccVpcForCreate = `
resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}
`

const testAccVpcForUpdate = `
resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
  dns_servers = ["8.8.8.8", "114.114.114.114"]

  tags {
    key = "k2"
    value = "v2"
  }

  tags {
    key = "k1"
    value = "v1"
  }
}
`

func TestAccVolcengineVpcResource_Basic(t *testing.T) {
	resourceName := "volcengine_vpc.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &vpc.VolcengineVpcService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVpcForCreate,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "vpc_name", "acc-test-vpc"),
					resource.TestCheckResourceAttr(acc.ResourceId, "cidr_block", "172.16.0.0/16"),
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

func TestAccVolcengineVpcResource_Update(t *testing.T) {
	resourceName := "volcengine_vpc.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &vpc.VolcengineVpcService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVpcForCreate,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "vpc_name", "acc-test-vpc"),
					resource.TestCheckResourceAttr(acc.ResourceId, "cidr_block", "172.16.0.0/16"),
					// compute status
					resource.TestCheckResourceAttr(acc.ResourceId, "status", "Available"),
				),
			},
			{
				Config: testAccVpcForUpdate,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "vpc_name", "acc-test-vpc"),
					resource.TestCheckResourceAttr(acc.ResourceId, "cidr_block", "172.16.0.0/16"),
					// update attr check
					resource.TestCheckResourceAttr(acc.ResourceId, "dns_servers.#", "2"),
					resource.TestCheckResourceAttr(acc.ResourceId, "tags.#", "2"),
					volcengine.TestCheckTypeSetElemAttr(acc.ResourceId, "dns_servers.*", "8.8.8.8"),
					volcengine.TestCheckTypeSetElemAttr(acc.ResourceId, "dns_servers.*", "114.114.114.114"),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "tags.*", map[string]string{
						"key":   "k1",
						"value": "v1",
					}),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "tags.*", map[string]string{
						"key":   "k2",
						"value": "v2",
					}),
					// compute status check
					resource.TestCheckResourceAttr(acc.ResourceId, "status", "Available"),
				),
			},
			{
				Config:             testAccVpcForUpdate,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // 修改之后，不应该再产生diff
			},
		},
	})
}
