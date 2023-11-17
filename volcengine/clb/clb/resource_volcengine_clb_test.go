package clb_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/clb/clb"
)

const testAccVolcengineClbCreateConfig = `
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
  	description = "acc-test-demo"
  	load_balancer_name = "acc-test-clb"
	load_balancer_billing_type = "PostPaid"
  	eip_billing_config {
    	isp = "BGP"
    	eip_billing_type = "PostPaidByBandwidth"
    	bandwidth = 1
  	}
	tags {
		key = "k1"
		value = "v1"
	}
}
`

func TestAccVolcengineClbResource_Basic(t *testing.T) {
	resourceName := "volcengine_clb.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &clb.VolcengineClbService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineClbCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "load_balancer_name", "acc-test-clb"),
					resource.TestCheckResourceAttr(acc.ResourceId, "load_balancer_spec", "small_1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "load_balancer_billing_type", "PostPaid"),
					resource.TestCheckResourceAttr(acc.ResourceId, "type", "public"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test-demo"),
					resource.TestCheckResourceAttr(acc.ResourceId, "modification_protection_reason", ""),
					resource.TestCheckResourceAttr(acc.ResourceId, "modification_protection_status", ""),
					resource.TestCheckResourceAttr(acc.ResourceId, "project_name", "default"),
					resource.TestCheckResourceAttr(acc.ResourceId, "eip_billing_config.#", "1"),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "eip_billing_config.*", map[string]string{
						"isp":              "BGP",
						"eip_billing_type": "PostPaidByBandwidth",
						"bandwidth":        "1",
					}),
					resource.TestCheckResourceAttr(acc.ResourceId, "tags.#", "1"),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "tags.*", map[string]string{
						"key":   "k1",
						"value": "v1",
					}),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "vpc_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "subnet_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "eip_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "eip_address"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "eni_address"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "master_zone_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "region_id"),
					resource.TestCheckNoResourceAttr(acc.ResourceId, "period"),
					resource.TestCheckNoResourceAttr(acc.ResourceId, "renew_type"),
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

const testAccVolcengineClbUpdateBasicAttributeConfig = `
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
  	load_balancer_spec = "small_2"
  	description = "acc-test-demo-new"
  	load_balancer_name = "acc-test-clb-new"
	load_balancer_billing_type = "PostPaid"
	modification_protection_status = "ConsoleProtection"
	modification_protection_reason = "reason"
  	eip_billing_config {
    	isp = "BGP"
    	eip_billing_type = "PostPaidByBandwidth"
    	bandwidth = 1
  	}
	tags {
		key = "k1"
		value = "v1"
	}
}
`

func TestAccVolcengineClbResource_UpdateBasicAttribute(t *testing.T) {
	resourceName := "volcengine_clb.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &clb.VolcengineClbService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineClbCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "load_balancer_name", "acc-test-clb"),
					resource.TestCheckResourceAttr(acc.ResourceId, "load_balancer_spec", "small_1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "load_balancer_billing_type", "PostPaid"),
					resource.TestCheckResourceAttr(acc.ResourceId, "type", "public"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test-demo"),
					resource.TestCheckResourceAttr(acc.ResourceId, "modification_protection_reason", ""),
					resource.TestCheckResourceAttr(acc.ResourceId, "modification_protection_status", ""),
					resource.TestCheckResourceAttr(acc.ResourceId, "project_name", ""),
					resource.TestCheckResourceAttr(acc.ResourceId, "eip_billing_config.#", "1"),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "eip_billing_config.*", map[string]string{
						"isp":              "BGP",
						"eip_billing_type": "PostPaidByBandwidth",
						"bandwidth":        "1",
					}),
					resource.TestCheckResourceAttr(acc.ResourceId, "tags.#", "1"),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "tags.*", map[string]string{
						"key":   "k1",
						"value": "v1",
					}),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "vpc_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "subnet_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "eip_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "eip_address"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "eni_address"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "master_zone_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "region_id"),
					resource.TestCheckNoResourceAttr(acc.ResourceId, "period"),
					resource.TestCheckNoResourceAttr(acc.ResourceId, "renew_type"),
				),
			},
			{
				Config: testAccVolcengineClbUpdateBasicAttributeConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "load_balancer_name", "acc-test-clb-new"),
					resource.TestCheckResourceAttr(acc.ResourceId, "load_balancer_spec", "small_2"),
					resource.TestCheckResourceAttr(acc.ResourceId, "load_balancer_billing_type", "PostPaid"),
					resource.TestCheckResourceAttr(acc.ResourceId, "type", "public"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test-demo-new"),
					resource.TestCheckResourceAttr(acc.ResourceId, "modification_protection_status", "ConsoleProtection"),
					resource.TestCheckResourceAttr(acc.ResourceId, "modification_protection_reason", "reason"),
					resource.TestCheckResourceAttr(acc.ResourceId, "project_name", ""),
					resource.TestCheckResourceAttr(acc.ResourceId, "eip_billing_config.#", "1"),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "eip_billing_config.*", map[string]string{
						"isp":              "BGP",
						"eip_billing_type": "PostPaidByBandwidth",
						"bandwidth":        "1",
					}),
					resource.TestCheckResourceAttr(acc.ResourceId, "tags.#", "1"),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "tags.*", map[string]string{
						"key":   "k1",
						"value": "v1",
					}),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "vpc_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "subnet_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "eip_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "eip_address"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "eni_address"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "master_zone_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "region_id"),
					resource.TestCheckNoResourceAttr(acc.ResourceId, "period"),
					resource.TestCheckNoResourceAttr(acc.ResourceId, "renew_type"),
				),
			},
			{
				Config:             testAccVolcengineClbUpdateBasicAttributeConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // 修改之后，不应该再产生diff
			},
		},
	})
}

