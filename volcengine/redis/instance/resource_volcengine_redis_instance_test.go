package instance_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/redis/instance"
	"testing"
)

const testAccVolcengineRedisInstanceCreateConfig = `
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

resource "volcengine_redis_instance" "foo"{
     zone_ids = ["${data.volcengine_zones.foo.zones[0].id}"]
     instance_name = "acc-test-tf-redis"
     sharded_cluster = 1
     password = "1qaz!QAZ12"
     node_number = 2
     shard_capacity = 1024
     shard_number = 2
     engine_version = "5.0"
     subnet_id = "${volcengine_subnet.foo.id}"
     deletion_protection = "disabled"
     vpc_auth_mode = "close"
     charge_type = "PostPaid"
     port = 6381
     project_name = "default"
}
`

const testAccVolcengineRedisInstanceUpdateConfig = `
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

resource "volcengine_redis_instance" "foo"{
     zone_ids = ["${data.volcengine_zones.foo.zones[0].id}"]
     instance_name = "acc-test-tf-redis-new"
     sharded_cluster = 1
     password = "1qaz!QAZ12"
     node_number = 2
     shard_capacity = 1024
     shard_number = 2
     engine_version = "5.0"
     subnet_id = "${volcengine_subnet.foo.id}"
     deletion_protection = "disabled"
     vpc_auth_mode = "close"
     charge_type = "PostPaid"
     port = 6381
     project_name = "default"
}
`

func TestAccVolcengineRedisInstanceResource_Basic(t *testing.T) {
	resourceName := "volcengine_redis_instance.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &instance.VolcengineRedisDbInstanceService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineRedisInstanceCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "charge_type", "PostPaid"),
					resource.TestCheckResourceAttr(acc.ResourceId, "deletion_protection", "disabled"),
					resource.TestCheckResourceAttr(acc.ResourceId, "engine_version", "5.0"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_name", "acc-test-tf-redis"),
					resource.TestCheckResourceAttr(acc.ResourceId, "node_number", "2"),
					resource.TestCheckResourceAttr(acc.ResourceId, "port", "6381"),
					resource.TestCheckResourceAttr(acc.ResourceId, "project_name", "default"),
					resource.TestCheckResourceAttr(acc.ResourceId, "shard_capacity", "1024"),
					resource.TestCheckResourceAttr(acc.ResourceId, "shard_number", "2"),
					resource.TestCheckResourceAttr(acc.ResourceId, "sharded_cluster", "1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "vpc_auth_mode", "close"),
					resource.TestCheckResourceAttr(acc.ResourceId, "zone_ids.#", "1"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"password"},
			},
		},
	})
}

func TestAccVolcengineRedisInstanceResource_Update(t *testing.T) {
	resourceName := "volcengine_redis_instance.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &instance.VolcengineRedisDbInstanceService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineRedisInstanceCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "charge_type", "PostPaid"),
					resource.TestCheckResourceAttr(acc.ResourceId, "deletion_protection", "disabled"),
					resource.TestCheckResourceAttr(acc.ResourceId, "engine_version", "5.0"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_name", "acc-test-tf-redis"),
					resource.TestCheckResourceAttr(acc.ResourceId, "node_number", "2"),
					resource.TestCheckResourceAttr(acc.ResourceId, "port", "6381"),
					resource.TestCheckResourceAttr(acc.ResourceId, "project_name", "default"),
					resource.TestCheckResourceAttr(acc.ResourceId, "shard_capacity", "1024"),
					resource.TestCheckResourceAttr(acc.ResourceId, "shard_number", "2"),
					resource.TestCheckResourceAttr(acc.ResourceId, "sharded_cluster", "1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "vpc_auth_mode", "close"),
					resource.TestCheckResourceAttr(acc.ResourceId, "zone_ids.#", "1"),
				),
			},
			{
				Config: testAccVolcengineRedisInstanceUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "charge_type", "PostPaid"),
					resource.TestCheckResourceAttr(acc.ResourceId, "deletion_protection", "disabled"),
					resource.TestCheckResourceAttr(acc.ResourceId, "engine_version", "5.0"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_name", "acc-test-tf-redis-new"),
					resource.TestCheckResourceAttr(acc.ResourceId, "node_number", "2"),
					resource.TestCheckResourceAttr(acc.ResourceId, "port", "6381"),
					resource.TestCheckResourceAttr(acc.ResourceId, "project_name", "default"),
					resource.TestCheckResourceAttr(acc.ResourceId, "shard_capacity", "1024"),
					resource.TestCheckResourceAttr(acc.ResourceId, "shard_number", "2"),
					resource.TestCheckResourceAttr(acc.ResourceId, "sharded_cluster", "1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "vpc_auth_mode", "close"),
					resource.TestCheckResourceAttr(acc.ResourceId, "zone_ids.#", "1"),
				),
			},
			{
				Config:             testAccVolcengineRedisInstanceUpdateConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // 修改之后，不应该再产生diff
			},
		},
	})
}
