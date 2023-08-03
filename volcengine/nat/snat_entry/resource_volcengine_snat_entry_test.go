package snat_entry_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/nat/snat_entry"
	"testing"
)

const testAccVolcengineSnatEntryCreateConfig = `
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

resource "volcengine_snat_entry" "foo" {
	snat_entry_name = "acc-test-snat-entry"
    nat_gateway_id = "${volcengine_nat_gateway.foo.id}"
	eip_id = "${volcengine_eip_address.foo.id}"
	subnet_id = "${volcengine_subnet.foo.id}"
	depends_on = [volcengine_eip_associate.foo]
}
`

func TestAccVolcengineSnatEntryResource_Basic(t *testing.T) {
	resourceName := "volcengine_snat_entry.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &snat_entry.VolcengineSnatEntryService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineSnatEntryCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "snat_entry_name", "acc-test-snat-entry"),
					resource.TestCheckResourceAttr(acc.ResourceId, "status", "Available"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "eip_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "nat_gateway_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "subnet_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "source_cidr"),
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

const testAccVolcengineSnatEntryCreateSourceCidrConfig = `
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

resource "volcengine_snat_entry" "foo" {
	snat_entry_name = "acc-test-snat-entry"
    nat_gateway_id = "${volcengine_nat_gateway.foo.id}"
	eip_id = "${volcengine_eip_address.foo.id}"
	source_cidr = "172.16.0.0/24"
	depends_on = [volcengine_eip_associate.foo]
}
`

func TestAccVolcengineSnatEntryResource_SourceCidr(t *testing.T) {
	resourceName := "volcengine_snat_entry.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &snat_entry.VolcengineSnatEntryService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineSnatEntryCreateSourceCidrConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "snat_entry_name", "acc-test-snat-entry"),
					resource.TestCheckResourceAttr(acc.ResourceId, "status", "Available"),
					resource.TestCheckResourceAttr(acc.ResourceId, "source_cidr", "172.16.0.0/24"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "eip_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "nat_gateway_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "subnet_id"),
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

const testAccVolcengineSnatEntryUpdateConfig = `
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
	name = "acc-test-eip1"
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

resource "volcengine_snat_entry" "foo" {
	snat_entry_name = "acc-test-snat-entry-new"
    nat_gateway_id = "${volcengine_nat_gateway.foo.id}"
	eip_id = "${volcengine_eip_address.foo1.id}"
	subnet_id = "${volcengine_subnet.foo.id}"
	depends_on = [volcengine_eip_associate.foo1]
}
`

func TestAccVolcengineSnatEntryResource_Update(t *testing.T) {
	resourceName := "volcengine_snat_entry.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &snat_entry.VolcengineSnatEntryService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineSnatEntryCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "snat_entry_name", "acc-test-snat-entry"),
					resource.TestCheckResourceAttr(acc.ResourceId, "status", "Available"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "eip_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "nat_gateway_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "subnet_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "source_cidr"),
				),
			},
			{
				Config: testAccVolcengineSnatEntryUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "snat_entry_name", "acc-test-snat-entry-new"),
					resource.TestCheckResourceAttr(acc.ResourceId, "status", "Available"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "eip_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "nat_gateway_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "subnet_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "source_cidr"),
				),
			},
			{
				Config:             testAccVolcengineSnatEntryUpdateConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // 修改之后，不应该再产生diff
			},
		},
	})
}
