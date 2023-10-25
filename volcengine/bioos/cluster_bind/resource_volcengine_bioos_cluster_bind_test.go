package cluster_bind_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	ve "github.com/volcengine/terraform-provider-volcengine/common"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/bioos/cluster_bind"
)

const testAccVolcengineBioosClusterBindCreateConfig = `
resource "volcengine_bioos_workspace" "foo" {
    description = "acc-test-workspace3"
    name = "acc-test-workspace3"
}

resource "volcengine_bioos_cluster" "foo" {
    name = "acc-test-cluster3"
  	description = "acc-test-description"
	shared_config {
    enable = true
  }
}

resource "volcengine_bioos_cluster_bind" "foo" {
    cluster_id = volcengine_bioos_cluster.foo.id
    type = "workflow"
    workspace_id = volcengine_bioos_workspace.foo.id
}

`

func TestAccVolcengineBioosClusterBindResource_Basic(t *testing.T) {
	resourceName := "volcengine_bioos_cluster_bind.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		SvcInitFunc: func(client *ve.SdkClient) ve.ResourceService {
			return cluster_bind.NewVolcengineBioosClusterBindService(client)
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
				Config: testAccVolcengineBioosClusterBindCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "type", "workflow"),
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
