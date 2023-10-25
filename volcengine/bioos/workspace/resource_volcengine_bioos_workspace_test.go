package workspace_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/bioos/workspace"
)

const testAccVolcengineBioosWorkspaceCreateConfig = `
resource "volcengine_bioos_workspace" "foo" {
    description = "acc-test-workspace"
    name = "acc-test-workspace"
}
`

const testAccVolcengineBioosWorkspaceUpdateConfig = `
resource "volcengine_bioos_workspace" "foo" {
    cover_path = "template-cover/pic5.png"
    description = "acc-test-workspace-modify"
    name = "acc-test-workspace-modify"
}
`

func TestAccVolcengineBioosWorkspaceResource_Basic(t *testing.T) {
	resourceName := "volcengine_bioos_workspace.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return workspace.NewVolcengineBioosWorkspaceService(client)
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
				Config: testAccVolcengineBioosWorkspaceCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test-workspace"),
					resource.TestCheckResourceAttr(acc.ResourceId, "name", "acc-test-workspace"),
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

func TestAccVolcengineBioosWorkspaceResource_Update(t *testing.T) {
	resourceName := "volcengine_bioos_workspace.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return workspace.NewVolcengineBioosWorkspaceService(client)
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
				Config: testAccVolcengineBioosWorkspaceCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test-workspace"),
					resource.TestCheckResourceAttr(acc.ResourceId, "name", "acc-test-workspace"),
				),
			},
			{
				Config: testAccVolcengineBioosWorkspaceUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "cover_path", "template-cover/pic5.png"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test-workspace-modify"),
					resource.TestCheckResourceAttr(acc.ResourceId, "name", "acc-test-workspace-modify"),
				),
			},
			{
				Config:             testAccVolcengineBioosWorkspaceUpdateConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // 修改之后，不应该再产生diff
			},
		},
	})
}
