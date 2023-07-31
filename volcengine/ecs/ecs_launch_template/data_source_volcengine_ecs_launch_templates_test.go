package ecs_launch_template_test

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/volcengine/terraform-provider-volcengine/volcengine"
	"github.com/volcengine/terraform-provider-volcengine/volcengine/ecs/ecs_launch_template"
	"testing"
)

const testAccVolcengineEcsLaunchTemplatesDatasourceConfig = `
resource "volcengine_ecs_launch_template" "foo" {
    description = "acc-test-desc"
    eip_bandwidth = 1
    eip_billing_type = "PostPaidByBandwidth"
    eip_isp = "ChinaMobile"
    host_name = "acc-xx"
    hpc_cluster_id = "acc-xx"
    image_id = "acc-xx"
    instance_charge_type = "acc-xx"
    instance_name = "acc-xx"
    instance_type_id = "acc-xx"
    key_pair_name = "acc-xx"
    launch_template_name = "acc-test-template2"
}

data "volcengine_ecs_launch_templates" "foo"{
    ids = ["${volcengine_ecs_launch_template.foo.id}"]
}
`

func TestAccVolcengineEcsLaunchTemplatesDatasource_Basic(t *testing.T) {
	resourceName := "data.volcengine_ecs_launch_templates.foo"

	acc := &volcengine.AccTestResource{
		ResourceId: resourceName,
		Svc:        &ecs_launch_template.VolcengineEcsLaunchTemplateService{},
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			volcengine.AccTestPreCheck(t)
		},
		Providers: volcengine.GetTestAccProviders(),
		Steps: []resource.TestStep{
			{
				Config: testAccVolcengineEcsLaunchTemplatesDatasourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(acc.ResourceId, "launch_templates.#", "1"),
				),
			},
		},
	})
}
