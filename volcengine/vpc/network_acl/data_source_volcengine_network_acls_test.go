package network_acl_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vpc/network_acl"
)

const testAccNetworkAclConfig = `
	
`

func TestAccVolcengineIpv6GatewayDataSource_Basic(t *testing.T) {
	resourceName := "data.volcengine_vpc_network_acls.foo"
	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &network_acl.VolcengineNetworkAclService{},
	}
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers: volcengine.GetTestAccProviders(),
		Steps: []resource.TestStep{
			{
				Config: testAccNetworkAclConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(acc.ResourceId, "network_acls.#", "1"),
				),
			},
		},
	})
}
