package allowlist_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/rds_mysql/allowlist"
)

const testAccVolcengineRdsMysqlAllowlistCreateConfig = `
resource "volcengine_rds_mysql_allowlist" "foo" {
    allow_list_name = "acc-test-allowlist"
	allow_list_desc = "acc-test"
	allow_list_type = "IPv4"
	allow_list = ["192.168.0.0/24", "192.168.1.0/24"]
}
`

func TestAccVolcengineRdsMysqlAllowlistResource_Basic(t *testing.T) {
	resourceName := "volcengine_rds_mysql_allowlist.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return allowlist.NewRdsMysqlAllowListService(client)
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
				Config: testAccVolcengineRdsMysqlAllowlistCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "allow_list_name", "acc-test-allowlist"),
					resource.TestCheckResourceAttr(acc.ResourceId, "allow_list_type", "IPv4"),
					resource.TestCheckResourceAttr(acc.ResourceId, "allow_list_desc", "acc-test"),
					resource.TestCheckResourceAttr(acc.ResourceId, "allow_list.#", "2"),
					volcengine.TestCheckTypeSetElemAttr(acc.ResourceId, "allow_list.*", "192.168.0.0/24"),
					volcengine.TestCheckTypeSetElemAttr(acc.ResourceId, "allow_list.*", "192.168.1.0/24"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "allow_list_id"),
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

const testAccVolcengineRdsMysqlAllowlistUpdateConfig = `
resource "volcengine_rds_mysql_allowlist" "foo" {
    allow_list_name = "acc-test-allowlist-new"
	allow_list_desc = "acc-test-new"
	allow_list_type = "IPv4"
	allow_list = ["192.168.0.0/24", "192.168.3.0/24", "192.168.4.0/24"]
}
`

func TestAccVolcengineRdsMysqlAllowlistResource_Update(t *testing.T) {
	resourceName := "volcengine_rds_mysql_allowlist.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return allowlist.NewRdsMysqlAllowListService(client)
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
				Config: testAccVolcengineRdsMysqlAllowlistCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "allow_list_name", "acc-test-allowlist"),
					resource.TestCheckResourceAttr(acc.ResourceId, "allow_list_type", "IPv4"),
					resource.TestCheckResourceAttr(acc.ResourceId, "allow_list_desc", "acc-test"),
					resource.TestCheckResourceAttr(acc.ResourceId, "allow_list.#", "2"),
					volcengine.TestCheckTypeSetElemAttr(acc.ResourceId, "allow_list.*", "192.168.0.0/24"),
					volcengine.TestCheckTypeSetElemAttr(acc.ResourceId, "allow_list.*", "192.168.1.0/24"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "allow_list_id"),
				),
			},
			{
				Config: testAccVolcengineRdsMysqlAllowlistUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "allow_list_name", "acc-test-allowlist-new"),
					resource.TestCheckResourceAttr(acc.ResourceId, "allow_list_type", "IPv4"),
					resource.TestCheckResourceAttr(acc.ResourceId, "allow_list_desc", "acc-test-new"),
					resource.TestCheckResourceAttr(acc.ResourceId, "allow_list.#", "3"),
					volcengine.TestCheckTypeSetElemAttr(acc.ResourceId, "allow_list.*", "192.168.0.0/24"),
					volcengine.TestCheckTypeSetElemAttr(acc.ResourceId, "allow_list.*", "192.168.3.0/24"),
					volcengine.TestCheckTypeSetElemAttr(acc.ResourceId, "allow_list.*", "192.168.4.0/24"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "allow_list_id"),
				),
			},
			{
				Config:             testAccVolcengineRdsMysqlAllowlistUpdateConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // 修改之后，不应该再产生diff
			},
		},
	})
}
