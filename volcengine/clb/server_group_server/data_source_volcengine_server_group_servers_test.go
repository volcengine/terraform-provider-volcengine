package server_group_server_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/clb/server_group_server"
	"testing"
)

const testAccVolcengineServerGroupServersDatasourceConfig = `
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

resource "volcengine_clb" "foo" {
  type = "public"
  subnet_id = "${volcengine_subnet.foo.id}"
  load_balancer_spec = "small_1"
  description = "acc0Demo"
  load_balancer_name = "acc-test-create"
  eip_billing_config {
    isp = "BGP"
    eip_billing_type = "PostPaidByBandwidth"
    bandwidth = 1
  }
}

resource "volcengine_server_group" "foo" {
  load_balancer_id = "${volcengine_clb.foo.id}"
  server_group_name = "acc-test-create"
  description = "hello demo11"
}

resource "volcengine_security_group" "foo" {
	  vpc_id = "${volcengine_vpc.foo.id}"
	  security_group_name = "acc-test-security-group"
}

resource "volcengine_ecs_instance" "foo" {
	  image_id = "image-ycjwwciuzy5pkh54xx8f"
	  instance_type = "ecs.c3i.large"
	  instance_name = "acc-test-ecs-name"
	  password = "93f0cb0614Aab12"
	  instance_charge_type = "PostPaid"
	  system_volume_type = "ESSD_PL0"
	  system_volume_size = 40
	  subnet_id = volcengine_subnet.foo.id
	  security_group_ids = [volcengine_security_group.foo.id]
}

resource "volcengine_server_group_server" "foo" {
  server_group_id = "${volcengine_server_group.foo.id}"
  instance_id = "${volcengine_ecs_instance.foo.id}"
  type = "ecs"
  weight = 100
  port = 80
  description = "This is a acc test server"
}

data "volcengine_server_group_servers" "foo"{
    ids = [element(split(":", volcengine_server_group_server.foo.id), length(split(":", volcengine_server_group_server.foo.id))-1)]
	server_group_id = "${volcengine_server_group.foo.id}"
}
`

func TestAccVolcengineServerGroupServersDatasource_Basic(t *testing.T) {
	resourceName := "data.volcengine_server_group_servers.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &server_group_server.VolcengineServerGroupServerService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers: volcengine.GetTestAccProviders(),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineServerGroupServersDatasourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(acc.ResourceId, "servers.#", "1"),
				),
			},
		},
	})
}
