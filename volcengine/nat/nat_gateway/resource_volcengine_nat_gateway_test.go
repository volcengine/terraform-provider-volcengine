package nat_gateway_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/nat/nat_gateway"
	"testing"
)

const testAccVolcengineNatGatewayCreateConfig = `
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
`

func TestAccVolcengineNatGatewayResource_Basic(t *testing.T) {
	resourceName := "volcengine_nat_gateway.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &nat_gateway.VolcengineNatGatewayService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineNatGatewayCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "nat_gateway_name", "acc-test-ng"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test"),
					resource.TestCheckResourceAttr(acc.ResourceId, "billing_type", "PostPaid"),
					resource.TestCheckResourceAttr(acc.ResourceId, "project_name", "default"),
					resource.TestCheckResourceAttr(acc.ResourceId, "spec", "Small"),
					resource.TestCheckResourceAttr(acc.ResourceId, "tags.#", "1"),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "tags.*", map[string]string{
						"key":   "k1",
						"value": "v1",
					}),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "subnet_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "vpc_id"),
					resource.TestCheckNoResourceAttr(acc.ResourceId, "period"),
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

const testAccVolcengineNatGatewayCreatePrePaidConfig = `
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
	billing_type = "PrePaid"
	period = "2"
	project_name = "default"
	tags {
		key = "k1"
		value = "v1"
	}
}
`

func TestAccVolcengineNatGatewayResource_CreatePrePaid(t *testing.T) {
	resourceName := "volcengine_nat_gateway.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &nat_gateway.VolcengineNatGatewayService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineNatGatewayCreatePrePaidConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "nat_gateway_name", "acc-test-ng"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test"),
					resource.TestCheckResourceAttr(acc.ResourceId, "billing_type", "PrePaid"),
					resource.TestCheckResourceAttr(acc.ResourceId, "project_name", "default"),
					resource.TestCheckResourceAttr(acc.ResourceId, "spec", "Small"),
					resource.TestCheckResourceAttr(acc.ResourceId, "tags.#", "1"),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "tags.*", map[string]string{
						"key":   "k1",
						"value": "v1",
					}),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "subnet_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "vpc_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "period"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"period"},
			},
		},
	})
}

const testAccVolcengineNatGatewayUpdateBasicAttributeConfig = `
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
	spec = "Medium"
	nat_gateway_name = "acc-test-ng-new"
	description = "acc-test-new"
	billing_type = "PostPaid"
	project_name = "default"
	tags {
		key = "k1"
		value = "v1"
	}
}
`

func TestAccVolcengineNatGatewayResource_UpdateBasicAttribute(t *testing.T) {
	resourceName := "volcengine_nat_gateway.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &nat_gateway.VolcengineNatGatewayService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineNatGatewayCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "nat_gateway_name", "acc-test-ng"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test"),
					resource.TestCheckResourceAttr(acc.ResourceId, "billing_type", "PostPaid"),
					resource.TestCheckResourceAttr(acc.ResourceId, "project_name", "default"),
					resource.TestCheckResourceAttr(acc.ResourceId, "spec", "Small"),
					resource.TestCheckResourceAttr(acc.ResourceId, "tags.#", "1"),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "tags.*", map[string]string{
						"key":   "k1",
						"value": "v1",
					}),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "subnet_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "vpc_id"),
					resource.TestCheckNoResourceAttr(acc.ResourceId, "period"),
				),
			},
			{
				Config: testAccVolcengineNatGatewayUpdateBasicAttributeConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "nat_gateway_name", "acc-test-ng-new"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test-new"),
					resource.TestCheckResourceAttr(acc.ResourceId, "billing_type", "PostPaid"),
					resource.TestCheckResourceAttr(acc.ResourceId, "project_name", "default"),
					resource.TestCheckResourceAttr(acc.ResourceId, "spec", "Medium"),
					resource.TestCheckResourceAttr(acc.ResourceId, "tags.#", "1"),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "tags.*", map[string]string{
						"key":   "k1",
						"value": "v1",
					}),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "subnet_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "vpc_id"),
					resource.TestCheckNoResourceAttr(acc.ResourceId, "period"),
				),
			},
			{
				Config:             testAccVolcengineNatGatewayUpdateBasicAttributeConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // 修改之后，不应该再产生diff
			},
		},
	})
}

const testAccVolcengineNatGatewayUpdateTagsConfig = `
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
	tags {
		key = "k2"
		value = "v2"
	}
}
`

func TestAccVolcengineNatGatewayResource_UpdateBasicTags(t *testing.T) {
	resourceName := "volcengine_nat_gateway.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &nat_gateway.VolcengineNatGatewayService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineNatGatewayCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "nat_gateway_name", "acc-test-ng"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test"),
					resource.TestCheckResourceAttr(acc.ResourceId, "billing_type", "PostPaid"),
					resource.TestCheckResourceAttr(acc.ResourceId, "project_name", "default"),
					resource.TestCheckResourceAttr(acc.ResourceId, "spec", "Small"),
					resource.TestCheckResourceAttr(acc.ResourceId, "tags.#", "1"),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "tags.*", map[string]string{
						"key":   "k1",
						"value": "v1",
					}),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "subnet_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "vpc_id"),
					resource.TestCheckNoResourceAttr(acc.ResourceId, "period"),
				),
			},
			{
				Config: testAccVolcengineNatGatewayUpdateTagsConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "nat_gateway_name", "acc-test-ng"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test"),
					resource.TestCheckResourceAttr(acc.ResourceId, "billing_type", "PostPaid"),
					resource.TestCheckResourceAttr(acc.ResourceId, "project_name", "default"),
					resource.TestCheckResourceAttr(acc.ResourceId, "spec", "Small"),
					resource.TestCheckResourceAttr(acc.ResourceId, "tags.#", "2"),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "tags.*", map[string]string{
						"key":   "k1",
						"value": "v1",
					}),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "tags.*", map[string]string{
						"key":   "k2",
						"value": "v2",
					}),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "subnet_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "vpc_id"),
					resource.TestCheckNoResourceAttr(acc.ResourceId, "period"),
				),
			},
			{
				Config:             testAccVolcengineNatGatewayUpdateTagsConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // 修改之后，不应该再产生diff
			},
		},
	})
}
