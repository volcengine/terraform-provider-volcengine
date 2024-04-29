package iam_access_key_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/iam/iam_access_key"
)

const testAccVolcengineIamAccessKeysDatasourceConfig = `
data "volcengine_iam_access_keys" "foo"{
  user_name = "inner-user"
}
`

func TestAccVolcengineIamAccessKeysDatasource_Basic(t *testing.T) {
	resourceName := "data.volcengine_iam_access_keys.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &iam_access_key.VolcengineIamAccessKeyService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers: volcengine.GetTestAccProviders(),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineIamAccessKeysDatasourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(acc.ResourceId, "access_key_metadata.#", "1"),
				),
			},
		},
	})
}
