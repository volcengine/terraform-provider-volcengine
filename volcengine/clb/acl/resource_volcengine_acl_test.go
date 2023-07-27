package acl_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/clb/acl"
	"testing"
)

const testAccVolcengineAclCreateConfig = `
resource "volcengine_acl" "foo" {
	acl_name = "acc-test-acl"
	description = "acc-test-demo"
	project_name = "default"
	acl_entries {
    	entry = "172.20.1.0/24"
    	description = "e1"
  	}
}
`

func TestAccVolcengineAclResource_Basic(t *testing.T) {
	resourceName := "volcengine_acl.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &acl.VolcengineAclService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineAclCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "acl_name", "acc-test-acl"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test-demo"),
					resource.TestCheckResourceAttr(acc.ResourceId, "project_name", "default"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "create_time"),
					resource.TestCheckResourceAttr(acc.ResourceId, "acl_entries.#", "1"),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "acl_entries.*", map[string]string{
						"entry":       "172.20.1.0/24",
						"description": "e1",
					}),
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

const testAccVolcengineAclUpdateConfig = `
resource "volcengine_acl" "foo" {
    acl_name = "acc-test-acl-new"
    description = "acc-test-demo-new"
    project_name = "default"
	acl_entries {
    	entry = "172.20.2.0/24"
    	description = "e2"
  	}
	acl_entries {
    	entry = "172.20.3.0/24"
    	description = "e3"
  	}
}
`

func TestAccVolcengineAclResource_Update(t *testing.T) {
	resourceName := "volcengine_acl.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &acl.VolcengineAclService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineAclCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "acl_name", "acc-test-acl"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test-demo"),
					resource.TestCheckResourceAttr(acc.ResourceId, "project_name", "default"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "create_time"),
					resource.TestCheckResourceAttr(acc.ResourceId, "acl_entries.#", "1"),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "acl_entries.*", map[string]string{
						"entry":       "172.20.1.0/24",
						"description": "e1",
					}),
				),
			},
			{
				Config: testAccVolcengineAclUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "acl_name", "acc-test-acl-new"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test-demo-new"),
					resource.TestCheckResourceAttr(acc.ResourceId, "project_name", "default"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "create_time"),
					resource.TestCheckResourceAttr(acc.ResourceId, "acl_entries.#", "2"),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "acl_entries.*", map[string]string{
						"entry":       "172.20.2.0/24",
						"description": "e2",
					}),
					volcengine.TestCheckTypeSetElemNestedAttrs(acc.ResourceId, "acl_entries.*", map[string]string{
						"entry":       "172.20.3.0/24",
						"description": "e3",
					}),
				),
			},
			{
				Config:             testAccVolcengineAclUpdateConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // 修改之后，不应该再产生diff
			},
		},
	})
}
