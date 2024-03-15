package cen_service_route_entry_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/cen/cen_service_route_entry"
)

const testAccVolcengineCenServiceRouteEntryCreateConfig = `
resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
  count      = 3
}

resource "volcengine_cen" "foo" {
  cen_name     = "acc-test-cen"
  description  = "acc-test"
  project_name = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}

resource "volcengine_cen_attach_instance" "foo" {
  cen_id             = volcengine_cen.foo.id
  instance_id        = volcengine_vpc.foo[count.index].id
  instance_region_id = "cn-beijing"
  instance_type      = "VPC"
  count              = 3
}

resource "volcengine_cen_service_route_entry" "foo" {
  cen_id                 = volcengine_cen.foo.id
  destination_cidr_block = "100.64.0.0/11"
  service_region_id      = "cn-beijing"
  service_vpc_id         = volcengine_cen_attach_instance.foo[0].instance_id
  description            = "acc-test"
  publish_mode           = "Custom"
  publish_to_instances {
    instance_region_id = "cn-beijing"
    instance_type      = "VPC"
    instance_id        = volcengine_cen_attach_instance.foo[1].instance_id
  }
  publish_to_instances {
    instance_region_id = "cn-beijing"
    instance_type      = "VPC"
    instance_id        = volcengine_cen_attach_instance.foo[2].instance_id
  }
}
`

func TestAccVolcengineCenServiceRouteEntryResource_Basic(t *testing.T) {
	resourceName := "volcengine_cen_service_route_entry.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return cen_service_route_entry.NewCenServiceRouteEntryService(client)
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
				Config: testAccVolcengineCenServiceRouteEntryCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test"),
					resource.TestCheckResourceAttr(acc.ResourceId, "destination_cidr_block", "100.64.0.0/11"),
					resource.TestCheckResourceAttr(acc.ResourceId, "publish_mode", "Custom"),
					resource.TestCheckResourceAttr(acc.ResourceId, "service_region_id", "cn-beijing"),
					resource.TestCheckResourceAttr(acc.ResourceId, "publish_to_instances.#", "2"),
					resource.TestCheckResourceAttr(acc.ResourceId, "publish_to_instances.0.instance_type", "VPC"),
					resource.TestCheckResourceAttr(acc.ResourceId, "publish_to_instances.0.instance_region_id", "cn-beijing"),
					resource.TestCheckResourceAttr(acc.ResourceId, "publish_to_instances.1.instance_type", "VPC"),
					resource.TestCheckResourceAttr(acc.ResourceId, "publish_to_instances.1.instance_region_id", "cn-beijing"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "cen_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "service_vpc_id"),
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

const testAccVolcengineCenServiceRouteEntryUpdateConfig = `
resource "volcengine_vpc" "foo" {
  vpc_name   = "acc-test-vpc"
  cidr_block = "172.16.0.0/16"
  count      = 3
}

resource "volcengine_cen" "foo" {
  cen_name     = "acc-test-cen"
  description  = "acc-test"
  project_name = "default"
  tags {
    key   = "k1"
    value = "v1"
  }
}

resource "volcengine_cen_attach_instance" "foo" {
  cen_id             = volcengine_cen.foo.id
  instance_id        = volcengine_vpc.foo[count.index].id
  instance_region_id = "cn-beijing"
  instance_type      = "VPC"
  count              = 3
}

resource "volcengine_cen_service_route_entry" "foo" {
  cen_id                 = volcengine_cen.foo.id
  destination_cidr_block = "100.64.0.0/11"
  service_region_id      = "cn-beijing"
  service_vpc_id         = volcengine_cen_attach_instance.foo[0].instance_id
  description            = "acc-test-new"
  publish_mode           = "LocalDCGW"
}
`

func TestAccVolcengineCenServiceRouteEntryResource_Update(t *testing.T) {
	resourceName := "volcengine_cen_service_route_entry.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return cen_service_route_entry.NewCenServiceRouteEntryService(client)
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
				Config: testAccVolcengineCenServiceRouteEntryCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test"),
					resource.TestCheckResourceAttr(acc.ResourceId, "destination_cidr_block", "100.64.0.0/11"),
					resource.TestCheckResourceAttr(acc.ResourceId, "publish_mode", "Custom"),
					resource.TestCheckResourceAttr(acc.ResourceId, "service_region_id", "cn-beijing"),
					resource.TestCheckResourceAttr(acc.ResourceId, "publish_to_instances.#", "2"),
					resource.TestCheckResourceAttr(acc.ResourceId, "publish_to_instances.0.instance_type", "VPC"),
					resource.TestCheckResourceAttr(acc.ResourceId, "publish_to_instances.0.instance_region_id", "cn-beijing"),
					resource.TestCheckResourceAttr(acc.ResourceId, "publish_to_instances.1.instance_type", "VPC"),
					resource.TestCheckResourceAttr(acc.ResourceId, "publish_to_instances.1.instance_region_id", "cn-beijing"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "cen_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "service_vpc_id"),
				),
			},
			{
				Config: testAccVolcengineCenServiceRouteEntryUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test-new"),
					resource.TestCheckResourceAttr(acc.ResourceId, "destination_cidr_block", "100.64.0.0/11"),
					resource.TestCheckResourceAttr(acc.ResourceId, "publish_mode", "LocalDCGW"),
					resource.TestCheckResourceAttr(acc.ResourceId, "service_region_id", "cn-beijing"),
					resource.TestCheckResourceAttr(acc.ResourceId, "publish_to_instances.#", "0"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "cen_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "service_vpc_id"),
				),
			},
			{
				Config:             testAccVolcengineCenServiceRouteEntryUpdateConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // 修改之后，不应该再产生diff
			},
		},
	})
}
