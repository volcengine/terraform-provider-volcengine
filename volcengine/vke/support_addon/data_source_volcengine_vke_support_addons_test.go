package support_addon_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vke/support_addon"
	"testing"
)

const testAccVolcengineVkeSupportAddonsDatasourceConfig = `
data "volcengine_vke_support_addons" "foo"{
}
`

func TestAccVolcengineVkeSupportAddonsDatasource_Basic(t *testing.T) {
	resourceName := "data.volcengine_vke_support_addons.foo"

	_ = &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &support_addon.VolcengineVkeSupportAddonService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers: volcengine.GetTestAccProviders(),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineVkeSupportAddonsDatasourceConfig,
				Check:  resource.ComposeTestCheckFunc(),
			},
		},
	})
}
