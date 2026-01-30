package health_check_log_project_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/clb/health_check_log_project"
)

const testAccVolcengineHealthCheckLogProjectsDatasourceConfig = `
data "volcengine_health_check_log_projects" "foo" {
}
`

func TestAccVolcengineHealthCheckLogProjectsDatasource_Basic(t *testing.T) {
	resourceName := "data.volcengine_health_check_log_projects.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &health_check_log_project.VolcengineHealthCheckLogProjectService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers: volcengine.GetTestAccProviders(),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineHealthCheckLogProjectsDatasourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(acc.ResourceId, "health_check_log_projects.#"),
					resource.TestCheckResourceAttrSet(acc.ResourceId, "total_count"),
				),
			},
		},
	})
}
