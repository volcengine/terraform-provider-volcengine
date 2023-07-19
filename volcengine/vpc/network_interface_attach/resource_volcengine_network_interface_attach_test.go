package network_interface_attach_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/network_interface_attach"
)

const testAccVolcengineNetworkInterfaceAttachCreateConfig = `
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

resource "volcengine_network_interface" "foo" {
  network_interface_name = "acc-test-eni"
  description = "acc-test"
  subnet_id = "${volcengine_subnet.foo.id}"
  security_group_ids = ["${volcengine_security_group.foo.id}"]
  primary_ip_address = "172.16.0.253"
  port_security_enabled = false
  private_ip_address = ["172.16.0.2"]
  project_name = "default"
}

resource "volcengine_network_interface_attach" "foo" {
    instance_id = "${volcengine_ecs_instance.foo.id}"
    network_interface_id = "${volcengine_network_interface.foo.id}"
}
`

func TestAccVolcengineNetworkInterfaceAttachResource_Basic(t *testing.T) {
	resourceName := "volcengine_network_interface_attach.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &network_interface_attach.VolcengineNetworkInterfaceAttachService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineNetworkInterfaceAttachCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
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
