package ecs_instance_state_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/ecs_instance_state"
	"testing"
)

const testAccVolcengineEcsInstanceStateCreateConfig = `
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

resource "volcengine_security_group" "foo" {
  	security_group_name = "acc-test-security-group"
  	vpc_id = "${volcengine_vpc.foo.id}"
}

data "volcengine_images" "foo" {
  	os_type = "Linux"
  	visibility = "public"
  	instance_type_id = "ecs.g1.large"
}

resource "volcengine_ecs_instance" "foo" {
 	instance_name = "acc-test-ecs"
  	image_id = "${data.volcengine_images.foo.images[0].image_id}"
  	instance_type = "ecs.g1.large"
  	password = "93f0cb0614Aab12"
  	instance_charge_type = "PostPaid"
  	system_volume_type = "ESSD_PL0"
  	system_volume_size = 40
	subnet_id = "${volcengine_subnet.foo.id}"
	security_group_ids = ["${volcengine_security_group.foo.id}"]
}

resource "volcengine_ecs_instance_state" "foo" {
  	instance_id = "${volcengine_ecs_instance.foo.id}"
  	action = "Stop"
  	stopped_mode = "KeepCharging"
}
`

func TestAccVolcengineEcsInstanceStateResource_Basic(t *testing.T) {
	resourceName := "volcengine_ecs_instance_state.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &ecs_instance_state.VolcengineInstanceStateService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineEcsInstanceStateCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "stopped_mode", "KeepCharging"),
					resource.TestCheckResourceAttr(acc.ResourceId, "status", "STOPPED"),
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

const testAccVolcengineEcsInstanceStateUpdateConfig = `
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

resource "volcengine_security_group" "foo" {
  	security_group_name = "acc-test-security-group"
  	vpc_id = "${volcengine_vpc.foo.id}"
}

data "volcengine_images" "foo" {
  	os_type = "Linux"
  	visibility = "public"
  	instance_type_id = "ecs.g1.large"
}

resource "volcengine_ecs_instance" "foo" {
 	instance_name = "acc-test-ecs"
  	image_id = "${data.volcengine_images.foo.images[0].image_id}"
  	instance_type = "ecs.g1.large"
  	password = "93f0cb0614Aab12"
  	instance_charge_type = "PostPaid"
  	system_volume_type = "ESSD_PL0"
  	system_volume_size = 40
	subnet_id = "${volcengine_subnet.foo.id}"
	security_group_ids = ["${volcengine_security_group.foo.id}"]
}

resource "volcengine_ecs_instance_state" "foo" {
  	instance_id = "${volcengine_ecs_instance.foo.id}"
  	action = "Start"
}
`

func TestAccVolcengineEcsInstanceStateResource_Update(t *testing.T) {
	resourceName := "volcengine_ecs_instance_state.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &ecs_instance_state.VolcengineInstanceStateService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineEcsInstanceStateCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "stopped_mode", "KeepCharging"),
					resource.TestCheckResourceAttr(acc.ResourceId, "status", "STOPPED"),
				),
			},
			{
				Config: testAccVolcengineEcsInstanceStateUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "status", "RUNNING"),
				),
			},
			{
				Config:             testAccVolcengineEcsInstanceStateUpdateConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // 修改之后，不应该再产生diff
			},
		},
	})
}
