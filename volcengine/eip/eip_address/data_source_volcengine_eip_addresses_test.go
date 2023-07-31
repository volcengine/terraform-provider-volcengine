package eip_address_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/eip/eip_address"
	"testing"
)

const testAccVolcengineEipAddressesDatasourceConfig = `
resource "volcengine_eip_address" "foo" {
    billing_type = "PostPaidByTraffic"
}
data "volcengine_eip_addresses" "foo"{
    ids = ["${volcengine_eip_address.foo.id}"]
}
`

func TestAccVolcengineEipAddressesDatasource_Basic(t *testing.T) {
	resourceName := "data.volcengine_eip_addresses.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &eip_address.VolcengineEipAddressService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers: volcengine.GetTestAccProviders(),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineEipAddressesDatasourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(acc.ResourceId, "addresses.#", "1"),
				),
			},
		},
	})
}
