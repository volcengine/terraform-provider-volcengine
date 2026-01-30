package health_check_log_topic_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/clb/health_check_log_topic"
)

const testAccVolcengineHealthCheckLogTopicsDatasourceConfig = `
data "volcengine_health_check_log_topics" "foo"{
    log_topic_id = "05df22e2-f561-4081-8cf3-b201af564407"
}
`

func TestAccVolcengineHealthCheckLogTopicsDatasource_Basic(t *testing.T) {
	resourceName := "data.volcengine_health_check_log_topics.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &health_check_log_topic.VolcengineHealthCheckLogTopicService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers: volcengine.GetTestAccProviders(),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineHealthCheckLogTopicsDatasourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(acc.ResourceId, "health_check_log_topics.#"),
				),
			},
		},
	})
}
