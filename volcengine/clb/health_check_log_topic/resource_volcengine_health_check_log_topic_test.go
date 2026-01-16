package health_check_log_topic_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/clb/health_check_log_topic"
)

const testAccVolcengineHealthCheckLogTopicCreateConfig = `
resource "volcengine_health_check_log_topic" "foo" {
  log_topic_id     = "05df22e2-f561-4081-8cf3-b201af564407"
  load_balancer_id = "clb-mim12q0soe805smt1bebim25"
}
`

func TestAccVolcengineHealthCheckLogTopicResource_Basic(t *testing.T) {
	resourceName := "volcengine_health_check_log_topic.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &health_check_log_topic.VolcengineHealthCheckLogTopicService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers:    volcengine.GetTestAccProviders(),
		CheckDestroy: volcengine.AccTestCheckResourceRemove(acc),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineHealthCheckLogTopicCreateConfig,
				Check: resource.ComposeTestCheckFunc(
					volcengine.AccTestCheckResourceExists(acc),
					resource.TestCheckResourceAttr(acc.ResourceId, "log_topic_id", "05df22e2-f561-4081-8cf3-b201af564407"),
					resource.TestCheckResourceAttr(acc.ResourceId, "load_balancer_id", "clb-mim12q0soe805smt1bebim25"),
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
