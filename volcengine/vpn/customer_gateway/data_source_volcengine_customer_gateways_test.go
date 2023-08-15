package customer_gateway_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpn/customer_gateway"
	"testing"
)

const testAccVolcengineCustomerGatewaysDatasourceConfig = `
resource "volcengine_customer_gateway" "foo" {
  ip_address = "192.0.1.3"
  customer_gateway_name = "acc-test"
  description = "acc-test"
  project_name = "default"
}
data "volcengine_customer_gateways" "foo"{
    ids = ["${volcengine_customer_gateway.foo.id}"]
}
`

func TestAccVolcengineCustomerGatewaysDatasource_Basic(t *testing.T) {
	resourceName := "data.volcengine_customer_gateways.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &customer_gateway.VolcengineCustomerGatewayService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers: volcengine.GetTestAccProviders(),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineCustomerGatewaysDatasourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(acc.ResourceId, "customer_gateways.#", "1"),
				),
			},
		},
	})
}
