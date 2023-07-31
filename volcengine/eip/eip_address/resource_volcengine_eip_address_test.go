package eip_address_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/eip/eip_address"
	"testing"
)

const testAccVolcengineEipAddressCreateConfig = `
resource "volcengine_eip_address" "foo" {
    billing_type = "PostPaidByTraffic"
}
`

const testAccVolcengineEipAddressUpdateConfig = `
resource "volcengine_eip_address" "foo" {
    bandwidth = 1
    billing_type = "PostPaidByBandwidth"
    description = "acc-test"
    isp = "BGP"
    name = "acc-test-eip"
}
`

func TestAccVolcengineEipAddressResource_Basic(t *testing.T) {
	resourceName := "volcengine_eip_address.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &eip_address.VolcengineEipAddressService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineEipAddressCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "billing_type", "PostPaidByTraffic"),
					resource.TestCheckResourceAttr(acc.ResourceId, "status", "Available"),
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

func TestAccVolcengineEipAddressResource_Update(t *testing.T) {
	resourceName := "volcengine_eip_address.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &eip_address.VolcengineEipAddressService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineEipAddressCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "billing_type", "PostPaidByTraffic"),
					resource.TestCheckResourceAttr(acc.ResourceId, "status", "Available"),
				),
			},
			{
				Config: testAccVolcengineEipAddressUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "bandwidth", "1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "billing_type", "PostPaidByBandwidth"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test"),
					resource.TestCheckResourceAttr(acc.ResourceId, "isp", "BGP"),
					resource.TestCheckResourceAttr(acc.ResourceId, "name", "acc-test-eip"),
					resource.TestCheckResourceAttr(acc.ResourceId, "status", "Available"),
				),
			},
			{
				Config:             testAccVolcengineEipAddressUpdateConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // 修改之后，不应该再产生diff
			},
		},
	})
}
