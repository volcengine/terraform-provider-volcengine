package iam_identity_provider_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/iam/iam_identity_provider"
)

func TestAccVolcengineIamIdentityProvidersDataSource_Basic(t *testing.T) {
	resourceName := "data.volcengine_iam_identity_providers.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &iam_identity_provider.VolcengineIamIdentityProviderService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers: volcengine.GetTestAccProviders(),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineIamIdentityProvidersDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(acc.ResourceId, "total_count"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "providers.#"),
				),
			},
		},
	})
}

func testAccVolcengineIamIdentityProvidersDataSourceConfig() string {
	return `
data "volcengine_iam_identity_providers" "foo" {
}
`
}
