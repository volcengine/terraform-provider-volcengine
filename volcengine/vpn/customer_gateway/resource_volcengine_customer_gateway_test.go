package customer_gateway_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpn/customer_gateway"
	"testing"
)

const testAccVolcengineCustomerGatewayCreateConfig = `
resource "volcengine_customer_gateway" "foo" {
  ip_address = "192.0.1.3"
  customer_gateway_name = "acc-test"
  description = "acc-test"
  project_name = "default"
}
`

const testAccVolcengineCustomerGatewayUpdateConfig = `
resource "volcengine_customer_gateway" "foo" {
    customer_gateway_name = "acc-test1"
    description = "acc-test1"
    ip_address = "192.0.1.3"
    project_name = "default"
}
`

func TestAccVolcengineCustomerGatewayResource_Basic(t *testing.T) {
	resourceName := "volcengine_customer_gateway.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &customer_gateway.VolcengineCustomerGatewayService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineCustomerGatewayCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "customer_gateway_name", "acc-test"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test"),
					resource.TestCheckResourceAttr(acc.ResourceId, "ip_address", "192.0.1.3"),
					resource.TestCheckResourceAttr(acc.ResourceId, "project_name", "default"),
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

func TestAccVolcengineCustomerGatewayResource_Update(t *testing.T) {
	resourceName := "volcengine_customer_gateway.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &customer_gateway.VolcengineCustomerGatewayService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineCustomerGatewayCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "customer_gateway_name", "acc-test"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test"),
					resource.TestCheckResourceAttr(acc.ResourceId, "ip_address", "192.0.1.3"),
					resource.TestCheckResourceAttr(acc.ResourceId, "project_name", "default"),
				),
			},
			{
				Config: testAccVolcengineCustomerGatewayUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "customer_gateway_name", "acc-test1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "ip_address", "192.0.1.3"),
					resource.TestCheckResourceAttr(acc.ResourceId, "project_name", "default"),
				),
			},
			{
				Config:             testAccVolcengineCustomerGatewayUpdateConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // 修改之后，不应该再产生diff
			},
		},
	})
}
