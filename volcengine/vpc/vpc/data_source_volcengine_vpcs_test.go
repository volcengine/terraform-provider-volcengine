package vpc_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/vpc"
	"testing"
)

const testAccVpcDatasourceConfig = `
resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_vpc" "foo1" {
  vpc_name   = "acc-test-vpc1"
  cidr_block = "172.16.0.0/16"
}

data "volcengine_vpcs" "foo"{
  ids = ["${volcengine_vpc.foo1.id}", "${volcengine_vpc.foo.id}"]
}
`

func TestAccVolcengineVpcDatasource_Basic(t *testing.T) {
	resourceName := "data.volcengine_vpcs.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &vpc.VolcengineVpcService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers: volcengine.GetTestAccProviders(),
		Steps: []resource.TestStep{
			{
				Config: testAccVpcDatasourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(acc.ResourceId, "vpcs.#", "2"),
				),
			},
		},
	})
}