const testAccVolcengineClbUpdateBillingTypeConfig1 = `
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
  	description = "acc-test-demo"
  	load_balancer_name = "acc-test-clb"
	load_balancer_billing_type = "PrePaid"
	period = 1
  	eip_billing_config {
    	isp = "BGP"
    	eip_billing_type = "PostPaidByBandwidth"
    	bandwidth = 1
  	}
	tags {
		key = "k1"
		value = "v1"
	}
}
`

const testAccVolcengineClbUpdateBillingTypeConfig2 = `
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
  	description = "acc-test-demo"
  	load_balancer_name = "acc-test-clb"
	load_balancer_billing_type = "PrePaid"
	period = 2
  	eip_billing_config {
    	isp = "BGP"
    	eip_billing_type = "PrePaid"
    	bandwidth = 1
  	}
	tags {
		key = "k1"
		value = "v1"
	}
}
`

func TestAccVolcengineClbResource_UpdateBillingType(t *testing.T) {
	resourceName := "volcengine_clb.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &clb.VolcengineClbService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineClbCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "load_balancer_name", "acc-test-clb"),
					resource.TestCheckResourceAttr(acc.ResourceId, "load_balancer_spec", "small_1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "load_balancer_billing_type", "PostPaid"),
					resource.TestCheckResourceAttr(acc.ResourceId, "type", "public"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test-demo"),
					resource.TestCheckResourceAttr(acc.ResourceId, "modification_protection_reason", ""),
					resource.TestCheckResourceAttr(acc.ResourceId, "modification_protection_status", ""),
					resource.TestCheckResourceAttr(acc.ResourceId, "project_name", ""),
					resource.TestCheckResourceAttr(acc.ResourceId, "eip_billing_config.#", "1"),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "eip_billing_config.*", map[string]string{
						"isp":              "BGP",
						"eip_billing_type": "PostPaidByBandwidth",
						"bandwidth":        "1",
					}),
					resource.TestCheckResourceAttr(acc.ResourceId, "tags.#", "1"),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "tags.*", map[string]string{
						"key":   "k1",
						"value": "v1",
					}),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "vpc_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "subnet_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "eip_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "eip_address"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "eni_address"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "master_zone_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "region_id"),
					resource.TestCheckNoResourceAttr(acc.ResourceId, "period"),
					resource.TestCheckNoResourceAttr(acc.ResourceId, "renew_type"),
				),
			},
			{
				Config:             testAccVolcengineClbUpdateBillingTypeConfig1,
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "load_balancer_name", "acc-test-clb"),
					resource.TestCheckResourceAttr(acc.ResourceId, "load_balancer_spec", "small_1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "load_balancer_billing_type", "PrePaid"),
					resource.TestCheckResourceAttr(acc.ResourceId, "period", "1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "renew_type", "ManualRenew"),
					resource.TestCheckResourceAttr(acc.ResourceId, "type", "public"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test-demo"),
					resource.TestCheckResourceAttr(acc.ResourceId, "modification_protection_reason", ""),
					resource.TestCheckResourceAttr(acc.ResourceId, "modification_protection_status", ""),
					resource.TestCheckResourceAttr(acc.ResourceId, "project_name", ""),
					resource.TestCheckResourceAttr(acc.ResourceId, "eip_billing_config.#", "1"),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "eip_billing_config.*", map[string]string{
						"isp":              "BGP",
						"eip_billing_type": "PrePaid",
						"bandwidth":        "1",
					}),
					resource.TestCheckResourceAttr(acc.ResourceId, "tags.#", "1"),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "tags.*", map[string]string{
						"key":   "k1",
						"value": "v1",
					}),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "vpc_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "subnet_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "eip_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "eip_address"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "eni_address"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "master_zone_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "region_id"),
				),
			},
			{
				Config: testAccVolcengineClbUpdateBillingTypeConfig2,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "load_balancer_name", "acc-test-clb"),
					resource.TestCheckResourceAttr(acc.ResourceId, "load_balancer_spec", "small_1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "load_balancer_billing_type", "PrePaid"),
					resource.TestCheckResourceAttr(acc.ResourceId, "period", "2"),
					resource.TestCheckResourceAttr(acc.ResourceId, "renew_type", "ManualRenew"),
					resource.TestCheckResourceAttr(acc.ResourceId, "type", "public"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test-demo"),
					resource.TestCheckResourceAttr(acc.ResourceId, "modification_protection_reason", ""),
					resource.TestCheckResourceAttr(acc.ResourceId, "modification_protection_status", ""),
					resource.TestCheckResourceAttr(acc.ResourceId, "project_name", ""),
					resource.TestCheckResourceAttr(acc.ResourceId, "eip_billing_config.#", "1"),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "eip_billing_config.*", map[string]string{
						"isp":              "BGP",
						"eip_billing_type": "PrePaid",
						"bandwidth":        "1",
					}),
					resource.TestCheckResourceAttr(acc.ResourceId, "tags.#", "1"),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "tags.*", map[string]string{
						"key":   "k1",
						"value": "v1",
					}),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "vpc_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "subnet_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "eip_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "eip_address"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "eni_address"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "master_zone_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "region_id"),
				),
			},
			{
				Config:             testAccVolcengineClbUpdateBillingTypeConfig2,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // 修改之后，不应该再产生diff
			},
		},
	})
}

