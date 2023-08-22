package nas_permission_group_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/nas/nas_permission_group"
)

const testAccVolcengineNasPermissionGroupsDatasourceConfig = `
resource "volcengine_nas_permission_group" "foo" {
  permission_group_name = "acc-test"
  description = "acctest"
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

data "volcengine_nas_permission_groups" "foo" {
  filters {
    key = "PermissionGroupId"
    value = volcengine_nas_permission_group.foo.id
  }
}
`

func TestAccVolcengineNasPermissionGroupsDatasource_Basic(t *testing.T) {
	resourceName := "data.volcengine_nas_permission_groups.foo"

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
		Providers: volcengine.GetTestAccProviders(),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineNasPermissionGroupsDatasourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(acc.ResourceId, "permission_groups.#", "1"),
				),
			},
		},
	})
}
