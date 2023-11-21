package allow_list_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/mongodb/allow_list"
)

const testAccVolcengineMongodbAllowListsDatasourceConfig = `
resource "volcengine_mongodb_allow_list" "foo" {
    allow_list_name="acc-test"
    allow_list_desc="acc-test"
    allow_list_type="IPv4"
    allow_list="10.1.1.3,10.2.3.0/24,10.1.1.1"
}

data "volcengine_mongodb_allow_lists" "foo"{
    allow_list_ids = [volcengine_mongodb_allow_list.foo.id]
    region_id = "cn-beijing"
}
`

func TestAccVolcengineMongodbAllowListsDatasource_Basic(t *testing.T) {
	resourceName := "data.volcengine_mongodb_allow_lists.foo"

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
		Providers: volcengine.GetTestAccProviders(),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineMongodbAllowListsDatasourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(acc.ResourceId, "allow_lists.#", "1"),
				),
			},
		},
	})
}
