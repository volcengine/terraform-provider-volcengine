package scaling_group_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/autoscaling/scaling_group"
	"testing"
)

const testAccScalingGroupForCreate = `
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
  scaling_group_name = "acc-test-scaling-group"
  subnet_ids = ["${volcengine_subnet.foo.id}"]
  multi_az_policy = "BALANCE"
  desire_instance_number = 0
  min_instance_number = 0
  max_instance_number = 1
  instance_terminate_policy = "OldestInstance"
  default_cooldown = 10
}
`

const testAccScalingGroupForUpdate = `
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
  scaling_group_name = "acc-test-scaling-group-new"
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
`

func TestAccVolcengineScalingGroupResource_Basic(t *testing.T) {
	resourceName := "volcengine_scaling_group.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &scaling_group.VolcengineScalingGroupService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccScalingGroupForCreate,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "scaling_group_name", "acc-test-scaling-group"),
					resource.TestCheckResourceAttr(acc.ResourceId, "multi_az_policy", "BALANCE"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_terminate_policy", "OldestInstance"),
					resource.TestCheckResourceAttr(acc.ResourceId, "max_instance_number", "1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "min_instance_number", "0"),
					resource.TestCheckResourceAttr(acc.ResourceId, "desire_instance_number", "0"),
					resource.TestCheckResourceAttr(acc.ResourceId, "default_cooldown", "10"),
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

func TestAccVolcengineScalingGroupResource_Update(t *testing.T) {
	resourceName := "volcengine_scaling_group.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &scaling_group.VolcengineScalingGroupService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccScalingGroupForCreate,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "scaling_group_name", "acc-test-scaling-group"),
					resource.TestCheckResourceAttr(acc.ResourceId, "multi_az_policy", "BALANCE"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_terminate_policy", "OldestInstance"),
					resource.TestCheckResourceAttr(acc.ResourceId, "max_instance_number", "1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "min_instance_number", "0"),
					resource.TestCheckResourceAttr(acc.ResourceId, "desire_instance_number", "0"),
					resource.TestCheckResourceAttr(acc.ResourceId, "default_cooldown", "10"),
				),
			},
			{
				Config: testAccScalingGroupForUpdate,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "scaling_group_name", "acc-test-scaling-group-new"),
					resource.TestCheckResourceAttr(acc.ResourceId, "multi_az_policy", "BALANCE"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_terminate_policy", "OldestInstance"),
					resource.TestCheckResourceAttr(acc.ResourceId, "max_instance_number", "10"),
					resource.TestCheckResourceAttr(acc.ResourceId, "min_instance_number", "0"),
					resource.TestCheckResourceAttr(acc.ResourceId, "desire_instance_number", "0"),
					resource.TestCheckResourceAttr(acc.ResourceId, "default_cooldown", "30"),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "tags.*", map[string]string{
						"key":   "k1",
						"value": "v1",
					}),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "tags.*", map[string]string{
						"key":   "k2",
						"value": "v2",
					}),
				),
			},
			{
				Config:             testAccScalingGroupForUpdate,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // 修改之后，不应该再产生diff
			},
		},
	})
}
