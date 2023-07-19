package security_group_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/security_group"
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
  count = 3
}

data "volcengine_security_groups" "foo"{
  ids = ["${volcengine_security_group.foo[0].id}", "${volcengine_security_group.foo[1].id}", "${volcengine_security_group.foo[2].id}"]
}
`

func TestAccVolcengineSecurityGroupDatasource_Basic(t *testing.T) {
	resourceName := "data.volcengine_security_groups.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &security_group.VolcengineSecurityGroupService{},
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
					resource.TestCheckResourceAttr(acc.ResourceId, "security_groups.#", "3"),
				),
			},
		},
	})
}
