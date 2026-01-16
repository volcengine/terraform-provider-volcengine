package health_check_log_project_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/clb/health_check_log_project"
)

const testAccVolcengineHealthCheckLogProjectCreateConfig = `
resource "volcengine_health_check_log_project" "foo" {
}
`

func TestAccVolcengineHealthCheckLogProjectResource_Basic(t *testing.T) {
	resourceName := "volcengine_health_check_log_project.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &health_check_log_project.VolcengineHealthCheckLogProjectService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineHealthCheckLogProjectCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "log_project_id"),
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
