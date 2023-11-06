package scaling_lifecycle_hook_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/autoscaling/scaling_lifecycle_hook"
	"testing"
)

const testAccVolcengineScalingLifecycleHookCreateConfig = `
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
  scaling_group_name = "acc-test-scaling-group-lifecycle"
  subnet_ids = ["${volcengine_subnet.foo.id}"]
  multi_az_policy = "BALANCE"
  desire_instance_number = 0
  min_instance_number = 0
  max_instance_number = 1
  instance_terminate_policy = "OldestInstance"
  default_cooldown = 10
}

resource "volcengine_scaling_lifecycle_hook" "foo" {
    lifecycle_hook_name = "acc-test-lifecycle"
    lifecycle_hook_policy = "CONTINUE"
    lifecycle_hook_timeout = 30
    lifecycle_hook_type = "SCALE_IN"
    scaling_group_id = "${volcengine_scaling_group.foo.id}"
}
`

const testAccVolcengineScalingLifecycleHookUpdateConfig = `
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

resource "volcengine_ecs_command" "foo" {
  name = "acc-test-command"
  description = "tf"
  working_dir = "/home"
  username = "root"
  timeout = 100
  command_content = "IyEvYmluL2Jhc2gKCgplY2hvICJvcGVyYXRpb24gc3VjY2VzcyEi"
}

resource "volcengine_scaling_group" "foo" {
  scaling_group_name = "acc-test-scaling-group-lifecycle"
  subnet_ids = ["${volcengine_subnet.foo.id}"]
  multi_az_policy = "BALANCE"
  desire_instance_number = 0
  min_instance_number = 0
  max_instance_number = 1
  instance_terminate_policy = "OldestInstance"
  default_cooldown = 10
}

resource "volcengine_scaling_lifecycle_hook" "foo" {
    lifecycle_hook_name = "acc-test-lifecycle"
    lifecycle_hook_policy = "ROLLBACK"
    lifecycle_hook_timeout = 300
    lifecycle_hook_type = "SCALE_OUT"
    scaling_group_id = "${volcengine_scaling_group.foo.id}"
	lifecycle_command {
    command_id = volcengine_ecs_command.foo.id
    parameters = "{}"
  }
}
`

func TestAccVolcengineScalingLifecycleHookResource_Basic(t *testing.T) {
	resourceName := "volcengine_scaling_lifecycle_hook.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &scaling_lifecycle_hook.VolcengineScalingLifecycleHookService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineScalingLifecycleHookCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "lifecycle_hook_name", "acc-test-lifecycle"),
					resource.TestCheckResourceAttr(acc.ResourceId, "lifecycle_hook_policy", "CONTINUE"),
					resource.TestCheckResourceAttr(acc.ResourceId, "lifecycle_hook_timeout", "30"),
					resource.TestCheckResourceAttr(acc.ResourceId, "lifecycle_hook_type", "SCALE_IN"),
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

func TestAccVolcengineScalingLifecycleHookResource_Update(t *testing.T) {
	resourceName := "volcengine_scaling_lifecycle_hook.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &scaling_lifecycle_hook.VolcengineScalingLifecycleHookService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineScalingLifecycleHookCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "lifecycle_hook_name", "acc-test-lifecycle"),
					resource.TestCheckResourceAttr(acc.ResourceId, "lifecycle_hook_policy", "CONTINUE"),
					resource.TestCheckResourceAttr(acc.ResourceId, "lifecycle_hook_timeout", "30"),
					resource.TestCheckResourceAttr(acc.ResourceId, "lifecycle_hook_type", "SCALE_IN"),
				),
			},
			{
				Config: testAccVolcengineScalingLifecycleHookUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "lifecycle_hook_name", "acc-test-lifecycle"),
					resource.TestCheckResourceAttr(acc.ResourceId, "lifecycle_hook_policy", "ROLLBACK"),
					resource.TestCheckResourceAttr(acc.ResourceId, "lifecycle_hook_timeout", "300"),
					resource.TestCheckResourceAttr(acc.ResourceId, "lifecycle_hook_type", "SCALE_OUT"),
					resource.TestCheckResourceAttr(acc.ResourceId, "lifecycle_command.#", "1"),
				),
			},
			{
				Config:             testAccVolcengineScalingLifecycleHookUpdateConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // 修改之后，不应该再产生diff
			},
		},
	})
}
