package cen_inter_region_bandwidth_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/cen/cen_inter_region_bandwidth"
)

const testAccVolcengineCenInterRegionBandwidthsDatasourceConfig = `
resource "volcengine_cen" "foo" {
  cen_name     = "acc-test-cen"
  description  = "acc-test"
  project_name = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}

resource "volcengine_cen_bandwidth_package" "foo" {
  local_geographic_region_set_id = "China"
  peer_geographic_region_set_id  = "China"
  bandwidth                      = 5
  cen_bandwidth_package_name     = "acc-test-cen-bp"
  description                    = "acc-test"
  billing_type                   = "PrePaid"
  period_unit                    = "Month"
  period                         = 1
  project_name                   = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}

resource "volcengine_cen_bandwidth_package_associate" "foo" {
  cen_bandwidth_package_id = volcengine_cen_bandwidth_package.foo.id
  cen_id                   = volcengine_cen.foo.id
}

resource "volcengine_cen_inter_region_bandwidth" "foo" {
  cen_id          = volcengine_cen.foo.id
  local_region_id = "cn-beijing"
  peer_region_id  = "cn-shanghai"
  bandwidth       = 2
  depends_on      = [volcengine_cen_bandwidth_package_associate.foo]
}

data "volcengine_cen_inter_region_bandwidths" "foo"{
  ids = [volcengine_cen_inter_region_bandwidth.foo.id]
}
`

func TestAccVolcengineCenInterRegionBandwidthsDatasource_Basic(t *testing.T) {
	resourceName := "data.volcengine_cen_inter_region_bandwidths.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return cen_inter_region_bandwidth.NewCenInterRegionBandwidthService(client)
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers: volcengine.GetTestAccProviders(),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineCenInterRegionBandwidthsDatasourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(acc.ResourceId, "inter_region_bandwidths.#", "1"),
				),
			},
		},
	})
}
