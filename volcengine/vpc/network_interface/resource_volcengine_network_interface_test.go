package network_interface_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/network_interface"
)

const testAccVolcengineNetworkInterfaceCreateConfig = `
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
  security_group_name = "acc-test-sg"
  vpc_id = "${volcengine_vpc.foo.id}"
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
  tags {
    key = "k1"
    value = "v1"
  }
}
`

func TestAccVolcengineNetworkInterfaceResource_Basic(t *testing.T) {
	resourceName := "volcengine_network_interface.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &network_interface.VolcengineNetworkInterfaceService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineNetworkInterfaceCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "network_interface_name", "acc-test-eni"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test"),
					resource.TestCheckResourceAttr(acc.ResourceId, "port_security_enabled", "false"),
					resource.TestCheckResourceAttr(acc.ResourceId, "primary_ip_address", "172.16.0.253"),
					resource.TestCheckResourceAttr(acc.ResourceId, "project_name", "default"),
					resource.TestCheckResourceAttr(acc.ResourceId, "security_group_ids.#", "1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "status", "Available"),
					resource.TestCheckResourceAttr(acc.ResourceId, "private_ip_address.#", "1"),
					volcengine.TestCheckTypeSetElemAttr(acc.ResourceId, "private_ip_address.*", "172.16.0.2"),
					resource.TestCheckResourceAttr(acc.ResourceId, "tags.#", "1"),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "tags.*", map[string]string{
						"key":   "k1",
						"value": "v1",
					}),
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

const testAccVolcengineNetworkInterfaceUpdateConfig = `
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
  security_group_name = "acc-test-sg"
  vpc_id = "${volcengine_vpc.foo.id}"
}

resource "volcengine_network_interface" "foo" {
  network_interface_name = "acc-test-eni-new"
  description = "acc-test-new"
  subnet_id = "${volcengine_subnet.foo.id}"
  security_group_ids = ["${volcengine_security_group.foo.id}"]
  primary_ip_address = "172.16.0.253"
  port_security_enabled = false
  private_ip_address = ["172.16.0.2", "172.16.0.3"]
  project_name = "default"
  tags {
    key = "k1"
    value = "v1"
  }
  tags {
    key = "k2"
    value = "v2"
  }
}
`

func TestAccVolcengineNetworkInterfaceResource_Update(t *testing.T) {
	resourceName := "volcengine_network_interface.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &network_interface.VolcengineNetworkInterfaceService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineNetworkInterfaceCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "network_interface_name", "acc-test-eni"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test"),
					resource.TestCheckResourceAttr(acc.ResourceId, "port_security_enabled", "false"),
					resource.TestCheckResourceAttr(acc.ResourceId, "primary_ip_address", "172.16.0.253"),
					resource.TestCheckResourceAttr(acc.ResourceId, "project_name", "default"),
					resource.TestCheckResourceAttr(acc.ResourceId, "security_group_ids.#", "1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "status", "Available"),
					resource.TestCheckResourceAttr(acc.ResourceId, "private_ip_address.#", "1"),
					volcengine.TestCheckTypeSetElemAttr(acc.ResourceId, "private_ip_address.*", "172.16.0.2"),
					resource.TestCheckResourceAttr(acc.ResourceId, "tags.#", "1"),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "tags.*", map[string]string{
						"key":   "k1",
						"value": "v1",
					}),
				),
			},
			{
				Config: testAccVolcengineNetworkInterfaceUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "network_interface_name", "acc-test-eni-new"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test-new"),
					resource.TestCheckResourceAttr(acc.ResourceId, "port_security_enabled", "false"),
					resource.TestCheckResourceAttr(acc.ResourceId, "primary_ip_address", "172.16.0.253"),
					resource.TestCheckResourceAttr(acc.ResourceId, "project_name", "default"),
					resource.TestCheckResourceAttr(acc.ResourceId, "security_group_ids.#", "1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "status", "Available"),
					resource.TestCheckResourceAttr(acc.ResourceId, "private_ip_address.#", "2"),
					volcengine.TestCheckTypeSetElemAttr(acc.ResourceId, "private_ip_address.*", "172.16.0.2"),
					volcengine.TestCheckTypeSetElemAttr(acc.ResourceId, "private_ip_address.*", "172.16.0.3"),
					resource.TestCheckResourceAttr(acc.ResourceId, "tags.#", "2"),
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
				Config:             testAccVolcengineNetworkInterfaceUpdateConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // 修改之后，不应该再产生diff
			},
		},
	})
}