const testAccVolcengineClbUpdateTagsConfig = `
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
  	description = "acc-test-demo"
  	load_balancer_name = "acc-test-clb"
	load_balancer_billing_type = "PostPaid"
  	eip_billing_config {
    	isp = "BGP"
    	eip_billing_type = "PostPaidByBandwidth"
    	bandwidth = 1
  	}
	tags {
		key = "k2"
		value = "v2"
	}
	tags {
		key = "k3"
		value = "v3"
	}
}
`

func TestAccVolcengineClbResource_UpdateTags(t *testing.T) {
	resourceName := "volcengine_clb.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &clb.VolcengineClbService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineClbCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "load_balancer_name", "acc-test-clb"),
					resource.TestCheckResourceAttr(acc.ResourceId, "load_balancer_spec", "small_1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "load_balancer_billing_type", "PostPaid"),
					resource.TestCheckResourceAttr(acc.ResourceId, "type", "public"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test-demo"),
					resource.TestCheckResourceAttr(acc.ResourceId, "modification_protection_reason", ""),
					resource.TestCheckResourceAttr(acc.ResourceId, "modification_protection_status", ""),
					resource.TestCheckResourceAttr(acc.ResourceId, "project_name", ""),
					resource.TestCheckResourceAttr(acc.ResourceId, "eip_billing_config.#", "1"),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "eip_billing_config.*", map[string]string{
						"isp":              "BGP",
						"eip_billing_type": "PostPaidByBandwidth",
						"bandwidth":        "1",
					}),
					resource.TestCheckResourceAttr(acc.ResourceId, "tags.#", "1"),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "tags.*", map[string]string{
						"key":   "k1",
						"value": "v1",
					}),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "vpc_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "subnet_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "eip_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "eip_address"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "eni_address"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "master_zone_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "region_id"),
					resource.TestCheckNoResourceAttr(acc.ResourceId, "period"),
					resource.TestCheckNoResourceAttr(acc.ResourceId, "renew_type"),
				),
			},
			{
				Config: testAccVolcengineClbUpdateTagsConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "load_balancer_name", "acc-test-clb"),
					resource.TestCheckResourceAttr(acc.ResourceId, "load_balancer_spec", "small_1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "load_balancer_billing_type", "PostPaid"),
					resource.TestCheckResourceAttr(acc.ResourceId, "type", "public"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test-demo"),
					resource.TestCheckResourceAttr(acc.ResourceId, "modification_protection_reason", ""),
					resource.TestCheckResourceAttr(acc.ResourceId, "modification_protection_status", ""),
					resource.TestCheckResourceAttr(acc.ResourceId, "project_name", ""),
					resource.TestCheckResourceAttr(acc.ResourceId, "eip_billing_config.#", "1"),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "eip_billing_config.*", map[string]string{
						"isp":              "BGP",
						"eip_billing_type": "PostPaidByBandwidth",
						"bandwidth":        "1",
					}),
					resource.TestCheckResourceAttr(acc.ResourceId, "tags.#", "2"),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "tags.*", map[string]string{
						"key":   "k2",
						"value": "v2",
					}),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "tags.*", map[string]string{
						"key":   "k3",
						"value": "v3",
					}),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "vpc_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "subnet_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "eip_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "eip_address"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "eni_address"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "master_zone_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "region_id"),
					resource.TestCheckNoResourceAttr(acc.ResourceId, "period"),
					resource.TestCheckNoResourceAttr(acc.ResourceId, "renew_type"),
				),
			},
			{
				Config:             testAccVolcengineClbUpdateTagsConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // 修改之后，不应该再产生diff
			},
		},
	})
}

