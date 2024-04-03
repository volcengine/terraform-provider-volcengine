package bandwidth_package_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/bandwidth_package/bandwidth_package"
)

const testAccVolcengineBandwidthPackagesDatasourceConfig = `
resource "volcengine_bandwidth_package" "foo" {
  bandwidth_package_name    = "acc-test-bp"
  billing_type              = "PostPaidByBandwidth"
  isp                       = "BGP"
  description               = "acc-test"
  bandwidth                 = 2
  protocol                  = "IPv4"
  security_protection_types = ["AntiDDoS_Enhanced"]
  tags {
    key   = "k1"
    value = "v1"
  }
  count = 2
}

data "volcengine_bandwidth_packages" "foo" {
  ids = volcengine_bandwidth_package.foo[*].id
}
`

func TestAccVolcengineBandwidthPackagesDatasource_Basic(t *testing.T) {
	resourceName := "data.volcengine_bandwidth_packages.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return bandwidth_package.NewBandwidthPackageService(client)
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers: volcengine.GetTestAccProviders(),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineBandwidthPackagesDatasourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(acc.ResourceId, "packages.#", "2"),
				),
			},
		},
	})
}
