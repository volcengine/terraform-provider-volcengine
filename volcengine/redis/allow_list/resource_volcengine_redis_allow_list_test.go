package allow_list_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/redis/allow_list"
)

const testAccVolcengineRedisAllowListCreateConfig = `
resource "volcengine_redis_allow_list" "foo" {
    allow_list = ["192.168.0.0/24"]
    allow_list_name = "acc-test-allowlist"
}
`

const testAccVolcengineRedisAllowListUpdateConfig = `
resource "volcengine_redis_allow_list" "foo" {
    allow_list = ["192.168.0.0/24", "192.168.1.0/24"]
    allow_list_desc = "acctest"
    allow_list_name = "acc-test-allowlist1"
}
`

func TestAccVolcengineRedisAllowListResource_Basic(t *testing.T) {
	resourceName := "volcengine_redis_allow_list.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return allow_list.NewRedisAllowListService(client)
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
				Config: testAccVolcengineRedisAllowListCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "allow_list.#", "1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "allow_list_desc", ""),
					resource.TestCheckResourceAttr(acc.ResourceId, "allow_list_name", "acc-test-allowlist"),
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

func TestAccVolcengineRedisAllowListResource_Update(t *testing.T) {
	resourceName := "volcengine_redis_allow_list.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return allow_list.NewRedisAllowListService(client)
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
				Config: testAccVolcengineRedisAllowListCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "allow_list.#", "1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "allow_list_desc", ""),
					resource.TestCheckResourceAttr(acc.ResourceId, "allow_list_name", "acc-test-allowlist"),
				),
			},
			{
				Config: testAccVolcengineRedisAllowListUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "allow_list.#", "2"),
					resource.TestCheckResourceAttr(acc.ResourceId, "allow_list_desc", "acctest"),
					resource.TestCheckResourceAttr(acc.ResourceId, "allow_list_name", "acc-test-allowlist1"),
				),
			},
			{
				Config:             testAccVolcengineRedisAllowListUpdateConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // 修改之后，不应该再产生diff
			},
		},
	})
}
