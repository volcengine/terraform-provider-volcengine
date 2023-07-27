package acl_entry_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/clb/acl_entry"
	"testing"
)

const testAccVolcengineAclEntryCreateConfig = `
resource "volcengine_acl" "foo" {
	acl_name = "acc-test-acl"
	description = "acc-test-demo"
	project_name = "default"
}

resource "volcengine_acl_entry" "foo" {
    acl_id = "${volcengine_acl.foo.id}"
    entry = "172.20.1.0/24"
	description = "entry"
}
`

func TestAccVolcengineAclEntryResource_Basic(t *testing.T) {
	resourceName := "volcengine_acl_entry.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &acl_entry.VolcengineAclEntryService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineAclEntryCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "entry", "172.20.1.0/24"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "entry"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "acl_id"),
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
