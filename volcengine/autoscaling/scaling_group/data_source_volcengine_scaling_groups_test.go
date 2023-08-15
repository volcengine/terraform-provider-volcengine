package scaling_group_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/autoscaling/scaling_group"
	"testing"
)

const testAccScalingGroupDatasourceConfig = `
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

resource "volcengine_scaling_group" "foo" {
  count = 3
  scaling_group_name = "acc-test-scaling-group-${count.index}"
  subnet_ids = ["${volcengine_subnet.foo.id}"]
  multi_az_policy = "BALANCE"
  desire_instance_number = 0
  min_instance_number = 0
  max_instance_number = 10
  instance_terminate_policy = "OldestInstance"
  default_cooldown = 30

  tags {
    key = "k2"
    value = "v2"
  }

  tags {
    key = "k1"
    value = "v1"
  }
}

data "volcengine_scaling_groups" "foo"{
  ids = volcengine_scaling_group.foo[*].id
}
`

func TestAccVolcengineScalingGroupDatasource_Basic(t *testing.T) {
	resourceName := "data.volcengine_scaling_groups.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &scaling_group.VolcengineScalingGroupService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers: volcengine.GetTestAccProviders(),
		Steps: []resource.TestStep{
			{
				Config: testAccScalingGroupDatasourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(acc.ResourceId, "scaling_groups.#", "3"),
				),
			},
		},
	})
}
