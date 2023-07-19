package security_group_rule_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/security_group_rule"
	"testing"
)

const testAccSecurityGroupDatasourceConfig = `
data "volcengine_zones" "foo"{
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
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
  cidr_ip           = "172.16.0.0/24"
}

data "volcengine_security_group_rules" "foo"{
  security_group_id = "${volcengine_security_group.foo.id}"
  direction = "${volcengine_security_group_rule.foo.direction}"
  cidr_ip = "${volcengine_security_group_rule.foo.cidr_ip}"
}
`

func TestAccVolcengineSecurityGroupRulesDatasource_Basic(t *testing.T) {
	resourceName := "data.volcengine_security_group_rules.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &security_group_rule.VolcengineSecurityGroupRuleService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers: volcengine.GetTestAccProviders(),
		Steps: []resource.TestStep{
			{
				Config: testAccSecurityGroupDatasourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(acc.ResourceId, "security_group_rules.#", "1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "security_group_rules.0.direction", "egress"),
				),
			},
		},
	})
}
