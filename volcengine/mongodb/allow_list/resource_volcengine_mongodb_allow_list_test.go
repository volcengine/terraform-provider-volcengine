package allow_list_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/mongodb/allow_list"
)

const testAccVolcengineMongodbAllowListCreateConfig = `
resource "volcengine_mongodb_allow_list" "foo" {
    allow_list_name="acc-test"
    allow_list_desc="acc-test"
    allow_list_type="IPv4"
    allow_list="10.1.1.3,10.2.3.0/24,10.1.1.1"
}
`

const testAccVolcengineMongodbAllowListUpdateConfig = `
resource "volcengine_mongodb_allow_list" "foo" {
    allow_list = "10.2.3.0/24"
    allow_list_desc = "acc-test-modify"
    allow_list_name = "acc-test-modify"
    allow_list_type = "IPv4"
}
`

func TestAccVolcengineMongodbAllowListResource_Basic(t *testing.T) {
	resourceName := "volcengine_mongodb_allow_list.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return allow_list.NewMongoDBAllowListService(client)
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
				Config: testAccVolcengineMongodbAllowListCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "allow_list", "10.1.1.3,10.2.3.0/24,10.1.1.1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "allow_list_desc", "acc-test"),
					resource.TestCheckResourceAttr(acc.ResourceId, "allow_list_name", "acc-test"),
					resource.TestCheckResourceAttr(acc.ResourceId, "allow_list_type", "IPv4"),
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

func TestAccVolcengineMongodbAllowListResource_Update(t *testing.T) {
	resourceName := "volcengine_mongodb_allow_list.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return allow_list.NewMongoDBAllowListService(client)
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
				Config: testAccVolcengineMongodbAllowListCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "allow_list", "10.1.1.3,10.2.3.0/24,10.1.1.1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "allow_list_desc", "acc-test"),
					resource.TestCheckResourceAttr(acc.ResourceId, "allow_list_name", "acc-test"),
					resource.TestCheckResourceAttr(acc.ResourceId, "allow_list_type", "IPv4"),
				),
			},
			{
				Config: testAccVolcengineMongodbAllowListUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "allow_list", "10.2.3.0/24"),
					resource.TestCheckResourceAttr(acc.ResourceId, "allow_list_desc", "acc-test-modify"),
					resource.TestCheckResourceAttr(acc.ResourceId, "allow_list_name", "acc-test-modify"),
					resource.TestCheckResourceAttr(acc.ResourceId, "allow_list_type", "IPv4"),
				),
			},
			{
				Config:             testAccVolcengineMongodbAllowListUpdateConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // 修改之后，不应该再产生diff
			},
		},
	})
}