const testAccVolcengineClbCreateConfigIpv6 = `
data "volcengine_zones" "foo"{
}

resource "volcengine_vpc" "vpc_ipv6" {
  vpc_name = "acc-test-vpc-ipv6"
  cidr_block = "172.16.0.0/16"
  enable_ipv6 = true
}

resource "volcengine_subnet" "subnet_ipv6" {
  subnet_name = "acc-test-subnet-ipv6"
  cidr_block = "172.16.0.0/24"
  zone_id = data.volcengine_zones.foo.zones[1].id
  vpc_id = volcengine_vpc.vpc_ipv6.id
  ipv6_cidr_block = 1
}

resource "volcengine_clb" "private_clb_ipv6" {
  type = "private"
  subnet_id = volcengine_subnet.subnet_ipv6.id
  load_balancer_name = "acc-test-clb-ipv6"
  load_balancer_spec = "small_1"
  description = "acc-test-demo"
  project_name = "default"
  address_ip_version = "DualStack"
  tags {
    key = "k1"
    value = "v1"
  }
}
`

func TestAccVolcengineClbResource_CreateIpv6(t *testing.T) {
	resourceName := "volcengine_clb.private_clb_ipv6"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &clb.VolcengineClbService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineClbCreateConfigIpv6,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "load_balancer_name", "acc-test-clb-ipv6"),
					resource.TestCheckResourceAttr(acc.ResourceId, "load_balancer_spec", "small_1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "load_balancer_billing_type", "PostPaid"),
					resource.TestCheckResourceAttr(acc.ResourceId, "type", "private"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test-demo"),
					resource.TestCheckResourceAttr(acc.ResourceId, "address_ip_version", "DualStack"),
					resource.TestCheckResourceAttr(acc.ResourceId, "eip_id", ""),
					resource.TestCheckResourceAttr(acc.ResourceId, "eip_address", ""),
					resource.TestCheckResourceAttr(acc.ResourceId, "modification_protection_reason", ""),
					resource.TestCheckResourceAttr(acc.ResourceId, "modification_protection_status", ""),
					resource.TestCheckResourceAttr(acc.ResourceId, "project_name", "default"),
					resource.TestCheckResourceAttr(acc.ResourceId, "eip_billing_config.#", "0"),
					resource.TestCheckResourceAttr(acc.ResourceId, "tags.#", "1"),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "tags.*", map[string]string{
						"key":   "k1",
						"value": "v1",
					}),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "vpc_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "subnet_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "eni_address"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "eni_ipv6_address"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "master_zone_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "region_id"),
					resource.TestCheckNoResourceAttr(acc.ResourceId, "period"),
					resource.TestCheckNoResourceAttr(acc.ResourceId, "renew_type"),
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
