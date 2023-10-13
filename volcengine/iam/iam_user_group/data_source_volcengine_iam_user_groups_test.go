package iam_user_group_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/iam/iam_user_group"
)

const testAccVolcengineIamUserGroupsDatasourceConfig = `
resource "volcengine_iam_user_group" "foo" {
  user_group_name = "acc-test-group"
  description = "acc-test"
  display_name = "acc-test"
}

data "volcengine_iam_user_groups" "foo"{
    query = "acc-test-group"
}
`

func TestAccVolcengineIamUserGroupsDatasource_Basic(t *testing.T) {
	resourceName := "data.volcengine_iam_user_groups.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return iam_user_group.NewIamUserGroupService(client)
		},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers: volcengine.GetTestAccProviders(),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineIamUserGroupsDatasourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(acc.ResourceId, "user_groups"),
				),
			},
		},
	})
}
