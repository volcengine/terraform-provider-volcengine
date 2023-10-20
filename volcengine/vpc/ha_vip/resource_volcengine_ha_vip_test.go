package ha_vip_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/ha_vip"
)

const testAccVolcengineHaVipCreateConfig = `
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

resource "volcengine_ha_vip" "foo" {
  ha_vip_name = "acc-test-ha-vip"
  description = "acc-test"
  subnet_id = volcengine_subnet.foo.id
  ip_address = "172.16.0.5"
}

resource "volcengine_eip_address" "foo" {
  billing_type = "PostPaidByTraffic"
}

resource "volcengine_eip_associate" "foo" {
  allocation_id = volcengine_eip_address.foo.id
  instance_id = volcengine_ha_vip.foo.id
  instance_type = "HaVip"
}
`

func TestAccVolcengineHaVipResource_Basic(t *testing.T) {
	resourceName := "volcengine_ha_vip.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return ha_vip.NewHaVipService(client)
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
				Config: testAccVolcengineHaVipCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "ha_vip_name", "acc-test-ha-vip"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test"),
					resource.TestCheckResourceAttr(acc.ResourceId, "ip_address", "172.16.0.5"),
					resource.TestCheckResourceAttr(acc.ResourceId, "project_name", "default"),
					resource.TestCheckResourceAttr(acc.ResourceId, "status", "Available"),
					resource.TestCheckResourceAttr(acc.ResourceId, "associated_instance_ids.#", "0"),
					resource.TestCheckResourceAttr(acc.ResourceId, "associated_instance_type", ""),
					resource.TestCheckResourceAttr(acc.ResourceId, "associated_eip_address", ""),
					resource.TestCheckResourceAttr(acc.ResourceId, "associated_eip_id", ""),
					resource.TestCheckResourceAttr(acc.ResourceId, "master_instance_id", ""),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "vpc_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "subnet_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "created_at"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "updated_at"),
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

const testAccVolcengineHaVipUpdateConfig = `
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

resource "volcengine_ha_vip" "foo" {
  ha_vip_name = "acc-test-ha-vip-new"
  description = "acc-test-new"
  subnet_id = volcengine_subnet.foo.id
  ip_address = "172.16.0.5"
}

resource "volcengine_eip_address" "foo" {
  billing_type = "PostPaidByTraffic"
}

resource "volcengine_eip_associate" "foo" {
  allocation_id = volcengine_eip_address.foo.id
  instance_id = volcengine_ha_vip.foo.id
  instance_type = "HaVip"
}
`

func TestAccVolcengineHaVipResource_Update(t *testing.T) {
	resourceName := "volcengine_ha_vip.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return ha_vip.NewHaVipService(client)
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
				Config: testAccVolcengineHaVipCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "ha_vip_name", "acc-test-ha-vip"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test"),
					resource.TestCheckResourceAttr(acc.ResourceId, "ip_address", "172.16.0.5"),
					resource.TestCheckResourceAttr(acc.ResourceId, "project_name", "default"),
					resource.TestCheckResourceAttr(acc.ResourceId, "status", "Available"),
					resource.TestCheckResourceAttr(acc.ResourceId, "associated_instance_ids.#", "0"),
					resource.TestCheckResourceAttr(acc.ResourceId, "associated_instance_type", ""),
					resource.TestCheckResourceAttr(acc.ResourceId, "associated_eip_address", ""),
					resource.TestCheckResourceAttr(acc.ResourceId, "associated_eip_id", ""),
					resource.TestCheckResourceAttr(acc.ResourceId, "master_instance_id", ""),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "vpc_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "subnet_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "created_at"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "updated_at"),
				),
			},
			{
				Config: testAccVolcengineHaVipUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "ha_vip_name", "acc-test-ha-vip-new"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test-new"),
					resource.TestCheckResourceAttr(acc.ResourceId, "ip_address", "172.16.0.5"),
					resource.TestCheckResourceAttr(acc.ResourceId, "project_name", "default"),
					resource.TestCheckResourceAttr(acc.ResourceId, "status", "Available"),
					resource.TestCheckResourceAttr(acc.ResourceId, "associated_instance_ids.#", "0"),
					resource.TestCheckResourceAttr(acc.ResourceId, "associated_instance_type", ""),
					resource.TestCheckResourceAttr(acc.ResourceId, "master_instance_id", ""),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "vpc_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "subnet_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "created_at"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "updated_at"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "associated_eip_address"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "associated_eip_id"),
				),
			},
			{
				Config:             testAccVolcengineHaVipUpdateConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // 修改之后，不应该再产生diff
			},
		},
	})
}
