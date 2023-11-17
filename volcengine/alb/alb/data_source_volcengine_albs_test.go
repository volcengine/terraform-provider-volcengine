package alb_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/alb/alb"
)

const testAccVolcengineAlbsDatasourceConfig = `
data "volcengine_alb_zones" "foo"{
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "subnet_1" {
  subnet_name = "acc-test-subnet-1"
  cidr_block = "172.16.1.0/24"
  zone_id = data.volcengine_alb_zones.foo.zones[0].id
  vpc_id = volcengine_vpc.foo.id
}

resource "volcengine_subnet" "subnet_2" {
  subnet_name = "acc-test-subnet-2"
  cidr_block = "172.16.2.0/24"
  zone_id = data.volcengine_alb_zones.foo.zones[1].id
  vpc_id = volcengine_vpc.foo.id
}

resource "volcengine_alb" "foo" {
  address_ip_version = "IPv4"
  type = "private"
  load_balancer_name = "acc-test-alb-private-${count.index}"
  description = "acc-test"
  subnet_ids = [volcengine_subnet.subnet_1.id, volcengine_subnet.subnet_2.id]
  project_name = "default"
  delete_protection = "off"
  tags {
    key = "k1"
    value = "v1"
  }
  count = 3
}

data "volcengine_albs" "foo" {
  ids = volcengine_alb.foo[*].id
}
`

func TestAccVolcengineAlbsDatasource_Basic(t *testing.T) {
	resourceName := "data.volcengine_albs.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return alb.NewAlbService(client)
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers: volcengine.GetTestAccProviders(),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineAlbsDatasourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(acc.ResourceId, "albs.#", "3"),
				),
			},
		},
	})
}
