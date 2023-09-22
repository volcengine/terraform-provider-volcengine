package nas_permission_group_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/nas/nas_permission_group"
)

const testAccVolcengineNasPermissionGroupCreateConfig = `
resource "volcengine_nas_permission_group" "foo" {
  permission_group_name = "acc-test"
  description = "acctest"
  permission_rules {
    cidr_ip = "*"
    rw_mode = "RW"
    use_mode = "All_squash"
  }
}
`

const testAccVolcengineNasPermissionGroupUpdateConfig = `
resource "volcengine_nas_permission_group" "foo" {
  permission_group_name = "acc-test1"
  description = "acctest1"
  permission_rules {
    cidr_ip = "*"
    rw_mode = "RW"
    use_mode = "All_squash"
  }
  permission_rules {
    cidr_ip = "192.168.0.0"
    rw_mode = "RO"
    use_mode = "All_squash"
  }
}
`

const testAccVolcengineNasPermissionGroupUpdate2Config = `
resource "volcengine_nas_permission_group" "foo" {
  permission_group_name = "acc-test1"
  description = "acctest1"
}
`

func TestAccVolcengineNasPermissionGroupResource_Basic(t *testing.T) {
	resourceName := "volcengine_nas_permission_group.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return nas_permission_group.NewVolcengineNasPermissionGroupService(client)
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
				Config: testAccVolcengineNasPermissionGroupCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acctest"),
					resource.TestCheckResourceAttr(acc.ResourceId, "permission_group_name", "acc-test"),
					resource.TestCheckResourceAttr(acc.ResourceId, "permission_rules.#", "1"),
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

func TestAccVolcengineNasPermissionGroupResource_Update(t *testing.T) {
	resourceName := "volcengine_nas_permission_group.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return nas_permission_group.NewVolcengineNasPermissionGroupService(client)
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
				Config: testAccVolcengineNasPermissionGroupCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acctest"),
					resource.TestCheckResourceAttr(acc.ResourceId, "permission_group_name", "acc-test"),
					resource.TestCheckResourceAttr(acc.ResourceId, "permission_rules.#", "1"),
				),
			},
			{
				Config: testAccVolcengineNasPermissionGroupUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acctest1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "permission_group_name", "acc-test1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "permission_rules.#", "2"),
				),
			},
			{
				Config: testAccVolcengineNasPermissionGroupUpdate2Config,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acctest1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "permission_group_name", "acc-test1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "permission_rules.#", "0"),
				),
			},
			{
				Config:             testAccVolcengineNasPermissionGroupUpdate2Config,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // 修改之后，不应该再产生diff
			},
		},
	})
}
