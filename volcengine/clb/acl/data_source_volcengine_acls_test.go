package acl_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/clb/acl"
	"testing"
)

const testAccVolcengineAclsDatasourceConfig = `
resource "volcengine_acl" "foo" {
	acl_name = "acc-test-acl-${count.index}"
	description = "acc-test-demo"
	project_name = "default"
	acl_entries {
    	entry = "172.20.1.0/24"
    	description = "e1"
  	}
	count = 3
}

data "volcengine_acls" "foo"{
    ids = volcengine_acl.foo[*].id
}
`

func TestAccVolcengineAclsDatasource_Basic(t *testing.T) {
	resourceName := "data.volcengine_acls.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &acl.VolcengineAclService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers: volcengine.GetTestAccProviders(),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineAclsDatasourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(acc.ResourceId, "acls.#", "3"),
				),
			},
		},
	})
}
