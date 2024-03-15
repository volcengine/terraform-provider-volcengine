package cen_inter_region_bandwidth_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/cen/cen_inter_region_bandwidth"
)

const testAccVolcengineCenInterRegionBandwidthCreateConfig = `
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
`

func TestAccVolcengineCenInterRegionBandwidthResource_Basic(t *testing.T) {
	resourceName := "volcengine_cen_inter_region_bandwidth.foo"

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
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineCenInterRegionBandwidthCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "bandwidth", "2"),
					resource.TestCheckResourceAttr(acc.ResourceId, "local_region_id", "cn-beijing"),
					resource.TestCheckResourceAttr(acc.ResourceId, "peer_region_id", "cn-shanghai"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "cen_id"),
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

const testAccVolcengineCenInterRegionBandwidthUpdateConfig = `
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
  bandwidth       = 5
  depends_on      = [volcengine_cen_bandwidth_package_associate.foo]
}
`

func TestAccVolcengineCenInterRegionBandwidthResource_Update(t *testing.T) {
	resourceName := "volcengine_cen_inter_region_bandwidth.foo"

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
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineCenInterRegionBandwidthCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "bandwidth", "2"),
					resource.TestCheckResourceAttr(acc.ResourceId, "local_region_id", "cn-beijing"),
					resource.TestCheckResourceAttr(acc.ResourceId, "peer_region_id", "cn-shanghai"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "cen_id"),
				),
			},
			{
				Config: testAccVolcengineCenInterRegionBandwidthUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "bandwidth", "5"),
					resource.TestCheckResourceAttr(acc.ResourceId, "local_region_id", "cn-beijing"),
					resource.TestCheckResourceAttr(acc.ResourceId, "peer_region_id", "cn-shanghai"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "cen_id"),
				),
			},
			{
				Config:             testAccVolcengineCenInterRegionBandwidthUpdateConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // 修改之后，不应该再产生diff
			},
		},
	})
}
