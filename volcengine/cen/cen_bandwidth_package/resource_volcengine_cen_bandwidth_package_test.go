package cen_bandwidth_package_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/cen/cen_bandwidth_package"
)

const testAccVolcengineCenBandwidthPackageCreateConfig = `
resource "volcengine_cen_bandwidth_package" "foo" {
  local_geographic_region_set_id = "China"
  peer_geographic_region_set_id  = "China"
  bandwidth                      = 2
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
`

func TestAccVolcengineCenBandwidthPackageResource_Basic(t *testing.T) {
	resourceName := "volcengine_cen_bandwidth_package.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return cen_bandwidth_package.NewCenBandwidthPackageService(client)
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
				Config: testAccVolcengineCenBandwidthPackageCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "bandwidth", "2"),
					resource.TestCheckResourceAttr(acc.ResourceId, "billing_type", "PrePaid"),
					resource.TestCheckResourceAttr(acc.ResourceId, "cen_bandwidth_package_name", "acc-test-cen-bp"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test"),
					resource.TestCheckResourceAttr(acc.ResourceId, "local_geographic_region_set_id", "China"),
					resource.TestCheckResourceAttr(acc.ResourceId, "peer_geographic_region_set_id", "China"),
					resource.TestCheckResourceAttr(acc.ResourceId, "project_name", "default"),
					resource.TestCheckResourceAttr(acc.ResourceId, "tags.#", "1"),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "tags.*", map[string]string{
						"key":   "k1",
						"value": "v1",
					}),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"period_unit", "period"},
			},
		},
	})
}

const testAccVolcengineCenBandwidthPackageUpdateConfig = `
resource "volcengine_cen_bandwidth_package" "foo" {
  local_geographic_region_set_id = "China"
  peer_geographic_region_set_id  = "China"
  bandwidth                      = 5
  cen_bandwidth_package_name     = "acc-test-cen-bp-new"
  description                    = "acc-test-new"
  billing_type                   = "PrePaid"
  period_unit                    = "Month"
  period                         = 1
  project_name                   = "default"
  tags {
    key   = "k2"
    value = "v2"
  }
  tags {
    key   = "k3"
    value = "v3"
  }
}
`

func TestAccVolcengineCenBandwidthPackageResource_Update(t *testing.T) {
	resourceName := "volcengine_cen_bandwidth_package.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return cen_bandwidth_package.NewCenBandwidthPackageService(client)
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
				Config: testAccVolcengineCenBandwidthPackageCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "bandwidth", "2"),
					resource.TestCheckResourceAttr(acc.ResourceId, "billing_type", "PrePaid"),
					resource.TestCheckResourceAttr(acc.ResourceId, "cen_bandwidth_package_name", "acc-test-cen-bp"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test"),
					resource.TestCheckResourceAttr(acc.ResourceId, "local_geographic_region_set_id", "China"),
					resource.TestCheckResourceAttr(acc.ResourceId, "peer_geographic_region_set_id", "China"),
					resource.TestCheckResourceAttr(acc.ResourceId, "project_name", "default"),
					resource.TestCheckResourceAttr(acc.ResourceId, "tags.#", "1"),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "tags.*", map[string]string{
						"key":   "k1",
						"value": "v1",
					}),
				),
			},
			{
				Config: testAccVolcengineCenBandwidthPackageUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "bandwidth", "5"),
					resource.TestCheckResourceAttr(acc.ResourceId, "billing_type", "PrePaid"),
					resource.TestCheckResourceAttr(acc.ResourceId, "cen_bandwidth_package_name", "acc-test-cen-bp-new"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test-new"),
					resource.TestCheckResourceAttr(acc.ResourceId, "local_geographic_region_set_id", "China"),
					resource.TestCheckResourceAttr(acc.ResourceId, "peer_geographic_region_set_id", "China"),
					resource.TestCheckResourceAttr(acc.ResourceId, "project_name", "default"),
					resource.TestCheckResourceAttr(acc.ResourceId, "tags.#", "2"),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "tags.*", map[string]string{
						"key":   "k2",
						"value": "v2",
					}),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "tags.*", map[string]string{
						"key":   "k3",
						"value": "v3",
					}),
				),
			},
			{
				Config:             testAccVolcengineCenBandwidthPackageUpdateConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // 修改之后，不应该再产生diff
			},
		},
	})
}
