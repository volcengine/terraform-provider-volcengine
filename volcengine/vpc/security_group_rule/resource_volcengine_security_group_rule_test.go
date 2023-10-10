package security_group_rule_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/security_group_rule"
	"testing"
)

const testAccSecurityGroupRuleForCreate = `
data "volcengine_zones" "foo"{
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
  enable_ipv6 = true
}

resource "volcengine_security_group" "foo" {
  vpc_id = "${volcengine_vpc.foo.id}"
  security_group_name = "acc-test-security-group"
}

resource "volcengine_security_group_rule" "foo" {
  direction         = "egress"
  security_group_id = "${volcengine_security_group.foo.id}"
  protocol          = "tcp"
  port_start        = 8000
  port_end          = 9003
  cidr_ip           = "2406:d440:10d:ff00::/64"
}
`

const testAccSecurityGroupRuleForUpdate = `
data "volcengine_zones" "foo"{
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
  enable_ipv6 = true
}

resource "volcengine_security_group" "foo" {
  vpc_id = "${volcengine_vpc.foo.id}"
  security_group_name = "acc-test-security-group"
}

resource "volcengine_security_group_rule" "foo" {
  direction         = "egress"
  security_group_id = "${volcengine_security_group.foo.id}"
  protocol          = "tcp"
  port_start        = 8000
  port_end          = 9003
  cidr_ip           = "2406:d440:10d:ff00::/64"
  description       = "tfdesc"
}
`

func TestAccVolcengineSecurityGroupRuleResource_Basic(t *testing.T) {
	resourceName := "volcengine_security_group_rule.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &security_group_rule.VolcengineSecurityGroupRuleService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityGroupRuleForCreate,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "direction", "egress"),
					resource.TestCheckResourceAttr(acc.ResourceId, "protocol", "tcp"),
					resource.TestCheckResourceAttr(acc.ResourceId, "port_start", "8000"),
					resource.TestCheckResourceAttr(acc.ResourceId, "port_end", "9003"),
					resource.TestCheckResourceAttr(acc.ResourceId, "cidr_ip", "2406:d440:10d:ff00::/64"),
					// compute status
					resource.TestCheckResourceAttr(acc.ResourceId, "status", "Available"),
				),
			},
			{
				Config:             testAccSecurityGroupRuleForCreate,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false,
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccVolcengineSubnetResource_Update(t *testing.T) {
	resourceName := "volcengine_security_group_rule.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &security_group_rule.VolcengineSecurityGroupRuleService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityGroupRuleForCreate,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "direction", "egress"),
					resource.TestCheckResourceAttr(acc.ResourceId, "protocol", "tcp"),
					resource.TestCheckResourceAttr(acc.ResourceId, "port_start", "8000"),
					resource.TestCheckResourceAttr(acc.ResourceId, "port_end", "9003"),
					resource.TestCheckResourceAttr(acc.ResourceId, "cidr_ip", "2406:d440:10d:ff00::/64"),
					// compute status
					resource.TestCheckResourceAttr(acc.ResourceId, "status", "Available"),
				),
			},
			{
				Config: testAccSecurityGroupRuleForUpdate,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "direction", "egress"),
					resource.TestCheckResourceAttr(acc.ResourceId, "protocol", "tcp"),
					resource.TestCheckResourceAttr(acc.ResourceId, "port_start", "8000"),
					resource.TestCheckResourceAttr(acc.ResourceId, "port_end", "9003"),
					resource.TestCheckResourceAttr(acc.ResourceId, "cidr_ip", "2406:d440:10d:ff00::/64"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "tfdesc"),
					// compute status
					resource.TestCheckResourceAttr(acc.ResourceId, "status", "Available"),
				),
			},
			{
				Config:             testAccSecurityGroupRuleForUpdate,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // 修改之后，不应该再产生diff
			},
		},
	})
}
