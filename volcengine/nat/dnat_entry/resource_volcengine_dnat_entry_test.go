package dnat_entry_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/nat/dnat_entry"
	"testing"
)

const testAccVolcengineDnatEntryCreateConfig = `
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

resource "volcengine_nat_gateway" "foo" {
	vpc_id = "${volcengine_vpc.foo.id}"
    subnet_id = "${volcengine_subnet.foo.id}"
	spec = "Small"
	nat_gateway_name = "acc-test-ng"
	description = "acc-test"
	billing_type = "PostPaid"
	project_name = "default"
	tags {
		key = "k1"
		value = "v1"
	}
}

resource "volcengine_eip_address" "foo" {
	name = "acc-test-eip"
    description = "acc-test"
    bandwidth = 1
    billing_type = "PostPaidByBandwidth"
    isp = "BGP"
}

resource "volcengine_eip_associate" "foo" {
	allocation_id = "${volcengine_eip_address.foo.id}"
	instance_id = "${volcengine_nat_gateway.foo.id}"
	instance_type = "Nat"
}

resource "volcengine_dnat_entry" "foo" {
	dnat_entry_name = "acc-test-dnat-entry"
    external_ip = "${volcengine_eip_address.foo.eip_address}"
    external_port = 80
    internal_ip = "172.16.0.10"
    internal_port = 80
    nat_gateway_id = "${volcengine_nat_gateway.foo.id}"
    protocol = "tcp"
	depends_on = [volcengine_eip_associate.foo]
}
`

func TestAccVolcengineDnatEntryResource_Basic(t *testing.T) {
	resourceName := "volcengine_dnat_entry.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &dnat_entry.VolcengineDnatEntryService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineDnatEntryCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "dnat_entry_name", "acc-test-dnat-entry"),
					resource.TestCheckResourceAttr(acc.ResourceId, "external_port", "80"),
					resource.TestCheckResourceAttr(acc.ResourceId, "internal_ip", "172.16.0.10"),
					resource.TestCheckResourceAttr(acc.ResourceId, "internal_port", "80"),
					resource.TestCheckResourceAttr(acc.ResourceId, "protocol", "tcp"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "external_ip"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "nat_gateway_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "dnat_entry_id"),
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

const testAccVolcengineDnatEntryUpdateConfig = `
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

resource "volcengine_nat_gateway" "foo" {
	vpc_id = "${volcengine_vpc.foo.id}"
    subnet_id = "${volcengine_subnet.foo.id}"
	spec = "Small"
	nat_gateway_name = "acc-test-ng"
	description = "acc-test"
	billing_type = "PostPaid"
	project_name = "default"
	tags {
		key = "k1"
		value = "v1"
	}
}

resource "volcengine_eip_address" "foo" {
	name = "acc-test-eip"
    description = "acc-test"
    bandwidth = 1
    billing_type = "PostPaidByBandwidth"
    isp = "BGP"
}

resource "volcengine_eip_associate" "foo" {
	allocation_id = "${volcengine_eip_address.foo.id}"
	instance_id = "${volcengine_nat_gateway.foo.id}"
	instance_type = "Nat"
}

resource "volcengine_eip_address" "foo1" {
	name = "acc-test-eip"
    description = "acc-test"
    bandwidth = 1
    billing_type = "PostPaidByBandwidth"
    isp = "BGP"
}

resource "volcengine_eip_associate" "foo1" {
	allocation_id = "${volcengine_eip_address.foo1.id}"
	instance_id = "${volcengine_nat_gateway.foo.id}"
	instance_type = "Nat"
}

resource "volcengine_dnat_entry" "foo" {
	dnat_entry_name = "acc-test-dnat-entry-new"
    external_ip = "${volcengine_eip_address.foo1.eip_address}"
    external_port = 90
    internal_ip = "172.16.0.17"
    internal_port = 90
    nat_gateway_id = "${volcengine_nat_gateway.foo.id}"
    protocol = "udp"
	depends_on = [volcengine_eip_associate.foo1]
}
`

func TestAccVolcengineDnatEntryResource_Update(t *testing.T) {
	resourceName := "volcengine_dnat_entry.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &dnat_entry.VolcengineDnatEntryService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineDnatEntryCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "dnat_entry_name", "acc-test-dnat-entry"),
					resource.TestCheckResourceAttr(acc.ResourceId, "external_port", "80"),
					resource.TestCheckResourceAttr(acc.ResourceId, "internal_ip", "172.16.0.10"),
					resource.TestCheckResourceAttr(acc.ResourceId, "internal_port", "80"),
					resource.TestCheckResourceAttr(acc.ResourceId, "protocol", "tcp"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "external_ip"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "nat_gateway_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "dnat_entry_id"),
				),
			},
			{
				Config: testAccVolcengineDnatEntryUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "dnat_entry_name", "acc-test-dnat-entry-new"),
					resource.TestCheckResourceAttr(acc.ResourceId, "external_port", "90"),
					resource.TestCheckResourceAttr(acc.ResourceId, "internal_ip", "172.16.0.17"),
					resource.TestCheckResourceAttr(acc.ResourceId, "internal_port", "90"),
					resource.TestCheckResourceAttr(acc.ResourceId, "protocol", "udp"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "external_ip"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "nat_gateway_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "dnat_entry_id"),
				),
			},
			{
				Config:             testAccVolcengineDnatEntryUpdateConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // 修改之后，不应该再产生diff
			},
		},
	})
}
