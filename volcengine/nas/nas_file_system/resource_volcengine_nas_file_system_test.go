package nas_file_system_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/nas/nas_file_system"
)

const testAccVolcengineNasFileSystemCreateConfig = `
data "volcengine_nas_zones" "foo" {

}

resource "volcengine_nas_file_system" "foo" {
    file_system_name = "acc-test-fs"
  	description = "acc-test"
  	zone_id = "${data.volcengine_nas_zones.foo.zones[0].id}"
  	capacity = 103
  	project_name = "default"
  	tags {
    	key = "k1"
    	value = "v1"
  	}
}
`

func TestAccVolcengineNasFileSystemResource_Basic(t *testing.T) {
	resourceName := "volcengine_nas_file_system.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return nas_file_system.NewNasFileSystemService(client)
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
				Config: testAccVolcengineNasFileSystemCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "file_system_name", "acc-test-fs"),
					resource.TestCheckResourceAttr(acc.ResourceId, "capacity", "103"),
					resource.TestCheckResourceAttr(acc.ResourceId, "charge_type", "PayAsYouGo"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test"),
					resource.TestCheckResourceAttr(acc.ResourceId, "file_system_type", "Extreme"),
					resource.TestCheckResourceAttr(acc.ResourceId, "project_name", "default"),
					resource.TestCheckResourceAttr(acc.ResourceId, "protocol_type", "NFS"),
					resource.TestCheckResourceAttr(acc.ResourceId, "snapshot_count", "0"),
					resource.TestCheckResourceAttr(acc.ResourceId, "status", "Running"),
					resource.TestCheckResourceAttr(acc.ResourceId, "storage_type", "Standard"),
					resource.TestCheckResourceAttr(acc.ResourceId, "tags.#", "1"),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "tags.*", map[string]string{
						"key":   "k1",
						"value": "v1",
					}),
					resource.TestCheckNoResourceAttr(acc.ResourceId, "snapshot_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "create_time"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "update_time"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "version"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "region_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "zone_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "zone_name"),
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

const testAccVolcengineNasFileSystemUpdateConfig = `
data "volcengine_nas_zones" "foo" {

}

resource "volcengine_nas_file_system" "foo" {
    file_system_name = "acc-test-fs-new"
  	description = "acc-test-new"
  	zone_id = "${data.volcengine_nas_zones.foo.zones[0].id}"
  	capacity = 105
  	project_name = "default"
  	tags {
    	key = "k2"
    	value = "v2"
  	}
	tags {
    	key = "k3"
    	value = "v3"
  	}
}
`

func TestAccVolcengineNasFileSystemResource_Update(t *testing.T) {
	resourceName := "volcengine_nas_file_system.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return nas_file_system.NewNasFileSystemService(client)
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
				Config: testAccVolcengineNasFileSystemCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "file_system_name", "acc-test-fs"),
					resource.TestCheckResourceAttr(acc.ResourceId, "capacity", "103"),
					resource.TestCheckResourceAttr(acc.ResourceId, "charge_type", "PayAsYouGo"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test"),
					resource.TestCheckResourceAttr(acc.ResourceId, "file_system_type", "Extreme"),
					resource.TestCheckResourceAttr(acc.ResourceId, "project_name", "default"),
					resource.TestCheckResourceAttr(acc.ResourceId, "protocol_type", "NFS"),
					resource.TestCheckResourceAttr(acc.ResourceId, "snapshot_count", "0"),
					resource.TestCheckResourceAttr(acc.ResourceId, "status", "Running"),
					resource.TestCheckResourceAttr(acc.ResourceId, "storage_type", "Standard"),
					resource.TestCheckResourceAttr(acc.ResourceId, "tags.#", "1"),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "tags.*", map[string]string{
						"key":   "k1",
						"value": "v1",
					}),
					resource.TestCheckNoResourceAttr(acc.ResourceId, "snapshot_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "create_time"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "update_time"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "version"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "region_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "zone_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "zone_name"),
				),
			},
			{
				Config: testAccVolcengineNasFileSystemUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "file_system_name", "acc-test-fs-new"),
					resource.TestCheckResourceAttr(acc.ResourceId, "capacity", "105"),
					resource.TestCheckResourceAttr(acc.ResourceId, "charge_type", "PayAsYouGo"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test-new"),
					resource.TestCheckResourceAttr(acc.ResourceId, "file_system_type", "Extreme"),
					resource.TestCheckResourceAttr(acc.ResourceId, "project_name", "default"),
					resource.TestCheckResourceAttr(acc.ResourceId, "protocol_type", "NFS"),
					resource.TestCheckResourceAttr(acc.ResourceId, "snapshot_count", "0"),
					resource.TestCheckResourceAttr(acc.ResourceId, "status", "Running"),
					resource.TestCheckResourceAttr(acc.ResourceId, "storage_type", "Standard"),
					resource.TestCheckResourceAttr(acc.ResourceId, "tags.#", "2"),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "tags.*", map[string]string{
						"key":   "k2",
						"value": "v2",
					}),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "tags.*", map[string]string{
						"key":   "k3",
						"value": "v3",
					}),
					resource.TestCheckNoResourceAttr(acc.ResourceId, "snapshot_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "create_time"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "update_time"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "version"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "region_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "zone_id"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "zone_name"),
				),
			},
			{
				Config:             testAccVolcengineNasFileSystemUpdateConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // 修改之后，不应该再产生diff
			},
		},
	})
}
