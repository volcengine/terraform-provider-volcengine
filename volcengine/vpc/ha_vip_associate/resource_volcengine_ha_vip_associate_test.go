package ha_vip_associate_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/ha_vip_associate"
)

const testAccVolcengineHaVipAssociateEcsCreateConfig = `
data "volcengine_zones" "foo"{
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet"
  cidr_block = "172.16.0.0/24"
  zone_id = data.volcengine_zones.foo.zones[0].id
  vpc_id = volcengine_vpc.foo.id
}

resource "volcengine_security_group" "foo" {
  security_group_name = "acc-test-sg"
  vpc_id = volcengine_vpc.foo.id
}

data "volcengine_images" "foo" {
  	os_type = "Linux"
  	visibility = "public"
  	instance_type_id = "ecs.g1.large"
}

resource "volcengine_ecs_instance" "foo" {
 	instance_name = "acc-test-ecs"
	description = "acc-test"
	host_name = "tf-acc-test"
  	image_id = "${data.volcengine_images.foo.images[0].image_id}"
  	instance_type = "ecs.g1.large"
  	password = "93f0cb0614Aab12"
  	instance_charge_type = "PostPaid"
  	system_volume_type = "ESSD_PL0"
  	system_volume_size = 40
	data_volumes {
    	volume_type = "ESSD_PL0"
    	size = 50
    	delete_with_instance = true
  	}
	subnet_id = "${volcengine_subnet.foo.id}"
	security_group_ids = ["${volcengine_security_group.foo.id}"]
	project_name = "default"
	tags {
    	key = "k1"
    	value = "v1"
  	}
}

resource "volcengine_ha_vip" "foo" {
  ha_vip_name = "acc-test-ha-vip"
  description = "acc-test"
  subnet_id = volcengine_subnet.foo.id
  ip_address = "172.16.0.5"
}

resource "volcengine_ha_vip_associate" "foo" {
  ha_vip_id = volcengine_ha_vip.foo.id
  instance_type = "EcsInstance"
  instance_id = volcengine_ecs_instance.foo.id
}
`

func TestAccVolcengineHaVipAssociateResource_BasicEcs(t *testing.T) {
	resourceName := "volcengine_ha_vip_associate.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return ha_vip_associate.NewHaVipAssociateService(client)
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineHaVipAssociateEcsCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_type", "EcsInstance"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "ha_vip_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "instance_id"),
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

const testAccVolcengineHaVipAssociateEniCreateConfig = `
data "volcengine_zones" "foo"{
}

resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
}

resource "volcengine_subnet" "foo" {
  subnet_name = "acc-test-subnet"
  cidr_block = "172.16.0.0/24"
  zone_id = data.volcengine_zones.foo.zones[0].id
  vpc_id = volcengine_vpc.foo.id
}

resource "volcengine_security_group" "foo" {
  security_group_name = "acc-test-sg"
  vpc_id = volcengine_vpc.foo.id
}

resource "volcengine_network_interface" "foo" {
  network_interface_name = "acc-test-eni"
  description = "acc-test"
  subnet_id = volcengine_subnet.foo.id
  security_group_ids = [volcengine_security_group.foo.id]
  port_security_enabled = false
  project_name = "default"
  tags {
    key = "k1"
    value = "v1"
  }
}

resource "volcengine_ha_vip" "foo" {
  ha_vip_name = "acc-test-ha-vip"
  description = "acc-test"
  subnet_id = volcengine_subnet.foo.id
  ip_address = "172.16.0.5"
}

resource "volcengine_ha_vip_associate" "foo" {
  ha_vip_id = volcengine_ha_vip.foo.id
  instance_type = "NetworkInterface"
  instance_id = volcengine_network_interface.foo.id
}
`

func TestAccVolcengineHaVipAssociateResource_BasicEni(t *testing.T) {
	resourceName := "volcengine_ha_vip_associate.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return ha_vip_associate.NewHaVipAssociateService(client)
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineHaVipAssociateEniCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_type", "NetworkInterface"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "ha_vip_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "instance_id"),
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
