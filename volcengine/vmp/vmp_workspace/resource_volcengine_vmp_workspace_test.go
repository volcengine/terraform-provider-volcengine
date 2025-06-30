package vmp_workspace_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/vmp/vmp_workspace"
)

const testAccVolcengineVmpWorkspaceCreateConfig = `
resource "volcengine_vmp_workspace" "foo" {
  name                      = "acc-test-1"
  instance_type_id          = "vmp.standard.15d"
  delete_protection_enabled = false
  description               = "acc-test-1"
  username                  = "admin123"
  password                  = "admin1239A82"
}

`

const testAccVolcengineVmpWorkspaceUpdateConfig = `
resource "volcengine_vmp_workspace" "foo" {
  name                      = "acc-test-2"
  instance_type_id          = "vmp.standard.15d"
  delete_protection_enabled = false
  description               = "acc-test-2"
  username                  = "admin1234"
  password                  = "admin1239A82"
}
`

func TestAccVolcengineVmpWorkspaceResource_Basic(t *testing.T) {
	resourceName := "volcengine_vmp_workspace.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return vmp_workspace.NewService(client)
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
				Config: testAccVolcengineVmpWorkspaceCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "delete_protection_enabled", "false"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test-1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_type_id", "vmp.standard.15d"),
					resource.TestCheckResourceAttr(acc.ResourceId, "name", "acc-test-1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "username", "admin123"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"password", "username"},
			},
		},
	})
}

func TestAccVolcengineVmpWorkspaceResource_Update(t *testing.T) {
	resourceName := "volcengine_vmp_workspace.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return vmp_workspace.NewService(client)
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
				Config: testAccVolcengineVmpWorkspaceCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "delete_protection_enabled", "false"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test-1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_type_id", "vmp.standard.15d"),
					resource.TestCheckResourceAttr(acc.ResourceId, "name", "acc-test-1"),
					resource.TestCheckResourceAttr(acc.ResourceId, "username", "admin123"),
				),
			},
			{
				Config: testAccVolcengineVmpWorkspaceUpdateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "delete_protection_enabled", "false"),
					resource.TestCheckResourceAttr(acc.ResourceId, "description", "acc-test-2"),
					resource.TestCheckResourceAttr(acc.ResourceId, "instance_type_id", "vmp.standard.15d"),
					resource.TestCheckResourceAttr(acc.ResourceId, "name", "acc-test-2"),
					resource.TestCheckResourceAttr(acc.ResourceId, "username", "admin1234"),
				),
			},
			{
				Config:             testAccVolcengineVmpWorkspaceUpdateConfig,
				PlanOnly:           true,
				ExpectNonEmptyPlan: false, // 修改之后，不应该再产生diff
			},
		},
	})
}
